package helpers

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	openapi "github.com/factpulse/sdk-go"
)

const (
	// DefaultAPIURL est l'URL par défaut de l'API FactPulse.
	DefaultAPIURL = "https://factpulse.fr"
	// DefaultPollingInterval est l'intervalle de polling par défaut (en millisecondes).
	DefaultPollingInterval = 2000
	// DefaultPollingTimeout est le timeout de polling par défaut (en millisecondes).
	DefaultPollingTimeout = 120000
	// DefaultMaxRetries est le nombre maximum de tentatives par défaut.
	DefaultMaxRetries = 1
)

// ClientConfig contient la configuration du client FactPulse.
type ClientConfig struct {
	Email           string
	Password        string
	APIURL          string
	ClientUID       string
	PollingInterval int64
	PollingTimeout  int64
	MaxRetries      int
}

// Client est le client simplifié pour l'API FactPulse avec authentification
// JWT et polling intégrés.
type Client struct {
	config         ClientConfig
	httpClient     *http.Client
	accessToken    string
	refreshToken   string
	tokenExpiresAt time.Time
	mu             sync.Mutex
}

// TokenResponse représente la réponse de l'endpoint de token.
type TokenResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

// TaskStatus représente le statut d'une tâche.
type TaskStatus struct {
	IDTache  string                 `json:"id_tache"`
	Statut   string                 `json:"statut"`
	Resultat map[string]interface{} `json:"resultat"`
}

// NewClient crée un nouveau client FactPulse.
func NewClient(config ClientConfig) *Client {
	if config.APIURL == "" {
		config.APIURL = DefaultAPIURL
	}
	config.APIURL = strings.TrimRight(config.APIURL, "/")

	if config.PollingInterval == 0 {
		config.PollingInterval = DefaultPollingInterval
	}
	if config.PollingTimeout == 0 {
		config.PollingTimeout = DefaultPollingTimeout
	}
	if config.MaxRetries == 0 {
		config.MaxRetries = DefaultMaxRetries
	}

	return &Client{
		config: config,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// obtainToken obtient un nouveau token JWT.
func (c *Client) obtainToken() (*TokenResponse, error) {
	tokenURL := c.config.APIURL + "/api/token/"

	payload := map[string]string{
		"username": c.config.Email,
		"password": c.config.Password,
	}
	if c.config.ClientUID != "" {
		payload["client_uid"] = c.config.ClientUID
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", tokenURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, NewFactPulseAuthError(fmt.Sprintf("Erreur réseau: %v", err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		var errData map[string]interface{}
		json.Unmarshal(respBody, &errData)
		detail := ""
		if d, ok := errData["detail"].(string); ok {
			detail = d
		} else {
			detail = string(respBody)
		}
		return nil, NewFactPulseAuthError(fmt.Sprintf("Impossible d'obtenir le token JWT: %s", detail))
	}

	var tokens TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokens); err != nil {
		return nil, err
	}

	log.Printf("Token JWT obtenu pour %s", c.config.Email)
	return &tokens, nil
}

// refreshAccessToken rafraîchit le token d'accès.
func (c *Client) refreshAccessToken() (string, error) {
	if c.refreshToken == "" {
		return "", NewFactPulseAuthError("Aucun refresh token disponible")
	}

	refreshURL := c.config.APIURL + "/api/token/refresh/"

	payload := map[string]string{
		"refresh": c.refreshToken,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", refreshURL, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("Refresh token échoué, ré-obtention d'un nouveau token")
		tokens, err := c.obtainToken()
		if err != nil {
			return "", err
		}
		return tokens.Access, nil
	}
	defer resp.Body.Close()

	var data map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	log.Println("Token rafraîchi avec succès")
	return data["access"], nil
}

// EnsureAuthenticated s'assure que le client est authentifié.
func (c *Client) EnsureAuthenticated(forceRefresh bool) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()

	if forceRefresh || c.accessToken == "" || now.After(c.tokenExpiresAt) {
		if c.refreshToken != "" && !c.tokenExpiresAt.IsZero() && !forceRefresh {
			token, err := c.refreshAccessToken()
			if err == nil {
				c.accessToken = token
				c.tokenExpiresAt = now.Add(28 * time.Minute)
				return nil
			}
		}

		tokens, err := c.obtainToken()
		if err != nil {
			return err
		}
		c.accessToken = tokens.Access
		c.refreshToken = tokens.Refresh
		c.tokenExpiresAt = now.Add(28 * time.Minute)
	}

	return nil
}

// ResetAuth réinitialise l'authentification.
func (c *Client) ResetAuth() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.accessToken = ""
	c.refreshToken = ""
	c.tokenExpiresAt = time.Time{}
	log.Println("Authentification réinitialisée")
}

// GetConfiguration retourne une configuration pour le SDK généré.
func (c *Client) GetConfiguration() *openapi.Configuration {
	cfg := openapi.NewConfiguration()
	cfg.Servers = openapi.ServerConfigurations{
		{URL: c.config.APIURL + "/api/facturation"},
	}
	cfg.AddDefaultHeader("Authorization", "Bearer "+c.accessToken)
	return cfg
}

// GetTraitementAPI retourne l'API de traitement de factures.
func (c *Client) GetTraitementAPI() (*openapi.APIClient, error) {
	if err := c.EnsureAuthenticated(false); err != nil {
		return nil, err
	}
	return openapi.NewAPIClient(c.GetConfiguration()), nil
}

// PollTask poll une tâche jusqu'à son achèvement.
func (c *Client) PollTask(ctx context.Context, taskID string, timeout, interval *int64) (map[string]interface{}, error) {
	timeoutMs := c.config.PollingTimeout
	if timeout != nil {
		timeoutMs = *timeout
	}
	intervalMs := c.config.PollingInterval
	if interval != nil {
		intervalMs = *interval
	}

	startTime := time.Now()
	currentInterval := float64(intervalMs)

	log.Printf("Début du polling pour la tâche %s (timeout: %dms)", taskID, timeoutMs)

	for {
		elapsed := time.Since(startTime).Milliseconds()

		if elapsed > timeoutMs {
			return nil, NewFactPulsePollingTimeout(taskID, timeoutMs)
		}

		apiClient, err := c.GetTraitementAPI()
		if err != nil {
			return nil, err
		}

		resp, httpResp, err := apiClient.TraitementFactureAPI.
			ObtenirStatutTacheApiV1TraitementTachesIdTacheStatutGet(ctx, taskID).
			Execute()

		if err != nil {
			if httpResp != nil && httpResp.StatusCode == 401 {
				log.Println("Token expiré pendant le polling, re-authentification...")
				c.ResetAuth()
				continue
			}
			return nil, NewFactPulseValidationError(fmt.Sprintf("Erreur API: %v", err), nil)
		}

		statut := ""
		if resp.Statut != nil {
			statut = string(*resp.Statut)
		}
		log.Printf("Tâche %s: statut=%s (%dms écoulées)", taskID, statut, elapsed)

		if statut == "SUCCESS" {
			log.Printf("Tâche %s terminée avec succès", taskID)
			// Convert resultat to map
			if resp.Resultat != nil {
				data, _ := json.Marshal(resp.Resultat)
				var result map[string]interface{}
				json.Unmarshal(data, &result)
				return result, nil
			}
			return make(map[string]interface{}), nil
		}

		if statut == "FAILURE" {
			errorMsg := "Erreur inconnue"
			var errors []ValidationErrorDetail
			if resp.Resultat != nil {
				data, _ := json.Marshal(resp.Resultat)
				var result map[string]interface{}
				json.Unmarshal(data, &result)
				if msg, ok := result["message_erreur"].(string); ok {
					errorMsg = msg
				}
				if errs, ok := result["erreurs"].([]interface{}); ok {
					for _, e := range errs {
						if errMap, ok := e.(map[string]interface{}); ok {
							detail := ValidationErrorDetail{}
							if v, ok := errMap["level"].(string); ok {
								detail.Level = v
							}
							if v, ok := errMap["item"].(string); ok {
								detail.Item = v
							}
							if v, ok := errMap["reason"].(string); ok {
								detail.Reason = v
							}
							errors = append(errors, detail)
						}
					}
				}
			}
			return nil, NewFactPulseValidationError(
				fmt.Sprintf("La tâche %s a échoué: %s", taskID, errorMsg),
				errors,
			)
		}

		// Backoff exponentiel (max 10 secondes)
		time.Sleep(time.Duration(currentInterval) * time.Millisecond)
		currentInterval = math.Min(currentInterval*1.5, 10000)
	}
}

// GenererFacturx génère une facture Factur-X.
func (c *Client) GenererFacturx(
	ctx context.Context,
	factureData interface{},
	pdfSource []byte,
	profil string,
	formatSortie string,
	sync bool,
	timeout *int64,
) ([]byte, error) {
	var jsonData string
	switch v := factureData.(type) {
	case string:
		jsonData = v
	case []byte:
		jsonData = string(v)
	default:
		data, err := json.Marshal(v)
		if err != nil {
			return nil, NewFactPulseValidationError(fmt.Sprintf("Erreur de sérialisation: %v", err), nil)
		}
		jsonData = string(data)
	}

	if profil == "" {
		profil = "EN16931"
	}
	if formatSortie == "" {
		formatSortie = "pdf"
	}

	// Créer un fichier temporaire pour le PDF
	tmpFile, err := os.CreateTemp("", "factpulse_*.pdf")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(pdfSource); err != nil {
		return nil, err
	}
	tmpFile.Close()

	var taskID string
	for attempt := 0; attempt <= c.config.MaxRetries; attempt++ {
		apiClient, err := c.GetTraitementAPI()
		if err != nil {
			return nil, err
		}

		pdfFile, err := os.Open(tmpFile.Name())
		if err != nil {
			return nil, err
		}
		defer pdfFile.Close()

		resp, httpResp, err := apiClient.TraitementFactureAPI.
			GenererFactureApiV1TraitementGenererFacturePost(ctx).
			DonneesFacture(jsonData).
			Profil(openapi.ProfilApi(profil)).
			FormatSortie(openapi.FormatSortie(formatSortie)).
			SourcePdf(pdfFile).
			Execute()

		if err != nil {
			if httpResp != nil && httpResp.StatusCode == 401 && attempt < c.config.MaxRetries {
				log.Printf("Erreur 401, réinitialisation du token (tentative %d/%d)",
					attempt+1, c.config.MaxRetries+1)
				c.ResetAuth()
				continue
			}
			return nil, NewFactPulseValidationError(fmt.Sprintf("Erreur API: %v", err), nil)
		}

		taskID = resp.GetIdTache()
		break
	}

	if !sync {
		return []byte(taskID), nil
	}

	result, err := c.PollTask(ctx, taskID, timeout, nil)
	if err != nil {
		return nil, err
	}

	if statut, ok := result["statut"].(string); ok && statut == "ERREUR" {
		errorMsg := "Erreur de validation"
		if msg, ok := result["message_erreur"].(string); ok {
			errorMsg = msg
		}
		var errors []ValidationErrorDetail
		if errs, ok := result["erreurs"].([]interface{}); ok {
			for _, e := range errs {
				if errMap, ok := e.(map[string]interface{}); ok {
					detail := ValidationErrorDetail{}
					if v, ok := errMap["level"].(string); ok {
						detail.Level = v
					}
					if v, ok := errMap["item"].(string); ok {
						detail.Item = v
					}
					if v, ok := errMap["reason"].(string); ok {
						detail.Reason = v
					}
					errors = append(errors, detail)
				}
			}
		}
		return nil, NewFactPulseValidationError(errorMsg, errors)
	}

	if contenuB64, ok := result["contenu_b64"].(string); ok {
		return base64.StdEncoding.DecodeString(contenuB64)
	}

	return nil, NewFactPulseValidationError("Le résultat ne contient pas de contenu", nil)
}

// FormatMontant formate un montant pour l'API FactPulse.
func FormatMontant(montant interface{}) string {
	if montant == nil {
		return "0.00"
	}

	switch v := montant.(type) {
	case float64:
		return fmt.Sprintf("%.2f", v)
	case float32:
		return fmt.Sprintf("%.2f", v)
	case int:
		return fmt.Sprintf("%.2f", float64(v))
	case int64:
		return fmt.Sprintf("%.2f", float64(v))
	case string:
		return v
	default:
		return fmt.Sprintf("%v", montant)
	}
}

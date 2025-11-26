package helpers

import (
    "bytes"; "encoding/json"; "fmt"; "io"; "net/http"; "strings"; "sync"; "time"
)

const DefaultAPIURL = "https://factpulse.fr"
const DefaultPollingInterval = 2000
const DefaultPollingTimeout = 120000

// ChorusProCredentials contient les credentials Chorus Pro pour le mode Zero-Trust.
// Ces credentials sont passés dans chaque requête et ne sont jamais stockés côté serveur.
type ChorusProCredentials struct {
    PisteClientID     string `json:"piste_client_id"`
    PisteClientSecret string `json:"piste_client_secret"`
    ChorusProLogin    string `json:"chorus_pro_login"`
    ChorusProPassword string `json:"chorus_pro_password"`
    Sandbox           bool   `json:"sandbox"`
}

func NewChorusProCredentials(pisteClientID, pisteClientSecret, chorusProLogin, chorusProPassword string, sandbox bool) *ChorusProCredentials {
    return &ChorusProCredentials{PisteClientID: pisteClientID, PisteClientSecret: pisteClientSecret, ChorusProLogin: chorusProLogin, ChorusProPassword: chorusProPassword, Sandbox: sandbox}
}

func (c *ChorusProCredentials) ToMap() map[string]interface{} {
    return map[string]interface{}{"piste_client_id": c.PisteClientID, "piste_client_secret": c.PisteClientSecret, "chorus_pro_login": c.ChorusProLogin, "chorus_pro_password": c.ChorusProPassword, "sandbox": c.Sandbox}
}

// AFNORCredentials contient les credentials AFNOR PDP pour le mode Zero-Trust.
// Ces credentials sont passés dans chaque requête et ne sont jamais stockés côté serveur.
type AFNORCredentials struct {
    ClientID       string `json:"client_id"`
    ClientSecret   string `json:"client_secret"`
    FlowServiceURL string `json:"flow_service_url"`
}

func NewAFNORCredentials(clientID, clientSecret, flowServiceURL string) *AFNORCredentials {
    return &AFNORCredentials{ClientID: clientID, ClientSecret: clientSecret, FlowServiceURL: flowServiceURL}
}

func (c *AFNORCredentials) ToMap() map[string]interface{} {
    return map[string]interface{}{"client_id": c.ClientID, "client_secret": c.ClientSecret, "flow_service_url": c.FlowServiceURL}
}

// Montant formate une valeur en montant avec 2 décimales
func Montant(m interface{}) string {
    if m == nil { return "0.00" }
    switch v := m.(type) {
    case float64: return fmt.Sprintf("%.2f", v)
    case float32: return fmt.Sprintf("%.2f", v)
    case int: return fmt.Sprintf("%.2f", float64(v))
    case int64: return fmt.Sprintf("%.2f", float64(v))
    case string: return v
    default: return "0.00"
    }
}

// MontantTotal crée un objet montant total simplifié
func MontantTotal(ht, tva, ttc, aPayer interface{}) map[string]interface{} {
    return MontantTotalAvecOptions(ht, tva, ttc, aPayer, nil, "", nil)
}

// MontantTotalAvecOptions crée un objet montant total avec options
func MontantTotalAvecOptions(ht, tva, ttc, aPayer, remiseTtc interface{}, motifRemise string, acompte interface{}) map[string]interface{} {
    result := map[string]interface{}{"montantHtTotal": Montant(ht), "montantTva": Montant(tva), "montantTtcTotal": Montant(ttc), "montantAPayer": Montant(aPayer)}
    if remiseTtc != nil { result["montantRemiseGlobaleTtc"] = Montant(remiseTtc) }
    if motifRemise != "" { result["motifRemiseGlobaleTtc"] = motifRemise }
    if acompte != nil { result["acompte"] = Montant(acompte) }
    return result
}

// LigneDePoste crée une ligne de poste simplifiée
func LigneDePoste(numero int, denomination string, quantite, montantUnitaireHt, montantLigneHt interface{}) map[string]interface{} {
    return LigneDePosteAvecOptions(numero, denomination, quantite, montantUnitaireHt, montantLigneHt, "20.00", "C62", nil)
}

// LigneDePosteAvecOptions crée une ligne de poste avec options
func LigneDePosteAvecOptions(numero int, denomination string, quantite, montantUnitaireHt, montantLigneHt interface{}, tauxTva, unite string, options map[string]interface{}) map[string]interface{} {
    result := map[string]interface{}{"numero": numero, "denomination": denomination, "quantite": Montant(quantite), "montantUnitaireHt": Montant(montantUnitaireHt), "montantTotalLigneHt": Montant(montantLigneHt), "tauxTva": Montant(tauxTva), "unite": unite}
    if options != nil {
        if v, ok := options["montantTvaLigne"]; ok { result["montantTvaLigne"] = Montant(v) }
        if v, ok := options["montantRemiseHt"]; ok { result["montantRemiseHt"] = Montant(v) }
        if v, ok := options["codeRaisonRemise"]; ok { result["codeRaisonReduction"] = v }
        if v, ok := options["motifRemise"]; ok { result["motifRemise"] = v }
        if v, ok := options["description"]; ok { result["description"] = v }
    }
    return result
}

// LigneDeTva crée une ligne de TVA simplifiée
func LigneDeTva(taux, baseHt, montantTva interface{}) map[string]interface{} {
    return LigneDeTvaAvecOptions(taux, baseHt, montantTva, "S", "")
}

// LigneDeTvaAvecOptions crée une ligne de TVA avec options
func LigneDeTvaAvecOptions(taux, baseHt, montantTva interface{}, categorie, motifExoneration string) map[string]interface{} {
    result := map[string]interface{}{"tauxTva": Montant(taux), "montantBaseHt": Montant(baseHt), "montantTva": Montant(montantTva), "categorieTva": categorie}
    if motifExoneration != "" { result["motifExoneration"] = motifExoneration }
    return result
}

type Client struct {
    Email, Password, APIURL, ClientUID string
    PollingInterval, PollingTimeout int64
    ChorusCredentials *ChorusProCredentials
    AFNORCredentials  *AFNORCredentials
    accessToken, refreshToken string
    tokenExpiresAt time.Time
    mu sync.Mutex
}

func NewClient(email, password string) *Client {
    return &Client{Email: email, Password: password, APIURL: DefaultAPIURL, PollingInterval: DefaultPollingInterval, PollingTimeout: DefaultPollingTimeout}
}

func NewClientWithCredentials(email, password, apiURL, clientUID string, chorusCredentials *ChorusProCredentials, afnorCredentials *AFNORCredentials) *Client {
    url := apiURL; if url == "" { url = DefaultAPIURL }
    return &Client{Email: email, Password: password, APIURL: strings.TrimRight(url, "/"), ClientUID: clientUID, ChorusCredentials: chorusCredentials, AFNORCredentials: afnorCredentials, PollingInterval: DefaultPollingInterval, PollingTimeout: DefaultPollingTimeout}
}

func (c *Client) GetChorusCredentialsForAPI() map[string]interface{} {
    if c.ChorusCredentials == nil { return nil }
    return c.ChorusCredentials.ToMap()
}

func (c *Client) GetAFNORCredentialsForAPI() map[string]interface{} {
    if c.AFNORCredentials == nil { return nil }
    return c.AFNORCredentials.ToMap()
}

func (c *Client) EnsureAuthenticated(forceRefresh bool) error {
    c.mu.Lock(); defer c.mu.Unlock()
    now := time.Now()
    if forceRefresh || c.accessToken == "" || now.After(c.tokenExpiresAt) {
        tokens, err := c.obtainToken(); if err != nil { return err }
        c.accessToken, c.refreshToken = tokens["access"], tokens["refresh"]
        c.tokenExpiresAt = now.Add(28 * time.Minute)
    }
    return nil
}

func (c *Client) obtainToken() (map[string]string, error) {
    payload := map[string]string{"username": c.Email, "password": c.Password}
    if c.ClientUID != "" { payload["client_uid"] = c.ClientUID }
    body, _ := json.Marshal(payload)
    resp, err := http.Post(strings.TrimRight(c.APIURL, "/")+"/api/token/", "application/json", bytes.NewReader(body))
    if err != nil { return nil, NewFactPulseAuthError(fmt.Sprintf("Network error: %v", err)) }
    defer resp.Body.Close()
    if resp.StatusCode != 200 { return nil, NewFactPulseAuthError("Auth failed") }
    var tokens map[string]string; json.NewDecoder(resp.Body).Decode(&tokens); return tokens, nil
}

func (c *Client) ResetAuth() { c.mu.Lock(); defer c.mu.Unlock(); c.accessToken, c.refreshToken = "", ""; c.tokenExpiresAt = time.Time{} }

func (c *Client) PollTask(taskID string, timeout, interval *int64) (map[string]interface{}, error) {
    timeoutMs := c.PollingTimeout; if timeout != nil { timeoutMs = *timeout }
    intervalMs := c.PollingInterval; if interval != nil { intervalMs = *interval }
    startTime := time.Now(); currentInterval := float64(intervalMs)
    for {
        if time.Since(startTime).Milliseconds() > timeoutMs { return nil, NewFactPulsePollingTimeout(taskID, timeoutMs) }
        if err := c.EnsureAuthenticated(false); err != nil { return nil, err }
        req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/facturation/v1/traitement/taches/%s/statut", strings.TrimRight(c.APIURL, "/"), taskID), nil)
        req.Header.Set("Authorization", "Bearer "+c.accessToken)
        resp, err := http.DefaultClient.Do(req)
        if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("API error: %v", err), nil) }
        defer resp.Body.Close()
        if resp.StatusCode == 401 { c.ResetAuth(); continue }
        body, _ := io.ReadAll(resp.Body); var data map[string]interface{}; json.Unmarshal(body, &data)
        statut, _ := data["statut"].(string)
        if statut == "SUCCESS" { if r, ok := data["resultat"].(map[string]interface{}); ok { return r, nil }; return map[string]interface{}{}, nil }
        if statut == "FAILURE" { return nil, NewFactPulseValidationError("Task failed", nil) }
        time.Sleep(time.Duration(currentInterval) * time.Millisecond); currentInterval = minFloat(currentInterval*1.5, 10000)
    }
}

func FormatMontant(m interface{}) string { return Montant(m) }
func minFloat(a, b float64) float64 { if a < b { return a }; return b }

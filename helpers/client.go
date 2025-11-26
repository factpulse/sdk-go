package helpers

import (
    "bytes"; "encoding/base64"; "encoding/json"; "fmt"; "io"; "net/http"; "os"; "strings"; "sync"; "time"
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
    return LigneDePosteAvecOptions(numero, denomination, quantite, montantUnitaireHt, montantLigneHt, "20.00", "S", "C62", nil)
}

// LigneDePosteAvecOptions crée une ligne de poste avec options
func LigneDePosteAvecOptions(numero int, denomination string, quantite, montantUnitaireHt, montantLigneHt interface{}, tauxTva, categorieTva, unite string, options map[string]interface{}) map[string]interface{} {
    result := map[string]interface{}{"numero": numero, "denomination": denomination, "quantite": Montant(quantite), "montantUnitaireHt": Montant(montantUnitaireHt), "montantTotalLigneHt": Montant(montantLigneHt), "tauxTvaManuel": Montant(tauxTva), "categorieTva": categorieTva, "unite": unite}
    if options != nil {
        if v, ok := options["reference"]; ok { result["reference"] = v }
        if v, ok := options["montantTvaLigne"]; ok { result["montantTvaLigne"] = Montant(v) }
        if v, ok := options["montantRemiseHt"]; ok { result["montantRemiseHt"] = Montant(v) }
        if v, ok := options["codeRaisonReduction"]; ok { result["codeRaisonReduction"] = v }
        if v, ok := options["raisonReduction"]; ok { result["raisonReduction"] = v }
        if v, ok := options["motifExoneration"]; ok { result["motifExoneration"] = v }
        if v, ok := options["dateDebutPeriode"]; ok { result["dateDebutPeriode"] = v }
        if v, ok := options["dateFinPeriode"]; ok { result["dateFinPeriode"] = v }
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
    result := map[string]interface{}{"tauxManuel": Montant(taux), "montantBaseHt": Montant(baseHt), "montantTva": Montant(montantTva), "categorie": categorie}
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

// Alias plus courts
func (c *Client) GetChorusProCredentials() map[string]interface{} { return c.GetChorusCredentialsForAPI() }
func (c *Client) GetAfnorCredentials() map[string]interface{} { return c.GetAFNORCredentialsForAPI() }

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
        req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/traitement/taches/%s/statut", strings.TrimRight(c.APIURL, "/"), taskID), nil)
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

// GenererFacturx génère une facture Factur-X à partir d'un map/struct et d'un PDF source.
// Accepte un map[string]interface{}, une string JSON, ou tout objet sérialisable en JSON.
func (c *Client) GenererFacturx(factureData interface{}, pdfPath string) ([]byte, error) {
    return c.GenererFacturxWithOptions(factureData, pdfPath, "EN16931", "pdf", true, nil)
}

// GenererFacturxWithOptions génère une facture Factur-X avec options avancées.
func (c *Client) GenererFacturxWithOptions(factureData interface{}, pdfPath, profil, formatSortie string, sync bool, timeout *int64) ([]byte, error) {
    // Conversion des données en JSON string
    var jsonData string
    switch v := factureData.(type) {
    case string:
        jsonData = v
    case []byte:
        jsonData = string(v)
    default:
        data, err := json.Marshal(factureData)
        if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Erreur sérialisation JSON: %v", err), nil) }
        jsonData = string(data)
    }

    // Lecture du fichier PDF
    pdfContent, err := os.ReadFile(pdfPath)
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Erreur lecture PDF: %v", err), nil) }

    if err := c.EnsureAuthenticated(false); err != nil { return nil, err }

    // Construire la requête multipart
    var body bytes.Buffer
    writer := NewMultipartWriter(&body)

    // Champ donnees_facture
    _ = writer.WriteField("donnees_facture", jsonData)
    _ = writer.WriteField("profil", profil)
    _ = writer.WriteField("format_sortie", formatSortie)

    // Champ source_pdf (fichier)
    part, _ := writer.CreateFormFile("source_pdf", "facture.pdf")
    part.Write(pdfContent)
    writer.Close()

    req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/traitement/generer-facture", strings.TrimRight(c.APIURL, "/")), &body)
    req.Header.Set("Authorization", "Bearer "+c.accessToken)
    req.Header.Set("Content-Type", writer.FormDataContentType())

    client := &http.Client{Timeout: 120 * time.Second}
    resp, err := client.Do(req)
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Erreur réseau: %v", err), nil) }
    defer resp.Body.Close()

    if resp.StatusCode == 401 {
        c.ResetAuth()
        if err := c.EnsureAuthenticated(false); err != nil { return nil, err }
        req.Header.Set("Authorization", "Bearer "+c.accessToken)
        resp, err = client.Do(req)
        if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Erreur réseau: %v", err), nil) }
        defer resp.Body.Close()
    }

    respBody, _ := io.ReadAll(resp.Body)
    if resp.StatusCode >= 400 {
        return nil, NewFactPulseValidationError(fmt.Sprintf("Erreur API (%d): %s", resp.StatusCode, string(respBody)), nil)
    }

    var data map[string]interface{}
    json.Unmarshal(respBody, &data)

    if sync {
        if taskID, ok := data["id_tache"].(string); ok {
            result, err := c.PollTask(taskID, timeout, nil)
            if err != nil { return nil, err }
            if contenuB64, ok := result["contenu_b64"].(string); ok {
                decoded, _ := base64.StdEncoding.DecodeString(contenuB64)
                return decoded, nil
            }
            if xml, ok := result["contenu_xml"].(string); ok {
                return []byte(xml), nil
            }
            return nil, NewFactPulseValidationError("Résultat inattendu", nil)
        }
    }

    return respBody, nil
}

// MultipartWriter wrapper pour créer des requêtes multipart
type MultipartWriter struct { *multipartWriter }
type multipartWriter struct { w *bytes.Buffer; boundary string }

func NewMultipartWriter(w *bytes.Buffer) *MultipartWriter {
    boundary := fmt.Sprintf("----GoFormBoundary%x", time.Now().UnixNano())
    return &MultipartWriter{&multipartWriter{w: w, boundary: boundary}}
}
func (m *MultipartWriter) WriteField(name, value string) error {
    fmt.Fprintf(m.w, "--%s\r\nContent-Disposition: form-data; name=\"%s\"\r\n\r\n%s\r\n", m.boundary, name, value)
    return nil
}
func (m *MultipartWriter) CreateFormFile(fieldname, filename string) (io.Writer, error) {
    fmt.Fprintf(m.w, "--%s\r\nContent-Disposition: form-data; name=\"%s\"; filename=\"%s\"\r\nContent-Type: application/pdf\r\n\r\n", m.boundary, fieldname, filename)
    return m.w, nil
}
func (m *MultipartWriter) Close() error { fmt.Fprintf(m.w, "--%s--\r\n", m.boundary); return nil }
func (m *MultipartWriter) FormDataContentType() string { return "multipart/form-data; boundary=" + m.boundary }

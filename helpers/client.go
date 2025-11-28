package helpers

import (
    "bytes"; "crypto/sha256"; "encoding/base64"; "encoding/hex"; "encoding/json"; "fmt"; "io"; "mime/multipart"
    "net/http"; "net/textproto"; "net/url"; "os"; "path/filepath"; "strconv"; "strings"; "sync"; "time"
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
// L'API FactPulse utilise ces credentials pour s'authentifier auprès de la PDP AFNOR.
type AFNORCredentials struct {
    FlowServiceURL      string `json:"flow_service_url"`
    TokenURL            string `json:"token_url"`
    ClientID            string `json:"client_id"`
    ClientSecret        string `json:"client_secret"`
    DirectoryServiceURL string `json:"directory_service_url,omitempty"`
}

func NewAFNORCredentials(flowServiceURL, tokenURL, clientID, clientSecret string) *AFNORCredentials {
    return &AFNORCredentials{FlowServiceURL: flowServiceURL, TokenURL: tokenURL, ClientID: clientID, ClientSecret: clientSecret}
}

func NewAFNORCredentialsWithDirectory(flowServiceURL, tokenURL, clientID, clientSecret, directoryServiceURL string) *AFNORCredentials {
    return &AFNORCredentials{FlowServiceURL: flowServiceURL, TokenURL: tokenURL, ClientID: clientID, ClientSecret: clientSecret, DirectoryServiceURL: directoryServiceURL}
}

func (c *AFNORCredentials) ToMap() map[string]interface{} {
    m := map[string]interface{}{"flow_service_url": c.FlowServiceURL, "token_url": c.TokenURL, "client_id": c.ClientID, "client_secret": c.ClientSecret}
    if c.DirectoryServiceURL != "" { m["directory_service_url"] = c.DirectoryServiceURL }
    return m
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

// LigneDePoste crée une ligne de poste (aligné sur LigneDePoste de models.py)
func LigneDePoste(numero int, denomination string, quantite, montantUnitaireHt, montantTotalLigneHt interface{}) map[string]interface{} {
    return LigneDePosteAvecOptions(numero, denomination, quantite, montantUnitaireHt, montantTotalLigneHt, "20.00", "S", "FORFAIT", nil)
}

// LigneDePosteAvecOptions crée une ligne de poste avec options
func LigneDePosteAvecOptions(numero int, denomination string, quantite, montantUnitaireHt, montantTotalLigneHt interface{}, tauxTva, categorieTva, unite string, options map[string]interface{}) map[string]interface{} {
    result := map[string]interface{}{"numero": numero, "denomination": denomination, "quantite": Montant(quantite), "montantUnitaireHt": Montant(montantUnitaireHt), "montantTotalLigneHt": Montant(montantTotalLigneHt), "tauxTvaManuel": Montant(tauxTva), "categorieTva": categorieTva, "unite": unite}
    if options != nil {
        if v, ok := options["reference"]; ok { result["reference"] = v }
        if v, ok := options["montantRemiseHt"]; ok { result["montantRemiseHt"] = Montant(v) }
        if v, ok := options["codeRaisonReduction"]; ok { result["codeRaisonReduction"] = v }
        if v, ok := options["raisonReduction"]; ok { result["raisonReduction"] = v }
        if v, ok := options["dateDebutPeriode"]; ok { result["dateDebutPeriode"] = v }
        if v, ok := options["dateFinPeriode"]; ok { result["dateFinPeriode"] = v }
    }
    return result
}

// LigneDeTva crée une ligne de TVA (aligné sur LigneDeTVA de models.py)
func LigneDeTva(tauxManuel, montantBaseHt, montantTva interface{}) map[string]interface{} {
    return LigneDeTvaAvecOptions(tauxManuel, montantBaseHt, montantTva, "S")
}

// LigneDeTvaAvecOptions crée une ligne de TVA avec options
func LigneDeTvaAvecOptions(tauxManuel, montantBaseHt, montantTva interface{}, categorie string) map[string]interface{} {
    return map[string]interface{}{"tauxManuel": Montant(tauxManuel), "montantBaseHt": Montant(montantBaseHt), "montantTva": Montant(montantTva), "categorie": categorie}
}

// AdressePostale crée une adresse postale pour l'API FactPulse
func AdressePostale(ligne1, codePostal, ville string) map[string]interface{} {
    return AdressePostaleAvecOptions(ligne1, codePostal, ville, "FR", "", "")
}

// AdressePostaleAvecOptions crée une adresse postale avec options
func AdressePostaleAvecOptions(ligne1, codePostal, ville, pays, ligne2, ligne3 string) map[string]interface{} {
    result := map[string]interface{}{"ligneUn": ligne1, "codePostal": codePostal, "nomVille": ville, "paysCodeIso": pays}
    if ligne2 != "" { result["ligneDeux"] = ligne2 }
    if ligne3 != "" { result["ligneTrois"] = ligne3 }
    return result
}

// AdresseElectronique crée une adresse électronique. schemeID: "0009"=SIREN, "0225"=SIRET
func AdresseElectronique(identifiant, schemeID string) map[string]interface{} {
    if schemeID == "" { schemeID = "0009" }
    return map[string]interface{}{"identifiant": identifiant, "schemeId": schemeID}
}

// calculerTvaIntra calcule le numéro TVA intracommunautaire français depuis un SIREN
func calculerTvaIntra(siren string) string {
    if len(siren) != 9 { return "" }
    sirenNum, err := strconv.ParseInt(siren, 10, 64)
    if err != nil { return "" }
    cle := (12 + 3*(sirenNum%97)) % 97
    return fmt.Sprintf("FR%02d%s", cle, siren)
}

// FournisseurOptions contient les options pour créer un fournisseur
type FournisseurOptions struct {
    IdFournisseur int
    Siren, NumeroTvaIntra, Iban, Pays, AdresseLigne2 string
    CodeService, CodeCoordonnesBancaires int
}

// Fournisseur crée un fournisseur avec auto-calcul SIREN, TVA intracommunautaire et adresses
func Fournisseur(nom, siret, adresseLigne1, codePostal, ville string, opts *FournisseurOptions) map[string]interface{} {
    if opts == nil { opts = &FournisseurOptions{} }
    siren := opts.Siren
    if siren == "" && len(siret) == 14 { siren = siret[:9] }
    numeroTvaIntra := opts.NumeroTvaIntra
    if numeroTvaIntra == "" && siren != "" { numeroTvaIntra = calculerTvaIntra(siren) }
    pays := opts.Pays; if pays == "" { pays = "FR" }
    result := map[string]interface{}{
        "nom": nom, "idFournisseur": opts.IdFournisseur, "siret": siret,
        "adresseElectronique": AdresseElectronique(siret, "0225"),
        "adressePostale": AdressePostaleAvecOptions(adresseLigne1, codePostal, ville, pays, opts.AdresseLigne2, ""),
    }
    if siren != "" { result["siren"] = siren }
    if numeroTvaIntra != "" { result["numeroTvaIntra"] = numeroTvaIntra }
    if opts.Iban != "" { result["iban"] = opts.Iban }
    if opts.CodeService != 0 { result["idServiceFournisseur"] = opts.CodeService }
    if opts.CodeCoordonnesBancaires != 0 { result["codeCoordonnesBancairesFournisseur"] = opts.CodeCoordonnesBancaires }
    return result
}

// DestinataireOptions contient les options pour créer un destinataire
type DestinataireOptions struct {
    Siren, Pays, AdresseLigne2, CodeServiceExecutant string
}

// Destinataire crée un destinataire avec auto-calcul SIREN et adresses
func Destinataire(nom, siret, adresseLigne1, codePostal, ville string, opts *DestinataireOptions) map[string]interface{} {
    if opts == nil { opts = &DestinataireOptions{} }
    siren := opts.Siren
    if siren == "" && len(siret) == 14 { siren = siret[:9] }
    pays := opts.Pays; if pays == "" { pays = "FR" }
    result := map[string]interface{}{
        "nom": nom, "siret": siret,
        "adresseElectronique": AdresseElectronique(siret, "0225"),
        "adressePostale": AdressePostaleAvecOptions(adresseLigne1, codePostal, ville, pays, opts.AdresseLigne2, ""),
    }
    if siren != "" { result["siren"] = siren }
    if opts.CodeServiceExecutant != "" { result["codeServiceExecutant"] = opts.CodeServiceExecutant }
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
    writer := multipart.NewWriter(&body)
    writer.WriteField("donnees_facture", jsonData)
    writer.WriteField("profil", profil)
    writer.WriteField("format_sortie", formatSortie)
    part, _ := writer.CreateFormFile("source_pdf", filepath.Base(pdfPath))
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

// =========================================================================
// AFNOR PDP - Authentication et helpers internes
// =========================================================================

func (c *Client) getAfnorCredentialsInternal() (*AFNORCredentials, error) {
    if c.AFNORCredentials != nil { return c.AFNORCredentials, nil }

    if err := c.EnsureAuthenticated(false); err != nil { return nil, err }
    req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/afnor/credentials", strings.TrimRight(c.APIURL, "/")), nil)
    req.Header.Set("Authorization", "Bearer "+c.accessToken)
    resp, err := http.DefaultClient.Do(req)
    if err != nil { return nil, NewFactPulseAuthError(fmt.Sprintf("Network error: %v", err)) }
    defer resp.Body.Close()
    if resp.StatusCode != 200 { return nil, NewFactPulseAuthError("Failed to get AFNOR credentials") }
    var creds map[string]string; json.NewDecoder(resp.Body).Decode(&creds)
    return &AFNORCredentials{FlowServiceURL: creds["flow_service_url"], TokenURL: creds["token_url"], ClientID: creds["client_id"], ClientSecret: creds["client_secret"], DirectoryServiceURL: creds["directory_service_url"]}, nil
}

func (c *Client) getAfnorTokenAndUrl() (string, string, error) {
    credentials, err := c.getAfnorCredentialsInternal()
    if err != nil { return "", "", err }

    formData := url.Values{}
    formData.Set("grant_type", "client_credentials")
    formData.Set("client_id", credentials.ClientID)
    formData.Set("client_secret", credentials.ClientSecret)

    req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/afnor/oauth/token", strings.TrimRight(c.APIURL, "/")), strings.NewReader(formData.Encode()))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("X-PDP-Token-URL", credentials.TokenURL)

    resp, err := http.DefaultClient.Do(req)
    if err != nil { return "", "", NewFactPulseAuthError(fmt.Sprintf("Network error: %v", err)) }
    defer resp.Body.Close()
    if resp.StatusCode != 200 { return "", "", NewFactPulseAuthError("AFNOR OAuth2 failed") }

    var tokenData map[string]interface{}; json.NewDecoder(resp.Body).Decode(&tokenData)
    token, ok := tokenData["access_token"].(string)
    if !ok { return "", "", NewFactPulseAuthError("Invalid AFNOR OAuth2 response") }

    return token, credentials.FlowServiceURL, nil
}

func (c *Client) makeAfnorRequest(method, endpoint string, jsonData map[string]interface{}, multipartBody *bytes.Buffer, multipartContentType string) (map[string]interface{}, error) {
    token, pdpBaseURL, err := c.getAfnorTokenAndUrl()
    if err != nil { return nil, err }

    url := fmt.Sprintf("%s/api/v1/afnor%s", strings.TrimRight(c.APIURL, "/"), endpoint)

    var req *http.Request
    if multipartBody != nil {
        req, _ = http.NewRequest(method, url, multipartBody)
        req.Header.Set("Content-Type", multipartContentType)
    } else if jsonData != nil {
        body, _ := json.Marshal(jsonData)
        req, _ = http.NewRequest(method, url, bytes.NewReader(body))
        req.Header.Set("Content-Type", "application/json")
    } else {
        req, _ = http.NewRequest(method, url, nil)
    }

    req.Header.Set("Authorization", "Bearer "+token)
    req.Header.Set("X-PDP-Base-URL", pdpBaseURL)

    client := &http.Client{Timeout: 60 * time.Second}
    resp, err := client.Do(req)
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Network error: %v", err), nil) }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    if resp.StatusCode >= 400 {
        return nil, NewFactPulseValidationError(fmt.Sprintf("AFNOR error (%d): %s", resp.StatusCode, string(body)), nil)
    }

    contentType := resp.Header.Get("Content-Type")
    if strings.Contains(contentType, "application/json") {
        var result map[string]interface{}; json.Unmarshal(body, &result); return result, nil
    }
    return map[string]interface{}{"_raw": body}, nil
}

// ==================== AFNOR Flow Service ====================

// SoumettreFactureAfnorOptions contient les options pour soumettreFactureAfnor
type SoumettreFactureAfnorOptions struct {
    FlowSyntax, FlowProfile, TrackingID string
}

// SoumettreFactureAfnor soumet une facture à une PDP via l'API AFNOR
func (c *Client) SoumettreFactureAfnor(pdfPath, flowName string, opts *SoumettreFactureAfnorOptions) (map[string]interface{}, error) {
    if opts == nil { opts = &SoumettreFactureAfnorOptions{} }
    pdfContent, err := os.ReadFile(pdfPath)
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Error reading PDF: %v", err), nil) }

    hash := sha256.Sum256(pdfContent)
    sha256Hex := hex.EncodeToString(hash[:])

    flowInfo := map[string]interface{}{
        "name": flowName,
        "flowSyntax": ifEmpty(opts.FlowSyntax, "CII"),
        "flowProfile": ifEmpty(opts.FlowProfile, "EN16931"),
        "sha256": sha256Hex,
    }
    if opts.TrackingID != "" { flowInfo["trackingId"] = opts.TrackingID }

    var body bytes.Buffer
    writer := multipart.NewWriter(&body)
    part, _ := writer.CreateFormFile("file", filepath.Base(pdfPath))
    part.Write(pdfContent)
    flowInfoJSON, _ := json.Marshal(flowInfo)
    h := make(textproto.MIMEHeader)
    h.Set("Content-Disposition", `form-data; name="flowInfo"`)
    h.Set("Content-Type", "application/json")
    part2, _ := writer.CreatePart(h)
    part2.Write(flowInfoJSON)
    writer.Close()

    return c.makeAfnorRequest("POST", "/flow/v1/flows", nil, &body, writer.FormDataContentType())
}

// RechercherFluxAfnor recherche des flux de facturation AFNOR
func (c *Client) RechercherFluxAfnor(criteria map[string]interface{}) (map[string]interface{}, error) {
    if criteria == nil { criteria = map[string]interface{}{} }
    searchBody := map[string]interface{}{
        "offset": getIntOrDefault(criteria, "offset", 0),
        "limit": getIntOrDefault(criteria, "limit", 25),
        "where": map[string]interface{}{},
    }
    where := searchBody["where"].(map[string]interface{})
    if v, ok := criteria["trackingId"]; ok { where["trackingId"] = v }
    if v, ok := criteria["status"]; ok { where["status"] = v }

    return c.makeAfnorRequest("POST", "/flow/v1/flows/search", searchBody, nil, "")
}

// TelechargerFluxAfnor télécharge le fichier PDF d'un flux AFNOR
func (c *Client) TelechargerFluxAfnor(flowID string) ([]byte, error) {
    result, err := c.makeAfnorRequest("GET", fmt.Sprintf("/flow/v1/flows/%s", flowID), nil, nil, "")
    if err != nil { return nil, err }
    if raw, ok := result["_raw"].([]byte); ok { return raw, nil }
    return nil, nil
}

// HealthcheckAfnor vérifie la disponibilité du Flow Service AFNOR
func (c *Client) HealthcheckAfnor() (map[string]interface{}, error) {
    return c.makeAfnorRequest("GET", "/flow/v1/healthcheck", nil, nil, "")
}

// ==================== AFNOR Directory ====================

// RechercherSiretAfnor recherche une entreprise par SIRET dans l'annuaire AFNOR
func (c *Client) RechercherSiretAfnor(siret string) (map[string]interface{}, error) {
    return c.makeAfnorRequest("GET", fmt.Sprintf("/directory/siret/%s", siret), nil, nil, "")
}

// RechercherSirenAfnor recherche une entreprise par SIREN dans l'annuaire AFNOR
func (c *Client) RechercherSirenAfnor(siren string) (map[string]interface{}, error) {
    return c.makeAfnorRequest("GET", fmt.Sprintf("/directory/siren/%s", siren), nil, nil, "")
}

// ListerCodesRoutageAfnor liste les codes de routage disponibles pour un SIREN
func (c *Client) ListerCodesRoutageAfnor(siren string) (map[string]interface{}, error) {
    return c.makeAfnorRequest("GET", fmt.Sprintf("/directory/siren/%s/routing-codes", siren), nil, nil, "")
}

// =========================================================================
// Chorus Pro
// =========================================================================

func (c *Client) makeChorusRequest(method, endpoint string, jsonData map[string]interface{}) (map[string]interface{}, error) {
    if err := c.EnsureAuthenticated(false); err != nil { return nil, err }

    url := fmt.Sprintf("%s/api/v1/chorus-pro%s", strings.TrimRight(c.APIURL, "/"), endpoint)
    body := jsonData
    if body == nil { body = map[string]interface{}{} }
    if c.ChorusCredentials != nil { body["credentials"] = c.ChorusCredentials.ToMap() }

    var req *http.Request
    if method == "GET" {
        req, _ = http.NewRequest(method, url, nil)
    } else {
        jsonBody, _ := json.Marshal(body)
        req, _ = http.NewRequest(method, url, bytes.NewReader(jsonBody))
        req.Header.Set("Content-Type", "application/json")
    }
    req.Header.Set("Authorization", "Bearer "+c.accessToken)

    client := &http.Client{Timeout: 60 * time.Second}
    resp, err := client.Do(req)
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Network error: %v", err), nil) }
    defer resp.Body.Close()

    respBody, _ := io.ReadAll(resp.Body)
    if resp.StatusCode >= 400 {
        return nil, NewFactPulseValidationError(fmt.Sprintf("Chorus Pro error (%d): %s", resp.StatusCode, string(respBody)), nil)
    }

    var result map[string]interface{}; json.Unmarshal(respBody, &result); return result, nil
}

// RechercherStructureChorus recherche des structures sur Chorus Pro
func (c *Client) RechercherStructureChorus(identifiantStructure, raisonSociale, typeIdentifiant string, restreindrePrivees bool) (map[string]interface{}, error) {
    body := map[string]interface{}{"restreindre_structures_privees": restreindrePrivees}
    if identifiantStructure != "" { body["identifiant_structure"] = identifiantStructure }
    if raisonSociale != "" { body["raison_sociale_structure"] = raisonSociale }
    if typeIdentifiant != "" { body["type_identifiant_structure"] = typeIdentifiant }
    return c.makeChorusRequest("POST", "/structures/rechercher", body)
}

// ConsulterStructureChorus consulte les détails d'une structure Chorus Pro
func (c *Client) ConsulterStructureChorus(idStructureCpp int) (map[string]interface{}, error) {
    return c.makeChorusRequest("POST", "/structures/consulter", map[string]interface{}{"id_structure_cpp": idStructureCpp})
}

// ObtenirIdChorusDepuisSiret obtient l'ID Chorus Pro d'une structure depuis son SIRET
func (c *Client) ObtenirIdChorusDepuisSiret(siret, typeIdentifiant string) (map[string]interface{}, error) {
    if typeIdentifiant == "" { typeIdentifiant = "SIRET" }
    return c.makeChorusRequest("POST", "/structures/obtenir-id-depuis-siret", map[string]interface{}{"siret": siret, "type_identifiant": typeIdentifiant})
}

// ListerServicesStructureChorus liste les services d'une structure Chorus Pro
func (c *Client) ListerServicesStructureChorus(idStructureCpp int) (map[string]interface{}, error) {
    return c.makeChorusRequest("GET", fmt.Sprintf("/structures/%d/services", idStructureCpp), nil)
}

// SoumettreFactureChorus soumet une facture à Chorus Pro
func (c *Client) SoumettreFactureChorus(factureData map[string]interface{}) (map[string]interface{}, error) {
    return c.makeChorusRequest("POST", "/factures/soumettre", factureData)
}

// ConsulterFactureChorus consulte le statut d'une facture Chorus Pro
func (c *Client) ConsulterFactureChorus(identifiantFactureCpp int) (map[string]interface{}, error) {
    return c.makeChorusRequest("POST", "/factures/consulter", map[string]interface{}{"identifiant_facture_cpp": identifiantFactureCpp})
}

// =========================================================================
// Validation
// =========================================================================

// ValiderPdfFacturx valide un PDF Factur-X
func (c *Client) ValiderPdfFacturx(pdfPath, profil string) (map[string]interface{}, error) {
    if profil == "" { profil = "EN16931" }
    if err := c.EnsureAuthenticated(false); err != nil { return nil, err }

    pdfContent, err := os.ReadFile(pdfPath)
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Error reading PDF: %v", err), nil) }

    var body bytes.Buffer
    writer := multipart.NewWriter(&body)
    part, _ := writer.CreateFormFile("fichier_pdf", filepath.Base(pdfPath))
    part.Write(pdfContent)
    writer.WriteField("profil", profil)
    writer.Close()

    req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/traitement/valider-pdf-facturx", strings.TrimRight(c.APIURL, "/")), &body)
    req.Header.Set("Authorization", "Bearer "+c.accessToken)
    req.Header.Set("Content-Type", writer.FormDataContentType())

    resp, err := http.DefaultClient.Do(req)
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Network error: %v", err), nil) }
    defer resp.Body.Close()

    var result map[string]interface{}; json.NewDecoder(resp.Body).Decode(&result); return result, nil
}

// ValiderXmlFacturx valide un XML Factur-X
func (c *Client) ValiderXmlFacturx(xmlContent, profil string) (map[string]interface{}, error) {
    if profil == "" { profil = "EN16931" }
    if err := c.EnsureAuthenticated(false); err != nil { return nil, err }

    var body bytes.Buffer
    writer := multipart.NewWriter(&body)
    part, _ := writer.CreateFormFile("fichier_xml", "facture.xml")
    part.Write([]byte(xmlContent))
    writer.WriteField("profil", profil)
    writer.Close()

    req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/traitement/valider-xml", strings.TrimRight(c.APIURL, "/")), &body)
    req.Header.Set("Authorization", "Bearer "+c.accessToken)
    req.Header.Set("Content-Type", writer.FormDataContentType())

    resp, err := http.DefaultClient.Do(req)
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Network error: %v", err), nil) }
    defer resp.Body.Close()

    var result map[string]interface{}; json.NewDecoder(resp.Body).Decode(&result); return result, nil
}

// ValiderSignaturePdf valide la signature d'un PDF signé
func (c *Client) ValiderSignaturePdf(pdfPath string) (map[string]interface{}, error) {
    if err := c.EnsureAuthenticated(false); err != nil { return nil, err }

    pdfContent, err := os.ReadFile(pdfPath)
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Error reading PDF: %v", err), nil) }

    var body bytes.Buffer
    writer := multipart.NewWriter(&body)
    part, _ := writer.CreateFormFile("fichier_pdf", filepath.Base(pdfPath))
    part.Write(pdfContent)
    writer.Close()

    req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/traitement/valider-signature-pdf", strings.TrimRight(c.APIURL, "/")), &body)
    req.Header.Set("Authorization", "Bearer "+c.accessToken)
    req.Header.Set("Content-Type", writer.FormDataContentType())

    resp, err := http.DefaultClient.Do(req)
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Network error: %v", err), nil) }
    defer resp.Body.Close()

    var result map[string]interface{}; json.NewDecoder(resp.Body).Decode(&result); return result, nil
}

// =========================================================================
// Signature
// =========================================================================

// SignerPdfOptions contient les options pour signerPdf
type SignerPdfOptions struct {
    Raison, Localisation, Contact string
    UsePadesLt, UseTimestamp bool
}

// SignerPdf signe un PDF avec le certificat configuré côté serveur
func (c *Client) SignerPdf(pdfPath string, opts *SignerPdfOptions) ([]byte, error) {
    if opts == nil { opts = &SignerPdfOptions{UseTimestamp: true} }
    if err := c.EnsureAuthenticated(false); err != nil { return nil, err }

    pdfContent, err := os.ReadFile(pdfPath)
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Error reading PDF: %v", err), nil) }

    var body bytes.Buffer
    writer := multipart.NewWriter(&body)
    part, _ := writer.CreateFormFile("fichier_pdf", filepath.Base(pdfPath))
    part.Write(pdfContent)
    if opts.Raison != "" { writer.WriteField("raison", opts.Raison) }
    if opts.Localisation != "" { writer.WriteField("localisation", opts.Localisation) }
    if opts.Contact != "" { writer.WriteField("contact", opts.Contact) }
    writer.WriteField("use_pades_lt", fmt.Sprintf("%t", opts.UsePadesLt))
    writer.WriteField("use_timestamp", fmt.Sprintf("%t", opts.UseTimestamp))
    writer.Close()

    req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/traitement/signer-pdf", strings.TrimRight(c.APIURL, "/")), &body)
    req.Header.Set("Authorization", "Bearer "+c.accessToken)
    req.Header.Set("Content-Type", writer.FormDataContentType())

    resp, err := http.DefaultClient.Do(req)
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Network error: %v", err), nil) }
    defer resp.Body.Close()

    var result map[string]interface{}; json.NewDecoder(resp.Body).Decode(&result)
    if pdfBase64, ok := result["pdf_signe_base64"].(string); ok {
        return base64.StdEncoding.DecodeString(pdfBase64)
    }
    return nil, NewFactPulseValidationError("Invalid signature response", nil)
}

// GenererCertificatTestOptions contient les options pour genererCertificatTest
type GenererCertificatTestOptions struct {
    CN, Organisation, Email string
    DureeJours, TailleCle int
}

// GenererCertificatTest génère un certificat de test (NON PRODUCTION)
func (c *Client) GenererCertificatTest(opts *GenererCertificatTestOptions) (map[string]interface{}, error) {
    if opts == nil { opts = &GenererCertificatTestOptions{} }
    if err := c.EnsureAuthenticated(false); err != nil { return nil, err }

    body := map[string]interface{}{
        "cn": ifEmpty(opts.CN, "Test Organisation"),
        "organisation": ifEmpty(opts.Organisation, "Test Organisation"),
        "email": ifEmpty(opts.Email, "test@example.com"),
        "duree_jours": ifZero(opts.DureeJours, 365),
        "taille_cle": ifZero(opts.TailleCle, 2048),
    }
    jsonBody, _ := json.Marshal(body)

    req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/traitement/generer-certificat-test", strings.TrimRight(c.APIURL, "/")), bytes.NewReader(jsonBody))
    req.Header.Set("Authorization", "Bearer "+c.accessToken)
    req.Header.Set("Content-Type", "application/json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Network error: %v", err), nil) }
    defer resp.Body.Close()

    var result map[string]interface{}; json.NewDecoder(resp.Body).Decode(&result); return result, nil
}

// =========================================================================
// Workflow complet
// =========================================================================

// GenererFacturxCompletOptions contient les options pour genererFacturxComplet
type GenererFacturxCompletOptions struct {
    Profil, OutputPath, AfnorFlowName, AfnorTrackingID string
    Valider, Signer, SoumettreAfnor bool
    Timeout int64
    SignerPdfOptions *SignerPdfOptions
}

// GenererFacturxComplet génère un PDF Factur-X complet avec validation, signature et soumission optionnelles
func (c *Client) GenererFacturxComplet(facture map[string]interface{}, pdfSourcePath string, opts *GenererFacturxCompletOptions) (map[string]interface{}, error) {
    if opts == nil { opts = &GenererFacturxCompletOptions{Valider: true} }
    profil := ifEmpty(opts.Profil, "EN16931")
    timeout := opts.Timeout; if timeout == 0 { timeout = 120000 }

    result := map[string]interface{}{}

    // 1. Génération
    timeoutPtr := &timeout
    pdfBytes, err := c.GenererFacturxWithOptions(facture, pdfSourcePath, profil, "pdf", true, timeoutPtr)
    if err != nil { return nil, err }
    result["pdfBytes"] = pdfBytes

    // Créer un fichier temporaire
    tempFile, err := os.CreateTemp("", "facturx_*.pdf")
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Error creating temp file: %v", err), nil) }
    defer os.Remove(tempFile.Name())
    tempFile.Write(pdfBytes)
    tempFile.Close()

    // 2. Validation
    if opts.Valider {
        validation, err := c.ValiderPdfFacturx(tempFile.Name(), profil)
        if err != nil { return nil, err }
        result["validation"] = validation
        if estConforme, ok := validation["est_conforme"].(bool); !ok || !estConforme {
            if opts.OutputPath != "" {
                os.WriteFile(opts.OutputPath, pdfBytes, 0644)
                result["pdfPath"] = opts.OutputPath
            }
            return result, nil
        }
    }

    // 3. Signature
    if opts.Signer {
        signedPdf, err := c.SignerPdf(tempFile.Name(), opts.SignerPdfOptions)
        if err != nil { return nil, err }
        pdfBytes = signedPdf
        result["pdfBytes"] = pdfBytes
        result["signature"] = map[string]interface{}{"signe": true}
        os.WriteFile(tempFile.Name(), pdfBytes, 0644)
    }

    // 4. Soumission AFNOR
    if opts.SoumettreAfnor {
        numeroFacture := getStringOrDefault(facture, "numeroFacture", getStringOrDefault(facture, "numero_facture", "FACTURE"))
        flowName := ifEmpty(opts.AfnorFlowName, fmt.Sprintf("Facture %s", numeroFacture))
        trackingID := ifEmpty(opts.AfnorTrackingID, numeroFacture)
        afnorResult, err := c.SoumettreFactureAfnor(tempFile.Name(), flowName, &SoumettreFactureAfnorOptions{TrackingID: trackingID})
        if err != nil { return nil, err }
        result["afnor"] = afnorResult
    }

    // Sauvegarde finale
    if opts.OutputPath != "" {
        os.WriteFile(opts.OutputPath, pdfBytes, 0644)
        result["pdfPath"] = opts.OutputPath
    }

    return result, nil
}

// Utilitaires
func ifEmpty(s, def string) string { if s == "" { return def }; return s }
func ifZero(n, def int) int { if n == 0 { return def }; return n }
func getIntOrDefault(m map[string]interface{}, key string, def int) int {
    if v, ok := m[key]; ok {
        if i, ok := v.(int); ok { return i }
        if f, ok := v.(float64); ok { return int(f) }
    }
    return def
}
func getStringOrDefault(m map[string]interface{}, key, def string) string {
    if v, ok := m[key].(string); ok { return v }
    return def
}

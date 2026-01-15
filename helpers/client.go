package helpers

import (
    "bytes"; "crypto/sha256"; "encoding/base64"; "encoding/hex"; "encoding/json"; "fmt"; "io"; "mime/multipart"
    "net/http"; "net/textproto"; "net/url"; "os"; "path/filepath"; "strconv"; "strings"; "sync"; "time"
)

const DefaultAPIURL = "https://factpulse.fr"
const DefaultPollingInterval = 2000
const DefaultPollingTimeout = 120000

// ChorusProCredentials contains Chorus Pro credentials for Zero-Trust mode.
// These credentials are passed in each request and are never stored server-side.
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

// AFNORCredentials contains AFNOR PDP credentials for Zero-Trust mode.
// The FactPulse API uses these credentials to authenticate with the AFNOR PDP.
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

// Amount formats a value as an amount with 2 decimal places
func Amount(m interface{}) string {
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

// TotalAmount creates a simplified total amount object
func TotalAmount(exclTax, vat, inclTax, amountDue interface{}) map[string]interface{} {
    return TotalAmountWithOptions(exclTax, vat, inclTax, amountDue, nil, "", nil)
}

// TotalAmountWithOptions creates a total amount object with options
func TotalAmountWithOptions(exclTax, vat, inclTax, amountDue, discountInclTax interface{}, discountReason string, prepayment interface{}) map[string]interface{} {
    result := map[string]interface{}{"totalExclTax": Amount(exclTax), "vatAmount": Amount(vat), "totalInclTax": Amount(inclTax), "amountDue": Amount(amountDue)}
    if discountInclTax != nil { result["globalDiscountInclTax"] = Amount(discountInclTax) }
    if discountReason != "" { result["globalDiscountReason"] = discountReason }
    if prepayment != nil { result["prepayment"] = Amount(prepayment) }
    return result
}

// InvoiceLine creates an invoice line (aligned with InvoiceLine in models.py)
func InvoiceLine(number int, description string, quantity, unitPriceExclTax, lineTotalExclTax interface{}) map[string]interface{} {
    return InvoiceLineWithOptions(number, description, quantity, unitPriceExclTax, lineTotalExclTax, "20.00", "S", "LUMP_SUM", nil)
}

// InvoiceLineWithOptions creates an invoice line with options
func InvoiceLineWithOptions(number int, description string, quantity, unitPriceExclTax, lineTotalExclTax interface{}, vatRate, vatCategory, unit string, options map[string]interface{}) map[string]interface{} {
    result := map[string]interface{}{"number": number, "description": description, "quantity": Amount(quantity), "unitPriceExclTax": Amount(unitPriceExclTax), "lineTotalExclTax": Amount(lineTotalExclTax), "vatRateManual": Amount(vatRate), "vatCategory": vatCategory, "unit": unit}
    if options != nil {
        if v, ok := options["reference"]; ok { result["reference"] = v }
        if v, ok := options["discountExclTax"]; ok { result["discountExclTax"] = Amount(v) }
        if v, ok := options["discountReasonCode"]; ok { result["discountReasonCode"] = v }
        if v, ok := options["discountReason"]; ok { result["discountReason"] = v }
        if v, ok := options["periodStartDate"]; ok { result["periodStartDate"] = v }
        if v, ok := options["periodEndDate"]; ok { result["periodEndDate"] = v }
    }
    return result
}

// VatLine creates a VAT line (aligned with VatLine in models.py)
func VatLine(rateManual, baseAmountExclTax, vatAmount interface{}) map[string]interface{} {
    return VatLineWithOptions(rateManual, baseAmountExclTax, vatAmount, "S")
}

// VatLineWithOptions creates a VAT line with options
func VatLineWithOptions(rateManual, baseAmountExclTax, vatAmount interface{}, category string) map[string]interface{} {
    return map[string]interface{}{"rateManual": Amount(rateManual), "baseAmountExclTax": Amount(baseAmountExclTax), "vatAmount": Amount(vatAmount), "category": category}
}

// PostalAddress creates a postal address for the FactPulse API
func PostalAddress(line1, postalCode, city string) map[string]interface{} {
    return PostalAddressWithOptions(line1, postalCode, city, "FR", "", "")
}

// PostalAddressWithOptions creates a postal address with options
func PostalAddressWithOptions(line1, postalCode, city, country, line2, line3 string) map[string]interface{} {
    result := map[string]interface{}{"line1": line1, "postalCode": postalCode, "city": city, "countryCode": country}
    if line2 != "" { result["line2"] = line2 }
    if line3 != "" { result["line3"] = line3 }
    return result
}

// ElectronicAddress creates an electronic address. schemeID: "0009"=SIREN, "0225"=SIRET
func ElectronicAddress(identifier, schemeID string) map[string]interface{} {
    if schemeID == "" { schemeID = "0009" }
    return map[string]interface{}{"identifier": identifier, "schemeId": schemeID}
}

// computeVatIntra computes the French intra-community VAT number from a SIREN
func computeVatIntra(siren string) string {
    if len(siren) != 9 { return "" }
    sirenNum, err := strconv.ParseInt(siren, 10, 64)
    if err != nil { return "" }
    cle := (12 + 3*(sirenNum%97)) % 97
    return fmt.Sprintf("FR%02d%s", cle, siren)
}

// SupplierOptions contains options for creating a supplier
type SupplierOptions struct {
    SupplierID int
    Siren, VatIntra, Iban, Country, AddressLine2 string
    ServiceCode, BankCoordinatesCode int
}

// Supplier creates a supplier with auto-computed SIREN, intra-EU VAT number and addresses
func Supplier(name, siret, addressLine1, postalCode, city string, opts *SupplierOptions) map[string]interface{} {
    if opts == nil { opts = &SupplierOptions{} }
    siren := opts.Siren
    if siren == "" && len(siret) == 14 { siren = siret[:9] }
    vatIntra := opts.VatIntra
    if vatIntra == "" && siren != "" { vatIntra = computeVatIntra(siren) }
    country := opts.Country; if country == "" { country = "FR" }
    result := map[string]interface{}{
        "name": name, "supplierId": opts.SupplierID, "siret": siret,
        "electronicAddress": ElectronicAddress(siret, "0225"),
        "postalAddress": PostalAddressWithOptions(addressLine1, postalCode, city, country, opts.AddressLine2, ""),
    }
    if siren != "" { result["siren"] = siren }
    if vatIntra != "" { result["vatIntra"] = vatIntra }
    if opts.Iban != "" { result["iban"] = opts.Iban }
    if opts.ServiceCode != 0 { result["supplierServiceId"] = opts.ServiceCode }
    if opts.BankCoordinatesCode != 0 { result["supplierBankCoordinatesCode"] = opts.BankCoordinatesCode }
    return result
}

// RecipientOptions contains options for creating a recipient
type RecipientOptions struct {
    Siren, Country, AddressLine2, ExecutingServiceCode string
}

// Recipient creates a recipient with auto-computed SIREN and addresses
func Recipient(name, siret, addressLine1, postalCode, city string, opts *RecipientOptions) map[string]interface{} {
    if opts == nil { opts = &RecipientOptions{} }
    siren := opts.Siren
    if siren == "" && len(siret) == 14 { siren = siret[:9] }
    country := opts.Country; if country == "" { country = "FR" }
    result := map[string]interface{}{
        "name": name, "siret": siret,
        "electronicAddress": ElectronicAddress(siret, "0225"),
        "postalAddress": PostalAddressWithOptions(addressLine1, postalCode, city, country, opts.AddressLine2, ""),
    }
    if siren != "" { result["siren"] = siren }
    if opts.ExecutingServiceCode != "" { result["executingServiceCode"] = opts.ExecutingServiceCode }
    return result
}

// BeneficiaryOptions contains options for creating a beneficiary (factor)
type BeneficiaryOptions struct {
    Siret, Siren, Iban, Bic string
}

// Beneficiary creates a beneficiary (factor) for factoring.
//
// The beneficiary (BG-10 / PayeeTradeParty) is used when payment
// must be made to a third party different from the supplier, typically
// a factor (factoring company).
//
// For factored invoices, you also need to:
// - Use a factored document type (393, 396, 501, 502, 472, 473)
// - Add an ACC note with the subrogation mention
// - The beneficiary's IBAN will be used for payment
func Beneficiary(name string, opts *BeneficiaryOptions) map[string]interface{} {
    if opts == nil { opts = &BeneficiaryOptions{} }
    // Auto-compute SIREN from SIRET
    siren := opts.Siren
    if siren == "" && len(opts.Siret) == 14 { siren = opts.Siret[:9] }

    result := map[string]interface{}{"name": name}
    if opts.Siret != "" { result["siret"] = opts.Siret }
    if siren != "" { result["siren"] = siren }
    if opts.Iban != "" { result["iban"] = opts.Iban }
    if opts.Bic != "" { result["bic"] = opts.Bic }
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
        req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/processing/tasks/%s/status", strings.TrimRight(c.APIURL, "/"), taskID), nil)
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

func FormatAmount(m interface{}) string { return Amount(m) }
func minFloat(a, b float64) float64 { if a < b { return a }; return b }

// GenerateFacturx generates a Factur-X invoice from a map/struct and a source PDF.
func (c *Client) GenerateFacturx(invoiceData interface{}, pdfPath string) ([]byte, error) {
    return c.GenerateFacturxWithOptions(invoiceData, pdfPath, "EN16931", "pdf", true, nil)
}

// GenerateFacturxWithOptions generates a Factur-X invoice with advanced options.
func (c *Client) GenerateFacturxWithOptions(invoiceData interface{}, pdfPath, profile, outputFormat string, sync bool, timeout *int64) ([]byte, error) {
    // Convert data to JSON string
    var jsonData string
    switch v := invoiceData.(type) {
    case string:
        jsonData = v
    case []byte:
        jsonData = string(v)
    default:
        data, err := json.Marshal(invoiceData)
        if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("JSON serialization error: %v", err), nil) }
        jsonData = string(data)
    }

    // Read PDF file
    pdfContent, err := os.ReadFile(pdfPath)
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Error reading PDF: %v", err), nil) }

    if err := c.EnsureAuthenticated(false); err != nil { return nil, err }

    // Build multipart request
    var body bytes.Buffer
    writer := multipart.NewWriter(&body)
    writer.WriteField("invoice_data", jsonData)
    writer.WriteField("profile", profile)
    writer.WriteField("output_format", outputFormat)
    part, _ := writer.CreateFormFile("source_pdf", filepath.Base(pdfPath))
    part.Write(pdfContent)
    writer.Close()

    req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/processing/generate-invoice", strings.TrimRight(c.APIURL, "/")), &body)
    req.Header.Set("Authorization", "Bearer "+c.accessToken)
    req.Header.Set("Content-Type", writer.FormDataContentType())

    client := &http.Client{Timeout: 120 * time.Second}
    resp, err := client.Do(req)
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Network error: %v", err), nil) }
    defer resp.Body.Close()

    if resp.StatusCode == 401 {
        c.ResetAuth()
        if err := c.EnsureAuthenticated(false); err != nil { return nil, err }
        req.Header.Set("Authorization", "Bearer "+c.accessToken)
        resp, err = client.Do(req)
        if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Network error: %v", err), nil) }
        defer resp.Body.Close()
    }

    respBody, _ := io.ReadAll(resp.Body)
    if resp.StatusCode >= 400 {
        // Extract error details from response body
        errorMsg := fmt.Sprintf("API Error (%d)", resp.StatusCode)
        var errors []ValidationErrorDetail

        var errorData map[string]interface{}
        if err := json.Unmarshal(respBody, &errorData); err == nil {
            // FastAPI/Pydantic format: {"detail": [{"loc": [...], "msg": "...", "type": "..."}]}
            if detail, ok := errorData["detail"]; ok {
                if detailList, ok := detail.([]interface{}); ok {
                    errorMsg = "Validation error"
                    for _, item := range detailList {
                        if errItem, ok := item.(map[string]interface{}); ok {
                            loc := ""
                            if locList, ok := errItem["loc"].([]interface{}); ok {
                                locParts := make([]string, len(locList))
                                for i, l := range locList { locParts[i] = fmt.Sprintf("%v", l) }
                                loc = strings.Join(locParts, " -> ")
                            }
                            reason := ""
                            if msg, ok := errItem["msg"].(string); ok { reason = msg }
                            code := ""
                            if t, ok := errItem["type"].(string); ok { code = t }
                            errors = append(errors, ValidationErrorDetail{Level: "ERROR", Item: loc, Reason: reason, Source: "validation", Code: code})
                        }
                    }
                } else if detailStr, ok := detail.(string); ok {
                    errorMsg = detailStr
                }
            } else if errMsg, ok := errorData["errorMessage"].(string); ok {
                errorMsg = errMsg
            }
        }

        return nil, NewFactPulseValidationError(errorMsg, errors)
    }

    var data map[string]interface{}
    json.Unmarshal(respBody, &data)

    if sync {
        if taskID, ok := data["task_id"].(string); ok {
            result, err := c.PollTask(taskID, timeout, nil)
            if err != nil { return nil, err }
            if contenuB64, ok := result["contenu_b64"].(string); ok {
                decoded, _ := base64.StdEncoding.DecodeString(contenuB64)
                return decoded, nil
            }
            if xml, ok := result["contenu_xml"].(string); ok {
                return []byte(xml), nil
            }
            return nil, NewFactPulseValidationError("Unexpected result", nil)
        }
    }

    return respBody, nil
}

// =========================================================================
// AFNOR PDP - Authentication and internal helpers
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

// SubmitInvoiceAfnorOptions contains options for SubmitInvoiceAfnor
type SubmitInvoiceAfnorOptions struct {
    FlowSyntax, FlowProfile, TrackingID string
}

// SubmitInvoiceAfnor submits an invoice to a PDP via the AFNOR API
func (c *Client) SubmitInvoiceAfnor(pdfPath, flowName string, opts *SubmitInvoiceAfnorOptions) (map[string]interface{}, error) {
    if opts == nil { opts = &SubmitInvoiceAfnorOptions{} }
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

// SearchFlowsAfnor searches for AFNOR invoicing flows
func (c *Client) SearchFlowsAfnor(criteria map[string]interface{}) (map[string]interface{}, error) {
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

// DownloadFlowAfnor downloads the PDF file of an AFNOR flow
func (c *Client) DownloadFlowAfnor(flowID string) ([]byte, error) {
    result, err := c.makeAfnorRequest("GET", fmt.Sprintf("/flow/v1/flows/%s", flowID), nil, nil, "")
    if err != nil { return nil, err }
    if raw, ok := result["_raw"].([]byte); ok { return raw, nil }
    return nil, nil
}

// GetIncomingInvoiceAfnor retrieves JSON metadata of an incoming flow (supplier invoice).
// Downloads an incoming flow from the AFNOR PDP and extracts invoice metadata
// into a unified JSON format. Supports Factur-X, CII and UBL.
//
// Note: This endpoint uses FactPulse JWT authentication (not AFNOR OAuth).
// The FactPulse server handles calling the PDP with stored credentials.
//
// flowID: Flow identifier (UUID)
// includeDocument: If true, includes the original document encoded in base64
//
// Returns invoice metadata (supplier, amounts, dates, etc.)
func (c *Client) GetIncomingInvoiceAfnor(flowID string, includeDocument bool) (map[string]interface{}, error) {
    if err := c.EnsureAuthenticated(false); err != nil { return nil, err }

    url := fmt.Sprintf("%s/api/v1/afnor/incoming-flows/%s", strings.TrimRight(c.APIURL, "/"), flowID)
    if includeDocument {
        url += "?include_document=true"
    }

    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", "Bearer "+c.accessToken)

    client := &http.Client{Timeout: 60 * time.Second}
    resp, err := client.Do(req)
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Network error: %v", err), nil) }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    if resp.StatusCode >= 400 {
        return nil, NewFactPulseValidationError(fmt.Sprintf("Incoming flow error (%d): %s", resp.StatusCode, string(body)), nil)
    }

    var result map[string]interface{}; json.Unmarshal(body, &result); return result, nil
}

// HealthcheckAfnor checks the availability of the AFNOR Flow Service
func (c *Client) HealthcheckAfnor() (map[string]interface{}, error) {
    return c.makeAfnorRequest("GET", "/flow/v1/healthcheck", nil, nil, "")
}

// ==================== AFNOR Directory ====================

// SearchSiretAfnor searches for a company by SIRET in the AFNOR directory
func (c *Client) SearchSiretAfnor(siret string) (map[string]interface{}, error) {
    return c.makeAfnorRequest("GET", fmt.Sprintf("/directory/siret/%s", siret), nil, nil, "")
}

// SearchSirenAfnor searches for a company by SIREN in the AFNOR directory
func (c *Client) SearchSirenAfnor(siren string) (map[string]interface{}, error) {
    return c.makeAfnorRequest("GET", fmt.Sprintf("/directory/siren/%s", siren), nil, nil, "")
}

// ListRoutingCodesAfnor lists available routing codes for a SIREN
func (c *Client) ListRoutingCodesAfnor(siren string) (map[string]interface{}, error) {
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

// SearchStructureChorus searches for structures on Chorus Pro
func (c *Client) SearchStructureChorus(structureIdentifier, businessName, identifierType string, restrictPrivate bool) (map[string]interface{}, error) {
    body := map[string]interface{}{"restrict_private_structures": restrictPrivate}
    if structureIdentifier != "" { body["structure_identifier"] = structureIdentifier }
    if businessName != "" { body["structure_business_name"] = businessName }
    if identifierType != "" { body["structure_identifier_type"] = identifierType }
    return c.makeChorusRequest("POST", "/structures/search", body)
}

// GetStructureChorus gets the details of a Chorus Pro structure
func (c *Client) GetStructureChorus(cppStructureId int) (map[string]interface{}, error) {
    return c.makeChorusRequest("POST", "/structures/get", map[string]interface{}{"cpp_structure_id": cppStructureId})
}

// GetChorusIdFromSiret gets the Chorus Pro ID of a structure from its SIRET
func (c *Client) GetChorusIdFromSiret(siret, identifierType string) (map[string]interface{}, error) {
    if identifierType == "" { identifierType = "SIRET" }
    return c.makeChorusRequest("POST", "/structures/get-id-from-siret", map[string]interface{}{"siret": siret, "identifier_type": identifierType})
}

// ListStructureServicesChorus lists the services of a Chorus Pro structure
func (c *Client) ListStructureServicesChorus(cppStructureId int) (map[string]interface{}, error) {
    return c.makeChorusRequest("GET", fmt.Sprintf("/structures/%d/services", cppStructureId), nil)
}

// SubmitInvoiceChorus submits an invoice to Chorus Pro
func (c *Client) SubmitInvoiceChorus(invoiceData map[string]interface{}) (map[string]interface{}, error) {
    return c.makeChorusRequest("POST", "/invoices/submit", invoiceData)
}

// GetInvoiceChorus gets the status of a Chorus Pro invoice
func (c *Client) GetInvoiceChorus(cppInvoiceIdentifier int) (map[string]interface{}, error) {
    return c.makeChorusRequest("POST", "/invoices/get", map[string]interface{}{"cpp_invoice_identifier": cppInvoiceIdentifier})
}

// =========================================================================
// Validation
// =========================================================================

// ValidateFacturxPdfOptions options for ValidateFacturxPdf
type ValidateFacturxPdfOptions struct {
    Profile    string // Factur-X profile (MINIMUM, BASIC, EN16931, EXTENDED). If empty, auto-detected.
    UseVerapdf bool   // Enable strict PDF/A validation with VeraPDF (default: false)
}

// ValidateFacturxPdf validates a Factur-X PDF
func (c *Client) ValidateFacturxPdf(pdfPath string, opts *ValidateFacturxPdfOptions) (map[string]interface{}, error) {
    if err := c.EnsureAuthenticated(false); err != nil { return nil, err }

    pdfContent, err := os.ReadFile(pdfPath)
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Error reading PDF: %v", err), nil) }

    var body bytes.Buffer
    writer := multipart.NewWriter(&body)
    part, _ := writer.CreateFormFile("pdf_file", filepath.Base(pdfPath))
    part.Write(pdfContent)
    if opts != nil && opts.Profile != "" {
        writer.WriteField("profile", opts.Profile)
    }
    useVerapdf := "false"
    if opts != nil && opts.UseVerapdf {
        useVerapdf = "true"
    }
    writer.WriteField("use_verapdf", useVerapdf)
    writer.Close()

    req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/processing/validate-facturx-pdf", strings.TrimRight(c.APIURL, "/")), &body)
    req.Header.Set("Authorization", "Bearer "+c.accessToken)
    req.Header.Set("Content-Type", writer.FormDataContentType())

    resp, err := http.DefaultClient.Do(req)
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Network error: %v", err), nil) }
    defer resp.Body.Close()

    var result map[string]interface{}; json.NewDecoder(resp.Body).Decode(&result); return result, nil
}

// ValidateFacturxXml validates a Factur-X XML
func (c *Client) ValidateFacturxXml(xmlContent, profile string) (map[string]interface{}, error) {
    if profile == "" { profile = "EN16931" }
    if err := c.EnsureAuthenticated(false); err != nil { return nil, err }

    var body bytes.Buffer
    writer := multipart.NewWriter(&body)
    part, _ := writer.CreateFormFile("xml_file", "invoice.xml")
    part.Write([]byte(xmlContent))
    writer.WriteField("profile", profile)
    writer.Close()

    req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/processing/validate-xml", strings.TrimRight(c.APIURL, "/")), &body)
    req.Header.Set("Authorization", "Bearer "+c.accessToken)
    req.Header.Set("Content-Type", writer.FormDataContentType())

    resp, err := http.DefaultClient.Do(req)
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Network error: %v", err), nil) }
    defer resp.Body.Close()

    var result map[string]interface{}; json.NewDecoder(resp.Body).Decode(&result); return result, nil
}

// ValidatePdfSignature validates the signature of a signed PDF
func (c *Client) ValidatePdfSignature(pdfPath string) (map[string]interface{}, error) {
    if err := c.EnsureAuthenticated(false); err != nil { return nil, err }

    pdfContent, err := os.ReadFile(pdfPath)
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Error reading PDF: %v", err), nil) }

    var body bytes.Buffer
    writer := multipart.NewWriter(&body)
    part, _ := writer.CreateFormFile("pdf_file", filepath.Base(pdfPath))
    part.Write(pdfContent)
    writer.Close()

    req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/processing/validate-pdf-signature", strings.TrimRight(c.APIURL, "/")), &body)
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

// SignPdfOptions contains options for SignPdf
type SignPdfOptions struct {
    Reason, Location, Contact string
    UsePadesLt, UseTimestamp bool
}

// SignPdf signs a PDF with the server-configured certificate
func (c *Client) SignPdf(pdfPath string, opts *SignPdfOptions) ([]byte, error) {
    if opts == nil { opts = &SignPdfOptions{UseTimestamp: true} }
    if err := c.EnsureAuthenticated(false); err != nil { return nil, err }

    pdfContent, err := os.ReadFile(pdfPath)
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Error reading PDF: %v", err), nil) }

    var body bytes.Buffer
    writer := multipart.NewWriter(&body)
    part, _ := writer.CreateFormFile("pdf_file", filepath.Base(pdfPath))
    part.Write(pdfContent)
    if opts.Reason != "" { writer.WriteField("reason", opts.Reason) }
    if opts.Location != "" { writer.WriteField("location", opts.Location) }
    if opts.Contact != "" { writer.WriteField("contact", opts.Contact) }
    writer.WriteField("use_pades_lt", fmt.Sprintf("%t", opts.UsePadesLt))
    writer.WriteField("use_timestamp", fmt.Sprintf("%t", opts.UseTimestamp))
    writer.Close()

    req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/processing/sign-pdf", strings.TrimRight(c.APIURL, "/")), &body)
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

// GenerateTestCertificateOptions contains options for GenerateTestCertificate
type GenerateTestCertificateOptions struct {
    CN, Organisation, Email string
    ValidityDays, KeySize int
}

// GenerateTestCertificate generates a test certificate (NOT FOR PRODUCTION)
func (c *Client) GenerateTestCertificate(opts *GenerateTestCertificateOptions) (map[string]interface{}, error) {
    if opts == nil { opts = &GenerateTestCertificateOptions{} }
    if err := c.EnsureAuthenticated(false); err != nil { return nil, err }

    body := map[string]interface{}{
        "cn": ifEmpty(opts.CN, "Test Organisation"),
        "organisation": ifEmpty(opts.Organisation, "Test Organisation"),
        "email": ifEmpty(opts.Email, "test@example.com"),
        "validity_days": ifZero(opts.ValidityDays, 365),
        "key_size": ifZero(opts.KeySize, 2048),
    }
    jsonBody, _ := json.Marshal(body)

    req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/processing/generate-test-certificate", strings.TrimRight(c.APIURL, "/")), bytes.NewReader(jsonBody))
    req.Header.Set("Authorization", "Bearer "+c.accessToken)
    req.Header.Set("Content-Type", "application/json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Network error: %v", err), nil) }
    defer resp.Body.Close()

    var result map[string]interface{}; json.NewDecoder(resp.Body).Decode(&result); return result, nil
}

// =========================================================================
// Complete workflow
// =========================================================================

// GenerateFacturxCompleteOptions contains options for GenerateFacturxComplete
type GenerateFacturxCompleteOptions struct {
    Profile, OutputPath, AfnorFlowName, AfnorTrackingID string
    Validate, Sign, SubmitAfnor bool
    Timeout int64
    SignPdfOptions *SignPdfOptions
}

// GenerateFacturxComplete generates a complete Factur-X PDF with optional validation, signature and submission
func (c *Client) GenerateFacturxComplete(invoice map[string]interface{}, pdfSourcePath string, opts *GenerateFacturxCompleteOptions) (map[string]interface{}, error) {
    if opts == nil { opts = &GenerateFacturxCompleteOptions{Validate: true} }
    profile := ifEmpty(opts.Profile, "EN16931")
    timeout := opts.Timeout; if timeout == 0 { timeout = 120000 }

    result := map[string]interface{}{}

    // 1. Generation
    timeoutPtr := &timeout
    pdfBytes, err := c.GenerateFacturxWithOptions(invoice, pdfSourcePath, profile, "pdf", true, timeoutPtr)
    if err != nil { return nil, err }
    result["pdfBytes"] = pdfBytes

    // Create temporary file
    tempFile, err := os.CreateTemp("", "facturx_*.pdf")
    if err != nil { return nil, NewFactPulseValidationError(fmt.Sprintf("Error creating temp file: %v", err), nil) }
    defer os.Remove(tempFile.Name())
    tempFile.Write(pdfBytes)
    tempFile.Close()

    // 2. Validation
    if opts.Validate {
        validation, err := c.ValidateFacturxPdf(tempFile.Name(), &ValidateFacturxPdfOptions{Profile: profile})
        if err != nil { return nil, err }
        result["validation"] = validation
        if isCompliant, ok := validation["is_compliant"].(bool); !ok || !isCompliant {
            if opts.OutputPath != "" {
                os.WriteFile(opts.OutputPath, pdfBytes, 0644)
                result["pdfPath"] = opts.OutputPath
            }
            return result, nil
        }
    }

    // 3. Signature
    if opts.Sign {
        signedPdf, err := c.SignPdf(tempFile.Name(), opts.SignPdfOptions)
        if err != nil { return nil, err }
        pdfBytes = signedPdf
        result["pdfBytes"] = pdfBytes
        result["signature"] = map[string]interface{}{"signed": true}
        os.WriteFile(tempFile.Name(), pdfBytes, 0644)
    }

    // 4. AFNOR submission
    if opts.SubmitAfnor {
        invoiceNumber := getStringOrDefault(invoice, "number", getStringOrDefault(invoice, "invoice_number", "INVOICE"))
        flowName := ifEmpty(opts.AfnorFlowName, fmt.Sprintf("Invoice %s", invoiceNumber))
        trackingID := ifEmpty(opts.AfnorTrackingID, invoiceNumber)
        afnorResult, err := c.SubmitInvoiceAfnor(tempFile.Name(), flowName, &SubmitInvoiceAfnorOptions{TrackingID: trackingID})
        if err != nil { return nil, err }
        result["afnor"] = afnorResult
    }

    // Final save
    if opts.OutputPath != "" {
        os.WriteFile(opts.OutputPath, pdfBytes, 0644)
        result["pdfPath"] = opts.OutputPath
    }

    return result, nil
}

// Utilities
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

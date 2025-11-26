package helpers

import (
    "bytes"; "encoding/json"; "fmt"; "io"; "net/http"; "strings"; "sync"; "time"
)

const DefaultAPIURL = "https://factpulse.fr"
const DefaultPollingInterval = 2000
const DefaultPollingTimeout = 120000

type Client struct {
    Email, Password, APIURL, ClientUID string
    PollingInterval, PollingTimeout int64
    accessToken, refreshToken string
    tokenExpiresAt time.Time
    mu sync.Mutex
}

func NewClient(email, password string) *Client {
    return &Client{Email: email, Password: password, APIURL: DefaultAPIURL, PollingInterval: DefaultPollingInterval, PollingTimeout: DefaultPollingTimeout}
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
        time.Sleep(time.Duration(currentInterval) * time.Millisecond); currentInterval = min(currentInterval*1.5, 10000)
    }
}

func FormatMontant(m interface{}) string {
    if m == nil { return "0.00" }
    switch v := m.(type) {
    case float64: return fmt.Sprintf("%.2f", v)
    case int: return fmt.Sprintf("%.2f", float64(v))
    case string: return v
    default: return "0.00"
    }
}
func min(a, b float64) float64 { if a < b { return a }; return b }

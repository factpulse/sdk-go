// FactPulse SDK - Thin HTTP wrapper with auto-polling.
//
// Usage:
//
//	client := NewClient("email", "password", "client_uid")
//
//	// POST /api/v1/processing/invoices/submit-complete-async
//	result, err := client.Post("processing/invoices/submit-complete-async", map[string]any{
//	    "invoiceData": invoiceData,
//	    "destination": map[string]any{"type": "afnor"},
//	})
//	pdfBytes := result["content"].([]byte) // auto-decoded, auto-polled
package helpers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"sync"
	"time"
)

const DefaultAPIURL = "https://factpulse.fr"

type Client struct {
	apiURL, email, password, clientUID string
	timeout, pollingTimeout time.Duration
	httpClient              *http.Client
	token                   string
	tokenExpiresAt          time.Time
	mu                      sync.Mutex
}

func NewClient(email, password, clientUID string) *Client {
	return NewClientWithConfig(email, password, clientUID, DefaultAPIURL, 60*time.Second, 120*time.Second)
}

func NewClientWithConfig(email, password, clientUID, apiURL string, timeout, pollingTimeout time.Duration) *Client {
	if apiURL == "" {
		apiURL = DefaultAPIURL
	}
	return &Client{
		apiURL:         apiURL,
		email:          email,
		password:       password,
		clientUID:      clientUID,
		timeout:        timeout,
		pollingTimeout: pollingTimeout,
		httpClient:     &http.Client{Timeout: timeout},
	}
}

// Post sends a POST request to /api/v1/{path} (JSON body)
func (c *Client) Post(path string, data map[string]any) (map[string]any, error) {
	return c.request("POST", path, data, true)
}

// PostMultipart sends a POST request with multipart/form-data (for file uploads)
func (c *Client) PostMultipart(path string, data map[string]string, files map[string][]byte) (map[string]any, error) {
	return c.requestMultipart(path, data, files, true)
}

// Get sends a GET request to /api/v1/{path}
func (c *Client) Get(path string) (map[string]any, error) {
	return c.request("GET", path, nil, true)
}

func (c *Client) request(method, path string, data map[string]any, retryAuth bool) (map[string]any, error) {
	if err := c.ensureAuth(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/api/v1/%s", c.apiURL, path)
	var body io.Reader
	if data != nil {
		jsonData, _ := json.Marshal(data)
		body = bytes.NewReader(jsonData)
	}

	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Authorization", "Bearer "+c.token)
	if method == "POST" {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, &FactPulseError{Message: fmt.Sprintf("Network error: %v", err)}
	}
	defer resp.Body.Close()

	if resp.StatusCode == 401 && retryAuth {
		c.invalidateToken()
		return c.request(method, path, data, false)
	}

	result, err := c.parseResponse(resp)
	if err != nil {
		return nil, err
	}

	// Auto-poll: support both taskId (camelCase) and task_id (snake_case)
	taskID, _ := result["taskId"].(string)
	if taskID == "" {
		taskID, _ = result["task_id"].(string)
	}
	if taskID != "" {
		return c.poll(taskID)
	}

	// Auto-decode: support both content_b64 and contentB64
	b64, _ := result["content_b64"].(string)
	if b64 == "" {
		b64, _ = result["contentB64"].(string)
	}
	if b64 != "" {
		decoded, _ := base64.StdEncoding.DecodeString(b64)
		result["content"] = decoded
		delete(result, "content_b64")
		delete(result, "contentB64")
	}

	return result, nil
}

func (c *Client) requestMultipart(path string, data map[string]string, files map[string][]byte, retryAuth bool) (map[string]any, error) {
	if err := c.ensureAuth(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/api/v1/%s", c.apiURL, path)

	// Create multipart form
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Add form fields
	for key, value := range data {
		writer.WriteField(key, value)
	}

	// Add files
	for key, content := range files {
		part, _ := writer.CreateFormFile(key, key)
		part.Write(content)
	}
	writer.Close()

	req, _ := http.NewRequest("POST", url, &buf)
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, &FactPulseError{Message: fmt.Sprintf("Network error: %v", err)}
	}
	defer resp.Body.Close()

	if resp.StatusCode == 401 && retryAuth {
		c.invalidateToken()
		return c.requestMultipart(path, data, files, false)
	}

	result, err := c.parseResponse(resp)
	if err != nil {
		return nil, err
	}

	// Auto-poll
	taskID, _ := result["taskId"].(string)
	if taskID == "" {
		taskID, _ = result["task_id"].(string)
	}
	if taskID != "" {
		return c.poll(taskID)
	}

	// Auto-decode
	b64, _ := result["content_b64"].(string)
	if b64 == "" {
		b64, _ = result["contentB64"].(string)
	}
	if b64 != "" {
		decoded, _ := base64.StdEncoding.DecodeString(b64)
		result["content"] = decoded
		delete(result, "content_b64")
		delete(result, "contentB64")
	}

	return result, nil
}

func (c *Client) parseResponse(resp *http.Response) (map[string]any, error) {
	body, _ := io.ReadAll(resp.Body)
	var data map[string]any
	json.Unmarshal(body, &data)

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		if data == nil {
			return map[string]any{}, nil
		}
		return data, nil
	}

	msg := fmt.Sprintf("HTTP %d", resp.StatusCode)
	var details []any

	if data != nil {
		if detail, ok := data["detail"].([]any); ok {
			details = detail
			var msgs []string
			for _, e := range detail {
				if err, ok := e.(map[string]any); ok {
					loc := err["loc"].([]any)
					msgs = append(msgs, fmt.Sprintf("%v: %v", loc[len(loc)-1], err["msg"]))
				}
			}
			msg = "Validation error: " + fmt.Sprint(msgs)
		} else if detail, ok := data["detail"].(string); ok {
			msg = detail
		} else if errMsg, ok := data["errorMessage"].(string); ok {
			msg = errMsg
		}
	}

	return nil, &FactPulseError{Message: msg, StatusCode: resp.StatusCode, Details: details}
}

func (c *Client) poll(taskID string) (map[string]any, error) {
	start := time.Now()
	interval := time.Second

	for {
		elapsed := time.Since(start)
		if elapsed >= c.pollingTimeout {
			return nil, &FactPulseError{Message: fmt.Sprintf("Polling timeout after %v for task %s", c.pollingTimeout, taskID)}
		}

		if err := c.ensureAuth(); err != nil {
			return nil, err
		}

		req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/processing/tasks/%s/status", c.apiURL, taskID), nil)
		req.Header.Set("Authorization", "Bearer "+c.token)

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, &FactPulseError{Message: fmt.Sprintf("Network error while polling: %v", err)}
		}

		if resp.StatusCode == 401 {
			resp.Body.Close()
			c.invalidateToken()
			continue
		}

		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		var data map[string]any
		json.Unmarshal(body, &data)

		status, _ := data["status"].(string)

		if status == "SUCCESS" {
			result, _ := data["result"].(map[string]any)
			if result == nil {
				result = map[string]any{}
			}
			// Support both content_b64 and contentB64
			b64, _ := result["content_b64"].(string)
			if b64 == "" {
				b64, _ = result["contentB64"].(string)
			}
			if b64 != "" {
				decoded, _ := base64.StdEncoding.DecodeString(b64)
				result["content"] = decoded
				delete(result, "content_b64")
				delete(result, "contentB64")
			}
			return result, nil
		}

		if status == "FAILURE" {
			res, _ := data["result"].(map[string]any)
			errMsg := "Task failed"
			if res != nil {
				if m, ok := res["errorMessage"].(string); ok {
					errMsg = m
				}
			}
			return nil, &FactPulseError{Message: errMsg}
		}

		sleepDuration := interval
		if remaining := c.pollingTimeout - elapsed; remaining < interval {
			sleepDuration = remaining
		}
		time.Sleep(sleepDuration)
		if interval < 10*time.Second {
			interval = time.Duration(float64(interval) * 1.5)
		}
	}
}

func (c *Client) ensureAuth() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if time.Now().After(c.tokenExpiresAt) {
		return c.refreshToken()
	}
	return nil
}

func (c *Client) refreshToken() error {
	payload, _ := json.Marshal(map[string]string{"username": c.email, "password": c.password, "client_uid": c.clientUID})
	resp, err := c.httpClient.Post(c.apiURL+"/api/token/", "application/json", bytes.NewReader(payload))
	if err != nil {
		return &FactPulseError{Message: fmt.Sprintf("Auth network error: %v", err)}
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return &FactPulseError{Message: fmt.Sprintf("Authentication failed: HTTP %d", resp.StatusCode), StatusCode: resp.StatusCode}
	}

	var data map[string]string
	json.NewDecoder(resp.Body).Decode(&data)
	c.token = data["access"]
	c.tokenExpiresAt = time.Now().Add(28 * time.Minute)
	return nil
}

func (c *Client) invalidateToken() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.tokenExpiresAt = time.Time{}
}

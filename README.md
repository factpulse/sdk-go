# FactPulse SDK Go

Official Go client for the FactPulse API - French electronic invoicing.

## Features

- **Factur-X**: Generation and validation of electronic invoices (MINIMUM, BASIC, EN16931, EXTENDED profiles)
- **Chorus Pro**: Integration with the French public invoicing platform
- **AFNOR PDP/PA**: Submission of flows compliant with XP Z12-013 standard
- **Electronic signature**: PDF signing (PAdES-B-B, PAdES-B-T, PAdES-B-LT)
- **Thin HTTP wrapper**: Generic `Post()` and `Get()` methods with automatic JWT auth and polling

## Installation

```bash
go get github.com/factpulse/sdk-go/v3
```

## Quick Start

```go
package main

import (
    "encoding/base64"
    "fmt"
    "log"
    "os"

    factpulse "github.com/factpulse/sdk-go/v3"
)

func main() {
    // Create the client
    client := factpulse.NewClient(
        "your_email@example.com",
        "your_password",
        "your-client-uuid", // From dashboard: Configuration > Clients
    )

    // Read your source PDF
    pdfBytes, _ := os.ReadFile("source_invoice.pdf")
    pdfB64 := base64.StdEncoding.EncodeToString(pdfBytes)

    // Generate Factur-X and submit to PDP in one call
    result, err := client.Post("processing/invoices/submit-complete-async", map[string]interface{}{
        "invoiceData": map[string]interface{}{
            "number": "INV-2025-001",
            "supplier": map[string]interface{}{
                "siret":          "12345678901234",
                "iban":           "FR7630001007941234567890185",
                "routingAddress": "12345678901234",
            },
            "recipient": map[string]interface{}{
                "siret":          "98765432109876",
                "routingAddress": "98765432109876",
            },
            "lines": []map[string]interface{}{
                {
                    "description": "Consulting services",
                    "quantity":    10,
                    "unitPrice":   100.0,
                    "vatRate":     20.0,
                },
            },
        },
        "sourcePdf":   pdfB64,
        "profile":     "EN16931",
        "destination": map[string]interface{}{"type": "afnor"},
    })
    if err != nil {
        log.Fatal(err)
    }

    // PDF is in result["content"] (auto-polled, auto-decoded)
    facturxPdf := result["content"].([]byte)
    os.WriteFile("facturx_invoice.pdf", facturxPdf, 0644)

    afnorResult := result["afnorResult"].(map[string]interface{})
    fmt.Printf("Flow ID: %s\n", afnorResult["flowId"])
}
```

## API Methods

The SDK provides two generic methods that map directly to API endpoints:

```go
// POST /api/v1/{path}
result, err := client.Post("path/to/endpoint", data)

// GET /api/v1/{path}
result, err := client.Get("path/to/endpoint", params)
```

### Common Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `processing/invoices/submit-complete-async` | POST | Generate Factur-X + submit to PDP |
| `processing/generate-invoice` | POST | Generate Factur-X XML or PDF |
| `processing/validate-xml` | POST | Validate Factur-X XML |
| `processing/validate-facturx-pdf` | POST | Validate Factur-X PDF |
| `processing/sign-pdf` | POST | Sign PDF with certificate |
| `afnor/flow/v1/flows` | POST | Submit flow to AFNOR PDP |
| `afnor/incoming-flows/{flow_id}` | GET | Get incoming invoice |
| `chorus-pro/factures/soumettre` | POST | Submit to Chorus Pro |

## Webhooks

Instead of polling, you can receive results via webhook by adding `callbackUrl`:

```go
result, err := client.Post("processing/invoices/submit-complete-async", map[string]interface{}{
    "invoiceData":  invoiceData,
    "sourcePdf":    pdfB64,
    "destination":  map[string]interface{}{"type": "afnor"},
    "callbackUrl":  "https://your-server.com/webhook/factpulse",
    "webhookMode":  "INLINE", // or "DOWNLOAD_URL"
})

taskId := result["taskId"].(string)
// Result will be POSTed to your webhook URL
```

### Webhook Receiver Example

```go
package main

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "io"
    "net/http"
    "strings"
)

const webhookSecret = "your-shared-secret"

func verifySignature(payload []byte, signature string) bool {
    if !strings.HasPrefix(signature, "sha256=") {
        return false
    }
    mac := hmac.New(sha256.New, []byte(webhookSecret))
    mac.Write(payload)
    expected := hex.EncodeToString(mac.Sum(nil))
    return hmac.Equal([]byte(expected), []byte(signature[7:]))
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
    payload, _ := io.ReadAll(r.Body)
    signature := r.Header.Get("X-Webhook-Signature")

    if !verifySignature(payload, signature) {
        http.Error(w, `{"error":"Invalid signature"}`, http.StatusUnauthorized)
        return
    }

    var event map[string]interface{}
    json.Unmarshal(payload, &event)

    eventType := event["event_type"].(string)
    data := event["data"].(map[string]interface{})

    switch eventType {
    case "submission.completed":
        afnorResult := data["afnorResult"].(map[string]interface{})
        fmt.Printf("Invoice submitted: %s\n", afnorResult["flowId"])
    case "submission.failed":
        fmt.Printf("Submission failed: %v\n", data["error"])
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"status":"received"}`))
}

func main() {
    http.HandleFunc("/webhook/factpulse", webhookHandler)
    http.ListenAndServe(":8080", nil)
}
```

### Webhook Event Types

| Event | Description |
|-------|-------------|
| `generation.completed` | Factur-X generated successfully |
| `generation.failed` | Generation failed |
| `validation.completed` | Validation passed |
| `validation.failed` | Validation failed |
| `signature.completed` | PDF signed |
| `submission.completed` | Submitted to PDP/Chorus |
| `submission.failed` | Submission failed |

## Zero-Storage Mode

Pass PDP credentials directly in the request (no server-side storage):

```go
result, err := client.Post("processing/invoices/submit-complete-async", map[string]interface{}{
    "invoiceData": invoiceData,
    "sourcePdf":   pdfB64,
    "destination": map[string]interface{}{
        "type":           "afnor",
        "flowServiceUrl": "https://api.pdp.example.com/flow/v1",
        "tokenUrl":       "https://auth.pdp.example.com/oauth/token",
        "clientId":       "your_pdp_client_id",
        "clientSecret":   "your_pdp_client_secret",
    },
})
```

## Error Handling

```go
import factpulse "github.com/factpulse/sdk-go/v3"

result, err := client.Post("processing/validate-xml", data)
if err != nil {
    if fpErr, ok := err.(*factpulse.Error); ok {
        fmt.Printf("Error: %s\n", fpErr.Message)
        fmt.Printf("Status code: %d\n", fpErr.StatusCode)
        fmt.Printf("Details: %v\n", fpErr.Details)
    }
}
```

## Resources

- **API Documentation**: https://factpulse.fr/api/facturation/documentation
- **Support**: contact@factpulse.fr

## License

MIT License - Copyright (c) 2025 FactPulse

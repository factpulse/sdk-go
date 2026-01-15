# FactPulse SDK Go

Official Go client for the FactPulse API - French electronic invoicing.

## Features

- **Factur-X**: Generation and validation of electronic invoices (MINIMUM, BASIC, EN16931, EXTENDED profiles)
- **Chorus Pro**: Integration with the French public sector invoicing platform
- **AFNOR PDP/PA**: Submission of flows compliant with the XP Z12-013 standard
- **Electronic signature**: PDF signature (PAdES-B-B, PAdES-B-T, PAdES-B-LT)
- **Simplified client**: JWT authentication and integrated polling via `helpers`

## Installation

```bash
go get github.com/factpulse/sdk-go
```

## Quick Start

The `helpers` package provides a simplified API with automatic authentication and polling:

```go
package main

import (
    "log"
    "os"

    "github.com/factpulse/sdk-go/helpers"
)

func main() {
    // Create the client
    client := helpers.NewClient(
        "your_email@example.com",
        "your_password",
    )

    // Build the invoice using simplified format (auto-calculates totals)
    invoiceData := map[string]interface{}{
        "number": "INV-2025-001",
        "supplier": map[string]interface{}{
            "name":  "My Company SAS",
            "siret": "12345678901234",
            "iban":  "FR7630001007941234567890185",
        },
        "recipient": map[string]interface{}{
            "name":  "Client SARL",
            "siret": "98765432109876",
        },
        "lines": []map[string]interface{}{
            {
                "description": "Consulting services",
                "quantity":    10,
                "unitPrice":   100.0,
                "vatRate":     20,
            },
        },
    }

    // Generate the Factur-X PDF
    pdfBytes, err := client.GenerateFacturx(invoiceData, "source_invoice.pdf")
    if err != nil {
        log.Fatal(err)
    }

    os.WriteFile("invoice_facturx.pdf", pdfBytes, 0644)
}
```

## Available Helpers (helpers package)

### Amount(value)

Converts a value to a formatted string for monetary amounts.

```go
helpers.Amount(1234.5)      // "1234.50"
helpers.Amount("1234.56")   // "1234.56"
helpers.Amount(nil)         // "0.00"
```

### TotalAmount(excludingTax, vat, includingTax, due)

Creates a complete TotalAmount object.

```go
total := helpers.TotalAmount(1000.00, 200.00, 1200.00, 1200.00)

// With options
total := helpers.TotalAmountWithOptions(
    1000.00, 200.00, 1200.00, 1200.00,
    &helpers.TotalAmountOptions{
        DiscountIncludingTax: 50.00,
        DiscountReason:       "Loyalty",
        Prepayment:           100.00,
    },
)
```

### InvoiceLine(number, description, quantity, unitPrice, lineTotal)

Creates an invoice line.

```go
line := helpers.InvoiceLineWithOptions(
    1,
    "Consulting services",
    5,
    200.00,
    1000.00,  // lineTotal required
    &helpers.InvoiceLineOptions{
        VatRate:     "VAT20",      // Or ManualVatRate: "20.00"
        VatCategory: "S",          // S, Z, E, AE, K
        Unit:        "HOUR",       // FIXED, PIECE, HOUR, DAY...
        Reference:   "REF-001",
    },
)
```

### VatLine(baseExcludingTax, vatAmount, options)

Creates a VAT breakdown line.

```go
vat := helpers.VatLine(1000.00, 200.00, &helpers.VatLineOptions{
    Rate:     "VAT20",       // Or ManualRate: "20.00"
    Category: "S",           // S, Z, E, AE, K
})
```

### PostalAddress(line1, postalCode, city, options)

Creates a structured postal address.

```go
address := helpers.PostalAddress(
    "123 Republic Street",
    "75001",
    "Paris",
    &helpers.PostalAddressOptions{
        Country: "FR",          // Default: "FR"
        Line2:   "Building A",  // Optional
    },
)
```

### Supplier(name, siret, addressLine1, postalCode, city, options)

Creates a complete supplier with automatic calculation of SIREN and intra-community VAT.

```go
s := helpers.Supplier(
    "My Company SAS",
    "12345678901234",
    "123 Example Street",
    "75001",
    "Paris",
    &helpers.SupplierOptions{
        IBAN: "FR7630006000011234567890189",
    },
)
// SIREN and intra-community VAT automatically calculated
```

### Recipient(name, siret, addressLine1, postalCode, city, options)

Creates a recipient (customer) with automatic calculation of SIREN.

```go
r := helpers.Recipient(
    "Client SARL",
    "98765432109876",
    "456 Test Avenue",
    "69001",
    "Lyon",
    nil,
)
```

## Zero-Trust Mode (Chorus Pro / AFNOR)

To pass your own credentials without server-side storage:

```go
chorusCreds := &helpers.ChorusProCredentials{
    PisteClientID:     "your_client_id",
    PisteClientSecret: "your_client_secret",
    ChorusProLogin:    "your_login",
    ChorusProPassword: "your_password",
    Sandbox:           true,
}

afnorCreds := &helpers.AFNORCredentials{
    FlowServiceURL: "https://api.pdp.fr/flow/v1",
    TokenURL:       "https://auth.pdp.fr/oauth/token",
    ClientID:       "your_client_id",
    ClientSecret:   "your_client_secret",
}

client := helpers.NewClientWithCredentials(
    "your_email@example.com",
    "your_password",
    chorusCreds,
    afnorCreds,
)
```

## Resources

- **API Documentation**: https://factpulse.fr/api/facturation/documentation
- **Support**: contact@factpulse.fr

## License

MIT License - Copyright (c) 2025 FactPulse

# FactPulse SDK Go

Official Go client for the FactPulse API.

## Installation

```bash
go get github.com/factpulse/sdk-go
```

## Quick Start

```go
package main

import (
    "fmt"
    "os"
    factpulse "github.com/factpulse/sdk-go"
)

func main() {
    client, err := factpulse.NewClient(
        factpulse.WithEmail("your_email@example.com"),
        factpulse.WithPassword("your_password"),
    )
    if err != nil {
        panic(err)
    }

    // Generate a Factur-X invoice
    sourcePdf, _ := os.ReadFile("source.pdf")

    pdfBytes, err := client.GenerateFacturx(factpulse.GenerateFacturxRequest{
        InvoiceData: factpulse.InvoiceData{
            Number: "INV-2025-001",
            Supplier: factpulse.Supplier{Name: "My Company", Siret: "12345678901234"},
            Recipient: factpulse.Recipient{Name: "Client", Siret: "98765432109876"},
        },
        PdfSource: sourcePdf,
        Profile:   "EN16931",
    })
    if err != nil {
        panic(err)
    }

    os.WriteFile("facturx.pdf", pdfBytes, 0644)
}
```

## License

MIT

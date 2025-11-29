# FactPulse SDK Go

Client Go officiel pour l'API FactPulse - Facturation électronique française.

## Fonctionnalités

- **Factur-X** : Génération et validation de factures électroniques (profils MINIMUM, BASIC, EN16931, EXTENDED)
- **Chorus Pro** : Intégration avec la plateforme de facturation publique française
- **AFNOR PDP/PA** : Soumission de flux conformes à la norme XP Z12-013
- **Signature électronique** : Signature PDF (PAdES-B-B, PAdES-B-T, PAdES-B-LT)
- **Client simplifié** : Authentification JWT et polling intégrés via `helpers`

## Installation

```bash
go get github.com/factpulse/sdk-go
```

## Démarrage rapide

Le package `helpers` offre une API simplifiée avec authentification et polling automatiques :

```go
package main

import (
    "context"
    "log"
    "os"

    "github.com/factpulse/sdk-go/helpers"
)

func main() {
    // Créer le client
    client := helpers.NewClient(
        "votre_email@example.com",
        "votre_mot_de_passe",
    )

    // Construire la facture avec les helpers
    factureData := map[string]interface{}{
        "numeroFacture": "FAC-2025-001",
        "dateFacture":   "2025-01-15",
        "fournisseur": helpers.Fournisseur(
            "Mon Entreprise SAS", "12345678901234",
            "123 Rue Example", "75001", "Paris", nil,
        ),
        "destinataire": helpers.Destinataire(
            "Client SARL", "98765432109876",
            "456 Avenue Test", "69001", "Lyon", nil,
        ),
        "montantTotal": helpers.MontantTotal(1000.00, 200.00, 1200.00, 1200.00),
        "lignesDePoste": []interface{}{
            helpers.LigneDePoste(1, "Prestation de conseil", 10, 100.00, 1000.00),
        },
        "lignesDeTva": []interface{}{
            helpers.LigneDeTva(1000.00, 200.00, nil),
        },
    }

    // Générer le PDF Factur-X
    ctx := context.Background()
    pdfBytes, err := client.GenererFacturx(ctx, factureData, "facture_source.pdf", "EN16931")
    if err != nil {
        log.Fatal(err)
    }

    os.WriteFile("facture_facturx.pdf", pdfBytes, 0644)
}
```

## Helpers disponibles (package helpers)

### Montant(value)

Convertit une valeur en string formaté pour les montants monétaires.

```go
helpers.Montant(1234.5)      // "1234.50"
helpers.Montant("1234.56")   // "1234.56"
helpers.Montant(nil)         // "0.00"
```

### MontantTotal(ht, tva, ttc, aPayer)

Crée un objet MontantTotal complet.

```go
total := helpers.MontantTotal(1000.00, 200.00, 1200.00, 1200.00)

// Avec options
total := helpers.MontantTotalWithOptions(
    1000.00, 200.00, 1200.00, 1200.00,
    &helpers.MontantTotalOptions{
        RemiseTtc:   50.00,
        MotifRemise: "Fidélité",
        Acompte:     100.00,
    },
)
```

### LigneDePoste(numero, denomination, quantite, montantUnitaireHt, montantTotalLigneHt)

Crée une ligne de facturation.

```go
ligne := helpers.LigneDePosteWithOptions(
    1,
    "Prestation de conseil",
    5,
    200.00,
    1000.00,  // montantTotalLigneHt requis
    &helpers.LigneDePosteOptions{
        TauxTva:     "TVA20",      // Ou TauxTvaManuel: "20.00"
        CategorieTva: "S",         // S, Z, E, AE, K
        Unite:       "HEURE",      // FORFAIT, PIECE, HEURE, JOUR...
        Reference:   "REF-001",
    },
)
```

### LigneDeTva(montantBaseHt, montantTva, options)

Crée une ligne de ventilation TVA.

```go
tva := helpers.LigneDeTva(1000.00, 200.00, &helpers.LigneDeTvaOptions{
    Taux:      "TVA20",       // Ou TauxManuel: "20.00"
    Categorie: "S",           // S, Z, E, AE, K
})
```

### AdressePostale(ligne1, codePostal, ville, options)

Crée une adresse postale structurée.

```go
adresse := helpers.AdressePostale(
    "123 Rue de la République",
    "75001",
    "Paris",
    &helpers.AdressePostaleOptions{
        Pays:   "FR",          // Défaut: "FR"
        Ligne2: "Bâtiment A",  // Optionnel
    },
)
```

### Fournisseur(nom, siret, adresseLigne1, codePostal, ville, options)

Crée un fournisseur complet avec calcul automatique du SIREN et TVA intra.

```go
f := helpers.Fournisseur(
    "Ma Société SAS",
    "12345678901234",
    "123 Rue Example",
    "75001",
    "Paris",
    &helpers.FournisseurOptions{
        IBAN: "FR7630006000011234567890189",
    },
)
// SIREN et TVA intracommunautaire calculés automatiquement
```

### Destinataire(nom, siret, adresseLigne1, codePostal, ville, options)

Crée un destinataire (client) avec calcul automatique du SIREN.

```go
d := helpers.Destinataire(
    "Client SARL",
    "98765432109876",
    "456 Avenue Test",
    "69001",
    "Lyon",
    nil,
)
```

## Mode Zero-Trust (Chorus Pro / AFNOR)

Pour passer vos propres credentials sans stockage côté serveur :

```go
chorusCreds := &helpers.ChorusProCredentials{
    PisteClientID:     "votre_client_id",
    PisteClientSecret: "votre_client_secret",
    ChorusProLogin:    "votre_login",
    ChorusProPassword: "votre_password",
    Sandbox:           true,
}

afnorCreds := &helpers.AFNORCredentials{
    FlowServiceURL: "https://api.pdp.fr/flow/v1",
    TokenURL:       "https://auth.pdp.fr/oauth/token",
    ClientID:       "votre_client_id",
    ClientSecret:   "votre_client_secret",
}

client := helpers.NewClientWithCredentials(
    "votre_email@example.com",
    "votre_mot_de_passe",
    chorusCreds,
    afnorCreds,
)
```

## Ressources

- **Documentation API** : https://factpulse.fr/api/facturation/documentation
- **Support** : contact@factpulse.fr

## Licence

MIT License - Copyright (c) 2025 FactPulse

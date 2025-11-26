# FactPulse SDK Go

Client Go officiel pour l'API FactPulse - Facturation électronique française.

## 🎯 Fonctionnalités

- **Factur-X** : Génération et validation de factures électroniques (profils MINIMUM, BASIC, EN16931, EXTENDED)
- **Chorus Pro** : Intégration avec la plateforme de facturation publique française
- **AFNOR PDP/PA** : Soumission de flux conformes à la norme XP Z12-013
- **Signature électronique** : Signature PDF (PAdES-B-B, PAdES-B-T, PAdES-B-LT)
- **Client simplifié** : Authentification JWT et polling intégrés via `helpers`
- **Go 1.18+** : Support des generics et modules modernes

## 🚀 Installation

```bash
go get github.com/factpulse/sdk-go
```

## 📖 Démarrage rapide

### Méthode recommandée : Client simplifié avec helpers

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
    // Créer le client (authentification automatique)
    client := helpers.NewClient(helpers.ClientConfig{
        Email:    "votre_email@example.com",
        Password: "votre_mot_de_passe",
    })

    // Données de la facture
    factureData := map[string]interface{}{
        "numero_facture": "FAC-2025-001",
        "date_facture":   "2025-01-15",
        "fournisseur": map[string]interface{}{
            "nom":   "Mon Entreprise SAS",
            "siret": "12345678901234",
            "adresse_postale": map[string]string{
                "ligne_un":      "123 Rue Example",
                "code_postal":   "75001",
                "nom_ville":     "Paris",
                "pays_code_iso": "FR",
            },
        },
        "destinataire": map[string]interface{}{
            "nom":   "Client SARL",
            "siret": "98765432109876",
            "adresse_postale": map[string]string{
                "ligne_un":      "456 Avenue Test",
                "code_postal":   "69001",
                "nom_ville":     "Lyon",
                "pays_code_iso": "FR",
            },
        },
        "montant_total": map[string]string{
            "montant_ht_total":  "1000.00",
            "montant_tva":       "200.00",
            "montant_ttc_total": "1200.00",
            "montant_a_payer":   "1200.00",
        },
        "lignes_de_poste": []map[string]interface{}{
            {
                "numero":            1,
                "denomination":      "Prestation de conseil",
                "quantite":          "10.00",
                "unite":             "PIECE",
                "montant_unitaire_ht": "100.00",
            },
        },
    }

    // Lire le PDF source
    pdfSource, err := os.ReadFile("facture_source.pdf")
    if err != nil {
        log.Fatal(err)
    }

    // Générer le PDF Factur-X (polling automatique)
    ctx := context.Background()
    pdfBytes, err := client.GenererFacturx(
        ctx,
        factureData,
        pdfSource,
        "EN16931", // profil
        "pdf",     // format
        true,      // sync (attend le résultat)
        nil,       // timeout (utilise la valeur par défaut)
    )
    if err != nil {
        log.Fatal(err)
    }

    // Sauvegarder
    os.WriteFile("facture_facturx.pdf", pdfBytes, 0644)
}
```

### Méthode alternative : SDK brut

Pour un contrôle total, utilisez le SDK généré directement :

```go
package main

import (
    "bytes"
    "context"
    "encoding/json"
    "net/http"
    "os"

    factpulse "github.com/factpulse/sdk-go"
)

func main() {
    // 1. Obtenir le token JWT
    credentials := map[string]string{
        "username": "votre_email@example.com",
        "password": "votre_mot_de_passe",
    }
    jsonData, _ := json.Marshal(credentials)
    resp, _ := http.Post(
        "https://factpulse.fr/api/token/",
        "application/json",
        bytes.NewBuffer(jsonData),
    )
    defer resp.Body.Close()

    var result map[string]string
    json.NewDecoder(resp.Body).Decode(&result)
    token := result["access"]

    // 2. Configurer le client
    cfg := factpulse.NewConfiguration()
    cfg.Servers = factpulse.ServerConfigurations{
        {URL: "https://factpulse.fr/api/facturation"},
    }
    cfg.AddDefaultHeader("Authorization", "Bearer "+token)

    client := factpulse.NewAPIClient(cfg)
    ctx := context.Background()

    // 3. Appeler l'API
    pdfFile, _ := os.Open("facture_source.pdf")
    defer pdfFile.Close()

    response, _, _ := client.TraitementFactureAPI.
        GenererFactureApiV1TraitementGenererFacturePost(ctx).
        DonneesFacture(string(jsonData)).
        Profil("EN16931").
        FormatSortie("pdf").
        SourcePdf(pdfFile).
        Execute()

    // 4. Polling manuel pour récupérer le résultat
    taskID := response.GetIdTache()
    // ... (implémenter le polling)
}
```

## 🔧 Avantages des helpers

| Fonctionnalité | SDK brut | helpers |
|----------------|----------|---------|
| Authentification | Manuelle | Automatique |
| Refresh token | Manuel | Automatique |
| Polling tâches async | Manuel | Automatique (backoff) |
| Retry sur 401 | Manuel | Automatique |

## 🔑 Options d'authentification

### Client UID (multi-clients)

Si vous gérez plusieurs clients :

```go
client := helpers.NewClient(helpers.ClientConfig{
    Email:     "votre_email@example.com",
    Password:  "votre_mot_de_passe",
    ClientUID: "identifiant_client",  // UID du client cible
})
```

### Configuration avancée

```go
client := helpers.NewClient(helpers.ClientConfig{
    Email:           "votre_email@example.com",
    Password:        "votre_mot_de_passe",
    APIURL:          "https://factpulse.fr",  // URL personnalisée
    PollingInterval: 2000,  // Intervalle de polling initial (ms)
    PollingTimeout:  120000,  // Timeout de polling (ms)
    MaxRetries:      2,  // Tentatives en cas de 401
})
```

## 💡 Formats de montants acceptés

L'API accepte plusieurs formats pour les montants :

```go
// String (recommandé pour la précision)
montant := "1234.56"

// Float
montant := 1234.56

// Integer
montant := 1234

// Helper de formatage
montantFormate := helpers.FormatMontant(1234.5)  // "1234.50"
```

## 📚 Ressources

- **Documentation API** : https://factpulse.fr/api/facturation/documentation
- **Code source** : https://github.com/factpulse/sdk-go
- **Issues** : https://github.com/factpulse/sdk-go/issues
- **Support** : contact@factpulse.fr

## 📄 Licence

MIT License - Copyright (c) 2025 FactPulse

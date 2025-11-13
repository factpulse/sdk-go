# FactPulse SDK Go

Client Go officiel pour l'API FactPulse - Facturation électronique française.

## 🎯 Fonctionnalités

- **Factur-X** : Génération et validation de factures électroniques (profils MINIMUM, BASIC, EN16931, EXTENDED)
- **Chorus Pro** : Intégration avec la plateforme de facturation publique française
- **AFNOR PDP/PA** : Soumission de flux conformes à la norme XP Z12-013
- **Signature électronique** : Signature PDF (PAdES-B-B, PAdES-B-T, PAdES-B-LT)
- **Traitement asynchrone** : Support Celery pour opérations longues
- **Go 1.18+** : Support des generics et modules modernes

## 🚀 Installation

```bash
go get github.com/factpulse/sdk-go
```

## 📖 Démarrage rapide

### 1. Authentification

```go
package main

import (
    "context"
    "github.com/factpulse/sdk-go"
)

func main() {
    // Configuration du client
    cfg := factpulse.NewConfiguration()
    cfg.Host = "https://factpulse.fr/api/facturation"
    cfg.AddDefaultHeader("Authorization", "Bearer votre_token_jwt")

    client := factpulse.NewAPIClient(cfg)
    ctx := context.Background()
}
```

### 2. Générer une facture Factur-X

```go
import (
    "encoding/json"
    "os"
)

// Données de la facture
factureData := map[string]interface{}{
    "numero_facture": "FAC-2025-001",
    "date_facture": "2025-01-15",
    "montant_total_ht": "1000.00",
    "montant_total_ttc": "1200.00",
    "fournisseur": map[string]interface{}{
        "nom": "Mon Entreprise SAS",
        "siret": "12345678901234",
        "adresse_postale": map[string]string{
            "ligne_un": "123 Rue Example",
            "code_postal": "75001",
            "nom_ville": "Paris",
            "pays_code_iso": "FR",
        },
    },
    "destinataire": map[string]interface{}{
        "nom": "Client SARL",
        "siret": "98765432109876",
        "adresse_postale": map[string]string{
            "ligne_un": "456 Avenue Test",
            "code_postal": "69001",
            "nom_ville": "Lyon",
            "pays_code_iso": "FR",
        },
    },
    "lignes_de_poste": []map[string]interface{}{
        {
            "numero": 1,
            "denomination": "Prestation de conseil",
            "quantite": "10.00",
            "montant_unitaire_ht": "100.00",
            "montant_ligne_ht": "1000.00",
        },
    },
}

jsonData, _ := json.Marshal(factureData)

// Générer le PDF Factur-X
pdfFile, _, err := client.TraitementFactureApi.GenererFactureApiV1TraitementGenererFacturePost(ctx).
    DonneesFacture(string(jsonData)).
    Profil("EN16931").
    FormatSortie("pdf").
    Execute()

if err != nil {
    panic(err)
}

// Sauvegarder
os.WriteFile("facture.pdf", pdfFile, 0644)
```

### 3. Soumettre une facture complète (Chorus Pro / AFNOR PDP)

```go
requestBody := map[string]interface{}{
    "facture": factureData,
    "destination": map[string]interface{}{
        "type": "chorus_pro",
        "credentials": map[string]string{
            "login": "votre_login_chorus",
            "password": "votre_password_chorus",
        },
    },
}

response, _, err := client.TraitementFactureApi.SoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePost(ctx).
    Body(requestBody).
    Execute()

if err != nil {
    panic(err)
}

println("Facture soumise :", response.IdFactureChorus)
```

## 🔑 Obtention du token JWT

### Via l'API

```go
import (
    "bytes"
    "encoding/json"
    "net/http"
)

credentials := map[string]string{
    "username": "votre_email@example.com",
    "password": "votre_mot_de_passe",
}

jsonData, _ := json.Marshal(credentials)
resp, err := http.Post(
    "https://factpulse.fr/api/token/",
    "application/json",
    bytes.NewBuffer(jsonData),
)

if err != nil {
    panic(err)
}
defer resp.Body.Close()

var result map[string]string
json.NewDecoder(resp.Body).Decode(&result)
token := result["access"]
```

**Accès aux credentials d'un client spécifique :**

Si vous gérez plusieurs clients et souhaitez accéder aux credentials (Chorus Pro, AFNOR PDP) d'un client particulier, ajoutez le champ `client_uid` :

```go
credentials := map[string]string{
    "username": "votre_email@example.com",
    "password": "votre_mot_de_passe",
    "client_uid": "identifiant_client",  // UID du client cible
}
```

### Via le Dashboard

1. Connectez-vous sur https://factpulse.fr/api/dashboard/
2. Générez un token API
3. Copiez et utilisez le token dans votre configuration

## 📚 Ressources

- **Documentation API** : https://factpulse.fr/api/facturation/documentation
- **Code source** : https://github.com/factpulse/sdk-go
- **Issues** : https://github.com/factpulse/sdk-go/issues
- **Support** : contact@factpulse.fr

## 📄 Licence

MIT License - Copyright (c) 2025 FactPulse

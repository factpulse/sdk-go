/*
API REST FactPulse

 API REST pour la facturation électronique en France : Factur-X, AFNOR PDP/PA, signatures électroniques.  ## 🎯 Fonctionnalités principales  ### 📄 Génération de factures Factur-X - **Formats** : XML seul ou PDF/A-3 avec XML embarqué - **Profils** : MINIMUM, BASIC, EN16931, EXTENDED - **Normes** : EN 16931 (directive UE 2014/55), ISO 19005-3 (PDF/A-3), CII (UN/CEFACT) - **🆕 Format simplifié** : Génération à partir de SIRET + auto-enrichissement (API Chorus Pro + Recherche Entreprises)  ### ✅ Validation et conformité - **Validation XML** : Schematron (45 à 210+ règles selon profil) - **Validation PDF** : PDF/A-3, métadonnées XMP Factur-X, signatures électroniques - **VeraPDF** : Validation stricte PDF/A (146+ règles ISO 19005-3) - **Traitement asynchrone** : Support Celery pour validations lourdes (VeraPDF)  ### 📡 Intégration AFNOR PDP/PA (XP Z12-013) - **Soumission de flux** : Envoi de factures vers Plateformes de Dématérialisation Partenaires - **Recherche de flux** : Consultation des factures soumises - **Téléchargement** : Récupération des PDF/A-3 avec XML - **Directory Service** : Recherche d'entreprises (SIREN/SIRET) - **Multi-client** : Support de plusieurs configs PDP par utilisateur (stored credentials ou zero-storage)  ### ✍️ Signature électronique PDF - **Standards** : PAdES-B-B, PAdES-B-T (horodatage RFC 3161), PAdES-B-LT (archivage long terme) - **Niveaux eIDAS** : SES (auto-signé), AdES (CA commerciale), QES (PSCO) - **Validation** : Vérification intégrité cryptographique et certificats - **Génération de certificats** : Certificats X.509 auto-signés pour tests  ### 🔄 Traitement asynchrone - **Celery** : Génération, validation et signature asynchrones - **Polling** : Suivi d'état via `/taches/{id_tache}/statut` - **Pas de timeout** : Idéal pour gros fichiers ou validations lourdes  ## 🔒 Authentification  Toutes les requêtes nécessitent un **token JWT** dans le header Authorization : ``` Authorization: Bearer YOUR_JWT_TOKEN ```  ### Comment obtenir un token JWT ?  #### 🔑 Méthode 1 : API `/api/token/` (Recommandée)  **URL :** `https://www.factpulse.fr/api/token/`  Cette méthode est **recommandée** pour l'intégration dans vos applications et workflows CI/CD.  **Prérequis :** Avoir défini un mot de passe sur votre compte  **Pour les utilisateurs inscrits via email/password :** - Vous avez déjà un mot de passe, utilisez-le directement  **Pour les utilisateurs inscrits via OAuth (Google/GitHub) :** - Vous devez d'abord définir un mot de passe sur : https://www.factpulse.fr/accounts/password/set/ - Une fois le mot de passe créé, vous pourrez utiliser l'API  **Exemple de requête :** ```bash curl -X POST https://www.factpulse.fr/api/token/ \\   -H \"Content-Type: application/json\" \\   -d '{     \"username\": \"votre_email@example.com\",     \"password\": \"votre_mot_de_passe\"   }' ```  **Paramètre optionnel `client_uid` :**  Pour sélectionner les credentials d'un client spécifique (PA/PDP, Chorus Pro, certificats de signature), ajoutez `client_uid` :  ```bash curl -X POST https://www.factpulse.fr/api/token/ \\   -H \"Content-Type: application/json\" \\   -d '{     \"username\": \"votre_email@example.com\",     \"password\": \"votre_mot_de_passe\",     \"client_uid\": \"550e8400-e29b-41d4-a716-446655440000\"   }' ```  Le `client_uid` sera inclus dans le JWT et permettra à l'API d'utiliser automatiquement : - Les credentials AFNOR/PDP configurés pour ce client - Les credentials Chorus Pro configurés pour ce client - Les certificats de signature électronique configurés pour ce client  **Réponse :** ```json {   \"access\": \"eyJ0eXAiOiJKV1QiLCJhbGc...\",  // Token d'accès (validité: 30 min)   \"refresh\": \"eyJ0eXAiOiJKV1QiLCJhbGc...\"  // Token de rafraîchissement (validité: 7 jours) } ```  **Avantages :** - ✅ Automatisation complète (CI/CD, scripts) - ✅ Gestion programmatique des tokens - ✅ Support du refresh token pour renouveler automatiquement l'accès - ✅ Intégration facile dans n'importe quel langage/outil  #### 🖥️ Méthode 2 : Génération via Dashboard (Alternative)  **URL :** https://www.factpulse.fr/dashboard/  Cette méthode convient pour des tests rapides ou une utilisation occasionnelle via l'interface graphique.  **Fonctionnement :** - Connectez-vous au dashboard - Utilisez les boutons \"Generate Test Token\" ou \"Generate Production Token\" - Fonctionne pour **tous** les utilisateurs (OAuth et email/password), sans nécessiter de mot de passe  **Types de tokens :** - **Token Test** : Validité 24h, quota 1000 appels/jour (gratuit) - **Token Production** : Validité 7 jours, quota selon votre forfait  **Avantages :** - ✅ Rapide pour tester l'API - ✅ Aucun mot de passe requis - ✅ Interface visuelle simple  **Inconvénients :** - ❌ Nécessite une action manuelle - ❌ Pas de refresh token - ❌ Moins adapté pour l'automatisation  ### 📚 Documentation complète  Pour plus d'informations sur l'authentification et l'utilisation de l'API : https://www.factpulse.fr/documentation-api/     

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package factpulse

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"os"
)


// TraitementFactureAPIService TraitementFactureAPI service
type TraitementFactureAPIService service

type ApiGenererCertificatTestApiV1TraitementGenererCertificatTestPostRequest struct {
	ctx context.Context
	ApiService *TraitementFactureAPIService
	generateCertificateRequest *GenerateCertificateRequest
}

func (r ApiGenererCertificatTestApiV1TraitementGenererCertificatTestPostRequest) GenerateCertificateRequest(generateCertificateRequest GenerateCertificateRequest) ApiGenererCertificatTestApiV1TraitementGenererCertificatTestPostRequest {
	r.generateCertificateRequest = &generateCertificateRequest
	return r
}

func (r ApiGenererCertificatTestApiV1TraitementGenererCertificatTestPostRequest) Execute() (*GenerateCertificateResponse, *http.Response, error) {
	return r.ApiService.GenererCertificatTestApiV1TraitementGenererCertificatTestPostExecute(r)
}

/*
GenererCertificatTestApiV1TraitementGenererCertificatTestPost Générer un certificat X.509 auto-signé de test

Génère un certificat X.509 auto-signé pour les tests de signature électronique PDF.

    **⚠️ ATTENTION : Certificat de TEST uniquement !**

    Ce certificat est :
    - ✅ Adapté pour tests et développement
    - ✅ Compatible signature PDF (PAdES)
    - ✅ Conforme eIDAS niveau **SES** (Simple Electronic Signature)
    - ❌ **JAMAIS utilisable en production**
    - ❌ **Non reconnu** par les navigateurs et lecteurs PDF
    - ❌ **Aucune valeur juridique**

    ## Niveaux eIDAS

    - **SES** (Simple) : Certificat auto-signé ← Généré par cet endpoint
    - **AdES** (Advanced) : Certificat CA commerciale (Let's Encrypt, etc.)
    - **QES** (Qualified) : Certificat qualifié PSCO (CertEurope, Universign, etc.)

    ## Utilisation

    Une fois généré, le certificat peut être :

    1. **Enregistré dans Django** (recommandé) :
       - Django Admin > Certificats de signature
       - Upload `certificat_pem` et `cle_privee_pem`

    2. **Utilisé directement** :
       - Signer un PDF avec `/signer-pdf`
       - Le certificat sera automatiquement utilisé

    ## Exemple d'appel

    ```bash
    curl -X POST "https://www.factpulse.fr/api/facturation/generer-certificat-test" \
      -H "Authorization: Bearer eyJ0eXAi..." \
      -H "Content-Type: application/json" \
      -d '{
        "cn": "Test Client XYZ",
        "organisation": "Client XYZ SARL",
        "email": "contact@xyz.fr",
        "duree_jours": 365
      }'
    ```

    ## Cas d'usage

    - Tests de signature PDF en développement
    - POC de signature électronique
    - Formation et démos
    - Tests d'intégration automatisés

    ## Conformité technique

    Certificat généré avec :
    - Clé RSA 2048 ou 4096 bits
    - Algorithme SHA-256
    - Extensions Key Usage : `digitalSignature`, `contentCommitment` (non-repudiation)
    - Extensions Extended Key Usage : `codeSigning`, `emailProtection`
    - Validité : 1 jour à 10 ans (configurable)
    - Format : PEM (certificat et clé)
    - Optionnel : PKCS#12 (.p12)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGenererCertificatTestApiV1TraitementGenererCertificatTestPostRequest
*/
func (a *TraitementFactureAPIService) GenererCertificatTestApiV1TraitementGenererCertificatTestPost(ctx context.Context) ApiGenererCertificatTestApiV1TraitementGenererCertificatTestPostRequest {
	return ApiGenererCertificatTestApiV1TraitementGenererCertificatTestPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GenerateCertificateResponse
func (a *TraitementFactureAPIService) GenererCertificatTestApiV1TraitementGenererCertificatTestPostExecute(r ApiGenererCertificatTestApiV1TraitementGenererCertificatTestPostRequest) (*GenerateCertificateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GenerateCertificateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TraitementFactureAPIService.GenererCertificatTestApiV1TraitementGenererCertificatTestPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/traitement/generer-certificat-test"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.generateCertificateRequest == nil {
		return localVarReturnValue, nil, reportError("generateCertificateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.generateCertificateRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGenererFactureApiV1TraitementGenererFacturePostRequest struct {
	ctx context.Context
	ApiService *TraitementFactureAPIService
	donneesFacture *string
	profil *ProfilAPI
	formatSortie *FormatSortie
	autoEnrichir *bool
	sourcePdf *os.File
}

// Données de la facture au format JSON.              Deux formats acceptés :             1. **Format classique** : Structure complète FactureFacturX (tous les champs)             2. **Format simplifié** (🆕 P0.1) : Structure minimale avec auto-enrichissement              Le format est détecté automatiquement !             
func (r ApiGenererFactureApiV1TraitementGenererFacturePostRequest) DonneesFacture(donneesFacture string) ApiGenererFactureApiV1TraitementGenererFacturePostRequest {
	r.donneesFacture = &donneesFacture
	return r
}

// Profil Factur-X : MINIMUM, BASIC, EN16931 ou EXTENDED.
func (r ApiGenererFactureApiV1TraitementGenererFacturePostRequest) Profil(profil ProfilAPI) ApiGenererFactureApiV1TraitementGenererFacturePostRequest {
	r.profil = &profil
	return r
}

// Format de sortie : &#39;xml&#39; (XML seul) ou &#39;pdf&#39; (PDF Factur-X avec XML embarqué).
func (r ApiGenererFactureApiV1TraitementGenererFacturePostRequest) FormatSortie(formatSortie FormatSortie) ApiGenererFactureApiV1TraitementGenererFacturePostRequest {
	r.formatSortie = &formatSortie
	return r
}

// 🆕 Activer l&#39;auto-enrichissement depuis SIRET/SIREN (format simplifié uniquement)
func (r ApiGenererFactureApiV1TraitementGenererFacturePostRequest) AutoEnrichir(autoEnrichir bool) ApiGenererFactureApiV1TraitementGenererFacturePostRequest {
	r.autoEnrichir = &autoEnrichir
	return r
}

func (r ApiGenererFactureApiV1TraitementGenererFacturePostRequest) SourcePdf(sourcePdf *os.File) ApiGenererFactureApiV1TraitementGenererFacturePostRequest {
	r.sourcePdf = sourcePdf
	return r
}

func (r ApiGenererFactureApiV1TraitementGenererFacturePostRequest) Execute() (*ReponseTache, *http.Response, error) {
	return r.ApiService.GenererFactureApiV1TraitementGenererFacturePostExecute(r)
}

/*
GenererFactureApiV1TraitementGenererFacturePost Générer une facture Factur-X

Génère une facture électronique au format Factur-X conforme aux normes européennes.

## Normes appliquées

- **Factur-X** (France) : Norme FNFE-MPE (Forum National de la Facture Électronique)
- **ZUGFeRD** (Allemagne) : Format allemand compatible Factur-X
- **EN 16931** : Norme sémantique européenne pour la facturation électronique
- **ISO 19005-3** (PDF/A-3) : Archivage électronique à long terme
- **Cross Industry Invoice (CII)** : Syntaxe XML UN/CEFACT

## 🆕 Nouveau : Format simplifié avec auto-enrichissement (P0.1)

Vous pouvez désormais créer une facture en fournissant uniquement :
- Un numéro de facture
- Un SIRET émetteur + **IBAN** (obligatoire)
- Un SIRET destinataire
- Les lignes de facture (description, quantité, prix HT)

**Exemple format simplifié** :
```json
{
  "numero": "FACT-2025-001",
  "emetteur": {
    "siret": "92019522900017",
    "iban": "FR7630001007941234567890185"
  },
  "destinataire": {"siret": "35600000000048"},
  "lignes": [
    {"description": "Prestation", "quantite": 10, "prix_ht": 100.00, "tva": 20.0}
  ]
}
```

**⚠️ Champs obligatoires (format simplifié)** :
- `numero` : Numéro de facture unique
- `emetteur.siret` : SIRET de l'émetteur (14 chiffres)
- `emetteur.iban` : IBAN du compte bancaire (pas d'API publique pour le récupérer)
- `destinataire.siret` : SIRET du destinataire
- `lignes[]` : Au moins une ligne de facture

**Ce qui se passe automatiquement avec `auto_enrichir=True`** :
- ✅ Enrichissement des noms depuis API Chorus Pro
- ✅ Enrichissement des adresses depuis API Recherche Entreprises (gratuite, publique)
- ✅ Calcul automatique de la TVA intracommunautaire (FR + clé + SIREN)
- ✅ Récupération de l'ID Chorus Pro pour la facturation électronique
- ✅ Calcul des totaux HT/TVA/TTC
- ✅ Génération des dates (aujourd'hui + échéance 30j)
- ✅ Gestion multi-taux de TVA

**Identifiants supportés** :
- SIRET (14 chiffres) : Établissement précis ⭐ Recommandé
- SIREN (9 chiffres) : Entreprise (sélection auto du siège)
- Types spéciaux : UE_HORS_FRANCE, RIDET, TAHITI, etc.

## Contrôles effectués lors de la génération

### 1. Validation des données (Pydantic)
- Types de données (montants en Decimal, dates ISO 8601)
- Formats (SIRET 14 chiffres, SIREN 9 chiffres, IBAN)
- Champs obligatoires selon le profil
- Cohérence des montants (HT + TVA = TTC)

### 2. Génération XML conforme CII
- Sérialisation selon schéma XSD Cross Industry Invoice
- Namespaces UN/CEFACT corrects
- Structure hiérarchique respectée
- Encodage UTF-8 sans BOM

### 3. Validation Schematron
- Règles métier du profil sélectionné (MINIMUM, BASIC, EN16931, EXTENDED)
- Cardinalité des éléments (obligatoire, optionnel, répétable)
- Règles de calcul (totaux, TVA, remises)
- Conformité européenne EN 16931

### 4. Conversion PDF/A-3 (si format_sortie='pdf')
- Conversion du PDF source en PDF/A-3 via Ghostscript
- Embarquement du XML Factur-X dans le PDF
- Métadonnées XMP conformes
- Profil ICC sRGB pour les couleurs
- Suppression des éléments interdits (JavaScript, formulaires)

## Fonctionnement

1. **Soumission** : La facture est mise en file d'attente Celery pour traitement asynchrone
2. **Retour immédiat** : Vous recevez un `id_tache` (HTTP 202 Accepted)
3. **Suivi** : Utilisez l'endpoint `/taches/{id_tache}/statut` pour suivre l'avancement

## Formats de sortie

- **xml** : Génère uniquement le XML Factur-X (recommandé pour les tests)
- **pdf** : Génère un PDF/A-3 avec XML embarqué (nécessite `source_pdf`)

## Profils Factur-X

- **MINIMUM** : Données minimales (facture simplifiée)
- **BASIC** : Informations de base (PME)
- **EN16931** : Standard européen (recommandé, conforme directive 2014/55/UE)
- **EXTENDED** : Toutes les données disponibles (grands comptes)

## Ce que vous obtenez

Après traitement réussi (statut `completed`) :
- **XML seul** : Fichier XML encodé base64 conforme Factur-X
- **PDF/A-3** : PDF avec XML embarqué, prêt pour envoi/archivage
- **Métadonnées** : Profil, version Factur-X, taille fichier
- **Validation** : Confirmation de conformité Schematron

## Validation

Les données sont validées automatiquement selon le format détecté.
En cas d'erreur, un statut 422 est retourné avec les détails des champs invalides.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGenererFactureApiV1TraitementGenererFacturePostRequest
*/
func (a *TraitementFactureAPIService) GenererFactureApiV1TraitementGenererFacturePost(ctx context.Context) ApiGenererFactureApiV1TraitementGenererFacturePostRequest {
	return ApiGenererFactureApiV1TraitementGenererFacturePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReponseTache
func (a *TraitementFactureAPIService) GenererFactureApiV1TraitementGenererFacturePostExecute(r ApiGenererFactureApiV1TraitementGenererFacturePostRequest) (*ReponseTache, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReponseTache
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TraitementFactureAPIService.GenererFactureApiV1TraitementGenererFacturePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/traitement/generer-facture"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.donneesFacture == nil {
		return localVarReturnValue, nil, reportError("donneesFacture is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "donnees_facture", r.donneesFacture, "", "")
	if r.profil != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "profil", r.profil, "", "")
	}
	if r.formatSortie != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "format_sortie", r.formatSortie, "", "")
	}
	if r.autoEnrichir != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "auto_enrichir", r.autoEnrichir, "", "")
	}
	var sourcePdfLocalVarFormFileName string
	var sourcePdfLocalVarFileName     string
	var sourcePdfLocalVarFileBytes    []byte

	sourcePdfLocalVarFormFileName = "source_pdf"
	sourcePdfLocalVarFile := r.sourcePdf

	if sourcePdfLocalVarFile != nil {
		fbs, _ := io.ReadAll(sourcePdfLocalVarFile)

		sourcePdfLocalVarFileBytes = fbs
		sourcePdfLocalVarFileName = sourcePdfLocalVarFile.Name()
		sourcePdfLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: sourcePdfLocalVarFileBytes, fileName: sourcePdfLocalVarFileName, formFileName: sourcePdfLocalVarFormFileName})
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiObtenirStatutTacheApiV1TraitementTachesIdTacheStatutGetRequest struct {
	ctx context.Context
	ApiService *TraitementFactureAPIService
	idTache string
}

func (r ApiObtenirStatutTacheApiV1TraitementTachesIdTacheStatutGetRequest) Execute() (*StatutTache, *http.Response, error) {
	return r.ApiService.ObtenirStatutTacheApiV1TraitementTachesIdTacheStatutGetExecute(r)
}

/*
ObtenirStatutTacheApiV1TraitementTachesIdTacheStatutGet Obtenir le statut d'une tâche de génération

Récupère l'état d'avancement d'une tâche de génération de facture.

## États possibles

Le champ `statut` utilise l'enum `StatutCelery` avec les valeurs :
- **PENDING, STARTED, SUCCESS, FAILURE, RETRY**

Voir la documentation du schéma `StatutCelery` pour les détails.

## Résultat métier

Quand `statut="SUCCESS"`, le champ `resultat` contient :
- `statut` : "SUCCES" ou "ERREUR" (résultat métier)
- `chemin_fichier` : Chemin du fichier généré (si succès)
- `message_erreur` : Détails de l'erreur (si échec métier)

## Usage

Appelez cet endpoint en boucle (polling) toutes les 2-3 secondes jusqu'à ce que
`statut` soit `SUCCESS` ou `FAILURE`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param idTache
 @return ApiObtenirStatutTacheApiV1TraitementTachesIdTacheStatutGetRequest
*/
func (a *TraitementFactureAPIService) ObtenirStatutTacheApiV1TraitementTachesIdTacheStatutGet(ctx context.Context, idTache string) ApiObtenirStatutTacheApiV1TraitementTachesIdTacheStatutGetRequest {
	return ApiObtenirStatutTacheApiV1TraitementTachesIdTacheStatutGetRequest{
		ApiService: a,
		ctx: ctx,
		idTache: idTache,
	}
}

// Execute executes the request
//  @return StatutTache
func (a *TraitementFactureAPIService) ObtenirStatutTacheApiV1TraitementTachesIdTacheStatutGetExecute(r ApiObtenirStatutTacheApiV1TraitementTachesIdTacheStatutGetRequest) (*StatutTache, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StatutTache
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TraitementFactureAPIService.ObtenirStatutTacheApiV1TraitementTachesIdTacheStatutGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/traitement/taches/{id_tache}/statut"
	localVarPath = strings.Replace(localVarPath, "{"+"id_tache"+"}", url.PathEscape(parameterValueToString(r.idTache, "idTache")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSignerPdfApiV1TraitementSignerPdfPostRequest struct {
	ctx context.Context
	ApiService *TraitementFactureAPIService
	fichierPdf *os.File
	raison *string
	localisation *string
	contact *string
	fieldName *string
	usePadesLt *bool
	useTimestamp *bool
}

// Fichier PDF à signer (sera traité puis retourné signé en base64)
func (r ApiSignerPdfApiV1TraitementSignerPdfPostRequest) FichierPdf(fichierPdf *os.File) ApiSignerPdfApiV1TraitementSignerPdfPostRequest {
	r.fichierPdf = fichierPdf
	return r
}

func (r ApiSignerPdfApiV1TraitementSignerPdfPostRequest) Raison(raison string) ApiSignerPdfApiV1TraitementSignerPdfPostRequest {
	r.raison = &raison
	return r
}

func (r ApiSignerPdfApiV1TraitementSignerPdfPostRequest) Localisation(localisation string) ApiSignerPdfApiV1TraitementSignerPdfPostRequest {
	r.localisation = &localisation
	return r
}

func (r ApiSignerPdfApiV1TraitementSignerPdfPostRequest) Contact(contact string) ApiSignerPdfApiV1TraitementSignerPdfPostRequest {
	r.contact = &contact
	return r
}

// Nom du champ de signature PDF
func (r ApiSignerPdfApiV1TraitementSignerPdfPostRequest) FieldName(fieldName string) ApiSignerPdfApiV1TraitementSignerPdfPostRequest {
	r.fieldName = &fieldName
	return r
}

// Activer PAdES-B-LT (archivage long terme avec données de validation embarquées). NÉCESSITE un certificat avec accès OCSP/CRL.
func (r ApiSignerPdfApiV1TraitementSignerPdfPostRequest) UsePadesLt(usePadesLt bool) ApiSignerPdfApiV1TraitementSignerPdfPostRequest {
	r.usePadesLt = &usePadesLt
	return r
}

// Activer l&#39;horodatage RFC 3161 avec FreeTSA (PAdES-B-T)
func (r ApiSignerPdfApiV1TraitementSignerPdfPostRequest) UseTimestamp(useTimestamp bool) ApiSignerPdfApiV1TraitementSignerPdfPostRequest {
	r.useTimestamp = &useTimestamp
	return r
}

func (r ApiSignerPdfApiV1TraitementSignerPdfPostRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.SignerPdfApiV1TraitementSignerPdfPostExecute(r)
}

/*
SignerPdfApiV1TraitementSignerPdfPost Signer un PDF avec le certificat du client (PAdES-B-LT)

Signe un PDF uploadé avec le certificat électronique configuré pour le client (via client_uid du JWT).

    **Standards supportés** : PAdES-B-B, PAdES-B-T (horodatage), PAdES-B-LT (archivage long terme).

    **Niveaux eIDAS** : SES (auto-signé), AdES (CA commerciale), QES (PSCO - hors scope).

    **Sécurité** : Double authentification X-Internal-Secret + JWT Bearer pour récupérer le certificat.

    **⚠️ Disclaimer légal** : Les signatures générées sont des cachets électroniques au sens
    du règlement eIDAS. Le niveau de validité juridique dépend du certificat utilisé (SES/AdES/QES).
    FactPulse ne fournit pas de certificats qualifiés QES - vous devez obtenir un certificat auprès
    d'un PSCO (Prestataire de Services de Confiance qualifié) pour une validité juridique maximale.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSignerPdfApiV1TraitementSignerPdfPostRequest
*/
func (a *TraitementFactureAPIService) SignerPdfApiV1TraitementSignerPdfPost(ctx context.Context) ApiSignerPdfApiV1TraitementSignerPdfPostRequest {
	return ApiSignerPdfApiV1TraitementSignerPdfPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return interface{}
func (a *TraitementFactureAPIService) SignerPdfApiV1TraitementSignerPdfPostExecute(r ApiSignerPdfApiV1TraitementSignerPdfPostRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TraitementFactureAPIService.SignerPdfApiV1TraitementSignerPdfPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/traitement/signer-pdf"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fichierPdf == nil {
		return localVarReturnValue, nil, reportError("fichierPdf is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	var fichierPdfLocalVarFormFileName string
	var fichierPdfLocalVarFileName     string
	var fichierPdfLocalVarFileBytes    []byte

	fichierPdfLocalVarFormFileName = "fichier_pdf"
	fichierPdfLocalVarFile := r.fichierPdf

	if fichierPdfLocalVarFile != nil {
		fbs, _ := io.ReadAll(fichierPdfLocalVarFile)

		fichierPdfLocalVarFileBytes = fbs
		fichierPdfLocalVarFileName = fichierPdfLocalVarFile.Name()
		fichierPdfLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fichierPdfLocalVarFileBytes, fileName: fichierPdfLocalVarFileName, formFileName: fichierPdfLocalVarFormFileName})
	}
	if r.raison != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "raison", r.raison, "", "")
	}
	if r.localisation != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "localisation", r.localisation, "", "")
	}
	if r.contact != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "contact", r.contact, "", "")
	}
	if r.fieldName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "field_name", r.fieldName, "", "")
	}
	if r.usePadesLt != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "use_pades_lt", r.usePadesLt, "", "")
	}
	if r.useTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "use_timestamp", r.useTimestamp, "", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSignerPdfAsyncApiV1TraitementSignerPdfAsyncPostRequest struct {
	ctx context.Context
	ApiService *TraitementFactureAPIService
	fichierPdf *os.File
	raison *string
	localisation *string
	contact *string
	fieldName *string
	usePadesLt *bool
	useTimestamp *bool
}

// Fichier PDF à signer (traité de manière asynchrone)
func (r ApiSignerPdfAsyncApiV1TraitementSignerPdfAsyncPostRequest) FichierPdf(fichierPdf *os.File) ApiSignerPdfAsyncApiV1TraitementSignerPdfAsyncPostRequest {
	r.fichierPdf = fichierPdf
	return r
}

func (r ApiSignerPdfAsyncApiV1TraitementSignerPdfAsyncPostRequest) Raison(raison string) ApiSignerPdfAsyncApiV1TraitementSignerPdfAsyncPostRequest {
	r.raison = &raison
	return r
}

func (r ApiSignerPdfAsyncApiV1TraitementSignerPdfAsyncPostRequest) Localisation(localisation string) ApiSignerPdfAsyncApiV1TraitementSignerPdfAsyncPostRequest {
	r.localisation = &localisation
	return r
}

func (r ApiSignerPdfAsyncApiV1TraitementSignerPdfAsyncPostRequest) Contact(contact string) ApiSignerPdfAsyncApiV1TraitementSignerPdfAsyncPostRequest {
	r.contact = &contact
	return r
}

// Nom du champ de signature PDF
func (r ApiSignerPdfAsyncApiV1TraitementSignerPdfAsyncPostRequest) FieldName(fieldName string) ApiSignerPdfAsyncApiV1TraitementSignerPdfAsyncPostRequest {
	r.fieldName = &fieldName
	return r
}

// Activer PAdES-B-LT (archivage long terme avec données de validation embarquées). NÉCESSITE un certificat avec accès OCSP/CRL.
func (r ApiSignerPdfAsyncApiV1TraitementSignerPdfAsyncPostRequest) UsePadesLt(usePadesLt bool) ApiSignerPdfAsyncApiV1TraitementSignerPdfAsyncPostRequest {
	r.usePadesLt = &usePadesLt
	return r
}

// Activer l&#39;horodatage RFC 3161 avec FreeTSA (PAdES-B-T)
func (r ApiSignerPdfAsyncApiV1TraitementSignerPdfAsyncPostRequest) UseTimestamp(useTimestamp bool) ApiSignerPdfAsyncApiV1TraitementSignerPdfAsyncPostRequest {
	r.useTimestamp = &useTimestamp
	return r
}

func (r ApiSignerPdfAsyncApiV1TraitementSignerPdfAsyncPostRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.SignerPdfAsyncApiV1TraitementSignerPdfAsyncPostExecute(r)
}

/*
SignerPdfAsyncApiV1TraitementSignerPdfAsyncPost Signer un PDF de manière asynchrone (Celery)

Signe un PDF uploadé de manière asynchrone via une tâche Celery.

    **Différence avec /signer-pdf** :
    - `/signer-pdf` : Signature synchrone (blocage jusqu'à la fin)
    - `/signer-pdf-async` : Signature asynchrone (retourne immédiatement un task_id)

    **Avantages de l'async** :
    - Pas de timeout pour les gros fichiers
    - Pas de blocage du worker FastAPI
    - Possibilité de suivre la progression via le task_id
    - Idéal pour les traitements par lot

    **Standards supportés** : PAdES-B-B, PAdES-B-T (horodatage), PAdES-B-LT (archivage long terme).

    **⚠️ Disclaimer légal** : Identique à /signer-pdf (voir documentation de cet endpoint).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSignerPdfAsyncApiV1TraitementSignerPdfAsyncPostRequest
*/
func (a *TraitementFactureAPIService) SignerPdfAsyncApiV1TraitementSignerPdfAsyncPost(ctx context.Context) ApiSignerPdfAsyncApiV1TraitementSignerPdfAsyncPostRequest {
	return ApiSignerPdfAsyncApiV1TraitementSignerPdfAsyncPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return interface{}
func (a *TraitementFactureAPIService) SignerPdfAsyncApiV1TraitementSignerPdfAsyncPostExecute(r ApiSignerPdfAsyncApiV1TraitementSignerPdfAsyncPostRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TraitementFactureAPIService.SignerPdfAsyncApiV1TraitementSignerPdfAsyncPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/traitement/signer-pdf-async"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fichierPdf == nil {
		return localVarReturnValue, nil, reportError("fichierPdf is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	var fichierPdfLocalVarFormFileName string
	var fichierPdfLocalVarFileName     string
	var fichierPdfLocalVarFileBytes    []byte

	fichierPdfLocalVarFormFileName = "fichier_pdf"
	fichierPdfLocalVarFile := r.fichierPdf

	if fichierPdfLocalVarFile != nil {
		fbs, _ := io.ReadAll(fichierPdfLocalVarFile)

		fichierPdfLocalVarFileBytes = fbs
		fichierPdfLocalVarFileName = fichierPdfLocalVarFile.Name()
		fichierPdfLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fichierPdfLocalVarFileBytes, fileName: fichierPdfLocalVarFileName, formFileName: fichierPdfLocalVarFormFileName})
	}
	if r.raison != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "raison", r.raison, "", "")
	}
	if r.localisation != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "localisation", r.localisation, "", "")
	}
	if r.contact != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "contact", r.contact, "", "")
	}
	if r.fieldName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "field_name", r.fieldName, "", "")
	}
	if r.usePadesLt != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "use_pades_lt", r.usePadesLt, "", "")
	}
	if r.useTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "use_timestamp", r.useTimestamp, "", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePostRequest struct {
	ctx context.Context
	ApiService *TraitementFactureAPIService
	soumettreFactureCompleteRequest *SoumettreFactureCompleteRequest
}

func (r ApiSoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePostRequest) SoumettreFactureCompleteRequest(soumettreFactureCompleteRequest SoumettreFactureCompleteRequest) ApiSoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePostRequest {
	r.soumettreFactureCompleteRequest = &soumettreFactureCompleteRequest
	return r
}

func (r ApiSoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePostRequest) Execute() (*SoumettreFactureCompleteResponse, *http.Response, error) {
	return r.ApiService.SoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePostExecute(r)
}

/*
SoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePost Soumettre une facture complète (génération + signature + soumission)

Endpoint unifié pour soumettre une facture complète vers différentes destinations.

    **Workflow automatisé :**
    1. **Auto-enrichissement** (optionnel) : récupère les données via APIs publiques et Chorus Pro/AFNOR
    2. **Génération PDF Factur-X** : crée un PDF/A-3 avec XML embarqué
    3. **Signature électronique** (optionnelle) : signe le PDF avec un certificat
    4. **Soumission** : envoie vers la destination choisie (Chorus Pro ou AFNOR PDP)

    **Destinations supportées :**
    - **Chorus Pro** : plateforme B2G française (factures vers secteur public)
    - **AFNOR PDP** : Plateformes de Dématérialisation Partenaires

    **Credentials de destination - 2 modes disponibles :**

    **Mode 1 - Récupération via JWT (recommandé) :**
    - Les credentials sont récupérés automatiquement via le `client_uid` du JWT
    - Ne pas fournir le champ `credentials` dans `destination`
    - Architecture 0-trust : aucun secret dans le payload
    - Exemple : `"destination": {"type": "chorus_pro"}`

    **Mode 2 - Credentials dans le payload :**
    - Fournir les credentials directement dans le payload
    - Utile pour tests ou intégrations tierces
    - Exemple : `"destination": {"type": "chorus_pro", "credentials": {...}}`


    **Signature électronique (optionnelle) - 2 modes disponibles :**

    **Mode 1 - Certificat stocké (recommandé) :**
    - Le certificat est récupéré automatiquement via le `client_uid` du JWT
    - Aucune clé à fournir dans le payload
    - Signature PAdES-B-LT avec horodatage (conforme eIDAS)
    - Exemple : `"signature": {"raison": "Conformité Factur-X"}`

    **Mode 2 - Clés dans le payload (pour tests) :**
    - Fournir `key_pem` et `cert_pem` directement
    - Format PEM accepté : brut ou base64
    - Utile pour tests ou cas spéciaux sans certificat stocké
    - Exemple : `"signature": {"key_pem": "-----BEGIN...", "cert_pem": "-----BEGIN..."}`

    Si `key_pem` et `cert_pem` sont fournis → Mode 2
    Sinon → Mode 1 (certificat récupéré via `client_uid`)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePostRequest
*/
func (a *TraitementFactureAPIService) SoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePost(ctx context.Context) ApiSoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePostRequest {
	return ApiSoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SoumettreFactureCompleteResponse
func (a *TraitementFactureAPIService) SoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePostExecute(r ApiSoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePostRequest) (*SoumettreFactureCompleteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SoumettreFactureCompleteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TraitementFactureAPIService.SoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/traitement/factures/soumettre-complete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.soumettreFactureCompleteRequest == nil {
		return localVarReturnValue, nil, reportError("soumettreFactureCompleteRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.soumettreFactureCompleteRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPostRequest struct {
	ctx context.Context
	ApiService *TraitementFactureAPIService
	soumettreFactureCompleteRequest *SoumettreFactureCompleteRequest
}

func (r ApiSoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPostRequest) SoumettreFactureCompleteRequest(soumettreFactureCompleteRequest SoumettreFactureCompleteRequest) ApiSoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPostRequest {
	r.soumettreFactureCompleteRequest = &soumettreFactureCompleteRequest
	return r
}

func (r ApiSoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPostRequest) Execute() (*ReponseTache, *http.Response, error) {
	return r.ApiService.SoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPostExecute(r)
}

/*
SoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPost Soumettre une facture complète (asynchrone avec Celery)

Version asynchrone de l'endpoint `/factures/soumettre-complete` utilisant Celery pour le traitement en arrière-plan.

    **Workflow automatisé (identique à la version synchrone) :**
    1. **Auto-enrichissement** (optionnel) : récupère les données via APIs publiques et Chorus Pro/AFNOR
    2. **Génération PDF Factur-X** : crée un PDF/A-3 avec XML embarqué
    3. **Signature électronique** (optionnelle) : signe le PDF avec un certificat
    4. **Soumission** : envoie vers la destination choisie (Chorus Pro ou AFNOR PDP)

    **Destinations supportées :**
    - **Chorus Pro** : plateforme B2G française (factures vers secteur public)
    - **AFNOR PDP** : Plateformes de Dématérialisation Partenaires

    **Différences avec la version synchrone :**
    - ✅ **Non-bloquant** : Retourne immédiatement un `id_tache` (HTTP 202 Accepted)
    - ✅ **Traitement en arrière-plan** : La facture est traitée par un worker Celery
    - ✅ **Suivi d'avancement** : Utilisez `/taches/{id_tache}/statut` pour suivre le statut
    - ✅ **Idéal pour gros volumes** : Permet de traiter de nombreuses factures en parallèle

    **Comment utiliser :**
    1. **Soumission** : Appelez cet endpoint avec vos données de facture
    2. **Retour immédiat** : Vous recevez un `id_tache` (ex: "abc123-def456")
    3. **Suivi** : Appelez `/taches/{id_tache}/statut` pour vérifier l'avancement
    4. **Résultat** : Quand `statut = "SUCCESS"`, le champ `resultat` contient la réponse complète

    **Credentials et signature** : Mêmes modes que la version synchrone (JWT ou payload).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPostRequest
*/
func (a *TraitementFactureAPIService) SoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPost(ctx context.Context) ApiSoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPostRequest {
	return ApiSoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReponseTache
func (a *TraitementFactureAPIService) SoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPostExecute(r ApiSoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPostRequest) (*ReponseTache, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReponseTache
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TraitementFactureAPIService.SoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/traitement/factures/soumettre-complete-async"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.soumettreFactureCompleteRequest == nil {
		return localVarReturnValue, nil, reportError("soumettreFactureCompleteRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.soumettreFactureCompleteRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiValiderPdfFacturxApiV1TraitementValiderPdfFacturxPostRequest struct {
	ctx context.Context
	ApiService *TraitementFactureAPIService
	fichierPdf *os.File
	profil *ProfilAPI
	useVerapdf *bool
}

// Fichier PDF Factur-X à valider (format .pdf).
func (r ApiValiderPdfFacturxApiV1TraitementValiderPdfFacturxPostRequest) FichierPdf(fichierPdf *os.File) ApiValiderPdfFacturxApiV1TraitementValiderPdfFacturxPostRequest {
	r.fichierPdf = fichierPdf
	return r
}

func (r ApiValiderPdfFacturxApiV1TraitementValiderPdfFacturxPostRequest) Profil(profil ProfilAPI) ApiValiderPdfFacturxApiV1TraitementValiderPdfFacturxPostRequest {
	r.profil = &profil
	return r
}

// Active la validation stricte PDF/A avec VeraPDF (recommandé pour la production). Si False, utilise une validation basique par métadonnées.
func (r ApiValiderPdfFacturxApiV1TraitementValiderPdfFacturxPostRequest) UseVerapdf(useVerapdf bool) ApiValiderPdfFacturxApiV1TraitementValiderPdfFacturxPostRequest {
	r.useVerapdf = &useVerapdf
	return r
}

func (r ApiValiderPdfFacturxApiV1TraitementValiderPdfFacturxPostRequest) Execute() (*ResultatValidationPDFAPI, *http.Response, error) {
	return r.ApiService.ValiderPdfFacturxApiV1TraitementValiderPdfFacturxPostExecute(r)
}

/*
ValiderPdfFacturxApiV1TraitementValiderPdfFacturxPost Valider un PDF Factur-X complet

Valide un PDF Factur-X complet selon les normes européennes et françaises.

## Normes de validation appliquées

- **EN 16931** : Norme sémantique européenne (directive 2014/55/UE)
- **ISO 19005-3** (PDF/A-3) : Archivage électronique à long terme
- **Factur-X / ZUGFeRD** : Spécification franco-allemande
- **Schematron** : Validation des règles métier XML
- **eIDAS** : Règlement européen sur l'identification électronique (signatures)

## Contrôles effectués

### 1. Extraction et validation du XML Factur-X
**Contrôles réalisés :**
- Présence d'un fichier XML embarqué (`factur-x.xml` ou `zugferd-invoice.xml`)
- Détection automatique du profil (MINIMUM, BASIC, EN16931, EXTENDED)
- Parsing XML avec validation UTF-8
- Extraction du GuidelineSpecifiedDocumentContextParameter/ID

**Validation Schematron :**
- Règles métier du profil détecté (MINIMUM : 45 règles, EN16931 : 178 règles)
- Cardinalité des éléments obligatoires
- Cohérence des calculs (montants HT, TVA, TTC, remises)
- Formats des identifiants (SIRET, TVA intracommunautaire, IBAN)
- Codes normalisés (codes pays ISO, unités UN/ECE, codes TVA)

**Ce qui est vérifié :**
- ✅ Structure XML conforme XSD Cross Industry Invoice
- ✅ Namespace UN/CEFACT correct
- ✅ Règles de gestion européennes (BR-xx)
- ✅ Règles françaises spécifiques (FR-xx)

### 2. Conformité PDF/A-3
**Validation de base (métadonnées) :**
- Présence du champ `/Type` à `Catalog`
- Métadonnée `pdfaid:part` = 3 (PDF/A-3)
- Métadonnée `pdfaid:conformance` = B ou U
- Version PDF >= 1.4

**Validation stricte VeraPDF (si use_verapdf=True) :**
- 146+ règles ISO 19005-3 (PDF/A-3B)
- Absence de contenu interdit (JavaScript, multimedia, formulaires dynamiques)
- Polices embarquées et sous-ensembles corrects
- Espaces colorimétriques conformes (sRGB, DeviceGray)
- Structure de fichier valide (cross-reference table)
- Métadonnées XMP conformes ISO 16684-1

**Ce qui est vérifié :**
- ✅ Fichier archivable à long terme (20+ ans)
- ✅ Lisibilité garantie (polices embarquées)
- ✅ Conformité légale (France, Allemagne, UE)

### 3. Métadonnées XMP (eXtensible Metadata Platform)
**Contrôles réalisés :**
- Présence du bloc `<?xpacket>` avec métadonnées XMP
- Namespace `fx:` pour Factur-X : `urn:factur-x:pdfa:CrossIndustryDocument:invoice:1p0#`
- Champs Factur-X obligatoires :
  - `fx:ConformanceLevel` : Profil (MINIMUM, BASIC, EN16931, EXTENDED)
  - `fx:DocumentFileName` : Nom du XML embarqué
  - `fx:DocumentType` : "INVOICE"
  - `fx:Version` : Version Factur-X (1.0.07)

**Ce qui est vérifié :**
- ✅ Métadonnées conformes ISO 16684-1
- ✅ Profil Factur-X déclaré correct
- ✅ Version Factur-X supportée

### 4. Signatures électroniques
**Détection et analyse :**
- Présence de dictionnaires `/Sig` dans le PDF
- Type de signature : PAdES (PDF Advanced Electronic Signature)
- Extraction des informations :
  - Nom du signataire (`/Name`)
  - Date de signature (`/M`)
  - Raison de la signature (`/Reason`)
  - Lieu de signature (`/Location`)
  - Type de signature (approval, certification)

**Ce qui est vérifié :**
- ✅ Présence de signatures ou cachets
- ✅ Nombre de signatures (mono ou multi-signature)
- ℹ️ Pas de vérification cryptographique (nécessite certificats)

## Paramètres

- **fichier_pdf** (requis) : Le fichier PDF Factur-X à valider
- **profil** (optionnel) : Profil attendu. Si absent, détection automatique depuis le XML
- **use_verapdf** (optionnel, défaut=false) : Active la validation stricte PDF/A avec VeraPDF
  - `false` : Validation rapide par métadonnées (2-3 secondes)
  - `true` : Validation complète ISO 19005-3 (15-30 secondes, **recommandé en production**)

## Réponse détaillée

```json
{
  "est_conforme": true,
  "xml": {
    "present": true,
    "conforme": true,
    "profil": "EN16931",
    "erreurs": []
  },
  "pdfa": {
    "conforme": true,
    "version": "PDF/A-3B",
    "methode": "verapdf",
    "erreurs": []
  },
  "xmp": {
    "present": true,
    "conforme": true,
    "metadonnees": {...}
  },
  "signatures": {
    "present": true,
    "nombre": 1,
    "details": [...]
  }
}
```

## Cas d'usage

- **Avant envoi** : Valider la facture générée avant transmission à un client
- **À réception** : Vérifier la conformité d'une facture reçue d'un fournisseur
- **Audit** : Contrôler la qualité de lots de factures
- **Conformité légale** : S'assurer du respect des obligations B2B/B2G en France
- **Debugging** : Identifier les problèmes dans le processus de génération
- **Archivage** : Garantir la validité à long terme (PDF/A-3)

## Temps de traitement

- Validation basique : 2-3 secondes
- Validation VeraPDF : 15-30 secondes (dépend de la taille du PDF)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiValiderPdfFacturxApiV1TraitementValiderPdfFacturxPostRequest
*/
func (a *TraitementFactureAPIService) ValiderPdfFacturxApiV1TraitementValiderPdfFacturxPost(ctx context.Context) ApiValiderPdfFacturxApiV1TraitementValiderPdfFacturxPostRequest {
	return ApiValiderPdfFacturxApiV1TraitementValiderPdfFacturxPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ResultatValidationPDFAPI
func (a *TraitementFactureAPIService) ValiderPdfFacturxApiV1TraitementValiderPdfFacturxPostExecute(r ApiValiderPdfFacturxApiV1TraitementValiderPdfFacturxPostRequest) (*ResultatValidationPDFAPI, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResultatValidationPDFAPI
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TraitementFactureAPIService.ValiderPdfFacturxApiV1TraitementValiderPdfFacturxPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/traitement/valider-pdf-facturx"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fichierPdf == nil {
		return localVarReturnValue, nil, reportError("fichierPdf is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	var fichierPdfLocalVarFormFileName string
	var fichierPdfLocalVarFileName     string
	var fichierPdfLocalVarFileBytes    []byte

	fichierPdfLocalVarFormFileName = "fichier_pdf"
	fichierPdfLocalVarFile := r.fichierPdf

	if fichierPdfLocalVarFile != nil {
		fbs, _ := io.ReadAll(fichierPdfLocalVarFile)

		fichierPdfLocalVarFileBytes = fbs
		fichierPdfLocalVarFileName = fichierPdfLocalVarFile.Name()
		fichierPdfLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fichierPdfLocalVarFileBytes, fileName: fichierPdfLocalVarFileName, formFileName: fichierPdfLocalVarFormFileName})
	}
	if r.profil != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "profil", r.profil, "", "")
	}
	if r.useVerapdf != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "use_verapdf", r.useVerapdf, "", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPostRequest struct {
	ctx context.Context
	ApiService *TraitementFactureAPIService
	fichierPdf *os.File
	profil *ProfilAPI
	useVerapdf *bool
}

// Fichier PDF Factur-X à valider (format .pdf).
func (r ApiValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPostRequest) FichierPdf(fichierPdf *os.File) ApiValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPostRequest {
	r.fichierPdf = fichierPdf
	return r
}

func (r ApiValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPostRequest) Profil(profil ProfilAPI) ApiValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPostRequest {
	r.profil = &profil
	return r
}

// Active la validation stricte PDF/A avec VeraPDF (recommandé pour la production). Peut prendre plusieurs secondes.
func (r ApiValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPostRequest) UseVerapdf(useVerapdf bool) ApiValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPostRequest {
	r.useVerapdf = &useVerapdf
	return r
}

func (r ApiValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPostRequest) Execute() (*ReponseTache, *http.Response, error) {
	return r.ApiService.ValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPostExecute(r)
}

/*
ValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPost Valider un PDF Factur-X (asynchrone avec polling)

Valide un PDF Factur-X de manière asynchrone avec système de polling.

## Fonctionnement

1. **Soumission** : Le PDF est mis en file d'attente pour validation asynchrone
2. **Retour immédiat** : Vous recevez un `id_tache` (HTTP 202)
3. **Suivi** : Utilisez l'endpoint `/taches/{id_tache}/statut` pour suivre l'avancement

## Avantages du mode asynchrone

- **Pas de timeout** : Idéal pour les gros PDFs ou la validation VeraPDF (qui peut prendre plusieurs secondes)
- **Scalabilité** : Les validations sont traitées par des workers Celery dédiés
- **Suivi d'état** : Permet de suivre la progression de la validation
- **Non-bloquant** : Votre client ne reste pas en attente pendant la validation

## Quand utiliser ce mode ?

- **Validation VeraPDF activée** (`use_verapdf=True`) : La validation stricte peut prendre 2-10 secondes
- **Gros fichiers PDF** : PDFs > 1 MB
- **Traitement par lots** : Validation de multiples factures en parallèle
- **Intégration asynchrone** : Votre système supporte le polling

## Contrôles effectués

### 1. Extraction et validation du XML Factur-X
- Vérifie la présence d'un fichier XML embarqué conforme Factur-X
- Détecte automatiquement le profil utilisé (MINIMUM, BASIC, EN16931, EXTENDED)
- Valide le XML contre les règles Schematron du profil détecté

### 2. Conformité PDF/A
- **Sans VeraPDF** : Validation basique par métadonnées (rapide, ~100ms)
- **Avec VeraPDF** : Validation stricte selon ISO 19005 (146+ règles, 2-10s)
  - Détecte la version PDF/A (PDF/A-1, PDF/A-3, etc.)
  - Rapports détaillés des non-conformités

### 3. Métadonnées XMP
- Vérifie la présence de métadonnées XMP dans le PDF
- Valide la conformité des métadonnées Factur-X (profil, version)
- Extrait toutes les métadonnées XMP disponibles

### 4. Signatures électroniques
- Détecte la présence de signatures ou cachets électroniques
- Extrait les informations sur chaque signature (signataire, date, raison)
- Compte le nombre de signatures présentes

## Paramètres

- **fichier_pdf** : Le fichier PDF Factur-X à valider
- **profil** : Le profil Factur-X attendu (optionnel). Si non spécifié, le profil
  sera automatiquement détecté depuis le fichier XML embarqué.
- **use_verapdf** : Active la validation stricte PDF/A avec VeraPDF.
  ⚠️ **Attention** : VeraPDF peut prendre 2-10 secondes selon la taille du PDF.
  Recommandé uniquement en mode asynchrone pour éviter les timeouts.

## Récupération du résultat

Après soumission, utilisez l'endpoint `GET /taches/{id_tache}/statut` pour récupérer le résultat.

**Exemple de polling** :
```python
import requests
import time

# 1. Soumettre la tâche
response = requests.post("/valider-facturx-async", files={"fichier_pdf": pdf_file})
task_id = response.json()["id_tache"]

# 2. Polling toutes les 2 secondes
while True:
    status_response = requests.get(f"/taches/{task_id}/statut")
    status = status_response.json()

    if status["statut"] == "SUCCESS":
        resultat = status["resultat"]["resultat_validation"]
        print(f"Conforme: {resultat['est_conforme']}")
        break
    elif status["statut"] == "FAILURE":
        print(f"Erreur: {status['resultat']['message_erreur']}")
        break

    time.sleep(2)  # Attendre 2 secondes avant le prochain check
```

## Cas d'usage

- Valider des factures avant envoi avec VeraPDF (validation stricte)
- Traiter des lots de factures en parallèle
- Intégrer la validation dans un pipeline asynchrone
- Valider des PDFs volumineux sans risque de timeout

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPostRequest
*/
func (a *TraitementFactureAPIService) ValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPost(ctx context.Context) ApiValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPostRequest {
	return ApiValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReponseTache
func (a *TraitementFactureAPIService) ValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPostExecute(r ApiValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPostRequest) (*ReponseTache, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReponseTache
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TraitementFactureAPIService.ValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/traitement/valider-facturx-async"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fichierPdf == nil {
		return localVarReturnValue, nil, reportError("fichierPdf is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	var fichierPdfLocalVarFormFileName string
	var fichierPdfLocalVarFileName     string
	var fichierPdfLocalVarFileBytes    []byte

	fichierPdfLocalVarFormFileName = "fichier_pdf"
	fichierPdfLocalVarFile := r.fichierPdf

	if fichierPdfLocalVarFile != nil {
		fbs, _ := io.ReadAll(fichierPdfLocalVarFile)

		fichierPdfLocalVarFileBytes = fbs
		fichierPdfLocalVarFileName = fichierPdfLocalVarFile.Name()
		fichierPdfLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fichierPdfLocalVarFileBytes, fileName: fichierPdfLocalVarFileName, formFileName: fichierPdfLocalVarFormFileName})
	}
	if r.profil != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "profil", r.profil, "", "")
	}
	if r.useVerapdf != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "use_verapdf", r.useVerapdf, "", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPostRequest struct {
	ctx context.Context
	ApiService *TraitementFactureAPIService
	fichierPdf *os.File
}

// Fichier PDF à valider (sera analysé pour détecter et valider les signatures)
func (r ApiValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPostRequest) FichierPdf(fichierPdf *os.File) ApiValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPostRequest {
	r.fichierPdf = fichierPdf
	return r
}

func (r ApiValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPostRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.ValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPostExecute(r)
}

/*
ValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPost Valider les signatures électroniques d'un PDF

Valide les signatures électroniques présentes dans un PDF uploadé.

    **Vérifications effectuées** :
    - Présence de signatures
    - Intégrité du document (non modifié depuis signature)
    - Validité des certificats
    - Chaîne de confiance (si disponible)
    - Présence d'horodatage (PAdES-B-T)
    - Données de validation (PAdES-B-LT)

    **Standards supportés** : PAdES-B-B, PAdES-B-T, PAdES-B-LT, ISO 32000-2.

    **⚠️ Note** : Cette validation est technique (intégrité cryptographique). La validité juridique
    dépend du niveau eIDAS du certificat (SES/AdES/QES) et du contexte d'utilisation.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPostRequest
*/
func (a *TraitementFactureAPIService) ValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPost(ctx context.Context) ApiValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPostRequest {
	return ApiValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return interface{}
func (a *TraitementFactureAPIService) ValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPostExecute(r ApiValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPostRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TraitementFactureAPIService.ValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/traitement/valider-signature-pdf"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fichierPdf == nil {
		return localVarReturnValue, nil, reportError("fichierPdf is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	var fichierPdfLocalVarFormFileName string
	var fichierPdfLocalVarFileName     string
	var fichierPdfLocalVarFileBytes    []byte

	fichierPdfLocalVarFormFileName = "fichier_pdf"
	fichierPdfLocalVarFile := r.fichierPdf

	if fichierPdfLocalVarFile != nil {
		fbs, _ := io.ReadAll(fichierPdfLocalVarFile)

		fichierPdfLocalVarFileBytes = fbs
		fichierPdfLocalVarFileName = fichierPdfLocalVarFile.Name()
		fichierPdfLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fichierPdfLocalVarFileBytes, fileName: fichierPdfLocalVarFileName, formFileName: fichierPdfLocalVarFormFileName})
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiValiderXmlApiV1TraitementValiderXmlPostRequest struct {
	ctx context.Context
	ApiService *TraitementFactureAPIService
	fichierXml *os.File
	profil *ProfilAPI
}

// Fichier XML Factur-X à valider (format .xml).
func (r ApiValiderXmlApiV1TraitementValiderXmlPostRequest) FichierXml(fichierXml *os.File) ApiValiderXmlApiV1TraitementValiderXmlPostRequest {
	r.fichierXml = fichierXml
	return r
}

// Profil de validation (MINIMUM, BASIC, EN16931, EXTENDED).
func (r ApiValiderXmlApiV1TraitementValiderXmlPostRequest) Profil(profil ProfilAPI) ApiValiderXmlApiV1TraitementValiderXmlPostRequest {
	r.profil = &profil
	return r
}

func (r ApiValiderXmlApiV1TraitementValiderXmlPostRequest) Execute() (*ReponseValidationSucces, *http.Response, error) {
	return r.ApiService.ValiderXmlApiV1TraitementValiderXmlPostExecute(r)
}

/*
ValiderXmlApiV1TraitementValiderXmlPost Valider un XML Factur-X existant

Valide un fichier XML Factur-X contre les règles métier Schematron selon la norme EN 16931.

## Norme appliquée

**Schematron ISO/IEC 19757-3** : Langage de validation de règles métier pour XML
- Validation sémantique (au-delà de la syntaxe XSD)
- Règles métier européennes EN 16931
- Règles françaises spécifiques Factur-X
- Calculs arithmétiques et cohérence des données

## Profils et règles validées

### MINIMUM (45 règles)
- Identifiant de facture unique
- Dates (émission, échéance)
- Identifiants parties (SIRET/SIREN)
- Montant total TTC

### BASIC (102 règles)
- Toutes les règles MINIMUM
- Lignes de facture détaillées
- Calculs de TVA basiques
- Modes de paiement
- Références (commande, contrat)

### EN16931 (178 règles)
- Toutes les règles BASIC
- **Règles européennes (BR-xx)** : 81 règles business
- **Règles françaises (FR-xx)** : 12 règles spécifiques France
- **Calculs avancés (CR-xx)** : 32 règles de calcul
- **Codes normalisés (CL-xx)** : 52 listes de codes

### EXTENDED (210+ règles)
- Toutes les règles EN16931
- Informations logistiques
- Données comptables avancées
- Références externes multiples

## Contrôles effectués

### 1. Validation syntaxique
- Parsing XML correct (UTF-8, bien formé)
- Namespaces UN/CEFACT présents
- Structure hiérarchique respectée

### 2. Règles business (BR-xx)
Exemples :
- `BR-1` : Le total de la facture doit être égal à la somme des totaux de lignes + montants au niveau document
- `BR-CO-10` : La somme des montants de base de TVA doit être égale au total net de la facture
- `BR-16` : Le code de devise de la facture doit figurer dans la liste ISO 4217

### 3. Règles françaises (FR-xx)
Exemples :
- `FR-1` : Le SIRET fournisseur doit avoir 14 chiffres
- `FR-2` : Le SIRET client doit avoir 14 chiffres (si présent)
- `FR-5` : Le numéro de TVA intracommunautaire doit être au format FRxx999999999

### 4. Règles de calcul (CR-xx)
- Montants HT + TVA = TTC
- Somme des lignes = Total document
- Remises et majorations correctement appliquées
- Arrondis conformes (2 décimales pour les montants)

### 5. Codes normalisés (CL-xx)
- Codes pays ISO 3166-1 alpha-2
- Codes devises ISO 4217
- Unités de mesure UN/ECE Rec 20
- Codes TVA (types, catégories, exonérations)
- SchemeID pour identifiants (0002=SIREN, 0009=SIRET, etc.)

## Processus de validation

1. **Chargement XSLT** : Fichier Schematron converti en XSLT (Saxon-HE)
2. **Transformation** : Application des règles sur le XML
3. **Analyse résultats** : Extraction des erreurs (`failed-assert`) et avertissements (`successful-report`)
4. **Rapport** : Liste structurée des non-conformités

## Réponses

**200 OK** : XML conforme
```json
{
  "message": "Le XML est conforme au profil EN16931"
}
```

**400 Bad Request** : XML non conforme
```json
{
  "detail": [
    "[BR-1] Le total de la facture (120.00) ne correspond pas à la somme calculée (100.00 + 20.00)",
    "[FR-1] Le SIRET fournisseur doit contenir exactement 14 chiffres"
  ]
}
```

## Cas d'usage

- **Pré-validation** : Vérifier un XML avant intégration dans un PDF/A
- **Debugging** : Identifier précisément les erreurs de génération
- **Tests** : Valider des XMLs de test ou d'exemple
- **Conformité** : S'assurer du respect des règles européennes et françaises
- **Développement** : Tester rapidement sans générer de PDF

## Temps de traitement

- Profil MINIMUM : ~0.5 seconde
- Profil EN16931 : ~1-2 secondes
- Profil EXTENDED : ~2-3 secondes

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiValiderXmlApiV1TraitementValiderXmlPostRequest
*/
func (a *TraitementFactureAPIService) ValiderXmlApiV1TraitementValiderXmlPost(ctx context.Context) ApiValiderXmlApiV1TraitementValiderXmlPostRequest {
	return ApiValiderXmlApiV1TraitementValiderXmlPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReponseValidationSucces
func (a *TraitementFactureAPIService) ValiderXmlApiV1TraitementValiderXmlPostExecute(r ApiValiderXmlApiV1TraitementValiderXmlPostRequest) (*ReponseValidationSucces, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReponseValidationSucces
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TraitementFactureAPIService.ValiderXmlApiV1TraitementValiderXmlPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/traitement/valider-xml"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fichierXml == nil {
		return localVarReturnValue, nil, reportError("fichierXml is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	var fichierXmlLocalVarFormFileName string
	var fichierXmlLocalVarFileName     string
	var fichierXmlLocalVarFileBytes    []byte

	fichierXmlLocalVarFormFileName = "fichier_xml"
	fichierXmlLocalVarFile := r.fichierXml

	if fichierXmlLocalVarFile != nil {
		fbs, _ := io.ReadAll(fichierXmlLocalVarFile)

		fichierXmlLocalVarFileBytes = fbs
		fichierXmlLocalVarFileName = fichierXmlLocalVarFile.Name()
		fichierXmlLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fichierXmlLocalVarFileBytes, fileName: fichierXmlLocalVarFileName, formFileName: fichierXmlLocalVarFormFileName})
	}
	if r.profil != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "profil", r.profil, "", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ReponseValidationErreur
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

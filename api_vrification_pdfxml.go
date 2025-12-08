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


// VrificationPDFXMLAPIService VrificationPDFXMLAPI service
type VrificationPDFXMLAPIService service

type ApiObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGetRequest struct {
	ctx context.Context
	ApiService *VrificationPDFXMLAPIService
	idTache string
}

func (r ApiObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGetRequest) Execute() (*StatutTache, *http.Response, error) {
	return r.ApiService.ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGetExecute(r)
}

/*
ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet Obtenir le statut d'une vérification asynchrone

Récupère le statut et le résultat d'une tâche de vérification asynchrone.

**Statuts possibles:**
- `PENDING`: Tâche en attente dans la file
- `STARTED`: Tâche en cours d'exécution
- `SUCCESS`: Tâche terminée avec succès (voir `resultat`)
- `FAILURE`: Erreur système (exception non gérée)

**Note:** Le champ `resultat.statut` peut être "SUCCES" ou "ERREUR"
indépendamment du statut Celery (qui sera toujours SUCCESS si la tâche s'est exécutée).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param idTache
 @return ApiObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGetRequest
*/
func (a *VrificationPDFXMLAPIService) ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet(ctx context.Context, idTache string) ApiObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGetRequest {
	return ApiObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGetRequest{
		ApiService: a,
		ctx: ctx,
		idTache: idTache,
	}
}

// Execute executes the request
//  @return StatutTache
func (a *VrificationPDFXMLAPIService) ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGetExecute(r ApiObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGetRequest) (*StatutTache, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StatutTache
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VrificationPDFXMLAPIService.ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/verification/verifier-async/{id_tache}/statut"
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

type ApiObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet_0Request struct {
	ctx context.Context
	ApiService *VrificationPDFXMLAPIService
	idTache string
}

func (r ApiObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet_0Request) Execute() (*StatutTache, *http.Response, error) {
	return r.ApiService.ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet_1Execute(r)
}

/*
ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet_0 Obtenir le statut d'une vérification asynchrone

Récupère le statut et le résultat d'une tâche de vérification asynchrone.

**Statuts possibles:**
- `PENDING`: Tâche en attente dans la file
- `STARTED`: Tâche en cours d'exécution
- `SUCCESS`: Tâche terminée avec succès (voir `resultat`)
- `FAILURE`: Erreur système (exception non gérée)

**Note:** Le champ `resultat.statut` peut être "SUCCES" ou "ERREUR"
indépendamment du statut Celery (qui sera toujours SUCCESS si la tâche s'est exécutée).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param idTache
 @return ApiObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet_0Request
*/
func (a *VrificationPDFXMLAPIService) ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet_1(ctx context.Context, idTache string) ApiObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet_0Request {
	return ApiObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet_0Request{
		ApiService: a,
		ctx: ctx,
		idTache: idTache,
	}
}

// Execute executes the request
//  @return StatutTache
func (a *VrificationPDFXMLAPIService) ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet_1Execute(r ApiObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet_0Request) (*StatutTache, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StatutTache
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VrificationPDFXMLAPIService.ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet_1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/verification/verifier-async/{id_tache}/statut"
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

type ApiVerifierPdfAsyncApiV1VerificationVerifierAsyncPostRequest struct {
	ctx context.Context
	ApiService *VrificationPDFXMLAPIService
	fichierPdf *os.File
	forcerOcr *bool
}

// Fichier PDF Factur-X à vérifier
func (r ApiVerifierPdfAsyncApiV1VerificationVerifierAsyncPostRequest) FichierPdf(fichierPdf *os.File) ApiVerifierPdfAsyncApiV1VerificationVerifierAsyncPostRequest {
	r.fichierPdf = fichierPdf
	return r
}

// Forcer l&#39;utilisation de l&#39;OCR même si le PDF contient du texte natif
func (r ApiVerifierPdfAsyncApiV1VerificationVerifierAsyncPostRequest) ForcerOcr(forcerOcr bool) ApiVerifierPdfAsyncApiV1VerificationVerifierAsyncPostRequest {
	r.forcerOcr = &forcerOcr
	return r
}

func (r ApiVerifierPdfAsyncApiV1VerificationVerifierAsyncPostRequest) Execute() (*ReponseTache, *http.Response, error) {
	return r.ApiService.VerifierPdfAsyncApiV1VerificationVerifierAsyncPostExecute(r)
}

/*
VerifierPdfAsyncApiV1VerificationVerifierAsyncPost Vérifier la conformité PDF/XML Factur-X (asynchrone)

Vérifie la conformité PDF/XML Factur-X de manière asynchrone.

**IMPORTANT**: Seuls les PDF Factur-X (avec XML embarqué) sont acceptés.
Les PDF sans XML Factur-X retourneront une erreur `NOT_FACTURX` dans le résultat.

Cette version utilise une tâche Celery et peut faire appel au service OCR
si le PDF est une image ou si `forcer_ocr=true`.

**Retourne immédiatement** un ID de tâche. Utilisez `/verifier-async/{id_tache}/statut`
pour récupérer le résultat.

**Principe de vérification (Factur-X 1.08):**
- Principe n°2: Le XML ne peut contenir que des infos présentes dans le PDF
- Principe n°4: Toute info XML doit être présente et conforme dans le PDF

**Champs vérifiés:**
- Identification: BT-1 (n° facture), BT-2 (date), BT-3 (type), BT-5 (devise), BT-23 (cadre)
- Vendeur: BT-27 (nom), BT-29 (SIRET), BT-30 (SIREN), BT-31 (TVA)
- Acheteur: BT-44 (nom), BT-46 (SIRET), BT-47 (SIREN), BT-48 (TVA)
- Montants: BT-109 (HT), BT-110 (TVA), BT-112 (TTC), BT-115 (à payer)
- Ventilation TVA: BT-116, BT-117, BT-118, BT-119
- Lignes de facture: BT-153, BT-129, BT-146, BT-131
- Notes obligatoires: PMT, PMD, AAB
- Règle BR-FR-09: cohérence SIRET/SIREN

**Avantages par rapport à la version synchrone:**
- Support OCR pour les PDF images (via service DocTR)
- Timeout plus long pour les gros documents
- Ne bloque pas le serveur

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiVerifierPdfAsyncApiV1VerificationVerifierAsyncPostRequest
*/
func (a *VrificationPDFXMLAPIService) VerifierPdfAsyncApiV1VerificationVerifierAsyncPost(ctx context.Context) ApiVerifierPdfAsyncApiV1VerificationVerifierAsyncPostRequest {
	return ApiVerifierPdfAsyncApiV1VerificationVerifierAsyncPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReponseTache
func (a *VrificationPDFXMLAPIService) VerifierPdfAsyncApiV1VerificationVerifierAsyncPostExecute(r ApiVerifierPdfAsyncApiV1VerificationVerifierAsyncPostRequest) (*ReponseTache, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReponseTache
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VrificationPDFXMLAPIService.VerifierPdfAsyncApiV1VerificationVerifierAsyncPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/verification/verifier-async"

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
	if r.forcerOcr != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "forcer_ocr", r.forcerOcr, "", "")
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

type ApiVerifierPdfAsyncApiV1VerificationVerifierAsyncPost_0Request struct {
	ctx context.Context
	ApiService *VrificationPDFXMLAPIService
	fichierPdf *os.File
	forcerOcr *bool
}

// Fichier PDF Factur-X à vérifier
func (r ApiVerifierPdfAsyncApiV1VerificationVerifierAsyncPost_0Request) FichierPdf(fichierPdf *os.File) ApiVerifierPdfAsyncApiV1VerificationVerifierAsyncPost_0Request {
	r.fichierPdf = fichierPdf
	return r
}

// Forcer l&#39;utilisation de l&#39;OCR même si le PDF contient du texte natif
func (r ApiVerifierPdfAsyncApiV1VerificationVerifierAsyncPost_0Request) ForcerOcr(forcerOcr bool) ApiVerifierPdfAsyncApiV1VerificationVerifierAsyncPost_0Request {
	r.forcerOcr = &forcerOcr
	return r
}

func (r ApiVerifierPdfAsyncApiV1VerificationVerifierAsyncPost_0Request) Execute() (*ReponseTache, *http.Response, error) {
	return r.ApiService.VerifierPdfAsyncApiV1VerificationVerifierAsyncPost_2Execute(r)
}

/*
VerifierPdfAsyncApiV1VerificationVerifierAsyncPost_0 Vérifier la conformité PDF/XML Factur-X (asynchrone)

Vérifie la conformité PDF/XML Factur-X de manière asynchrone.

**IMPORTANT**: Seuls les PDF Factur-X (avec XML embarqué) sont acceptés.
Les PDF sans XML Factur-X retourneront une erreur `NOT_FACTURX` dans le résultat.

Cette version utilise une tâche Celery et peut faire appel au service OCR
si le PDF est une image ou si `forcer_ocr=true`.

**Retourne immédiatement** un ID de tâche. Utilisez `/verifier-async/{id_tache}/statut`
pour récupérer le résultat.

**Principe de vérification (Factur-X 1.08):**
- Principe n°2: Le XML ne peut contenir que des infos présentes dans le PDF
- Principe n°4: Toute info XML doit être présente et conforme dans le PDF

**Champs vérifiés:**
- Identification: BT-1 (n° facture), BT-2 (date), BT-3 (type), BT-5 (devise), BT-23 (cadre)
- Vendeur: BT-27 (nom), BT-29 (SIRET), BT-30 (SIREN), BT-31 (TVA)
- Acheteur: BT-44 (nom), BT-46 (SIRET), BT-47 (SIREN), BT-48 (TVA)
- Montants: BT-109 (HT), BT-110 (TVA), BT-112 (TTC), BT-115 (à payer)
- Ventilation TVA: BT-116, BT-117, BT-118, BT-119
- Lignes de facture: BT-153, BT-129, BT-146, BT-131
- Notes obligatoires: PMT, PMD, AAB
- Règle BR-FR-09: cohérence SIRET/SIREN

**Avantages par rapport à la version synchrone:**
- Support OCR pour les PDF images (via service DocTR)
- Timeout plus long pour les gros documents
- Ne bloque pas le serveur

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiVerifierPdfAsyncApiV1VerificationVerifierAsyncPost_0Request
*/
func (a *VrificationPDFXMLAPIService) VerifierPdfAsyncApiV1VerificationVerifierAsyncPost_2(ctx context.Context) ApiVerifierPdfAsyncApiV1VerificationVerifierAsyncPost_0Request {
	return ApiVerifierPdfAsyncApiV1VerificationVerifierAsyncPost_0Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReponseTache
func (a *VrificationPDFXMLAPIService) VerifierPdfAsyncApiV1VerificationVerifierAsyncPost_2Execute(r ApiVerifierPdfAsyncApiV1VerificationVerifierAsyncPost_0Request) (*ReponseTache, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReponseTache
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VrificationPDFXMLAPIService.VerifierPdfAsyncApiV1VerificationVerifierAsyncPost_2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/verification/verifier-async"

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
	if r.forcerOcr != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "forcer_ocr", r.forcerOcr, "", "")
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

type ApiVerifierPdfSyncApiV1VerificationVerifierPostRequest struct {
	ctx context.Context
	ApiService *VrificationPDFXMLAPIService
	fichierPdf *os.File
}

// Fichier PDF Factur-X à vérifier
func (r ApiVerifierPdfSyncApiV1VerificationVerifierPostRequest) FichierPdf(fichierPdf *os.File) ApiVerifierPdfSyncApiV1VerificationVerifierPostRequest {
	r.fichierPdf = fichierPdf
	return r
}

func (r ApiVerifierPdfSyncApiV1VerificationVerifierPostRequest) Execute() (*ReponseVerificationSucces, *http.Response, error) {
	return r.ApiService.VerifierPdfSyncApiV1VerificationVerifierPostExecute(r)
}

/*
VerifierPdfSyncApiV1VerificationVerifierPost Vérifier la conformité PDF/XML Factur-X (synchrone)

Vérifie la conformité entre le PDF et son XML Factur-X embarqué.

**IMPORTANT**: Seuls les PDF Factur-X (avec XML embarqué) sont acceptés.
Les PDF sans XML Factur-X seront rejetés avec une erreur 400.

Cette version synchrone utilise uniquement l'extraction PDF native (pdfplumber).
Pour les PDF images nécessitant de l'OCR, utilisez l'endpoint `/verifier-async`.

**Principe de vérification (Factur-X 1.08):**
- Principe n°2: Le XML ne peut contenir que des infos présentes dans le PDF
- Principe n°4: Toute info XML doit être présente et conforme dans le PDF

**Champs vérifiés:**
- Identification: BT-1 (n° facture), BT-2 (date), BT-3 (type), BT-5 (devise), BT-23 (cadre)
- Vendeur: BT-27 (nom), BT-29 (SIRET), BT-30 (SIREN), BT-31 (TVA)
- Acheteur: BT-44 (nom), BT-46 (SIRET), BT-47 (SIREN), BT-48 (TVA)
- Montants: BT-109 (HT), BT-110 (TVA), BT-112 (TTC), BT-115 (à payer)
- Ventilation TVA: BT-116, BT-117, BT-118, BT-119
- Lignes de facture: BT-153, BT-129, BT-146, BT-131
- Notes obligatoires: PMT, PMD, AAB
- Règle BR-FR-09: cohérence SIRET/SIREN

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiVerifierPdfSyncApiV1VerificationVerifierPostRequest
*/
func (a *VrificationPDFXMLAPIService) VerifierPdfSyncApiV1VerificationVerifierPost(ctx context.Context) ApiVerifierPdfSyncApiV1VerificationVerifierPostRequest {
	return ApiVerifierPdfSyncApiV1VerificationVerifierPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReponseVerificationSucces
func (a *VrificationPDFXMLAPIService) VerifierPdfSyncApiV1VerificationVerifierPostExecute(r ApiVerifierPdfSyncApiV1VerificationVerifierPostRequest) (*ReponseVerificationSucces, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReponseVerificationSucces
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VrificationPDFXMLAPIService.VerifierPdfSyncApiV1VerificationVerifierPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/verification/verifier"

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 413 {
			var v APIError
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

type ApiVerifierPdfSyncApiV1VerificationVerifierPost_0Request struct {
	ctx context.Context
	ApiService *VrificationPDFXMLAPIService
	fichierPdf *os.File
}

// Fichier PDF Factur-X à vérifier
func (r ApiVerifierPdfSyncApiV1VerificationVerifierPost_0Request) FichierPdf(fichierPdf *os.File) ApiVerifierPdfSyncApiV1VerificationVerifierPost_0Request {
	r.fichierPdf = fichierPdf
	return r
}

func (r ApiVerifierPdfSyncApiV1VerificationVerifierPost_0Request) Execute() (*ReponseVerificationSucces, *http.Response, error) {
	return r.ApiService.VerifierPdfSyncApiV1VerificationVerifierPost_3Execute(r)
}

/*
VerifierPdfSyncApiV1VerificationVerifierPost_0 Vérifier la conformité PDF/XML Factur-X (synchrone)

Vérifie la conformité entre le PDF et son XML Factur-X embarqué.

**IMPORTANT**: Seuls les PDF Factur-X (avec XML embarqué) sont acceptés.
Les PDF sans XML Factur-X seront rejetés avec une erreur 400.

Cette version synchrone utilise uniquement l'extraction PDF native (pdfplumber).
Pour les PDF images nécessitant de l'OCR, utilisez l'endpoint `/verifier-async`.

**Principe de vérification (Factur-X 1.08):**
- Principe n°2: Le XML ne peut contenir que des infos présentes dans le PDF
- Principe n°4: Toute info XML doit être présente et conforme dans le PDF

**Champs vérifiés:**
- Identification: BT-1 (n° facture), BT-2 (date), BT-3 (type), BT-5 (devise), BT-23 (cadre)
- Vendeur: BT-27 (nom), BT-29 (SIRET), BT-30 (SIREN), BT-31 (TVA)
- Acheteur: BT-44 (nom), BT-46 (SIRET), BT-47 (SIREN), BT-48 (TVA)
- Montants: BT-109 (HT), BT-110 (TVA), BT-112 (TTC), BT-115 (à payer)
- Ventilation TVA: BT-116, BT-117, BT-118, BT-119
- Lignes de facture: BT-153, BT-129, BT-146, BT-131
- Notes obligatoires: PMT, PMD, AAB
- Règle BR-FR-09: cohérence SIRET/SIREN

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiVerifierPdfSyncApiV1VerificationVerifierPost_0Request
*/
func (a *VrificationPDFXMLAPIService) VerifierPdfSyncApiV1VerificationVerifierPost_3(ctx context.Context) ApiVerifierPdfSyncApiV1VerificationVerifierPost_0Request {
	return ApiVerifierPdfSyncApiV1VerificationVerifierPost_0Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReponseVerificationSucces
func (a *VrificationPDFXMLAPIService) VerifierPdfSyncApiV1VerificationVerifierPost_3Execute(r ApiVerifierPdfSyncApiV1VerificationVerifierPost_0Request) (*ReponseVerificationSucces, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReponseVerificationSucces
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VrificationPDFXMLAPIService.VerifierPdfSyncApiV1VerificationVerifierPost_3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/verification/verifier"

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 413 {
			var v APIError
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

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
)


// ChorusProAPIService ChorusProAPI service
type ChorusProAPIService service

type ApiAjouterFichierApiV1ChorusProTransversesAjouterFichierPostRequest struct {
	ctx context.Context
	ApiService *ChorusProAPIService
	bodyAjouterFichierApiV1ChorusProTransversesAjouterFichierPost *BodyAjouterFichierApiV1ChorusProTransversesAjouterFichierPost
}

func (r ApiAjouterFichierApiV1ChorusProTransversesAjouterFichierPostRequest) BodyAjouterFichierApiV1ChorusProTransversesAjouterFichierPost(bodyAjouterFichierApiV1ChorusProTransversesAjouterFichierPost BodyAjouterFichierApiV1ChorusProTransversesAjouterFichierPost) ApiAjouterFichierApiV1ChorusProTransversesAjouterFichierPostRequest {
	r.bodyAjouterFichierApiV1ChorusProTransversesAjouterFichierPost = &bodyAjouterFichierApiV1ChorusProTransversesAjouterFichierPost
	return r
}

func (r ApiAjouterFichierApiV1ChorusProTransversesAjouterFichierPostRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.AjouterFichierApiV1ChorusProTransversesAjouterFichierPostExecute(r)
}

/*
AjouterFichierApiV1ChorusProTransversesAjouterFichierPost Ajouter une pièce jointe

Ajoute une pièce jointe au compte utilisateur courant.

    **Taille max** : 10 Mo par fichier

    **Payload exemple** :
    ```json
    {
      "pieceJointeFichier": "JVBERi0xLjQKJeLjz9MKNSAwIG9iago8P...",
      "pieceJointeNom": "bon_commande.pdf",
      "pieceJointeTypeMime": "application/pdf",
      "pieceJointeExtension": "PDF"
    }
    ```

    **Retour** : L'ID de la pièce jointe (`pieceJointeIdFichier`) à utiliser ensuite dans `/factures/completer`.

    **Extensions acceptées** : PDF, JPG, PNG, ZIP, XML, etc.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAjouterFichierApiV1ChorusProTransversesAjouterFichierPostRequest
*/
func (a *ChorusProAPIService) AjouterFichierApiV1ChorusProTransversesAjouterFichierPost(ctx context.Context) ApiAjouterFichierApiV1ChorusProTransversesAjouterFichierPostRequest {
	return ApiAjouterFichierApiV1ChorusProTransversesAjouterFichierPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return interface{}
func (a *ChorusProAPIService) AjouterFichierApiV1ChorusProTransversesAjouterFichierPostExecute(r ApiAjouterFichierApiV1ChorusProTransversesAjouterFichierPostRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChorusProAPIService.AjouterFichierApiV1ChorusProTransversesAjouterFichierPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/chorus-pro/transverses/ajouter-fichier"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bodyAjouterFichierApiV1ChorusProTransversesAjouterFichierPost == nil {
		return localVarReturnValue, nil, reportError("bodyAjouterFichierApiV1ChorusProTransversesAjouterFichierPost is required and must be specified")
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
	localVarPostBody = r.bodyAjouterFichierApiV1ChorusProTransversesAjouterFichierPost
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

type ApiCompleterFactureApiV1ChorusProFacturesCompleterPostRequest struct {
	ctx context.Context
	ApiService *ChorusProAPIService
	bodyCompleterFactureApiV1ChorusProFacturesCompleterPost *BodyCompleterFactureApiV1ChorusProFacturesCompleterPost
}

func (r ApiCompleterFactureApiV1ChorusProFacturesCompleterPostRequest) BodyCompleterFactureApiV1ChorusProFacturesCompleterPost(bodyCompleterFactureApiV1ChorusProFacturesCompleterPost BodyCompleterFactureApiV1ChorusProFacturesCompleterPost) ApiCompleterFactureApiV1ChorusProFacturesCompleterPostRequest {
	r.bodyCompleterFactureApiV1ChorusProFacturesCompleterPost = &bodyCompleterFactureApiV1ChorusProFacturesCompleterPost
	return r
}

func (r ApiCompleterFactureApiV1ChorusProFacturesCompleterPostRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.CompleterFactureApiV1ChorusProFacturesCompleterPostExecute(r)
}

/*
CompleterFactureApiV1ChorusProFacturesCompleterPost Compléter une facture suspendue (Fournisseur)

Complète une facture au statut SUSPENDUE en ajoutant des pièces jointes ou un commentaire.

    **Statut requis** : SUSPENDUE

    **Actions possibles** :
    - Ajouter des pièces jointes (justificatifs, bons de commande, etc.)
    - Modifier le commentaire

    **Payload exemple** :
    ```json
    {
      "identifiantFactureCPP": 12345,
      "commentaire": "Voici les justificatifs demandés",
      "listePiecesJointes": [
        {
          "pieceJointeIdFichier": 98765,
          "pieceJointeNom": "bon_commande.pdf"
        }
      ]
    }
    ```

    **Note** : Les pièces jointes doivent d'abord être uploadées via `/transverses/ajouter-fichier`.

    **Après complétion** : La facture repasse au statut MISE_A_DISPOSITION.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCompleterFactureApiV1ChorusProFacturesCompleterPostRequest
*/
func (a *ChorusProAPIService) CompleterFactureApiV1ChorusProFacturesCompleterPost(ctx context.Context) ApiCompleterFactureApiV1ChorusProFacturesCompleterPostRequest {
	return ApiCompleterFactureApiV1ChorusProFacturesCompleterPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return interface{}
func (a *ChorusProAPIService) CompleterFactureApiV1ChorusProFacturesCompleterPostExecute(r ApiCompleterFactureApiV1ChorusProFacturesCompleterPostRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChorusProAPIService.CompleterFactureApiV1ChorusProFacturesCompleterPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/chorus-pro/factures/completer"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bodyCompleterFactureApiV1ChorusProFacturesCompleterPost == nil {
		return localVarReturnValue, nil, reportError("bodyCompleterFactureApiV1ChorusProFacturesCompleterPost is required and must be specified")
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
	localVarPostBody = r.bodyCompleterFactureApiV1ChorusProFacturesCompleterPost
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

type ApiConsulterFactureApiV1ChorusProFacturesConsulterPostRequest struct {
	ctx context.Context
	ApiService *ChorusProAPIService
	consulterFactureRequest *ConsulterFactureRequest
}

func (r ApiConsulterFactureApiV1ChorusProFacturesConsulterPostRequest) ConsulterFactureRequest(consulterFactureRequest ConsulterFactureRequest) ApiConsulterFactureApiV1ChorusProFacturesConsulterPostRequest {
	r.consulterFactureRequest = &consulterFactureRequest
	return r
}

func (r ApiConsulterFactureApiV1ChorusProFacturesConsulterPostRequest) Execute() (*ConsulterFactureResponse, *http.Response, error) {
	return r.ApiService.ConsulterFactureApiV1ChorusProFacturesConsulterPostExecute(r)
}

/*
ConsulterFactureApiV1ChorusProFacturesConsulterPost Consulter le statut d'une facture

Récupère les informations et le statut actuel d'une facture soumise à Chorus Pro.

    **Retour** :
    - Numéro et date de facture
    - Montant TTC
    - **Statut courant** : SOUMISE, VALIDEE, REJETEE, SUSPENDUE, MANDATEE, MISE_EN_PAIEMENT, etc.
    - Structure destinataire

    **Cas d'usage** :
    - Suivre l'évolution du traitement d'une facture
    - Vérifier si une facture a été validée ou rejetée
    - Obtenir la date de mise en paiement

    **Polling** : Appelez cet endpoint régulièrement pour suivre l'évolution du statut.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiConsulterFactureApiV1ChorusProFacturesConsulterPostRequest
*/
func (a *ChorusProAPIService) ConsulterFactureApiV1ChorusProFacturesConsulterPost(ctx context.Context) ApiConsulterFactureApiV1ChorusProFacturesConsulterPostRequest {
	return ApiConsulterFactureApiV1ChorusProFacturesConsulterPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ConsulterFactureResponse
func (a *ChorusProAPIService) ConsulterFactureApiV1ChorusProFacturesConsulterPostExecute(r ApiConsulterFactureApiV1ChorusProFacturesConsulterPostRequest) (*ConsulterFactureResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConsulterFactureResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChorusProAPIService.ConsulterFactureApiV1ChorusProFacturesConsulterPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/chorus-pro/factures/consulter"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.consulterFactureRequest == nil {
		return localVarReturnValue, nil, reportError("consulterFactureRequest is required and must be specified")
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
	localVarPostBody = r.consulterFactureRequest
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

type ApiConsulterStructureApiV1ChorusProStructuresConsulterPostRequest struct {
	ctx context.Context
	ApiService *ChorusProAPIService
	consulterStructureRequest *ConsulterStructureRequest
}

func (r ApiConsulterStructureApiV1ChorusProStructuresConsulterPostRequest) ConsulterStructureRequest(consulterStructureRequest ConsulterStructureRequest) ApiConsulterStructureApiV1ChorusProStructuresConsulterPostRequest {
	r.consulterStructureRequest = &consulterStructureRequest
	return r
}

func (r ApiConsulterStructureApiV1ChorusProStructuresConsulterPostRequest) Execute() (*ConsulterStructureResponse, *http.Response, error) {
	return r.ApiService.ConsulterStructureApiV1ChorusProStructuresConsulterPostExecute(r)
}

/*
ConsulterStructureApiV1ChorusProStructuresConsulterPost Consulter les détails d'une structure

Récupère les informations détaillées d'une structure Chorus Pro.


    **Retour** :
    - Raison sociale
    - Numéro de TVA intracommunautaire
    - Email de contact
    - **Paramètres obligatoires** : Indique si le code service et/ou numéro d'engagement sont requis pour soumettre une facture

    **Étape typique** : Appelée après `rechercher-structures` pour savoir quels champs sont obligatoires avant de soumettre une facture.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiConsulterStructureApiV1ChorusProStructuresConsulterPostRequest
*/
func (a *ChorusProAPIService) ConsulterStructureApiV1ChorusProStructuresConsulterPost(ctx context.Context) ApiConsulterStructureApiV1ChorusProStructuresConsulterPostRequest {
	return ApiConsulterStructureApiV1ChorusProStructuresConsulterPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ConsulterStructureResponse
func (a *ChorusProAPIService) ConsulterStructureApiV1ChorusProStructuresConsulterPostExecute(r ApiConsulterStructureApiV1ChorusProStructuresConsulterPostRequest) (*ConsulterStructureResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConsulterStructureResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChorusProAPIService.ConsulterStructureApiV1ChorusProStructuresConsulterPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/chorus-pro/structures/consulter"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.consulterStructureRequest == nil {
		return localVarReturnValue, nil, reportError("consulterStructureRequest is required and must be specified")
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
	localVarPostBody = r.consulterStructureRequest
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

type ApiListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGetRequest struct {
	ctx context.Context
	ApiService *ChorusProAPIService
	idStructureCpp int32
	bodyListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet *BodyListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet
}

func (r ApiListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGetRequest) BodyListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet(bodyListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet BodyListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet) ApiListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGetRequest {
	r.bodyListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet = &bodyListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet
	return r
}

func (r ApiListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGetRequest) Execute() (*RechercherServicesResponse, *http.Response, error) {
	return r.ApiService.ListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGetExecute(r)
}

/*
ListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet Lister les services d'une structure

Récupère la liste des services actifs d'une structure publique.

    **Cas d'usage** :
    - Lister les services disponibles pour une administration
    - Vérifier qu'un code service existe avant de soumettre une facture

    **Retour** :
    - Liste des services avec leur code, libellé et statut (actif/inactif)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param idStructureCpp
 @return ApiListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGetRequest
*/
func (a *ChorusProAPIService) ListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet(ctx context.Context, idStructureCpp int32) ApiListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGetRequest {
	return ApiListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGetRequest{
		ApiService: a,
		ctx: ctx,
		idStructureCpp: idStructureCpp,
	}
}

// Execute executes the request
//  @return RechercherServicesResponse
func (a *ChorusProAPIService) ListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGetExecute(r ApiListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGetRequest) (*RechercherServicesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RechercherServicesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChorusProAPIService.ListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/chorus-pro/structures/{id_structure_cpp}/services"
	localVarPath = strings.Replace(localVarPath, "{"+"id_structure_cpp"+"}", url.PathEscape(parameterValueToString(r.idStructureCpp, "idStructureCpp")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bodyListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet == nil {
		return localVarReturnValue, nil, reportError("bodyListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet is required and must be specified")
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
	localVarPostBody = r.bodyListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet
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

type ApiObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPostRequest struct {
	ctx context.Context
	ApiService *ChorusProAPIService
	obtenirIdChorusProRequest *ObtenirIdChorusProRequest
}

func (r ApiObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPostRequest) ObtenirIdChorusProRequest(obtenirIdChorusProRequest ObtenirIdChorusProRequest) ApiObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPostRequest {
	r.obtenirIdChorusProRequest = &obtenirIdChorusProRequest
	return r
}

func (r ApiObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPostRequest) Execute() (*ObtenirIdChorusProResponse, *http.Response, error) {
	return r.ApiService.ObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPostExecute(r)
}

/*
ObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPost Utilitaire : Obtenir l'ID Chorus Pro depuis un SIRET

**Utilitaire pratique** pour obtenir l'ID Chorus Pro d'une structure à partir de son SIRET.


    Cette fonction wrapper combine :
    1. Recherche de la structure par SIRET
    2. Extraction de l'`id_structure_cpp` si une seule structure est trouvée

    **Retour** :
    - `id_structure_cpp` : ID Chorus Pro (0 si non trouvé ou si plusieurs résultats)
    - `designation_structure` : Nom de la structure (si trouvée)
    - `message` : Message explicatif

    **Cas d'usage** :
    - Raccourci pour obtenir directement l'ID Chorus Pro avant de soumettre une facture
    - Alternative simplifiée à `rechercher-structures` + extraction manuelle de l'ID

    **Note** : Si plusieurs structures correspondent au SIRET (rare), retourne 0 et un message d'erreur.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPostRequest
*/
func (a *ChorusProAPIService) ObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPost(ctx context.Context) ApiObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPostRequest {
	return ApiObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ObtenirIdChorusProResponse
func (a *ChorusProAPIService) ObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPostExecute(r ApiObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPostRequest) (*ObtenirIdChorusProResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ObtenirIdChorusProResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChorusProAPIService.ObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/chorus-pro/structures/obtenir-id-depuis-siret"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.obtenirIdChorusProRequest == nil {
		return localVarReturnValue, nil, reportError("obtenirIdChorusProRequest is required and must be specified")
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
	localVarPostBody = r.obtenirIdChorusProRequest
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

type ApiRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePostRequest struct {
	ctx context.Context
	ApiService *ChorusProAPIService
	bodyRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost *BodyRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost
}

func (r ApiRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePostRequest) BodyRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost(bodyRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost BodyRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost) ApiRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePostRequest {
	r.bodyRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost = &bodyRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost
	return r
}

func (r ApiRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePostRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.RechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePostExecute(r)
}

/*
RechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost Rechercher factures reçues (Destinataire)

Recherche les factures reçues par le destinataire connecté.

    **Filtres** :
    - Téléchargée / non téléchargée
    - Dates de réception
    - Statut (MISE_A_DISPOSITION, SUSPENDUE, etc.)
    - Fournisseur

    **Indicateur utile** : `factureTelechargeeParDestinataire` permet de savoir si la facture a déjà été téléchargée.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePostRequest
*/
func (a *ChorusProAPIService) RechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost(ctx context.Context) ApiRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePostRequest {
	return ApiRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return interface{}
func (a *ChorusProAPIService) RechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePostExecute(r ApiRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePostRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChorusProAPIService.RechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/chorus-pro/factures/rechercher-destinataire"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bodyRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost == nil {
		return localVarReturnValue, nil, reportError("bodyRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost is required and must be specified")
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
	localVarPostBody = r.bodyRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost
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

type ApiRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPostRequest struct {
	ctx context.Context
	ApiService *ChorusProAPIService
	bodyRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost *BodyRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost
}

func (r ApiRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPostRequest) BodyRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost(bodyRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost BodyRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost) ApiRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPostRequest {
	r.bodyRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost = &bodyRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost
	return r
}

func (r ApiRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPostRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.RechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPostExecute(r)
}

/*
RechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost Rechercher factures émises (Fournisseur)

Recherche les factures émises par le fournisseur connecté.

    **Filtres disponibles** :
    - Numéro de facture
    - Dates (début/fin)
    - Statut
    - Structure destinataire
    - Montant

    **Cas d'usage** :
    - Suivi des factures émises
    - Vérification des statuts
    - Export pour comptabilité

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPostRequest
*/
func (a *ChorusProAPIService) RechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost(ctx context.Context) ApiRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPostRequest {
	return ApiRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return interface{}
func (a *ChorusProAPIService) RechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPostExecute(r ApiRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPostRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChorusProAPIService.RechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/chorus-pro/factures/rechercher-fournisseur"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bodyRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost == nil {
		return localVarReturnValue, nil, reportError("bodyRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost is required and must be specified")
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
	localVarPostBody = r.bodyRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost
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

type ApiRechercherStructuresApiV1ChorusProStructuresRechercherPostRequest struct {
	ctx context.Context
	ApiService *ChorusProAPIService
	rechercherStructureRequest *RechercherStructureRequest
}

func (r ApiRechercherStructuresApiV1ChorusProStructuresRechercherPostRequest) RechercherStructureRequest(rechercherStructureRequest RechercherStructureRequest) ApiRechercherStructuresApiV1ChorusProStructuresRechercherPostRequest {
	r.rechercherStructureRequest = &rechercherStructureRequest
	return r
}

func (r ApiRechercherStructuresApiV1ChorusProStructuresRechercherPostRequest) Execute() (*RechercherStructureResponse, *http.Response, error) {
	return r.ApiService.RechercherStructuresApiV1ChorusProStructuresRechercherPostExecute(r)
}

/*
RechercherStructuresApiV1ChorusProStructuresRechercherPost Rechercher des structures Chorus Pro

Recherche des structures (entreprises, administrations) enregistrées sur Chorus Pro.

    **Cas d'usage** :
    - Trouver l'ID Chorus Pro d'une structure à partir de son SIRET
    - Vérifier si une structure est enregistrée sur Chorus Pro
    - Lister les structures correspondant à des critères

    **Filtres disponibles** :
    - Identifiant (SIRET, SIREN, etc.)
    - Raison sociale
    - Type d'identifiant
    - Structures privées uniquement

    **Étape typique** : Appelée avant `soumettre-facture` pour obtenir l'`id_structure_cpp` du destinataire.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRechercherStructuresApiV1ChorusProStructuresRechercherPostRequest
*/
func (a *ChorusProAPIService) RechercherStructuresApiV1ChorusProStructuresRechercherPost(ctx context.Context) ApiRechercherStructuresApiV1ChorusProStructuresRechercherPostRequest {
	return ApiRechercherStructuresApiV1ChorusProStructuresRechercherPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RechercherStructureResponse
func (a *ChorusProAPIService) RechercherStructuresApiV1ChorusProStructuresRechercherPostExecute(r ApiRechercherStructuresApiV1ChorusProStructuresRechercherPostRequest) (*RechercherStructureResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RechercherStructureResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChorusProAPIService.RechercherStructuresApiV1ChorusProStructuresRechercherPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/chorus-pro/structures/rechercher"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rechercherStructureRequest == nil {
		return localVarReturnValue, nil, reportError("rechercherStructureRequest is required and must be specified")
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
	localVarPostBody = r.rechercherStructureRequest
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

type ApiRecyclerFactureApiV1ChorusProFacturesRecyclerPostRequest struct {
	ctx context.Context
	ApiService *ChorusProAPIService
	bodyRecyclerFactureApiV1ChorusProFacturesRecyclerPost *BodyRecyclerFactureApiV1ChorusProFacturesRecyclerPost
}

func (r ApiRecyclerFactureApiV1ChorusProFacturesRecyclerPostRequest) BodyRecyclerFactureApiV1ChorusProFacturesRecyclerPost(bodyRecyclerFactureApiV1ChorusProFacturesRecyclerPost BodyRecyclerFactureApiV1ChorusProFacturesRecyclerPost) ApiRecyclerFactureApiV1ChorusProFacturesRecyclerPostRequest {
	r.bodyRecyclerFactureApiV1ChorusProFacturesRecyclerPost = &bodyRecyclerFactureApiV1ChorusProFacturesRecyclerPost
	return r
}

func (r ApiRecyclerFactureApiV1ChorusProFacturesRecyclerPostRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.RecyclerFactureApiV1ChorusProFacturesRecyclerPostExecute(r)
}

/*
RecyclerFactureApiV1ChorusProFacturesRecyclerPost Recycler une facture (Fournisseur)

Recycle une facture au statut A_RECYCLER en modifiant les données d'acheminement.

    **Statut requis** : A_RECYCLER

    **Champs modifiables** :
    - Destinataire (`idStructureCPP`)
    - Code service
    - Numéro d'engagement

    **Cas d'usage** :
    - Erreur de destinataire
    - Changement de service facturation
    - Mise à jour du numéro d'engagement

    **Payload exemple** :
    ```json
    {
      "identifiantFactureCPP": 12345,
      "idStructureCPP": 67890,
      "codeService": "SERVICE_01",
      "numeroEngagement": "ENG2024001"
    }
    ```

    **Note** : La facture conserve son numéro et ses montants, seuls les champs d'acheminement changent.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRecyclerFactureApiV1ChorusProFacturesRecyclerPostRequest
*/
func (a *ChorusProAPIService) RecyclerFactureApiV1ChorusProFacturesRecyclerPost(ctx context.Context) ApiRecyclerFactureApiV1ChorusProFacturesRecyclerPostRequest {
	return ApiRecyclerFactureApiV1ChorusProFacturesRecyclerPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return interface{}
func (a *ChorusProAPIService) RecyclerFactureApiV1ChorusProFacturesRecyclerPostExecute(r ApiRecyclerFactureApiV1ChorusProFacturesRecyclerPostRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChorusProAPIService.RecyclerFactureApiV1ChorusProFacturesRecyclerPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/chorus-pro/factures/recycler"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bodyRecyclerFactureApiV1ChorusProFacturesRecyclerPost == nil {
		return localVarReturnValue, nil, reportError("bodyRecyclerFactureApiV1ChorusProFacturesRecyclerPost is required and must be specified")
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
	localVarPostBody = r.bodyRecyclerFactureApiV1ChorusProFacturesRecyclerPost
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

type ApiSoumettreFactureApiV1ChorusProFacturesSoumettrePostRequest struct {
	ctx context.Context
	ApiService *ChorusProAPIService
	soumettreFactureRequest *SoumettreFactureRequest
}

func (r ApiSoumettreFactureApiV1ChorusProFacturesSoumettrePostRequest) SoumettreFactureRequest(soumettreFactureRequest SoumettreFactureRequest) ApiSoumettreFactureApiV1ChorusProFacturesSoumettrePostRequest {
	r.soumettreFactureRequest = &soumettreFactureRequest
	return r
}

func (r ApiSoumettreFactureApiV1ChorusProFacturesSoumettrePostRequest) Execute() (*SoumettreFactureResponse, *http.Response, error) {
	return r.ApiService.SoumettreFactureApiV1ChorusProFacturesSoumettrePostExecute(r)
}

/*
SoumettreFactureApiV1ChorusProFacturesSoumettrePost Soumettre une facture à Chorus Pro

Soumet une facture électronique à une structure publique via Chorus Pro.


    **📋 Workflow complet** :
    1. **Uploader le PDF Factur-X** via `/transverses/ajouter-fichier` → récupérer `pieceJointeId`
    2. **Obtenir l'ID structure** via `/structures/rechercher` ou `/structures/obtenir-id-depuis-siret`
    3. **Vérifier les paramètres obligatoires** via `/structures/consulter`
    4. **Soumettre la facture** avec le `piece_jointe_principale_id` obtenu à l'étape 1

    **Pré-requis** :
    1. Avoir l'`id_structure_cpp` du destinataire (via `/structures/rechercher`)
    2. Connaître les paramètres obligatoires (via `/structures/consulter`) :
       - Code service si `code_service_doit_etre_renseigne=true`
       - Numéro d'engagement si `numero_ej_doit_etre_renseigne=true`
    3. Avoir uploadé le PDF Factur-X (via `/transverses/ajouter-fichier`)

    **Format attendu** :
    - `piece_jointe_principale_id` : ID retourné par `/transverses/ajouter-fichier`
    - Montants : Chaînes de caractères avec 2 décimales (ex: "1250.50")
    - Dates : Format ISO 8601 (YYYY-MM-DD)

    **Retour** :
    - `identifiant_facture_cpp` : ID Chorus Pro de la facture créée
    - `numero_flux_depot` : Numéro de suivi du dépôt

    **Statuts possibles après soumission** :
    - SOUMISE : En attente de validation
    - VALIDEE : Validée par le destinataire
    - REJETEE : Rejetée (erreur de données ou refus métier)
    - SUSPENDUE : En attente d'informations complémentaires

    **Note** : Utilisez `/factures/consulter` pour suivre l'évolution du statut.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSoumettreFactureApiV1ChorusProFacturesSoumettrePostRequest
*/
func (a *ChorusProAPIService) SoumettreFactureApiV1ChorusProFacturesSoumettrePost(ctx context.Context) ApiSoumettreFactureApiV1ChorusProFacturesSoumettrePostRequest {
	return ApiSoumettreFactureApiV1ChorusProFacturesSoumettrePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SoumettreFactureResponse
func (a *ChorusProAPIService) SoumettreFactureApiV1ChorusProFacturesSoumettrePostExecute(r ApiSoumettreFactureApiV1ChorusProFacturesSoumettrePostRequest) (*SoumettreFactureResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SoumettreFactureResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChorusProAPIService.SoumettreFactureApiV1ChorusProFacturesSoumettrePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/chorus-pro/factures/soumettre"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.soumettreFactureRequest == nil {
		return localVarReturnValue, nil, reportError("soumettreFactureRequest is required and must be specified")
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
	localVarPostBody = r.soumettreFactureRequest
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

type ApiTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePostRequest struct {
	ctx context.Context
	ApiService *ChorusProAPIService
	bodyTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost *BodyTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost
}

func (r ApiTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePostRequest) BodyTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost(bodyTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost BodyTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost) ApiTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePostRequest {
	r.bodyTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost = &bodyTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost
	return r
}

func (r ApiTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePostRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.TelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePostExecute(r)
}

/*
TelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost Télécharger un groupe de factures

Télécharge une ou plusieurs factures (max 10 recommandé) avec leurs pièces jointes.

    **Formats disponibles** :
    - PDF : Fichier PDF uniquement
    - XML : Fichier XML uniquement
    - ZIP : Archive contenant PDF + XML + pièces jointes

    **Taille maximale** : 120 Mo par téléchargement

    **Payload exemple** :
    ```json
    {
      "listeIdentifiantsFactureCPP": [12345, 12346],
      "inclurePiecesJointes": true,
      "formatFichier": "ZIP"
    }
    ```

    **Retour** : Le fichier est encodé en base64 dans le champ `fichierBase64`.

    **Note** : Le flag `factureTelechargeeParDestinataire` est mis à jour automatiquement.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePostRequest
*/
func (a *ChorusProAPIService) TelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost(ctx context.Context) ApiTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePostRequest {
	return ApiTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return interface{}
func (a *ChorusProAPIService) TelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePostExecute(r ApiTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePostRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChorusProAPIService.TelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/chorus-pro/factures/telecharger-groupe"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bodyTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost == nil {
		return localVarReturnValue, nil, reportError("bodyTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost is required and must be specified")
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
	localVarPostBody = r.bodyTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost
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

type ApiTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePostRequest struct {
	ctx context.Context
	ApiService *ChorusProAPIService
	bodyTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost *BodyTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost
}

func (r ApiTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePostRequest) BodyTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost(bodyTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost BodyTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost) ApiTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePostRequest {
	r.bodyTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost = &bodyTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost
	return r
}

func (r ApiTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePostRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.TraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePostExecute(r)
}

/*
TraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost Traiter une facture reçue (Destinataire)

Change le statut d'une facture reçue.

    **Statuts possibles** :
    - MISE_A_DISPOSITION : Facture acceptée
    - SUSPENDUE : En attente d'informations complémentaires (motif obligatoire)
    - REJETEE : Facture refusée (motif obligatoire)
    - MANDATEE : Facture mandatée
    - MISE_EN_PAIEMENT : Facture en cours de paiement
    - COMPTABILISEE : Facture comptabilisée
    - MISE_A_DISPOSITION_COMPTABLE : Mise à disposition comptable
    - A_RECYCLER : À recycler
    - COMPLETEE : Complétée
    - SERVICE-FAIT : Service fait
    - PRISE_EN_COMPTE_DESTINATAIRE : Prise en compte
    - TRANSMISE_MOA : Transmise à la MOA

    **Payload exemple** :
    ```json
    {
      "identifiantFactureCPP": 12345,
      "nouveauStatut": "REJETEE",
      "motifRejet": "Facture en double",
      "commentaire": "Facture déjà reçue sous la référence ABC123"
    }
    ```

    **Règles** :
    - Un motif est **obligatoire** pour SUSPENDUE et REJETEE
    - Seuls certains statuts sont autorisés selon le statut actuel de la facture

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePostRequest
*/
func (a *ChorusProAPIService) TraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost(ctx context.Context) ApiTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePostRequest {
	return ApiTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return interface{}
func (a *ChorusProAPIService) TraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePostExecute(r ApiTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePostRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChorusProAPIService.TraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/chorus-pro/factures/traiter-facture-recue"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bodyTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost == nil {
		return localVarReturnValue, nil, reportError("bodyTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost is required and must be specified")
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
	localVarPostBody = r.bodyTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost
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

type ApiValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPostRequest struct {
	ctx context.Context
	ApiService *ChorusProAPIService
	bodyValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost *BodyValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost
}

func (r ApiValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPostRequest) BodyValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost(bodyValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost BodyValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost) ApiValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPostRequest {
	r.bodyValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost = &bodyValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost
	return r
}

func (r ApiValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPostRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.ValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPostExecute(r)
}

/*
ValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost Consulter une facture (Valideur)

Consulte facture (valideur).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPostRequest
*/
func (a *ChorusProAPIService) ValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost(ctx context.Context) ApiValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPostRequest {
	return ApiValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return interface{}
func (a *ChorusProAPIService) ValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPostExecute(r ApiValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPostRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChorusProAPIService.ValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/chorus-pro/factures/valideur/consulter"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bodyValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost == nil {
		return localVarReturnValue, nil, reportError("bodyValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost is required and must be specified")
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
	localVarPostBody = r.bodyValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost
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

type ApiValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPostRequest struct {
	ctx context.Context
	ApiService *ChorusProAPIService
	bodyValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost *BodyValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost
}

func (r ApiValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPostRequest) BodyValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost(bodyValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost BodyValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost) ApiValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPostRequest {
	r.bodyValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost = &bodyValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost
	return r
}

func (r ApiValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPostRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.ValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPostExecute(r)
}

/*
ValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost Rechercher factures à valider (Valideur)

Recherche les factures en attente de validation par le valideur connecté.

    **Rôle** : Valideur dans le circuit de validation interne.

    **Filtres** : Dates, structure, service, etc.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPostRequest
*/
func (a *ChorusProAPIService) ValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost(ctx context.Context) ApiValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPostRequest {
	return ApiValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return interface{}
func (a *ChorusProAPIService) ValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPostExecute(r ApiValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPostRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChorusProAPIService.ValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/chorus-pro/factures/valideur/rechercher"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bodyValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost == nil {
		return localVarReturnValue, nil, reportError("bodyValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost is required and must be specified")
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
	localVarPostBody = r.bodyValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost
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

type ApiValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPostRequest struct {
	ctx context.Context
	ApiService *ChorusProAPIService
	bodyValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost *BodyValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost
}

func (r ApiValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPostRequest) BodyValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost(bodyValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost BodyValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost) ApiValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPostRequest {
	r.bodyValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost = &bodyValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost
	return r
}

func (r ApiValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPostRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.ValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPostExecute(r)
}

/*
ValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost Valider ou refuser une facture (Valideur)

Valide ou refuse une facture en attente de validation.

    **Actions** :
    - Valider : La facture passe au statut suivant du circuit
    - Refuser : La facture est rejetée (motif obligatoire)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPostRequest
*/
func (a *ChorusProAPIService) ValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost(ctx context.Context) ApiValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPostRequest {
	return ApiValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return interface{}
func (a *ChorusProAPIService) ValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPostExecute(r ApiValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPostRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChorusProAPIService.ValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/chorus-pro/factures/valideur/traiter"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bodyValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost == nil {
		return localVarReturnValue, nil, reportError("bodyValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost is required and must be specified")
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
	localVarPostBody = r.bodyValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost
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

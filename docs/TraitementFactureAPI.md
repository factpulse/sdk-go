# \TraitementFactureAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenererCertificatTestApiV1TraitementGenererCertificatTestPost**](TraitementFactureAPI.md#GenererCertificatTestApiV1TraitementGenererCertificatTestPost) | **Post** /api/v1/traitement/generer-certificat-test | Générer un certificat X.509 auto-signé de test
[**GenererFactureApiV1TraitementGenererFacturePost**](TraitementFactureAPI.md#GenererFactureApiV1TraitementGenererFacturePost) | **Post** /api/v1/traitement/generer-facture | Générer une facture Factur-X
[**ObtenirStatutTacheApiV1TraitementTachesIdTacheStatutGet**](TraitementFactureAPI.md#ObtenirStatutTacheApiV1TraitementTachesIdTacheStatutGet) | **Get** /api/v1/traitement/taches/{id_tache}/statut | Obtenir le statut d&#39;une tâche de génération
[**SignerPdfApiV1TraitementSignerPdfPost**](TraitementFactureAPI.md#SignerPdfApiV1TraitementSignerPdfPost) | **Post** /api/v1/traitement/signer-pdf | Signer un PDF avec le certificat du client (PAdES-B-LT)
[**SignerPdfAsyncApiV1TraitementSignerPdfAsyncPost**](TraitementFactureAPI.md#SignerPdfAsyncApiV1TraitementSignerPdfAsyncPost) | **Post** /api/v1/traitement/signer-pdf-async | Signer un PDF de manière asynchrone (Celery)
[**SoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePost**](TraitementFactureAPI.md#SoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePost) | **Post** /api/v1/traitement/factures/soumettre-complete | Soumettre une facture complète (génération + signature + soumission)
[**SoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPost**](TraitementFactureAPI.md#SoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPost) | **Post** /api/v1/traitement/factures/soumettre-complete-async | Soumettre une facture complète (asynchrone avec Celery)
[**ValiderPdfFacturxApiV1TraitementValiderPdfFacturxPost**](TraitementFactureAPI.md#ValiderPdfFacturxApiV1TraitementValiderPdfFacturxPost) | **Post** /api/v1/traitement/valider-pdf-facturx | Valider un PDF Factur-X complet
[**ValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPost**](TraitementFactureAPI.md#ValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPost) | **Post** /api/v1/traitement/valider-facturx-async | Valider un PDF Factur-X (asynchrone avec polling)
[**ValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPost**](TraitementFactureAPI.md#ValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPost) | **Post** /api/v1/traitement/valider-signature-pdf | Valider les signatures électroniques d&#39;un PDF
[**ValiderXmlApiV1TraitementValiderXmlPost**](TraitementFactureAPI.md#ValiderXmlApiV1TraitementValiderXmlPost) | **Post** /api/v1/traitement/valider-xml | Valider un XML Factur-X existant



## GenererCertificatTestApiV1TraitementGenererCertificatTestPost

> GenerateCertificateResponse GenererCertificatTestApiV1TraitementGenererCertificatTestPost(ctx).GenerateCertificateRequest(generateCertificateRequest).Execute()

Générer un certificat X.509 auto-signé de test



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	generateCertificateRequest := *openapiclient.NewGenerateCertificateRequest() // GenerateCertificateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TraitementFactureAPI.GenererCertificatTestApiV1TraitementGenererCertificatTestPost(context.Background()).GenerateCertificateRequest(generateCertificateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TraitementFactureAPI.GenererCertificatTestApiV1TraitementGenererCertificatTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenererCertificatTestApiV1TraitementGenererCertificatTestPost`: GenerateCertificateResponse
	fmt.Fprintf(os.Stdout, "Response from `TraitementFactureAPI.GenererCertificatTestApiV1TraitementGenererCertificatTestPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenererCertificatTestApiV1TraitementGenererCertificatTestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateCertificateRequest** | [**GenerateCertificateRequest**](GenerateCertificateRequest.md) |  | 

### Return type

[**GenerateCertificateResponse**](GenerateCertificateResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenererFactureApiV1TraitementGenererFacturePost

> ReponseTache GenererFactureApiV1TraitementGenererFacturePost(ctx).DonneesFacture(donneesFacture).Profil(profil).FormatSortie(formatSortie).AutoEnrichir(autoEnrichir).SourcePdf(sourcePdf).Execute()

Générer une facture Factur-X



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	donneesFacture := "donneesFacture_example" // string | Données de la facture au format JSON.              Deux formats acceptés :             1. **Format classique** : Structure complète FactureFacturX (tous les champs)             2. **Format simplifié** (🆕 P0.1) : Structure minimale avec auto-enrichissement              Le format est détecté automatiquement !             
	profil := openapiclient.ProfilAPI("MINIMUM") // ProfilAPI | Profil Factur-X : MINIMUM, BASIC, EN16931 ou EXTENDED. (optional)
	formatSortie := openapiclient.FormatSortie("xml") // FormatSortie | Format de sortie : 'xml' (XML seul) ou 'pdf' (PDF Factur-X avec XML embarqué). (optional)
	autoEnrichir := true // bool | 🆕 Activer l'auto-enrichissement depuis SIRET/SIREN (format simplifié uniquement) (optional) (default to true)
	sourcePdf := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TraitementFactureAPI.GenererFactureApiV1TraitementGenererFacturePost(context.Background()).DonneesFacture(donneesFacture).Profil(profil).FormatSortie(formatSortie).AutoEnrichir(autoEnrichir).SourcePdf(sourcePdf).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TraitementFactureAPI.GenererFactureApiV1TraitementGenererFacturePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenererFactureApiV1TraitementGenererFacturePost`: ReponseTache
	fmt.Fprintf(os.Stdout, "Response from `TraitementFactureAPI.GenererFactureApiV1TraitementGenererFacturePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenererFactureApiV1TraitementGenererFacturePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **donneesFacture** | **string** | Données de la facture au format JSON.              Deux formats acceptés :             1. **Format classique** : Structure complète FactureFacturX (tous les champs)             2. **Format simplifié** (🆕 P0.1) : Structure minimale avec auto-enrichissement              Le format est détecté automatiquement !              | 
 **profil** | [**ProfilAPI**](ProfilAPI.md) | Profil Factur-X : MINIMUM, BASIC, EN16931 ou EXTENDED. | 
 **formatSortie** | [**FormatSortie**](FormatSortie.md) | Format de sortie : &#39;xml&#39; (XML seul) ou &#39;pdf&#39; (PDF Factur-X avec XML embarqué). | 
 **autoEnrichir** | **bool** | 🆕 Activer l&#39;auto-enrichissement depuis SIRET/SIREN (format simplifié uniquement) | [default to true]
 **sourcePdf** | ***os.File** |  | 

### Return type

[**ReponseTache**](ReponseTache.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObtenirStatutTacheApiV1TraitementTachesIdTacheStatutGet

> StatutTache ObtenirStatutTacheApiV1TraitementTachesIdTacheStatutGet(ctx, idTache).Execute()

Obtenir le statut d'une tâche de génération



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	idTache := "idTache_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TraitementFactureAPI.ObtenirStatutTacheApiV1TraitementTachesIdTacheStatutGet(context.Background(), idTache).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TraitementFactureAPI.ObtenirStatutTacheApiV1TraitementTachesIdTacheStatutGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObtenirStatutTacheApiV1TraitementTachesIdTacheStatutGet`: StatutTache
	fmt.Fprintf(os.Stdout, "Response from `TraitementFactureAPI.ObtenirStatutTacheApiV1TraitementTachesIdTacheStatutGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idTache** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiObtenirStatutTacheApiV1TraitementTachesIdTacheStatutGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StatutTache**](StatutTache.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignerPdfApiV1TraitementSignerPdfPost

> interface{} SignerPdfApiV1TraitementSignerPdfPost(ctx).FichierPdf(fichierPdf).Raison(raison).Localisation(localisation).Contact(contact).FieldName(fieldName).UsePadesLt(usePadesLt).UseTimestamp(useTimestamp).Execute()

Signer un PDF avec le certificat du client (PAdES-B-LT)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	fichierPdf := os.NewFile(1234, "some_file") // *os.File | Fichier PDF à signer (sera traité puis retourné signé en base64)
	raison := "raison_example" // string |  (optional)
	localisation := "localisation_example" // string |  (optional)
	contact := "contact_example" // string |  (optional)
	fieldName := "fieldName_example" // string | Nom du champ de signature PDF (optional) (default to "FactPulseSignature")
	usePadesLt := true // bool | Activer PAdES-B-LT (archivage long terme avec données de validation embarquées). NÉCESSITE un certificat avec accès OCSP/CRL. (optional) (default to false)
	useTimestamp := true // bool | Activer l'horodatage RFC 3161 avec FreeTSA (PAdES-B-T) (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TraitementFactureAPI.SignerPdfApiV1TraitementSignerPdfPost(context.Background()).FichierPdf(fichierPdf).Raison(raison).Localisation(localisation).Contact(contact).FieldName(fieldName).UsePadesLt(usePadesLt).UseTimestamp(useTimestamp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TraitementFactureAPI.SignerPdfApiV1TraitementSignerPdfPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignerPdfApiV1TraitementSignerPdfPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TraitementFactureAPI.SignerPdfApiV1TraitementSignerPdfPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignerPdfApiV1TraitementSignerPdfPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fichierPdf** | ***os.File** | Fichier PDF à signer (sera traité puis retourné signé en base64) | 
 **raison** | **string** |  | 
 **localisation** | **string** |  | 
 **contact** | **string** |  | 
 **fieldName** | **string** | Nom du champ de signature PDF | [default to &quot;FactPulseSignature&quot;]
 **usePadesLt** | **bool** | Activer PAdES-B-LT (archivage long terme avec données de validation embarquées). NÉCESSITE un certificat avec accès OCSP/CRL. | [default to false]
 **useTimestamp** | **bool** | Activer l&#39;horodatage RFC 3161 avec FreeTSA (PAdES-B-T) | [default to true]

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignerPdfAsyncApiV1TraitementSignerPdfAsyncPost

> interface{} SignerPdfAsyncApiV1TraitementSignerPdfAsyncPost(ctx).FichierPdf(fichierPdf).Raison(raison).Localisation(localisation).Contact(contact).FieldName(fieldName).UsePadesLt(usePadesLt).UseTimestamp(useTimestamp).Execute()

Signer un PDF de manière asynchrone (Celery)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	fichierPdf := os.NewFile(1234, "some_file") // *os.File | Fichier PDF à signer (traité de manière asynchrone)
	raison := "raison_example" // string |  (optional)
	localisation := "localisation_example" // string |  (optional)
	contact := "contact_example" // string |  (optional)
	fieldName := "fieldName_example" // string | Nom du champ de signature PDF (optional) (default to "FactPulseSignature")
	usePadesLt := true // bool | Activer PAdES-B-LT (archivage long terme avec données de validation embarquées). NÉCESSITE un certificat avec accès OCSP/CRL. (optional) (default to false)
	useTimestamp := true // bool | Activer l'horodatage RFC 3161 avec FreeTSA (PAdES-B-T) (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TraitementFactureAPI.SignerPdfAsyncApiV1TraitementSignerPdfAsyncPost(context.Background()).FichierPdf(fichierPdf).Raison(raison).Localisation(localisation).Contact(contact).FieldName(fieldName).UsePadesLt(usePadesLt).UseTimestamp(useTimestamp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TraitementFactureAPI.SignerPdfAsyncApiV1TraitementSignerPdfAsyncPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignerPdfAsyncApiV1TraitementSignerPdfAsyncPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TraitementFactureAPI.SignerPdfAsyncApiV1TraitementSignerPdfAsyncPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignerPdfAsyncApiV1TraitementSignerPdfAsyncPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fichierPdf** | ***os.File** | Fichier PDF à signer (traité de manière asynchrone) | 
 **raison** | **string** |  | 
 **localisation** | **string** |  | 
 **contact** | **string** |  | 
 **fieldName** | **string** | Nom du champ de signature PDF | [default to &quot;FactPulseSignature&quot;]
 **usePadesLt** | **bool** | Activer PAdES-B-LT (archivage long terme avec données de validation embarquées). NÉCESSITE un certificat avec accès OCSP/CRL. | [default to false]
 **useTimestamp** | **bool** | Activer l&#39;horodatage RFC 3161 avec FreeTSA (PAdES-B-T) | [default to true]

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePost

> SoumettreFactureCompleteResponse SoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePost(ctx).SoumettreFactureCompleteRequest(soumettreFactureCompleteRequest).Execute()

Soumettre une facture complète (génération + signature + soumission)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	soumettreFactureCompleteRequest := *openapiclient.NewSoumettreFactureCompleteRequest(*openapiclient.NewDonneesFactureSimplifiees("Numero_example", map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}), "PdfSource_example", openapiclient.Destination{DestinationAFNOR: openapiclient.NewDestinationAFNOR()}) // SoumettreFactureCompleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TraitementFactureAPI.SoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePost(context.Background()).SoumettreFactureCompleteRequest(soumettreFactureCompleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TraitementFactureAPI.SoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePost`: SoumettreFactureCompleteResponse
	fmt.Fprintf(os.Stdout, "Response from `TraitementFactureAPI.SoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **soumettreFactureCompleteRequest** | [**SoumettreFactureCompleteRequest**](SoumettreFactureCompleteRequest.md) |  | 

### Return type

[**SoumettreFactureCompleteResponse**](SoumettreFactureCompleteResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPost

> ReponseTache SoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPost(ctx).SoumettreFactureCompleteRequest(soumettreFactureCompleteRequest).Execute()

Soumettre une facture complète (asynchrone avec Celery)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	soumettreFactureCompleteRequest := *openapiclient.NewSoumettreFactureCompleteRequest(*openapiclient.NewDonneesFactureSimplifiees("Numero_example", map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}), "PdfSource_example", openapiclient.Destination{DestinationAFNOR: openapiclient.NewDestinationAFNOR()}) // SoumettreFactureCompleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TraitementFactureAPI.SoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPost(context.Background()).SoumettreFactureCompleteRequest(soumettreFactureCompleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TraitementFactureAPI.SoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPost`: ReponseTache
	fmt.Fprintf(os.Stdout, "Response from `TraitementFactureAPI.SoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **soumettreFactureCompleteRequest** | [**SoumettreFactureCompleteRequest**](SoumettreFactureCompleteRequest.md) |  | 

### Return type

[**ReponseTache**](ReponseTache.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValiderPdfFacturxApiV1TraitementValiderPdfFacturxPost

> ResultatValidationPDFAPI ValiderPdfFacturxApiV1TraitementValiderPdfFacturxPost(ctx).FichierPdf(fichierPdf).Profil(profil).UseVerapdf(useVerapdf).Execute()

Valider un PDF Factur-X complet



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	fichierPdf := os.NewFile(1234, "some_file") // *os.File | Fichier PDF Factur-X à valider (format .pdf).
	profil := openapiclient.ProfilAPI("MINIMUM") // ProfilAPI |  (optional)
	useVerapdf := true // bool | Active la validation stricte PDF/A avec VeraPDF (recommandé pour la production). Si False, utilise une validation basique par métadonnées. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TraitementFactureAPI.ValiderPdfFacturxApiV1TraitementValiderPdfFacturxPost(context.Background()).FichierPdf(fichierPdf).Profil(profil).UseVerapdf(useVerapdf).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TraitementFactureAPI.ValiderPdfFacturxApiV1TraitementValiderPdfFacturxPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValiderPdfFacturxApiV1TraitementValiderPdfFacturxPost`: ResultatValidationPDFAPI
	fmt.Fprintf(os.Stdout, "Response from `TraitementFactureAPI.ValiderPdfFacturxApiV1TraitementValiderPdfFacturxPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValiderPdfFacturxApiV1TraitementValiderPdfFacturxPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fichierPdf** | ***os.File** | Fichier PDF Factur-X à valider (format .pdf). | 
 **profil** | [**ProfilAPI**](ProfilAPI.md) |  | 
 **useVerapdf** | **bool** | Active la validation stricte PDF/A avec VeraPDF (recommandé pour la production). Si False, utilise une validation basique par métadonnées. | [default to false]

### Return type

[**ResultatValidationPDFAPI**](ResultatValidationPDFAPI.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPost

> ReponseTache ValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPost(ctx).FichierPdf(fichierPdf).Profil(profil).UseVerapdf(useVerapdf).Execute()

Valider un PDF Factur-X (asynchrone avec polling)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	fichierPdf := os.NewFile(1234, "some_file") // *os.File | Fichier PDF Factur-X à valider (format .pdf).
	profil := openapiclient.ProfilAPI("MINIMUM") // ProfilAPI |  (optional)
	useVerapdf := true // bool | Active la validation stricte PDF/A avec VeraPDF (recommandé pour la production). Peut prendre plusieurs secondes. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TraitementFactureAPI.ValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPost(context.Background()).FichierPdf(fichierPdf).Profil(profil).UseVerapdf(useVerapdf).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TraitementFactureAPI.ValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPost`: ReponseTache
	fmt.Fprintf(os.Stdout, "Response from `TraitementFactureAPI.ValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValiderPdfFacturxAsyncApiV1TraitementValiderFacturxAsyncPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fichierPdf** | ***os.File** | Fichier PDF Factur-X à valider (format .pdf). | 
 **profil** | [**ProfilAPI**](ProfilAPI.md) |  | 
 **useVerapdf** | **bool** | Active la validation stricte PDF/A avec VeraPDF (recommandé pour la production). Peut prendre plusieurs secondes. | [default to false]

### Return type

[**ReponseTache**](ReponseTache.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPost

> interface{} ValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPost(ctx).FichierPdf(fichierPdf).Execute()

Valider les signatures électroniques d'un PDF



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	fichierPdf := os.NewFile(1234, "some_file") // *os.File | Fichier PDF à valider (sera analysé pour détecter et valider les signatures)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TraitementFactureAPI.ValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPost(context.Background()).FichierPdf(fichierPdf).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TraitementFactureAPI.ValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TraitementFactureAPI.ValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fichierPdf** | ***os.File** | Fichier PDF à valider (sera analysé pour détecter et valider les signatures) | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValiderXmlApiV1TraitementValiderXmlPost

> ReponseValidationSucces ValiderXmlApiV1TraitementValiderXmlPost(ctx).FichierXml(fichierXml).Profil(profil).Execute()

Valider un XML Factur-X existant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	fichierXml := os.NewFile(1234, "some_file") // *os.File | Fichier XML Factur-X à valider (format .xml).
	profil := openapiclient.ProfilAPI("MINIMUM") // ProfilAPI | Profil de validation (MINIMUM, BASIC, EN16931, EXTENDED). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TraitementFactureAPI.ValiderXmlApiV1TraitementValiderXmlPost(context.Background()).FichierXml(fichierXml).Profil(profil).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TraitementFactureAPI.ValiderXmlApiV1TraitementValiderXmlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValiderXmlApiV1TraitementValiderXmlPost`: ReponseValidationSucces
	fmt.Fprintf(os.Stdout, "Response from `TraitementFactureAPI.ValiderXmlApiV1TraitementValiderXmlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValiderXmlApiV1TraitementValiderXmlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fichierXml** | ***os.File** | Fichier XML Factur-X à valider (format .xml). | 
 **profil** | [**ProfilAPI**](ProfilAPI.md) | Profil de validation (MINIMUM, BASIC, EN16931, EXTENDED). | 

### Return type

[**ReponseValidationSucces**](ReponseValidationSucces.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


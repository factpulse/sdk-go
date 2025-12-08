# \VrificationPDFXMLAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet**](VrificationPDFXMLAPI.md#ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet) | **Get** /api/v1/verification/verifier-async/{id_tache}/statut | Obtenir le statut d&#39;une vérification asynchrone
[**ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet_0**](VrificationPDFXMLAPI.md#ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet_0) | **Get** /api/v1/verification/verifier-async/{id_tache}/statut | Obtenir le statut d&#39;une vérification asynchrone
[**VerifierPdfAsyncApiV1VerificationVerifierAsyncPost**](VrificationPDFXMLAPI.md#VerifierPdfAsyncApiV1VerificationVerifierAsyncPost) | **Post** /api/v1/verification/verifier-async | Vérifier la conformité PDF/XML Factur-X (asynchrone)
[**VerifierPdfAsyncApiV1VerificationVerifierAsyncPost_0**](VrificationPDFXMLAPI.md#VerifierPdfAsyncApiV1VerificationVerifierAsyncPost_0) | **Post** /api/v1/verification/verifier-async | Vérifier la conformité PDF/XML Factur-X (asynchrone)
[**VerifierPdfSyncApiV1VerificationVerifierPost**](VrificationPDFXMLAPI.md#VerifierPdfSyncApiV1VerificationVerifierPost) | **Post** /api/v1/verification/verifier | Vérifier la conformité PDF/XML Factur-X (synchrone)
[**VerifierPdfSyncApiV1VerificationVerifierPost_0**](VrificationPDFXMLAPI.md#VerifierPdfSyncApiV1VerificationVerifierPost_0) | **Post** /api/v1/verification/verifier | Vérifier la conformité PDF/XML Factur-X (synchrone)



## ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet

> StatutTache ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet(ctx, idTache).Execute()

Obtenir le statut d'une vérification asynchrone



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
	resp, r, err := apiClient.VrificationPDFXMLAPI.ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet(context.Background(), idTache).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VrificationPDFXMLAPI.ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet`: StatutTache
	fmt.Fprintf(os.Stdout, "Response from `VrificationPDFXMLAPI.ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idTache** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGetRequest struct via the builder pattern


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


## ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet_0

> StatutTache ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet_0(ctx, idTache).Execute()

Obtenir le statut d'une vérification asynchrone



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
	resp, r, err := apiClient.VrificationPDFXMLAPI.ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet_0(context.Background(), idTache).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VrificationPDFXMLAPI.ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet_0`: StatutTache
	fmt.Fprintf(os.Stdout, "Response from `VrificationPDFXMLAPI.ObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idTache** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiObtenirStatutVerificationApiV1VerificationVerifierAsyncIdTacheStatutGet_1Request struct via the builder pattern


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


## VerifierPdfAsyncApiV1VerificationVerifierAsyncPost

> ReponseTache VerifierPdfAsyncApiV1VerificationVerifierAsyncPost(ctx).FichierPdf(fichierPdf).ForcerOcr(forcerOcr).Execute()

Vérifier la conformité PDF/XML Factur-X (asynchrone)



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
	fichierPdf := os.NewFile(1234, "some_file") // *os.File | Fichier PDF Factur-X à vérifier
	forcerOcr := true // bool | Forcer l'utilisation de l'OCR même si le PDF contient du texte natif (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VrificationPDFXMLAPI.VerifierPdfAsyncApiV1VerificationVerifierAsyncPost(context.Background()).FichierPdf(fichierPdf).ForcerOcr(forcerOcr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VrificationPDFXMLAPI.VerifierPdfAsyncApiV1VerificationVerifierAsyncPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifierPdfAsyncApiV1VerificationVerifierAsyncPost`: ReponseTache
	fmt.Fprintf(os.Stdout, "Response from `VrificationPDFXMLAPI.VerifierPdfAsyncApiV1VerificationVerifierAsyncPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifierPdfAsyncApiV1VerificationVerifierAsyncPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fichierPdf** | ***os.File** | Fichier PDF Factur-X à vérifier | 
 **forcerOcr** | **bool** | Forcer l&#39;utilisation de l&#39;OCR même si le PDF contient du texte natif | [default to false]

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


## VerifierPdfAsyncApiV1VerificationVerifierAsyncPost_0

> ReponseTache VerifierPdfAsyncApiV1VerificationVerifierAsyncPost_0(ctx).FichierPdf(fichierPdf).ForcerOcr(forcerOcr).Execute()

Vérifier la conformité PDF/XML Factur-X (asynchrone)



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
	fichierPdf := os.NewFile(1234, "some_file") // *os.File | Fichier PDF Factur-X à vérifier
	forcerOcr := true // bool | Forcer l'utilisation de l'OCR même si le PDF contient du texte natif (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VrificationPDFXMLAPI.VerifierPdfAsyncApiV1VerificationVerifierAsyncPost_0(context.Background()).FichierPdf(fichierPdf).ForcerOcr(forcerOcr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VrificationPDFXMLAPI.VerifierPdfAsyncApiV1VerificationVerifierAsyncPost_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifierPdfAsyncApiV1VerificationVerifierAsyncPost_0`: ReponseTache
	fmt.Fprintf(os.Stdout, "Response from `VrificationPDFXMLAPI.VerifierPdfAsyncApiV1VerificationVerifierAsyncPost_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifierPdfAsyncApiV1VerificationVerifierAsyncPost_2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fichierPdf** | ***os.File** | Fichier PDF Factur-X à vérifier | 
 **forcerOcr** | **bool** | Forcer l&#39;utilisation de l&#39;OCR même si le PDF contient du texte natif | [default to false]

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


## VerifierPdfSyncApiV1VerificationVerifierPost

> ReponseVerificationSucces VerifierPdfSyncApiV1VerificationVerifierPost(ctx).FichierPdf(fichierPdf).Execute()

Vérifier la conformité PDF/XML Factur-X (synchrone)



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
	fichierPdf := os.NewFile(1234, "some_file") // *os.File | Fichier PDF Factur-X à vérifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VrificationPDFXMLAPI.VerifierPdfSyncApiV1VerificationVerifierPost(context.Background()).FichierPdf(fichierPdf).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VrificationPDFXMLAPI.VerifierPdfSyncApiV1VerificationVerifierPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifierPdfSyncApiV1VerificationVerifierPost`: ReponseVerificationSucces
	fmt.Fprintf(os.Stdout, "Response from `VrificationPDFXMLAPI.VerifierPdfSyncApiV1VerificationVerifierPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifierPdfSyncApiV1VerificationVerifierPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fichierPdf** | ***os.File** | Fichier PDF Factur-X à vérifier | 

### Return type

[**ReponseVerificationSucces**](ReponseVerificationSucces.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifierPdfSyncApiV1VerificationVerifierPost_0

> ReponseVerificationSucces VerifierPdfSyncApiV1VerificationVerifierPost_0(ctx).FichierPdf(fichierPdf).Execute()

Vérifier la conformité PDF/XML Factur-X (synchrone)



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
	fichierPdf := os.NewFile(1234, "some_file") // *os.File | Fichier PDF Factur-X à vérifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VrificationPDFXMLAPI.VerifierPdfSyncApiV1VerificationVerifierPost_0(context.Background()).FichierPdf(fichierPdf).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VrificationPDFXMLAPI.VerifierPdfSyncApiV1VerificationVerifierPost_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifierPdfSyncApiV1VerificationVerifierPost_0`: ReponseVerificationSucces
	fmt.Fprintf(os.Stdout, "Response from `VrificationPDFXMLAPI.VerifierPdfSyncApiV1VerificationVerifierPost_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifierPdfSyncApiV1VerificationVerifierPost_3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fichierPdf** | ***os.File** | Fichier PDF Factur-X à vérifier | 

### Return type

[**ReponseVerificationSucces**](ReponseVerificationSucces.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


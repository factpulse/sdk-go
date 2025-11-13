# \SignatureLectroniqueAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenererCertificatTestApiV1TraitementGenererCertificatTestPost**](SignatureLectroniqueAPI.md#GenererCertificatTestApiV1TraitementGenererCertificatTestPost) | **Post** /api/v1/traitement/generer-certificat-test | Générer un certificat X.509 auto-signé de test
[**SignerPdfApiV1TraitementSignerPdfPost**](SignatureLectroniqueAPI.md#SignerPdfApiV1TraitementSignerPdfPost) | **Post** /api/v1/traitement/signer-pdf | Signer un PDF avec le certificat du client (PAdES-B-LT)
[**SignerPdfAsyncApiV1TraitementSignerPdfAsyncPost**](SignatureLectroniqueAPI.md#SignerPdfAsyncApiV1TraitementSignerPdfAsyncPost) | **Post** /api/v1/traitement/signer-pdf-async | Signer un PDF de manière asynchrone (Celery)
[**ValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPost**](SignatureLectroniqueAPI.md#ValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPost) | **Post** /api/v1/traitement/valider-signature-pdf | Valider les signatures électroniques d&#39;un PDF



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
	resp, r, err := apiClient.SignatureLectroniqueAPI.GenererCertificatTestApiV1TraitementGenererCertificatTestPost(context.Background()).GenerateCertificateRequest(generateCertificateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignatureLectroniqueAPI.GenererCertificatTestApiV1TraitementGenererCertificatTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenererCertificatTestApiV1TraitementGenererCertificatTestPost`: GenerateCertificateResponse
	fmt.Fprintf(os.Stdout, "Response from `SignatureLectroniqueAPI.GenererCertificatTestApiV1TraitementGenererCertificatTestPost`: %v\n", resp)
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
	resp, r, err := apiClient.SignatureLectroniqueAPI.SignerPdfApiV1TraitementSignerPdfPost(context.Background()).FichierPdf(fichierPdf).Raison(raison).Localisation(localisation).Contact(contact).FieldName(fieldName).UsePadesLt(usePadesLt).UseTimestamp(useTimestamp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignatureLectroniqueAPI.SignerPdfApiV1TraitementSignerPdfPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignerPdfApiV1TraitementSignerPdfPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SignatureLectroniqueAPI.SignerPdfApiV1TraitementSignerPdfPost`: %v\n", resp)
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
	resp, r, err := apiClient.SignatureLectroniqueAPI.SignerPdfAsyncApiV1TraitementSignerPdfAsyncPost(context.Background()).FichierPdf(fichierPdf).Raison(raison).Localisation(localisation).Contact(contact).FieldName(fieldName).UsePadesLt(usePadesLt).UseTimestamp(useTimestamp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignatureLectroniqueAPI.SignerPdfAsyncApiV1TraitementSignerPdfAsyncPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignerPdfAsyncApiV1TraitementSignerPdfAsyncPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SignatureLectroniqueAPI.SignerPdfAsyncApiV1TraitementSignerPdfAsyncPost`: %v\n", resp)
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
	resp, r, err := apiClient.SignatureLectroniqueAPI.ValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPost(context.Background()).FichierPdf(fichierPdf).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignatureLectroniqueAPI.ValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SignatureLectroniqueAPI.ValiderSignaturePdfEndpointApiV1TraitementValiderSignaturePdfPost`: %v\n", resp)
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


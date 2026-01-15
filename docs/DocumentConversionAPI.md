# \DocumentConversionAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConvertDocumentAsyncApiV1ConvertAsyncPost**](DocumentConversionAPI.md#ConvertDocumentAsyncApiV1ConvertAsyncPost) | **Post** /api/v1/convert/async | Convertir un document en Factur-X (mode asynchrone)
[**DownloadFileApiV1ConvertConversionIdDownloadFilenameGet**](DocumentConversionAPI.md#DownloadFileApiV1ConvertConversionIdDownloadFilenameGet) | **Get** /api/v1/convert/{conversion_id}/download/{filename} | Télécharger un fichier généré
[**GetConversionStatusApiV1ConvertConversionIdStatusGet**](DocumentConversionAPI.md#GetConversionStatusApiV1ConvertConversionIdStatusGet) | **Get** /api/v1/convert/{conversion_id}/status | Vérifier le statut d&#39;une conversion
[**ResumeConversionApiV1ConvertConversionIdResumePost**](DocumentConversionAPI.md#ResumeConversionApiV1ConvertConversionIdResumePost) | **Post** /api/v1/convert/{conversion_id}/resume | Reprendre une conversion avec corrections



## ConvertDocumentAsyncApiV1ConvertAsyncPost

> interface{} ConvertDocumentAsyncApiV1ConvertAsyncPost(ctx).File(file).Output(output).CallbackUrl(callbackUrl).WebhookMode(webhookMode).Execute()

Convertir un document en Factur-X (mode asynchrone)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/factpulse/sdk-go/v3"
)

func main() {
	file := os.NewFile(1234, "some_file") // *os.File | Document à convertir (PDF, DOCX, XLSX, JPG, PNG)
	output := "output_example" // string | Format de sortie: pdf, xml, both (optional) (default to "pdf")
	callbackUrl := "callbackUrl_example" // string |  (optional)
	webhookMode := "webhookMode_example" // string | Mode de livraison du contenu: 'inline' (base64 dans webhook) ou 'download_url' (URL temporaire 1h) (optional) (default to "inline")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentConversionAPI.ConvertDocumentAsyncApiV1ConvertAsyncPost(context.Background()).File(file).Output(output).CallbackUrl(callbackUrl).WebhookMode(webhookMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentConversionAPI.ConvertDocumentAsyncApiV1ConvertAsyncPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConvertDocumentAsyncApiV1ConvertAsyncPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DocumentConversionAPI.ConvertDocumentAsyncApiV1ConvertAsyncPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConvertDocumentAsyncApiV1ConvertAsyncPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | Document à convertir (PDF, DOCX, XLSX, JPG, PNG) | 
 **output** | **string** | Format de sortie: pdf, xml, both | [default to &quot;pdf&quot;]
 **callbackUrl** | **string** |  | 
 **webhookMode** | **string** | Mode de livraison du contenu: &#39;inline&#39; (base64 dans webhook) ou &#39;download_url&#39; (URL temporaire 1h) | [default to &quot;inline&quot;]

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


## DownloadFileApiV1ConvertConversionIdDownloadFilenameGet

> interface{} DownloadFileApiV1ConvertConversionIdDownloadFilenameGet(ctx, conversionId, filename).Execute()

Télécharger un fichier généré



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/factpulse/sdk-go/v3"
)

func main() {
	conversionId := "conversionId_example" // string | Conversion ID returned by POST /convert (UUID format)
	filename := "filename_example" // string | File to download: 'facturx.pdf' or 'facturx.xml'

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentConversionAPI.DownloadFileApiV1ConvertConversionIdDownloadFilenameGet(context.Background(), conversionId, filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentConversionAPI.DownloadFileApiV1ConvertConversionIdDownloadFilenameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadFileApiV1ConvertConversionIdDownloadFilenameGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DocumentConversionAPI.DownloadFileApiV1ConvertConversionIdDownloadFilenameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversionId** | **string** | Conversion ID returned by POST /convert (UUID format) | 
**filename** | **string** | File to download: &#39;facturx.pdf&#39; or &#39;facturx.xml&#39; | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadFileApiV1ConvertConversionIdDownloadFilenameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversionStatusApiV1ConvertConversionIdStatusGet

> map[string]interface{} GetConversionStatusApiV1ConvertConversionIdStatusGet(ctx, conversionId).Execute()

Vérifier le statut d'une conversion



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/factpulse/sdk-go/v3"
)

func main() {
	conversionId := "conversionId_example" // string | Conversion ID returned by POST /convert (UUID format)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentConversionAPI.GetConversionStatusApiV1ConvertConversionIdStatusGet(context.Background(), conversionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentConversionAPI.GetConversionStatusApiV1ConvertConversionIdStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversionStatusApiV1ConvertConversionIdStatusGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DocumentConversionAPI.GetConversionStatusApiV1ConvertConversionIdStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversionId** | **string** | Conversion ID returned by POST /convert (UUID format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversionStatusApiV1ConvertConversionIdStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeConversionApiV1ConvertConversionIdResumePost

> ConvertSuccessResponse ResumeConversionApiV1ConvertConversionIdResumePost(ctx, conversionId).ConvertResumeRequest(convertResumeRequest).Execute()

Reprendre une conversion avec corrections



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/factpulse/sdk-go/v3"
)

func main() {
	conversionId := "conversionId_example" // string | Conversion ID returned by POST /convert (UUID format)
	convertResumeRequest := *openapiclient.NewConvertResumeRequest() // ConvertResumeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentConversionAPI.ResumeConversionApiV1ConvertConversionIdResumePost(context.Background(), conversionId).ConvertResumeRequest(convertResumeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentConversionAPI.ResumeConversionApiV1ConvertConversionIdResumePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResumeConversionApiV1ConvertConversionIdResumePost`: ConvertSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `DocumentConversionAPI.ResumeConversionApiV1ConvertConversionIdResumePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversionId** | **string** | Conversion ID returned by POST /convert (UUID format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeConversionApiV1ConvertConversionIdResumePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **convertResumeRequest** | [**ConvertResumeRequest**](ConvertResumeRequest.md) |  | 

### Return type

[**ConvertSuccessResponse**](ConvertSuccessResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


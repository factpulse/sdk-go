# \DocumentConversionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConvertDocumentApiV1ConvertPost**](DocumentConversionAPI.md#ConvertDocumentApiV1ConvertPost) | **Post** /api/v1/convert | Convertir un document en Factur-X
[**ConvertDocumentAsyncApiV1ConvertAsyncPost**](DocumentConversionAPI.md#ConvertDocumentAsyncApiV1ConvertAsyncPost) | **Post** /api/v1/convert/async | Convertir un document en Factur-X (mode asynchrone)
[**DownloadFileApiV1ConvertConversionIdDownloadFilenameGet**](DocumentConversionAPI.md#DownloadFileApiV1ConvertConversionIdDownloadFilenameGet) | **Get** /api/v1/convert/{conversion_id}/download/{filename} | Télécharger un fichier généré
[**GetConversionStatusApiV1ConvertConversionIdStatusGet**](DocumentConversionAPI.md#GetConversionStatusApiV1ConvertConversionIdStatusGet) | **Get** /api/v1/convert/{conversion_id}/status | Vérifier le statut d&#39;une conversion
[**ResumeConversionApiV1ConvertConversionIdResumePost**](DocumentConversionAPI.md#ResumeConversionApiV1ConvertConversionIdResumePost) | **Post** /api/v1/convert/{conversion_id}/resume | Reprendre une conversion avec corrections



## ConvertDocumentApiV1ConvertPost

> ConvertSuccessResponse ConvertDocumentApiV1ConvertPost(ctx).File(file).Output(output).CallbackUrl(callbackUrl).Execute()

Convertir un document en Factur-X



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
	file := os.NewFile(1234, "some_file") // *os.File | Document à convertir (PDF, DOCX, XLSX, JPG, PNG)
	output := "output_example" // string | Format de sortie: pdf, xml, both (optional) (default to "pdf")
	callbackUrl := "callbackUrl_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentConversionAPI.ConvertDocumentApiV1ConvertPost(context.Background()).File(file).Output(output).CallbackUrl(callbackUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentConversionAPI.ConvertDocumentApiV1ConvertPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConvertDocumentApiV1ConvertPost`: ConvertSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `DocumentConversionAPI.ConvertDocumentApiV1ConvertPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConvertDocumentApiV1ConvertPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | Document à convertir (PDF, DOCX, XLSX, JPG, PNG) | 
 **output** | **string** | Format de sortie: pdf, xml, both | [default to &quot;pdf&quot;]
 **callbackUrl** | **string** |  | 

### Return type

[**ConvertSuccessResponse**](ConvertSuccessResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConvertDocumentAsyncApiV1ConvertAsyncPost

> interface{} ConvertDocumentAsyncApiV1ConvertAsyncPost(ctx).File(file).Output(output).CallbackUrl(callbackUrl).Execute()

Convertir un document en Factur-X (mode asynchrone)



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
	file := os.NewFile(1234, "some_file") // *os.File | Document à convertir (PDF, DOCX, XLSX, JPG, PNG)
	output := "output_example" // string | Format de sortie: pdf, xml, both (optional) (default to "pdf")
	callbackUrl := "callbackUrl_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentConversionAPI.ConvertDocumentAsyncApiV1ConvertAsyncPost(context.Background()).File(file).Output(output).CallbackUrl(callbackUrl).Execute()
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	conversionId := "conversionId_example" // string | 
	filename := "filename_example" // string | 

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
**conversionId** | **string** |  | 
**filename** | **string** |  | 

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	conversionId := "conversionId_example" // string | 

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
**conversionId** | **string** |  | 

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	conversionId := "conversionId_example" // string | 
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
**conversionId** | **string** |  | 

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


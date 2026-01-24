# \FacturXConversionAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConvertDocumentAsyncApiV1ConvertAsyncPost**](FacturXConversionAPI.md#ConvertDocumentAsyncApiV1ConvertAsyncPost) | **Post** /api/v1/convert/async | Convert a document to Factur-X (async mode)
[**DownloadFileApiV1ConvertConversionIdDownloadFilenameGet**](FacturXConversionAPI.md#DownloadFileApiV1ConvertConversionIdDownloadFilenameGet) | **Get** /api/v1/convert/{conversion_id}/download/{filename} | Download a generated file
[**GetConversionStatusApiV1ConvertConversionIdStatusGet**](FacturXConversionAPI.md#GetConversionStatusApiV1ConvertConversionIdStatusGet) | **Get** /api/v1/convert/{conversion_id}/status | Check conversion status
[**ResumeConversionApiV1ConvertConversionIdResumePost**](FacturXConversionAPI.md#ResumeConversionApiV1ConvertConversionIdResumePost) | **Post** /api/v1/convert/{conversion_id}/resume | Resume a conversion with corrections



## ConvertDocumentAsyncApiV1ConvertAsyncPost

> interface{} ConvertDocumentAsyncApiV1ConvertAsyncPost(ctx).File(file).Output(output).CallbackUrl(callbackUrl).WebhookMode(webhookMode).Execute()

Convert a document to Factur-X (async mode)



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
	file := os.NewFile(1234, "some_file") // *os.File | Document to convert (PDF, DOCX, XLSX, JPG, PNG)
	output := "output_example" // string | Output format: pdf, xml, both (optional) (default to "pdf")
	callbackUrl := "callbackUrl_example" // string |  (optional)
	webhookMode := "webhookMode_example" // string | Content delivery mode: 'inline' (base64 in webhook) or 'download_url' (temporary URL, 1h TTL) (optional) (default to "inline")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FacturXConversionAPI.ConvertDocumentAsyncApiV1ConvertAsyncPost(context.Background()).File(file).Output(output).CallbackUrl(callbackUrl).WebhookMode(webhookMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FacturXConversionAPI.ConvertDocumentAsyncApiV1ConvertAsyncPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConvertDocumentAsyncApiV1ConvertAsyncPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FacturXConversionAPI.ConvertDocumentAsyncApiV1ConvertAsyncPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConvertDocumentAsyncApiV1ConvertAsyncPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | Document to convert (PDF, DOCX, XLSX, JPG, PNG) | 
 **output** | **string** | Output format: pdf, xml, both | [default to &quot;pdf&quot;]
 **callbackUrl** | **string** |  | 
 **webhookMode** | **string** | Content delivery mode: &#39;inline&#39; (base64 in webhook) or &#39;download_url&#39; (temporary URL, 1h TTL) | [default to &quot;inline&quot;]

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

Download a generated file



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
	resp, r, err := apiClient.FacturXConversionAPI.DownloadFileApiV1ConvertConversionIdDownloadFilenameGet(context.Background(), conversionId, filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FacturXConversionAPI.DownloadFileApiV1ConvertConversionIdDownloadFilenameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadFileApiV1ConvertConversionIdDownloadFilenameGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FacturXConversionAPI.DownloadFileApiV1ConvertConversionIdDownloadFilenameGet`: %v\n", resp)
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

Check conversion status



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
	resp, r, err := apiClient.FacturXConversionAPI.GetConversionStatusApiV1ConvertConversionIdStatusGet(context.Background(), conversionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FacturXConversionAPI.GetConversionStatusApiV1ConvertConversionIdStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversionStatusApiV1ConvertConversionIdStatusGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FacturXConversionAPI.GetConversionStatusApiV1ConvertConversionIdStatusGet`: %v\n", resp)
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

Resume a conversion with corrections



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
	resp, r, err := apiClient.FacturXConversionAPI.ResumeConversionApiV1ConvertConversionIdResumePost(context.Background(), conversionId).ConvertResumeRequest(convertResumeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FacturXConversionAPI.ResumeConversionApiV1ConvertConversionIdResumePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResumeConversionApiV1ConvertConversionIdResumePost`: ConvertSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `FacturXConversionAPI.ResumeConversionApiV1ConvertConversionIdResumePost`: %v\n", resp)
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


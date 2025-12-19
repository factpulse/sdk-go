# \PDFXMLVerificationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVerificationStatusApiV1VerificationVerifyAsyncTaskIdStatusGet**](PDFXMLVerificationAPI.md#GetVerificationStatusApiV1VerificationVerifyAsyncTaskIdStatusGet) | **Get** /api/v1/verification/verify-async/{task_id}/status | Get status of an asynchronous verification
[**GetVerificationStatusApiV1VerificationVerifyAsyncTaskIdStatusGet_0**](PDFXMLVerificationAPI.md#GetVerificationStatusApiV1VerificationVerifyAsyncTaskIdStatusGet_0) | **Get** /api/v1/verification/verify-async/{task_id}/status | Get status of an asynchronous verification
[**VerifyPdfAsyncApiV1VerificationVerifyAsyncPost**](PDFXMLVerificationAPI.md#VerifyPdfAsyncApiV1VerificationVerifyAsyncPost) | **Post** /api/v1/verification/verify-async | Verify PDF/XML Factur-X compliance (asynchronous)
[**VerifyPdfAsyncApiV1VerificationVerifyAsyncPost_0**](PDFXMLVerificationAPI.md#VerifyPdfAsyncApiV1VerificationVerifyAsyncPost_0) | **Post** /api/v1/verification/verify-async | Verify PDF/XML Factur-X compliance (asynchronous)
[**VerifyPdfSyncApiV1VerificationVerifyPost**](PDFXMLVerificationAPI.md#VerifyPdfSyncApiV1VerificationVerifyPost) | **Post** /api/v1/verification/verify | Verify PDF/XML Factur-X compliance (synchronous)
[**VerifyPdfSyncApiV1VerificationVerifyPost_0**](PDFXMLVerificationAPI.md#VerifyPdfSyncApiV1VerificationVerifyPost_0) | **Post** /api/v1/verification/verify | Verify PDF/XML Factur-X compliance (synchronous)



## GetVerificationStatusApiV1VerificationVerifyAsyncTaskIdStatusGet

> TaskStatus GetVerificationStatusApiV1VerificationVerifyAsyncTaskIdStatusGet(ctx, taskId).Execute()

Get status of an asynchronous verification



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
	taskId := "taskId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFXMLVerificationAPI.GetVerificationStatusApiV1VerificationVerifyAsyncTaskIdStatusGet(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFXMLVerificationAPI.GetVerificationStatusApiV1VerificationVerifyAsyncTaskIdStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVerificationStatusApiV1VerificationVerifyAsyncTaskIdStatusGet`: TaskStatus
	fmt.Fprintf(os.Stdout, "Response from `PDFXMLVerificationAPI.GetVerificationStatusApiV1VerificationVerifyAsyncTaskIdStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVerificationStatusApiV1VerificationVerifyAsyncTaskIdStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskStatus**](TaskStatus.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVerificationStatusApiV1VerificationVerifyAsyncTaskIdStatusGet_0

> TaskStatus GetVerificationStatusApiV1VerificationVerifyAsyncTaskIdStatusGet_0(ctx, taskId).Execute()

Get status of an asynchronous verification



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
	taskId := "taskId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFXMLVerificationAPI.GetVerificationStatusApiV1VerificationVerifyAsyncTaskIdStatusGet_0(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFXMLVerificationAPI.GetVerificationStatusApiV1VerificationVerifyAsyncTaskIdStatusGet_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVerificationStatusApiV1VerificationVerifyAsyncTaskIdStatusGet_0`: TaskStatus
	fmt.Fprintf(os.Stdout, "Response from `PDFXMLVerificationAPI.GetVerificationStatusApiV1VerificationVerifyAsyncTaskIdStatusGet_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVerificationStatusApiV1VerificationVerifyAsyncTaskIdStatusGet_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskStatus**](TaskStatus.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyPdfAsyncApiV1VerificationVerifyAsyncPost

> TaskResponse VerifyPdfAsyncApiV1VerificationVerifyAsyncPost(ctx).PdfFile(pdfFile).ForceOcr(forceOcr).Execute()

Verify PDF/XML Factur-X compliance (asynchronous)



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
	pdfFile := os.NewFile(1234, "some_file") // *os.File | Factur-X PDF file to verify
	forceOcr := true // bool | Force OCR usage even if PDF contains native text (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFXMLVerificationAPI.VerifyPdfAsyncApiV1VerificationVerifyAsyncPost(context.Background()).PdfFile(pdfFile).ForceOcr(forceOcr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFXMLVerificationAPI.VerifyPdfAsyncApiV1VerificationVerifyAsyncPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyPdfAsyncApiV1VerificationVerifyAsyncPost`: TaskResponse
	fmt.Fprintf(os.Stdout, "Response from `PDFXMLVerificationAPI.VerifyPdfAsyncApiV1VerificationVerifyAsyncPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyPdfAsyncApiV1VerificationVerifyAsyncPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pdfFile** | ***os.File** | Factur-X PDF file to verify | 
 **forceOcr** | **bool** | Force OCR usage even if PDF contains native text | [default to false]

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyPdfAsyncApiV1VerificationVerifyAsyncPost_0

> TaskResponse VerifyPdfAsyncApiV1VerificationVerifyAsyncPost_0(ctx).PdfFile(pdfFile).ForceOcr(forceOcr).Execute()

Verify PDF/XML Factur-X compliance (asynchronous)



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
	pdfFile := os.NewFile(1234, "some_file") // *os.File | Factur-X PDF file to verify
	forceOcr := true // bool | Force OCR usage even if PDF contains native text (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFXMLVerificationAPI.VerifyPdfAsyncApiV1VerificationVerifyAsyncPost_0(context.Background()).PdfFile(pdfFile).ForceOcr(forceOcr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFXMLVerificationAPI.VerifyPdfAsyncApiV1VerificationVerifyAsyncPost_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyPdfAsyncApiV1VerificationVerifyAsyncPost_0`: TaskResponse
	fmt.Fprintf(os.Stdout, "Response from `PDFXMLVerificationAPI.VerifyPdfAsyncApiV1VerificationVerifyAsyncPost_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyPdfAsyncApiV1VerificationVerifyAsyncPost_2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pdfFile** | ***os.File** | Factur-X PDF file to verify | 
 **forceOcr** | **bool** | Force OCR usage even if PDF contains native text | [default to false]

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyPdfSyncApiV1VerificationVerifyPost

> VerificationSuccessResponse VerifyPdfSyncApiV1VerificationVerifyPost(ctx).PdfFile(pdfFile).Execute()

Verify PDF/XML Factur-X compliance (synchronous)



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
	pdfFile := os.NewFile(1234, "some_file") // *os.File | Factur-X PDF file to verify

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFXMLVerificationAPI.VerifyPdfSyncApiV1VerificationVerifyPost(context.Background()).PdfFile(pdfFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFXMLVerificationAPI.VerifyPdfSyncApiV1VerificationVerifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyPdfSyncApiV1VerificationVerifyPost`: VerificationSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `PDFXMLVerificationAPI.VerifyPdfSyncApiV1VerificationVerifyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyPdfSyncApiV1VerificationVerifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pdfFile** | ***os.File** | Factur-X PDF file to verify | 

### Return type

[**VerificationSuccessResponse**](VerificationSuccessResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyPdfSyncApiV1VerificationVerifyPost_0

> VerificationSuccessResponse VerifyPdfSyncApiV1VerificationVerifyPost_0(ctx).PdfFile(pdfFile).Execute()

Verify PDF/XML Factur-X compliance (synchronous)



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
	pdfFile := os.NewFile(1234, "some_file") // *os.File | Factur-X PDF file to verify

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PDFXMLVerificationAPI.VerifyPdfSyncApiV1VerificationVerifyPost_0(context.Background()).PdfFile(pdfFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PDFXMLVerificationAPI.VerifyPdfSyncApiV1VerificationVerifyPost_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyPdfSyncApiV1VerificationVerifyPost_0`: VerificationSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `PDFXMLVerificationAPI.VerifyPdfSyncApiV1VerificationVerifyPost_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyPdfSyncApiV1VerificationVerifyPost_3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pdfFile** | ***os.File** | Factur-X PDF file to verify | 

### Return type

[**VerificationSuccessResponse**](VerificationSuccessResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


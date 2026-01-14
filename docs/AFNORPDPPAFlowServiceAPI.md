# \AFNORPDPPAFlowServiceAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadFlowProxyApiV1AfnorFlowV1FlowsFlowIdGet**](AFNORPDPPAFlowServiceAPI.md#DownloadFlowProxyApiV1AfnorFlowV1FlowsFlowIdGet) | **Get** /api/v1/afnor/flow/v1/flows/{flowId} | Download a flow
[**FlowHealthcheckProxyApiV1AfnorFlowV1HealthcheckGet**](AFNORPDPPAFlowServiceAPI.md#FlowHealthcheckProxyApiV1AfnorFlowV1HealthcheckGet) | **Get** /api/v1/afnor/flow/v1/healthcheck | Healthcheck Flow Service
[**SearchFlowsProxyApiV1AfnorFlowV1FlowsSearchPost**](AFNORPDPPAFlowServiceAPI.md#SearchFlowsProxyApiV1AfnorFlowV1FlowsSearchPost) | **Post** /api/v1/afnor/flow/v1/flows/search | Search flows
[**SubmitFlowProxyApiV1AfnorFlowV1FlowsPost**](AFNORPDPPAFlowServiceAPI.md#SubmitFlowProxyApiV1AfnorFlowV1FlowsPost) | **Post** /api/v1/afnor/flow/v1/flows | Submit an invoicing flow



## DownloadFlowProxyApiV1AfnorFlowV1FlowsFlowIdGet

> AFNORFlow DownloadFlowProxyApiV1AfnorFlowV1FlowsFlowIdGet(ctx, flowId).DocType(docType).Execute()

Download a flow



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/factpulse/sdk-go"
)

func main() {
	flowId := "flowId_example" // string | AFNOR flow identifier (UUID)
	docType := openapiclient.DocType("Metadata") // DocType | Type of file to download: Metadata (default, JSON), Original, Converted, or ReadableView (optional) (default to "Metadata")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPAFlowServiceAPI.DownloadFlowProxyApiV1AfnorFlowV1FlowsFlowIdGet(context.Background(), flowId).DocType(docType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPAFlowServiceAPI.DownloadFlowProxyApiV1AfnorFlowV1FlowsFlowIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadFlowProxyApiV1AfnorFlowV1FlowsFlowIdGet`: AFNORFlow
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPAFlowServiceAPI.DownloadFlowProxyApiV1AfnorFlowV1FlowsFlowIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | AFNOR flow identifier (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadFlowProxyApiV1AfnorFlowV1FlowsFlowIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **docType** | [**DocType**](DocType.md) | Type of file to download: Metadata (default, JSON), Original, Converted, or ReadableView | [default to &quot;Metadata&quot;]

### Return type

[**AFNORFlow**](AFNORFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/pdf, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowHealthcheckProxyApiV1AfnorFlowV1HealthcheckGet

> map[string]interface{} FlowHealthcheckProxyApiV1AfnorFlowV1HealthcheckGet(ctx).Execute()

Healthcheck Flow Service



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/factpulse/sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPAFlowServiceAPI.FlowHealthcheckProxyApiV1AfnorFlowV1HealthcheckGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPAFlowServiceAPI.FlowHealthcheckProxyApiV1AfnorFlowV1HealthcheckGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowHealthcheckProxyApiV1AfnorFlowV1HealthcheckGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPAFlowServiceAPI.FlowHealthcheckProxyApiV1AfnorFlowV1HealthcheckGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFlowHealthcheckProxyApiV1AfnorFlowV1HealthcheckGetRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchFlowsProxyApiV1AfnorFlowV1FlowsSearchPost

> AFNORSearchFlowContent SearchFlowsProxyApiV1AfnorFlowV1FlowsSearchPost(ctx).AFNORSearchFlowParams(aFNORSearchFlowParams).Execute()

Search flows



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/factpulse/sdk-go"
)

func main() {
	aFNORSearchFlowParams := *openapiclient.NewAFNORSearchFlowParams(*openapiclient.NewAFNORSearchFlowFilters()) // AFNORSearchFlowParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPAFlowServiceAPI.SearchFlowsProxyApiV1AfnorFlowV1FlowsSearchPost(context.Background()).AFNORSearchFlowParams(aFNORSearchFlowParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPAFlowServiceAPI.SearchFlowsProxyApiV1AfnorFlowV1FlowsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchFlowsProxyApiV1AfnorFlowV1FlowsSearchPost`: AFNORSearchFlowContent
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPAFlowServiceAPI.SearchFlowsProxyApiV1AfnorFlowV1FlowsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchFlowsProxyApiV1AfnorFlowV1FlowsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aFNORSearchFlowParams** | [**AFNORSearchFlowParams**](AFNORSearchFlowParams.md) |  | 

### Return type

[**AFNORSearchFlowContent**](AFNORSearchFlowContent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitFlowProxyApiV1AfnorFlowV1FlowsPost

> interface{} SubmitFlowProxyApiV1AfnorFlowV1FlowsPost(ctx).FlowInfo(flowInfo).File(file).Execute()

Submit an invoicing flow



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/factpulse/sdk-go"
)

func main() {
	flowInfo := *openapiclient.NewAFNORFlowInfo(openapiclient.AFNOR_FlowSyntax("CII")) // AFNORFlowInfo | 
	file := os.NewFile(1234, "some_file") // *os.File | Flow file (PDF/A-3 with embedded XML or XML)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPAFlowServiceAPI.SubmitFlowProxyApiV1AfnorFlowV1FlowsPost(context.Background()).FlowInfo(flowInfo).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPAFlowServiceAPI.SubmitFlowProxyApiV1AfnorFlowV1FlowsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitFlowProxyApiV1AfnorFlowV1FlowsPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPAFlowServiceAPI.SubmitFlowProxyApiV1AfnorFlowV1FlowsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitFlowProxyApiV1AfnorFlowV1FlowsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowInfo** | [**AFNORFlowInfo**](AFNORFlowInfo.md) |  | 
 **file** | ***os.File** | Flow file (PDF/A-3 with embedded XML or XML) | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


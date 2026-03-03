# \AFNORPDPPAFlowServiceAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhookProxyApiV1AfnorFlowV1WebhooksPost**](AFNORPDPPAFlowServiceAPI.md#CreateWebhookProxyApiV1AfnorFlowV1WebhooksPost) | **Post** /api/v1/afnor/flow/v1/webhooks | Create a webhook
[**DeleteWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidDelete**](AFNORPDPPAFlowServiceAPI.md#DeleteWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidDelete) | **Delete** /api/v1/afnor/flow/v1/webhooks/{webhookUid} | Delete a webhook
[**DownloadFlowProxyApiV1AfnorFlowV1FlowsFlowIdGet**](AFNORPDPPAFlowServiceAPI.md#DownloadFlowProxyApiV1AfnorFlowV1FlowsFlowIdGet) | **Get** /api/v1/afnor/flow/v1/flows/{flowId} | Download a flow
[**FlowHealthcheckProxyApiV1AfnorFlowV1HealthcheckGet**](AFNORPDPPAFlowServiceAPI.md#FlowHealthcheckProxyApiV1AfnorFlowV1HealthcheckGet) | **Get** /api/v1/afnor/flow/v1/healthcheck | Healthcheck Flow Service
[**GetWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidGet**](AFNORPDPPAFlowServiceAPI.md#GetWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidGet) | **Get** /api/v1/afnor/flow/v1/webhooks/{webhookUid} | Get a webhook
[**ListWebhooksProxyApiV1AfnorFlowV1WebhooksGet**](AFNORPDPPAFlowServiceAPI.md#ListWebhooksProxyApiV1AfnorFlowV1WebhooksGet) | **Get** /api/v1/afnor/flow/v1/webhooks | List webhooks
[**SearchFlowsProxyApiV1AfnorFlowV1FlowsSearchPost**](AFNORPDPPAFlowServiceAPI.md#SearchFlowsProxyApiV1AfnorFlowV1FlowsSearchPost) | **Post** /api/v1/afnor/flow/v1/flows/search | Search flows
[**SubmitFlowProxyApiV1AfnorFlowV1FlowsPost**](AFNORPDPPAFlowServiceAPI.md#SubmitFlowProxyApiV1AfnorFlowV1FlowsPost) | **Post** /api/v1/afnor/flow/v1/flows | Submit an invoicing flow
[**UpdateWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidPatch**](AFNORPDPPAFlowServiceAPI.md#UpdateWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidPatch) | **Patch** /api/v1/afnor/flow/v1/webhooks/{webhookUid} | Update a webhook



## CreateWebhookProxyApiV1AfnorFlowV1WebhooksPost

> interface{} CreateWebhookProxyApiV1AfnorFlowV1WebhooksPost(ctx).Execute()

Create a webhook



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPAFlowServiceAPI.CreateWebhookProxyApiV1AfnorFlowV1WebhooksPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPAFlowServiceAPI.CreateWebhookProxyApiV1AfnorFlowV1WebhooksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhookProxyApiV1AfnorFlowV1WebhooksPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPAFlowServiceAPI.CreateWebhookProxyApiV1AfnorFlowV1WebhooksPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookProxyApiV1AfnorFlowV1WebhooksPostRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidDelete

> interface{} DeleteWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidDelete(ctx, webhookUid).Execute()

Delete a webhook



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
	webhookUid := "webhookUid_example" // string | Webhook unique identifier (UUID)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPAFlowServiceAPI.DeleteWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidDelete(context.Background(), webhookUid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPAFlowServiceAPI.DeleteWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPAFlowServiceAPI.DeleteWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookUid** | **string** | Webhook unique identifier (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "github.com/factpulse/sdk-go/v3"
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
	openapiclient "github.com/factpulse/sdk-go/v3"
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


## GetWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidGet

> AFNORWebhook GetWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidGet(ctx, webhookUid).Execute()

Get a webhook



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
	webhookUid := "webhookUid_example" // string | Webhook unique identifier (UUID)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPAFlowServiceAPI.GetWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidGet(context.Background(), webhookUid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPAFlowServiceAPI.GetWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidGet`: AFNORWebhook
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPAFlowServiceAPI.GetWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookUid** | **string** | Webhook unique identifier (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AFNORWebhook**](AFNORWebhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhooksProxyApiV1AfnorFlowV1WebhooksGet

> []AFNORWebhook ListWebhooksProxyApiV1AfnorFlowV1WebhooksGet(ctx).Execute()

List webhooks



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPAFlowServiceAPI.ListWebhooksProxyApiV1AfnorFlowV1WebhooksGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPAFlowServiceAPI.ListWebhooksProxyApiV1AfnorFlowV1WebhooksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWebhooksProxyApiV1AfnorFlowV1WebhooksGet`: []AFNORWebhook
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPAFlowServiceAPI.ListWebhooksProxyApiV1AfnorFlowV1WebhooksGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWebhooksProxyApiV1AfnorFlowV1WebhooksGetRequest struct via the builder pattern


### Return type

[**[]AFNORWebhook**](AFNORWebhook.md)

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
	openapiclient "github.com/factpulse/sdk-go/v3"
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
	openapiclient "github.com/factpulse/sdk-go/v3"
)

func main() {
	flowInfo := *openapiclient.NewAFNORFlowInfo("Name_example", openapiclient.AFNOR_FlowSyntax("CII")) // AFNORFlowInfo | 
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


## UpdateWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidPatch

> AFNORWebhook UpdateWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidPatch(ctx, webhookUid).Execute()

Update a webhook



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
	webhookUid := "webhookUid_example" // string | Webhook unique identifier (UUID)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPAFlowServiceAPI.UpdateWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidPatch(context.Background(), webhookUid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPAFlowServiceAPI.UpdateWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidPatch`: AFNORWebhook
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPAFlowServiceAPI.UpdateWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookUid** | **string** | Webhook unique identifier (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookProxyApiV1AfnorFlowV1WebhooksWebhookUidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AFNORWebhook**](AFNORWebhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


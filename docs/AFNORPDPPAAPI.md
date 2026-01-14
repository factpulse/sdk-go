# \AFNORPDPPAAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAfnorCredentialsApiV1AfnorCredentialsGet**](AFNORPDPPAAPI.md#GetAfnorCredentialsApiV1AfnorCredentialsGet) | **Get** /api/v1/afnor/credentials | Retrieve stored AFNOR credentials
[**GetFluxEntrantApiV1AfnorIncomingFlowsFlowIdGet**](AFNORPDPPAAPI.md#GetFluxEntrantApiV1AfnorIncomingFlowsFlowIdGet) | **Get** /api/v1/afnor/incoming-flows/{flow_id} | Retrieve and extract an incoming invoice
[**OauthTokenProxyApiV1AfnorOauthTokenPost**](AFNORPDPPAAPI.md#OauthTokenProxyApiV1AfnorOauthTokenPost) | **Post** /api/v1/afnor/oauth/token | OAuth2 endpoint for AFNOR authentication



## GetAfnorCredentialsApiV1AfnorCredentialsGet

> interface{} GetAfnorCredentialsApiV1AfnorCredentialsGet(ctx).Execute()

Retrieve stored AFNOR credentials



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
	resp, r, err := apiClient.AFNORPDPPAAPI.GetAfnorCredentialsApiV1AfnorCredentialsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPAAPI.GetAfnorCredentialsApiV1AfnorCredentialsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAfnorCredentialsApiV1AfnorCredentialsGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPAAPI.GetAfnorCredentialsApiV1AfnorCredentialsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAfnorCredentialsApiV1AfnorCredentialsGetRequest struct via the builder pattern


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


## GetFluxEntrantApiV1AfnorIncomingFlowsFlowIdGet

> IncomingInvoice GetFluxEntrantApiV1AfnorIncomingFlowsFlowIdGet(ctx, flowId).IncludeDocument(includeDocument).Execute()

Retrieve and extract an incoming invoice



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
	flowId := "flowId_example" // string | AFNOR flow ID (UUID format)
	includeDocument := true // bool | Include base64-encoded document in response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPAAPI.GetFluxEntrantApiV1AfnorIncomingFlowsFlowIdGet(context.Background(), flowId).IncludeDocument(includeDocument).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPAAPI.GetFluxEntrantApiV1AfnorIncomingFlowsFlowIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFluxEntrantApiV1AfnorIncomingFlowsFlowIdGet`: IncomingInvoice
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPAAPI.GetFluxEntrantApiV1AfnorIncomingFlowsFlowIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | AFNOR flow ID (UUID format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFluxEntrantApiV1AfnorIncomingFlowsFlowIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDocument** | **bool** | Include base64-encoded document in response | [default to false]

### Return type

[**IncomingInvoice**](IncomingInvoice.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OauthTokenProxyApiV1AfnorOauthTokenPost

> interface{} OauthTokenProxyApiV1AfnorOauthTokenPost(ctx).Execute()

OAuth2 endpoint for AFNOR authentication



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
	resp, r, err := apiClient.AFNORPDPPAAPI.OauthTokenProxyApiV1AfnorOauthTokenPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPAAPI.OauthTokenProxyApiV1AfnorOauthTokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OauthTokenProxyApiV1AfnorOauthTokenPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPAAPI.OauthTokenProxyApiV1AfnorOauthTokenPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOauthTokenProxyApiV1AfnorOauthTokenPostRequest struct via the builder pattern


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


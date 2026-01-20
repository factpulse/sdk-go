# \ClientManagementAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateClientApiV1ClientsUidActiverPost**](ClientManagementAPI.md#ActivateClientApiV1ClientsUidActiverPost) | **Post** /api/v1/clients/{uid}/activer | Activate a client
[**CreateClientApiV1ClientsPost**](ClientManagementAPI.md#CreateClientApiV1ClientsPost) | **Post** /api/v1/clients | Create a client
[**DeactivateClientApiV1ClientsUidDesactiverPost**](ClientManagementAPI.md#DeactivateClientApiV1ClientsUidDesactiverPost) | **Post** /api/v1/clients/{uid}/desactiver | Deactivate a client
[**GetClientApiV1ClientsUidGet**](ClientManagementAPI.md#GetClientApiV1ClientsUidGet) | **Get** /api/v1/clients/{uid} | Get client details
[**GetPdpConfigApiV1ClientsUidPdpConfigGet**](ClientManagementAPI.md#GetPdpConfigApiV1ClientsUidPdpConfigGet) | **Get** /api/v1/clients/{uid}/pdp-config | Get client PDP configuration
[**ListClientsApiV1ClientsGet**](ClientManagementAPI.md#ListClientsApiV1ClientsGet) | **Get** /api/v1/clients | List clients
[**UpdateClientApiV1ClientsUidPatch**](ClientManagementAPI.md#UpdateClientApiV1ClientsUidPatch) | **Patch** /api/v1/clients/{uid} | Update a client
[**UpdatePdpConfigApiV1ClientsUidPdpConfigPut**](ClientManagementAPI.md#UpdatePdpConfigApiV1ClientsUidPdpConfigPut) | **Put** /api/v1/clients/{uid}/pdp-config | Configure client PDP



## ActivateClientApiV1ClientsUidActiverPost

> ClientActivateResponse ActivateClientApiV1ClientsUidActiverPost(ctx, uid).Execute()

Activate a client



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.ActivateClientApiV1ClientsUidActiverPost(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.ActivateClientApiV1ClientsUidActiverPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateClientApiV1ClientsUidActiverPost`: ClientActivateResponse
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.ActivateClientApiV1ClientsUidActiverPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateClientApiV1ClientsUidActiverPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientActivateResponse**](ClientActivateResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateClientApiV1ClientsPost

> ClientDetail CreateClientApiV1ClientsPost(ctx).ClientCreateRequest(clientCreateRequest).Execute()

Create a client



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
	clientCreateRequest := *openapiclient.NewClientCreateRequest("Name_example") // ClientCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.CreateClientApiV1ClientsPost(context.Background()).ClientCreateRequest(clientCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.CreateClientApiV1ClientsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateClientApiV1ClientsPost`: ClientDetail
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.CreateClientApiV1ClientsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClientApiV1ClientsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientCreateRequest** | [**ClientCreateRequest**](ClientCreateRequest.md) |  | 

### Return type

[**ClientDetail**](ClientDetail.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateClientApiV1ClientsUidDesactiverPost

> ClientActivateResponse DeactivateClientApiV1ClientsUidDesactiverPost(ctx, uid).Execute()

Deactivate a client



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.DeactivateClientApiV1ClientsUidDesactiverPost(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.DeactivateClientApiV1ClientsUidDesactiverPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeactivateClientApiV1ClientsUidDesactiverPost`: ClientActivateResponse
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.DeactivateClientApiV1ClientsUidDesactiverPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateClientApiV1ClientsUidDesactiverPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientActivateResponse**](ClientActivateResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientApiV1ClientsUidGet

> ClientDetail GetClientApiV1ClientsUidGet(ctx, uid).Execute()

Get client details



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.GetClientApiV1ClientsUidGet(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.GetClientApiV1ClientsUidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientApiV1ClientsUidGet`: ClientDetail
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.GetClientApiV1ClientsUidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientApiV1ClientsUidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientDetail**](ClientDetail.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPdpConfigApiV1ClientsUidPdpConfigGet

> PDPConfigResponse GetPdpConfigApiV1ClientsUidPdpConfigGet(ctx, uid).Execute()

Get client PDP configuration



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.GetPdpConfigApiV1ClientsUidPdpConfigGet(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.GetPdpConfigApiV1ClientsUidPdpConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPdpConfigApiV1ClientsUidPdpConfigGet`: PDPConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.GetPdpConfigApiV1ClientsUidPdpConfigGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPdpConfigApiV1ClientsUidPdpConfigGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PDPConfigResponse**](PDPConfigResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClientsApiV1ClientsGet

> ClientListResponse ListClientsApiV1ClientsGet(ctx).Page(page).PageSize(pageSize).Execute()

List clients



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
	page := int32(56) // int32 | Page number (optional) (default to 1)
	pageSize := int32(56) // int32 | Page size (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.ListClientsApiV1ClientsGet(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.ListClientsApiV1ClientsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClientsApiV1ClientsGet`: ClientListResponse
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.ListClientsApiV1ClientsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListClientsApiV1ClientsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number | [default to 1]
 **pageSize** | **int32** | Page size | [default to 20]

### Return type

[**ClientListResponse**](ClientListResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClientApiV1ClientsUidPatch

> ClientDetail UpdateClientApiV1ClientsUidPatch(ctx, uid).ClientUpdateRequest(clientUpdateRequest).Execute()

Update a client



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	clientUpdateRequest := *openapiclient.NewClientUpdateRequest() // ClientUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.UpdateClientApiV1ClientsUidPatch(context.Background(), uid).ClientUpdateRequest(clientUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.UpdateClientApiV1ClientsUidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateClientApiV1ClientsUidPatch`: ClientDetail
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.UpdateClientApiV1ClientsUidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientApiV1ClientsUidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientUpdateRequest** | [**ClientUpdateRequest**](ClientUpdateRequest.md) |  | 

### Return type

[**ClientDetail**](ClientDetail.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePdpConfigApiV1ClientsUidPdpConfigPut

> PDPConfigResponse UpdatePdpConfigApiV1ClientsUidPdpConfigPut(ctx, uid).PDPConfigUpdateRequest(pDPConfigUpdateRequest).Execute()

Configure client PDP



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	pDPConfigUpdateRequest := *openapiclient.NewPDPConfigUpdateRequest("FlowServiceUrl_example", "TokenUrl_example", "OauthClientId_example", "ClientSecret_example") // PDPConfigUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.UpdatePdpConfigApiV1ClientsUidPdpConfigPut(context.Background(), uid).PDPConfigUpdateRequest(pDPConfigUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.UpdatePdpConfigApiV1ClientsUidPdpConfigPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePdpConfigApiV1ClientsUidPdpConfigPut`: PDPConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.UpdatePdpConfigApiV1ClientsUidPdpConfigPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePdpConfigApiV1ClientsUidPdpConfigPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pDPConfigUpdateRequest** | [**PDPConfigUpdateRequest**](PDPConfigUpdateRequest.md) |  | 

### Return type

[**PDPConfigResponse**](PDPConfigResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


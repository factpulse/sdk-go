# \ClientManagementAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateClientApiV1ClientsUidActiverPost**](ClientManagementAPI.md#ActivateClientApiV1ClientsUidActiverPost) | **Post** /api/v1/clients/{uid}/activer | Activate a client
[**CreateClientApiV1ClientsPost**](ClientManagementAPI.md#CreateClientApiV1ClientsPost) | **Post** /api/v1/clients | Create a client
[**DeactivateClientApiV1ClientsUidDesactiverPost**](ClientManagementAPI.md#DeactivateClientApiV1ClientsUidDesactiverPost) | **Post** /api/v1/clients/{uid}/desactiver | Deactivate a client
[**DeleteWebhookSecretApiV1ClientsUidWebhookSecretDelete**](ClientManagementAPI.md#DeleteWebhookSecretApiV1ClientsUidWebhookSecretDelete) | **Delete** /api/v1/clients/{uid}/webhook-secret | Delete webhook secret
[**GenerateWebhookSecretApiV1ClientsUidWebhookSecretGeneratePost**](ClientManagementAPI.md#GenerateWebhookSecretApiV1ClientsUidWebhookSecretGeneratePost) | **Post** /api/v1/clients/{uid}/webhook-secret/generate | Generate webhook secret
[**GetClientApiV1ClientsUidGet**](ClientManagementAPI.md#GetClientApiV1ClientsUidGet) | **Get** /api/v1/clients/{uid} | Get client details
[**GetPdpConfigApiV1ClientsUidPdpConfigGet**](ClientManagementAPI.md#GetPdpConfigApiV1ClientsUidPdpConfigGet) | **Get** /api/v1/clients/{uid}/pdp-config | Get client PDP configuration
[**GetWebhookSecretStatusApiV1ClientsUidWebhookSecretStatusGet**](ClientManagementAPI.md#GetWebhookSecretStatusApiV1ClientsUidWebhookSecretStatusGet) | **Get** /api/v1/clients/{uid}/webhook-secret/status | Get webhook secret status
[**ListClientsApiV1ClientsGet**](ClientManagementAPI.md#ListClientsApiV1ClientsGet) | **Get** /api/v1/clients | List clients
[**RotateEncryptionKeyApiV1ClientsUidRotateEncryptionKeyPost**](ClientManagementAPI.md#RotateEncryptionKeyApiV1ClientsUidRotateEncryptionKeyPost) | **Post** /api/v1/clients/{uid}/rotate-encryption-key | Rotate client encryption key
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

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

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

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

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

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhookSecretApiV1ClientsUidWebhookSecretDelete

> WebhookSecretDeleteResponse DeleteWebhookSecretApiV1ClientsUidWebhookSecretDelete(ctx, uid).Execute()

Delete webhook secret



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
	resp, r, err := apiClient.ClientManagementAPI.DeleteWebhookSecretApiV1ClientsUidWebhookSecretDelete(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.DeleteWebhookSecretApiV1ClientsUidWebhookSecretDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWebhookSecretApiV1ClientsUidWebhookSecretDelete`: WebhookSecretDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.DeleteWebhookSecretApiV1ClientsUidWebhookSecretDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookSecretApiV1ClientsUidWebhookSecretDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookSecretDeleteResponse**](WebhookSecretDeleteResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateWebhookSecretApiV1ClientsUidWebhookSecretGeneratePost

> WebhookSecretGenerateResponse GenerateWebhookSecretApiV1ClientsUidWebhookSecretGeneratePost(ctx, uid).Execute()

Generate webhook secret



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
	resp, r, err := apiClient.ClientManagementAPI.GenerateWebhookSecretApiV1ClientsUidWebhookSecretGeneratePost(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.GenerateWebhookSecretApiV1ClientsUidWebhookSecretGeneratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateWebhookSecretApiV1ClientsUidWebhookSecretGeneratePost`: WebhookSecretGenerateResponse
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.GenerateWebhookSecretApiV1ClientsUidWebhookSecretGeneratePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateWebhookSecretApiV1ClientsUidWebhookSecretGeneratePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookSecretGenerateResponse**](WebhookSecretGenerateResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

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

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

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

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookSecretStatusApiV1ClientsUidWebhookSecretStatusGet

> WebhookSecretStatusResponse GetWebhookSecretStatusApiV1ClientsUidWebhookSecretStatusGet(ctx, uid).Execute()

Get webhook secret status



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
	resp, r, err := apiClient.ClientManagementAPI.GetWebhookSecretStatusApiV1ClientsUidWebhookSecretStatusGet(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.GetWebhookSecretStatusApiV1ClientsUidWebhookSecretStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookSecretStatusApiV1ClientsUidWebhookSecretStatusGet`: WebhookSecretStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.GetWebhookSecretStatusApiV1ClientsUidWebhookSecretStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookSecretStatusApiV1ClientsUidWebhookSecretStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookSecretStatusResponse**](WebhookSecretStatusResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

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

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateEncryptionKeyApiV1ClientsUidRotateEncryptionKeyPost

> KeyRotationResponse RotateEncryptionKeyApiV1ClientsUidRotateEncryptionKeyPost(ctx, uid).KeyRotationRequest(keyRotationRequest).Execute()

Rotate client encryption key



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
	keyRotationRequest := *openapiclient.NewKeyRotationRequest("OldKey_example", "NewKey_example") // KeyRotationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.RotateEncryptionKeyApiV1ClientsUidRotateEncryptionKeyPost(context.Background(), uid).KeyRotationRequest(keyRotationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.RotateEncryptionKeyApiV1ClientsUidRotateEncryptionKeyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateEncryptionKeyApiV1ClientsUidRotateEncryptionKeyPost`: KeyRotationResponse
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.RotateEncryptionKeyApiV1ClientsUidRotateEncryptionKeyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateEncryptionKeyApiV1ClientsUidRotateEncryptionKeyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keyRotationRequest** | [**KeyRotationRequest**](KeyRotationRequest.md) |  | 

### Return type

[**KeyRotationResponse**](KeyRotationResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
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

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePdpConfigApiV1ClientsUidPdpConfigPut

> PDPConfigResponse UpdatePdpConfigApiV1ClientsUidPdpConfigPut(ctx, uid).PDPConfigUpdateRequest(pDPConfigUpdateRequest).XEncryptionKey(xEncryptionKey).Execute()

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
	xEncryptionKey := "xEncryptionKey_example" // string | Client encryption key for double encryption mode. Must be a base64-encoded AES-256 key (32 bytes). Required only when accessing resources encrypted with encryption_mode='double'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.UpdatePdpConfigApiV1ClientsUidPdpConfigPut(context.Background(), uid).PDPConfigUpdateRequest(pDPConfigUpdateRequest).XEncryptionKey(xEncryptionKey).Execute()
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
 **xEncryptionKey** | **string** | Client encryption key for double encryption mode. Must be a base64-encoded AES-256 key (32 bytes). Required only when accessing resources encrypted with encryption_mode&#x3D;&#39;double&#39;. | 

### Return type

[**PDPConfigResponse**](PDPConfigResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


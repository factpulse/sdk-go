# \AFNORPDPPAAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAfnorCredentialsApiV1AfnorCredentialsGet**](AFNORPDPPAAPI.md#GetAfnorCredentialsApiV1AfnorCredentialsGet) | **Get** /api/v1/afnor/credentials | Récupérer les credentials AFNOR stockés
[**GetFluxEntrantApiV1AfnorFluxEntrantsFlowIdGet**](AFNORPDPPAAPI.md#GetFluxEntrantApiV1AfnorFluxEntrantsFlowIdGet) | **Get** /api/v1/afnor/flux-entrants/{flow_id} | Récupérer et extraire une facture entrante
[**OauthTokenProxyApiV1AfnorOauthTokenPost**](AFNORPDPPAAPI.md#OauthTokenProxyApiV1AfnorOauthTokenPost) | **Post** /api/v1/afnor/oauth/token | Endpoint OAuth2 pour authentification AFNOR



## GetAfnorCredentialsApiV1AfnorCredentialsGet

> interface{} GetAfnorCredentialsApiV1AfnorCredentialsGet(ctx).Execute()

Récupérer les credentials AFNOR stockés



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


## GetFluxEntrantApiV1AfnorFluxEntrantsFlowIdGet

> FactureEntrante GetFluxEntrantApiV1AfnorFluxEntrantsFlowIdGet(ctx, flowId).Execute()

Récupérer et extraire une facture entrante



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
	flowId := "flowId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPAAPI.GetFluxEntrantApiV1AfnorFluxEntrantsFlowIdGet(context.Background(), flowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPAAPI.GetFluxEntrantApiV1AfnorFluxEntrantsFlowIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFluxEntrantApiV1AfnorFluxEntrantsFlowIdGet`: FactureEntrante
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPAAPI.GetFluxEntrantApiV1AfnorFluxEntrantsFlowIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFluxEntrantApiV1AfnorFluxEntrantsFlowIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FactureEntrante**](FactureEntrante.md)

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

Endpoint OAuth2 pour authentification AFNOR



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


# \AFNORDirectoryServiceMtierAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSirenMetierApiV1AfnorDirectorySirenSirenGet**](AFNORDirectoryServiceMtierAPI.md#GetSirenMetierApiV1AfnorDirectorySirenSirenGet) | **Get** /api/v1/afnor/directory/siren/{siren} | Récupérer une entreprise par SIREN (multi-tenant)
[**GetSiretMetierApiV1AfnorDirectorySiretSiretGet**](AFNORDirectoryServiceMtierAPI.md#GetSiretMetierApiV1AfnorDirectorySiretSiretGet) | **Get** /api/v1/afnor/directory/siret/{siret} | Récupérer un établissement par SIRET (multi-tenant)
[**SearchSirenMetierApiV1AfnorDirectorySirenSearchPost**](AFNORDirectoryServiceMtierAPI.md#SearchSirenMetierApiV1AfnorDirectorySirenSearchPost) | **Post** /api/v1/afnor/directory/siren/search | Rechercher des entreprises (multi-tenant)
[**SearchSiretMetierApiV1AfnorDirectorySiretSearchPost**](AFNORDirectoryServiceMtierAPI.md#SearchSiretMetierApiV1AfnorDirectorySiretSearchPost) | **Post** /api/v1/afnor/directory/siret/search | Rechercher des établissements (multi-tenant)



## GetSirenMetierApiV1AfnorDirectorySirenSirenGet

> interface{} GetSirenMetierApiV1AfnorDirectorySirenSirenGet(ctx, siren).PDPCredentials(pDPCredentials).Execute()

Récupérer une entreprise par SIREN (multi-tenant)



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
	siren := "siren_example" // string | 
	pDPCredentials := *openapiclient.NewPDPCredentials("https://api.pdp-example.fr/flow/v1", "https://auth.pdp-example.fr/oauth/token", "factpulse_prod_abc123", "secret_xyz789") // PDPCredentials |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORDirectoryServiceMtierAPI.GetSirenMetierApiV1AfnorDirectorySirenSirenGet(context.Background(), siren).PDPCredentials(pDPCredentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORDirectoryServiceMtierAPI.GetSirenMetierApiV1AfnorDirectorySirenSirenGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSirenMetierApiV1AfnorDirectorySirenSirenGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AFNORDirectoryServiceMtierAPI.GetSirenMetierApiV1AfnorDirectorySirenSirenGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siren** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSirenMetierApiV1AfnorDirectorySirenSirenGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pDPCredentials** | [**PDPCredentials**](PDPCredentials.md) |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiretMetierApiV1AfnorDirectorySiretSiretGet

> interface{} GetSiretMetierApiV1AfnorDirectorySiretSiretGet(ctx, siret).PDPCredentials(pDPCredentials).Execute()

Récupérer un établissement par SIRET (multi-tenant)



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
	siret := "siret_example" // string | 
	pDPCredentials := *openapiclient.NewPDPCredentials("https://api.pdp-example.fr/flow/v1", "https://auth.pdp-example.fr/oauth/token", "factpulse_prod_abc123", "secret_xyz789") // PDPCredentials |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORDirectoryServiceMtierAPI.GetSiretMetierApiV1AfnorDirectorySiretSiretGet(context.Background(), siret).PDPCredentials(pDPCredentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORDirectoryServiceMtierAPI.GetSiretMetierApiV1AfnorDirectorySiretSiretGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiretMetierApiV1AfnorDirectorySiretSiretGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AFNORDirectoryServiceMtierAPI.GetSiretMetierApiV1AfnorDirectorySiretSiretGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siret** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiretMetierApiV1AfnorDirectorySiretSiretGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pDPCredentials** | [**PDPCredentials**](PDPCredentials.md) |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSirenMetierApiV1AfnorDirectorySirenSearchPost

> interface{} SearchSirenMetierApiV1AfnorDirectorySirenSearchPost(ctx).BodySearchSirenMetierApiV1AfnorDirectorySirenSearchPost(bodySearchSirenMetierApiV1AfnorDirectorySirenSearchPost).Execute()

Rechercher des entreprises (multi-tenant)



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
	bodySearchSirenMetierApiV1AfnorDirectorySirenSearchPost := *openapiclient.NewBodySearchSirenMetierApiV1AfnorDirectorySirenSearchPost(map[string]interface{}{"key": interface{}(123)}) // BodySearchSirenMetierApiV1AfnorDirectorySirenSearchPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORDirectoryServiceMtierAPI.SearchSirenMetierApiV1AfnorDirectorySirenSearchPost(context.Background()).BodySearchSirenMetierApiV1AfnorDirectorySirenSearchPost(bodySearchSirenMetierApiV1AfnorDirectorySirenSearchPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORDirectoryServiceMtierAPI.SearchSirenMetierApiV1AfnorDirectorySirenSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSirenMetierApiV1AfnorDirectorySirenSearchPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AFNORDirectoryServiceMtierAPI.SearchSirenMetierApiV1AfnorDirectorySirenSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSirenMetierApiV1AfnorDirectorySirenSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodySearchSirenMetierApiV1AfnorDirectorySirenSearchPost** | [**BodySearchSirenMetierApiV1AfnorDirectorySirenSearchPost**](BodySearchSirenMetierApiV1AfnorDirectorySirenSearchPost.md) |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiretMetierApiV1AfnorDirectorySiretSearchPost

> interface{} SearchSiretMetierApiV1AfnorDirectorySiretSearchPost(ctx).BodySearchSiretMetierApiV1AfnorDirectorySiretSearchPost(bodySearchSiretMetierApiV1AfnorDirectorySiretSearchPost).Execute()

Rechercher des établissements (multi-tenant)



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
	bodySearchSiretMetierApiV1AfnorDirectorySiretSearchPost := *openapiclient.NewBodySearchSiretMetierApiV1AfnorDirectorySiretSearchPost(map[string]interface{}{"key": interface{}(123)}) // BodySearchSiretMetierApiV1AfnorDirectorySiretSearchPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORDirectoryServiceMtierAPI.SearchSiretMetierApiV1AfnorDirectorySiretSearchPost(context.Background()).BodySearchSiretMetierApiV1AfnorDirectorySiretSearchPost(bodySearchSiretMetierApiV1AfnorDirectorySiretSearchPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORDirectoryServiceMtierAPI.SearchSiretMetierApiV1AfnorDirectorySiretSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiretMetierApiV1AfnorDirectorySiretSearchPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AFNORDirectoryServiceMtierAPI.SearchSiretMetierApiV1AfnorDirectorySiretSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiretMetierApiV1AfnorDirectorySiretSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodySearchSiretMetierApiV1AfnorDirectorySiretSearchPost** | [**BodySearchSiretMetierApiV1AfnorDirectorySiretSearchPost**](BodySearchSiretMetierApiV1AfnorDirectorySiretSearchPost.md) |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


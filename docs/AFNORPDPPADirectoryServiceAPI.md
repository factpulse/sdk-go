# \AFNORPDPPADirectoryServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryHealthcheckProxyApiV1AfnorDirectoryV1HealthcheckGet**](AFNORPDPPADirectoryServiceAPI.md#DirectoryHealthcheckProxyApiV1AfnorDirectoryV1HealthcheckGet) | **Get** /api/v1/afnor/directory/v1/healthcheck | Healthcheck Directory Service
[**GetCompanyProxyApiV1AfnorDirectoryV1CompaniesSirenGet**](AFNORPDPPADirectoryServiceAPI.md#GetCompanyProxyApiV1AfnorDirectoryV1CompaniesSirenGet) | **Get** /api/v1/afnor/directory/v1/companies/{siren} | Récupérer une entreprise
[**SearchCompaniesProxyApiV1AfnorDirectoryV1CompaniesSearchPost**](AFNORPDPPADirectoryServiceAPI.md#SearchCompaniesProxyApiV1AfnorDirectoryV1CompaniesSearchPost) | **Post** /api/v1/afnor/directory/v1/companies/search | Rechercher des entreprises



## DirectoryHealthcheckProxyApiV1AfnorDirectoryV1HealthcheckGet

> interface{} DirectoryHealthcheckProxyApiV1AfnorDirectoryV1HealthcheckGet(ctx).Execute()

Healthcheck Directory Service



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
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.DirectoryHealthcheckProxyApiV1AfnorDirectoryV1HealthcheckGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.DirectoryHealthcheckProxyApiV1AfnorDirectoryV1HealthcheckGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DirectoryHealthcheckProxyApiV1AfnorDirectoryV1HealthcheckGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.DirectoryHealthcheckProxyApiV1AfnorDirectoryV1HealthcheckGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryHealthcheckProxyApiV1AfnorDirectoryV1HealthcheckGetRequest struct via the builder pattern


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


## GetCompanyProxyApiV1AfnorDirectoryV1CompaniesSirenGet

> interface{} GetCompanyProxyApiV1AfnorDirectoryV1CompaniesSirenGet(ctx, siren).Execute()

Récupérer une entreprise



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.GetCompanyProxyApiV1AfnorDirectoryV1CompaniesSirenGet(context.Background(), siren).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.GetCompanyProxyApiV1AfnorDirectoryV1CompaniesSirenGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyProxyApiV1AfnorDirectoryV1CompaniesSirenGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.GetCompanyProxyApiV1AfnorDirectoryV1CompaniesSirenGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siren** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyProxyApiV1AfnorDirectoryV1CompaniesSirenGetRequest struct via the builder pattern


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


## SearchCompaniesProxyApiV1AfnorDirectoryV1CompaniesSearchPost

> interface{} SearchCompaniesProxyApiV1AfnorDirectoryV1CompaniesSearchPost(ctx).Execute()

Rechercher des entreprises



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
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.SearchCompaniesProxyApiV1AfnorDirectoryV1CompaniesSearchPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.SearchCompaniesProxyApiV1AfnorDirectoryV1CompaniesSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCompaniesProxyApiV1AfnorDirectoryV1CompaniesSearchPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.SearchCompaniesProxyApiV1AfnorDirectoryV1CompaniesSearchPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchCompaniesProxyApiV1AfnorDirectoryV1CompaniesSearchPostRequest struct via the builder pattern


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


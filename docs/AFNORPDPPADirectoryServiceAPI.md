# \AFNORPDPPADirectoryServiceAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryHealthcheckProxyApiV1AfnorDirectoryV1HealthcheckGet**](AFNORPDPPADirectoryServiceAPI.md#DirectoryHealthcheckProxyApiV1AfnorDirectoryV1HealthcheckGet) | **Get** /api/v1/afnor/directory/v1/healthcheck | Healthcheck Directory Service
[**GetDirectoryLineByCodeProxyApiV1AfnorDirectoryV1DirectoryLineCodeAddressingIdentifierGet**](AFNORPDPPADirectoryServiceAPI.md#GetDirectoryLineByCodeProxyApiV1AfnorDirectoryV1DirectoryLineCodeAddressingIdentifierGet) | **Get** /api/v1/afnor/directory/v1/directory-line/code:{addressing_identifier} | Get a directory line.
[**GetRoutingCodeBySiretAndCodeProxyApiV1AfnorDirectoryV1RoutingCodeSiretSiretCodeRoutingIdentifierGet**](AFNORPDPPADirectoryServiceAPI.md#GetRoutingCodeBySiretAndCodeProxyApiV1AfnorDirectoryV1RoutingCodeSiretSiretCodeRoutingIdentifierGet) | **Get** /api/v1/afnor/directory/v1/routing-code/siret:{siret}/code:{routing_identifier} | Get a routing code by SIRET and routing identifier
[**GetSirenByCodeInseeProxyApiV1AfnorDirectoryV1SirenCodeInseeSirenGet**](AFNORPDPPADirectoryServiceAPI.md#GetSirenByCodeInseeProxyApiV1AfnorDirectoryV1SirenCodeInseeSirenGet) | **Get** /api/v1/afnor/directory/v1/siren/code-insee:{siren} | Consult a siren (legal unit) by SIREN number
[**GetSiretByCodeInseeProxyApiV1AfnorDirectoryV1SiretCodeInseeSiretGet**](AFNORPDPPADirectoryServiceAPI.md#GetSiretByCodeInseeProxyApiV1AfnorDirectoryV1SiretCodeInseeSiretGet) | **Get** /api/v1/afnor/directory/v1/siret/code-insee:{siret} | Gets a siret (facility) by SIRET number
[**SearchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineSearchPost**](AFNORPDPPADirectoryServiceAPI.md#SearchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineSearchPost) | **Post** /api/v1/afnor/directory/v1/directory-line/search | Search for a directory line
[**SearchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeSearchPost**](AFNORPDPPADirectoryServiceAPI.md#SearchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeSearchPost) | **Post** /api/v1/afnor/directory/v1/routing-code/search | Search for a routing code
[**SearchSirenProxyApiV1AfnorDirectoryV1SirenSearchPost**](AFNORPDPPADirectoryServiceAPI.md#SearchSirenProxyApiV1AfnorDirectoryV1SirenSearchPost) | **Post** /api/v1/afnor/directory/v1/siren/search | SIREN search (or legal unit)
[**SearchSiretProxyApiV1AfnorDirectoryV1SiretSearchPost**](AFNORPDPPADirectoryServiceAPI.md#SearchSiretProxyApiV1AfnorDirectoryV1SiretSearchPost) | **Post** /api/v1/afnor/directory/v1/siret/search | Search for a SIRET (facility)



## DirectoryHealthcheckProxyApiV1AfnorDirectoryV1HealthcheckGet

> map[string]interface{} DirectoryHealthcheckProxyApiV1AfnorDirectoryV1HealthcheckGet(ctx).Execute()

Healthcheck Directory Service



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
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.DirectoryHealthcheckProxyApiV1AfnorDirectoryV1HealthcheckGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.DirectoryHealthcheckProxyApiV1AfnorDirectoryV1HealthcheckGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DirectoryHealthcheckProxyApiV1AfnorDirectoryV1HealthcheckGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.DirectoryHealthcheckProxyApiV1AfnorDirectoryV1HealthcheckGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryHealthcheckProxyApiV1AfnorDirectoryV1HealthcheckGetRequest struct via the builder pattern


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


## GetDirectoryLineByCodeProxyApiV1AfnorDirectoryV1DirectoryLineCodeAddressingIdentifierGet

> AFNORDirectoryLinePayloadLegalUnitFacilityRoutingCode GetDirectoryLineByCodeProxyApiV1AfnorDirectoryV1DirectoryLineCodeAddressingIdentifierGet(ctx, addressingIdentifier).Execute()

Get a directory line.



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
	addressingIdentifier := "addressingIdentifier_example" // string | Addressing identifier (SIREN, SIRET or routing code)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.GetDirectoryLineByCodeProxyApiV1AfnorDirectoryV1DirectoryLineCodeAddressingIdentifierGet(context.Background(), addressingIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.GetDirectoryLineByCodeProxyApiV1AfnorDirectoryV1DirectoryLineCodeAddressingIdentifierGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDirectoryLineByCodeProxyApiV1AfnorDirectoryV1DirectoryLineCodeAddressingIdentifierGet`: AFNORDirectoryLinePayloadLegalUnitFacilityRoutingCode
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.GetDirectoryLineByCodeProxyApiV1AfnorDirectoryV1DirectoryLineCodeAddressingIdentifierGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addressingIdentifier** | **string** | Addressing identifier (SIREN, SIRET or routing code) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDirectoryLineByCodeProxyApiV1AfnorDirectoryV1DirectoryLineCodeAddressingIdentifierGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AFNORDirectoryLinePayloadLegalUnitFacilityRoutingCode**](AFNORDirectoryLinePayloadLegalUnitFacilityRoutingCode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoutingCodeBySiretAndCodeProxyApiV1AfnorDirectoryV1RoutingCodeSiretSiretCodeRoutingIdentifierGet

> AFNORRoutingCodePayloadHistoryLegalUnitFacility GetRoutingCodeBySiretAndCodeProxyApiV1AfnorDirectoryV1RoutingCodeSiretSiretCodeRoutingIdentifierGet(ctx, siret, routingIdentifier).Execute()

Get a routing code by SIRET and routing identifier



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
	siret := "siret_example" // string | 14-digit SIRET number (INSEE establishment identifier)
	routingIdentifier := "routingIdentifier_example" // string | Routing code identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.GetRoutingCodeBySiretAndCodeProxyApiV1AfnorDirectoryV1RoutingCodeSiretSiretCodeRoutingIdentifierGet(context.Background(), siret, routingIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.GetRoutingCodeBySiretAndCodeProxyApiV1AfnorDirectoryV1RoutingCodeSiretSiretCodeRoutingIdentifierGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoutingCodeBySiretAndCodeProxyApiV1AfnorDirectoryV1RoutingCodeSiretSiretCodeRoutingIdentifierGet`: AFNORRoutingCodePayloadHistoryLegalUnitFacility
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.GetRoutingCodeBySiretAndCodeProxyApiV1AfnorDirectoryV1RoutingCodeSiretSiretCodeRoutingIdentifierGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siret** | **string** | 14-digit SIRET number (INSEE establishment identifier) | 
**routingIdentifier** | **string** | Routing code identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoutingCodeBySiretAndCodeProxyApiV1AfnorDirectoryV1RoutingCodeSiretSiretCodeRoutingIdentifierGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AFNORRoutingCodePayloadHistoryLegalUnitFacility**](AFNORRoutingCodePayloadHistoryLegalUnitFacility.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSirenByCodeInseeProxyApiV1AfnorDirectoryV1SirenCodeInseeSirenGet

> AFNORLegalUnitPayloadHistory GetSirenByCodeInseeProxyApiV1AfnorDirectoryV1SirenCodeInseeSirenGet(ctx, siren).Execute()

Consult a siren (legal unit) by SIREN number



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
	siren := "siren_example" // string | 9-digit SIREN number (INSEE company identifier)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.GetSirenByCodeInseeProxyApiV1AfnorDirectoryV1SirenCodeInseeSirenGet(context.Background(), siren).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.GetSirenByCodeInseeProxyApiV1AfnorDirectoryV1SirenCodeInseeSirenGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSirenByCodeInseeProxyApiV1AfnorDirectoryV1SirenCodeInseeSirenGet`: AFNORLegalUnitPayloadHistory
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.GetSirenByCodeInseeProxyApiV1AfnorDirectoryV1SirenCodeInseeSirenGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siren** | **string** | 9-digit SIREN number (INSEE company identifier) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSirenByCodeInseeProxyApiV1AfnorDirectoryV1SirenCodeInseeSirenGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AFNORLegalUnitPayloadHistory**](AFNORLegalUnitPayloadHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiretByCodeInseeProxyApiV1AfnorDirectoryV1SiretCodeInseeSiretGet

> AFNORFacilityPayloadHistory GetSiretByCodeInseeProxyApiV1AfnorDirectoryV1SiretCodeInseeSiretGet(ctx, siret).Execute()

Gets a siret (facility) by SIRET number



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
	siret := "siret_example" // string | 14-digit SIRET number (INSEE establishment identifier)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.GetSiretByCodeInseeProxyApiV1AfnorDirectoryV1SiretCodeInseeSiretGet(context.Background(), siret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.GetSiretByCodeInseeProxyApiV1AfnorDirectoryV1SiretCodeInseeSiretGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiretByCodeInseeProxyApiV1AfnorDirectoryV1SiretCodeInseeSiretGet`: AFNORFacilityPayloadHistory
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.GetSiretByCodeInseeProxyApiV1AfnorDirectoryV1SiretCodeInseeSiretGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siret** | **string** | 14-digit SIRET number (INSEE establishment identifier) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiretByCodeInseeProxyApiV1AfnorDirectoryV1SiretCodeInseeSiretGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AFNORFacilityPayloadHistory**](AFNORFacilityPayloadHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineSearchPost

> AFNORDirectoryLineSearchPost200Response SearchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineSearchPost(ctx).Execute()

Search for a directory line



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
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.SearchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineSearchPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.SearchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineSearchPost`: AFNORDirectoryLineSearchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.SearchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineSearchPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineSearchPostRequest struct via the builder pattern


### Return type

[**AFNORDirectoryLineSearchPost200Response**](AFNORDirectoryLineSearchPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeSearchPost

> AFNORRoutingCodeSearchPost200Response SearchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeSearchPost(ctx).Execute()

Search for a routing code



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
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.SearchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeSearchPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.SearchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeSearchPost`: AFNORRoutingCodeSearchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.SearchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeSearchPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeSearchPostRequest struct via the builder pattern


### Return type

[**AFNORRoutingCodeSearchPost200Response**](AFNORRoutingCodeSearchPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSirenProxyApiV1AfnorDirectoryV1SirenSearchPost

> AFNORSirenSearchPost200Response SearchSirenProxyApiV1AfnorDirectoryV1SirenSearchPost(ctx).Execute()

SIREN search (or legal unit)



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
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.SearchSirenProxyApiV1AfnorDirectoryV1SirenSearchPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.SearchSirenProxyApiV1AfnorDirectoryV1SirenSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSirenProxyApiV1AfnorDirectoryV1SirenSearchPost`: AFNORSirenSearchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.SearchSirenProxyApiV1AfnorDirectoryV1SirenSearchPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSirenProxyApiV1AfnorDirectoryV1SirenSearchPostRequest struct via the builder pattern


### Return type

[**AFNORSirenSearchPost200Response**](AFNORSirenSearchPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiretProxyApiV1AfnorDirectoryV1SiretSearchPost

> AFNORSiretSearchPost200Response SearchSiretProxyApiV1AfnorDirectoryV1SiretSearchPost(ctx).Execute()

Search for a SIRET (facility)



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
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.SearchSiretProxyApiV1AfnorDirectoryV1SiretSearchPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.SearchSiretProxyApiV1AfnorDirectoryV1SiretSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiretProxyApiV1AfnorDirectoryV1SiretSearchPost`: AFNORSiretSearchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.SearchSiretProxyApiV1AfnorDirectoryV1SiretSearchPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiretProxyApiV1AfnorDirectoryV1SiretSearchPostRequest struct via the builder pattern


### Return type

[**AFNORSiretSearchPost200Response**](AFNORSiretSearchPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


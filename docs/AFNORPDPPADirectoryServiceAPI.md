# \AFNORPDPPADirectoryServiceAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLinePost**](AFNORPDPPADirectoryServiceAPI.md#CreateDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLinePost) | **Post** /api/v1/afnor/directory/v1/directory-line | Creating a directory line
[**CreateRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodePost**](AFNORPDPPADirectoryServiceAPI.md#CreateRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodePost) | **Post** /api/v1/afnor/directory/v1/routing-code | Create a routing code
[**DeleteDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstanceDelete**](AFNORPDPPADirectoryServiceAPI.md#DeleteDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstanceDelete) | **Delete** /api/v1/afnor/directory/v1/directory-line/id-instance:{id_instance} | Delete a directory line
[**DirectoryHealthcheckProxyApiV1AfnorDirectoryV1HealthcheckGet**](AFNORPDPPADirectoryServiceAPI.md#DirectoryHealthcheckProxyApiV1AfnorDirectoryV1HealthcheckGet) | **Get** /api/v1/afnor/directory/v1/healthcheck | Healthcheck Directory Service
[**GetDirectoryLineByCodeProxyApiV1AfnorDirectoryV1DirectoryLineCodeAddressingIdentifierGet**](AFNORPDPPADirectoryServiceAPI.md#GetDirectoryLineByCodeProxyApiV1AfnorDirectoryV1DirectoryLineCodeAddressingIdentifierGet) | **Get** /api/v1/afnor/directory/v1/directory-line/code:{addressing_identifier} | Get a directory line.
[**GetDirectoryLineByIdInstanceProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstanceGet**](AFNORPDPPADirectoryServiceAPI.md#GetDirectoryLineByIdInstanceProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstanceGet) | **Get** /api/v1/afnor/directory/v1/directory-line/id-instance:{id_instance} | Get a directory line.
[**GetRoutingCodeByIdInstanceProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstanceGet**](AFNORPDPPADirectoryServiceAPI.md#GetRoutingCodeByIdInstanceProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstanceGet) | **Get** /api/v1/afnor/directory/v1/routing-code/id-instance:{id_instance} | Get a routing code by instance-id.
[**GetRoutingCodeBySiretAndCodeProxyApiV1AfnorDirectoryV1RoutingCodeSiretSiretCodeRoutingIdentifierGet**](AFNORPDPPADirectoryServiceAPI.md#GetRoutingCodeBySiretAndCodeProxyApiV1AfnorDirectoryV1RoutingCodeSiretSiretCodeRoutingIdentifierGet) | **Get** /api/v1/afnor/directory/v1/routing-code/siret:{siret}/code:{routing_identifier} | Get a routing code by SIRET and routing identifier
[**GetSirenByCodeInseeProxyApiV1AfnorDirectoryV1SirenCodeInseeSirenGet**](AFNORPDPPADirectoryServiceAPI.md#GetSirenByCodeInseeProxyApiV1AfnorDirectoryV1SirenCodeInseeSirenGet) | **Get** /api/v1/afnor/directory/v1/siren/code-insee:{siren} | Consult a siren (legal unit) by SIREN number
[**GetSirenByIdInstanceProxyApiV1AfnorDirectoryV1SirenIdInstanceIdInstanceGet**](AFNORPDPPADirectoryServiceAPI.md#GetSirenByIdInstanceProxyApiV1AfnorDirectoryV1SirenIdInstanceIdInstanceGet) | **Get** /api/v1/afnor/directory/v1/siren/id-instance:{id_instance} | Gets a siren (legal unit) by instance ID
[**GetSiretByCodeInseeProxyApiV1AfnorDirectoryV1SiretCodeInseeSiretGet**](AFNORPDPPADirectoryServiceAPI.md#GetSiretByCodeInseeProxyApiV1AfnorDirectoryV1SiretCodeInseeSiretGet) | **Get** /api/v1/afnor/directory/v1/siret/code-insee:{siret} | Gets a siret (facility) by SIRET number
[**GetSiretByIdInstanceProxyApiV1AfnorDirectoryV1SiretIdInstanceIdInstanceGet**](AFNORPDPPADirectoryServiceAPI.md#GetSiretByIdInstanceProxyApiV1AfnorDirectoryV1SiretIdInstanceIdInstanceGet) | **Get** /api/v1/afnor/directory/v1/siret/id-instance:{id_instance} | Gets a siret (facility) by id-instance
[**PatchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstancePatch**](AFNORPDPPADirectoryServiceAPI.md#PatchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstancePatch) | **Patch** /api/v1/afnor/directory/v1/directory-line/id-instance:{id_instance} | Partially updates a directory line..
[**PatchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstancePatch**](AFNORPDPPADirectoryServiceAPI.md#PatchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstancePatch) | **Patch** /api/v1/afnor/directory/v1/routing-code/id-instance:{id_instance} | Partially update a private routing code.
[**PutRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstancePut**](AFNORPDPPADirectoryServiceAPI.md#PutRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstancePut) | **Put** /api/v1/afnor/directory/v1/routing-code/id-instance:{id_instance} | Completely update a private routing code.
[**SearchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineSearchPost**](AFNORPDPPADirectoryServiceAPI.md#SearchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineSearchPost) | **Post** /api/v1/afnor/directory/v1/directory-line/search | Search for a directory line
[**SearchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeSearchPost**](AFNORPDPPADirectoryServiceAPI.md#SearchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeSearchPost) | **Post** /api/v1/afnor/directory/v1/routing-code/search | Search for a routing code
[**SearchSirenProxyApiV1AfnorDirectoryV1SirenSearchPost**](AFNORPDPPADirectoryServiceAPI.md#SearchSirenProxyApiV1AfnorDirectoryV1SirenSearchPost) | **Post** /api/v1/afnor/directory/v1/siren/search | SIREN search (or legal unit)
[**SearchSiretProxyApiV1AfnorDirectoryV1SiretSearchPost**](AFNORPDPPADirectoryServiceAPI.md#SearchSiretProxyApiV1AfnorDirectoryV1SiretSearchPost) | **Post** /api/v1/afnor/directory/v1/siret/search | Search for a SIRET (facility)



## CreateDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLinePost

> interface{} CreateDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLinePost(ctx).AcceptLanguage(acceptLanguage).Execute()

Creating a directory line



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
	acceptLanguage := openapiclient.AcceptLanguage("fr") // AcceptLanguage | Specifies the language in which the resource is requested. (optional) (default to "fr")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.CreateDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLinePost(context.Background()).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.CreateDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLinePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLinePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.CreateDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLinePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLinePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | [**AcceptLanguage**](AcceptLanguage.md) | Specifies the language in which the resource is requested. | [default to &quot;fr&quot;]

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


## CreateRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodePost

> interface{} CreateRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodePost(ctx).AcceptLanguage(acceptLanguage).Execute()

Create a routing code



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
	acceptLanguage := openapiclient.AcceptLanguage("fr") // AcceptLanguage | Specifies the language in which the resource is requested. (optional) (default to "fr")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.CreateRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodePost(context.Background()).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.CreateRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.CreateRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | [**AcceptLanguage**](AcceptLanguage.md) | Specifies the language in which the resource is requested. | [default to &quot;fr&quot;]

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


## DeleteDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstanceDelete

> interface{} DeleteDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstanceDelete(ctx, idInstance).AcceptLanguage(acceptLanguage).Execute()

Delete a directory line



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
	idInstance := "idInstance_example" // string | AFNOR instance ID (UUID)
	acceptLanguage := openapiclient.AcceptLanguage("fr") // AcceptLanguage | Specifies the language in which the resource is requested. (optional) (default to "fr")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.DeleteDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstanceDelete(context.Background(), idInstance).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.DeleteDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstanceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstanceDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.DeleteDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstanceDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idInstance** | **string** | AFNOR instance ID (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstanceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | [**AcceptLanguage**](AcceptLanguage.md) | Specifies the language in which the resource is requested. | [default to &quot;fr&quot;]

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
	openapiclient "github.com/factpulse/sdk-go"
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

> AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode GetDirectoryLineByCodeProxyApiV1AfnorDirectoryV1DirectoryLineCodeAddressingIdentifierGet(ctx, addressingIdentifier).Fields(fields).Include(include).AcceptLanguage(acceptLanguage).Execute()

Get a directory line.



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
	addressingIdentifier := "addressingIdentifier_example" // string | Addressing identifier (SIREN, SIRET or routing code)
	fields := []string{"Inner_example"} // []string | Fields of the Directory Line resource. (optional)
	include := []openapiclient.DirectoryLineInclude{openapiclient.DirectoryLineInclude("siren")} // []DirectoryLineInclude | Relations to include in the response. (optional)
	acceptLanguage := openapiclient.AcceptLanguage("fr") // AcceptLanguage | Specifies the language in which the resource is requested. (optional) (default to "fr")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.GetDirectoryLineByCodeProxyApiV1AfnorDirectoryV1DirectoryLineCodeAddressingIdentifierGet(context.Background(), addressingIdentifier).Fields(fields).Include(include).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.GetDirectoryLineByCodeProxyApiV1AfnorDirectoryV1DirectoryLineCodeAddressingIdentifierGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDirectoryLineByCodeProxyApiV1AfnorDirectoryV1DirectoryLineCodeAddressingIdentifierGet`: AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode
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

 **fields** | **[]string** | Fields of the Directory Line resource. | 
 **include** | [**[]DirectoryLineInclude**](DirectoryLineInclude.md) | Relations to include in the response. | 
 **acceptLanguage** | [**AcceptLanguage**](AcceptLanguage.md) | Specifies the language in which the resource is requested. | [default to &quot;fr&quot;]

### Return type

[**AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode**](AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDirectoryLineByIdInstanceProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstanceGet

> AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode GetDirectoryLineByIdInstanceProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstanceGet(ctx, idInstance).Fields(fields).AcceptLanguage(acceptLanguage).Execute()

Get a directory line.



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
	idInstance := "idInstance_example" // string | AFNOR instance ID (UUID)
	fields := []string{"Inner_example"} // []string | Fields of the Directory Line resource. (optional)
	acceptLanguage := openapiclient.AcceptLanguage("fr") // AcceptLanguage | Specifies the language in which the resource is requested. (optional) (default to "fr")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.GetDirectoryLineByIdInstanceProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstanceGet(context.Background(), idInstance).Fields(fields).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.GetDirectoryLineByIdInstanceProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstanceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDirectoryLineByIdInstanceProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstanceGet`: AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.GetDirectoryLineByIdInstanceProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstanceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idInstance** | **string** | AFNOR instance ID (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDirectoryLineByIdInstanceProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | Fields of the Directory Line resource. | 
 **acceptLanguage** | [**AcceptLanguage**](AcceptLanguage.md) | Specifies the language in which the resource is requested. | [default to &quot;fr&quot;]

### Return type

[**AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode**](AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoutingCodeByIdInstanceProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstanceGet

> AFNORRoutingCodePayloadHistoryLegalUnitFacility GetRoutingCodeByIdInstanceProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstanceGet(ctx, idInstance).Fields(fields).AcceptLanguage(acceptLanguage).Execute()

Get a routing code by instance-id.



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
	idInstance := "idInstance_example" // string | AFNOR instance ID (UUID)
	fields := []string{"Inner_example"} // []string | Fields of the Routing Code resource (optional)
	acceptLanguage := openapiclient.AcceptLanguage("fr") // AcceptLanguage | Specifies the language in which the resource is requested. (optional) (default to "fr")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.GetRoutingCodeByIdInstanceProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstanceGet(context.Background(), idInstance).Fields(fields).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.GetRoutingCodeByIdInstanceProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstanceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoutingCodeByIdInstanceProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstanceGet`: AFNORRoutingCodePayloadHistoryLegalUnitFacility
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.GetRoutingCodeByIdInstanceProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstanceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idInstance** | **string** | AFNOR instance ID (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoutingCodeByIdInstanceProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | Fields of the Routing Code resource | 
 **acceptLanguage** | [**AcceptLanguage**](AcceptLanguage.md) | Specifies the language in which the resource is requested. | [default to &quot;fr&quot;]

### Return type

[**AFNORRoutingCodePayloadHistoryLegalUnitFacility**](AFNORRoutingCodePayloadHistoryLegalUnitFacility.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoutingCodeBySiretAndCodeProxyApiV1AfnorDirectoryV1RoutingCodeSiretSiretCodeRoutingIdentifierGet

> AFNORRoutingCodePayloadHistoryLegalUnitFacility GetRoutingCodeBySiretAndCodeProxyApiV1AfnorDirectoryV1RoutingCodeSiretSiretCodeRoutingIdentifierGet(ctx, siret, routingIdentifier).Fields(fields).Include(include).AcceptLanguage(acceptLanguage).Execute()

Get a routing code by SIRET and routing identifier



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
	siret := "siret_example" // string | 14-digit SIRET number (INSEE establishment identifier)
	routingIdentifier := "routingIdentifier_example" // string | Routing code identifier
	fields := []string{"Inner_example"} // []string | Fields of the Routing Code resource (optional)
	include := []openapiclient.RoutingCodeInclude{openapiclient.RoutingCodeInclude("siren")} // []RoutingCodeInclude | Relations to include in the response. (optional)
	acceptLanguage := openapiclient.AcceptLanguage("fr") // AcceptLanguage | Specifies the language in which the resource is requested. (optional) (default to "fr")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.GetRoutingCodeBySiretAndCodeProxyApiV1AfnorDirectoryV1RoutingCodeSiretSiretCodeRoutingIdentifierGet(context.Background(), siret, routingIdentifier).Fields(fields).Include(include).AcceptLanguage(acceptLanguage).Execute()
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


 **fields** | **[]string** | Fields of the Routing Code resource | 
 **include** | [**[]RoutingCodeInclude**](RoutingCodeInclude.md) | Relations to include in the response. | 
 **acceptLanguage** | [**AcceptLanguage**](AcceptLanguage.md) | Specifies the language in which the resource is requested. | [default to &quot;fr&quot;]

### Return type

[**AFNORRoutingCodePayloadHistoryLegalUnitFacility**](AFNORRoutingCodePayloadHistoryLegalUnitFacility.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSirenByCodeInseeProxyApiV1AfnorDirectoryV1SirenCodeInseeSirenGet

> AFNORLegalUnitPayloadHistory GetSirenByCodeInseeProxyApiV1AfnorDirectoryV1SirenCodeInseeSirenGet(ctx, siren).Fields(fields).AcceptLanguage(acceptLanguage).Execute()

Consult a siren (legal unit) by SIREN number



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
	siren := "siren_example" // string | 9-digit SIREN number (INSEE company identifier)
	fields := []*string{"Inner_example"} // []*string | Fields of the SIREN resource (optional)
	acceptLanguage := openapiclient.AcceptLanguage("fr") // AcceptLanguage | Specifies the language in which the resource is requested. (optional) (default to "fr")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.GetSirenByCodeInseeProxyApiV1AfnorDirectoryV1SirenCodeInseeSirenGet(context.Background(), siren).Fields(fields).AcceptLanguage(acceptLanguage).Execute()
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

 **fields** | **[]string** | Fields of the SIREN resource | 
 **acceptLanguage** | [**AcceptLanguage**](AcceptLanguage.md) | Specifies the language in which the resource is requested. | [default to &quot;fr&quot;]

### Return type

[**AFNORLegalUnitPayloadHistory**](AFNORLegalUnitPayloadHistory.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSirenByIdInstanceProxyApiV1AfnorDirectoryV1SirenIdInstanceIdInstanceGet

> AFNORLegalUnitPayloadHistory GetSirenByIdInstanceProxyApiV1AfnorDirectoryV1SirenIdInstanceIdInstanceGet(ctx, idInstance).Fields(fields).AcceptLanguage(acceptLanguage).Execute()

Gets a siren (legal unit) by instance ID



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
	idInstance := "idInstance_example" // string | AFNOR instance ID (UUID)
	fields := []string{"Inner_example"} // []string | Fields of the SIREN resource (optional)
	acceptLanguage := openapiclient.AcceptLanguage("fr") // AcceptLanguage | Specifies the language in which the resource is requested. (optional) (default to "fr")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.GetSirenByIdInstanceProxyApiV1AfnorDirectoryV1SirenIdInstanceIdInstanceGet(context.Background(), idInstance).Fields(fields).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.GetSirenByIdInstanceProxyApiV1AfnorDirectoryV1SirenIdInstanceIdInstanceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSirenByIdInstanceProxyApiV1AfnorDirectoryV1SirenIdInstanceIdInstanceGet`: AFNORLegalUnitPayloadHistory
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.GetSirenByIdInstanceProxyApiV1AfnorDirectoryV1SirenIdInstanceIdInstanceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idInstance** | **string** | AFNOR instance ID (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSirenByIdInstanceProxyApiV1AfnorDirectoryV1SirenIdInstanceIdInstanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | Fields of the SIREN resource | 
 **acceptLanguage** | [**AcceptLanguage**](AcceptLanguage.md) | Specifies the language in which the resource is requested. | [default to &quot;fr&quot;]

### Return type

[**AFNORLegalUnitPayloadHistory**](AFNORLegalUnitPayloadHistory.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiretByCodeInseeProxyApiV1AfnorDirectoryV1SiretCodeInseeSiretGet

> AFNORFacilityPayloadHistory GetSiretByCodeInseeProxyApiV1AfnorDirectoryV1SiretCodeInseeSiretGet(ctx, siret).Fields(fields).Include(include).AcceptLanguage(acceptLanguage).Execute()

Gets a siret (facility) by SIRET number



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
	siret := "siret_example" // string | 14-digit SIRET number (INSEE establishment identifier)
	fields := []string{"Inner_example"} // []string | Fields of a SIRET resource. (optional)
	include := []openapiclient.SiretInclude{openapiclient.SiretInclude("siren")} // []SiretInclude | Relations to include in the response. (optional)
	acceptLanguage := openapiclient.AcceptLanguage("fr") // AcceptLanguage | Specifies the language in which the resource is requested. (optional) (default to "fr")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.GetSiretByCodeInseeProxyApiV1AfnorDirectoryV1SiretCodeInseeSiretGet(context.Background(), siret).Fields(fields).Include(include).AcceptLanguage(acceptLanguage).Execute()
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

 **fields** | **[]string** | Fields of a SIRET resource. | 
 **include** | [**[]SiretInclude**](SiretInclude.md) | Relations to include in the response. | 
 **acceptLanguage** | [**AcceptLanguage**](AcceptLanguage.md) | Specifies the language in which the resource is requested. | [default to &quot;fr&quot;]

### Return type

[**AFNORFacilityPayloadHistory**](AFNORFacilityPayloadHistory.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiretByIdInstanceProxyApiV1AfnorDirectoryV1SiretIdInstanceIdInstanceGet

> AFNORFacilityPayloadHistory GetSiretByIdInstanceProxyApiV1AfnorDirectoryV1SiretIdInstanceIdInstanceGet(ctx, idInstance).Fields(fields).AcceptLanguage(acceptLanguage).Execute()

Gets a siret (facility) by id-instance



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
	idInstance := "idInstance_example" // string | AFNOR instance ID (UUID)
	fields := []string{"Inner_example"} // []string | Fields of a SIRET resource. (optional)
	acceptLanguage := openapiclient.AcceptLanguage("fr") // AcceptLanguage | Specifies the language in which the resource is requested. (optional) (default to "fr")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.GetSiretByIdInstanceProxyApiV1AfnorDirectoryV1SiretIdInstanceIdInstanceGet(context.Background(), idInstance).Fields(fields).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.GetSiretByIdInstanceProxyApiV1AfnorDirectoryV1SiretIdInstanceIdInstanceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiretByIdInstanceProxyApiV1AfnorDirectoryV1SiretIdInstanceIdInstanceGet`: AFNORFacilityPayloadHistory
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.GetSiretByIdInstanceProxyApiV1AfnorDirectoryV1SiretIdInstanceIdInstanceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idInstance** | **string** | AFNOR instance ID (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiretByIdInstanceProxyApiV1AfnorDirectoryV1SiretIdInstanceIdInstanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | Fields of a SIRET resource. | 
 **acceptLanguage** | [**AcceptLanguage**](AcceptLanguage.md) | Specifies the language in which the resource is requested. | [default to &quot;fr&quot;]

### Return type

[**AFNORFacilityPayloadHistory**](AFNORFacilityPayloadHistory.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstancePatch

> AFNORDirectoryLinePost201Response PatchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstancePatch(ctx, idInstance).AcceptLanguage(acceptLanguage).Execute()

Partially updates a directory line..



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
	idInstance := "idInstance_example" // string | AFNOR instance ID (UUID)
	acceptLanguage := openapiclient.AcceptLanguage("fr") // AcceptLanguage | Specifies the language in which the resource is requested. (optional) (default to "fr")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.PatchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstancePatch(context.Background(), idInstance).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.PatchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstancePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstancePatch`: AFNORDirectoryLinePost201Response
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.PatchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstancePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idInstance** | **string** | AFNOR instance ID (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineIdInstanceIdInstancePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | [**AcceptLanguage**](AcceptLanguage.md) | Specifies the language in which the resource is requested. | [default to &quot;fr&quot;]

### Return type

[**AFNORDirectoryLinePost201Response**](AFNORDirectoryLinePost201Response.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstancePatch

> AFNORRoutingCodePost201Response PatchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstancePatch(ctx, idInstance).AcceptLanguage(acceptLanguage).Execute()

Partially update a private routing code.



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
	idInstance := "idInstance_example" // string | AFNOR instance ID (UUID)
	acceptLanguage := openapiclient.AcceptLanguage("fr") // AcceptLanguage | Specifies the language in which the resource is requested. (optional) (default to "fr")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.PatchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstancePatch(context.Background(), idInstance).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.PatchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstancePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstancePatch`: AFNORRoutingCodePost201Response
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.PatchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstancePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idInstance** | **string** | AFNOR instance ID (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstancePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | [**AcceptLanguage**](AcceptLanguage.md) | Specifies the language in which the resource is requested. | [default to &quot;fr&quot;]

### Return type

[**AFNORRoutingCodePost201Response**](AFNORRoutingCodePost201Response.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstancePut

> AFNORRoutingCodePost201Response PutRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstancePut(ctx, idInstance).AcceptLanguage(acceptLanguage).Execute()

Completely update a private routing code.



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
	idInstance := "idInstance_example" // string | AFNOR instance ID (UUID)
	acceptLanguage := openapiclient.AcceptLanguage("fr") // AcceptLanguage | Specifies the language in which the resource is requested. (optional) (default to "fr")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.PutRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstancePut(context.Background(), idInstance).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.PutRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstancePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstancePut`: AFNORRoutingCodePost201Response
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.PutRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstancePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idInstance** | **string** | AFNOR instance ID (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeIdInstanceIdInstancePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | [**AcceptLanguage**](AcceptLanguage.md) | Specifies the language in which the resource is requested. | [default to &quot;fr&quot;]

### Return type

[**AFNORRoutingCodePost201Response**](AFNORRoutingCodePost201Response.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineSearchPost

> AFNORDirectoryLineSearchPost200Response SearchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineSearchPost(ctx).AcceptLanguage(acceptLanguage).Execute()

Search for a directory line



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
	acceptLanguage := openapiclient.AcceptLanguage("fr") // AcceptLanguage | Specifies the language in which the resource is requested. (optional) (default to "fr")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.SearchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineSearchPost(context.Background()).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.SearchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineSearchPost`: AFNORDirectoryLineSearchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.SearchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchDirectoryLineProxyApiV1AfnorDirectoryV1DirectoryLineSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | [**AcceptLanguage**](AcceptLanguage.md) | Specifies the language in which the resource is requested. | [default to &quot;fr&quot;]

### Return type

[**AFNORDirectoryLineSearchPost200Response**](AFNORDirectoryLineSearchPost200Response.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeSearchPost

> AFNORRoutingCodeSearchPost200Response SearchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeSearchPost(ctx).AcceptLanguage(acceptLanguage).Execute()

Search for a routing code



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
	acceptLanguage := openapiclient.AcceptLanguage("fr") // AcceptLanguage | Specifies the language in which the resource is requested. (optional) (default to "fr")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.SearchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeSearchPost(context.Background()).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.SearchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeSearchPost`: AFNORRoutingCodeSearchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.SearchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRoutingCodeProxyApiV1AfnorDirectoryV1RoutingCodeSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | [**AcceptLanguage**](AcceptLanguage.md) | Specifies the language in which the resource is requested. | [default to &quot;fr&quot;]

### Return type

[**AFNORRoutingCodeSearchPost200Response**](AFNORRoutingCodeSearchPost200Response.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSirenProxyApiV1AfnorDirectoryV1SirenSearchPost

> AFNORSirenSearchPost200Response SearchSirenProxyApiV1AfnorDirectoryV1SirenSearchPost(ctx).AcceptLanguage(acceptLanguage).Execute()

SIREN search (or legal unit)



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
	acceptLanguage := openapiclient.AcceptLanguage("fr") // AcceptLanguage | Specifies the language in which the resource is requested. (optional) (default to "fr")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.SearchSirenProxyApiV1AfnorDirectoryV1SirenSearchPost(context.Background()).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.SearchSirenProxyApiV1AfnorDirectoryV1SirenSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSirenProxyApiV1AfnorDirectoryV1SirenSearchPost`: AFNORSirenSearchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.SearchSirenProxyApiV1AfnorDirectoryV1SirenSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSirenProxyApiV1AfnorDirectoryV1SirenSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | [**AcceptLanguage**](AcceptLanguage.md) | Specifies the language in which the resource is requested. | [default to &quot;fr&quot;]

### Return type

[**AFNORSirenSearchPost200Response**](AFNORSirenSearchPost200Response.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiretProxyApiV1AfnorDirectoryV1SiretSearchPost

> AFNORSiretSearchPost200Response SearchSiretProxyApiV1AfnorDirectoryV1SiretSearchPost(ctx).AcceptLanguage(acceptLanguage).Execute()

Search for a SIRET (facility)



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
	acceptLanguage := openapiclient.AcceptLanguage("fr") // AcceptLanguage | Specifies the language in which the resource is requested. (optional) (default to "fr")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AFNORPDPPADirectoryServiceAPI.SearchSiretProxyApiV1AfnorDirectoryV1SiretSearchPost(context.Background()).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPADirectoryServiceAPI.SearchSiretProxyApiV1AfnorDirectoryV1SiretSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiretProxyApiV1AfnorDirectoryV1SiretSearchPost`: AFNORSiretSearchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPADirectoryServiceAPI.SearchSiretProxyApiV1AfnorDirectoryV1SiretSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiretProxyApiV1AfnorDirectoryV1SiretSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | [**AcceptLanguage**](AcceptLanguage.md) | Specifies the language in which the resource is requested. | [default to &quot;fr&quot;]

### Return type

[**AFNORSiretSearchPost200Response**](AFNORSiretSearchPost200Response.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \Flux10EReportingAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateAggregatedEreportingApiV1EreportingGenerateAggregatedPost**](Flux10EReportingAPI.md#GenerateAggregatedEreportingApiV1EreportingGenerateAggregatedPost) | **Post** /api/v1/ereporting/generate-aggregated | Generate aggregated e-reporting XML (PPF-compliant)
[**GenerateEreportingApiV1EreportingGeneratePost**](Flux10EReportingAPI.md#GenerateEreportingApiV1EreportingGeneratePost) | **Post** /api/v1/ereporting/generate | Generate e-reporting XML
[**GenerateEreportingDownloadApiV1EreportingGenerateDownloadPost**](Flux10EReportingAPI.md#GenerateEreportingDownloadApiV1EreportingGenerateDownloadPost) | **Post** /api/v1/ereporting/generate/download | Generate and download e-reporting XML
[**ListCategoryCodesApiV1EreportingCategoryCodesGet**](Flux10EReportingAPI.md#ListCategoryCodesApiV1EreportingCategoryCodesGet) | **Get** /api/v1/ereporting/category-codes | List PPF-compliant category codes
[**ListFlowTypesApiV1EreportingFlowTypesGet**](Flux10EReportingAPI.md#ListFlowTypesApiV1EreportingFlowTypesGet) | **Get** /api/v1/ereporting/flow-types | List available flow types
[**SubmitAggregatedEreportingApiV1EreportingSubmitAggregatedPost**](Flux10EReportingAPI.md#SubmitAggregatedEreportingApiV1EreportingSubmitAggregatedPost) | **Post** /api/v1/ereporting/submit-aggregated | Submit aggregated e-reporting to PA/PDP
[**SubmitEreportingApiV1EreportingSubmitPost**](Flux10EReportingAPI.md#SubmitEreportingApiV1EreportingSubmitPost) | **Post** /api/v1/ereporting/submit | Submit e-reporting to PA/PDP
[**SubmitXmlEreportingApiV1EreportingSubmitXmlPost**](Flux10EReportingAPI.md#SubmitXmlEreportingApiV1EreportingSubmitXmlPost) | **Post** /api/v1/ereporting/submit-xml | Submit pre-generated e-reporting XML
[**ValidateAggregatedEreportingApiV1EreportingValidateAggregatedPost**](Flux10EReportingAPI.md#ValidateAggregatedEreportingApiV1EreportingValidateAggregatedPost) | **Post** /api/v1/ereporting/validate-aggregated | Validate aggregated e-reporting data
[**ValidateEreportingApiV1EreportingValidatePost**](Flux10EReportingAPI.md#ValidateEreportingApiV1EreportingValidatePost) | **Post** /api/v1/ereporting/validate | Validate e-reporting data
[**ValidateXmlEreportingApiV1EreportingValidateXmlPost**](Flux10EReportingAPI.md#ValidateXmlEreportingApiV1EreportingValidateXmlPost) | **Post** /api/v1/ereporting/validate-xml | Validate e-reporting XML (PPF Annexe 6 v1.9 compliant)



## GenerateAggregatedEreportingApiV1EreportingGenerateAggregatedPost

> GenerateAggregatedReportResponse GenerateAggregatedEreportingApiV1EreportingGenerateAggregatedPost(ctx).CreateAggregatedReportRequest(createAggregatedReportRequest).Execute()

Generate aggregated e-reporting XML (PPF-compliant)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/factpulse/sdk-go/v3"
)

func main() {
	createAggregatedReportRequest := *openapiclient.NewCreateAggregatedReportRequest("EREPORT-2025-001", *openapiclient.NewReportSender("123456789", "Ma Société SARL"), *openapiclient.NewReportPeriod(time.Now(), time.Now())) // CreateAggregatedReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Flux10EReportingAPI.GenerateAggregatedEreportingApiV1EreportingGenerateAggregatedPost(context.Background()).CreateAggregatedReportRequest(createAggregatedReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Flux10EReportingAPI.GenerateAggregatedEreportingApiV1EreportingGenerateAggregatedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateAggregatedEreportingApiV1EreportingGenerateAggregatedPost`: GenerateAggregatedReportResponse
	fmt.Fprintf(os.Stdout, "Response from `Flux10EReportingAPI.GenerateAggregatedEreportingApiV1EreportingGenerateAggregatedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateAggregatedEreportingApiV1EreportingGenerateAggregatedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAggregatedReportRequest** | [**CreateAggregatedReportRequest**](CreateAggregatedReportRequest.md) |  | 

### Return type

[**GenerateAggregatedReportResponse**](GenerateAggregatedReportResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateEreportingApiV1EreportingGeneratePost

> GenerateEReportingResponse GenerateEreportingApiV1EreportingGeneratePost(ctx).CreateEReportingRequest(createEReportingRequest).Execute()

Generate e-reporting XML



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/factpulse/sdk-go/v3"
)

func main() {
	createEReportingRequest := *openapiclient.NewCreateEReportingRequest("EREPORT-2025-001", openapiclient.EReportingFlowType("10.1"), *openapiclient.NewReportSender("123456789", "Ma Société SARL"), *openapiclient.NewReportPeriod(time.Now(), time.Now())) // CreateEReportingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Flux10EReportingAPI.GenerateEreportingApiV1EreportingGeneratePost(context.Background()).CreateEReportingRequest(createEReportingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Flux10EReportingAPI.GenerateEreportingApiV1EreportingGeneratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateEreportingApiV1EreportingGeneratePost`: GenerateEReportingResponse
	fmt.Fprintf(os.Stdout, "Response from `Flux10EReportingAPI.GenerateEreportingApiV1EreportingGeneratePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateEreportingApiV1EreportingGeneratePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEReportingRequest** | [**CreateEReportingRequest**](CreateEReportingRequest.md) |  | 

### Return type

[**GenerateEReportingResponse**](GenerateEReportingResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateEreportingDownloadApiV1EreportingGenerateDownloadPost

> GenerateEreportingDownloadApiV1EreportingGenerateDownloadPost(ctx).CreateEReportingRequest(createEReportingRequest).Filename(filename).Execute()

Generate and download e-reporting XML



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/factpulse/sdk-go/v3"
)

func main() {
	createEReportingRequest := *openapiclient.NewCreateEReportingRequest("EREPORT-2025-001", openapiclient.EReportingFlowType("10.1"), *openapiclient.NewReportSender("123456789", "Ma Société SARL"), *openapiclient.NewReportPeriod(time.Now(), time.Now())) // CreateEReportingRequest | 
	filename := "filename_example" // string | Output filename (default: ereporting_{reportId}.xml) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Flux10EReportingAPI.GenerateEreportingDownloadApiV1EreportingGenerateDownloadPost(context.Background()).CreateEReportingRequest(createEReportingRequest).Filename(filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Flux10EReportingAPI.GenerateEreportingDownloadApiV1EreportingGenerateDownloadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateEreportingDownloadApiV1EreportingGenerateDownloadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEReportingRequest** | [**CreateEReportingRequest**](CreateEReportingRequest.md) |  | 
 **filename** | **string** | Output filename (default: ereporting_{reportId}.xml) | 

### Return type

 (empty response body)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCategoryCodesApiV1EreportingCategoryCodesGet

> map[string]interface{} ListCategoryCodesApiV1EreportingCategoryCodesGet(ctx).Execute()

List PPF-compliant category codes



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
	resp, r, err := apiClient.Flux10EReportingAPI.ListCategoryCodesApiV1EreportingCategoryCodesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Flux10EReportingAPI.ListCategoryCodesApiV1EreportingCategoryCodesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCategoryCodesApiV1EreportingCategoryCodesGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Flux10EReportingAPI.ListCategoryCodesApiV1EreportingCategoryCodesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCategoryCodesApiV1EreportingCategoryCodesGetRequest struct via the builder pattern


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


## ListFlowTypesApiV1EreportingFlowTypesGet

> map[string]interface{} ListFlowTypesApiV1EreportingFlowTypesGet(ctx).Execute()

List available flow types



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
	resp, r, err := apiClient.Flux10EReportingAPI.ListFlowTypesApiV1EreportingFlowTypesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Flux10EReportingAPI.ListFlowTypesApiV1EreportingFlowTypesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFlowTypesApiV1EreportingFlowTypesGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Flux10EReportingAPI.ListFlowTypesApiV1EreportingFlowTypesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFlowTypesApiV1EreportingFlowTypesGetRequest struct via the builder pattern


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


## SubmitAggregatedEreportingApiV1EreportingSubmitAggregatedPost

> SubmitEReportingResponse SubmitAggregatedEreportingApiV1EreportingSubmitAggregatedPost(ctx).SubmitAggregatedReportRequest(submitAggregatedReportRequest).Execute()

Submit aggregated e-reporting to PA/PDP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/factpulse/sdk-go/v3"
)

func main() {
	submitAggregatedReportRequest := *openapiclient.NewSubmitAggregatedReportRequest(*openapiclient.NewCreateAggregatedReportRequest("EREPORT-2025-001", *openapiclient.NewReportSender("123456789", "Ma Société SARL"), *openapiclient.NewReportPeriod(time.Now(), time.Now()))) // SubmitAggregatedReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Flux10EReportingAPI.SubmitAggregatedEreportingApiV1EreportingSubmitAggregatedPost(context.Background()).SubmitAggregatedReportRequest(submitAggregatedReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Flux10EReportingAPI.SubmitAggregatedEreportingApiV1EreportingSubmitAggregatedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitAggregatedEreportingApiV1EreportingSubmitAggregatedPost`: SubmitEReportingResponse
	fmt.Fprintf(os.Stdout, "Response from `Flux10EReportingAPI.SubmitAggregatedEreportingApiV1EreportingSubmitAggregatedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitAggregatedEreportingApiV1EreportingSubmitAggregatedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitAggregatedReportRequest** | [**SubmitAggregatedReportRequest**](SubmitAggregatedReportRequest.md) |  | 

### Return type

[**SubmitEReportingResponse**](SubmitEReportingResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitEreportingApiV1EreportingSubmitPost

> SubmitEReportingResponse SubmitEreportingApiV1EreportingSubmitPost(ctx).SubmitEReportingRequest(submitEReportingRequest).Execute()

Submit e-reporting to PA/PDP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/factpulse/sdk-go/v3"
)

func main() {
	submitEReportingRequest := *openapiclient.NewSubmitEReportingRequest(*openapiclient.NewCreateEReportingRequest("EREPORT-2025-001", openapiclient.EReportingFlowType("10.1"), *openapiclient.NewReportSender("123456789", "Ma Société SARL"), *openapiclient.NewReportPeriod(time.Now(), time.Now()))) // SubmitEReportingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Flux10EReportingAPI.SubmitEreportingApiV1EreportingSubmitPost(context.Background()).SubmitEReportingRequest(submitEReportingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Flux10EReportingAPI.SubmitEreportingApiV1EreportingSubmitPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitEreportingApiV1EreportingSubmitPost`: SubmitEReportingResponse
	fmt.Fprintf(os.Stdout, "Response from `Flux10EReportingAPI.SubmitEreportingApiV1EreportingSubmitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitEreportingApiV1EreportingSubmitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitEReportingRequest** | [**SubmitEReportingRequest**](SubmitEReportingRequest.md) |  | 

### Return type

[**SubmitEReportingResponse**](SubmitEReportingResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitXmlEreportingApiV1EreportingSubmitXmlPost

> SubmitEReportingResponse SubmitXmlEreportingApiV1EreportingSubmitXmlPost(ctx).XmlFile(xmlFile).TrackingId(trackingId).SkipValidation(skipValidation).PdpFlowServiceUrl(pdpFlowServiceUrl).PdpTokenUrl(pdpTokenUrl).PdpClientId(pdpClientId).PdpClientSecret(pdpClientSecret).Execute()

Submit pre-generated e-reporting XML



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
	xmlFile := os.NewFile(1234, "some_file") // *os.File | E-reporting XML file
	trackingId := "trackingId_example" // string |  (optional)
	skipValidation := true // bool | Skip XSD validation (optional) (default to false)
	pdpFlowServiceUrl := "pdpFlowServiceUrl_example" // string |  (optional)
	pdpTokenUrl := "pdpTokenUrl_example" // string |  (optional)
	pdpClientId := "pdpClientId_example" // string |  (optional)
	pdpClientSecret := "pdpClientSecret_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Flux10EReportingAPI.SubmitXmlEreportingApiV1EreportingSubmitXmlPost(context.Background()).XmlFile(xmlFile).TrackingId(trackingId).SkipValidation(skipValidation).PdpFlowServiceUrl(pdpFlowServiceUrl).PdpTokenUrl(pdpTokenUrl).PdpClientId(pdpClientId).PdpClientSecret(pdpClientSecret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Flux10EReportingAPI.SubmitXmlEreportingApiV1EreportingSubmitXmlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitXmlEreportingApiV1EreportingSubmitXmlPost`: SubmitEReportingResponse
	fmt.Fprintf(os.Stdout, "Response from `Flux10EReportingAPI.SubmitXmlEreportingApiV1EreportingSubmitXmlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitXmlEreportingApiV1EreportingSubmitXmlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xmlFile** | ***os.File** | E-reporting XML file | 
 **trackingId** | **string** |  | 
 **skipValidation** | **bool** | Skip XSD validation | [default to false]
 **pdpFlowServiceUrl** | **string** |  | 
 **pdpTokenUrl** | **string** |  | 
 **pdpClientId** | **string** |  | 
 **pdpClientSecret** | **string** |  | 

### Return type

[**SubmitEReportingResponse**](SubmitEReportingResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateAggregatedEreportingApiV1EreportingValidateAggregatedPost

> map[string]interface{} ValidateAggregatedEreportingApiV1EreportingValidateAggregatedPost(ctx).CreateAggregatedReportRequest(createAggregatedReportRequest).Execute()

Validate aggregated e-reporting data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/factpulse/sdk-go/v3"
)

func main() {
	createAggregatedReportRequest := *openapiclient.NewCreateAggregatedReportRequest("EREPORT-2025-001", *openapiclient.NewReportSender("123456789", "Ma Société SARL"), *openapiclient.NewReportPeriod(time.Now(), time.Now())) // CreateAggregatedReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Flux10EReportingAPI.ValidateAggregatedEreportingApiV1EreportingValidateAggregatedPost(context.Background()).CreateAggregatedReportRequest(createAggregatedReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Flux10EReportingAPI.ValidateAggregatedEreportingApiV1EreportingValidateAggregatedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateAggregatedEreportingApiV1EreportingValidateAggregatedPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Flux10EReportingAPI.ValidateAggregatedEreportingApiV1EreportingValidateAggregatedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAggregatedEreportingApiV1EreportingValidateAggregatedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAggregatedReportRequest** | [**CreateAggregatedReportRequest**](CreateAggregatedReportRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateEreportingApiV1EreportingValidatePost

> ValidateEReportingResponse ValidateEreportingApiV1EreportingValidatePost(ctx).ValidateEReportingRequest(validateEReportingRequest).Execute()

Validate e-reporting data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/factpulse/sdk-go/v3"
)

func main() {
	validateEReportingRequest := *openapiclient.NewValidateEReportingRequest(*openapiclient.NewCreateEReportingRequest("EREPORT-2025-001", openapiclient.EReportingFlowType("10.1"), *openapiclient.NewReportSender("123456789", "Ma Société SARL"), *openapiclient.NewReportPeriod(time.Now(), time.Now()))) // ValidateEReportingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Flux10EReportingAPI.ValidateEreportingApiV1EreportingValidatePost(context.Background()).ValidateEReportingRequest(validateEReportingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Flux10EReportingAPI.ValidateEreportingApiV1EreportingValidatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateEreportingApiV1EreportingValidatePost`: ValidateEReportingResponse
	fmt.Fprintf(os.Stdout, "Response from `Flux10EReportingAPI.ValidateEreportingApiV1EreportingValidatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateEreportingApiV1EreportingValidatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateEReportingRequest** | [**ValidateEReportingRequest**](ValidateEReportingRequest.md) |  | 

### Return type

[**ValidateEReportingResponse**](ValidateEReportingResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateXmlEreportingApiV1EreportingValidateXmlPost

> map[string]interface{} ValidateXmlEreportingApiV1EreportingValidateXmlPost(ctx).XmlFile(xmlFile).ValidateCoherence(validateCoherence).ValidatePeriod(validatePeriod).Execute()

Validate e-reporting XML (PPF Annexe 6 v1.9 compliant)



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
	xmlFile := os.NewFile(1234, "some_file") // *os.File | E-reporting XML file to validate
	validateCoherence := true // bool | Validate data coherence (REJ_COH) (optional) (default to true)
	validatePeriod := true // bool | Validate period coherence (REJ_PER) (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Flux10EReportingAPI.ValidateXmlEreportingApiV1EreportingValidateXmlPost(context.Background()).XmlFile(xmlFile).ValidateCoherence(validateCoherence).ValidatePeriod(validatePeriod).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Flux10EReportingAPI.ValidateXmlEreportingApiV1EreportingValidateXmlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateXmlEreportingApiV1EreportingValidateXmlPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Flux10EReportingAPI.ValidateXmlEreportingApiV1EreportingValidateXmlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateXmlEreportingApiV1EreportingValidateXmlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xmlFile** | ***os.File** | E-reporting XML file to validate | 
 **validateCoherence** | **bool** | Validate data coherence (REJ_COH) | [default to true]
 **validatePeriod** | **bool** | Validate period coherence (REJ_PER) | [default to true]

### Return type

**map[string]interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


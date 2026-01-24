# \Flux6InvoiceLifecycleCDARAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateCdarApiV1CdarGeneratePost**](Flux6InvoiceLifecycleCDARAPI.md#GenerateCdarApiV1CdarGeneratePost) | **Post** /api/v1/cdar/generate | Generate a CDAR message
[**GetActionCodesApiV1CdarActionCodesGet**](Flux6InvoiceLifecycleCDARAPI.md#GetActionCodesApiV1CdarActionCodesGet) | **Get** /api/v1/cdar/action-codes | List of CDAR action codes
[**GetReasonCodesApiV1CdarReasonCodesGet**](Flux6InvoiceLifecycleCDARAPI.md#GetReasonCodesApiV1CdarReasonCodesGet) | **Get** /api/v1/cdar/reason-codes | List of CDAR reason codes
[**GetStatusCodesApiV1CdarStatusCodesGet**](Flux6InvoiceLifecycleCDARAPI.md#GetStatusCodesApiV1CdarStatusCodesGet) | **Get** /api/v1/cdar/status-codes | List of CDAR status codes
[**SubmitCdarApiV1CdarSubmitPost**](Flux6InvoiceLifecycleCDARAPI.md#SubmitCdarApiV1CdarSubmitPost) | **Post** /api/v1/cdar/submit | Generate and submit a CDAR message
[**SubmitCdarXmlApiV1CdarSubmitXmlPost**](Flux6InvoiceLifecycleCDARAPI.md#SubmitCdarXmlApiV1CdarSubmitXmlPost) | **Post** /api/v1/cdar/submit-xml | Submit a pre-generated CDAR XML
[**SubmitEncaisseeApiV1CdarEncaisseePost**](Flux6InvoiceLifecycleCDARAPI.md#SubmitEncaisseeApiV1CdarEncaisseePost) | **Post** /api/v1/cdar/encaissee | [Simplified] Submit PAID status (212) - Issued invoice
[**SubmitRefuseeApiV1CdarRefuseePost**](Flux6InvoiceLifecycleCDARAPI.md#SubmitRefuseeApiV1CdarRefuseePost) | **Post** /api/v1/cdar/refusee | [Simplified] Submit REFUSED status (210) - Received invoice
[**ValidateCdarApiV1CdarValidatePost**](Flux6InvoiceLifecycleCDARAPI.md#ValidateCdarApiV1CdarValidatePost) | **Post** /api/v1/cdar/validate | Validate CDAR structured data
[**ValidateXmlCdarApiV1CdarValidateXmlPost**](Flux6InvoiceLifecycleCDARAPI.md#ValidateXmlCdarApiV1CdarValidateXmlPost) | **Post** /api/v1/cdar/validate-xml | Validate CDAR XML against XSD and Schematron BR-FR-CDV



## GenerateCdarApiV1CdarGeneratePost

> GenerateCDARResponse GenerateCdarApiV1CdarGeneratePost(ctx).CreateCDARRequest(createCDARRequest).Execute()

Generate a CDAR message



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
	createCDARRequest := *openapiclient.NewCreateCDARRequest("DocumentId_example", "SenderSiren_example", "InvoiceId_example", time.Now(), "Status_example") // CreateCDARRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Flux6InvoiceLifecycleCDARAPI.GenerateCdarApiV1CdarGeneratePost(context.Background()).CreateCDARRequest(createCDARRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Flux6InvoiceLifecycleCDARAPI.GenerateCdarApiV1CdarGeneratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateCdarApiV1CdarGeneratePost`: GenerateCDARResponse
	fmt.Fprintf(os.Stdout, "Response from `Flux6InvoiceLifecycleCDARAPI.GenerateCdarApiV1CdarGeneratePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateCdarApiV1CdarGeneratePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCDARRequest** | [**CreateCDARRequest**](CreateCDARRequest.md) |  | 

### Return type

[**GenerateCDARResponse**](GenerateCDARResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActionCodesApiV1CdarActionCodesGet

> ActionCodesResponse GetActionCodesApiV1CdarActionCodesGet(ctx).Execute()

List of CDAR action codes



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
	resp, r, err := apiClient.Flux6InvoiceLifecycleCDARAPI.GetActionCodesApiV1CdarActionCodesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Flux6InvoiceLifecycleCDARAPI.GetActionCodesApiV1CdarActionCodesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActionCodesApiV1CdarActionCodesGet`: ActionCodesResponse
	fmt.Fprintf(os.Stdout, "Response from `Flux6InvoiceLifecycleCDARAPI.GetActionCodesApiV1CdarActionCodesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetActionCodesApiV1CdarActionCodesGetRequest struct via the builder pattern


### Return type

[**ActionCodesResponse**](ActionCodesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReasonCodesApiV1CdarReasonCodesGet

> ReasonCodesResponse GetReasonCodesApiV1CdarReasonCodesGet(ctx).Execute()

List of CDAR reason codes



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
	resp, r, err := apiClient.Flux6InvoiceLifecycleCDARAPI.GetReasonCodesApiV1CdarReasonCodesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Flux6InvoiceLifecycleCDARAPI.GetReasonCodesApiV1CdarReasonCodesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReasonCodesApiV1CdarReasonCodesGet`: ReasonCodesResponse
	fmt.Fprintf(os.Stdout, "Response from `Flux6InvoiceLifecycleCDARAPI.GetReasonCodesApiV1CdarReasonCodesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetReasonCodesApiV1CdarReasonCodesGetRequest struct via the builder pattern


### Return type

[**ReasonCodesResponse**](ReasonCodesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatusCodesApiV1CdarStatusCodesGet

> StatusCodesResponse GetStatusCodesApiV1CdarStatusCodesGet(ctx).Execute()

List of CDAR status codes



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
	resp, r, err := apiClient.Flux6InvoiceLifecycleCDARAPI.GetStatusCodesApiV1CdarStatusCodesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Flux6InvoiceLifecycleCDARAPI.GetStatusCodesApiV1CdarStatusCodesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatusCodesApiV1CdarStatusCodesGet`: StatusCodesResponse
	fmt.Fprintf(os.Stdout, "Response from `Flux6InvoiceLifecycleCDARAPI.GetStatusCodesApiV1CdarStatusCodesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusCodesApiV1CdarStatusCodesGetRequest struct via the builder pattern


### Return type

[**StatusCodesResponse**](StatusCodesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitCdarApiV1CdarSubmitPost

> SubmitCDARResponse SubmitCdarApiV1CdarSubmitPost(ctx).SubmitCDARRequest(submitCDARRequest).Execute()

Generate and submit a CDAR message



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
	submitCDARRequest := *openapiclient.NewSubmitCDARRequest("DocumentId_example", "SenderSiren_example", "InvoiceId_example", time.Now(), "Status_example") // SubmitCDARRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Flux6InvoiceLifecycleCDARAPI.SubmitCdarApiV1CdarSubmitPost(context.Background()).SubmitCDARRequest(submitCDARRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Flux6InvoiceLifecycleCDARAPI.SubmitCdarApiV1CdarSubmitPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitCdarApiV1CdarSubmitPost`: SubmitCDARResponse
	fmt.Fprintf(os.Stdout, "Response from `Flux6InvoiceLifecycleCDARAPI.SubmitCdarApiV1CdarSubmitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitCdarApiV1CdarSubmitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitCDARRequest** | [**SubmitCDARRequest**](SubmitCDARRequest.md) |  | 

### Return type

[**SubmitCDARResponse**](SubmitCDARResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitCdarXmlApiV1CdarSubmitXmlPost

> SubmitCDARResponse SubmitCdarXmlApiV1CdarSubmitXmlPost(ctx).SubmitCDARXMLRequest(submitCDARXMLRequest).Execute()

Submit a pre-generated CDAR XML



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
	submitCDARXMLRequest := *openapiclient.NewSubmitCDARXMLRequest("Xml_example") // SubmitCDARXMLRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Flux6InvoiceLifecycleCDARAPI.SubmitCdarXmlApiV1CdarSubmitXmlPost(context.Background()).SubmitCDARXMLRequest(submitCDARXMLRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Flux6InvoiceLifecycleCDARAPI.SubmitCdarXmlApiV1CdarSubmitXmlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitCdarXmlApiV1CdarSubmitXmlPost`: SubmitCDARResponse
	fmt.Fprintf(os.Stdout, "Response from `Flux6InvoiceLifecycleCDARAPI.SubmitCdarXmlApiV1CdarSubmitXmlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitCdarXmlApiV1CdarSubmitXmlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitCDARXMLRequest** | [**SubmitCDARXMLRequest**](SubmitCDARXMLRequest.md) |  | 

### Return type

[**SubmitCDARResponse**](SubmitCDARResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitEncaisseeApiV1CdarEncaisseePost

> SimplifiedCDARResponse SubmitEncaisseeApiV1CdarEncaisseePost(ctx).EncaisseeRequest(encaisseeRequest).Execute()

[Simplified] Submit PAID status (212) - Issued invoice



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
	encaisseeRequest := *openapiclient.NewEncaisseeRequest("InvoiceId_example", time.Now(), "InvoiceBuyerSiren_example", "InvoiceBuyerElectronicAddress_example", *openapiclient.NewAmount()) // EncaisseeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Flux6InvoiceLifecycleCDARAPI.SubmitEncaisseeApiV1CdarEncaisseePost(context.Background()).EncaisseeRequest(encaisseeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Flux6InvoiceLifecycleCDARAPI.SubmitEncaisseeApiV1CdarEncaisseePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitEncaisseeApiV1CdarEncaisseePost`: SimplifiedCDARResponse
	fmt.Fprintf(os.Stdout, "Response from `Flux6InvoiceLifecycleCDARAPI.SubmitEncaisseeApiV1CdarEncaisseePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitEncaisseeApiV1CdarEncaisseePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **encaisseeRequest** | [**EncaisseeRequest**](EncaisseeRequest.md) |  | 

### Return type

[**SimplifiedCDARResponse**](SimplifiedCDARResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitRefuseeApiV1CdarRefuseePost

> SimplifiedCDARResponse SubmitRefuseeApiV1CdarRefuseePost(ctx).RefuseeRequest(refuseeRequest).Execute()

[Simplified] Submit REFUSED status (210) - Received invoice



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
	refuseeRequest := *openapiclient.NewRefuseeRequest("InvoiceId_example", time.Now(), "InvoiceSellerSiren_example", "InvoiceSellerElectronicAddress_example", "ReasonCode_example") // RefuseeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Flux6InvoiceLifecycleCDARAPI.SubmitRefuseeApiV1CdarRefuseePost(context.Background()).RefuseeRequest(refuseeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Flux6InvoiceLifecycleCDARAPI.SubmitRefuseeApiV1CdarRefuseePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitRefuseeApiV1CdarRefuseePost`: SimplifiedCDARResponse
	fmt.Fprintf(os.Stdout, "Response from `Flux6InvoiceLifecycleCDARAPI.SubmitRefuseeApiV1CdarRefuseePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitRefuseeApiV1CdarRefuseePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refuseeRequest** | [**RefuseeRequest**](RefuseeRequest.md) |  | 

### Return type

[**SimplifiedCDARResponse**](SimplifiedCDARResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateCdarApiV1CdarValidatePost

> ValidateCDARResponse ValidateCdarApiV1CdarValidatePost(ctx).ValidateCDARRequest(validateCDARRequest).Execute()

Validate CDAR structured data



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
	validateCDARRequest := *openapiclient.NewValidateCDARRequest() // ValidateCDARRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Flux6InvoiceLifecycleCDARAPI.ValidateCdarApiV1CdarValidatePost(context.Background()).ValidateCDARRequest(validateCDARRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Flux6InvoiceLifecycleCDARAPI.ValidateCdarApiV1CdarValidatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateCdarApiV1CdarValidatePost`: ValidateCDARResponse
	fmt.Fprintf(os.Stdout, "Response from `Flux6InvoiceLifecycleCDARAPI.ValidateCdarApiV1CdarValidatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCdarApiV1CdarValidatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateCDARRequest** | [**ValidateCDARRequest**](ValidateCDARRequest.md) |  | 

### Return type

[**ValidateCDARResponse**](ValidateCDARResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateXmlCdarApiV1CdarValidateXmlPost

> map[string]interface{} ValidateXmlCdarApiV1CdarValidateXmlPost(ctx).XmlFile(xmlFile).Execute()

Validate CDAR XML against XSD and Schematron BR-FR-CDV



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
	xmlFile := os.NewFile(1234, "some_file") // *os.File | CDAR XML file to validate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Flux6InvoiceLifecycleCDARAPI.ValidateXmlCdarApiV1CdarValidateXmlPost(context.Background()).XmlFile(xmlFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Flux6InvoiceLifecycleCDARAPI.ValidateXmlCdarApiV1CdarValidateXmlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateXmlCdarApiV1CdarValidateXmlPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Flux6InvoiceLifecycleCDARAPI.ValidateXmlCdarApiV1CdarValidateXmlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateXmlCdarApiV1CdarValidateXmlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xmlFile** | ***os.File** | CDAR XML file to validate | 

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


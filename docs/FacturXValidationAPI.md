# \FacturXValidationAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidateFacturxPdfApiV1ProcessingValidateFacturxPdfPost**](FacturXValidationAPI.md#ValidateFacturxPdfApiV1ProcessingValidateFacturxPdfPost) | **Post** /api/v1/processing/validate-facturx-pdf | Validate a complete Factur-X PDF
[**ValidateFacturxPdfAsyncApiV1ProcessingValidateFacturxAsyncPost**](FacturXValidationAPI.md#ValidateFacturxPdfAsyncApiV1ProcessingValidateFacturxAsyncPost) | **Post** /api/v1/processing/validate-facturx-async | Validate a Factur-X PDF (asynchronous with polling)
[**ValidateXmlApiV1ProcessingValidateXmlPost**](FacturXValidationAPI.md#ValidateXmlApiV1ProcessingValidateXmlPost) | **Post** /api/v1/processing/validate-xml | Validate an existing Factur-X XML



## ValidateFacturxPdfApiV1ProcessingValidateFacturxPdfPost

> PDFValidationResultAPI ValidateFacturxPdfApiV1ProcessingValidateFacturxPdfPost(ctx).PdfFile(pdfFile).Profile(profile).UseVerapdf(useVerapdf).SkipBrFr(skipBrFr).Execute()

Validate a complete Factur-X PDF



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
	pdfFile := os.NewFile(1234, "some_file") // *os.File | Factur-X PDF file to validate (.pdf format).
	profile := openapiclient.APIProfile("MINIMUM") // APIProfile |  (optional)
	useVerapdf := true // bool | Enable strict PDF/A validation with VeraPDF (recommended for production). If False, uses basic metadata validation. (optional) (default to false)
	skipBrFr := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FacturXValidationAPI.ValidateFacturxPdfApiV1ProcessingValidateFacturxPdfPost(context.Background()).PdfFile(pdfFile).Profile(profile).UseVerapdf(useVerapdf).SkipBrFr(skipBrFr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FacturXValidationAPI.ValidateFacturxPdfApiV1ProcessingValidateFacturxPdfPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateFacturxPdfApiV1ProcessingValidateFacturxPdfPost`: PDFValidationResultAPI
	fmt.Fprintf(os.Stdout, "Response from `FacturXValidationAPI.ValidateFacturxPdfApiV1ProcessingValidateFacturxPdfPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateFacturxPdfApiV1ProcessingValidateFacturxPdfPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pdfFile** | ***os.File** | Factur-X PDF file to validate (.pdf format). | 
 **profile** | [**APIProfile**](APIProfile.md) |  | 
 **useVerapdf** | **bool** | Enable strict PDF/A validation with VeraPDF (recommended for production). If False, uses basic metadata validation. | [default to false]
 **skipBrFr** | **bool** |  | 

### Return type

[**PDFValidationResultAPI**](PDFValidationResultAPI.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateFacturxPdfAsyncApiV1ProcessingValidateFacturxAsyncPost

> TaskResponse ValidateFacturxPdfAsyncApiV1ProcessingValidateFacturxAsyncPost(ctx).PdfFile(pdfFile).Profile(profile).UseVerapdf(useVerapdf).CallbackUrl(callbackUrl).WebhookMode(webhookMode).Execute()

Validate a Factur-X PDF (asynchronous with polling)



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
	pdfFile := os.NewFile(1234, "some_file") // *os.File | Factur-X PDF file to validate (.pdf format).
	profile := openapiclient.APIProfile("MINIMUM") // APIProfile |  (optional)
	useVerapdf := true // bool | Enable strict PDF/A validation with VeraPDF (recommended for production). May take several seconds. (optional) (default to false)
	callbackUrl := "callbackUrl_example" // string |  (optional)
	webhookMode := "webhookMode_example" // string | Webhook content delivery: 'inline' (base64 in payload) or 'download_url' (temporary URL, 1h TTL) (optional) (default to "inline")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FacturXValidationAPI.ValidateFacturxPdfAsyncApiV1ProcessingValidateFacturxAsyncPost(context.Background()).PdfFile(pdfFile).Profile(profile).UseVerapdf(useVerapdf).CallbackUrl(callbackUrl).WebhookMode(webhookMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FacturXValidationAPI.ValidateFacturxPdfAsyncApiV1ProcessingValidateFacturxAsyncPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateFacturxPdfAsyncApiV1ProcessingValidateFacturxAsyncPost`: TaskResponse
	fmt.Fprintf(os.Stdout, "Response from `FacturXValidationAPI.ValidateFacturxPdfAsyncApiV1ProcessingValidateFacturxAsyncPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateFacturxPdfAsyncApiV1ProcessingValidateFacturxAsyncPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pdfFile** | ***os.File** | Factur-X PDF file to validate (.pdf format). | 
 **profile** | [**APIProfile**](APIProfile.md) |  | 
 **useVerapdf** | **bool** | Enable strict PDF/A validation with VeraPDF (recommended for production). May take several seconds. | [default to false]
 **callbackUrl** | **string** |  | 
 **webhookMode** | **string** | Webhook content delivery: &#39;inline&#39; (base64 in payload) or &#39;download_url&#39; (temporary URL, 1h TTL) | [default to &quot;inline&quot;]

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateXmlApiV1ProcessingValidateXmlPost

> ValidationSuccessResponse ValidateXmlApiV1ProcessingValidateXmlPost(ctx).XmlFile(xmlFile).Profile(profile).SkipBrFr(skipBrFr).Execute()

Validate an existing Factur-X XML



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
	xmlFile := os.NewFile(1234, "some_file") // *os.File | Factur-X XML file to validate (.xml format).
	profile := openapiclient.APIProfile("MINIMUM") // APIProfile | Validation profile (MINIMUM, BASIC, EN16931, EXTENDED). (optional) (default to "EXTENDED")
	skipBrFr := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FacturXValidationAPI.ValidateXmlApiV1ProcessingValidateXmlPost(context.Background()).XmlFile(xmlFile).Profile(profile).SkipBrFr(skipBrFr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FacturXValidationAPI.ValidateXmlApiV1ProcessingValidateXmlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateXmlApiV1ProcessingValidateXmlPost`: ValidationSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `FacturXValidationAPI.ValidateXmlApiV1ProcessingValidateXmlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateXmlApiV1ProcessingValidateXmlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xmlFile** | ***os.File** | Factur-X XML file to validate (.xml format). | 
 **profile** | [**APIProfile**](APIProfile.md) | Validation profile (MINIMUM, BASIC, EN16931, EXTENDED). | [default to &quot;EXTENDED&quot;]
 **skipBrFr** | **bool** |  | 

### Return type

[**ValidationSuccessResponse**](ValidationSuccessResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


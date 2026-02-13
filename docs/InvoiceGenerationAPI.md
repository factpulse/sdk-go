# \InvoiceGenerationAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateInvoiceApiV1ProcessingGenerateInvoicePost**](InvoiceGenerationAPI.md#GenerateInvoiceApiV1ProcessingGenerateInvoicePost) | **Post** /api/v1/processing/generate-invoice | Generate an electronic invoice (CII / UBL / Factur-X PDF)
[**SubmitCompleteInvoiceApiV1ProcessingInvoicesSubmitCompletePost**](InvoiceGenerationAPI.md#SubmitCompleteInvoiceApiV1ProcessingInvoicesSubmitCompletePost) | **Post** /api/v1/processing/invoices/submit-complete | Submit a complete invoice (generation + signature + submission)
[**SubmitCompleteInvoiceAsyncApiV1ProcessingInvoicesSubmitCompleteAsyncPost**](InvoiceGenerationAPI.md#SubmitCompleteInvoiceAsyncApiV1ProcessingInvoicesSubmitCompleteAsyncPost) | **Post** /api/v1/processing/invoices/submit-complete-async | Submit a complete invoice (asynchronous with Celery)



## GenerateInvoiceApiV1ProcessingGenerateInvoicePost

> TaskResponse GenerateInvoiceApiV1ProcessingGenerateInvoicePost(ctx).InvoiceData(invoiceData).Profile(profile).OutputFormat(outputFormat).AutoEnrich(autoEnrich).SourcePdf(sourcePdf).CallbackUrl(callbackUrl).WebhookMode(webhookMode).SkipBrFr(skipBrFr).Execute()

Generate an electronic invoice (CII / UBL / Factur-X PDF)



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
	invoiceData := "invoiceData_example" // string | Invoice data in JSON format.              Two formats accepted:             1. **Classic format**: Complete FacturXInvoice structure (all fields)             2. **Simplified format** (🆕 P0.1): Minimal structure with auto-enrichment              Format is detected automatically!             
	profile := openapiclient.APIProfile("MINIMUM") // APIProfile | Factur-X/CII profile: MINIMUM, BASIC, EN16931 or EXTENDED. Ignored when output_format='ubl' (always EN16931). (optional) (default to "EXTENDED")
	outputFormat := openapiclient.OutputFormat("xml") // OutputFormat | Output format: 'xml' or 'cii' (CII/Factur-X XML), 'ubl' (UBL 2.1 XML), 'pdf' (Factur-X PDF/A-3). (optional) (default to "xml")
	autoEnrich := true // bool | 🆕 Enable auto-enrichment from SIRET/SIREN (simplified format only) (optional) (default to true)
	sourcePdf := os.NewFile(1234, "some_file") // *os.File |  (optional)
	callbackUrl := "callbackUrl_example" // string |  (optional)
	webhookMode := "webhookMode_example" // string | Webhook content delivery: 'inline' (base64 in payload) or 'download_url' (temporary URL, 1h TTL) (optional) (default to "inline")
	skipBrFr := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceGenerationAPI.GenerateInvoiceApiV1ProcessingGenerateInvoicePost(context.Background()).InvoiceData(invoiceData).Profile(profile).OutputFormat(outputFormat).AutoEnrich(autoEnrich).SourcePdf(sourcePdf).CallbackUrl(callbackUrl).WebhookMode(webhookMode).SkipBrFr(skipBrFr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceGenerationAPI.GenerateInvoiceApiV1ProcessingGenerateInvoicePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateInvoiceApiV1ProcessingGenerateInvoicePost`: TaskResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoiceGenerationAPI.GenerateInvoiceApiV1ProcessingGenerateInvoicePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateInvoiceApiV1ProcessingGenerateInvoicePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceData** | **string** | Invoice data in JSON format.              Two formats accepted:             1. **Classic format**: Complete FacturXInvoice structure (all fields)             2. **Simplified format** (🆕 P0.1): Minimal structure with auto-enrichment              Format is detected automatically!              | 
 **profile** | [**APIProfile**](APIProfile.md) | Factur-X/CII profile: MINIMUM, BASIC, EN16931 or EXTENDED. Ignored when output_format&#x3D;&#39;ubl&#39; (always EN16931). | [default to &quot;EXTENDED&quot;]
 **outputFormat** | [**OutputFormat**](OutputFormat.md) | Output format: &#39;xml&#39; or &#39;cii&#39; (CII/Factur-X XML), &#39;ubl&#39; (UBL 2.1 XML), &#39;pdf&#39; (Factur-X PDF/A-3). | [default to &quot;xml&quot;]
 **autoEnrich** | **bool** | 🆕 Enable auto-enrichment from SIRET/SIREN (simplified format only) | [default to true]
 **sourcePdf** | ***os.File** |  | 
 **callbackUrl** | **string** |  | 
 **webhookMode** | **string** | Webhook content delivery: &#39;inline&#39; (base64 in payload) or &#39;download_url&#39; (temporary URL, 1h TTL) | [default to &quot;inline&quot;]
 **skipBrFr** | **bool** |  | 

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitCompleteInvoiceApiV1ProcessingInvoicesSubmitCompletePost

> SubmitCompleteInvoiceResponse SubmitCompleteInvoiceApiV1ProcessingInvoicesSubmitCompletePost(ctx).SubmitCompleteInvoiceRequest(submitCompleteInvoiceRequest).Execute()

Submit a complete invoice (generation + signature + submission)



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
	submitCompleteInvoiceRequest := *openapiclient.NewSubmitCompleteInvoiceRequest(*openapiclient.NewSimplifiedInvoiceData("Number_example", map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}), "SourcePdf_example", openapiclient.Destination{AFNORDestination: openapiclient.NewAFNORDestination()}) // SubmitCompleteInvoiceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceGenerationAPI.SubmitCompleteInvoiceApiV1ProcessingInvoicesSubmitCompletePost(context.Background()).SubmitCompleteInvoiceRequest(submitCompleteInvoiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceGenerationAPI.SubmitCompleteInvoiceApiV1ProcessingInvoicesSubmitCompletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitCompleteInvoiceApiV1ProcessingInvoicesSubmitCompletePost`: SubmitCompleteInvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoiceGenerationAPI.SubmitCompleteInvoiceApiV1ProcessingInvoicesSubmitCompletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitCompleteInvoiceApiV1ProcessingInvoicesSubmitCompletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitCompleteInvoiceRequest** | [**SubmitCompleteInvoiceRequest**](SubmitCompleteInvoiceRequest.md) |  | 

### Return type

[**SubmitCompleteInvoiceResponse**](SubmitCompleteInvoiceResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitCompleteInvoiceAsyncApiV1ProcessingInvoicesSubmitCompleteAsyncPost

> TaskResponse SubmitCompleteInvoiceAsyncApiV1ProcessingInvoicesSubmitCompleteAsyncPost(ctx).SubmitCompleteInvoiceRequest(submitCompleteInvoiceRequest).CallbackUrl(callbackUrl).WebhookMode(webhookMode).Execute()

Submit a complete invoice (asynchronous with Celery)



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
	submitCompleteInvoiceRequest := *openapiclient.NewSubmitCompleteInvoiceRequest(*openapiclient.NewSimplifiedInvoiceData("Number_example", map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}), "SourcePdf_example", openapiclient.Destination{AFNORDestination: openapiclient.NewAFNORDestination()}) // SubmitCompleteInvoiceRequest | 
	callbackUrl := "callbackUrl_example" // string | Webhook URL for async notification when submission completes. (optional)
	webhookMode := "webhookMode_example" // string | Webhook content delivery: 'inline' (base64 in payload) or 'download_url' (temporary URL, 1h TTL) (optional) (default to "inline")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceGenerationAPI.SubmitCompleteInvoiceAsyncApiV1ProcessingInvoicesSubmitCompleteAsyncPost(context.Background()).SubmitCompleteInvoiceRequest(submitCompleteInvoiceRequest).CallbackUrl(callbackUrl).WebhookMode(webhookMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceGenerationAPI.SubmitCompleteInvoiceAsyncApiV1ProcessingInvoicesSubmitCompleteAsyncPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitCompleteInvoiceAsyncApiV1ProcessingInvoicesSubmitCompleteAsyncPost`: TaskResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoiceGenerationAPI.SubmitCompleteInvoiceAsyncApiV1ProcessingInvoicesSubmitCompleteAsyncPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitCompleteInvoiceAsyncApiV1ProcessingInvoicesSubmitCompleteAsyncPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitCompleteInvoiceRequest** | [**SubmitCompleteInvoiceRequest**](SubmitCompleteInvoiceRequest.md) |  | 
 **callbackUrl** | **string** | Webhook URL for async notification when submission completes. | 
 **webhookMode** | **string** | Webhook content delivery: &#39;inline&#39; (base64 in payload) or &#39;download_url&#39; (temporary URL, 1h TTL) | [default to &quot;inline&quot;]

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


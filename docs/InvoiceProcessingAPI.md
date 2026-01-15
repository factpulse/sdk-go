# \InvoiceProcessingAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateInvoiceApiV1ProcessingGenerateInvoicePost**](InvoiceProcessingAPI.md#GenerateInvoiceApiV1ProcessingGenerateInvoicePost) | **Post** /api/v1/processing/generate-invoice | Generate a Factur-X invoice
[**GenerateTestCertificateApiV1ProcessingGenerateTestCertificatePost**](InvoiceProcessingAPI.md#GenerateTestCertificateApiV1ProcessingGenerateTestCertificatePost) | **Post** /api/v1/processing/generate-test-certificate | Generate a self-signed X.509 test certificate
[**GetTaskStatusApiV1ProcessingTasksTaskIdStatusGet**](InvoiceProcessingAPI.md#GetTaskStatusApiV1ProcessingTasksTaskIdStatusGet) | **Get** /api/v1/processing/tasks/{task_id}/status | Get task generation status
[**SignPdfApiV1ProcessingSignPdfPost**](InvoiceProcessingAPI.md#SignPdfApiV1ProcessingSignPdfPost) | **Post** /api/v1/processing/sign-pdf | Sign a PDF with client&#39;s certificate (PAdES-B-LT)
[**SignPdfAsyncApiV1ProcessingSignPdfAsyncPost**](InvoiceProcessingAPI.md#SignPdfAsyncApiV1ProcessingSignPdfAsyncPost) | **Post** /api/v1/processing/sign-pdf-async | Sign a PDF asynchronously (Celery)
[**SubmitCompleteInvoiceApiV1ProcessingInvoicesSubmitCompletePost**](InvoiceProcessingAPI.md#SubmitCompleteInvoiceApiV1ProcessingInvoicesSubmitCompletePost) | **Post** /api/v1/processing/invoices/submit-complete | Submit a complete invoice (generation + signature + submission)
[**SubmitCompleteInvoiceAsyncApiV1ProcessingInvoicesSubmitCompleteAsyncPost**](InvoiceProcessingAPI.md#SubmitCompleteInvoiceAsyncApiV1ProcessingInvoicesSubmitCompleteAsyncPost) | **Post** /api/v1/processing/invoices/submit-complete-async | Submit a complete invoice (asynchronous with Celery)
[**ValidateFacturxPdfApiV1ProcessingValidateFacturxPdfPost**](InvoiceProcessingAPI.md#ValidateFacturxPdfApiV1ProcessingValidateFacturxPdfPost) | **Post** /api/v1/processing/validate-facturx-pdf | Validate a complete Factur-X PDF
[**ValidateFacturxPdfAsyncApiV1ProcessingValidateFacturxAsyncPost**](InvoiceProcessingAPI.md#ValidateFacturxPdfAsyncApiV1ProcessingValidateFacturxAsyncPost) | **Post** /api/v1/processing/validate-facturx-async | Validate a Factur-X PDF (asynchronous with polling)
[**ValidatePdfSignatureEndpointApiV1ProcessingValidatePdfSignaturePost**](InvoiceProcessingAPI.md#ValidatePdfSignatureEndpointApiV1ProcessingValidatePdfSignaturePost) | **Post** /api/v1/processing/validate-pdf-signature | Validate electronic signatures of a PDF
[**ValidateXmlApiV1ProcessingValidateXmlPost**](InvoiceProcessingAPI.md#ValidateXmlApiV1ProcessingValidateXmlPost) | **Post** /api/v1/processing/validate-xml | Validate an existing Factur-X XML



## GenerateInvoiceApiV1ProcessingGenerateInvoicePost

> TaskResponse GenerateInvoiceApiV1ProcessingGenerateInvoicePost(ctx).InvoiceData(invoiceData).Profile(profile).OutputFormat(outputFormat).AutoEnrich(autoEnrich).SourcePdf(sourcePdf).CallbackUrl(callbackUrl).WebhookMode(webhookMode).SkipBrFr(skipBrFr).Execute()

Generate a Factur-X invoice



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
	profile := openapiclient.APIProfile("MINIMUM") // APIProfile | Factur-X profile: MINIMUM, BASIC, EN16931 or EXTENDED. (optional) (default to "EXTENDED")
	outputFormat := openapiclient.OutputFormat("xml") // OutputFormat | Output format: 'xml' (XML only) or 'pdf' (Factur-X PDF with embedded XML). (optional) (default to "xml")
	autoEnrich := true // bool | 🆕 Enable auto-enrichment from SIRET/SIREN (simplified format only) (optional) (default to true)
	sourcePdf := os.NewFile(1234, "some_file") // *os.File |  (optional)
	callbackUrl := "callbackUrl_example" // string |  (optional)
	webhookMode := "webhookMode_example" // string | Webhook content delivery: 'inline' (base64 in payload) or 'download_url' (temporary URL, 1h TTL) (optional) (default to "inline")
	skipBrFr := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceProcessingAPI.GenerateInvoiceApiV1ProcessingGenerateInvoicePost(context.Background()).InvoiceData(invoiceData).Profile(profile).OutputFormat(outputFormat).AutoEnrich(autoEnrich).SourcePdf(sourcePdf).CallbackUrl(callbackUrl).WebhookMode(webhookMode).SkipBrFr(skipBrFr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceProcessingAPI.GenerateInvoiceApiV1ProcessingGenerateInvoicePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateInvoiceApiV1ProcessingGenerateInvoicePost`: TaskResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoiceProcessingAPI.GenerateInvoiceApiV1ProcessingGenerateInvoicePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateInvoiceApiV1ProcessingGenerateInvoicePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceData** | **string** | Invoice data in JSON format.              Two formats accepted:             1. **Classic format**: Complete FacturXInvoice structure (all fields)             2. **Simplified format** (🆕 P0.1): Minimal structure with auto-enrichment              Format is detected automatically!              | 
 **profile** | [**APIProfile**](APIProfile.md) | Factur-X profile: MINIMUM, BASIC, EN16931 or EXTENDED. | [default to &quot;EXTENDED&quot;]
 **outputFormat** | [**OutputFormat**](OutputFormat.md) | Output format: &#39;xml&#39; (XML only) or &#39;pdf&#39; (Factur-X PDF with embedded XML). | [default to &quot;xml&quot;]
 **autoEnrich** | **bool** | 🆕 Enable auto-enrichment from SIRET/SIREN (simplified format only) | [default to true]
 **sourcePdf** | ***os.File** |  | 
 **callbackUrl** | **string** |  | 
 **webhookMode** | **string** | Webhook content delivery: &#39;inline&#39; (base64 in payload) or &#39;download_url&#39; (temporary URL, 1h TTL) | [default to &quot;inline&quot;]
 **skipBrFr** | **bool** |  | 

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


## GenerateTestCertificateApiV1ProcessingGenerateTestCertificatePost

> GenerateCertificateResponse GenerateTestCertificateApiV1ProcessingGenerateTestCertificatePost(ctx).GenerateCertificateRequest(generateCertificateRequest).Execute()

Generate a self-signed X.509 test certificate



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
	generateCertificateRequest := *openapiclient.NewGenerateCertificateRequest() // GenerateCertificateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceProcessingAPI.GenerateTestCertificateApiV1ProcessingGenerateTestCertificatePost(context.Background()).GenerateCertificateRequest(generateCertificateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceProcessingAPI.GenerateTestCertificateApiV1ProcessingGenerateTestCertificatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateTestCertificateApiV1ProcessingGenerateTestCertificatePost`: GenerateCertificateResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoiceProcessingAPI.GenerateTestCertificateApiV1ProcessingGenerateTestCertificatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateTestCertificateApiV1ProcessingGenerateTestCertificatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateCertificateRequest** | [**GenerateCertificateRequest**](GenerateCertificateRequest.md) |  | 

### Return type

[**GenerateCertificateResponse**](GenerateCertificateResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskStatusApiV1ProcessingTasksTaskIdStatusGet

> AsyncTaskStatus GetTaskStatusApiV1ProcessingTasksTaskIdStatusGet(ctx, taskId).Execute()

Get task generation status



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
	taskId := "taskId_example" // string | Celery task ID returned by async endpoints (UUID format)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceProcessingAPI.GetTaskStatusApiV1ProcessingTasksTaskIdStatusGet(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceProcessingAPI.GetTaskStatusApiV1ProcessingTasksTaskIdStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskStatusApiV1ProcessingTasksTaskIdStatusGet`: AsyncTaskStatus
	fmt.Fprintf(os.Stdout, "Response from `InvoiceProcessingAPI.GetTaskStatusApiV1ProcessingTasksTaskIdStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | Celery task ID returned by async endpoints (UUID format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskStatusApiV1ProcessingTasksTaskIdStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncTaskStatus**](AsyncTaskStatus.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignPdfApiV1ProcessingSignPdfPost

> interface{} SignPdfApiV1ProcessingSignPdfPost(ctx).PdfFile(pdfFile).Reason(reason).Location(location).Contact(contact).FieldName(fieldName).UsePadesLt(usePadesLt).UseTimestamp(useTimestamp).Execute()

Sign a PDF with client's certificate (PAdES-B-LT)



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
	pdfFile := os.NewFile(1234, "some_file") // *os.File | PDF file to sign (will be processed and returned signed in base64)
	reason := "reason_example" // string |  (optional)
	location := "location_example" // string |  (optional)
	contact := "contact_example" // string |  (optional)
	fieldName := "fieldName_example" // string | PDF signature field name (optional) (default to "FactPulseSignature")
	usePadesLt := true // bool | Enable PAdES-B-LT (long-term archiving with embedded validation data). REQUIRES a certificate with OCSP/CRL access. (optional) (default to false)
	useTimestamp := true // bool | Enable RFC 3161 timestamping with FreeTSA (PAdES-B-T) (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceProcessingAPI.SignPdfApiV1ProcessingSignPdfPost(context.Background()).PdfFile(pdfFile).Reason(reason).Location(location).Contact(contact).FieldName(fieldName).UsePadesLt(usePadesLt).UseTimestamp(useTimestamp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceProcessingAPI.SignPdfApiV1ProcessingSignPdfPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignPdfApiV1ProcessingSignPdfPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InvoiceProcessingAPI.SignPdfApiV1ProcessingSignPdfPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignPdfApiV1ProcessingSignPdfPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pdfFile** | ***os.File** | PDF file to sign (will be processed and returned signed in base64) | 
 **reason** | **string** |  | 
 **location** | **string** |  | 
 **contact** | **string** |  | 
 **fieldName** | **string** | PDF signature field name | [default to &quot;FactPulseSignature&quot;]
 **usePadesLt** | **bool** | Enable PAdES-B-LT (long-term archiving with embedded validation data). REQUIRES a certificate with OCSP/CRL access. | [default to false]
 **useTimestamp** | **bool** | Enable RFC 3161 timestamping with FreeTSA (PAdES-B-T) | [default to true]

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignPdfAsyncApiV1ProcessingSignPdfAsyncPost

> interface{} SignPdfAsyncApiV1ProcessingSignPdfAsyncPost(ctx).PdfFile(pdfFile).CallbackUrl(callbackUrl).WebhookMode(webhookMode).Reason(reason).Location(location).Contact(contact).FieldName(fieldName).UsePadesLt(usePadesLt).UseTimestamp(useTimestamp).Execute()

Sign a PDF asynchronously (Celery)



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
	pdfFile := os.NewFile(1234, "some_file") // *os.File | PDF file to sign (processed asynchronously)
	callbackUrl := "callbackUrl_example" // string |  (optional)
	webhookMode := "webhookMode_example" // string | Webhook content delivery: 'inline' (base64 in payload) or 'download_url' (temporary URL, 1h TTL) (optional) (default to "inline")
	reason := "reason_example" // string |  (optional)
	location := "location_example" // string |  (optional)
	contact := "contact_example" // string |  (optional)
	fieldName := "fieldName_example" // string | PDF signature field name (optional) (default to "FactPulseSignature")
	usePadesLt := true // bool | Enable PAdES-B-LT (long-term archiving with embedded validation data). REQUIRES a certificate with OCSP/CRL access. (optional) (default to false)
	useTimestamp := true // bool | Enable RFC 3161 timestamping with FreeTSA (PAdES-B-T) (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceProcessingAPI.SignPdfAsyncApiV1ProcessingSignPdfAsyncPost(context.Background()).PdfFile(pdfFile).CallbackUrl(callbackUrl).WebhookMode(webhookMode).Reason(reason).Location(location).Contact(contact).FieldName(fieldName).UsePadesLt(usePadesLt).UseTimestamp(useTimestamp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceProcessingAPI.SignPdfAsyncApiV1ProcessingSignPdfAsyncPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignPdfAsyncApiV1ProcessingSignPdfAsyncPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InvoiceProcessingAPI.SignPdfAsyncApiV1ProcessingSignPdfAsyncPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignPdfAsyncApiV1ProcessingSignPdfAsyncPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pdfFile** | ***os.File** | PDF file to sign (processed asynchronously) | 
 **callbackUrl** | **string** |  | 
 **webhookMode** | **string** | Webhook content delivery: &#39;inline&#39; (base64 in payload) or &#39;download_url&#39; (temporary URL, 1h TTL) | [default to &quot;inline&quot;]
 **reason** | **string** |  | 
 **location** | **string** |  | 
 **contact** | **string** |  | 
 **fieldName** | **string** | PDF signature field name | [default to &quot;FactPulseSignature&quot;]
 **usePadesLt** | **bool** | Enable PAdES-B-LT (long-term archiving with embedded validation data). REQUIRES a certificate with OCSP/CRL access. | [default to false]
 **useTimestamp** | **bool** | Enable RFC 3161 timestamping with FreeTSA (PAdES-B-T) | [default to true]

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

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
	resp, r, err := apiClient.InvoiceProcessingAPI.SubmitCompleteInvoiceApiV1ProcessingInvoicesSubmitCompletePost(context.Background()).SubmitCompleteInvoiceRequest(submitCompleteInvoiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceProcessingAPI.SubmitCompleteInvoiceApiV1ProcessingInvoicesSubmitCompletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitCompleteInvoiceApiV1ProcessingInvoicesSubmitCompletePost`: SubmitCompleteInvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoiceProcessingAPI.SubmitCompleteInvoiceApiV1ProcessingInvoicesSubmitCompletePost`: %v\n", resp)
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

[HTTPBearer](../README.md#HTTPBearer)

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
	resp, r, err := apiClient.InvoiceProcessingAPI.SubmitCompleteInvoiceAsyncApiV1ProcessingInvoicesSubmitCompleteAsyncPost(context.Background()).SubmitCompleteInvoiceRequest(submitCompleteInvoiceRequest).CallbackUrl(callbackUrl).WebhookMode(webhookMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceProcessingAPI.SubmitCompleteInvoiceAsyncApiV1ProcessingInvoicesSubmitCompleteAsyncPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitCompleteInvoiceAsyncApiV1ProcessingInvoicesSubmitCompleteAsyncPost`: TaskResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoiceProcessingAPI.SubmitCompleteInvoiceAsyncApiV1ProcessingInvoicesSubmitCompleteAsyncPost`: %v\n", resp)
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

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.InvoiceProcessingAPI.ValidateFacturxPdfApiV1ProcessingValidateFacturxPdfPost(context.Background()).PdfFile(pdfFile).Profile(profile).UseVerapdf(useVerapdf).SkipBrFr(skipBrFr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceProcessingAPI.ValidateFacturxPdfApiV1ProcessingValidateFacturxPdfPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateFacturxPdfApiV1ProcessingValidateFacturxPdfPost`: PDFValidationResultAPI
	fmt.Fprintf(os.Stdout, "Response from `InvoiceProcessingAPI.ValidateFacturxPdfApiV1ProcessingValidateFacturxPdfPost`: %v\n", resp)
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
	resp, r, err := apiClient.InvoiceProcessingAPI.ValidateFacturxPdfAsyncApiV1ProcessingValidateFacturxAsyncPost(context.Background()).PdfFile(pdfFile).Profile(profile).UseVerapdf(useVerapdf).CallbackUrl(callbackUrl).WebhookMode(webhookMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceProcessingAPI.ValidateFacturxPdfAsyncApiV1ProcessingValidateFacturxAsyncPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateFacturxPdfAsyncApiV1ProcessingValidateFacturxAsyncPost`: TaskResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoiceProcessingAPI.ValidateFacturxPdfAsyncApiV1ProcessingValidateFacturxAsyncPost`: %v\n", resp)
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


## ValidatePdfSignatureEndpointApiV1ProcessingValidatePdfSignaturePost

> interface{} ValidatePdfSignatureEndpointApiV1ProcessingValidatePdfSignaturePost(ctx).PdfFile(pdfFile).Execute()

Validate electronic signatures of a PDF



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
	pdfFile := os.NewFile(1234, "some_file") // *os.File | PDF file to validate (will be analyzed to detect and validate signatures)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceProcessingAPI.ValidatePdfSignatureEndpointApiV1ProcessingValidatePdfSignaturePost(context.Background()).PdfFile(pdfFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceProcessingAPI.ValidatePdfSignatureEndpointApiV1ProcessingValidatePdfSignaturePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidatePdfSignatureEndpointApiV1ProcessingValidatePdfSignaturePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InvoiceProcessingAPI.ValidatePdfSignatureEndpointApiV1ProcessingValidatePdfSignaturePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidatePdfSignatureEndpointApiV1ProcessingValidatePdfSignaturePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pdfFile** | ***os.File** | PDF file to validate (will be analyzed to detect and validate signatures) | 

### Return type

**interface{}**

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
	resp, r, err := apiClient.InvoiceProcessingAPI.ValidateXmlApiV1ProcessingValidateXmlPost(context.Background()).XmlFile(xmlFile).Profile(profile).SkipBrFr(skipBrFr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceProcessingAPI.ValidateXmlApiV1ProcessingValidateXmlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateXmlApiV1ProcessingValidateXmlPost`: ValidationSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoiceProcessingAPI.ValidateXmlApiV1ProcessingValidateXmlPost`: %v\n", resp)
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


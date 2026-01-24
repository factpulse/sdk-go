# \ElectronicSignatureAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateTestCertificateApiV1ProcessingGenerateTestCertificatePost**](ElectronicSignatureAPI.md#GenerateTestCertificateApiV1ProcessingGenerateTestCertificatePost) | **Post** /api/v1/processing/generate-test-certificate | Generate a self-signed X.509 test certificate
[**SignPdfApiV1ProcessingSignPdfPost**](ElectronicSignatureAPI.md#SignPdfApiV1ProcessingSignPdfPost) | **Post** /api/v1/processing/sign-pdf | Sign a PDF with client&#39;s certificate (PAdES-B-LT)
[**SignPdfAsyncApiV1ProcessingSignPdfAsyncPost**](ElectronicSignatureAPI.md#SignPdfAsyncApiV1ProcessingSignPdfAsyncPost) | **Post** /api/v1/processing/sign-pdf-async | Sign a PDF asynchronously (Celery)
[**ValidatePdfSignatureEndpointApiV1ProcessingValidatePdfSignaturePost**](ElectronicSignatureAPI.md#ValidatePdfSignatureEndpointApiV1ProcessingValidatePdfSignaturePost) | **Post** /api/v1/processing/validate-pdf-signature | Validate electronic signatures of a PDF



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
	resp, r, err := apiClient.ElectronicSignatureAPI.GenerateTestCertificateApiV1ProcessingGenerateTestCertificatePost(context.Background()).GenerateCertificateRequest(generateCertificateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElectronicSignatureAPI.GenerateTestCertificateApiV1ProcessingGenerateTestCertificatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateTestCertificateApiV1ProcessingGenerateTestCertificatePost`: GenerateCertificateResponse
	fmt.Fprintf(os.Stdout, "Response from `ElectronicSignatureAPI.GenerateTestCertificateApiV1ProcessingGenerateTestCertificatePost`: %v\n", resp)
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
	resp, r, err := apiClient.ElectronicSignatureAPI.SignPdfApiV1ProcessingSignPdfPost(context.Background()).PdfFile(pdfFile).Reason(reason).Location(location).Contact(contact).FieldName(fieldName).UsePadesLt(usePadesLt).UseTimestamp(useTimestamp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElectronicSignatureAPI.SignPdfApiV1ProcessingSignPdfPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignPdfApiV1ProcessingSignPdfPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ElectronicSignatureAPI.SignPdfApiV1ProcessingSignPdfPost`: %v\n", resp)
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
	resp, r, err := apiClient.ElectronicSignatureAPI.SignPdfAsyncApiV1ProcessingSignPdfAsyncPost(context.Background()).PdfFile(pdfFile).CallbackUrl(callbackUrl).WebhookMode(webhookMode).Reason(reason).Location(location).Contact(contact).FieldName(fieldName).UsePadesLt(usePadesLt).UseTimestamp(useTimestamp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElectronicSignatureAPI.SignPdfAsyncApiV1ProcessingSignPdfAsyncPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignPdfAsyncApiV1ProcessingSignPdfAsyncPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ElectronicSignatureAPI.SignPdfAsyncApiV1ProcessingSignPdfAsyncPost`: %v\n", resp)
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
	resp, r, err := apiClient.ElectronicSignatureAPI.ValidatePdfSignatureEndpointApiV1ProcessingValidatePdfSignaturePost(context.Background()).PdfFile(pdfFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElectronicSignatureAPI.ValidatePdfSignatureEndpointApiV1ProcessingValidatePdfSignaturePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidatePdfSignatureEndpointApiV1ProcessingValidatePdfSignaturePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ElectronicSignatureAPI.ValidatePdfSignatureEndpointApiV1ProcessingValidatePdfSignaturePost`: %v\n", resp)
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


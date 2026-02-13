# \ParseAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ParseFacturxAsyncApiV1ProcessingParseFacturxAsyncPost**](ParseAPI.md#ParseFacturxAsyncApiV1ProcessingParseFacturxAsyncPost) | **Post** /api/v1/processing/parse-facturx/async | Parse Factur-X XML or PDF (async)
[**ParseFacturxSyncApiV1ProcessingParseFacturxPost**](ParseAPI.md#ParseFacturxSyncApiV1ProcessingParseFacturxPost) | **Post** /api/v1/processing/parse-facturx | Parse CII, UBL or Factur-X PDF (sync)



## ParseFacturxAsyncApiV1ProcessingParseFacturxAsyncPost

> TaskResponse ParseFacturxAsyncApiV1ProcessingParseFacturxAsyncPost(ctx).File(file).CallbackUrl(callbackUrl).Execute()

Parse Factur-X XML or PDF (async)



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
	file := os.NewFile(1234, "some_file") // *os.File | Factur-X PDF or XML file to parse
	callbackUrl := "callbackUrl_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParseAPI.ParseFacturxAsyncApiV1ProcessingParseFacturxAsyncPost(context.Background()).File(file).CallbackUrl(callbackUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParseAPI.ParseFacturxAsyncApiV1ProcessingParseFacturxAsyncPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParseFacturxAsyncApiV1ProcessingParseFacturxAsyncPost`: TaskResponse
	fmt.Fprintf(os.Stdout, "Response from `ParseAPI.ParseFacturxAsyncApiV1ProcessingParseFacturxAsyncPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiParseFacturxAsyncApiV1ProcessingParseFacturxAsyncPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | Factur-X PDF or XML file to parse | 
 **callbackUrl** | **string** |  | 

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


## ParseFacturxSyncApiV1ProcessingParseFacturxPost

> ParseFacturXResponse ParseFacturxSyncApiV1ProcessingParseFacturxPost(ctx).File(file).Execute()

Parse CII, UBL or Factur-X PDF (sync)



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
	file := os.NewFile(1234, "some_file") // *os.File | Factur-X PDF or XML file to parse

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParseAPI.ParseFacturxSyncApiV1ProcessingParseFacturxPost(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParseAPI.ParseFacturxSyncApiV1ProcessingParseFacturxPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParseFacturxSyncApiV1ProcessingParseFacturxPost`: ParseFacturXResponse
	fmt.Fprintf(os.Stdout, "Response from `ParseAPI.ParseFacturxSyncApiV1ProcessingParseFacturxPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiParseFacturxSyncApiV1ProcessingParseFacturxPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | Factur-X PDF or XML file to parse | 

### Return type

[**ParseFacturXResponse**](ParseFacturXResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


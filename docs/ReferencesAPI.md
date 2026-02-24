# \ReferencesAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVatexCodesApiV1ReferencesVatexCodesGet**](ReferencesAPI.md#GetVatexCodesApiV1ReferencesVatexCodesGet) | **Get** /api/v1/references/vatex-codes | VATEX exemption reason codes



## GetVatexCodesApiV1ReferencesVatexCodesGet

> VATEXCodesResponse GetVatexCodesApiV1ReferencesVatexCodesGet(ctx).Category(category).Execute()

VATEX exemption reason codes



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
	category := "category_example" // string | Filter by VAT category code (E, AE, K, G, O). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferencesAPI.GetVatexCodesApiV1ReferencesVatexCodesGet(context.Background()).Category(category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferencesAPI.GetVatexCodesApiV1ReferencesVatexCodesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVatexCodesApiV1ReferencesVatexCodesGet`: VATEXCodesResponse
	fmt.Fprintf(os.Stdout, "Response from `ReferencesAPI.GetVatexCodesApiV1ReferencesVatexCodesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVatexCodesApiV1ReferencesVatexCodesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** | Filter by VAT category code (E, AE, K, G, O). | 

### Return type

[**VATEXCodesResponse**](VATEXCodesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


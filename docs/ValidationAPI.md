# \ValidationAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidateXmlApiV1ProcessingValidateXmlPost**](ValidationAPI.md#ValidateXmlApiV1ProcessingValidateXmlPost) | **Post** /api/v1/processing/validate-xml | Validate a CII or UBL XML against Schematron rules



## ValidateXmlApiV1ProcessingValidateXmlPost

> ValidationSuccessResponse ValidateXmlApiV1ProcessingValidateXmlPost(ctx).XmlFile(xmlFile).Profile(profile).SkipBrFr(skipBrFr).Execute()

Validate a CII or UBL XML against Schematron rules



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
	xmlFile := os.NewFile(1234, "some_file") // *os.File | CII (Factur-X) or UBL XML file to validate (.xml format). Format is auto-detected.
	profile := openapiclient.APIProfile("MINIMUM") // APIProfile | Validation profile (MINIMUM, BASIC, EN16931, EXTENDED). Used for CII only; ignored for UBL (always EN16931). (optional) (default to "EXTENDED")
	skipBrFr := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidationAPI.ValidateXmlApiV1ProcessingValidateXmlPost(context.Background()).XmlFile(xmlFile).Profile(profile).SkipBrFr(skipBrFr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidationAPI.ValidateXmlApiV1ProcessingValidateXmlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateXmlApiV1ProcessingValidateXmlPost`: ValidationSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `ValidationAPI.ValidateXmlApiV1ProcessingValidateXmlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateXmlApiV1ProcessingValidateXmlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xmlFile** | ***os.File** | CII (Factur-X) or UBL XML file to validate (.xml format). Format is auto-detected. | 
 **profile** | [**APIProfile**](APIProfile.md) | Validation profile (MINIMUM, BASIC, EN16931, EXTENDED). Used for CII only; ignored for UBL (always EN16931). | [default to &quot;EXTENDED&quot;]
 **skipBrFr** | **bool** |  | 

### Return type

[**ValidationSuccessResponse**](ValidationSuccessResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


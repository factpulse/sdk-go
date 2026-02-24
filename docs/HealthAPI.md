# \HealthAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserInfoApiV1MeGet**](HealthAPI.md#GetUserInfoApiV1MeGet) | **Get** /api/v1/me | Get current user information



## GetUserInfoApiV1MeGet

> interface{} GetUserInfoApiV1MeGet(ctx).Execute()

Get current user information



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
	resp, r, err := apiClient.HealthAPI.GetUserInfoApiV1MeGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GetUserInfoApiV1MeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserInfoApiV1MeGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GetUserInfoApiV1MeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserInfoApiV1MeGetRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


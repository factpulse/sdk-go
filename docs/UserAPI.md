# \UserAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserInfoApiV1MeGet**](UserAPI.md#GetUserInfoApiV1MeGet) | **Get** /api/v1/me | Get current user information



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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUserInfoApiV1MeGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserInfoApiV1MeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserInfoApiV1MeGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserInfoApiV1MeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserInfoApiV1MeGetRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


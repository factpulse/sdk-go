# \AFNORPDPPAAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OauthTokenProxyApiV1AfnorOauthTokenPost**](AFNORPDPPAAPI.md#OauthTokenProxyApiV1AfnorOauthTokenPost) | **Post** /api/v1/afnor/oauth/token | Endpoint OAuth2 pour authentification AFNOR



## OauthTokenProxyApiV1AfnorOauthTokenPost

> interface{} OauthTokenProxyApiV1AfnorOauthTokenPost(ctx).Execute()

Endpoint OAuth2 pour authentification AFNOR



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
	resp, r, err := apiClient.AFNORPDPPAAPI.OauthTokenProxyApiV1AfnorOauthTokenPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AFNORPDPPAAPI.OauthTokenProxyApiV1AfnorOauthTokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OauthTokenProxyApiV1AfnorOauthTokenPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AFNORPDPPAAPI.OauthTokenProxyApiV1AfnorOauthTokenPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOauthTokenProxyApiV1AfnorOauthTokenPostRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


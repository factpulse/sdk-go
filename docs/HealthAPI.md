# \HealthAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HealthcheckHealthcheckGet**](HealthAPI.md#HealthcheckHealthcheckGet) | **Get** /healthcheck | Docker healthcheck endpoint
[**RootGet**](HealthAPI.md#RootGet) | **Get** / | Check API status



## HealthcheckHealthcheckGet

> interface{} HealthcheckHealthcheckGet(ctx).Execute()

Docker healthcheck endpoint



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
	resp, r, err := apiClient.HealthAPI.HealthcheckHealthcheckGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.HealthcheckHealthcheckGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HealthcheckHealthcheckGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.HealthcheckHealthcheckGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthcheckHealthcheckGetRequest struct via the builder pattern


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


## RootGet

> interface{} RootGet(ctx).Execute()

Check API status



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
	resp, r, err := apiClient.HealthAPI.RootGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.RootGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RootGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.RootGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRootGetRequest struct via the builder pattern


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


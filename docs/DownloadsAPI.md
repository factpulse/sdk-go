# \DownloadsAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckFileApiV1DownloadDownloadIdHead**](DownloadsAPI.md#CheckFileApiV1DownloadDownloadIdHead) | **Head** /api/v1/download/{download_id} | Check if a file exists
[**DownloadFileApiV1DownloadDownloadIdGet**](DownloadsAPI.md#DownloadFileApiV1DownloadDownloadIdGet) | **Get** /api/v1/download/{download_id} | Download a temporary file



## CheckFileApiV1DownloadDownloadIdHead

> interface{} CheckFileApiV1DownloadDownloadIdHead(ctx, downloadId).Execute()

Check if a file exists



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
	downloadId := "downloadId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DownloadsAPI.CheckFileApiV1DownloadDownloadIdHead(context.Background(), downloadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadsAPI.CheckFileApiV1DownloadDownloadIdHead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckFileApiV1DownloadDownloadIdHead`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DownloadsAPI.CheckFileApiV1DownloadDownloadIdHead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**downloadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckFileApiV1DownloadDownloadIdHeadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DownloadFileApiV1DownloadDownloadIdGet

> interface{} DownloadFileApiV1DownloadDownloadIdGet(ctx, downloadId).DeleteAfter(deleteAfter).Execute()

Download a temporary file



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
	downloadId := "downloadId_example" // string | 
	deleteAfter := true // bool | If true, delete the file after download (one-time download) (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DownloadsAPI.DownloadFileApiV1DownloadDownloadIdGet(context.Background(), downloadId).DeleteAfter(deleteAfter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadsAPI.DownloadFileApiV1DownloadDownloadIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadFileApiV1DownloadDownloadIdGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DownloadsAPI.DownloadFileApiV1DownloadDownloadIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**downloadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadFileApiV1DownloadDownloadIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteAfter** | **bool** | If true, delete the file after download (one-time download) | [default to false]

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/pdf, application/xml, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


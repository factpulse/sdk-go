# \AsyncTasksAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTaskStatusApiV1ProcessingTasksTaskIdStatusGet**](AsyncTasksAPI.md#GetTaskStatusApiV1ProcessingTasksTaskIdStatusGet) | **Get** /api/v1/processing/tasks/{task_id}/status | Get task generation status



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
	resp, r, err := apiClient.AsyncTasksAPI.GetTaskStatusApiV1ProcessingTasksTaskIdStatusGet(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AsyncTasksAPI.GetTaskStatusApiV1ProcessingTasksTaskIdStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskStatusApiV1ProcessingTasksTaskIdStatusGet`: AsyncTaskStatus
	fmt.Fprintf(os.Stdout, "Response from `AsyncTasksAPI.GetTaskStatusApiV1ProcessingTasksTaskIdStatusGet`: %v\n", resp)
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


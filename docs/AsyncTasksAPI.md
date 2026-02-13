# \AsyncTasksAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGenerationTaskStatusApiV1ProcessingTasksTaskIdGenerationStatusGet**](AsyncTasksAPI.md#GetGenerationTaskStatusApiV1ProcessingTasksTaskIdGenerationStatusGet) | **Get** /api/v1/processing/tasks/{task_id}/generation-status | Get typed generation task status
[**GetSignatureTaskStatusApiV1ProcessingTasksTaskIdSignatureStatusGet**](AsyncTasksAPI.md#GetSignatureTaskStatusApiV1ProcessingTasksTaskIdSignatureStatusGet) | **Get** /api/v1/processing/tasks/{task_id}/signature-status | Get typed signature task status
[**GetTaskStatusApiV1ProcessingTasksTaskIdStatusGet**](AsyncTasksAPI.md#GetTaskStatusApiV1ProcessingTasksTaskIdStatusGet) | **Get** /api/v1/processing/tasks/{task_id}/status | Get task generation status
[**GetValidationTaskStatusApiV1ProcessingTasksTaskIdValidationStatusGet**](AsyncTasksAPI.md#GetValidationTaskStatusApiV1ProcessingTasksTaskIdValidationStatusGet) | **Get** /api/v1/processing/tasks/{task_id}/validation-status | Get typed validation task status



## GetGenerationTaskStatusApiV1ProcessingTasksTaskIdGenerationStatusGet

> GenerationTaskStatus GetGenerationTaskStatusApiV1ProcessingTasksTaskIdGenerationStatusGet(ctx, taskId).Execute()

Get typed generation task status



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
	taskId := "taskId_example" // string | Celery task ID returned by /generate-invoice

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AsyncTasksAPI.GetGenerationTaskStatusApiV1ProcessingTasksTaskIdGenerationStatusGet(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AsyncTasksAPI.GetGenerationTaskStatusApiV1ProcessingTasksTaskIdGenerationStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGenerationTaskStatusApiV1ProcessingTasksTaskIdGenerationStatusGet`: GenerationTaskStatus
	fmt.Fprintf(os.Stdout, "Response from `AsyncTasksAPI.GetGenerationTaskStatusApiV1ProcessingTasksTaskIdGenerationStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | Celery task ID returned by /generate-invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGenerationTaskStatusApiV1ProcessingTasksTaskIdGenerationStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GenerationTaskStatus**](GenerationTaskStatus.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignatureTaskStatusApiV1ProcessingTasksTaskIdSignatureStatusGet

> SignatureTaskStatus GetSignatureTaskStatusApiV1ProcessingTasksTaskIdSignatureStatusGet(ctx, taskId).Execute()

Get typed signature task status



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
	taskId := "taskId_example" // string | Celery task ID returned by /sign-pdf-async

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AsyncTasksAPI.GetSignatureTaskStatusApiV1ProcessingTasksTaskIdSignatureStatusGet(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AsyncTasksAPI.GetSignatureTaskStatusApiV1ProcessingTasksTaskIdSignatureStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSignatureTaskStatusApiV1ProcessingTasksTaskIdSignatureStatusGet`: SignatureTaskStatus
	fmt.Fprintf(os.Stdout, "Response from `AsyncTasksAPI.GetSignatureTaskStatusApiV1ProcessingTasksTaskIdSignatureStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | Celery task ID returned by /sign-pdf-async | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSignatureTaskStatusApiV1ProcessingTasksTaskIdSignatureStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SignatureTaskStatus**](SignatureTaskStatus.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetValidationTaskStatusApiV1ProcessingTasksTaskIdValidationStatusGet

> ValidationTaskStatus GetValidationTaskStatusApiV1ProcessingTasksTaskIdValidationStatusGet(ctx, taskId).Execute()

Get typed validation task status



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
	taskId := "taskId_example" // string | Celery task ID returned by /validate-pdf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AsyncTasksAPI.GetValidationTaskStatusApiV1ProcessingTasksTaskIdValidationStatusGet(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AsyncTasksAPI.GetValidationTaskStatusApiV1ProcessingTasksTaskIdValidationStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetValidationTaskStatusApiV1ProcessingTasksTaskIdValidationStatusGet`: ValidationTaskStatus
	fmt.Fprintf(os.Stdout, "Response from `AsyncTasksAPI.GetValidationTaskStatusApiV1ProcessingTasksTaskIdValidationStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | Celery task ID returned by /validate-pdf | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetValidationTaskStatusApiV1ProcessingTasksTaskIdValidationStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ValidationTaskStatus**](ValidationTaskStatus.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


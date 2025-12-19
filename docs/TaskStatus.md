# TaskStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** | Unique task identifier | 
**Status** | [**CeleryStatus**](CeleryStatus.md) | Celery task status (PENDING, STARTED, SUCCESS, FAILURE, RETRY) | 
**Result** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTaskStatus

`func NewTaskStatus(taskId string, status CeleryStatus, ) *TaskStatus`

NewTaskStatus instantiates a new TaskStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskStatusWithDefaults

`func NewTaskStatusWithDefaults() *TaskStatus`

NewTaskStatusWithDefaults instantiates a new TaskStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *TaskStatus) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TaskStatus) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TaskStatus) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetStatus

`func (o *TaskStatus) GetStatus() CeleryStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskStatus) GetStatusOk() (*CeleryStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskStatus) SetStatus(v CeleryStatus)`

SetStatus sets Status field to given value.


### GetResult

`func (o *TaskStatus) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TaskStatus) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TaskStatus) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *TaskStatus) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *TaskStatus) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *TaskStatus) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GenerationTaskStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** |  | 
**Status** | [**CeleryStatus**](CeleryStatus.md) |  | 
**Result** | Pointer to [**NullableGenerationTaskStatusResult**](GenerationTaskStatusResult.md) |  | [optional] 

## Methods

### NewGenerationTaskStatus

`func NewGenerationTaskStatus(taskId string, status CeleryStatus, ) *GenerationTaskStatus`

NewGenerationTaskStatus instantiates a new GenerationTaskStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerationTaskStatusWithDefaults

`func NewGenerationTaskStatusWithDefaults() *GenerationTaskStatus`

NewGenerationTaskStatusWithDefaults instantiates a new GenerationTaskStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *GenerationTaskStatus) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *GenerationTaskStatus) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *GenerationTaskStatus) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetStatus

`func (o *GenerationTaskStatus) GetStatus() CeleryStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GenerationTaskStatus) GetStatusOk() (*CeleryStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GenerationTaskStatus) SetStatus(v CeleryStatus)`

SetStatus sets Status field to given value.


### GetResult

`func (o *GenerationTaskStatus) GetResult() GenerationTaskStatusResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GenerationTaskStatus) GetResultOk() (*GenerationTaskStatusResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GenerationTaskStatus) SetResult(v GenerationTaskStatusResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *GenerationTaskStatus) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *GenerationTaskStatus) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *GenerationTaskStatus) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



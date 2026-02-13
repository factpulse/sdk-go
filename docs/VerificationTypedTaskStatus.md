# VerificationTypedTaskStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** |  | 
**Status** | [**CeleryStatus**](CeleryStatus.md) |  | 
**Result** | Pointer to [**NullableVerificationTypedTaskStatusResult**](VerificationTypedTaskStatusResult.md) |  | [optional] 

## Methods

### NewVerificationTypedTaskStatus

`func NewVerificationTypedTaskStatus(taskId string, status CeleryStatus, ) *VerificationTypedTaskStatus`

NewVerificationTypedTaskStatus instantiates a new VerificationTypedTaskStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationTypedTaskStatusWithDefaults

`func NewVerificationTypedTaskStatusWithDefaults() *VerificationTypedTaskStatus`

NewVerificationTypedTaskStatusWithDefaults instantiates a new VerificationTypedTaskStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *VerificationTypedTaskStatus) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *VerificationTypedTaskStatus) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *VerificationTypedTaskStatus) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetStatus

`func (o *VerificationTypedTaskStatus) GetStatus() CeleryStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VerificationTypedTaskStatus) GetStatusOk() (*CeleryStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VerificationTypedTaskStatus) SetStatus(v CeleryStatus)`

SetStatus sets Status field to given value.


### GetResult

`func (o *VerificationTypedTaskStatus) GetResult() VerificationTypedTaskStatusResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *VerificationTypedTaskStatus) GetResultOk() (*VerificationTypedTaskStatusResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *VerificationTypedTaskStatus) SetResult(v VerificationTypedTaskStatusResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *VerificationTypedTaskStatus) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *VerificationTypedTaskStatus) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *VerificationTypedTaskStatus) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



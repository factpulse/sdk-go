# ConversionTaskStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** |  | 
**Status** | [**CeleryStatus**](CeleryStatus.md) |  | 
**Result** | Pointer to [**NullableConversionTaskStatusResult**](ConversionTaskStatusResult.md) |  | [optional] 

## Methods

### NewConversionTaskStatus

`func NewConversionTaskStatus(taskId string, status CeleryStatus, ) *ConversionTaskStatus`

NewConversionTaskStatus instantiates a new ConversionTaskStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversionTaskStatusWithDefaults

`func NewConversionTaskStatusWithDefaults() *ConversionTaskStatus`

NewConversionTaskStatusWithDefaults instantiates a new ConversionTaskStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *ConversionTaskStatus) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ConversionTaskStatus) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *ConversionTaskStatus) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetStatus

`func (o *ConversionTaskStatus) GetStatus() CeleryStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConversionTaskStatus) GetStatusOk() (*CeleryStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConversionTaskStatus) SetStatus(v CeleryStatus)`

SetStatus sets Status field to given value.


### GetResult

`func (o *ConversionTaskStatus) GetResult() ConversionTaskStatusResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ConversionTaskStatus) GetResultOk() (*ConversionTaskStatusResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ConversionTaskStatus) SetResult(v ConversionTaskStatusResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ConversionTaskStatus) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *ConversionTaskStatus) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *ConversionTaskStatus) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



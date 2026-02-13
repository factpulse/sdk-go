# TaskErrorResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "ERROR"]
**ErrorCode** | **string** |  | 
**ErrorMessage** | **string** |  | 
**Details** | Pointer to [**[]AFNORErrorDetail**](AFNORErrorDetail.md) |  | [optional] 
**Traceback** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTaskErrorResult

`func NewTaskErrorResult(errorCode string, errorMessage string, ) *TaskErrorResult`

NewTaskErrorResult instantiates a new TaskErrorResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskErrorResultWithDefaults

`func NewTaskErrorResultWithDefaults() *TaskErrorResult`

NewTaskErrorResultWithDefaults instantiates a new TaskErrorResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TaskErrorResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskErrorResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskErrorResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskErrorResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorCode

`func (o *TaskErrorResult) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *TaskErrorResult) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *TaskErrorResult) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorMessage

`func (o *TaskErrorResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TaskErrorResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TaskErrorResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### GetDetails

`func (o *TaskErrorResult) GetDetails() []AFNORErrorDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *TaskErrorResult) GetDetailsOk() (*[]AFNORErrorDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *TaskErrorResult) SetDetails(v []AFNORErrorDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *TaskErrorResult) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTraceback

`func (o *TaskErrorResult) GetTraceback() string`

GetTraceback returns the Traceback field if non-nil, zero value otherwise.

### GetTracebackOk

`func (o *TaskErrorResult) GetTracebackOk() (*string, bool)`

GetTracebackOk returns a tuple with the Traceback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceback

`func (o *TaskErrorResult) SetTraceback(v string)`

SetTraceback sets Traceback field to given value.

### HasTraceback

`func (o *TaskErrorResult) HasTraceback() bool`

HasTraceback returns a boolean if a field has been set.

### SetTracebackNil

`func (o *TaskErrorResult) SetTracebackNil(b bool)`

 SetTracebackNil sets the value for Traceback to be an explicit nil

### UnsetTraceback
`func (o *TaskErrorResult) UnsetTraceback()`

UnsetTraceback ensures that no value is present for Traceback, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



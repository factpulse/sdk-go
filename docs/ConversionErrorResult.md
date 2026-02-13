# ConversionErrorResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "ERROR"]
**ConversionId** | Pointer to **NullableString** |  | [optional] 
**ErrorCode** | **string** |  | 
**ErrorMessage** | **string** |  | 
**Details** | Pointer to [**[]AFNORErrorDetail**](AFNORErrorDetail.md) |  | [optional] 
**Traceback** | Pointer to **NullableString** |  | [optional] 
**ProcessingTimeMs** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewConversionErrorResult

`func NewConversionErrorResult(errorCode string, errorMessage string, ) *ConversionErrorResult`

NewConversionErrorResult instantiates a new ConversionErrorResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversionErrorResultWithDefaults

`func NewConversionErrorResultWithDefaults() *ConversionErrorResult`

NewConversionErrorResultWithDefaults instantiates a new ConversionErrorResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ConversionErrorResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConversionErrorResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConversionErrorResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConversionErrorResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConversionId

`func (o *ConversionErrorResult) GetConversionId() string`

GetConversionId returns the ConversionId field if non-nil, zero value otherwise.

### GetConversionIdOk

`func (o *ConversionErrorResult) GetConversionIdOk() (*string, bool)`

GetConversionIdOk returns a tuple with the ConversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionId

`func (o *ConversionErrorResult) SetConversionId(v string)`

SetConversionId sets ConversionId field to given value.

### HasConversionId

`func (o *ConversionErrorResult) HasConversionId() bool`

HasConversionId returns a boolean if a field has been set.

### SetConversionIdNil

`func (o *ConversionErrorResult) SetConversionIdNil(b bool)`

 SetConversionIdNil sets the value for ConversionId to be an explicit nil

### UnsetConversionId
`func (o *ConversionErrorResult) UnsetConversionId()`

UnsetConversionId ensures that no value is present for ConversionId, not even an explicit nil
### GetErrorCode

`func (o *ConversionErrorResult) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ConversionErrorResult) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ConversionErrorResult) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorMessage

`func (o *ConversionErrorResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ConversionErrorResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ConversionErrorResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### GetDetails

`func (o *ConversionErrorResult) GetDetails() []AFNORErrorDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ConversionErrorResult) GetDetailsOk() (*[]AFNORErrorDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ConversionErrorResult) SetDetails(v []AFNORErrorDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ConversionErrorResult) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTraceback

`func (o *ConversionErrorResult) GetTraceback() string`

GetTraceback returns the Traceback field if non-nil, zero value otherwise.

### GetTracebackOk

`func (o *ConversionErrorResult) GetTracebackOk() (*string, bool)`

GetTracebackOk returns a tuple with the Traceback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceback

`func (o *ConversionErrorResult) SetTraceback(v string)`

SetTraceback sets Traceback field to given value.

### HasTraceback

`func (o *ConversionErrorResult) HasTraceback() bool`

HasTraceback returns a boolean if a field has been set.

### SetTracebackNil

`func (o *ConversionErrorResult) SetTracebackNil(b bool)`

 SetTracebackNil sets the value for Traceback to be an explicit nil

### UnsetTraceback
`func (o *ConversionErrorResult) UnsetTraceback()`

UnsetTraceback ensures that no value is present for Traceback, not even an explicit nil
### GetProcessingTimeMs

`func (o *ConversionErrorResult) GetProcessingTimeMs() int32`

GetProcessingTimeMs returns the ProcessingTimeMs field if non-nil, zero value otherwise.

### GetProcessingTimeMsOk

`func (o *ConversionErrorResult) GetProcessingTimeMsOk() (*int32, bool)`

GetProcessingTimeMsOk returns a tuple with the ProcessingTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeMs

`func (o *ConversionErrorResult) SetProcessingTimeMs(v int32)`

SetProcessingTimeMs sets ProcessingTimeMs field to given value.

### HasProcessingTimeMs

`func (o *ConversionErrorResult) HasProcessingTimeMs() bool`

HasProcessingTimeMs returns a boolean if a field has been set.

### SetProcessingTimeMsNil

`func (o *ConversionErrorResult) SetProcessingTimeMsNil(b bool)`

 SetProcessingTimeMsNil sets the value for ProcessingTimeMs to be an explicit nil

### UnsetProcessingTimeMs
`func (o *ConversionErrorResult) UnsetProcessingTimeMs()`

UnsetProcessingTimeMs ensures that no value is present for ProcessingTimeMs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ValidationTaskStatusResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "ERROR"]
**ValidationResult** | [**PDFValidationResultAPI**](PDFValidationResultAPI.md) |  | 
**ErrorCode** | **string** |  | 
**ErrorMessage** | **string** |  | 
**Details** | Pointer to [**[]AFNORErrorDetail**](AFNORErrorDetail.md) |  | [optional] 
**Traceback** | Pointer to **string** |  | [optional] 

## Methods

### NewValidationTaskStatusResult

`func NewValidationTaskStatusResult(validationResult PDFValidationResultAPI, errorCode string, errorMessage string, ) *ValidationTaskStatusResult`

NewValidationTaskStatusResult instantiates a new ValidationTaskStatusResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationTaskStatusResultWithDefaults

`func NewValidationTaskStatusResultWithDefaults() *ValidationTaskStatusResult`

NewValidationTaskStatusResultWithDefaults instantiates a new ValidationTaskStatusResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ValidationTaskStatusResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ValidationTaskStatusResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ValidationTaskStatusResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ValidationTaskStatusResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetValidationResult

`func (o *ValidationTaskStatusResult) GetValidationResult() PDFValidationResultAPI`

GetValidationResult returns the ValidationResult field if non-nil, zero value otherwise.

### GetValidationResultOk

`func (o *ValidationTaskStatusResult) GetValidationResultOk() (*PDFValidationResultAPI, bool)`

GetValidationResultOk returns a tuple with the ValidationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationResult

`func (o *ValidationTaskStatusResult) SetValidationResult(v PDFValidationResultAPI)`

SetValidationResult sets ValidationResult field to given value.


### GetErrorCode

`func (o *ValidationTaskStatusResult) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ValidationTaskStatusResult) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ValidationTaskStatusResult) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorMessage

`func (o *ValidationTaskStatusResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ValidationTaskStatusResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ValidationTaskStatusResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### GetDetails

`func (o *ValidationTaskStatusResult) GetDetails() []AFNORErrorDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ValidationTaskStatusResult) GetDetailsOk() (*[]AFNORErrorDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ValidationTaskStatusResult) SetDetails(v []AFNORErrorDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ValidationTaskStatusResult) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTraceback

`func (o *ValidationTaskStatusResult) GetTraceback() string`

GetTraceback returns the Traceback field if non-nil, zero value otherwise.

### GetTracebackOk

`func (o *ValidationTaskStatusResult) GetTracebackOk() (*string, bool)`

GetTracebackOk returns a tuple with the Traceback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceback

`func (o *ValidationTaskStatusResult) SetTraceback(v string)`

SetTraceback sets Traceback field to given value.

### HasTraceback

`func (o *ValidationTaskStatusResult) HasTraceback() bool`

HasTraceback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



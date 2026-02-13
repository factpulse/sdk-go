# ValidationSuccessTaskResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "SUCCESS"]
**ValidationResult** | [**PDFValidationResultAPI**](PDFValidationResultAPI.md) |  | 

## Methods

### NewValidationSuccessTaskResult

`func NewValidationSuccessTaskResult(validationResult PDFValidationResultAPI, ) *ValidationSuccessTaskResult`

NewValidationSuccessTaskResult instantiates a new ValidationSuccessTaskResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationSuccessTaskResultWithDefaults

`func NewValidationSuccessTaskResultWithDefaults() *ValidationSuccessTaskResult`

NewValidationSuccessTaskResultWithDefaults instantiates a new ValidationSuccessTaskResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ValidationSuccessTaskResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ValidationSuccessTaskResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ValidationSuccessTaskResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ValidationSuccessTaskResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetValidationResult

`func (o *ValidationSuccessTaskResult) GetValidationResult() PDFValidationResultAPI`

GetValidationResult returns the ValidationResult field if non-nil, zero value otherwise.

### GetValidationResultOk

`func (o *ValidationSuccessTaskResult) GetValidationResultOk() (*PDFValidationResultAPI, bool)`

GetValidationResultOk returns a tuple with the ValidationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationResult

`func (o *ValidationSuccessTaskResult) SetValidationResult(v PDFValidationResultAPI)`

SetValidationResult sets ValidationResult field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



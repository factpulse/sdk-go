# EReportingValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Field path with error | 
**Message** | **string** | Error message | 
**Code** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEReportingValidationError

`func NewEReportingValidationError(field string, message string, ) *EReportingValidationError`

NewEReportingValidationError instantiates a new EReportingValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEReportingValidationErrorWithDefaults

`func NewEReportingValidationErrorWithDefaults() *EReportingValidationError`

NewEReportingValidationErrorWithDefaults instantiates a new EReportingValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *EReportingValidationError) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *EReportingValidationError) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *EReportingValidationError) SetField(v string)`

SetField sets Field field to given value.


### GetMessage

`func (o *EReportingValidationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EReportingValidationError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EReportingValidationError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCode

`func (o *EReportingValidationError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EReportingValidationError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EReportingValidationError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *EReportingValidationError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *EReportingValidationError) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *EReportingValidationError) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



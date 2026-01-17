# ValidationErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Champ concerné | 
**Message** | **string** | Message d&#39;erreur | 
**Rule** | Pointer to **NullableString** |  | [optional] 
**Severity** | Pointer to **string** | Sévérité (error/warning) | [optional] [default to "error"]

## Methods

### NewValidationErrorResponse

`func NewValidationErrorResponse(field string, message string, ) *ValidationErrorResponse`

NewValidationErrorResponse instantiates a new ValidationErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorResponseWithDefaults

`func NewValidationErrorResponseWithDefaults() *ValidationErrorResponse`

NewValidationErrorResponseWithDefaults instantiates a new ValidationErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *ValidationErrorResponse) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ValidationErrorResponse) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ValidationErrorResponse) SetField(v string)`

SetField sets Field field to given value.


### GetMessage

`func (o *ValidationErrorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidationErrorResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidationErrorResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRule

`func (o *ValidationErrorResponse) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *ValidationErrorResponse) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *ValidationErrorResponse) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *ValidationErrorResponse) HasRule() bool`

HasRule returns a boolean if a field has been set.

### SetRuleNil

`func (o *ValidationErrorResponse) SetRuleNil(b bool)`

 SetRuleNil sets the value for Rule to be an explicit nil

### UnsetRule
`func (o *ValidationErrorResponse) UnsetRule()`

UnsetRule ensures that no value is present for Rule, not even an explicit nil
### GetSeverity

`func (o *ValidationErrorResponse) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ValidationErrorResponse) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ValidationErrorResponse) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ValidationErrorResponse) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



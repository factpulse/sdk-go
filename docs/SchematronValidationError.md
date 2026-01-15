# SchematronValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | **string** | Code de la regle (BR-XX, BR-FR-XX) | 
**BtCode** | Pointer to **NullableString** |  | [optional] 
**Severity** | **string** | Gravite: error, warning | 
**Message** | **string** | Message d&#39;erreur | 
**SuggestedValue** | Pointer to **NullableString** |  | [optional] 
**SuggestedField** | Pointer to **NullableString** |  | [optional] 
**Explanation** | Pointer to **NullableString** |  | [optional] 
**Confidence** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewSchematronValidationError

`func NewSchematronValidationError(rule string, severity string, message string, ) *SchematronValidationError`

NewSchematronValidationError instantiates a new SchematronValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchematronValidationErrorWithDefaults

`func NewSchematronValidationErrorWithDefaults() *SchematronValidationError`

NewSchematronValidationErrorWithDefaults instantiates a new SchematronValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRule

`func (o *SchematronValidationError) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *SchematronValidationError) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *SchematronValidationError) SetRule(v string)`

SetRule sets Rule field to given value.


### GetBtCode

`func (o *SchematronValidationError) GetBtCode() string`

GetBtCode returns the BtCode field if non-nil, zero value otherwise.

### GetBtCodeOk

`func (o *SchematronValidationError) GetBtCodeOk() (*string, bool)`

GetBtCodeOk returns a tuple with the BtCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtCode

`func (o *SchematronValidationError) SetBtCode(v string)`

SetBtCode sets BtCode field to given value.

### HasBtCode

`func (o *SchematronValidationError) HasBtCode() bool`

HasBtCode returns a boolean if a field has been set.

### SetBtCodeNil

`func (o *SchematronValidationError) SetBtCodeNil(b bool)`

 SetBtCodeNil sets the value for BtCode to be an explicit nil

### UnsetBtCode
`func (o *SchematronValidationError) UnsetBtCode()`

UnsetBtCode ensures that no value is present for BtCode, not even an explicit nil
### GetSeverity

`func (o *SchematronValidationError) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *SchematronValidationError) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *SchematronValidationError) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetMessage

`func (o *SchematronValidationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SchematronValidationError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SchematronValidationError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSuggestedValue

`func (o *SchematronValidationError) GetSuggestedValue() string`

GetSuggestedValue returns the SuggestedValue field if non-nil, zero value otherwise.

### GetSuggestedValueOk

`func (o *SchematronValidationError) GetSuggestedValueOk() (*string, bool)`

GetSuggestedValueOk returns a tuple with the SuggestedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedValue

`func (o *SchematronValidationError) SetSuggestedValue(v string)`

SetSuggestedValue sets SuggestedValue field to given value.

### HasSuggestedValue

`func (o *SchematronValidationError) HasSuggestedValue() bool`

HasSuggestedValue returns a boolean if a field has been set.

### SetSuggestedValueNil

`func (o *SchematronValidationError) SetSuggestedValueNil(b bool)`

 SetSuggestedValueNil sets the value for SuggestedValue to be an explicit nil

### UnsetSuggestedValue
`func (o *SchematronValidationError) UnsetSuggestedValue()`

UnsetSuggestedValue ensures that no value is present for SuggestedValue, not even an explicit nil
### GetSuggestedField

`func (o *SchematronValidationError) GetSuggestedField() string`

GetSuggestedField returns the SuggestedField field if non-nil, zero value otherwise.

### GetSuggestedFieldOk

`func (o *SchematronValidationError) GetSuggestedFieldOk() (*string, bool)`

GetSuggestedFieldOk returns a tuple with the SuggestedField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedField

`func (o *SchematronValidationError) SetSuggestedField(v string)`

SetSuggestedField sets SuggestedField field to given value.

### HasSuggestedField

`func (o *SchematronValidationError) HasSuggestedField() bool`

HasSuggestedField returns a boolean if a field has been set.

### SetSuggestedFieldNil

`func (o *SchematronValidationError) SetSuggestedFieldNil(b bool)`

 SetSuggestedFieldNil sets the value for SuggestedField to be an explicit nil

### UnsetSuggestedField
`func (o *SchematronValidationError) UnsetSuggestedField()`

UnsetSuggestedField ensures that no value is present for SuggestedField, not even an explicit nil
### GetExplanation

`func (o *SchematronValidationError) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *SchematronValidationError) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *SchematronValidationError) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.

### HasExplanation

`func (o *SchematronValidationError) HasExplanation() bool`

HasExplanation returns a boolean if a field has been set.

### SetExplanationNil

`func (o *SchematronValidationError) SetExplanationNil(b bool)`

 SetExplanationNil sets the value for Explanation to be an explicit nil

### UnsetExplanation
`func (o *SchematronValidationError) UnsetExplanation()`

UnsetExplanation ensures that no value is present for Explanation, not even an explicit nil
### GetConfidence

`func (o *SchematronValidationError) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *SchematronValidationError) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *SchematronValidationError) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *SchematronValidationError) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidenceNil

`func (o *SchematronValidationError) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *SchematronValidationError) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



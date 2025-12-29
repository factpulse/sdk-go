# MissingField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Nom du champ | 
**BtCode** | **string** | Code Business Term (BT-XX) | 
**Description** | **string** | Description du champ | 
**RequiredFor** | **[]string** | Profils necessitant ce champ | 
**SuggestedValue** | Pointer to **NullableString** |  | [optional] 
**Confidence** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewMissingField

`func NewMissingField(field string, btCode string, description string, requiredFor []string, ) *MissingField`

NewMissingField instantiates a new MissingField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMissingFieldWithDefaults

`func NewMissingFieldWithDefaults() *MissingField`

NewMissingFieldWithDefaults instantiates a new MissingField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *MissingField) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *MissingField) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *MissingField) SetField(v string)`

SetField sets Field field to given value.


### GetBtCode

`func (o *MissingField) GetBtCode() string`

GetBtCode returns the BtCode field if non-nil, zero value otherwise.

### GetBtCodeOk

`func (o *MissingField) GetBtCodeOk() (*string, bool)`

GetBtCodeOk returns a tuple with the BtCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtCode

`func (o *MissingField) SetBtCode(v string)`

SetBtCode sets BtCode field to given value.


### GetDescription

`func (o *MissingField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MissingField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MissingField) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRequiredFor

`func (o *MissingField) GetRequiredFor() []string`

GetRequiredFor returns the RequiredFor field if non-nil, zero value otherwise.

### GetRequiredForOk

`func (o *MissingField) GetRequiredForOk() (*[]string, bool)`

GetRequiredForOk returns a tuple with the RequiredFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFor

`func (o *MissingField) SetRequiredFor(v []string)`

SetRequiredFor sets RequiredFor field to given value.


### GetSuggestedValue

`func (o *MissingField) GetSuggestedValue() string`

GetSuggestedValue returns the SuggestedValue field if non-nil, zero value otherwise.

### GetSuggestedValueOk

`func (o *MissingField) GetSuggestedValueOk() (*string, bool)`

GetSuggestedValueOk returns a tuple with the SuggestedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedValue

`func (o *MissingField) SetSuggestedValue(v string)`

SetSuggestedValue sets SuggestedValue field to given value.

### HasSuggestedValue

`func (o *MissingField) HasSuggestedValue() bool`

HasSuggestedValue returns a boolean if a field has been set.

### SetSuggestedValueNil

`func (o *MissingField) SetSuggestedValueNil(b bool)`

 SetSuggestedValueNil sets the value for SuggestedValue to be an explicit nil

### UnsetSuggestedValue
`func (o *MissingField) UnsetSuggestedValue()`

UnsetSuggestedValue ensures that no value is present for SuggestedValue, not even an explicit nil
### GetConfidence

`func (o *MissingField) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *MissingField) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *MissingField) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *MissingField) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidenceNil

`func (o *MissingField) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *MissingField) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



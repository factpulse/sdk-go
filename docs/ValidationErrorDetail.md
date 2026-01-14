# ValidationErrorDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to [**ErrorLevel**](ErrorLevel.md) | Severity level: &#39;Error&#39; or &#39;Warning&#39; | [optional] [default to ERROR]
**Item** | **string** | Identifier of the concerned element (XPath, field, BR-FR rule, etc.) | 
**Reason** | **string** | Error description | 
**Source** | Pointer to [**NullableErrorSource**](ErrorSource.md) |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewValidationErrorDetail

`func NewValidationErrorDetail(item string, reason string, ) *ValidationErrorDetail`

NewValidationErrorDetail instantiates a new ValidationErrorDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorDetailWithDefaults

`func NewValidationErrorDetailWithDefaults() *ValidationErrorDetail`

NewValidationErrorDetailWithDefaults instantiates a new ValidationErrorDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *ValidationErrorDetail) GetLevel() ErrorLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ValidationErrorDetail) GetLevelOk() (*ErrorLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ValidationErrorDetail) SetLevel(v ErrorLevel)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ValidationErrorDetail) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetItem

`func (o *ValidationErrorDetail) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ValidationErrorDetail) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ValidationErrorDetail) SetItem(v string)`

SetItem sets Item field to given value.


### GetReason

`func (o *ValidationErrorDetail) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ValidationErrorDetail) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ValidationErrorDetail) SetReason(v string)`

SetReason sets Reason field to given value.


### GetSource

`func (o *ValidationErrorDetail) GetSource() ErrorSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ValidationErrorDetail) GetSourceOk() (*ErrorSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ValidationErrorDetail) SetSource(v ErrorSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *ValidationErrorDetail) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *ValidationErrorDetail) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *ValidationErrorDetail) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetCode

`func (o *ValidationErrorDetail) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ValidationErrorDetail) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ValidationErrorDetail) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ValidationErrorDetail) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ValidationErrorDetail) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ValidationErrorDetail) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



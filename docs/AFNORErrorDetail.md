# AFNORErrorDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | **string** |  | 
**Item** | **string** |  | 
**Reason** | **string** |  | 
**Source** | Pointer to **NullableString** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAFNORErrorDetail

`func NewAFNORErrorDetail(level string, item string, reason string, ) *AFNORErrorDetail`

NewAFNORErrorDetail instantiates a new AFNORErrorDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORErrorDetailWithDefaults

`func NewAFNORErrorDetailWithDefaults() *AFNORErrorDetail`

NewAFNORErrorDetailWithDefaults instantiates a new AFNORErrorDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *AFNORErrorDetail) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AFNORErrorDetail) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AFNORErrorDetail) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetItem

`func (o *AFNORErrorDetail) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *AFNORErrorDetail) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *AFNORErrorDetail) SetItem(v string)`

SetItem sets Item field to given value.


### GetReason

`func (o *AFNORErrorDetail) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AFNORErrorDetail) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AFNORErrorDetail) SetReason(v string)`

SetReason sets Reason field to given value.


### GetSource

`func (o *AFNORErrorDetail) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AFNORErrorDetail) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AFNORErrorDetail) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AFNORErrorDetail) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *AFNORErrorDetail) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *AFNORErrorDetail) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetCode

`func (o *AFNORErrorDetail) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AFNORErrorDetail) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AFNORErrorDetail) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AFNORErrorDetail) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *AFNORErrorDetail) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *AFNORErrorDetail) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



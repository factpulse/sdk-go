# SignatureInfoAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | **string** | Signature field name in the PDF | 
**Signer** | Pointer to **NullableString** |  | [optional] 
**SigningDate** | Pointer to **NullableString** |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 
**Location** | Pointer to **NullableString** |  | [optional] 
**IsValid** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewSignatureInfoAPI

`func NewSignatureInfoAPI(fieldName string, ) *SignatureInfoAPI`

NewSignatureInfoAPI instantiates a new SignatureInfoAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureInfoAPIWithDefaults

`func NewSignatureInfoAPIWithDefaults() *SignatureInfoAPI`

NewSignatureInfoAPIWithDefaults instantiates a new SignatureInfoAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *SignatureInfoAPI) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *SignatureInfoAPI) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *SignatureInfoAPI) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetSigner

`func (o *SignatureInfoAPI) GetSigner() string`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *SignatureInfoAPI) GetSignerOk() (*string, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *SignatureInfoAPI) SetSigner(v string)`

SetSigner sets Signer field to given value.

### HasSigner

`func (o *SignatureInfoAPI) HasSigner() bool`

HasSigner returns a boolean if a field has been set.

### SetSignerNil

`func (o *SignatureInfoAPI) SetSignerNil(b bool)`

 SetSignerNil sets the value for Signer to be an explicit nil

### UnsetSigner
`func (o *SignatureInfoAPI) UnsetSigner()`

UnsetSigner ensures that no value is present for Signer, not even an explicit nil
### GetSigningDate

`func (o *SignatureInfoAPI) GetSigningDate() string`

GetSigningDate returns the SigningDate field if non-nil, zero value otherwise.

### GetSigningDateOk

`func (o *SignatureInfoAPI) GetSigningDateOk() (*string, bool)`

GetSigningDateOk returns a tuple with the SigningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningDate

`func (o *SignatureInfoAPI) SetSigningDate(v string)`

SetSigningDate sets SigningDate field to given value.

### HasSigningDate

`func (o *SignatureInfoAPI) HasSigningDate() bool`

HasSigningDate returns a boolean if a field has been set.

### SetSigningDateNil

`func (o *SignatureInfoAPI) SetSigningDateNil(b bool)`

 SetSigningDateNil sets the value for SigningDate to be an explicit nil

### UnsetSigningDate
`func (o *SignatureInfoAPI) UnsetSigningDate()`

UnsetSigningDate ensures that no value is present for SigningDate, not even an explicit nil
### GetReason

`func (o *SignatureInfoAPI) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SignatureInfoAPI) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SignatureInfoAPI) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SignatureInfoAPI) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *SignatureInfoAPI) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *SignatureInfoAPI) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetLocation

`func (o *SignatureInfoAPI) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SignatureInfoAPI) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SignatureInfoAPI) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SignatureInfoAPI) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *SignatureInfoAPI) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *SignatureInfoAPI) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetIsValid

`func (o *SignatureInfoAPI) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *SignatureInfoAPI) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *SignatureInfoAPI) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *SignatureInfoAPI) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### SetIsValidNil

`func (o *SignatureInfoAPI) SetIsValidNil(b bool)`

 SetIsValidNil sets the value for IsValid to be an explicit nil

### UnsetIsValid
`func (o *SignatureInfoAPI) UnsetIsValid()`

UnsetIsValid ensures that no value is present for IsValid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



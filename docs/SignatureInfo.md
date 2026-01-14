# SignatureInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signed** | **bool** | PDF was signed | 
**Cn** | Pointer to **NullableString** |  | [optional] 
**Expiration** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSignatureInfo

`func NewSignatureInfo(signed bool, ) *SignatureInfo`

NewSignatureInfo instantiates a new SignatureInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureInfoWithDefaults

`func NewSignatureInfoWithDefaults() *SignatureInfo`

NewSignatureInfoWithDefaults instantiates a new SignatureInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigned

`func (o *SignatureInfo) GetSigned() bool`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *SignatureInfo) GetSignedOk() (*bool, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *SignatureInfo) SetSigned(v bool)`

SetSigned sets Signed field to given value.


### GetCn

`func (o *SignatureInfo) GetCn() string`

GetCn returns the Cn field if non-nil, zero value otherwise.

### GetCnOk

`func (o *SignatureInfo) GetCnOk() (*string, bool)`

GetCnOk returns a tuple with the Cn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCn

`func (o *SignatureInfo) SetCn(v string)`

SetCn sets Cn field to given value.

### HasCn

`func (o *SignatureInfo) HasCn() bool`

HasCn returns a boolean if a field has been set.

### SetCnNil

`func (o *SignatureInfo) SetCnNil(b bool)`

 SetCnNil sets the value for Cn to be an explicit nil

### UnsetCn
`func (o *SignatureInfo) UnsetCn()`

UnsetCn ensures that no value is present for Cn, not even an explicit nil
### GetExpiration

`func (o *SignatureInfo) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *SignatureInfo) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *SignatureInfo) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *SignatureInfo) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### SetExpirationNil

`func (o *SignatureInfo) SetExpirationNil(b bool)`

 SetExpirationNil sets the value for Expiration to be an explicit nil

### UnsetExpiration
`func (o *SignatureInfo) UnsetExpiration()`

UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



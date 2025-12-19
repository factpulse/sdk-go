# Recipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElectronicAddress** | [**NullableElectronicAddress**](ElectronicAddress.md) |  | 
**ExecutingServiceCode** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Siren** | Pointer to **NullableString** |  | [optional] 
**Siret** | Pointer to **NullableString** |  | [optional] 
**PostalAddress** | Pointer to [**NullablePostalAddress**](PostalAddress.md) |  | [optional] 

## Methods

### NewRecipient

`func NewRecipient(electronicAddress NullableElectronicAddress, ) *Recipient`

NewRecipient instantiates a new Recipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipientWithDefaults

`func NewRecipientWithDefaults() *Recipient`

NewRecipientWithDefaults instantiates a new Recipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElectronicAddress

`func (o *Recipient) GetElectronicAddress() ElectronicAddress`

GetElectronicAddress returns the ElectronicAddress field if non-nil, zero value otherwise.

### GetElectronicAddressOk

`func (o *Recipient) GetElectronicAddressOk() (*ElectronicAddress, bool)`

GetElectronicAddressOk returns a tuple with the ElectronicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectronicAddress

`func (o *Recipient) SetElectronicAddress(v ElectronicAddress)`

SetElectronicAddress sets ElectronicAddress field to given value.


### SetElectronicAddressNil

`func (o *Recipient) SetElectronicAddressNil(b bool)`

 SetElectronicAddressNil sets the value for ElectronicAddress to be an explicit nil

### UnsetElectronicAddress
`func (o *Recipient) UnsetElectronicAddress()`

UnsetElectronicAddress ensures that no value is present for ElectronicAddress, not even an explicit nil
### GetExecutingServiceCode

`func (o *Recipient) GetExecutingServiceCode() string`

GetExecutingServiceCode returns the ExecutingServiceCode field if non-nil, zero value otherwise.

### GetExecutingServiceCodeOk

`func (o *Recipient) GetExecutingServiceCodeOk() (*string, bool)`

GetExecutingServiceCodeOk returns a tuple with the ExecutingServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutingServiceCode

`func (o *Recipient) SetExecutingServiceCode(v string)`

SetExecutingServiceCode sets ExecutingServiceCode field to given value.

### HasExecutingServiceCode

`func (o *Recipient) HasExecutingServiceCode() bool`

HasExecutingServiceCode returns a boolean if a field has been set.

### SetExecutingServiceCodeNil

`func (o *Recipient) SetExecutingServiceCodeNil(b bool)`

 SetExecutingServiceCodeNil sets the value for ExecutingServiceCode to be an explicit nil

### UnsetExecutingServiceCode
`func (o *Recipient) UnsetExecutingServiceCode()`

UnsetExecutingServiceCode ensures that no value is present for ExecutingServiceCode, not even an explicit nil
### GetName

`func (o *Recipient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Recipient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Recipient) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Recipient) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Recipient) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Recipient) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSiren

`func (o *Recipient) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *Recipient) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *Recipient) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *Recipient) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### SetSirenNil

`func (o *Recipient) SetSirenNil(b bool)`

 SetSirenNil sets the value for Siren to be an explicit nil

### UnsetSiren
`func (o *Recipient) UnsetSiren()`

UnsetSiren ensures that no value is present for Siren, not even an explicit nil
### GetSiret

`func (o *Recipient) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *Recipient) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *Recipient) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *Recipient) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### SetSiretNil

`func (o *Recipient) SetSiretNil(b bool)`

 SetSiretNil sets the value for Siret to be an explicit nil

### UnsetSiret
`func (o *Recipient) UnsetSiret()`

UnsetSiret ensures that no value is present for Siret, not even an explicit nil
### GetPostalAddress

`func (o *Recipient) GetPostalAddress() PostalAddress`

GetPostalAddress returns the PostalAddress field if non-nil, zero value otherwise.

### GetPostalAddressOk

`func (o *Recipient) GetPostalAddressOk() (*PostalAddress, bool)`

GetPostalAddressOk returns a tuple with the PostalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalAddress

`func (o *Recipient) SetPostalAddress(v PostalAddress)`

SetPostalAddress sets PostalAddress field to given value.

### HasPostalAddress

`func (o *Recipient) HasPostalAddress() bool`

HasPostalAddress returns a boolean if a field has been set.

### SetPostalAddressNil

`func (o *Recipient) SetPostalAddressNil(b bool)`

 SetPostalAddressNil sets the value for PostalAddress to be an explicit nil

### UnsetPostalAddress
`func (o *Recipient) UnsetPostalAddress()`

UnsetPostalAddress ensures that no value is present for PostalAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



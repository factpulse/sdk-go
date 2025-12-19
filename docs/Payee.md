# Payee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nom** | **string** | Payee name (BT-59). Mandatory. | 
**Siret** | Pointer to **NullableString** |  | [optional] 
**Siren** | Pointer to **NullableString** |  | [optional] 
**ElectronicAddress** | Pointer to [**NullableElectronicAddress**](ElectronicAddress.md) |  | [optional] 
**Iban** | Pointer to **NullableString** |  | [optional] 
**Bic** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPayee

`func NewPayee(nom string, ) *Payee`

NewPayee instantiates a new Payee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayeeWithDefaults

`func NewPayeeWithDefaults() *Payee`

NewPayeeWithDefaults instantiates a new Payee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNom

`func (o *Payee) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *Payee) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *Payee) SetNom(v string)`

SetNom sets Nom field to given value.


### GetSiret

`func (o *Payee) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *Payee) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *Payee) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *Payee) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### SetSiretNil

`func (o *Payee) SetSiretNil(b bool)`

 SetSiretNil sets the value for Siret to be an explicit nil

### UnsetSiret
`func (o *Payee) UnsetSiret()`

UnsetSiret ensures that no value is present for Siret, not even an explicit nil
### GetSiren

`func (o *Payee) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *Payee) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *Payee) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *Payee) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### SetSirenNil

`func (o *Payee) SetSirenNil(b bool)`

 SetSirenNil sets the value for Siren to be an explicit nil

### UnsetSiren
`func (o *Payee) UnsetSiren()`

UnsetSiren ensures that no value is present for Siren, not even an explicit nil
### GetElectronicAddress

`func (o *Payee) GetElectronicAddress() ElectronicAddress`

GetElectronicAddress returns the ElectronicAddress field if non-nil, zero value otherwise.

### GetElectronicAddressOk

`func (o *Payee) GetElectronicAddressOk() (*ElectronicAddress, bool)`

GetElectronicAddressOk returns a tuple with the ElectronicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectronicAddress

`func (o *Payee) SetElectronicAddress(v ElectronicAddress)`

SetElectronicAddress sets ElectronicAddress field to given value.

### HasElectronicAddress

`func (o *Payee) HasElectronicAddress() bool`

HasElectronicAddress returns a boolean if a field has been set.

### SetElectronicAddressNil

`func (o *Payee) SetElectronicAddressNil(b bool)`

 SetElectronicAddressNil sets the value for ElectronicAddress to be an explicit nil

### UnsetElectronicAddress
`func (o *Payee) UnsetElectronicAddress()`

UnsetElectronicAddress ensures that no value is present for ElectronicAddress, not even an explicit nil
### GetIban

`func (o *Payee) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *Payee) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *Payee) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *Payee) HasIban() bool`

HasIban returns a boolean if a field has been set.

### SetIbanNil

`func (o *Payee) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *Payee) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetBic

`func (o *Payee) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *Payee) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *Payee) SetBic(v string)`

SetBic sets Bic field to given value.

### HasBic

`func (o *Payee) HasBic() bool`

HasBic returns a boolean if a field has been set.

### SetBicNil

`func (o *Payee) SetBicNil(b bool)`

 SetBicNil sets the value for Bic to be an explicit nil

### UnsetBic
`func (o *Payee) UnsetBic()`

UnsetBic ensures that no value is present for Bic, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Beneficiaire

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nom** | **string** | Nom du bénéficiaire (BT-59). Obligatoire. | 
**Siret** | Pointer to **NullableString** |  | [optional] 
**Siren** | Pointer to **NullableString** |  | [optional] 
**AdresseElectronique** | Pointer to [**NullableAdresseElectronique**](AdresseElectronique.md) |  | [optional] 
**Iban** | Pointer to **NullableString** |  | [optional] 
**Bic** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBeneficiaire

`func NewBeneficiaire(nom string, ) *Beneficiaire`

NewBeneficiaire instantiates a new Beneficiaire object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBeneficiaireWithDefaults

`func NewBeneficiaireWithDefaults() *Beneficiaire`

NewBeneficiaireWithDefaults instantiates a new Beneficiaire object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNom

`func (o *Beneficiaire) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *Beneficiaire) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *Beneficiaire) SetNom(v string)`

SetNom sets Nom field to given value.


### GetSiret

`func (o *Beneficiaire) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *Beneficiaire) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *Beneficiaire) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *Beneficiaire) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### SetSiretNil

`func (o *Beneficiaire) SetSiretNil(b bool)`

 SetSiretNil sets the value for Siret to be an explicit nil

### UnsetSiret
`func (o *Beneficiaire) UnsetSiret()`

UnsetSiret ensures that no value is present for Siret, not even an explicit nil
### GetSiren

`func (o *Beneficiaire) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *Beneficiaire) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *Beneficiaire) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *Beneficiaire) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### SetSirenNil

`func (o *Beneficiaire) SetSirenNil(b bool)`

 SetSirenNil sets the value for Siren to be an explicit nil

### UnsetSiren
`func (o *Beneficiaire) UnsetSiren()`

UnsetSiren ensures that no value is present for Siren, not even an explicit nil
### GetAdresseElectronique

`func (o *Beneficiaire) GetAdresseElectronique() AdresseElectronique`

GetAdresseElectronique returns the AdresseElectronique field if non-nil, zero value otherwise.

### GetAdresseElectroniqueOk

`func (o *Beneficiaire) GetAdresseElectroniqueOk() (*AdresseElectronique, bool)`

GetAdresseElectroniqueOk returns a tuple with the AdresseElectronique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseElectronique

`func (o *Beneficiaire) SetAdresseElectronique(v AdresseElectronique)`

SetAdresseElectronique sets AdresseElectronique field to given value.

### HasAdresseElectronique

`func (o *Beneficiaire) HasAdresseElectronique() bool`

HasAdresseElectronique returns a boolean if a field has been set.

### SetAdresseElectroniqueNil

`func (o *Beneficiaire) SetAdresseElectroniqueNil(b bool)`

 SetAdresseElectroniqueNil sets the value for AdresseElectronique to be an explicit nil

### UnsetAdresseElectronique
`func (o *Beneficiaire) UnsetAdresseElectronique()`

UnsetAdresseElectronique ensures that no value is present for AdresseElectronique, not even an explicit nil
### GetIban

`func (o *Beneficiaire) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *Beneficiaire) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *Beneficiaire) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *Beneficiaire) HasIban() bool`

HasIban returns a boolean if a field has been set.

### SetIbanNil

`func (o *Beneficiaire) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *Beneficiaire) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetBic

`func (o *Beneficiaire) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *Beneficiaire) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *Beneficiaire) SetBic(v string)`

SetBic sets Bic field to given value.

### HasBic

`func (o *Beneficiaire) HasBic() bool`

HasBic returns a boolean if a field has been set.

### SetBicNil

`func (o *Beneficiaire) SetBicNil(b bool)`

 SetBicNil sets the value for Bic to be an explicit nil

### UnsetBic
`func (o *Beneficiaire) UnsetBic()`

UnsetBic ensures that no value is present for Bic, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



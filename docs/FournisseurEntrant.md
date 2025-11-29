# FournisseurEntrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nom** | **string** | Raison sociale du fournisseur (BT-27) | 
**Siren** | Pointer to **NullableString** |  | [optional] 
**Siret** | Pointer to **NullableString** |  | [optional] 
**NumeroTvaIntra** | Pointer to **NullableString** |  | [optional] 
**AdressePostale** | Pointer to [**NullableAdressePostale**](AdressePostale.md) |  | [optional] 
**AdresseElectronique** | Pointer to [**NullableAdresseElectronique**](AdresseElectronique.md) |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Telephone** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFournisseurEntrant

`func NewFournisseurEntrant(nom string, ) *FournisseurEntrant`

NewFournisseurEntrant instantiates a new FournisseurEntrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFournisseurEntrantWithDefaults

`func NewFournisseurEntrantWithDefaults() *FournisseurEntrant`

NewFournisseurEntrantWithDefaults instantiates a new FournisseurEntrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNom

`func (o *FournisseurEntrant) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *FournisseurEntrant) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *FournisseurEntrant) SetNom(v string)`

SetNom sets Nom field to given value.


### GetSiren

`func (o *FournisseurEntrant) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *FournisseurEntrant) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *FournisseurEntrant) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *FournisseurEntrant) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### SetSirenNil

`func (o *FournisseurEntrant) SetSirenNil(b bool)`

 SetSirenNil sets the value for Siren to be an explicit nil

### UnsetSiren
`func (o *FournisseurEntrant) UnsetSiren()`

UnsetSiren ensures that no value is present for Siren, not even an explicit nil
### GetSiret

`func (o *FournisseurEntrant) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *FournisseurEntrant) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *FournisseurEntrant) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *FournisseurEntrant) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### SetSiretNil

`func (o *FournisseurEntrant) SetSiretNil(b bool)`

 SetSiretNil sets the value for Siret to be an explicit nil

### UnsetSiret
`func (o *FournisseurEntrant) UnsetSiret()`

UnsetSiret ensures that no value is present for Siret, not even an explicit nil
### GetNumeroTvaIntra

`func (o *FournisseurEntrant) GetNumeroTvaIntra() string`

GetNumeroTvaIntra returns the NumeroTvaIntra field if non-nil, zero value otherwise.

### GetNumeroTvaIntraOk

`func (o *FournisseurEntrant) GetNumeroTvaIntraOk() (*string, bool)`

GetNumeroTvaIntraOk returns a tuple with the NumeroTvaIntra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroTvaIntra

`func (o *FournisseurEntrant) SetNumeroTvaIntra(v string)`

SetNumeroTvaIntra sets NumeroTvaIntra field to given value.

### HasNumeroTvaIntra

`func (o *FournisseurEntrant) HasNumeroTvaIntra() bool`

HasNumeroTvaIntra returns a boolean if a field has been set.

### SetNumeroTvaIntraNil

`func (o *FournisseurEntrant) SetNumeroTvaIntraNil(b bool)`

 SetNumeroTvaIntraNil sets the value for NumeroTvaIntra to be an explicit nil

### UnsetNumeroTvaIntra
`func (o *FournisseurEntrant) UnsetNumeroTvaIntra()`

UnsetNumeroTvaIntra ensures that no value is present for NumeroTvaIntra, not even an explicit nil
### GetAdressePostale

`func (o *FournisseurEntrant) GetAdressePostale() AdressePostale`

GetAdressePostale returns the AdressePostale field if non-nil, zero value otherwise.

### GetAdressePostaleOk

`func (o *FournisseurEntrant) GetAdressePostaleOk() (*AdressePostale, bool)`

GetAdressePostaleOk returns a tuple with the AdressePostale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdressePostale

`func (o *FournisseurEntrant) SetAdressePostale(v AdressePostale)`

SetAdressePostale sets AdressePostale field to given value.

### HasAdressePostale

`func (o *FournisseurEntrant) HasAdressePostale() bool`

HasAdressePostale returns a boolean if a field has been set.

### SetAdressePostaleNil

`func (o *FournisseurEntrant) SetAdressePostaleNil(b bool)`

 SetAdressePostaleNil sets the value for AdressePostale to be an explicit nil

### UnsetAdressePostale
`func (o *FournisseurEntrant) UnsetAdressePostale()`

UnsetAdressePostale ensures that no value is present for AdressePostale, not even an explicit nil
### GetAdresseElectronique

`func (o *FournisseurEntrant) GetAdresseElectronique() AdresseElectronique`

GetAdresseElectronique returns the AdresseElectronique field if non-nil, zero value otherwise.

### GetAdresseElectroniqueOk

`func (o *FournisseurEntrant) GetAdresseElectroniqueOk() (*AdresseElectronique, bool)`

GetAdresseElectroniqueOk returns a tuple with the AdresseElectronique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseElectronique

`func (o *FournisseurEntrant) SetAdresseElectronique(v AdresseElectronique)`

SetAdresseElectronique sets AdresseElectronique field to given value.

### HasAdresseElectronique

`func (o *FournisseurEntrant) HasAdresseElectronique() bool`

HasAdresseElectronique returns a boolean if a field has been set.

### SetAdresseElectroniqueNil

`func (o *FournisseurEntrant) SetAdresseElectroniqueNil(b bool)`

 SetAdresseElectroniqueNil sets the value for AdresseElectronique to be an explicit nil

### UnsetAdresseElectronique
`func (o *FournisseurEntrant) UnsetAdresseElectronique()`

UnsetAdresseElectronique ensures that no value is present for AdresseElectronique, not even an explicit nil
### GetEmail

`func (o *FournisseurEntrant) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FournisseurEntrant) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FournisseurEntrant) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FournisseurEntrant) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *FournisseurEntrant) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *FournisseurEntrant) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetTelephone

`func (o *FournisseurEntrant) GetTelephone() string`

GetTelephone returns the Telephone field if non-nil, zero value otherwise.

### GetTelephoneOk

`func (o *FournisseurEntrant) GetTelephoneOk() (*string, bool)`

GetTelephoneOk returns a tuple with the Telephone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephone

`func (o *FournisseurEntrant) SetTelephone(v string)`

SetTelephone sets Telephone field to given value.

### HasTelephone

`func (o *FournisseurEntrant) HasTelephone() bool`

HasTelephone returns a boolean if a field has been set.

### SetTelephoneNil

`func (o *FournisseurEntrant) SetTelephoneNil(b bool)`

 SetTelephoneNil sets the value for Telephone to be an explicit nil

### UnsetTelephone
`func (o *FournisseurEntrant) UnsetTelephone()`

UnsetTelephone ensures that no value is present for Telephone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



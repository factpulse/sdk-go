# Destinataire

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdresseElectronique** | [**NullableAdresseElectronique**](AdresseElectronique.md) |  | 
**CodeServiceExecutant** | Pointer to **NullableString** |  | [optional] 
**Nom** | Pointer to **NullableString** |  | [optional] 
**Siren** | Pointer to **NullableString** |  | [optional] 
**Siret** | Pointer to **NullableString** |  | [optional] 
**AdressePostale** | Pointer to [**NullableAdressePostale**](AdressePostale.md) |  | [optional] 

## Methods

### NewDestinataire

`func NewDestinataire(adresseElectronique NullableAdresseElectronique, ) *Destinataire`

NewDestinataire instantiates a new Destinataire object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinataireWithDefaults

`func NewDestinataireWithDefaults() *Destinataire`

NewDestinataireWithDefaults instantiates a new Destinataire object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdresseElectronique

`func (o *Destinataire) GetAdresseElectronique() AdresseElectronique`

GetAdresseElectronique returns the AdresseElectronique field if non-nil, zero value otherwise.

### GetAdresseElectroniqueOk

`func (o *Destinataire) GetAdresseElectroniqueOk() (*AdresseElectronique, bool)`

GetAdresseElectroniqueOk returns a tuple with the AdresseElectronique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseElectronique

`func (o *Destinataire) SetAdresseElectronique(v AdresseElectronique)`

SetAdresseElectronique sets AdresseElectronique field to given value.


### SetAdresseElectroniqueNil

`func (o *Destinataire) SetAdresseElectroniqueNil(b bool)`

 SetAdresseElectroniqueNil sets the value for AdresseElectronique to be an explicit nil

### UnsetAdresseElectronique
`func (o *Destinataire) UnsetAdresseElectronique()`

UnsetAdresseElectronique ensures that no value is present for AdresseElectronique, not even an explicit nil
### GetCodeServiceExecutant

`func (o *Destinataire) GetCodeServiceExecutant() string`

GetCodeServiceExecutant returns the CodeServiceExecutant field if non-nil, zero value otherwise.

### GetCodeServiceExecutantOk

`func (o *Destinataire) GetCodeServiceExecutantOk() (*string, bool)`

GetCodeServiceExecutantOk returns a tuple with the CodeServiceExecutant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeServiceExecutant

`func (o *Destinataire) SetCodeServiceExecutant(v string)`

SetCodeServiceExecutant sets CodeServiceExecutant field to given value.

### HasCodeServiceExecutant

`func (o *Destinataire) HasCodeServiceExecutant() bool`

HasCodeServiceExecutant returns a boolean if a field has been set.

### SetCodeServiceExecutantNil

`func (o *Destinataire) SetCodeServiceExecutantNil(b bool)`

 SetCodeServiceExecutantNil sets the value for CodeServiceExecutant to be an explicit nil

### UnsetCodeServiceExecutant
`func (o *Destinataire) UnsetCodeServiceExecutant()`

UnsetCodeServiceExecutant ensures that no value is present for CodeServiceExecutant, not even an explicit nil
### GetNom

`func (o *Destinataire) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *Destinataire) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *Destinataire) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *Destinataire) HasNom() bool`

HasNom returns a boolean if a field has been set.

### SetNomNil

`func (o *Destinataire) SetNomNil(b bool)`

 SetNomNil sets the value for Nom to be an explicit nil

### UnsetNom
`func (o *Destinataire) UnsetNom()`

UnsetNom ensures that no value is present for Nom, not even an explicit nil
### GetSiren

`func (o *Destinataire) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *Destinataire) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *Destinataire) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *Destinataire) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### SetSirenNil

`func (o *Destinataire) SetSirenNil(b bool)`

 SetSirenNil sets the value for Siren to be an explicit nil

### UnsetSiren
`func (o *Destinataire) UnsetSiren()`

UnsetSiren ensures that no value is present for Siren, not even an explicit nil
### GetSiret

`func (o *Destinataire) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *Destinataire) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *Destinataire) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *Destinataire) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### SetSiretNil

`func (o *Destinataire) SetSiretNil(b bool)`

 SetSiretNil sets the value for Siret to be an explicit nil

### UnsetSiret
`func (o *Destinataire) UnsetSiret()`

UnsetSiret ensures that no value is present for Siret, not even an explicit nil
### GetAdressePostale

`func (o *Destinataire) GetAdressePostale() AdressePostale`

GetAdressePostale returns the AdressePostale field if non-nil, zero value otherwise.

### GetAdressePostaleOk

`func (o *Destinataire) GetAdressePostaleOk() (*AdressePostale, bool)`

GetAdressePostaleOk returns a tuple with the AdressePostale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdressePostale

`func (o *Destinataire) SetAdressePostale(v AdressePostale)`

SetAdressePostale sets AdressePostale field to given value.

### HasAdressePostale

`func (o *Destinataire) HasAdressePostale() bool`

HasAdressePostale returns a boolean if a field has been set.

### SetAdressePostaleNil

`func (o *Destinataire) SetAdressePostaleNil(b bool)`

 SetAdressePostaleNil sets the value for AdressePostale to be an explicit nil

### UnsetAdressePostale
`func (o *Destinataire) UnsetAdressePostale()`

UnsetAdressePostale ensures that no value is present for AdressePostale, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



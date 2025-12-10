# Fournisseur

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdresseElectronique** | [**NullableAdresseElectronique**](AdresseElectronique.md) |  | 
**IdFournisseur** | **int32** |  | 
**CodeCoordonneesBancairesFournisseur** | Pointer to **NullableInt32** |  | [optional] 
**IdServiceFournisseur** | Pointer to **NullableInt32** |  | [optional] 
**Nom** | Pointer to **NullableString** |  | [optional] 
**Siren** | Pointer to **NullableString** |  | [optional] 
**Siret** | Pointer to **NullableString** |  | [optional] 
**NumeroTvaIntra** | Pointer to **NullableString** |  | [optional] 
**Iban** | Pointer to **NullableString** |  | [optional] 
**AdressePostale** | Pointer to [**NullableAdressePostale**](AdressePostale.md) |  | [optional] 

## Methods

### NewFournisseur

`func NewFournisseur(adresseElectronique NullableAdresseElectronique, idFournisseur int32, ) *Fournisseur`

NewFournisseur instantiates a new Fournisseur object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFournisseurWithDefaults

`func NewFournisseurWithDefaults() *Fournisseur`

NewFournisseurWithDefaults instantiates a new Fournisseur object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdresseElectronique

`func (o *Fournisseur) GetAdresseElectronique() AdresseElectronique`

GetAdresseElectronique returns the AdresseElectronique field if non-nil, zero value otherwise.

### GetAdresseElectroniqueOk

`func (o *Fournisseur) GetAdresseElectroniqueOk() (*AdresseElectronique, bool)`

GetAdresseElectroniqueOk returns a tuple with the AdresseElectronique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseElectronique

`func (o *Fournisseur) SetAdresseElectronique(v AdresseElectronique)`

SetAdresseElectronique sets AdresseElectronique field to given value.


### SetAdresseElectroniqueNil

`func (o *Fournisseur) SetAdresseElectroniqueNil(b bool)`

 SetAdresseElectroniqueNil sets the value for AdresseElectronique to be an explicit nil

### UnsetAdresseElectronique
`func (o *Fournisseur) UnsetAdresseElectronique()`

UnsetAdresseElectronique ensures that no value is present for AdresseElectronique, not even an explicit nil
### GetIdFournisseur

`func (o *Fournisseur) GetIdFournisseur() int32`

GetIdFournisseur returns the IdFournisseur field if non-nil, zero value otherwise.

### GetIdFournisseurOk

`func (o *Fournisseur) GetIdFournisseurOk() (*int32, bool)`

GetIdFournisseurOk returns a tuple with the IdFournisseur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdFournisseur

`func (o *Fournisseur) SetIdFournisseur(v int32)`

SetIdFournisseur sets IdFournisseur field to given value.


### GetCodeCoordonneesBancairesFournisseur

`func (o *Fournisseur) GetCodeCoordonneesBancairesFournisseur() int32`

GetCodeCoordonneesBancairesFournisseur returns the CodeCoordonneesBancairesFournisseur field if non-nil, zero value otherwise.

### GetCodeCoordonneesBancairesFournisseurOk

`func (o *Fournisseur) GetCodeCoordonneesBancairesFournisseurOk() (*int32, bool)`

GetCodeCoordonneesBancairesFournisseurOk returns a tuple with the CodeCoordonneesBancairesFournisseur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeCoordonneesBancairesFournisseur

`func (o *Fournisseur) SetCodeCoordonneesBancairesFournisseur(v int32)`

SetCodeCoordonneesBancairesFournisseur sets CodeCoordonneesBancairesFournisseur field to given value.

### HasCodeCoordonneesBancairesFournisseur

`func (o *Fournisseur) HasCodeCoordonneesBancairesFournisseur() bool`

HasCodeCoordonneesBancairesFournisseur returns a boolean if a field has been set.

### SetCodeCoordonneesBancairesFournisseurNil

`func (o *Fournisseur) SetCodeCoordonneesBancairesFournisseurNil(b bool)`

 SetCodeCoordonneesBancairesFournisseurNil sets the value for CodeCoordonneesBancairesFournisseur to be an explicit nil

### UnsetCodeCoordonneesBancairesFournisseur
`func (o *Fournisseur) UnsetCodeCoordonneesBancairesFournisseur()`

UnsetCodeCoordonneesBancairesFournisseur ensures that no value is present for CodeCoordonneesBancairesFournisseur, not even an explicit nil
### GetIdServiceFournisseur

`func (o *Fournisseur) GetIdServiceFournisseur() int32`

GetIdServiceFournisseur returns the IdServiceFournisseur field if non-nil, zero value otherwise.

### GetIdServiceFournisseurOk

`func (o *Fournisseur) GetIdServiceFournisseurOk() (*int32, bool)`

GetIdServiceFournisseurOk returns a tuple with the IdServiceFournisseur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdServiceFournisseur

`func (o *Fournisseur) SetIdServiceFournisseur(v int32)`

SetIdServiceFournisseur sets IdServiceFournisseur field to given value.

### HasIdServiceFournisseur

`func (o *Fournisseur) HasIdServiceFournisseur() bool`

HasIdServiceFournisseur returns a boolean if a field has been set.

### SetIdServiceFournisseurNil

`func (o *Fournisseur) SetIdServiceFournisseurNil(b bool)`

 SetIdServiceFournisseurNil sets the value for IdServiceFournisseur to be an explicit nil

### UnsetIdServiceFournisseur
`func (o *Fournisseur) UnsetIdServiceFournisseur()`

UnsetIdServiceFournisseur ensures that no value is present for IdServiceFournisseur, not even an explicit nil
### GetNom

`func (o *Fournisseur) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *Fournisseur) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *Fournisseur) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *Fournisseur) HasNom() bool`

HasNom returns a boolean if a field has been set.

### SetNomNil

`func (o *Fournisseur) SetNomNil(b bool)`

 SetNomNil sets the value for Nom to be an explicit nil

### UnsetNom
`func (o *Fournisseur) UnsetNom()`

UnsetNom ensures that no value is present for Nom, not even an explicit nil
### GetSiren

`func (o *Fournisseur) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *Fournisseur) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *Fournisseur) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *Fournisseur) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### SetSirenNil

`func (o *Fournisseur) SetSirenNil(b bool)`

 SetSirenNil sets the value for Siren to be an explicit nil

### UnsetSiren
`func (o *Fournisseur) UnsetSiren()`

UnsetSiren ensures that no value is present for Siren, not even an explicit nil
### GetSiret

`func (o *Fournisseur) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *Fournisseur) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *Fournisseur) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *Fournisseur) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### SetSiretNil

`func (o *Fournisseur) SetSiretNil(b bool)`

 SetSiretNil sets the value for Siret to be an explicit nil

### UnsetSiret
`func (o *Fournisseur) UnsetSiret()`

UnsetSiret ensures that no value is present for Siret, not even an explicit nil
### GetNumeroTvaIntra

`func (o *Fournisseur) GetNumeroTvaIntra() string`

GetNumeroTvaIntra returns the NumeroTvaIntra field if non-nil, zero value otherwise.

### GetNumeroTvaIntraOk

`func (o *Fournisseur) GetNumeroTvaIntraOk() (*string, bool)`

GetNumeroTvaIntraOk returns a tuple with the NumeroTvaIntra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroTvaIntra

`func (o *Fournisseur) SetNumeroTvaIntra(v string)`

SetNumeroTvaIntra sets NumeroTvaIntra field to given value.

### HasNumeroTvaIntra

`func (o *Fournisseur) HasNumeroTvaIntra() bool`

HasNumeroTvaIntra returns a boolean if a field has been set.

### SetNumeroTvaIntraNil

`func (o *Fournisseur) SetNumeroTvaIntraNil(b bool)`

 SetNumeroTvaIntraNil sets the value for NumeroTvaIntra to be an explicit nil

### UnsetNumeroTvaIntra
`func (o *Fournisseur) UnsetNumeroTvaIntra()`

UnsetNumeroTvaIntra ensures that no value is present for NumeroTvaIntra, not even an explicit nil
### GetIban

`func (o *Fournisseur) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *Fournisseur) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *Fournisseur) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *Fournisseur) HasIban() bool`

HasIban returns a boolean if a field has been set.

### SetIbanNil

`func (o *Fournisseur) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *Fournisseur) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetAdressePostale

`func (o *Fournisseur) GetAdressePostale() AdressePostale`

GetAdressePostale returns the AdressePostale field if non-nil, zero value otherwise.

### GetAdressePostaleOk

`func (o *Fournisseur) GetAdressePostaleOk() (*AdressePostale, bool)`

GetAdressePostaleOk returns a tuple with the AdressePostale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdressePostale

`func (o *Fournisseur) SetAdressePostale(v AdressePostale)`

SetAdressePostale sets AdressePostale field to given value.

### HasAdressePostale

`func (o *Fournisseur) HasAdressePostale() bool`

HasAdressePostale returns a boolean if a field has been set.

### SetAdressePostaleNil

`func (o *Fournisseur) SetAdressePostaleNil(b bool)`

 SetAdressePostaleNil sets the value for AdressePostale to be an explicit nil

### UnsetAdressePostale
`func (o *Fournisseur) UnsetAdressePostale()`

UnsetAdressePostale ensures that no value is present for AdressePostale, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



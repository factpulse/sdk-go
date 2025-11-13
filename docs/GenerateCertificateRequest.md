# GenerateCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cn** | Pointer to **string** | Common Name (CN) - Nom du certificat | [optional] [default to "Test Signature FactPulse"]
**Organisation** | Pointer to **string** | Organisation (O) | [optional] [default to "FactPulse Test"]
**Pays** | Pointer to **string** | Code pays ISO 2 lettres (C) | [optional] [default to "FR"]
**Ville** | Pointer to **string** | Ville (L) | [optional] [default to "Paris"]
**Province** | Pointer to **string** | Province/État (ST) | [optional] [default to "Ile-de-France"]
**Email** | Pointer to **NullableString** |  | [optional] 
**DureeJours** | Pointer to **int32** | Durée de validité en jours | [optional] [default to 365]
**TailleCle** | Pointer to **int32** | Taille de la clé RSA en bits | [optional] [default to 2048]
**PassphraseCle** | Pointer to **NullableString** |  | [optional] 
**GenererP12** | Pointer to **bool** | Générer aussi un fichier PKCS#12 (.p12) | [optional] [default to false]
**PassphraseP12** | Pointer to **string** | Passphrase pour le fichier PKCS#12 | [optional] [default to "changeme"]

## Methods

### NewGenerateCertificateRequest

`func NewGenerateCertificateRequest() *GenerateCertificateRequest`

NewGenerateCertificateRequest instantiates a new GenerateCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateCertificateRequestWithDefaults

`func NewGenerateCertificateRequestWithDefaults() *GenerateCertificateRequest`

NewGenerateCertificateRequestWithDefaults instantiates a new GenerateCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCn

`func (o *GenerateCertificateRequest) GetCn() string`

GetCn returns the Cn field if non-nil, zero value otherwise.

### GetCnOk

`func (o *GenerateCertificateRequest) GetCnOk() (*string, bool)`

GetCnOk returns a tuple with the Cn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCn

`func (o *GenerateCertificateRequest) SetCn(v string)`

SetCn sets Cn field to given value.

### HasCn

`func (o *GenerateCertificateRequest) HasCn() bool`

HasCn returns a boolean if a field has been set.

### GetOrganisation

`func (o *GenerateCertificateRequest) GetOrganisation() string`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *GenerateCertificateRequest) GetOrganisationOk() (*string, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *GenerateCertificateRequest) SetOrganisation(v string)`

SetOrganisation sets Organisation field to given value.

### HasOrganisation

`func (o *GenerateCertificateRequest) HasOrganisation() bool`

HasOrganisation returns a boolean if a field has been set.

### GetPays

`func (o *GenerateCertificateRequest) GetPays() string`

GetPays returns the Pays field if non-nil, zero value otherwise.

### GetPaysOk

`func (o *GenerateCertificateRequest) GetPaysOk() (*string, bool)`

GetPaysOk returns a tuple with the Pays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPays

`func (o *GenerateCertificateRequest) SetPays(v string)`

SetPays sets Pays field to given value.

### HasPays

`func (o *GenerateCertificateRequest) HasPays() bool`

HasPays returns a boolean if a field has been set.

### GetVille

`func (o *GenerateCertificateRequest) GetVille() string`

GetVille returns the Ville field if non-nil, zero value otherwise.

### GetVilleOk

`func (o *GenerateCertificateRequest) GetVilleOk() (*string, bool)`

GetVilleOk returns a tuple with the Ville field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVille

`func (o *GenerateCertificateRequest) SetVille(v string)`

SetVille sets Ville field to given value.

### HasVille

`func (o *GenerateCertificateRequest) HasVille() bool`

HasVille returns a boolean if a field has been set.

### GetProvince

`func (o *GenerateCertificateRequest) GetProvince() string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *GenerateCertificateRequest) GetProvinceOk() (*string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *GenerateCertificateRequest) SetProvince(v string)`

SetProvince sets Province field to given value.

### HasProvince

`func (o *GenerateCertificateRequest) HasProvince() bool`

HasProvince returns a boolean if a field has been set.

### GetEmail

`func (o *GenerateCertificateRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GenerateCertificateRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GenerateCertificateRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GenerateCertificateRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *GenerateCertificateRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *GenerateCertificateRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetDureeJours

`func (o *GenerateCertificateRequest) GetDureeJours() int32`

GetDureeJours returns the DureeJours field if non-nil, zero value otherwise.

### GetDureeJoursOk

`func (o *GenerateCertificateRequest) GetDureeJoursOk() (*int32, bool)`

GetDureeJoursOk returns a tuple with the DureeJours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDureeJours

`func (o *GenerateCertificateRequest) SetDureeJours(v int32)`

SetDureeJours sets DureeJours field to given value.

### HasDureeJours

`func (o *GenerateCertificateRequest) HasDureeJours() bool`

HasDureeJours returns a boolean if a field has been set.

### GetTailleCle

`func (o *GenerateCertificateRequest) GetTailleCle() int32`

GetTailleCle returns the TailleCle field if non-nil, zero value otherwise.

### GetTailleCleOk

`func (o *GenerateCertificateRequest) GetTailleCleOk() (*int32, bool)`

GetTailleCleOk returns a tuple with the TailleCle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTailleCle

`func (o *GenerateCertificateRequest) SetTailleCle(v int32)`

SetTailleCle sets TailleCle field to given value.

### HasTailleCle

`func (o *GenerateCertificateRequest) HasTailleCle() bool`

HasTailleCle returns a boolean if a field has been set.

### GetPassphraseCle

`func (o *GenerateCertificateRequest) GetPassphraseCle() string`

GetPassphraseCle returns the PassphraseCle field if non-nil, zero value otherwise.

### GetPassphraseCleOk

`func (o *GenerateCertificateRequest) GetPassphraseCleOk() (*string, bool)`

GetPassphraseCleOk returns a tuple with the PassphraseCle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseCle

`func (o *GenerateCertificateRequest) SetPassphraseCle(v string)`

SetPassphraseCle sets PassphraseCle field to given value.

### HasPassphraseCle

`func (o *GenerateCertificateRequest) HasPassphraseCle() bool`

HasPassphraseCle returns a boolean if a field has been set.

### SetPassphraseCleNil

`func (o *GenerateCertificateRequest) SetPassphraseCleNil(b bool)`

 SetPassphraseCleNil sets the value for PassphraseCle to be an explicit nil

### UnsetPassphraseCle
`func (o *GenerateCertificateRequest) UnsetPassphraseCle()`

UnsetPassphraseCle ensures that no value is present for PassphraseCle, not even an explicit nil
### GetGenererP12

`func (o *GenerateCertificateRequest) GetGenererP12() bool`

GetGenererP12 returns the GenererP12 field if non-nil, zero value otherwise.

### GetGenererP12Ok

`func (o *GenerateCertificateRequest) GetGenererP12Ok() (*bool, bool)`

GetGenererP12Ok returns a tuple with the GenererP12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenererP12

`func (o *GenerateCertificateRequest) SetGenererP12(v bool)`

SetGenererP12 sets GenererP12 field to given value.

### HasGenererP12

`func (o *GenerateCertificateRequest) HasGenererP12() bool`

HasGenererP12 returns a boolean if a field has been set.

### GetPassphraseP12

`func (o *GenerateCertificateRequest) GetPassphraseP12() string`

GetPassphraseP12 returns the PassphraseP12 field if non-nil, zero value otherwise.

### GetPassphraseP12Ok

`func (o *GenerateCertificateRequest) GetPassphraseP12Ok() (*string, bool)`

GetPassphraseP12Ok returns a tuple with the PassphraseP12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseP12

`func (o *GenerateCertificateRequest) SetPassphraseP12(v string)`

SetPassphraseP12 sets PassphraseP12 field to given value.

### HasPassphraseP12

`func (o *GenerateCertificateRequest) HasPassphraseP12() bool`

HasPassphraseP12 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



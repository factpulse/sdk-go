# CertificateInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cn** | **string** | Common Name | 
**Organisation** | **string** | Organisation | 
**Pays** | **string** | Code pays | 
**Ville** | **string** | Ville | 
**Province** | **string** | Province | 
**Email** | Pointer to **NullableString** |  | [optional] 
**Sujet** | **string** | Sujet complet (RFC4514) | 
**Emetteur** | **string** | Émetteur (auto-signé &#x3D; même que sujet) | 
**NumeroSerie** | **int32** | Numéro de série du certificat | 
**ValideDu** | **string** | Date de début de validité (ISO 8601) | 
**ValideAu** | **string** | Date de fin de validité (ISO 8601) | 
**Algorithme** | **string** | Algorithme de signature | 

## Methods

### NewCertificateInfoResponse

`func NewCertificateInfoResponse(cn string, organisation string, pays string, ville string, province string, sujet string, emetteur string, numeroSerie int32, valideDu string, valideAu string, algorithme string, ) *CertificateInfoResponse`

NewCertificateInfoResponse instantiates a new CertificateInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateInfoResponseWithDefaults

`func NewCertificateInfoResponseWithDefaults() *CertificateInfoResponse`

NewCertificateInfoResponseWithDefaults instantiates a new CertificateInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCn

`func (o *CertificateInfoResponse) GetCn() string`

GetCn returns the Cn field if non-nil, zero value otherwise.

### GetCnOk

`func (o *CertificateInfoResponse) GetCnOk() (*string, bool)`

GetCnOk returns a tuple with the Cn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCn

`func (o *CertificateInfoResponse) SetCn(v string)`

SetCn sets Cn field to given value.


### GetOrganisation

`func (o *CertificateInfoResponse) GetOrganisation() string`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *CertificateInfoResponse) GetOrganisationOk() (*string, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *CertificateInfoResponse) SetOrganisation(v string)`

SetOrganisation sets Organisation field to given value.


### GetPays

`func (o *CertificateInfoResponse) GetPays() string`

GetPays returns the Pays field if non-nil, zero value otherwise.

### GetPaysOk

`func (o *CertificateInfoResponse) GetPaysOk() (*string, bool)`

GetPaysOk returns a tuple with the Pays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPays

`func (o *CertificateInfoResponse) SetPays(v string)`

SetPays sets Pays field to given value.


### GetVille

`func (o *CertificateInfoResponse) GetVille() string`

GetVille returns the Ville field if non-nil, zero value otherwise.

### GetVilleOk

`func (o *CertificateInfoResponse) GetVilleOk() (*string, bool)`

GetVilleOk returns a tuple with the Ville field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVille

`func (o *CertificateInfoResponse) SetVille(v string)`

SetVille sets Ville field to given value.


### GetProvince

`func (o *CertificateInfoResponse) GetProvince() string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *CertificateInfoResponse) GetProvinceOk() (*string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *CertificateInfoResponse) SetProvince(v string)`

SetProvince sets Province field to given value.


### GetEmail

`func (o *CertificateInfoResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CertificateInfoResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CertificateInfoResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CertificateInfoResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CertificateInfoResponse) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CertificateInfoResponse) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetSujet

`func (o *CertificateInfoResponse) GetSujet() string`

GetSujet returns the Sujet field if non-nil, zero value otherwise.

### GetSujetOk

`func (o *CertificateInfoResponse) GetSujetOk() (*string, bool)`

GetSujetOk returns a tuple with the Sujet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSujet

`func (o *CertificateInfoResponse) SetSujet(v string)`

SetSujet sets Sujet field to given value.


### GetEmetteur

`func (o *CertificateInfoResponse) GetEmetteur() string`

GetEmetteur returns the Emetteur field if non-nil, zero value otherwise.

### GetEmetteurOk

`func (o *CertificateInfoResponse) GetEmetteurOk() (*string, bool)`

GetEmetteurOk returns a tuple with the Emetteur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmetteur

`func (o *CertificateInfoResponse) SetEmetteur(v string)`

SetEmetteur sets Emetteur field to given value.


### GetNumeroSerie

`func (o *CertificateInfoResponse) GetNumeroSerie() int32`

GetNumeroSerie returns the NumeroSerie field if non-nil, zero value otherwise.

### GetNumeroSerieOk

`func (o *CertificateInfoResponse) GetNumeroSerieOk() (*int32, bool)`

GetNumeroSerieOk returns a tuple with the NumeroSerie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroSerie

`func (o *CertificateInfoResponse) SetNumeroSerie(v int32)`

SetNumeroSerie sets NumeroSerie field to given value.


### GetValideDu

`func (o *CertificateInfoResponse) GetValideDu() string`

GetValideDu returns the ValideDu field if non-nil, zero value otherwise.

### GetValideDuOk

`func (o *CertificateInfoResponse) GetValideDuOk() (*string, bool)`

GetValideDuOk returns a tuple with the ValideDu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValideDu

`func (o *CertificateInfoResponse) SetValideDu(v string)`

SetValideDu sets ValideDu field to given value.


### GetValideAu

`func (o *CertificateInfoResponse) GetValideAu() string`

GetValideAu returns the ValideAu field if non-nil, zero value otherwise.

### GetValideAuOk

`func (o *CertificateInfoResponse) GetValideAuOk() (*string, bool)`

GetValideAuOk returns a tuple with the ValideAu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValideAu

`func (o *CertificateInfoResponse) SetValideAu(v string)`

SetValideAu sets ValideAu field to given value.


### GetAlgorithme

`func (o *CertificateInfoResponse) GetAlgorithme() string`

GetAlgorithme returns the Algorithme field if non-nil, zero value otherwise.

### GetAlgorithmeOk

`func (o *CertificateInfoResponse) GetAlgorithmeOk() (*string, bool)`

GetAlgorithmeOk returns a tuple with the Algorithme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithme

`func (o *CertificateInfoResponse) SetAlgorithme(v string)`

SetAlgorithme sets Algorithme field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ObtenirIdChorusProRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**NullableChorusProCredentials**](ChorusProCredentials.md) |  | [optional] 
**Siret** | **string** | SIRET de la structure (14 chiffres) | 
**TypeIdentifiant** | Pointer to **string** | Type d&#39;identifiant (SIRET, SIREN, UE_HORS_FRANCE, etc.) | [optional] [default to "SIRET"]

## Methods

### NewObtenirIdChorusProRequest

`func NewObtenirIdChorusProRequest(siret string, ) *ObtenirIdChorusProRequest`

NewObtenirIdChorusProRequest instantiates a new ObtenirIdChorusProRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObtenirIdChorusProRequestWithDefaults

`func NewObtenirIdChorusProRequestWithDefaults() *ObtenirIdChorusProRequest`

NewObtenirIdChorusProRequestWithDefaults instantiates a new ObtenirIdChorusProRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *ObtenirIdChorusProRequest) GetCredentials() ChorusProCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ObtenirIdChorusProRequest) GetCredentialsOk() (*ChorusProCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ObtenirIdChorusProRequest) SetCredentials(v ChorusProCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ObtenirIdChorusProRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *ObtenirIdChorusProRequest) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *ObtenirIdChorusProRequest) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetSiret

`func (o *ObtenirIdChorusProRequest) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *ObtenirIdChorusProRequest) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *ObtenirIdChorusProRequest) SetSiret(v string)`

SetSiret sets Siret field to given value.


### GetTypeIdentifiant

`func (o *ObtenirIdChorusProRequest) GetTypeIdentifiant() string`

GetTypeIdentifiant returns the TypeIdentifiant field if non-nil, zero value otherwise.

### GetTypeIdentifiantOk

`func (o *ObtenirIdChorusProRequest) GetTypeIdentifiantOk() (*string, bool)`

GetTypeIdentifiantOk returns a tuple with the TypeIdentifiant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeIdentifiant

`func (o *ObtenirIdChorusProRequest) SetTypeIdentifiant(v string)`

SetTypeIdentifiant sets TypeIdentifiant field to given value.

### HasTypeIdentifiant

`func (o *ObtenirIdChorusProRequest) HasTypeIdentifiant() bool`

HasTypeIdentifiant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



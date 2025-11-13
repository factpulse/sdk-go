# ChorusProCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PisteClientId** | **string** | Client ID PISTE (portail API gouvernement) | 
**PisteClientSecret** | **string** | Client Secret PISTE | 
**ChorusProLogin** | **string** | Login Chorus Pro | 
**ChorusProPassword** | **string** | Mot de passe Chorus Pro | 
**Sandbox** | Pointer to **bool** | Utiliser l&#39;environnement sandbox (true) ou production (false) | [optional] [default to true]

## Methods

### NewChorusProCredentials

`func NewChorusProCredentials(pisteClientId string, pisteClientSecret string, chorusProLogin string, chorusProPassword string, ) *ChorusProCredentials`

NewChorusProCredentials instantiates a new ChorusProCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChorusProCredentialsWithDefaults

`func NewChorusProCredentialsWithDefaults() *ChorusProCredentials`

NewChorusProCredentialsWithDefaults instantiates a new ChorusProCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPisteClientId

`func (o *ChorusProCredentials) GetPisteClientId() string`

GetPisteClientId returns the PisteClientId field if non-nil, zero value otherwise.

### GetPisteClientIdOk

`func (o *ChorusProCredentials) GetPisteClientIdOk() (*string, bool)`

GetPisteClientIdOk returns a tuple with the PisteClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPisteClientId

`func (o *ChorusProCredentials) SetPisteClientId(v string)`

SetPisteClientId sets PisteClientId field to given value.


### GetPisteClientSecret

`func (o *ChorusProCredentials) GetPisteClientSecret() string`

GetPisteClientSecret returns the PisteClientSecret field if non-nil, zero value otherwise.

### GetPisteClientSecretOk

`func (o *ChorusProCredentials) GetPisteClientSecretOk() (*string, bool)`

GetPisteClientSecretOk returns a tuple with the PisteClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPisteClientSecret

`func (o *ChorusProCredentials) SetPisteClientSecret(v string)`

SetPisteClientSecret sets PisteClientSecret field to given value.


### GetChorusProLogin

`func (o *ChorusProCredentials) GetChorusProLogin() string`

GetChorusProLogin returns the ChorusProLogin field if non-nil, zero value otherwise.

### GetChorusProLoginOk

`func (o *ChorusProCredentials) GetChorusProLoginOk() (*string, bool)`

GetChorusProLoginOk returns a tuple with the ChorusProLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChorusProLogin

`func (o *ChorusProCredentials) SetChorusProLogin(v string)`

SetChorusProLogin sets ChorusProLogin field to given value.


### GetChorusProPassword

`func (o *ChorusProCredentials) GetChorusProPassword() string`

GetChorusProPassword returns the ChorusProPassword field if non-nil, zero value otherwise.

### GetChorusProPasswordOk

`func (o *ChorusProCredentials) GetChorusProPasswordOk() (*string, bool)`

GetChorusProPasswordOk returns a tuple with the ChorusProPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChorusProPassword

`func (o *ChorusProCredentials) SetChorusProPassword(v string)`

SetChorusProPassword sets ChorusProPassword field to given value.


### GetSandbox

`func (o *ChorusProCredentials) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *ChorusProCredentials) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *ChorusProCredentials) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *ChorusProCredentials) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



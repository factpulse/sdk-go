# ChorusProCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PisteClientId** | Pointer to **NullableString** |  | [optional] 
**PisteClientSecret** | Pointer to **NullableString** |  | [optional] 
**ChorusLogin** | Pointer to **NullableString** |  | [optional] 
**ChorusPassword** | Pointer to **NullableString** |  | [optional] 
**SandboxMode** | Pointer to **bool** | [MODE 2] Use sandbox mode (default: True) | [optional] [default to true]

## Methods

### NewChorusProCredentials

`func NewChorusProCredentials() *ChorusProCredentials`

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

### HasPisteClientId

`func (o *ChorusProCredentials) HasPisteClientId() bool`

HasPisteClientId returns a boolean if a field has been set.

### SetPisteClientIdNil

`func (o *ChorusProCredentials) SetPisteClientIdNil(b bool)`

 SetPisteClientIdNil sets the value for PisteClientId to be an explicit nil

### UnsetPisteClientId
`func (o *ChorusProCredentials) UnsetPisteClientId()`

UnsetPisteClientId ensures that no value is present for PisteClientId, not even an explicit nil
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

### HasPisteClientSecret

`func (o *ChorusProCredentials) HasPisteClientSecret() bool`

HasPisteClientSecret returns a boolean if a field has been set.

### SetPisteClientSecretNil

`func (o *ChorusProCredentials) SetPisteClientSecretNil(b bool)`

 SetPisteClientSecretNil sets the value for PisteClientSecret to be an explicit nil

### UnsetPisteClientSecret
`func (o *ChorusProCredentials) UnsetPisteClientSecret()`

UnsetPisteClientSecret ensures that no value is present for PisteClientSecret, not even an explicit nil
### GetChorusLogin

`func (o *ChorusProCredentials) GetChorusLogin() string`

GetChorusLogin returns the ChorusLogin field if non-nil, zero value otherwise.

### GetChorusLoginOk

`func (o *ChorusProCredentials) GetChorusLoginOk() (*string, bool)`

GetChorusLoginOk returns a tuple with the ChorusLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChorusLogin

`func (o *ChorusProCredentials) SetChorusLogin(v string)`

SetChorusLogin sets ChorusLogin field to given value.

### HasChorusLogin

`func (o *ChorusProCredentials) HasChorusLogin() bool`

HasChorusLogin returns a boolean if a field has been set.

### SetChorusLoginNil

`func (o *ChorusProCredentials) SetChorusLoginNil(b bool)`

 SetChorusLoginNil sets the value for ChorusLogin to be an explicit nil

### UnsetChorusLogin
`func (o *ChorusProCredentials) UnsetChorusLogin()`

UnsetChorusLogin ensures that no value is present for ChorusLogin, not even an explicit nil
### GetChorusPassword

`func (o *ChorusProCredentials) GetChorusPassword() string`

GetChorusPassword returns the ChorusPassword field if non-nil, zero value otherwise.

### GetChorusPasswordOk

`func (o *ChorusProCredentials) GetChorusPasswordOk() (*string, bool)`

GetChorusPasswordOk returns a tuple with the ChorusPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChorusPassword

`func (o *ChorusProCredentials) SetChorusPassword(v string)`

SetChorusPassword sets ChorusPassword field to given value.

### HasChorusPassword

`func (o *ChorusProCredentials) HasChorusPassword() bool`

HasChorusPassword returns a boolean if a field has been set.

### SetChorusPasswordNil

`func (o *ChorusProCredentials) SetChorusPasswordNil(b bool)`

 SetChorusPasswordNil sets the value for ChorusPassword to be an explicit nil

### UnsetChorusPassword
`func (o *ChorusProCredentials) UnsetChorusPassword()`

UnsetChorusPassword ensures that no value is present for ChorusPassword, not even an explicit nil
### GetSandboxMode

`func (o *ChorusProCredentials) GetSandboxMode() bool`

GetSandboxMode returns the SandboxMode field if non-nil, zero value otherwise.

### GetSandboxModeOk

`func (o *ChorusProCredentials) GetSandboxModeOk() (*bool, bool)`

GetSandboxModeOk returns a tuple with the SandboxMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxMode

`func (o *ChorusProCredentials) SetSandboxMode(v bool)`

SetSandboxMode sets SandboxMode field to given value.

### HasSandboxMode

`func (o *ChorusProCredentials) HasSandboxMode() bool`

HasSandboxMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



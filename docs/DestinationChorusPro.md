# DestinationChorusPro

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "chorus_pro"]
**Credentials** | Pointer to [**NullableCredentialsChorusPro**](CredentialsChorusPro.md) |  | [optional] 

## Methods

### NewDestinationChorusPro

`func NewDestinationChorusPro() *DestinationChorusPro`

NewDestinationChorusPro instantiates a new DestinationChorusPro object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationChorusProWithDefaults

`func NewDestinationChorusProWithDefaults() *DestinationChorusPro`

NewDestinationChorusProWithDefaults instantiates a new DestinationChorusPro object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DestinationChorusPro) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationChorusPro) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationChorusPro) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DestinationChorusPro) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCredentials

`func (o *DestinationChorusPro) GetCredentials() CredentialsChorusPro`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *DestinationChorusPro) GetCredentialsOk() (*CredentialsChorusPro, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *DestinationChorusPro) SetCredentials(v CredentialsChorusPro)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *DestinationChorusPro) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *DestinationChorusPro) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *DestinationChorusPro) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



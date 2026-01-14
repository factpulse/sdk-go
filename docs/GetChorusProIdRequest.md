# GetChorusProIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**NullableChorusProCredentials**](ChorusProCredentials.md) |  | [optional] 
**Siret** | **string** | Structure SIRET (14 digits) | 
**IdentifierType** | Pointer to **string** | Identifier type (SIRET, SIREN, UE_HORS_FRANCE, etc.) | [optional] [default to "SIRET"]

## Methods

### NewGetChorusProIdRequest

`func NewGetChorusProIdRequest(siret string, ) *GetChorusProIdRequest`

NewGetChorusProIdRequest instantiates a new GetChorusProIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetChorusProIdRequestWithDefaults

`func NewGetChorusProIdRequestWithDefaults() *GetChorusProIdRequest`

NewGetChorusProIdRequestWithDefaults instantiates a new GetChorusProIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *GetChorusProIdRequest) GetCredentials() ChorusProCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GetChorusProIdRequest) GetCredentialsOk() (*ChorusProCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GetChorusProIdRequest) SetCredentials(v ChorusProCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *GetChorusProIdRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *GetChorusProIdRequest) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *GetChorusProIdRequest) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetSiret

`func (o *GetChorusProIdRequest) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *GetChorusProIdRequest) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *GetChorusProIdRequest) SetSiret(v string)`

SetSiret sets Siret field to given value.


### GetIdentifierType

`func (o *GetChorusProIdRequest) GetIdentifierType() string`

GetIdentifierType returns the IdentifierType field if non-nil, zero value otherwise.

### GetIdentifierTypeOk

`func (o *GetChorusProIdRequest) GetIdentifierTypeOk() (*string, bool)`

GetIdentifierTypeOk returns a tuple with the IdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierType

`func (o *GetChorusProIdRequest) SetIdentifierType(v string)`

SetIdentifierType sets IdentifierType field to given value.

### HasIdentifierType

`func (o *GetChorusProIdRequest) HasIdentifierType() bool`

HasIdentifierType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



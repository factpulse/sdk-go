# ConsulterFactureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**NullableChorusProCredentials**](ChorusProCredentials.md) |  | [optional] 
**IdentifiantFactureCpp** | **int32** | ID Chorus Pro de la facture | 

## Methods

### NewConsulterFactureRequest

`func NewConsulterFactureRequest(identifiantFactureCpp int32, ) *ConsulterFactureRequest`

NewConsulterFactureRequest instantiates a new ConsulterFactureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsulterFactureRequestWithDefaults

`func NewConsulterFactureRequestWithDefaults() *ConsulterFactureRequest`

NewConsulterFactureRequestWithDefaults instantiates a new ConsulterFactureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *ConsulterFactureRequest) GetCredentials() ChorusProCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ConsulterFactureRequest) GetCredentialsOk() (*ChorusProCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ConsulterFactureRequest) SetCredentials(v ChorusProCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ConsulterFactureRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *ConsulterFactureRequest) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *ConsulterFactureRequest) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetIdentifiantFactureCpp

`func (o *ConsulterFactureRequest) GetIdentifiantFactureCpp() int32`

GetIdentifiantFactureCpp returns the IdentifiantFactureCpp field if non-nil, zero value otherwise.

### GetIdentifiantFactureCppOk

`func (o *ConsulterFactureRequest) GetIdentifiantFactureCppOk() (*int32, bool)`

GetIdentifiantFactureCppOk returns a tuple with the IdentifiantFactureCpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiantFactureCpp

`func (o *ConsulterFactureRequest) SetIdentifiantFactureCpp(v int32)`

SetIdentifiantFactureCpp sets IdentifiantFactureCpp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ConsulterStructureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**NullableChorusProCredentials**](ChorusProCredentials.md) |  | [optional] 
**IdStructureCpp** | **int32** | ID Chorus Pro de la structure | 
**CodeLangue** | Pointer to **string** | Code langue (fr, en) | [optional] [default to "fr"]

## Methods

### NewConsulterStructureRequest

`func NewConsulterStructureRequest(idStructureCpp int32, ) *ConsulterStructureRequest`

NewConsulterStructureRequest instantiates a new ConsulterStructureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsulterStructureRequestWithDefaults

`func NewConsulterStructureRequestWithDefaults() *ConsulterStructureRequest`

NewConsulterStructureRequestWithDefaults instantiates a new ConsulterStructureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *ConsulterStructureRequest) GetCredentials() ChorusProCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ConsulterStructureRequest) GetCredentialsOk() (*ChorusProCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ConsulterStructureRequest) SetCredentials(v ChorusProCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ConsulterStructureRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *ConsulterStructureRequest) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *ConsulterStructureRequest) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetIdStructureCpp

`func (o *ConsulterStructureRequest) GetIdStructureCpp() int32`

GetIdStructureCpp returns the IdStructureCpp field if non-nil, zero value otherwise.

### GetIdStructureCppOk

`func (o *ConsulterStructureRequest) GetIdStructureCppOk() (*int32, bool)`

GetIdStructureCppOk returns a tuple with the IdStructureCpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdStructureCpp

`func (o *ConsulterStructureRequest) SetIdStructureCpp(v int32)`

SetIdStructureCpp sets IdStructureCpp field to given value.


### GetCodeLangue

`func (o *ConsulterStructureRequest) GetCodeLangue() string`

GetCodeLangue returns the CodeLangue field if non-nil, zero value otherwise.

### GetCodeLangueOk

`func (o *ConsulterStructureRequest) GetCodeLangueOk() (*string, bool)`

GetCodeLangueOk returns a tuple with the CodeLangue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeLangue

`func (o *ConsulterStructureRequest) SetCodeLangue(v string)`

SetCodeLangue sets CodeLangue field to given value.

### HasCodeLangue

`func (o *ConsulterStructureRequest) HasCodeLangue() bool`

HasCodeLangue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# RechercherStructureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**NullableChorusProCredentials**](ChorusProCredentials.md) |  | [optional] 
**IdentifiantStructure** | Pointer to **NullableString** |  | [optional] 
**TypeIdentifiantStructure** | Pointer to **NullableString** |  | [optional] 
**RaisonSocialeStructure** | Pointer to **NullableString** |  | [optional] 
**RestreindreStructuresPrivees** | Pointer to **bool** | Limiter la recherche aux structures privées uniquement | [optional] [default to false]

## Methods

### NewRechercherStructureRequest

`func NewRechercherStructureRequest() *RechercherStructureRequest`

NewRechercherStructureRequest instantiates a new RechercherStructureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRechercherStructureRequestWithDefaults

`func NewRechercherStructureRequestWithDefaults() *RechercherStructureRequest`

NewRechercherStructureRequestWithDefaults instantiates a new RechercherStructureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *RechercherStructureRequest) GetCredentials() ChorusProCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *RechercherStructureRequest) GetCredentialsOk() (*ChorusProCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *RechercherStructureRequest) SetCredentials(v ChorusProCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *RechercherStructureRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *RechercherStructureRequest) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *RechercherStructureRequest) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetIdentifiantStructure

`func (o *RechercherStructureRequest) GetIdentifiantStructure() string`

GetIdentifiantStructure returns the IdentifiantStructure field if non-nil, zero value otherwise.

### GetIdentifiantStructureOk

`func (o *RechercherStructureRequest) GetIdentifiantStructureOk() (*string, bool)`

GetIdentifiantStructureOk returns a tuple with the IdentifiantStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiantStructure

`func (o *RechercherStructureRequest) SetIdentifiantStructure(v string)`

SetIdentifiantStructure sets IdentifiantStructure field to given value.

### HasIdentifiantStructure

`func (o *RechercherStructureRequest) HasIdentifiantStructure() bool`

HasIdentifiantStructure returns a boolean if a field has been set.

### SetIdentifiantStructureNil

`func (o *RechercherStructureRequest) SetIdentifiantStructureNil(b bool)`

 SetIdentifiantStructureNil sets the value for IdentifiantStructure to be an explicit nil

### UnsetIdentifiantStructure
`func (o *RechercherStructureRequest) UnsetIdentifiantStructure()`

UnsetIdentifiantStructure ensures that no value is present for IdentifiantStructure, not even an explicit nil
### GetTypeIdentifiantStructure

`func (o *RechercherStructureRequest) GetTypeIdentifiantStructure() string`

GetTypeIdentifiantStructure returns the TypeIdentifiantStructure field if non-nil, zero value otherwise.

### GetTypeIdentifiantStructureOk

`func (o *RechercherStructureRequest) GetTypeIdentifiantStructureOk() (*string, bool)`

GetTypeIdentifiantStructureOk returns a tuple with the TypeIdentifiantStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeIdentifiantStructure

`func (o *RechercherStructureRequest) SetTypeIdentifiantStructure(v string)`

SetTypeIdentifiantStructure sets TypeIdentifiantStructure field to given value.

### HasTypeIdentifiantStructure

`func (o *RechercherStructureRequest) HasTypeIdentifiantStructure() bool`

HasTypeIdentifiantStructure returns a boolean if a field has been set.

### SetTypeIdentifiantStructureNil

`func (o *RechercherStructureRequest) SetTypeIdentifiantStructureNil(b bool)`

 SetTypeIdentifiantStructureNil sets the value for TypeIdentifiantStructure to be an explicit nil

### UnsetTypeIdentifiantStructure
`func (o *RechercherStructureRequest) UnsetTypeIdentifiantStructure()`

UnsetTypeIdentifiantStructure ensures that no value is present for TypeIdentifiantStructure, not even an explicit nil
### GetRaisonSocialeStructure

`func (o *RechercherStructureRequest) GetRaisonSocialeStructure() string`

GetRaisonSocialeStructure returns the RaisonSocialeStructure field if non-nil, zero value otherwise.

### GetRaisonSocialeStructureOk

`func (o *RechercherStructureRequest) GetRaisonSocialeStructureOk() (*string, bool)`

GetRaisonSocialeStructureOk returns a tuple with the RaisonSocialeStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaisonSocialeStructure

`func (o *RechercherStructureRequest) SetRaisonSocialeStructure(v string)`

SetRaisonSocialeStructure sets RaisonSocialeStructure field to given value.

### HasRaisonSocialeStructure

`func (o *RechercherStructureRequest) HasRaisonSocialeStructure() bool`

HasRaisonSocialeStructure returns a boolean if a field has been set.

### SetRaisonSocialeStructureNil

`func (o *RechercherStructureRequest) SetRaisonSocialeStructureNil(b bool)`

 SetRaisonSocialeStructureNil sets the value for RaisonSocialeStructure to be an explicit nil

### UnsetRaisonSocialeStructure
`func (o *RechercherStructureRequest) UnsetRaisonSocialeStructure()`

UnsetRaisonSocialeStructure ensures that no value is present for RaisonSocialeStructure, not even an explicit nil
### GetRestreindreStructuresPrivees

`func (o *RechercherStructureRequest) GetRestreindreStructuresPrivees() bool`

GetRestreindreStructuresPrivees returns the RestreindreStructuresPrivees field if non-nil, zero value otherwise.

### GetRestreindreStructuresPriveesOk

`func (o *RechercherStructureRequest) GetRestreindreStructuresPriveesOk() (*bool, bool)`

GetRestreindreStructuresPriveesOk returns a tuple with the RestreindreStructuresPrivees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestreindreStructuresPrivees

`func (o *RechercherStructureRequest) SetRestreindreStructuresPrivees(v bool)`

SetRestreindreStructuresPrivees sets RestreindreStructuresPrivees field to given value.

### HasRestreindreStructuresPrivees

`func (o *RechercherStructureRequest) HasRestreindreStructuresPrivees() bool`

HasRestreindreStructuresPrivees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SearchStructureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**NullableFactureElectroniqueRestApiSchemasChorusProChorusProCredentials**](FactureElectroniqueRestApiSchemasChorusProChorusProCredentials.md) |  | [optional] 
**StructureIdentifier** | Pointer to **NullableString** |  | [optional] 
**StructureIdentifierType** | Pointer to **NullableString** |  | [optional] 
**CompanyName** | Pointer to **NullableString** |  | [optional] 
**RestrictPrivateStructures** | Pointer to **bool** | Limit search to private structures only | [optional] [default to false]

## Methods

### NewSearchStructureRequest

`func NewSearchStructureRequest() *SearchStructureRequest`

NewSearchStructureRequest instantiates a new SearchStructureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchStructureRequestWithDefaults

`func NewSearchStructureRequestWithDefaults() *SearchStructureRequest`

NewSearchStructureRequestWithDefaults instantiates a new SearchStructureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *SearchStructureRequest) GetCredentials() FactureElectroniqueRestApiSchemasChorusProChorusProCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *SearchStructureRequest) GetCredentialsOk() (*FactureElectroniqueRestApiSchemasChorusProChorusProCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *SearchStructureRequest) SetCredentials(v FactureElectroniqueRestApiSchemasChorusProChorusProCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *SearchStructureRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *SearchStructureRequest) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *SearchStructureRequest) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetStructureIdentifier

`func (o *SearchStructureRequest) GetStructureIdentifier() string`

GetStructureIdentifier returns the StructureIdentifier field if non-nil, zero value otherwise.

### GetStructureIdentifierOk

`func (o *SearchStructureRequest) GetStructureIdentifierOk() (*string, bool)`

GetStructureIdentifierOk returns a tuple with the StructureIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureIdentifier

`func (o *SearchStructureRequest) SetStructureIdentifier(v string)`

SetStructureIdentifier sets StructureIdentifier field to given value.

### HasStructureIdentifier

`func (o *SearchStructureRequest) HasStructureIdentifier() bool`

HasStructureIdentifier returns a boolean if a field has been set.

### SetStructureIdentifierNil

`func (o *SearchStructureRequest) SetStructureIdentifierNil(b bool)`

 SetStructureIdentifierNil sets the value for StructureIdentifier to be an explicit nil

### UnsetStructureIdentifier
`func (o *SearchStructureRequest) UnsetStructureIdentifier()`

UnsetStructureIdentifier ensures that no value is present for StructureIdentifier, not even an explicit nil
### GetStructureIdentifierType

`func (o *SearchStructureRequest) GetStructureIdentifierType() string`

GetStructureIdentifierType returns the StructureIdentifierType field if non-nil, zero value otherwise.

### GetStructureIdentifierTypeOk

`func (o *SearchStructureRequest) GetStructureIdentifierTypeOk() (*string, bool)`

GetStructureIdentifierTypeOk returns a tuple with the StructureIdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureIdentifierType

`func (o *SearchStructureRequest) SetStructureIdentifierType(v string)`

SetStructureIdentifierType sets StructureIdentifierType field to given value.

### HasStructureIdentifierType

`func (o *SearchStructureRequest) HasStructureIdentifierType() bool`

HasStructureIdentifierType returns a boolean if a field has been set.

### SetStructureIdentifierTypeNil

`func (o *SearchStructureRequest) SetStructureIdentifierTypeNil(b bool)`

 SetStructureIdentifierTypeNil sets the value for StructureIdentifierType to be an explicit nil

### UnsetStructureIdentifierType
`func (o *SearchStructureRequest) UnsetStructureIdentifierType()`

UnsetStructureIdentifierType ensures that no value is present for StructureIdentifierType, not even an explicit nil
### GetCompanyName

`func (o *SearchStructureRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *SearchStructureRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *SearchStructureRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *SearchStructureRequest) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *SearchStructureRequest) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *SearchStructureRequest) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetRestrictPrivateStructures

`func (o *SearchStructureRequest) GetRestrictPrivateStructures() bool`

GetRestrictPrivateStructures returns the RestrictPrivateStructures field if non-nil, zero value otherwise.

### GetRestrictPrivateStructuresOk

`func (o *SearchStructureRequest) GetRestrictPrivateStructuresOk() (*bool, bool)`

GetRestrictPrivateStructuresOk returns a tuple with the RestrictPrivateStructures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictPrivateStructures

`func (o *SearchStructureRequest) SetRestrictPrivateStructures(v bool)`

SetRestrictPrivateStructures sets RestrictPrivateStructures field to given value.

### HasRestrictPrivateStructures

`func (o *SearchStructureRequest) HasRestrictPrivateStructures() bool`

HasRestrictPrivateStructures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



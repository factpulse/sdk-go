# StructureInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdStructureCpp** | **int32** | ID Chorus Pro de la structure | 
**IdentifiantStructure** | **string** | Identifiant (SIRET, SIREN) | 
**DesignationStructure** | **string** | Nom de la structure | 
**TypeIdentifiantStructure** | **string** | Type d&#39;identifiant | 
**Statut** | **string** | Statut (ACTIVE, INACTIVE) | 

## Methods

### NewStructureInfo

`func NewStructureInfo(idStructureCpp int32, identifiantStructure string, designationStructure string, typeIdentifiantStructure string, statut string, ) *StructureInfo`

NewStructureInfo instantiates a new StructureInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructureInfoWithDefaults

`func NewStructureInfoWithDefaults() *StructureInfo`

NewStructureInfoWithDefaults instantiates a new StructureInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdStructureCpp

`func (o *StructureInfo) GetIdStructureCpp() int32`

GetIdStructureCpp returns the IdStructureCpp field if non-nil, zero value otherwise.

### GetIdStructureCppOk

`func (o *StructureInfo) GetIdStructureCppOk() (*int32, bool)`

GetIdStructureCppOk returns a tuple with the IdStructureCpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdStructureCpp

`func (o *StructureInfo) SetIdStructureCpp(v int32)`

SetIdStructureCpp sets IdStructureCpp field to given value.


### GetIdentifiantStructure

`func (o *StructureInfo) GetIdentifiantStructure() string`

GetIdentifiantStructure returns the IdentifiantStructure field if non-nil, zero value otherwise.

### GetIdentifiantStructureOk

`func (o *StructureInfo) GetIdentifiantStructureOk() (*string, bool)`

GetIdentifiantStructureOk returns a tuple with the IdentifiantStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiantStructure

`func (o *StructureInfo) SetIdentifiantStructure(v string)`

SetIdentifiantStructure sets IdentifiantStructure field to given value.


### GetDesignationStructure

`func (o *StructureInfo) GetDesignationStructure() string`

GetDesignationStructure returns the DesignationStructure field if non-nil, zero value otherwise.

### GetDesignationStructureOk

`func (o *StructureInfo) GetDesignationStructureOk() (*string, bool)`

GetDesignationStructureOk returns a tuple with the DesignationStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignationStructure

`func (o *StructureInfo) SetDesignationStructure(v string)`

SetDesignationStructure sets DesignationStructure field to given value.


### GetTypeIdentifiantStructure

`func (o *StructureInfo) GetTypeIdentifiantStructure() string`

GetTypeIdentifiantStructure returns the TypeIdentifiantStructure field if non-nil, zero value otherwise.

### GetTypeIdentifiantStructureOk

`func (o *StructureInfo) GetTypeIdentifiantStructureOk() (*string, bool)`

GetTypeIdentifiantStructureOk returns a tuple with the TypeIdentifiantStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeIdentifiantStructure

`func (o *StructureInfo) SetTypeIdentifiantStructure(v string)`

SetTypeIdentifiantStructure sets TypeIdentifiantStructure field to given value.


### GetStatut

`func (o *StructureInfo) GetStatut() string`

GetStatut returns the Statut field if non-nil, zero value otherwise.

### GetStatutOk

`func (o *StructureInfo) GetStatutOk() (*string, bool)`

GetStatutOk returns a tuple with the Statut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatut

`func (o *StructureInfo) SetStatut(v string)`

SetStatut sets Statut field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# StructureInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StructureId** | **int32** | Chorus Pro structure ID | 
**StructureIdentifier** | **string** | Identifier (SIRET, SIREN) | 
**StructureName** | **string** | Structure name | 
**StructureIdentifierType** | **string** | Identifier type | 
**Status** | **string** | Status (ACTIVE, INACTIVE) | 

## Methods

### NewStructureInfo

`func NewStructureInfo(structureId int32, structureIdentifier string, structureName string, structureIdentifierType string, status string, ) *StructureInfo`

NewStructureInfo instantiates a new StructureInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructureInfoWithDefaults

`func NewStructureInfoWithDefaults() *StructureInfo`

NewStructureInfoWithDefaults instantiates a new StructureInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStructureId

`func (o *StructureInfo) GetStructureId() int32`

GetStructureId returns the StructureId field if non-nil, zero value otherwise.

### GetStructureIdOk

`func (o *StructureInfo) GetStructureIdOk() (*int32, bool)`

GetStructureIdOk returns a tuple with the StructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureId

`func (o *StructureInfo) SetStructureId(v int32)`

SetStructureId sets StructureId field to given value.


### GetStructureIdentifier

`func (o *StructureInfo) GetStructureIdentifier() string`

GetStructureIdentifier returns the StructureIdentifier field if non-nil, zero value otherwise.

### GetStructureIdentifierOk

`func (o *StructureInfo) GetStructureIdentifierOk() (*string, bool)`

GetStructureIdentifierOk returns a tuple with the StructureIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureIdentifier

`func (o *StructureInfo) SetStructureIdentifier(v string)`

SetStructureIdentifier sets StructureIdentifier field to given value.


### GetStructureName

`func (o *StructureInfo) GetStructureName() string`

GetStructureName returns the StructureName field if non-nil, zero value otherwise.

### GetStructureNameOk

`func (o *StructureInfo) GetStructureNameOk() (*string, bool)`

GetStructureNameOk returns a tuple with the StructureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureName

`func (o *StructureInfo) SetStructureName(v string)`

SetStructureName sets StructureName field to given value.


### GetStructureIdentifierType

`func (o *StructureInfo) GetStructureIdentifierType() string`

GetStructureIdentifierType returns the StructureIdentifierType field if non-nil, zero value otherwise.

### GetStructureIdentifierTypeOk

`func (o *StructureInfo) GetStructureIdentifierTypeOk() (*string, bool)`

GetStructureIdentifierTypeOk returns a tuple with the StructureIdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureIdentifierType

`func (o *StructureInfo) SetStructureIdentifierType(v string)`

SetStructureIdentifierType sets StructureIdentifierType field to given value.


### GetStatus

`func (o *StructureInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StructureInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StructureInfo) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetChorusProIdResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StructureId** | **int32** | Chorus Pro ID (0 if not found) | 
**StructureIdentifier** | **string** | Searched SIRET/SIREN | 
**StructureName** | Pointer to **NullableString** |  | [optional] 
**Message** | **string** | Return message | 

## Methods

### NewGetChorusProIdResponse

`func NewGetChorusProIdResponse(structureId int32, structureIdentifier string, message string, ) *GetChorusProIdResponse`

NewGetChorusProIdResponse instantiates a new GetChorusProIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetChorusProIdResponseWithDefaults

`func NewGetChorusProIdResponseWithDefaults() *GetChorusProIdResponse`

NewGetChorusProIdResponseWithDefaults instantiates a new GetChorusProIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStructureId

`func (o *GetChorusProIdResponse) GetStructureId() int32`

GetStructureId returns the StructureId field if non-nil, zero value otherwise.

### GetStructureIdOk

`func (o *GetChorusProIdResponse) GetStructureIdOk() (*int32, bool)`

GetStructureIdOk returns a tuple with the StructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureId

`func (o *GetChorusProIdResponse) SetStructureId(v int32)`

SetStructureId sets StructureId field to given value.


### GetStructureIdentifier

`func (o *GetChorusProIdResponse) GetStructureIdentifier() string`

GetStructureIdentifier returns the StructureIdentifier field if non-nil, zero value otherwise.

### GetStructureIdentifierOk

`func (o *GetChorusProIdResponse) GetStructureIdentifierOk() (*string, bool)`

GetStructureIdentifierOk returns a tuple with the StructureIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureIdentifier

`func (o *GetChorusProIdResponse) SetStructureIdentifier(v string)`

SetStructureIdentifier sets StructureIdentifier field to given value.


### GetStructureName

`func (o *GetChorusProIdResponse) GetStructureName() string`

GetStructureName returns the StructureName field if non-nil, zero value otherwise.

### GetStructureNameOk

`func (o *GetChorusProIdResponse) GetStructureNameOk() (*string, bool)`

GetStructureNameOk returns a tuple with the StructureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureName

`func (o *GetChorusProIdResponse) SetStructureName(v string)`

SetStructureName sets StructureName field to given value.

### HasStructureName

`func (o *GetChorusProIdResponse) HasStructureName() bool`

HasStructureName returns a boolean if a field has been set.

### SetStructureNameNil

`func (o *GetChorusProIdResponse) SetStructureNameNil(b bool)`

 SetStructureNameNil sets the value for StructureName to be an explicit nil

### UnsetStructureName
`func (o *GetChorusProIdResponse) UnsetStructureName()`

UnsetStructureName ensures that no value is present for StructureName, not even an explicit nil
### GetMessage

`func (o *GetChorusProIdResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetChorusProIdResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetChorusProIdResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



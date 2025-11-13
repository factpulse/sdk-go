# ObtenirIdChorusProResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdStructureCpp** | **int32** | ID Chorus Pro (0 si non trouvé) | 
**IdentifiantStructure** | **string** | SIRET/SIREN recherché | 
**DesignationStructure** | Pointer to **NullableString** |  | [optional] 
**Message** | **string** | Message de retour | 

## Methods

### NewObtenirIdChorusProResponse

`func NewObtenirIdChorusProResponse(idStructureCpp int32, identifiantStructure string, message string, ) *ObtenirIdChorusProResponse`

NewObtenirIdChorusProResponse instantiates a new ObtenirIdChorusProResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObtenirIdChorusProResponseWithDefaults

`func NewObtenirIdChorusProResponseWithDefaults() *ObtenirIdChorusProResponse`

NewObtenirIdChorusProResponseWithDefaults instantiates a new ObtenirIdChorusProResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdStructureCpp

`func (o *ObtenirIdChorusProResponse) GetIdStructureCpp() int32`

GetIdStructureCpp returns the IdStructureCpp field if non-nil, zero value otherwise.

### GetIdStructureCppOk

`func (o *ObtenirIdChorusProResponse) GetIdStructureCppOk() (*int32, bool)`

GetIdStructureCppOk returns a tuple with the IdStructureCpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdStructureCpp

`func (o *ObtenirIdChorusProResponse) SetIdStructureCpp(v int32)`

SetIdStructureCpp sets IdStructureCpp field to given value.


### GetIdentifiantStructure

`func (o *ObtenirIdChorusProResponse) GetIdentifiantStructure() string`

GetIdentifiantStructure returns the IdentifiantStructure field if non-nil, zero value otherwise.

### GetIdentifiantStructureOk

`func (o *ObtenirIdChorusProResponse) GetIdentifiantStructureOk() (*string, bool)`

GetIdentifiantStructureOk returns a tuple with the IdentifiantStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiantStructure

`func (o *ObtenirIdChorusProResponse) SetIdentifiantStructure(v string)`

SetIdentifiantStructure sets IdentifiantStructure field to given value.


### GetDesignationStructure

`func (o *ObtenirIdChorusProResponse) GetDesignationStructure() string`

GetDesignationStructure returns the DesignationStructure field if non-nil, zero value otherwise.

### GetDesignationStructureOk

`func (o *ObtenirIdChorusProResponse) GetDesignationStructureOk() (*string, bool)`

GetDesignationStructureOk returns a tuple with the DesignationStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignationStructure

`func (o *ObtenirIdChorusProResponse) SetDesignationStructure(v string)`

SetDesignationStructure sets DesignationStructure field to given value.

### HasDesignationStructure

`func (o *ObtenirIdChorusProResponse) HasDesignationStructure() bool`

HasDesignationStructure returns a boolean if a field has been set.

### SetDesignationStructureNil

`func (o *ObtenirIdChorusProResponse) SetDesignationStructureNil(b bool)`

 SetDesignationStructureNil sets the value for DesignationStructure to be an explicit nil

### UnsetDesignationStructure
`func (o *ObtenirIdChorusProResponse) UnsetDesignationStructure()`

UnsetDesignationStructure ensures that no value is present for DesignationStructure, not even an explicit nil
### GetMessage

`func (o *ObtenirIdChorusProResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ObtenirIdChorusProResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ObtenirIdChorusProResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



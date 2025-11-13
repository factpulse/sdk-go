# RechercherStructureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodeRetour** | **int32** | Code retour (0 &#x3D; succès) | 
**Libelle** | **string** | Message de retour | 
**ListeStructures** | Pointer to [**[]StructureInfo**](StructureInfo.md) |  | [optional] 
**Total** | Pointer to **int32** | Nombre total de résultats | [optional] [default to 0]

## Methods

### NewRechercherStructureResponse

`func NewRechercherStructureResponse(codeRetour int32, libelle string, ) *RechercherStructureResponse`

NewRechercherStructureResponse instantiates a new RechercherStructureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRechercherStructureResponseWithDefaults

`func NewRechercherStructureResponseWithDefaults() *RechercherStructureResponse`

NewRechercherStructureResponseWithDefaults instantiates a new RechercherStructureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodeRetour

`func (o *RechercherStructureResponse) GetCodeRetour() int32`

GetCodeRetour returns the CodeRetour field if non-nil, zero value otherwise.

### GetCodeRetourOk

`func (o *RechercherStructureResponse) GetCodeRetourOk() (*int32, bool)`

GetCodeRetourOk returns a tuple with the CodeRetour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeRetour

`func (o *RechercherStructureResponse) SetCodeRetour(v int32)`

SetCodeRetour sets CodeRetour field to given value.


### GetLibelle

`func (o *RechercherStructureResponse) GetLibelle() string`

GetLibelle returns the Libelle field if non-nil, zero value otherwise.

### GetLibelleOk

`func (o *RechercherStructureResponse) GetLibelleOk() (*string, bool)`

GetLibelleOk returns a tuple with the Libelle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibelle

`func (o *RechercherStructureResponse) SetLibelle(v string)`

SetLibelle sets Libelle field to given value.


### GetListeStructures

`func (o *RechercherStructureResponse) GetListeStructures() []StructureInfo`

GetListeStructures returns the ListeStructures field if non-nil, zero value otherwise.

### GetListeStructuresOk

`func (o *RechercherStructureResponse) GetListeStructuresOk() (*[]StructureInfo, bool)`

GetListeStructuresOk returns a tuple with the ListeStructures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeStructures

`func (o *RechercherStructureResponse) SetListeStructures(v []StructureInfo)`

SetListeStructures sets ListeStructures field to given value.

### HasListeStructures

`func (o *RechercherStructureResponse) HasListeStructures() bool`

HasListeStructures returns a boolean if a field has been set.

### GetTotal

`func (o *RechercherStructureResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RechercherStructureResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RechercherStructureResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *RechercherStructureResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



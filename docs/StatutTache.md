# StatutTache

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdTache** | **string** |  | 
**Statut** | [**StatutCelery**](StatutCelery.md) | Statut Celery de la tâche (PENDING, STARTED, SUCCESS, FAILURE, RETRY) | 
**Resultat** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewStatutTache

`func NewStatutTache(idTache string, statut StatutCelery, ) *StatutTache`

NewStatutTache instantiates a new StatutTache object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatutTacheWithDefaults

`func NewStatutTacheWithDefaults() *StatutTache`

NewStatutTacheWithDefaults instantiates a new StatutTache object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdTache

`func (o *StatutTache) GetIdTache() string`

GetIdTache returns the IdTache field if non-nil, zero value otherwise.

### GetIdTacheOk

`func (o *StatutTache) GetIdTacheOk() (*string, bool)`

GetIdTacheOk returns a tuple with the IdTache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTache

`func (o *StatutTache) SetIdTache(v string)`

SetIdTache sets IdTache field to given value.


### GetStatut

`func (o *StatutTache) GetStatut() StatutCelery`

GetStatut returns the Statut field if non-nil, zero value otherwise.

### GetStatutOk

`func (o *StatutTache) GetStatutOk() (*StatutCelery, bool)`

GetStatutOk returns a tuple with the Statut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatut

`func (o *StatutTache) SetStatut(v StatutCelery)`

SetStatut sets Statut field to given value.


### GetResultat

`func (o *StatutTache) GetResultat() map[string]interface{}`

GetResultat returns the Resultat field if non-nil, zero value otherwise.

### GetResultatOk

`func (o *StatutTache) GetResultatOk() (*map[string]interface{}, bool)`

GetResultatOk returns a tuple with the Resultat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultat

`func (o *StatutTache) SetResultat(v map[string]interface{})`

SetResultat sets Resultat field to given value.

### HasResultat

`func (o *StatutTache) HasResultat() bool`

HasResultat returns a boolean if a field has been set.

### SetResultatNil

`func (o *StatutTache) SetResultatNil(b bool)`

 SetResultatNil sets the value for Resultat to be an explicit nil

### UnsetResultat
`func (o *StatutTache) UnsetResultat()`

UnsetResultat ensures that no value is present for Resultat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



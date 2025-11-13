# ReponseRechercheFlux

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | Nombre total de résultats | 
**Offset** | **int32** | Décalage appliqué | 
**Limit** | **int32** | Limite de résultats | 
**Resultats** | [**[]FluxResume**](FluxResume.md) | Liste des flux trouvés | 

## Methods

### NewReponseRechercheFlux

`func NewReponseRechercheFlux(total int32, offset int32, limit int32, resultats []FluxResume, ) *ReponseRechercheFlux`

NewReponseRechercheFlux instantiates a new ReponseRechercheFlux object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReponseRechercheFluxWithDefaults

`func NewReponseRechercheFluxWithDefaults() *ReponseRechercheFlux`

NewReponseRechercheFluxWithDefaults instantiates a new ReponseRechercheFlux object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ReponseRechercheFlux) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ReponseRechercheFlux) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ReponseRechercheFlux) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetOffset

`func (o *ReponseRechercheFlux) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ReponseRechercheFlux) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ReponseRechercheFlux) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *ReponseRechercheFlux) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ReponseRechercheFlux) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ReponseRechercheFlux) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetResultats

`func (o *ReponseRechercheFlux) GetResultats() []FluxResume`

GetResultats returns the Resultats field if non-nil, zero value otherwise.

### GetResultatsOk

`func (o *ReponseRechercheFlux) GetResultatsOk() (*[]FluxResume, bool)`

GetResultatsOk returns a tuple with the Resultats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultats

`func (o *ReponseRechercheFlux) SetResultats(v []FluxResume)`

SetResultats sets Resultats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



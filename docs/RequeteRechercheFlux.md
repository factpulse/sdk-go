# RequeteRechercheFlux

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateMajApres** | Pointer to **NullableTime** |  | [optional] 
**DateMajAvant** | Pointer to **NullableTime** |  | [optional] 
**TypeFlux** | Pointer to [**[]TypeFlux**](TypeFlux.md) |  | [optional] 
**DirectionFlux** | Pointer to [**[]DirectionFlux**](DirectionFlux.md) |  | [optional] 
**TrackingId** | Pointer to **NullableString** |  | [optional] 
**FlowId** | Pointer to **NullableString** |  | [optional] 
**StatutAcquittement** | Pointer to [**NullableStatutAcquittement**](StatutAcquittement.md) |  | [optional] 
**Offset** | Pointer to **int32** | Décalage pour la pagination | [optional] [default to 0]
**Limit** | Pointer to **int32** | Nombre maximum de résultats (max 100) | [optional] [default to 25]

## Methods

### NewRequeteRechercheFlux

`func NewRequeteRechercheFlux() *RequeteRechercheFlux`

NewRequeteRechercheFlux instantiates a new RequeteRechercheFlux object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequeteRechercheFluxWithDefaults

`func NewRequeteRechercheFluxWithDefaults() *RequeteRechercheFlux`

NewRequeteRechercheFluxWithDefaults instantiates a new RequeteRechercheFlux object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateMajApres

`func (o *RequeteRechercheFlux) GetDateMajApres() time.Time`

GetDateMajApres returns the DateMajApres field if non-nil, zero value otherwise.

### GetDateMajApresOk

`func (o *RequeteRechercheFlux) GetDateMajApresOk() (*time.Time, bool)`

GetDateMajApresOk returns a tuple with the DateMajApres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateMajApres

`func (o *RequeteRechercheFlux) SetDateMajApres(v time.Time)`

SetDateMajApres sets DateMajApres field to given value.

### HasDateMajApres

`func (o *RequeteRechercheFlux) HasDateMajApres() bool`

HasDateMajApres returns a boolean if a field has been set.

### SetDateMajApresNil

`func (o *RequeteRechercheFlux) SetDateMajApresNil(b bool)`

 SetDateMajApresNil sets the value for DateMajApres to be an explicit nil

### UnsetDateMajApres
`func (o *RequeteRechercheFlux) UnsetDateMajApres()`

UnsetDateMajApres ensures that no value is present for DateMajApres, not even an explicit nil
### GetDateMajAvant

`func (o *RequeteRechercheFlux) GetDateMajAvant() time.Time`

GetDateMajAvant returns the DateMajAvant field if non-nil, zero value otherwise.

### GetDateMajAvantOk

`func (o *RequeteRechercheFlux) GetDateMajAvantOk() (*time.Time, bool)`

GetDateMajAvantOk returns a tuple with the DateMajAvant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateMajAvant

`func (o *RequeteRechercheFlux) SetDateMajAvant(v time.Time)`

SetDateMajAvant sets DateMajAvant field to given value.

### HasDateMajAvant

`func (o *RequeteRechercheFlux) HasDateMajAvant() bool`

HasDateMajAvant returns a boolean if a field has been set.

### SetDateMajAvantNil

`func (o *RequeteRechercheFlux) SetDateMajAvantNil(b bool)`

 SetDateMajAvantNil sets the value for DateMajAvant to be an explicit nil

### UnsetDateMajAvant
`func (o *RequeteRechercheFlux) UnsetDateMajAvant()`

UnsetDateMajAvant ensures that no value is present for DateMajAvant, not even an explicit nil
### GetTypeFlux

`func (o *RequeteRechercheFlux) GetTypeFlux() []TypeFlux`

GetTypeFlux returns the TypeFlux field if non-nil, zero value otherwise.

### GetTypeFluxOk

`func (o *RequeteRechercheFlux) GetTypeFluxOk() (*[]TypeFlux, bool)`

GetTypeFluxOk returns a tuple with the TypeFlux field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeFlux

`func (o *RequeteRechercheFlux) SetTypeFlux(v []TypeFlux)`

SetTypeFlux sets TypeFlux field to given value.

### HasTypeFlux

`func (o *RequeteRechercheFlux) HasTypeFlux() bool`

HasTypeFlux returns a boolean if a field has been set.

### SetTypeFluxNil

`func (o *RequeteRechercheFlux) SetTypeFluxNil(b bool)`

 SetTypeFluxNil sets the value for TypeFlux to be an explicit nil

### UnsetTypeFlux
`func (o *RequeteRechercheFlux) UnsetTypeFlux()`

UnsetTypeFlux ensures that no value is present for TypeFlux, not even an explicit nil
### GetDirectionFlux

`func (o *RequeteRechercheFlux) GetDirectionFlux() []DirectionFlux`

GetDirectionFlux returns the DirectionFlux field if non-nil, zero value otherwise.

### GetDirectionFluxOk

`func (o *RequeteRechercheFlux) GetDirectionFluxOk() (*[]DirectionFlux, bool)`

GetDirectionFluxOk returns a tuple with the DirectionFlux field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectionFlux

`func (o *RequeteRechercheFlux) SetDirectionFlux(v []DirectionFlux)`

SetDirectionFlux sets DirectionFlux field to given value.

### HasDirectionFlux

`func (o *RequeteRechercheFlux) HasDirectionFlux() bool`

HasDirectionFlux returns a boolean if a field has been set.

### SetDirectionFluxNil

`func (o *RequeteRechercheFlux) SetDirectionFluxNil(b bool)`

 SetDirectionFluxNil sets the value for DirectionFlux to be an explicit nil

### UnsetDirectionFlux
`func (o *RequeteRechercheFlux) UnsetDirectionFlux()`

UnsetDirectionFlux ensures that no value is present for DirectionFlux, not even an explicit nil
### GetTrackingId

`func (o *RequeteRechercheFlux) GetTrackingId() string`

GetTrackingId returns the TrackingId field if non-nil, zero value otherwise.

### GetTrackingIdOk

`func (o *RequeteRechercheFlux) GetTrackingIdOk() (*string, bool)`

GetTrackingIdOk returns a tuple with the TrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingId

`func (o *RequeteRechercheFlux) SetTrackingId(v string)`

SetTrackingId sets TrackingId field to given value.

### HasTrackingId

`func (o *RequeteRechercheFlux) HasTrackingId() bool`

HasTrackingId returns a boolean if a field has been set.

### SetTrackingIdNil

`func (o *RequeteRechercheFlux) SetTrackingIdNil(b bool)`

 SetTrackingIdNil sets the value for TrackingId to be an explicit nil

### UnsetTrackingId
`func (o *RequeteRechercheFlux) UnsetTrackingId()`

UnsetTrackingId ensures that no value is present for TrackingId, not even an explicit nil
### GetFlowId

`func (o *RequeteRechercheFlux) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *RequeteRechercheFlux) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *RequeteRechercheFlux) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *RequeteRechercheFlux) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### SetFlowIdNil

`func (o *RequeteRechercheFlux) SetFlowIdNil(b bool)`

 SetFlowIdNil sets the value for FlowId to be an explicit nil

### UnsetFlowId
`func (o *RequeteRechercheFlux) UnsetFlowId()`

UnsetFlowId ensures that no value is present for FlowId, not even an explicit nil
### GetStatutAcquittement

`func (o *RequeteRechercheFlux) GetStatutAcquittement() StatutAcquittement`

GetStatutAcquittement returns the StatutAcquittement field if non-nil, zero value otherwise.

### GetStatutAcquittementOk

`func (o *RequeteRechercheFlux) GetStatutAcquittementOk() (*StatutAcquittement, bool)`

GetStatutAcquittementOk returns a tuple with the StatutAcquittement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatutAcquittement

`func (o *RequeteRechercheFlux) SetStatutAcquittement(v StatutAcquittement)`

SetStatutAcquittement sets StatutAcquittement field to given value.

### HasStatutAcquittement

`func (o *RequeteRechercheFlux) HasStatutAcquittement() bool`

HasStatutAcquittement returns a boolean if a field has been set.

### SetStatutAcquittementNil

`func (o *RequeteRechercheFlux) SetStatutAcquittementNil(b bool)`

 SetStatutAcquittementNil sets the value for StatutAcquittement to be an explicit nil

### UnsetStatutAcquittement
`func (o *RequeteRechercheFlux) UnsetStatutAcquittement()`

UnsetStatutAcquittement ensures that no value is present for StatutAcquittement, not even an explicit nil
### GetOffset

`func (o *RequeteRechercheFlux) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *RequeteRechercheFlux) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *RequeteRechercheFlux) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *RequeteRechercheFlux) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *RequeteRechercheFlux) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RequeteRechercheFlux) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RequeteRechercheFlux) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *RequeteRechercheFlux) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



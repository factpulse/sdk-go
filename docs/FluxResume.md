# FluxResume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowId** | **string** |  | 
**TrackingId** | Pointer to **NullableString** |  | [optional] 
**Nom** | **string** |  | 
**TypeFlux** | Pointer to **NullableString** |  | [optional] 
**DirectionFlux** | Pointer to **NullableString** |  | [optional] 
**StatutAcquittement** | Pointer to **NullableString** |  | [optional] 
**DateCreation** | Pointer to **NullableString** |  | [optional] 
**DateMaj** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFluxResume

`func NewFluxResume(flowId string, nom string, ) *FluxResume`

NewFluxResume instantiates a new FluxResume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFluxResumeWithDefaults

`func NewFluxResumeWithDefaults() *FluxResume`

NewFluxResumeWithDefaults instantiates a new FluxResume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowId

`func (o *FluxResume) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *FluxResume) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *FluxResume) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.


### GetTrackingId

`func (o *FluxResume) GetTrackingId() string`

GetTrackingId returns the TrackingId field if non-nil, zero value otherwise.

### GetTrackingIdOk

`func (o *FluxResume) GetTrackingIdOk() (*string, bool)`

GetTrackingIdOk returns a tuple with the TrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingId

`func (o *FluxResume) SetTrackingId(v string)`

SetTrackingId sets TrackingId field to given value.

### HasTrackingId

`func (o *FluxResume) HasTrackingId() bool`

HasTrackingId returns a boolean if a field has been set.

### SetTrackingIdNil

`func (o *FluxResume) SetTrackingIdNil(b bool)`

 SetTrackingIdNil sets the value for TrackingId to be an explicit nil

### UnsetTrackingId
`func (o *FluxResume) UnsetTrackingId()`

UnsetTrackingId ensures that no value is present for TrackingId, not even an explicit nil
### GetNom

`func (o *FluxResume) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *FluxResume) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *FluxResume) SetNom(v string)`

SetNom sets Nom field to given value.


### GetTypeFlux

`func (o *FluxResume) GetTypeFlux() string`

GetTypeFlux returns the TypeFlux field if non-nil, zero value otherwise.

### GetTypeFluxOk

`func (o *FluxResume) GetTypeFluxOk() (*string, bool)`

GetTypeFluxOk returns a tuple with the TypeFlux field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeFlux

`func (o *FluxResume) SetTypeFlux(v string)`

SetTypeFlux sets TypeFlux field to given value.

### HasTypeFlux

`func (o *FluxResume) HasTypeFlux() bool`

HasTypeFlux returns a boolean if a field has been set.

### SetTypeFluxNil

`func (o *FluxResume) SetTypeFluxNil(b bool)`

 SetTypeFluxNil sets the value for TypeFlux to be an explicit nil

### UnsetTypeFlux
`func (o *FluxResume) UnsetTypeFlux()`

UnsetTypeFlux ensures that no value is present for TypeFlux, not even an explicit nil
### GetDirectionFlux

`func (o *FluxResume) GetDirectionFlux() string`

GetDirectionFlux returns the DirectionFlux field if non-nil, zero value otherwise.

### GetDirectionFluxOk

`func (o *FluxResume) GetDirectionFluxOk() (*string, bool)`

GetDirectionFluxOk returns a tuple with the DirectionFlux field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectionFlux

`func (o *FluxResume) SetDirectionFlux(v string)`

SetDirectionFlux sets DirectionFlux field to given value.

### HasDirectionFlux

`func (o *FluxResume) HasDirectionFlux() bool`

HasDirectionFlux returns a boolean if a field has been set.

### SetDirectionFluxNil

`func (o *FluxResume) SetDirectionFluxNil(b bool)`

 SetDirectionFluxNil sets the value for DirectionFlux to be an explicit nil

### UnsetDirectionFlux
`func (o *FluxResume) UnsetDirectionFlux()`

UnsetDirectionFlux ensures that no value is present for DirectionFlux, not even an explicit nil
### GetStatutAcquittement

`func (o *FluxResume) GetStatutAcquittement() string`

GetStatutAcquittement returns the StatutAcquittement field if non-nil, zero value otherwise.

### GetStatutAcquittementOk

`func (o *FluxResume) GetStatutAcquittementOk() (*string, bool)`

GetStatutAcquittementOk returns a tuple with the StatutAcquittement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatutAcquittement

`func (o *FluxResume) SetStatutAcquittement(v string)`

SetStatutAcquittement sets StatutAcquittement field to given value.

### HasStatutAcquittement

`func (o *FluxResume) HasStatutAcquittement() bool`

HasStatutAcquittement returns a boolean if a field has been set.

### SetStatutAcquittementNil

`func (o *FluxResume) SetStatutAcquittementNil(b bool)`

 SetStatutAcquittementNil sets the value for StatutAcquittement to be an explicit nil

### UnsetStatutAcquittement
`func (o *FluxResume) UnsetStatutAcquittement()`

UnsetStatutAcquittement ensures that no value is present for StatutAcquittement, not even an explicit nil
### GetDateCreation

`func (o *FluxResume) GetDateCreation() string`

GetDateCreation returns the DateCreation field if non-nil, zero value otherwise.

### GetDateCreationOk

`func (o *FluxResume) GetDateCreationOk() (*string, bool)`

GetDateCreationOk returns a tuple with the DateCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreation

`func (o *FluxResume) SetDateCreation(v string)`

SetDateCreation sets DateCreation field to given value.

### HasDateCreation

`func (o *FluxResume) HasDateCreation() bool`

HasDateCreation returns a boolean if a field has been set.

### SetDateCreationNil

`func (o *FluxResume) SetDateCreationNil(b bool)`

 SetDateCreationNil sets the value for DateCreation to be an explicit nil

### UnsetDateCreation
`func (o *FluxResume) UnsetDateCreation()`

UnsetDateCreation ensures that no value is present for DateCreation, not even an explicit nil
### GetDateMaj

`func (o *FluxResume) GetDateMaj() string`

GetDateMaj returns the DateMaj field if non-nil, zero value otherwise.

### GetDateMajOk

`func (o *FluxResume) GetDateMajOk() (*string, bool)`

GetDateMajOk returns a tuple with the DateMaj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateMaj

`func (o *FluxResume) SetDateMaj(v string)`

SetDateMaj sets DateMaj field to given value.

### HasDateMaj

`func (o *FluxResume) HasDateMaj() bool`

HasDateMaj returns a boolean if a field has been set.

### SetDateMajNil

`func (o *FluxResume) SetDateMajNil(b bool)`

 SetDateMajNil sets the value for DateMaj to be an explicit nil

### UnsetDateMaj
`func (o *FluxResume) UnsetDateMaj()`

UnsetDateMaj ensures that no value is present for DateMaj, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SearchFlowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAfter** | Pointer to **NullableTime** |  | [optional] 
**UpdatedBefore** | Pointer to **NullableTime** |  | [optional] 
**FlowTypes** | Pointer to [**[]FlowType**](FlowType.md) |  | [optional] 
**FlowDirections** | Pointer to [**[]FlowDirection**](FlowDirection.md) |  | [optional] 
**TrackingId** | Pointer to **NullableString** |  | [optional] 
**FlowId** | Pointer to **NullableString** |  | [optional] 
**AcknowledgmentStatus** | Pointer to [**NullableAcknowledgmentStatus**](AcknowledgmentStatus.md) |  | [optional] 
**Offset** | Pointer to **int32** | Offset for pagination | [optional] [default to 0]
**Limit** | Pointer to **int32** | Maximum number of results (max 100) | [optional] [default to 25]

## Methods

### NewSearchFlowRequest

`func NewSearchFlowRequest() *SearchFlowRequest`

NewSearchFlowRequest instantiates a new SearchFlowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFlowRequestWithDefaults

`func NewSearchFlowRequestWithDefaults() *SearchFlowRequest`

NewSearchFlowRequestWithDefaults instantiates a new SearchFlowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAfter

`func (o *SearchFlowRequest) GetUpdatedAfter() time.Time`

GetUpdatedAfter returns the UpdatedAfter field if non-nil, zero value otherwise.

### GetUpdatedAfterOk

`func (o *SearchFlowRequest) GetUpdatedAfterOk() (*time.Time, bool)`

GetUpdatedAfterOk returns a tuple with the UpdatedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAfter

`func (o *SearchFlowRequest) SetUpdatedAfter(v time.Time)`

SetUpdatedAfter sets UpdatedAfter field to given value.

### HasUpdatedAfter

`func (o *SearchFlowRequest) HasUpdatedAfter() bool`

HasUpdatedAfter returns a boolean if a field has been set.

### SetUpdatedAfterNil

`func (o *SearchFlowRequest) SetUpdatedAfterNil(b bool)`

 SetUpdatedAfterNil sets the value for UpdatedAfter to be an explicit nil

### UnsetUpdatedAfter
`func (o *SearchFlowRequest) UnsetUpdatedAfter()`

UnsetUpdatedAfter ensures that no value is present for UpdatedAfter, not even an explicit nil
### GetUpdatedBefore

`func (o *SearchFlowRequest) GetUpdatedBefore() time.Time`

GetUpdatedBefore returns the UpdatedBefore field if non-nil, zero value otherwise.

### GetUpdatedBeforeOk

`func (o *SearchFlowRequest) GetUpdatedBeforeOk() (*time.Time, bool)`

GetUpdatedBeforeOk returns a tuple with the UpdatedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBefore

`func (o *SearchFlowRequest) SetUpdatedBefore(v time.Time)`

SetUpdatedBefore sets UpdatedBefore field to given value.

### HasUpdatedBefore

`func (o *SearchFlowRequest) HasUpdatedBefore() bool`

HasUpdatedBefore returns a boolean if a field has been set.

### SetUpdatedBeforeNil

`func (o *SearchFlowRequest) SetUpdatedBeforeNil(b bool)`

 SetUpdatedBeforeNil sets the value for UpdatedBefore to be an explicit nil

### UnsetUpdatedBefore
`func (o *SearchFlowRequest) UnsetUpdatedBefore()`

UnsetUpdatedBefore ensures that no value is present for UpdatedBefore, not even an explicit nil
### GetFlowTypes

`func (o *SearchFlowRequest) GetFlowTypes() []FlowType`

GetFlowTypes returns the FlowTypes field if non-nil, zero value otherwise.

### GetFlowTypesOk

`func (o *SearchFlowRequest) GetFlowTypesOk() (*[]FlowType, bool)`

GetFlowTypesOk returns a tuple with the FlowTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowTypes

`func (o *SearchFlowRequest) SetFlowTypes(v []FlowType)`

SetFlowTypes sets FlowTypes field to given value.

### HasFlowTypes

`func (o *SearchFlowRequest) HasFlowTypes() bool`

HasFlowTypes returns a boolean if a field has been set.

### SetFlowTypesNil

`func (o *SearchFlowRequest) SetFlowTypesNil(b bool)`

 SetFlowTypesNil sets the value for FlowTypes to be an explicit nil

### UnsetFlowTypes
`func (o *SearchFlowRequest) UnsetFlowTypes()`

UnsetFlowTypes ensures that no value is present for FlowTypes, not even an explicit nil
### GetFlowDirections

`func (o *SearchFlowRequest) GetFlowDirections() []FlowDirection`

GetFlowDirections returns the FlowDirections field if non-nil, zero value otherwise.

### GetFlowDirectionsOk

`func (o *SearchFlowRequest) GetFlowDirectionsOk() (*[]FlowDirection, bool)`

GetFlowDirectionsOk returns a tuple with the FlowDirections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDirections

`func (o *SearchFlowRequest) SetFlowDirections(v []FlowDirection)`

SetFlowDirections sets FlowDirections field to given value.

### HasFlowDirections

`func (o *SearchFlowRequest) HasFlowDirections() bool`

HasFlowDirections returns a boolean if a field has been set.

### SetFlowDirectionsNil

`func (o *SearchFlowRequest) SetFlowDirectionsNil(b bool)`

 SetFlowDirectionsNil sets the value for FlowDirections to be an explicit nil

### UnsetFlowDirections
`func (o *SearchFlowRequest) UnsetFlowDirections()`

UnsetFlowDirections ensures that no value is present for FlowDirections, not even an explicit nil
### GetTrackingId

`func (o *SearchFlowRequest) GetTrackingId() string`

GetTrackingId returns the TrackingId field if non-nil, zero value otherwise.

### GetTrackingIdOk

`func (o *SearchFlowRequest) GetTrackingIdOk() (*string, bool)`

GetTrackingIdOk returns a tuple with the TrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingId

`func (o *SearchFlowRequest) SetTrackingId(v string)`

SetTrackingId sets TrackingId field to given value.

### HasTrackingId

`func (o *SearchFlowRequest) HasTrackingId() bool`

HasTrackingId returns a boolean if a field has been set.

### SetTrackingIdNil

`func (o *SearchFlowRequest) SetTrackingIdNil(b bool)`

 SetTrackingIdNil sets the value for TrackingId to be an explicit nil

### UnsetTrackingId
`func (o *SearchFlowRequest) UnsetTrackingId()`

UnsetTrackingId ensures that no value is present for TrackingId, not even an explicit nil
### GetFlowId

`func (o *SearchFlowRequest) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *SearchFlowRequest) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *SearchFlowRequest) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *SearchFlowRequest) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### SetFlowIdNil

`func (o *SearchFlowRequest) SetFlowIdNil(b bool)`

 SetFlowIdNil sets the value for FlowId to be an explicit nil

### UnsetFlowId
`func (o *SearchFlowRequest) UnsetFlowId()`

UnsetFlowId ensures that no value is present for FlowId, not even an explicit nil
### GetAcknowledgmentStatus

`func (o *SearchFlowRequest) GetAcknowledgmentStatus() AcknowledgmentStatus`

GetAcknowledgmentStatus returns the AcknowledgmentStatus field if non-nil, zero value otherwise.

### GetAcknowledgmentStatusOk

`func (o *SearchFlowRequest) GetAcknowledgmentStatusOk() (*AcknowledgmentStatus, bool)`

GetAcknowledgmentStatusOk returns a tuple with the AcknowledgmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgmentStatus

`func (o *SearchFlowRequest) SetAcknowledgmentStatus(v AcknowledgmentStatus)`

SetAcknowledgmentStatus sets AcknowledgmentStatus field to given value.

### HasAcknowledgmentStatus

`func (o *SearchFlowRequest) HasAcknowledgmentStatus() bool`

HasAcknowledgmentStatus returns a boolean if a field has been set.

### SetAcknowledgmentStatusNil

`func (o *SearchFlowRequest) SetAcknowledgmentStatusNil(b bool)`

 SetAcknowledgmentStatusNil sets the value for AcknowledgmentStatus to be an explicit nil

### UnsetAcknowledgmentStatus
`func (o *SearchFlowRequest) UnsetAcknowledgmentStatus()`

UnsetAcknowledgmentStatus ensures that no value is present for AcknowledgmentStatus, not even an explicit nil
### GetOffset

`func (o *SearchFlowRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SearchFlowRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SearchFlowRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SearchFlowRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *SearchFlowRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchFlowRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchFlowRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchFlowRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# FlowSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowId** | **string** | Unique flow identifier | 
**TrackingId** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** | Flow name | 
**FlowType** | Pointer to **NullableString** |  | [optional] 
**FlowDirection** | Pointer to **NullableString** |  | [optional] 
**AcknowledgmentStatus** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFlowSummary

`func NewFlowSummary(flowId string, name string, ) *FlowSummary`

NewFlowSummary instantiates a new FlowSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowSummaryWithDefaults

`func NewFlowSummaryWithDefaults() *FlowSummary`

NewFlowSummaryWithDefaults instantiates a new FlowSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowId

`func (o *FlowSummary) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *FlowSummary) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *FlowSummary) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.


### GetTrackingId

`func (o *FlowSummary) GetTrackingId() string`

GetTrackingId returns the TrackingId field if non-nil, zero value otherwise.

### GetTrackingIdOk

`func (o *FlowSummary) GetTrackingIdOk() (*string, bool)`

GetTrackingIdOk returns a tuple with the TrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingId

`func (o *FlowSummary) SetTrackingId(v string)`

SetTrackingId sets TrackingId field to given value.

### HasTrackingId

`func (o *FlowSummary) HasTrackingId() bool`

HasTrackingId returns a boolean if a field has been set.

### SetTrackingIdNil

`func (o *FlowSummary) SetTrackingIdNil(b bool)`

 SetTrackingIdNil sets the value for TrackingId to be an explicit nil

### UnsetTrackingId
`func (o *FlowSummary) UnsetTrackingId()`

UnsetTrackingId ensures that no value is present for TrackingId, not even an explicit nil
### GetName

`func (o *FlowSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlowSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlowSummary) SetName(v string)`

SetName sets Name field to given value.


### GetFlowType

`func (o *FlowSummary) GetFlowType() string`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *FlowSummary) GetFlowTypeOk() (*string, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *FlowSummary) SetFlowType(v string)`

SetFlowType sets FlowType field to given value.

### HasFlowType

`func (o *FlowSummary) HasFlowType() bool`

HasFlowType returns a boolean if a field has been set.

### SetFlowTypeNil

`func (o *FlowSummary) SetFlowTypeNil(b bool)`

 SetFlowTypeNil sets the value for FlowType to be an explicit nil

### UnsetFlowType
`func (o *FlowSummary) UnsetFlowType()`

UnsetFlowType ensures that no value is present for FlowType, not even an explicit nil
### GetFlowDirection

`func (o *FlowSummary) GetFlowDirection() string`

GetFlowDirection returns the FlowDirection field if non-nil, zero value otherwise.

### GetFlowDirectionOk

`func (o *FlowSummary) GetFlowDirectionOk() (*string, bool)`

GetFlowDirectionOk returns a tuple with the FlowDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDirection

`func (o *FlowSummary) SetFlowDirection(v string)`

SetFlowDirection sets FlowDirection field to given value.

### HasFlowDirection

`func (o *FlowSummary) HasFlowDirection() bool`

HasFlowDirection returns a boolean if a field has been set.

### SetFlowDirectionNil

`func (o *FlowSummary) SetFlowDirectionNil(b bool)`

 SetFlowDirectionNil sets the value for FlowDirection to be an explicit nil

### UnsetFlowDirection
`func (o *FlowSummary) UnsetFlowDirection()`

UnsetFlowDirection ensures that no value is present for FlowDirection, not even an explicit nil
### GetAcknowledgmentStatus

`func (o *FlowSummary) GetAcknowledgmentStatus() string`

GetAcknowledgmentStatus returns the AcknowledgmentStatus field if non-nil, zero value otherwise.

### GetAcknowledgmentStatusOk

`func (o *FlowSummary) GetAcknowledgmentStatusOk() (*string, bool)`

GetAcknowledgmentStatusOk returns a tuple with the AcknowledgmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgmentStatus

`func (o *FlowSummary) SetAcknowledgmentStatus(v string)`

SetAcknowledgmentStatus sets AcknowledgmentStatus field to given value.

### HasAcknowledgmentStatus

`func (o *FlowSummary) HasAcknowledgmentStatus() bool`

HasAcknowledgmentStatus returns a boolean if a field has been set.

### SetAcknowledgmentStatusNil

`func (o *FlowSummary) SetAcknowledgmentStatusNil(b bool)`

 SetAcknowledgmentStatusNil sets the value for AcknowledgmentStatus to be an explicit nil

### UnsetAcknowledgmentStatus
`func (o *FlowSummary) UnsetAcknowledgmentStatus()`

UnsetAcknowledgmentStatus ensures that no value is present for AcknowledgmentStatus, not even an explicit nil
### GetCreatedAt

`func (o *FlowSummary) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FlowSummary) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FlowSummary) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FlowSummary) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *FlowSummary) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *FlowSummary) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *FlowSummary) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FlowSummary) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FlowSummary) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FlowSummary) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *FlowSummary) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *FlowSummary) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



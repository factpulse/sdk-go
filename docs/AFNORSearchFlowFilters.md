# AFNORSearchFlowFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAfter** | Pointer to **time.Time** |  | [optional] 
**UpdatedBefore** | Pointer to **time.Time** |  | [optional] 
**ProcessingRule** | Pointer to [**[]AFNORProcessingRule**](AFNORProcessingRule.md) |  | [optional] 
**FlowType** | Pointer to [**[]AFNORFlowType**](AFNORFlowType.md) |  | [optional] 
**FlowDirection** | Pointer to [**[]AFNORFlowDirection**](AFNORFlowDirection.md) |  | [optional] 
**TrackingId** | Pointer to **string** | Unique identifier supporting UUID but not only, for flexibility purpose | [optional] 
**AckStatus** | Pointer to [**AFNORFlowAckStatus**](AFNORFlowAckStatus.md) |  | [optional] 

## Methods

### NewAFNORSearchFlowFilters

`func NewAFNORSearchFlowFilters() *AFNORSearchFlowFilters`

NewAFNORSearchFlowFilters instantiates a new AFNORSearchFlowFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORSearchFlowFiltersWithDefaults

`func NewAFNORSearchFlowFiltersWithDefaults() *AFNORSearchFlowFilters`

NewAFNORSearchFlowFiltersWithDefaults instantiates a new AFNORSearchFlowFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAfter

`func (o *AFNORSearchFlowFilters) GetUpdatedAfter() time.Time`

GetUpdatedAfter returns the UpdatedAfter field if non-nil, zero value otherwise.

### GetUpdatedAfterOk

`func (o *AFNORSearchFlowFilters) GetUpdatedAfterOk() (*time.Time, bool)`

GetUpdatedAfterOk returns a tuple with the UpdatedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAfter

`func (o *AFNORSearchFlowFilters) SetUpdatedAfter(v time.Time)`

SetUpdatedAfter sets UpdatedAfter field to given value.

### HasUpdatedAfter

`func (o *AFNORSearchFlowFilters) HasUpdatedAfter() bool`

HasUpdatedAfter returns a boolean if a field has been set.

### GetUpdatedBefore

`func (o *AFNORSearchFlowFilters) GetUpdatedBefore() time.Time`

GetUpdatedBefore returns the UpdatedBefore field if non-nil, zero value otherwise.

### GetUpdatedBeforeOk

`func (o *AFNORSearchFlowFilters) GetUpdatedBeforeOk() (*time.Time, bool)`

GetUpdatedBeforeOk returns a tuple with the UpdatedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBefore

`func (o *AFNORSearchFlowFilters) SetUpdatedBefore(v time.Time)`

SetUpdatedBefore sets UpdatedBefore field to given value.

### HasUpdatedBefore

`func (o *AFNORSearchFlowFilters) HasUpdatedBefore() bool`

HasUpdatedBefore returns a boolean if a field has been set.

### GetProcessingRule

`func (o *AFNORSearchFlowFilters) GetProcessingRule() []AFNORProcessingRule`

GetProcessingRule returns the ProcessingRule field if non-nil, zero value otherwise.

### GetProcessingRuleOk

`func (o *AFNORSearchFlowFilters) GetProcessingRuleOk() (*[]AFNORProcessingRule, bool)`

GetProcessingRuleOk returns a tuple with the ProcessingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingRule

`func (o *AFNORSearchFlowFilters) SetProcessingRule(v []AFNORProcessingRule)`

SetProcessingRule sets ProcessingRule field to given value.

### HasProcessingRule

`func (o *AFNORSearchFlowFilters) HasProcessingRule() bool`

HasProcessingRule returns a boolean if a field has been set.

### GetFlowType

`func (o *AFNORSearchFlowFilters) GetFlowType() []AFNORFlowType`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *AFNORSearchFlowFilters) GetFlowTypeOk() (*[]AFNORFlowType, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *AFNORSearchFlowFilters) SetFlowType(v []AFNORFlowType)`

SetFlowType sets FlowType field to given value.

### HasFlowType

`func (o *AFNORSearchFlowFilters) HasFlowType() bool`

HasFlowType returns a boolean if a field has been set.

### GetFlowDirection

`func (o *AFNORSearchFlowFilters) GetFlowDirection() []AFNORFlowDirection`

GetFlowDirection returns the FlowDirection field if non-nil, zero value otherwise.

### GetFlowDirectionOk

`func (o *AFNORSearchFlowFilters) GetFlowDirectionOk() (*[]AFNORFlowDirection, bool)`

GetFlowDirectionOk returns a tuple with the FlowDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDirection

`func (o *AFNORSearchFlowFilters) SetFlowDirection(v []AFNORFlowDirection)`

SetFlowDirection sets FlowDirection field to given value.

### HasFlowDirection

`func (o *AFNORSearchFlowFilters) HasFlowDirection() bool`

HasFlowDirection returns a boolean if a field has been set.

### GetTrackingId

`func (o *AFNORSearchFlowFilters) GetTrackingId() string`

GetTrackingId returns the TrackingId field if non-nil, zero value otherwise.

### GetTrackingIdOk

`func (o *AFNORSearchFlowFilters) GetTrackingIdOk() (*string, bool)`

GetTrackingIdOk returns a tuple with the TrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingId

`func (o *AFNORSearchFlowFilters) SetTrackingId(v string)`

SetTrackingId sets TrackingId field to given value.

### HasTrackingId

`func (o *AFNORSearchFlowFilters) HasTrackingId() bool`

HasTrackingId returns a boolean if a field has been set.

### GetAckStatus

`func (o *AFNORSearchFlowFilters) GetAckStatus() AFNORFlowAckStatus`

GetAckStatus returns the AckStatus field if non-nil, zero value otherwise.

### GetAckStatusOk

`func (o *AFNORSearchFlowFilters) GetAckStatusOk() (*AFNORFlowAckStatus, bool)`

GetAckStatusOk returns a tuple with the AckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckStatus

`func (o *AFNORSearchFlowFilters) SetAckStatus(v AFNORFlowAckStatus)`

SetAckStatus sets AckStatus field to given value.

### HasAckStatus

`func (o *AFNORSearchFlowFilters) HasAckStatus() bool`

HasAckStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



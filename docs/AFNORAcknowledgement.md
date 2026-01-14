# AFNORAcknowledgement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**AFNORFlowAckStatus**](AFNORFlowAckStatus.md) |  | 
**Details** | Pointer to [**[]AFNORAcknowledgementDetail**](AFNORAcknowledgementDetail.md) |  | [optional] 

## Methods

### NewAFNORAcknowledgement

`func NewAFNORAcknowledgement(status AFNORFlowAckStatus, ) *AFNORAcknowledgement`

NewAFNORAcknowledgement instantiates a new AFNORAcknowledgement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORAcknowledgementWithDefaults

`func NewAFNORAcknowledgementWithDefaults() *AFNORAcknowledgement`

NewAFNORAcknowledgementWithDefaults instantiates a new AFNORAcknowledgement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AFNORAcknowledgement) GetStatus() AFNORFlowAckStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AFNORAcknowledgement) GetStatusOk() (*AFNORFlowAckStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AFNORAcknowledgement) SetStatus(v AFNORFlowAckStatus)`

SetStatus sets Status field to given value.


### GetDetails

`func (o *AFNORAcknowledgement) GetDetails() []AFNORAcknowledgementDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AFNORAcknowledgement) GetDetailsOk() (*[]AFNORAcknowledgementDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AFNORAcknowledgement) SetDetails(v []AFNORAcknowledgementDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AFNORAcknowledgement) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



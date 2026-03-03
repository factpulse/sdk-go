# AFNORWebhookMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessingRule** | Pointer to [**AFNORProcessingRule**](AFNORProcessingRule.md) |  | [optional] 
**FlowType** | [**AFNORFlowType**](AFNORFlowType.md) |  | 
**FlowDirection** | [**AFNORFlowDirection**](AFNORFlowDirection.md) |  | 
**AckStatus** | Pointer to [**AFNORFlowAckStatus**](AFNORFlowAckStatus.md) |  | [optional] 

## Methods

### NewAFNORWebhookMetadata

`func NewAFNORWebhookMetadata(flowType AFNORFlowType, flowDirection AFNORFlowDirection, ) *AFNORWebhookMetadata`

NewAFNORWebhookMetadata instantiates a new AFNORWebhookMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORWebhookMetadataWithDefaults

`func NewAFNORWebhookMetadataWithDefaults() *AFNORWebhookMetadata`

NewAFNORWebhookMetadataWithDefaults instantiates a new AFNORWebhookMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessingRule

`func (o *AFNORWebhookMetadata) GetProcessingRule() AFNORProcessingRule`

GetProcessingRule returns the ProcessingRule field if non-nil, zero value otherwise.

### GetProcessingRuleOk

`func (o *AFNORWebhookMetadata) GetProcessingRuleOk() (*AFNORProcessingRule, bool)`

GetProcessingRuleOk returns a tuple with the ProcessingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingRule

`func (o *AFNORWebhookMetadata) SetProcessingRule(v AFNORProcessingRule)`

SetProcessingRule sets ProcessingRule field to given value.

### HasProcessingRule

`func (o *AFNORWebhookMetadata) HasProcessingRule() bool`

HasProcessingRule returns a boolean if a field has been set.

### GetFlowType

`func (o *AFNORWebhookMetadata) GetFlowType() AFNORFlowType`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *AFNORWebhookMetadata) GetFlowTypeOk() (*AFNORFlowType, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *AFNORWebhookMetadata) SetFlowType(v AFNORFlowType)`

SetFlowType sets FlowType field to given value.


### GetFlowDirection

`func (o *AFNORWebhookMetadata) GetFlowDirection() AFNORFlowDirection`

GetFlowDirection returns the FlowDirection field if non-nil, zero value otherwise.

### GetFlowDirectionOk

`func (o *AFNORWebhookMetadata) GetFlowDirectionOk() (*AFNORFlowDirection, bool)`

GetFlowDirectionOk returns a tuple with the FlowDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDirection

`func (o *AFNORWebhookMetadata) SetFlowDirection(v AFNORFlowDirection)`

SetFlowDirection sets FlowDirection field to given value.


### GetAckStatus

`func (o *AFNORWebhookMetadata) GetAckStatus() AFNORFlowAckStatus`

GetAckStatus returns the AckStatus field if non-nil, zero value otherwise.

### GetAckStatusOk

`func (o *AFNORWebhookMetadata) GetAckStatusOk() (*AFNORFlowAckStatus, bool)`

GetAckStatusOk returns a tuple with the AckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckStatus

`func (o *AFNORWebhookMetadata) SetAckStatus(v AFNORFlowAckStatus)`

SetAckStatus sets AckStatus field to given value.

### HasAckStatus

`func (o *AFNORWebhookMetadata) HasAckStatus() bool`

HasAckStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



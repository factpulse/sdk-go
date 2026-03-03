# AFNORFlowExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAt** | **time.Time** | The last update date and time of the flow. When the flow is submitted updatedAt is equal to submittedAt. When the flow acknowledgment status is changed updatedAt date and time is updated.  | 
**FlowType** | [**AFNORFlowType**](AFNORFlowType.md) |  | 
**ProcessingRuleSource** | **string** | Says whether the processing rule has been computed or the processing rule was an input parameter | 
**FlowDirection** | [**AFNORFlowDirection**](AFNORFlowDirection.md) |  | 
**Acknowledgement** | [**AFNORAcknowledgement**](AFNORAcknowledgement.md) |  | 

## Methods

### NewAFNORFlowExtension

`func NewAFNORFlowExtension(updatedAt time.Time, flowType AFNORFlowType, processingRuleSource string, flowDirection AFNORFlowDirection, acknowledgement AFNORAcknowledgement, ) *AFNORFlowExtension`

NewAFNORFlowExtension instantiates a new AFNORFlowExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORFlowExtensionWithDefaults

`func NewAFNORFlowExtensionWithDefaults() *AFNORFlowExtension`

NewAFNORFlowExtensionWithDefaults instantiates a new AFNORFlowExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAt

`func (o *AFNORFlowExtension) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AFNORFlowExtension) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AFNORFlowExtension) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetFlowType

`func (o *AFNORFlowExtension) GetFlowType() AFNORFlowType`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *AFNORFlowExtension) GetFlowTypeOk() (*AFNORFlowType, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *AFNORFlowExtension) SetFlowType(v AFNORFlowType)`

SetFlowType sets FlowType field to given value.


### GetProcessingRuleSource

`func (o *AFNORFlowExtension) GetProcessingRuleSource() string`

GetProcessingRuleSource returns the ProcessingRuleSource field if non-nil, zero value otherwise.

### GetProcessingRuleSourceOk

`func (o *AFNORFlowExtension) GetProcessingRuleSourceOk() (*string, bool)`

GetProcessingRuleSourceOk returns a tuple with the ProcessingRuleSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingRuleSource

`func (o *AFNORFlowExtension) SetProcessingRuleSource(v string)`

SetProcessingRuleSource sets ProcessingRuleSource field to given value.


### GetFlowDirection

`func (o *AFNORFlowExtension) GetFlowDirection() AFNORFlowDirection`

GetFlowDirection returns the FlowDirection field if non-nil, zero value otherwise.

### GetFlowDirectionOk

`func (o *AFNORFlowExtension) GetFlowDirectionOk() (*AFNORFlowDirection, bool)`

GetFlowDirectionOk returns a tuple with the FlowDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDirection

`func (o *AFNORFlowExtension) SetFlowDirection(v AFNORFlowDirection)`

SetFlowDirection sets FlowDirection field to given value.


### GetAcknowledgement

`func (o *AFNORFlowExtension) GetAcknowledgement() AFNORAcknowledgement`

GetAcknowledgement returns the Acknowledgement field if non-nil, zero value otherwise.

### GetAcknowledgementOk

`func (o *AFNORFlowExtension) GetAcknowledgementOk() (*AFNORAcknowledgement, bool)`

GetAcknowledgementOk returns a tuple with the Acknowledgement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgement

`func (o *AFNORFlowExtension) SetAcknowledgement(v AFNORAcknowledgement)`

SetAcknowledgement sets Acknowledgement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



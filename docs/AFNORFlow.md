# AFNORFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingId** | Pointer to **string** | The tracking id is an external identifier and is used to track the flow by the sender | [optional] 
**Name** | **string** | Name of the file | 
**ProcessingRule** | Pointer to [**AFNORProcessingRule**](AFNORProcessingRule.md) |  | [optional] 
**FlowSyntax** | [**AFNORFlowSyntax**](AFNORFlowSyntax.md) |  | 
**FlowProfile** | Pointer to [**AFNORFlowProfile**](AFNORFlowProfile.md) |  | [optional] 
**FlowId** | **string** | Unique identifier supporting UUID but not only, for flexibility purpose | 
**SubmittedAt** | **time.Time** | The flow submission date and time (the date and time when the flow was created on the system) This property should be used by the API consumer as a time reference to avoid clock synchronization issues  | 
**UpdatedAt** | **time.Time** | The last update date and time of the flow. When the flow is submitted updatedAt is equal to submittedAt. When the flow acknowledgment status is changed updatedAt date and time is updated.  | 
**FlowType** | [**AFNORFlowType**](AFNORFlowType.md) |  | 
**ProcessingRuleSource** | **string** | Says whether the processing rule has been computed or the processing rule was an input parameter | 
**FlowDirection** | [**AFNORFlowDirection**](AFNORFlowDirection.md) |  | 
**Acknowledgement** | [**AFNORAcknowledgement**](AFNORAcknowledgement.md) |  | 

## Methods

### NewAFNORFlow

`func NewAFNORFlow(name string, flowSyntax AFNORFlowSyntax, flowId string, submittedAt time.Time, updatedAt time.Time, flowType AFNORFlowType, processingRuleSource string, flowDirection AFNORFlowDirection, acknowledgement AFNORAcknowledgement, ) *AFNORFlow`

NewAFNORFlow instantiates a new AFNORFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORFlowWithDefaults

`func NewAFNORFlowWithDefaults() *AFNORFlow`

NewAFNORFlowWithDefaults instantiates a new AFNORFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingId

`func (o *AFNORFlow) GetTrackingId() string`

GetTrackingId returns the TrackingId field if non-nil, zero value otherwise.

### GetTrackingIdOk

`func (o *AFNORFlow) GetTrackingIdOk() (*string, bool)`

GetTrackingIdOk returns a tuple with the TrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingId

`func (o *AFNORFlow) SetTrackingId(v string)`

SetTrackingId sets TrackingId field to given value.

### HasTrackingId

`func (o *AFNORFlow) HasTrackingId() bool`

HasTrackingId returns a boolean if a field has been set.

### GetName

`func (o *AFNORFlow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AFNORFlow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AFNORFlow) SetName(v string)`

SetName sets Name field to given value.


### GetProcessingRule

`func (o *AFNORFlow) GetProcessingRule() AFNORProcessingRule`

GetProcessingRule returns the ProcessingRule field if non-nil, zero value otherwise.

### GetProcessingRuleOk

`func (o *AFNORFlow) GetProcessingRuleOk() (*AFNORProcessingRule, bool)`

GetProcessingRuleOk returns a tuple with the ProcessingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingRule

`func (o *AFNORFlow) SetProcessingRule(v AFNORProcessingRule)`

SetProcessingRule sets ProcessingRule field to given value.

### HasProcessingRule

`func (o *AFNORFlow) HasProcessingRule() bool`

HasProcessingRule returns a boolean if a field has been set.

### GetFlowSyntax

`func (o *AFNORFlow) GetFlowSyntax() AFNORFlowSyntax`

GetFlowSyntax returns the FlowSyntax field if non-nil, zero value otherwise.

### GetFlowSyntaxOk

`func (o *AFNORFlow) GetFlowSyntaxOk() (*AFNORFlowSyntax, bool)`

GetFlowSyntaxOk returns a tuple with the FlowSyntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSyntax

`func (o *AFNORFlow) SetFlowSyntax(v AFNORFlowSyntax)`

SetFlowSyntax sets FlowSyntax field to given value.


### GetFlowProfile

`func (o *AFNORFlow) GetFlowProfile() AFNORFlowProfile`

GetFlowProfile returns the FlowProfile field if non-nil, zero value otherwise.

### GetFlowProfileOk

`func (o *AFNORFlow) GetFlowProfileOk() (*AFNORFlowProfile, bool)`

GetFlowProfileOk returns a tuple with the FlowProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowProfile

`func (o *AFNORFlow) SetFlowProfile(v AFNORFlowProfile)`

SetFlowProfile sets FlowProfile field to given value.

### HasFlowProfile

`func (o *AFNORFlow) HasFlowProfile() bool`

HasFlowProfile returns a boolean if a field has been set.

### GetFlowId

`func (o *AFNORFlow) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *AFNORFlow) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *AFNORFlow) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.


### GetSubmittedAt

`func (o *AFNORFlow) GetSubmittedAt() time.Time`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *AFNORFlow) GetSubmittedAtOk() (*time.Time, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *AFNORFlow) SetSubmittedAt(v time.Time)`

SetSubmittedAt sets SubmittedAt field to given value.


### GetUpdatedAt

`func (o *AFNORFlow) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AFNORFlow) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AFNORFlow) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetFlowType

`func (o *AFNORFlow) GetFlowType() AFNORFlowType`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *AFNORFlow) GetFlowTypeOk() (*AFNORFlowType, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *AFNORFlow) SetFlowType(v AFNORFlowType)`

SetFlowType sets FlowType field to given value.


### GetProcessingRuleSource

`func (o *AFNORFlow) GetProcessingRuleSource() string`

GetProcessingRuleSource returns the ProcessingRuleSource field if non-nil, zero value otherwise.

### GetProcessingRuleSourceOk

`func (o *AFNORFlow) GetProcessingRuleSourceOk() (*string, bool)`

GetProcessingRuleSourceOk returns a tuple with the ProcessingRuleSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingRuleSource

`func (o *AFNORFlow) SetProcessingRuleSource(v string)`

SetProcessingRuleSource sets ProcessingRuleSource field to given value.


### GetFlowDirection

`func (o *AFNORFlow) GetFlowDirection() AFNORFlowDirection`

GetFlowDirection returns the FlowDirection field if non-nil, zero value otherwise.

### GetFlowDirectionOk

`func (o *AFNORFlow) GetFlowDirectionOk() (*AFNORFlowDirection, bool)`

GetFlowDirectionOk returns a tuple with the FlowDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDirection

`func (o *AFNORFlow) SetFlowDirection(v AFNORFlowDirection)`

SetFlowDirection sets FlowDirection field to given value.


### GetAcknowledgement

`func (o *AFNORFlow) GetAcknowledgement() AFNORAcknowledgement`

GetAcknowledgement returns the Acknowledgement field if non-nil, zero value otherwise.

### GetAcknowledgementOk

`func (o *AFNORFlow) GetAcknowledgementOk() (*AFNORAcknowledgement, bool)`

GetAcknowledgementOk returns a tuple with the Acknowledgement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgement

`func (o *AFNORFlow) SetAcknowledgement(v AFNORAcknowledgement)`

SetAcknowledgement sets Acknowledgement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



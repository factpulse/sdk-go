# AFNORCoreFlowInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingId** | Pointer to **string** | The tracking id is an external identifier and is used to track the flow by the sender | [optional] 
**Name** | **string** | Name of the file | 
**ProcessingRule** | Pointer to [**AFNORProcessingRule**](AFNORProcessingRule.md) |  | [optional] 
**FlowSyntax** | [**AFNORFlowSyntax**](AFNORFlowSyntax.md) |  | 
**FlowProfile** | Pointer to [**AFNORFlowProfile**](AFNORFlowProfile.md) |  | [optional] 

## Methods

### NewAFNORCoreFlowInfo

`func NewAFNORCoreFlowInfo(name string, flowSyntax AFNORFlowSyntax, ) *AFNORCoreFlowInfo`

NewAFNORCoreFlowInfo instantiates a new AFNORCoreFlowInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORCoreFlowInfoWithDefaults

`func NewAFNORCoreFlowInfoWithDefaults() *AFNORCoreFlowInfo`

NewAFNORCoreFlowInfoWithDefaults instantiates a new AFNORCoreFlowInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingId

`func (o *AFNORCoreFlowInfo) GetTrackingId() string`

GetTrackingId returns the TrackingId field if non-nil, zero value otherwise.

### GetTrackingIdOk

`func (o *AFNORCoreFlowInfo) GetTrackingIdOk() (*string, bool)`

GetTrackingIdOk returns a tuple with the TrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingId

`func (o *AFNORCoreFlowInfo) SetTrackingId(v string)`

SetTrackingId sets TrackingId field to given value.

### HasTrackingId

`func (o *AFNORCoreFlowInfo) HasTrackingId() bool`

HasTrackingId returns a boolean if a field has been set.

### GetName

`func (o *AFNORCoreFlowInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AFNORCoreFlowInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AFNORCoreFlowInfo) SetName(v string)`

SetName sets Name field to given value.


### GetProcessingRule

`func (o *AFNORCoreFlowInfo) GetProcessingRule() AFNORProcessingRule`

GetProcessingRule returns the ProcessingRule field if non-nil, zero value otherwise.

### GetProcessingRuleOk

`func (o *AFNORCoreFlowInfo) GetProcessingRuleOk() (*AFNORProcessingRule, bool)`

GetProcessingRuleOk returns a tuple with the ProcessingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingRule

`func (o *AFNORCoreFlowInfo) SetProcessingRule(v AFNORProcessingRule)`

SetProcessingRule sets ProcessingRule field to given value.

### HasProcessingRule

`func (o *AFNORCoreFlowInfo) HasProcessingRule() bool`

HasProcessingRule returns a boolean if a field has been set.

### GetFlowSyntax

`func (o *AFNORCoreFlowInfo) GetFlowSyntax() AFNORFlowSyntax`

GetFlowSyntax returns the FlowSyntax field if non-nil, zero value otherwise.

### GetFlowSyntaxOk

`func (o *AFNORCoreFlowInfo) GetFlowSyntaxOk() (*AFNORFlowSyntax, bool)`

GetFlowSyntaxOk returns a tuple with the FlowSyntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSyntax

`func (o *AFNORCoreFlowInfo) SetFlowSyntax(v AFNORFlowSyntax)`

SetFlowSyntax sets FlowSyntax field to given value.


### GetFlowProfile

`func (o *AFNORCoreFlowInfo) GetFlowProfile() AFNORFlowProfile`

GetFlowProfile returns the FlowProfile field if non-nil, zero value otherwise.

### GetFlowProfileOk

`func (o *AFNORCoreFlowInfo) GetFlowProfileOk() (*AFNORFlowProfile, bool)`

GetFlowProfileOk returns a tuple with the FlowProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowProfile

`func (o *AFNORCoreFlowInfo) SetFlowProfile(v AFNORFlowProfile)`

SetFlowProfile sets FlowProfile field to given value.

### HasFlowProfile

`func (o *AFNORCoreFlowInfo) HasFlowProfile() bool`

HasFlowProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



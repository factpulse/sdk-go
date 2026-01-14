# AFNORResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowId** | **string** | Submitted flow identifier | 
**TrackingId** | Pointer to **NullableString** |  | [optional] 
**Sha256** | **string** | SHA-256 hash of submitted file | 
**FlowSyntax** | **string** | Flow syntax | 
**FlowProfile** | Pointer to **NullableString** |  | [optional] 
**FlowType** | Pointer to **NullableString** |  | [optional] 
**ProcessingRule** | Pointer to **NullableString** |  | [optional] 
**ProcessingRuleSource** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAFNORResult

`func NewAFNORResult(flowId string, sha256 string, flowSyntax string, ) *AFNORResult`

NewAFNORResult instantiates a new AFNORResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORResultWithDefaults

`func NewAFNORResultWithDefaults() *AFNORResult`

NewAFNORResultWithDefaults instantiates a new AFNORResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowId

`func (o *AFNORResult) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *AFNORResult) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *AFNORResult) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.


### GetTrackingId

`func (o *AFNORResult) GetTrackingId() string`

GetTrackingId returns the TrackingId field if non-nil, zero value otherwise.

### GetTrackingIdOk

`func (o *AFNORResult) GetTrackingIdOk() (*string, bool)`

GetTrackingIdOk returns a tuple with the TrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingId

`func (o *AFNORResult) SetTrackingId(v string)`

SetTrackingId sets TrackingId field to given value.

### HasTrackingId

`func (o *AFNORResult) HasTrackingId() bool`

HasTrackingId returns a boolean if a field has been set.

### SetTrackingIdNil

`func (o *AFNORResult) SetTrackingIdNil(b bool)`

 SetTrackingIdNil sets the value for TrackingId to be an explicit nil

### UnsetTrackingId
`func (o *AFNORResult) UnsetTrackingId()`

UnsetTrackingId ensures that no value is present for TrackingId, not even an explicit nil
### GetSha256

`func (o *AFNORResult) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *AFNORResult) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *AFNORResult) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.


### GetFlowSyntax

`func (o *AFNORResult) GetFlowSyntax() string`

GetFlowSyntax returns the FlowSyntax field if non-nil, zero value otherwise.

### GetFlowSyntaxOk

`func (o *AFNORResult) GetFlowSyntaxOk() (*string, bool)`

GetFlowSyntaxOk returns a tuple with the FlowSyntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSyntax

`func (o *AFNORResult) SetFlowSyntax(v string)`

SetFlowSyntax sets FlowSyntax field to given value.


### GetFlowProfile

`func (o *AFNORResult) GetFlowProfile() string`

GetFlowProfile returns the FlowProfile field if non-nil, zero value otherwise.

### GetFlowProfileOk

`func (o *AFNORResult) GetFlowProfileOk() (*string, bool)`

GetFlowProfileOk returns a tuple with the FlowProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowProfile

`func (o *AFNORResult) SetFlowProfile(v string)`

SetFlowProfile sets FlowProfile field to given value.

### HasFlowProfile

`func (o *AFNORResult) HasFlowProfile() bool`

HasFlowProfile returns a boolean if a field has been set.

### SetFlowProfileNil

`func (o *AFNORResult) SetFlowProfileNil(b bool)`

 SetFlowProfileNil sets the value for FlowProfile to be an explicit nil

### UnsetFlowProfile
`func (o *AFNORResult) UnsetFlowProfile()`

UnsetFlowProfile ensures that no value is present for FlowProfile, not even an explicit nil
### GetFlowType

`func (o *AFNORResult) GetFlowType() string`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *AFNORResult) GetFlowTypeOk() (*string, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *AFNORResult) SetFlowType(v string)`

SetFlowType sets FlowType field to given value.

### HasFlowType

`func (o *AFNORResult) HasFlowType() bool`

HasFlowType returns a boolean if a field has been set.

### SetFlowTypeNil

`func (o *AFNORResult) SetFlowTypeNil(b bool)`

 SetFlowTypeNil sets the value for FlowType to be an explicit nil

### UnsetFlowType
`func (o *AFNORResult) UnsetFlowType()`

UnsetFlowType ensures that no value is present for FlowType, not even an explicit nil
### GetProcessingRule

`func (o *AFNORResult) GetProcessingRule() string`

GetProcessingRule returns the ProcessingRule field if non-nil, zero value otherwise.

### GetProcessingRuleOk

`func (o *AFNORResult) GetProcessingRuleOk() (*string, bool)`

GetProcessingRuleOk returns a tuple with the ProcessingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingRule

`func (o *AFNORResult) SetProcessingRule(v string)`

SetProcessingRule sets ProcessingRule field to given value.

### HasProcessingRule

`func (o *AFNORResult) HasProcessingRule() bool`

HasProcessingRule returns a boolean if a field has been set.

### SetProcessingRuleNil

`func (o *AFNORResult) SetProcessingRuleNil(b bool)`

 SetProcessingRuleNil sets the value for ProcessingRule to be an explicit nil

### UnsetProcessingRule
`func (o *AFNORResult) UnsetProcessingRule()`

UnsetProcessingRule ensures that no value is present for ProcessingRule, not even an explicit nil
### GetProcessingRuleSource

`func (o *AFNORResult) GetProcessingRuleSource() string`

GetProcessingRuleSource returns the ProcessingRuleSource field if non-nil, zero value otherwise.

### GetProcessingRuleSourceOk

`func (o *AFNORResult) GetProcessingRuleSourceOk() (*string, bool)`

GetProcessingRuleSourceOk returns a tuple with the ProcessingRuleSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingRuleSource

`func (o *AFNORResult) SetProcessingRuleSource(v string)`

SetProcessingRuleSource sets ProcessingRuleSource field to given value.

### HasProcessingRuleSource

`func (o *AFNORResult) HasProcessingRuleSource() bool`

HasProcessingRuleSource returns a boolean if a field has been set.

### SetProcessingRuleSourceNil

`func (o *AFNORResult) SetProcessingRuleSourceNil(b bool)`

 SetProcessingRuleSourceNil sets the value for ProcessingRuleSource to be an explicit nil

### UnsetProcessingRuleSource
`func (o *AFNORResult) UnsetProcessingRuleSource()`

UnsetProcessingRuleSource ensures that no value is present for ProcessingRuleSource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



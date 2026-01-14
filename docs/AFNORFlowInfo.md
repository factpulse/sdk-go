# AFNORFlowInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingId** | Pointer to **string** | Unique identifier supporting UUID but not only, for flexibility purpose | [optional] 
**Name** | Pointer to **string** | Name of the file | [optional] 
**ProcessingRule** | Pointer to [**AFNORProcessingRule**](AFNORProcessingRule.md) |  | [optional] 
**FlowSyntax** | [**AFNORFlowSyntax**](AFNORFlowSyntax.md) |  | 
**FlowProfile** | Pointer to [**AFNORFlowProfile**](AFNORFlowProfile.md) |  | [optional] 
**Sha256** | Pointer to **string** |  | [optional] 

## Methods

### NewAFNORFlowInfo

`func NewAFNORFlowInfo(flowSyntax AFNORFlowSyntax, ) *AFNORFlowInfo`

NewAFNORFlowInfo instantiates a new AFNORFlowInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORFlowInfoWithDefaults

`func NewAFNORFlowInfoWithDefaults() *AFNORFlowInfo`

NewAFNORFlowInfoWithDefaults instantiates a new AFNORFlowInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingId

`func (o *AFNORFlowInfo) GetTrackingId() string`

GetTrackingId returns the TrackingId field if non-nil, zero value otherwise.

### GetTrackingIdOk

`func (o *AFNORFlowInfo) GetTrackingIdOk() (*string, bool)`

GetTrackingIdOk returns a tuple with the TrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingId

`func (o *AFNORFlowInfo) SetTrackingId(v string)`

SetTrackingId sets TrackingId field to given value.

### HasTrackingId

`func (o *AFNORFlowInfo) HasTrackingId() bool`

HasTrackingId returns a boolean if a field has been set.

### GetName

`func (o *AFNORFlowInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AFNORFlowInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AFNORFlowInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AFNORFlowInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProcessingRule

`func (o *AFNORFlowInfo) GetProcessingRule() AFNORProcessingRule`

GetProcessingRule returns the ProcessingRule field if non-nil, zero value otherwise.

### GetProcessingRuleOk

`func (o *AFNORFlowInfo) GetProcessingRuleOk() (*AFNORProcessingRule, bool)`

GetProcessingRuleOk returns a tuple with the ProcessingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingRule

`func (o *AFNORFlowInfo) SetProcessingRule(v AFNORProcessingRule)`

SetProcessingRule sets ProcessingRule field to given value.

### HasProcessingRule

`func (o *AFNORFlowInfo) HasProcessingRule() bool`

HasProcessingRule returns a boolean if a field has been set.

### GetFlowSyntax

`func (o *AFNORFlowInfo) GetFlowSyntax() AFNORFlowSyntax`

GetFlowSyntax returns the FlowSyntax field if non-nil, zero value otherwise.

### GetFlowSyntaxOk

`func (o *AFNORFlowInfo) GetFlowSyntaxOk() (*AFNORFlowSyntax, bool)`

GetFlowSyntaxOk returns a tuple with the FlowSyntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSyntax

`func (o *AFNORFlowInfo) SetFlowSyntax(v AFNORFlowSyntax)`

SetFlowSyntax sets FlowSyntax field to given value.


### GetFlowProfile

`func (o *AFNORFlowInfo) GetFlowProfile() AFNORFlowProfile`

GetFlowProfile returns the FlowProfile field if non-nil, zero value otherwise.

### GetFlowProfileOk

`func (o *AFNORFlowInfo) GetFlowProfileOk() (*AFNORFlowProfile, bool)`

GetFlowProfileOk returns a tuple with the FlowProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowProfile

`func (o *AFNORFlowInfo) SetFlowProfile(v AFNORFlowProfile)`

SetFlowProfile sets FlowProfile field to given value.

### HasFlowProfile

`func (o *AFNORFlowInfo) HasFlowProfile() bool`

HasFlowProfile returns a boolean if a field has been set.

### GetSha256

`func (o *AFNORFlowInfo) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *AFNORFlowInfo) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *AFNORFlowInfo) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *AFNORFlowInfo) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



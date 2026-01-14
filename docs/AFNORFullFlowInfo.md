# AFNORFullFlowInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingId** | Pointer to **string** | Unique identifier supporting UUID but not only, for flexibility purpose | [optional] 
**Name** | Pointer to **string** | Name of the file | [optional] 
**ProcessingRule** | Pointer to [**AFNORProcessingRule**](AFNORProcessingRule.md) |  | [optional] 
**FlowSyntax** | [**AFNORFlowSyntax**](AFNORFlowSyntax.md) |  | 
**FlowProfile** | Pointer to [**AFNORFlowProfile**](AFNORFlowProfile.md) |  | [optional] 
**Sha256** | Pointer to **string** |  | [optional] 
**FlowId** | Pointer to **string** | Unique identifier supporting UUID but not only, for flexibility purpose | [optional] 
**SubmittedAt** | Pointer to **time.Time** | The flow submission date and time (the date and time when the flow was created on the system) This property should be used by the API consumer as a time reference to avoid clock synchronization issues  | [optional] 

## Methods

### NewAFNORFullFlowInfo

`func NewAFNORFullFlowInfo(flowSyntax AFNORFlowSyntax, ) *AFNORFullFlowInfo`

NewAFNORFullFlowInfo instantiates a new AFNORFullFlowInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORFullFlowInfoWithDefaults

`func NewAFNORFullFlowInfoWithDefaults() *AFNORFullFlowInfo`

NewAFNORFullFlowInfoWithDefaults instantiates a new AFNORFullFlowInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingId

`func (o *AFNORFullFlowInfo) GetTrackingId() string`

GetTrackingId returns the TrackingId field if non-nil, zero value otherwise.

### GetTrackingIdOk

`func (o *AFNORFullFlowInfo) GetTrackingIdOk() (*string, bool)`

GetTrackingIdOk returns a tuple with the TrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingId

`func (o *AFNORFullFlowInfo) SetTrackingId(v string)`

SetTrackingId sets TrackingId field to given value.

### HasTrackingId

`func (o *AFNORFullFlowInfo) HasTrackingId() bool`

HasTrackingId returns a boolean if a field has been set.

### GetName

`func (o *AFNORFullFlowInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AFNORFullFlowInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AFNORFullFlowInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AFNORFullFlowInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProcessingRule

`func (o *AFNORFullFlowInfo) GetProcessingRule() AFNORProcessingRule`

GetProcessingRule returns the ProcessingRule field if non-nil, zero value otherwise.

### GetProcessingRuleOk

`func (o *AFNORFullFlowInfo) GetProcessingRuleOk() (*AFNORProcessingRule, bool)`

GetProcessingRuleOk returns a tuple with the ProcessingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingRule

`func (o *AFNORFullFlowInfo) SetProcessingRule(v AFNORProcessingRule)`

SetProcessingRule sets ProcessingRule field to given value.

### HasProcessingRule

`func (o *AFNORFullFlowInfo) HasProcessingRule() bool`

HasProcessingRule returns a boolean if a field has been set.

### GetFlowSyntax

`func (o *AFNORFullFlowInfo) GetFlowSyntax() AFNORFlowSyntax`

GetFlowSyntax returns the FlowSyntax field if non-nil, zero value otherwise.

### GetFlowSyntaxOk

`func (o *AFNORFullFlowInfo) GetFlowSyntaxOk() (*AFNORFlowSyntax, bool)`

GetFlowSyntaxOk returns a tuple with the FlowSyntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSyntax

`func (o *AFNORFullFlowInfo) SetFlowSyntax(v AFNORFlowSyntax)`

SetFlowSyntax sets FlowSyntax field to given value.


### GetFlowProfile

`func (o *AFNORFullFlowInfo) GetFlowProfile() AFNORFlowProfile`

GetFlowProfile returns the FlowProfile field if non-nil, zero value otherwise.

### GetFlowProfileOk

`func (o *AFNORFullFlowInfo) GetFlowProfileOk() (*AFNORFlowProfile, bool)`

GetFlowProfileOk returns a tuple with the FlowProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowProfile

`func (o *AFNORFullFlowInfo) SetFlowProfile(v AFNORFlowProfile)`

SetFlowProfile sets FlowProfile field to given value.

### HasFlowProfile

`func (o *AFNORFullFlowInfo) HasFlowProfile() bool`

HasFlowProfile returns a boolean if a field has been set.

### GetSha256

`func (o *AFNORFullFlowInfo) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *AFNORFullFlowInfo) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *AFNORFullFlowInfo) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *AFNORFullFlowInfo) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetFlowId

`func (o *AFNORFullFlowInfo) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *AFNORFullFlowInfo) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *AFNORFullFlowInfo) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *AFNORFullFlowInfo) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetSubmittedAt

`func (o *AFNORFullFlowInfo) GetSubmittedAt() time.Time`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *AFNORFullFlowInfo) GetSubmittedAtOk() (*time.Time, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *AFNORFullFlowInfo) SetSubmittedAt(v time.Time)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *AFNORFullFlowInfo) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



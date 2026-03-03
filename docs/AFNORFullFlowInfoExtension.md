# AFNORFullFlowInfoExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowId** | **string** | Unique identifier supporting UUID but not only, for flexibility purpose | 
**SubmittedAt** | **time.Time** | The flow submission date and time (the date and time when the flow was created on the system) This property should be used by the API consumer as a time reference to avoid clock synchronization issues  | 

## Methods

### NewAFNORFullFlowInfoExtension

`func NewAFNORFullFlowInfoExtension(flowId string, submittedAt time.Time, ) *AFNORFullFlowInfoExtension`

NewAFNORFullFlowInfoExtension instantiates a new AFNORFullFlowInfoExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORFullFlowInfoExtensionWithDefaults

`func NewAFNORFullFlowInfoExtensionWithDefaults() *AFNORFullFlowInfoExtension`

NewAFNORFullFlowInfoExtensionWithDefaults instantiates a new AFNORFullFlowInfoExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowId

`func (o *AFNORFullFlowInfoExtension) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *AFNORFullFlowInfoExtension) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *AFNORFullFlowInfoExtension) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.


### GetSubmittedAt

`func (o *AFNORFullFlowInfoExtension) GetSubmittedAt() time.Time`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *AFNORFullFlowInfoExtension) GetSubmittedAtOk() (*time.Time, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *AFNORFullFlowInfoExtension) SetSubmittedAt(v time.Time)`

SetSubmittedAt sets SubmittedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



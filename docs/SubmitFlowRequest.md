# SubmitFlowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowName** | **string** | Flow name (e.g., &#39;Invoice 2025-001&#39;) | 
**FlowSyntax** | Pointer to [**FlowSyntax**](FlowSyntax.md) | Flow syntax (CII for Factur-X) | [optional] [default to CII]
**FlowProfile** | Pointer to [**NullableFlowProfile**](FlowProfile.md) |  | [optional] 
**TrackingId** | Pointer to **NullableString** |  | [optional] 
**RequestId** | Pointer to **NullableString** |  | [optional] 
**PdpCredentials** | Pointer to [**NullablePDPCredentials**](PDPCredentials.md) |  | [optional] 

## Methods

### NewSubmitFlowRequest

`func NewSubmitFlowRequest(flowName string, ) *SubmitFlowRequest`

NewSubmitFlowRequest instantiates a new SubmitFlowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitFlowRequestWithDefaults

`func NewSubmitFlowRequestWithDefaults() *SubmitFlowRequest`

NewSubmitFlowRequestWithDefaults instantiates a new SubmitFlowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowName

`func (o *SubmitFlowRequest) GetFlowName() string`

GetFlowName returns the FlowName field if non-nil, zero value otherwise.

### GetFlowNameOk

`func (o *SubmitFlowRequest) GetFlowNameOk() (*string, bool)`

GetFlowNameOk returns a tuple with the FlowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowName

`func (o *SubmitFlowRequest) SetFlowName(v string)`

SetFlowName sets FlowName field to given value.


### GetFlowSyntax

`func (o *SubmitFlowRequest) GetFlowSyntax() FlowSyntax`

GetFlowSyntax returns the FlowSyntax field if non-nil, zero value otherwise.

### GetFlowSyntaxOk

`func (o *SubmitFlowRequest) GetFlowSyntaxOk() (*FlowSyntax, bool)`

GetFlowSyntaxOk returns a tuple with the FlowSyntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSyntax

`func (o *SubmitFlowRequest) SetFlowSyntax(v FlowSyntax)`

SetFlowSyntax sets FlowSyntax field to given value.

### HasFlowSyntax

`func (o *SubmitFlowRequest) HasFlowSyntax() bool`

HasFlowSyntax returns a boolean if a field has been set.

### GetFlowProfile

`func (o *SubmitFlowRequest) GetFlowProfile() FlowProfile`

GetFlowProfile returns the FlowProfile field if non-nil, zero value otherwise.

### GetFlowProfileOk

`func (o *SubmitFlowRequest) GetFlowProfileOk() (*FlowProfile, bool)`

GetFlowProfileOk returns a tuple with the FlowProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowProfile

`func (o *SubmitFlowRequest) SetFlowProfile(v FlowProfile)`

SetFlowProfile sets FlowProfile field to given value.

### HasFlowProfile

`func (o *SubmitFlowRequest) HasFlowProfile() bool`

HasFlowProfile returns a boolean if a field has been set.

### SetFlowProfileNil

`func (o *SubmitFlowRequest) SetFlowProfileNil(b bool)`

 SetFlowProfileNil sets the value for FlowProfile to be an explicit nil

### UnsetFlowProfile
`func (o *SubmitFlowRequest) UnsetFlowProfile()`

UnsetFlowProfile ensures that no value is present for FlowProfile, not even an explicit nil
### GetTrackingId

`func (o *SubmitFlowRequest) GetTrackingId() string`

GetTrackingId returns the TrackingId field if non-nil, zero value otherwise.

### GetTrackingIdOk

`func (o *SubmitFlowRequest) GetTrackingIdOk() (*string, bool)`

GetTrackingIdOk returns a tuple with the TrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingId

`func (o *SubmitFlowRequest) SetTrackingId(v string)`

SetTrackingId sets TrackingId field to given value.

### HasTrackingId

`func (o *SubmitFlowRequest) HasTrackingId() bool`

HasTrackingId returns a boolean if a field has been set.

### SetTrackingIdNil

`func (o *SubmitFlowRequest) SetTrackingIdNil(b bool)`

 SetTrackingIdNil sets the value for TrackingId to be an explicit nil

### UnsetTrackingId
`func (o *SubmitFlowRequest) UnsetTrackingId()`

UnsetTrackingId ensures that no value is present for TrackingId, not even an explicit nil
### GetRequestId

`func (o *SubmitFlowRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SubmitFlowRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SubmitFlowRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *SubmitFlowRequest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestIdNil

`func (o *SubmitFlowRequest) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *SubmitFlowRequest) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetPdpCredentials

`func (o *SubmitFlowRequest) GetPdpCredentials() PDPCredentials`

GetPdpCredentials returns the PdpCredentials field if non-nil, zero value otherwise.

### GetPdpCredentialsOk

`func (o *SubmitFlowRequest) GetPdpCredentialsOk() (*PDPCredentials, bool)`

GetPdpCredentialsOk returns a tuple with the PdpCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpCredentials

`func (o *SubmitFlowRequest) SetPdpCredentials(v PDPCredentials)`

SetPdpCredentials sets PdpCredentials field to given value.

### HasPdpCredentials

`func (o *SubmitFlowRequest) HasPdpCredentials() bool`

HasPdpCredentials returns a boolean if a field has been set.

### SetPdpCredentialsNil

`func (o *SubmitFlowRequest) SetPdpCredentialsNil(b bool)`

 SetPdpCredentialsNil sets the value for PdpCredentials to be an explicit nil

### UnsetPdpCredentials
`func (o *SubmitFlowRequest) UnsetPdpCredentials()`

UnsetPdpCredentials ensures that no value is present for PdpCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



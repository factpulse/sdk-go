# SubmitEReportingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**CreateEReportingRequest**](CreateEReportingRequest.md) | E-Reporting data to submit | 
**TrackingId** | Pointer to **NullableString** |  | [optional] 
**PdpFlowServiceUrl** | Pointer to **NullableString** |  | [optional] 
**PdpTokenUrl** | Pointer to **NullableString** |  | [optional] 
**PdpClientId** | Pointer to **NullableString** |  | [optional] 
**PdpClientSecret** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSubmitEReportingRequest

`func NewSubmitEReportingRequest(data CreateEReportingRequest, ) *SubmitEReportingRequest`

NewSubmitEReportingRequest instantiates a new SubmitEReportingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitEReportingRequestWithDefaults

`func NewSubmitEReportingRequestWithDefaults() *SubmitEReportingRequest`

NewSubmitEReportingRequestWithDefaults instantiates a new SubmitEReportingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SubmitEReportingRequest) GetData() CreateEReportingRequest`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubmitEReportingRequest) GetDataOk() (*CreateEReportingRequest, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubmitEReportingRequest) SetData(v CreateEReportingRequest)`

SetData sets Data field to given value.


### GetTrackingId

`func (o *SubmitEReportingRequest) GetTrackingId() string`

GetTrackingId returns the TrackingId field if non-nil, zero value otherwise.

### GetTrackingIdOk

`func (o *SubmitEReportingRequest) GetTrackingIdOk() (*string, bool)`

GetTrackingIdOk returns a tuple with the TrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingId

`func (o *SubmitEReportingRequest) SetTrackingId(v string)`

SetTrackingId sets TrackingId field to given value.

### HasTrackingId

`func (o *SubmitEReportingRequest) HasTrackingId() bool`

HasTrackingId returns a boolean if a field has been set.

### SetTrackingIdNil

`func (o *SubmitEReportingRequest) SetTrackingIdNil(b bool)`

 SetTrackingIdNil sets the value for TrackingId to be an explicit nil

### UnsetTrackingId
`func (o *SubmitEReportingRequest) UnsetTrackingId()`

UnsetTrackingId ensures that no value is present for TrackingId, not even an explicit nil
### GetPdpFlowServiceUrl

`func (o *SubmitEReportingRequest) GetPdpFlowServiceUrl() string`

GetPdpFlowServiceUrl returns the PdpFlowServiceUrl field if non-nil, zero value otherwise.

### GetPdpFlowServiceUrlOk

`func (o *SubmitEReportingRequest) GetPdpFlowServiceUrlOk() (*string, bool)`

GetPdpFlowServiceUrlOk returns a tuple with the PdpFlowServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpFlowServiceUrl

`func (o *SubmitEReportingRequest) SetPdpFlowServiceUrl(v string)`

SetPdpFlowServiceUrl sets PdpFlowServiceUrl field to given value.

### HasPdpFlowServiceUrl

`func (o *SubmitEReportingRequest) HasPdpFlowServiceUrl() bool`

HasPdpFlowServiceUrl returns a boolean if a field has been set.

### SetPdpFlowServiceUrlNil

`func (o *SubmitEReportingRequest) SetPdpFlowServiceUrlNil(b bool)`

 SetPdpFlowServiceUrlNil sets the value for PdpFlowServiceUrl to be an explicit nil

### UnsetPdpFlowServiceUrl
`func (o *SubmitEReportingRequest) UnsetPdpFlowServiceUrl()`

UnsetPdpFlowServiceUrl ensures that no value is present for PdpFlowServiceUrl, not even an explicit nil
### GetPdpTokenUrl

`func (o *SubmitEReportingRequest) GetPdpTokenUrl() string`

GetPdpTokenUrl returns the PdpTokenUrl field if non-nil, zero value otherwise.

### GetPdpTokenUrlOk

`func (o *SubmitEReportingRequest) GetPdpTokenUrlOk() (*string, bool)`

GetPdpTokenUrlOk returns a tuple with the PdpTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpTokenUrl

`func (o *SubmitEReportingRequest) SetPdpTokenUrl(v string)`

SetPdpTokenUrl sets PdpTokenUrl field to given value.

### HasPdpTokenUrl

`func (o *SubmitEReportingRequest) HasPdpTokenUrl() bool`

HasPdpTokenUrl returns a boolean if a field has been set.

### SetPdpTokenUrlNil

`func (o *SubmitEReportingRequest) SetPdpTokenUrlNil(b bool)`

 SetPdpTokenUrlNil sets the value for PdpTokenUrl to be an explicit nil

### UnsetPdpTokenUrl
`func (o *SubmitEReportingRequest) UnsetPdpTokenUrl()`

UnsetPdpTokenUrl ensures that no value is present for PdpTokenUrl, not even an explicit nil
### GetPdpClientId

`func (o *SubmitEReportingRequest) GetPdpClientId() string`

GetPdpClientId returns the PdpClientId field if non-nil, zero value otherwise.

### GetPdpClientIdOk

`func (o *SubmitEReportingRequest) GetPdpClientIdOk() (*string, bool)`

GetPdpClientIdOk returns a tuple with the PdpClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpClientId

`func (o *SubmitEReportingRequest) SetPdpClientId(v string)`

SetPdpClientId sets PdpClientId field to given value.

### HasPdpClientId

`func (o *SubmitEReportingRequest) HasPdpClientId() bool`

HasPdpClientId returns a boolean if a field has been set.

### SetPdpClientIdNil

`func (o *SubmitEReportingRequest) SetPdpClientIdNil(b bool)`

 SetPdpClientIdNil sets the value for PdpClientId to be an explicit nil

### UnsetPdpClientId
`func (o *SubmitEReportingRequest) UnsetPdpClientId()`

UnsetPdpClientId ensures that no value is present for PdpClientId, not even an explicit nil
### GetPdpClientSecret

`func (o *SubmitEReportingRequest) GetPdpClientSecret() string`

GetPdpClientSecret returns the PdpClientSecret field if non-nil, zero value otherwise.

### GetPdpClientSecretOk

`func (o *SubmitEReportingRequest) GetPdpClientSecretOk() (*string, bool)`

GetPdpClientSecretOk returns a tuple with the PdpClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpClientSecret

`func (o *SubmitEReportingRequest) SetPdpClientSecret(v string)`

SetPdpClientSecret sets PdpClientSecret field to given value.

### HasPdpClientSecret

`func (o *SubmitEReportingRequest) HasPdpClientSecret() bool`

HasPdpClientSecret returns a boolean if a field has been set.

### SetPdpClientSecretNil

`func (o *SubmitEReportingRequest) SetPdpClientSecretNil(b bool)`

 SetPdpClientSecretNil sets the value for PdpClientSecret to be an explicit nil

### UnsetPdpClientSecret
`func (o *SubmitEReportingRequest) UnsetPdpClientSecret()`

UnsetPdpClientSecret ensures that no value is present for PdpClientSecret, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



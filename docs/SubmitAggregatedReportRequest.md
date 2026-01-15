# SubmitAggregatedReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**CreateAggregatedReportRequest**](CreateAggregatedReportRequest.md) | Aggregated e-reporting data | 
**TrackingId** | Pointer to **NullableString** |  | [optional] 
**PdpFlowServiceUrl** | Pointer to **NullableString** |  | [optional] 
**PdpTokenUrl** | Pointer to **NullableString** |  | [optional] 
**PdpClientId** | Pointer to **NullableString** |  | [optional] 
**PdpClientSecret** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSubmitAggregatedReportRequest

`func NewSubmitAggregatedReportRequest(data CreateAggregatedReportRequest, ) *SubmitAggregatedReportRequest`

NewSubmitAggregatedReportRequest instantiates a new SubmitAggregatedReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitAggregatedReportRequestWithDefaults

`func NewSubmitAggregatedReportRequestWithDefaults() *SubmitAggregatedReportRequest`

NewSubmitAggregatedReportRequestWithDefaults instantiates a new SubmitAggregatedReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SubmitAggregatedReportRequest) GetData() CreateAggregatedReportRequest`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubmitAggregatedReportRequest) GetDataOk() (*CreateAggregatedReportRequest, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubmitAggregatedReportRequest) SetData(v CreateAggregatedReportRequest)`

SetData sets Data field to given value.


### GetTrackingId

`func (o *SubmitAggregatedReportRequest) GetTrackingId() string`

GetTrackingId returns the TrackingId field if non-nil, zero value otherwise.

### GetTrackingIdOk

`func (o *SubmitAggregatedReportRequest) GetTrackingIdOk() (*string, bool)`

GetTrackingIdOk returns a tuple with the TrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingId

`func (o *SubmitAggregatedReportRequest) SetTrackingId(v string)`

SetTrackingId sets TrackingId field to given value.

### HasTrackingId

`func (o *SubmitAggregatedReportRequest) HasTrackingId() bool`

HasTrackingId returns a boolean if a field has been set.

### SetTrackingIdNil

`func (o *SubmitAggregatedReportRequest) SetTrackingIdNil(b bool)`

 SetTrackingIdNil sets the value for TrackingId to be an explicit nil

### UnsetTrackingId
`func (o *SubmitAggregatedReportRequest) UnsetTrackingId()`

UnsetTrackingId ensures that no value is present for TrackingId, not even an explicit nil
### GetPdpFlowServiceUrl

`func (o *SubmitAggregatedReportRequest) GetPdpFlowServiceUrl() string`

GetPdpFlowServiceUrl returns the PdpFlowServiceUrl field if non-nil, zero value otherwise.

### GetPdpFlowServiceUrlOk

`func (o *SubmitAggregatedReportRequest) GetPdpFlowServiceUrlOk() (*string, bool)`

GetPdpFlowServiceUrlOk returns a tuple with the PdpFlowServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpFlowServiceUrl

`func (o *SubmitAggregatedReportRequest) SetPdpFlowServiceUrl(v string)`

SetPdpFlowServiceUrl sets PdpFlowServiceUrl field to given value.

### HasPdpFlowServiceUrl

`func (o *SubmitAggregatedReportRequest) HasPdpFlowServiceUrl() bool`

HasPdpFlowServiceUrl returns a boolean if a field has been set.

### SetPdpFlowServiceUrlNil

`func (o *SubmitAggregatedReportRequest) SetPdpFlowServiceUrlNil(b bool)`

 SetPdpFlowServiceUrlNil sets the value for PdpFlowServiceUrl to be an explicit nil

### UnsetPdpFlowServiceUrl
`func (o *SubmitAggregatedReportRequest) UnsetPdpFlowServiceUrl()`

UnsetPdpFlowServiceUrl ensures that no value is present for PdpFlowServiceUrl, not even an explicit nil
### GetPdpTokenUrl

`func (o *SubmitAggregatedReportRequest) GetPdpTokenUrl() string`

GetPdpTokenUrl returns the PdpTokenUrl field if non-nil, zero value otherwise.

### GetPdpTokenUrlOk

`func (o *SubmitAggregatedReportRequest) GetPdpTokenUrlOk() (*string, bool)`

GetPdpTokenUrlOk returns a tuple with the PdpTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpTokenUrl

`func (o *SubmitAggregatedReportRequest) SetPdpTokenUrl(v string)`

SetPdpTokenUrl sets PdpTokenUrl field to given value.

### HasPdpTokenUrl

`func (o *SubmitAggregatedReportRequest) HasPdpTokenUrl() bool`

HasPdpTokenUrl returns a boolean if a field has been set.

### SetPdpTokenUrlNil

`func (o *SubmitAggregatedReportRequest) SetPdpTokenUrlNil(b bool)`

 SetPdpTokenUrlNil sets the value for PdpTokenUrl to be an explicit nil

### UnsetPdpTokenUrl
`func (o *SubmitAggregatedReportRequest) UnsetPdpTokenUrl()`

UnsetPdpTokenUrl ensures that no value is present for PdpTokenUrl, not even an explicit nil
### GetPdpClientId

`func (o *SubmitAggregatedReportRequest) GetPdpClientId() string`

GetPdpClientId returns the PdpClientId field if non-nil, zero value otherwise.

### GetPdpClientIdOk

`func (o *SubmitAggregatedReportRequest) GetPdpClientIdOk() (*string, bool)`

GetPdpClientIdOk returns a tuple with the PdpClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpClientId

`func (o *SubmitAggregatedReportRequest) SetPdpClientId(v string)`

SetPdpClientId sets PdpClientId field to given value.

### HasPdpClientId

`func (o *SubmitAggregatedReportRequest) HasPdpClientId() bool`

HasPdpClientId returns a boolean if a field has been set.

### SetPdpClientIdNil

`func (o *SubmitAggregatedReportRequest) SetPdpClientIdNil(b bool)`

 SetPdpClientIdNil sets the value for PdpClientId to be an explicit nil

### UnsetPdpClientId
`func (o *SubmitAggregatedReportRequest) UnsetPdpClientId()`

UnsetPdpClientId ensures that no value is present for PdpClientId, not even an explicit nil
### GetPdpClientSecret

`func (o *SubmitAggregatedReportRequest) GetPdpClientSecret() string`

GetPdpClientSecret returns the PdpClientSecret field if non-nil, zero value otherwise.

### GetPdpClientSecretOk

`func (o *SubmitAggregatedReportRequest) GetPdpClientSecretOk() (*string, bool)`

GetPdpClientSecretOk returns a tuple with the PdpClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpClientSecret

`func (o *SubmitAggregatedReportRequest) SetPdpClientSecret(v string)`

SetPdpClientSecret sets PdpClientSecret field to given value.

### HasPdpClientSecret

`func (o *SubmitAggregatedReportRequest) HasPdpClientSecret() bool`

HasPdpClientSecret returns a boolean if a field has been set.

### SetPdpClientSecretNil

`func (o *SubmitAggregatedReportRequest) SetPdpClientSecretNil(b bool)`

 SetPdpClientSecretNil sets the value for PdpClientSecret to be an explicit nil

### UnsetPdpClientSecret
`func (o *SubmitAggregatedReportRequest) UnsetPdpClientSecret()`

UnsetPdpClientSecret ensures that no value is present for PdpClientSecret, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



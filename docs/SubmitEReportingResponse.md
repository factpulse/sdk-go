# SubmitEReportingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowId** | **string** | Flow identifier from PA/PDP | 
**ReportId** | **string** | Report identifier | 
**TrackingId** | Pointer to **NullableString** |  | [optional] 
**FlowType** | **string** | Flux type | 
**Sha256** | **string** | SHA256 hash of submitted XML | 
**AfnorFlowType** | Pointer to **NullableString** |  | [optional] 
**AfnorResponse** | Pointer to **map[string]interface{}** |  | [optional] 
**Message** | **string** | Status message | 

## Methods

### NewSubmitEReportingResponse

`func NewSubmitEReportingResponse(flowId string, reportId string, flowType string, sha256 string, message string, ) *SubmitEReportingResponse`

NewSubmitEReportingResponse instantiates a new SubmitEReportingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitEReportingResponseWithDefaults

`func NewSubmitEReportingResponseWithDefaults() *SubmitEReportingResponse`

NewSubmitEReportingResponseWithDefaults instantiates a new SubmitEReportingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowId

`func (o *SubmitEReportingResponse) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *SubmitEReportingResponse) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *SubmitEReportingResponse) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.


### GetReportId

`func (o *SubmitEReportingResponse) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *SubmitEReportingResponse) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *SubmitEReportingResponse) SetReportId(v string)`

SetReportId sets ReportId field to given value.


### GetTrackingId

`func (o *SubmitEReportingResponse) GetTrackingId() string`

GetTrackingId returns the TrackingId field if non-nil, zero value otherwise.

### GetTrackingIdOk

`func (o *SubmitEReportingResponse) GetTrackingIdOk() (*string, bool)`

GetTrackingIdOk returns a tuple with the TrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingId

`func (o *SubmitEReportingResponse) SetTrackingId(v string)`

SetTrackingId sets TrackingId field to given value.

### HasTrackingId

`func (o *SubmitEReportingResponse) HasTrackingId() bool`

HasTrackingId returns a boolean if a field has been set.

### SetTrackingIdNil

`func (o *SubmitEReportingResponse) SetTrackingIdNil(b bool)`

 SetTrackingIdNil sets the value for TrackingId to be an explicit nil

### UnsetTrackingId
`func (o *SubmitEReportingResponse) UnsetTrackingId()`

UnsetTrackingId ensures that no value is present for TrackingId, not even an explicit nil
### GetFlowType

`func (o *SubmitEReportingResponse) GetFlowType() string`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *SubmitEReportingResponse) GetFlowTypeOk() (*string, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *SubmitEReportingResponse) SetFlowType(v string)`

SetFlowType sets FlowType field to given value.


### GetSha256

`func (o *SubmitEReportingResponse) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *SubmitEReportingResponse) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *SubmitEReportingResponse) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.


### GetAfnorFlowType

`func (o *SubmitEReportingResponse) GetAfnorFlowType() string`

GetAfnorFlowType returns the AfnorFlowType field if non-nil, zero value otherwise.

### GetAfnorFlowTypeOk

`func (o *SubmitEReportingResponse) GetAfnorFlowTypeOk() (*string, bool)`

GetAfnorFlowTypeOk returns a tuple with the AfnorFlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfnorFlowType

`func (o *SubmitEReportingResponse) SetAfnorFlowType(v string)`

SetAfnorFlowType sets AfnorFlowType field to given value.

### HasAfnorFlowType

`func (o *SubmitEReportingResponse) HasAfnorFlowType() bool`

HasAfnorFlowType returns a boolean if a field has been set.

### SetAfnorFlowTypeNil

`func (o *SubmitEReportingResponse) SetAfnorFlowTypeNil(b bool)`

 SetAfnorFlowTypeNil sets the value for AfnorFlowType to be an explicit nil

### UnsetAfnorFlowType
`func (o *SubmitEReportingResponse) UnsetAfnorFlowType()`

UnsetAfnorFlowType ensures that no value is present for AfnorFlowType, not even an explicit nil
### GetAfnorResponse

`func (o *SubmitEReportingResponse) GetAfnorResponse() map[string]interface{}`

GetAfnorResponse returns the AfnorResponse field if non-nil, zero value otherwise.

### GetAfnorResponseOk

`func (o *SubmitEReportingResponse) GetAfnorResponseOk() (*map[string]interface{}, bool)`

GetAfnorResponseOk returns a tuple with the AfnorResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfnorResponse

`func (o *SubmitEReportingResponse) SetAfnorResponse(v map[string]interface{})`

SetAfnorResponse sets AfnorResponse field to given value.

### HasAfnorResponse

`func (o *SubmitEReportingResponse) HasAfnorResponse() bool`

HasAfnorResponse returns a boolean if a field has been set.

### SetAfnorResponseNil

`func (o *SubmitEReportingResponse) SetAfnorResponseNil(b bool)`

 SetAfnorResponseNil sets the value for AfnorResponse to be an explicit nil

### UnsetAfnorResponse
`func (o *SubmitEReportingResponse) UnsetAfnorResponse()`

UnsetAfnorResponse ensures that no value is present for AfnorResponse, not even an explicit nil
### GetMessage

`func (o *SubmitEReportingResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SubmitEReportingResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SubmitEReportingResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AFNORHealthCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowServiceOk** | **bool** | Flow Service API status | 
**DirectoryServiceOk** | **bool** | Directory Service API status | 
**Message** | **string** | Descriptive status message | 

## Methods

### NewAFNORHealthCheckResponse

`func NewAFNORHealthCheckResponse(flowServiceOk bool, directoryServiceOk bool, message string, ) *AFNORHealthCheckResponse`

NewAFNORHealthCheckResponse instantiates a new AFNORHealthCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORHealthCheckResponseWithDefaults

`func NewAFNORHealthCheckResponseWithDefaults() *AFNORHealthCheckResponse`

NewAFNORHealthCheckResponseWithDefaults instantiates a new AFNORHealthCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowServiceOk

`func (o *AFNORHealthCheckResponse) GetFlowServiceOk() bool`

GetFlowServiceOk returns the FlowServiceOk field if non-nil, zero value otherwise.

### GetFlowServiceOkOk

`func (o *AFNORHealthCheckResponse) GetFlowServiceOkOk() (*bool, bool)`

GetFlowServiceOkOk returns a tuple with the FlowServiceOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowServiceOk

`func (o *AFNORHealthCheckResponse) SetFlowServiceOk(v bool)`

SetFlowServiceOk sets FlowServiceOk field to given value.


### GetDirectoryServiceOk

`func (o *AFNORHealthCheckResponse) GetDirectoryServiceOk() bool`

GetDirectoryServiceOk returns the DirectoryServiceOk field if non-nil, zero value otherwise.

### GetDirectoryServiceOkOk

`func (o *AFNORHealthCheckResponse) GetDirectoryServiceOkOk() (*bool, bool)`

GetDirectoryServiceOkOk returns a tuple with the DirectoryServiceOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryServiceOk

`func (o *AFNORHealthCheckResponse) SetDirectoryServiceOk(v bool)`

SetDirectoryServiceOk sets DirectoryServiceOk field to given value.


### GetMessage

`func (o *AFNORHealthCheckResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AFNORHealthCheckResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AFNORHealthCheckResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ValidateEReportingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | **bool** | Whether the data is valid | 
**ReportId** | **string** | Report identifier | 
**FlowType** | **string** | Flux type | 
**Errors** | Pointer to [**[]ValidationError**](ValidationError.md) | List of validation errors (if any) | [optional] 
**Warnings** | Pointer to [**[]ValidationError**](ValidationError.md) | List of validation warnings (if any) | [optional] 
**Message** | **string** | Status message | 

## Methods

### NewValidateEReportingResponse

`func NewValidateEReportingResponse(valid bool, reportId string, flowType string, message string, ) *ValidateEReportingResponse`

NewValidateEReportingResponse instantiates a new ValidateEReportingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateEReportingResponseWithDefaults

`func NewValidateEReportingResponseWithDefaults() *ValidateEReportingResponse`

NewValidateEReportingResponseWithDefaults instantiates a new ValidateEReportingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *ValidateEReportingResponse) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *ValidateEReportingResponse) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *ValidateEReportingResponse) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetReportId

`func (o *ValidateEReportingResponse) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *ValidateEReportingResponse) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *ValidateEReportingResponse) SetReportId(v string)`

SetReportId sets ReportId field to given value.


### GetFlowType

`func (o *ValidateEReportingResponse) GetFlowType() string`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *ValidateEReportingResponse) GetFlowTypeOk() (*string, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *ValidateEReportingResponse) SetFlowType(v string)`

SetFlowType sets FlowType field to given value.


### GetErrors

`func (o *ValidateEReportingResponse) GetErrors() []ValidationError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ValidateEReportingResponse) GetErrorsOk() (*[]ValidationError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ValidateEReportingResponse) SetErrors(v []ValidationError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ValidateEReportingResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *ValidateEReportingResponse) GetWarnings() []ValidationError`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ValidateEReportingResponse) GetWarningsOk() (*[]ValidationError, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ValidateEReportingResponse) SetWarnings(v []ValidationError)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ValidateEReportingResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetMessage

`func (o *ValidateEReportingResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidateEReportingResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidateEReportingResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GenerateEReportingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportId** | **string** | Report identifier | 
**FlowType** | **string** | Flux type | 
**Xml** | **string** | Generated XML content | 
**XmlSize** | **int32** | XML size in bytes | 
**Message** | **string** | Status message | 

## Methods

### NewGenerateEReportingResponse

`func NewGenerateEReportingResponse(reportId string, flowType string, xml string, xmlSize int32, message string, ) *GenerateEReportingResponse`

NewGenerateEReportingResponse instantiates a new GenerateEReportingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateEReportingResponseWithDefaults

`func NewGenerateEReportingResponseWithDefaults() *GenerateEReportingResponse`

NewGenerateEReportingResponseWithDefaults instantiates a new GenerateEReportingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportId

`func (o *GenerateEReportingResponse) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *GenerateEReportingResponse) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *GenerateEReportingResponse) SetReportId(v string)`

SetReportId sets ReportId field to given value.


### GetFlowType

`func (o *GenerateEReportingResponse) GetFlowType() string`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *GenerateEReportingResponse) GetFlowTypeOk() (*string, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *GenerateEReportingResponse) SetFlowType(v string)`

SetFlowType sets FlowType field to given value.


### GetXml

`func (o *GenerateEReportingResponse) GetXml() string`

GetXml returns the Xml field if non-nil, zero value otherwise.

### GetXmlOk

`func (o *GenerateEReportingResponse) GetXmlOk() (*string, bool)`

GetXmlOk returns a tuple with the Xml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXml

`func (o *GenerateEReportingResponse) SetXml(v string)`

SetXml sets Xml field to given value.


### GetXmlSize

`func (o *GenerateEReportingResponse) GetXmlSize() int32`

GetXmlSize returns the XmlSize field if non-nil, zero value otherwise.

### GetXmlSizeOk

`func (o *GenerateEReportingResponse) GetXmlSizeOk() (*int32, bool)`

GetXmlSizeOk returns a tuple with the XmlSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlSize

`func (o *GenerateEReportingResponse) SetXmlSize(v int32)`

SetXmlSize sets XmlSize field to given value.


### GetMessage

`func (o *GenerateEReportingResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GenerateEReportingResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GenerateEReportingResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



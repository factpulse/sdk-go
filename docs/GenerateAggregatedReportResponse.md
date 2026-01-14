# GenerateAggregatedReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportId** | **string** | Report identifier | 
**TransmissionType** | **string** | Transmission type (IN or RE) | 
**FlowType** | **string** | AFNOR FlowType determined from content | 
**Xml** | **string** | Generated XML content | 
**XmlSize** | **int32** | XML size in bytes | 
**ContentSummary** | **map[string]interface{}** | Summary of content (counts by flux type) | 
**Message** | **string** | Status message | 

## Methods

### NewGenerateAggregatedReportResponse

`func NewGenerateAggregatedReportResponse(reportId string, transmissionType string, flowType string, xml string, xmlSize int32, contentSummary map[string]interface{}, message string, ) *GenerateAggregatedReportResponse`

NewGenerateAggregatedReportResponse instantiates a new GenerateAggregatedReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateAggregatedReportResponseWithDefaults

`func NewGenerateAggregatedReportResponseWithDefaults() *GenerateAggregatedReportResponse`

NewGenerateAggregatedReportResponseWithDefaults instantiates a new GenerateAggregatedReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportId

`func (o *GenerateAggregatedReportResponse) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *GenerateAggregatedReportResponse) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *GenerateAggregatedReportResponse) SetReportId(v string)`

SetReportId sets ReportId field to given value.


### GetTransmissionType

`func (o *GenerateAggregatedReportResponse) GetTransmissionType() string`

GetTransmissionType returns the TransmissionType field if non-nil, zero value otherwise.

### GetTransmissionTypeOk

`func (o *GenerateAggregatedReportResponse) GetTransmissionTypeOk() (*string, bool)`

GetTransmissionTypeOk returns a tuple with the TransmissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmissionType

`func (o *GenerateAggregatedReportResponse) SetTransmissionType(v string)`

SetTransmissionType sets TransmissionType field to given value.


### GetFlowType

`func (o *GenerateAggregatedReportResponse) GetFlowType() string`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *GenerateAggregatedReportResponse) GetFlowTypeOk() (*string, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *GenerateAggregatedReportResponse) SetFlowType(v string)`

SetFlowType sets FlowType field to given value.


### GetXml

`func (o *GenerateAggregatedReportResponse) GetXml() string`

GetXml returns the Xml field if non-nil, zero value otherwise.

### GetXmlOk

`func (o *GenerateAggregatedReportResponse) GetXmlOk() (*string, bool)`

GetXmlOk returns a tuple with the Xml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXml

`func (o *GenerateAggregatedReportResponse) SetXml(v string)`

SetXml sets Xml field to given value.


### GetXmlSize

`func (o *GenerateAggregatedReportResponse) GetXmlSize() int32`

GetXmlSize returns the XmlSize field if non-nil, zero value otherwise.

### GetXmlSizeOk

`func (o *GenerateAggregatedReportResponse) GetXmlSizeOk() (*int32, bool)`

GetXmlSizeOk returns a tuple with the XmlSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlSize

`func (o *GenerateAggregatedReportResponse) SetXmlSize(v int32)`

SetXmlSize sets XmlSize field to given value.


### GetContentSummary

`func (o *GenerateAggregatedReportResponse) GetContentSummary() map[string]interface{}`

GetContentSummary returns the ContentSummary field if non-nil, zero value otherwise.

### GetContentSummaryOk

`func (o *GenerateAggregatedReportResponse) GetContentSummaryOk() (*map[string]interface{}, bool)`

GetContentSummaryOk returns a tuple with the ContentSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSummary

`func (o *GenerateAggregatedReportResponse) SetContentSummary(v map[string]interface{})`

SetContentSummary sets ContentSummary field to given value.


### GetMessage

`func (o *GenerateAggregatedReportResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GenerateAggregatedReportResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GenerateAggregatedReportResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



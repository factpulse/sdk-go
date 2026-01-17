# SubmitCDARXMLRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Xml** | **string** | XML CDAR à soumettre | 
**FlowType** | Pointer to **string** | Type de flux AFNOR | [optional] [default to "CustomerInvoiceLC"]
**Filename** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSubmitCDARXMLRequest

`func NewSubmitCDARXMLRequest(xml string, ) *SubmitCDARXMLRequest`

NewSubmitCDARXMLRequest instantiates a new SubmitCDARXMLRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitCDARXMLRequestWithDefaults

`func NewSubmitCDARXMLRequestWithDefaults() *SubmitCDARXMLRequest`

NewSubmitCDARXMLRequestWithDefaults instantiates a new SubmitCDARXMLRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetXml

`func (o *SubmitCDARXMLRequest) GetXml() string`

GetXml returns the Xml field if non-nil, zero value otherwise.

### GetXmlOk

`func (o *SubmitCDARXMLRequest) GetXmlOk() (*string, bool)`

GetXmlOk returns a tuple with the Xml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXml

`func (o *SubmitCDARXMLRequest) SetXml(v string)`

SetXml sets Xml field to given value.


### GetFlowType

`func (o *SubmitCDARXMLRequest) GetFlowType() string`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *SubmitCDARXMLRequest) GetFlowTypeOk() (*string, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *SubmitCDARXMLRequest) SetFlowType(v string)`

SetFlowType sets FlowType field to given value.

### HasFlowType

`func (o *SubmitCDARXMLRequest) HasFlowType() bool`

HasFlowType returns a boolean if a field has been set.

### GetFilename

`func (o *SubmitCDARXMLRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *SubmitCDARXMLRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *SubmitCDARXMLRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *SubmitCDARXMLRequest) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### SetFilenameNil

`func (o *SubmitCDARXMLRequest) SetFilenameNil(b bool)`

 SetFilenameNil sets the value for Filename to be an explicit nil

### UnsetFilename
`func (o *SubmitCDARXMLRequest) UnsetFilename()`

UnsetFilename ensures that no value is present for Filename, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



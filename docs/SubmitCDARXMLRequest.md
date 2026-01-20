# SubmitCDARXMLRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Xml** | **string** | XML CDAR à soumettre | 
**FlowType** | Pointer to **string** | Type de flux AFNOR | [optional] [default to "CustomerInvoiceLC"]
**Filename** | Pointer to **NullableString** |  | [optional] 
**PdpFlowServiceUrl** | Pointer to **NullableString** |  | [optional] 
**PdpTokenUrl** | Pointer to **NullableString** |  | [optional] 
**PdpClientId** | Pointer to **NullableString** |  | [optional] 
**PdpClientSecret** | Pointer to **NullableString** |  | [optional] 

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
### GetPdpFlowServiceUrl

`func (o *SubmitCDARXMLRequest) GetPdpFlowServiceUrl() string`

GetPdpFlowServiceUrl returns the PdpFlowServiceUrl field if non-nil, zero value otherwise.

### GetPdpFlowServiceUrlOk

`func (o *SubmitCDARXMLRequest) GetPdpFlowServiceUrlOk() (*string, bool)`

GetPdpFlowServiceUrlOk returns a tuple with the PdpFlowServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpFlowServiceUrl

`func (o *SubmitCDARXMLRequest) SetPdpFlowServiceUrl(v string)`

SetPdpFlowServiceUrl sets PdpFlowServiceUrl field to given value.

### HasPdpFlowServiceUrl

`func (o *SubmitCDARXMLRequest) HasPdpFlowServiceUrl() bool`

HasPdpFlowServiceUrl returns a boolean if a field has been set.

### SetPdpFlowServiceUrlNil

`func (o *SubmitCDARXMLRequest) SetPdpFlowServiceUrlNil(b bool)`

 SetPdpFlowServiceUrlNil sets the value for PdpFlowServiceUrl to be an explicit nil

### UnsetPdpFlowServiceUrl
`func (o *SubmitCDARXMLRequest) UnsetPdpFlowServiceUrl()`

UnsetPdpFlowServiceUrl ensures that no value is present for PdpFlowServiceUrl, not even an explicit nil
### GetPdpTokenUrl

`func (o *SubmitCDARXMLRequest) GetPdpTokenUrl() string`

GetPdpTokenUrl returns the PdpTokenUrl field if non-nil, zero value otherwise.

### GetPdpTokenUrlOk

`func (o *SubmitCDARXMLRequest) GetPdpTokenUrlOk() (*string, bool)`

GetPdpTokenUrlOk returns a tuple with the PdpTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpTokenUrl

`func (o *SubmitCDARXMLRequest) SetPdpTokenUrl(v string)`

SetPdpTokenUrl sets PdpTokenUrl field to given value.

### HasPdpTokenUrl

`func (o *SubmitCDARXMLRequest) HasPdpTokenUrl() bool`

HasPdpTokenUrl returns a boolean if a field has been set.

### SetPdpTokenUrlNil

`func (o *SubmitCDARXMLRequest) SetPdpTokenUrlNil(b bool)`

 SetPdpTokenUrlNil sets the value for PdpTokenUrl to be an explicit nil

### UnsetPdpTokenUrl
`func (o *SubmitCDARXMLRequest) UnsetPdpTokenUrl()`

UnsetPdpTokenUrl ensures that no value is present for PdpTokenUrl, not even an explicit nil
### GetPdpClientId

`func (o *SubmitCDARXMLRequest) GetPdpClientId() string`

GetPdpClientId returns the PdpClientId field if non-nil, zero value otherwise.

### GetPdpClientIdOk

`func (o *SubmitCDARXMLRequest) GetPdpClientIdOk() (*string, bool)`

GetPdpClientIdOk returns a tuple with the PdpClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpClientId

`func (o *SubmitCDARXMLRequest) SetPdpClientId(v string)`

SetPdpClientId sets PdpClientId field to given value.

### HasPdpClientId

`func (o *SubmitCDARXMLRequest) HasPdpClientId() bool`

HasPdpClientId returns a boolean if a field has been set.

### SetPdpClientIdNil

`func (o *SubmitCDARXMLRequest) SetPdpClientIdNil(b bool)`

 SetPdpClientIdNil sets the value for PdpClientId to be an explicit nil

### UnsetPdpClientId
`func (o *SubmitCDARXMLRequest) UnsetPdpClientId()`

UnsetPdpClientId ensures that no value is present for PdpClientId, not even an explicit nil
### GetPdpClientSecret

`func (o *SubmitCDARXMLRequest) GetPdpClientSecret() string`

GetPdpClientSecret returns the PdpClientSecret field if non-nil, zero value otherwise.

### GetPdpClientSecretOk

`func (o *SubmitCDARXMLRequest) GetPdpClientSecretOk() (*string, bool)`

GetPdpClientSecretOk returns a tuple with the PdpClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpClientSecret

`func (o *SubmitCDARXMLRequest) SetPdpClientSecret(v string)`

SetPdpClientSecret sets PdpClientSecret field to given value.

### HasPdpClientSecret

`func (o *SubmitCDARXMLRequest) HasPdpClientSecret() bool`

HasPdpClientSecret returns a boolean if a field has been set.

### SetPdpClientSecretNil

`func (o *SubmitCDARXMLRequest) SetPdpClientSecretNil(b bool)`

 SetPdpClientSecretNil sets the value for PdpClientSecret to be an explicit nil

### UnsetPdpClientSecret
`func (o *SubmitCDARXMLRequest) UnsetPdpClientSecret()`

UnsetPdpClientSecret ensures that no value is present for PdpClientSecret, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



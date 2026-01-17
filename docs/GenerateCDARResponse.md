# GenerateCDARResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | **string** |  | 
**Xml** | **string** | XML CDAR généré | 
**XmlSize** | **int32** | Taille du XML en octets | 
**Sha256** | **string** | Hash SHA256 du XML | 
**Message** | **string** | Message de succès | 

## Methods

### NewGenerateCDARResponse

`func NewGenerateCDARResponse(documentId string, xml string, xmlSize int32, sha256 string, message string, ) *GenerateCDARResponse`

NewGenerateCDARResponse instantiates a new GenerateCDARResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateCDARResponseWithDefaults

`func NewGenerateCDARResponseWithDefaults() *GenerateCDARResponse`

NewGenerateCDARResponseWithDefaults instantiates a new GenerateCDARResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *GenerateCDARResponse) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *GenerateCDARResponse) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *GenerateCDARResponse) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.


### GetXml

`func (o *GenerateCDARResponse) GetXml() string`

GetXml returns the Xml field if non-nil, zero value otherwise.

### GetXmlOk

`func (o *GenerateCDARResponse) GetXmlOk() (*string, bool)`

GetXmlOk returns a tuple with the Xml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXml

`func (o *GenerateCDARResponse) SetXml(v string)`

SetXml sets Xml field to given value.


### GetXmlSize

`func (o *GenerateCDARResponse) GetXmlSize() int32`

GetXmlSize returns the XmlSize field if non-nil, zero value otherwise.

### GetXmlSizeOk

`func (o *GenerateCDARResponse) GetXmlSizeOk() (*int32, bool)`

GetXmlSizeOk returns a tuple with the XmlSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlSize

`func (o *GenerateCDARResponse) SetXmlSize(v int32)`

SetXmlSize sets XmlSize field to given value.


### GetSha256

`func (o *GenerateCDARResponse) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *GenerateCDARResponse) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *GenerateCDARResponse) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.


### GetMessage

`func (o *GenerateCDARResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GenerateCDARResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GenerateCDARResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



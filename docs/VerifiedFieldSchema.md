# VerifiedFieldSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessTerm** | **string** | EN16931 Business Term (e.g., BT-1) | 
**Label** | **string** | Field label (e.g., Invoice Number) | 
**PdfValue** | Pointer to **NullableString** |  | [optional] 
**XmlValue** | Pointer to **NullableString** |  | [optional] 
**Status** | [**FieldStatus**](FieldStatus.md) | Compliance status | 
**Message** | Pointer to **NullableString** |  | [optional] 
**Confidence** | Pointer to **float32** | Confidence score (0-1) | [optional] [default to 1.0]
**Source** | Pointer to **string** | Extraction source | [optional] [default to "native_pdf"]
**Bbox** | Pointer to [**NullableBoundingBoxSchema**](BoundingBoxSchema.md) |  | [optional] 

## Methods

### NewVerifiedFieldSchema

`func NewVerifiedFieldSchema(businessTerm string, label string, status FieldStatus, ) *VerifiedFieldSchema`

NewVerifiedFieldSchema instantiates a new VerifiedFieldSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifiedFieldSchemaWithDefaults

`func NewVerifiedFieldSchemaWithDefaults() *VerifiedFieldSchema`

NewVerifiedFieldSchemaWithDefaults instantiates a new VerifiedFieldSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessTerm

`func (o *VerifiedFieldSchema) GetBusinessTerm() string`

GetBusinessTerm returns the BusinessTerm field if non-nil, zero value otherwise.

### GetBusinessTermOk

`func (o *VerifiedFieldSchema) GetBusinessTermOk() (*string, bool)`

GetBusinessTermOk returns a tuple with the BusinessTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessTerm

`func (o *VerifiedFieldSchema) SetBusinessTerm(v string)`

SetBusinessTerm sets BusinessTerm field to given value.


### GetLabel

`func (o *VerifiedFieldSchema) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VerifiedFieldSchema) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VerifiedFieldSchema) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetPdfValue

`func (o *VerifiedFieldSchema) GetPdfValue() string`

GetPdfValue returns the PdfValue field if non-nil, zero value otherwise.

### GetPdfValueOk

`func (o *VerifiedFieldSchema) GetPdfValueOk() (*string, bool)`

GetPdfValueOk returns a tuple with the PdfValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfValue

`func (o *VerifiedFieldSchema) SetPdfValue(v string)`

SetPdfValue sets PdfValue field to given value.

### HasPdfValue

`func (o *VerifiedFieldSchema) HasPdfValue() bool`

HasPdfValue returns a boolean if a field has been set.

### SetPdfValueNil

`func (o *VerifiedFieldSchema) SetPdfValueNil(b bool)`

 SetPdfValueNil sets the value for PdfValue to be an explicit nil

### UnsetPdfValue
`func (o *VerifiedFieldSchema) UnsetPdfValue()`

UnsetPdfValue ensures that no value is present for PdfValue, not even an explicit nil
### GetXmlValue

`func (o *VerifiedFieldSchema) GetXmlValue() string`

GetXmlValue returns the XmlValue field if non-nil, zero value otherwise.

### GetXmlValueOk

`func (o *VerifiedFieldSchema) GetXmlValueOk() (*string, bool)`

GetXmlValueOk returns a tuple with the XmlValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlValue

`func (o *VerifiedFieldSchema) SetXmlValue(v string)`

SetXmlValue sets XmlValue field to given value.

### HasXmlValue

`func (o *VerifiedFieldSchema) HasXmlValue() bool`

HasXmlValue returns a boolean if a field has been set.

### SetXmlValueNil

`func (o *VerifiedFieldSchema) SetXmlValueNil(b bool)`

 SetXmlValueNil sets the value for XmlValue to be an explicit nil

### UnsetXmlValue
`func (o *VerifiedFieldSchema) UnsetXmlValue()`

UnsetXmlValue ensures that no value is present for XmlValue, not even an explicit nil
### GetStatus

`func (o *VerifiedFieldSchema) GetStatus() FieldStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VerifiedFieldSchema) GetStatusOk() (*FieldStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VerifiedFieldSchema) SetStatus(v FieldStatus)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *VerifiedFieldSchema) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VerifiedFieldSchema) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VerifiedFieldSchema) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *VerifiedFieldSchema) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *VerifiedFieldSchema) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *VerifiedFieldSchema) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetConfidence

`func (o *VerifiedFieldSchema) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *VerifiedFieldSchema) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *VerifiedFieldSchema) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *VerifiedFieldSchema) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetSource

`func (o *VerifiedFieldSchema) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *VerifiedFieldSchema) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *VerifiedFieldSchema) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *VerifiedFieldSchema) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetBbox

`func (o *VerifiedFieldSchema) GetBbox() BoundingBoxSchema`

GetBbox returns the Bbox field if non-nil, zero value otherwise.

### GetBboxOk

`func (o *VerifiedFieldSchema) GetBboxOk() (*BoundingBoxSchema, bool)`

GetBboxOk returns a tuple with the Bbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBbox

`func (o *VerifiedFieldSchema) SetBbox(v BoundingBoxSchema)`

SetBbox sets Bbox field to given value.

### HasBbox

`func (o *VerifiedFieldSchema) HasBbox() bool`

HasBbox returns a boolean if a field has been set.

### SetBboxNil

`func (o *VerifiedFieldSchema) SetBboxNil(b bool)`

 SetBboxNil sets the value for Bbox to be an explicit nil

### UnsetBbox
`func (o *VerifiedFieldSchema) UnsetBbox()`

UnsetBbox ensures that no value is present for Bbox, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



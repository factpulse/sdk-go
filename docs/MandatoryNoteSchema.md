# MandatoryNoteSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectCode** | **string** | Subject code (PMT, PMD, AAB) | 
**Label** | **string** | Label (e.g., Recovery indemnity) | 
**PdfValue** | Pointer to **NullableString** |  | [optional] 
**XmlValue** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**FieldStatus**](FieldStatus.md) | Compliance status (COMPLIANT if XML found in PDF) | [optional] [default to NOT_VERIFIED]
**Message** | Pointer to **NullableString** |  | [optional] 
**Bbox** | Pointer to [**NullableBoundingBoxSchema**](BoundingBoxSchema.md) |  | [optional] 

## Methods

### NewMandatoryNoteSchema

`func NewMandatoryNoteSchema(subjectCode string, label string, ) *MandatoryNoteSchema`

NewMandatoryNoteSchema instantiates a new MandatoryNoteSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMandatoryNoteSchemaWithDefaults

`func NewMandatoryNoteSchemaWithDefaults() *MandatoryNoteSchema`

NewMandatoryNoteSchemaWithDefaults instantiates a new MandatoryNoteSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectCode

`func (o *MandatoryNoteSchema) GetSubjectCode() string`

GetSubjectCode returns the SubjectCode field if non-nil, zero value otherwise.

### GetSubjectCodeOk

`func (o *MandatoryNoteSchema) GetSubjectCodeOk() (*string, bool)`

GetSubjectCodeOk returns a tuple with the SubjectCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCode

`func (o *MandatoryNoteSchema) SetSubjectCode(v string)`

SetSubjectCode sets SubjectCode field to given value.


### GetLabel

`func (o *MandatoryNoteSchema) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MandatoryNoteSchema) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MandatoryNoteSchema) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetPdfValue

`func (o *MandatoryNoteSchema) GetPdfValue() string`

GetPdfValue returns the PdfValue field if non-nil, zero value otherwise.

### GetPdfValueOk

`func (o *MandatoryNoteSchema) GetPdfValueOk() (*string, bool)`

GetPdfValueOk returns a tuple with the PdfValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfValue

`func (o *MandatoryNoteSchema) SetPdfValue(v string)`

SetPdfValue sets PdfValue field to given value.

### HasPdfValue

`func (o *MandatoryNoteSchema) HasPdfValue() bool`

HasPdfValue returns a boolean if a field has been set.

### SetPdfValueNil

`func (o *MandatoryNoteSchema) SetPdfValueNil(b bool)`

 SetPdfValueNil sets the value for PdfValue to be an explicit nil

### UnsetPdfValue
`func (o *MandatoryNoteSchema) UnsetPdfValue()`

UnsetPdfValue ensures that no value is present for PdfValue, not even an explicit nil
### GetXmlValue

`func (o *MandatoryNoteSchema) GetXmlValue() string`

GetXmlValue returns the XmlValue field if non-nil, zero value otherwise.

### GetXmlValueOk

`func (o *MandatoryNoteSchema) GetXmlValueOk() (*string, bool)`

GetXmlValueOk returns a tuple with the XmlValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlValue

`func (o *MandatoryNoteSchema) SetXmlValue(v string)`

SetXmlValue sets XmlValue field to given value.

### HasXmlValue

`func (o *MandatoryNoteSchema) HasXmlValue() bool`

HasXmlValue returns a boolean if a field has been set.

### SetXmlValueNil

`func (o *MandatoryNoteSchema) SetXmlValueNil(b bool)`

 SetXmlValueNil sets the value for XmlValue to be an explicit nil

### UnsetXmlValue
`func (o *MandatoryNoteSchema) UnsetXmlValue()`

UnsetXmlValue ensures that no value is present for XmlValue, not even an explicit nil
### GetStatus

`func (o *MandatoryNoteSchema) GetStatus() FieldStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MandatoryNoteSchema) GetStatusOk() (*FieldStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MandatoryNoteSchema) SetStatus(v FieldStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MandatoryNoteSchema) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *MandatoryNoteSchema) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MandatoryNoteSchema) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MandatoryNoteSchema) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MandatoryNoteSchema) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *MandatoryNoteSchema) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *MandatoryNoteSchema) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetBbox

`func (o *MandatoryNoteSchema) GetBbox() BoundingBoxSchema`

GetBbox returns the Bbox field if non-nil, zero value otherwise.

### GetBboxOk

`func (o *MandatoryNoteSchema) GetBboxOk() (*BoundingBoxSchema, bool)`

GetBboxOk returns a tuple with the Bbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBbox

`func (o *MandatoryNoteSchema) SetBbox(v BoundingBoxSchema)`

SetBbox sets Bbox field to given value.

### HasBbox

`func (o *MandatoryNoteSchema) HasBbox() bool`

HasBbox returns a boolean if a field has been set.

### SetBboxNil

`func (o *MandatoryNoteSchema) SetBboxNil(b bool)`

 SetBboxNil sets the value for Bbox to be an explicit nil

### UnsetBbox
`func (o *MandatoryNoteSchema) UnsetBbox()`

UnsetBbox ensures that no value is present for Bbox, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



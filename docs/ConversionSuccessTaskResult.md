# ConversionSuccessTaskResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "SUCCESS"]
**ConversionId** | **string** |  | 
**DocumentTypeCode** | **int32** |  | 
**Profile** | **string** |  | 
**Extraction** | [**ConversionExtractionInfo**](ConversionExtractionInfo.md) |  | 
**ProcessingTimeMs** | **int32** |  | 
**PdfRegenerated** | Pointer to **bool** |  | [optional] [default to false]
**PdfRegeneratedReason** | Pointer to **NullableString** |  | [optional] 
**ContentB64** | Pointer to **NullableString** |  | [optional] 
**XmlContent** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewConversionSuccessTaskResult

`func NewConversionSuccessTaskResult(conversionId string, documentTypeCode int32, profile string, extraction ConversionExtractionInfo, processingTimeMs int32, ) *ConversionSuccessTaskResult`

NewConversionSuccessTaskResult instantiates a new ConversionSuccessTaskResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversionSuccessTaskResultWithDefaults

`func NewConversionSuccessTaskResultWithDefaults() *ConversionSuccessTaskResult`

NewConversionSuccessTaskResultWithDefaults instantiates a new ConversionSuccessTaskResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ConversionSuccessTaskResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConversionSuccessTaskResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConversionSuccessTaskResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConversionSuccessTaskResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConversionId

`func (o *ConversionSuccessTaskResult) GetConversionId() string`

GetConversionId returns the ConversionId field if non-nil, zero value otherwise.

### GetConversionIdOk

`func (o *ConversionSuccessTaskResult) GetConversionIdOk() (*string, bool)`

GetConversionIdOk returns a tuple with the ConversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionId

`func (o *ConversionSuccessTaskResult) SetConversionId(v string)`

SetConversionId sets ConversionId field to given value.


### GetDocumentTypeCode

`func (o *ConversionSuccessTaskResult) GetDocumentTypeCode() int32`

GetDocumentTypeCode returns the DocumentTypeCode field if non-nil, zero value otherwise.

### GetDocumentTypeCodeOk

`func (o *ConversionSuccessTaskResult) GetDocumentTypeCodeOk() (*int32, bool)`

GetDocumentTypeCodeOk returns a tuple with the DocumentTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypeCode

`func (o *ConversionSuccessTaskResult) SetDocumentTypeCode(v int32)`

SetDocumentTypeCode sets DocumentTypeCode field to given value.


### GetProfile

`func (o *ConversionSuccessTaskResult) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ConversionSuccessTaskResult) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ConversionSuccessTaskResult) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetExtraction

`func (o *ConversionSuccessTaskResult) GetExtraction() ConversionExtractionInfo`

GetExtraction returns the Extraction field if non-nil, zero value otherwise.

### GetExtractionOk

`func (o *ConversionSuccessTaskResult) GetExtractionOk() (*ConversionExtractionInfo, bool)`

GetExtractionOk returns a tuple with the Extraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraction

`func (o *ConversionSuccessTaskResult) SetExtraction(v ConversionExtractionInfo)`

SetExtraction sets Extraction field to given value.


### GetProcessingTimeMs

`func (o *ConversionSuccessTaskResult) GetProcessingTimeMs() int32`

GetProcessingTimeMs returns the ProcessingTimeMs field if non-nil, zero value otherwise.

### GetProcessingTimeMsOk

`func (o *ConversionSuccessTaskResult) GetProcessingTimeMsOk() (*int32, bool)`

GetProcessingTimeMsOk returns a tuple with the ProcessingTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeMs

`func (o *ConversionSuccessTaskResult) SetProcessingTimeMs(v int32)`

SetProcessingTimeMs sets ProcessingTimeMs field to given value.


### GetPdfRegenerated

`func (o *ConversionSuccessTaskResult) GetPdfRegenerated() bool`

GetPdfRegenerated returns the PdfRegenerated field if non-nil, zero value otherwise.

### GetPdfRegeneratedOk

`func (o *ConversionSuccessTaskResult) GetPdfRegeneratedOk() (*bool, bool)`

GetPdfRegeneratedOk returns a tuple with the PdfRegenerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfRegenerated

`func (o *ConversionSuccessTaskResult) SetPdfRegenerated(v bool)`

SetPdfRegenerated sets PdfRegenerated field to given value.

### HasPdfRegenerated

`func (o *ConversionSuccessTaskResult) HasPdfRegenerated() bool`

HasPdfRegenerated returns a boolean if a field has been set.

### GetPdfRegeneratedReason

`func (o *ConversionSuccessTaskResult) GetPdfRegeneratedReason() string`

GetPdfRegeneratedReason returns the PdfRegeneratedReason field if non-nil, zero value otherwise.

### GetPdfRegeneratedReasonOk

`func (o *ConversionSuccessTaskResult) GetPdfRegeneratedReasonOk() (*string, bool)`

GetPdfRegeneratedReasonOk returns a tuple with the PdfRegeneratedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfRegeneratedReason

`func (o *ConversionSuccessTaskResult) SetPdfRegeneratedReason(v string)`

SetPdfRegeneratedReason sets PdfRegeneratedReason field to given value.

### HasPdfRegeneratedReason

`func (o *ConversionSuccessTaskResult) HasPdfRegeneratedReason() bool`

HasPdfRegeneratedReason returns a boolean if a field has been set.

### SetPdfRegeneratedReasonNil

`func (o *ConversionSuccessTaskResult) SetPdfRegeneratedReasonNil(b bool)`

 SetPdfRegeneratedReasonNil sets the value for PdfRegeneratedReason to be an explicit nil

### UnsetPdfRegeneratedReason
`func (o *ConversionSuccessTaskResult) UnsetPdfRegeneratedReason()`

UnsetPdfRegeneratedReason ensures that no value is present for PdfRegeneratedReason, not even an explicit nil
### GetContentB64

`func (o *ConversionSuccessTaskResult) GetContentB64() string`

GetContentB64 returns the ContentB64 field if non-nil, zero value otherwise.

### GetContentB64Ok

`func (o *ConversionSuccessTaskResult) GetContentB64Ok() (*string, bool)`

GetContentB64Ok returns a tuple with the ContentB64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentB64

`func (o *ConversionSuccessTaskResult) SetContentB64(v string)`

SetContentB64 sets ContentB64 field to given value.

### HasContentB64

`func (o *ConversionSuccessTaskResult) HasContentB64() bool`

HasContentB64 returns a boolean if a field has been set.

### SetContentB64Nil

`func (o *ConversionSuccessTaskResult) SetContentB64Nil(b bool)`

 SetContentB64Nil sets the value for ContentB64 to be an explicit nil

### UnsetContentB64
`func (o *ConversionSuccessTaskResult) UnsetContentB64()`

UnsetContentB64 ensures that no value is present for ContentB64, not even an explicit nil
### GetXmlContent

`func (o *ConversionSuccessTaskResult) GetXmlContent() string`

GetXmlContent returns the XmlContent field if non-nil, zero value otherwise.

### GetXmlContentOk

`func (o *ConversionSuccessTaskResult) GetXmlContentOk() (*string, bool)`

GetXmlContentOk returns a tuple with the XmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlContent

`func (o *ConversionSuccessTaskResult) SetXmlContent(v string)`

SetXmlContent sets XmlContent field to given value.

### HasXmlContent

`func (o *ConversionSuccessTaskResult) HasXmlContent() bool`

HasXmlContent returns a boolean if a field has been set.

### SetXmlContentNil

`func (o *ConversionSuccessTaskResult) SetXmlContentNil(b bool)`

 SetXmlContentNil sets the value for XmlContent to be an explicit nil

### UnsetXmlContent
`func (o *ConversionSuccessTaskResult) UnsetXmlContent()`

UnsetXmlContent ensures that no value is present for XmlContent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



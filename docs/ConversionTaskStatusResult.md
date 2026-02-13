# ConversionTaskStatusResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "ERROR"]
**ConversionId** | **string** |  | 
**DocumentTypeCode** | **int32** |  | 
**Profile** | **string** |  | 
**Extraction** | [**ConversionExtractionInfo**](ConversionExtractionInfo.md) |  | 
**ProcessingTimeMs** | **int32** |  | 
**PdfRegenerated** | Pointer to **bool** |  | [optional] [default to false]
**PdfRegeneratedReason** | Pointer to **string** |  | [optional] 
**ContentB64** | Pointer to **string** |  | [optional] 
**XmlContent** | Pointer to **string** |  | [optional] 
**Message** | **string** |  | 
**MissingFields** | **[]string** |  | 
**ExtractedData** | **map[string]interface{}** |  | 
**ConfidenceScore** | **float32** |  | 
**ValidationErrors** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CorrectionAttempted** | Pointer to **bool** |  | [optional] [default to false]
**ErrorCode** | **string** |  | 
**ErrorMessage** | **string** |  | 
**Details** | Pointer to [**[]AFNORErrorDetail**](AFNORErrorDetail.md) |  | [optional] 
**Traceback** | Pointer to **string** |  | [optional] 

## Methods

### NewConversionTaskStatusResult

`func NewConversionTaskStatusResult(conversionId string, documentTypeCode int32, profile string, extraction ConversionExtractionInfo, processingTimeMs int32, message string, missingFields []string, extractedData map[string]interface{}, confidenceScore float32, errorCode string, errorMessage string, ) *ConversionTaskStatusResult`

NewConversionTaskStatusResult instantiates a new ConversionTaskStatusResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversionTaskStatusResultWithDefaults

`func NewConversionTaskStatusResultWithDefaults() *ConversionTaskStatusResult`

NewConversionTaskStatusResultWithDefaults instantiates a new ConversionTaskStatusResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ConversionTaskStatusResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConversionTaskStatusResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConversionTaskStatusResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConversionTaskStatusResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConversionId

`func (o *ConversionTaskStatusResult) GetConversionId() string`

GetConversionId returns the ConversionId field if non-nil, zero value otherwise.

### GetConversionIdOk

`func (o *ConversionTaskStatusResult) GetConversionIdOk() (*string, bool)`

GetConversionIdOk returns a tuple with the ConversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionId

`func (o *ConversionTaskStatusResult) SetConversionId(v string)`

SetConversionId sets ConversionId field to given value.


### GetDocumentTypeCode

`func (o *ConversionTaskStatusResult) GetDocumentTypeCode() int32`

GetDocumentTypeCode returns the DocumentTypeCode field if non-nil, zero value otherwise.

### GetDocumentTypeCodeOk

`func (o *ConversionTaskStatusResult) GetDocumentTypeCodeOk() (*int32, bool)`

GetDocumentTypeCodeOk returns a tuple with the DocumentTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypeCode

`func (o *ConversionTaskStatusResult) SetDocumentTypeCode(v int32)`

SetDocumentTypeCode sets DocumentTypeCode field to given value.


### GetProfile

`func (o *ConversionTaskStatusResult) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ConversionTaskStatusResult) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ConversionTaskStatusResult) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetExtraction

`func (o *ConversionTaskStatusResult) GetExtraction() ConversionExtractionInfo`

GetExtraction returns the Extraction field if non-nil, zero value otherwise.

### GetExtractionOk

`func (o *ConversionTaskStatusResult) GetExtractionOk() (*ConversionExtractionInfo, bool)`

GetExtractionOk returns a tuple with the Extraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraction

`func (o *ConversionTaskStatusResult) SetExtraction(v ConversionExtractionInfo)`

SetExtraction sets Extraction field to given value.


### GetProcessingTimeMs

`func (o *ConversionTaskStatusResult) GetProcessingTimeMs() int32`

GetProcessingTimeMs returns the ProcessingTimeMs field if non-nil, zero value otherwise.

### GetProcessingTimeMsOk

`func (o *ConversionTaskStatusResult) GetProcessingTimeMsOk() (*int32, bool)`

GetProcessingTimeMsOk returns a tuple with the ProcessingTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeMs

`func (o *ConversionTaskStatusResult) SetProcessingTimeMs(v int32)`

SetProcessingTimeMs sets ProcessingTimeMs field to given value.


### GetPdfRegenerated

`func (o *ConversionTaskStatusResult) GetPdfRegenerated() bool`

GetPdfRegenerated returns the PdfRegenerated field if non-nil, zero value otherwise.

### GetPdfRegeneratedOk

`func (o *ConversionTaskStatusResult) GetPdfRegeneratedOk() (*bool, bool)`

GetPdfRegeneratedOk returns a tuple with the PdfRegenerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfRegenerated

`func (o *ConversionTaskStatusResult) SetPdfRegenerated(v bool)`

SetPdfRegenerated sets PdfRegenerated field to given value.

### HasPdfRegenerated

`func (o *ConversionTaskStatusResult) HasPdfRegenerated() bool`

HasPdfRegenerated returns a boolean if a field has been set.

### GetPdfRegeneratedReason

`func (o *ConversionTaskStatusResult) GetPdfRegeneratedReason() string`

GetPdfRegeneratedReason returns the PdfRegeneratedReason field if non-nil, zero value otherwise.

### GetPdfRegeneratedReasonOk

`func (o *ConversionTaskStatusResult) GetPdfRegeneratedReasonOk() (*string, bool)`

GetPdfRegeneratedReasonOk returns a tuple with the PdfRegeneratedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfRegeneratedReason

`func (o *ConversionTaskStatusResult) SetPdfRegeneratedReason(v string)`

SetPdfRegeneratedReason sets PdfRegeneratedReason field to given value.

### HasPdfRegeneratedReason

`func (o *ConversionTaskStatusResult) HasPdfRegeneratedReason() bool`

HasPdfRegeneratedReason returns a boolean if a field has been set.

### GetContentB64

`func (o *ConversionTaskStatusResult) GetContentB64() string`

GetContentB64 returns the ContentB64 field if non-nil, zero value otherwise.

### GetContentB64Ok

`func (o *ConversionTaskStatusResult) GetContentB64Ok() (*string, bool)`

GetContentB64Ok returns a tuple with the ContentB64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentB64

`func (o *ConversionTaskStatusResult) SetContentB64(v string)`

SetContentB64 sets ContentB64 field to given value.

### HasContentB64

`func (o *ConversionTaskStatusResult) HasContentB64() bool`

HasContentB64 returns a boolean if a field has been set.

### GetXmlContent

`func (o *ConversionTaskStatusResult) GetXmlContent() string`

GetXmlContent returns the XmlContent field if non-nil, zero value otherwise.

### GetXmlContentOk

`func (o *ConversionTaskStatusResult) GetXmlContentOk() (*string, bool)`

GetXmlContentOk returns a tuple with the XmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlContent

`func (o *ConversionTaskStatusResult) SetXmlContent(v string)`

SetXmlContent sets XmlContent field to given value.

### HasXmlContent

`func (o *ConversionTaskStatusResult) HasXmlContent() bool`

HasXmlContent returns a boolean if a field has been set.

### GetMessage

`func (o *ConversionTaskStatusResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConversionTaskStatusResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConversionTaskStatusResult) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMissingFields

`func (o *ConversionTaskStatusResult) GetMissingFields() []string`

GetMissingFields returns the MissingFields field if non-nil, zero value otherwise.

### GetMissingFieldsOk

`func (o *ConversionTaskStatusResult) GetMissingFieldsOk() (*[]string, bool)`

GetMissingFieldsOk returns a tuple with the MissingFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingFields

`func (o *ConversionTaskStatusResult) SetMissingFields(v []string)`

SetMissingFields sets MissingFields field to given value.


### GetExtractedData

`func (o *ConversionTaskStatusResult) GetExtractedData() map[string]interface{}`

GetExtractedData returns the ExtractedData field if non-nil, zero value otherwise.

### GetExtractedDataOk

`func (o *ConversionTaskStatusResult) GetExtractedDataOk() (*map[string]interface{}, bool)`

GetExtractedDataOk returns a tuple with the ExtractedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedData

`func (o *ConversionTaskStatusResult) SetExtractedData(v map[string]interface{})`

SetExtractedData sets ExtractedData field to given value.


### GetConfidenceScore

`func (o *ConversionTaskStatusResult) GetConfidenceScore() float32`

GetConfidenceScore returns the ConfidenceScore field if non-nil, zero value otherwise.

### GetConfidenceScoreOk

`func (o *ConversionTaskStatusResult) GetConfidenceScoreOk() (*float32, bool)`

GetConfidenceScoreOk returns a tuple with the ConfidenceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceScore

`func (o *ConversionTaskStatusResult) SetConfidenceScore(v float32)`

SetConfidenceScore sets ConfidenceScore field to given value.


### GetValidationErrors

`func (o *ConversionTaskStatusResult) GetValidationErrors() []map[string]interface{}`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *ConversionTaskStatusResult) GetValidationErrorsOk() (*[]map[string]interface{}, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *ConversionTaskStatusResult) SetValidationErrors(v []map[string]interface{})`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *ConversionTaskStatusResult) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### GetCorrectionAttempted

`func (o *ConversionTaskStatusResult) GetCorrectionAttempted() bool`

GetCorrectionAttempted returns the CorrectionAttempted field if non-nil, zero value otherwise.

### GetCorrectionAttemptedOk

`func (o *ConversionTaskStatusResult) GetCorrectionAttemptedOk() (*bool, bool)`

GetCorrectionAttemptedOk returns a tuple with the CorrectionAttempted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrectionAttempted

`func (o *ConversionTaskStatusResult) SetCorrectionAttempted(v bool)`

SetCorrectionAttempted sets CorrectionAttempted field to given value.

### HasCorrectionAttempted

`func (o *ConversionTaskStatusResult) HasCorrectionAttempted() bool`

HasCorrectionAttempted returns a boolean if a field has been set.

### GetErrorCode

`func (o *ConversionTaskStatusResult) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ConversionTaskStatusResult) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ConversionTaskStatusResult) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorMessage

`func (o *ConversionTaskStatusResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ConversionTaskStatusResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ConversionTaskStatusResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### GetDetails

`func (o *ConversionTaskStatusResult) GetDetails() []AFNORErrorDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ConversionTaskStatusResult) GetDetailsOk() (*[]AFNORErrorDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ConversionTaskStatusResult) SetDetails(v []AFNORErrorDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ConversionTaskStatusResult) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTraceback

`func (o *ConversionTaskStatusResult) GetTraceback() string`

GetTraceback returns the Traceback field if non-nil, zero value otherwise.

### GetTracebackOk

`func (o *ConversionTaskStatusResult) GetTracebackOk() (*string, bool)`

GetTracebackOk returns a tuple with the Traceback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceback

`func (o *ConversionTaskStatusResult) SetTraceback(v string)`

SetTraceback sets Traceback field to given value.

### HasTraceback

`func (o *ConversionTaskStatusResult) HasTraceback() bool`

HasTraceback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



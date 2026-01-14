# ConvertSuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Statut de la conversion | [optional] [default to "success"]
**ConversionId** | **string** | Identifiant unique de conversion | 
**DocumentType** | [**DocumentTypeInfo**](DocumentTypeInfo.md) |  | 
**Invoice** | **map[string]interface{}** | Donnees facture au format FacturXInvoice (cf. models.py) | 
**Extraction** | [**ExtractionInfo**](ExtractionInfo.md) |  | 
**Validation** | [**ValidationInfo**](ValidationInfo.md) |  | 
**Files** | [**FilesInfo**](FilesInfo.md) |  | 
**ProcessingTimeMs** | **int32** | Temps de traitement en ms | 
**PdfRegenerated** | Pointer to **bool** | True si le PDF a ete regenere (source non exploitable) | [optional] [default to false]
**PdfRegeneratedReason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewConvertSuccessResponse

`func NewConvertSuccessResponse(conversionId string, documentType DocumentTypeInfo, invoice map[string]interface{}, extraction ExtractionInfo, validation ValidationInfo, files FilesInfo, processingTimeMs int32, ) *ConvertSuccessResponse`

NewConvertSuccessResponse instantiates a new ConvertSuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertSuccessResponseWithDefaults

`func NewConvertSuccessResponseWithDefaults() *ConvertSuccessResponse`

NewConvertSuccessResponseWithDefaults instantiates a new ConvertSuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ConvertSuccessResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConvertSuccessResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConvertSuccessResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConvertSuccessResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConversionId

`func (o *ConvertSuccessResponse) GetConversionId() string`

GetConversionId returns the ConversionId field if non-nil, zero value otherwise.

### GetConversionIdOk

`func (o *ConvertSuccessResponse) GetConversionIdOk() (*string, bool)`

GetConversionIdOk returns a tuple with the ConversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionId

`func (o *ConvertSuccessResponse) SetConversionId(v string)`

SetConversionId sets ConversionId field to given value.


### GetDocumentType

`func (o *ConvertSuccessResponse) GetDocumentType() DocumentTypeInfo`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *ConvertSuccessResponse) GetDocumentTypeOk() (*DocumentTypeInfo, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *ConvertSuccessResponse) SetDocumentType(v DocumentTypeInfo)`

SetDocumentType sets DocumentType field to given value.


### GetInvoice

`func (o *ConvertSuccessResponse) GetInvoice() map[string]interface{}`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *ConvertSuccessResponse) GetInvoiceOk() (*map[string]interface{}, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *ConvertSuccessResponse) SetInvoice(v map[string]interface{})`

SetInvoice sets Invoice field to given value.


### GetExtraction

`func (o *ConvertSuccessResponse) GetExtraction() ExtractionInfo`

GetExtraction returns the Extraction field if non-nil, zero value otherwise.

### GetExtractionOk

`func (o *ConvertSuccessResponse) GetExtractionOk() (*ExtractionInfo, bool)`

GetExtractionOk returns a tuple with the Extraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraction

`func (o *ConvertSuccessResponse) SetExtraction(v ExtractionInfo)`

SetExtraction sets Extraction field to given value.


### GetValidation

`func (o *ConvertSuccessResponse) GetValidation() ValidationInfo`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *ConvertSuccessResponse) GetValidationOk() (*ValidationInfo, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *ConvertSuccessResponse) SetValidation(v ValidationInfo)`

SetValidation sets Validation field to given value.


### GetFiles

`func (o *ConvertSuccessResponse) GetFiles() FilesInfo`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ConvertSuccessResponse) GetFilesOk() (*FilesInfo, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ConvertSuccessResponse) SetFiles(v FilesInfo)`

SetFiles sets Files field to given value.


### GetProcessingTimeMs

`func (o *ConvertSuccessResponse) GetProcessingTimeMs() int32`

GetProcessingTimeMs returns the ProcessingTimeMs field if non-nil, zero value otherwise.

### GetProcessingTimeMsOk

`func (o *ConvertSuccessResponse) GetProcessingTimeMsOk() (*int32, bool)`

GetProcessingTimeMsOk returns a tuple with the ProcessingTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeMs

`func (o *ConvertSuccessResponse) SetProcessingTimeMs(v int32)`

SetProcessingTimeMs sets ProcessingTimeMs field to given value.


### GetPdfRegenerated

`func (o *ConvertSuccessResponse) GetPdfRegenerated() bool`

GetPdfRegenerated returns the PdfRegenerated field if non-nil, zero value otherwise.

### GetPdfRegeneratedOk

`func (o *ConvertSuccessResponse) GetPdfRegeneratedOk() (*bool, bool)`

GetPdfRegeneratedOk returns a tuple with the PdfRegenerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfRegenerated

`func (o *ConvertSuccessResponse) SetPdfRegenerated(v bool)`

SetPdfRegenerated sets PdfRegenerated field to given value.

### HasPdfRegenerated

`func (o *ConvertSuccessResponse) HasPdfRegenerated() bool`

HasPdfRegenerated returns a boolean if a field has been set.

### GetPdfRegeneratedReason

`func (o *ConvertSuccessResponse) GetPdfRegeneratedReason() string`

GetPdfRegeneratedReason returns the PdfRegeneratedReason field if non-nil, zero value otherwise.

### GetPdfRegeneratedReasonOk

`func (o *ConvertSuccessResponse) GetPdfRegeneratedReasonOk() (*string, bool)`

GetPdfRegeneratedReasonOk returns a tuple with the PdfRegeneratedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfRegeneratedReason

`func (o *ConvertSuccessResponse) SetPdfRegeneratedReason(v string)`

SetPdfRegeneratedReason sets PdfRegeneratedReason field to given value.

### HasPdfRegeneratedReason

`func (o *ConvertSuccessResponse) HasPdfRegeneratedReason() bool`

HasPdfRegeneratedReason returns a boolean if a field has been set.

### SetPdfRegeneratedReasonNil

`func (o *ConvertSuccessResponse) SetPdfRegeneratedReasonNil(b bool)`

 SetPdfRegeneratedReasonNil sets the value for PdfRegeneratedReason to be an explicit nil

### UnsetPdfRegeneratedReason
`func (o *ConvertSuccessResponse) UnsetPdfRegeneratedReason()`

UnsetPdfRegeneratedReason ensures that no value is present for PdfRegeneratedReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



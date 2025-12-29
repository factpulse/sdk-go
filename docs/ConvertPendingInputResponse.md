# ConvertPendingInputResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "pending_input"]
**ConversionId** | **string** |  | 
**Message** | Pointer to **string** |  | [optional] [default to "Donnees manquantes requises pour la conformite"]
**Extraction** | [**ExtractionInfo**](ExtractionInfo.md) |  | 
**ExtractedData** | **map[string]interface{}** | Donnees extraites par OCR au format FacturXInvoice | 
**MissingFields** | [**[]MissingField**](MissingField.md) |  | 
**ResumeUrl** | **string** |  | 
**ExpiresAt** | **time.Time** |  | 

## Methods

### NewConvertPendingInputResponse

`func NewConvertPendingInputResponse(conversionId string, extraction ExtractionInfo, extractedData map[string]interface{}, missingFields []MissingField, resumeUrl string, expiresAt time.Time, ) *ConvertPendingInputResponse`

NewConvertPendingInputResponse instantiates a new ConvertPendingInputResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertPendingInputResponseWithDefaults

`func NewConvertPendingInputResponseWithDefaults() *ConvertPendingInputResponse`

NewConvertPendingInputResponseWithDefaults instantiates a new ConvertPendingInputResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ConvertPendingInputResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConvertPendingInputResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConvertPendingInputResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConvertPendingInputResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConversionId

`func (o *ConvertPendingInputResponse) GetConversionId() string`

GetConversionId returns the ConversionId field if non-nil, zero value otherwise.

### GetConversionIdOk

`func (o *ConvertPendingInputResponse) GetConversionIdOk() (*string, bool)`

GetConversionIdOk returns a tuple with the ConversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionId

`func (o *ConvertPendingInputResponse) SetConversionId(v string)`

SetConversionId sets ConversionId field to given value.


### GetMessage

`func (o *ConvertPendingInputResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConvertPendingInputResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConvertPendingInputResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConvertPendingInputResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetExtraction

`func (o *ConvertPendingInputResponse) GetExtraction() ExtractionInfo`

GetExtraction returns the Extraction field if non-nil, zero value otherwise.

### GetExtractionOk

`func (o *ConvertPendingInputResponse) GetExtractionOk() (*ExtractionInfo, bool)`

GetExtractionOk returns a tuple with the Extraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraction

`func (o *ConvertPendingInputResponse) SetExtraction(v ExtractionInfo)`

SetExtraction sets Extraction field to given value.


### GetExtractedData

`func (o *ConvertPendingInputResponse) GetExtractedData() map[string]interface{}`

GetExtractedData returns the ExtractedData field if non-nil, zero value otherwise.

### GetExtractedDataOk

`func (o *ConvertPendingInputResponse) GetExtractedDataOk() (*map[string]interface{}, bool)`

GetExtractedDataOk returns a tuple with the ExtractedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedData

`func (o *ConvertPendingInputResponse) SetExtractedData(v map[string]interface{})`

SetExtractedData sets ExtractedData field to given value.


### GetMissingFields

`func (o *ConvertPendingInputResponse) GetMissingFields() []MissingField`

GetMissingFields returns the MissingFields field if non-nil, zero value otherwise.

### GetMissingFieldsOk

`func (o *ConvertPendingInputResponse) GetMissingFieldsOk() (*[]MissingField, bool)`

GetMissingFieldsOk returns a tuple with the MissingFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingFields

`func (o *ConvertPendingInputResponse) SetMissingFields(v []MissingField)`

SetMissingFields sets MissingFields field to given value.


### GetResumeUrl

`func (o *ConvertPendingInputResponse) GetResumeUrl() string`

GetResumeUrl returns the ResumeUrl field if non-nil, zero value otherwise.

### GetResumeUrlOk

`func (o *ConvertPendingInputResponse) GetResumeUrlOk() (*string, bool)`

GetResumeUrlOk returns a tuple with the ResumeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeUrl

`func (o *ConvertPendingInputResponse) SetResumeUrl(v string)`

SetResumeUrl sets ResumeUrl field to given value.


### GetExpiresAt

`func (o *ConvertPendingInputResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ConvertPendingInputResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ConvertPendingInputResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



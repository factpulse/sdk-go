# ConvertValidationFailedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "validation_failed"]
**ConversionId** | **string** |  | 
**Message** | Pointer to **string** |  | [optional] [default to "Donnees extraites avec erreurs de validation. Completez le formulaire et appelez /resume."]
**Extraction** | [**ExtractionInfo**](ExtractionInfo.md) |  | 
**ExtractedData** | **map[string]interface{}** | Donnees extraites par OCR au format FacturXInvoice (a completer/corriger) | 
**MissingFields** | Pointer to [**[]MissingField**](MissingField.md) | Champs manquants pour conformite Factur-X | [optional] 
**Validation** | [**ValidationInfo**](ValidationInfo.md) |  | 
**ResumeUrl** | **string** | URL pour reprendre la conversion avec corrections | 
**ExpiresAt** | **time.Time** | Expiration de la conversion (1h) | 

## Methods

### NewConvertValidationFailedResponse

`func NewConvertValidationFailedResponse(conversionId string, extraction ExtractionInfo, extractedData map[string]interface{}, validation ValidationInfo, resumeUrl string, expiresAt time.Time, ) *ConvertValidationFailedResponse`

NewConvertValidationFailedResponse instantiates a new ConvertValidationFailedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertValidationFailedResponseWithDefaults

`func NewConvertValidationFailedResponseWithDefaults() *ConvertValidationFailedResponse`

NewConvertValidationFailedResponseWithDefaults instantiates a new ConvertValidationFailedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ConvertValidationFailedResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConvertValidationFailedResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConvertValidationFailedResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConvertValidationFailedResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConversionId

`func (o *ConvertValidationFailedResponse) GetConversionId() string`

GetConversionId returns the ConversionId field if non-nil, zero value otherwise.

### GetConversionIdOk

`func (o *ConvertValidationFailedResponse) GetConversionIdOk() (*string, bool)`

GetConversionIdOk returns a tuple with the ConversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionId

`func (o *ConvertValidationFailedResponse) SetConversionId(v string)`

SetConversionId sets ConversionId field to given value.


### GetMessage

`func (o *ConvertValidationFailedResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConvertValidationFailedResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConvertValidationFailedResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConvertValidationFailedResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetExtraction

`func (o *ConvertValidationFailedResponse) GetExtraction() ExtractionInfo`

GetExtraction returns the Extraction field if non-nil, zero value otherwise.

### GetExtractionOk

`func (o *ConvertValidationFailedResponse) GetExtractionOk() (*ExtractionInfo, bool)`

GetExtractionOk returns a tuple with the Extraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraction

`func (o *ConvertValidationFailedResponse) SetExtraction(v ExtractionInfo)`

SetExtraction sets Extraction field to given value.


### GetExtractedData

`func (o *ConvertValidationFailedResponse) GetExtractedData() map[string]interface{}`

GetExtractedData returns the ExtractedData field if non-nil, zero value otherwise.

### GetExtractedDataOk

`func (o *ConvertValidationFailedResponse) GetExtractedDataOk() (*map[string]interface{}, bool)`

GetExtractedDataOk returns a tuple with the ExtractedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedData

`func (o *ConvertValidationFailedResponse) SetExtractedData(v map[string]interface{})`

SetExtractedData sets ExtractedData field to given value.


### GetMissingFields

`func (o *ConvertValidationFailedResponse) GetMissingFields() []MissingField`

GetMissingFields returns the MissingFields field if non-nil, zero value otherwise.

### GetMissingFieldsOk

`func (o *ConvertValidationFailedResponse) GetMissingFieldsOk() (*[]MissingField, bool)`

GetMissingFieldsOk returns a tuple with the MissingFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingFields

`func (o *ConvertValidationFailedResponse) SetMissingFields(v []MissingField)`

SetMissingFields sets MissingFields field to given value.

### HasMissingFields

`func (o *ConvertValidationFailedResponse) HasMissingFields() bool`

HasMissingFields returns a boolean if a field has been set.

### GetValidation

`func (o *ConvertValidationFailedResponse) GetValidation() ValidationInfo`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *ConvertValidationFailedResponse) GetValidationOk() (*ValidationInfo, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *ConvertValidationFailedResponse) SetValidation(v ValidationInfo)`

SetValidation sets Validation field to given value.


### GetResumeUrl

`func (o *ConvertValidationFailedResponse) GetResumeUrl() string`

GetResumeUrl returns the ResumeUrl field if non-nil, zero value otherwise.

### GetResumeUrlOk

`func (o *ConvertValidationFailedResponse) GetResumeUrlOk() (*string, bool)`

GetResumeUrlOk returns a tuple with the ResumeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeUrl

`func (o *ConvertValidationFailedResponse) SetResumeUrl(v string)`

SetResumeUrl sets ResumeUrl field to given value.


### GetExpiresAt

`func (o *ConvertValidationFailedResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ConvertValidationFailedResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ConvertValidationFailedResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



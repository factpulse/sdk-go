# ConversionValidationFailedResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "VALIDATION_FAILED"]
**ConversionId** | **string** |  | 
**Message** | **string** |  | 
**ExtractedData** | **map[string]interface{}** |  | 
**Extraction** | Pointer to [**NullableConversionExtractionInfo**](ConversionExtractionInfo.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Profile** | **string** |  | 
**ProcessingTimeMs** | **int32** |  | 
**CorrectionAttempted** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewConversionValidationFailedResult

`func NewConversionValidationFailedResult(conversionId string, message string, extractedData map[string]interface{}, profile string, processingTimeMs int32, ) *ConversionValidationFailedResult`

NewConversionValidationFailedResult instantiates a new ConversionValidationFailedResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversionValidationFailedResultWithDefaults

`func NewConversionValidationFailedResultWithDefaults() *ConversionValidationFailedResult`

NewConversionValidationFailedResultWithDefaults instantiates a new ConversionValidationFailedResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ConversionValidationFailedResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConversionValidationFailedResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConversionValidationFailedResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConversionValidationFailedResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConversionId

`func (o *ConversionValidationFailedResult) GetConversionId() string`

GetConversionId returns the ConversionId field if non-nil, zero value otherwise.

### GetConversionIdOk

`func (o *ConversionValidationFailedResult) GetConversionIdOk() (*string, bool)`

GetConversionIdOk returns a tuple with the ConversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionId

`func (o *ConversionValidationFailedResult) SetConversionId(v string)`

SetConversionId sets ConversionId field to given value.


### GetMessage

`func (o *ConversionValidationFailedResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConversionValidationFailedResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConversionValidationFailedResult) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetExtractedData

`func (o *ConversionValidationFailedResult) GetExtractedData() map[string]interface{}`

GetExtractedData returns the ExtractedData field if non-nil, zero value otherwise.

### GetExtractedDataOk

`func (o *ConversionValidationFailedResult) GetExtractedDataOk() (*map[string]interface{}, bool)`

GetExtractedDataOk returns a tuple with the ExtractedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedData

`func (o *ConversionValidationFailedResult) SetExtractedData(v map[string]interface{})`

SetExtractedData sets ExtractedData field to given value.


### GetExtraction

`func (o *ConversionValidationFailedResult) GetExtraction() ConversionExtractionInfo`

GetExtraction returns the Extraction field if non-nil, zero value otherwise.

### GetExtractionOk

`func (o *ConversionValidationFailedResult) GetExtractionOk() (*ConversionExtractionInfo, bool)`

GetExtractionOk returns a tuple with the Extraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraction

`func (o *ConversionValidationFailedResult) SetExtraction(v ConversionExtractionInfo)`

SetExtraction sets Extraction field to given value.

### HasExtraction

`func (o *ConversionValidationFailedResult) HasExtraction() bool`

HasExtraction returns a boolean if a field has been set.

### SetExtractionNil

`func (o *ConversionValidationFailedResult) SetExtractionNil(b bool)`

 SetExtractionNil sets the value for Extraction to be an explicit nil

### UnsetExtraction
`func (o *ConversionValidationFailedResult) UnsetExtraction()`

UnsetExtraction ensures that no value is present for Extraction, not even an explicit nil
### GetValidationErrors

`func (o *ConversionValidationFailedResult) GetValidationErrors() []map[string]interface{}`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *ConversionValidationFailedResult) GetValidationErrorsOk() (*[]map[string]interface{}, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *ConversionValidationFailedResult) SetValidationErrors(v []map[string]interface{})`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *ConversionValidationFailedResult) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### GetProfile

`func (o *ConversionValidationFailedResult) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ConversionValidationFailedResult) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ConversionValidationFailedResult) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetProcessingTimeMs

`func (o *ConversionValidationFailedResult) GetProcessingTimeMs() int32`

GetProcessingTimeMs returns the ProcessingTimeMs field if non-nil, zero value otherwise.

### GetProcessingTimeMsOk

`func (o *ConversionValidationFailedResult) GetProcessingTimeMsOk() (*int32, bool)`

GetProcessingTimeMsOk returns a tuple with the ProcessingTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeMs

`func (o *ConversionValidationFailedResult) SetProcessingTimeMs(v int32)`

SetProcessingTimeMs sets ProcessingTimeMs field to given value.


### GetCorrectionAttempted

`func (o *ConversionValidationFailedResult) GetCorrectionAttempted() bool`

GetCorrectionAttempted returns the CorrectionAttempted field if non-nil, zero value otherwise.

### GetCorrectionAttemptedOk

`func (o *ConversionValidationFailedResult) GetCorrectionAttemptedOk() (*bool, bool)`

GetCorrectionAttemptedOk returns a tuple with the CorrectionAttempted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrectionAttempted

`func (o *ConversionValidationFailedResult) SetCorrectionAttempted(v bool)`

SetCorrectionAttempted sets CorrectionAttempted field to given value.

### HasCorrectionAttempted

`func (o *ConversionValidationFailedResult) HasCorrectionAttempted() bool`

HasCorrectionAttempted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



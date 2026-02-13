# ConversionPendingInputResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "PENDING_INPUT"]
**ConversionId** | **string** |  | 
**Message** | **string** |  | 
**MissingFields** | **[]string** |  | 
**ExtractedData** | **map[string]interface{}** |  | 
**ConfidenceScore** | **float32** |  | 
**ProcessingTimeMs** | **int32** |  | 

## Methods

### NewConversionPendingInputResult

`func NewConversionPendingInputResult(conversionId string, message string, missingFields []string, extractedData map[string]interface{}, confidenceScore float32, processingTimeMs int32, ) *ConversionPendingInputResult`

NewConversionPendingInputResult instantiates a new ConversionPendingInputResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversionPendingInputResultWithDefaults

`func NewConversionPendingInputResultWithDefaults() *ConversionPendingInputResult`

NewConversionPendingInputResultWithDefaults instantiates a new ConversionPendingInputResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ConversionPendingInputResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConversionPendingInputResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConversionPendingInputResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConversionPendingInputResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConversionId

`func (o *ConversionPendingInputResult) GetConversionId() string`

GetConversionId returns the ConversionId field if non-nil, zero value otherwise.

### GetConversionIdOk

`func (o *ConversionPendingInputResult) GetConversionIdOk() (*string, bool)`

GetConversionIdOk returns a tuple with the ConversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionId

`func (o *ConversionPendingInputResult) SetConversionId(v string)`

SetConversionId sets ConversionId field to given value.


### GetMessage

`func (o *ConversionPendingInputResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConversionPendingInputResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConversionPendingInputResult) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMissingFields

`func (o *ConversionPendingInputResult) GetMissingFields() []string`

GetMissingFields returns the MissingFields field if non-nil, zero value otherwise.

### GetMissingFieldsOk

`func (o *ConversionPendingInputResult) GetMissingFieldsOk() (*[]string, bool)`

GetMissingFieldsOk returns a tuple with the MissingFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingFields

`func (o *ConversionPendingInputResult) SetMissingFields(v []string)`

SetMissingFields sets MissingFields field to given value.


### GetExtractedData

`func (o *ConversionPendingInputResult) GetExtractedData() map[string]interface{}`

GetExtractedData returns the ExtractedData field if non-nil, zero value otherwise.

### GetExtractedDataOk

`func (o *ConversionPendingInputResult) GetExtractedDataOk() (*map[string]interface{}, bool)`

GetExtractedDataOk returns a tuple with the ExtractedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedData

`func (o *ConversionPendingInputResult) SetExtractedData(v map[string]interface{})`

SetExtractedData sets ExtractedData field to given value.


### GetConfidenceScore

`func (o *ConversionPendingInputResult) GetConfidenceScore() float32`

GetConfidenceScore returns the ConfidenceScore field if non-nil, zero value otherwise.

### GetConfidenceScoreOk

`func (o *ConversionPendingInputResult) GetConfidenceScoreOk() (*float32, bool)`

GetConfidenceScoreOk returns a tuple with the ConfidenceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceScore

`func (o *ConversionPendingInputResult) SetConfidenceScore(v float32)`

SetConfidenceScore sets ConfidenceScore field to given value.


### GetProcessingTimeMs

`func (o *ConversionPendingInputResult) GetProcessingTimeMs() int32`

GetProcessingTimeMs returns the ProcessingTimeMs field if non-nil, zero value otherwise.

### GetProcessingTimeMsOk

`func (o *ConversionPendingInputResult) GetProcessingTimeMsOk() (*int32, bool)`

GetProcessingTimeMsOk returns a tuple with the ProcessingTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeMs

`func (o *ConversionPendingInputResult) SetProcessingTimeMs(v int32)`

SetProcessingTimeMs sets ProcessingTimeMs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



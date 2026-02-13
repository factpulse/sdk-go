# ConversionExtractionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfidenceScore** | **float32** |  | 
**OcrProvider** | Pointer to **string** |  | [optional] [default to "mistral_ocr"]

## Methods

### NewConversionExtractionInfo

`func NewConversionExtractionInfo(confidenceScore float32, ) *ConversionExtractionInfo`

NewConversionExtractionInfo instantiates a new ConversionExtractionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversionExtractionInfoWithDefaults

`func NewConversionExtractionInfoWithDefaults() *ConversionExtractionInfo`

NewConversionExtractionInfoWithDefaults instantiates a new ConversionExtractionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfidenceScore

`func (o *ConversionExtractionInfo) GetConfidenceScore() float32`

GetConfidenceScore returns the ConfidenceScore field if non-nil, zero value otherwise.

### GetConfidenceScoreOk

`func (o *ConversionExtractionInfo) GetConfidenceScoreOk() (*float32, bool)`

GetConfidenceScoreOk returns a tuple with the ConfidenceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceScore

`func (o *ConversionExtractionInfo) SetConfidenceScore(v float32)`

SetConfidenceScore sets ConfidenceScore field to given value.


### GetOcrProvider

`func (o *ConversionExtractionInfo) GetOcrProvider() string`

GetOcrProvider returns the OcrProvider field if non-nil, zero value otherwise.

### GetOcrProviderOk

`func (o *ConversionExtractionInfo) GetOcrProviderOk() (*string, bool)`

GetOcrProviderOk returns a tuple with the OcrProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcrProvider

`func (o *ConversionExtractionInfo) SetOcrProvider(v string)`

SetOcrProvider sets OcrProvider field to given value.

### HasOcrProvider

`func (o *ConversionExtractionInfo) HasOcrProvider() bool`

HasOcrProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



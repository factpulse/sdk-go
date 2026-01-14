# ExtractionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfidenceScore** | **float32** | Score de confiance global (0-1) | 
**FieldsExtracted** | **int32** | Nombre de champs extraits | 
**FieldsEnriched** | Pointer to **int32** | Nombre de champs enrichis automatiquement | [optional] [default to 0]
**FieldsMissing** | Pointer to **int32** | Nombre de champs manquants | [optional] [default to 0]

## Methods

### NewExtractionInfo

`func NewExtractionInfo(confidenceScore float32, fieldsExtracted int32, ) *ExtractionInfo`

NewExtractionInfo instantiates a new ExtractionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtractionInfoWithDefaults

`func NewExtractionInfoWithDefaults() *ExtractionInfo`

NewExtractionInfoWithDefaults instantiates a new ExtractionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfidenceScore

`func (o *ExtractionInfo) GetConfidenceScore() float32`

GetConfidenceScore returns the ConfidenceScore field if non-nil, zero value otherwise.

### GetConfidenceScoreOk

`func (o *ExtractionInfo) GetConfidenceScoreOk() (*float32, bool)`

GetConfidenceScoreOk returns a tuple with the ConfidenceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceScore

`func (o *ExtractionInfo) SetConfidenceScore(v float32)`

SetConfidenceScore sets ConfidenceScore field to given value.


### GetFieldsExtracted

`func (o *ExtractionInfo) GetFieldsExtracted() int32`

GetFieldsExtracted returns the FieldsExtracted field if non-nil, zero value otherwise.

### GetFieldsExtractedOk

`func (o *ExtractionInfo) GetFieldsExtractedOk() (*int32, bool)`

GetFieldsExtractedOk returns a tuple with the FieldsExtracted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsExtracted

`func (o *ExtractionInfo) SetFieldsExtracted(v int32)`

SetFieldsExtracted sets FieldsExtracted field to given value.


### GetFieldsEnriched

`func (o *ExtractionInfo) GetFieldsEnriched() int32`

GetFieldsEnriched returns the FieldsEnriched field if non-nil, zero value otherwise.

### GetFieldsEnrichedOk

`func (o *ExtractionInfo) GetFieldsEnrichedOk() (*int32, bool)`

GetFieldsEnrichedOk returns a tuple with the FieldsEnriched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsEnriched

`func (o *ExtractionInfo) SetFieldsEnriched(v int32)`

SetFieldsEnriched sets FieldsEnriched field to given value.

### HasFieldsEnriched

`func (o *ExtractionInfo) HasFieldsEnriched() bool`

HasFieldsEnriched returns a boolean if a field has been set.

### GetFieldsMissing

`func (o *ExtractionInfo) GetFieldsMissing() int32`

GetFieldsMissing returns the FieldsMissing field if non-nil, zero value otherwise.

### GetFieldsMissingOk

`func (o *ExtractionInfo) GetFieldsMissingOk() (*int32, bool)`

GetFieldsMissingOk returns a tuple with the FieldsMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsMissing

`func (o *ExtractionInfo) SetFieldsMissing(v int32)`

SetFieldsMissing sets FieldsMissing field to given value.

### HasFieldsMissing

`func (o *ExtractionInfo) HasFieldsMissing() bool`

HasFieldsMissing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



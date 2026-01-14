# AFNORSearchFlowContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** |  | [optional] 
**Filters** | Pointer to [**AFNORSearchFlowFilters**](AFNORSearchFlowFilters.md) |  | [optional] 
**Results** | Pointer to [**[]AFNORFlow**](AFNORFlow.md) |  | [optional] 

## Methods

### NewAFNORSearchFlowContent

`func NewAFNORSearchFlowContent() *AFNORSearchFlowContent`

NewAFNORSearchFlowContent instantiates a new AFNORSearchFlowContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORSearchFlowContentWithDefaults

`func NewAFNORSearchFlowContentWithDefaults() *AFNORSearchFlowContent`

NewAFNORSearchFlowContentWithDefaults instantiates a new AFNORSearchFlowContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *AFNORSearchFlowContent) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AFNORSearchFlowContent) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AFNORSearchFlowContent) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AFNORSearchFlowContent) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetFilters

`func (o *AFNORSearchFlowContent) GetFilters() AFNORSearchFlowFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *AFNORSearchFlowContent) GetFiltersOk() (*AFNORSearchFlowFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *AFNORSearchFlowContent) SetFilters(v AFNORSearchFlowFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *AFNORSearchFlowContent) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetResults

`func (o *AFNORSearchFlowContent) GetResults() []AFNORFlow`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AFNORSearchFlowContent) GetResultsOk() (*[]AFNORFlow, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AFNORSearchFlowContent) SetResults(v []AFNORFlow)`

SetResults sets Results field to given value.

### HasResults

`func (o *AFNORSearchFlowContent) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



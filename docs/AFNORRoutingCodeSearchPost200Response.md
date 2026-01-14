# AFNORRoutingCodeSearchPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Search** | Pointer to [**AFNORRoutingCodeSearch**](AFNORRoutingCodeSearch.md) |  | [optional] 
**TotalNumberOfResults** | Pointer to **int32** | The total number of results | [optional] 
**Results** | Pointer to [**[]AFNORRoutingCodePayloadHistoryLegalUnitFacility**](AFNORRoutingCodePayloadHistoryLegalUnitFacility.md) |  | [optional] 

## Methods

### NewAFNORRoutingCodeSearchPost200Response

`func NewAFNORRoutingCodeSearchPost200Response() *AFNORRoutingCodeSearchPost200Response`

NewAFNORRoutingCodeSearchPost200Response instantiates a new AFNORRoutingCodeSearchPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORRoutingCodeSearchPost200ResponseWithDefaults

`func NewAFNORRoutingCodeSearchPost200ResponseWithDefaults() *AFNORRoutingCodeSearchPost200Response`

NewAFNORRoutingCodeSearchPost200ResponseWithDefaults instantiates a new AFNORRoutingCodeSearchPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearch

`func (o *AFNORRoutingCodeSearchPost200Response) GetSearch() AFNORRoutingCodeSearch`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *AFNORRoutingCodeSearchPost200Response) GetSearchOk() (*AFNORRoutingCodeSearch, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *AFNORRoutingCodeSearchPost200Response) SetSearch(v AFNORRoutingCodeSearch)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *AFNORRoutingCodeSearchPost200Response) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetTotalNumberOfResults

`func (o *AFNORRoutingCodeSearchPost200Response) GetTotalNumberOfResults() int32`

GetTotalNumberOfResults returns the TotalNumberOfResults field if non-nil, zero value otherwise.

### GetTotalNumberOfResultsOk

`func (o *AFNORRoutingCodeSearchPost200Response) GetTotalNumberOfResultsOk() (*int32, bool)`

GetTotalNumberOfResultsOk returns a tuple with the TotalNumberOfResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfResults

`func (o *AFNORRoutingCodeSearchPost200Response) SetTotalNumberOfResults(v int32)`

SetTotalNumberOfResults sets TotalNumberOfResults field to given value.

### HasTotalNumberOfResults

`func (o *AFNORRoutingCodeSearchPost200Response) HasTotalNumberOfResults() bool`

HasTotalNumberOfResults returns a boolean if a field has been set.

### GetResults

`func (o *AFNORRoutingCodeSearchPost200Response) GetResults() []AFNORRoutingCodePayloadHistoryLegalUnitFacility`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AFNORRoutingCodeSearchPost200Response) GetResultsOk() (*[]AFNORRoutingCodePayloadHistoryLegalUnitFacility, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AFNORRoutingCodeSearchPost200Response) SetResults(v []AFNORRoutingCodePayloadHistoryLegalUnitFacility)`

SetResults sets Results field to given value.

### HasResults

`func (o *AFNORRoutingCodeSearchPost200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



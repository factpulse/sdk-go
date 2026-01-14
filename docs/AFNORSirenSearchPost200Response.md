# AFNORSirenSearchPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Search** | Pointer to [**AFNORSearchSiren**](AFNORSearchSiren.md) |  | [optional] 
**TotalNumberOfResults** | Pointer to **int32** | The total number of results | [optional] 
**Results** | Pointer to [**[]AFNORLegalUnitPayloadHistory**](AFNORLegalUnitPayloadHistory.md) |  | [optional] 

## Methods

### NewAFNORSirenSearchPost200Response

`func NewAFNORSirenSearchPost200Response() *AFNORSirenSearchPost200Response`

NewAFNORSirenSearchPost200Response instantiates a new AFNORSirenSearchPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORSirenSearchPost200ResponseWithDefaults

`func NewAFNORSirenSearchPost200ResponseWithDefaults() *AFNORSirenSearchPost200Response`

NewAFNORSirenSearchPost200ResponseWithDefaults instantiates a new AFNORSirenSearchPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearch

`func (o *AFNORSirenSearchPost200Response) GetSearch() AFNORSearchSiren`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *AFNORSirenSearchPost200Response) GetSearchOk() (*AFNORSearchSiren, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *AFNORSirenSearchPost200Response) SetSearch(v AFNORSearchSiren)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *AFNORSirenSearchPost200Response) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetTotalNumberOfResults

`func (o *AFNORSirenSearchPost200Response) GetTotalNumberOfResults() int32`

GetTotalNumberOfResults returns the TotalNumberOfResults field if non-nil, zero value otherwise.

### GetTotalNumberOfResultsOk

`func (o *AFNORSirenSearchPost200Response) GetTotalNumberOfResultsOk() (*int32, bool)`

GetTotalNumberOfResultsOk returns a tuple with the TotalNumberOfResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfResults

`func (o *AFNORSirenSearchPost200Response) SetTotalNumberOfResults(v int32)`

SetTotalNumberOfResults sets TotalNumberOfResults field to given value.

### HasTotalNumberOfResults

`func (o *AFNORSirenSearchPost200Response) HasTotalNumberOfResults() bool`

HasTotalNumberOfResults returns a boolean if a field has been set.

### GetResults

`func (o *AFNORSirenSearchPost200Response) GetResults() []AFNORLegalUnitPayloadHistory`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AFNORSirenSearchPost200Response) GetResultsOk() (*[]AFNORLegalUnitPayloadHistory, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AFNORSirenSearchPost200Response) SetResults(v []AFNORLegalUnitPayloadHistory)`

SetResults sets Results field to given value.

### HasResults

`func (o *AFNORSirenSearchPost200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



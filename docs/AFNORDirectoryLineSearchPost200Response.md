# AFNORDirectoryLineSearchPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Search** | Pointer to [**AFNORSearchDirectoryLine**](AFNORSearchDirectoryLine.md) |  | [optional] 
**TotalNumberOfResults** | Pointer to **int32** | The total number of results | [optional] 
**Results** | Pointer to [**[]AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode**](AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode.md) |  | [optional] 

## Methods

### NewAFNORDirectoryLineSearchPost200Response

`func NewAFNORDirectoryLineSearchPost200Response() *AFNORDirectoryLineSearchPost200Response`

NewAFNORDirectoryLineSearchPost200Response instantiates a new AFNORDirectoryLineSearchPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORDirectoryLineSearchPost200ResponseWithDefaults

`func NewAFNORDirectoryLineSearchPost200ResponseWithDefaults() *AFNORDirectoryLineSearchPost200Response`

NewAFNORDirectoryLineSearchPost200ResponseWithDefaults instantiates a new AFNORDirectoryLineSearchPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearch

`func (o *AFNORDirectoryLineSearchPost200Response) GetSearch() AFNORSearchDirectoryLine`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *AFNORDirectoryLineSearchPost200Response) GetSearchOk() (*AFNORSearchDirectoryLine, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *AFNORDirectoryLineSearchPost200Response) SetSearch(v AFNORSearchDirectoryLine)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *AFNORDirectoryLineSearchPost200Response) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetTotalNumberOfResults

`func (o *AFNORDirectoryLineSearchPost200Response) GetTotalNumberOfResults() int32`

GetTotalNumberOfResults returns the TotalNumberOfResults field if non-nil, zero value otherwise.

### GetTotalNumberOfResultsOk

`func (o *AFNORDirectoryLineSearchPost200Response) GetTotalNumberOfResultsOk() (*int32, bool)`

GetTotalNumberOfResultsOk returns a tuple with the TotalNumberOfResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfResults

`func (o *AFNORDirectoryLineSearchPost200Response) SetTotalNumberOfResults(v int32)`

SetTotalNumberOfResults sets TotalNumberOfResults field to given value.

### HasTotalNumberOfResults

`func (o *AFNORDirectoryLineSearchPost200Response) HasTotalNumberOfResults() bool`

HasTotalNumberOfResults returns a boolean if a field has been set.

### GetResults

`func (o *AFNORDirectoryLineSearchPost200Response) GetResults() []AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AFNORDirectoryLineSearchPost200Response) GetResultsOk() (*[]AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AFNORDirectoryLineSearchPost200Response) SetResults(v []AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode)`

SetResults sets Results field to given value.

### HasResults

`func (o *AFNORDirectoryLineSearchPost200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



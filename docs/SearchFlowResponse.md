# SearchFlowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | Total number of results | 
**Offset** | **int32** | Applied offset | 
**Limit** | **int32** | Results limit | 
**Results** | [**[]FlowSummary**](FlowSummary.md) | List of found flows | 

## Methods

### NewSearchFlowResponse

`func NewSearchFlowResponse(total int32, offset int32, limit int32, results []FlowSummary, ) *SearchFlowResponse`

NewSearchFlowResponse instantiates a new SearchFlowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFlowResponseWithDefaults

`func NewSearchFlowResponseWithDefaults() *SearchFlowResponse`

NewSearchFlowResponseWithDefaults instantiates a new SearchFlowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *SearchFlowResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SearchFlowResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SearchFlowResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetOffset

`func (o *SearchFlowResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SearchFlowResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SearchFlowResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *SearchFlowResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchFlowResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchFlowResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetResults

`func (o *SearchFlowResponse) GetResults() []FlowSummary`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchFlowResponse) GetResultsOk() (*[]FlowSummary, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchFlowResponse) SetResults(v []FlowSummary)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



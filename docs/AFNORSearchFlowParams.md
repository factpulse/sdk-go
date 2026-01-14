# AFNORSearchFlowParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | Maximum number of results that may be returned | [optional] [default to 25]
**Where** | [**AFNORSearchFlowFilters**](AFNORSearchFlowFilters.md) |  | 

## Methods

### NewAFNORSearchFlowParams

`func NewAFNORSearchFlowParams(where AFNORSearchFlowFilters, ) *AFNORSearchFlowParams`

NewAFNORSearchFlowParams instantiates a new AFNORSearchFlowParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORSearchFlowParamsWithDefaults

`func NewAFNORSearchFlowParamsWithDefaults() *AFNORSearchFlowParams`

NewAFNORSearchFlowParamsWithDefaults instantiates a new AFNORSearchFlowParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *AFNORSearchFlowParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AFNORSearchFlowParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AFNORSearchFlowParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AFNORSearchFlowParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetWhere

`func (o *AFNORSearchFlowParams) GetWhere() AFNORSearchFlowFilters`

GetWhere returns the Where field if non-nil, zero value otherwise.

### GetWhereOk

`func (o *AFNORSearchFlowParams) GetWhereOk() (*AFNORSearchFlowFilters, bool)`

GetWhereOk returns a tuple with the Where field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhere

`func (o *AFNORSearchFlowParams) SetWhere(v AFNORSearchFlowFilters)`

SetWhere sets Where field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



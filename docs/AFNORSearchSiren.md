# AFNORSearchSiren

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**AFNORSearchSirenFilters**](AFNORSearchSirenFilters.md) |  | [optional] 
**Sorting** | Pointer to [**[]AFNORSearchSirenSortingInner**](AFNORSearchSirenSortingInner.md) | Sorting criteria on a field and an order (ascending or descending). | [optional] 
**Fields** | Pointer to [**[]AFNORSirenField**](AFNORSirenField.md) | Allows you to filter the desired fields in the response. | [optional] 
**Limit** | Pointer to **int32** | Maximum number of results | [optional] 
**Ignore** | Pointer to **int32** | Number of results to skip | [optional] 

## Methods

### NewAFNORSearchSiren

`func NewAFNORSearchSiren() *AFNORSearchSiren`

NewAFNORSearchSiren instantiates a new AFNORSearchSiren object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORSearchSirenWithDefaults

`func NewAFNORSearchSirenWithDefaults() *AFNORSearchSiren`

NewAFNORSearchSirenWithDefaults instantiates a new AFNORSearchSiren object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *AFNORSearchSiren) GetFilters() AFNORSearchSirenFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *AFNORSearchSiren) GetFiltersOk() (*AFNORSearchSirenFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *AFNORSearchSiren) SetFilters(v AFNORSearchSirenFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *AFNORSearchSiren) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSorting

`func (o *AFNORSearchSiren) GetSorting() []AFNORSearchSirenSortingInner`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *AFNORSearchSiren) GetSortingOk() (*[]AFNORSearchSirenSortingInner, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *AFNORSearchSiren) SetSorting(v []AFNORSearchSirenSortingInner)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *AFNORSearchSiren) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### GetFields

`func (o *AFNORSearchSiren) GetFields() []AFNORSirenField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AFNORSearchSiren) GetFieldsOk() (*[]AFNORSirenField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AFNORSearchSiren) SetFields(v []AFNORSirenField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AFNORSearchSiren) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetLimit

`func (o *AFNORSearchSiren) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AFNORSearchSiren) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AFNORSearchSiren) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AFNORSearchSiren) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetIgnore

`func (o *AFNORSearchSiren) GetIgnore() int32`

GetIgnore returns the Ignore field if non-nil, zero value otherwise.

### GetIgnoreOk

`func (o *AFNORSearchSiren) GetIgnoreOk() (*int32, bool)`

GetIgnoreOk returns a tuple with the Ignore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnore

`func (o *AFNORSearchSiren) SetIgnore(v int32)`

SetIgnore sets Ignore field to given value.

### HasIgnore

`func (o *AFNORSearchSiren) HasIgnore() bool`

HasIgnore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



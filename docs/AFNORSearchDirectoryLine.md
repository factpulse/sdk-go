# AFNORSearchDirectoryLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**AFNORSearchDirectoryLineFilters**](AFNORSearchDirectoryLineFilters.md) |  | [optional] 
**Sorting** | Pointer to [**[]AFNORSearchDirectoryLineSortingInner**](AFNORSearchDirectoryLineSortingInner.md) | Sorting criteria on a field and an order (ascending or descending). | [optional] 
**Fields** | Pointer to [**[]AFNORDirectoryLineField**](AFNORDirectoryLineField.md) | Allows you to filter the desired fields in the response. | [optional] 
**Limit** | Pointer to **int32** | Maximum number of results | [optional] 
**Ignore** | Pointer to **int32** | Number of results to skip | [optional] 

## Methods

### NewAFNORSearchDirectoryLine

`func NewAFNORSearchDirectoryLine() *AFNORSearchDirectoryLine`

NewAFNORSearchDirectoryLine instantiates a new AFNORSearchDirectoryLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORSearchDirectoryLineWithDefaults

`func NewAFNORSearchDirectoryLineWithDefaults() *AFNORSearchDirectoryLine`

NewAFNORSearchDirectoryLineWithDefaults instantiates a new AFNORSearchDirectoryLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *AFNORSearchDirectoryLine) GetFilters() AFNORSearchDirectoryLineFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *AFNORSearchDirectoryLine) GetFiltersOk() (*AFNORSearchDirectoryLineFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *AFNORSearchDirectoryLine) SetFilters(v AFNORSearchDirectoryLineFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *AFNORSearchDirectoryLine) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSorting

`func (o *AFNORSearchDirectoryLine) GetSorting() []AFNORSearchDirectoryLineSortingInner`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *AFNORSearchDirectoryLine) GetSortingOk() (*[]AFNORSearchDirectoryLineSortingInner, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *AFNORSearchDirectoryLine) SetSorting(v []AFNORSearchDirectoryLineSortingInner)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *AFNORSearchDirectoryLine) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### GetFields

`func (o *AFNORSearchDirectoryLine) GetFields() []AFNORDirectoryLineField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AFNORSearchDirectoryLine) GetFieldsOk() (*[]AFNORDirectoryLineField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AFNORSearchDirectoryLine) SetFields(v []AFNORDirectoryLineField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AFNORSearchDirectoryLine) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetLimit

`func (o *AFNORSearchDirectoryLine) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AFNORSearchDirectoryLine) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AFNORSearchDirectoryLine) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AFNORSearchDirectoryLine) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetIgnore

`func (o *AFNORSearchDirectoryLine) GetIgnore() int32`

GetIgnore returns the Ignore field if non-nil, zero value otherwise.

### GetIgnoreOk

`func (o *AFNORSearchDirectoryLine) GetIgnoreOk() (*int32, bool)`

GetIgnoreOk returns a tuple with the Ignore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnore

`func (o *AFNORSearchDirectoryLine) SetIgnore(v int32)`

SetIgnore sets Ignore field to given value.

### HasIgnore

`func (o *AFNORSearchDirectoryLine) HasIgnore() bool`

HasIgnore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



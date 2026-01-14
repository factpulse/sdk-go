# AFNORRoutingCodeSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**AFNORRoutingCodeSearchFilters**](AFNORRoutingCodeSearchFilters.md) |  | [optional] 
**Sorting** | Pointer to [**[]AFNORRoutingCodeSearchSortingInner**](AFNORRoutingCodeSearchSortingInner.md) | Sorting criteria on a field and an order (ascending or descending). | [optional] 
**Fields** | Pointer to [**[]AFNORRoutingCodeField**](AFNORRoutingCodeField.md) | Allows you to filter the desired fields in the response. | [optional] 
**Include** | Pointer to **[]string** |  | [optional] 
**Limit** | Pointer to **int32** | Maximum number of results | [optional] 
**Ignore** | Pointer to **int32** | Number of results to skip | [optional] 

## Methods

### NewAFNORRoutingCodeSearch

`func NewAFNORRoutingCodeSearch() *AFNORRoutingCodeSearch`

NewAFNORRoutingCodeSearch instantiates a new AFNORRoutingCodeSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORRoutingCodeSearchWithDefaults

`func NewAFNORRoutingCodeSearchWithDefaults() *AFNORRoutingCodeSearch`

NewAFNORRoutingCodeSearchWithDefaults instantiates a new AFNORRoutingCodeSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *AFNORRoutingCodeSearch) GetFilters() AFNORRoutingCodeSearchFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *AFNORRoutingCodeSearch) GetFiltersOk() (*AFNORRoutingCodeSearchFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *AFNORRoutingCodeSearch) SetFilters(v AFNORRoutingCodeSearchFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *AFNORRoutingCodeSearch) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSorting

`func (o *AFNORRoutingCodeSearch) GetSorting() []AFNORRoutingCodeSearchSortingInner`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *AFNORRoutingCodeSearch) GetSortingOk() (*[]AFNORRoutingCodeSearchSortingInner, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *AFNORRoutingCodeSearch) SetSorting(v []AFNORRoutingCodeSearchSortingInner)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *AFNORRoutingCodeSearch) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### GetFields

`func (o *AFNORRoutingCodeSearch) GetFields() []AFNORRoutingCodeField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AFNORRoutingCodeSearch) GetFieldsOk() (*[]AFNORRoutingCodeField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AFNORRoutingCodeSearch) SetFields(v []AFNORRoutingCodeField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AFNORRoutingCodeSearch) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetInclude

`func (o *AFNORRoutingCodeSearch) GetInclude() []string`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *AFNORRoutingCodeSearch) GetIncludeOk() (*[]string, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *AFNORRoutingCodeSearch) SetInclude(v []string)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *AFNORRoutingCodeSearch) HasInclude() bool`

HasInclude returns a boolean if a field has been set.

### GetLimit

`func (o *AFNORRoutingCodeSearch) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AFNORRoutingCodeSearch) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AFNORRoutingCodeSearch) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AFNORRoutingCodeSearch) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetIgnore

`func (o *AFNORRoutingCodeSearch) GetIgnore() int32`

GetIgnore returns the Ignore field if non-nil, zero value otherwise.

### GetIgnoreOk

`func (o *AFNORRoutingCodeSearch) GetIgnoreOk() (*int32, bool)`

GetIgnoreOk returns a tuple with the Ignore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnore

`func (o *AFNORRoutingCodeSearch) SetIgnore(v int32)`

SetIgnore sets Ignore field to given value.

### HasIgnore

`func (o *AFNORRoutingCodeSearch) HasIgnore() bool`

HasIgnore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



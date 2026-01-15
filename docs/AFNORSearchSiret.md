# AFNORSearchSiret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**AFNORSearchSiretFilters**](AFNORSearchSiretFilters.md) |  | [optional] 
**Sorting** | Pointer to [**[]AFNORSearchSiretSortingInner**](AFNORSearchSiretSortingInner.md) | Sorting criteria on a field and an order (ascending or descending). | [optional] 
**Fields** | Pointer to [**[]AFNORSiretField**](AFNORSiretField.md) | Allows you to filter the desired fields in the response. | [optional] 
**Include** | Pointer to **[]string** |  | [optional] 
**Limit** | Pointer to **int32** | Maximum number of results | [optional] 
**Ignore** | Pointer to **int32** | Number of results to skip | [optional] 

## Methods

### NewAFNORSearchSiret

`func NewAFNORSearchSiret() *AFNORSearchSiret`

NewAFNORSearchSiret instantiates a new AFNORSearchSiret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORSearchSiretWithDefaults

`func NewAFNORSearchSiretWithDefaults() *AFNORSearchSiret`

NewAFNORSearchSiretWithDefaults instantiates a new AFNORSearchSiret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *AFNORSearchSiret) GetFilters() AFNORSearchSiretFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *AFNORSearchSiret) GetFiltersOk() (*AFNORSearchSiretFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *AFNORSearchSiret) SetFilters(v AFNORSearchSiretFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *AFNORSearchSiret) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSorting

`func (o *AFNORSearchSiret) GetSorting() []AFNORSearchSiretSortingInner`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *AFNORSearchSiret) GetSortingOk() (*[]AFNORSearchSiretSortingInner, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *AFNORSearchSiret) SetSorting(v []AFNORSearchSiretSortingInner)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *AFNORSearchSiret) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### GetFields

`func (o *AFNORSearchSiret) GetFields() []AFNORSiretField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AFNORSearchSiret) GetFieldsOk() (*[]AFNORSiretField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AFNORSearchSiret) SetFields(v []AFNORSiretField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AFNORSearchSiret) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetInclude

`func (o *AFNORSearchSiret) GetInclude() []string`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *AFNORSearchSiret) GetIncludeOk() (*[]string, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *AFNORSearchSiret) SetInclude(v []string)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *AFNORSearchSiret) HasInclude() bool`

HasInclude returns a boolean if a field has been set.

### GetLimit

`func (o *AFNORSearchSiret) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AFNORSearchSiret) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AFNORSearchSiret) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AFNORSearchSiret) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetIgnore

`func (o *AFNORSearchSiret) GetIgnore() int32`

GetIgnore returns the Ignore field if non-nil, zero value otherwise.

### GetIgnoreOk

`func (o *AFNORSearchSiret) GetIgnoreOk() (*int32, bool)`

GetIgnoreOk returns a tuple with the Ignore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnore

`func (o *AFNORSearchSiret) SetIgnore(v int32)`

SetIgnore sets Ignore field to given value.

### HasIgnore

`func (o *AFNORSearchSiret) HasIgnore() bool`

HasIgnore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



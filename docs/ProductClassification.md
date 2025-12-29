# ProductClassification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassCode** | **string** | Classification code (BT-158). | 
**ListId** | Pointer to **NullableString** |  | [optional] 
**ListVersionId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProductClassification

`func NewProductClassification(classCode string, ) *ProductClassification`

NewProductClassification instantiates a new ProductClassification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductClassificationWithDefaults

`func NewProductClassificationWithDefaults() *ProductClassification`

NewProductClassificationWithDefaults instantiates a new ProductClassification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassCode

`func (o *ProductClassification) GetClassCode() string`

GetClassCode returns the ClassCode field if non-nil, zero value otherwise.

### GetClassCodeOk

`func (o *ProductClassification) GetClassCodeOk() (*string, bool)`

GetClassCodeOk returns a tuple with the ClassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassCode

`func (o *ProductClassification) SetClassCode(v string)`

SetClassCode sets ClassCode field to given value.


### GetListId

`func (o *ProductClassification) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *ProductClassification) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *ProductClassification) SetListId(v string)`

SetListId sets ListId field to given value.

### HasListId

`func (o *ProductClassification) HasListId() bool`

HasListId returns a boolean if a field has been set.

### SetListIdNil

`func (o *ProductClassification) SetListIdNil(b bool)`

 SetListIdNil sets the value for ListId to be an explicit nil

### UnsetListId
`func (o *ProductClassification) UnsetListId()`

UnsetListId ensures that no value is present for ListId, not even an explicit nil
### GetListVersionId

`func (o *ProductClassification) GetListVersionId() string`

GetListVersionId returns the ListVersionId field if non-nil, zero value otherwise.

### GetListVersionIdOk

`func (o *ProductClassification) GetListVersionIdOk() (*string, bool)`

GetListVersionIdOk returns a tuple with the ListVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListVersionId

`func (o *ProductClassification) SetListVersionId(v string)`

SetListVersionId sets ListVersionId field to given value.

### HasListVersionId

`func (o *ProductClassification) HasListVersionId() bool`

HasListVersionId returns a boolean if a field has been set.

### SetListVersionIdNil

`func (o *ProductClassification) SetListVersionIdNil(b bool)`

 SetListVersionIdNil sets the value for ListVersionId to be an explicit nil

### UnsetListVersionId
`func (o *ProductClassification) UnsetListVersionId()`

UnsetListVersionId ensures that no value is present for ListVersionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



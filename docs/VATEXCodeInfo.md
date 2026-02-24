# VATEXCodeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | VATEX code (e.g. VATEX-EU-IC) | 
**Name** | **string** | Short name (e.g. &#39;Intra-Community supply&#39;) | 
**Description** | **string** | Detailed description / remark | 
**Category** | **string** | Associated VAT category code (E, AE, K, G, O) | 

## Methods

### NewVATEXCodeInfo

`func NewVATEXCodeInfo(code string, name string, description string, category string, ) *VATEXCodeInfo`

NewVATEXCodeInfo instantiates a new VATEXCodeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVATEXCodeInfoWithDefaults

`func NewVATEXCodeInfoWithDefaults() *VATEXCodeInfo`

NewVATEXCodeInfoWithDefaults instantiates a new VATEXCodeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *VATEXCodeInfo) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VATEXCodeInfo) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VATEXCodeInfo) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *VATEXCodeInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VATEXCodeInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VATEXCodeInfo) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VATEXCodeInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VATEXCodeInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VATEXCodeInfo) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCategory

`func (o *VATEXCodeInfo) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *VATEXCodeInfo) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *VATEXCodeInfo) SetCategory(v string)`

SetCategory sets Category field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



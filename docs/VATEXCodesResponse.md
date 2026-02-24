# VATEXCodesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Codes** | [**[]VATEXCodeInfo**](VATEXCodeInfo.md) |  | 
**Count** | **int32** |  | 
**Source** | Pointer to **string** |  | [optional] [default to "https://docs.peppol.eu/poacc/billing/3.0/codelist/vatex/"]

## Methods

### NewVATEXCodesResponse

`func NewVATEXCodesResponse(codes []VATEXCodeInfo, count int32, ) *VATEXCodesResponse`

NewVATEXCodesResponse instantiates a new VATEXCodesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVATEXCodesResponseWithDefaults

`func NewVATEXCodesResponseWithDefaults() *VATEXCodesResponse`

NewVATEXCodesResponseWithDefaults instantiates a new VATEXCodesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodes

`func (o *VATEXCodesResponse) GetCodes() []VATEXCodeInfo`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *VATEXCodesResponse) GetCodesOk() (*[]VATEXCodeInfo, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *VATEXCodesResponse) SetCodes(v []VATEXCodeInfo)`

SetCodes sets Codes field to given value.


### GetCount

`func (o *VATEXCodesResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VATEXCodesResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VATEXCodesResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetSource

`func (o *VATEXCodesResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *VATEXCodesResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *VATEXCodesResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *VATEXCodesResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ActionCodesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Codes** | [**[]ActionCodeInfo**](ActionCodeInfo.md) | Liste des codes action | 
**Count** | **int32** | Nombre de codes | 
**Source** | Pointer to **string** | Règle source | [optional] [default to "BR-FR-CDV-CL-10"]

## Methods

### NewActionCodesResponse

`func NewActionCodesResponse(codes []ActionCodeInfo, count int32, ) *ActionCodesResponse`

NewActionCodesResponse instantiates a new ActionCodesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionCodesResponseWithDefaults

`func NewActionCodesResponseWithDefaults() *ActionCodesResponse`

NewActionCodesResponseWithDefaults instantiates a new ActionCodesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodes

`func (o *ActionCodesResponse) GetCodes() []ActionCodeInfo`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *ActionCodesResponse) GetCodesOk() (*[]ActionCodeInfo, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *ActionCodesResponse) SetCodes(v []ActionCodeInfo)`

SetCodes sets Codes field to given value.


### GetCount

`func (o *ActionCodesResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ActionCodesResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ActionCodesResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetSource

`func (o *ActionCodesResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ActionCodesResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ActionCodesResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ActionCodesResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



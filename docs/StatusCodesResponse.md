# StatusCodesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Codes** | [**[]StatusCodeInfo**](StatusCodeInfo.md) | Liste des codes statut | 
**Count** | **int32** | Nombre de codes | 
**Source** | Pointer to **string** | Règle source | [optional] [default to "BR-FR-CDV-CL-06"]

## Methods

### NewStatusCodesResponse

`func NewStatusCodesResponse(codes []StatusCodeInfo, count int32, ) *StatusCodesResponse`

NewStatusCodesResponse instantiates a new StatusCodesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusCodesResponseWithDefaults

`func NewStatusCodesResponseWithDefaults() *StatusCodesResponse`

NewStatusCodesResponseWithDefaults instantiates a new StatusCodesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodes

`func (o *StatusCodesResponse) GetCodes() []StatusCodeInfo`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *StatusCodesResponse) GetCodesOk() (*[]StatusCodeInfo, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *StatusCodesResponse) SetCodes(v []StatusCodeInfo)`

SetCodes sets Codes field to given value.


### GetCount

`func (o *StatusCodesResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *StatusCodesResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *StatusCodesResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetSource

`func (o *StatusCodesResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StatusCodesResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StatusCodesResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *StatusCodesResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



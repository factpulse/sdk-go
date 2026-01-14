# FacturXPDFInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **int32** | PDF size in bytes | 
**Profile** | **string** | Factur-X profile used | 
**Signed** | Pointer to **bool** | PDF electronically signed | [optional] [default to false]

## Methods

### NewFacturXPDFInfo

`func NewFacturXPDFInfo(size int32, profile string, ) *FacturXPDFInfo`

NewFacturXPDFInfo instantiates a new FacturXPDFInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacturXPDFInfoWithDefaults

`func NewFacturXPDFInfoWithDefaults() *FacturXPDFInfo`

NewFacturXPDFInfoWithDefaults instantiates a new FacturXPDFInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *FacturXPDFInfo) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FacturXPDFInfo) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FacturXPDFInfo) SetSize(v int32)`

SetSize sets Size field to given value.


### GetProfile

`func (o *FacturXPDFInfo) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *FacturXPDFInfo) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *FacturXPDFInfo) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetSigned

`func (o *FacturXPDFInfo) GetSigned() bool`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *FacturXPDFInfo) GetSignedOk() (*bool, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *FacturXPDFInfo) SetSigned(v bool)`

SetSigned sets Signed field to given value.

### HasSigned

`func (o *FacturXPDFInfo) HasSigned() bool`

HasSigned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



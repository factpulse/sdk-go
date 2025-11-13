# PDFFacturXInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Taille** | **int32** | Taille du PDF en octets | 
**Profil** | **string** | Profil Factur-X utilisé | 
**Signe** | Pointer to **bool** | PDF signé électroniquement | [optional] [default to false]

## Methods

### NewPDFFacturXInfo

`func NewPDFFacturXInfo(taille int32, profil string, ) *PDFFacturXInfo`

NewPDFFacturXInfo instantiates a new PDFFacturXInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPDFFacturXInfoWithDefaults

`func NewPDFFacturXInfoWithDefaults() *PDFFacturXInfo`

NewPDFFacturXInfoWithDefaults instantiates a new PDFFacturXInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaille

`func (o *PDFFacturXInfo) GetTaille() int32`

GetTaille returns the Taille field if non-nil, zero value otherwise.

### GetTailleOk

`func (o *PDFFacturXInfo) GetTailleOk() (*int32, bool)`

GetTailleOk returns a tuple with the Taille field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaille

`func (o *PDFFacturXInfo) SetTaille(v int32)`

SetTaille sets Taille field to given value.


### GetProfil

`func (o *PDFFacturXInfo) GetProfil() string`

GetProfil returns the Profil field if non-nil, zero value otherwise.

### GetProfilOk

`func (o *PDFFacturXInfo) GetProfilOk() (*string, bool)`

GetProfilOk returns a tuple with the Profil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfil

`func (o *PDFFacturXInfo) SetProfil(v string)`

SetProfil sets Profil field to given value.


### GetSigne

`func (o *PDFFacturXInfo) GetSigne() bool`

GetSigne returns the Signe field if non-nil, zero value otherwise.

### GetSigneOk

`func (o *PDFFacturXInfo) GetSigneOk() (*bool, bool)`

GetSigneOk returns a tuple with the Signe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigne

`func (o *PDFFacturXInfo) SetSigne(v bool)`

SetSigne sets Signe field to given value.

### HasSigne

`func (o *PDFFacturXInfo) HasSigne() bool`

HasSigne returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



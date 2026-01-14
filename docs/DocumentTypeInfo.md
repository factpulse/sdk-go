# DocumentTypeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | Code UNTDID 1001 | 
**Label** | **string** | Libelle (Facture, Avoir, etc.) | 
**DetectedAs** | **string** | Classification interne | 

## Methods

### NewDocumentTypeInfo

`func NewDocumentTypeInfo(code int32, label string, detectedAs string, ) *DocumentTypeInfo`

NewDocumentTypeInfo instantiates a new DocumentTypeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentTypeInfoWithDefaults

`func NewDocumentTypeInfoWithDefaults() *DocumentTypeInfo`

NewDocumentTypeInfoWithDefaults instantiates a new DocumentTypeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *DocumentTypeInfo) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DocumentTypeInfo) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DocumentTypeInfo) SetCode(v int32)`

SetCode sets Code field to given value.


### GetLabel

`func (o *DocumentTypeInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DocumentTypeInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DocumentTypeInfo) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDetectedAs

`func (o *DocumentTypeInfo) GetDetectedAs() string`

GetDetectedAs returns the DetectedAs field if non-nil, zero value otherwise.

### GetDetectedAsOk

`func (o *DocumentTypeInfo) GetDetectedAsOk() (*string, bool)`

GetDetectedAsOk returns a tuple with the DetectedAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedAs

`func (o *DocumentTypeInfo) SetDetectedAs(v string)`

SetDetectedAs sets DetectedAs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



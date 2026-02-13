# GenerationSuccessResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "SUCCESS"]
**ContentB64** | Pointer to **NullableString** |  | [optional] 
**XmlContent** | Pointer to **NullableString** |  | [optional] 
**Filename** | **string** |  | 

## Methods

### NewGenerationSuccessResult

`func NewGenerationSuccessResult(filename string, ) *GenerationSuccessResult`

NewGenerationSuccessResult instantiates a new GenerationSuccessResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerationSuccessResultWithDefaults

`func NewGenerationSuccessResultWithDefaults() *GenerationSuccessResult`

NewGenerationSuccessResultWithDefaults instantiates a new GenerationSuccessResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GenerationSuccessResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GenerationSuccessResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GenerationSuccessResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GenerationSuccessResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetContentB64

`func (o *GenerationSuccessResult) GetContentB64() string`

GetContentB64 returns the ContentB64 field if non-nil, zero value otherwise.

### GetContentB64Ok

`func (o *GenerationSuccessResult) GetContentB64Ok() (*string, bool)`

GetContentB64Ok returns a tuple with the ContentB64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentB64

`func (o *GenerationSuccessResult) SetContentB64(v string)`

SetContentB64 sets ContentB64 field to given value.

### HasContentB64

`func (o *GenerationSuccessResult) HasContentB64() bool`

HasContentB64 returns a boolean if a field has been set.

### SetContentB64Nil

`func (o *GenerationSuccessResult) SetContentB64Nil(b bool)`

 SetContentB64Nil sets the value for ContentB64 to be an explicit nil

### UnsetContentB64
`func (o *GenerationSuccessResult) UnsetContentB64()`

UnsetContentB64 ensures that no value is present for ContentB64, not even an explicit nil
### GetXmlContent

`func (o *GenerationSuccessResult) GetXmlContent() string`

GetXmlContent returns the XmlContent field if non-nil, zero value otherwise.

### GetXmlContentOk

`func (o *GenerationSuccessResult) GetXmlContentOk() (*string, bool)`

GetXmlContentOk returns a tuple with the XmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlContent

`func (o *GenerationSuccessResult) SetXmlContent(v string)`

SetXmlContent sets XmlContent field to given value.

### HasXmlContent

`func (o *GenerationSuccessResult) HasXmlContent() bool`

HasXmlContent returns a boolean if a field has been set.

### SetXmlContentNil

`func (o *GenerationSuccessResult) SetXmlContentNil(b bool)`

 SetXmlContentNil sets the value for XmlContent to be an explicit nil

### UnsetXmlContent
`func (o *GenerationSuccessResult) UnsetXmlContent()`

UnsetXmlContent ensures that no value is present for XmlContent, not even an explicit nil
### GetFilename

`func (o *GenerationSuccessResult) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *GenerationSuccessResult) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *GenerationSuccessResult) SetFilename(v string)`

SetFilename sets Filename field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



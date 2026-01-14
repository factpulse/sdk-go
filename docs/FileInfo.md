# FileInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentB64** | Pointer to **NullableString** |  | [optional] 
**SizeBytes** | Pointer to **int32** | Taille en bytes | [optional] [default to 0]

## Methods

### NewFileInfo

`func NewFileInfo() *FileInfo`

NewFileInfo instantiates a new FileInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileInfoWithDefaults

`func NewFileInfoWithDefaults() *FileInfo`

NewFileInfoWithDefaults instantiates a new FileInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentB64

`func (o *FileInfo) GetContentB64() string`

GetContentB64 returns the ContentB64 field if non-nil, zero value otherwise.

### GetContentB64Ok

`func (o *FileInfo) GetContentB64Ok() (*string, bool)`

GetContentB64Ok returns a tuple with the ContentB64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentB64

`func (o *FileInfo) SetContentB64(v string)`

SetContentB64 sets ContentB64 field to given value.

### HasContentB64

`func (o *FileInfo) HasContentB64() bool`

HasContentB64 returns a boolean if a field has been set.

### SetContentB64Nil

`func (o *FileInfo) SetContentB64Nil(b bool)`

 SetContentB64Nil sets the value for ContentB64 to be an explicit nil

### UnsetContentB64
`func (o *FileInfo) UnsetContentB64()`

UnsetContentB64 ensures that no value is present for ContentB64, not even an explicit nil
### GetSizeBytes

`func (o *FileInfo) GetSizeBytes() int32`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *FileInfo) GetSizeBytesOk() (*int32, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *FileInfo) SetSizeBytes(v int32)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *FileInfo) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



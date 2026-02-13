# GenerationTaskStatusResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "ERROR"]
**ContentB64** | Pointer to **string** |  | [optional] 
**XmlContent** | Pointer to **string** |  | [optional] 
**Filename** | **string** |  | 
**ErrorCode** | **string** |  | 
**ErrorMessage** | **string** |  | 
**Details** | Pointer to [**[]AFNORErrorDetail**](AFNORErrorDetail.md) |  | [optional] 
**Traceback** | Pointer to **string** |  | [optional] 

## Methods

### NewGenerationTaskStatusResult

`func NewGenerationTaskStatusResult(filename string, errorCode string, errorMessage string, ) *GenerationTaskStatusResult`

NewGenerationTaskStatusResult instantiates a new GenerationTaskStatusResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerationTaskStatusResultWithDefaults

`func NewGenerationTaskStatusResultWithDefaults() *GenerationTaskStatusResult`

NewGenerationTaskStatusResultWithDefaults instantiates a new GenerationTaskStatusResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GenerationTaskStatusResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GenerationTaskStatusResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GenerationTaskStatusResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GenerationTaskStatusResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetContentB64

`func (o *GenerationTaskStatusResult) GetContentB64() string`

GetContentB64 returns the ContentB64 field if non-nil, zero value otherwise.

### GetContentB64Ok

`func (o *GenerationTaskStatusResult) GetContentB64Ok() (*string, bool)`

GetContentB64Ok returns a tuple with the ContentB64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentB64

`func (o *GenerationTaskStatusResult) SetContentB64(v string)`

SetContentB64 sets ContentB64 field to given value.

### HasContentB64

`func (o *GenerationTaskStatusResult) HasContentB64() bool`

HasContentB64 returns a boolean if a field has been set.

### GetXmlContent

`func (o *GenerationTaskStatusResult) GetXmlContent() string`

GetXmlContent returns the XmlContent field if non-nil, zero value otherwise.

### GetXmlContentOk

`func (o *GenerationTaskStatusResult) GetXmlContentOk() (*string, bool)`

GetXmlContentOk returns a tuple with the XmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlContent

`func (o *GenerationTaskStatusResult) SetXmlContent(v string)`

SetXmlContent sets XmlContent field to given value.

### HasXmlContent

`func (o *GenerationTaskStatusResult) HasXmlContent() bool`

HasXmlContent returns a boolean if a field has been set.

### GetFilename

`func (o *GenerationTaskStatusResult) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *GenerationTaskStatusResult) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *GenerationTaskStatusResult) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetErrorCode

`func (o *GenerationTaskStatusResult) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GenerationTaskStatusResult) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GenerationTaskStatusResult) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorMessage

`func (o *GenerationTaskStatusResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GenerationTaskStatusResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GenerationTaskStatusResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### GetDetails

`func (o *GenerationTaskStatusResult) GetDetails() []AFNORErrorDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GenerationTaskStatusResult) GetDetailsOk() (*[]AFNORErrorDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GenerationTaskStatusResult) SetDetails(v []AFNORErrorDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GenerationTaskStatusResult) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTraceback

`func (o *GenerationTaskStatusResult) GetTraceback() string`

GetTraceback returns the Traceback field if non-nil, zero value otherwise.

### GetTracebackOk

`func (o *GenerationTaskStatusResult) GetTracebackOk() (*string, bool)`

GetTracebackOk returns a tuple with the Traceback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceback

`func (o *GenerationTaskStatusResult) SetTraceback(v string)`

SetTraceback sets Traceback field to given value.

### HasTraceback

`func (o *GenerationTaskStatusResult) HasTraceback() bool`

HasTraceback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



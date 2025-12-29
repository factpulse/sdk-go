# ConvertErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "error"]
**Error** | **string** | Code erreur | 
**Message** | **string** | Message d&#39;erreur | 
**ConversionId** | Pointer to **NullableString** |  | [optional] 
**ResumeUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewConvertErrorResponse

`func NewConvertErrorResponse(error_ string, message string, ) *ConvertErrorResponse`

NewConvertErrorResponse instantiates a new ConvertErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertErrorResponseWithDefaults

`func NewConvertErrorResponseWithDefaults() *ConvertErrorResponse`

NewConvertErrorResponseWithDefaults instantiates a new ConvertErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ConvertErrorResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConvertErrorResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConvertErrorResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConvertErrorResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetError

`func (o *ConvertErrorResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ConvertErrorResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ConvertErrorResponse) SetError(v string)`

SetError sets Error field to given value.


### GetMessage

`func (o *ConvertErrorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConvertErrorResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConvertErrorResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetConversionId

`func (o *ConvertErrorResponse) GetConversionId() string`

GetConversionId returns the ConversionId field if non-nil, zero value otherwise.

### GetConversionIdOk

`func (o *ConvertErrorResponse) GetConversionIdOk() (*string, bool)`

GetConversionIdOk returns a tuple with the ConversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionId

`func (o *ConvertErrorResponse) SetConversionId(v string)`

SetConversionId sets ConversionId field to given value.

### HasConversionId

`func (o *ConvertErrorResponse) HasConversionId() bool`

HasConversionId returns a boolean if a field has been set.

### SetConversionIdNil

`func (o *ConvertErrorResponse) SetConversionIdNil(b bool)`

 SetConversionIdNil sets the value for ConversionId to be an explicit nil

### UnsetConversionId
`func (o *ConvertErrorResponse) UnsetConversionId()`

UnsetConversionId ensures that no value is present for ConversionId, not even an explicit nil
### GetResumeUrl

`func (o *ConvertErrorResponse) GetResumeUrl() string`

GetResumeUrl returns the ResumeUrl field if non-nil, zero value otherwise.

### GetResumeUrlOk

`func (o *ConvertErrorResponse) GetResumeUrlOk() (*string, bool)`

GetResumeUrlOk returns a tuple with the ResumeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeUrl

`func (o *ConvertErrorResponse) SetResumeUrl(v string)`

SetResumeUrl sets ResumeUrl field to given value.

### HasResumeUrl

`func (o *ConvertErrorResponse) HasResumeUrl() bool`

HasResumeUrl returns a boolean if a field has been set.

### SetResumeUrlNil

`func (o *ConvertErrorResponse) SetResumeUrlNil(b bool)`

 SetResumeUrlNil sets the value for ResumeUrl to be an explicit nil

### UnsetResumeUrl
`func (o *ConvertErrorResponse) UnsetResumeUrl()`

UnsetResumeUrl ensures that no value is present for ResumeUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



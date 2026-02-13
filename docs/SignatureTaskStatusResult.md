# SignatureTaskStatusResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "ERROR"]
**ContentB64** | **string** |  | 
**Filename** | **string** |  | 
**SizeBytes** | **int32** |  | 
**IsSigned** | **bool** |  | 
**SignatureCount** | **int32** |  | 
**SignerCn** | **string** |  | 
**SigningDate** | **string** |  | 
**ValidationDetails** | **map[string]interface{}** |  | 
**ErrorCode** | **string** |  | 
**ErrorMessage** | **string** |  | 
**Details** | Pointer to [**[]AFNORErrorDetail**](AFNORErrorDetail.md) |  | [optional] 
**Traceback** | Pointer to **string** |  | [optional] 

## Methods

### NewSignatureTaskStatusResult

`func NewSignatureTaskStatusResult(contentB64 string, filename string, sizeBytes int32, isSigned bool, signatureCount int32, signerCn string, signingDate string, validationDetails map[string]interface{}, errorCode string, errorMessage string, ) *SignatureTaskStatusResult`

NewSignatureTaskStatusResult instantiates a new SignatureTaskStatusResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureTaskStatusResultWithDefaults

`func NewSignatureTaskStatusResultWithDefaults() *SignatureTaskStatusResult`

NewSignatureTaskStatusResultWithDefaults instantiates a new SignatureTaskStatusResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SignatureTaskStatusResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SignatureTaskStatusResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SignatureTaskStatusResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SignatureTaskStatusResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetContentB64

`func (o *SignatureTaskStatusResult) GetContentB64() string`

GetContentB64 returns the ContentB64 field if non-nil, zero value otherwise.

### GetContentB64Ok

`func (o *SignatureTaskStatusResult) GetContentB64Ok() (*string, bool)`

GetContentB64Ok returns a tuple with the ContentB64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentB64

`func (o *SignatureTaskStatusResult) SetContentB64(v string)`

SetContentB64 sets ContentB64 field to given value.


### GetFilename

`func (o *SignatureTaskStatusResult) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *SignatureTaskStatusResult) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *SignatureTaskStatusResult) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetSizeBytes

`func (o *SignatureTaskStatusResult) GetSizeBytes() int32`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *SignatureTaskStatusResult) GetSizeBytesOk() (*int32, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *SignatureTaskStatusResult) SetSizeBytes(v int32)`

SetSizeBytes sets SizeBytes field to given value.


### GetIsSigned

`func (o *SignatureTaskStatusResult) GetIsSigned() bool`

GetIsSigned returns the IsSigned field if non-nil, zero value otherwise.

### GetIsSignedOk

`func (o *SignatureTaskStatusResult) GetIsSignedOk() (*bool, bool)`

GetIsSignedOk returns a tuple with the IsSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSigned

`func (o *SignatureTaskStatusResult) SetIsSigned(v bool)`

SetIsSigned sets IsSigned field to given value.


### GetSignatureCount

`func (o *SignatureTaskStatusResult) GetSignatureCount() int32`

GetSignatureCount returns the SignatureCount field if non-nil, zero value otherwise.

### GetSignatureCountOk

`func (o *SignatureTaskStatusResult) GetSignatureCountOk() (*int32, bool)`

GetSignatureCountOk returns a tuple with the SignatureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureCount

`func (o *SignatureTaskStatusResult) SetSignatureCount(v int32)`

SetSignatureCount sets SignatureCount field to given value.


### GetSignerCn

`func (o *SignatureTaskStatusResult) GetSignerCn() string`

GetSignerCn returns the SignerCn field if non-nil, zero value otherwise.

### GetSignerCnOk

`func (o *SignatureTaskStatusResult) GetSignerCnOk() (*string, bool)`

GetSignerCnOk returns a tuple with the SignerCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerCn

`func (o *SignatureTaskStatusResult) SetSignerCn(v string)`

SetSignerCn sets SignerCn field to given value.


### GetSigningDate

`func (o *SignatureTaskStatusResult) GetSigningDate() string`

GetSigningDate returns the SigningDate field if non-nil, zero value otherwise.

### GetSigningDateOk

`func (o *SignatureTaskStatusResult) GetSigningDateOk() (*string, bool)`

GetSigningDateOk returns a tuple with the SigningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningDate

`func (o *SignatureTaskStatusResult) SetSigningDate(v string)`

SetSigningDate sets SigningDate field to given value.


### GetValidationDetails

`func (o *SignatureTaskStatusResult) GetValidationDetails() map[string]interface{}`

GetValidationDetails returns the ValidationDetails field if non-nil, zero value otherwise.

### GetValidationDetailsOk

`func (o *SignatureTaskStatusResult) GetValidationDetailsOk() (*map[string]interface{}, bool)`

GetValidationDetailsOk returns a tuple with the ValidationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationDetails

`func (o *SignatureTaskStatusResult) SetValidationDetails(v map[string]interface{})`

SetValidationDetails sets ValidationDetails field to given value.


### GetErrorCode

`func (o *SignatureTaskStatusResult) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *SignatureTaskStatusResult) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *SignatureTaskStatusResult) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorMessage

`func (o *SignatureTaskStatusResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *SignatureTaskStatusResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *SignatureTaskStatusResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### GetDetails

`func (o *SignatureTaskStatusResult) GetDetails() []AFNORErrorDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SignatureTaskStatusResult) GetDetailsOk() (*[]AFNORErrorDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SignatureTaskStatusResult) SetDetails(v []AFNORErrorDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *SignatureTaskStatusResult) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTraceback

`func (o *SignatureTaskStatusResult) GetTraceback() string`

GetTraceback returns the Traceback field if non-nil, zero value otherwise.

### GetTracebackOk

`func (o *SignatureTaskStatusResult) GetTracebackOk() (*string, bool)`

GetTracebackOk returns a tuple with the Traceback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceback

`func (o *SignatureTaskStatusResult) SetTraceback(v string)`

SetTraceback sets Traceback field to given value.

### HasTraceback

`func (o *SignatureTaskStatusResult) HasTraceback() bool`

HasTraceback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



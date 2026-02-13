# SignatureSuccessResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "SUCCESS"]
**ContentB64** | **string** |  | 
**Filename** | **string** |  | 
**SizeBytes** | **int32** |  | 
**IsSigned** | **bool** |  | 
**SignatureCount** | **int32** |  | 
**SignerCn** | **string** |  | 
**SigningDate** | **string** |  | 
**ValidationDetails** | **map[string]interface{}** |  | 

## Methods

### NewSignatureSuccessResult

`func NewSignatureSuccessResult(contentB64 string, filename string, sizeBytes int32, isSigned bool, signatureCount int32, signerCn string, signingDate string, validationDetails map[string]interface{}, ) *SignatureSuccessResult`

NewSignatureSuccessResult instantiates a new SignatureSuccessResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureSuccessResultWithDefaults

`func NewSignatureSuccessResultWithDefaults() *SignatureSuccessResult`

NewSignatureSuccessResultWithDefaults instantiates a new SignatureSuccessResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SignatureSuccessResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SignatureSuccessResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SignatureSuccessResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SignatureSuccessResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetContentB64

`func (o *SignatureSuccessResult) GetContentB64() string`

GetContentB64 returns the ContentB64 field if non-nil, zero value otherwise.

### GetContentB64Ok

`func (o *SignatureSuccessResult) GetContentB64Ok() (*string, bool)`

GetContentB64Ok returns a tuple with the ContentB64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentB64

`func (o *SignatureSuccessResult) SetContentB64(v string)`

SetContentB64 sets ContentB64 field to given value.


### GetFilename

`func (o *SignatureSuccessResult) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *SignatureSuccessResult) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *SignatureSuccessResult) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetSizeBytes

`func (o *SignatureSuccessResult) GetSizeBytes() int32`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *SignatureSuccessResult) GetSizeBytesOk() (*int32, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *SignatureSuccessResult) SetSizeBytes(v int32)`

SetSizeBytes sets SizeBytes field to given value.


### GetIsSigned

`func (o *SignatureSuccessResult) GetIsSigned() bool`

GetIsSigned returns the IsSigned field if non-nil, zero value otherwise.

### GetIsSignedOk

`func (o *SignatureSuccessResult) GetIsSignedOk() (*bool, bool)`

GetIsSignedOk returns a tuple with the IsSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSigned

`func (o *SignatureSuccessResult) SetIsSigned(v bool)`

SetIsSigned sets IsSigned field to given value.


### GetSignatureCount

`func (o *SignatureSuccessResult) GetSignatureCount() int32`

GetSignatureCount returns the SignatureCount field if non-nil, zero value otherwise.

### GetSignatureCountOk

`func (o *SignatureSuccessResult) GetSignatureCountOk() (*int32, bool)`

GetSignatureCountOk returns a tuple with the SignatureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureCount

`func (o *SignatureSuccessResult) SetSignatureCount(v int32)`

SetSignatureCount sets SignatureCount field to given value.


### GetSignerCn

`func (o *SignatureSuccessResult) GetSignerCn() string`

GetSignerCn returns the SignerCn field if non-nil, zero value otherwise.

### GetSignerCnOk

`func (o *SignatureSuccessResult) GetSignerCnOk() (*string, bool)`

GetSignerCnOk returns a tuple with the SignerCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerCn

`func (o *SignatureSuccessResult) SetSignerCn(v string)`

SetSignerCn sets SignerCn field to given value.


### GetSigningDate

`func (o *SignatureSuccessResult) GetSigningDate() string`

GetSigningDate returns the SigningDate field if non-nil, zero value otherwise.

### GetSigningDateOk

`func (o *SignatureSuccessResult) GetSigningDateOk() (*string, bool)`

GetSigningDateOk returns a tuple with the SigningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningDate

`func (o *SignatureSuccessResult) SetSigningDate(v string)`

SetSigningDate sets SigningDate field to given value.


### GetValidationDetails

`func (o *SignatureSuccessResult) GetValidationDetails() map[string]interface{}`

GetValidationDetails returns the ValidationDetails field if non-nil, zero value otherwise.

### GetValidationDetailsOk

`func (o *SignatureSuccessResult) GetValidationDetailsOk() (*map[string]interface{}, bool)`

GetValidationDetailsOk returns a tuple with the ValidationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationDetails

`func (o *SignatureSuccessResult) SetValidationDetails(v map[string]interface{})`

SetValidationDetails sets ValidationDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



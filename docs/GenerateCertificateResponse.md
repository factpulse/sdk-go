# GenerateCertificateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Operation status | [optional] [default to "success"]
**CertificatePem** | **string** | X.509 certificate in PEM format | 
**PrivateKeyPem** | **string** | RSA private key in PEM format | 
**Pkcs12Base64** | Pointer to **NullableString** |  | [optional] 
**Info** | [**CertificateInfoResponse**](CertificateInfoResponse.md) | Generated certificate information | 
**Warning** | Pointer to **string** | Warning about certificate usage | [optional] [default to "WARNING: This certificate is SELF-SIGNED and intended for TESTING only. DO NOT use in production. eIDAS level: SES (Simple Electronic Signature)"]

## Methods

### NewGenerateCertificateResponse

`func NewGenerateCertificateResponse(certificatePem string, privateKeyPem string, info CertificateInfoResponse, ) *GenerateCertificateResponse`

NewGenerateCertificateResponse instantiates a new GenerateCertificateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateCertificateResponseWithDefaults

`func NewGenerateCertificateResponseWithDefaults() *GenerateCertificateResponse`

NewGenerateCertificateResponseWithDefaults instantiates a new GenerateCertificateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GenerateCertificateResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GenerateCertificateResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GenerateCertificateResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GenerateCertificateResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCertificatePem

`func (o *GenerateCertificateResponse) GetCertificatePem() string`

GetCertificatePem returns the CertificatePem field if non-nil, zero value otherwise.

### GetCertificatePemOk

`func (o *GenerateCertificateResponse) GetCertificatePemOk() (*string, bool)`

GetCertificatePemOk returns a tuple with the CertificatePem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePem

`func (o *GenerateCertificateResponse) SetCertificatePem(v string)`

SetCertificatePem sets CertificatePem field to given value.


### GetPrivateKeyPem

`func (o *GenerateCertificateResponse) GetPrivateKeyPem() string`

GetPrivateKeyPem returns the PrivateKeyPem field if non-nil, zero value otherwise.

### GetPrivateKeyPemOk

`func (o *GenerateCertificateResponse) GetPrivateKeyPemOk() (*string, bool)`

GetPrivateKeyPemOk returns a tuple with the PrivateKeyPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPem

`func (o *GenerateCertificateResponse) SetPrivateKeyPem(v string)`

SetPrivateKeyPem sets PrivateKeyPem field to given value.


### GetPkcs12Base64

`func (o *GenerateCertificateResponse) GetPkcs12Base64() string`

GetPkcs12Base64 returns the Pkcs12Base64 field if non-nil, zero value otherwise.

### GetPkcs12Base64Ok

`func (o *GenerateCertificateResponse) GetPkcs12Base64Ok() (*string, bool)`

GetPkcs12Base64Ok returns a tuple with the Pkcs12Base64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs12Base64

`func (o *GenerateCertificateResponse) SetPkcs12Base64(v string)`

SetPkcs12Base64 sets Pkcs12Base64 field to given value.

### HasPkcs12Base64

`func (o *GenerateCertificateResponse) HasPkcs12Base64() bool`

HasPkcs12Base64 returns a boolean if a field has been set.

### SetPkcs12Base64Nil

`func (o *GenerateCertificateResponse) SetPkcs12Base64Nil(b bool)`

 SetPkcs12Base64Nil sets the value for Pkcs12Base64 to be an explicit nil

### UnsetPkcs12Base64
`func (o *GenerateCertificateResponse) UnsetPkcs12Base64()`

UnsetPkcs12Base64 ensures that no value is present for Pkcs12Base64, not even an explicit nil
### GetInfo

`func (o *GenerateCertificateResponse) GetInfo() CertificateInfoResponse`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *GenerateCertificateResponse) GetInfoOk() (*CertificateInfoResponse, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *GenerateCertificateResponse) SetInfo(v CertificateInfoResponse)`

SetInfo sets Info field to given value.


### GetWarning

`func (o *GenerateCertificateResponse) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *GenerateCertificateResponse) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *GenerateCertificateResponse) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *GenerateCertificateResponse) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



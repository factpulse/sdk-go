# GenerateCertificateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Statut de l&#39;opération | [optional] [default to "success"]
**CertificatPem** | **string** | Certificat X.509 au format PEM | 
**ClePriveePem** | **string** | Clé privée RSA au format PEM | 
**Pkcs12Base64** | Pointer to **NullableString** |  | [optional] 
**Info** | [**CertificateInfoResponse**](CertificateInfoResponse.md) | Informations sur le certificat généré | 
**Avertissement** | Pointer to **string** | Avertissement sur l&#39;utilisation du certificat | [optional] [default to "⚠️ Ce certificat est AUTO-SIGNÉ et destiné uniquement aux TESTS. Ne PAS utiliser en production. Niveau eIDAS : SES (Simple Electronic Signature)"]

## Methods

### NewGenerateCertificateResponse

`func NewGenerateCertificateResponse(certificatPem string, clePriveePem string, info CertificateInfoResponse, ) *GenerateCertificateResponse`

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

### GetCertificatPem

`func (o *GenerateCertificateResponse) GetCertificatPem() string`

GetCertificatPem returns the CertificatPem field if non-nil, zero value otherwise.

### GetCertificatPemOk

`func (o *GenerateCertificateResponse) GetCertificatPemOk() (*string, bool)`

GetCertificatPemOk returns a tuple with the CertificatPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatPem

`func (o *GenerateCertificateResponse) SetCertificatPem(v string)`

SetCertificatPem sets CertificatPem field to given value.


### GetClePriveePem

`func (o *GenerateCertificateResponse) GetClePriveePem() string`

GetClePriveePem returns the ClePriveePem field if non-nil, zero value otherwise.

### GetClePriveePemOk

`func (o *GenerateCertificateResponse) GetClePriveePemOk() (*string, bool)`

GetClePriveePemOk returns a tuple with the ClePriveePem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClePriveePem

`func (o *GenerateCertificateResponse) SetClePriveePem(v string)`

SetClePriveePem sets ClePriveePem field to given value.


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


### GetAvertissement

`func (o *GenerateCertificateResponse) GetAvertissement() string`

GetAvertissement returns the Avertissement field if non-nil, zero value otherwise.

### GetAvertissementOk

`func (o *GenerateCertificateResponse) GetAvertissementOk() (*string, bool)`

GetAvertissementOk returns a tuple with the Avertissement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvertissement

`func (o *GenerateCertificateResponse) SetAvertissement(v string)`

SetAvertissement sets Avertissement field to given value.

### HasAvertissement

`func (o *GenerateCertificateResponse) HasAvertissement() bool`

HasAvertissement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



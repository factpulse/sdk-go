# GenerateCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cn** | Pointer to **string** | Common Name (CN) - Certificate name | [optional] [default to "Test Signature FactPulse"]
**Organization** | Pointer to **string** | Organization (O) | [optional] [default to "FactPulse Test"]
**Country** | Pointer to **string** | ISO 2-letter country code (C) | [optional] [default to "FR"]
**City** | Pointer to **string** | City (L) | [optional] [default to "Paris"]
**State** | Pointer to **string** | State/Province (ST) | [optional] [default to "Ile-de-France"]
**Email** | Pointer to **NullableString** |  | [optional] 
**ValidityDays** | Pointer to **int32** | Validity duration in days | [optional] [default to 365]
**KeySize** | Pointer to **int32** | RSA key size in bits | [optional] [default to 2048]
**KeyPassphrase** | Pointer to **NullableString** |  | [optional] 
**GenerateP12** | Pointer to **bool** | Also generate a PKCS#12 (.p12) file | [optional] [default to false]
**P12Passphrase** | Pointer to **string** | Passphrase for PKCS#12 file | [optional] [default to "changeme"]

## Methods

### NewGenerateCertificateRequest

`func NewGenerateCertificateRequest() *GenerateCertificateRequest`

NewGenerateCertificateRequest instantiates a new GenerateCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateCertificateRequestWithDefaults

`func NewGenerateCertificateRequestWithDefaults() *GenerateCertificateRequest`

NewGenerateCertificateRequestWithDefaults instantiates a new GenerateCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCn

`func (o *GenerateCertificateRequest) GetCn() string`

GetCn returns the Cn field if non-nil, zero value otherwise.

### GetCnOk

`func (o *GenerateCertificateRequest) GetCnOk() (*string, bool)`

GetCnOk returns a tuple with the Cn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCn

`func (o *GenerateCertificateRequest) SetCn(v string)`

SetCn sets Cn field to given value.

### HasCn

`func (o *GenerateCertificateRequest) HasCn() bool`

HasCn returns a boolean if a field has been set.

### GetOrganization

`func (o *GenerateCertificateRequest) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *GenerateCertificateRequest) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *GenerateCertificateRequest) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *GenerateCertificateRequest) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetCountry

`func (o *GenerateCertificateRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GenerateCertificateRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GenerateCertificateRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GenerateCertificateRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCity

`func (o *GenerateCertificateRequest) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GenerateCertificateRequest) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GenerateCertificateRequest) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *GenerateCertificateRequest) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *GenerateCertificateRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GenerateCertificateRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GenerateCertificateRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GenerateCertificateRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetEmail

`func (o *GenerateCertificateRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GenerateCertificateRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GenerateCertificateRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GenerateCertificateRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *GenerateCertificateRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *GenerateCertificateRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetValidityDays

`func (o *GenerateCertificateRequest) GetValidityDays() int32`

GetValidityDays returns the ValidityDays field if non-nil, zero value otherwise.

### GetValidityDaysOk

`func (o *GenerateCertificateRequest) GetValidityDaysOk() (*int32, bool)`

GetValidityDaysOk returns a tuple with the ValidityDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityDays

`func (o *GenerateCertificateRequest) SetValidityDays(v int32)`

SetValidityDays sets ValidityDays field to given value.

### HasValidityDays

`func (o *GenerateCertificateRequest) HasValidityDays() bool`

HasValidityDays returns a boolean if a field has been set.

### GetKeySize

`func (o *GenerateCertificateRequest) GetKeySize() int32`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *GenerateCertificateRequest) GetKeySizeOk() (*int32, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *GenerateCertificateRequest) SetKeySize(v int32)`

SetKeySize sets KeySize field to given value.

### HasKeySize

`func (o *GenerateCertificateRequest) HasKeySize() bool`

HasKeySize returns a boolean if a field has been set.

### GetKeyPassphrase

`func (o *GenerateCertificateRequest) GetKeyPassphrase() string`

GetKeyPassphrase returns the KeyPassphrase field if non-nil, zero value otherwise.

### GetKeyPassphraseOk

`func (o *GenerateCertificateRequest) GetKeyPassphraseOk() (*string, bool)`

GetKeyPassphraseOk returns a tuple with the KeyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPassphrase

`func (o *GenerateCertificateRequest) SetKeyPassphrase(v string)`

SetKeyPassphrase sets KeyPassphrase field to given value.

### HasKeyPassphrase

`func (o *GenerateCertificateRequest) HasKeyPassphrase() bool`

HasKeyPassphrase returns a boolean if a field has been set.

### SetKeyPassphraseNil

`func (o *GenerateCertificateRequest) SetKeyPassphraseNil(b bool)`

 SetKeyPassphraseNil sets the value for KeyPassphrase to be an explicit nil

### UnsetKeyPassphrase
`func (o *GenerateCertificateRequest) UnsetKeyPassphrase()`

UnsetKeyPassphrase ensures that no value is present for KeyPassphrase, not even an explicit nil
### GetGenerateP12

`func (o *GenerateCertificateRequest) GetGenerateP12() bool`

GetGenerateP12 returns the GenerateP12 field if non-nil, zero value otherwise.

### GetGenerateP12Ok

`func (o *GenerateCertificateRequest) GetGenerateP12Ok() (*bool, bool)`

GetGenerateP12Ok returns a tuple with the GenerateP12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateP12

`func (o *GenerateCertificateRequest) SetGenerateP12(v bool)`

SetGenerateP12 sets GenerateP12 field to given value.

### HasGenerateP12

`func (o *GenerateCertificateRequest) HasGenerateP12() bool`

HasGenerateP12 returns a boolean if a field has been set.

### GetP12Passphrase

`func (o *GenerateCertificateRequest) GetP12Passphrase() string`

GetP12Passphrase returns the P12Passphrase field if non-nil, zero value otherwise.

### GetP12PassphraseOk

`func (o *GenerateCertificateRequest) GetP12PassphraseOk() (*string, bool)`

GetP12PassphraseOk returns a tuple with the P12Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP12Passphrase

`func (o *GenerateCertificateRequest) SetP12Passphrase(v string)`

SetP12Passphrase sets P12Passphrase field to given value.

### HasP12Passphrase

`func (o *GenerateCertificateRequest) HasP12Passphrase() bool`

HasP12Passphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



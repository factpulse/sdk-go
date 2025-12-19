# SignatureParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyPem** | Pointer to **NullableString** |  | [optional] 
**CertPem** | Pointer to **NullableString** |  | [optional] 
**KeyPassphrase** | Pointer to **NullableString** |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 
**Location** | Pointer to **NullableString** |  | [optional] 
**Contact** | Pointer to **NullableString** |  | [optional] 
**FieldName** | Pointer to **string** | PDF signature field name | [optional] [default to "FactPulseSignature"]
**UsePadesLt** | Pointer to **bool** | Enable PAdES-B-LT (long-term archival). REQUIRES certificate with OCSP/CRL access | [optional] [default to false]
**UseTimestamp** | Pointer to **bool** | Enable RFC 3161 timestamping with FreeTSA (PAdES-B-T) | [optional] [default to true]

## Methods

### NewSignatureParameters

`func NewSignatureParameters() *SignatureParameters`

NewSignatureParameters instantiates a new SignatureParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureParametersWithDefaults

`func NewSignatureParametersWithDefaults() *SignatureParameters`

NewSignatureParametersWithDefaults instantiates a new SignatureParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyPem

`func (o *SignatureParameters) GetKeyPem() string`

GetKeyPem returns the KeyPem field if non-nil, zero value otherwise.

### GetKeyPemOk

`func (o *SignatureParameters) GetKeyPemOk() (*string, bool)`

GetKeyPemOk returns a tuple with the KeyPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPem

`func (o *SignatureParameters) SetKeyPem(v string)`

SetKeyPem sets KeyPem field to given value.

### HasKeyPem

`func (o *SignatureParameters) HasKeyPem() bool`

HasKeyPem returns a boolean if a field has been set.

### SetKeyPemNil

`func (o *SignatureParameters) SetKeyPemNil(b bool)`

 SetKeyPemNil sets the value for KeyPem to be an explicit nil

### UnsetKeyPem
`func (o *SignatureParameters) UnsetKeyPem()`

UnsetKeyPem ensures that no value is present for KeyPem, not even an explicit nil
### GetCertPem

`func (o *SignatureParameters) GetCertPem() string`

GetCertPem returns the CertPem field if non-nil, zero value otherwise.

### GetCertPemOk

`func (o *SignatureParameters) GetCertPemOk() (*string, bool)`

GetCertPemOk returns a tuple with the CertPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertPem

`func (o *SignatureParameters) SetCertPem(v string)`

SetCertPem sets CertPem field to given value.

### HasCertPem

`func (o *SignatureParameters) HasCertPem() bool`

HasCertPem returns a boolean if a field has been set.

### SetCertPemNil

`func (o *SignatureParameters) SetCertPemNil(b bool)`

 SetCertPemNil sets the value for CertPem to be an explicit nil

### UnsetCertPem
`func (o *SignatureParameters) UnsetCertPem()`

UnsetCertPem ensures that no value is present for CertPem, not even an explicit nil
### GetKeyPassphrase

`func (o *SignatureParameters) GetKeyPassphrase() string`

GetKeyPassphrase returns the KeyPassphrase field if non-nil, zero value otherwise.

### GetKeyPassphraseOk

`func (o *SignatureParameters) GetKeyPassphraseOk() (*string, bool)`

GetKeyPassphraseOk returns a tuple with the KeyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPassphrase

`func (o *SignatureParameters) SetKeyPassphrase(v string)`

SetKeyPassphrase sets KeyPassphrase field to given value.

### HasKeyPassphrase

`func (o *SignatureParameters) HasKeyPassphrase() bool`

HasKeyPassphrase returns a boolean if a field has been set.

### SetKeyPassphraseNil

`func (o *SignatureParameters) SetKeyPassphraseNil(b bool)`

 SetKeyPassphraseNil sets the value for KeyPassphrase to be an explicit nil

### UnsetKeyPassphrase
`func (o *SignatureParameters) UnsetKeyPassphrase()`

UnsetKeyPassphrase ensures that no value is present for KeyPassphrase, not even an explicit nil
### GetReason

`func (o *SignatureParameters) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SignatureParameters) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SignatureParameters) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SignatureParameters) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *SignatureParameters) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *SignatureParameters) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetLocation

`func (o *SignatureParameters) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SignatureParameters) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SignatureParameters) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SignatureParameters) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *SignatureParameters) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *SignatureParameters) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetContact

`func (o *SignatureParameters) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *SignatureParameters) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *SignatureParameters) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *SignatureParameters) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *SignatureParameters) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *SignatureParameters) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetFieldName

`func (o *SignatureParameters) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *SignatureParameters) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *SignatureParameters) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *SignatureParameters) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetUsePadesLt

`func (o *SignatureParameters) GetUsePadesLt() bool`

GetUsePadesLt returns the UsePadesLt field if non-nil, zero value otherwise.

### GetUsePadesLtOk

`func (o *SignatureParameters) GetUsePadesLtOk() (*bool, bool)`

GetUsePadesLtOk returns a tuple with the UsePadesLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePadesLt

`func (o *SignatureParameters) SetUsePadesLt(v bool)`

SetUsePadesLt sets UsePadesLt field to given value.

### HasUsePadesLt

`func (o *SignatureParameters) HasUsePadesLt() bool`

HasUsePadesLt returns a boolean if a field has been set.

### GetUseTimestamp

`func (o *SignatureParameters) GetUseTimestamp() bool`

GetUseTimestamp returns the UseTimestamp field if non-nil, zero value otherwise.

### GetUseTimestampOk

`func (o *SignatureParameters) GetUseTimestampOk() (*bool, bool)`

GetUseTimestampOk returns a tuple with the UseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTimestamp

`func (o *SignatureParameters) SetUseTimestamp(v bool)`

SetUseTimestamp sets UseTimestamp field to given value.

### HasUseTimestamp

`func (o *SignatureParameters) HasUseTimestamp() bool`

HasUseTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



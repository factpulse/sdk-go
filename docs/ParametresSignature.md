# ParametresSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyPem** | Pointer to **NullableString** |  | [optional] 
**CertPem** | Pointer to **NullableString** |  | [optional] 
**KeyPassphrase** | Pointer to **NullableString** |  | [optional] 
**Raison** | Pointer to **NullableString** |  | [optional] 
**Localisation** | Pointer to **NullableString** |  | [optional] 
**Contact** | Pointer to **NullableString** |  | [optional] 
**FieldName** | Pointer to **string** | Nom du champ de signature PDF | [optional] [default to "FactPulseSignature"]
**UsePadesLt** | Pointer to **bool** | Activer PAdES-B-LT (archivage long terme). NÉCESSITE certificat avec accès OCSP/CRL | [optional] [default to false]
**UseTimestamp** | Pointer to **bool** | Activer l&#39;horodatage RFC 3161 avec FreeTSA (PAdES-B-T) | [optional] [default to true]

## Methods

### NewParametresSignature

`func NewParametresSignature() *ParametresSignature`

NewParametresSignature instantiates a new ParametresSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParametresSignatureWithDefaults

`func NewParametresSignatureWithDefaults() *ParametresSignature`

NewParametresSignatureWithDefaults instantiates a new ParametresSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyPem

`func (o *ParametresSignature) GetKeyPem() string`

GetKeyPem returns the KeyPem field if non-nil, zero value otherwise.

### GetKeyPemOk

`func (o *ParametresSignature) GetKeyPemOk() (*string, bool)`

GetKeyPemOk returns a tuple with the KeyPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPem

`func (o *ParametresSignature) SetKeyPem(v string)`

SetKeyPem sets KeyPem field to given value.

### HasKeyPem

`func (o *ParametresSignature) HasKeyPem() bool`

HasKeyPem returns a boolean if a field has been set.

### SetKeyPemNil

`func (o *ParametresSignature) SetKeyPemNil(b bool)`

 SetKeyPemNil sets the value for KeyPem to be an explicit nil

### UnsetKeyPem
`func (o *ParametresSignature) UnsetKeyPem()`

UnsetKeyPem ensures that no value is present for KeyPem, not even an explicit nil
### GetCertPem

`func (o *ParametresSignature) GetCertPem() string`

GetCertPem returns the CertPem field if non-nil, zero value otherwise.

### GetCertPemOk

`func (o *ParametresSignature) GetCertPemOk() (*string, bool)`

GetCertPemOk returns a tuple with the CertPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertPem

`func (o *ParametresSignature) SetCertPem(v string)`

SetCertPem sets CertPem field to given value.

### HasCertPem

`func (o *ParametresSignature) HasCertPem() bool`

HasCertPem returns a boolean if a field has been set.

### SetCertPemNil

`func (o *ParametresSignature) SetCertPemNil(b bool)`

 SetCertPemNil sets the value for CertPem to be an explicit nil

### UnsetCertPem
`func (o *ParametresSignature) UnsetCertPem()`

UnsetCertPem ensures that no value is present for CertPem, not even an explicit nil
### GetKeyPassphrase

`func (o *ParametresSignature) GetKeyPassphrase() string`

GetKeyPassphrase returns the KeyPassphrase field if non-nil, zero value otherwise.

### GetKeyPassphraseOk

`func (o *ParametresSignature) GetKeyPassphraseOk() (*string, bool)`

GetKeyPassphraseOk returns a tuple with the KeyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPassphrase

`func (o *ParametresSignature) SetKeyPassphrase(v string)`

SetKeyPassphrase sets KeyPassphrase field to given value.

### HasKeyPassphrase

`func (o *ParametresSignature) HasKeyPassphrase() bool`

HasKeyPassphrase returns a boolean if a field has been set.

### SetKeyPassphraseNil

`func (o *ParametresSignature) SetKeyPassphraseNil(b bool)`

 SetKeyPassphraseNil sets the value for KeyPassphrase to be an explicit nil

### UnsetKeyPassphrase
`func (o *ParametresSignature) UnsetKeyPassphrase()`

UnsetKeyPassphrase ensures that no value is present for KeyPassphrase, not even an explicit nil
### GetRaison

`func (o *ParametresSignature) GetRaison() string`

GetRaison returns the Raison field if non-nil, zero value otherwise.

### GetRaisonOk

`func (o *ParametresSignature) GetRaisonOk() (*string, bool)`

GetRaisonOk returns a tuple with the Raison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaison

`func (o *ParametresSignature) SetRaison(v string)`

SetRaison sets Raison field to given value.

### HasRaison

`func (o *ParametresSignature) HasRaison() bool`

HasRaison returns a boolean if a field has been set.

### SetRaisonNil

`func (o *ParametresSignature) SetRaisonNil(b bool)`

 SetRaisonNil sets the value for Raison to be an explicit nil

### UnsetRaison
`func (o *ParametresSignature) UnsetRaison()`

UnsetRaison ensures that no value is present for Raison, not even an explicit nil
### GetLocalisation

`func (o *ParametresSignature) GetLocalisation() string`

GetLocalisation returns the Localisation field if non-nil, zero value otherwise.

### GetLocalisationOk

`func (o *ParametresSignature) GetLocalisationOk() (*string, bool)`

GetLocalisationOk returns a tuple with the Localisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalisation

`func (o *ParametresSignature) SetLocalisation(v string)`

SetLocalisation sets Localisation field to given value.

### HasLocalisation

`func (o *ParametresSignature) HasLocalisation() bool`

HasLocalisation returns a boolean if a field has been set.

### SetLocalisationNil

`func (o *ParametresSignature) SetLocalisationNil(b bool)`

 SetLocalisationNil sets the value for Localisation to be an explicit nil

### UnsetLocalisation
`func (o *ParametresSignature) UnsetLocalisation()`

UnsetLocalisation ensures that no value is present for Localisation, not even an explicit nil
### GetContact

`func (o *ParametresSignature) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ParametresSignature) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ParametresSignature) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ParametresSignature) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *ParametresSignature) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *ParametresSignature) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetFieldName

`func (o *ParametresSignature) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *ParametresSignature) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *ParametresSignature) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *ParametresSignature) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetUsePadesLt

`func (o *ParametresSignature) GetUsePadesLt() bool`

GetUsePadesLt returns the UsePadesLt field if non-nil, zero value otherwise.

### GetUsePadesLtOk

`func (o *ParametresSignature) GetUsePadesLtOk() (*bool, bool)`

GetUsePadesLtOk returns a tuple with the UsePadesLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePadesLt

`func (o *ParametresSignature) SetUsePadesLt(v bool)`

SetUsePadesLt sets UsePadesLt field to given value.

### HasUsePadesLt

`func (o *ParametresSignature) HasUsePadesLt() bool`

HasUsePadesLt returns a boolean if a field has been set.

### GetUseTimestamp

`func (o *ParametresSignature) GetUseTimestamp() bool`

GetUseTimestamp returns the UseTimestamp field if non-nil, zero value otherwise.

### GetUseTimestampOk

`func (o *ParametresSignature) GetUseTimestampOk() (*bool, bool)`

GetUseTimestampOk returns a tuple with the UseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTimestamp

`func (o *ParametresSignature) SetUseTimestamp(v bool)`

SetUseTimestamp sets UseTimestamp field to given value.

### HasUseTimestamp

`func (o *ParametresSignature) HasUseTimestamp() bool`

HasUseTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



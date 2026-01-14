# CertificateInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cn** | **string** | Common Name | 
**Organization** | **string** | Organization | 
**Country** | **string** | Country code | 
**City** | **string** | City | 
**State** | **string** | State/Province | 
**Email** | Pointer to **NullableString** |  | [optional] 
**Subject** | **string** | Full subject (RFC4514) | 
**Issuer** | **string** | Issuer (self-signed &#x3D; same as subject) | 
**SerialNumber** | **int32** | Certificate serial number | 
**ValidFrom** | **string** | Validity start date (ISO 8601) | 
**ValidTo** | **string** | Validity end date (ISO 8601) | 
**Algorithm** | **string** | Signature algorithm | 

## Methods

### NewCertificateInfoResponse

`func NewCertificateInfoResponse(cn string, organization string, country string, city string, state string, subject string, issuer string, serialNumber int32, validFrom string, validTo string, algorithm string, ) *CertificateInfoResponse`

NewCertificateInfoResponse instantiates a new CertificateInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateInfoResponseWithDefaults

`func NewCertificateInfoResponseWithDefaults() *CertificateInfoResponse`

NewCertificateInfoResponseWithDefaults instantiates a new CertificateInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCn

`func (o *CertificateInfoResponse) GetCn() string`

GetCn returns the Cn field if non-nil, zero value otherwise.

### GetCnOk

`func (o *CertificateInfoResponse) GetCnOk() (*string, bool)`

GetCnOk returns a tuple with the Cn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCn

`func (o *CertificateInfoResponse) SetCn(v string)`

SetCn sets Cn field to given value.


### GetOrganization

`func (o *CertificateInfoResponse) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CertificateInfoResponse) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CertificateInfoResponse) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetCountry

`func (o *CertificateInfoResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CertificateInfoResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CertificateInfoResponse) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCity

`func (o *CertificateInfoResponse) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CertificateInfoResponse) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CertificateInfoResponse) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *CertificateInfoResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CertificateInfoResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CertificateInfoResponse) SetState(v string)`

SetState sets State field to given value.


### GetEmail

`func (o *CertificateInfoResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CertificateInfoResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CertificateInfoResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CertificateInfoResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CertificateInfoResponse) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CertificateInfoResponse) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetSubject

`func (o *CertificateInfoResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CertificateInfoResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CertificateInfoResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetIssuer

`func (o *CertificateInfoResponse) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CertificateInfoResponse) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CertificateInfoResponse) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetSerialNumber

`func (o *CertificateInfoResponse) GetSerialNumber() int32`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CertificateInfoResponse) GetSerialNumberOk() (*int32, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CertificateInfoResponse) SetSerialNumber(v int32)`

SetSerialNumber sets SerialNumber field to given value.


### GetValidFrom

`func (o *CertificateInfoResponse) GetValidFrom() string`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *CertificateInfoResponse) GetValidFromOk() (*string, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *CertificateInfoResponse) SetValidFrom(v string)`

SetValidFrom sets ValidFrom field to given value.


### GetValidTo

`func (o *CertificateInfoResponse) GetValidTo() string`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *CertificateInfoResponse) GetValidToOk() (*string, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *CertificateInfoResponse) SetValidTo(v string)`

SetValidTo sets ValidTo field to given value.


### GetAlgorithm

`func (o *CertificateInfoResponse) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *CertificateInfoResponse) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *CertificateInfoResponse) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AFNORFacilityPayloadHistoryUleB2gAdditionalData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pm** | Pointer to **bool** | Indicates whether the public structure acts as project manager for work invoices in addition to receiving simple invoices. This attribute is only returned if the directory line is defined for a public structure at the SIREN / SIRET or SIREN / SIRET / Routing code level. | [optional] 
**PmOnly** | Pointer to **bool** | Indicates whether the public structure only acts as a project manager; if so, it can only receive invoices for work. This attribute is only returned if the directory line is defined for a public structure at the SIREN/SIRET or SIREN/SIRET/Routing code level. | [optional] 
**ManagesPaymentStatus** | Pointer to **bool** | Indicates whether the public structure manages the payment status of invoices. This attribute is only returned if the directory line is defined for a public structure at the SIREN / SIRET or SIREN / SIRET / Routing code level. | [optional] 
**ManagesLegalCommitmentCode** | Pointer to **bool** | Indicates whether the public structure requires a legal commitment number. This attribute is only returned if the directory line is defined for a public structure at the SIREN / SIRET or SIREN / SIRET / Routing code level. | [optional] 
**ManagesLegalCommitmentOrServiceCode** | Pointer to **bool** | Indicates whether the public structure requires a service code or a commitment code in its exchanges. This attribute is only returned if the directory line is defined for a public structure at the SIREN / SIRET or SIREN / SIRET / Routing code level. | [optional] 
**ServiceCodeStatus** | Pointer to **bool** | Indicates whether the structure requires a service code. This attribute is only returned if the directory line is defined for a public structure at the SIREN / SIRET or SIREN / SIRET / Routing code level. | [optional] 

## Methods

### NewAFNORFacilityPayloadHistoryUleB2gAdditionalData

`func NewAFNORFacilityPayloadHistoryUleB2gAdditionalData() *AFNORFacilityPayloadHistoryUleB2gAdditionalData`

NewAFNORFacilityPayloadHistoryUleB2gAdditionalData instantiates a new AFNORFacilityPayloadHistoryUleB2gAdditionalData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORFacilityPayloadHistoryUleB2gAdditionalDataWithDefaults

`func NewAFNORFacilityPayloadHistoryUleB2gAdditionalDataWithDefaults() *AFNORFacilityPayloadHistoryUleB2gAdditionalData`

NewAFNORFacilityPayloadHistoryUleB2gAdditionalDataWithDefaults instantiates a new AFNORFacilityPayloadHistoryUleB2gAdditionalData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPm

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) GetPm() bool`

GetPm returns the Pm field if non-nil, zero value otherwise.

### GetPmOk

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) GetPmOk() (*bool, bool)`

GetPmOk returns a tuple with the Pm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) SetPm(v bool)`

SetPm sets Pm field to given value.

### HasPm

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) HasPm() bool`

HasPm returns a boolean if a field has been set.

### GetPmOnly

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) GetPmOnly() bool`

GetPmOnly returns the PmOnly field if non-nil, zero value otherwise.

### GetPmOnlyOk

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) GetPmOnlyOk() (*bool, bool)`

GetPmOnlyOk returns a tuple with the PmOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmOnly

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) SetPmOnly(v bool)`

SetPmOnly sets PmOnly field to given value.

### HasPmOnly

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) HasPmOnly() bool`

HasPmOnly returns a boolean if a field has been set.

### GetManagesPaymentStatus

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) GetManagesPaymentStatus() bool`

GetManagesPaymentStatus returns the ManagesPaymentStatus field if non-nil, zero value otherwise.

### GetManagesPaymentStatusOk

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) GetManagesPaymentStatusOk() (*bool, bool)`

GetManagesPaymentStatusOk returns a tuple with the ManagesPaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagesPaymentStatus

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) SetManagesPaymentStatus(v bool)`

SetManagesPaymentStatus sets ManagesPaymentStatus field to given value.

### HasManagesPaymentStatus

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) HasManagesPaymentStatus() bool`

HasManagesPaymentStatus returns a boolean if a field has been set.

### GetManagesLegalCommitmentCode

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) GetManagesLegalCommitmentCode() bool`

GetManagesLegalCommitmentCode returns the ManagesLegalCommitmentCode field if non-nil, zero value otherwise.

### GetManagesLegalCommitmentCodeOk

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) GetManagesLegalCommitmentCodeOk() (*bool, bool)`

GetManagesLegalCommitmentCodeOk returns a tuple with the ManagesLegalCommitmentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagesLegalCommitmentCode

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) SetManagesLegalCommitmentCode(v bool)`

SetManagesLegalCommitmentCode sets ManagesLegalCommitmentCode field to given value.

### HasManagesLegalCommitmentCode

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) HasManagesLegalCommitmentCode() bool`

HasManagesLegalCommitmentCode returns a boolean if a field has been set.

### GetManagesLegalCommitmentOrServiceCode

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) GetManagesLegalCommitmentOrServiceCode() bool`

GetManagesLegalCommitmentOrServiceCode returns the ManagesLegalCommitmentOrServiceCode field if non-nil, zero value otherwise.

### GetManagesLegalCommitmentOrServiceCodeOk

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) GetManagesLegalCommitmentOrServiceCodeOk() (*bool, bool)`

GetManagesLegalCommitmentOrServiceCodeOk returns a tuple with the ManagesLegalCommitmentOrServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagesLegalCommitmentOrServiceCode

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) SetManagesLegalCommitmentOrServiceCode(v bool)`

SetManagesLegalCommitmentOrServiceCode sets ManagesLegalCommitmentOrServiceCode field to given value.

### HasManagesLegalCommitmentOrServiceCode

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) HasManagesLegalCommitmentOrServiceCode() bool`

HasManagesLegalCommitmentOrServiceCode returns a boolean if a field has been set.

### GetServiceCodeStatus

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) GetServiceCodeStatus() bool`

GetServiceCodeStatus returns the ServiceCodeStatus field if non-nil, zero value otherwise.

### GetServiceCodeStatusOk

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) GetServiceCodeStatusOk() (*bool, bool)`

GetServiceCodeStatusOk returns a tuple with the ServiceCodeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCodeStatus

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) SetServiceCodeStatus(v bool)`

SetServiceCodeStatus sets ServiceCodeStatus field to given value.

### HasServiceCodeStatus

`func (o *AFNORFacilityPayloadHistoryUleB2gAdditionalData) HasServiceCodeStatus() bool`

HasServiceCodeStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



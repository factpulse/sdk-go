# ReportSender

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Siren** | **string** | SIREN or SIRET number | 
**Name** | **string** | Company name | 
**VatId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewReportSender

`func NewReportSender(siren string, name string, ) *ReportSender`

NewReportSender instantiates a new ReportSender object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportSenderWithDefaults

`func NewReportSenderWithDefaults() *ReportSender`

NewReportSenderWithDefaults instantiates a new ReportSender object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiren

`func (o *ReportSender) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *ReportSender) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *ReportSender) SetSiren(v string)`

SetSiren sets Siren field to given value.


### GetName

`func (o *ReportSender) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportSender) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportSender) SetName(v string)`

SetName sets Name field to given value.


### GetVatId

`func (o *ReportSender) GetVatId() string`

GetVatId returns the VatId field if non-nil, zero value otherwise.

### GetVatIdOk

`func (o *ReportSender) GetVatIdOk() (*string, bool)`

GetVatIdOk returns a tuple with the VatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatId

`func (o *ReportSender) SetVatId(v string)`

SetVatId sets VatId field to given value.

### HasVatId

`func (o *ReportSender) HasVatId() bool`

HasVatId returns a boolean if a field has been set.

### SetVatIdNil

`func (o *ReportSender) SetVatIdNil(b bool)`

 SetVatIdNil sets the value for VatId to be an explicit nil

### UnsetVatId
`func (o *ReportSender) UnsetVatId()`

UnsetVatId ensures that no value is present for VatId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



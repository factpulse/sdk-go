# ReportIssuer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Siren** | **string** | SIREN or SIRET of the declarant (French company) | 
**Name** | Pointer to **NullableString** |  | [optional] 
**VatId** | Pointer to **NullableString** |  | [optional] 
**RoleCode** | Pointer to [**IssuerRoleCode**](IssuerRoleCode.md) | Role of the declarant (TT-15). SE&#x3D;Seller (B2Bi: French seller to foreign buyer). BY&#x3D;Buyer (Bi2B: French buyer from foreign seller). | [optional] [default to SE]

## Methods

### NewReportIssuer

`func NewReportIssuer(siren string, ) *ReportIssuer`

NewReportIssuer instantiates a new ReportIssuer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportIssuerWithDefaults

`func NewReportIssuerWithDefaults() *ReportIssuer`

NewReportIssuerWithDefaults instantiates a new ReportIssuer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiren

`func (o *ReportIssuer) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *ReportIssuer) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *ReportIssuer) SetSiren(v string)`

SetSiren sets Siren field to given value.


### GetName

`func (o *ReportIssuer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportIssuer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportIssuer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReportIssuer) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ReportIssuer) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ReportIssuer) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVatId

`func (o *ReportIssuer) GetVatId() string`

GetVatId returns the VatId field if non-nil, zero value otherwise.

### GetVatIdOk

`func (o *ReportIssuer) GetVatIdOk() (*string, bool)`

GetVatIdOk returns a tuple with the VatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatId

`func (o *ReportIssuer) SetVatId(v string)`

SetVatId sets VatId field to given value.

### HasVatId

`func (o *ReportIssuer) HasVatId() bool`

HasVatId returns a boolean if a field has been set.

### SetVatIdNil

`func (o *ReportIssuer) SetVatIdNil(b bool)`

 SetVatIdNil sets the value for VatId to be an explicit nil

### UnsetVatId
`func (o *ReportIssuer) UnsetVatId()`

UnsetVatId ensures that no value is present for VatId, not even an explicit nil
### GetRoleCode

`func (o *ReportIssuer) GetRoleCode() IssuerRoleCode`

GetRoleCode returns the RoleCode field if non-nil, zero value otherwise.

### GetRoleCodeOk

`func (o *ReportIssuer) GetRoleCodeOk() (*IssuerRoleCode, bool)`

GetRoleCodeOk returns a tuple with the RoleCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleCode

`func (o *ReportIssuer) SetRoleCode(v IssuerRoleCode)`

SetRoleCode sets RoleCode field to given value.

### HasRoleCode

`func (o *ReportIssuer) HasRoleCode() bool`

HasRoleCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



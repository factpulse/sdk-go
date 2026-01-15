# AFNORFacilityPayloadHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Siret** | Pointer to **string** | SIRET Number | [optional] 
**Siren** | Pointer to **string** | SIREN number | [optional] 
**Name** | Pointer to **string** | business name | [optional] 
**FacilityType** | Pointer to [**AFNORFacilityType**](AFNORFacilityType.md) |  | [optional] 
**Diffusible** | Pointer to [**AFNORDiffusionStatus**](AFNORDiffusionStatus.md) |  | [optional] 
**AdministrativeStatus** | Pointer to [**AFNORFacilityAdministrativeStatus**](AFNORFacilityAdministrativeStatus.md) |  | [optional] 
**Address** | Pointer to [**AFNORAddressRead**](AFNORAddressRead.md) |  | [optional] 
**B2gAdditionalData** | Pointer to [**AFNORFacilityPayloadHistoryUleB2gAdditionalData**](AFNORFacilityPayloadHistoryUleB2gAdditionalData.md) |  | [optional] 
**LegalUnit** | Pointer to [**AFNORLegalUnitPayloadIncludedNoSiren**](AFNORLegalUnitPayloadIncludedNoSiren.md) |  | [optional] 

## Methods

### NewAFNORFacilityPayloadHistory

`func NewAFNORFacilityPayloadHistory() *AFNORFacilityPayloadHistory`

NewAFNORFacilityPayloadHistory instantiates a new AFNORFacilityPayloadHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORFacilityPayloadHistoryWithDefaults

`func NewAFNORFacilityPayloadHistoryWithDefaults() *AFNORFacilityPayloadHistory`

NewAFNORFacilityPayloadHistoryWithDefaults instantiates a new AFNORFacilityPayloadHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiret

`func (o *AFNORFacilityPayloadHistory) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *AFNORFacilityPayloadHistory) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *AFNORFacilityPayloadHistory) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *AFNORFacilityPayloadHistory) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### GetSiren

`func (o *AFNORFacilityPayloadHistory) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *AFNORFacilityPayloadHistory) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *AFNORFacilityPayloadHistory) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *AFNORFacilityPayloadHistory) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### GetName

`func (o *AFNORFacilityPayloadHistory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AFNORFacilityPayloadHistory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AFNORFacilityPayloadHistory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AFNORFacilityPayloadHistory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFacilityType

`func (o *AFNORFacilityPayloadHistory) GetFacilityType() AFNORFacilityType`

GetFacilityType returns the FacilityType field if non-nil, zero value otherwise.

### GetFacilityTypeOk

`func (o *AFNORFacilityPayloadHistory) GetFacilityTypeOk() (*AFNORFacilityType, bool)`

GetFacilityTypeOk returns a tuple with the FacilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityType

`func (o *AFNORFacilityPayloadHistory) SetFacilityType(v AFNORFacilityType)`

SetFacilityType sets FacilityType field to given value.

### HasFacilityType

`func (o *AFNORFacilityPayloadHistory) HasFacilityType() bool`

HasFacilityType returns a boolean if a field has been set.

### GetDiffusible

`func (o *AFNORFacilityPayloadHistory) GetDiffusible() AFNORDiffusionStatus`

GetDiffusible returns the Diffusible field if non-nil, zero value otherwise.

### GetDiffusibleOk

`func (o *AFNORFacilityPayloadHistory) GetDiffusibleOk() (*AFNORDiffusionStatus, bool)`

GetDiffusibleOk returns a tuple with the Diffusible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffusible

`func (o *AFNORFacilityPayloadHistory) SetDiffusible(v AFNORDiffusionStatus)`

SetDiffusible sets Diffusible field to given value.

### HasDiffusible

`func (o *AFNORFacilityPayloadHistory) HasDiffusible() bool`

HasDiffusible returns a boolean if a field has been set.

### GetAdministrativeStatus

`func (o *AFNORFacilityPayloadHistory) GetAdministrativeStatus() AFNORFacilityAdministrativeStatus`

GetAdministrativeStatus returns the AdministrativeStatus field if non-nil, zero value otherwise.

### GetAdministrativeStatusOk

`func (o *AFNORFacilityPayloadHistory) GetAdministrativeStatusOk() (*AFNORFacilityAdministrativeStatus, bool)`

GetAdministrativeStatusOk returns a tuple with the AdministrativeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeStatus

`func (o *AFNORFacilityPayloadHistory) SetAdministrativeStatus(v AFNORFacilityAdministrativeStatus)`

SetAdministrativeStatus sets AdministrativeStatus field to given value.

### HasAdministrativeStatus

`func (o *AFNORFacilityPayloadHistory) HasAdministrativeStatus() bool`

HasAdministrativeStatus returns a boolean if a field has been set.

### GetAddress

`func (o *AFNORFacilityPayloadHistory) GetAddress() AFNORAddressRead`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AFNORFacilityPayloadHistory) GetAddressOk() (*AFNORAddressRead, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AFNORFacilityPayloadHistory) SetAddress(v AFNORAddressRead)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AFNORFacilityPayloadHistory) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetB2gAdditionalData

`func (o *AFNORFacilityPayloadHistory) GetB2gAdditionalData() AFNORFacilityPayloadHistoryUleB2gAdditionalData`

GetB2gAdditionalData returns the B2gAdditionalData field if non-nil, zero value otherwise.

### GetB2gAdditionalDataOk

`func (o *AFNORFacilityPayloadHistory) GetB2gAdditionalDataOk() (*AFNORFacilityPayloadHistoryUleB2gAdditionalData, bool)`

GetB2gAdditionalDataOk returns a tuple with the B2gAdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB2gAdditionalData

`func (o *AFNORFacilityPayloadHistory) SetB2gAdditionalData(v AFNORFacilityPayloadHistoryUleB2gAdditionalData)`

SetB2gAdditionalData sets B2gAdditionalData field to given value.

### HasB2gAdditionalData

`func (o *AFNORFacilityPayloadHistory) HasB2gAdditionalData() bool`

HasB2gAdditionalData returns a boolean if a field has been set.

### GetLegalUnit

`func (o *AFNORFacilityPayloadHistory) GetLegalUnit() AFNORLegalUnitPayloadIncludedNoSiren`

GetLegalUnit returns the LegalUnit field if non-nil, zero value otherwise.

### GetLegalUnitOk

`func (o *AFNORFacilityPayloadHistory) GetLegalUnitOk() (*AFNORLegalUnitPayloadIncludedNoSiren, bool)`

GetLegalUnitOk returns a tuple with the LegalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalUnit

`func (o *AFNORFacilityPayloadHistory) SetLegalUnit(v AFNORLegalUnitPayloadIncludedNoSiren)`

SetLegalUnit sets LegalUnit field to given value.

### HasLegalUnit

`func (o *AFNORFacilityPayloadHistory) HasLegalUnit() bool`

HasLegalUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



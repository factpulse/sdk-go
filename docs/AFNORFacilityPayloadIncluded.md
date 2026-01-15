# AFNORFacilityPayloadIncluded

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

## Methods

### NewAFNORFacilityPayloadIncluded

`func NewAFNORFacilityPayloadIncluded() *AFNORFacilityPayloadIncluded`

NewAFNORFacilityPayloadIncluded instantiates a new AFNORFacilityPayloadIncluded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORFacilityPayloadIncludedWithDefaults

`func NewAFNORFacilityPayloadIncludedWithDefaults() *AFNORFacilityPayloadIncluded`

NewAFNORFacilityPayloadIncludedWithDefaults instantiates a new AFNORFacilityPayloadIncluded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiret

`func (o *AFNORFacilityPayloadIncluded) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *AFNORFacilityPayloadIncluded) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *AFNORFacilityPayloadIncluded) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *AFNORFacilityPayloadIncluded) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### GetSiren

`func (o *AFNORFacilityPayloadIncluded) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *AFNORFacilityPayloadIncluded) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *AFNORFacilityPayloadIncluded) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *AFNORFacilityPayloadIncluded) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### GetName

`func (o *AFNORFacilityPayloadIncluded) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AFNORFacilityPayloadIncluded) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AFNORFacilityPayloadIncluded) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AFNORFacilityPayloadIncluded) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFacilityType

`func (o *AFNORFacilityPayloadIncluded) GetFacilityType() AFNORFacilityType`

GetFacilityType returns the FacilityType field if non-nil, zero value otherwise.

### GetFacilityTypeOk

`func (o *AFNORFacilityPayloadIncluded) GetFacilityTypeOk() (*AFNORFacilityType, bool)`

GetFacilityTypeOk returns a tuple with the FacilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityType

`func (o *AFNORFacilityPayloadIncluded) SetFacilityType(v AFNORFacilityType)`

SetFacilityType sets FacilityType field to given value.

### HasFacilityType

`func (o *AFNORFacilityPayloadIncluded) HasFacilityType() bool`

HasFacilityType returns a boolean if a field has been set.

### GetDiffusible

`func (o *AFNORFacilityPayloadIncluded) GetDiffusible() AFNORDiffusionStatus`

GetDiffusible returns the Diffusible field if non-nil, zero value otherwise.

### GetDiffusibleOk

`func (o *AFNORFacilityPayloadIncluded) GetDiffusibleOk() (*AFNORDiffusionStatus, bool)`

GetDiffusibleOk returns a tuple with the Diffusible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffusible

`func (o *AFNORFacilityPayloadIncluded) SetDiffusible(v AFNORDiffusionStatus)`

SetDiffusible sets Diffusible field to given value.

### HasDiffusible

`func (o *AFNORFacilityPayloadIncluded) HasDiffusible() bool`

HasDiffusible returns a boolean if a field has been set.

### GetAdministrativeStatus

`func (o *AFNORFacilityPayloadIncluded) GetAdministrativeStatus() AFNORFacilityAdministrativeStatus`

GetAdministrativeStatus returns the AdministrativeStatus field if non-nil, zero value otherwise.

### GetAdministrativeStatusOk

`func (o *AFNORFacilityPayloadIncluded) GetAdministrativeStatusOk() (*AFNORFacilityAdministrativeStatus, bool)`

GetAdministrativeStatusOk returns a tuple with the AdministrativeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeStatus

`func (o *AFNORFacilityPayloadIncluded) SetAdministrativeStatus(v AFNORFacilityAdministrativeStatus)`

SetAdministrativeStatus sets AdministrativeStatus field to given value.

### HasAdministrativeStatus

`func (o *AFNORFacilityPayloadIncluded) HasAdministrativeStatus() bool`

HasAdministrativeStatus returns a boolean if a field has been set.

### GetAddress

`func (o *AFNORFacilityPayloadIncluded) GetAddress() AFNORAddressRead`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AFNORFacilityPayloadIncluded) GetAddressOk() (*AFNORAddressRead, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AFNORFacilityPayloadIncluded) SetAddress(v AFNORAddressRead)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AFNORFacilityPayloadIncluded) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetB2gAdditionalData

`func (o *AFNORFacilityPayloadIncluded) GetB2gAdditionalData() AFNORFacilityPayloadHistoryUleB2gAdditionalData`

GetB2gAdditionalData returns the B2gAdditionalData field if non-nil, zero value otherwise.

### GetB2gAdditionalDataOk

`func (o *AFNORFacilityPayloadIncluded) GetB2gAdditionalDataOk() (*AFNORFacilityPayloadHistoryUleB2gAdditionalData, bool)`

GetB2gAdditionalDataOk returns a tuple with the B2gAdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB2gAdditionalData

`func (o *AFNORFacilityPayloadIncluded) SetB2gAdditionalData(v AFNORFacilityPayloadHistoryUleB2gAdditionalData)`

SetB2gAdditionalData sets B2gAdditionalData field to given value.

### HasB2gAdditionalData

`func (o *AFNORFacilityPayloadIncluded) HasB2gAdditionalData() bool`

HasB2gAdditionalData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



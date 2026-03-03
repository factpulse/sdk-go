# AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressingIdentifier** | Pointer to **string** | Addressing identifier of the directory line. | [optional] 
**PlatformType** | Pointer to [**AFNORRecipientPlatformType**](AFNORRecipientPlatformType.md) |  | [optional] 
**DirectoryLineStatus** | Pointer to [**AFNORDirectoryLineStatus**](AFNORDirectoryLineStatus.md) |  | [optional] 
**Siren** | Pointer to **string** | SIREN number | [optional] 
**Siret** | Pointer to **string** | SIRET Number | [optional] 
**AddressingSuffix** | Pointer to **string** | suffix of the directory line which defines an address mesh not attached to a facility | [optional] 
**RoutingCode** | Pointer to [**AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode**](AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode.md) |  | [optional] 
**LegalUnit** | Pointer to [**AFNORLegalUnitPayloadIncluded**](AFNORLegalUnitPayloadIncluded.md) |  | [optional] 
**Facility** | Pointer to [**AFNORFacilityPayloadIncluded**](AFNORFacilityPayloadIncluded.md) |  | [optional] 

## Methods

### NewAFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode

`func NewAFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode() *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode`

NewAFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode instantiates a new AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCodeWithDefaults

`func NewAFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCodeWithDefaults() *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode`

NewAFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCodeWithDefaults instantiates a new AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressingIdentifier

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) GetAddressingIdentifier() string`

GetAddressingIdentifier returns the AddressingIdentifier field if non-nil, zero value otherwise.

### GetAddressingIdentifierOk

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) GetAddressingIdentifierOk() (*string, bool)`

GetAddressingIdentifierOk returns a tuple with the AddressingIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressingIdentifier

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) SetAddressingIdentifier(v string)`

SetAddressingIdentifier sets AddressingIdentifier field to given value.

### HasAddressingIdentifier

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) HasAddressingIdentifier() bool`

HasAddressingIdentifier returns a boolean if a field has been set.

### GetPlatformType

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) GetPlatformType() AFNORRecipientPlatformType`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) GetPlatformTypeOk() (*AFNORRecipientPlatformType, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) SetPlatformType(v AFNORRecipientPlatformType)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetDirectoryLineStatus

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) GetDirectoryLineStatus() AFNORDirectoryLineStatus`

GetDirectoryLineStatus returns the DirectoryLineStatus field if non-nil, zero value otherwise.

### GetDirectoryLineStatusOk

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) GetDirectoryLineStatusOk() (*AFNORDirectoryLineStatus, bool)`

GetDirectoryLineStatusOk returns a tuple with the DirectoryLineStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryLineStatus

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) SetDirectoryLineStatus(v AFNORDirectoryLineStatus)`

SetDirectoryLineStatus sets DirectoryLineStatus field to given value.

### HasDirectoryLineStatus

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) HasDirectoryLineStatus() bool`

HasDirectoryLineStatus returns a boolean if a field has been set.

### GetSiren

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### GetSiret

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### GetAddressingSuffix

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) GetAddressingSuffix() string`

GetAddressingSuffix returns the AddressingSuffix field if non-nil, zero value otherwise.

### GetAddressingSuffixOk

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) GetAddressingSuffixOk() (*string, bool)`

GetAddressingSuffixOk returns a tuple with the AddressingSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressingSuffix

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) SetAddressingSuffix(v string)`

SetAddressingSuffix sets AddressingSuffix field to given value.

### HasAddressingSuffix

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) HasAddressingSuffix() bool`

HasAddressingSuffix returns a boolean if a field has been set.

### GetRoutingCode

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) GetRoutingCode() AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode`

GetRoutingCode returns the RoutingCode field if non-nil, zero value otherwise.

### GetRoutingCodeOk

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) GetRoutingCodeOk() (*AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode, bool)`

GetRoutingCodeOk returns a tuple with the RoutingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingCode

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) SetRoutingCode(v AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode)`

SetRoutingCode sets RoutingCode field to given value.

### HasRoutingCode

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) HasRoutingCode() bool`

HasRoutingCode returns a boolean if a field has been set.

### GetLegalUnit

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) GetLegalUnit() AFNORLegalUnitPayloadIncluded`

GetLegalUnit returns the LegalUnit field if non-nil, zero value otherwise.

### GetLegalUnitOk

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) GetLegalUnitOk() (*AFNORLegalUnitPayloadIncluded, bool)`

GetLegalUnitOk returns a tuple with the LegalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalUnit

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) SetLegalUnit(v AFNORLegalUnitPayloadIncluded)`

SetLegalUnit sets LegalUnit field to given value.

### HasLegalUnit

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) HasLegalUnit() bool`

HasLegalUnit returns a boolean if a field has been set.

### GetFacility

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) GetFacility() AFNORFacilityPayloadIncluded`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) GetFacilityOk() (*AFNORFacilityPayloadIncluded, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) SetFacility(v AFNORFacilityPayloadIncluded)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *AFNORDirectoryLinePayloadStatusLegalUnitFacilityRoutingCode) HasFacility() bool`

HasFacility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



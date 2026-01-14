# AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressingIdentifier** | Pointer to **string** | Addressing identifier of the directory line. | [optional] 
**Siren** | Pointer to **string** | SIREN number | [optional] 
**Siret** | Pointer to **string** | SIRET Number | [optional] 
**AddressingSuffix** | Pointer to **string** | suffix of the directory line which defines an address mesh not attached to a facility | [optional] 
**RoutingCode** | Pointer to [**AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode**](AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode.md) |  | [optional] 
**Platform** | Pointer to [**AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodePlatform**](AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodePlatform.md) |  | [optional] 
**LegalUnit** | Pointer to [**AFNORLegalUnitPayloadIncludedNoSiren**](AFNORLegalUnitPayloadIncludedNoSiren.md) |  | [optional] 
**Facility** | Pointer to [**AFNORFacilityPayloadIncluded**](AFNORFacilityPayloadIncluded.md) |  | [optional] 

## Methods

### NewAFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode

`func NewAFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode() *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode`

NewAFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode instantiates a new AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeWithDefaults

`func NewAFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeWithDefaults() *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode`

NewAFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeWithDefaults instantiates a new AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressingIdentifier

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) GetAddressingIdentifier() string`

GetAddressingIdentifier returns the AddressingIdentifier field if non-nil, zero value otherwise.

### GetAddressingIdentifierOk

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) GetAddressingIdentifierOk() (*string, bool)`

GetAddressingIdentifierOk returns a tuple with the AddressingIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressingIdentifier

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) SetAddressingIdentifier(v string)`

SetAddressingIdentifier sets AddressingIdentifier field to given value.

### HasAddressingIdentifier

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) HasAddressingIdentifier() bool`

HasAddressingIdentifier returns a boolean if a field has been set.

### GetSiren

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### GetSiret

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### GetAddressingSuffix

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) GetAddressingSuffix() string`

GetAddressingSuffix returns the AddressingSuffix field if non-nil, zero value otherwise.

### GetAddressingSuffixOk

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) GetAddressingSuffixOk() (*string, bool)`

GetAddressingSuffixOk returns a tuple with the AddressingSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressingSuffix

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) SetAddressingSuffix(v string)`

SetAddressingSuffix sets AddressingSuffix field to given value.

### HasAddressingSuffix

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) HasAddressingSuffix() bool`

HasAddressingSuffix returns a boolean if a field has been set.

### GetRoutingCode

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) GetRoutingCode() AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode`

GetRoutingCode returns the RoutingCode field if non-nil, zero value otherwise.

### GetRoutingCodeOk

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) GetRoutingCodeOk() (*AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode, bool)`

GetRoutingCodeOk returns a tuple with the RoutingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingCode

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) SetRoutingCode(v AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode)`

SetRoutingCode sets RoutingCode field to given value.

### HasRoutingCode

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) HasRoutingCode() bool`

HasRoutingCode returns a boolean if a field has been set.

### GetPlatform

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) GetPlatform() AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodePlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) GetPlatformOk() (*AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodePlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) SetPlatform(v AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodePlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetLegalUnit

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) GetLegalUnit() AFNORLegalUnitPayloadIncludedNoSiren`

GetLegalUnit returns the LegalUnit field if non-nil, zero value otherwise.

### GetLegalUnitOk

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) GetLegalUnitOk() (*AFNORLegalUnitPayloadIncludedNoSiren, bool)`

GetLegalUnitOk returns a tuple with the LegalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalUnit

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) SetLegalUnit(v AFNORLegalUnitPayloadIncludedNoSiren)`

SetLegalUnit sets LegalUnit field to given value.

### HasLegalUnit

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) HasLegalUnit() bool`

HasLegalUnit returns a boolean if a field has been set.

### GetFacility

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) GetFacility() AFNORFacilityPayloadIncluded`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) GetFacilityOk() (*AFNORFacilityPayloadIncluded, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) SetFacility(v AFNORFacilityPayloadIncluded)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCode) HasFacility() bool`

HasFacility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AFNORRoutingCodePayloadHistoryLegalUnitFacility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoutingIdentifier** | Pointer to **string** | Routing identifier od a routing code. | [optional] 
**Siret** | Pointer to **string** | SIRET Number | [optional] 
**RoutingIdentifierType** | Pointer to **string** | Routing Identifier type. | [optional] 
**RoutingCodeName** | Pointer to **string** | Name of the directory line routing code. This attribute is only returned if the directory line is defined at the SIREN / SIRET / Routing code mesh. | [optional] 
**ManagesLegalCommitmentCode** | Pointer to **bool** | Indicates whether the public structure requires a legal commitment number. This attribute is only returned if the directory line is defined for a public structure at the SIREN / SIRET or SIREN / SIRET / Routing code level. | [optional] 
**AdministrativeStatus** | Pointer to [**AFNORRoutingCodeAdministrativeStatus**](AFNORRoutingCodeAdministrativeStatus.md) |  | [optional] 
**Address** | Pointer to [**AFNORAddressRead**](AFNORAddressRead.md) |  | [optional] 
**LegalUnit** | Pointer to [**AFNORLegalUnitPayloadIncluded**](AFNORLegalUnitPayloadIncluded.md) |  | [optional] 
**Facility** | Pointer to [**AFNORFacilityPayloadIncluded**](AFNORFacilityPayloadIncluded.md) |  | [optional] 

## Methods

### NewAFNORRoutingCodePayloadHistoryLegalUnitFacility

`func NewAFNORRoutingCodePayloadHistoryLegalUnitFacility() *AFNORRoutingCodePayloadHistoryLegalUnitFacility`

NewAFNORRoutingCodePayloadHistoryLegalUnitFacility instantiates a new AFNORRoutingCodePayloadHistoryLegalUnitFacility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORRoutingCodePayloadHistoryLegalUnitFacilityWithDefaults

`func NewAFNORRoutingCodePayloadHistoryLegalUnitFacilityWithDefaults() *AFNORRoutingCodePayloadHistoryLegalUnitFacility`

NewAFNORRoutingCodePayloadHistoryLegalUnitFacilityWithDefaults instantiates a new AFNORRoutingCodePayloadHistoryLegalUnitFacility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutingIdentifier

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) GetRoutingIdentifier() string`

GetRoutingIdentifier returns the RoutingIdentifier field if non-nil, zero value otherwise.

### GetRoutingIdentifierOk

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) GetRoutingIdentifierOk() (*string, bool)`

GetRoutingIdentifierOk returns a tuple with the RoutingIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIdentifier

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) SetRoutingIdentifier(v string)`

SetRoutingIdentifier sets RoutingIdentifier field to given value.

### HasRoutingIdentifier

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) HasRoutingIdentifier() bool`

HasRoutingIdentifier returns a boolean if a field has been set.

### GetSiret

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### GetRoutingIdentifierType

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) GetRoutingIdentifierType() string`

GetRoutingIdentifierType returns the RoutingIdentifierType field if non-nil, zero value otherwise.

### GetRoutingIdentifierTypeOk

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) GetRoutingIdentifierTypeOk() (*string, bool)`

GetRoutingIdentifierTypeOk returns a tuple with the RoutingIdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIdentifierType

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) SetRoutingIdentifierType(v string)`

SetRoutingIdentifierType sets RoutingIdentifierType field to given value.

### HasRoutingIdentifierType

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) HasRoutingIdentifierType() bool`

HasRoutingIdentifierType returns a boolean if a field has been set.

### GetRoutingCodeName

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) GetRoutingCodeName() string`

GetRoutingCodeName returns the RoutingCodeName field if non-nil, zero value otherwise.

### GetRoutingCodeNameOk

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) GetRoutingCodeNameOk() (*string, bool)`

GetRoutingCodeNameOk returns a tuple with the RoutingCodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingCodeName

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) SetRoutingCodeName(v string)`

SetRoutingCodeName sets RoutingCodeName field to given value.

### HasRoutingCodeName

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) HasRoutingCodeName() bool`

HasRoutingCodeName returns a boolean if a field has been set.

### GetManagesLegalCommitmentCode

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) GetManagesLegalCommitmentCode() bool`

GetManagesLegalCommitmentCode returns the ManagesLegalCommitmentCode field if non-nil, zero value otherwise.

### GetManagesLegalCommitmentCodeOk

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) GetManagesLegalCommitmentCodeOk() (*bool, bool)`

GetManagesLegalCommitmentCodeOk returns a tuple with the ManagesLegalCommitmentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagesLegalCommitmentCode

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) SetManagesLegalCommitmentCode(v bool)`

SetManagesLegalCommitmentCode sets ManagesLegalCommitmentCode field to given value.

### HasManagesLegalCommitmentCode

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) HasManagesLegalCommitmentCode() bool`

HasManagesLegalCommitmentCode returns a boolean if a field has been set.

### GetAdministrativeStatus

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) GetAdministrativeStatus() AFNORRoutingCodeAdministrativeStatus`

GetAdministrativeStatus returns the AdministrativeStatus field if non-nil, zero value otherwise.

### GetAdministrativeStatusOk

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) GetAdministrativeStatusOk() (*AFNORRoutingCodeAdministrativeStatus, bool)`

GetAdministrativeStatusOk returns a tuple with the AdministrativeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeStatus

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) SetAdministrativeStatus(v AFNORRoutingCodeAdministrativeStatus)`

SetAdministrativeStatus sets AdministrativeStatus field to given value.

### HasAdministrativeStatus

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) HasAdministrativeStatus() bool`

HasAdministrativeStatus returns a boolean if a field has been set.

### GetAddress

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) GetAddress() AFNORAddressRead`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) GetAddressOk() (*AFNORAddressRead, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) SetAddress(v AFNORAddressRead)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLegalUnit

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) GetLegalUnit() AFNORLegalUnitPayloadIncluded`

GetLegalUnit returns the LegalUnit field if non-nil, zero value otherwise.

### GetLegalUnitOk

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) GetLegalUnitOk() (*AFNORLegalUnitPayloadIncluded, bool)`

GetLegalUnitOk returns a tuple with the LegalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalUnit

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) SetLegalUnit(v AFNORLegalUnitPayloadIncluded)`

SetLegalUnit sets LegalUnit field to given value.

### HasLegalUnit

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) HasLegalUnit() bool`

HasLegalUnit returns a boolean if a field has been set.

### GetFacility

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) GetFacility() AFNORFacilityPayloadIncluded`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) GetFacilityOk() (*AFNORFacilityPayloadIncluded, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) SetFacility(v AFNORFacilityPayloadIncluded)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *AFNORRoutingCodePayloadHistoryLegalUnitFacility) HasFacility() bool`

HasFacility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



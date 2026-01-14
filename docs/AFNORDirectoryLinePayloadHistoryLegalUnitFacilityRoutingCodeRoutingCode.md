# AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoutingIdentifier** | Pointer to **string** | Routing identifier od a routing code. | [optional] 
**RoutingIdentifierType** | Pointer to **string** | Routing Identifier type. | [optional] 
**RoutingCodeName** | Pointer to **string** | Name of the directory line routing code. This attribute is only returned if the directory line is defined at the SIREN / SIRET / Routing code mesh. | [optional] 
**ManagesLegalCommitment** | Pointer to **bool** | Indicates whether the public structure requires a legal commitment number. This attribute is only returned if the directory line is defined for a public structure at the SIREN / SIRET or SIREN / SIRET / Routing code level. | [optional] 
**AdministrativeStatus** | Pointer to [**AFNORRoutingCodeAdministrativeStatus**](AFNORRoutingCodeAdministrativeStatus.md) |  | [optional] 
**Address** | Pointer to [**AFNORAddressRead**](AFNORAddressRead.md) |  | [optional] 

## Methods

### NewAFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode

`func NewAFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode() *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode`

NewAFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode instantiates a new AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCodeWithDefaults

`func NewAFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCodeWithDefaults() *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode`

NewAFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCodeWithDefaults instantiates a new AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutingIdentifier

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) GetRoutingIdentifier() string`

GetRoutingIdentifier returns the RoutingIdentifier field if non-nil, zero value otherwise.

### GetRoutingIdentifierOk

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) GetRoutingIdentifierOk() (*string, bool)`

GetRoutingIdentifierOk returns a tuple with the RoutingIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIdentifier

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) SetRoutingIdentifier(v string)`

SetRoutingIdentifier sets RoutingIdentifier field to given value.

### HasRoutingIdentifier

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) HasRoutingIdentifier() bool`

HasRoutingIdentifier returns a boolean if a field has been set.

### GetRoutingIdentifierType

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) GetRoutingIdentifierType() string`

GetRoutingIdentifierType returns the RoutingIdentifierType field if non-nil, zero value otherwise.

### GetRoutingIdentifierTypeOk

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) GetRoutingIdentifierTypeOk() (*string, bool)`

GetRoutingIdentifierTypeOk returns a tuple with the RoutingIdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIdentifierType

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) SetRoutingIdentifierType(v string)`

SetRoutingIdentifierType sets RoutingIdentifierType field to given value.

### HasRoutingIdentifierType

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) HasRoutingIdentifierType() bool`

HasRoutingIdentifierType returns a boolean if a field has been set.

### GetRoutingCodeName

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) GetRoutingCodeName() string`

GetRoutingCodeName returns the RoutingCodeName field if non-nil, zero value otherwise.

### GetRoutingCodeNameOk

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) GetRoutingCodeNameOk() (*string, bool)`

GetRoutingCodeNameOk returns a tuple with the RoutingCodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingCodeName

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) SetRoutingCodeName(v string)`

SetRoutingCodeName sets RoutingCodeName field to given value.

### HasRoutingCodeName

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) HasRoutingCodeName() bool`

HasRoutingCodeName returns a boolean if a field has been set.

### GetManagesLegalCommitment

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) GetManagesLegalCommitment() bool`

GetManagesLegalCommitment returns the ManagesLegalCommitment field if non-nil, zero value otherwise.

### GetManagesLegalCommitmentOk

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) GetManagesLegalCommitmentOk() (*bool, bool)`

GetManagesLegalCommitmentOk returns a tuple with the ManagesLegalCommitment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagesLegalCommitment

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) SetManagesLegalCommitment(v bool)`

SetManagesLegalCommitment sets ManagesLegalCommitment field to given value.

### HasManagesLegalCommitment

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) HasManagesLegalCommitment() bool`

HasManagesLegalCommitment returns a boolean if a field has been set.

### GetAdministrativeStatus

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) GetAdministrativeStatus() AFNORRoutingCodeAdministrativeStatus`

GetAdministrativeStatus returns the AdministrativeStatus field if non-nil, zero value otherwise.

### GetAdministrativeStatusOk

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) GetAdministrativeStatusOk() (*AFNORRoutingCodeAdministrativeStatus, bool)`

GetAdministrativeStatusOk returns a tuple with the AdministrativeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeStatus

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) SetAdministrativeStatus(v AFNORRoutingCodeAdministrativeStatus)`

SetAdministrativeStatus sets AdministrativeStatus field to given value.

### HasAdministrativeStatus

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) HasAdministrativeStatus() bool`

HasAdministrativeStatus returns a boolean if a field has been set.

### GetAddress

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) GetAddress() AFNORAddressRead`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) GetAddressOk() (*AFNORAddressRead, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) SetAddress(v AFNORAddressRead)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AFNORDirectoryLinePayloadHistoryLegalUnitFacilityRoutingCodeRoutingCode) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AFNORCreateRoutingCodeBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FacilityNature** | [**AFNORFacilityNature**](AFNORFacilityNature.md) |  | 
**RoutingIdentifier** | **string** | Routing identifier od a routing code. | 
**Siret** | **string** | SIRET Number | 
**RoutingIdentifierType** | Pointer to **string** | Routing Identifier type. | [optional] 
**RoutingCodeName** | **string** | Name of the directory line routing code. This attribute is only returned if the directory line is defined at the SIREN / SIRET / Routing code mesh. | 
**ManagesLegalCommitmentCode** | Pointer to **bool** | Indicates whether the public structure requires a legal commitment number. This attribute is only returned if the directory line is defined for a public structure at the SIREN / SIRET or SIREN / SIRET / Routing code level. | [optional] 
**AdministrativeStatus** | [**AFNORRoutingCodeAdministrativeStatus**](AFNORRoutingCodeAdministrativeStatus.md) |  | 
**Address** | Pointer to [**AFNORAddressEdit**](AFNORAddressEdit.md) |  | [optional] 

## Methods

### NewAFNORCreateRoutingCodeBody

`func NewAFNORCreateRoutingCodeBody(facilityNature AFNORFacilityNature, routingIdentifier string, siret string, routingCodeName string, administrativeStatus AFNORRoutingCodeAdministrativeStatus, ) *AFNORCreateRoutingCodeBody`

NewAFNORCreateRoutingCodeBody instantiates a new AFNORCreateRoutingCodeBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORCreateRoutingCodeBodyWithDefaults

`func NewAFNORCreateRoutingCodeBodyWithDefaults() *AFNORCreateRoutingCodeBody`

NewAFNORCreateRoutingCodeBodyWithDefaults instantiates a new AFNORCreateRoutingCodeBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacilityNature

`func (o *AFNORCreateRoutingCodeBody) GetFacilityNature() AFNORFacilityNature`

GetFacilityNature returns the FacilityNature field if non-nil, zero value otherwise.

### GetFacilityNatureOk

`func (o *AFNORCreateRoutingCodeBody) GetFacilityNatureOk() (*AFNORFacilityNature, bool)`

GetFacilityNatureOk returns a tuple with the FacilityNature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityNature

`func (o *AFNORCreateRoutingCodeBody) SetFacilityNature(v AFNORFacilityNature)`

SetFacilityNature sets FacilityNature field to given value.


### GetRoutingIdentifier

`func (o *AFNORCreateRoutingCodeBody) GetRoutingIdentifier() string`

GetRoutingIdentifier returns the RoutingIdentifier field if non-nil, zero value otherwise.

### GetRoutingIdentifierOk

`func (o *AFNORCreateRoutingCodeBody) GetRoutingIdentifierOk() (*string, bool)`

GetRoutingIdentifierOk returns a tuple with the RoutingIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIdentifier

`func (o *AFNORCreateRoutingCodeBody) SetRoutingIdentifier(v string)`

SetRoutingIdentifier sets RoutingIdentifier field to given value.


### GetSiret

`func (o *AFNORCreateRoutingCodeBody) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *AFNORCreateRoutingCodeBody) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *AFNORCreateRoutingCodeBody) SetSiret(v string)`

SetSiret sets Siret field to given value.


### GetRoutingIdentifierType

`func (o *AFNORCreateRoutingCodeBody) GetRoutingIdentifierType() string`

GetRoutingIdentifierType returns the RoutingIdentifierType field if non-nil, zero value otherwise.

### GetRoutingIdentifierTypeOk

`func (o *AFNORCreateRoutingCodeBody) GetRoutingIdentifierTypeOk() (*string, bool)`

GetRoutingIdentifierTypeOk returns a tuple with the RoutingIdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIdentifierType

`func (o *AFNORCreateRoutingCodeBody) SetRoutingIdentifierType(v string)`

SetRoutingIdentifierType sets RoutingIdentifierType field to given value.

### HasRoutingIdentifierType

`func (o *AFNORCreateRoutingCodeBody) HasRoutingIdentifierType() bool`

HasRoutingIdentifierType returns a boolean if a field has been set.

### GetRoutingCodeName

`func (o *AFNORCreateRoutingCodeBody) GetRoutingCodeName() string`

GetRoutingCodeName returns the RoutingCodeName field if non-nil, zero value otherwise.

### GetRoutingCodeNameOk

`func (o *AFNORCreateRoutingCodeBody) GetRoutingCodeNameOk() (*string, bool)`

GetRoutingCodeNameOk returns a tuple with the RoutingCodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingCodeName

`func (o *AFNORCreateRoutingCodeBody) SetRoutingCodeName(v string)`

SetRoutingCodeName sets RoutingCodeName field to given value.


### GetManagesLegalCommitmentCode

`func (o *AFNORCreateRoutingCodeBody) GetManagesLegalCommitmentCode() bool`

GetManagesLegalCommitmentCode returns the ManagesLegalCommitmentCode field if non-nil, zero value otherwise.

### GetManagesLegalCommitmentCodeOk

`func (o *AFNORCreateRoutingCodeBody) GetManagesLegalCommitmentCodeOk() (*bool, bool)`

GetManagesLegalCommitmentCodeOk returns a tuple with the ManagesLegalCommitmentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagesLegalCommitmentCode

`func (o *AFNORCreateRoutingCodeBody) SetManagesLegalCommitmentCode(v bool)`

SetManagesLegalCommitmentCode sets ManagesLegalCommitmentCode field to given value.

### HasManagesLegalCommitmentCode

`func (o *AFNORCreateRoutingCodeBody) HasManagesLegalCommitmentCode() bool`

HasManagesLegalCommitmentCode returns a boolean if a field has been set.

### GetAdministrativeStatus

`func (o *AFNORCreateRoutingCodeBody) GetAdministrativeStatus() AFNORRoutingCodeAdministrativeStatus`

GetAdministrativeStatus returns the AdministrativeStatus field if non-nil, zero value otherwise.

### GetAdministrativeStatusOk

`func (o *AFNORCreateRoutingCodeBody) GetAdministrativeStatusOk() (*AFNORRoutingCodeAdministrativeStatus, bool)`

GetAdministrativeStatusOk returns a tuple with the AdministrativeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeStatus

`func (o *AFNORCreateRoutingCodeBody) SetAdministrativeStatus(v AFNORRoutingCodeAdministrativeStatus)`

SetAdministrativeStatus sets AdministrativeStatus field to given value.


### GetAddress

`func (o *AFNORCreateRoutingCodeBody) GetAddress() AFNORAddressEdit`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AFNORCreateRoutingCodeBody) GetAddressOk() (*AFNORAddressEdit, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AFNORCreateRoutingCodeBody) SetAddress(v AFNORAddressEdit)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AFNORCreateRoutingCodeBody) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



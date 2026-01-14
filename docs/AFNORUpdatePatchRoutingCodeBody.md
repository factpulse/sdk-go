# AFNORUpdatePatchRoutingCodeBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoutingIdentifierType** | Pointer to **string** | Routing Identifier type. | [optional] 
**RoutingCodeName** | Pointer to **string** | Name of the directory line routing code. This attribute is only returned if the directory line is defined at the SIREN / SIRET / Routing code mesh. | [optional] 
**AdministrativeStatus** | Pointer to [**AFNORRoutingCodeAdministrativeStatus**](AFNORRoutingCodeAdministrativeStatus.md) |  | [optional] 
**Address** | Pointer to [**AFNORAddressPatch**](AFNORAddressPatch.md) |  | [optional] 

## Methods

### NewAFNORUpdatePatchRoutingCodeBody

`func NewAFNORUpdatePatchRoutingCodeBody() *AFNORUpdatePatchRoutingCodeBody`

NewAFNORUpdatePatchRoutingCodeBody instantiates a new AFNORUpdatePatchRoutingCodeBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORUpdatePatchRoutingCodeBodyWithDefaults

`func NewAFNORUpdatePatchRoutingCodeBodyWithDefaults() *AFNORUpdatePatchRoutingCodeBody`

NewAFNORUpdatePatchRoutingCodeBodyWithDefaults instantiates a new AFNORUpdatePatchRoutingCodeBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutingIdentifierType

`func (o *AFNORUpdatePatchRoutingCodeBody) GetRoutingIdentifierType() string`

GetRoutingIdentifierType returns the RoutingIdentifierType field if non-nil, zero value otherwise.

### GetRoutingIdentifierTypeOk

`func (o *AFNORUpdatePatchRoutingCodeBody) GetRoutingIdentifierTypeOk() (*string, bool)`

GetRoutingIdentifierTypeOk returns a tuple with the RoutingIdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIdentifierType

`func (o *AFNORUpdatePatchRoutingCodeBody) SetRoutingIdentifierType(v string)`

SetRoutingIdentifierType sets RoutingIdentifierType field to given value.

### HasRoutingIdentifierType

`func (o *AFNORUpdatePatchRoutingCodeBody) HasRoutingIdentifierType() bool`

HasRoutingIdentifierType returns a boolean if a field has been set.

### GetRoutingCodeName

`func (o *AFNORUpdatePatchRoutingCodeBody) GetRoutingCodeName() string`

GetRoutingCodeName returns the RoutingCodeName field if non-nil, zero value otherwise.

### GetRoutingCodeNameOk

`func (o *AFNORUpdatePatchRoutingCodeBody) GetRoutingCodeNameOk() (*string, bool)`

GetRoutingCodeNameOk returns a tuple with the RoutingCodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingCodeName

`func (o *AFNORUpdatePatchRoutingCodeBody) SetRoutingCodeName(v string)`

SetRoutingCodeName sets RoutingCodeName field to given value.

### HasRoutingCodeName

`func (o *AFNORUpdatePatchRoutingCodeBody) HasRoutingCodeName() bool`

HasRoutingCodeName returns a boolean if a field has been set.

### GetAdministrativeStatus

`func (o *AFNORUpdatePatchRoutingCodeBody) GetAdministrativeStatus() AFNORRoutingCodeAdministrativeStatus`

GetAdministrativeStatus returns the AdministrativeStatus field if non-nil, zero value otherwise.

### GetAdministrativeStatusOk

`func (o *AFNORUpdatePatchRoutingCodeBody) GetAdministrativeStatusOk() (*AFNORRoutingCodeAdministrativeStatus, bool)`

GetAdministrativeStatusOk returns a tuple with the AdministrativeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeStatus

`func (o *AFNORUpdatePatchRoutingCodeBody) SetAdministrativeStatus(v AFNORRoutingCodeAdministrativeStatus)`

SetAdministrativeStatus sets AdministrativeStatus field to given value.

### HasAdministrativeStatus

`func (o *AFNORUpdatePatchRoutingCodeBody) HasAdministrativeStatus() bool`

HasAdministrativeStatus returns a boolean if a field has been set.

### GetAddress

`func (o *AFNORUpdatePatchRoutingCodeBody) GetAddress() AFNORAddressPatch`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AFNORUpdatePatchRoutingCodeBody) GetAddressOk() (*AFNORAddressPatch, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AFNORUpdatePatchRoutingCodeBody) SetAddress(v AFNORAddressPatch)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AFNORUpdatePatchRoutingCodeBody) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



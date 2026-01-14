# AFNORUpdatePutRoutingCodeBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoutingIdentifierType** | **string** | Routing Identifier type. | 
**RoutingCodeName** | **string** | Name of the directory line routing code. This attribute is only returned if the directory line is defined at the SIREN / SIRET / Routing code mesh. | 
**AdministrativeStatus** | [**AFNORRoutingCodeAdministrativeStatus**](AFNORRoutingCodeAdministrativeStatus.md) |  | 
**Address** | Pointer to [**AFNORAddressPut**](AFNORAddressPut.md) |  | [optional] 

## Methods

### NewAFNORUpdatePutRoutingCodeBody

`func NewAFNORUpdatePutRoutingCodeBody(routingIdentifierType string, routingCodeName string, administrativeStatus AFNORRoutingCodeAdministrativeStatus, ) *AFNORUpdatePutRoutingCodeBody`

NewAFNORUpdatePutRoutingCodeBody instantiates a new AFNORUpdatePutRoutingCodeBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORUpdatePutRoutingCodeBodyWithDefaults

`func NewAFNORUpdatePutRoutingCodeBodyWithDefaults() *AFNORUpdatePutRoutingCodeBody`

NewAFNORUpdatePutRoutingCodeBodyWithDefaults instantiates a new AFNORUpdatePutRoutingCodeBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutingIdentifierType

`func (o *AFNORUpdatePutRoutingCodeBody) GetRoutingIdentifierType() string`

GetRoutingIdentifierType returns the RoutingIdentifierType field if non-nil, zero value otherwise.

### GetRoutingIdentifierTypeOk

`func (o *AFNORUpdatePutRoutingCodeBody) GetRoutingIdentifierTypeOk() (*string, bool)`

GetRoutingIdentifierTypeOk returns a tuple with the RoutingIdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIdentifierType

`func (o *AFNORUpdatePutRoutingCodeBody) SetRoutingIdentifierType(v string)`

SetRoutingIdentifierType sets RoutingIdentifierType field to given value.


### GetRoutingCodeName

`func (o *AFNORUpdatePutRoutingCodeBody) GetRoutingCodeName() string`

GetRoutingCodeName returns the RoutingCodeName field if non-nil, zero value otherwise.

### GetRoutingCodeNameOk

`func (o *AFNORUpdatePutRoutingCodeBody) GetRoutingCodeNameOk() (*string, bool)`

GetRoutingCodeNameOk returns a tuple with the RoutingCodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingCodeName

`func (o *AFNORUpdatePutRoutingCodeBody) SetRoutingCodeName(v string)`

SetRoutingCodeName sets RoutingCodeName field to given value.


### GetAdministrativeStatus

`func (o *AFNORUpdatePutRoutingCodeBody) GetAdministrativeStatus() AFNORRoutingCodeAdministrativeStatus`

GetAdministrativeStatus returns the AdministrativeStatus field if non-nil, zero value otherwise.

### GetAdministrativeStatusOk

`func (o *AFNORUpdatePutRoutingCodeBody) GetAdministrativeStatusOk() (*AFNORRoutingCodeAdministrativeStatus, bool)`

GetAdministrativeStatusOk returns a tuple with the AdministrativeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeStatus

`func (o *AFNORUpdatePutRoutingCodeBody) SetAdministrativeStatus(v AFNORRoutingCodeAdministrativeStatus)`

SetAdministrativeStatus sets AdministrativeStatus field to given value.


### GetAddress

`func (o *AFNORUpdatePutRoutingCodeBody) GetAddress() AFNORAddressPut`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AFNORUpdatePutRoutingCodeBody) GetAddressOk() (*AFNORAddressPut, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AFNORUpdatePutRoutingCodeBody) SetAddress(v AFNORAddressPut)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AFNORUpdatePutRoutingCodeBody) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



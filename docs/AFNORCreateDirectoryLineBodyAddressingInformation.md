# AFNORCreateDirectoryLineBodyAddressingInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Siren** | **string** | SIREN number | 
**Siret** | Pointer to **string** | SIRET Number | [optional] 
**RoutingIdentifier** | Pointer to **string** | Routing identifier od a routing code. | [optional] 
**AddressingSuffix** | Pointer to **string** | suffix of the directory line which defines an address mesh not attached to a facility | [optional] 

## Methods

### NewAFNORCreateDirectoryLineBodyAddressingInformation

`func NewAFNORCreateDirectoryLineBodyAddressingInformation(siren string, ) *AFNORCreateDirectoryLineBodyAddressingInformation`

NewAFNORCreateDirectoryLineBodyAddressingInformation instantiates a new AFNORCreateDirectoryLineBodyAddressingInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORCreateDirectoryLineBodyAddressingInformationWithDefaults

`func NewAFNORCreateDirectoryLineBodyAddressingInformationWithDefaults() *AFNORCreateDirectoryLineBodyAddressingInformation`

NewAFNORCreateDirectoryLineBodyAddressingInformationWithDefaults instantiates a new AFNORCreateDirectoryLineBodyAddressingInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiren

`func (o *AFNORCreateDirectoryLineBodyAddressingInformation) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *AFNORCreateDirectoryLineBodyAddressingInformation) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *AFNORCreateDirectoryLineBodyAddressingInformation) SetSiren(v string)`

SetSiren sets Siren field to given value.


### GetSiret

`func (o *AFNORCreateDirectoryLineBodyAddressingInformation) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *AFNORCreateDirectoryLineBodyAddressingInformation) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *AFNORCreateDirectoryLineBodyAddressingInformation) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *AFNORCreateDirectoryLineBodyAddressingInformation) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### GetRoutingIdentifier

`func (o *AFNORCreateDirectoryLineBodyAddressingInformation) GetRoutingIdentifier() string`

GetRoutingIdentifier returns the RoutingIdentifier field if non-nil, zero value otherwise.

### GetRoutingIdentifierOk

`func (o *AFNORCreateDirectoryLineBodyAddressingInformation) GetRoutingIdentifierOk() (*string, bool)`

GetRoutingIdentifierOk returns a tuple with the RoutingIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIdentifier

`func (o *AFNORCreateDirectoryLineBodyAddressingInformation) SetRoutingIdentifier(v string)`

SetRoutingIdentifier sets RoutingIdentifier field to given value.

### HasRoutingIdentifier

`func (o *AFNORCreateDirectoryLineBodyAddressingInformation) HasRoutingIdentifier() bool`

HasRoutingIdentifier returns a boolean if a field has been set.

### GetAddressingSuffix

`func (o *AFNORCreateDirectoryLineBodyAddressingInformation) GetAddressingSuffix() string`

GetAddressingSuffix returns the AddressingSuffix field if non-nil, zero value otherwise.

### GetAddressingSuffixOk

`func (o *AFNORCreateDirectoryLineBodyAddressingInformation) GetAddressingSuffixOk() (*string, bool)`

GetAddressingSuffixOk returns a tuple with the AddressingSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressingSuffix

`func (o *AFNORCreateDirectoryLineBodyAddressingInformation) SetAddressingSuffix(v string)`

SetAddressingSuffix sets AddressingSuffix field to given value.

### HasAddressingSuffix

`func (o *AFNORCreateDirectoryLineBodyAddressingInformation) HasAddressingSuffix() bool`

HasAddressingSuffix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



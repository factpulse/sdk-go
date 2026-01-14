# AFNORLegalUnitPayloadIncluded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Siren** | Pointer to **string** | SIREN number | [optional] 
**BusinessName** | Pointer to **string** | Business name | [optional] 
**EntityType** | Pointer to [**AFNOREntityType**](AFNOREntityType.md) |  | [optional] 
**AdministrativeStatus** | Pointer to [**AFNORLegalUnitAdministrativeStatus**](AFNORLegalUnitAdministrativeStatus.md) |  | [optional] 

## Methods

### NewAFNORLegalUnitPayloadIncluded

`func NewAFNORLegalUnitPayloadIncluded() *AFNORLegalUnitPayloadIncluded`

NewAFNORLegalUnitPayloadIncluded instantiates a new AFNORLegalUnitPayloadIncluded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORLegalUnitPayloadIncludedWithDefaults

`func NewAFNORLegalUnitPayloadIncludedWithDefaults() *AFNORLegalUnitPayloadIncluded`

NewAFNORLegalUnitPayloadIncludedWithDefaults instantiates a new AFNORLegalUnitPayloadIncluded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiren

`func (o *AFNORLegalUnitPayloadIncluded) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *AFNORLegalUnitPayloadIncluded) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *AFNORLegalUnitPayloadIncluded) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *AFNORLegalUnitPayloadIncluded) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### GetBusinessName

`func (o *AFNORLegalUnitPayloadIncluded) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *AFNORLegalUnitPayloadIncluded) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *AFNORLegalUnitPayloadIncluded) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.

### HasBusinessName

`func (o *AFNORLegalUnitPayloadIncluded) HasBusinessName() bool`

HasBusinessName returns a boolean if a field has been set.

### GetEntityType

`func (o *AFNORLegalUnitPayloadIncluded) GetEntityType() AFNOREntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *AFNORLegalUnitPayloadIncluded) GetEntityTypeOk() (*AFNOREntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *AFNORLegalUnitPayloadIncluded) SetEntityType(v AFNOREntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *AFNORLegalUnitPayloadIncluded) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetAdministrativeStatus

`func (o *AFNORLegalUnitPayloadIncluded) GetAdministrativeStatus() AFNORLegalUnitAdministrativeStatus`

GetAdministrativeStatus returns the AdministrativeStatus field if non-nil, zero value otherwise.

### GetAdministrativeStatusOk

`func (o *AFNORLegalUnitPayloadIncluded) GetAdministrativeStatusOk() (*AFNORLegalUnitAdministrativeStatus, bool)`

GetAdministrativeStatusOk returns a tuple with the AdministrativeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeStatus

`func (o *AFNORLegalUnitPayloadIncluded) SetAdministrativeStatus(v AFNORLegalUnitAdministrativeStatus)`

SetAdministrativeStatus sets AdministrativeStatus field to given value.

### HasAdministrativeStatus

`func (o *AFNORLegalUnitPayloadIncluded) HasAdministrativeStatus() bool`

HasAdministrativeStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



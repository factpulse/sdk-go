# AFNORRoutingCodeSearchFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoutingIdentifier** | Pointer to [**AFNORRoutingCodeSearchFiltersRoutingIdentifier**](AFNORRoutingCodeSearchFiltersRoutingIdentifier.md) |  | [optional] 
**Siret** | Pointer to [**AFNORSearchSiretFiltersSiret**](AFNORSearchSiretFiltersSiret.md) |  | [optional] 
**RoutingCodeName** | Pointer to [**AFNORRoutingCodeSearchFiltersRoutingCodeName**](AFNORRoutingCodeSearchFiltersRoutingCodeName.md) |  | [optional] 
**AdministrativeStatus** | Pointer to [**AFNORRoutingCodeSearchFiltersAdministrativeStatus**](AFNORRoutingCodeSearchFiltersAdministrativeStatus.md) |  | [optional] 
**AddressLines** | Pointer to [**AFNORSearchSiretFiltersAddressLines**](AFNORSearchSiretFiltersAddressLines.md) |  | [optional] 
**PostalCode** | Pointer to [**AFNORSearchSiretFiltersPostalCode**](AFNORSearchSiretFiltersPostalCode.md) |  | [optional] 
**Locality** | Pointer to [**AFNORSearchSiretFiltersLocality**](AFNORSearchSiretFiltersLocality.md) |  | [optional] 

## Methods

### NewAFNORRoutingCodeSearchFilters

`func NewAFNORRoutingCodeSearchFilters() *AFNORRoutingCodeSearchFilters`

NewAFNORRoutingCodeSearchFilters instantiates a new AFNORRoutingCodeSearchFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORRoutingCodeSearchFiltersWithDefaults

`func NewAFNORRoutingCodeSearchFiltersWithDefaults() *AFNORRoutingCodeSearchFilters`

NewAFNORRoutingCodeSearchFiltersWithDefaults instantiates a new AFNORRoutingCodeSearchFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutingIdentifier

`func (o *AFNORRoutingCodeSearchFilters) GetRoutingIdentifier() AFNORRoutingCodeSearchFiltersRoutingIdentifier`

GetRoutingIdentifier returns the RoutingIdentifier field if non-nil, zero value otherwise.

### GetRoutingIdentifierOk

`func (o *AFNORRoutingCodeSearchFilters) GetRoutingIdentifierOk() (*AFNORRoutingCodeSearchFiltersRoutingIdentifier, bool)`

GetRoutingIdentifierOk returns a tuple with the RoutingIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIdentifier

`func (o *AFNORRoutingCodeSearchFilters) SetRoutingIdentifier(v AFNORRoutingCodeSearchFiltersRoutingIdentifier)`

SetRoutingIdentifier sets RoutingIdentifier field to given value.

### HasRoutingIdentifier

`func (o *AFNORRoutingCodeSearchFilters) HasRoutingIdentifier() bool`

HasRoutingIdentifier returns a boolean if a field has been set.

### GetSiret

`func (o *AFNORRoutingCodeSearchFilters) GetSiret() AFNORSearchSiretFiltersSiret`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *AFNORRoutingCodeSearchFilters) GetSiretOk() (*AFNORSearchSiretFiltersSiret, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *AFNORRoutingCodeSearchFilters) SetSiret(v AFNORSearchSiretFiltersSiret)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *AFNORRoutingCodeSearchFilters) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### GetRoutingCodeName

`func (o *AFNORRoutingCodeSearchFilters) GetRoutingCodeName() AFNORRoutingCodeSearchFiltersRoutingCodeName`

GetRoutingCodeName returns the RoutingCodeName field if non-nil, zero value otherwise.

### GetRoutingCodeNameOk

`func (o *AFNORRoutingCodeSearchFilters) GetRoutingCodeNameOk() (*AFNORRoutingCodeSearchFiltersRoutingCodeName, bool)`

GetRoutingCodeNameOk returns a tuple with the RoutingCodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingCodeName

`func (o *AFNORRoutingCodeSearchFilters) SetRoutingCodeName(v AFNORRoutingCodeSearchFiltersRoutingCodeName)`

SetRoutingCodeName sets RoutingCodeName field to given value.

### HasRoutingCodeName

`func (o *AFNORRoutingCodeSearchFilters) HasRoutingCodeName() bool`

HasRoutingCodeName returns a boolean if a field has been set.

### GetAdministrativeStatus

`func (o *AFNORRoutingCodeSearchFilters) GetAdministrativeStatus() AFNORRoutingCodeSearchFiltersAdministrativeStatus`

GetAdministrativeStatus returns the AdministrativeStatus field if non-nil, zero value otherwise.

### GetAdministrativeStatusOk

`func (o *AFNORRoutingCodeSearchFilters) GetAdministrativeStatusOk() (*AFNORRoutingCodeSearchFiltersAdministrativeStatus, bool)`

GetAdministrativeStatusOk returns a tuple with the AdministrativeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeStatus

`func (o *AFNORRoutingCodeSearchFilters) SetAdministrativeStatus(v AFNORRoutingCodeSearchFiltersAdministrativeStatus)`

SetAdministrativeStatus sets AdministrativeStatus field to given value.

### HasAdministrativeStatus

`func (o *AFNORRoutingCodeSearchFilters) HasAdministrativeStatus() bool`

HasAdministrativeStatus returns a boolean if a field has been set.

### GetAddressLines

`func (o *AFNORRoutingCodeSearchFilters) GetAddressLines() AFNORSearchSiretFiltersAddressLines`

GetAddressLines returns the AddressLines field if non-nil, zero value otherwise.

### GetAddressLinesOk

`func (o *AFNORRoutingCodeSearchFilters) GetAddressLinesOk() (*AFNORSearchSiretFiltersAddressLines, bool)`

GetAddressLinesOk returns a tuple with the AddressLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLines

`func (o *AFNORRoutingCodeSearchFilters) SetAddressLines(v AFNORSearchSiretFiltersAddressLines)`

SetAddressLines sets AddressLines field to given value.

### HasAddressLines

`func (o *AFNORRoutingCodeSearchFilters) HasAddressLines() bool`

HasAddressLines returns a boolean if a field has been set.

### GetPostalCode

`func (o *AFNORRoutingCodeSearchFilters) GetPostalCode() AFNORSearchSiretFiltersPostalCode`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AFNORRoutingCodeSearchFilters) GetPostalCodeOk() (*AFNORSearchSiretFiltersPostalCode, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AFNORRoutingCodeSearchFilters) SetPostalCode(v AFNORSearchSiretFiltersPostalCode)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *AFNORRoutingCodeSearchFilters) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetLocality

`func (o *AFNORRoutingCodeSearchFilters) GetLocality() AFNORSearchSiretFiltersLocality`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *AFNORRoutingCodeSearchFilters) GetLocalityOk() (*AFNORSearchSiretFiltersLocality, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *AFNORRoutingCodeSearchFilters) SetLocality(v AFNORSearchSiretFiltersLocality)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *AFNORRoutingCodeSearchFilters) HasLocality() bool`

HasLocality returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



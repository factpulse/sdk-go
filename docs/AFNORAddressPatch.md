# AFNORAddressPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LigneAdresse1** | Pointer to **string** | Corresponds to the address of the recipient structure having defined the directory line(s). | [optional] 
**LigneAdresse2** | Pointer to **string** | Corresponds to the address of the recipient structure having defined the directory line(s). | [optional] 
**LigneAdresse3** | Pointer to **string** | Corresponds to the address of the recipient structure having defined the directory line(s). | [optional] 
**PostalCode** | Pointer to **string** | Service postal code | [optional] 
**CountrySubdivision** | Pointer to **string** | Subdivision of the country | [optional] 
**Locality** | Pointer to **string** | Corresponds to the municipality of the recipient structure having defined the directory line(s). | [optional] 
**CodePays** | Pointer to **string** | Corresponds to the country code of the recipient structure | [optional] 

## Methods

### NewAFNORAddressPatch

`func NewAFNORAddressPatch() *AFNORAddressPatch`

NewAFNORAddressPatch instantiates a new AFNORAddressPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORAddressPatchWithDefaults

`func NewAFNORAddressPatchWithDefaults() *AFNORAddressPatch`

NewAFNORAddressPatchWithDefaults instantiates a new AFNORAddressPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLigneAdresse1

`func (o *AFNORAddressPatch) GetLigneAdresse1() string`

GetLigneAdresse1 returns the LigneAdresse1 field if non-nil, zero value otherwise.

### GetLigneAdresse1Ok

`func (o *AFNORAddressPatch) GetLigneAdresse1Ok() (*string, bool)`

GetLigneAdresse1Ok returns a tuple with the LigneAdresse1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLigneAdresse1

`func (o *AFNORAddressPatch) SetLigneAdresse1(v string)`

SetLigneAdresse1 sets LigneAdresse1 field to given value.

### HasLigneAdresse1

`func (o *AFNORAddressPatch) HasLigneAdresse1() bool`

HasLigneAdresse1 returns a boolean if a field has been set.

### GetLigneAdresse2

`func (o *AFNORAddressPatch) GetLigneAdresse2() string`

GetLigneAdresse2 returns the LigneAdresse2 field if non-nil, zero value otherwise.

### GetLigneAdresse2Ok

`func (o *AFNORAddressPatch) GetLigneAdresse2Ok() (*string, bool)`

GetLigneAdresse2Ok returns a tuple with the LigneAdresse2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLigneAdresse2

`func (o *AFNORAddressPatch) SetLigneAdresse2(v string)`

SetLigneAdresse2 sets LigneAdresse2 field to given value.

### HasLigneAdresse2

`func (o *AFNORAddressPatch) HasLigneAdresse2() bool`

HasLigneAdresse2 returns a boolean if a field has been set.

### GetLigneAdresse3

`func (o *AFNORAddressPatch) GetLigneAdresse3() string`

GetLigneAdresse3 returns the LigneAdresse3 field if non-nil, zero value otherwise.

### GetLigneAdresse3Ok

`func (o *AFNORAddressPatch) GetLigneAdresse3Ok() (*string, bool)`

GetLigneAdresse3Ok returns a tuple with the LigneAdresse3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLigneAdresse3

`func (o *AFNORAddressPatch) SetLigneAdresse3(v string)`

SetLigneAdresse3 sets LigneAdresse3 field to given value.

### HasLigneAdresse3

`func (o *AFNORAddressPatch) HasLigneAdresse3() bool`

HasLigneAdresse3 returns a boolean if a field has been set.

### GetPostalCode

`func (o *AFNORAddressPatch) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AFNORAddressPatch) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AFNORAddressPatch) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *AFNORAddressPatch) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountrySubdivision

`func (o *AFNORAddressPatch) GetCountrySubdivision() string`

GetCountrySubdivision returns the CountrySubdivision field if non-nil, zero value otherwise.

### GetCountrySubdivisionOk

`func (o *AFNORAddressPatch) GetCountrySubdivisionOk() (*string, bool)`

GetCountrySubdivisionOk returns a tuple with the CountrySubdivision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountrySubdivision

`func (o *AFNORAddressPatch) SetCountrySubdivision(v string)`

SetCountrySubdivision sets CountrySubdivision field to given value.

### HasCountrySubdivision

`func (o *AFNORAddressPatch) HasCountrySubdivision() bool`

HasCountrySubdivision returns a boolean if a field has been set.

### GetLocality

`func (o *AFNORAddressPatch) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *AFNORAddressPatch) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *AFNORAddressPatch) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *AFNORAddressPatch) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetCodePays

`func (o *AFNORAddressPatch) GetCodePays() string`

GetCodePays returns the CodePays field if non-nil, zero value otherwise.

### GetCodePaysOk

`func (o *AFNORAddressPatch) GetCodePaysOk() (*string, bool)`

GetCodePaysOk returns a tuple with the CodePays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePays

`func (o *AFNORAddressPatch) SetCodePays(v string)`

SetCodePays sets CodePays field to given value.

### HasCodePays

`func (o *AFNORAddressPatch) HasCodePays() bool`

HasCodePays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



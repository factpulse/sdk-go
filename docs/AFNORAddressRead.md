# AFNORAddressRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | Pointer to **string** | Corresponds to the address of the recipient structure having defined the directory line(s). | [optional] 
**AddressLine2** | Pointer to **string** | Corresponds to the address of the recipient structure having defined the directory line(s). | [optional] 
**AddressLine3** | Pointer to **string** | Corresponds to the address of the recipient structure having defined the directory line(s). | [optional] 
**PostalCode** | Pointer to **string** | Service postal code | [optional] 
**CountrySubdivision** | Pointer to **string** | Subdivision of the country | [optional] 
**Locality** | Pointer to **string** | Municipality of the recipient structure having defined the directory line(s). | [optional] 
**CountryCode** | Pointer to **string** | Corresponds to the country of the recipient structure. | [optional] 
**CountryName** | Pointer to **string** | Corresponds to the country of the recipient structure having defined the directory line(s). | [optional] 

## Methods

### NewAFNORAddressRead

`func NewAFNORAddressRead() *AFNORAddressRead`

NewAFNORAddressRead instantiates a new AFNORAddressRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORAddressReadWithDefaults

`func NewAFNORAddressReadWithDefaults() *AFNORAddressRead`

NewAFNORAddressReadWithDefaults instantiates a new AFNORAddressRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *AFNORAddressRead) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *AFNORAddressRead) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *AFNORAddressRead) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *AFNORAddressRead) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *AFNORAddressRead) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *AFNORAddressRead) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *AFNORAddressRead) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *AFNORAddressRead) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressLine3

`func (o *AFNORAddressRead) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *AFNORAddressRead) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *AFNORAddressRead) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3

`func (o *AFNORAddressRead) HasAddressLine3() bool`

HasAddressLine3 returns a boolean if a field has been set.

### GetPostalCode

`func (o *AFNORAddressRead) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AFNORAddressRead) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AFNORAddressRead) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *AFNORAddressRead) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountrySubdivision

`func (o *AFNORAddressRead) GetCountrySubdivision() string`

GetCountrySubdivision returns the CountrySubdivision field if non-nil, zero value otherwise.

### GetCountrySubdivisionOk

`func (o *AFNORAddressRead) GetCountrySubdivisionOk() (*string, bool)`

GetCountrySubdivisionOk returns a tuple with the CountrySubdivision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountrySubdivision

`func (o *AFNORAddressRead) SetCountrySubdivision(v string)`

SetCountrySubdivision sets CountrySubdivision field to given value.

### HasCountrySubdivision

`func (o *AFNORAddressRead) HasCountrySubdivision() bool`

HasCountrySubdivision returns a boolean if a field has been set.

### GetLocality

`func (o *AFNORAddressRead) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *AFNORAddressRead) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *AFNORAddressRead) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *AFNORAddressRead) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetCountryCode

`func (o *AFNORAddressRead) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AFNORAddressRead) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AFNORAddressRead) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *AFNORAddressRead) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountryName

`func (o *AFNORAddressRead) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *AFNORAddressRead) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *AFNORAddressRead) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *AFNORAddressRead) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



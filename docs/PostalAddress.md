# PostalAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostalCode** | Pointer to **NullableString** |  | [optional] 
**LineOne** | Pointer to **NullableString** |  | [optional] 
**LineTwo** | Pointer to **NullableString** |  | [optional] 
**LineThree** | Pointer to **NullableString** |  | [optional] 
**City** | Pointer to **NullableString** |  | [optional] 
**CountryCode** | Pointer to **NullableString** |  | [optional] 
**CountrySubdivision** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPostalAddress

`func NewPostalAddress() *PostalAddress`

NewPostalAddress instantiates a new PostalAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostalAddressWithDefaults

`func NewPostalAddressWithDefaults() *PostalAddress`

NewPostalAddressWithDefaults instantiates a new PostalAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostalCode

`func (o *PostalAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PostalAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PostalAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *PostalAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *PostalAddress) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *PostalAddress) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetLineOne

`func (o *PostalAddress) GetLineOne() string`

GetLineOne returns the LineOne field if non-nil, zero value otherwise.

### GetLineOneOk

`func (o *PostalAddress) GetLineOneOk() (*string, bool)`

GetLineOneOk returns a tuple with the LineOne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineOne

`func (o *PostalAddress) SetLineOne(v string)`

SetLineOne sets LineOne field to given value.

### HasLineOne

`func (o *PostalAddress) HasLineOne() bool`

HasLineOne returns a boolean if a field has been set.

### SetLineOneNil

`func (o *PostalAddress) SetLineOneNil(b bool)`

 SetLineOneNil sets the value for LineOne to be an explicit nil

### UnsetLineOne
`func (o *PostalAddress) UnsetLineOne()`

UnsetLineOne ensures that no value is present for LineOne, not even an explicit nil
### GetLineTwo

`func (o *PostalAddress) GetLineTwo() string`

GetLineTwo returns the LineTwo field if non-nil, zero value otherwise.

### GetLineTwoOk

`func (o *PostalAddress) GetLineTwoOk() (*string, bool)`

GetLineTwoOk returns a tuple with the LineTwo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineTwo

`func (o *PostalAddress) SetLineTwo(v string)`

SetLineTwo sets LineTwo field to given value.

### HasLineTwo

`func (o *PostalAddress) HasLineTwo() bool`

HasLineTwo returns a boolean if a field has been set.

### SetLineTwoNil

`func (o *PostalAddress) SetLineTwoNil(b bool)`

 SetLineTwoNil sets the value for LineTwo to be an explicit nil

### UnsetLineTwo
`func (o *PostalAddress) UnsetLineTwo()`

UnsetLineTwo ensures that no value is present for LineTwo, not even an explicit nil
### GetLineThree

`func (o *PostalAddress) GetLineThree() string`

GetLineThree returns the LineThree field if non-nil, zero value otherwise.

### GetLineThreeOk

`func (o *PostalAddress) GetLineThreeOk() (*string, bool)`

GetLineThreeOk returns a tuple with the LineThree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineThree

`func (o *PostalAddress) SetLineThree(v string)`

SetLineThree sets LineThree field to given value.

### HasLineThree

`func (o *PostalAddress) HasLineThree() bool`

HasLineThree returns a boolean if a field has been set.

### SetLineThreeNil

`func (o *PostalAddress) SetLineThreeNil(b bool)`

 SetLineThreeNil sets the value for LineThree to be an explicit nil

### UnsetLineThree
`func (o *PostalAddress) UnsetLineThree()`

UnsetLineThree ensures that no value is present for LineThree, not even an explicit nil
### GetCity

`func (o *PostalAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PostalAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PostalAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *PostalAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *PostalAddress) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *PostalAddress) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountryCode

`func (o *PostalAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PostalAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PostalAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PostalAddress) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *PostalAddress) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *PostalAddress) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetCountrySubdivision

`func (o *PostalAddress) GetCountrySubdivision() string`

GetCountrySubdivision returns the CountrySubdivision field if non-nil, zero value otherwise.

### GetCountrySubdivisionOk

`func (o *PostalAddress) GetCountrySubdivisionOk() (*string, bool)`

GetCountrySubdivisionOk returns a tuple with the CountrySubdivision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountrySubdivision

`func (o *PostalAddress) SetCountrySubdivision(v string)`

SetCountrySubdivision sets CountrySubdivision field to given value.

### HasCountrySubdivision

`func (o *PostalAddress) HasCountrySubdivision() bool`

HasCountrySubdivision returns a boolean if a field has been set.

### SetCountrySubdivisionNil

`func (o *PostalAddress) SetCountrySubdivisionNil(b bool)`

 SetCountrySubdivisionNil sets the value for CountrySubdivision to be an explicit nil

### UnsetCountrySubdivision
`func (o *PostalAddress) UnsetCountrySubdivision()`

UnsetCountrySubdivision ensures that no value is present for CountrySubdivision, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TaxRepresentative

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Tax representative name (BT-62). | 
**VatNumber** | **string** | Tax representative VAT identifier (BT-63). | 
**PostalAddress** | [**PostalAddress**](PostalAddress.md) | Tax representative postal address (BG-12). | 

## Methods

### NewTaxRepresentative

`func NewTaxRepresentative(name string, vatNumber string, postalAddress PostalAddress, ) *TaxRepresentative`

NewTaxRepresentative instantiates a new TaxRepresentative object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxRepresentativeWithDefaults

`func NewTaxRepresentativeWithDefaults() *TaxRepresentative`

NewTaxRepresentativeWithDefaults instantiates a new TaxRepresentative object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TaxRepresentative) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxRepresentative) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxRepresentative) SetName(v string)`

SetName sets Name field to given value.


### GetVatNumber

`func (o *TaxRepresentative) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *TaxRepresentative) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *TaxRepresentative) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.


### GetPostalAddress

`func (o *TaxRepresentative) GetPostalAddress() PostalAddress`

GetPostalAddress returns the PostalAddress field if non-nil, zero value otherwise.

### GetPostalAddressOk

`func (o *TaxRepresentative) GetPostalAddressOk() (*PostalAddress, bool)`

GetPostalAddressOk returns a tuple with the PostalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalAddress

`func (o *TaxRepresentative) SetPostalAddress(v PostalAddress)`

SetPostalAddress sets PostalAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SuperPDPVatRegimeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VatRegime** | Pointer to **NullableString** |  | [optional] 
**HasVatOnDebits** | Pointer to **bool** | Whether VAT on debits is opted for | [optional] [default to false]

## Methods

### NewSuperPDPVatRegimeResponse

`func NewSuperPDPVatRegimeResponse() *SuperPDPVatRegimeResponse`

NewSuperPDPVatRegimeResponse instantiates a new SuperPDPVatRegimeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuperPDPVatRegimeResponseWithDefaults

`func NewSuperPDPVatRegimeResponseWithDefaults() *SuperPDPVatRegimeResponse`

NewSuperPDPVatRegimeResponseWithDefaults instantiates a new SuperPDPVatRegimeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVatRegime

`func (o *SuperPDPVatRegimeResponse) GetVatRegime() string`

GetVatRegime returns the VatRegime field if non-nil, zero value otherwise.

### GetVatRegimeOk

`func (o *SuperPDPVatRegimeResponse) GetVatRegimeOk() (*string, bool)`

GetVatRegimeOk returns a tuple with the VatRegime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatRegime

`func (o *SuperPDPVatRegimeResponse) SetVatRegime(v string)`

SetVatRegime sets VatRegime field to given value.

### HasVatRegime

`func (o *SuperPDPVatRegimeResponse) HasVatRegime() bool`

HasVatRegime returns a boolean if a field has been set.

### SetVatRegimeNil

`func (o *SuperPDPVatRegimeResponse) SetVatRegimeNil(b bool)`

 SetVatRegimeNil sets the value for VatRegime to be an explicit nil

### UnsetVatRegime
`func (o *SuperPDPVatRegimeResponse) UnsetVatRegime()`

UnsetVatRegime ensures that no value is present for VatRegime, not even an explicit nil
### GetHasVatOnDebits

`func (o *SuperPDPVatRegimeResponse) GetHasVatOnDebits() bool`

GetHasVatOnDebits returns the HasVatOnDebits field if non-nil, zero value otherwise.

### GetHasVatOnDebitsOk

`func (o *SuperPDPVatRegimeResponse) GetHasVatOnDebitsOk() (*bool, bool)`

GetHasVatOnDebitsOk returns a tuple with the HasVatOnDebits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVatOnDebits

`func (o *SuperPDPVatRegimeResponse) SetHasVatOnDebits(v bool)`

SetHasVatOnDebits sets HasVatOnDebits field to given value.

### HasHasVatOnDebits

`func (o *SuperPDPVatRegimeResponse) HasHasVatOnDebits() bool`

HasHasVatOnDebits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



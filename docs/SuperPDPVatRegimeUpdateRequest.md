# SuperPDPVatRegimeUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VatRegime** | **string** | monthly (réel normal mensuel), quarterly (réel normal trimestriel), simplified (régime simplifié), vat_exemption (franchise en base) | 
**HasVatOnDebits** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewSuperPDPVatRegimeUpdateRequest

`func NewSuperPDPVatRegimeUpdateRequest(vatRegime string, ) *SuperPDPVatRegimeUpdateRequest`

NewSuperPDPVatRegimeUpdateRequest instantiates a new SuperPDPVatRegimeUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuperPDPVatRegimeUpdateRequestWithDefaults

`func NewSuperPDPVatRegimeUpdateRequestWithDefaults() *SuperPDPVatRegimeUpdateRequest`

NewSuperPDPVatRegimeUpdateRequestWithDefaults instantiates a new SuperPDPVatRegimeUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVatRegime

`func (o *SuperPDPVatRegimeUpdateRequest) GetVatRegime() string`

GetVatRegime returns the VatRegime field if non-nil, zero value otherwise.

### GetVatRegimeOk

`func (o *SuperPDPVatRegimeUpdateRequest) GetVatRegimeOk() (*string, bool)`

GetVatRegimeOk returns a tuple with the VatRegime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatRegime

`func (o *SuperPDPVatRegimeUpdateRequest) SetVatRegime(v string)`

SetVatRegime sets VatRegime field to given value.


### GetHasVatOnDebits

`func (o *SuperPDPVatRegimeUpdateRequest) GetHasVatOnDebits() bool`

GetHasVatOnDebits returns the HasVatOnDebits field if non-nil, zero value otherwise.

### GetHasVatOnDebitsOk

`func (o *SuperPDPVatRegimeUpdateRequest) GetHasVatOnDebitsOk() (*bool, bool)`

GetHasVatOnDebitsOk returns a tuple with the HasVatOnDebits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVatOnDebits

`func (o *SuperPDPVatRegimeUpdateRequest) SetHasVatOnDebits(v bool)`

SetHasVatOnDebits sets HasVatOnDebits field to given value.

### HasHasVatOnDebits

`func (o *SuperPDPVatRegimeUpdateRequest) HasHasVatOnDebits() bool`

HasHasVatOnDebits returns a boolean if a field has been set.

### SetHasVatOnDebitsNil

`func (o *SuperPDPVatRegimeUpdateRequest) SetHasVatOnDebitsNil(b bool)`

 SetHasVatOnDebitsNil sets the value for HasVatOnDebits to be an explicit nil

### UnsetHasVatOnDebits
`func (o *SuperPDPVatRegimeUpdateRequest) UnsetHasVatOnDebits()`

UnsetHasVatOnDebits ensures that no value is present for HasVatOnDebits, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



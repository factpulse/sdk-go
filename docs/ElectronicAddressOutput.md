# ElectronicAddressOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | Electronic address identifier (e.g., SIREN, SIRET, routing code) | 
**SchemeId** | Pointer to [**SchemeID**](SchemeID.md) | Identification scheme (EAS code) | [optional] [default to FR_ELECTRONIC_ADDRESS]

## Methods

### NewElectronicAddressOutput

`func NewElectronicAddressOutput(identifier string, ) *ElectronicAddressOutput`

NewElectronicAddressOutput instantiates a new ElectronicAddressOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElectronicAddressOutputWithDefaults

`func NewElectronicAddressOutputWithDefaults() *ElectronicAddressOutput`

NewElectronicAddressOutputWithDefaults instantiates a new ElectronicAddressOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *ElectronicAddressOutput) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ElectronicAddressOutput) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ElectronicAddressOutput) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetSchemeId

`func (o *ElectronicAddressOutput) GetSchemeId() SchemeID`

GetSchemeId returns the SchemeId field if non-nil, zero value otherwise.

### GetSchemeIdOk

`func (o *ElectronicAddressOutput) GetSchemeIdOk() (*SchemeID, bool)`

GetSchemeIdOk returns a tuple with the SchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeId

`func (o *ElectronicAddressOutput) SetSchemeId(v SchemeID)`

SetSchemeId sets SchemeId field to given value.

### HasSchemeId

`func (o *ElectronicAddressOutput) HasSchemeId() bool`

HasSchemeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



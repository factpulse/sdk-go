# ElectronicAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** |  | 
**SchemeId** | Pointer to [**SchemeID**](SchemeID.md) |  | [optional] [default to FR_SIREN]

## Methods

### NewElectronicAddress

`func NewElectronicAddress(identifier string, ) *ElectronicAddress`

NewElectronicAddress instantiates a new ElectronicAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElectronicAddressWithDefaults

`func NewElectronicAddressWithDefaults() *ElectronicAddress`

NewElectronicAddressWithDefaults instantiates a new ElectronicAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *ElectronicAddress) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ElectronicAddress) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ElectronicAddress) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetSchemeId

`func (o *ElectronicAddress) GetSchemeId() SchemeID`

GetSchemeId returns the SchemeId field if non-nil, zero value otherwise.

### GetSchemeIdOk

`func (o *ElectronicAddress) GetSchemeIdOk() (*SchemeID, bool)`

GetSchemeIdOk returns a tuple with the SchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeId

`func (o *ElectronicAddress) SetSchemeId(v SchemeID)`

SetSchemeId sets SchemeId field to given value.

### HasSchemeId

`func (o *ElectronicAddress) HasSchemeId() bool`

HasSchemeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



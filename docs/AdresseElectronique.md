# AdresseElectronique

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifiant** | **string** |  | 
**SchemeId** | Pointer to [**SchemeID**](SchemeID.md) |  | [optional] [default to FR_SIREN]

## Methods

### NewAdresseElectronique

`func NewAdresseElectronique(identifiant string, ) *AdresseElectronique`

NewAdresseElectronique instantiates a new AdresseElectronique object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdresseElectroniqueWithDefaults

`func NewAdresseElectroniqueWithDefaults() *AdresseElectronique`

NewAdresseElectroniqueWithDefaults instantiates a new AdresseElectronique object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifiant

`func (o *AdresseElectronique) GetIdentifiant() string`

GetIdentifiant returns the Identifiant field if non-nil, zero value otherwise.

### GetIdentifiantOk

`func (o *AdresseElectronique) GetIdentifiantOk() (*string, bool)`

GetIdentifiantOk returns a tuple with the Identifiant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiant

`func (o *AdresseElectronique) SetIdentifiant(v string)`

SetIdentifiant sets Identifiant field to given value.


### GetSchemeId

`func (o *AdresseElectronique) GetSchemeId() SchemeID`

GetSchemeId returns the SchemeId field if non-nil, zero value otherwise.

### GetSchemeIdOk

`func (o *AdresseElectronique) GetSchemeIdOk() (*SchemeID, bool)`

GetSchemeIdOk returns a tuple with the SchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeId

`func (o *AdresseElectronique) SetSchemeId(v SchemeID)`

SetSchemeId sets SchemeId field to given value.

### HasSchemeId

`func (o *AdresseElectronique) HasSchemeId() bool`

HasSchemeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SoumettreFactureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodeRetour** | **int32** | Code retour (0 &#x3D; succès) | 
**Libelle** | **string** | Message de retour | 
**IdentifiantFactureCpp** | Pointer to **NullableInt32** |  | [optional] 
**NumeroFluxDepot** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSoumettreFactureResponse

`func NewSoumettreFactureResponse(codeRetour int32, libelle string, ) *SoumettreFactureResponse`

NewSoumettreFactureResponse instantiates a new SoumettreFactureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoumettreFactureResponseWithDefaults

`func NewSoumettreFactureResponseWithDefaults() *SoumettreFactureResponse`

NewSoumettreFactureResponseWithDefaults instantiates a new SoumettreFactureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodeRetour

`func (o *SoumettreFactureResponse) GetCodeRetour() int32`

GetCodeRetour returns the CodeRetour field if non-nil, zero value otherwise.

### GetCodeRetourOk

`func (o *SoumettreFactureResponse) GetCodeRetourOk() (*int32, bool)`

GetCodeRetourOk returns a tuple with the CodeRetour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeRetour

`func (o *SoumettreFactureResponse) SetCodeRetour(v int32)`

SetCodeRetour sets CodeRetour field to given value.


### GetLibelle

`func (o *SoumettreFactureResponse) GetLibelle() string`

GetLibelle returns the Libelle field if non-nil, zero value otherwise.

### GetLibelleOk

`func (o *SoumettreFactureResponse) GetLibelleOk() (*string, bool)`

GetLibelleOk returns a tuple with the Libelle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibelle

`func (o *SoumettreFactureResponse) SetLibelle(v string)`

SetLibelle sets Libelle field to given value.


### GetIdentifiantFactureCpp

`func (o *SoumettreFactureResponse) GetIdentifiantFactureCpp() int32`

GetIdentifiantFactureCpp returns the IdentifiantFactureCpp field if non-nil, zero value otherwise.

### GetIdentifiantFactureCppOk

`func (o *SoumettreFactureResponse) GetIdentifiantFactureCppOk() (*int32, bool)`

GetIdentifiantFactureCppOk returns a tuple with the IdentifiantFactureCpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiantFactureCpp

`func (o *SoumettreFactureResponse) SetIdentifiantFactureCpp(v int32)`

SetIdentifiantFactureCpp sets IdentifiantFactureCpp field to given value.

### HasIdentifiantFactureCpp

`func (o *SoumettreFactureResponse) HasIdentifiantFactureCpp() bool`

HasIdentifiantFactureCpp returns a boolean if a field has been set.

### SetIdentifiantFactureCppNil

`func (o *SoumettreFactureResponse) SetIdentifiantFactureCppNil(b bool)`

 SetIdentifiantFactureCppNil sets the value for IdentifiantFactureCpp to be an explicit nil

### UnsetIdentifiantFactureCpp
`func (o *SoumettreFactureResponse) UnsetIdentifiantFactureCpp()`

UnsetIdentifiantFactureCpp ensures that no value is present for IdentifiantFactureCpp, not even an explicit nil
### GetNumeroFluxDepot

`func (o *SoumettreFactureResponse) GetNumeroFluxDepot() string`

GetNumeroFluxDepot returns the NumeroFluxDepot field if non-nil, zero value otherwise.

### GetNumeroFluxDepotOk

`func (o *SoumettreFactureResponse) GetNumeroFluxDepotOk() (*string, bool)`

GetNumeroFluxDepotOk returns a tuple with the NumeroFluxDepot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroFluxDepot

`func (o *SoumettreFactureResponse) SetNumeroFluxDepot(v string)`

SetNumeroFluxDepot sets NumeroFluxDepot field to given value.

### HasNumeroFluxDepot

`func (o *SoumettreFactureResponse) HasNumeroFluxDepot() bool`

HasNumeroFluxDepot returns a boolean if a field has been set.

### SetNumeroFluxDepotNil

`func (o *SoumettreFactureResponse) SetNumeroFluxDepotNil(b bool)`

 SetNumeroFluxDepotNil sets the value for NumeroFluxDepot to be an explicit nil

### UnsetNumeroFluxDepot
`func (o *SoumettreFactureResponse) UnsetNumeroFluxDepot()`

UnsetNumeroFluxDepot ensures that no value is present for NumeroFluxDepot, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



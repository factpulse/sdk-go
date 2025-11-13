# References

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviseFacture** | Pointer to **string** |  | [optional] [default to "EUR"]
**ModePaiement** | [**ModePaiement**](ModePaiement.md) |  | 
**TypeFacture** | [**TypeFacture**](TypeFacture.md) |  | 
**TypeTva** | [**TypeTVA**](TypeTVA.md) |  | 
**NumeroMarche** | Pointer to **NullableString** |  | [optional] 
**MotifExonerationTva** | Pointer to **NullableString** |  | [optional] 
**NumeroBonCommande** | Pointer to **NullableString** |  | [optional] 
**NumeroFactureOrigine** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewReferences

`func NewReferences(modePaiement ModePaiement, typeFacture TypeFacture, typeTva TypeTVA, ) *References`

NewReferences instantiates a new References object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferencesWithDefaults

`func NewReferencesWithDefaults() *References`

NewReferencesWithDefaults instantiates a new References object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviseFacture

`func (o *References) GetDeviseFacture() string`

GetDeviseFacture returns the DeviseFacture field if non-nil, zero value otherwise.

### GetDeviseFactureOk

`func (o *References) GetDeviseFactureOk() (*string, bool)`

GetDeviseFactureOk returns a tuple with the DeviseFacture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviseFacture

`func (o *References) SetDeviseFacture(v string)`

SetDeviseFacture sets DeviseFacture field to given value.

### HasDeviseFacture

`func (o *References) HasDeviseFacture() bool`

HasDeviseFacture returns a boolean if a field has been set.

### GetModePaiement

`func (o *References) GetModePaiement() ModePaiement`

GetModePaiement returns the ModePaiement field if non-nil, zero value otherwise.

### GetModePaiementOk

`func (o *References) GetModePaiementOk() (*ModePaiement, bool)`

GetModePaiementOk returns a tuple with the ModePaiement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModePaiement

`func (o *References) SetModePaiement(v ModePaiement)`

SetModePaiement sets ModePaiement field to given value.


### GetTypeFacture

`func (o *References) GetTypeFacture() TypeFacture`

GetTypeFacture returns the TypeFacture field if non-nil, zero value otherwise.

### GetTypeFactureOk

`func (o *References) GetTypeFactureOk() (*TypeFacture, bool)`

GetTypeFactureOk returns a tuple with the TypeFacture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeFacture

`func (o *References) SetTypeFacture(v TypeFacture)`

SetTypeFacture sets TypeFacture field to given value.


### GetTypeTva

`func (o *References) GetTypeTva() TypeTVA`

GetTypeTva returns the TypeTva field if non-nil, zero value otherwise.

### GetTypeTvaOk

`func (o *References) GetTypeTvaOk() (*TypeTVA, bool)`

GetTypeTvaOk returns a tuple with the TypeTva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeTva

`func (o *References) SetTypeTva(v TypeTVA)`

SetTypeTva sets TypeTva field to given value.


### GetNumeroMarche

`func (o *References) GetNumeroMarche() string`

GetNumeroMarche returns the NumeroMarche field if non-nil, zero value otherwise.

### GetNumeroMarcheOk

`func (o *References) GetNumeroMarcheOk() (*string, bool)`

GetNumeroMarcheOk returns a tuple with the NumeroMarche field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroMarche

`func (o *References) SetNumeroMarche(v string)`

SetNumeroMarche sets NumeroMarche field to given value.

### HasNumeroMarche

`func (o *References) HasNumeroMarche() bool`

HasNumeroMarche returns a boolean if a field has been set.

### SetNumeroMarcheNil

`func (o *References) SetNumeroMarcheNil(b bool)`

 SetNumeroMarcheNil sets the value for NumeroMarche to be an explicit nil

### UnsetNumeroMarche
`func (o *References) UnsetNumeroMarche()`

UnsetNumeroMarche ensures that no value is present for NumeroMarche, not even an explicit nil
### GetMotifExonerationTva

`func (o *References) GetMotifExonerationTva() string`

GetMotifExonerationTva returns the MotifExonerationTva field if non-nil, zero value otherwise.

### GetMotifExonerationTvaOk

`func (o *References) GetMotifExonerationTvaOk() (*string, bool)`

GetMotifExonerationTvaOk returns a tuple with the MotifExonerationTva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotifExonerationTva

`func (o *References) SetMotifExonerationTva(v string)`

SetMotifExonerationTva sets MotifExonerationTva field to given value.

### HasMotifExonerationTva

`func (o *References) HasMotifExonerationTva() bool`

HasMotifExonerationTva returns a boolean if a field has been set.

### SetMotifExonerationTvaNil

`func (o *References) SetMotifExonerationTvaNil(b bool)`

 SetMotifExonerationTvaNil sets the value for MotifExonerationTva to be an explicit nil

### UnsetMotifExonerationTva
`func (o *References) UnsetMotifExonerationTva()`

UnsetMotifExonerationTva ensures that no value is present for MotifExonerationTva, not even an explicit nil
### GetNumeroBonCommande

`func (o *References) GetNumeroBonCommande() string`

GetNumeroBonCommande returns the NumeroBonCommande field if non-nil, zero value otherwise.

### GetNumeroBonCommandeOk

`func (o *References) GetNumeroBonCommandeOk() (*string, bool)`

GetNumeroBonCommandeOk returns a tuple with the NumeroBonCommande field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroBonCommande

`func (o *References) SetNumeroBonCommande(v string)`

SetNumeroBonCommande sets NumeroBonCommande field to given value.

### HasNumeroBonCommande

`func (o *References) HasNumeroBonCommande() bool`

HasNumeroBonCommande returns a boolean if a field has been set.

### SetNumeroBonCommandeNil

`func (o *References) SetNumeroBonCommandeNil(b bool)`

 SetNumeroBonCommandeNil sets the value for NumeroBonCommande to be an explicit nil

### UnsetNumeroBonCommande
`func (o *References) UnsetNumeroBonCommande()`

UnsetNumeroBonCommande ensures that no value is present for NumeroBonCommande, not even an explicit nil
### GetNumeroFactureOrigine

`func (o *References) GetNumeroFactureOrigine() string`

GetNumeroFactureOrigine returns the NumeroFactureOrigine field if non-nil, zero value otherwise.

### GetNumeroFactureOrigineOk

`func (o *References) GetNumeroFactureOrigineOk() (*string, bool)`

GetNumeroFactureOrigineOk returns a tuple with the NumeroFactureOrigine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroFactureOrigine

`func (o *References) SetNumeroFactureOrigine(v string)`

SetNumeroFactureOrigine sets NumeroFactureOrigine field to given value.

### HasNumeroFactureOrigine

`func (o *References) HasNumeroFactureOrigine() bool`

HasNumeroFactureOrigine returns a boolean if a field has been set.

### SetNumeroFactureOrigineNil

`func (o *References) SetNumeroFactureOrigineNil(b bool)`

 SetNumeroFactureOrigineNil sets the value for NumeroFactureOrigine to be an explicit nil

### UnsetNumeroFactureOrigine
`func (o *References) UnsetNumeroFactureOrigine()`

UnsetNumeroFactureOrigine ensures that no value is present for NumeroFactureOrigine, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



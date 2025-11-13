# LigneDeTVA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MontantBaseHt** | [**Montantbaseht**](Montantbaseht.md) |  | 
**MontantTva** | [**Montanttva**](Montanttva.md) |  | 
**Taux** | Pointer to **NullableString** |  | [optional] 
**TauxManuel** | Pointer to [**Tauxmanuel**](Tauxmanuel.md) |  | [optional] 
**Categorie** | Pointer to [**NullableCategorieTVA**](CategorieTVA.md) |  | [optional] 

## Methods

### NewLigneDeTVA

`func NewLigneDeTVA(montantBaseHt Montantbaseht, montantTva Montanttva, ) *LigneDeTVA`

NewLigneDeTVA instantiates a new LigneDeTVA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLigneDeTVAWithDefaults

`func NewLigneDeTVAWithDefaults() *LigneDeTVA`

NewLigneDeTVAWithDefaults instantiates a new LigneDeTVA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMontantBaseHt

`func (o *LigneDeTVA) GetMontantBaseHt() Montantbaseht`

GetMontantBaseHt returns the MontantBaseHt field if non-nil, zero value otherwise.

### GetMontantBaseHtOk

`func (o *LigneDeTVA) GetMontantBaseHtOk() (*Montantbaseht, bool)`

GetMontantBaseHtOk returns a tuple with the MontantBaseHt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantBaseHt

`func (o *LigneDeTVA) SetMontantBaseHt(v Montantbaseht)`

SetMontantBaseHt sets MontantBaseHt field to given value.


### GetMontantTva

`func (o *LigneDeTVA) GetMontantTva() Montanttva`

GetMontantTva returns the MontantTva field if non-nil, zero value otherwise.

### GetMontantTvaOk

`func (o *LigneDeTVA) GetMontantTvaOk() (*Montanttva, bool)`

GetMontantTvaOk returns a tuple with the MontantTva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantTva

`func (o *LigneDeTVA) SetMontantTva(v Montanttva)`

SetMontantTva sets MontantTva field to given value.


### GetTaux

`func (o *LigneDeTVA) GetTaux() string`

GetTaux returns the Taux field if non-nil, zero value otherwise.

### GetTauxOk

`func (o *LigneDeTVA) GetTauxOk() (*string, bool)`

GetTauxOk returns a tuple with the Taux field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaux

`func (o *LigneDeTVA) SetTaux(v string)`

SetTaux sets Taux field to given value.

### HasTaux

`func (o *LigneDeTVA) HasTaux() bool`

HasTaux returns a boolean if a field has been set.

### SetTauxNil

`func (o *LigneDeTVA) SetTauxNil(b bool)`

 SetTauxNil sets the value for Taux to be an explicit nil

### UnsetTaux
`func (o *LigneDeTVA) UnsetTaux()`

UnsetTaux ensures that no value is present for Taux, not even an explicit nil
### GetTauxManuel

`func (o *LigneDeTVA) GetTauxManuel() Tauxmanuel`

GetTauxManuel returns the TauxManuel field if non-nil, zero value otherwise.

### GetTauxManuelOk

`func (o *LigneDeTVA) GetTauxManuelOk() (*Tauxmanuel, bool)`

GetTauxManuelOk returns a tuple with the TauxManuel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTauxManuel

`func (o *LigneDeTVA) SetTauxManuel(v Tauxmanuel)`

SetTauxManuel sets TauxManuel field to given value.

### HasTauxManuel

`func (o *LigneDeTVA) HasTauxManuel() bool`

HasTauxManuel returns a boolean if a field has been set.

### GetCategorie

`func (o *LigneDeTVA) GetCategorie() CategorieTVA`

GetCategorie returns the Categorie field if non-nil, zero value otherwise.

### GetCategorieOk

`func (o *LigneDeTVA) GetCategorieOk() (*CategorieTVA, bool)`

GetCategorieOk returns a tuple with the Categorie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorie

`func (o *LigneDeTVA) SetCategorie(v CategorieTVA)`

SetCategorie sets Categorie field to given value.

### HasCategorie

`func (o *LigneDeTVA) HasCategorie() bool`

HasCategorie returns a boolean if a field has been set.

### SetCategorieNil

`func (o *LigneDeTVA) SetCategorieNil(b bool)`

 SetCategorieNil sets the value for Categorie to be an explicit nil

### UnsetCategorie
`func (o *LigneDeTVA) UnsetCategorie()`

UnsetCategorie ensures that no value is present for Categorie, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



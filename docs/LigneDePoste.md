# LigneDePoste

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Numero** | **int32** |  | 
**Reference** | Pointer to **NullableString** |  | [optional] 
**Denomination** | **string** |  | 
**Quantite** | [**Quantite**](Quantite.md) |  | 
**Unite** | [**Unite**](Unite.md) |  | 
**MontantUnitaireHt** | [**Montantunitaireht**](Montantunitaireht.md) |  | 
**MontantRemiseHt** | Pointer to [**NullableLigneDePosteMontantRemiseHt**](LigneDePosteMontantRemiseHt.md) |  | [optional] 
**MontantTotalLigneHt** | Pointer to [**NullableLigneDePosteMontantTotalLigneHt**](LigneDePosteMontantTotalLigneHt.md) |  | [optional] 
**TauxTva** | Pointer to **NullableString** |  | [optional] 
**TauxTvaManuel** | Pointer to [**NullableLigneDePosteTauxTvaManuel**](LigneDePosteTauxTvaManuel.md) |  | [optional] 
**CategorieTva** | Pointer to [**NullableCategorieTVA**](CategorieTVA.md) |  | [optional] 
**DateDebutPeriode** | Pointer to **NullableString** |  | [optional] 
**DateFinPeriode** | Pointer to **NullableString** |  | [optional] 
**CodeRaisonReduction** | Pointer to [**NullableCodeRaisonReduction**](CodeRaisonReduction.md) |  | [optional] 
**RaisonReduction** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLigneDePoste

`func NewLigneDePoste(numero int32, denomination string, quantite Quantite, unite Unite, montantUnitaireHt Montantunitaireht, ) *LigneDePoste`

NewLigneDePoste instantiates a new LigneDePoste object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLigneDePosteWithDefaults

`func NewLigneDePosteWithDefaults() *LigneDePoste`

NewLigneDePosteWithDefaults instantiates a new LigneDePoste object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumero

`func (o *LigneDePoste) GetNumero() int32`

GetNumero returns the Numero field if non-nil, zero value otherwise.

### GetNumeroOk

`func (o *LigneDePoste) GetNumeroOk() (*int32, bool)`

GetNumeroOk returns a tuple with the Numero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumero

`func (o *LigneDePoste) SetNumero(v int32)`

SetNumero sets Numero field to given value.


### GetReference

`func (o *LigneDePoste) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *LigneDePoste) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *LigneDePoste) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *LigneDePoste) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *LigneDePoste) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *LigneDePoste) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetDenomination

`func (o *LigneDePoste) GetDenomination() string`

GetDenomination returns the Denomination field if non-nil, zero value otherwise.

### GetDenominationOk

`func (o *LigneDePoste) GetDenominationOk() (*string, bool)`

GetDenominationOk returns a tuple with the Denomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenomination

`func (o *LigneDePoste) SetDenomination(v string)`

SetDenomination sets Denomination field to given value.


### GetQuantite

`func (o *LigneDePoste) GetQuantite() Quantite`

GetQuantite returns the Quantite field if non-nil, zero value otherwise.

### GetQuantiteOk

`func (o *LigneDePoste) GetQuantiteOk() (*Quantite, bool)`

GetQuantiteOk returns a tuple with the Quantite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantite

`func (o *LigneDePoste) SetQuantite(v Quantite)`

SetQuantite sets Quantite field to given value.


### GetUnite

`func (o *LigneDePoste) GetUnite() Unite`

GetUnite returns the Unite field if non-nil, zero value otherwise.

### GetUniteOk

`func (o *LigneDePoste) GetUniteOk() (*Unite, bool)`

GetUniteOk returns a tuple with the Unite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnite

`func (o *LigneDePoste) SetUnite(v Unite)`

SetUnite sets Unite field to given value.


### GetMontantUnitaireHt

`func (o *LigneDePoste) GetMontantUnitaireHt() Montantunitaireht`

GetMontantUnitaireHt returns the MontantUnitaireHt field if non-nil, zero value otherwise.

### GetMontantUnitaireHtOk

`func (o *LigneDePoste) GetMontantUnitaireHtOk() (*Montantunitaireht, bool)`

GetMontantUnitaireHtOk returns a tuple with the MontantUnitaireHt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantUnitaireHt

`func (o *LigneDePoste) SetMontantUnitaireHt(v Montantunitaireht)`

SetMontantUnitaireHt sets MontantUnitaireHt field to given value.


### GetMontantRemiseHt

`func (o *LigneDePoste) GetMontantRemiseHt() LigneDePosteMontantRemiseHt`

GetMontantRemiseHt returns the MontantRemiseHt field if non-nil, zero value otherwise.

### GetMontantRemiseHtOk

`func (o *LigneDePoste) GetMontantRemiseHtOk() (*LigneDePosteMontantRemiseHt, bool)`

GetMontantRemiseHtOk returns a tuple with the MontantRemiseHt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantRemiseHt

`func (o *LigneDePoste) SetMontantRemiseHt(v LigneDePosteMontantRemiseHt)`

SetMontantRemiseHt sets MontantRemiseHt field to given value.

### HasMontantRemiseHt

`func (o *LigneDePoste) HasMontantRemiseHt() bool`

HasMontantRemiseHt returns a boolean if a field has been set.

### SetMontantRemiseHtNil

`func (o *LigneDePoste) SetMontantRemiseHtNil(b bool)`

 SetMontantRemiseHtNil sets the value for MontantRemiseHt to be an explicit nil

### UnsetMontantRemiseHt
`func (o *LigneDePoste) UnsetMontantRemiseHt()`

UnsetMontantRemiseHt ensures that no value is present for MontantRemiseHt, not even an explicit nil
### GetMontantTotalLigneHt

`func (o *LigneDePoste) GetMontantTotalLigneHt() LigneDePosteMontantTotalLigneHt`

GetMontantTotalLigneHt returns the MontantTotalLigneHt field if non-nil, zero value otherwise.

### GetMontantTotalLigneHtOk

`func (o *LigneDePoste) GetMontantTotalLigneHtOk() (*LigneDePosteMontantTotalLigneHt, bool)`

GetMontantTotalLigneHtOk returns a tuple with the MontantTotalLigneHt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantTotalLigneHt

`func (o *LigneDePoste) SetMontantTotalLigneHt(v LigneDePosteMontantTotalLigneHt)`

SetMontantTotalLigneHt sets MontantTotalLigneHt field to given value.

### HasMontantTotalLigneHt

`func (o *LigneDePoste) HasMontantTotalLigneHt() bool`

HasMontantTotalLigneHt returns a boolean if a field has been set.

### SetMontantTotalLigneHtNil

`func (o *LigneDePoste) SetMontantTotalLigneHtNil(b bool)`

 SetMontantTotalLigneHtNil sets the value for MontantTotalLigneHt to be an explicit nil

### UnsetMontantTotalLigneHt
`func (o *LigneDePoste) UnsetMontantTotalLigneHt()`

UnsetMontantTotalLigneHt ensures that no value is present for MontantTotalLigneHt, not even an explicit nil
### GetTauxTva

`func (o *LigneDePoste) GetTauxTva() string`

GetTauxTva returns the TauxTva field if non-nil, zero value otherwise.

### GetTauxTvaOk

`func (o *LigneDePoste) GetTauxTvaOk() (*string, bool)`

GetTauxTvaOk returns a tuple with the TauxTva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTauxTva

`func (o *LigneDePoste) SetTauxTva(v string)`

SetTauxTva sets TauxTva field to given value.

### HasTauxTva

`func (o *LigneDePoste) HasTauxTva() bool`

HasTauxTva returns a boolean if a field has been set.

### SetTauxTvaNil

`func (o *LigneDePoste) SetTauxTvaNil(b bool)`

 SetTauxTvaNil sets the value for TauxTva to be an explicit nil

### UnsetTauxTva
`func (o *LigneDePoste) UnsetTauxTva()`

UnsetTauxTva ensures that no value is present for TauxTva, not even an explicit nil
### GetTauxTvaManuel

`func (o *LigneDePoste) GetTauxTvaManuel() LigneDePosteTauxTvaManuel`

GetTauxTvaManuel returns the TauxTvaManuel field if non-nil, zero value otherwise.

### GetTauxTvaManuelOk

`func (o *LigneDePoste) GetTauxTvaManuelOk() (*LigneDePosteTauxTvaManuel, bool)`

GetTauxTvaManuelOk returns a tuple with the TauxTvaManuel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTauxTvaManuel

`func (o *LigneDePoste) SetTauxTvaManuel(v LigneDePosteTauxTvaManuel)`

SetTauxTvaManuel sets TauxTvaManuel field to given value.

### HasTauxTvaManuel

`func (o *LigneDePoste) HasTauxTvaManuel() bool`

HasTauxTvaManuel returns a boolean if a field has been set.

### SetTauxTvaManuelNil

`func (o *LigneDePoste) SetTauxTvaManuelNil(b bool)`

 SetTauxTvaManuelNil sets the value for TauxTvaManuel to be an explicit nil

### UnsetTauxTvaManuel
`func (o *LigneDePoste) UnsetTauxTvaManuel()`

UnsetTauxTvaManuel ensures that no value is present for TauxTvaManuel, not even an explicit nil
### GetCategorieTva

`func (o *LigneDePoste) GetCategorieTva() CategorieTVA`

GetCategorieTva returns the CategorieTva field if non-nil, zero value otherwise.

### GetCategorieTvaOk

`func (o *LigneDePoste) GetCategorieTvaOk() (*CategorieTVA, bool)`

GetCategorieTvaOk returns a tuple with the CategorieTva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorieTva

`func (o *LigneDePoste) SetCategorieTva(v CategorieTVA)`

SetCategorieTva sets CategorieTva field to given value.

### HasCategorieTva

`func (o *LigneDePoste) HasCategorieTva() bool`

HasCategorieTva returns a boolean if a field has been set.

### SetCategorieTvaNil

`func (o *LigneDePoste) SetCategorieTvaNil(b bool)`

 SetCategorieTvaNil sets the value for CategorieTva to be an explicit nil

### UnsetCategorieTva
`func (o *LigneDePoste) UnsetCategorieTva()`

UnsetCategorieTva ensures that no value is present for CategorieTva, not even an explicit nil
### GetDateDebutPeriode

`func (o *LigneDePoste) GetDateDebutPeriode() string`

GetDateDebutPeriode returns the DateDebutPeriode field if non-nil, zero value otherwise.

### GetDateDebutPeriodeOk

`func (o *LigneDePoste) GetDateDebutPeriodeOk() (*string, bool)`

GetDateDebutPeriodeOk returns a tuple with the DateDebutPeriode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDebutPeriode

`func (o *LigneDePoste) SetDateDebutPeriode(v string)`

SetDateDebutPeriode sets DateDebutPeriode field to given value.

### HasDateDebutPeriode

`func (o *LigneDePoste) HasDateDebutPeriode() bool`

HasDateDebutPeriode returns a boolean if a field has been set.

### SetDateDebutPeriodeNil

`func (o *LigneDePoste) SetDateDebutPeriodeNil(b bool)`

 SetDateDebutPeriodeNil sets the value for DateDebutPeriode to be an explicit nil

### UnsetDateDebutPeriode
`func (o *LigneDePoste) UnsetDateDebutPeriode()`

UnsetDateDebutPeriode ensures that no value is present for DateDebutPeriode, not even an explicit nil
### GetDateFinPeriode

`func (o *LigneDePoste) GetDateFinPeriode() string`

GetDateFinPeriode returns the DateFinPeriode field if non-nil, zero value otherwise.

### GetDateFinPeriodeOk

`func (o *LigneDePoste) GetDateFinPeriodeOk() (*string, bool)`

GetDateFinPeriodeOk returns a tuple with the DateFinPeriode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFinPeriode

`func (o *LigneDePoste) SetDateFinPeriode(v string)`

SetDateFinPeriode sets DateFinPeriode field to given value.

### HasDateFinPeriode

`func (o *LigneDePoste) HasDateFinPeriode() bool`

HasDateFinPeriode returns a boolean if a field has been set.

### SetDateFinPeriodeNil

`func (o *LigneDePoste) SetDateFinPeriodeNil(b bool)`

 SetDateFinPeriodeNil sets the value for DateFinPeriode to be an explicit nil

### UnsetDateFinPeriode
`func (o *LigneDePoste) UnsetDateFinPeriode()`

UnsetDateFinPeriode ensures that no value is present for DateFinPeriode, not even an explicit nil
### GetCodeRaisonReduction

`func (o *LigneDePoste) GetCodeRaisonReduction() CodeRaisonReduction`

GetCodeRaisonReduction returns the CodeRaisonReduction field if non-nil, zero value otherwise.

### GetCodeRaisonReductionOk

`func (o *LigneDePoste) GetCodeRaisonReductionOk() (*CodeRaisonReduction, bool)`

GetCodeRaisonReductionOk returns a tuple with the CodeRaisonReduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeRaisonReduction

`func (o *LigneDePoste) SetCodeRaisonReduction(v CodeRaisonReduction)`

SetCodeRaisonReduction sets CodeRaisonReduction field to given value.

### HasCodeRaisonReduction

`func (o *LigneDePoste) HasCodeRaisonReduction() bool`

HasCodeRaisonReduction returns a boolean if a field has been set.

### SetCodeRaisonReductionNil

`func (o *LigneDePoste) SetCodeRaisonReductionNil(b bool)`

 SetCodeRaisonReductionNil sets the value for CodeRaisonReduction to be an explicit nil

### UnsetCodeRaisonReduction
`func (o *LigneDePoste) UnsetCodeRaisonReduction()`

UnsetCodeRaisonReduction ensures that no value is present for CodeRaisonReduction, not even an explicit nil
### GetRaisonReduction

`func (o *LigneDePoste) GetRaisonReduction() string`

GetRaisonReduction returns the RaisonReduction field if non-nil, zero value otherwise.

### GetRaisonReductionOk

`func (o *LigneDePoste) GetRaisonReductionOk() (*string, bool)`

GetRaisonReductionOk returns a tuple with the RaisonReduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaisonReduction

`func (o *LigneDePoste) SetRaisonReduction(v string)`

SetRaisonReduction sets RaisonReduction field to given value.

### HasRaisonReduction

`func (o *LigneDePoste) HasRaisonReduction() bool`

HasRaisonReduction returns a boolean if a field has been set.

### SetRaisonReductionNil

`func (o *LigneDePoste) SetRaisonReductionNil(b bool)`

 SetRaisonReductionNil sets the value for RaisonReduction to be an explicit nil

### UnsetRaisonReduction
`func (o *LigneDePoste) UnsetRaisonReduction()`

UnsetRaisonReduction ensures that no value is present for RaisonReduction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DonneesFactureSimplifiees

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Numero** | **string** | Numéro de facture unique | 
**Emetteur** | **map[string]interface{}** | Informations sur l&#39;émetteur (siret, iban, ...) | 
**Destinataire** | **map[string]interface{}** | Informations sur le destinataire (siret, ...) | 
**Lignes** | **[]map[string]interface{}** | Lignes de la facture | 
**Date** | Pointer to **NullableString** |  | [optional] 
**EcheanceJours** | Pointer to **int32** | Échéance en jours (défaut: 30) | [optional] [default to 30]
**Commentaire** | Pointer to **NullableString** |  | [optional] 
**NumeroBonCommande** | Pointer to **NullableString** |  | [optional] 
**NumeroMarche** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDonneesFactureSimplifiees

`func NewDonneesFactureSimplifiees(numero string, emetteur map[string]interface{}, destinataire map[string]interface{}, lignes []map[string]interface{}, ) *DonneesFactureSimplifiees`

NewDonneesFactureSimplifiees instantiates a new DonneesFactureSimplifiees object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDonneesFactureSimplifieesWithDefaults

`func NewDonneesFactureSimplifieesWithDefaults() *DonneesFactureSimplifiees`

NewDonneesFactureSimplifieesWithDefaults instantiates a new DonneesFactureSimplifiees object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumero

`func (o *DonneesFactureSimplifiees) GetNumero() string`

GetNumero returns the Numero field if non-nil, zero value otherwise.

### GetNumeroOk

`func (o *DonneesFactureSimplifiees) GetNumeroOk() (*string, bool)`

GetNumeroOk returns a tuple with the Numero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumero

`func (o *DonneesFactureSimplifiees) SetNumero(v string)`

SetNumero sets Numero field to given value.


### GetEmetteur

`func (o *DonneesFactureSimplifiees) GetEmetteur() map[string]interface{}`

GetEmetteur returns the Emetteur field if non-nil, zero value otherwise.

### GetEmetteurOk

`func (o *DonneesFactureSimplifiees) GetEmetteurOk() (*map[string]interface{}, bool)`

GetEmetteurOk returns a tuple with the Emetteur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmetteur

`func (o *DonneesFactureSimplifiees) SetEmetteur(v map[string]interface{})`

SetEmetteur sets Emetteur field to given value.


### GetDestinataire

`func (o *DonneesFactureSimplifiees) GetDestinataire() map[string]interface{}`

GetDestinataire returns the Destinataire field if non-nil, zero value otherwise.

### GetDestinataireOk

`func (o *DonneesFactureSimplifiees) GetDestinataireOk() (*map[string]interface{}, bool)`

GetDestinataireOk returns a tuple with the Destinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinataire

`func (o *DonneesFactureSimplifiees) SetDestinataire(v map[string]interface{})`

SetDestinataire sets Destinataire field to given value.


### GetLignes

`func (o *DonneesFactureSimplifiees) GetLignes() []map[string]interface{}`

GetLignes returns the Lignes field if non-nil, zero value otherwise.

### GetLignesOk

`func (o *DonneesFactureSimplifiees) GetLignesOk() (*[]map[string]interface{}, bool)`

GetLignesOk returns a tuple with the Lignes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLignes

`func (o *DonneesFactureSimplifiees) SetLignes(v []map[string]interface{})`

SetLignes sets Lignes field to given value.


### GetDate

`func (o *DonneesFactureSimplifiees) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DonneesFactureSimplifiees) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DonneesFactureSimplifiees) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *DonneesFactureSimplifiees) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *DonneesFactureSimplifiees) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *DonneesFactureSimplifiees) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetEcheanceJours

`func (o *DonneesFactureSimplifiees) GetEcheanceJours() int32`

GetEcheanceJours returns the EcheanceJours field if non-nil, zero value otherwise.

### GetEcheanceJoursOk

`func (o *DonneesFactureSimplifiees) GetEcheanceJoursOk() (*int32, bool)`

GetEcheanceJoursOk returns a tuple with the EcheanceJours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcheanceJours

`func (o *DonneesFactureSimplifiees) SetEcheanceJours(v int32)`

SetEcheanceJours sets EcheanceJours field to given value.

### HasEcheanceJours

`func (o *DonneesFactureSimplifiees) HasEcheanceJours() bool`

HasEcheanceJours returns a boolean if a field has been set.

### GetCommentaire

`func (o *DonneesFactureSimplifiees) GetCommentaire() string`

GetCommentaire returns the Commentaire field if non-nil, zero value otherwise.

### GetCommentaireOk

`func (o *DonneesFactureSimplifiees) GetCommentaireOk() (*string, bool)`

GetCommentaireOk returns a tuple with the Commentaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentaire

`func (o *DonneesFactureSimplifiees) SetCommentaire(v string)`

SetCommentaire sets Commentaire field to given value.

### HasCommentaire

`func (o *DonneesFactureSimplifiees) HasCommentaire() bool`

HasCommentaire returns a boolean if a field has been set.

### SetCommentaireNil

`func (o *DonneesFactureSimplifiees) SetCommentaireNil(b bool)`

 SetCommentaireNil sets the value for Commentaire to be an explicit nil

### UnsetCommentaire
`func (o *DonneesFactureSimplifiees) UnsetCommentaire()`

UnsetCommentaire ensures that no value is present for Commentaire, not even an explicit nil
### GetNumeroBonCommande

`func (o *DonneesFactureSimplifiees) GetNumeroBonCommande() string`

GetNumeroBonCommande returns the NumeroBonCommande field if non-nil, zero value otherwise.

### GetNumeroBonCommandeOk

`func (o *DonneesFactureSimplifiees) GetNumeroBonCommandeOk() (*string, bool)`

GetNumeroBonCommandeOk returns a tuple with the NumeroBonCommande field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroBonCommande

`func (o *DonneesFactureSimplifiees) SetNumeroBonCommande(v string)`

SetNumeroBonCommande sets NumeroBonCommande field to given value.

### HasNumeroBonCommande

`func (o *DonneesFactureSimplifiees) HasNumeroBonCommande() bool`

HasNumeroBonCommande returns a boolean if a field has been set.

### SetNumeroBonCommandeNil

`func (o *DonneesFactureSimplifiees) SetNumeroBonCommandeNil(b bool)`

 SetNumeroBonCommandeNil sets the value for NumeroBonCommande to be an explicit nil

### UnsetNumeroBonCommande
`func (o *DonneesFactureSimplifiees) UnsetNumeroBonCommande()`

UnsetNumeroBonCommande ensures that no value is present for NumeroBonCommande, not even an explicit nil
### GetNumeroMarche

`func (o *DonneesFactureSimplifiees) GetNumeroMarche() string`

GetNumeroMarche returns the NumeroMarche field if non-nil, zero value otherwise.

### GetNumeroMarcheOk

`func (o *DonneesFactureSimplifiees) GetNumeroMarcheOk() (*string, bool)`

GetNumeroMarcheOk returns a tuple with the NumeroMarche field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroMarche

`func (o *DonneesFactureSimplifiees) SetNumeroMarche(v string)`

SetNumeroMarche sets NumeroMarche field to given value.

### HasNumeroMarche

`func (o *DonneesFactureSimplifiees) HasNumeroMarche() bool`

HasNumeroMarche returns a boolean if a field has been set.

### SetNumeroMarcheNil

`func (o *DonneesFactureSimplifiees) SetNumeroMarcheNil(b bool)`

 SetNumeroMarcheNil sets the value for NumeroMarche to be an explicit nil

### UnsetNumeroMarche
`func (o *DonneesFactureSimplifiees) UnsetNumeroMarche()`

UnsetNumeroMarche ensures that no value is present for NumeroMarche, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



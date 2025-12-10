# FactureFacturX

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumeroFacture** | **string** |  | 
**DateEcheancePaiement** | **string** |  | 
**DateFacture** | Pointer to **string** |  | [optional] 
**ModeDepot** | [**ModeDepot**](ModeDepot.md) |  | 
**Destinataire** | [**Destinataire**](Destinataire.md) |  | 
**Fournisseur** | [**Fournisseur**](Fournisseur.md) |  | 
**CadreDeFacturation** | [**CadreDeFacturation**](CadreDeFacturation.md) |  | 
**References** | [**References**](References.md) |  | 
**MontantTotal** | [**MontantTotal**](MontantTotal.md) |  | 
**LignesDePoste** | Pointer to [**[]LigneDePoste**](LigneDePoste.md) |  | [optional] 
**LignesDeTva** | Pointer to [**[]LigneDeTVA**](LigneDeTVA.md) |  | [optional] 
**Notes** | Pointer to [**[]Note**](Note.md) |  | [optional] 
**Commentaire** | Pointer to **NullableString** |  | [optional] 
**IdUtilisateurCourant** | Pointer to **NullableInt32** |  | [optional] 
**PiecesJointesComplementaires** | Pointer to [**[]PieceJointeComplementaire**](PieceJointeComplementaire.md) |  | [optional] 
**Beneficiaire** | Pointer to [**NullableBeneficiaire**](Beneficiaire.md) |  | [optional] 

## Methods

### NewFactureFacturX

`func NewFactureFacturX(numeroFacture string, dateEcheancePaiement string, modeDepot ModeDepot, destinataire Destinataire, fournisseur Fournisseur, cadreDeFacturation CadreDeFacturation, references References, montantTotal MontantTotal, ) *FactureFacturX`

NewFactureFacturX instantiates a new FactureFacturX object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFactureFacturXWithDefaults

`func NewFactureFacturXWithDefaults() *FactureFacturX`

NewFactureFacturXWithDefaults instantiates a new FactureFacturX object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumeroFacture

`func (o *FactureFacturX) GetNumeroFacture() string`

GetNumeroFacture returns the NumeroFacture field if non-nil, zero value otherwise.

### GetNumeroFactureOk

`func (o *FactureFacturX) GetNumeroFactureOk() (*string, bool)`

GetNumeroFactureOk returns a tuple with the NumeroFacture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroFacture

`func (o *FactureFacturX) SetNumeroFacture(v string)`

SetNumeroFacture sets NumeroFacture field to given value.


### GetDateEcheancePaiement

`func (o *FactureFacturX) GetDateEcheancePaiement() string`

GetDateEcheancePaiement returns the DateEcheancePaiement field if non-nil, zero value otherwise.

### GetDateEcheancePaiementOk

`func (o *FactureFacturX) GetDateEcheancePaiementOk() (*string, bool)`

GetDateEcheancePaiementOk returns a tuple with the DateEcheancePaiement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEcheancePaiement

`func (o *FactureFacturX) SetDateEcheancePaiement(v string)`

SetDateEcheancePaiement sets DateEcheancePaiement field to given value.


### GetDateFacture

`func (o *FactureFacturX) GetDateFacture() string`

GetDateFacture returns the DateFacture field if non-nil, zero value otherwise.

### GetDateFactureOk

`func (o *FactureFacturX) GetDateFactureOk() (*string, bool)`

GetDateFactureOk returns a tuple with the DateFacture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFacture

`func (o *FactureFacturX) SetDateFacture(v string)`

SetDateFacture sets DateFacture field to given value.

### HasDateFacture

`func (o *FactureFacturX) HasDateFacture() bool`

HasDateFacture returns a boolean if a field has been set.

### GetModeDepot

`func (o *FactureFacturX) GetModeDepot() ModeDepot`

GetModeDepot returns the ModeDepot field if non-nil, zero value otherwise.

### GetModeDepotOk

`func (o *FactureFacturX) GetModeDepotOk() (*ModeDepot, bool)`

GetModeDepotOk returns a tuple with the ModeDepot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeDepot

`func (o *FactureFacturX) SetModeDepot(v ModeDepot)`

SetModeDepot sets ModeDepot field to given value.


### GetDestinataire

`func (o *FactureFacturX) GetDestinataire() Destinataire`

GetDestinataire returns the Destinataire field if non-nil, zero value otherwise.

### GetDestinataireOk

`func (o *FactureFacturX) GetDestinataireOk() (*Destinataire, bool)`

GetDestinataireOk returns a tuple with the Destinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinataire

`func (o *FactureFacturX) SetDestinataire(v Destinataire)`

SetDestinataire sets Destinataire field to given value.


### GetFournisseur

`func (o *FactureFacturX) GetFournisseur() Fournisseur`

GetFournisseur returns the Fournisseur field if non-nil, zero value otherwise.

### GetFournisseurOk

`func (o *FactureFacturX) GetFournisseurOk() (*Fournisseur, bool)`

GetFournisseurOk returns a tuple with the Fournisseur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFournisseur

`func (o *FactureFacturX) SetFournisseur(v Fournisseur)`

SetFournisseur sets Fournisseur field to given value.


### GetCadreDeFacturation

`func (o *FactureFacturX) GetCadreDeFacturation() CadreDeFacturation`

GetCadreDeFacturation returns the CadreDeFacturation field if non-nil, zero value otherwise.

### GetCadreDeFacturationOk

`func (o *FactureFacturX) GetCadreDeFacturationOk() (*CadreDeFacturation, bool)`

GetCadreDeFacturationOk returns a tuple with the CadreDeFacturation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCadreDeFacturation

`func (o *FactureFacturX) SetCadreDeFacturation(v CadreDeFacturation)`

SetCadreDeFacturation sets CadreDeFacturation field to given value.


### GetReferences

`func (o *FactureFacturX) GetReferences() References`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *FactureFacturX) GetReferencesOk() (*References, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *FactureFacturX) SetReferences(v References)`

SetReferences sets References field to given value.


### GetMontantTotal

`func (o *FactureFacturX) GetMontantTotal() MontantTotal`

GetMontantTotal returns the MontantTotal field if non-nil, zero value otherwise.

### GetMontantTotalOk

`func (o *FactureFacturX) GetMontantTotalOk() (*MontantTotal, bool)`

GetMontantTotalOk returns a tuple with the MontantTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantTotal

`func (o *FactureFacturX) SetMontantTotal(v MontantTotal)`

SetMontantTotal sets MontantTotal field to given value.


### GetLignesDePoste

`func (o *FactureFacturX) GetLignesDePoste() []LigneDePoste`

GetLignesDePoste returns the LignesDePoste field if non-nil, zero value otherwise.

### GetLignesDePosteOk

`func (o *FactureFacturX) GetLignesDePosteOk() (*[]LigneDePoste, bool)`

GetLignesDePosteOk returns a tuple with the LignesDePoste field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLignesDePoste

`func (o *FactureFacturX) SetLignesDePoste(v []LigneDePoste)`

SetLignesDePoste sets LignesDePoste field to given value.

### HasLignesDePoste

`func (o *FactureFacturX) HasLignesDePoste() bool`

HasLignesDePoste returns a boolean if a field has been set.

### GetLignesDeTva

`func (o *FactureFacturX) GetLignesDeTva() []LigneDeTVA`

GetLignesDeTva returns the LignesDeTva field if non-nil, zero value otherwise.

### GetLignesDeTvaOk

`func (o *FactureFacturX) GetLignesDeTvaOk() (*[]LigneDeTVA, bool)`

GetLignesDeTvaOk returns a tuple with the LignesDeTva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLignesDeTva

`func (o *FactureFacturX) SetLignesDeTva(v []LigneDeTVA)`

SetLignesDeTva sets LignesDeTva field to given value.

### HasLignesDeTva

`func (o *FactureFacturX) HasLignesDeTva() bool`

HasLignesDeTva returns a boolean if a field has been set.

### GetNotes

`func (o *FactureFacturX) GetNotes() []Note`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *FactureFacturX) GetNotesOk() (*[]Note, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *FactureFacturX) SetNotes(v []Note)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *FactureFacturX) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetCommentaire

`func (o *FactureFacturX) GetCommentaire() string`

GetCommentaire returns the Commentaire field if non-nil, zero value otherwise.

### GetCommentaireOk

`func (o *FactureFacturX) GetCommentaireOk() (*string, bool)`

GetCommentaireOk returns a tuple with the Commentaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentaire

`func (o *FactureFacturX) SetCommentaire(v string)`

SetCommentaire sets Commentaire field to given value.

### HasCommentaire

`func (o *FactureFacturX) HasCommentaire() bool`

HasCommentaire returns a boolean if a field has been set.

### SetCommentaireNil

`func (o *FactureFacturX) SetCommentaireNil(b bool)`

 SetCommentaireNil sets the value for Commentaire to be an explicit nil

### UnsetCommentaire
`func (o *FactureFacturX) UnsetCommentaire()`

UnsetCommentaire ensures that no value is present for Commentaire, not even an explicit nil
### GetIdUtilisateurCourant

`func (o *FactureFacturX) GetIdUtilisateurCourant() int32`

GetIdUtilisateurCourant returns the IdUtilisateurCourant field if non-nil, zero value otherwise.

### GetIdUtilisateurCourantOk

`func (o *FactureFacturX) GetIdUtilisateurCourantOk() (*int32, bool)`

GetIdUtilisateurCourantOk returns a tuple with the IdUtilisateurCourant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdUtilisateurCourant

`func (o *FactureFacturX) SetIdUtilisateurCourant(v int32)`

SetIdUtilisateurCourant sets IdUtilisateurCourant field to given value.

### HasIdUtilisateurCourant

`func (o *FactureFacturX) HasIdUtilisateurCourant() bool`

HasIdUtilisateurCourant returns a boolean if a field has been set.

### SetIdUtilisateurCourantNil

`func (o *FactureFacturX) SetIdUtilisateurCourantNil(b bool)`

 SetIdUtilisateurCourantNil sets the value for IdUtilisateurCourant to be an explicit nil

### UnsetIdUtilisateurCourant
`func (o *FactureFacturX) UnsetIdUtilisateurCourant()`

UnsetIdUtilisateurCourant ensures that no value is present for IdUtilisateurCourant, not even an explicit nil
### GetPiecesJointesComplementaires

`func (o *FactureFacturX) GetPiecesJointesComplementaires() []PieceJointeComplementaire`

GetPiecesJointesComplementaires returns the PiecesJointesComplementaires field if non-nil, zero value otherwise.

### GetPiecesJointesComplementairesOk

`func (o *FactureFacturX) GetPiecesJointesComplementairesOk() (*[]PieceJointeComplementaire, bool)`

GetPiecesJointesComplementairesOk returns a tuple with the PiecesJointesComplementaires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiecesJointesComplementaires

`func (o *FactureFacturX) SetPiecesJointesComplementaires(v []PieceJointeComplementaire)`

SetPiecesJointesComplementaires sets PiecesJointesComplementaires field to given value.

### HasPiecesJointesComplementaires

`func (o *FactureFacturX) HasPiecesJointesComplementaires() bool`

HasPiecesJointesComplementaires returns a boolean if a field has been set.

### SetPiecesJointesComplementairesNil

`func (o *FactureFacturX) SetPiecesJointesComplementairesNil(b bool)`

 SetPiecesJointesComplementairesNil sets the value for PiecesJointesComplementaires to be an explicit nil

### UnsetPiecesJointesComplementaires
`func (o *FactureFacturX) UnsetPiecesJointesComplementaires()`

UnsetPiecesJointesComplementaires ensures that no value is present for PiecesJointesComplementaires, not even an explicit nil
### GetBeneficiaire

`func (o *FactureFacturX) GetBeneficiaire() Beneficiaire`

GetBeneficiaire returns the Beneficiaire field if non-nil, zero value otherwise.

### GetBeneficiaireOk

`func (o *FactureFacturX) GetBeneficiaireOk() (*Beneficiaire, bool)`

GetBeneficiaireOk returns a tuple with the Beneficiaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaire

`func (o *FactureFacturX) SetBeneficiaire(v Beneficiaire)`

SetBeneficiaire sets Beneficiaire field to given value.

### HasBeneficiaire

`func (o *FactureFacturX) HasBeneficiaire() bool`

HasBeneficiaire returns a boolean if a field has been set.

### SetBeneficiaireNil

`func (o *FactureFacturX) SetBeneficiaireNil(b bool)`

 SetBeneficiaireNil sets the value for Beneficiaire to be an explicit nil

### UnsetBeneficiaire
`func (o *FactureFacturX) UnsetBeneficiaire()`

UnsetBeneficiaire ensures that no value is present for Beneficiaire, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



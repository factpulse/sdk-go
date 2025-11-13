# SoumettreFactureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**NullableChorusProCredentials**](ChorusProCredentials.md) |  | [optional] 
**NumeroFacture** | **string** | Numéro de la facture | 
**DateFacture** | **string** | Date de facture (format ISO: YYYY-MM-DD) | 
**DateEcheancePaiement** | Pointer to **NullableString** |  | [optional] 
**IdStructureCpp** | **int32** | ID Chorus Pro de la structure destinataire | 
**CodeService** | Pointer to **NullableString** |  | [optional] 
**NumeroEngagement** | Pointer to **NullableString** |  | [optional] 
**MontantHtTotal** | [**MontantHtTotal1**](MontantHtTotal1.md) |  | 
**MontantTva** | [**MontantTva1**](MontantTva1.md) |  | 
**MontantTtcTotal** | [**MontantTtcTotal1**](MontantTtcTotal1.md) |  | 
**PieceJointePrincipaleId** | Pointer to **NullableInt32** |  | [optional] 
**PieceJointePrincipaleDesignation** | Pointer to **NullableString** |  | [optional] 
**Commentaire** | Pointer to **NullableString** |  | [optional] 
**NumeroBonCommande** | Pointer to **NullableString** |  | [optional] 
**NumeroMarche** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSoumettreFactureRequest

`func NewSoumettreFactureRequest(numeroFacture string, dateFacture string, idStructureCpp int32, montantHtTotal MontantHtTotal1, montantTva MontantTva1, montantTtcTotal MontantTtcTotal1, ) *SoumettreFactureRequest`

NewSoumettreFactureRequest instantiates a new SoumettreFactureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoumettreFactureRequestWithDefaults

`func NewSoumettreFactureRequestWithDefaults() *SoumettreFactureRequest`

NewSoumettreFactureRequestWithDefaults instantiates a new SoumettreFactureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *SoumettreFactureRequest) GetCredentials() ChorusProCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *SoumettreFactureRequest) GetCredentialsOk() (*ChorusProCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *SoumettreFactureRequest) SetCredentials(v ChorusProCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *SoumettreFactureRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *SoumettreFactureRequest) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *SoumettreFactureRequest) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetNumeroFacture

`func (o *SoumettreFactureRequest) GetNumeroFacture() string`

GetNumeroFacture returns the NumeroFacture field if non-nil, zero value otherwise.

### GetNumeroFactureOk

`func (o *SoumettreFactureRequest) GetNumeroFactureOk() (*string, bool)`

GetNumeroFactureOk returns a tuple with the NumeroFacture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroFacture

`func (o *SoumettreFactureRequest) SetNumeroFacture(v string)`

SetNumeroFacture sets NumeroFacture field to given value.


### GetDateFacture

`func (o *SoumettreFactureRequest) GetDateFacture() string`

GetDateFacture returns the DateFacture field if non-nil, zero value otherwise.

### GetDateFactureOk

`func (o *SoumettreFactureRequest) GetDateFactureOk() (*string, bool)`

GetDateFactureOk returns a tuple with the DateFacture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFacture

`func (o *SoumettreFactureRequest) SetDateFacture(v string)`

SetDateFacture sets DateFacture field to given value.


### GetDateEcheancePaiement

`func (o *SoumettreFactureRequest) GetDateEcheancePaiement() string`

GetDateEcheancePaiement returns the DateEcheancePaiement field if non-nil, zero value otherwise.

### GetDateEcheancePaiementOk

`func (o *SoumettreFactureRequest) GetDateEcheancePaiementOk() (*string, bool)`

GetDateEcheancePaiementOk returns a tuple with the DateEcheancePaiement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEcheancePaiement

`func (o *SoumettreFactureRequest) SetDateEcheancePaiement(v string)`

SetDateEcheancePaiement sets DateEcheancePaiement field to given value.

### HasDateEcheancePaiement

`func (o *SoumettreFactureRequest) HasDateEcheancePaiement() bool`

HasDateEcheancePaiement returns a boolean if a field has been set.

### SetDateEcheancePaiementNil

`func (o *SoumettreFactureRequest) SetDateEcheancePaiementNil(b bool)`

 SetDateEcheancePaiementNil sets the value for DateEcheancePaiement to be an explicit nil

### UnsetDateEcheancePaiement
`func (o *SoumettreFactureRequest) UnsetDateEcheancePaiement()`

UnsetDateEcheancePaiement ensures that no value is present for DateEcheancePaiement, not even an explicit nil
### GetIdStructureCpp

`func (o *SoumettreFactureRequest) GetIdStructureCpp() int32`

GetIdStructureCpp returns the IdStructureCpp field if non-nil, zero value otherwise.

### GetIdStructureCppOk

`func (o *SoumettreFactureRequest) GetIdStructureCppOk() (*int32, bool)`

GetIdStructureCppOk returns a tuple with the IdStructureCpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdStructureCpp

`func (o *SoumettreFactureRequest) SetIdStructureCpp(v int32)`

SetIdStructureCpp sets IdStructureCpp field to given value.


### GetCodeService

`func (o *SoumettreFactureRequest) GetCodeService() string`

GetCodeService returns the CodeService field if non-nil, zero value otherwise.

### GetCodeServiceOk

`func (o *SoumettreFactureRequest) GetCodeServiceOk() (*string, bool)`

GetCodeServiceOk returns a tuple with the CodeService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeService

`func (o *SoumettreFactureRequest) SetCodeService(v string)`

SetCodeService sets CodeService field to given value.

### HasCodeService

`func (o *SoumettreFactureRequest) HasCodeService() bool`

HasCodeService returns a boolean if a field has been set.

### SetCodeServiceNil

`func (o *SoumettreFactureRequest) SetCodeServiceNil(b bool)`

 SetCodeServiceNil sets the value for CodeService to be an explicit nil

### UnsetCodeService
`func (o *SoumettreFactureRequest) UnsetCodeService()`

UnsetCodeService ensures that no value is present for CodeService, not even an explicit nil
### GetNumeroEngagement

`func (o *SoumettreFactureRequest) GetNumeroEngagement() string`

GetNumeroEngagement returns the NumeroEngagement field if non-nil, zero value otherwise.

### GetNumeroEngagementOk

`func (o *SoumettreFactureRequest) GetNumeroEngagementOk() (*string, bool)`

GetNumeroEngagementOk returns a tuple with the NumeroEngagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroEngagement

`func (o *SoumettreFactureRequest) SetNumeroEngagement(v string)`

SetNumeroEngagement sets NumeroEngagement field to given value.

### HasNumeroEngagement

`func (o *SoumettreFactureRequest) HasNumeroEngagement() bool`

HasNumeroEngagement returns a boolean if a field has been set.

### SetNumeroEngagementNil

`func (o *SoumettreFactureRequest) SetNumeroEngagementNil(b bool)`

 SetNumeroEngagementNil sets the value for NumeroEngagement to be an explicit nil

### UnsetNumeroEngagement
`func (o *SoumettreFactureRequest) UnsetNumeroEngagement()`

UnsetNumeroEngagement ensures that no value is present for NumeroEngagement, not even an explicit nil
### GetMontantHtTotal

`func (o *SoumettreFactureRequest) GetMontantHtTotal() MontantHtTotal1`

GetMontantHtTotal returns the MontantHtTotal field if non-nil, zero value otherwise.

### GetMontantHtTotalOk

`func (o *SoumettreFactureRequest) GetMontantHtTotalOk() (*MontantHtTotal1, bool)`

GetMontantHtTotalOk returns a tuple with the MontantHtTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantHtTotal

`func (o *SoumettreFactureRequest) SetMontantHtTotal(v MontantHtTotal1)`

SetMontantHtTotal sets MontantHtTotal field to given value.


### GetMontantTva

`func (o *SoumettreFactureRequest) GetMontantTva() MontantTva1`

GetMontantTva returns the MontantTva field if non-nil, zero value otherwise.

### GetMontantTvaOk

`func (o *SoumettreFactureRequest) GetMontantTvaOk() (*MontantTva1, bool)`

GetMontantTvaOk returns a tuple with the MontantTva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantTva

`func (o *SoumettreFactureRequest) SetMontantTva(v MontantTva1)`

SetMontantTva sets MontantTva field to given value.


### GetMontantTtcTotal

`func (o *SoumettreFactureRequest) GetMontantTtcTotal() MontantTtcTotal1`

GetMontantTtcTotal returns the MontantTtcTotal field if non-nil, zero value otherwise.

### GetMontantTtcTotalOk

`func (o *SoumettreFactureRequest) GetMontantTtcTotalOk() (*MontantTtcTotal1, bool)`

GetMontantTtcTotalOk returns a tuple with the MontantTtcTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantTtcTotal

`func (o *SoumettreFactureRequest) SetMontantTtcTotal(v MontantTtcTotal1)`

SetMontantTtcTotal sets MontantTtcTotal field to given value.


### GetPieceJointePrincipaleId

`func (o *SoumettreFactureRequest) GetPieceJointePrincipaleId() int32`

GetPieceJointePrincipaleId returns the PieceJointePrincipaleId field if non-nil, zero value otherwise.

### GetPieceJointePrincipaleIdOk

`func (o *SoumettreFactureRequest) GetPieceJointePrincipaleIdOk() (*int32, bool)`

GetPieceJointePrincipaleIdOk returns a tuple with the PieceJointePrincipaleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPieceJointePrincipaleId

`func (o *SoumettreFactureRequest) SetPieceJointePrincipaleId(v int32)`

SetPieceJointePrincipaleId sets PieceJointePrincipaleId field to given value.

### HasPieceJointePrincipaleId

`func (o *SoumettreFactureRequest) HasPieceJointePrincipaleId() bool`

HasPieceJointePrincipaleId returns a boolean if a field has been set.

### SetPieceJointePrincipaleIdNil

`func (o *SoumettreFactureRequest) SetPieceJointePrincipaleIdNil(b bool)`

 SetPieceJointePrincipaleIdNil sets the value for PieceJointePrincipaleId to be an explicit nil

### UnsetPieceJointePrincipaleId
`func (o *SoumettreFactureRequest) UnsetPieceJointePrincipaleId()`

UnsetPieceJointePrincipaleId ensures that no value is present for PieceJointePrincipaleId, not even an explicit nil
### GetPieceJointePrincipaleDesignation

`func (o *SoumettreFactureRequest) GetPieceJointePrincipaleDesignation() string`

GetPieceJointePrincipaleDesignation returns the PieceJointePrincipaleDesignation field if non-nil, zero value otherwise.

### GetPieceJointePrincipaleDesignationOk

`func (o *SoumettreFactureRequest) GetPieceJointePrincipaleDesignationOk() (*string, bool)`

GetPieceJointePrincipaleDesignationOk returns a tuple with the PieceJointePrincipaleDesignation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPieceJointePrincipaleDesignation

`func (o *SoumettreFactureRequest) SetPieceJointePrincipaleDesignation(v string)`

SetPieceJointePrincipaleDesignation sets PieceJointePrincipaleDesignation field to given value.

### HasPieceJointePrincipaleDesignation

`func (o *SoumettreFactureRequest) HasPieceJointePrincipaleDesignation() bool`

HasPieceJointePrincipaleDesignation returns a boolean if a field has been set.

### SetPieceJointePrincipaleDesignationNil

`func (o *SoumettreFactureRequest) SetPieceJointePrincipaleDesignationNil(b bool)`

 SetPieceJointePrincipaleDesignationNil sets the value for PieceJointePrincipaleDesignation to be an explicit nil

### UnsetPieceJointePrincipaleDesignation
`func (o *SoumettreFactureRequest) UnsetPieceJointePrincipaleDesignation()`

UnsetPieceJointePrincipaleDesignation ensures that no value is present for PieceJointePrincipaleDesignation, not even an explicit nil
### GetCommentaire

`func (o *SoumettreFactureRequest) GetCommentaire() string`

GetCommentaire returns the Commentaire field if non-nil, zero value otherwise.

### GetCommentaireOk

`func (o *SoumettreFactureRequest) GetCommentaireOk() (*string, bool)`

GetCommentaireOk returns a tuple with the Commentaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentaire

`func (o *SoumettreFactureRequest) SetCommentaire(v string)`

SetCommentaire sets Commentaire field to given value.

### HasCommentaire

`func (o *SoumettreFactureRequest) HasCommentaire() bool`

HasCommentaire returns a boolean if a field has been set.

### SetCommentaireNil

`func (o *SoumettreFactureRequest) SetCommentaireNil(b bool)`

 SetCommentaireNil sets the value for Commentaire to be an explicit nil

### UnsetCommentaire
`func (o *SoumettreFactureRequest) UnsetCommentaire()`

UnsetCommentaire ensures that no value is present for Commentaire, not even an explicit nil
### GetNumeroBonCommande

`func (o *SoumettreFactureRequest) GetNumeroBonCommande() string`

GetNumeroBonCommande returns the NumeroBonCommande field if non-nil, zero value otherwise.

### GetNumeroBonCommandeOk

`func (o *SoumettreFactureRequest) GetNumeroBonCommandeOk() (*string, bool)`

GetNumeroBonCommandeOk returns a tuple with the NumeroBonCommande field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroBonCommande

`func (o *SoumettreFactureRequest) SetNumeroBonCommande(v string)`

SetNumeroBonCommande sets NumeroBonCommande field to given value.

### HasNumeroBonCommande

`func (o *SoumettreFactureRequest) HasNumeroBonCommande() bool`

HasNumeroBonCommande returns a boolean if a field has been set.

### SetNumeroBonCommandeNil

`func (o *SoumettreFactureRequest) SetNumeroBonCommandeNil(b bool)`

 SetNumeroBonCommandeNil sets the value for NumeroBonCommande to be an explicit nil

### UnsetNumeroBonCommande
`func (o *SoumettreFactureRequest) UnsetNumeroBonCommande()`

UnsetNumeroBonCommande ensures that no value is present for NumeroBonCommande, not even an explicit nil
### GetNumeroMarche

`func (o *SoumettreFactureRequest) GetNumeroMarche() string`

GetNumeroMarche returns the NumeroMarche field if non-nil, zero value otherwise.

### GetNumeroMarcheOk

`func (o *SoumettreFactureRequest) GetNumeroMarcheOk() (*string, bool)`

GetNumeroMarcheOk returns a tuple with the NumeroMarche field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroMarche

`func (o *SoumettreFactureRequest) SetNumeroMarche(v string)`

SetNumeroMarche sets NumeroMarche field to given value.

### HasNumeroMarche

`func (o *SoumettreFactureRequest) HasNumeroMarche() bool`

HasNumeroMarche returns a boolean if a field has been set.

### SetNumeroMarcheNil

`func (o *SoumettreFactureRequest) SetNumeroMarcheNil(b bool)`

 SetNumeroMarcheNil sets the value for NumeroMarche to be an explicit nil

### UnsetNumeroMarche
`func (o *SoumettreFactureRequest) UnsetNumeroMarche()`

UnsetNumeroMarche ensures that no value is present for NumeroMarche, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



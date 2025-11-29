# FactureEntrante

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowId** | Pointer to **NullableString** |  | [optional] 
**FormatSource** | [**FormatFacture**](FormatFacture.md) | Format source de la facture | 
**RefFournisseur** | **string** | Numéro de facture émis par le fournisseur (BT-1) | 
**TypeDocument** | Pointer to [**TypeDocument**](TypeDocument.md) | Type de document (BT-3) | [optional] [default to _380]
**Fournisseur** | [**FournisseurEntrant**](FournisseurEntrant.md) | Émetteur de la facture (SellerTradeParty) | 
**SiteFacturationNom** | **string** | Nom du destinataire / votre entreprise (BT-44) | 
**SiteFacturationSiret** | Pointer to **NullableString** |  | [optional] 
**DateDePiece** | **string** | Date de la facture (BT-2) - YYYY-MM-DD | 
**DateReglement** | Pointer to **NullableString** |  | [optional] 
**Devise** | Pointer to **string** | Code devise ISO (BT-5) | [optional] [default to "EUR"]
**MontantHt** | **string** | Montant HT total (BT-109) | 
**MontantTva** | **string** | Montant TVA total (BT-110) | 
**MontantTtc** | **string** | Montant TTC total (BT-112) | 
**NumeroBonCommande** | Pointer to **NullableString** |  | [optional] 
**ReferenceContrat** | Pointer to **NullableString** |  | [optional] 
**ObjetFacture** | Pointer to **NullableString** |  | [optional] 
**DocumentBase64** | Pointer to **NullableString** |  | [optional] 
**DocumentContentType** | Pointer to **NullableString** |  | [optional] 
**DocumentFilename** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFactureEntrante

`func NewFactureEntrante(formatSource FormatFacture, refFournisseur string, fournisseur FournisseurEntrant, siteFacturationNom string, dateDePiece string, montantHt string, montantTva string, montantTtc string, ) *FactureEntrante`

NewFactureEntrante instantiates a new FactureEntrante object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFactureEntranteWithDefaults

`func NewFactureEntranteWithDefaults() *FactureEntrante`

NewFactureEntranteWithDefaults instantiates a new FactureEntrante object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowId

`func (o *FactureEntrante) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *FactureEntrante) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *FactureEntrante) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *FactureEntrante) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### SetFlowIdNil

`func (o *FactureEntrante) SetFlowIdNil(b bool)`

 SetFlowIdNil sets the value for FlowId to be an explicit nil

### UnsetFlowId
`func (o *FactureEntrante) UnsetFlowId()`

UnsetFlowId ensures that no value is present for FlowId, not even an explicit nil
### GetFormatSource

`func (o *FactureEntrante) GetFormatSource() FormatFacture`

GetFormatSource returns the FormatSource field if non-nil, zero value otherwise.

### GetFormatSourceOk

`func (o *FactureEntrante) GetFormatSourceOk() (*FormatFacture, bool)`

GetFormatSourceOk returns a tuple with the FormatSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatSource

`func (o *FactureEntrante) SetFormatSource(v FormatFacture)`

SetFormatSource sets FormatSource field to given value.


### GetRefFournisseur

`func (o *FactureEntrante) GetRefFournisseur() string`

GetRefFournisseur returns the RefFournisseur field if non-nil, zero value otherwise.

### GetRefFournisseurOk

`func (o *FactureEntrante) GetRefFournisseurOk() (*string, bool)`

GetRefFournisseurOk returns a tuple with the RefFournisseur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefFournisseur

`func (o *FactureEntrante) SetRefFournisseur(v string)`

SetRefFournisseur sets RefFournisseur field to given value.


### GetTypeDocument

`func (o *FactureEntrante) GetTypeDocument() TypeDocument`

GetTypeDocument returns the TypeDocument field if non-nil, zero value otherwise.

### GetTypeDocumentOk

`func (o *FactureEntrante) GetTypeDocumentOk() (*TypeDocument, bool)`

GetTypeDocumentOk returns a tuple with the TypeDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDocument

`func (o *FactureEntrante) SetTypeDocument(v TypeDocument)`

SetTypeDocument sets TypeDocument field to given value.

### HasTypeDocument

`func (o *FactureEntrante) HasTypeDocument() bool`

HasTypeDocument returns a boolean if a field has been set.

### GetFournisseur

`func (o *FactureEntrante) GetFournisseur() FournisseurEntrant`

GetFournisseur returns the Fournisseur field if non-nil, zero value otherwise.

### GetFournisseurOk

`func (o *FactureEntrante) GetFournisseurOk() (*FournisseurEntrant, bool)`

GetFournisseurOk returns a tuple with the Fournisseur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFournisseur

`func (o *FactureEntrante) SetFournisseur(v FournisseurEntrant)`

SetFournisseur sets Fournisseur field to given value.


### GetSiteFacturationNom

`func (o *FactureEntrante) GetSiteFacturationNom() string`

GetSiteFacturationNom returns the SiteFacturationNom field if non-nil, zero value otherwise.

### GetSiteFacturationNomOk

`func (o *FactureEntrante) GetSiteFacturationNomOk() (*string, bool)`

GetSiteFacturationNomOk returns a tuple with the SiteFacturationNom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteFacturationNom

`func (o *FactureEntrante) SetSiteFacturationNom(v string)`

SetSiteFacturationNom sets SiteFacturationNom field to given value.


### GetSiteFacturationSiret

`func (o *FactureEntrante) GetSiteFacturationSiret() string`

GetSiteFacturationSiret returns the SiteFacturationSiret field if non-nil, zero value otherwise.

### GetSiteFacturationSiretOk

`func (o *FactureEntrante) GetSiteFacturationSiretOk() (*string, bool)`

GetSiteFacturationSiretOk returns a tuple with the SiteFacturationSiret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteFacturationSiret

`func (o *FactureEntrante) SetSiteFacturationSiret(v string)`

SetSiteFacturationSiret sets SiteFacturationSiret field to given value.

### HasSiteFacturationSiret

`func (o *FactureEntrante) HasSiteFacturationSiret() bool`

HasSiteFacturationSiret returns a boolean if a field has been set.

### SetSiteFacturationSiretNil

`func (o *FactureEntrante) SetSiteFacturationSiretNil(b bool)`

 SetSiteFacturationSiretNil sets the value for SiteFacturationSiret to be an explicit nil

### UnsetSiteFacturationSiret
`func (o *FactureEntrante) UnsetSiteFacturationSiret()`

UnsetSiteFacturationSiret ensures that no value is present for SiteFacturationSiret, not even an explicit nil
### GetDateDePiece

`func (o *FactureEntrante) GetDateDePiece() string`

GetDateDePiece returns the DateDePiece field if non-nil, zero value otherwise.

### GetDateDePieceOk

`func (o *FactureEntrante) GetDateDePieceOk() (*string, bool)`

GetDateDePieceOk returns a tuple with the DateDePiece field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDePiece

`func (o *FactureEntrante) SetDateDePiece(v string)`

SetDateDePiece sets DateDePiece field to given value.


### GetDateReglement

`func (o *FactureEntrante) GetDateReglement() string`

GetDateReglement returns the DateReglement field if non-nil, zero value otherwise.

### GetDateReglementOk

`func (o *FactureEntrante) GetDateReglementOk() (*string, bool)`

GetDateReglementOk returns a tuple with the DateReglement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateReglement

`func (o *FactureEntrante) SetDateReglement(v string)`

SetDateReglement sets DateReglement field to given value.

### HasDateReglement

`func (o *FactureEntrante) HasDateReglement() bool`

HasDateReglement returns a boolean if a field has been set.

### SetDateReglementNil

`func (o *FactureEntrante) SetDateReglementNil(b bool)`

 SetDateReglementNil sets the value for DateReglement to be an explicit nil

### UnsetDateReglement
`func (o *FactureEntrante) UnsetDateReglement()`

UnsetDateReglement ensures that no value is present for DateReglement, not even an explicit nil
### GetDevise

`func (o *FactureEntrante) GetDevise() string`

GetDevise returns the Devise field if non-nil, zero value otherwise.

### GetDeviseOk

`func (o *FactureEntrante) GetDeviseOk() (*string, bool)`

GetDeviseOk returns a tuple with the Devise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevise

`func (o *FactureEntrante) SetDevise(v string)`

SetDevise sets Devise field to given value.

### HasDevise

`func (o *FactureEntrante) HasDevise() bool`

HasDevise returns a boolean if a field has been set.

### GetMontantHt

`func (o *FactureEntrante) GetMontantHt() string`

GetMontantHt returns the MontantHt field if non-nil, zero value otherwise.

### GetMontantHtOk

`func (o *FactureEntrante) GetMontantHtOk() (*string, bool)`

GetMontantHtOk returns a tuple with the MontantHt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantHt

`func (o *FactureEntrante) SetMontantHt(v string)`

SetMontantHt sets MontantHt field to given value.


### GetMontantTva

`func (o *FactureEntrante) GetMontantTva() string`

GetMontantTva returns the MontantTva field if non-nil, zero value otherwise.

### GetMontantTvaOk

`func (o *FactureEntrante) GetMontantTvaOk() (*string, bool)`

GetMontantTvaOk returns a tuple with the MontantTva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantTva

`func (o *FactureEntrante) SetMontantTva(v string)`

SetMontantTva sets MontantTva field to given value.


### GetMontantTtc

`func (o *FactureEntrante) GetMontantTtc() string`

GetMontantTtc returns the MontantTtc field if non-nil, zero value otherwise.

### GetMontantTtcOk

`func (o *FactureEntrante) GetMontantTtcOk() (*string, bool)`

GetMontantTtcOk returns a tuple with the MontantTtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantTtc

`func (o *FactureEntrante) SetMontantTtc(v string)`

SetMontantTtc sets MontantTtc field to given value.


### GetNumeroBonCommande

`func (o *FactureEntrante) GetNumeroBonCommande() string`

GetNumeroBonCommande returns the NumeroBonCommande field if non-nil, zero value otherwise.

### GetNumeroBonCommandeOk

`func (o *FactureEntrante) GetNumeroBonCommandeOk() (*string, bool)`

GetNumeroBonCommandeOk returns a tuple with the NumeroBonCommande field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroBonCommande

`func (o *FactureEntrante) SetNumeroBonCommande(v string)`

SetNumeroBonCommande sets NumeroBonCommande field to given value.

### HasNumeroBonCommande

`func (o *FactureEntrante) HasNumeroBonCommande() bool`

HasNumeroBonCommande returns a boolean if a field has been set.

### SetNumeroBonCommandeNil

`func (o *FactureEntrante) SetNumeroBonCommandeNil(b bool)`

 SetNumeroBonCommandeNil sets the value for NumeroBonCommande to be an explicit nil

### UnsetNumeroBonCommande
`func (o *FactureEntrante) UnsetNumeroBonCommande()`

UnsetNumeroBonCommande ensures that no value is present for NumeroBonCommande, not even an explicit nil
### GetReferenceContrat

`func (o *FactureEntrante) GetReferenceContrat() string`

GetReferenceContrat returns the ReferenceContrat field if non-nil, zero value otherwise.

### GetReferenceContratOk

`func (o *FactureEntrante) GetReferenceContratOk() (*string, bool)`

GetReferenceContratOk returns a tuple with the ReferenceContrat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceContrat

`func (o *FactureEntrante) SetReferenceContrat(v string)`

SetReferenceContrat sets ReferenceContrat field to given value.

### HasReferenceContrat

`func (o *FactureEntrante) HasReferenceContrat() bool`

HasReferenceContrat returns a boolean if a field has been set.

### SetReferenceContratNil

`func (o *FactureEntrante) SetReferenceContratNil(b bool)`

 SetReferenceContratNil sets the value for ReferenceContrat to be an explicit nil

### UnsetReferenceContrat
`func (o *FactureEntrante) UnsetReferenceContrat()`

UnsetReferenceContrat ensures that no value is present for ReferenceContrat, not even an explicit nil
### GetObjetFacture

`func (o *FactureEntrante) GetObjetFacture() string`

GetObjetFacture returns the ObjetFacture field if non-nil, zero value otherwise.

### GetObjetFactureOk

`func (o *FactureEntrante) GetObjetFactureOk() (*string, bool)`

GetObjetFactureOk returns a tuple with the ObjetFacture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjetFacture

`func (o *FactureEntrante) SetObjetFacture(v string)`

SetObjetFacture sets ObjetFacture field to given value.

### HasObjetFacture

`func (o *FactureEntrante) HasObjetFacture() bool`

HasObjetFacture returns a boolean if a field has been set.

### SetObjetFactureNil

`func (o *FactureEntrante) SetObjetFactureNil(b bool)`

 SetObjetFactureNil sets the value for ObjetFacture to be an explicit nil

### UnsetObjetFacture
`func (o *FactureEntrante) UnsetObjetFacture()`

UnsetObjetFacture ensures that no value is present for ObjetFacture, not even an explicit nil
### GetDocumentBase64

`func (o *FactureEntrante) GetDocumentBase64() string`

GetDocumentBase64 returns the DocumentBase64 field if non-nil, zero value otherwise.

### GetDocumentBase64Ok

`func (o *FactureEntrante) GetDocumentBase64Ok() (*string, bool)`

GetDocumentBase64Ok returns a tuple with the DocumentBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentBase64

`func (o *FactureEntrante) SetDocumentBase64(v string)`

SetDocumentBase64 sets DocumentBase64 field to given value.

### HasDocumentBase64

`func (o *FactureEntrante) HasDocumentBase64() bool`

HasDocumentBase64 returns a boolean if a field has been set.

### SetDocumentBase64Nil

`func (o *FactureEntrante) SetDocumentBase64Nil(b bool)`

 SetDocumentBase64Nil sets the value for DocumentBase64 to be an explicit nil

### UnsetDocumentBase64
`func (o *FactureEntrante) UnsetDocumentBase64()`

UnsetDocumentBase64 ensures that no value is present for DocumentBase64, not even an explicit nil
### GetDocumentContentType

`func (o *FactureEntrante) GetDocumentContentType() string`

GetDocumentContentType returns the DocumentContentType field if non-nil, zero value otherwise.

### GetDocumentContentTypeOk

`func (o *FactureEntrante) GetDocumentContentTypeOk() (*string, bool)`

GetDocumentContentTypeOk returns a tuple with the DocumentContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentContentType

`func (o *FactureEntrante) SetDocumentContentType(v string)`

SetDocumentContentType sets DocumentContentType field to given value.

### HasDocumentContentType

`func (o *FactureEntrante) HasDocumentContentType() bool`

HasDocumentContentType returns a boolean if a field has been set.

### SetDocumentContentTypeNil

`func (o *FactureEntrante) SetDocumentContentTypeNil(b bool)`

 SetDocumentContentTypeNil sets the value for DocumentContentType to be an explicit nil

### UnsetDocumentContentType
`func (o *FactureEntrante) UnsetDocumentContentType()`

UnsetDocumentContentType ensures that no value is present for DocumentContentType, not even an explicit nil
### GetDocumentFilename

`func (o *FactureEntrante) GetDocumentFilename() string`

GetDocumentFilename returns the DocumentFilename field if non-nil, zero value otherwise.

### GetDocumentFilenameOk

`func (o *FactureEntrante) GetDocumentFilenameOk() (*string, bool)`

GetDocumentFilenameOk returns a tuple with the DocumentFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentFilename

`func (o *FactureEntrante) SetDocumentFilename(v string)`

SetDocumentFilename sets DocumentFilename field to given value.

### HasDocumentFilename

`func (o *FactureEntrante) HasDocumentFilename() bool`

HasDocumentFilename returns a boolean if a field has been set.

### SetDocumentFilenameNil

`func (o *FactureEntrante) SetDocumentFilenameNil(b bool)`

 SetDocumentFilenameNil sets the value for DocumentFilename to be an explicit nil

### UnsetDocumentFilename
`func (o *FactureEntrante) UnsetDocumentFilename()`

UnsetDocumentFilename ensures that no value is present for DocumentFilename, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



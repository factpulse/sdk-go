# ReponseVerificationSucces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstConforme** | **bool** | True si aucun écart critique | 
**ScoreConformite** | **float32** | Score de conformité (0-100%) | 
**ChampsVerifies** | Pointer to **int32** | Nombre de champs vérifiés | [optional] [default to 0]
**ChampsConformes** | Pointer to **int32** | Nombre de champs conformes | [optional] [default to 0]
**EstFacturx** | Pointer to **bool** | True si le PDF contient du XML Factur-X | [optional] [default to false]
**ProfilFacturx** | Pointer to **NullableString** |  | [optional] 
**Champs** | Pointer to [**[]ChampVerifieSchema**](ChampVerifieSchema.md) | Liste des champs vérifiés avec valeurs, statuts et coordonnées PDF | [optional] 
**NotesObligatoires** | Pointer to [**[]NoteObligatoireSchema**](NoteObligatoireSchema.md) | Notes obligatoires (PMT, PMD, AAB) avec localisation PDF | [optional] 
**DimensionsPages** | Pointer to [**[]DimensionPageSchema**](DimensionPageSchema.md) | Dimensions de chaque page du PDF (largeur, hauteur) | [optional] 
**Avertissements** | Pointer to **[]string** | Avertissements non bloquants | [optional] 

## Methods

### NewReponseVerificationSucces

`func NewReponseVerificationSucces(estConforme bool, scoreConformite float32, ) *ReponseVerificationSucces`

NewReponseVerificationSucces instantiates a new ReponseVerificationSucces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReponseVerificationSuccesWithDefaults

`func NewReponseVerificationSuccesWithDefaults() *ReponseVerificationSucces`

NewReponseVerificationSuccesWithDefaults instantiates a new ReponseVerificationSucces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstConforme

`func (o *ReponseVerificationSucces) GetEstConforme() bool`

GetEstConforme returns the EstConforme field if non-nil, zero value otherwise.

### GetEstConformeOk

`func (o *ReponseVerificationSucces) GetEstConformeOk() (*bool, bool)`

GetEstConformeOk returns a tuple with the EstConforme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstConforme

`func (o *ReponseVerificationSucces) SetEstConforme(v bool)`

SetEstConforme sets EstConforme field to given value.


### GetScoreConformite

`func (o *ReponseVerificationSucces) GetScoreConformite() float32`

GetScoreConformite returns the ScoreConformite field if non-nil, zero value otherwise.

### GetScoreConformiteOk

`func (o *ReponseVerificationSucces) GetScoreConformiteOk() (*float32, bool)`

GetScoreConformiteOk returns a tuple with the ScoreConformite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreConformite

`func (o *ReponseVerificationSucces) SetScoreConformite(v float32)`

SetScoreConformite sets ScoreConformite field to given value.


### GetChampsVerifies

`func (o *ReponseVerificationSucces) GetChampsVerifies() int32`

GetChampsVerifies returns the ChampsVerifies field if non-nil, zero value otherwise.

### GetChampsVerifiesOk

`func (o *ReponseVerificationSucces) GetChampsVerifiesOk() (*int32, bool)`

GetChampsVerifiesOk returns a tuple with the ChampsVerifies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChampsVerifies

`func (o *ReponseVerificationSucces) SetChampsVerifies(v int32)`

SetChampsVerifies sets ChampsVerifies field to given value.

### HasChampsVerifies

`func (o *ReponseVerificationSucces) HasChampsVerifies() bool`

HasChampsVerifies returns a boolean if a field has been set.

### GetChampsConformes

`func (o *ReponseVerificationSucces) GetChampsConformes() int32`

GetChampsConformes returns the ChampsConformes field if non-nil, zero value otherwise.

### GetChampsConformesOk

`func (o *ReponseVerificationSucces) GetChampsConformesOk() (*int32, bool)`

GetChampsConformesOk returns a tuple with the ChampsConformes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChampsConformes

`func (o *ReponseVerificationSucces) SetChampsConformes(v int32)`

SetChampsConformes sets ChampsConformes field to given value.

### HasChampsConformes

`func (o *ReponseVerificationSucces) HasChampsConformes() bool`

HasChampsConformes returns a boolean if a field has been set.

### GetEstFacturx

`func (o *ReponseVerificationSucces) GetEstFacturx() bool`

GetEstFacturx returns the EstFacturx field if non-nil, zero value otherwise.

### GetEstFacturxOk

`func (o *ReponseVerificationSucces) GetEstFacturxOk() (*bool, bool)`

GetEstFacturxOk returns a tuple with the EstFacturx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstFacturx

`func (o *ReponseVerificationSucces) SetEstFacturx(v bool)`

SetEstFacturx sets EstFacturx field to given value.

### HasEstFacturx

`func (o *ReponseVerificationSucces) HasEstFacturx() bool`

HasEstFacturx returns a boolean if a field has been set.

### GetProfilFacturx

`func (o *ReponseVerificationSucces) GetProfilFacturx() string`

GetProfilFacturx returns the ProfilFacturx field if non-nil, zero value otherwise.

### GetProfilFacturxOk

`func (o *ReponseVerificationSucces) GetProfilFacturxOk() (*string, bool)`

GetProfilFacturxOk returns a tuple with the ProfilFacturx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilFacturx

`func (o *ReponseVerificationSucces) SetProfilFacturx(v string)`

SetProfilFacturx sets ProfilFacturx field to given value.

### HasProfilFacturx

`func (o *ReponseVerificationSucces) HasProfilFacturx() bool`

HasProfilFacturx returns a boolean if a field has been set.

### SetProfilFacturxNil

`func (o *ReponseVerificationSucces) SetProfilFacturxNil(b bool)`

 SetProfilFacturxNil sets the value for ProfilFacturx to be an explicit nil

### UnsetProfilFacturx
`func (o *ReponseVerificationSucces) UnsetProfilFacturx()`

UnsetProfilFacturx ensures that no value is present for ProfilFacturx, not even an explicit nil
### GetChamps

`func (o *ReponseVerificationSucces) GetChamps() []ChampVerifieSchema`

GetChamps returns the Champs field if non-nil, zero value otherwise.

### GetChampsOk

`func (o *ReponseVerificationSucces) GetChampsOk() (*[]ChampVerifieSchema, bool)`

GetChampsOk returns a tuple with the Champs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChamps

`func (o *ReponseVerificationSucces) SetChamps(v []ChampVerifieSchema)`

SetChamps sets Champs field to given value.

### HasChamps

`func (o *ReponseVerificationSucces) HasChamps() bool`

HasChamps returns a boolean if a field has been set.

### GetNotesObligatoires

`func (o *ReponseVerificationSucces) GetNotesObligatoires() []NoteObligatoireSchema`

GetNotesObligatoires returns the NotesObligatoires field if non-nil, zero value otherwise.

### GetNotesObligatoiresOk

`func (o *ReponseVerificationSucces) GetNotesObligatoiresOk() (*[]NoteObligatoireSchema, bool)`

GetNotesObligatoiresOk returns a tuple with the NotesObligatoires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesObligatoires

`func (o *ReponseVerificationSucces) SetNotesObligatoires(v []NoteObligatoireSchema)`

SetNotesObligatoires sets NotesObligatoires field to given value.

### HasNotesObligatoires

`func (o *ReponseVerificationSucces) HasNotesObligatoires() bool`

HasNotesObligatoires returns a boolean if a field has been set.

### GetDimensionsPages

`func (o *ReponseVerificationSucces) GetDimensionsPages() []DimensionPageSchema`

GetDimensionsPages returns the DimensionsPages field if non-nil, zero value otherwise.

### GetDimensionsPagesOk

`func (o *ReponseVerificationSucces) GetDimensionsPagesOk() (*[]DimensionPageSchema, bool)`

GetDimensionsPagesOk returns a tuple with the DimensionsPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionsPages

`func (o *ReponseVerificationSucces) SetDimensionsPages(v []DimensionPageSchema)`

SetDimensionsPages sets DimensionsPages field to given value.

### HasDimensionsPages

`func (o *ReponseVerificationSucces) HasDimensionsPages() bool`

HasDimensionsPages returns a boolean if a field has been set.

### GetAvertissements

`func (o *ReponseVerificationSucces) GetAvertissements() []string`

GetAvertissements returns the Avertissements field if non-nil, zero value otherwise.

### GetAvertissementsOk

`func (o *ReponseVerificationSucces) GetAvertissementsOk() (*[]string, bool)`

GetAvertissementsOk returns a tuple with the Avertissements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvertissements

`func (o *ReponseVerificationSucces) SetAvertissements(v []string)`

SetAvertissements sets Avertissements field to given value.

### HasAvertissements

`func (o *ReponseVerificationSucces) HasAvertissements() bool`

HasAvertissements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



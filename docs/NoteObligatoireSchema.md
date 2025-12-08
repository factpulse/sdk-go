# NoteObligatoireSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodeSujet** | **string** | Code sujet (PMT, PMD, AAB) | 
**Label** | **string** | Libellé (ex: Indemnité recouvrement) | 
**ValeurPdf** | Pointer to **NullableString** |  | [optional] 
**ValeurXml** | Pointer to **NullableString** |  | [optional] 
**Statut** | Pointer to [**StatutChampAPI**](StatutChampAPI.md) | Statut de conformité (CONFORME si XML trouvé dans PDF) | [optional] [default to NON_VERIFIE]
**Message** | Pointer to **NullableString** |  | [optional] 
**Bbox** | Pointer to [**NullableBoundingBoxSchema**](BoundingBoxSchema.md) |  | [optional] 

## Methods

### NewNoteObligatoireSchema

`func NewNoteObligatoireSchema(codeSujet string, label string, ) *NoteObligatoireSchema`

NewNoteObligatoireSchema instantiates a new NoteObligatoireSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteObligatoireSchemaWithDefaults

`func NewNoteObligatoireSchemaWithDefaults() *NoteObligatoireSchema`

NewNoteObligatoireSchemaWithDefaults instantiates a new NoteObligatoireSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodeSujet

`func (o *NoteObligatoireSchema) GetCodeSujet() string`

GetCodeSujet returns the CodeSujet field if non-nil, zero value otherwise.

### GetCodeSujetOk

`func (o *NoteObligatoireSchema) GetCodeSujetOk() (*string, bool)`

GetCodeSujetOk returns a tuple with the CodeSujet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeSujet

`func (o *NoteObligatoireSchema) SetCodeSujet(v string)`

SetCodeSujet sets CodeSujet field to given value.


### GetLabel

`func (o *NoteObligatoireSchema) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *NoteObligatoireSchema) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *NoteObligatoireSchema) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetValeurPdf

`func (o *NoteObligatoireSchema) GetValeurPdf() string`

GetValeurPdf returns the ValeurPdf field if non-nil, zero value otherwise.

### GetValeurPdfOk

`func (o *NoteObligatoireSchema) GetValeurPdfOk() (*string, bool)`

GetValeurPdfOk returns a tuple with the ValeurPdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValeurPdf

`func (o *NoteObligatoireSchema) SetValeurPdf(v string)`

SetValeurPdf sets ValeurPdf field to given value.

### HasValeurPdf

`func (o *NoteObligatoireSchema) HasValeurPdf() bool`

HasValeurPdf returns a boolean if a field has been set.

### SetValeurPdfNil

`func (o *NoteObligatoireSchema) SetValeurPdfNil(b bool)`

 SetValeurPdfNil sets the value for ValeurPdf to be an explicit nil

### UnsetValeurPdf
`func (o *NoteObligatoireSchema) UnsetValeurPdf()`

UnsetValeurPdf ensures that no value is present for ValeurPdf, not even an explicit nil
### GetValeurXml

`func (o *NoteObligatoireSchema) GetValeurXml() string`

GetValeurXml returns the ValeurXml field if non-nil, zero value otherwise.

### GetValeurXmlOk

`func (o *NoteObligatoireSchema) GetValeurXmlOk() (*string, bool)`

GetValeurXmlOk returns a tuple with the ValeurXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValeurXml

`func (o *NoteObligatoireSchema) SetValeurXml(v string)`

SetValeurXml sets ValeurXml field to given value.

### HasValeurXml

`func (o *NoteObligatoireSchema) HasValeurXml() bool`

HasValeurXml returns a boolean if a field has been set.

### SetValeurXmlNil

`func (o *NoteObligatoireSchema) SetValeurXmlNil(b bool)`

 SetValeurXmlNil sets the value for ValeurXml to be an explicit nil

### UnsetValeurXml
`func (o *NoteObligatoireSchema) UnsetValeurXml()`

UnsetValeurXml ensures that no value is present for ValeurXml, not even an explicit nil
### GetStatut

`func (o *NoteObligatoireSchema) GetStatut() StatutChampAPI`

GetStatut returns the Statut field if non-nil, zero value otherwise.

### GetStatutOk

`func (o *NoteObligatoireSchema) GetStatutOk() (*StatutChampAPI, bool)`

GetStatutOk returns a tuple with the Statut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatut

`func (o *NoteObligatoireSchema) SetStatut(v StatutChampAPI)`

SetStatut sets Statut field to given value.

### HasStatut

`func (o *NoteObligatoireSchema) HasStatut() bool`

HasStatut returns a boolean if a field has been set.

### GetMessage

`func (o *NoteObligatoireSchema) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NoteObligatoireSchema) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NoteObligatoireSchema) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NoteObligatoireSchema) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *NoteObligatoireSchema) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *NoteObligatoireSchema) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetBbox

`func (o *NoteObligatoireSchema) GetBbox() BoundingBoxSchema`

GetBbox returns the Bbox field if non-nil, zero value otherwise.

### GetBboxOk

`func (o *NoteObligatoireSchema) GetBboxOk() (*BoundingBoxSchema, bool)`

GetBboxOk returns a tuple with the Bbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBbox

`func (o *NoteObligatoireSchema) SetBbox(v BoundingBoxSchema)`

SetBbox sets Bbox field to given value.

### HasBbox

`func (o *NoteObligatoireSchema) HasBbox() bool`

HasBbox returns a boolean if a field has been set.

### SetBboxNil

`func (o *NoteObligatoireSchema) SetBboxNil(b bool)`

 SetBboxNil sets the value for Bbox to be an explicit nil

### UnsetBbox
`func (o *NoteObligatoireSchema) UnsetBbox()`

UnsetBbox ensures that no value is present for Bbox, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



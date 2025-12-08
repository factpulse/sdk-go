# ChampVerifieSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessTerm** | **string** | Business Term EN16931 (ex: BT-1) | 
**Label** | **string** | Libellé du champ (ex: N° Facture) | 
**ValeurPdf** | Pointer to **NullableString** |  | [optional] 
**ValeurXml** | Pointer to **NullableString** |  | [optional] 
**Statut** | [**StatutChampAPI**](StatutChampAPI.md) | Statut de conformité | 
**Message** | Pointer to **NullableString** |  | [optional] 
**Confiance** | Pointer to **float32** | Score de confiance (0-1) | [optional] [default to 1.0]
**Source** | Pointer to **string** | Source d&#39;extraction | [optional] [default to "pdf_natif"]
**Bbox** | Pointer to [**NullableBoundingBoxSchema**](BoundingBoxSchema.md) |  | [optional] 

## Methods

### NewChampVerifieSchema

`func NewChampVerifieSchema(businessTerm string, label string, statut StatutChampAPI, ) *ChampVerifieSchema`

NewChampVerifieSchema instantiates a new ChampVerifieSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChampVerifieSchemaWithDefaults

`func NewChampVerifieSchemaWithDefaults() *ChampVerifieSchema`

NewChampVerifieSchemaWithDefaults instantiates a new ChampVerifieSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessTerm

`func (o *ChampVerifieSchema) GetBusinessTerm() string`

GetBusinessTerm returns the BusinessTerm field if non-nil, zero value otherwise.

### GetBusinessTermOk

`func (o *ChampVerifieSchema) GetBusinessTermOk() (*string, bool)`

GetBusinessTermOk returns a tuple with the BusinessTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessTerm

`func (o *ChampVerifieSchema) SetBusinessTerm(v string)`

SetBusinessTerm sets BusinessTerm field to given value.


### GetLabel

`func (o *ChampVerifieSchema) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ChampVerifieSchema) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ChampVerifieSchema) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetValeurPdf

`func (o *ChampVerifieSchema) GetValeurPdf() string`

GetValeurPdf returns the ValeurPdf field if non-nil, zero value otherwise.

### GetValeurPdfOk

`func (o *ChampVerifieSchema) GetValeurPdfOk() (*string, bool)`

GetValeurPdfOk returns a tuple with the ValeurPdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValeurPdf

`func (o *ChampVerifieSchema) SetValeurPdf(v string)`

SetValeurPdf sets ValeurPdf field to given value.

### HasValeurPdf

`func (o *ChampVerifieSchema) HasValeurPdf() bool`

HasValeurPdf returns a boolean if a field has been set.

### SetValeurPdfNil

`func (o *ChampVerifieSchema) SetValeurPdfNil(b bool)`

 SetValeurPdfNil sets the value for ValeurPdf to be an explicit nil

### UnsetValeurPdf
`func (o *ChampVerifieSchema) UnsetValeurPdf()`

UnsetValeurPdf ensures that no value is present for ValeurPdf, not even an explicit nil
### GetValeurXml

`func (o *ChampVerifieSchema) GetValeurXml() string`

GetValeurXml returns the ValeurXml field if non-nil, zero value otherwise.

### GetValeurXmlOk

`func (o *ChampVerifieSchema) GetValeurXmlOk() (*string, bool)`

GetValeurXmlOk returns a tuple with the ValeurXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValeurXml

`func (o *ChampVerifieSchema) SetValeurXml(v string)`

SetValeurXml sets ValeurXml field to given value.

### HasValeurXml

`func (o *ChampVerifieSchema) HasValeurXml() bool`

HasValeurXml returns a boolean if a field has been set.

### SetValeurXmlNil

`func (o *ChampVerifieSchema) SetValeurXmlNil(b bool)`

 SetValeurXmlNil sets the value for ValeurXml to be an explicit nil

### UnsetValeurXml
`func (o *ChampVerifieSchema) UnsetValeurXml()`

UnsetValeurXml ensures that no value is present for ValeurXml, not even an explicit nil
### GetStatut

`func (o *ChampVerifieSchema) GetStatut() StatutChampAPI`

GetStatut returns the Statut field if non-nil, zero value otherwise.

### GetStatutOk

`func (o *ChampVerifieSchema) GetStatutOk() (*StatutChampAPI, bool)`

GetStatutOk returns a tuple with the Statut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatut

`func (o *ChampVerifieSchema) SetStatut(v StatutChampAPI)`

SetStatut sets Statut field to given value.


### GetMessage

`func (o *ChampVerifieSchema) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ChampVerifieSchema) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ChampVerifieSchema) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ChampVerifieSchema) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ChampVerifieSchema) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ChampVerifieSchema) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetConfiance

`func (o *ChampVerifieSchema) GetConfiance() float32`

GetConfiance returns the Confiance field if non-nil, zero value otherwise.

### GetConfianceOk

`func (o *ChampVerifieSchema) GetConfianceOk() (*float32, bool)`

GetConfianceOk returns a tuple with the Confiance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiance

`func (o *ChampVerifieSchema) SetConfiance(v float32)`

SetConfiance sets Confiance field to given value.

### HasConfiance

`func (o *ChampVerifieSchema) HasConfiance() bool`

HasConfiance returns a boolean if a field has been set.

### GetSource

`func (o *ChampVerifieSchema) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ChampVerifieSchema) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ChampVerifieSchema) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ChampVerifieSchema) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetBbox

`func (o *ChampVerifieSchema) GetBbox() BoundingBoxSchema`

GetBbox returns the Bbox field if non-nil, zero value otherwise.

### GetBboxOk

`func (o *ChampVerifieSchema) GetBboxOk() (*BoundingBoxSchema, bool)`

GetBboxOk returns a tuple with the Bbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBbox

`func (o *ChampVerifieSchema) SetBbox(v BoundingBoxSchema)`

SetBbox sets Bbox field to given value.

### HasBbox

`func (o *ChampVerifieSchema) HasBbox() bool`

HasBbox returns a boolean if a field has been set.

### SetBboxNil

`func (o *ChampVerifieSchema) SetBboxNil(b bool)`

 SetBboxNil sets the value for Bbox to be an explicit nil

### UnsetBbox
`func (o *ChampVerifieSchema) UnsetBbox()`

UnsetBbox ensures that no value is present for Bbox, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



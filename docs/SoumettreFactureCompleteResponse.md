# SoumettreFactureCompleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Succes** | **bool** | La facture a été soumise avec succès | 
**DestinationType** | **string** | Type de destination | 
**ResultatChorus** | Pointer to [**NullableResultatChorusPro**](ResultatChorusPro.md) |  | [optional] 
**ResultatAfnor** | Pointer to [**NullableResultatAFNOR**](ResultatAFNOR.md) |  | [optional] 
**FactureEnrichie** | [**FactureEnrichieInfoOutput**](FactureEnrichieInfoOutput.md) | Données de la facture enrichie | 
**PdfFacturx** | [**PDFFacturXInfo**](PDFFacturXInfo.md) | Informations sur le PDF généré | 
**Signature** | Pointer to [**NullableSignatureInfo**](SignatureInfo.md) |  | [optional] 
**PdfBase64** | **string** | PDF Factur-X généré (et signé si demandé) encodé en base64 | 
**Message** | **string** | Message de retour | 

## Methods

### NewSoumettreFactureCompleteResponse

`func NewSoumettreFactureCompleteResponse(succes bool, destinationType string, factureEnrichie FactureEnrichieInfoOutput, pdfFacturx PDFFacturXInfo, pdfBase64 string, message string, ) *SoumettreFactureCompleteResponse`

NewSoumettreFactureCompleteResponse instantiates a new SoumettreFactureCompleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoumettreFactureCompleteResponseWithDefaults

`func NewSoumettreFactureCompleteResponseWithDefaults() *SoumettreFactureCompleteResponse`

NewSoumettreFactureCompleteResponseWithDefaults instantiates a new SoumettreFactureCompleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSucces

`func (o *SoumettreFactureCompleteResponse) GetSucces() bool`

GetSucces returns the Succes field if non-nil, zero value otherwise.

### GetSuccesOk

`func (o *SoumettreFactureCompleteResponse) GetSuccesOk() (*bool, bool)`

GetSuccesOk returns a tuple with the Succes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucces

`func (o *SoumettreFactureCompleteResponse) SetSucces(v bool)`

SetSucces sets Succes field to given value.


### GetDestinationType

`func (o *SoumettreFactureCompleteResponse) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *SoumettreFactureCompleteResponse) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *SoumettreFactureCompleteResponse) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.


### GetResultatChorus

`func (o *SoumettreFactureCompleteResponse) GetResultatChorus() ResultatChorusPro`

GetResultatChorus returns the ResultatChorus field if non-nil, zero value otherwise.

### GetResultatChorusOk

`func (o *SoumettreFactureCompleteResponse) GetResultatChorusOk() (*ResultatChorusPro, bool)`

GetResultatChorusOk returns a tuple with the ResultatChorus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultatChorus

`func (o *SoumettreFactureCompleteResponse) SetResultatChorus(v ResultatChorusPro)`

SetResultatChorus sets ResultatChorus field to given value.

### HasResultatChorus

`func (o *SoumettreFactureCompleteResponse) HasResultatChorus() bool`

HasResultatChorus returns a boolean if a field has been set.

### SetResultatChorusNil

`func (o *SoumettreFactureCompleteResponse) SetResultatChorusNil(b bool)`

 SetResultatChorusNil sets the value for ResultatChorus to be an explicit nil

### UnsetResultatChorus
`func (o *SoumettreFactureCompleteResponse) UnsetResultatChorus()`

UnsetResultatChorus ensures that no value is present for ResultatChorus, not even an explicit nil
### GetResultatAfnor

`func (o *SoumettreFactureCompleteResponse) GetResultatAfnor() ResultatAFNOR`

GetResultatAfnor returns the ResultatAfnor field if non-nil, zero value otherwise.

### GetResultatAfnorOk

`func (o *SoumettreFactureCompleteResponse) GetResultatAfnorOk() (*ResultatAFNOR, bool)`

GetResultatAfnorOk returns a tuple with the ResultatAfnor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultatAfnor

`func (o *SoumettreFactureCompleteResponse) SetResultatAfnor(v ResultatAFNOR)`

SetResultatAfnor sets ResultatAfnor field to given value.

### HasResultatAfnor

`func (o *SoumettreFactureCompleteResponse) HasResultatAfnor() bool`

HasResultatAfnor returns a boolean if a field has been set.

### SetResultatAfnorNil

`func (o *SoumettreFactureCompleteResponse) SetResultatAfnorNil(b bool)`

 SetResultatAfnorNil sets the value for ResultatAfnor to be an explicit nil

### UnsetResultatAfnor
`func (o *SoumettreFactureCompleteResponse) UnsetResultatAfnor()`

UnsetResultatAfnor ensures that no value is present for ResultatAfnor, not even an explicit nil
### GetFactureEnrichie

`func (o *SoumettreFactureCompleteResponse) GetFactureEnrichie() FactureEnrichieInfoOutput`

GetFactureEnrichie returns the FactureEnrichie field if non-nil, zero value otherwise.

### GetFactureEnrichieOk

`func (o *SoumettreFactureCompleteResponse) GetFactureEnrichieOk() (*FactureEnrichieInfoOutput, bool)`

GetFactureEnrichieOk returns a tuple with the FactureEnrichie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactureEnrichie

`func (o *SoumettreFactureCompleteResponse) SetFactureEnrichie(v FactureEnrichieInfoOutput)`

SetFactureEnrichie sets FactureEnrichie field to given value.


### GetPdfFacturx

`func (o *SoumettreFactureCompleteResponse) GetPdfFacturx() PDFFacturXInfo`

GetPdfFacturx returns the PdfFacturx field if non-nil, zero value otherwise.

### GetPdfFacturxOk

`func (o *SoumettreFactureCompleteResponse) GetPdfFacturxOk() (*PDFFacturXInfo, bool)`

GetPdfFacturxOk returns a tuple with the PdfFacturx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfFacturx

`func (o *SoumettreFactureCompleteResponse) SetPdfFacturx(v PDFFacturXInfo)`

SetPdfFacturx sets PdfFacturx field to given value.


### GetSignature

`func (o *SoumettreFactureCompleteResponse) GetSignature() SignatureInfo`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *SoumettreFactureCompleteResponse) GetSignatureOk() (*SignatureInfo, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *SoumettreFactureCompleteResponse) SetSignature(v SignatureInfo)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *SoumettreFactureCompleteResponse) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### SetSignatureNil

`func (o *SoumettreFactureCompleteResponse) SetSignatureNil(b bool)`

 SetSignatureNil sets the value for Signature to be an explicit nil

### UnsetSignature
`func (o *SoumettreFactureCompleteResponse) UnsetSignature()`

UnsetSignature ensures that no value is present for Signature, not even an explicit nil
### GetPdfBase64

`func (o *SoumettreFactureCompleteResponse) GetPdfBase64() string`

GetPdfBase64 returns the PdfBase64 field if non-nil, zero value otherwise.

### GetPdfBase64Ok

`func (o *SoumettreFactureCompleteResponse) GetPdfBase64Ok() (*string, bool)`

GetPdfBase64Ok returns a tuple with the PdfBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfBase64

`func (o *SoumettreFactureCompleteResponse) SetPdfBase64(v string)`

SetPdfBase64 sets PdfBase64 field to given value.


### GetMessage

`func (o *SoumettreFactureCompleteResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SoumettreFactureCompleteResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SoumettreFactureCompleteResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



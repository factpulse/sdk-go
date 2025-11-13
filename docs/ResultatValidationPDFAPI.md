# ResultatValidationPDFAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstConforme** | **bool** | True si le PDF est conforme à tous les critères (XML, PDF/A, XMP) | 
**XmlPresent** | **bool** | True si un XML Factur-X est embarqué dans le PDF | 
**XmlConforme** | **bool** | True si le XML Factur-X est conforme aux règles Schematron | 
**ProfilDetecte** | Pointer to **NullableString** |  | [optional] 
**ErreursXml** | Pointer to **[]string** | Liste des erreurs de validation XML | [optional] 
**PdfaConforme** | **bool** | True si le PDF est conforme PDF/A | 
**VersionPdfa** | Pointer to **NullableString** |  | [optional] 
**MethodeValidationPdfa** | Pointer to **string** | Méthode utilisée pour la validation PDF/A (metadata ou verapdf) | [optional] [default to "metadata"]
**ReglesValidees** | Pointer to **NullableInt32** |  | [optional] 
**ReglesEchouees** | Pointer to **NullableInt32** |  | [optional] 
**ErreursPdfa** | Pointer to **[]string** | Liste des erreurs de conformité PDF/A | [optional] 
**AvertissementsPdfa** | Pointer to **[]string** | Liste des avertissements PDF/A | [optional] 
**XmpPresent** | **bool** | True si des métadonnées XMP sont présentes | 
**XmpConformeFacturx** | **bool** | True si les métadonnées XMP contiennent des informations Factur-X | 
**ProfilXmp** | Pointer to **NullableString** |  | [optional] 
**VersionXmp** | Pointer to **NullableString** |  | [optional] 
**ErreursXmp** | Pointer to **[]string** | Liste des erreurs de métadonnées XMP | [optional] 
**MetadonneesXmp** | Pointer to **map[string]interface{}** | Métadonnées XMP extraites du PDF | [optional] 
**EstSigne** | **bool** | True si le PDF contient au moins une signature | 
**NombreSignatures** | Pointer to **int32** | Nombre de signatures électroniques trouvées | [optional] [default to 0]
**Signatures** | Pointer to [**[]InformationSignatureAPI**](InformationSignatureAPI.md) | Liste des signatures trouvées avec leurs informations | [optional] 
**ErreursSignatures** | Pointer to **[]string** | Liste des erreurs lors de l&#39;analyse des signatures | [optional] 
**MessageResume** | **string** | Message résumant le résultat de la validation | 

## Methods

### NewResultatValidationPDFAPI

`func NewResultatValidationPDFAPI(estConforme bool, xmlPresent bool, xmlConforme bool, pdfaConforme bool, xmpPresent bool, xmpConformeFacturx bool, estSigne bool, messageResume string, ) *ResultatValidationPDFAPI`

NewResultatValidationPDFAPI instantiates a new ResultatValidationPDFAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultatValidationPDFAPIWithDefaults

`func NewResultatValidationPDFAPIWithDefaults() *ResultatValidationPDFAPI`

NewResultatValidationPDFAPIWithDefaults instantiates a new ResultatValidationPDFAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstConforme

`func (o *ResultatValidationPDFAPI) GetEstConforme() bool`

GetEstConforme returns the EstConforme field if non-nil, zero value otherwise.

### GetEstConformeOk

`func (o *ResultatValidationPDFAPI) GetEstConformeOk() (*bool, bool)`

GetEstConformeOk returns a tuple with the EstConforme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstConforme

`func (o *ResultatValidationPDFAPI) SetEstConforme(v bool)`

SetEstConforme sets EstConforme field to given value.


### GetXmlPresent

`func (o *ResultatValidationPDFAPI) GetXmlPresent() bool`

GetXmlPresent returns the XmlPresent field if non-nil, zero value otherwise.

### GetXmlPresentOk

`func (o *ResultatValidationPDFAPI) GetXmlPresentOk() (*bool, bool)`

GetXmlPresentOk returns a tuple with the XmlPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlPresent

`func (o *ResultatValidationPDFAPI) SetXmlPresent(v bool)`

SetXmlPresent sets XmlPresent field to given value.


### GetXmlConforme

`func (o *ResultatValidationPDFAPI) GetXmlConforme() bool`

GetXmlConforme returns the XmlConforme field if non-nil, zero value otherwise.

### GetXmlConformeOk

`func (o *ResultatValidationPDFAPI) GetXmlConformeOk() (*bool, bool)`

GetXmlConformeOk returns a tuple with the XmlConforme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlConforme

`func (o *ResultatValidationPDFAPI) SetXmlConforme(v bool)`

SetXmlConforme sets XmlConforme field to given value.


### GetProfilDetecte

`func (o *ResultatValidationPDFAPI) GetProfilDetecte() string`

GetProfilDetecte returns the ProfilDetecte field if non-nil, zero value otherwise.

### GetProfilDetecteOk

`func (o *ResultatValidationPDFAPI) GetProfilDetecteOk() (*string, bool)`

GetProfilDetecteOk returns a tuple with the ProfilDetecte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilDetecte

`func (o *ResultatValidationPDFAPI) SetProfilDetecte(v string)`

SetProfilDetecte sets ProfilDetecte field to given value.

### HasProfilDetecte

`func (o *ResultatValidationPDFAPI) HasProfilDetecte() bool`

HasProfilDetecte returns a boolean if a field has been set.

### SetProfilDetecteNil

`func (o *ResultatValidationPDFAPI) SetProfilDetecteNil(b bool)`

 SetProfilDetecteNil sets the value for ProfilDetecte to be an explicit nil

### UnsetProfilDetecte
`func (o *ResultatValidationPDFAPI) UnsetProfilDetecte()`

UnsetProfilDetecte ensures that no value is present for ProfilDetecte, not even an explicit nil
### GetErreursXml

`func (o *ResultatValidationPDFAPI) GetErreursXml() []string`

GetErreursXml returns the ErreursXml field if non-nil, zero value otherwise.

### GetErreursXmlOk

`func (o *ResultatValidationPDFAPI) GetErreursXmlOk() (*[]string, bool)`

GetErreursXmlOk returns a tuple with the ErreursXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErreursXml

`func (o *ResultatValidationPDFAPI) SetErreursXml(v []string)`

SetErreursXml sets ErreursXml field to given value.

### HasErreursXml

`func (o *ResultatValidationPDFAPI) HasErreursXml() bool`

HasErreursXml returns a boolean if a field has been set.

### GetPdfaConforme

`func (o *ResultatValidationPDFAPI) GetPdfaConforme() bool`

GetPdfaConforme returns the PdfaConforme field if non-nil, zero value otherwise.

### GetPdfaConformeOk

`func (o *ResultatValidationPDFAPI) GetPdfaConformeOk() (*bool, bool)`

GetPdfaConformeOk returns a tuple with the PdfaConforme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfaConforme

`func (o *ResultatValidationPDFAPI) SetPdfaConforme(v bool)`

SetPdfaConforme sets PdfaConforme field to given value.


### GetVersionPdfa

`func (o *ResultatValidationPDFAPI) GetVersionPdfa() string`

GetVersionPdfa returns the VersionPdfa field if non-nil, zero value otherwise.

### GetVersionPdfaOk

`func (o *ResultatValidationPDFAPI) GetVersionPdfaOk() (*string, bool)`

GetVersionPdfaOk returns a tuple with the VersionPdfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionPdfa

`func (o *ResultatValidationPDFAPI) SetVersionPdfa(v string)`

SetVersionPdfa sets VersionPdfa field to given value.

### HasVersionPdfa

`func (o *ResultatValidationPDFAPI) HasVersionPdfa() bool`

HasVersionPdfa returns a boolean if a field has been set.

### SetVersionPdfaNil

`func (o *ResultatValidationPDFAPI) SetVersionPdfaNil(b bool)`

 SetVersionPdfaNil sets the value for VersionPdfa to be an explicit nil

### UnsetVersionPdfa
`func (o *ResultatValidationPDFAPI) UnsetVersionPdfa()`

UnsetVersionPdfa ensures that no value is present for VersionPdfa, not even an explicit nil
### GetMethodeValidationPdfa

`func (o *ResultatValidationPDFAPI) GetMethodeValidationPdfa() string`

GetMethodeValidationPdfa returns the MethodeValidationPdfa field if non-nil, zero value otherwise.

### GetMethodeValidationPdfaOk

`func (o *ResultatValidationPDFAPI) GetMethodeValidationPdfaOk() (*string, bool)`

GetMethodeValidationPdfaOk returns a tuple with the MethodeValidationPdfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodeValidationPdfa

`func (o *ResultatValidationPDFAPI) SetMethodeValidationPdfa(v string)`

SetMethodeValidationPdfa sets MethodeValidationPdfa field to given value.

### HasMethodeValidationPdfa

`func (o *ResultatValidationPDFAPI) HasMethodeValidationPdfa() bool`

HasMethodeValidationPdfa returns a boolean if a field has been set.

### GetReglesValidees

`func (o *ResultatValidationPDFAPI) GetReglesValidees() int32`

GetReglesValidees returns the ReglesValidees field if non-nil, zero value otherwise.

### GetReglesValideesOk

`func (o *ResultatValidationPDFAPI) GetReglesValideesOk() (*int32, bool)`

GetReglesValideesOk returns a tuple with the ReglesValidees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReglesValidees

`func (o *ResultatValidationPDFAPI) SetReglesValidees(v int32)`

SetReglesValidees sets ReglesValidees field to given value.

### HasReglesValidees

`func (o *ResultatValidationPDFAPI) HasReglesValidees() bool`

HasReglesValidees returns a boolean if a field has been set.

### SetReglesValideesNil

`func (o *ResultatValidationPDFAPI) SetReglesValideesNil(b bool)`

 SetReglesValideesNil sets the value for ReglesValidees to be an explicit nil

### UnsetReglesValidees
`func (o *ResultatValidationPDFAPI) UnsetReglesValidees()`

UnsetReglesValidees ensures that no value is present for ReglesValidees, not even an explicit nil
### GetReglesEchouees

`func (o *ResultatValidationPDFAPI) GetReglesEchouees() int32`

GetReglesEchouees returns the ReglesEchouees field if non-nil, zero value otherwise.

### GetReglesEchoueesOk

`func (o *ResultatValidationPDFAPI) GetReglesEchoueesOk() (*int32, bool)`

GetReglesEchoueesOk returns a tuple with the ReglesEchouees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReglesEchouees

`func (o *ResultatValidationPDFAPI) SetReglesEchouees(v int32)`

SetReglesEchouees sets ReglesEchouees field to given value.

### HasReglesEchouees

`func (o *ResultatValidationPDFAPI) HasReglesEchouees() bool`

HasReglesEchouees returns a boolean if a field has been set.

### SetReglesEchoueesNil

`func (o *ResultatValidationPDFAPI) SetReglesEchoueesNil(b bool)`

 SetReglesEchoueesNil sets the value for ReglesEchouees to be an explicit nil

### UnsetReglesEchouees
`func (o *ResultatValidationPDFAPI) UnsetReglesEchouees()`

UnsetReglesEchouees ensures that no value is present for ReglesEchouees, not even an explicit nil
### GetErreursPdfa

`func (o *ResultatValidationPDFAPI) GetErreursPdfa() []string`

GetErreursPdfa returns the ErreursPdfa field if non-nil, zero value otherwise.

### GetErreursPdfaOk

`func (o *ResultatValidationPDFAPI) GetErreursPdfaOk() (*[]string, bool)`

GetErreursPdfaOk returns a tuple with the ErreursPdfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErreursPdfa

`func (o *ResultatValidationPDFAPI) SetErreursPdfa(v []string)`

SetErreursPdfa sets ErreursPdfa field to given value.

### HasErreursPdfa

`func (o *ResultatValidationPDFAPI) HasErreursPdfa() bool`

HasErreursPdfa returns a boolean if a field has been set.

### GetAvertissementsPdfa

`func (o *ResultatValidationPDFAPI) GetAvertissementsPdfa() []string`

GetAvertissementsPdfa returns the AvertissementsPdfa field if non-nil, zero value otherwise.

### GetAvertissementsPdfaOk

`func (o *ResultatValidationPDFAPI) GetAvertissementsPdfaOk() (*[]string, bool)`

GetAvertissementsPdfaOk returns a tuple with the AvertissementsPdfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvertissementsPdfa

`func (o *ResultatValidationPDFAPI) SetAvertissementsPdfa(v []string)`

SetAvertissementsPdfa sets AvertissementsPdfa field to given value.

### HasAvertissementsPdfa

`func (o *ResultatValidationPDFAPI) HasAvertissementsPdfa() bool`

HasAvertissementsPdfa returns a boolean if a field has been set.

### GetXmpPresent

`func (o *ResultatValidationPDFAPI) GetXmpPresent() bool`

GetXmpPresent returns the XmpPresent field if non-nil, zero value otherwise.

### GetXmpPresentOk

`func (o *ResultatValidationPDFAPI) GetXmpPresentOk() (*bool, bool)`

GetXmpPresentOk returns a tuple with the XmpPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmpPresent

`func (o *ResultatValidationPDFAPI) SetXmpPresent(v bool)`

SetXmpPresent sets XmpPresent field to given value.


### GetXmpConformeFacturx

`func (o *ResultatValidationPDFAPI) GetXmpConformeFacturx() bool`

GetXmpConformeFacturx returns the XmpConformeFacturx field if non-nil, zero value otherwise.

### GetXmpConformeFacturxOk

`func (o *ResultatValidationPDFAPI) GetXmpConformeFacturxOk() (*bool, bool)`

GetXmpConformeFacturxOk returns a tuple with the XmpConformeFacturx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmpConformeFacturx

`func (o *ResultatValidationPDFAPI) SetXmpConformeFacturx(v bool)`

SetXmpConformeFacturx sets XmpConformeFacturx field to given value.


### GetProfilXmp

`func (o *ResultatValidationPDFAPI) GetProfilXmp() string`

GetProfilXmp returns the ProfilXmp field if non-nil, zero value otherwise.

### GetProfilXmpOk

`func (o *ResultatValidationPDFAPI) GetProfilXmpOk() (*string, bool)`

GetProfilXmpOk returns a tuple with the ProfilXmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilXmp

`func (o *ResultatValidationPDFAPI) SetProfilXmp(v string)`

SetProfilXmp sets ProfilXmp field to given value.

### HasProfilXmp

`func (o *ResultatValidationPDFAPI) HasProfilXmp() bool`

HasProfilXmp returns a boolean if a field has been set.

### SetProfilXmpNil

`func (o *ResultatValidationPDFAPI) SetProfilXmpNil(b bool)`

 SetProfilXmpNil sets the value for ProfilXmp to be an explicit nil

### UnsetProfilXmp
`func (o *ResultatValidationPDFAPI) UnsetProfilXmp()`

UnsetProfilXmp ensures that no value is present for ProfilXmp, not even an explicit nil
### GetVersionXmp

`func (o *ResultatValidationPDFAPI) GetVersionXmp() string`

GetVersionXmp returns the VersionXmp field if non-nil, zero value otherwise.

### GetVersionXmpOk

`func (o *ResultatValidationPDFAPI) GetVersionXmpOk() (*string, bool)`

GetVersionXmpOk returns a tuple with the VersionXmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionXmp

`func (o *ResultatValidationPDFAPI) SetVersionXmp(v string)`

SetVersionXmp sets VersionXmp field to given value.

### HasVersionXmp

`func (o *ResultatValidationPDFAPI) HasVersionXmp() bool`

HasVersionXmp returns a boolean if a field has been set.

### SetVersionXmpNil

`func (o *ResultatValidationPDFAPI) SetVersionXmpNil(b bool)`

 SetVersionXmpNil sets the value for VersionXmp to be an explicit nil

### UnsetVersionXmp
`func (o *ResultatValidationPDFAPI) UnsetVersionXmp()`

UnsetVersionXmp ensures that no value is present for VersionXmp, not even an explicit nil
### GetErreursXmp

`func (o *ResultatValidationPDFAPI) GetErreursXmp() []string`

GetErreursXmp returns the ErreursXmp field if non-nil, zero value otherwise.

### GetErreursXmpOk

`func (o *ResultatValidationPDFAPI) GetErreursXmpOk() (*[]string, bool)`

GetErreursXmpOk returns a tuple with the ErreursXmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErreursXmp

`func (o *ResultatValidationPDFAPI) SetErreursXmp(v []string)`

SetErreursXmp sets ErreursXmp field to given value.

### HasErreursXmp

`func (o *ResultatValidationPDFAPI) HasErreursXmp() bool`

HasErreursXmp returns a boolean if a field has been set.

### GetMetadonneesXmp

`func (o *ResultatValidationPDFAPI) GetMetadonneesXmp() map[string]interface{}`

GetMetadonneesXmp returns the MetadonneesXmp field if non-nil, zero value otherwise.

### GetMetadonneesXmpOk

`func (o *ResultatValidationPDFAPI) GetMetadonneesXmpOk() (*map[string]interface{}, bool)`

GetMetadonneesXmpOk returns a tuple with the MetadonneesXmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadonneesXmp

`func (o *ResultatValidationPDFAPI) SetMetadonneesXmp(v map[string]interface{})`

SetMetadonneesXmp sets MetadonneesXmp field to given value.

### HasMetadonneesXmp

`func (o *ResultatValidationPDFAPI) HasMetadonneesXmp() bool`

HasMetadonneesXmp returns a boolean if a field has been set.

### GetEstSigne

`func (o *ResultatValidationPDFAPI) GetEstSigne() bool`

GetEstSigne returns the EstSigne field if non-nil, zero value otherwise.

### GetEstSigneOk

`func (o *ResultatValidationPDFAPI) GetEstSigneOk() (*bool, bool)`

GetEstSigneOk returns a tuple with the EstSigne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstSigne

`func (o *ResultatValidationPDFAPI) SetEstSigne(v bool)`

SetEstSigne sets EstSigne field to given value.


### GetNombreSignatures

`func (o *ResultatValidationPDFAPI) GetNombreSignatures() int32`

GetNombreSignatures returns the NombreSignatures field if non-nil, zero value otherwise.

### GetNombreSignaturesOk

`func (o *ResultatValidationPDFAPI) GetNombreSignaturesOk() (*int32, bool)`

GetNombreSignaturesOk returns a tuple with the NombreSignatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNombreSignatures

`func (o *ResultatValidationPDFAPI) SetNombreSignatures(v int32)`

SetNombreSignatures sets NombreSignatures field to given value.

### HasNombreSignatures

`func (o *ResultatValidationPDFAPI) HasNombreSignatures() bool`

HasNombreSignatures returns a boolean if a field has been set.

### GetSignatures

`func (o *ResultatValidationPDFAPI) GetSignatures() []InformationSignatureAPI`

GetSignatures returns the Signatures field if non-nil, zero value otherwise.

### GetSignaturesOk

`func (o *ResultatValidationPDFAPI) GetSignaturesOk() (*[]InformationSignatureAPI, bool)`

GetSignaturesOk returns a tuple with the Signatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatures

`func (o *ResultatValidationPDFAPI) SetSignatures(v []InformationSignatureAPI)`

SetSignatures sets Signatures field to given value.

### HasSignatures

`func (o *ResultatValidationPDFAPI) HasSignatures() bool`

HasSignatures returns a boolean if a field has been set.

### GetErreursSignatures

`func (o *ResultatValidationPDFAPI) GetErreursSignatures() []string`

GetErreursSignatures returns the ErreursSignatures field if non-nil, zero value otherwise.

### GetErreursSignaturesOk

`func (o *ResultatValidationPDFAPI) GetErreursSignaturesOk() (*[]string, bool)`

GetErreursSignaturesOk returns a tuple with the ErreursSignatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErreursSignatures

`func (o *ResultatValidationPDFAPI) SetErreursSignatures(v []string)`

SetErreursSignatures sets ErreursSignatures field to given value.

### HasErreursSignatures

`func (o *ResultatValidationPDFAPI) HasErreursSignatures() bool`

HasErreursSignatures returns a boolean if a field has been set.

### GetMessageResume

`func (o *ResultatValidationPDFAPI) GetMessageResume() string`

GetMessageResume returns the MessageResume field if non-nil, zero value otherwise.

### GetMessageResumeOk

`func (o *ResultatValidationPDFAPI) GetMessageResumeOk() (*string, bool)`

GetMessageResumeOk returns a tuple with the MessageResume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageResume

`func (o *ResultatValidationPDFAPI) SetMessageResume(v string)`

SetMessageResume sets MessageResume field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



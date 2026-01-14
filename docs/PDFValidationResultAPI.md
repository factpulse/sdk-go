# PDFValidationResultAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCompliant** | **bool** | True if PDF complies with all criteria (XML, PDF/A, XMP) | 
**XmlPresent** | **bool** | True if a Factur-X XML is embedded in the PDF | 
**XmlCompliant** | **bool** | True if Factur-X XML complies with Schematron rules | 
**DetectedProfile** | Pointer to **NullableString** |  | [optional] 
**XmlErrors** | Pointer to **[]string** | List of XML validation errors | [optional] 
**PdfaCompliant** | **bool** | True if PDF is PDF/A compliant | 
**PdfaVersion** | Pointer to **NullableString** |  | [optional] 
**PdfaValidationMethod** | Pointer to **string** | Method used for PDF/A validation (metadata or verapdf) | [optional] [default to "metadata"]
**ValidatedRules** | Pointer to **NullableInt32** |  | [optional] 
**FailedRules** | Pointer to **NullableInt32** |  | [optional] 
**PdfaErrors** | Pointer to **[]string** | List of PDF/A compliance errors | [optional] 
**PdfaWarnings** | Pointer to **[]string** | List of PDF/A warnings | [optional] 
**XmpPresent** | **bool** | True if XMP metadata is present | 
**XmpFacturxCompliant** | **bool** | True if XMP metadata contains Factur-X information | 
**XmpProfile** | Pointer to **NullableString** |  | [optional] 
**XmpVersion** | Pointer to **NullableString** |  | [optional] 
**XmpErrors** | Pointer to **[]string** | List of XMP metadata errors | [optional] 
**XmpMetadata** | Pointer to **map[string]interface{}** | XMP metadata extracted from PDF | [optional] 
**IsSigned** | **bool** | True if PDF contains at least one signature | 
**SignatureCount** | Pointer to **int32** | Number of electronic signatures found | [optional] [default to 0]
**Signatures** | Pointer to [**[]SignatureInfoAPI**](SignatureInfoAPI.md) | List of found signatures with their information | [optional] 
**SignatureErrors** | Pointer to **[]string** | List of errors during signature analysis | [optional] 
**SummaryMessage** | **string** | Message summarizing the validation result | 

## Methods

### NewPDFValidationResultAPI

`func NewPDFValidationResultAPI(isCompliant bool, xmlPresent bool, xmlCompliant bool, pdfaCompliant bool, xmpPresent bool, xmpFacturxCompliant bool, isSigned bool, summaryMessage string, ) *PDFValidationResultAPI`

NewPDFValidationResultAPI instantiates a new PDFValidationResultAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPDFValidationResultAPIWithDefaults

`func NewPDFValidationResultAPIWithDefaults() *PDFValidationResultAPI`

NewPDFValidationResultAPIWithDefaults instantiates a new PDFValidationResultAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCompliant

`func (o *PDFValidationResultAPI) GetIsCompliant() bool`

GetIsCompliant returns the IsCompliant field if non-nil, zero value otherwise.

### GetIsCompliantOk

`func (o *PDFValidationResultAPI) GetIsCompliantOk() (*bool, bool)`

GetIsCompliantOk returns a tuple with the IsCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompliant

`func (o *PDFValidationResultAPI) SetIsCompliant(v bool)`

SetIsCompliant sets IsCompliant field to given value.


### GetXmlPresent

`func (o *PDFValidationResultAPI) GetXmlPresent() bool`

GetXmlPresent returns the XmlPresent field if non-nil, zero value otherwise.

### GetXmlPresentOk

`func (o *PDFValidationResultAPI) GetXmlPresentOk() (*bool, bool)`

GetXmlPresentOk returns a tuple with the XmlPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlPresent

`func (o *PDFValidationResultAPI) SetXmlPresent(v bool)`

SetXmlPresent sets XmlPresent field to given value.


### GetXmlCompliant

`func (o *PDFValidationResultAPI) GetXmlCompliant() bool`

GetXmlCompliant returns the XmlCompliant field if non-nil, zero value otherwise.

### GetXmlCompliantOk

`func (o *PDFValidationResultAPI) GetXmlCompliantOk() (*bool, bool)`

GetXmlCompliantOk returns a tuple with the XmlCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlCompliant

`func (o *PDFValidationResultAPI) SetXmlCompliant(v bool)`

SetXmlCompliant sets XmlCompliant field to given value.


### GetDetectedProfile

`func (o *PDFValidationResultAPI) GetDetectedProfile() string`

GetDetectedProfile returns the DetectedProfile field if non-nil, zero value otherwise.

### GetDetectedProfileOk

`func (o *PDFValidationResultAPI) GetDetectedProfileOk() (*string, bool)`

GetDetectedProfileOk returns a tuple with the DetectedProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedProfile

`func (o *PDFValidationResultAPI) SetDetectedProfile(v string)`

SetDetectedProfile sets DetectedProfile field to given value.

### HasDetectedProfile

`func (o *PDFValidationResultAPI) HasDetectedProfile() bool`

HasDetectedProfile returns a boolean if a field has been set.

### SetDetectedProfileNil

`func (o *PDFValidationResultAPI) SetDetectedProfileNil(b bool)`

 SetDetectedProfileNil sets the value for DetectedProfile to be an explicit nil

### UnsetDetectedProfile
`func (o *PDFValidationResultAPI) UnsetDetectedProfile()`

UnsetDetectedProfile ensures that no value is present for DetectedProfile, not even an explicit nil
### GetXmlErrors

`func (o *PDFValidationResultAPI) GetXmlErrors() []string`

GetXmlErrors returns the XmlErrors field if non-nil, zero value otherwise.

### GetXmlErrorsOk

`func (o *PDFValidationResultAPI) GetXmlErrorsOk() (*[]string, bool)`

GetXmlErrorsOk returns a tuple with the XmlErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlErrors

`func (o *PDFValidationResultAPI) SetXmlErrors(v []string)`

SetXmlErrors sets XmlErrors field to given value.

### HasXmlErrors

`func (o *PDFValidationResultAPI) HasXmlErrors() bool`

HasXmlErrors returns a boolean if a field has been set.

### GetPdfaCompliant

`func (o *PDFValidationResultAPI) GetPdfaCompliant() bool`

GetPdfaCompliant returns the PdfaCompliant field if non-nil, zero value otherwise.

### GetPdfaCompliantOk

`func (o *PDFValidationResultAPI) GetPdfaCompliantOk() (*bool, bool)`

GetPdfaCompliantOk returns a tuple with the PdfaCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfaCompliant

`func (o *PDFValidationResultAPI) SetPdfaCompliant(v bool)`

SetPdfaCompliant sets PdfaCompliant field to given value.


### GetPdfaVersion

`func (o *PDFValidationResultAPI) GetPdfaVersion() string`

GetPdfaVersion returns the PdfaVersion field if non-nil, zero value otherwise.

### GetPdfaVersionOk

`func (o *PDFValidationResultAPI) GetPdfaVersionOk() (*string, bool)`

GetPdfaVersionOk returns a tuple with the PdfaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfaVersion

`func (o *PDFValidationResultAPI) SetPdfaVersion(v string)`

SetPdfaVersion sets PdfaVersion field to given value.

### HasPdfaVersion

`func (o *PDFValidationResultAPI) HasPdfaVersion() bool`

HasPdfaVersion returns a boolean if a field has been set.

### SetPdfaVersionNil

`func (o *PDFValidationResultAPI) SetPdfaVersionNil(b bool)`

 SetPdfaVersionNil sets the value for PdfaVersion to be an explicit nil

### UnsetPdfaVersion
`func (o *PDFValidationResultAPI) UnsetPdfaVersion()`

UnsetPdfaVersion ensures that no value is present for PdfaVersion, not even an explicit nil
### GetPdfaValidationMethod

`func (o *PDFValidationResultAPI) GetPdfaValidationMethod() string`

GetPdfaValidationMethod returns the PdfaValidationMethod field if non-nil, zero value otherwise.

### GetPdfaValidationMethodOk

`func (o *PDFValidationResultAPI) GetPdfaValidationMethodOk() (*string, bool)`

GetPdfaValidationMethodOk returns a tuple with the PdfaValidationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfaValidationMethod

`func (o *PDFValidationResultAPI) SetPdfaValidationMethod(v string)`

SetPdfaValidationMethod sets PdfaValidationMethod field to given value.

### HasPdfaValidationMethod

`func (o *PDFValidationResultAPI) HasPdfaValidationMethod() bool`

HasPdfaValidationMethod returns a boolean if a field has been set.

### GetValidatedRules

`func (o *PDFValidationResultAPI) GetValidatedRules() int32`

GetValidatedRules returns the ValidatedRules field if non-nil, zero value otherwise.

### GetValidatedRulesOk

`func (o *PDFValidationResultAPI) GetValidatedRulesOk() (*int32, bool)`

GetValidatedRulesOk returns a tuple with the ValidatedRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedRules

`func (o *PDFValidationResultAPI) SetValidatedRules(v int32)`

SetValidatedRules sets ValidatedRules field to given value.

### HasValidatedRules

`func (o *PDFValidationResultAPI) HasValidatedRules() bool`

HasValidatedRules returns a boolean if a field has been set.

### SetValidatedRulesNil

`func (o *PDFValidationResultAPI) SetValidatedRulesNil(b bool)`

 SetValidatedRulesNil sets the value for ValidatedRules to be an explicit nil

### UnsetValidatedRules
`func (o *PDFValidationResultAPI) UnsetValidatedRules()`

UnsetValidatedRules ensures that no value is present for ValidatedRules, not even an explicit nil
### GetFailedRules

`func (o *PDFValidationResultAPI) GetFailedRules() int32`

GetFailedRules returns the FailedRules field if non-nil, zero value otherwise.

### GetFailedRulesOk

`func (o *PDFValidationResultAPI) GetFailedRulesOk() (*int32, bool)`

GetFailedRulesOk returns a tuple with the FailedRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRules

`func (o *PDFValidationResultAPI) SetFailedRules(v int32)`

SetFailedRules sets FailedRules field to given value.

### HasFailedRules

`func (o *PDFValidationResultAPI) HasFailedRules() bool`

HasFailedRules returns a boolean if a field has been set.

### SetFailedRulesNil

`func (o *PDFValidationResultAPI) SetFailedRulesNil(b bool)`

 SetFailedRulesNil sets the value for FailedRules to be an explicit nil

### UnsetFailedRules
`func (o *PDFValidationResultAPI) UnsetFailedRules()`

UnsetFailedRules ensures that no value is present for FailedRules, not even an explicit nil
### GetPdfaErrors

`func (o *PDFValidationResultAPI) GetPdfaErrors() []string`

GetPdfaErrors returns the PdfaErrors field if non-nil, zero value otherwise.

### GetPdfaErrorsOk

`func (o *PDFValidationResultAPI) GetPdfaErrorsOk() (*[]string, bool)`

GetPdfaErrorsOk returns a tuple with the PdfaErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfaErrors

`func (o *PDFValidationResultAPI) SetPdfaErrors(v []string)`

SetPdfaErrors sets PdfaErrors field to given value.

### HasPdfaErrors

`func (o *PDFValidationResultAPI) HasPdfaErrors() bool`

HasPdfaErrors returns a boolean if a field has been set.

### GetPdfaWarnings

`func (o *PDFValidationResultAPI) GetPdfaWarnings() []string`

GetPdfaWarnings returns the PdfaWarnings field if non-nil, zero value otherwise.

### GetPdfaWarningsOk

`func (o *PDFValidationResultAPI) GetPdfaWarningsOk() (*[]string, bool)`

GetPdfaWarningsOk returns a tuple with the PdfaWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfaWarnings

`func (o *PDFValidationResultAPI) SetPdfaWarnings(v []string)`

SetPdfaWarnings sets PdfaWarnings field to given value.

### HasPdfaWarnings

`func (o *PDFValidationResultAPI) HasPdfaWarnings() bool`

HasPdfaWarnings returns a boolean if a field has been set.

### GetXmpPresent

`func (o *PDFValidationResultAPI) GetXmpPresent() bool`

GetXmpPresent returns the XmpPresent field if non-nil, zero value otherwise.

### GetXmpPresentOk

`func (o *PDFValidationResultAPI) GetXmpPresentOk() (*bool, bool)`

GetXmpPresentOk returns a tuple with the XmpPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmpPresent

`func (o *PDFValidationResultAPI) SetXmpPresent(v bool)`

SetXmpPresent sets XmpPresent field to given value.


### GetXmpFacturxCompliant

`func (o *PDFValidationResultAPI) GetXmpFacturxCompliant() bool`

GetXmpFacturxCompliant returns the XmpFacturxCompliant field if non-nil, zero value otherwise.

### GetXmpFacturxCompliantOk

`func (o *PDFValidationResultAPI) GetXmpFacturxCompliantOk() (*bool, bool)`

GetXmpFacturxCompliantOk returns a tuple with the XmpFacturxCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmpFacturxCompliant

`func (o *PDFValidationResultAPI) SetXmpFacturxCompliant(v bool)`

SetXmpFacturxCompliant sets XmpFacturxCompliant field to given value.


### GetXmpProfile

`func (o *PDFValidationResultAPI) GetXmpProfile() string`

GetXmpProfile returns the XmpProfile field if non-nil, zero value otherwise.

### GetXmpProfileOk

`func (o *PDFValidationResultAPI) GetXmpProfileOk() (*string, bool)`

GetXmpProfileOk returns a tuple with the XmpProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmpProfile

`func (o *PDFValidationResultAPI) SetXmpProfile(v string)`

SetXmpProfile sets XmpProfile field to given value.

### HasXmpProfile

`func (o *PDFValidationResultAPI) HasXmpProfile() bool`

HasXmpProfile returns a boolean if a field has been set.

### SetXmpProfileNil

`func (o *PDFValidationResultAPI) SetXmpProfileNil(b bool)`

 SetXmpProfileNil sets the value for XmpProfile to be an explicit nil

### UnsetXmpProfile
`func (o *PDFValidationResultAPI) UnsetXmpProfile()`

UnsetXmpProfile ensures that no value is present for XmpProfile, not even an explicit nil
### GetXmpVersion

`func (o *PDFValidationResultAPI) GetXmpVersion() string`

GetXmpVersion returns the XmpVersion field if non-nil, zero value otherwise.

### GetXmpVersionOk

`func (o *PDFValidationResultAPI) GetXmpVersionOk() (*string, bool)`

GetXmpVersionOk returns a tuple with the XmpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmpVersion

`func (o *PDFValidationResultAPI) SetXmpVersion(v string)`

SetXmpVersion sets XmpVersion field to given value.

### HasXmpVersion

`func (o *PDFValidationResultAPI) HasXmpVersion() bool`

HasXmpVersion returns a boolean if a field has been set.

### SetXmpVersionNil

`func (o *PDFValidationResultAPI) SetXmpVersionNil(b bool)`

 SetXmpVersionNil sets the value for XmpVersion to be an explicit nil

### UnsetXmpVersion
`func (o *PDFValidationResultAPI) UnsetXmpVersion()`

UnsetXmpVersion ensures that no value is present for XmpVersion, not even an explicit nil
### GetXmpErrors

`func (o *PDFValidationResultAPI) GetXmpErrors() []string`

GetXmpErrors returns the XmpErrors field if non-nil, zero value otherwise.

### GetXmpErrorsOk

`func (o *PDFValidationResultAPI) GetXmpErrorsOk() (*[]string, bool)`

GetXmpErrorsOk returns a tuple with the XmpErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmpErrors

`func (o *PDFValidationResultAPI) SetXmpErrors(v []string)`

SetXmpErrors sets XmpErrors field to given value.

### HasXmpErrors

`func (o *PDFValidationResultAPI) HasXmpErrors() bool`

HasXmpErrors returns a boolean if a field has been set.

### GetXmpMetadata

`func (o *PDFValidationResultAPI) GetXmpMetadata() map[string]interface{}`

GetXmpMetadata returns the XmpMetadata field if non-nil, zero value otherwise.

### GetXmpMetadataOk

`func (o *PDFValidationResultAPI) GetXmpMetadataOk() (*map[string]interface{}, bool)`

GetXmpMetadataOk returns a tuple with the XmpMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmpMetadata

`func (o *PDFValidationResultAPI) SetXmpMetadata(v map[string]interface{})`

SetXmpMetadata sets XmpMetadata field to given value.

### HasXmpMetadata

`func (o *PDFValidationResultAPI) HasXmpMetadata() bool`

HasXmpMetadata returns a boolean if a field has been set.

### GetIsSigned

`func (o *PDFValidationResultAPI) GetIsSigned() bool`

GetIsSigned returns the IsSigned field if non-nil, zero value otherwise.

### GetIsSignedOk

`func (o *PDFValidationResultAPI) GetIsSignedOk() (*bool, bool)`

GetIsSignedOk returns a tuple with the IsSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSigned

`func (o *PDFValidationResultAPI) SetIsSigned(v bool)`

SetIsSigned sets IsSigned field to given value.


### GetSignatureCount

`func (o *PDFValidationResultAPI) GetSignatureCount() int32`

GetSignatureCount returns the SignatureCount field if non-nil, zero value otherwise.

### GetSignatureCountOk

`func (o *PDFValidationResultAPI) GetSignatureCountOk() (*int32, bool)`

GetSignatureCountOk returns a tuple with the SignatureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureCount

`func (o *PDFValidationResultAPI) SetSignatureCount(v int32)`

SetSignatureCount sets SignatureCount field to given value.

### HasSignatureCount

`func (o *PDFValidationResultAPI) HasSignatureCount() bool`

HasSignatureCount returns a boolean if a field has been set.

### GetSignatures

`func (o *PDFValidationResultAPI) GetSignatures() []SignatureInfoAPI`

GetSignatures returns the Signatures field if non-nil, zero value otherwise.

### GetSignaturesOk

`func (o *PDFValidationResultAPI) GetSignaturesOk() (*[]SignatureInfoAPI, bool)`

GetSignaturesOk returns a tuple with the Signatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatures

`func (o *PDFValidationResultAPI) SetSignatures(v []SignatureInfoAPI)`

SetSignatures sets Signatures field to given value.

### HasSignatures

`func (o *PDFValidationResultAPI) HasSignatures() bool`

HasSignatures returns a boolean if a field has been set.

### GetSignatureErrors

`func (o *PDFValidationResultAPI) GetSignatureErrors() []string`

GetSignatureErrors returns the SignatureErrors field if non-nil, zero value otherwise.

### GetSignatureErrorsOk

`func (o *PDFValidationResultAPI) GetSignatureErrorsOk() (*[]string, bool)`

GetSignatureErrorsOk returns a tuple with the SignatureErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureErrors

`func (o *PDFValidationResultAPI) SetSignatureErrors(v []string)`

SetSignatureErrors sets SignatureErrors field to given value.

### HasSignatureErrors

`func (o *PDFValidationResultAPI) HasSignatureErrors() bool`

HasSignatureErrors returns a boolean if a field has been set.

### GetSummaryMessage

`func (o *PDFValidationResultAPI) GetSummaryMessage() string`

GetSummaryMessage returns the SummaryMessage field if non-nil, zero value otherwise.

### GetSummaryMessageOk

`func (o *PDFValidationResultAPI) GetSummaryMessageOk() (*string, bool)`

GetSummaryMessageOk returns a tuple with the SummaryMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryMessage

`func (o *PDFValidationResultAPI) SetSummaryMessage(v string)`

SetSummaryMessage sets SummaryMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



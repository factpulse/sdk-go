# ValidationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | **string** | Profil Factur-X utilise | 
**SchematronRulesPassed** | **int32** | Regles passees | 
**SchematronRulesTotal** | **int32** | Total regles | 
**PdfaCompliant** | Pointer to **bool** | PDF/A-3 conforme | [optional] [default to true]
**XmlEmbedded** | Pointer to **bool** | XML embarque dans PDF | [optional] [default to true]
**Errors** | Pointer to [**[]SchematronValidationError**](SchematronValidationError.md) |  | [optional] 

## Methods

### NewValidationInfo

`func NewValidationInfo(profile string, schematronRulesPassed int32, schematronRulesTotal int32, ) *ValidationInfo`

NewValidationInfo instantiates a new ValidationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationInfoWithDefaults

`func NewValidationInfoWithDefaults() *ValidationInfo`

NewValidationInfoWithDefaults instantiates a new ValidationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *ValidationInfo) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ValidationInfo) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ValidationInfo) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetSchematronRulesPassed

`func (o *ValidationInfo) GetSchematronRulesPassed() int32`

GetSchematronRulesPassed returns the SchematronRulesPassed field if non-nil, zero value otherwise.

### GetSchematronRulesPassedOk

`func (o *ValidationInfo) GetSchematronRulesPassedOk() (*int32, bool)`

GetSchematronRulesPassedOk returns a tuple with the SchematronRulesPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchematronRulesPassed

`func (o *ValidationInfo) SetSchematronRulesPassed(v int32)`

SetSchematronRulesPassed sets SchematronRulesPassed field to given value.


### GetSchematronRulesTotal

`func (o *ValidationInfo) GetSchematronRulesTotal() int32`

GetSchematronRulesTotal returns the SchematronRulesTotal field if non-nil, zero value otherwise.

### GetSchematronRulesTotalOk

`func (o *ValidationInfo) GetSchematronRulesTotalOk() (*int32, bool)`

GetSchematronRulesTotalOk returns a tuple with the SchematronRulesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchematronRulesTotal

`func (o *ValidationInfo) SetSchematronRulesTotal(v int32)`

SetSchematronRulesTotal sets SchematronRulesTotal field to given value.


### GetPdfaCompliant

`func (o *ValidationInfo) GetPdfaCompliant() bool`

GetPdfaCompliant returns the PdfaCompliant field if non-nil, zero value otherwise.

### GetPdfaCompliantOk

`func (o *ValidationInfo) GetPdfaCompliantOk() (*bool, bool)`

GetPdfaCompliantOk returns a tuple with the PdfaCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfaCompliant

`func (o *ValidationInfo) SetPdfaCompliant(v bool)`

SetPdfaCompliant sets PdfaCompliant field to given value.

### HasPdfaCompliant

`func (o *ValidationInfo) HasPdfaCompliant() bool`

HasPdfaCompliant returns a boolean if a field has been set.

### GetXmlEmbedded

`func (o *ValidationInfo) GetXmlEmbedded() bool`

GetXmlEmbedded returns the XmlEmbedded field if non-nil, zero value otherwise.

### GetXmlEmbeddedOk

`func (o *ValidationInfo) GetXmlEmbeddedOk() (*bool, bool)`

GetXmlEmbeddedOk returns a tuple with the XmlEmbedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlEmbedded

`func (o *ValidationInfo) SetXmlEmbedded(v bool)`

SetXmlEmbedded sets XmlEmbedded field to given value.

### HasXmlEmbedded

`func (o *ValidationInfo) HasXmlEmbedded() bool`

HasXmlEmbedded returns a boolean if a field has been set.

### GetErrors

`func (o *ValidationInfo) GetErrors() []SchematronValidationError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ValidationInfo) GetErrorsOk() (*[]SchematronValidationError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ValidationInfo) SetErrors(v []SchematronValidationError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ValidationInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



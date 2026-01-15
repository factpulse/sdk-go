# ProcessingOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FacturxProfile** | Pointer to [**APIProfile**](APIProfile.md) | Factur-X profile to use | [optional] [default to EN16931]
**AutoEnrich** | Pointer to **bool** | Auto-enrich data (Company APIs, Chorus Pro, etc.) | [optional] [default to true]
**ValidateXml** | Pointer to **bool** | Validate Factur-X XML with Schematron | [optional] [default to true]
**VerifyDestinationParameters** | Pointer to **bool** | Verify required parameters for destination (e.g., service_code for Chorus) | [optional] [default to true]

## Methods

### NewProcessingOptions

`func NewProcessingOptions() *ProcessingOptions`

NewProcessingOptions instantiates a new ProcessingOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessingOptionsWithDefaults

`func NewProcessingOptionsWithDefaults() *ProcessingOptions`

NewProcessingOptionsWithDefaults instantiates a new ProcessingOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacturxProfile

`func (o *ProcessingOptions) GetFacturxProfile() APIProfile`

GetFacturxProfile returns the FacturxProfile field if non-nil, zero value otherwise.

### GetFacturxProfileOk

`func (o *ProcessingOptions) GetFacturxProfileOk() (*APIProfile, bool)`

GetFacturxProfileOk returns a tuple with the FacturxProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacturxProfile

`func (o *ProcessingOptions) SetFacturxProfile(v APIProfile)`

SetFacturxProfile sets FacturxProfile field to given value.

### HasFacturxProfile

`func (o *ProcessingOptions) HasFacturxProfile() bool`

HasFacturxProfile returns a boolean if a field has been set.

### GetAutoEnrich

`func (o *ProcessingOptions) GetAutoEnrich() bool`

GetAutoEnrich returns the AutoEnrich field if non-nil, zero value otherwise.

### GetAutoEnrichOk

`func (o *ProcessingOptions) GetAutoEnrichOk() (*bool, bool)`

GetAutoEnrichOk returns a tuple with the AutoEnrich field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoEnrich

`func (o *ProcessingOptions) SetAutoEnrich(v bool)`

SetAutoEnrich sets AutoEnrich field to given value.

### HasAutoEnrich

`func (o *ProcessingOptions) HasAutoEnrich() bool`

HasAutoEnrich returns a boolean if a field has been set.

### GetValidateXml

`func (o *ProcessingOptions) GetValidateXml() bool`

GetValidateXml returns the ValidateXml field if non-nil, zero value otherwise.

### GetValidateXmlOk

`func (o *ProcessingOptions) GetValidateXmlOk() (*bool, bool)`

GetValidateXmlOk returns a tuple with the ValidateXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateXml

`func (o *ProcessingOptions) SetValidateXml(v bool)`

SetValidateXml sets ValidateXml field to given value.

### HasValidateXml

`func (o *ProcessingOptions) HasValidateXml() bool`

HasValidateXml returns a boolean if a field has been set.

### GetVerifyDestinationParameters

`func (o *ProcessingOptions) GetVerifyDestinationParameters() bool`

GetVerifyDestinationParameters returns the VerifyDestinationParameters field if non-nil, zero value otherwise.

### GetVerifyDestinationParametersOk

`func (o *ProcessingOptions) GetVerifyDestinationParametersOk() (*bool, bool)`

GetVerifyDestinationParametersOk returns a tuple with the VerifyDestinationParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyDestinationParameters

`func (o *ProcessingOptions) SetVerifyDestinationParameters(v bool)`

SetVerifyDestinationParameters sets VerifyDestinationParameters field to given value.

### HasVerifyDestinationParameters

`func (o *ProcessingOptions) HasVerifyDestinationParameters() bool`

HasVerifyDestinationParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



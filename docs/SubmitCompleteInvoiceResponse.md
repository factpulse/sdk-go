# SubmitCompleteInvoiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Invoice was successfully submitted | 
**DestinationType** | **string** | Destination type | 
**ChorusResult** | Pointer to [**NullableChorusProResult**](ChorusProResult.md) |  | [optional] 
**AfnorResult** | Pointer to [**NullableAFNORResult**](AFNORResult.md) |  | [optional] 
**EnrichedInvoice** | [**EnrichedInvoiceInfo**](EnrichedInvoiceInfo.md) | Enriched invoice data | 
**FacturxPdf** | [**FacturXPDFInfo**](FacturXPDFInfo.md) | Generated PDF information | 
**Signature** | Pointer to [**NullableSignatureInfo**](SignatureInfo.md) |  | [optional] 
**PdfBase64** | **string** | Generated Factur-X PDF (and signed if requested) base64-encoded | 
**Message** | **string** | Return message | 

## Methods

### NewSubmitCompleteInvoiceResponse

`func NewSubmitCompleteInvoiceResponse(success bool, destinationType string, enrichedInvoice EnrichedInvoiceInfo, facturxPdf FacturXPDFInfo, pdfBase64 string, message string, ) *SubmitCompleteInvoiceResponse`

NewSubmitCompleteInvoiceResponse instantiates a new SubmitCompleteInvoiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitCompleteInvoiceResponseWithDefaults

`func NewSubmitCompleteInvoiceResponseWithDefaults() *SubmitCompleteInvoiceResponse`

NewSubmitCompleteInvoiceResponseWithDefaults instantiates a new SubmitCompleteInvoiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *SubmitCompleteInvoiceResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SubmitCompleteInvoiceResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SubmitCompleteInvoiceResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetDestinationType

`func (o *SubmitCompleteInvoiceResponse) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *SubmitCompleteInvoiceResponse) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *SubmitCompleteInvoiceResponse) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.


### GetChorusResult

`func (o *SubmitCompleteInvoiceResponse) GetChorusResult() ChorusProResult`

GetChorusResult returns the ChorusResult field if non-nil, zero value otherwise.

### GetChorusResultOk

`func (o *SubmitCompleteInvoiceResponse) GetChorusResultOk() (*ChorusProResult, bool)`

GetChorusResultOk returns a tuple with the ChorusResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChorusResult

`func (o *SubmitCompleteInvoiceResponse) SetChorusResult(v ChorusProResult)`

SetChorusResult sets ChorusResult field to given value.

### HasChorusResult

`func (o *SubmitCompleteInvoiceResponse) HasChorusResult() bool`

HasChorusResult returns a boolean if a field has been set.

### SetChorusResultNil

`func (o *SubmitCompleteInvoiceResponse) SetChorusResultNil(b bool)`

 SetChorusResultNil sets the value for ChorusResult to be an explicit nil

### UnsetChorusResult
`func (o *SubmitCompleteInvoiceResponse) UnsetChorusResult()`

UnsetChorusResult ensures that no value is present for ChorusResult, not even an explicit nil
### GetAfnorResult

`func (o *SubmitCompleteInvoiceResponse) GetAfnorResult() AFNORResult`

GetAfnorResult returns the AfnorResult field if non-nil, zero value otherwise.

### GetAfnorResultOk

`func (o *SubmitCompleteInvoiceResponse) GetAfnorResultOk() (*AFNORResult, bool)`

GetAfnorResultOk returns a tuple with the AfnorResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfnorResult

`func (o *SubmitCompleteInvoiceResponse) SetAfnorResult(v AFNORResult)`

SetAfnorResult sets AfnorResult field to given value.

### HasAfnorResult

`func (o *SubmitCompleteInvoiceResponse) HasAfnorResult() bool`

HasAfnorResult returns a boolean if a field has been set.

### SetAfnorResultNil

`func (o *SubmitCompleteInvoiceResponse) SetAfnorResultNil(b bool)`

 SetAfnorResultNil sets the value for AfnorResult to be an explicit nil

### UnsetAfnorResult
`func (o *SubmitCompleteInvoiceResponse) UnsetAfnorResult()`

UnsetAfnorResult ensures that no value is present for AfnorResult, not even an explicit nil
### GetEnrichedInvoice

`func (o *SubmitCompleteInvoiceResponse) GetEnrichedInvoice() EnrichedInvoiceInfo`

GetEnrichedInvoice returns the EnrichedInvoice field if non-nil, zero value otherwise.

### GetEnrichedInvoiceOk

`func (o *SubmitCompleteInvoiceResponse) GetEnrichedInvoiceOk() (*EnrichedInvoiceInfo, bool)`

GetEnrichedInvoiceOk returns a tuple with the EnrichedInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrichedInvoice

`func (o *SubmitCompleteInvoiceResponse) SetEnrichedInvoice(v EnrichedInvoiceInfo)`

SetEnrichedInvoice sets EnrichedInvoice field to given value.


### GetFacturxPdf

`func (o *SubmitCompleteInvoiceResponse) GetFacturxPdf() FacturXPDFInfo`

GetFacturxPdf returns the FacturxPdf field if non-nil, zero value otherwise.

### GetFacturxPdfOk

`func (o *SubmitCompleteInvoiceResponse) GetFacturxPdfOk() (*FacturXPDFInfo, bool)`

GetFacturxPdfOk returns a tuple with the FacturxPdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacturxPdf

`func (o *SubmitCompleteInvoiceResponse) SetFacturxPdf(v FacturXPDFInfo)`

SetFacturxPdf sets FacturxPdf field to given value.


### GetSignature

`func (o *SubmitCompleteInvoiceResponse) GetSignature() SignatureInfo`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *SubmitCompleteInvoiceResponse) GetSignatureOk() (*SignatureInfo, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *SubmitCompleteInvoiceResponse) SetSignature(v SignatureInfo)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *SubmitCompleteInvoiceResponse) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### SetSignatureNil

`func (o *SubmitCompleteInvoiceResponse) SetSignatureNil(b bool)`

 SetSignatureNil sets the value for Signature to be an explicit nil

### UnsetSignature
`func (o *SubmitCompleteInvoiceResponse) UnsetSignature()`

UnsetSignature ensures that no value is present for Signature, not even an explicit nil
### GetPdfBase64

`func (o *SubmitCompleteInvoiceResponse) GetPdfBase64() string`

GetPdfBase64 returns the PdfBase64 field if non-nil, zero value otherwise.

### GetPdfBase64Ok

`func (o *SubmitCompleteInvoiceResponse) GetPdfBase64Ok() (*string, bool)`

GetPdfBase64Ok returns a tuple with the PdfBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfBase64

`func (o *SubmitCompleteInvoiceResponse) SetPdfBase64(v string)`

SetPdfBase64 sets PdfBase64 field to given value.


### GetMessage

`func (o *SubmitCompleteInvoiceResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SubmitCompleteInvoiceResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SubmitCompleteInvoiceResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



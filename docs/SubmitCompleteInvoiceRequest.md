# SubmitCompleteInvoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceData** | [**SimplifiedInvoiceData**](SimplifiedInvoiceData.md) | Invoice data in simplified format (see examples) | 
**SourcePdf** | **string** | Base64-encoded source PDF (will be transformed to Factur-X) | 
**Destination** | [**Destination**](Destination.md) |  | 
**Signature** | Pointer to [**NullableSignatureParameters**](SignatureParameters.md) |  | [optional] 
**Options** | Pointer to [**ProcessingOptions**](ProcessingOptions.md) | Processing options | [optional] 

## Methods

### NewSubmitCompleteInvoiceRequest

`func NewSubmitCompleteInvoiceRequest(invoiceData SimplifiedInvoiceData, sourcePdf string, destination Destination, ) *SubmitCompleteInvoiceRequest`

NewSubmitCompleteInvoiceRequest instantiates a new SubmitCompleteInvoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitCompleteInvoiceRequestWithDefaults

`func NewSubmitCompleteInvoiceRequestWithDefaults() *SubmitCompleteInvoiceRequest`

NewSubmitCompleteInvoiceRequestWithDefaults instantiates a new SubmitCompleteInvoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceData

`func (o *SubmitCompleteInvoiceRequest) GetInvoiceData() SimplifiedInvoiceData`

GetInvoiceData returns the InvoiceData field if non-nil, zero value otherwise.

### GetInvoiceDataOk

`func (o *SubmitCompleteInvoiceRequest) GetInvoiceDataOk() (*SimplifiedInvoiceData, bool)`

GetInvoiceDataOk returns a tuple with the InvoiceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceData

`func (o *SubmitCompleteInvoiceRequest) SetInvoiceData(v SimplifiedInvoiceData)`

SetInvoiceData sets InvoiceData field to given value.


### GetSourcePdf

`func (o *SubmitCompleteInvoiceRequest) GetSourcePdf() string`

GetSourcePdf returns the SourcePdf field if non-nil, zero value otherwise.

### GetSourcePdfOk

`func (o *SubmitCompleteInvoiceRequest) GetSourcePdfOk() (*string, bool)`

GetSourcePdfOk returns a tuple with the SourcePdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePdf

`func (o *SubmitCompleteInvoiceRequest) SetSourcePdf(v string)`

SetSourcePdf sets SourcePdf field to given value.


### GetDestination

`func (o *SubmitCompleteInvoiceRequest) GetDestination() Destination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *SubmitCompleteInvoiceRequest) GetDestinationOk() (*Destination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *SubmitCompleteInvoiceRequest) SetDestination(v Destination)`

SetDestination sets Destination field to given value.


### GetSignature

`func (o *SubmitCompleteInvoiceRequest) GetSignature() SignatureParameters`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *SubmitCompleteInvoiceRequest) GetSignatureOk() (*SignatureParameters, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *SubmitCompleteInvoiceRequest) SetSignature(v SignatureParameters)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *SubmitCompleteInvoiceRequest) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### SetSignatureNil

`func (o *SubmitCompleteInvoiceRequest) SetSignatureNil(b bool)`

 SetSignatureNil sets the value for Signature to be an explicit nil

### UnsetSignature
`func (o *SubmitCompleteInvoiceRequest) UnsetSignature()`

UnsetSignature ensures that no value is present for Signature, not even an explicit nil
### GetOptions

`func (o *SubmitCompleteInvoiceRequest) GetOptions() ProcessingOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SubmitCompleteInvoiceRequest) GetOptionsOk() (*ProcessingOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SubmitCompleteInvoiceRequest) SetOptions(v ProcessingOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SubmitCompleteInvoiceRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



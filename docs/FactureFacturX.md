# FactureFacturX

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceNumber** | **string** |  | 
**PaymentDueDate** | **string** |  | 
**InvoiceDate** | Pointer to **string** |  | [optional] 
**SubmissionMode** | [**SubmissionMode**](SubmissionMode.md) |  | 
**Recipient** | [**Recipient**](Recipient.md) |  | 
**Supplier** | [**Supplier**](Supplier.md) |  | 
**InvoicingFramework** | [**InvoicingFramework**](InvoicingFramework.md) |  | 
**References** | [**InvoiceReferences**](InvoiceReferences.md) |  | 
**Totals** | [**InvoiceTotals**](InvoiceTotals.md) |  | 
**InvoiceLines** | Pointer to [**[]InvoiceLine**](InvoiceLine.md) |  | [optional] 
**VatLines** | Pointer to [**[]VATLine**](VATLine.md) |  | [optional] 
**Notes** | Pointer to [**[]InvoiceNote**](InvoiceNote.md) |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**CurrentUserId** | Pointer to **NullableInt32** |  | [optional] 
**SupplementaryAttachments** | Pointer to [**[]SupplementaryAttachment**](SupplementaryAttachment.md) |  | [optional] 
**Payee** | Pointer to [**NullablePayee**](Payee.md) |  | [optional] 

## Methods

### NewFactureFacturX

`func NewFactureFacturX(invoiceNumber string, paymentDueDate string, submissionMode SubmissionMode, recipient Recipient, supplier Supplier, invoicingFramework InvoicingFramework, references InvoiceReferences, totals InvoiceTotals, ) *FactureFacturX`

NewFactureFacturX instantiates a new FactureFacturX object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFactureFacturXWithDefaults

`func NewFactureFacturXWithDefaults() *FactureFacturX`

NewFactureFacturXWithDefaults instantiates a new FactureFacturX object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceNumber

`func (o *FactureFacturX) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *FactureFacturX) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *FactureFacturX) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.


### GetPaymentDueDate

`func (o *FactureFacturX) GetPaymentDueDate() string`

GetPaymentDueDate returns the PaymentDueDate field if non-nil, zero value otherwise.

### GetPaymentDueDateOk

`func (o *FactureFacturX) GetPaymentDueDateOk() (*string, bool)`

GetPaymentDueDateOk returns a tuple with the PaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDueDate

`func (o *FactureFacturX) SetPaymentDueDate(v string)`

SetPaymentDueDate sets PaymentDueDate field to given value.


### GetInvoiceDate

`func (o *FactureFacturX) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *FactureFacturX) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *FactureFacturX) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *FactureFacturX) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetSubmissionMode

`func (o *FactureFacturX) GetSubmissionMode() SubmissionMode`

GetSubmissionMode returns the SubmissionMode field if non-nil, zero value otherwise.

### GetSubmissionModeOk

`func (o *FactureFacturX) GetSubmissionModeOk() (*SubmissionMode, bool)`

GetSubmissionModeOk returns a tuple with the SubmissionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionMode

`func (o *FactureFacturX) SetSubmissionMode(v SubmissionMode)`

SetSubmissionMode sets SubmissionMode field to given value.


### GetRecipient

`func (o *FactureFacturX) GetRecipient() Recipient`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *FactureFacturX) GetRecipientOk() (*Recipient, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *FactureFacturX) SetRecipient(v Recipient)`

SetRecipient sets Recipient field to given value.


### GetSupplier

`func (o *FactureFacturX) GetSupplier() Supplier`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *FactureFacturX) GetSupplierOk() (*Supplier, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *FactureFacturX) SetSupplier(v Supplier)`

SetSupplier sets Supplier field to given value.


### GetInvoicingFramework

`func (o *FactureFacturX) GetInvoicingFramework() InvoicingFramework`

GetInvoicingFramework returns the InvoicingFramework field if non-nil, zero value otherwise.

### GetInvoicingFrameworkOk

`func (o *FactureFacturX) GetInvoicingFrameworkOk() (*InvoicingFramework, bool)`

GetInvoicingFrameworkOk returns a tuple with the InvoicingFramework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicingFramework

`func (o *FactureFacturX) SetInvoicingFramework(v InvoicingFramework)`

SetInvoicingFramework sets InvoicingFramework field to given value.


### GetReferences

`func (o *FactureFacturX) GetReferences() InvoiceReferences`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *FactureFacturX) GetReferencesOk() (*InvoiceReferences, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *FactureFacturX) SetReferences(v InvoiceReferences)`

SetReferences sets References field to given value.


### GetTotals

`func (o *FactureFacturX) GetTotals() InvoiceTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *FactureFacturX) GetTotalsOk() (*InvoiceTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *FactureFacturX) SetTotals(v InvoiceTotals)`

SetTotals sets Totals field to given value.


### GetInvoiceLines

`func (o *FactureFacturX) GetInvoiceLines() []InvoiceLine`

GetInvoiceLines returns the InvoiceLines field if non-nil, zero value otherwise.

### GetInvoiceLinesOk

`func (o *FactureFacturX) GetInvoiceLinesOk() (*[]InvoiceLine, bool)`

GetInvoiceLinesOk returns a tuple with the InvoiceLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceLines

`func (o *FactureFacturX) SetInvoiceLines(v []InvoiceLine)`

SetInvoiceLines sets InvoiceLines field to given value.

### HasInvoiceLines

`func (o *FactureFacturX) HasInvoiceLines() bool`

HasInvoiceLines returns a boolean if a field has been set.

### GetVatLines

`func (o *FactureFacturX) GetVatLines() []VATLine`

GetVatLines returns the VatLines field if non-nil, zero value otherwise.

### GetVatLinesOk

`func (o *FactureFacturX) GetVatLinesOk() (*[]VATLine, bool)`

GetVatLinesOk returns a tuple with the VatLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatLines

`func (o *FactureFacturX) SetVatLines(v []VATLine)`

SetVatLines sets VatLines field to given value.

### HasVatLines

`func (o *FactureFacturX) HasVatLines() bool`

HasVatLines returns a boolean if a field has been set.

### GetNotes

`func (o *FactureFacturX) GetNotes() []InvoiceNote`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *FactureFacturX) GetNotesOk() (*[]InvoiceNote, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *FactureFacturX) SetNotes(v []InvoiceNote)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *FactureFacturX) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetComment

`func (o *FactureFacturX) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FactureFacturX) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FactureFacturX) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FactureFacturX) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *FactureFacturX) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *FactureFacturX) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetCurrentUserId

`func (o *FactureFacturX) GetCurrentUserId() int32`

GetCurrentUserId returns the CurrentUserId field if non-nil, zero value otherwise.

### GetCurrentUserIdOk

`func (o *FactureFacturX) GetCurrentUserIdOk() (*int32, bool)`

GetCurrentUserIdOk returns a tuple with the CurrentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUserId

`func (o *FactureFacturX) SetCurrentUserId(v int32)`

SetCurrentUserId sets CurrentUserId field to given value.

### HasCurrentUserId

`func (o *FactureFacturX) HasCurrentUserId() bool`

HasCurrentUserId returns a boolean if a field has been set.

### SetCurrentUserIdNil

`func (o *FactureFacturX) SetCurrentUserIdNil(b bool)`

 SetCurrentUserIdNil sets the value for CurrentUserId to be an explicit nil

### UnsetCurrentUserId
`func (o *FactureFacturX) UnsetCurrentUserId()`

UnsetCurrentUserId ensures that no value is present for CurrentUserId, not even an explicit nil
### GetSupplementaryAttachments

`func (o *FactureFacturX) GetSupplementaryAttachments() []SupplementaryAttachment`

GetSupplementaryAttachments returns the SupplementaryAttachments field if non-nil, zero value otherwise.

### GetSupplementaryAttachmentsOk

`func (o *FactureFacturX) GetSupplementaryAttachmentsOk() (*[]SupplementaryAttachment, bool)`

GetSupplementaryAttachmentsOk returns a tuple with the SupplementaryAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementaryAttachments

`func (o *FactureFacturX) SetSupplementaryAttachments(v []SupplementaryAttachment)`

SetSupplementaryAttachments sets SupplementaryAttachments field to given value.

### HasSupplementaryAttachments

`func (o *FactureFacturX) HasSupplementaryAttachments() bool`

HasSupplementaryAttachments returns a boolean if a field has been set.

### SetSupplementaryAttachmentsNil

`func (o *FactureFacturX) SetSupplementaryAttachmentsNil(b bool)`

 SetSupplementaryAttachmentsNil sets the value for SupplementaryAttachments to be an explicit nil

### UnsetSupplementaryAttachments
`func (o *FactureFacturX) UnsetSupplementaryAttachments()`

UnsetSupplementaryAttachments ensures that no value is present for SupplementaryAttachments, not even an explicit nil
### GetPayee

`func (o *FactureFacturX) GetPayee() Payee`

GetPayee returns the Payee field if non-nil, zero value otherwise.

### GetPayeeOk

`func (o *FactureFacturX) GetPayeeOk() (*Payee, bool)`

GetPayeeOk returns a tuple with the Payee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayee

`func (o *FactureFacturX) SetPayee(v Payee)`

SetPayee sets Payee field to given value.

### HasPayee

`func (o *FactureFacturX) HasPayee() bool`

HasPayee returns a boolean if a field has been set.

### SetPayeeNil

`func (o *FactureFacturX) SetPayeeNil(b bool)`

 SetPayeeNil sets the value for Payee to be an explicit nil

### UnsetPayee
`func (o *FactureFacturX) UnsetPayee()`

UnsetPayee ensures that no value is present for Payee, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



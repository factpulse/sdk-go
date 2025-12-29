# FacturXInvoice

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
**DeliveryParty** | Pointer to [**NullableDeliveryParty**](DeliveryParty.md) |  | [optional] 
**TaxRepresentative** | Pointer to [**NullableTaxRepresentative**](TaxRepresentative.md) |  | [optional] 
**DeliveryDate** | Pointer to **NullableString** |  | [optional] 
**BillingPeriodStart** | Pointer to **NullableString** |  | [optional] 
**BillingPeriodEnd** | Pointer to **NullableString** |  | [optional] 
**PaymentReference** | Pointer to **NullableString** |  | [optional] 
**CreditorReferenceId** | Pointer to **NullableString** |  | [optional] 
**DirectDebitMandateId** | Pointer to **NullableString** |  | [optional] 
**DebtorIban** | Pointer to **NullableString** |  | [optional] 
**PaymentTerms** | Pointer to **NullableString** |  | [optional] 
**AllowancesCharges** | Pointer to [**[]AllowanceCharge**](AllowanceCharge.md) |  | [optional] 
**AdditionalDocuments** | Pointer to [**[]AdditionalDocument**](AdditionalDocument.md) |  | [optional] 
**BuyerAccountingReference** | Pointer to **NullableString** |  | [optional] 
**PaymentCard** | Pointer to [**NullablePaymentCard**](PaymentCard.md) |  | [optional] 

## Methods

### NewFacturXInvoice

`func NewFacturXInvoice(invoiceNumber string, paymentDueDate string, submissionMode SubmissionMode, recipient Recipient, supplier Supplier, invoicingFramework InvoicingFramework, references InvoiceReferences, totals InvoiceTotals, ) *FacturXInvoice`

NewFacturXInvoice instantiates a new FacturXInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacturXInvoiceWithDefaults

`func NewFacturXInvoiceWithDefaults() *FacturXInvoice`

NewFacturXInvoiceWithDefaults instantiates a new FacturXInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceNumber

`func (o *FacturXInvoice) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *FacturXInvoice) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *FacturXInvoice) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.


### GetPaymentDueDate

`func (o *FacturXInvoice) GetPaymentDueDate() string`

GetPaymentDueDate returns the PaymentDueDate field if non-nil, zero value otherwise.

### GetPaymentDueDateOk

`func (o *FacturXInvoice) GetPaymentDueDateOk() (*string, bool)`

GetPaymentDueDateOk returns a tuple with the PaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDueDate

`func (o *FacturXInvoice) SetPaymentDueDate(v string)`

SetPaymentDueDate sets PaymentDueDate field to given value.


### GetInvoiceDate

`func (o *FacturXInvoice) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *FacturXInvoice) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *FacturXInvoice) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *FacturXInvoice) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetSubmissionMode

`func (o *FacturXInvoice) GetSubmissionMode() SubmissionMode`

GetSubmissionMode returns the SubmissionMode field if non-nil, zero value otherwise.

### GetSubmissionModeOk

`func (o *FacturXInvoice) GetSubmissionModeOk() (*SubmissionMode, bool)`

GetSubmissionModeOk returns a tuple with the SubmissionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionMode

`func (o *FacturXInvoice) SetSubmissionMode(v SubmissionMode)`

SetSubmissionMode sets SubmissionMode field to given value.


### GetRecipient

`func (o *FacturXInvoice) GetRecipient() Recipient`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *FacturXInvoice) GetRecipientOk() (*Recipient, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *FacturXInvoice) SetRecipient(v Recipient)`

SetRecipient sets Recipient field to given value.


### GetSupplier

`func (o *FacturXInvoice) GetSupplier() Supplier`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *FacturXInvoice) GetSupplierOk() (*Supplier, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *FacturXInvoice) SetSupplier(v Supplier)`

SetSupplier sets Supplier field to given value.


### GetInvoicingFramework

`func (o *FacturXInvoice) GetInvoicingFramework() InvoicingFramework`

GetInvoicingFramework returns the InvoicingFramework field if non-nil, zero value otherwise.

### GetInvoicingFrameworkOk

`func (o *FacturXInvoice) GetInvoicingFrameworkOk() (*InvoicingFramework, bool)`

GetInvoicingFrameworkOk returns a tuple with the InvoicingFramework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicingFramework

`func (o *FacturXInvoice) SetInvoicingFramework(v InvoicingFramework)`

SetInvoicingFramework sets InvoicingFramework field to given value.


### GetReferences

`func (o *FacturXInvoice) GetReferences() InvoiceReferences`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *FacturXInvoice) GetReferencesOk() (*InvoiceReferences, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *FacturXInvoice) SetReferences(v InvoiceReferences)`

SetReferences sets References field to given value.


### GetTotals

`func (o *FacturXInvoice) GetTotals() InvoiceTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *FacturXInvoice) GetTotalsOk() (*InvoiceTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *FacturXInvoice) SetTotals(v InvoiceTotals)`

SetTotals sets Totals field to given value.


### GetInvoiceLines

`func (o *FacturXInvoice) GetInvoiceLines() []InvoiceLine`

GetInvoiceLines returns the InvoiceLines field if non-nil, zero value otherwise.

### GetInvoiceLinesOk

`func (o *FacturXInvoice) GetInvoiceLinesOk() (*[]InvoiceLine, bool)`

GetInvoiceLinesOk returns a tuple with the InvoiceLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceLines

`func (o *FacturXInvoice) SetInvoiceLines(v []InvoiceLine)`

SetInvoiceLines sets InvoiceLines field to given value.

### HasInvoiceLines

`func (o *FacturXInvoice) HasInvoiceLines() bool`

HasInvoiceLines returns a boolean if a field has been set.

### GetVatLines

`func (o *FacturXInvoice) GetVatLines() []VATLine`

GetVatLines returns the VatLines field if non-nil, zero value otherwise.

### GetVatLinesOk

`func (o *FacturXInvoice) GetVatLinesOk() (*[]VATLine, bool)`

GetVatLinesOk returns a tuple with the VatLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatLines

`func (o *FacturXInvoice) SetVatLines(v []VATLine)`

SetVatLines sets VatLines field to given value.

### HasVatLines

`func (o *FacturXInvoice) HasVatLines() bool`

HasVatLines returns a boolean if a field has been set.

### GetNotes

`func (o *FacturXInvoice) GetNotes() []InvoiceNote`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *FacturXInvoice) GetNotesOk() (*[]InvoiceNote, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *FacturXInvoice) SetNotes(v []InvoiceNote)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *FacturXInvoice) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetComment

`func (o *FacturXInvoice) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FacturXInvoice) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FacturXInvoice) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FacturXInvoice) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *FacturXInvoice) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *FacturXInvoice) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetCurrentUserId

`func (o *FacturXInvoice) GetCurrentUserId() int32`

GetCurrentUserId returns the CurrentUserId field if non-nil, zero value otherwise.

### GetCurrentUserIdOk

`func (o *FacturXInvoice) GetCurrentUserIdOk() (*int32, bool)`

GetCurrentUserIdOk returns a tuple with the CurrentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUserId

`func (o *FacturXInvoice) SetCurrentUserId(v int32)`

SetCurrentUserId sets CurrentUserId field to given value.

### HasCurrentUserId

`func (o *FacturXInvoice) HasCurrentUserId() bool`

HasCurrentUserId returns a boolean if a field has been set.

### SetCurrentUserIdNil

`func (o *FacturXInvoice) SetCurrentUserIdNil(b bool)`

 SetCurrentUserIdNil sets the value for CurrentUserId to be an explicit nil

### UnsetCurrentUserId
`func (o *FacturXInvoice) UnsetCurrentUserId()`

UnsetCurrentUserId ensures that no value is present for CurrentUserId, not even an explicit nil
### GetSupplementaryAttachments

`func (o *FacturXInvoice) GetSupplementaryAttachments() []SupplementaryAttachment`

GetSupplementaryAttachments returns the SupplementaryAttachments field if non-nil, zero value otherwise.

### GetSupplementaryAttachmentsOk

`func (o *FacturXInvoice) GetSupplementaryAttachmentsOk() (*[]SupplementaryAttachment, bool)`

GetSupplementaryAttachmentsOk returns a tuple with the SupplementaryAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementaryAttachments

`func (o *FacturXInvoice) SetSupplementaryAttachments(v []SupplementaryAttachment)`

SetSupplementaryAttachments sets SupplementaryAttachments field to given value.

### HasSupplementaryAttachments

`func (o *FacturXInvoice) HasSupplementaryAttachments() bool`

HasSupplementaryAttachments returns a boolean if a field has been set.

### SetSupplementaryAttachmentsNil

`func (o *FacturXInvoice) SetSupplementaryAttachmentsNil(b bool)`

 SetSupplementaryAttachmentsNil sets the value for SupplementaryAttachments to be an explicit nil

### UnsetSupplementaryAttachments
`func (o *FacturXInvoice) UnsetSupplementaryAttachments()`

UnsetSupplementaryAttachments ensures that no value is present for SupplementaryAttachments, not even an explicit nil
### GetPayee

`func (o *FacturXInvoice) GetPayee() Payee`

GetPayee returns the Payee field if non-nil, zero value otherwise.

### GetPayeeOk

`func (o *FacturXInvoice) GetPayeeOk() (*Payee, bool)`

GetPayeeOk returns a tuple with the Payee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayee

`func (o *FacturXInvoice) SetPayee(v Payee)`

SetPayee sets Payee field to given value.

### HasPayee

`func (o *FacturXInvoice) HasPayee() bool`

HasPayee returns a boolean if a field has been set.

### SetPayeeNil

`func (o *FacturXInvoice) SetPayeeNil(b bool)`

 SetPayeeNil sets the value for Payee to be an explicit nil

### UnsetPayee
`func (o *FacturXInvoice) UnsetPayee()`

UnsetPayee ensures that no value is present for Payee, not even an explicit nil
### GetDeliveryParty

`func (o *FacturXInvoice) GetDeliveryParty() DeliveryParty`

GetDeliveryParty returns the DeliveryParty field if non-nil, zero value otherwise.

### GetDeliveryPartyOk

`func (o *FacturXInvoice) GetDeliveryPartyOk() (*DeliveryParty, bool)`

GetDeliveryPartyOk returns a tuple with the DeliveryParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryParty

`func (o *FacturXInvoice) SetDeliveryParty(v DeliveryParty)`

SetDeliveryParty sets DeliveryParty field to given value.

### HasDeliveryParty

`func (o *FacturXInvoice) HasDeliveryParty() bool`

HasDeliveryParty returns a boolean if a field has been set.

### SetDeliveryPartyNil

`func (o *FacturXInvoice) SetDeliveryPartyNil(b bool)`

 SetDeliveryPartyNil sets the value for DeliveryParty to be an explicit nil

### UnsetDeliveryParty
`func (o *FacturXInvoice) UnsetDeliveryParty()`

UnsetDeliveryParty ensures that no value is present for DeliveryParty, not even an explicit nil
### GetTaxRepresentative

`func (o *FacturXInvoice) GetTaxRepresentative() TaxRepresentative`

GetTaxRepresentative returns the TaxRepresentative field if non-nil, zero value otherwise.

### GetTaxRepresentativeOk

`func (o *FacturXInvoice) GetTaxRepresentativeOk() (*TaxRepresentative, bool)`

GetTaxRepresentativeOk returns a tuple with the TaxRepresentative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRepresentative

`func (o *FacturXInvoice) SetTaxRepresentative(v TaxRepresentative)`

SetTaxRepresentative sets TaxRepresentative field to given value.

### HasTaxRepresentative

`func (o *FacturXInvoice) HasTaxRepresentative() bool`

HasTaxRepresentative returns a boolean if a field has been set.

### SetTaxRepresentativeNil

`func (o *FacturXInvoice) SetTaxRepresentativeNil(b bool)`

 SetTaxRepresentativeNil sets the value for TaxRepresentative to be an explicit nil

### UnsetTaxRepresentative
`func (o *FacturXInvoice) UnsetTaxRepresentative()`

UnsetTaxRepresentative ensures that no value is present for TaxRepresentative, not even an explicit nil
### GetDeliveryDate

`func (o *FacturXInvoice) GetDeliveryDate() string`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *FacturXInvoice) GetDeliveryDateOk() (*string, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *FacturXInvoice) SetDeliveryDate(v string)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *FacturXInvoice) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### SetDeliveryDateNil

`func (o *FacturXInvoice) SetDeliveryDateNil(b bool)`

 SetDeliveryDateNil sets the value for DeliveryDate to be an explicit nil

### UnsetDeliveryDate
`func (o *FacturXInvoice) UnsetDeliveryDate()`

UnsetDeliveryDate ensures that no value is present for DeliveryDate, not even an explicit nil
### GetBillingPeriodStart

`func (o *FacturXInvoice) GetBillingPeriodStart() string`

GetBillingPeriodStart returns the BillingPeriodStart field if non-nil, zero value otherwise.

### GetBillingPeriodStartOk

`func (o *FacturXInvoice) GetBillingPeriodStartOk() (*string, bool)`

GetBillingPeriodStartOk returns a tuple with the BillingPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodStart

`func (o *FacturXInvoice) SetBillingPeriodStart(v string)`

SetBillingPeriodStart sets BillingPeriodStart field to given value.

### HasBillingPeriodStart

`func (o *FacturXInvoice) HasBillingPeriodStart() bool`

HasBillingPeriodStart returns a boolean if a field has been set.

### SetBillingPeriodStartNil

`func (o *FacturXInvoice) SetBillingPeriodStartNil(b bool)`

 SetBillingPeriodStartNil sets the value for BillingPeriodStart to be an explicit nil

### UnsetBillingPeriodStart
`func (o *FacturXInvoice) UnsetBillingPeriodStart()`

UnsetBillingPeriodStart ensures that no value is present for BillingPeriodStart, not even an explicit nil
### GetBillingPeriodEnd

`func (o *FacturXInvoice) GetBillingPeriodEnd() string`

GetBillingPeriodEnd returns the BillingPeriodEnd field if non-nil, zero value otherwise.

### GetBillingPeriodEndOk

`func (o *FacturXInvoice) GetBillingPeriodEndOk() (*string, bool)`

GetBillingPeriodEndOk returns a tuple with the BillingPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodEnd

`func (o *FacturXInvoice) SetBillingPeriodEnd(v string)`

SetBillingPeriodEnd sets BillingPeriodEnd field to given value.

### HasBillingPeriodEnd

`func (o *FacturXInvoice) HasBillingPeriodEnd() bool`

HasBillingPeriodEnd returns a boolean if a field has been set.

### SetBillingPeriodEndNil

`func (o *FacturXInvoice) SetBillingPeriodEndNil(b bool)`

 SetBillingPeriodEndNil sets the value for BillingPeriodEnd to be an explicit nil

### UnsetBillingPeriodEnd
`func (o *FacturXInvoice) UnsetBillingPeriodEnd()`

UnsetBillingPeriodEnd ensures that no value is present for BillingPeriodEnd, not even an explicit nil
### GetPaymentReference

`func (o *FacturXInvoice) GetPaymentReference() string`

GetPaymentReference returns the PaymentReference field if non-nil, zero value otherwise.

### GetPaymentReferenceOk

`func (o *FacturXInvoice) GetPaymentReferenceOk() (*string, bool)`

GetPaymentReferenceOk returns a tuple with the PaymentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentReference

`func (o *FacturXInvoice) SetPaymentReference(v string)`

SetPaymentReference sets PaymentReference field to given value.

### HasPaymentReference

`func (o *FacturXInvoice) HasPaymentReference() bool`

HasPaymentReference returns a boolean if a field has been set.

### SetPaymentReferenceNil

`func (o *FacturXInvoice) SetPaymentReferenceNil(b bool)`

 SetPaymentReferenceNil sets the value for PaymentReference to be an explicit nil

### UnsetPaymentReference
`func (o *FacturXInvoice) UnsetPaymentReference()`

UnsetPaymentReference ensures that no value is present for PaymentReference, not even an explicit nil
### GetCreditorReferenceId

`func (o *FacturXInvoice) GetCreditorReferenceId() string`

GetCreditorReferenceId returns the CreditorReferenceId field if non-nil, zero value otherwise.

### GetCreditorReferenceIdOk

`func (o *FacturXInvoice) GetCreditorReferenceIdOk() (*string, bool)`

GetCreditorReferenceIdOk returns a tuple with the CreditorReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditorReferenceId

`func (o *FacturXInvoice) SetCreditorReferenceId(v string)`

SetCreditorReferenceId sets CreditorReferenceId field to given value.

### HasCreditorReferenceId

`func (o *FacturXInvoice) HasCreditorReferenceId() bool`

HasCreditorReferenceId returns a boolean if a field has been set.

### SetCreditorReferenceIdNil

`func (o *FacturXInvoice) SetCreditorReferenceIdNil(b bool)`

 SetCreditorReferenceIdNil sets the value for CreditorReferenceId to be an explicit nil

### UnsetCreditorReferenceId
`func (o *FacturXInvoice) UnsetCreditorReferenceId()`

UnsetCreditorReferenceId ensures that no value is present for CreditorReferenceId, not even an explicit nil
### GetDirectDebitMandateId

`func (o *FacturXInvoice) GetDirectDebitMandateId() string`

GetDirectDebitMandateId returns the DirectDebitMandateId field if non-nil, zero value otherwise.

### GetDirectDebitMandateIdOk

`func (o *FacturXInvoice) GetDirectDebitMandateIdOk() (*string, bool)`

GetDirectDebitMandateIdOk returns a tuple with the DirectDebitMandateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebitMandateId

`func (o *FacturXInvoice) SetDirectDebitMandateId(v string)`

SetDirectDebitMandateId sets DirectDebitMandateId field to given value.

### HasDirectDebitMandateId

`func (o *FacturXInvoice) HasDirectDebitMandateId() bool`

HasDirectDebitMandateId returns a boolean if a field has been set.

### SetDirectDebitMandateIdNil

`func (o *FacturXInvoice) SetDirectDebitMandateIdNil(b bool)`

 SetDirectDebitMandateIdNil sets the value for DirectDebitMandateId to be an explicit nil

### UnsetDirectDebitMandateId
`func (o *FacturXInvoice) UnsetDirectDebitMandateId()`

UnsetDirectDebitMandateId ensures that no value is present for DirectDebitMandateId, not even an explicit nil
### GetDebtorIban

`func (o *FacturXInvoice) GetDebtorIban() string`

GetDebtorIban returns the DebtorIban field if non-nil, zero value otherwise.

### GetDebtorIbanOk

`func (o *FacturXInvoice) GetDebtorIbanOk() (*string, bool)`

GetDebtorIbanOk returns a tuple with the DebtorIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebtorIban

`func (o *FacturXInvoice) SetDebtorIban(v string)`

SetDebtorIban sets DebtorIban field to given value.

### HasDebtorIban

`func (o *FacturXInvoice) HasDebtorIban() bool`

HasDebtorIban returns a boolean if a field has been set.

### SetDebtorIbanNil

`func (o *FacturXInvoice) SetDebtorIbanNil(b bool)`

 SetDebtorIbanNil sets the value for DebtorIban to be an explicit nil

### UnsetDebtorIban
`func (o *FacturXInvoice) UnsetDebtorIban()`

UnsetDebtorIban ensures that no value is present for DebtorIban, not even an explicit nil
### GetPaymentTerms

`func (o *FacturXInvoice) GetPaymentTerms() string`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *FacturXInvoice) GetPaymentTermsOk() (*string, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *FacturXInvoice) SetPaymentTerms(v string)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *FacturXInvoice) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.

### SetPaymentTermsNil

`func (o *FacturXInvoice) SetPaymentTermsNil(b bool)`

 SetPaymentTermsNil sets the value for PaymentTerms to be an explicit nil

### UnsetPaymentTerms
`func (o *FacturXInvoice) UnsetPaymentTerms()`

UnsetPaymentTerms ensures that no value is present for PaymentTerms, not even an explicit nil
### GetAllowancesCharges

`func (o *FacturXInvoice) GetAllowancesCharges() []AllowanceCharge`

GetAllowancesCharges returns the AllowancesCharges field if non-nil, zero value otherwise.

### GetAllowancesChargesOk

`func (o *FacturXInvoice) GetAllowancesChargesOk() (*[]AllowanceCharge, bool)`

GetAllowancesChargesOk returns a tuple with the AllowancesCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowancesCharges

`func (o *FacturXInvoice) SetAllowancesCharges(v []AllowanceCharge)`

SetAllowancesCharges sets AllowancesCharges field to given value.

### HasAllowancesCharges

`func (o *FacturXInvoice) HasAllowancesCharges() bool`

HasAllowancesCharges returns a boolean if a field has been set.

### SetAllowancesChargesNil

`func (o *FacturXInvoice) SetAllowancesChargesNil(b bool)`

 SetAllowancesChargesNil sets the value for AllowancesCharges to be an explicit nil

### UnsetAllowancesCharges
`func (o *FacturXInvoice) UnsetAllowancesCharges()`

UnsetAllowancesCharges ensures that no value is present for AllowancesCharges, not even an explicit nil
### GetAdditionalDocuments

`func (o *FacturXInvoice) GetAdditionalDocuments() []AdditionalDocument`

GetAdditionalDocuments returns the AdditionalDocuments field if non-nil, zero value otherwise.

### GetAdditionalDocumentsOk

`func (o *FacturXInvoice) GetAdditionalDocumentsOk() (*[]AdditionalDocument, bool)`

GetAdditionalDocumentsOk returns a tuple with the AdditionalDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDocuments

`func (o *FacturXInvoice) SetAdditionalDocuments(v []AdditionalDocument)`

SetAdditionalDocuments sets AdditionalDocuments field to given value.

### HasAdditionalDocuments

`func (o *FacturXInvoice) HasAdditionalDocuments() bool`

HasAdditionalDocuments returns a boolean if a field has been set.

### SetAdditionalDocumentsNil

`func (o *FacturXInvoice) SetAdditionalDocumentsNil(b bool)`

 SetAdditionalDocumentsNil sets the value for AdditionalDocuments to be an explicit nil

### UnsetAdditionalDocuments
`func (o *FacturXInvoice) UnsetAdditionalDocuments()`

UnsetAdditionalDocuments ensures that no value is present for AdditionalDocuments, not even an explicit nil
### GetBuyerAccountingReference

`func (o *FacturXInvoice) GetBuyerAccountingReference() string`

GetBuyerAccountingReference returns the BuyerAccountingReference field if non-nil, zero value otherwise.

### GetBuyerAccountingReferenceOk

`func (o *FacturXInvoice) GetBuyerAccountingReferenceOk() (*string, bool)`

GetBuyerAccountingReferenceOk returns a tuple with the BuyerAccountingReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerAccountingReference

`func (o *FacturXInvoice) SetBuyerAccountingReference(v string)`

SetBuyerAccountingReference sets BuyerAccountingReference field to given value.

### HasBuyerAccountingReference

`func (o *FacturXInvoice) HasBuyerAccountingReference() bool`

HasBuyerAccountingReference returns a boolean if a field has been set.

### SetBuyerAccountingReferenceNil

`func (o *FacturXInvoice) SetBuyerAccountingReferenceNil(b bool)`

 SetBuyerAccountingReferenceNil sets the value for BuyerAccountingReference to be an explicit nil

### UnsetBuyerAccountingReference
`func (o *FacturXInvoice) UnsetBuyerAccountingReference()`

UnsetBuyerAccountingReference ensures that no value is present for BuyerAccountingReference, not even an explicit nil
### GetPaymentCard

`func (o *FacturXInvoice) GetPaymentCard() PaymentCard`

GetPaymentCard returns the PaymentCard field if non-nil, zero value otherwise.

### GetPaymentCardOk

`func (o *FacturXInvoice) GetPaymentCardOk() (*PaymentCard, bool)`

GetPaymentCardOk returns a tuple with the PaymentCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCard

`func (o *FacturXInvoice) SetPaymentCard(v PaymentCard)`

SetPaymentCard sets PaymentCard field to given value.

### HasPaymentCard

`func (o *FacturXInvoice) HasPaymentCard() bool`

HasPaymentCard returns a boolean if a field has been set.

### SetPaymentCardNil

`func (o *FacturXInvoice) SetPaymentCardNil(b bool)`

 SetPaymentCardNil sets the value for PaymentCard to be an explicit nil

### UnsetPaymentCard
`func (o *FacturXInvoice) UnsetPaymentCard()`

UnsetPaymentCard ensures that no value is present for PaymentCard, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Invoicedata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** | Unique invoice number | 
**Supplier** | [**Supplier**](Supplier.md) |  | 
**Recipient** | [**Recipient**](Recipient.md) |  | 
**Lines** | **[]map[string]interface{}** | Invoice lines | 
**Date** | Pointer to **string** |  | [optional] 
**DueDays** | Pointer to **int32** | Due date in days (default: 30) | [optional] [default to 30]
**Comment** | Pointer to **string** |  | [optional] 
**PurchaseOrderReference** | Pointer to **string** |  | [optional] 
**ContractReference** | Pointer to **string** |  | [optional] 
**InvoiceType** | Pointer to [**FactureElectroniqueModelsInvoiceTypeCode**](FactureElectroniqueModelsInvoiceTypeCode.md) | Document type (UNTDID 1001). Default: 380 (Invoice). | [optional] [default to INVOICE]
**PrecedingInvoiceReference** | Pointer to **string** |  | [optional] 
**OperationNature** | Pointer to [**OperationNature**](OperationNature.md) |  | [optional] 
**InvoicingFramework** | Pointer to [**InvoicingFrameworkCode**](InvoicingFrameworkCode.md) |  | [optional] 
**InvoiceNumber** | **string** |  | 
**PaymentDueDate** | **string** |  | 
**InvoiceDate** | Pointer to **string** |  | [optional] 
**SubmissionMode** | Pointer to [**SubmissionMode**](SubmissionMode.md) |  | [optional] [default to DEPOT_PDF_API]
**InvoicingFramework** | Pointer to [**InvoicingFramework**](InvoicingFramework.md) |  | [optional] 
**References** | [**InvoiceReferences**](InvoiceReferences.md) |  | 
**Totals** | [**InvoiceTotals**](InvoiceTotals.md) |  | 
**InvoiceLines** | Pointer to [**[]InvoiceLine**](InvoiceLine.md) |  | [optional] 
**VatLines** | Pointer to [**[]VATLine**](VATLine.md) |  | [optional] 
**Notes** | Pointer to [**[]InvoiceNote**](InvoiceNote.md) |  | [optional] 
**CurrentUserId** | Pointer to **int32** |  | [optional] 
**SupplementaryAttachments** | Pointer to [**[]SupplementaryAttachment**](SupplementaryAttachment.md) |  | [optional] 
**Payee** | Pointer to [**Payee**](Payee.md) |  | [optional] 
**DeliveryParty** | Pointer to [**DeliveryParty**](DeliveryParty.md) |  | [optional] 
**TaxRepresentative** | Pointer to [**TaxRepresentative**](TaxRepresentative.md) |  | [optional] 
**DeliveryDate** | Pointer to **string** |  | [optional] 
**BillingPeriodStart** | Pointer to **string** |  | [optional] 
**BillingPeriodEnd** | Pointer to **string** |  | [optional] 
**PaymentReference** | Pointer to **string** |  | [optional] 
**CreditorReferenceId** | Pointer to **string** |  | [optional] 
**DirectDebitMandateId** | Pointer to **string** |  | [optional] 
**DebtorIban** | Pointer to **string** |  | [optional] 
**PaymentTerms** | Pointer to **string** |  | [optional] 
**AllowancesCharges** | Pointer to [**[]AllowanceCharge**](AllowanceCharge.md) |  | [optional] 
**AdditionalDocuments** | Pointer to [**[]AdditionalDocument**](AdditionalDocument.md) |  | [optional] 
**BuyerAccountingReference** | Pointer to **string** |  | [optional] 
**PaymentCard** | Pointer to [**PaymentCard**](PaymentCard.md) |  | [optional] 

## Methods

### NewInvoicedata

`func NewInvoicedata(number string, supplier Supplier, recipient Recipient, lines []map[string]interface{}, invoiceNumber string, paymentDueDate string, references InvoiceReferences, totals InvoiceTotals, ) *Invoicedata`

NewInvoicedata instantiates a new Invoicedata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicedataWithDefaults

`func NewInvoicedataWithDefaults() *Invoicedata`

NewInvoicedataWithDefaults instantiates a new Invoicedata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *Invoicedata) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Invoicedata) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Invoicedata) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetSupplier

`func (o *Invoicedata) GetSupplier() Supplier`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *Invoicedata) GetSupplierOk() (*Supplier, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *Invoicedata) SetSupplier(v Supplier)`

SetSupplier sets Supplier field to given value.


### GetRecipient

`func (o *Invoicedata) GetRecipient() Recipient`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *Invoicedata) GetRecipientOk() (*Recipient, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *Invoicedata) SetRecipient(v Recipient)`

SetRecipient sets Recipient field to given value.


### GetLines

`func (o *Invoicedata) GetLines() []map[string]interface{}`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *Invoicedata) GetLinesOk() (*[]map[string]interface{}, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *Invoicedata) SetLines(v []map[string]interface{})`

SetLines sets Lines field to given value.


### GetDate

`func (o *Invoicedata) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Invoicedata) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Invoicedata) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Invoicedata) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDueDays

`func (o *Invoicedata) GetDueDays() int32`

GetDueDays returns the DueDays field if non-nil, zero value otherwise.

### GetDueDaysOk

`func (o *Invoicedata) GetDueDaysOk() (*int32, bool)`

GetDueDaysOk returns a tuple with the DueDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDays

`func (o *Invoicedata) SetDueDays(v int32)`

SetDueDays sets DueDays field to given value.

### HasDueDays

`func (o *Invoicedata) HasDueDays() bool`

HasDueDays returns a boolean if a field has been set.

### GetComment

`func (o *Invoicedata) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Invoicedata) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Invoicedata) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Invoicedata) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPurchaseOrderReference

`func (o *Invoicedata) GetPurchaseOrderReference() string`

GetPurchaseOrderReference returns the PurchaseOrderReference field if non-nil, zero value otherwise.

### GetPurchaseOrderReferenceOk

`func (o *Invoicedata) GetPurchaseOrderReferenceOk() (*string, bool)`

GetPurchaseOrderReferenceOk returns a tuple with the PurchaseOrderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderReference

`func (o *Invoicedata) SetPurchaseOrderReference(v string)`

SetPurchaseOrderReference sets PurchaseOrderReference field to given value.

### HasPurchaseOrderReference

`func (o *Invoicedata) HasPurchaseOrderReference() bool`

HasPurchaseOrderReference returns a boolean if a field has been set.

### GetContractReference

`func (o *Invoicedata) GetContractReference() string`

GetContractReference returns the ContractReference field if non-nil, zero value otherwise.

### GetContractReferenceOk

`func (o *Invoicedata) GetContractReferenceOk() (*string, bool)`

GetContractReferenceOk returns a tuple with the ContractReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractReference

`func (o *Invoicedata) SetContractReference(v string)`

SetContractReference sets ContractReference field to given value.

### HasContractReference

`func (o *Invoicedata) HasContractReference() bool`

HasContractReference returns a boolean if a field has been set.

### GetInvoiceType

`func (o *Invoicedata) GetInvoiceType() FactureElectroniqueModelsInvoiceTypeCode`

GetInvoiceType returns the InvoiceType field if non-nil, zero value otherwise.

### GetInvoiceTypeOk

`func (o *Invoicedata) GetInvoiceTypeOk() (*FactureElectroniqueModelsInvoiceTypeCode, bool)`

GetInvoiceTypeOk returns a tuple with the InvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceType

`func (o *Invoicedata) SetInvoiceType(v FactureElectroniqueModelsInvoiceTypeCode)`

SetInvoiceType sets InvoiceType field to given value.

### HasInvoiceType

`func (o *Invoicedata) HasInvoiceType() bool`

HasInvoiceType returns a boolean if a field has been set.

### GetPrecedingInvoiceReference

`func (o *Invoicedata) GetPrecedingInvoiceReference() string`

GetPrecedingInvoiceReference returns the PrecedingInvoiceReference field if non-nil, zero value otherwise.

### GetPrecedingInvoiceReferenceOk

`func (o *Invoicedata) GetPrecedingInvoiceReferenceOk() (*string, bool)`

GetPrecedingInvoiceReferenceOk returns a tuple with the PrecedingInvoiceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedingInvoiceReference

`func (o *Invoicedata) SetPrecedingInvoiceReference(v string)`

SetPrecedingInvoiceReference sets PrecedingInvoiceReference field to given value.

### HasPrecedingInvoiceReference

`func (o *Invoicedata) HasPrecedingInvoiceReference() bool`

HasPrecedingInvoiceReference returns a boolean if a field has been set.

### GetOperationNature

`func (o *Invoicedata) GetOperationNature() OperationNature`

GetOperationNature returns the OperationNature field if non-nil, zero value otherwise.

### GetOperationNatureOk

`func (o *Invoicedata) GetOperationNatureOk() (*OperationNature, bool)`

GetOperationNatureOk returns a tuple with the OperationNature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationNature

`func (o *Invoicedata) SetOperationNature(v OperationNature)`

SetOperationNature sets OperationNature field to given value.

### HasOperationNature

`func (o *Invoicedata) HasOperationNature() bool`

HasOperationNature returns a boolean if a field has been set.

### GetInvoicingFramework

`func (o *Invoicedata) GetInvoicingFramework() InvoicingFrameworkCode`

GetInvoicingFramework returns the InvoicingFramework field if non-nil, zero value otherwise.

### GetInvoicingFrameworkOk

`func (o *Invoicedata) GetInvoicingFrameworkOk() (*InvoicingFrameworkCode, bool)`

GetInvoicingFrameworkOk returns a tuple with the InvoicingFramework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicingFramework

`func (o *Invoicedata) SetInvoicingFramework(v InvoicingFrameworkCode)`

SetInvoicingFramework sets InvoicingFramework field to given value.

### HasInvoicingFramework

`func (o *Invoicedata) HasInvoicingFramework() bool`

HasInvoicingFramework returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *Invoicedata) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *Invoicedata) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *Invoicedata) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.


### GetPaymentDueDate

`func (o *Invoicedata) GetPaymentDueDate() string`

GetPaymentDueDate returns the PaymentDueDate field if non-nil, zero value otherwise.

### GetPaymentDueDateOk

`func (o *Invoicedata) GetPaymentDueDateOk() (*string, bool)`

GetPaymentDueDateOk returns a tuple with the PaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDueDate

`func (o *Invoicedata) SetPaymentDueDate(v string)`

SetPaymentDueDate sets PaymentDueDate field to given value.


### GetInvoiceDate

`func (o *Invoicedata) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *Invoicedata) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *Invoicedata) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *Invoicedata) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetSubmissionMode

`func (o *Invoicedata) GetSubmissionMode() SubmissionMode`

GetSubmissionMode returns the SubmissionMode field if non-nil, zero value otherwise.

### GetSubmissionModeOk

`func (o *Invoicedata) GetSubmissionModeOk() (*SubmissionMode, bool)`

GetSubmissionModeOk returns a tuple with the SubmissionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionMode

`func (o *Invoicedata) SetSubmissionMode(v SubmissionMode)`

SetSubmissionMode sets SubmissionMode field to given value.

### HasSubmissionMode

`func (o *Invoicedata) HasSubmissionMode() bool`

HasSubmissionMode returns a boolean if a field has been set.

### GetInvoicingFramework

`func (o *Invoicedata) GetInvoicingFramework() InvoicingFramework`

GetInvoicingFramework returns the InvoicingFramework field if non-nil, zero value otherwise.

### GetInvoicingFrameworkOk

`func (o *Invoicedata) GetInvoicingFrameworkOk() (*InvoicingFramework, bool)`

GetInvoicingFrameworkOk returns a tuple with the InvoicingFramework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicingFramework

`func (o *Invoicedata) SetInvoicingFramework(v InvoicingFramework)`

SetInvoicingFramework sets InvoicingFramework field to given value.

### HasInvoicingFramework

`func (o *Invoicedata) HasInvoicingFramework() bool`

HasInvoicingFramework returns a boolean if a field has been set.

### GetReferences

`func (o *Invoicedata) GetReferences() InvoiceReferences`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *Invoicedata) GetReferencesOk() (*InvoiceReferences, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *Invoicedata) SetReferences(v InvoiceReferences)`

SetReferences sets References field to given value.


### GetTotals

`func (o *Invoicedata) GetTotals() InvoiceTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *Invoicedata) GetTotalsOk() (*InvoiceTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *Invoicedata) SetTotals(v InvoiceTotals)`

SetTotals sets Totals field to given value.


### GetInvoiceLines

`func (o *Invoicedata) GetInvoiceLines() []InvoiceLine`

GetInvoiceLines returns the InvoiceLines field if non-nil, zero value otherwise.

### GetInvoiceLinesOk

`func (o *Invoicedata) GetInvoiceLinesOk() (*[]InvoiceLine, bool)`

GetInvoiceLinesOk returns a tuple with the InvoiceLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceLines

`func (o *Invoicedata) SetInvoiceLines(v []InvoiceLine)`

SetInvoiceLines sets InvoiceLines field to given value.

### HasInvoiceLines

`func (o *Invoicedata) HasInvoiceLines() bool`

HasInvoiceLines returns a boolean if a field has been set.

### GetVatLines

`func (o *Invoicedata) GetVatLines() []VATLine`

GetVatLines returns the VatLines field if non-nil, zero value otherwise.

### GetVatLinesOk

`func (o *Invoicedata) GetVatLinesOk() (*[]VATLine, bool)`

GetVatLinesOk returns a tuple with the VatLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatLines

`func (o *Invoicedata) SetVatLines(v []VATLine)`

SetVatLines sets VatLines field to given value.

### HasVatLines

`func (o *Invoicedata) HasVatLines() bool`

HasVatLines returns a boolean if a field has been set.

### GetNotes

`func (o *Invoicedata) GetNotes() []InvoiceNote`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Invoicedata) GetNotesOk() (*[]InvoiceNote, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Invoicedata) SetNotes(v []InvoiceNote)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Invoicedata) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetCurrentUserId

`func (o *Invoicedata) GetCurrentUserId() int32`

GetCurrentUserId returns the CurrentUserId field if non-nil, zero value otherwise.

### GetCurrentUserIdOk

`func (o *Invoicedata) GetCurrentUserIdOk() (*int32, bool)`

GetCurrentUserIdOk returns a tuple with the CurrentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUserId

`func (o *Invoicedata) SetCurrentUserId(v int32)`

SetCurrentUserId sets CurrentUserId field to given value.

### HasCurrentUserId

`func (o *Invoicedata) HasCurrentUserId() bool`

HasCurrentUserId returns a boolean if a field has been set.

### GetSupplementaryAttachments

`func (o *Invoicedata) GetSupplementaryAttachments() []SupplementaryAttachment`

GetSupplementaryAttachments returns the SupplementaryAttachments field if non-nil, zero value otherwise.

### GetSupplementaryAttachmentsOk

`func (o *Invoicedata) GetSupplementaryAttachmentsOk() (*[]SupplementaryAttachment, bool)`

GetSupplementaryAttachmentsOk returns a tuple with the SupplementaryAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementaryAttachments

`func (o *Invoicedata) SetSupplementaryAttachments(v []SupplementaryAttachment)`

SetSupplementaryAttachments sets SupplementaryAttachments field to given value.

### HasSupplementaryAttachments

`func (o *Invoicedata) HasSupplementaryAttachments() bool`

HasSupplementaryAttachments returns a boolean if a field has been set.

### GetPayee

`func (o *Invoicedata) GetPayee() Payee`

GetPayee returns the Payee field if non-nil, zero value otherwise.

### GetPayeeOk

`func (o *Invoicedata) GetPayeeOk() (*Payee, bool)`

GetPayeeOk returns a tuple with the Payee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayee

`func (o *Invoicedata) SetPayee(v Payee)`

SetPayee sets Payee field to given value.

### HasPayee

`func (o *Invoicedata) HasPayee() bool`

HasPayee returns a boolean if a field has been set.

### GetDeliveryParty

`func (o *Invoicedata) GetDeliveryParty() DeliveryParty`

GetDeliveryParty returns the DeliveryParty field if non-nil, zero value otherwise.

### GetDeliveryPartyOk

`func (o *Invoicedata) GetDeliveryPartyOk() (*DeliveryParty, bool)`

GetDeliveryPartyOk returns a tuple with the DeliveryParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryParty

`func (o *Invoicedata) SetDeliveryParty(v DeliveryParty)`

SetDeliveryParty sets DeliveryParty field to given value.

### HasDeliveryParty

`func (o *Invoicedata) HasDeliveryParty() bool`

HasDeliveryParty returns a boolean if a field has been set.

### GetTaxRepresentative

`func (o *Invoicedata) GetTaxRepresentative() TaxRepresentative`

GetTaxRepresentative returns the TaxRepresentative field if non-nil, zero value otherwise.

### GetTaxRepresentativeOk

`func (o *Invoicedata) GetTaxRepresentativeOk() (*TaxRepresentative, bool)`

GetTaxRepresentativeOk returns a tuple with the TaxRepresentative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRepresentative

`func (o *Invoicedata) SetTaxRepresentative(v TaxRepresentative)`

SetTaxRepresentative sets TaxRepresentative field to given value.

### HasTaxRepresentative

`func (o *Invoicedata) HasTaxRepresentative() bool`

HasTaxRepresentative returns a boolean if a field has been set.

### GetDeliveryDate

`func (o *Invoicedata) GetDeliveryDate() string`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *Invoicedata) GetDeliveryDateOk() (*string, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *Invoicedata) SetDeliveryDate(v string)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *Invoicedata) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetBillingPeriodStart

`func (o *Invoicedata) GetBillingPeriodStart() string`

GetBillingPeriodStart returns the BillingPeriodStart field if non-nil, zero value otherwise.

### GetBillingPeriodStartOk

`func (o *Invoicedata) GetBillingPeriodStartOk() (*string, bool)`

GetBillingPeriodStartOk returns a tuple with the BillingPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodStart

`func (o *Invoicedata) SetBillingPeriodStart(v string)`

SetBillingPeriodStart sets BillingPeriodStart field to given value.

### HasBillingPeriodStart

`func (o *Invoicedata) HasBillingPeriodStart() bool`

HasBillingPeriodStart returns a boolean if a field has been set.

### GetBillingPeriodEnd

`func (o *Invoicedata) GetBillingPeriodEnd() string`

GetBillingPeriodEnd returns the BillingPeriodEnd field if non-nil, zero value otherwise.

### GetBillingPeriodEndOk

`func (o *Invoicedata) GetBillingPeriodEndOk() (*string, bool)`

GetBillingPeriodEndOk returns a tuple with the BillingPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodEnd

`func (o *Invoicedata) SetBillingPeriodEnd(v string)`

SetBillingPeriodEnd sets BillingPeriodEnd field to given value.

### HasBillingPeriodEnd

`func (o *Invoicedata) HasBillingPeriodEnd() bool`

HasBillingPeriodEnd returns a boolean if a field has been set.

### GetPaymentReference

`func (o *Invoicedata) GetPaymentReference() string`

GetPaymentReference returns the PaymentReference field if non-nil, zero value otherwise.

### GetPaymentReferenceOk

`func (o *Invoicedata) GetPaymentReferenceOk() (*string, bool)`

GetPaymentReferenceOk returns a tuple with the PaymentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentReference

`func (o *Invoicedata) SetPaymentReference(v string)`

SetPaymentReference sets PaymentReference field to given value.

### HasPaymentReference

`func (o *Invoicedata) HasPaymentReference() bool`

HasPaymentReference returns a boolean if a field has been set.

### GetCreditorReferenceId

`func (o *Invoicedata) GetCreditorReferenceId() string`

GetCreditorReferenceId returns the CreditorReferenceId field if non-nil, zero value otherwise.

### GetCreditorReferenceIdOk

`func (o *Invoicedata) GetCreditorReferenceIdOk() (*string, bool)`

GetCreditorReferenceIdOk returns a tuple with the CreditorReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditorReferenceId

`func (o *Invoicedata) SetCreditorReferenceId(v string)`

SetCreditorReferenceId sets CreditorReferenceId field to given value.

### HasCreditorReferenceId

`func (o *Invoicedata) HasCreditorReferenceId() bool`

HasCreditorReferenceId returns a boolean if a field has been set.

### GetDirectDebitMandateId

`func (o *Invoicedata) GetDirectDebitMandateId() string`

GetDirectDebitMandateId returns the DirectDebitMandateId field if non-nil, zero value otherwise.

### GetDirectDebitMandateIdOk

`func (o *Invoicedata) GetDirectDebitMandateIdOk() (*string, bool)`

GetDirectDebitMandateIdOk returns a tuple with the DirectDebitMandateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebitMandateId

`func (o *Invoicedata) SetDirectDebitMandateId(v string)`

SetDirectDebitMandateId sets DirectDebitMandateId field to given value.

### HasDirectDebitMandateId

`func (o *Invoicedata) HasDirectDebitMandateId() bool`

HasDirectDebitMandateId returns a boolean if a field has been set.

### GetDebtorIban

`func (o *Invoicedata) GetDebtorIban() string`

GetDebtorIban returns the DebtorIban field if non-nil, zero value otherwise.

### GetDebtorIbanOk

`func (o *Invoicedata) GetDebtorIbanOk() (*string, bool)`

GetDebtorIbanOk returns a tuple with the DebtorIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebtorIban

`func (o *Invoicedata) SetDebtorIban(v string)`

SetDebtorIban sets DebtorIban field to given value.

### HasDebtorIban

`func (o *Invoicedata) HasDebtorIban() bool`

HasDebtorIban returns a boolean if a field has been set.

### GetPaymentTerms

`func (o *Invoicedata) GetPaymentTerms() string`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *Invoicedata) GetPaymentTermsOk() (*string, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *Invoicedata) SetPaymentTerms(v string)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *Invoicedata) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.

### GetAllowancesCharges

`func (o *Invoicedata) GetAllowancesCharges() []AllowanceCharge`

GetAllowancesCharges returns the AllowancesCharges field if non-nil, zero value otherwise.

### GetAllowancesChargesOk

`func (o *Invoicedata) GetAllowancesChargesOk() (*[]AllowanceCharge, bool)`

GetAllowancesChargesOk returns a tuple with the AllowancesCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowancesCharges

`func (o *Invoicedata) SetAllowancesCharges(v []AllowanceCharge)`

SetAllowancesCharges sets AllowancesCharges field to given value.

### HasAllowancesCharges

`func (o *Invoicedata) HasAllowancesCharges() bool`

HasAllowancesCharges returns a boolean if a field has been set.

### GetAdditionalDocuments

`func (o *Invoicedata) GetAdditionalDocuments() []AdditionalDocument`

GetAdditionalDocuments returns the AdditionalDocuments field if non-nil, zero value otherwise.

### GetAdditionalDocumentsOk

`func (o *Invoicedata) GetAdditionalDocumentsOk() (*[]AdditionalDocument, bool)`

GetAdditionalDocumentsOk returns a tuple with the AdditionalDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDocuments

`func (o *Invoicedata) SetAdditionalDocuments(v []AdditionalDocument)`

SetAdditionalDocuments sets AdditionalDocuments field to given value.

### HasAdditionalDocuments

`func (o *Invoicedata) HasAdditionalDocuments() bool`

HasAdditionalDocuments returns a boolean if a field has been set.

### GetBuyerAccountingReference

`func (o *Invoicedata) GetBuyerAccountingReference() string`

GetBuyerAccountingReference returns the BuyerAccountingReference field if non-nil, zero value otherwise.

### GetBuyerAccountingReferenceOk

`func (o *Invoicedata) GetBuyerAccountingReferenceOk() (*string, bool)`

GetBuyerAccountingReferenceOk returns a tuple with the BuyerAccountingReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerAccountingReference

`func (o *Invoicedata) SetBuyerAccountingReference(v string)`

SetBuyerAccountingReference sets BuyerAccountingReference field to given value.

### HasBuyerAccountingReference

`func (o *Invoicedata) HasBuyerAccountingReference() bool`

HasBuyerAccountingReference returns a boolean if a field has been set.

### GetPaymentCard

`func (o *Invoicedata) GetPaymentCard() PaymentCard`

GetPaymentCard returns the PaymentCard field if non-nil, zero value otherwise.

### GetPaymentCardOk

`func (o *Invoicedata) GetPaymentCardOk() (*PaymentCard, bool)`

GetPaymentCardOk returns a tuple with the PaymentCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCard

`func (o *Invoicedata) SetPaymentCard(v PaymentCard)`

SetPaymentCard sets PaymentCard field to given value.

### HasPaymentCard

`func (o *Invoicedata) HasPaymentCard() bool`

HasPaymentCard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# IncomingInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowId** | Pointer to **NullableString** |  | [optional] 
**SourceFormat** | [**InvoiceFormat**](InvoiceFormat.md) | Invoice source format | 
**SupplierReference** | **string** | Invoice number issued by the supplier (BT-1) | 
**DocumentType** | Pointer to [**InvoiceTypeCode**](InvoiceTypeCode.md) | Document type (BT-3) | [optional] [default to INVOICE]
**Supplier** | [**IncomingSupplier**](IncomingSupplier.md) | Invoice issuer (SellerTradeParty) | 
**BillingSiteName** | **string** | Recipient name / your company (BT-44) | 
**BillingSiteSiret** | Pointer to **NullableString** |  | [optional] 
**IssueDate** | **string** | Invoice date (BT-2) - YYYY-MM-DD | 
**DueDate** | Pointer to **NullableString** |  | [optional] 
**Currency** | Pointer to **string** | ISO currency code (BT-5) | [optional] [default to "EUR"]
**NetAmount** | **string** | Total net amount (BT-109) | 
**VatAmount** | **string** | Total VAT amount (BT-110) | 
**GrossAmount** | **string** | Total gross amount (BT-112) | 
**PurchaseOrderNumber** | Pointer to **NullableString** |  | [optional] 
**ContractReference** | Pointer to **NullableString** |  | [optional] 
**InvoiceSubject** | Pointer to **NullableString** |  | [optional] 
**DocumentBase64** | Pointer to **NullableString** |  | [optional] 
**DocumentContentType** | Pointer to **NullableString** |  | [optional] 
**DocumentFilename** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIncomingInvoice

`func NewIncomingInvoice(sourceFormat InvoiceFormat, supplierReference string, supplier IncomingSupplier, billingSiteName string, issueDate string, netAmount string, vatAmount string, grossAmount string, ) *IncomingInvoice`

NewIncomingInvoice instantiates a new IncomingInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomingInvoiceWithDefaults

`func NewIncomingInvoiceWithDefaults() *IncomingInvoice`

NewIncomingInvoiceWithDefaults instantiates a new IncomingInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowId

`func (o *IncomingInvoice) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *IncomingInvoice) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *IncomingInvoice) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *IncomingInvoice) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### SetFlowIdNil

`func (o *IncomingInvoice) SetFlowIdNil(b bool)`

 SetFlowIdNil sets the value for FlowId to be an explicit nil

### UnsetFlowId
`func (o *IncomingInvoice) UnsetFlowId()`

UnsetFlowId ensures that no value is present for FlowId, not even an explicit nil
### GetSourceFormat

`func (o *IncomingInvoice) GetSourceFormat() InvoiceFormat`

GetSourceFormat returns the SourceFormat field if non-nil, zero value otherwise.

### GetSourceFormatOk

`func (o *IncomingInvoice) GetSourceFormatOk() (*InvoiceFormat, bool)`

GetSourceFormatOk returns a tuple with the SourceFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFormat

`func (o *IncomingInvoice) SetSourceFormat(v InvoiceFormat)`

SetSourceFormat sets SourceFormat field to given value.


### GetSupplierReference

`func (o *IncomingInvoice) GetSupplierReference() string`

GetSupplierReference returns the SupplierReference field if non-nil, zero value otherwise.

### GetSupplierReferenceOk

`func (o *IncomingInvoice) GetSupplierReferenceOk() (*string, bool)`

GetSupplierReferenceOk returns a tuple with the SupplierReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierReference

`func (o *IncomingInvoice) SetSupplierReference(v string)`

SetSupplierReference sets SupplierReference field to given value.


### GetDocumentType

`func (o *IncomingInvoice) GetDocumentType() InvoiceTypeCode`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *IncomingInvoice) GetDocumentTypeOk() (*InvoiceTypeCode, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *IncomingInvoice) SetDocumentType(v InvoiceTypeCode)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *IncomingInvoice) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetSupplier

`func (o *IncomingInvoice) GetSupplier() IncomingSupplier`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *IncomingInvoice) GetSupplierOk() (*IncomingSupplier, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *IncomingInvoice) SetSupplier(v IncomingSupplier)`

SetSupplier sets Supplier field to given value.


### GetBillingSiteName

`func (o *IncomingInvoice) GetBillingSiteName() string`

GetBillingSiteName returns the BillingSiteName field if non-nil, zero value otherwise.

### GetBillingSiteNameOk

`func (o *IncomingInvoice) GetBillingSiteNameOk() (*string, bool)`

GetBillingSiteNameOk returns a tuple with the BillingSiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSiteName

`func (o *IncomingInvoice) SetBillingSiteName(v string)`

SetBillingSiteName sets BillingSiteName field to given value.


### GetBillingSiteSiret

`func (o *IncomingInvoice) GetBillingSiteSiret() string`

GetBillingSiteSiret returns the BillingSiteSiret field if non-nil, zero value otherwise.

### GetBillingSiteSiretOk

`func (o *IncomingInvoice) GetBillingSiteSiretOk() (*string, bool)`

GetBillingSiteSiretOk returns a tuple with the BillingSiteSiret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSiteSiret

`func (o *IncomingInvoice) SetBillingSiteSiret(v string)`

SetBillingSiteSiret sets BillingSiteSiret field to given value.

### HasBillingSiteSiret

`func (o *IncomingInvoice) HasBillingSiteSiret() bool`

HasBillingSiteSiret returns a boolean if a field has been set.

### SetBillingSiteSiretNil

`func (o *IncomingInvoice) SetBillingSiteSiretNil(b bool)`

 SetBillingSiteSiretNil sets the value for BillingSiteSiret to be an explicit nil

### UnsetBillingSiteSiret
`func (o *IncomingInvoice) UnsetBillingSiteSiret()`

UnsetBillingSiteSiret ensures that no value is present for BillingSiteSiret, not even an explicit nil
### GetIssueDate

`func (o *IncomingInvoice) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *IncomingInvoice) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *IncomingInvoice) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetDueDate

`func (o *IncomingInvoice) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *IncomingInvoice) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *IncomingInvoice) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *IncomingInvoice) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *IncomingInvoice) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *IncomingInvoice) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetCurrency

`func (o *IncomingInvoice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *IncomingInvoice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *IncomingInvoice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *IncomingInvoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetNetAmount

`func (o *IncomingInvoice) GetNetAmount() string`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *IncomingInvoice) GetNetAmountOk() (*string, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *IncomingInvoice) SetNetAmount(v string)`

SetNetAmount sets NetAmount field to given value.


### GetVatAmount

`func (o *IncomingInvoice) GetVatAmount() string`

GetVatAmount returns the VatAmount field if non-nil, zero value otherwise.

### GetVatAmountOk

`func (o *IncomingInvoice) GetVatAmountOk() (*string, bool)`

GetVatAmountOk returns a tuple with the VatAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatAmount

`func (o *IncomingInvoice) SetVatAmount(v string)`

SetVatAmount sets VatAmount field to given value.


### GetGrossAmount

`func (o *IncomingInvoice) GetGrossAmount() string`

GetGrossAmount returns the GrossAmount field if non-nil, zero value otherwise.

### GetGrossAmountOk

`func (o *IncomingInvoice) GetGrossAmountOk() (*string, bool)`

GetGrossAmountOk returns a tuple with the GrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossAmount

`func (o *IncomingInvoice) SetGrossAmount(v string)`

SetGrossAmount sets GrossAmount field to given value.


### GetPurchaseOrderNumber

`func (o *IncomingInvoice) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *IncomingInvoice) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *IncomingInvoice) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *IncomingInvoice) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### SetPurchaseOrderNumberNil

`func (o *IncomingInvoice) SetPurchaseOrderNumberNil(b bool)`

 SetPurchaseOrderNumberNil sets the value for PurchaseOrderNumber to be an explicit nil

### UnsetPurchaseOrderNumber
`func (o *IncomingInvoice) UnsetPurchaseOrderNumber()`

UnsetPurchaseOrderNumber ensures that no value is present for PurchaseOrderNumber, not even an explicit nil
### GetContractReference

`func (o *IncomingInvoice) GetContractReference() string`

GetContractReference returns the ContractReference field if non-nil, zero value otherwise.

### GetContractReferenceOk

`func (o *IncomingInvoice) GetContractReferenceOk() (*string, bool)`

GetContractReferenceOk returns a tuple with the ContractReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractReference

`func (o *IncomingInvoice) SetContractReference(v string)`

SetContractReference sets ContractReference field to given value.

### HasContractReference

`func (o *IncomingInvoice) HasContractReference() bool`

HasContractReference returns a boolean if a field has been set.

### SetContractReferenceNil

`func (o *IncomingInvoice) SetContractReferenceNil(b bool)`

 SetContractReferenceNil sets the value for ContractReference to be an explicit nil

### UnsetContractReference
`func (o *IncomingInvoice) UnsetContractReference()`

UnsetContractReference ensures that no value is present for ContractReference, not even an explicit nil
### GetInvoiceSubject

`func (o *IncomingInvoice) GetInvoiceSubject() string`

GetInvoiceSubject returns the InvoiceSubject field if non-nil, zero value otherwise.

### GetInvoiceSubjectOk

`func (o *IncomingInvoice) GetInvoiceSubjectOk() (*string, bool)`

GetInvoiceSubjectOk returns a tuple with the InvoiceSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceSubject

`func (o *IncomingInvoice) SetInvoiceSubject(v string)`

SetInvoiceSubject sets InvoiceSubject field to given value.

### HasInvoiceSubject

`func (o *IncomingInvoice) HasInvoiceSubject() bool`

HasInvoiceSubject returns a boolean if a field has been set.

### SetInvoiceSubjectNil

`func (o *IncomingInvoice) SetInvoiceSubjectNil(b bool)`

 SetInvoiceSubjectNil sets the value for InvoiceSubject to be an explicit nil

### UnsetInvoiceSubject
`func (o *IncomingInvoice) UnsetInvoiceSubject()`

UnsetInvoiceSubject ensures that no value is present for InvoiceSubject, not even an explicit nil
### GetDocumentBase64

`func (o *IncomingInvoice) GetDocumentBase64() string`

GetDocumentBase64 returns the DocumentBase64 field if non-nil, zero value otherwise.

### GetDocumentBase64Ok

`func (o *IncomingInvoice) GetDocumentBase64Ok() (*string, bool)`

GetDocumentBase64Ok returns a tuple with the DocumentBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentBase64

`func (o *IncomingInvoice) SetDocumentBase64(v string)`

SetDocumentBase64 sets DocumentBase64 field to given value.

### HasDocumentBase64

`func (o *IncomingInvoice) HasDocumentBase64() bool`

HasDocumentBase64 returns a boolean if a field has been set.

### SetDocumentBase64Nil

`func (o *IncomingInvoice) SetDocumentBase64Nil(b bool)`

 SetDocumentBase64Nil sets the value for DocumentBase64 to be an explicit nil

### UnsetDocumentBase64
`func (o *IncomingInvoice) UnsetDocumentBase64()`

UnsetDocumentBase64 ensures that no value is present for DocumentBase64, not even an explicit nil
### GetDocumentContentType

`func (o *IncomingInvoice) GetDocumentContentType() string`

GetDocumentContentType returns the DocumentContentType field if non-nil, zero value otherwise.

### GetDocumentContentTypeOk

`func (o *IncomingInvoice) GetDocumentContentTypeOk() (*string, bool)`

GetDocumentContentTypeOk returns a tuple with the DocumentContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentContentType

`func (o *IncomingInvoice) SetDocumentContentType(v string)`

SetDocumentContentType sets DocumentContentType field to given value.

### HasDocumentContentType

`func (o *IncomingInvoice) HasDocumentContentType() bool`

HasDocumentContentType returns a boolean if a field has been set.

### SetDocumentContentTypeNil

`func (o *IncomingInvoice) SetDocumentContentTypeNil(b bool)`

 SetDocumentContentTypeNil sets the value for DocumentContentType to be an explicit nil

### UnsetDocumentContentType
`func (o *IncomingInvoice) UnsetDocumentContentType()`

UnsetDocumentContentType ensures that no value is present for DocumentContentType, not even an explicit nil
### GetDocumentFilename

`func (o *IncomingInvoice) GetDocumentFilename() string`

GetDocumentFilename returns the DocumentFilename field if non-nil, zero value otherwise.

### GetDocumentFilenameOk

`func (o *IncomingInvoice) GetDocumentFilenameOk() (*string, bool)`

GetDocumentFilenameOk returns a tuple with the DocumentFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentFilename

`func (o *IncomingInvoice) SetDocumentFilename(v string)`

SetDocumentFilename sets DocumentFilename field to given value.

### HasDocumentFilename

`func (o *IncomingInvoice) HasDocumentFilename() bool`

HasDocumentFilename returns a boolean if a field has been set.

### SetDocumentFilenameNil

`func (o *IncomingInvoice) SetDocumentFilenameNil(b bool)`

 SetDocumentFilenameNil sets the value for DocumentFilename to be an explicit nil

### UnsetDocumentFilename
`func (o *IncomingInvoice) UnsetDocumentFilename()`

UnsetDocumentFilename ensures that no value is present for DocumentFilename, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



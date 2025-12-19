# SubmitInvoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**NullableFactureElectroniqueRestApiSchemasChorusProChorusProCredentials**](FactureElectroniqueRestApiSchemasChorusProChorusProCredentials.md) |  | [optional] 
**InvoiceNumber** | **string** | Invoice number | 
**InvoiceDate** | **string** | Invoice date (ISO format: YYYY-MM-DD) | 
**PaymentDueDate** | Pointer to **NullableString** |  | [optional] 
**StructureId** | **int32** | Chorus Pro recipient structure ID | 
**ServiceCode** | Pointer to **NullableString** |  | [optional] 
**EngagementNumber** | Pointer to **NullableString** |  | [optional] 
**TotalNetAmount** | [**TotalNetAmount**](TotalNetAmount.md) |  | 
**VatAmount** | [**VatAmount**](VatAmount.md) |  | 
**TotalGrossAmount** | [**TotalGrossAmount**](TotalGrossAmount.md) |  | 
**MainAttachmentId** | Pointer to **NullableInt32** |  | [optional] 
**MainAttachmentLabel** | Pointer to **NullableString** |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**PurchaseOrderReference** | Pointer to **NullableString** |  | [optional] 
**ContractReference** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSubmitInvoiceRequest

`func NewSubmitInvoiceRequest(invoiceNumber string, invoiceDate string, structureId int32, totalNetAmount TotalNetAmount, vatAmount VatAmount, totalGrossAmount TotalGrossAmount, ) *SubmitInvoiceRequest`

NewSubmitInvoiceRequest instantiates a new SubmitInvoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitInvoiceRequestWithDefaults

`func NewSubmitInvoiceRequestWithDefaults() *SubmitInvoiceRequest`

NewSubmitInvoiceRequestWithDefaults instantiates a new SubmitInvoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *SubmitInvoiceRequest) GetCredentials() FactureElectroniqueRestApiSchemasChorusProChorusProCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *SubmitInvoiceRequest) GetCredentialsOk() (*FactureElectroniqueRestApiSchemasChorusProChorusProCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *SubmitInvoiceRequest) SetCredentials(v FactureElectroniqueRestApiSchemasChorusProChorusProCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *SubmitInvoiceRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *SubmitInvoiceRequest) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *SubmitInvoiceRequest) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetInvoiceNumber

`func (o *SubmitInvoiceRequest) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *SubmitInvoiceRequest) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *SubmitInvoiceRequest) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.


### GetInvoiceDate

`func (o *SubmitInvoiceRequest) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *SubmitInvoiceRequest) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *SubmitInvoiceRequest) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.


### GetPaymentDueDate

`func (o *SubmitInvoiceRequest) GetPaymentDueDate() string`

GetPaymentDueDate returns the PaymentDueDate field if non-nil, zero value otherwise.

### GetPaymentDueDateOk

`func (o *SubmitInvoiceRequest) GetPaymentDueDateOk() (*string, bool)`

GetPaymentDueDateOk returns a tuple with the PaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDueDate

`func (o *SubmitInvoiceRequest) SetPaymentDueDate(v string)`

SetPaymentDueDate sets PaymentDueDate field to given value.

### HasPaymentDueDate

`func (o *SubmitInvoiceRequest) HasPaymentDueDate() bool`

HasPaymentDueDate returns a boolean if a field has been set.

### SetPaymentDueDateNil

`func (o *SubmitInvoiceRequest) SetPaymentDueDateNil(b bool)`

 SetPaymentDueDateNil sets the value for PaymentDueDate to be an explicit nil

### UnsetPaymentDueDate
`func (o *SubmitInvoiceRequest) UnsetPaymentDueDate()`

UnsetPaymentDueDate ensures that no value is present for PaymentDueDate, not even an explicit nil
### GetStructureId

`func (o *SubmitInvoiceRequest) GetStructureId() int32`

GetStructureId returns the StructureId field if non-nil, zero value otherwise.

### GetStructureIdOk

`func (o *SubmitInvoiceRequest) GetStructureIdOk() (*int32, bool)`

GetStructureIdOk returns a tuple with the StructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureId

`func (o *SubmitInvoiceRequest) SetStructureId(v int32)`

SetStructureId sets StructureId field to given value.


### GetServiceCode

`func (o *SubmitInvoiceRequest) GetServiceCode() string`

GetServiceCode returns the ServiceCode field if non-nil, zero value otherwise.

### GetServiceCodeOk

`func (o *SubmitInvoiceRequest) GetServiceCodeOk() (*string, bool)`

GetServiceCodeOk returns a tuple with the ServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCode

`func (o *SubmitInvoiceRequest) SetServiceCode(v string)`

SetServiceCode sets ServiceCode field to given value.

### HasServiceCode

`func (o *SubmitInvoiceRequest) HasServiceCode() bool`

HasServiceCode returns a boolean if a field has been set.

### SetServiceCodeNil

`func (o *SubmitInvoiceRequest) SetServiceCodeNil(b bool)`

 SetServiceCodeNil sets the value for ServiceCode to be an explicit nil

### UnsetServiceCode
`func (o *SubmitInvoiceRequest) UnsetServiceCode()`

UnsetServiceCode ensures that no value is present for ServiceCode, not even an explicit nil
### GetEngagementNumber

`func (o *SubmitInvoiceRequest) GetEngagementNumber() string`

GetEngagementNumber returns the EngagementNumber field if non-nil, zero value otherwise.

### GetEngagementNumberOk

`func (o *SubmitInvoiceRequest) GetEngagementNumberOk() (*string, bool)`

GetEngagementNumberOk returns a tuple with the EngagementNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementNumber

`func (o *SubmitInvoiceRequest) SetEngagementNumber(v string)`

SetEngagementNumber sets EngagementNumber field to given value.

### HasEngagementNumber

`func (o *SubmitInvoiceRequest) HasEngagementNumber() bool`

HasEngagementNumber returns a boolean if a field has been set.

### SetEngagementNumberNil

`func (o *SubmitInvoiceRequest) SetEngagementNumberNil(b bool)`

 SetEngagementNumberNil sets the value for EngagementNumber to be an explicit nil

### UnsetEngagementNumber
`func (o *SubmitInvoiceRequest) UnsetEngagementNumber()`

UnsetEngagementNumber ensures that no value is present for EngagementNumber, not even an explicit nil
### GetTotalNetAmount

`func (o *SubmitInvoiceRequest) GetTotalNetAmount() TotalNetAmount`

GetTotalNetAmount returns the TotalNetAmount field if non-nil, zero value otherwise.

### GetTotalNetAmountOk

`func (o *SubmitInvoiceRequest) GetTotalNetAmountOk() (*TotalNetAmount, bool)`

GetTotalNetAmountOk returns a tuple with the TotalNetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNetAmount

`func (o *SubmitInvoiceRequest) SetTotalNetAmount(v TotalNetAmount)`

SetTotalNetAmount sets TotalNetAmount field to given value.


### GetVatAmount

`func (o *SubmitInvoiceRequest) GetVatAmount() VatAmount`

GetVatAmount returns the VatAmount field if non-nil, zero value otherwise.

### GetVatAmountOk

`func (o *SubmitInvoiceRequest) GetVatAmountOk() (*VatAmount, bool)`

GetVatAmountOk returns a tuple with the VatAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatAmount

`func (o *SubmitInvoiceRequest) SetVatAmount(v VatAmount)`

SetVatAmount sets VatAmount field to given value.


### GetTotalGrossAmount

`func (o *SubmitInvoiceRequest) GetTotalGrossAmount() TotalGrossAmount`

GetTotalGrossAmount returns the TotalGrossAmount field if non-nil, zero value otherwise.

### GetTotalGrossAmountOk

`func (o *SubmitInvoiceRequest) GetTotalGrossAmountOk() (*TotalGrossAmount, bool)`

GetTotalGrossAmountOk returns a tuple with the TotalGrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossAmount

`func (o *SubmitInvoiceRequest) SetTotalGrossAmount(v TotalGrossAmount)`

SetTotalGrossAmount sets TotalGrossAmount field to given value.


### GetMainAttachmentId

`func (o *SubmitInvoiceRequest) GetMainAttachmentId() int32`

GetMainAttachmentId returns the MainAttachmentId field if non-nil, zero value otherwise.

### GetMainAttachmentIdOk

`func (o *SubmitInvoiceRequest) GetMainAttachmentIdOk() (*int32, bool)`

GetMainAttachmentIdOk returns a tuple with the MainAttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainAttachmentId

`func (o *SubmitInvoiceRequest) SetMainAttachmentId(v int32)`

SetMainAttachmentId sets MainAttachmentId field to given value.

### HasMainAttachmentId

`func (o *SubmitInvoiceRequest) HasMainAttachmentId() bool`

HasMainAttachmentId returns a boolean if a field has been set.

### SetMainAttachmentIdNil

`func (o *SubmitInvoiceRequest) SetMainAttachmentIdNil(b bool)`

 SetMainAttachmentIdNil sets the value for MainAttachmentId to be an explicit nil

### UnsetMainAttachmentId
`func (o *SubmitInvoiceRequest) UnsetMainAttachmentId()`

UnsetMainAttachmentId ensures that no value is present for MainAttachmentId, not even an explicit nil
### GetMainAttachmentLabel

`func (o *SubmitInvoiceRequest) GetMainAttachmentLabel() string`

GetMainAttachmentLabel returns the MainAttachmentLabel field if non-nil, zero value otherwise.

### GetMainAttachmentLabelOk

`func (o *SubmitInvoiceRequest) GetMainAttachmentLabelOk() (*string, bool)`

GetMainAttachmentLabelOk returns a tuple with the MainAttachmentLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainAttachmentLabel

`func (o *SubmitInvoiceRequest) SetMainAttachmentLabel(v string)`

SetMainAttachmentLabel sets MainAttachmentLabel field to given value.

### HasMainAttachmentLabel

`func (o *SubmitInvoiceRequest) HasMainAttachmentLabel() bool`

HasMainAttachmentLabel returns a boolean if a field has been set.

### SetMainAttachmentLabelNil

`func (o *SubmitInvoiceRequest) SetMainAttachmentLabelNil(b bool)`

 SetMainAttachmentLabelNil sets the value for MainAttachmentLabel to be an explicit nil

### UnsetMainAttachmentLabel
`func (o *SubmitInvoiceRequest) UnsetMainAttachmentLabel()`

UnsetMainAttachmentLabel ensures that no value is present for MainAttachmentLabel, not even an explicit nil
### GetComment

`func (o *SubmitInvoiceRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SubmitInvoiceRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SubmitInvoiceRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SubmitInvoiceRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *SubmitInvoiceRequest) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *SubmitInvoiceRequest) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetPurchaseOrderReference

`func (o *SubmitInvoiceRequest) GetPurchaseOrderReference() string`

GetPurchaseOrderReference returns the PurchaseOrderReference field if non-nil, zero value otherwise.

### GetPurchaseOrderReferenceOk

`func (o *SubmitInvoiceRequest) GetPurchaseOrderReferenceOk() (*string, bool)`

GetPurchaseOrderReferenceOk returns a tuple with the PurchaseOrderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderReference

`func (o *SubmitInvoiceRequest) SetPurchaseOrderReference(v string)`

SetPurchaseOrderReference sets PurchaseOrderReference field to given value.

### HasPurchaseOrderReference

`func (o *SubmitInvoiceRequest) HasPurchaseOrderReference() bool`

HasPurchaseOrderReference returns a boolean if a field has been set.

### SetPurchaseOrderReferenceNil

`func (o *SubmitInvoiceRequest) SetPurchaseOrderReferenceNil(b bool)`

 SetPurchaseOrderReferenceNil sets the value for PurchaseOrderReference to be an explicit nil

### UnsetPurchaseOrderReference
`func (o *SubmitInvoiceRequest) UnsetPurchaseOrderReference()`

UnsetPurchaseOrderReference ensures that no value is present for PurchaseOrderReference, not even an explicit nil
### GetContractReference

`func (o *SubmitInvoiceRequest) GetContractReference() string`

GetContractReference returns the ContractReference field if non-nil, zero value otherwise.

### GetContractReferenceOk

`func (o *SubmitInvoiceRequest) GetContractReferenceOk() (*string, bool)`

GetContractReferenceOk returns a tuple with the ContractReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractReference

`func (o *SubmitInvoiceRequest) SetContractReference(v string)`

SetContractReference sets ContractReference field to given value.

### HasContractReference

`func (o *SubmitInvoiceRequest) HasContractReference() bool`

HasContractReference returns a boolean if a field has been set.

### SetContractReferenceNil

`func (o *SubmitInvoiceRequest) SetContractReferenceNil(b bool)`

 SetContractReferenceNil sets the value for ContractReference to be an explicit nil

### UnsetContractReference
`func (o *SubmitInvoiceRequest) UnsetContractReference()`

UnsetContractReference ensures that no value is present for ContractReference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



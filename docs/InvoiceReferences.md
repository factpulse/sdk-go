# InvoiceReferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessProcessId** | Pointer to **NullableString** |  | [optional] 
**InvoiceCurrency** | Pointer to **string** | Invoice currency code (BT-5). ISO 4217. | [optional] [default to "EUR"]
**PaymentMeans** | [**PaymentMeans**](PaymentMeans.md) | Payment means type code (BT-81). | 
**PaymentMeansText** | Pointer to **NullableString** |  | [optional] 
**InvoiceType** | [**InvoiceTypeCode**](InvoiceTypeCode.md) |  | 
**VatAccountingCode** | [**VATAccountingCode**](VATAccountingCode.md) | VAT accounting code. | 
**BuyerReference** | Pointer to **NullableString** |  | [optional] 
**ContractReference** | Pointer to **NullableString** |  | [optional] 
**PurchaseOrderReference** | Pointer to **NullableString** |  | [optional] 
**SellerOrderReference** | Pointer to **NullableString** |  | [optional] 
**ReceivingAdviceReference** | Pointer to **NullableString** |  | [optional] 
**DespatchAdviceReference** | Pointer to **NullableString** |  | [optional] 
**TenderReference** | Pointer to **NullableString** |  | [optional] 
**PrecedingInvoiceReference** | Pointer to **NullableString** |  | [optional] 
**PrecedingInvoiceDate** | Pointer to **NullableString** |  | [optional] 
**ProjectReference** | Pointer to **NullableString** |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**VatExemptionReason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInvoiceReferences

`func NewInvoiceReferences(paymentMeans PaymentMeans, invoiceType InvoiceTypeCode, vatAccountingCode VATAccountingCode, ) *InvoiceReferences`

NewInvoiceReferences instantiates a new InvoiceReferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceReferencesWithDefaults

`func NewInvoiceReferencesWithDefaults() *InvoiceReferences`

NewInvoiceReferencesWithDefaults instantiates a new InvoiceReferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessProcessId

`func (o *InvoiceReferences) GetBusinessProcessId() string`

GetBusinessProcessId returns the BusinessProcessId field if non-nil, zero value otherwise.

### GetBusinessProcessIdOk

`func (o *InvoiceReferences) GetBusinessProcessIdOk() (*string, bool)`

GetBusinessProcessIdOk returns a tuple with the BusinessProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProcessId

`func (o *InvoiceReferences) SetBusinessProcessId(v string)`

SetBusinessProcessId sets BusinessProcessId field to given value.

### HasBusinessProcessId

`func (o *InvoiceReferences) HasBusinessProcessId() bool`

HasBusinessProcessId returns a boolean if a field has been set.

### SetBusinessProcessIdNil

`func (o *InvoiceReferences) SetBusinessProcessIdNil(b bool)`

 SetBusinessProcessIdNil sets the value for BusinessProcessId to be an explicit nil

### UnsetBusinessProcessId
`func (o *InvoiceReferences) UnsetBusinessProcessId()`

UnsetBusinessProcessId ensures that no value is present for BusinessProcessId, not even an explicit nil
### GetInvoiceCurrency

`func (o *InvoiceReferences) GetInvoiceCurrency() string`

GetInvoiceCurrency returns the InvoiceCurrency field if non-nil, zero value otherwise.

### GetInvoiceCurrencyOk

`func (o *InvoiceReferences) GetInvoiceCurrencyOk() (*string, bool)`

GetInvoiceCurrencyOk returns a tuple with the InvoiceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCurrency

`func (o *InvoiceReferences) SetInvoiceCurrency(v string)`

SetInvoiceCurrency sets InvoiceCurrency field to given value.

### HasInvoiceCurrency

`func (o *InvoiceReferences) HasInvoiceCurrency() bool`

HasInvoiceCurrency returns a boolean if a field has been set.

### GetPaymentMeans

`func (o *InvoiceReferences) GetPaymentMeans() PaymentMeans`

GetPaymentMeans returns the PaymentMeans field if non-nil, zero value otherwise.

### GetPaymentMeansOk

`func (o *InvoiceReferences) GetPaymentMeansOk() (*PaymentMeans, bool)`

GetPaymentMeansOk returns a tuple with the PaymentMeans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMeans

`func (o *InvoiceReferences) SetPaymentMeans(v PaymentMeans)`

SetPaymentMeans sets PaymentMeans field to given value.


### GetPaymentMeansText

`func (o *InvoiceReferences) GetPaymentMeansText() string`

GetPaymentMeansText returns the PaymentMeansText field if non-nil, zero value otherwise.

### GetPaymentMeansTextOk

`func (o *InvoiceReferences) GetPaymentMeansTextOk() (*string, bool)`

GetPaymentMeansTextOk returns a tuple with the PaymentMeansText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMeansText

`func (o *InvoiceReferences) SetPaymentMeansText(v string)`

SetPaymentMeansText sets PaymentMeansText field to given value.

### HasPaymentMeansText

`func (o *InvoiceReferences) HasPaymentMeansText() bool`

HasPaymentMeansText returns a boolean if a field has been set.

### SetPaymentMeansTextNil

`func (o *InvoiceReferences) SetPaymentMeansTextNil(b bool)`

 SetPaymentMeansTextNil sets the value for PaymentMeansText to be an explicit nil

### UnsetPaymentMeansText
`func (o *InvoiceReferences) UnsetPaymentMeansText()`

UnsetPaymentMeansText ensures that no value is present for PaymentMeansText, not even an explicit nil
### GetInvoiceType

`func (o *InvoiceReferences) GetInvoiceType() InvoiceTypeCode`

GetInvoiceType returns the InvoiceType field if non-nil, zero value otherwise.

### GetInvoiceTypeOk

`func (o *InvoiceReferences) GetInvoiceTypeOk() (*InvoiceTypeCode, bool)`

GetInvoiceTypeOk returns a tuple with the InvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceType

`func (o *InvoiceReferences) SetInvoiceType(v InvoiceTypeCode)`

SetInvoiceType sets InvoiceType field to given value.


### GetVatAccountingCode

`func (o *InvoiceReferences) GetVatAccountingCode() VATAccountingCode`

GetVatAccountingCode returns the VatAccountingCode field if non-nil, zero value otherwise.

### GetVatAccountingCodeOk

`func (o *InvoiceReferences) GetVatAccountingCodeOk() (*VATAccountingCode, bool)`

GetVatAccountingCodeOk returns a tuple with the VatAccountingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatAccountingCode

`func (o *InvoiceReferences) SetVatAccountingCode(v VATAccountingCode)`

SetVatAccountingCode sets VatAccountingCode field to given value.


### GetBuyerReference

`func (o *InvoiceReferences) GetBuyerReference() string`

GetBuyerReference returns the BuyerReference field if non-nil, zero value otherwise.

### GetBuyerReferenceOk

`func (o *InvoiceReferences) GetBuyerReferenceOk() (*string, bool)`

GetBuyerReferenceOk returns a tuple with the BuyerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerReference

`func (o *InvoiceReferences) SetBuyerReference(v string)`

SetBuyerReference sets BuyerReference field to given value.

### HasBuyerReference

`func (o *InvoiceReferences) HasBuyerReference() bool`

HasBuyerReference returns a boolean if a field has been set.

### SetBuyerReferenceNil

`func (o *InvoiceReferences) SetBuyerReferenceNil(b bool)`

 SetBuyerReferenceNil sets the value for BuyerReference to be an explicit nil

### UnsetBuyerReference
`func (o *InvoiceReferences) UnsetBuyerReference()`

UnsetBuyerReference ensures that no value is present for BuyerReference, not even an explicit nil
### GetContractReference

`func (o *InvoiceReferences) GetContractReference() string`

GetContractReference returns the ContractReference field if non-nil, zero value otherwise.

### GetContractReferenceOk

`func (o *InvoiceReferences) GetContractReferenceOk() (*string, bool)`

GetContractReferenceOk returns a tuple with the ContractReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractReference

`func (o *InvoiceReferences) SetContractReference(v string)`

SetContractReference sets ContractReference field to given value.

### HasContractReference

`func (o *InvoiceReferences) HasContractReference() bool`

HasContractReference returns a boolean if a field has been set.

### SetContractReferenceNil

`func (o *InvoiceReferences) SetContractReferenceNil(b bool)`

 SetContractReferenceNil sets the value for ContractReference to be an explicit nil

### UnsetContractReference
`func (o *InvoiceReferences) UnsetContractReference()`

UnsetContractReference ensures that no value is present for ContractReference, not even an explicit nil
### GetPurchaseOrderReference

`func (o *InvoiceReferences) GetPurchaseOrderReference() string`

GetPurchaseOrderReference returns the PurchaseOrderReference field if non-nil, zero value otherwise.

### GetPurchaseOrderReferenceOk

`func (o *InvoiceReferences) GetPurchaseOrderReferenceOk() (*string, bool)`

GetPurchaseOrderReferenceOk returns a tuple with the PurchaseOrderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderReference

`func (o *InvoiceReferences) SetPurchaseOrderReference(v string)`

SetPurchaseOrderReference sets PurchaseOrderReference field to given value.

### HasPurchaseOrderReference

`func (o *InvoiceReferences) HasPurchaseOrderReference() bool`

HasPurchaseOrderReference returns a boolean if a field has been set.

### SetPurchaseOrderReferenceNil

`func (o *InvoiceReferences) SetPurchaseOrderReferenceNil(b bool)`

 SetPurchaseOrderReferenceNil sets the value for PurchaseOrderReference to be an explicit nil

### UnsetPurchaseOrderReference
`func (o *InvoiceReferences) UnsetPurchaseOrderReference()`

UnsetPurchaseOrderReference ensures that no value is present for PurchaseOrderReference, not even an explicit nil
### GetSellerOrderReference

`func (o *InvoiceReferences) GetSellerOrderReference() string`

GetSellerOrderReference returns the SellerOrderReference field if non-nil, zero value otherwise.

### GetSellerOrderReferenceOk

`func (o *InvoiceReferences) GetSellerOrderReferenceOk() (*string, bool)`

GetSellerOrderReferenceOk returns a tuple with the SellerOrderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerOrderReference

`func (o *InvoiceReferences) SetSellerOrderReference(v string)`

SetSellerOrderReference sets SellerOrderReference field to given value.

### HasSellerOrderReference

`func (o *InvoiceReferences) HasSellerOrderReference() bool`

HasSellerOrderReference returns a boolean if a field has been set.

### SetSellerOrderReferenceNil

`func (o *InvoiceReferences) SetSellerOrderReferenceNil(b bool)`

 SetSellerOrderReferenceNil sets the value for SellerOrderReference to be an explicit nil

### UnsetSellerOrderReference
`func (o *InvoiceReferences) UnsetSellerOrderReference()`

UnsetSellerOrderReference ensures that no value is present for SellerOrderReference, not even an explicit nil
### GetReceivingAdviceReference

`func (o *InvoiceReferences) GetReceivingAdviceReference() string`

GetReceivingAdviceReference returns the ReceivingAdviceReference field if non-nil, zero value otherwise.

### GetReceivingAdviceReferenceOk

`func (o *InvoiceReferences) GetReceivingAdviceReferenceOk() (*string, bool)`

GetReceivingAdviceReferenceOk returns a tuple with the ReceivingAdviceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAdviceReference

`func (o *InvoiceReferences) SetReceivingAdviceReference(v string)`

SetReceivingAdviceReference sets ReceivingAdviceReference field to given value.

### HasReceivingAdviceReference

`func (o *InvoiceReferences) HasReceivingAdviceReference() bool`

HasReceivingAdviceReference returns a boolean if a field has been set.

### SetReceivingAdviceReferenceNil

`func (o *InvoiceReferences) SetReceivingAdviceReferenceNil(b bool)`

 SetReceivingAdviceReferenceNil sets the value for ReceivingAdviceReference to be an explicit nil

### UnsetReceivingAdviceReference
`func (o *InvoiceReferences) UnsetReceivingAdviceReference()`

UnsetReceivingAdviceReference ensures that no value is present for ReceivingAdviceReference, not even an explicit nil
### GetDespatchAdviceReference

`func (o *InvoiceReferences) GetDespatchAdviceReference() string`

GetDespatchAdviceReference returns the DespatchAdviceReference field if non-nil, zero value otherwise.

### GetDespatchAdviceReferenceOk

`func (o *InvoiceReferences) GetDespatchAdviceReferenceOk() (*string, bool)`

GetDespatchAdviceReferenceOk returns a tuple with the DespatchAdviceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDespatchAdviceReference

`func (o *InvoiceReferences) SetDespatchAdviceReference(v string)`

SetDespatchAdviceReference sets DespatchAdviceReference field to given value.

### HasDespatchAdviceReference

`func (o *InvoiceReferences) HasDespatchAdviceReference() bool`

HasDespatchAdviceReference returns a boolean if a field has been set.

### SetDespatchAdviceReferenceNil

`func (o *InvoiceReferences) SetDespatchAdviceReferenceNil(b bool)`

 SetDespatchAdviceReferenceNil sets the value for DespatchAdviceReference to be an explicit nil

### UnsetDespatchAdviceReference
`func (o *InvoiceReferences) UnsetDespatchAdviceReference()`

UnsetDespatchAdviceReference ensures that no value is present for DespatchAdviceReference, not even an explicit nil
### GetTenderReference

`func (o *InvoiceReferences) GetTenderReference() string`

GetTenderReference returns the TenderReference field if non-nil, zero value otherwise.

### GetTenderReferenceOk

`func (o *InvoiceReferences) GetTenderReferenceOk() (*string, bool)`

GetTenderReferenceOk returns a tuple with the TenderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenderReference

`func (o *InvoiceReferences) SetTenderReference(v string)`

SetTenderReference sets TenderReference field to given value.

### HasTenderReference

`func (o *InvoiceReferences) HasTenderReference() bool`

HasTenderReference returns a boolean if a field has been set.

### SetTenderReferenceNil

`func (o *InvoiceReferences) SetTenderReferenceNil(b bool)`

 SetTenderReferenceNil sets the value for TenderReference to be an explicit nil

### UnsetTenderReference
`func (o *InvoiceReferences) UnsetTenderReference()`

UnsetTenderReference ensures that no value is present for TenderReference, not even an explicit nil
### GetPrecedingInvoiceReference

`func (o *InvoiceReferences) GetPrecedingInvoiceReference() string`

GetPrecedingInvoiceReference returns the PrecedingInvoiceReference field if non-nil, zero value otherwise.

### GetPrecedingInvoiceReferenceOk

`func (o *InvoiceReferences) GetPrecedingInvoiceReferenceOk() (*string, bool)`

GetPrecedingInvoiceReferenceOk returns a tuple with the PrecedingInvoiceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedingInvoiceReference

`func (o *InvoiceReferences) SetPrecedingInvoiceReference(v string)`

SetPrecedingInvoiceReference sets PrecedingInvoiceReference field to given value.

### HasPrecedingInvoiceReference

`func (o *InvoiceReferences) HasPrecedingInvoiceReference() bool`

HasPrecedingInvoiceReference returns a boolean if a field has been set.

### SetPrecedingInvoiceReferenceNil

`func (o *InvoiceReferences) SetPrecedingInvoiceReferenceNil(b bool)`

 SetPrecedingInvoiceReferenceNil sets the value for PrecedingInvoiceReference to be an explicit nil

### UnsetPrecedingInvoiceReference
`func (o *InvoiceReferences) UnsetPrecedingInvoiceReference()`

UnsetPrecedingInvoiceReference ensures that no value is present for PrecedingInvoiceReference, not even an explicit nil
### GetPrecedingInvoiceDate

`func (o *InvoiceReferences) GetPrecedingInvoiceDate() string`

GetPrecedingInvoiceDate returns the PrecedingInvoiceDate field if non-nil, zero value otherwise.

### GetPrecedingInvoiceDateOk

`func (o *InvoiceReferences) GetPrecedingInvoiceDateOk() (*string, bool)`

GetPrecedingInvoiceDateOk returns a tuple with the PrecedingInvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedingInvoiceDate

`func (o *InvoiceReferences) SetPrecedingInvoiceDate(v string)`

SetPrecedingInvoiceDate sets PrecedingInvoiceDate field to given value.

### HasPrecedingInvoiceDate

`func (o *InvoiceReferences) HasPrecedingInvoiceDate() bool`

HasPrecedingInvoiceDate returns a boolean if a field has been set.

### SetPrecedingInvoiceDateNil

`func (o *InvoiceReferences) SetPrecedingInvoiceDateNil(b bool)`

 SetPrecedingInvoiceDateNil sets the value for PrecedingInvoiceDate to be an explicit nil

### UnsetPrecedingInvoiceDate
`func (o *InvoiceReferences) UnsetPrecedingInvoiceDate()`

UnsetPrecedingInvoiceDate ensures that no value is present for PrecedingInvoiceDate, not even an explicit nil
### GetProjectReference

`func (o *InvoiceReferences) GetProjectReference() string`

GetProjectReference returns the ProjectReference field if non-nil, zero value otherwise.

### GetProjectReferenceOk

`func (o *InvoiceReferences) GetProjectReferenceOk() (*string, bool)`

GetProjectReferenceOk returns a tuple with the ProjectReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectReference

`func (o *InvoiceReferences) SetProjectReference(v string)`

SetProjectReference sets ProjectReference field to given value.

### HasProjectReference

`func (o *InvoiceReferences) HasProjectReference() bool`

HasProjectReference returns a boolean if a field has been set.

### SetProjectReferenceNil

`func (o *InvoiceReferences) SetProjectReferenceNil(b bool)`

 SetProjectReferenceNil sets the value for ProjectReference to be an explicit nil

### UnsetProjectReference
`func (o *InvoiceReferences) UnsetProjectReference()`

UnsetProjectReference ensures that no value is present for ProjectReference, not even an explicit nil
### GetProjectName

`func (o *InvoiceReferences) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *InvoiceReferences) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *InvoiceReferences) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *InvoiceReferences) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *InvoiceReferences) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *InvoiceReferences) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetVatExemptionReason

`func (o *InvoiceReferences) GetVatExemptionReason() string`

GetVatExemptionReason returns the VatExemptionReason field if non-nil, zero value otherwise.

### GetVatExemptionReasonOk

`func (o *InvoiceReferences) GetVatExemptionReasonOk() (*string, bool)`

GetVatExemptionReasonOk returns a tuple with the VatExemptionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatExemptionReason

`func (o *InvoiceReferences) SetVatExemptionReason(v string)`

SetVatExemptionReason sets VatExemptionReason field to given value.

### HasVatExemptionReason

`func (o *InvoiceReferences) HasVatExemptionReason() bool`

HasVatExemptionReason returns a boolean if a field has been set.

### SetVatExemptionReasonNil

`func (o *InvoiceReferences) SetVatExemptionReasonNil(b bool)`

 SetVatExemptionReasonNil sets the value for VatExemptionReason to be an explicit nil

### UnsetVatExemptionReason
`func (o *InvoiceReferences) UnsetVatExemptionReason()`

UnsetVatExemptionReason ensures that no value is present for VatExemptionReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



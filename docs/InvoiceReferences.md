# InvoiceReferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceCurrency** | Pointer to **string** |  | [optional] [default to "EUR"]
**PaymentMeans** | [**PaymentMeans**](PaymentMeans.md) |  | 
**InvoiceType** | [**InvoiceTypeCode**](InvoiceTypeCode.md) |  | 
**VatAccountingCode** | [**VATAccountingCode**](VATAccountingCode.md) |  | 
**ContractReference** | Pointer to **NullableString** |  | [optional] 
**VatExemptionReason** | Pointer to **NullableString** |  | [optional] 
**PurchaseOrderReference** | Pointer to **NullableString** |  | [optional] 
**PrecedingInvoiceReference** | Pointer to **NullableString** |  | [optional] 

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



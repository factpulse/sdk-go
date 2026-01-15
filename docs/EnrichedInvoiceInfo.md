# EnrichedInvoiceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceNumber** | **string** | Invoice number | 
**SupplierId** | Pointer to **NullableInt32** |  | [optional] 
**RecipientId** | Pointer to **NullableInt32** |  | [optional] 
**SupplierName** | **string** | Supplier name | 
**RecipientName** | **string** | Recipient name | 
**TotalNetAmount** | **string** | Total net amount (HT) | 
**VatAmount** | **string** | VAT amount | 
**TotalGrossAmount** | **string** | Total gross amount (TTC) | 

## Methods

### NewEnrichedInvoiceInfo

`func NewEnrichedInvoiceInfo(invoiceNumber string, supplierName string, recipientName string, totalNetAmount string, vatAmount string, totalGrossAmount string, ) *EnrichedInvoiceInfo`

NewEnrichedInvoiceInfo instantiates a new EnrichedInvoiceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrichedInvoiceInfoWithDefaults

`func NewEnrichedInvoiceInfoWithDefaults() *EnrichedInvoiceInfo`

NewEnrichedInvoiceInfoWithDefaults instantiates a new EnrichedInvoiceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceNumber

`func (o *EnrichedInvoiceInfo) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *EnrichedInvoiceInfo) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *EnrichedInvoiceInfo) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.


### GetSupplierId

`func (o *EnrichedInvoiceInfo) GetSupplierId() int32`

GetSupplierId returns the SupplierId field if non-nil, zero value otherwise.

### GetSupplierIdOk

`func (o *EnrichedInvoiceInfo) GetSupplierIdOk() (*int32, bool)`

GetSupplierIdOk returns a tuple with the SupplierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierId

`func (o *EnrichedInvoiceInfo) SetSupplierId(v int32)`

SetSupplierId sets SupplierId field to given value.

### HasSupplierId

`func (o *EnrichedInvoiceInfo) HasSupplierId() bool`

HasSupplierId returns a boolean if a field has been set.

### SetSupplierIdNil

`func (o *EnrichedInvoiceInfo) SetSupplierIdNil(b bool)`

 SetSupplierIdNil sets the value for SupplierId to be an explicit nil

### UnsetSupplierId
`func (o *EnrichedInvoiceInfo) UnsetSupplierId()`

UnsetSupplierId ensures that no value is present for SupplierId, not even an explicit nil
### GetRecipientId

`func (o *EnrichedInvoiceInfo) GetRecipientId() int32`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *EnrichedInvoiceInfo) GetRecipientIdOk() (*int32, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *EnrichedInvoiceInfo) SetRecipientId(v int32)`

SetRecipientId sets RecipientId field to given value.

### HasRecipientId

`func (o *EnrichedInvoiceInfo) HasRecipientId() bool`

HasRecipientId returns a boolean if a field has been set.

### SetRecipientIdNil

`func (o *EnrichedInvoiceInfo) SetRecipientIdNil(b bool)`

 SetRecipientIdNil sets the value for RecipientId to be an explicit nil

### UnsetRecipientId
`func (o *EnrichedInvoiceInfo) UnsetRecipientId()`

UnsetRecipientId ensures that no value is present for RecipientId, not even an explicit nil
### GetSupplierName

`func (o *EnrichedInvoiceInfo) GetSupplierName() string`

GetSupplierName returns the SupplierName field if non-nil, zero value otherwise.

### GetSupplierNameOk

`func (o *EnrichedInvoiceInfo) GetSupplierNameOk() (*string, bool)`

GetSupplierNameOk returns a tuple with the SupplierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierName

`func (o *EnrichedInvoiceInfo) SetSupplierName(v string)`

SetSupplierName sets SupplierName field to given value.


### GetRecipientName

`func (o *EnrichedInvoiceInfo) GetRecipientName() string`

GetRecipientName returns the RecipientName field if non-nil, zero value otherwise.

### GetRecipientNameOk

`func (o *EnrichedInvoiceInfo) GetRecipientNameOk() (*string, bool)`

GetRecipientNameOk returns a tuple with the RecipientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientName

`func (o *EnrichedInvoiceInfo) SetRecipientName(v string)`

SetRecipientName sets RecipientName field to given value.


### GetTotalNetAmount

`func (o *EnrichedInvoiceInfo) GetTotalNetAmount() string`

GetTotalNetAmount returns the TotalNetAmount field if non-nil, zero value otherwise.

### GetTotalNetAmountOk

`func (o *EnrichedInvoiceInfo) GetTotalNetAmountOk() (*string, bool)`

GetTotalNetAmountOk returns a tuple with the TotalNetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNetAmount

`func (o *EnrichedInvoiceInfo) SetTotalNetAmount(v string)`

SetTotalNetAmount sets TotalNetAmount field to given value.


### GetVatAmount

`func (o *EnrichedInvoiceInfo) GetVatAmount() string`

GetVatAmount returns the VatAmount field if non-nil, zero value otherwise.

### GetVatAmountOk

`func (o *EnrichedInvoiceInfo) GetVatAmountOk() (*string, bool)`

GetVatAmountOk returns a tuple with the VatAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatAmount

`func (o *EnrichedInvoiceInfo) SetVatAmount(v string)`

SetVatAmount sets VatAmount field to given value.


### GetTotalGrossAmount

`func (o *EnrichedInvoiceInfo) GetTotalGrossAmount() string`

GetTotalGrossAmount returns the TotalGrossAmount field if non-nil, zero value otherwise.

### GetTotalGrossAmountOk

`func (o *EnrichedInvoiceInfo) GetTotalGrossAmountOk() (*string, bool)`

GetTotalGrossAmountOk returns a tuple with the TotalGrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossAmount

`func (o *EnrichedInvoiceInfo) SetTotalGrossAmount(v string)`

SetTotalGrossAmount sets TotalGrossAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InvoiceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** | Invoice identifier | 
**IssueDate** | **string** | Invoice issue date | 
**TypeCode** | Pointer to [**InvoiceTypeCode**](InvoiceTypeCode.md) | Invoice type code | [optional] [default to _380]
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] [default to EUR]
**DueDate** | Pointer to **NullableString** |  | [optional] 
**SellerSiren** | **string** | Seller SIREN/SIRET | 
**SellerVatId** | Pointer to **NullableString** |  | [optional] 
**SellerCountry** | Pointer to [**Sellercountry**](Sellercountry.md) |  | [optional] [default to FR]
**BuyerId** | Pointer to **NullableString** |  | [optional] 
**BuyerVatId** | Pointer to **NullableString** |  | [optional] 
**BuyerCountry** | [**Buyercountry**](Buyercountry.md) |  | 
**TaxExclusiveAmount** | [**Taxexclusiveamount1**](Taxexclusiveamount1.md) |  | 
**TaxAmount** | [**Taxamount1**](Taxamount1.md) |  | 
**TaxBreakdown** | [**[]TaxBreakdownInput**](TaxBreakdownInput.md) | VAT breakdown | 
**ReferencedInvoiceId** | Pointer to **NullableString** |  | [optional] 
**ReferencedInvoiceDate** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInvoiceInput

`func NewInvoiceInput(invoiceId string, issueDate string, sellerSiren string, buyerCountry Buyercountry, taxExclusiveAmount Taxexclusiveamount1, taxAmount Taxamount1, taxBreakdown []TaxBreakdownInput, ) *InvoiceInput`

NewInvoiceInput instantiates a new InvoiceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceInputWithDefaults

`func NewInvoiceInputWithDefaults() *InvoiceInput`

NewInvoiceInputWithDefaults instantiates a new InvoiceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *InvoiceInput) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *InvoiceInput) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *InvoiceInput) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetIssueDate

`func (o *InvoiceInput) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *InvoiceInput) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *InvoiceInput) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetTypeCode

`func (o *InvoiceInput) GetTypeCode() InvoiceTypeCode`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *InvoiceInput) GetTypeCodeOk() (*InvoiceTypeCode, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *InvoiceInput) SetTypeCode(v InvoiceTypeCode)`

SetTypeCode sets TypeCode field to given value.

### HasTypeCode

`func (o *InvoiceInput) HasTypeCode() bool`

HasTypeCode returns a boolean if a field has been set.

### GetCurrency

`func (o *InvoiceInput) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceInput) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceInput) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InvoiceInput) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDueDate

`func (o *InvoiceInput) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoiceInput) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoiceInput) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InvoiceInput) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *InvoiceInput) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *InvoiceInput) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetSellerSiren

`func (o *InvoiceInput) GetSellerSiren() string`

GetSellerSiren returns the SellerSiren field if non-nil, zero value otherwise.

### GetSellerSirenOk

`func (o *InvoiceInput) GetSellerSirenOk() (*string, bool)`

GetSellerSirenOk returns a tuple with the SellerSiren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerSiren

`func (o *InvoiceInput) SetSellerSiren(v string)`

SetSellerSiren sets SellerSiren field to given value.


### GetSellerVatId

`func (o *InvoiceInput) GetSellerVatId() string`

GetSellerVatId returns the SellerVatId field if non-nil, zero value otherwise.

### GetSellerVatIdOk

`func (o *InvoiceInput) GetSellerVatIdOk() (*string, bool)`

GetSellerVatIdOk returns a tuple with the SellerVatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerVatId

`func (o *InvoiceInput) SetSellerVatId(v string)`

SetSellerVatId sets SellerVatId field to given value.

### HasSellerVatId

`func (o *InvoiceInput) HasSellerVatId() bool`

HasSellerVatId returns a boolean if a field has been set.

### SetSellerVatIdNil

`func (o *InvoiceInput) SetSellerVatIdNil(b bool)`

 SetSellerVatIdNil sets the value for SellerVatId to be an explicit nil

### UnsetSellerVatId
`func (o *InvoiceInput) UnsetSellerVatId()`

UnsetSellerVatId ensures that no value is present for SellerVatId, not even an explicit nil
### GetSellerCountry

`func (o *InvoiceInput) GetSellerCountry() Sellercountry`

GetSellerCountry returns the SellerCountry field if non-nil, zero value otherwise.

### GetSellerCountryOk

`func (o *InvoiceInput) GetSellerCountryOk() (*Sellercountry, bool)`

GetSellerCountryOk returns a tuple with the SellerCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerCountry

`func (o *InvoiceInput) SetSellerCountry(v Sellercountry)`

SetSellerCountry sets SellerCountry field to given value.

### HasSellerCountry

`func (o *InvoiceInput) HasSellerCountry() bool`

HasSellerCountry returns a boolean if a field has been set.

### GetBuyerId

`func (o *InvoiceInput) GetBuyerId() string`

GetBuyerId returns the BuyerId field if non-nil, zero value otherwise.

### GetBuyerIdOk

`func (o *InvoiceInput) GetBuyerIdOk() (*string, bool)`

GetBuyerIdOk returns a tuple with the BuyerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerId

`func (o *InvoiceInput) SetBuyerId(v string)`

SetBuyerId sets BuyerId field to given value.

### HasBuyerId

`func (o *InvoiceInput) HasBuyerId() bool`

HasBuyerId returns a boolean if a field has been set.

### SetBuyerIdNil

`func (o *InvoiceInput) SetBuyerIdNil(b bool)`

 SetBuyerIdNil sets the value for BuyerId to be an explicit nil

### UnsetBuyerId
`func (o *InvoiceInput) UnsetBuyerId()`

UnsetBuyerId ensures that no value is present for BuyerId, not even an explicit nil
### GetBuyerVatId

`func (o *InvoiceInput) GetBuyerVatId() string`

GetBuyerVatId returns the BuyerVatId field if non-nil, zero value otherwise.

### GetBuyerVatIdOk

`func (o *InvoiceInput) GetBuyerVatIdOk() (*string, bool)`

GetBuyerVatIdOk returns a tuple with the BuyerVatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerVatId

`func (o *InvoiceInput) SetBuyerVatId(v string)`

SetBuyerVatId sets BuyerVatId field to given value.

### HasBuyerVatId

`func (o *InvoiceInput) HasBuyerVatId() bool`

HasBuyerVatId returns a boolean if a field has been set.

### SetBuyerVatIdNil

`func (o *InvoiceInput) SetBuyerVatIdNil(b bool)`

 SetBuyerVatIdNil sets the value for BuyerVatId to be an explicit nil

### UnsetBuyerVatId
`func (o *InvoiceInput) UnsetBuyerVatId()`

UnsetBuyerVatId ensures that no value is present for BuyerVatId, not even an explicit nil
### GetBuyerCountry

`func (o *InvoiceInput) GetBuyerCountry() Buyercountry`

GetBuyerCountry returns the BuyerCountry field if non-nil, zero value otherwise.

### GetBuyerCountryOk

`func (o *InvoiceInput) GetBuyerCountryOk() (*Buyercountry, bool)`

GetBuyerCountryOk returns a tuple with the BuyerCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerCountry

`func (o *InvoiceInput) SetBuyerCountry(v Buyercountry)`

SetBuyerCountry sets BuyerCountry field to given value.


### GetTaxExclusiveAmount

`func (o *InvoiceInput) GetTaxExclusiveAmount() Taxexclusiveamount1`

GetTaxExclusiveAmount returns the TaxExclusiveAmount field if non-nil, zero value otherwise.

### GetTaxExclusiveAmountOk

`func (o *InvoiceInput) GetTaxExclusiveAmountOk() (*Taxexclusiveamount1, bool)`

GetTaxExclusiveAmountOk returns a tuple with the TaxExclusiveAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxExclusiveAmount

`func (o *InvoiceInput) SetTaxExclusiveAmount(v Taxexclusiveamount1)`

SetTaxExclusiveAmount sets TaxExclusiveAmount field to given value.


### GetTaxAmount

`func (o *InvoiceInput) GetTaxAmount() Taxamount1`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *InvoiceInput) GetTaxAmountOk() (*Taxamount1, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *InvoiceInput) SetTaxAmount(v Taxamount1)`

SetTaxAmount sets TaxAmount field to given value.


### GetTaxBreakdown

`func (o *InvoiceInput) GetTaxBreakdown() []TaxBreakdownInput`

GetTaxBreakdown returns the TaxBreakdown field if non-nil, zero value otherwise.

### GetTaxBreakdownOk

`func (o *InvoiceInput) GetTaxBreakdownOk() (*[]TaxBreakdownInput, bool)`

GetTaxBreakdownOk returns a tuple with the TaxBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxBreakdown

`func (o *InvoiceInput) SetTaxBreakdown(v []TaxBreakdownInput)`

SetTaxBreakdown sets TaxBreakdown field to given value.


### GetReferencedInvoiceId

`func (o *InvoiceInput) GetReferencedInvoiceId() string`

GetReferencedInvoiceId returns the ReferencedInvoiceId field if non-nil, zero value otherwise.

### GetReferencedInvoiceIdOk

`func (o *InvoiceInput) GetReferencedInvoiceIdOk() (*string, bool)`

GetReferencedInvoiceIdOk returns a tuple with the ReferencedInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedInvoiceId

`func (o *InvoiceInput) SetReferencedInvoiceId(v string)`

SetReferencedInvoiceId sets ReferencedInvoiceId field to given value.

### HasReferencedInvoiceId

`func (o *InvoiceInput) HasReferencedInvoiceId() bool`

HasReferencedInvoiceId returns a boolean if a field has been set.

### SetReferencedInvoiceIdNil

`func (o *InvoiceInput) SetReferencedInvoiceIdNil(b bool)`

 SetReferencedInvoiceIdNil sets the value for ReferencedInvoiceId to be an explicit nil

### UnsetReferencedInvoiceId
`func (o *InvoiceInput) UnsetReferencedInvoiceId()`

UnsetReferencedInvoiceId ensures that no value is present for ReferencedInvoiceId, not even an explicit nil
### GetReferencedInvoiceDate

`func (o *InvoiceInput) GetReferencedInvoiceDate() string`

GetReferencedInvoiceDate returns the ReferencedInvoiceDate field if non-nil, zero value otherwise.

### GetReferencedInvoiceDateOk

`func (o *InvoiceInput) GetReferencedInvoiceDateOk() (*string, bool)`

GetReferencedInvoiceDateOk returns a tuple with the ReferencedInvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedInvoiceDate

`func (o *InvoiceInput) SetReferencedInvoiceDate(v string)`

SetReferencedInvoiceDate sets ReferencedInvoiceDate field to given value.

### HasReferencedInvoiceDate

`func (o *InvoiceInput) HasReferencedInvoiceDate() bool`

HasReferencedInvoiceDate returns a boolean if a field has been set.

### SetReferencedInvoiceDateNil

`func (o *InvoiceInput) SetReferencedInvoiceDateNil(b bool)`

 SetReferencedInvoiceDateNil sets the value for ReferencedInvoiceDate to be an explicit nil

### UnsetReferencedInvoiceDate
`func (o *InvoiceInput) UnsetReferencedInvoiceDate()`

UnsetReferencedInvoiceDate ensures that no value is present for ReferencedInvoiceDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



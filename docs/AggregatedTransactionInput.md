# AggregatedTransactionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | Transaction date (TT-77) | 
**CategoryCode** | [**TransactionCategory**](TransactionCategory.md) | Transaction category code (TT-81). Use TLB1 for goods, TPS1 for services. | 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] [default to EUR]
**TaxExclusiveAmount** | [**Taxexclusiveamount**](Taxexclusiveamount.md) |  | 
**TaxAmount** | [**Taxamount**](Taxamount.md) |  | 
**TaxBreakdown** | [**[]TaxBreakdownInput**](TaxBreakdownInput.md) | VAT breakdown by rate | 
**TransactionCount** | Pointer to **NullableInt32** |  | [optional] 
**TaxDueType** | Pointer to [**NullableTaxDueDateType**](TaxDueDateType.md) |  | [optional] 

## Methods

### NewAggregatedTransactionInput

`func NewAggregatedTransactionInput(date string, categoryCode TransactionCategory, taxExclusiveAmount Taxexclusiveamount, taxAmount Taxamount, taxBreakdown []TaxBreakdownInput, ) *AggregatedTransactionInput`

NewAggregatedTransactionInput instantiates a new AggregatedTransactionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregatedTransactionInputWithDefaults

`func NewAggregatedTransactionInputWithDefaults() *AggregatedTransactionInput`

NewAggregatedTransactionInputWithDefaults instantiates a new AggregatedTransactionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *AggregatedTransactionInput) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AggregatedTransactionInput) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AggregatedTransactionInput) SetDate(v string)`

SetDate sets Date field to given value.


### GetCategoryCode

`func (o *AggregatedTransactionInput) GetCategoryCode() TransactionCategory`

GetCategoryCode returns the CategoryCode field if non-nil, zero value otherwise.

### GetCategoryCodeOk

`func (o *AggregatedTransactionInput) GetCategoryCodeOk() (*TransactionCategory, bool)`

GetCategoryCodeOk returns a tuple with the CategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryCode

`func (o *AggregatedTransactionInput) SetCategoryCode(v TransactionCategory)`

SetCategoryCode sets CategoryCode field to given value.


### GetCurrency

`func (o *AggregatedTransactionInput) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AggregatedTransactionInput) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AggregatedTransactionInput) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AggregatedTransactionInput) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTaxExclusiveAmount

`func (o *AggregatedTransactionInput) GetTaxExclusiveAmount() Taxexclusiveamount`

GetTaxExclusiveAmount returns the TaxExclusiveAmount field if non-nil, zero value otherwise.

### GetTaxExclusiveAmountOk

`func (o *AggregatedTransactionInput) GetTaxExclusiveAmountOk() (*Taxexclusiveamount, bool)`

GetTaxExclusiveAmountOk returns a tuple with the TaxExclusiveAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxExclusiveAmount

`func (o *AggregatedTransactionInput) SetTaxExclusiveAmount(v Taxexclusiveamount)`

SetTaxExclusiveAmount sets TaxExclusiveAmount field to given value.


### GetTaxAmount

`func (o *AggregatedTransactionInput) GetTaxAmount() Taxamount`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *AggregatedTransactionInput) GetTaxAmountOk() (*Taxamount, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *AggregatedTransactionInput) SetTaxAmount(v Taxamount)`

SetTaxAmount sets TaxAmount field to given value.


### GetTaxBreakdown

`func (o *AggregatedTransactionInput) GetTaxBreakdown() []TaxBreakdownInput`

GetTaxBreakdown returns the TaxBreakdown field if non-nil, zero value otherwise.

### GetTaxBreakdownOk

`func (o *AggregatedTransactionInput) GetTaxBreakdownOk() (*[]TaxBreakdownInput, bool)`

GetTaxBreakdownOk returns a tuple with the TaxBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxBreakdown

`func (o *AggregatedTransactionInput) SetTaxBreakdown(v []TaxBreakdownInput)`

SetTaxBreakdown sets TaxBreakdown field to given value.


### GetTransactionCount

`func (o *AggregatedTransactionInput) GetTransactionCount() int32`

GetTransactionCount returns the TransactionCount field if non-nil, zero value otherwise.

### GetTransactionCountOk

`func (o *AggregatedTransactionInput) GetTransactionCountOk() (*int32, bool)`

GetTransactionCountOk returns a tuple with the TransactionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCount

`func (o *AggregatedTransactionInput) SetTransactionCount(v int32)`

SetTransactionCount sets TransactionCount field to given value.

### HasTransactionCount

`func (o *AggregatedTransactionInput) HasTransactionCount() bool`

HasTransactionCount returns a boolean if a field has been set.

### SetTransactionCountNil

`func (o *AggregatedTransactionInput) SetTransactionCountNil(b bool)`

 SetTransactionCountNil sets the value for TransactionCount to be an explicit nil

### UnsetTransactionCount
`func (o *AggregatedTransactionInput) UnsetTransactionCount()`

UnsetTransactionCount ensures that no value is present for TransactionCount, not even an explicit nil
### GetTaxDueType

`func (o *AggregatedTransactionInput) GetTaxDueType() TaxDueDateType`

GetTaxDueType returns the TaxDueType field if non-nil, zero value otherwise.

### GetTaxDueTypeOk

`func (o *AggregatedTransactionInput) GetTaxDueTypeOk() (*TaxDueDateType, bool)`

GetTaxDueTypeOk returns a tuple with the TaxDueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDueType

`func (o *AggregatedTransactionInput) SetTaxDueType(v TaxDueDateType)`

SetTaxDueType sets TaxDueType field to given value.

### HasTaxDueType

`func (o *AggregatedTransactionInput) HasTaxDueType() bool`

HasTaxDueType returns a boolean if a field has been set.

### SetTaxDueTypeNil

`func (o *AggregatedTransactionInput) SetTaxDueTypeNil(b bool)`

 SetTaxDueTypeNil sets the value for TaxDueType to be an explicit nil

### UnsetTaxDueType
`func (o *AggregatedTransactionInput) UnsetTaxDueType()`

UnsetTaxDueType ensures that no value is present for TaxDueType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AggregatedPaymentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentDate** | **string** | Payment date | 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] [default to EUR]
**AmountsByRate** | [**[]PaymentAmountByRate**](PaymentAmountByRate.md) | Payment amounts by VAT rate | 

## Methods

### NewAggregatedPaymentInput

`func NewAggregatedPaymentInput(paymentDate string, amountsByRate []PaymentAmountByRate, ) *AggregatedPaymentInput`

NewAggregatedPaymentInput instantiates a new AggregatedPaymentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregatedPaymentInputWithDefaults

`func NewAggregatedPaymentInputWithDefaults() *AggregatedPaymentInput`

NewAggregatedPaymentInputWithDefaults instantiates a new AggregatedPaymentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentDate

`func (o *AggregatedPaymentInput) GetPaymentDate() string`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *AggregatedPaymentInput) GetPaymentDateOk() (*string, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *AggregatedPaymentInput) SetPaymentDate(v string)`

SetPaymentDate sets PaymentDate field to given value.


### GetCurrency

`func (o *AggregatedPaymentInput) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AggregatedPaymentInput) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AggregatedPaymentInput) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AggregatedPaymentInput) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAmountsByRate

`func (o *AggregatedPaymentInput) GetAmountsByRate() []PaymentAmountByRate`

GetAmountsByRate returns the AmountsByRate field if non-nil, zero value otherwise.

### GetAmountsByRateOk

`func (o *AggregatedPaymentInput) GetAmountsByRateOk() (*[]PaymentAmountByRate, bool)`

GetAmountsByRateOk returns a tuple with the AmountsByRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountsByRate

`func (o *AggregatedPaymentInput) SetAmountsByRate(v []PaymentAmountByRate)`

SetAmountsByRate sets AmountsByRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



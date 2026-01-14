# InvoicePaymentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** | Invoice identifier | 
**InvoiceDate** | **string** | Original invoice date | 
**PaymentDate** | **string** | Payment date | 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] [default to EUR]
**AmountsByRate** | [**[]PaymentAmountByRate**](PaymentAmountByRate.md) | Payment amounts by VAT rate | 

## Methods

### NewInvoicePaymentInput

`func NewInvoicePaymentInput(invoiceId string, invoiceDate string, paymentDate string, amountsByRate []PaymentAmountByRate, ) *InvoicePaymentInput`

NewInvoicePaymentInput instantiates a new InvoicePaymentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicePaymentInputWithDefaults

`func NewInvoicePaymentInputWithDefaults() *InvoicePaymentInput`

NewInvoicePaymentInputWithDefaults instantiates a new InvoicePaymentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *InvoicePaymentInput) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *InvoicePaymentInput) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *InvoicePaymentInput) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetInvoiceDate

`func (o *InvoicePaymentInput) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *InvoicePaymentInput) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *InvoicePaymentInput) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.


### GetPaymentDate

`func (o *InvoicePaymentInput) GetPaymentDate() string`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *InvoicePaymentInput) GetPaymentDateOk() (*string, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *InvoicePaymentInput) SetPaymentDate(v string)`

SetPaymentDate sets PaymentDate field to given value.


### GetCurrency

`func (o *InvoicePaymentInput) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoicePaymentInput) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoicePaymentInput) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InvoicePaymentInput) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAmountsByRate

`func (o *InvoicePaymentInput) GetAmountsByRate() []PaymentAmountByRate`

GetAmountsByRate returns the AmountsByRate field if non-nil, zero value otherwise.

### GetAmountsByRateOk

`func (o *InvoicePaymentInput) GetAmountsByRateOk() (*[]PaymentAmountByRate, bool)`

GetAmountsByRateOk returns a tuple with the AmountsByRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountsByRate

`func (o *InvoicePaymentInput) SetAmountsByRate(v []PaymentAmountByRate)`

SetAmountsByRate sets AmountsByRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



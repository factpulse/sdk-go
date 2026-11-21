# PaymentAmountByRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rate** | [**Rate**](Rate.md) |  | 
**Amount** | [**Amount3**](Amount3.md) |  | 
**AmountEur** | Pointer to [**NullableAmounteur**](Amounteur.md) |  | [optional] 

## Methods

### NewPaymentAmountByRate

`func NewPaymentAmountByRate(rate Rate, amount Amount3, ) *PaymentAmountByRate`

NewPaymentAmountByRate instantiates a new PaymentAmountByRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentAmountByRateWithDefaults

`func NewPaymentAmountByRateWithDefaults() *PaymentAmountByRate`

NewPaymentAmountByRateWithDefaults instantiates a new PaymentAmountByRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRate

`func (o *PaymentAmountByRate) GetRate() Rate`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *PaymentAmountByRate) GetRateOk() (*Rate, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *PaymentAmountByRate) SetRate(v Rate)`

SetRate sets Rate field to given value.


### GetAmount

`func (o *PaymentAmountByRate) GetAmount() Amount3`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentAmountByRate) GetAmountOk() (*Amount3, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentAmountByRate) SetAmount(v Amount3)`

SetAmount sets Amount field to given value.


### GetAmountEur

`func (o *PaymentAmountByRate) GetAmountEur() Amounteur`

GetAmountEur returns the AmountEur field if non-nil, zero value otherwise.

### GetAmountEurOk

`func (o *PaymentAmountByRate) GetAmountEurOk() (*Amounteur, bool)`

GetAmountEurOk returns a tuple with the AmountEur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountEur

`func (o *PaymentAmountByRate) SetAmountEur(v Amounteur)`

SetAmountEur sets AmountEur field to given value.

### HasAmountEur

`func (o *PaymentAmountByRate) HasAmountEur() bool`

HasAmountEur returns a boolean if a field has been set.

### SetAmountEurNil

`func (o *PaymentAmountByRate) SetAmountEurNil(b bool)`

 SetAmountEurNil sets the value for AmountEur to be an explicit nil

### UnsetAmountEur
`func (o *PaymentAmountByRate) UnsetAmountEur()`

UnsetAmountEur ensures that no value is present for AmountEur, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# EncaisseAmountByRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rate** | **float32** | Taux de TVA applicable au montant encaissé, en % (MDT-224) | 
**Amount** | [**Amount1**](Amount1.md) |  | 

## Methods

### NewEncaisseAmountByRate

`func NewEncaisseAmountByRate(rate float32, amount Amount1, ) *EncaisseAmountByRate`

NewEncaisseAmountByRate instantiates a new EncaisseAmountByRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncaisseAmountByRateWithDefaults

`func NewEncaisseAmountByRateWithDefaults() *EncaisseAmountByRate`

NewEncaisseAmountByRateWithDefaults instantiates a new EncaisseAmountByRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRate

`func (o *EncaisseAmountByRate) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *EncaisseAmountByRate) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *EncaisseAmountByRate) SetRate(v float32)`

SetRate sets Rate field to given value.


### GetAmount

`func (o *EncaisseAmountByRate) GetAmount() Amount1`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EncaisseAmountByRate) GetAmountOk() (*Amount1, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EncaisseAmountByRate) SetAmount(v Amount1)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



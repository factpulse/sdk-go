# TaxBreakdownInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rate** | [**Rate1**](Rate1.md) |  | 
**TaxableAmount** | [**Taxableamount**](Taxableamount.md) |  | 
**TaxAmount** | [**Taxamount2**](Taxamount2.md) |  | 

## Methods

### NewTaxBreakdownInput

`func NewTaxBreakdownInput(rate Rate1, taxableAmount Taxableamount, taxAmount Taxamount2, ) *TaxBreakdownInput`

NewTaxBreakdownInput instantiates a new TaxBreakdownInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxBreakdownInputWithDefaults

`func NewTaxBreakdownInputWithDefaults() *TaxBreakdownInput`

NewTaxBreakdownInputWithDefaults instantiates a new TaxBreakdownInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRate

`func (o *TaxBreakdownInput) GetRate() Rate1`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *TaxBreakdownInput) GetRateOk() (*Rate1, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *TaxBreakdownInput) SetRate(v Rate1)`

SetRate sets Rate field to given value.


### GetTaxableAmount

`func (o *TaxBreakdownInput) GetTaxableAmount() Taxableamount`

GetTaxableAmount returns the TaxableAmount field if non-nil, zero value otherwise.

### GetTaxableAmountOk

`func (o *TaxBreakdownInput) GetTaxableAmountOk() (*Taxableamount, bool)`

GetTaxableAmountOk returns a tuple with the TaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableAmount

`func (o *TaxBreakdownInput) SetTaxableAmount(v Taxableamount)`

SetTaxableAmount sets TaxableAmount field to given value.


### GetTaxAmount

`func (o *TaxBreakdownInput) GetTaxAmount() Taxamount2`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *TaxBreakdownInput) GetTaxAmountOk() (*Taxamount2, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *TaxBreakdownInput) SetTaxAmount(v Taxamount2)`

SetTaxAmount sets TaxAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



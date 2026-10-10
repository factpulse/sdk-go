# TaxBreakdownInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rate** | [**Rate1**](Rate1.md) |  | 
**TaxableAmount** | [**Taxableamount**](Taxableamount.md) |  | 
**TaxAmount** | [**Taxamount2**](Taxamount2.md) |  | 
**TaxCategoryCode** | Pointer to **NullableString** |  | [optional] 
**ExemptionReasonCode** | Pointer to **NullableString** |  | [optional] 
**ExemptionReason** | Pointer to **NullableString** |  | [optional] 

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


### GetTaxCategoryCode

`func (o *TaxBreakdownInput) GetTaxCategoryCode() string`

GetTaxCategoryCode returns the TaxCategoryCode field if non-nil, zero value otherwise.

### GetTaxCategoryCodeOk

`func (o *TaxBreakdownInput) GetTaxCategoryCodeOk() (*string, bool)`

GetTaxCategoryCodeOk returns a tuple with the TaxCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCategoryCode

`func (o *TaxBreakdownInput) SetTaxCategoryCode(v string)`

SetTaxCategoryCode sets TaxCategoryCode field to given value.

### HasTaxCategoryCode

`func (o *TaxBreakdownInput) HasTaxCategoryCode() bool`

HasTaxCategoryCode returns a boolean if a field has been set.

### SetTaxCategoryCodeNil

`func (o *TaxBreakdownInput) SetTaxCategoryCodeNil(b bool)`

 SetTaxCategoryCodeNil sets the value for TaxCategoryCode to be an explicit nil

### UnsetTaxCategoryCode
`func (o *TaxBreakdownInput) UnsetTaxCategoryCode()`

UnsetTaxCategoryCode ensures that no value is present for TaxCategoryCode, not even an explicit nil
### GetExemptionReasonCode

`func (o *TaxBreakdownInput) GetExemptionReasonCode() string`

GetExemptionReasonCode returns the ExemptionReasonCode field if non-nil, zero value otherwise.

### GetExemptionReasonCodeOk

`func (o *TaxBreakdownInput) GetExemptionReasonCodeOk() (*string, bool)`

GetExemptionReasonCodeOk returns a tuple with the ExemptionReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptionReasonCode

`func (o *TaxBreakdownInput) SetExemptionReasonCode(v string)`

SetExemptionReasonCode sets ExemptionReasonCode field to given value.

### HasExemptionReasonCode

`func (o *TaxBreakdownInput) HasExemptionReasonCode() bool`

HasExemptionReasonCode returns a boolean if a field has been set.

### SetExemptionReasonCodeNil

`func (o *TaxBreakdownInput) SetExemptionReasonCodeNil(b bool)`

 SetExemptionReasonCodeNil sets the value for ExemptionReasonCode to be an explicit nil

### UnsetExemptionReasonCode
`func (o *TaxBreakdownInput) UnsetExemptionReasonCode()`

UnsetExemptionReasonCode ensures that no value is present for ExemptionReasonCode, not even an explicit nil
### GetExemptionReason

`func (o *TaxBreakdownInput) GetExemptionReason() string`

GetExemptionReason returns the ExemptionReason field if non-nil, zero value otherwise.

### GetExemptionReasonOk

`func (o *TaxBreakdownInput) GetExemptionReasonOk() (*string, bool)`

GetExemptionReasonOk returns a tuple with the ExemptionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptionReason

`func (o *TaxBreakdownInput) SetExemptionReason(v string)`

SetExemptionReason sets ExemptionReason field to given value.

### HasExemptionReason

`func (o *TaxBreakdownInput) HasExemptionReason() bool`

HasExemptionReason returns a boolean if a field has been set.

### SetExemptionReasonNil

`func (o *TaxBreakdownInput) SetExemptionReasonNil(b bool)`

 SetExemptionReasonNil sets the value for ExemptionReason to be an explicit nil

### UnsetExemptionReason
`func (o *TaxBreakdownInput) UnsetExemptionReason()`

UnsetExemptionReason ensures that no value is present for ExemptionReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



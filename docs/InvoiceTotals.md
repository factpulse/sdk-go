# InvoiceTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineTotalAmount** | Pointer to [**NullableLineTotalAmount**](LineTotalAmount.md) |  | [optional] 
**AllowanceTotalAmount** | Pointer to [**NullableAllowanceTotalAmount**](AllowanceTotalAmount.md) |  | [optional] 
**ChargeTotalAmount** | Pointer to [**NullableChargeTotalAmount**](ChargeTotalAmount.md) |  | [optional] 
**TotalNetAmount** | [**TotalNetAmount**](TotalNetAmount.md) |  | 
**VatAmount** | [**TotalVATAmount**](TotalVATAmount.md) |  | 
**TotalGrossAmount** | [**TotalGrossAmount**](TotalGrossAmount.md) |  | 
**Prepayment** | Pointer to [**NullableInvoiceTotalsPrepayment**](InvoiceTotalsPrepayment.md) |  | [optional] 
**RoundingAmount** | Pointer to [**NullableRoundingAmount**](RoundingAmount.md) |  | [optional] 
**AmountDue** | [**AmountDue**](AmountDue.md) |  | 
**GlobalAllowanceAmount** | Pointer to [**GlobalAllowanceAmount**](GlobalAllowanceAmount.md) |  | [optional] 
**GlobalAllowanceReason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInvoiceTotals

`func NewInvoiceTotals(totalNetAmount TotalNetAmount, vatAmount TotalVATAmount, totalGrossAmount TotalGrossAmount, amountDue AmountDue, ) *InvoiceTotals`

NewInvoiceTotals instantiates a new InvoiceTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceTotalsWithDefaults

`func NewInvoiceTotalsWithDefaults() *InvoiceTotals`

NewInvoiceTotalsWithDefaults instantiates a new InvoiceTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineTotalAmount

`func (o *InvoiceTotals) GetLineTotalAmount() LineTotalAmount`

GetLineTotalAmount returns the LineTotalAmount field if non-nil, zero value otherwise.

### GetLineTotalAmountOk

`func (o *InvoiceTotals) GetLineTotalAmountOk() (*LineTotalAmount, bool)`

GetLineTotalAmountOk returns a tuple with the LineTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineTotalAmount

`func (o *InvoiceTotals) SetLineTotalAmount(v LineTotalAmount)`

SetLineTotalAmount sets LineTotalAmount field to given value.

### HasLineTotalAmount

`func (o *InvoiceTotals) HasLineTotalAmount() bool`

HasLineTotalAmount returns a boolean if a field has been set.

### SetLineTotalAmountNil

`func (o *InvoiceTotals) SetLineTotalAmountNil(b bool)`

 SetLineTotalAmountNil sets the value for LineTotalAmount to be an explicit nil

### UnsetLineTotalAmount
`func (o *InvoiceTotals) UnsetLineTotalAmount()`

UnsetLineTotalAmount ensures that no value is present for LineTotalAmount, not even an explicit nil
### GetAllowanceTotalAmount

`func (o *InvoiceTotals) GetAllowanceTotalAmount() AllowanceTotalAmount`

GetAllowanceTotalAmount returns the AllowanceTotalAmount field if non-nil, zero value otherwise.

### GetAllowanceTotalAmountOk

`func (o *InvoiceTotals) GetAllowanceTotalAmountOk() (*AllowanceTotalAmount, bool)`

GetAllowanceTotalAmountOk returns a tuple with the AllowanceTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowanceTotalAmount

`func (o *InvoiceTotals) SetAllowanceTotalAmount(v AllowanceTotalAmount)`

SetAllowanceTotalAmount sets AllowanceTotalAmount field to given value.

### HasAllowanceTotalAmount

`func (o *InvoiceTotals) HasAllowanceTotalAmount() bool`

HasAllowanceTotalAmount returns a boolean if a field has been set.

### SetAllowanceTotalAmountNil

`func (o *InvoiceTotals) SetAllowanceTotalAmountNil(b bool)`

 SetAllowanceTotalAmountNil sets the value for AllowanceTotalAmount to be an explicit nil

### UnsetAllowanceTotalAmount
`func (o *InvoiceTotals) UnsetAllowanceTotalAmount()`

UnsetAllowanceTotalAmount ensures that no value is present for AllowanceTotalAmount, not even an explicit nil
### GetChargeTotalAmount

`func (o *InvoiceTotals) GetChargeTotalAmount() ChargeTotalAmount`

GetChargeTotalAmount returns the ChargeTotalAmount field if non-nil, zero value otherwise.

### GetChargeTotalAmountOk

`func (o *InvoiceTotals) GetChargeTotalAmountOk() (*ChargeTotalAmount, bool)`

GetChargeTotalAmountOk returns a tuple with the ChargeTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeTotalAmount

`func (o *InvoiceTotals) SetChargeTotalAmount(v ChargeTotalAmount)`

SetChargeTotalAmount sets ChargeTotalAmount field to given value.

### HasChargeTotalAmount

`func (o *InvoiceTotals) HasChargeTotalAmount() bool`

HasChargeTotalAmount returns a boolean if a field has been set.

### SetChargeTotalAmountNil

`func (o *InvoiceTotals) SetChargeTotalAmountNil(b bool)`

 SetChargeTotalAmountNil sets the value for ChargeTotalAmount to be an explicit nil

### UnsetChargeTotalAmount
`func (o *InvoiceTotals) UnsetChargeTotalAmount()`

UnsetChargeTotalAmount ensures that no value is present for ChargeTotalAmount, not even an explicit nil
### GetTotalNetAmount

`func (o *InvoiceTotals) GetTotalNetAmount() TotalNetAmount`

GetTotalNetAmount returns the TotalNetAmount field if non-nil, zero value otherwise.

### GetTotalNetAmountOk

`func (o *InvoiceTotals) GetTotalNetAmountOk() (*TotalNetAmount, bool)`

GetTotalNetAmountOk returns a tuple with the TotalNetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNetAmount

`func (o *InvoiceTotals) SetTotalNetAmount(v TotalNetAmount)`

SetTotalNetAmount sets TotalNetAmount field to given value.


### GetVatAmount

`func (o *InvoiceTotals) GetVatAmount() TotalVATAmount`

GetVatAmount returns the VatAmount field if non-nil, zero value otherwise.

### GetVatAmountOk

`func (o *InvoiceTotals) GetVatAmountOk() (*TotalVATAmount, bool)`

GetVatAmountOk returns a tuple with the VatAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatAmount

`func (o *InvoiceTotals) SetVatAmount(v TotalVATAmount)`

SetVatAmount sets VatAmount field to given value.


### GetTotalGrossAmount

`func (o *InvoiceTotals) GetTotalGrossAmount() TotalGrossAmount`

GetTotalGrossAmount returns the TotalGrossAmount field if non-nil, zero value otherwise.

### GetTotalGrossAmountOk

`func (o *InvoiceTotals) GetTotalGrossAmountOk() (*TotalGrossAmount, bool)`

GetTotalGrossAmountOk returns a tuple with the TotalGrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossAmount

`func (o *InvoiceTotals) SetTotalGrossAmount(v TotalGrossAmount)`

SetTotalGrossAmount sets TotalGrossAmount field to given value.


### GetPrepayment

`func (o *InvoiceTotals) GetPrepayment() InvoiceTotalsPrepayment`

GetPrepayment returns the Prepayment field if non-nil, zero value otherwise.

### GetPrepaymentOk

`func (o *InvoiceTotals) GetPrepaymentOk() (*InvoiceTotalsPrepayment, bool)`

GetPrepaymentOk returns a tuple with the Prepayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepayment

`func (o *InvoiceTotals) SetPrepayment(v InvoiceTotalsPrepayment)`

SetPrepayment sets Prepayment field to given value.

### HasPrepayment

`func (o *InvoiceTotals) HasPrepayment() bool`

HasPrepayment returns a boolean if a field has been set.

### SetPrepaymentNil

`func (o *InvoiceTotals) SetPrepaymentNil(b bool)`

 SetPrepaymentNil sets the value for Prepayment to be an explicit nil

### UnsetPrepayment
`func (o *InvoiceTotals) UnsetPrepayment()`

UnsetPrepayment ensures that no value is present for Prepayment, not even an explicit nil
### GetRoundingAmount

`func (o *InvoiceTotals) GetRoundingAmount() RoundingAmount`

GetRoundingAmount returns the RoundingAmount field if non-nil, zero value otherwise.

### GetRoundingAmountOk

`func (o *InvoiceTotals) GetRoundingAmountOk() (*RoundingAmount, bool)`

GetRoundingAmountOk returns a tuple with the RoundingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundingAmount

`func (o *InvoiceTotals) SetRoundingAmount(v RoundingAmount)`

SetRoundingAmount sets RoundingAmount field to given value.

### HasRoundingAmount

`func (o *InvoiceTotals) HasRoundingAmount() bool`

HasRoundingAmount returns a boolean if a field has been set.

### SetRoundingAmountNil

`func (o *InvoiceTotals) SetRoundingAmountNil(b bool)`

 SetRoundingAmountNil sets the value for RoundingAmount to be an explicit nil

### UnsetRoundingAmount
`func (o *InvoiceTotals) UnsetRoundingAmount()`

UnsetRoundingAmount ensures that no value is present for RoundingAmount, not even an explicit nil
### GetAmountDue

`func (o *InvoiceTotals) GetAmountDue() AmountDue`

GetAmountDue returns the AmountDue field if non-nil, zero value otherwise.

### GetAmountDueOk

`func (o *InvoiceTotals) GetAmountDueOk() (*AmountDue, bool)`

GetAmountDueOk returns a tuple with the AmountDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDue

`func (o *InvoiceTotals) SetAmountDue(v AmountDue)`

SetAmountDue sets AmountDue field to given value.


### GetGlobalAllowanceAmount

`func (o *InvoiceTotals) GetGlobalAllowanceAmount() GlobalAllowanceAmount`

GetGlobalAllowanceAmount returns the GlobalAllowanceAmount field if non-nil, zero value otherwise.

### GetGlobalAllowanceAmountOk

`func (o *InvoiceTotals) GetGlobalAllowanceAmountOk() (*GlobalAllowanceAmount, bool)`

GetGlobalAllowanceAmountOk returns a tuple with the GlobalAllowanceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAllowanceAmount

`func (o *InvoiceTotals) SetGlobalAllowanceAmount(v GlobalAllowanceAmount)`

SetGlobalAllowanceAmount sets GlobalAllowanceAmount field to given value.

### HasGlobalAllowanceAmount

`func (o *InvoiceTotals) HasGlobalAllowanceAmount() bool`

HasGlobalAllowanceAmount returns a boolean if a field has been set.

### GetGlobalAllowanceReason

`func (o *InvoiceTotals) GetGlobalAllowanceReason() string`

GetGlobalAllowanceReason returns the GlobalAllowanceReason field if non-nil, zero value otherwise.

### GetGlobalAllowanceReasonOk

`func (o *InvoiceTotals) GetGlobalAllowanceReasonOk() (*string, bool)`

GetGlobalAllowanceReasonOk returns a tuple with the GlobalAllowanceReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAllowanceReason

`func (o *InvoiceTotals) SetGlobalAllowanceReason(v string)`

SetGlobalAllowanceReason sets GlobalAllowanceReason field to given value.

### HasGlobalAllowanceReason

`func (o *InvoiceTotals) HasGlobalAllowanceReason() bool`

HasGlobalAllowanceReason returns a boolean if a field has been set.

### SetGlobalAllowanceReasonNil

`func (o *InvoiceTotals) SetGlobalAllowanceReasonNil(b bool)`

 SetGlobalAllowanceReasonNil sets the value for GlobalAllowanceReason to be an explicit nil

### UnsetGlobalAllowanceReason
`func (o *InvoiceTotals) UnsetGlobalAllowanceReason()`

UnsetGlobalAllowanceReason ensures that no value is present for GlobalAllowanceReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



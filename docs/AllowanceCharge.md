# AllowanceCharge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCharge** | **bool** | True for charge, False for allowance (ChargeIndicator). | 
**Amount** | [**Amount**](Amount.md) |  | 
**BaseAmount** | Pointer to [**NullableBaseAmount**](BaseAmount.md) |  | [optional] 
**Percentage** | Pointer to [**NullablePercentage**](Percentage.md) |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 
**ReasonCode** | Pointer to **NullableString** |  | [optional] 
**VatCategory** | Pointer to **NullableString** |  | [optional] 
**VatRate** | Pointer to [**NullableVatRate**](VatRate.md) |  | [optional] 

## Methods

### NewAllowanceCharge

`func NewAllowanceCharge(isCharge bool, amount Amount, ) *AllowanceCharge`

NewAllowanceCharge instantiates a new AllowanceCharge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowanceChargeWithDefaults

`func NewAllowanceChargeWithDefaults() *AllowanceCharge`

NewAllowanceChargeWithDefaults instantiates a new AllowanceCharge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCharge

`func (o *AllowanceCharge) GetIsCharge() bool`

GetIsCharge returns the IsCharge field if non-nil, zero value otherwise.

### GetIsChargeOk

`func (o *AllowanceCharge) GetIsChargeOk() (*bool, bool)`

GetIsChargeOk returns a tuple with the IsCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCharge

`func (o *AllowanceCharge) SetIsCharge(v bool)`

SetIsCharge sets IsCharge field to given value.


### GetAmount

`func (o *AllowanceCharge) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AllowanceCharge) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AllowanceCharge) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetBaseAmount

`func (o *AllowanceCharge) GetBaseAmount() BaseAmount`

GetBaseAmount returns the BaseAmount field if non-nil, zero value otherwise.

### GetBaseAmountOk

`func (o *AllowanceCharge) GetBaseAmountOk() (*BaseAmount, bool)`

GetBaseAmountOk returns a tuple with the BaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmount

`func (o *AllowanceCharge) SetBaseAmount(v BaseAmount)`

SetBaseAmount sets BaseAmount field to given value.

### HasBaseAmount

`func (o *AllowanceCharge) HasBaseAmount() bool`

HasBaseAmount returns a boolean if a field has been set.

### SetBaseAmountNil

`func (o *AllowanceCharge) SetBaseAmountNil(b bool)`

 SetBaseAmountNil sets the value for BaseAmount to be an explicit nil

### UnsetBaseAmount
`func (o *AllowanceCharge) UnsetBaseAmount()`

UnsetBaseAmount ensures that no value is present for BaseAmount, not even an explicit nil
### GetPercentage

`func (o *AllowanceCharge) GetPercentage() Percentage`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *AllowanceCharge) GetPercentageOk() (*Percentage, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *AllowanceCharge) SetPercentage(v Percentage)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *AllowanceCharge) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### SetPercentageNil

`func (o *AllowanceCharge) SetPercentageNil(b bool)`

 SetPercentageNil sets the value for Percentage to be an explicit nil

### UnsetPercentage
`func (o *AllowanceCharge) UnsetPercentage()`

UnsetPercentage ensures that no value is present for Percentage, not even an explicit nil
### GetReason

`func (o *AllowanceCharge) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AllowanceCharge) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AllowanceCharge) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AllowanceCharge) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *AllowanceCharge) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *AllowanceCharge) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetReasonCode

`func (o *AllowanceCharge) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *AllowanceCharge) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *AllowanceCharge) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *AllowanceCharge) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### SetReasonCodeNil

`func (o *AllowanceCharge) SetReasonCodeNil(b bool)`

 SetReasonCodeNil sets the value for ReasonCode to be an explicit nil

### UnsetReasonCode
`func (o *AllowanceCharge) UnsetReasonCode()`

UnsetReasonCode ensures that no value is present for ReasonCode, not even an explicit nil
### GetVatCategory

`func (o *AllowanceCharge) GetVatCategory() string`

GetVatCategory returns the VatCategory field if non-nil, zero value otherwise.

### GetVatCategoryOk

`func (o *AllowanceCharge) GetVatCategoryOk() (*string, bool)`

GetVatCategoryOk returns a tuple with the VatCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatCategory

`func (o *AllowanceCharge) SetVatCategory(v string)`

SetVatCategory sets VatCategory field to given value.

### HasVatCategory

`func (o *AllowanceCharge) HasVatCategory() bool`

HasVatCategory returns a boolean if a field has been set.

### SetVatCategoryNil

`func (o *AllowanceCharge) SetVatCategoryNil(b bool)`

 SetVatCategoryNil sets the value for VatCategory to be an explicit nil

### UnsetVatCategory
`func (o *AllowanceCharge) UnsetVatCategory()`

UnsetVatCategory ensures that no value is present for VatCategory, not even an explicit nil
### GetVatRate

`func (o *AllowanceCharge) GetVatRate() VatRate`

GetVatRate returns the VatRate field if non-nil, zero value otherwise.

### GetVatRateOk

`func (o *AllowanceCharge) GetVatRateOk() (*VatRate, bool)`

GetVatRateOk returns a tuple with the VatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatRate

`func (o *AllowanceCharge) SetVatRate(v VatRate)`

SetVatRate sets VatRate field to given value.

### HasVatRate

`func (o *AllowanceCharge) HasVatRate() bool`

HasVatRate returns a boolean if a field has been set.

### SetVatRateNil

`func (o *AllowanceCharge) SetVatRateNil(b bool)`

 SetVatRateNil sets the value for VatRate to be an explicit nil

### UnsetVatRate
`func (o *AllowanceCharge) UnsetVatRate()`

UnsetVatRate ensures that no value is present for VatRate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



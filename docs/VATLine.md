# VATLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxableAmount** | [**TaxableAmount**](TaxableAmount.md) |  | 
**VatAmount** | [**VATAmount**](VATAmount.md) |  | 
**Rate** | Pointer to **NullableString** |  | [optional] 
**ManualRate** | Pointer to [**ManualRate**](ManualRate.md) |  | [optional] 
**Category** | Pointer to [**NullableVATCategory**](VATCategory.md) |  | [optional] 
**ExemptionReason** | Pointer to **NullableString** |  | [optional] 
**VatexCode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVATLine

`func NewVATLine(taxableAmount TaxableAmount, vatAmount VATAmount, ) *VATLine`

NewVATLine instantiates a new VATLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVATLineWithDefaults

`func NewVATLineWithDefaults() *VATLine`

NewVATLineWithDefaults instantiates a new VATLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxableAmount

`func (o *VATLine) GetTaxableAmount() TaxableAmount`

GetTaxableAmount returns the TaxableAmount field if non-nil, zero value otherwise.

### GetTaxableAmountOk

`func (o *VATLine) GetTaxableAmountOk() (*TaxableAmount, bool)`

GetTaxableAmountOk returns a tuple with the TaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableAmount

`func (o *VATLine) SetTaxableAmount(v TaxableAmount)`

SetTaxableAmount sets TaxableAmount field to given value.


### GetVatAmount

`func (o *VATLine) GetVatAmount() VATAmount`

GetVatAmount returns the VatAmount field if non-nil, zero value otherwise.

### GetVatAmountOk

`func (o *VATLine) GetVatAmountOk() (*VATAmount, bool)`

GetVatAmountOk returns a tuple with the VatAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatAmount

`func (o *VATLine) SetVatAmount(v VATAmount)`

SetVatAmount sets VatAmount field to given value.


### GetRate

`func (o *VATLine) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *VATLine) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *VATLine) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *VATLine) HasRate() bool`

HasRate returns a boolean if a field has been set.

### SetRateNil

`func (o *VATLine) SetRateNil(b bool)`

 SetRateNil sets the value for Rate to be an explicit nil

### UnsetRate
`func (o *VATLine) UnsetRate()`

UnsetRate ensures that no value is present for Rate, not even an explicit nil
### GetManualRate

`func (o *VATLine) GetManualRate() ManualRate`

GetManualRate returns the ManualRate field if non-nil, zero value otherwise.

### GetManualRateOk

`func (o *VATLine) GetManualRateOk() (*ManualRate, bool)`

GetManualRateOk returns a tuple with the ManualRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualRate

`func (o *VATLine) SetManualRate(v ManualRate)`

SetManualRate sets ManualRate field to given value.

### HasManualRate

`func (o *VATLine) HasManualRate() bool`

HasManualRate returns a boolean if a field has been set.

### GetCategory

`func (o *VATLine) GetCategory() VATCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *VATLine) GetCategoryOk() (*VATCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *VATLine) SetCategory(v VATCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *VATLine) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *VATLine) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *VATLine) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetExemptionReason

`func (o *VATLine) GetExemptionReason() string`

GetExemptionReason returns the ExemptionReason field if non-nil, zero value otherwise.

### GetExemptionReasonOk

`func (o *VATLine) GetExemptionReasonOk() (*string, bool)`

GetExemptionReasonOk returns a tuple with the ExemptionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptionReason

`func (o *VATLine) SetExemptionReason(v string)`

SetExemptionReason sets ExemptionReason field to given value.

### HasExemptionReason

`func (o *VATLine) HasExemptionReason() bool`

HasExemptionReason returns a boolean if a field has been set.

### SetExemptionReasonNil

`func (o *VATLine) SetExemptionReasonNil(b bool)`

 SetExemptionReasonNil sets the value for ExemptionReason to be an explicit nil

### UnsetExemptionReason
`func (o *VATLine) UnsetExemptionReason()`

UnsetExemptionReason ensures that no value is present for ExemptionReason, not even an explicit nil
### GetVatexCode

`func (o *VATLine) GetVatexCode() string`

GetVatexCode returns the VatexCode field if non-nil, zero value otherwise.

### GetVatexCodeOk

`func (o *VATLine) GetVatexCodeOk() (*string, bool)`

GetVatexCodeOk returns a tuple with the VatexCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatexCode

`func (o *VATLine) SetVatexCode(v string)`

SetVatexCode sets VatexCode field to given value.

### HasVatexCode

`func (o *VATLine) HasVatexCode() bool`

HasVatexCode returns a boolean if a field has been set.

### SetVatexCodeNil

`func (o *VATLine) SetVatexCodeNil(b bool)`

 SetVatexCodeNil sets the value for VatexCode to be an explicit nil

### UnsetVatexCode
`func (o *VATLine) UnsetVatexCode()`

UnsetVatexCode ensures that no value is present for VatexCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



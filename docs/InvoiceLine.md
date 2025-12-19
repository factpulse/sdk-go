# InvoiceLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineNumber** | **int32** |  | 
**Reference** | Pointer to **NullableString** |  | [optional] 
**ItemName** | **string** |  | 
**Quantity** | [**Quantity**](Quantity.md) |  | 
**Unit** | [**UnitOfMeasure**](UnitOfMeasure.md) |  | 
**UnitNetPrice** | [**UnitNetPrice**](UnitNetPrice.md) |  | 
**AllowanceAmount** | Pointer to [**NullableInvoiceLineAllowanceAmount**](InvoiceLineAllowanceAmount.md) |  | [optional] 
**LineNetAmount** | Pointer to [**LineNetAmount**](LineNetAmount.md) |  | [optional] 
**VatRate** | Pointer to **NullableString** |  | [optional] 
**ManualVatRate** | Pointer to [**ManualVatRate**](ManualVatRate.md) |  | [optional] 
**VatCategory** | Pointer to [**NullableVATCategory**](VATCategory.md) |  | [optional] 
**PeriodStartDate** | Pointer to **NullableString** |  | [optional] 
**PeriodEndDate** | Pointer to **NullableString** |  | [optional] 
**AllowanceReasonCode** | Pointer to [**NullableAllowanceReasonCode**](AllowanceReasonCode.md) |  | [optional] 
**AllowanceReason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInvoiceLine

`func NewInvoiceLine(lineNumber int32, itemName string, quantity Quantity, unit UnitOfMeasure, unitNetPrice UnitNetPrice, ) *InvoiceLine`

NewInvoiceLine instantiates a new InvoiceLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineWithDefaults

`func NewInvoiceLineWithDefaults() *InvoiceLine`

NewInvoiceLineWithDefaults instantiates a new InvoiceLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineNumber

`func (o *InvoiceLine) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *InvoiceLine) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *InvoiceLine) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.


### GetReference

`func (o *InvoiceLine) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *InvoiceLine) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *InvoiceLine) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *InvoiceLine) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *InvoiceLine) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *InvoiceLine) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetItemName

`func (o *InvoiceLine) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *InvoiceLine) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *InvoiceLine) SetItemName(v string)`

SetItemName sets ItemName field to given value.


### GetQuantity

`func (o *InvoiceLine) GetQuantity() Quantity`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceLine) GetQuantityOk() (*Quantity, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceLine) SetQuantity(v Quantity)`

SetQuantity sets Quantity field to given value.


### GetUnit

`func (o *InvoiceLine) GetUnit() UnitOfMeasure`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *InvoiceLine) GetUnitOk() (*UnitOfMeasure, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *InvoiceLine) SetUnit(v UnitOfMeasure)`

SetUnit sets Unit field to given value.


### GetUnitNetPrice

`func (o *InvoiceLine) GetUnitNetPrice() UnitNetPrice`

GetUnitNetPrice returns the UnitNetPrice field if non-nil, zero value otherwise.

### GetUnitNetPriceOk

`func (o *InvoiceLine) GetUnitNetPriceOk() (*UnitNetPrice, bool)`

GetUnitNetPriceOk returns a tuple with the UnitNetPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNetPrice

`func (o *InvoiceLine) SetUnitNetPrice(v UnitNetPrice)`

SetUnitNetPrice sets UnitNetPrice field to given value.


### GetAllowanceAmount

`func (o *InvoiceLine) GetAllowanceAmount() InvoiceLineAllowanceAmount`

GetAllowanceAmount returns the AllowanceAmount field if non-nil, zero value otherwise.

### GetAllowanceAmountOk

`func (o *InvoiceLine) GetAllowanceAmountOk() (*InvoiceLineAllowanceAmount, bool)`

GetAllowanceAmountOk returns a tuple with the AllowanceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowanceAmount

`func (o *InvoiceLine) SetAllowanceAmount(v InvoiceLineAllowanceAmount)`

SetAllowanceAmount sets AllowanceAmount field to given value.

### HasAllowanceAmount

`func (o *InvoiceLine) HasAllowanceAmount() bool`

HasAllowanceAmount returns a boolean if a field has been set.

### SetAllowanceAmountNil

`func (o *InvoiceLine) SetAllowanceAmountNil(b bool)`

 SetAllowanceAmountNil sets the value for AllowanceAmount to be an explicit nil

### UnsetAllowanceAmount
`func (o *InvoiceLine) UnsetAllowanceAmount()`

UnsetAllowanceAmount ensures that no value is present for AllowanceAmount, not even an explicit nil
### GetLineNetAmount

`func (o *InvoiceLine) GetLineNetAmount() LineNetAmount`

GetLineNetAmount returns the LineNetAmount field if non-nil, zero value otherwise.

### GetLineNetAmountOk

`func (o *InvoiceLine) GetLineNetAmountOk() (*LineNetAmount, bool)`

GetLineNetAmountOk returns a tuple with the LineNetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNetAmount

`func (o *InvoiceLine) SetLineNetAmount(v LineNetAmount)`

SetLineNetAmount sets LineNetAmount field to given value.

### HasLineNetAmount

`func (o *InvoiceLine) HasLineNetAmount() bool`

HasLineNetAmount returns a boolean if a field has been set.

### GetVatRate

`func (o *InvoiceLine) GetVatRate() string`

GetVatRate returns the VatRate field if non-nil, zero value otherwise.

### GetVatRateOk

`func (o *InvoiceLine) GetVatRateOk() (*string, bool)`

GetVatRateOk returns a tuple with the VatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatRate

`func (o *InvoiceLine) SetVatRate(v string)`

SetVatRate sets VatRate field to given value.

### HasVatRate

`func (o *InvoiceLine) HasVatRate() bool`

HasVatRate returns a boolean if a field has been set.

### SetVatRateNil

`func (o *InvoiceLine) SetVatRateNil(b bool)`

 SetVatRateNil sets the value for VatRate to be an explicit nil

### UnsetVatRate
`func (o *InvoiceLine) UnsetVatRate()`

UnsetVatRate ensures that no value is present for VatRate, not even an explicit nil
### GetManualVatRate

`func (o *InvoiceLine) GetManualVatRate() ManualVatRate`

GetManualVatRate returns the ManualVatRate field if non-nil, zero value otherwise.

### GetManualVatRateOk

`func (o *InvoiceLine) GetManualVatRateOk() (*ManualVatRate, bool)`

GetManualVatRateOk returns a tuple with the ManualVatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualVatRate

`func (o *InvoiceLine) SetManualVatRate(v ManualVatRate)`

SetManualVatRate sets ManualVatRate field to given value.

### HasManualVatRate

`func (o *InvoiceLine) HasManualVatRate() bool`

HasManualVatRate returns a boolean if a field has been set.

### GetVatCategory

`func (o *InvoiceLine) GetVatCategory() VATCategory`

GetVatCategory returns the VatCategory field if non-nil, zero value otherwise.

### GetVatCategoryOk

`func (o *InvoiceLine) GetVatCategoryOk() (*VATCategory, bool)`

GetVatCategoryOk returns a tuple with the VatCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatCategory

`func (o *InvoiceLine) SetVatCategory(v VATCategory)`

SetVatCategory sets VatCategory field to given value.

### HasVatCategory

`func (o *InvoiceLine) HasVatCategory() bool`

HasVatCategory returns a boolean if a field has been set.

### SetVatCategoryNil

`func (o *InvoiceLine) SetVatCategoryNil(b bool)`

 SetVatCategoryNil sets the value for VatCategory to be an explicit nil

### UnsetVatCategory
`func (o *InvoiceLine) UnsetVatCategory()`

UnsetVatCategory ensures that no value is present for VatCategory, not even an explicit nil
### GetPeriodStartDate

`func (o *InvoiceLine) GetPeriodStartDate() string`

GetPeriodStartDate returns the PeriodStartDate field if non-nil, zero value otherwise.

### GetPeriodStartDateOk

`func (o *InvoiceLine) GetPeriodStartDateOk() (*string, bool)`

GetPeriodStartDateOk returns a tuple with the PeriodStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStartDate

`func (o *InvoiceLine) SetPeriodStartDate(v string)`

SetPeriodStartDate sets PeriodStartDate field to given value.

### HasPeriodStartDate

`func (o *InvoiceLine) HasPeriodStartDate() bool`

HasPeriodStartDate returns a boolean if a field has been set.

### SetPeriodStartDateNil

`func (o *InvoiceLine) SetPeriodStartDateNil(b bool)`

 SetPeriodStartDateNil sets the value for PeriodStartDate to be an explicit nil

### UnsetPeriodStartDate
`func (o *InvoiceLine) UnsetPeriodStartDate()`

UnsetPeriodStartDate ensures that no value is present for PeriodStartDate, not even an explicit nil
### GetPeriodEndDate

`func (o *InvoiceLine) GetPeriodEndDate() string`

GetPeriodEndDate returns the PeriodEndDate field if non-nil, zero value otherwise.

### GetPeriodEndDateOk

`func (o *InvoiceLine) GetPeriodEndDateOk() (*string, bool)`

GetPeriodEndDateOk returns a tuple with the PeriodEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEndDate

`func (o *InvoiceLine) SetPeriodEndDate(v string)`

SetPeriodEndDate sets PeriodEndDate field to given value.

### HasPeriodEndDate

`func (o *InvoiceLine) HasPeriodEndDate() bool`

HasPeriodEndDate returns a boolean if a field has been set.

### SetPeriodEndDateNil

`func (o *InvoiceLine) SetPeriodEndDateNil(b bool)`

 SetPeriodEndDateNil sets the value for PeriodEndDate to be an explicit nil

### UnsetPeriodEndDate
`func (o *InvoiceLine) UnsetPeriodEndDate()`

UnsetPeriodEndDate ensures that no value is present for PeriodEndDate, not even an explicit nil
### GetAllowanceReasonCode

`func (o *InvoiceLine) GetAllowanceReasonCode() AllowanceReasonCode`

GetAllowanceReasonCode returns the AllowanceReasonCode field if non-nil, zero value otherwise.

### GetAllowanceReasonCodeOk

`func (o *InvoiceLine) GetAllowanceReasonCodeOk() (*AllowanceReasonCode, bool)`

GetAllowanceReasonCodeOk returns a tuple with the AllowanceReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowanceReasonCode

`func (o *InvoiceLine) SetAllowanceReasonCode(v AllowanceReasonCode)`

SetAllowanceReasonCode sets AllowanceReasonCode field to given value.

### HasAllowanceReasonCode

`func (o *InvoiceLine) HasAllowanceReasonCode() bool`

HasAllowanceReasonCode returns a boolean if a field has been set.

### SetAllowanceReasonCodeNil

`func (o *InvoiceLine) SetAllowanceReasonCodeNil(b bool)`

 SetAllowanceReasonCodeNil sets the value for AllowanceReasonCode to be an explicit nil

### UnsetAllowanceReasonCode
`func (o *InvoiceLine) UnsetAllowanceReasonCode()`

UnsetAllowanceReasonCode ensures that no value is present for AllowanceReasonCode, not even an explicit nil
### GetAllowanceReason

`func (o *InvoiceLine) GetAllowanceReason() string`

GetAllowanceReason returns the AllowanceReason field if non-nil, zero value otherwise.

### GetAllowanceReasonOk

`func (o *InvoiceLine) GetAllowanceReasonOk() (*string, bool)`

GetAllowanceReasonOk returns a tuple with the AllowanceReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowanceReason

`func (o *InvoiceLine) SetAllowanceReason(v string)`

SetAllowanceReason sets AllowanceReason field to given value.

### HasAllowanceReason

`func (o *InvoiceLine) HasAllowanceReason() bool`

HasAllowanceReason returns a boolean if a field has been set.

### SetAllowanceReasonNil

`func (o *InvoiceLine) SetAllowanceReasonNil(b bool)`

 SetAllowanceReasonNil sets the value for AllowanceReason to be an explicit nil

### UnsetAllowanceReason
`func (o *InvoiceLine) UnsetAllowanceReason()`

UnsetAllowanceReason ensures that no value is present for AllowanceReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



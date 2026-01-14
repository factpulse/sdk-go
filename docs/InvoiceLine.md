# InvoiceLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineNumber** | **int32** | Invoice line identifier (BT-126). | 
**LineNote** | Pointer to **NullableString** |  | [optional] 
**ParentLineId** | Pointer to **NullableString** |  | [optional] 
**LineSubType** | Pointer to [**NullableLineSubType**](LineSubType.md) |  | [optional] 
**Reference** | Pointer to **NullableString** |  | [optional] 
**BuyerAssignedId** | Pointer to **NullableString** |  | [optional] 
**ProductGlobalId** | Pointer to **NullableString** |  | [optional] 
**ProductGlobalIdScheme** | Pointer to **NullableString** |  | [optional] 
**ItemName** | **string** | Item name (BT-153). | 
**ItemDescription** | Pointer to **NullableString** |  | [optional] 
**OriginCountry** | Pointer to **NullableString** |  | [optional] 
**Characteristics** | Pointer to [**[]ProductCharacteristic**](ProductCharacteristic.md) |  | [optional] 
**Classifications** | Pointer to [**[]ProductClassification**](ProductClassification.md) |  | [optional] 
**Quantity** | [**Quantity**](Quantity.md) |  | 
**Unit** | [**UnitOfMeasure**](UnitOfMeasure.md) | Invoiced quantity unit of measure code (BT-130). | 
**GrossUnitPrice** | Pointer to [**NullableGrossUnitPrice**](GrossUnitPrice.md) |  | [optional] 
**UnitNetPrice** | [**UnitNetPrice**](UnitNetPrice.md) |  | 
**PriceBasisQuantity** | Pointer to [**NullablePriceBasisQuantity**](PriceBasisQuantity.md) |  | [optional] 
**PriceBasisUnit** | Pointer to **NullableString** |  | [optional] 
**PriceAllowanceAmount** | Pointer to [**NullablePriceAllowanceAmount**](PriceAllowanceAmount.md) |  | [optional] 
**LineNetAmount** | Pointer to [**NullableLineNetAmount**](LineNetAmount.md) |  | [optional] 
**AllowanceAmount** | Pointer to [**NullableInvoiceLineAllowanceAmount**](InvoiceLineAllowanceAmount.md) |  | [optional] 
**AllowanceReasonCode** | Pointer to [**NullableAllowanceReasonCode**](AllowanceReasonCode.md) |  | [optional] 
**AllowanceReason** | Pointer to **NullableString** |  | [optional] 
**AllowancesCharges** | Pointer to [**[]AllowanceCharge**](AllowanceCharge.md) |  | [optional] 
**VatRate** | Pointer to **NullableString** |  | [optional] 
**ManualVatRate** | Pointer to [**ManualVatRate**](ManualVatRate.md) |  | [optional] 
**VatCategory** | Pointer to [**NullableVATCategory**](VATCategory.md) |  | [optional] 
**PeriodStartDate** | Pointer to **NullableString** |  | [optional] 
**PeriodEndDate** | Pointer to **NullableString** |  | [optional] 
**PurchaseOrderLineRef** | Pointer to **NullableString** |  | [optional] 
**AccountingAccount** | Pointer to **NullableString** |  | [optional] 
**AdditionalDocuments** | Pointer to [**[]AdditionalDocument**](AdditionalDocument.md) |  | [optional] 
**LineNotes** | Pointer to [**[]InvoiceNote**](InvoiceNote.md) |  | [optional] 

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


### GetLineNote

`func (o *InvoiceLine) GetLineNote() string`

GetLineNote returns the LineNote field if non-nil, zero value otherwise.

### GetLineNoteOk

`func (o *InvoiceLine) GetLineNoteOk() (*string, bool)`

GetLineNoteOk returns a tuple with the LineNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNote

`func (o *InvoiceLine) SetLineNote(v string)`

SetLineNote sets LineNote field to given value.

### HasLineNote

`func (o *InvoiceLine) HasLineNote() bool`

HasLineNote returns a boolean if a field has been set.

### SetLineNoteNil

`func (o *InvoiceLine) SetLineNoteNil(b bool)`

 SetLineNoteNil sets the value for LineNote to be an explicit nil

### UnsetLineNote
`func (o *InvoiceLine) UnsetLineNote()`

UnsetLineNote ensures that no value is present for LineNote, not even an explicit nil
### GetParentLineId

`func (o *InvoiceLine) GetParentLineId() string`

GetParentLineId returns the ParentLineId field if non-nil, zero value otherwise.

### GetParentLineIdOk

`func (o *InvoiceLine) GetParentLineIdOk() (*string, bool)`

GetParentLineIdOk returns a tuple with the ParentLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentLineId

`func (o *InvoiceLine) SetParentLineId(v string)`

SetParentLineId sets ParentLineId field to given value.

### HasParentLineId

`func (o *InvoiceLine) HasParentLineId() bool`

HasParentLineId returns a boolean if a field has been set.

### SetParentLineIdNil

`func (o *InvoiceLine) SetParentLineIdNil(b bool)`

 SetParentLineIdNil sets the value for ParentLineId to be an explicit nil

### UnsetParentLineId
`func (o *InvoiceLine) UnsetParentLineId()`

UnsetParentLineId ensures that no value is present for ParentLineId, not even an explicit nil
### GetLineSubType

`func (o *InvoiceLine) GetLineSubType() LineSubType`

GetLineSubType returns the LineSubType field if non-nil, zero value otherwise.

### GetLineSubTypeOk

`func (o *InvoiceLine) GetLineSubTypeOk() (*LineSubType, bool)`

GetLineSubTypeOk returns a tuple with the LineSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineSubType

`func (o *InvoiceLine) SetLineSubType(v LineSubType)`

SetLineSubType sets LineSubType field to given value.

### HasLineSubType

`func (o *InvoiceLine) HasLineSubType() bool`

HasLineSubType returns a boolean if a field has been set.

### SetLineSubTypeNil

`func (o *InvoiceLine) SetLineSubTypeNil(b bool)`

 SetLineSubTypeNil sets the value for LineSubType to be an explicit nil

### UnsetLineSubType
`func (o *InvoiceLine) UnsetLineSubType()`

UnsetLineSubType ensures that no value is present for LineSubType, not even an explicit nil
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
### GetBuyerAssignedId

`func (o *InvoiceLine) GetBuyerAssignedId() string`

GetBuyerAssignedId returns the BuyerAssignedId field if non-nil, zero value otherwise.

### GetBuyerAssignedIdOk

`func (o *InvoiceLine) GetBuyerAssignedIdOk() (*string, bool)`

GetBuyerAssignedIdOk returns a tuple with the BuyerAssignedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerAssignedId

`func (o *InvoiceLine) SetBuyerAssignedId(v string)`

SetBuyerAssignedId sets BuyerAssignedId field to given value.

### HasBuyerAssignedId

`func (o *InvoiceLine) HasBuyerAssignedId() bool`

HasBuyerAssignedId returns a boolean if a field has been set.

### SetBuyerAssignedIdNil

`func (o *InvoiceLine) SetBuyerAssignedIdNil(b bool)`

 SetBuyerAssignedIdNil sets the value for BuyerAssignedId to be an explicit nil

### UnsetBuyerAssignedId
`func (o *InvoiceLine) UnsetBuyerAssignedId()`

UnsetBuyerAssignedId ensures that no value is present for BuyerAssignedId, not even an explicit nil
### GetProductGlobalId

`func (o *InvoiceLine) GetProductGlobalId() string`

GetProductGlobalId returns the ProductGlobalId field if non-nil, zero value otherwise.

### GetProductGlobalIdOk

`func (o *InvoiceLine) GetProductGlobalIdOk() (*string, bool)`

GetProductGlobalIdOk returns a tuple with the ProductGlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGlobalId

`func (o *InvoiceLine) SetProductGlobalId(v string)`

SetProductGlobalId sets ProductGlobalId field to given value.

### HasProductGlobalId

`func (o *InvoiceLine) HasProductGlobalId() bool`

HasProductGlobalId returns a boolean if a field has been set.

### SetProductGlobalIdNil

`func (o *InvoiceLine) SetProductGlobalIdNil(b bool)`

 SetProductGlobalIdNil sets the value for ProductGlobalId to be an explicit nil

### UnsetProductGlobalId
`func (o *InvoiceLine) UnsetProductGlobalId()`

UnsetProductGlobalId ensures that no value is present for ProductGlobalId, not even an explicit nil
### GetProductGlobalIdScheme

`func (o *InvoiceLine) GetProductGlobalIdScheme() string`

GetProductGlobalIdScheme returns the ProductGlobalIdScheme field if non-nil, zero value otherwise.

### GetProductGlobalIdSchemeOk

`func (o *InvoiceLine) GetProductGlobalIdSchemeOk() (*string, bool)`

GetProductGlobalIdSchemeOk returns a tuple with the ProductGlobalIdScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGlobalIdScheme

`func (o *InvoiceLine) SetProductGlobalIdScheme(v string)`

SetProductGlobalIdScheme sets ProductGlobalIdScheme field to given value.

### HasProductGlobalIdScheme

`func (o *InvoiceLine) HasProductGlobalIdScheme() bool`

HasProductGlobalIdScheme returns a boolean if a field has been set.

### SetProductGlobalIdSchemeNil

`func (o *InvoiceLine) SetProductGlobalIdSchemeNil(b bool)`

 SetProductGlobalIdSchemeNil sets the value for ProductGlobalIdScheme to be an explicit nil

### UnsetProductGlobalIdScheme
`func (o *InvoiceLine) UnsetProductGlobalIdScheme()`

UnsetProductGlobalIdScheme ensures that no value is present for ProductGlobalIdScheme, not even an explicit nil
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


### GetItemDescription

`func (o *InvoiceLine) GetItemDescription() string`

GetItemDescription returns the ItemDescription field if non-nil, zero value otherwise.

### GetItemDescriptionOk

`func (o *InvoiceLine) GetItemDescriptionOk() (*string, bool)`

GetItemDescriptionOk returns a tuple with the ItemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDescription

`func (o *InvoiceLine) SetItemDescription(v string)`

SetItemDescription sets ItemDescription field to given value.

### HasItemDescription

`func (o *InvoiceLine) HasItemDescription() bool`

HasItemDescription returns a boolean if a field has been set.

### SetItemDescriptionNil

`func (o *InvoiceLine) SetItemDescriptionNil(b bool)`

 SetItemDescriptionNil sets the value for ItemDescription to be an explicit nil

### UnsetItemDescription
`func (o *InvoiceLine) UnsetItemDescription()`

UnsetItemDescription ensures that no value is present for ItemDescription, not even an explicit nil
### GetOriginCountry

`func (o *InvoiceLine) GetOriginCountry() string`

GetOriginCountry returns the OriginCountry field if non-nil, zero value otherwise.

### GetOriginCountryOk

`func (o *InvoiceLine) GetOriginCountryOk() (*string, bool)`

GetOriginCountryOk returns a tuple with the OriginCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCountry

`func (o *InvoiceLine) SetOriginCountry(v string)`

SetOriginCountry sets OriginCountry field to given value.

### HasOriginCountry

`func (o *InvoiceLine) HasOriginCountry() bool`

HasOriginCountry returns a boolean if a field has been set.

### SetOriginCountryNil

`func (o *InvoiceLine) SetOriginCountryNil(b bool)`

 SetOriginCountryNil sets the value for OriginCountry to be an explicit nil

### UnsetOriginCountry
`func (o *InvoiceLine) UnsetOriginCountry()`

UnsetOriginCountry ensures that no value is present for OriginCountry, not even an explicit nil
### GetCharacteristics

`func (o *InvoiceLine) GetCharacteristics() []ProductCharacteristic`

GetCharacteristics returns the Characteristics field if non-nil, zero value otherwise.

### GetCharacteristicsOk

`func (o *InvoiceLine) GetCharacteristicsOk() (*[]ProductCharacteristic, bool)`

GetCharacteristicsOk returns a tuple with the Characteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacteristics

`func (o *InvoiceLine) SetCharacteristics(v []ProductCharacteristic)`

SetCharacteristics sets Characteristics field to given value.

### HasCharacteristics

`func (o *InvoiceLine) HasCharacteristics() bool`

HasCharacteristics returns a boolean if a field has been set.

### SetCharacteristicsNil

`func (o *InvoiceLine) SetCharacteristicsNil(b bool)`

 SetCharacteristicsNil sets the value for Characteristics to be an explicit nil

### UnsetCharacteristics
`func (o *InvoiceLine) UnsetCharacteristics()`

UnsetCharacteristics ensures that no value is present for Characteristics, not even an explicit nil
### GetClassifications

`func (o *InvoiceLine) GetClassifications() []ProductClassification`

GetClassifications returns the Classifications field if non-nil, zero value otherwise.

### GetClassificationsOk

`func (o *InvoiceLine) GetClassificationsOk() (*[]ProductClassification, bool)`

GetClassificationsOk returns a tuple with the Classifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifications

`func (o *InvoiceLine) SetClassifications(v []ProductClassification)`

SetClassifications sets Classifications field to given value.

### HasClassifications

`func (o *InvoiceLine) HasClassifications() bool`

HasClassifications returns a boolean if a field has been set.

### SetClassificationsNil

`func (o *InvoiceLine) SetClassificationsNil(b bool)`

 SetClassificationsNil sets the value for Classifications to be an explicit nil

### UnsetClassifications
`func (o *InvoiceLine) UnsetClassifications()`

UnsetClassifications ensures that no value is present for Classifications, not even an explicit nil
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


### GetGrossUnitPrice

`func (o *InvoiceLine) GetGrossUnitPrice() GrossUnitPrice`

GetGrossUnitPrice returns the GrossUnitPrice field if non-nil, zero value otherwise.

### GetGrossUnitPriceOk

`func (o *InvoiceLine) GetGrossUnitPriceOk() (*GrossUnitPrice, bool)`

GetGrossUnitPriceOk returns a tuple with the GrossUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossUnitPrice

`func (o *InvoiceLine) SetGrossUnitPrice(v GrossUnitPrice)`

SetGrossUnitPrice sets GrossUnitPrice field to given value.

### HasGrossUnitPrice

`func (o *InvoiceLine) HasGrossUnitPrice() bool`

HasGrossUnitPrice returns a boolean if a field has been set.

### SetGrossUnitPriceNil

`func (o *InvoiceLine) SetGrossUnitPriceNil(b bool)`

 SetGrossUnitPriceNil sets the value for GrossUnitPrice to be an explicit nil

### UnsetGrossUnitPrice
`func (o *InvoiceLine) UnsetGrossUnitPrice()`

UnsetGrossUnitPrice ensures that no value is present for GrossUnitPrice, not even an explicit nil
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


### GetPriceBasisQuantity

`func (o *InvoiceLine) GetPriceBasisQuantity() PriceBasisQuantity`

GetPriceBasisQuantity returns the PriceBasisQuantity field if non-nil, zero value otherwise.

### GetPriceBasisQuantityOk

`func (o *InvoiceLine) GetPriceBasisQuantityOk() (*PriceBasisQuantity, bool)`

GetPriceBasisQuantityOk returns a tuple with the PriceBasisQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceBasisQuantity

`func (o *InvoiceLine) SetPriceBasisQuantity(v PriceBasisQuantity)`

SetPriceBasisQuantity sets PriceBasisQuantity field to given value.

### HasPriceBasisQuantity

`func (o *InvoiceLine) HasPriceBasisQuantity() bool`

HasPriceBasisQuantity returns a boolean if a field has been set.

### SetPriceBasisQuantityNil

`func (o *InvoiceLine) SetPriceBasisQuantityNil(b bool)`

 SetPriceBasisQuantityNil sets the value for PriceBasisQuantity to be an explicit nil

### UnsetPriceBasisQuantity
`func (o *InvoiceLine) UnsetPriceBasisQuantity()`

UnsetPriceBasisQuantity ensures that no value is present for PriceBasisQuantity, not even an explicit nil
### GetPriceBasisUnit

`func (o *InvoiceLine) GetPriceBasisUnit() string`

GetPriceBasisUnit returns the PriceBasisUnit field if non-nil, zero value otherwise.

### GetPriceBasisUnitOk

`func (o *InvoiceLine) GetPriceBasisUnitOk() (*string, bool)`

GetPriceBasisUnitOk returns a tuple with the PriceBasisUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceBasisUnit

`func (o *InvoiceLine) SetPriceBasisUnit(v string)`

SetPriceBasisUnit sets PriceBasisUnit field to given value.

### HasPriceBasisUnit

`func (o *InvoiceLine) HasPriceBasisUnit() bool`

HasPriceBasisUnit returns a boolean if a field has been set.

### SetPriceBasisUnitNil

`func (o *InvoiceLine) SetPriceBasisUnitNil(b bool)`

 SetPriceBasisUnitNil sets the value for PriceBasisUnit to be an explicit nil

### UnsetPriceBasisUnit
`func (o *InvoiceLine) UnsetPriceBasisUnit()`

UnsetPriceBasisUnit ensures that no value is present for PriceBasisUnit, not even an explicit nil
### GetPriceAllowanceAmount

`func (o *InvoiceLine) GetPriceAllowanceAmount() PriceAllowanceAmount`

GetPriceAllowanceAmount returns the PriceAllowanceAmount field if non-nil, zero value otherwise.

### GetPriceAllowanceAmountOk

`func (o *InvoiceLine) GetPriceAllowanceAmountOk() (*PriceAllowanceAmount, bool)`

GetPriceAllowanceAmountOk returns a tuple with the PriceAllowanceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAllowanceAmount

`func (o *InvoiceLine) SetPriceAllowanceAmount(v PriceAllowanceAmount)`

SetPriceAllowanceAmount sets PriceAllowanceAmount field to given value.

### HasPriceAllowanceAmount

`func (o *InvoiceLine) HasPriceAllowanceAmount() bool`

HasPriceAllowanceAmount returns a boolean if a field has been set.

### SetPriceAllowanceAmountNil

`func (o *InvoiceLine) SetPriceAllowanceAmountNil(b bool)`

 SetPriceAllowanceAmountNil sets the value for PriceAllowanceAmount to be an explicit nil

### UnsetPriceAllowanceAmount
`func (o *InvoiceLine) UnsetPriceAllowanceAmount()`

UnsetPriceAllowanceAmount ensures that no value is present for PriceAllowanceAmount, not even an explicit nil
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

### SetLineNetAmountNil

`func (o *InvoiceLine) SetLineNetAmountNil(b bool)`

 SetLineNetAmountNil sets the value for LineNetAmount to be an explicit nil

### UnsetLineNetAmount
`func (o *InvoiceLine) UnsetLineNetAmount()`

UnsetLineNetAmount ensures that no value is present for LineNetAmount, not even an explicit nil
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
### GetAllowancesCharges

`func (o *InvoiceLine) GetAllowancesCharges() []AllowanceCharge`

GetAllowancesCharges returns the AllowancesCharges field if non-nil, zero value otherwise.

### GetAllowancesChargesOk

`func (o *InvoiceLine) GetAllowancesChargesOk() (*[]AllowanceCharge, bool)`

GetAllowancesChargesOk returns a tuple with the AllowancesCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowancesCharges

`func (o *InvoiceLine) SetAllowancesCharges(v []AllowanceCharge)`

SetAllowancesCharges sets AllowancesCharges field to given value.

### HasAllowancesCharges

`func (o *InvoiceLine) HasAllowancesCharges() bool`

HasAllowancesCharges returns a boolean if a field has been set.

### SetAllowancesChargesNil

`func (o *InvoiceLine) SetAllowancesChargesNil(b bool)`

 SetAllowancesChargesNil sets the value for AllowancesCharges to be an explicit nil

### UnsetAllowancesCharges
`func (o *InvoiceLine) UnsetAllowancesCharges()`

UnsetAllowancesCharges ensures that no value is present for AllowancesCharges, not even an explicit nil
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
### GetPurchaseOrderLineRef

`func (o *InvoiceLine) GetPurchaseOrderLineRef() string`

GetPurchaseOrderLineRef returns the PurchaseOrderLineRef field if non-nil, zero value otherwise.

### GetPurchaseOrderLineRefOk

`func (o *InvoiceLine) GetPurchaseOrderLineRefOk() (*string, bool)`

GetPurchaseOrderLineRefOk returns a tuple with the PurchaseOrderLineRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderLineRef

`func (o *InvoiceLine) SetPurchaseOrderLineRef(v string)`

SetPurchaseOrderLineRef sets PurchaseOrderLineRef field to given value.

### HasPurchaseOrderLineRef

`func (o *InvoiceLine) HasPurchaseOrderLineRef() bool`

HasPurchaseOrderLineRef returns a boolean if a field has been set.

### SetPurchaseOrderLineRefNil

`func (o *InvoiceLine) SetPurchaseOrderLineRefNil(b bool)`

 SetPurchaseOrderLineRefNil sets the value for PurchaseOrderLineRef to be an explicit nil

### UnsetPurchaseOrderLineRef
`func (o *InvoiceLine) UnsetPurchaseOrderLineRef()`

UnsetPurchaseOrderLineRef ensures that no value is present for PurchaseOrderLineRef, not even an explicit nil
### GetAccountingAccount

`func (o *InvoiceLine) GetAccountingAccount() string`

GetAccountingAccount returns the AccountingAccount field if non-nil, zero value otherwise.

### GetAccountingAccountOk

`func (o *InvoiceLine) GetAccountingAccountOk() (*string, bool)`

GetAccountingAccountOk returns a tuple with the AccountingAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingAccount

`func (o *InvoiceLine) SetAccountingAccount(v string)`

SetAccountingAccount sets AccountingAccount field to given value.

### HasAccountingAccount

`func (o *InvoiceLine) HasAccountingAccount() bool`

HasAccountingAccount returns a boolean if a field has been set.

### SetAccountingAccountNil

`func (o *InvoiceLine) SetAccountingAccountNil(b bool)`

 SetAccountingAccountNil sets the value for AccountingAccount to be an explicit nil

### UnsetAccountingAccount
`func (o *InvoiceLine) UnsetAccountingAccount()`

UnsetAccountingAccount ensures that no value is present for AccountingAccount, not even an explicit nil
### GetAdditionalDocuments

`func (o *InvoiceLine) GetAdditionalDocuments() []AdditionalDocument`

GetAdditionalDocuments returns the AdditionalDocuments field if non-nil, zero value otherwise.

### GetAdditionalDocumentsOk

`func (o *InvoiceLine) GetAdditionalDocumentsOk() (*[]AdditionalDocument, bool)`

GetAdditionalDocumentsOk returns a tuple with the AdditionalDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDocuments

`func (o *InvoiceLine) SetAdditionalDocuments(v []AdditionalDocument)`

SetAdditionalDocuments sets AdditionalDocuments field to given value.

### HasAdditionalDocuments

`func (o *InvoiceLine) HasAdditionalDocuments() bool`

HasAdditionalDocuments returns a boolean if a field has been set.

### SetAdditionalDocumentsNil

`func (o *InvoiceLine) SetAdditionalDocumentsNil(b bool)`

 SetAdditionalDocumentsNil sets the value for AdditionalDocuments to be an explicit nil

### UnsetAdditionalDocuments
`func (o *InvoiceLine) UnsetAdditionalDocuments()`

UnsetAdditionalDocuments ensures that no value is present for AdditionalDocuments, not even an explicit nil
### GetLineNotes

`func (o *InvoiceLine) GetLineNotes() []InvoiceNote`

GetLineNotes returns the LineNotes field if non-nil, zero value otherwise.

### GetLineNotesOk

`func (o *InvoiceLine) GetLineNotesOk() (*[]InvoiceNote, bool)`

GetLineNotesOk returns a tuple with the LineNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNotes

`func (o *InvoiceLine) SetLineNotes(v []InvoiceNote)`

SetLineNotes sets LineNotes field to given value.

### HasLineNotes

`func (o *InvoiceLine) HasLineNotes() bool`

HasLineNotes returns a boolean if a field has been set.

### SetLineNotesNil

`func (o *InvoiceLine) SetLineNotesNil(b bool)`

 SetLineNotesNil sets the value for LineNotes to be an explicit nil

### UnsetLineNotes
`func (o *InvoiceLine) UnsetLineNotes()`

UnsetLineNotes ensures that no value is present for LineNotes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SimplifiedInvoiceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** | Unique invoice number | 
**Supplier** | **map[string]interface{}** | Supplier information (siret, iban, ...) | 
**Recipient** | **map[string]interface{}** | Recipient information (siret, ...) | 
**Lines** | **[]map[string]interface{}** | Invoice lines | 
**Date** | Pointer to **NullableString** |  | [optional] 
**DueDays** | Pointer to **int32** | Due date in days (default: 30) | [optional] [default to 30]
**Comment** | Pointer to **NullableString** |  | [optional] 
**PurchaseOrderReference** | Pointer to **NullableString** |  | [optional] 
**ContractReference** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSimplifiedInvoiceData

`func NewSimplifiedInvoiceData(number string, supplier map[string]interface{}, recipient map[string]interface{}, lines []map[string]interface{}, ) *SimplifiedInvoiceData`

NewSimplifiedInvoiceData instantiates a new SimplifiedInvoiceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedInvoiceDataWithDefaults

`func NewSimplifiedInvoiceDataWithDefaults() *SimplifiedInvoiceData`

NewSimplifiedInvoiceDataWithDefaults instantiates a new SimplifiedInvoiceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *SimplifiedInvoiceData) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *SimplifiedInvoiceData) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *SimplifiedInvoiceData) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetSupplier

`func (o *SimplifiedInvoiceData) GetSupplier() map[string]interface{}`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *SimplifiedInvoiceData) GetSupplierOk() (*map[string]interface{}, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *SimplifiedInvoiceData) SetSupplier(v map[string]interface{})`

SetSupplier sets Supplier field to given value.


### GetRecipient

`func (o *SimplifiedInvoiceData) GetRecipient() map[string]interface{}`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *SimplifiedInvoiceData) GetRecipientOk() (*map[string]interface{}, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *SimplifiedInvoiceData) SetRecipient(v map[string]interface{})`

SetRecipient sets Recipient field to given value.


### GetLines

`func (o *SimplifiedInvoiceData) GetLines() []map[string]interface{}`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *SimplifiedInvoiceData) GetLinesOk() (*[]map[string]interface{}, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *SimplifiedInvoiceData) SetLines(v []map[string]interface{})`

SetLines sets Lines field to given value.


### GetDate

`func (o *SimplifiedInvoiceData) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *SimplifiedInvoiceData) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *SimplifiedInvoiceData) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *SimplifiedInvoiceData) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *SimplifiedInvoiceData) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *SimplifiedInvoiceData) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetDueDays

`func (o *SimplifiedInvoiceData) GetDueDays() int32`

GetDueDays returns the DueDays field if non-nil, zero value otherwise.

### GetDueDaysOk

`func (o *SimplifiedInvoiceData) GetDueDaysOk() (*int32, bool)`

GetDueDaysOk returns a tuple with the DueDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDays

`func (o *SimplifiedInvoiceData) SetDueDays(v int32)`

SetDueDays sets DueDays field to given value.

### HasDueDays

`func (o *SimplifiedInvoiceData) HasDueDays() bool`

HasDueDays returns a boolean if a field has been set.

### GetComment

`func (o *SimplifiedInvoiceData) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SimplifiedInvoiceData) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SimplifiedInvoiceData) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SimplifiedInvoiceData) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *SimplifiedInvoiceData) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *SimplifiedInvoiceData) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetPurchaseOrderReference

`func (o *SimplifiedInvoiceData) GetPurchaseOrderReference() string`

GetPurchaseOrderReference returns the PurchaseOrderReference field if non-nil, zero value otherwise.

### GetPurchaseOrderReferenceOk

`func (o *SimplifiedInvoiceData) GetPurchaseOrderReferenceOk() (*string, bool)`

GetPurchaseOrderReferenceOk returns a tuple with the PurchaseOrderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderReference

`func (o *SimplifiedInvoiceData) SetPurchaseOrderReference(v string)`

SetPurchaseOrderReference sets PurchaseOrderReference field to given value.

### HasPurchaseOrderReference

`func (o *SimplifiedInvoiceData) HasPurchaseOrderReference() bool`

HasPurchaseOrderReference returns a boolean if a field has been set.

### SetPurchaseOrderReferenceNil

`func (o *SimplifiedInvoiceData) SetPurchaseOrderReferenceNil(b bool)`

 SetPurchaseOrderReferenceNil sets the value for PurchaseOrderReference to be an explicit nil

### UnsetPurchaseOrderReference
`func (o *SimplifiedInvoiceData) UnsetPurchaseOrderReference()`

UnsetPurchaseOrderReference ensures that no value is present for PurchaseOrderReference, not even an explicit nil
### GetContractReference

`func (o *SimplifiedInvoiceData) GetContractReference() string`

GetContractReference returns the ContractReference field if non-nil, zero value otherwise.

### GetContractReferenceOk

`func (o *SimplifiedInvoiceData) GetContractReferenceOk() (*string, bool)`

GetContractReferenceOk returns a tuple with the ContractReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractReference

`func (o *SimplifiedInvoiceData) SetContractReference(v string)`

SetContractReference sets ContractReference field to given value.

### HasContractReference

`func (o *SimplifiedInvoiceData) HasContractReference() bool`

HasContractReference returns a boolean if a field has been set.

### SetContractReferenceNil

`func (o *SimplifiedInvoiceData) SetContractReferenceNil(b bool)`

 SetContractReferenceNil sets the value for ContractReference to be an explicit nil

### UnsetContractReference
`func (o *SimplifiedInvoiceData) UnsetContractReference()`

UnsetContractReference ensures that no value is present for ContractReference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



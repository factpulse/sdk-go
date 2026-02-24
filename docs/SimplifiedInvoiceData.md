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
**InvoiceType** | Pointer to [**InvoiceTypeCode**](InvoiceTypeCode.md) | Document type (UNTDID 1001). Default: 380 (Invoice). | [optional] [default to INVOICE]
**PrecedingInvoiceReference** | Pointer to **NullableString** |  | [optional] 
**OperationNature** | Pointer to [**NullableOperationNature**](OperationNature.md) |  | [optional] 
**InvoicingFramework** | Pointer to [**NullableInvoicingFrameworkCode**](InvoicingFrameworkCode.md) |  | [optional] 

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
### GetInvoiceType

`func (o *SimplifiedInvoiceData) GetInvoiceType() InvoiceTypeCode`

GetInvoiceType returns the InvoiceType field if non-nil, zero value otherwise.

### GetInvoiceTypeOk

`func (o *SimplifiedInvoiceData) GetInvoiceTypeOk() (*InvoiceTypeCode, bool)`

GetInvoiceTypeOk returns a tuple with the InvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceType

`func (o *SimplifiedInvoiceData) SetInvoiceType(v InvoiceTypeCode)`

SetInvoiceType sets InvoiceType field to given value.

### HasInvoiceType

`func (o *SimplifiedInvoiceData) HasInvoiceType() bool`

HasInvoiceType returns a boolean if a field has been set.

### GetPrecedingInvoiceReference

`func (o *SimplifiedInvoiceData) GetPrecedingInvoiceReference() string`

GetPrecedingInvoiceReference returns the PrecedingInvoiceReference field if non-nil, zero value otherwise.

### GetPrecedingInvoiceReferenceOk

`func (o *SimplifiedInvoiceData) GetPrecedingInvoiceReferenceOk() (*string, bool)`

GetPrecedingInvoiceReferenceOk returns a tuple with the PrecedingInvoiceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedingInvoiceReference

`func (o *SimplifiedInvoiceData) SetPrecedingInvoiceReference(v string)`

SetPrecedingInvoiceReference sets PrecedingInvoiceReference field to given value.

### HasPrecedingInvoiceReference

`func (o *SimplifiedInvoiceData) HasPrecedingInvoiceReference() bool`

HasPrecedingInvoiceReference returns a boolean if a field has been set.

### SetPrecedingInvoiceReferenceNil

`func (o *SimplifiedInvoiceData) SetPrecedingInvoiceReferenceNil(b bool)`

 SetPrecedingInvoiceReferenceNil sets the value for PrecedingInvoiceReference to be an explicit nil

### UnsetPrecedingInvoiceReference
`func (o *SimplifiedInvoiceData) UnsetPrecedingInvoiceReference()`

UnsetPrecedingInvoiceReference ensures that no value is present for PrecedingInvoiceReference, not even an explicit nil
### GetOperationNature

`func (o *SimplifiedInvoiceData) GetOperationNature() OperationNature`

GetOperationNature returns the OperationNature field if non-nil, zero value otherwise.

### GetOperationNatureOk

`func (o *SimplifiedInvoiceData) GetOperationNatureOk() (*OperationNature, bool)`

GetOperationNatureOk returns a tuple with the OperationNature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationNature

`func (o *SimplifiedInvoiceData) SetOperationNature(v OperationNature)`

SetOperationNature sets OperationNature field to given value.

### HasOperationNature

`func (o *SimplifiedInvoiceData) HasOperationNature() bool`

HasOperationNature returns a boolean if a field has been set.

### SetOperationNatureNil

`func (o *SimplifiedInvoiceData) SetOperationNatureNil(b bool)`

 SetOperationNatureNil sets the value for OperationNature to be an explicit nil

### UnsetOperationNature
`func (o *SimplifiedInvoiceData) UnsetOperationNature()`

UnsetOperationNature ensures that no value is present for OperationNature, not even an explicit nil
### GetInvoicingFramework

`func (o *SimplifiedInvoiceData) GetInvoicingFramework() InvoicingFrameworkCode`

GetInvoicingFramework returns the InvoicingFramework field if non-nil, zero value otherwise.

### GetInvoicingFrameworkOk

`func (o *SimplifiedInvoiceData) GetInvoicingFrameworkOk() (*InvoicingFrameworkCode, bool)`

GetInvoicingFrameworkOk returns a tuple with the InvoicingFramework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicingFramework

`func (o *SimplifiedInvoiceData) SetInvoicingFramework(v InvoicingFrameworkCode)`

SetInvoicingFramework sets InvoicingFramework field to given value.

### HasInvoicingFramework

`func (o *SimplifiedInvoiceData) HasInvoicingFramework() bool`

HasInvoicingFramework returns a boolean if a field has been set.

### SetInvoicingFrameworkNil

`func (o *SimplifiedInvoiceData) SetInvoicingFrameworkNil(b bool)`

 SetInvoicingFrameworkNil sets the value for InvoicingFramework to be an explicit nil

### UnsetInvoicingFramework
`func (o *SimplifiedInvoiceData) UnsetInvoicingFramework()`

UnsetInvoicingFramework ensures that no value is present for InvoicingFramework, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



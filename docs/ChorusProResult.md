# ChorusProResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChorusInvoiceId** | **int32** | Chorus Pro invoice ID | 
**DepositFlowNumber** | Pointer to **NullableString** |  | [optional] 
**AttachmentId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewChorusProResult

`func NewChorusProResult(chorusInvoiceId int32, ) *ChorusProResult`

NewChorusProResult instantiates a new ChorusProResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChorusProResultWithDefaults

`func NewChorusProResultWithDefaults() *ChorusProResult`

NewChorusProResultWithDefaults instantiates a new ChorusProResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChorusInvoiceId

`func (o *ChorusProResult) GetChorusInvoiceId() int32`

GetChorusInvoiceId returns the ChorusInvoiceId field if non-nil, zero value otherwise.

### GetChorusInvoiceIdOk

`func (o *ChorusProResult) GetChorusInvoiceIdOk() (*int32, bool)`

GetChorusInvoiceIdOk returns a tuple with the ChorusInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChorusInvoiceId

`func (o *ChorusProResult) SetChorusInvoiceId(v int32)`

SetChorusInvoiceId sets ChorusInvoiceId field to given value.


### GetDepositFlowNumber

`func (o *ChorusProResult) GetDepositFlowNumber() string`

GetDepositFlowNumber returns the DepositFlowNumber field if non-nil, zero value otherwise.

### GetDepositFlowNumberOk

`func (o *ChorusProResult) GetDepositFlowNumberOk() (*string, bool)`

GetDepositFlowNumberOk returns a tuple with the DepositFlowNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositFlowNumber

`func (o *ChorusProResult) SetDepositFlowNumber(v string)`

SetDepositFlowNumber sets DepositFlowNumber field to given value.

### HasDepositFlowNumber

`func (o *ChorusProResult) HasDepositFlowNumber() bool`

HasDepositFlowNumber returns a boolean if a field has been set.

### SetDepositFlowNumberNil

`func (o *ChorusProResult) SetDepositFlowNumberNil(b bool)`

 SetDepositFlowNumberNil sets the value for DepositFlowNumber to be an explicit nil

### UnsetDepositFlowNumber
`func (o *ChorusProResult) UnsetDepositFlowNumber()`

UnsetDepositFlowNumber ensures that no value is present for DepositFlowNumber, not even an explicit nil
### GetAttachmentId

`func (o *ChorusProResult) GetAttachmentId() int32`

GetAttachmentId returns the AttachmentId field if non-nil, zero value otherwise.

### GetAttachmentIdOk

`func (o *ChorusProResult) GetAttachmentIdOk() (*int32, bool)`

GetAttachmentIdOk returns a tuple with the AttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentId

`func (o *ChorusProResult) SetAttachmentId(v int32)`

SetAttachmentId sets AttachmentId field to given value.

### HasAttachmentId

`func (o *ChorusProResult) HasAttachmentId() bool`

HasAttachmentId returns a boolean if a field has been set.

### SetAttachmentIdNil

`func (o *ChorusProResult) SetAttachmentIdNil(b bool)`

 SetAttachmentIdNil sets the value for AttachmentId to be an explicit nil

### UnsetAttachmentId
`func (o *ChorusProResult) UnsetAttachmentId()`

UnsetAttachmentId ensures that no value is present for AttachmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



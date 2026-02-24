# LifecycleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoices** | Pointer to [**[]InvoiceLifecycle**](InvoiceLifecycle.md) | Cycles de vie par facture | [optional] 
**TotalInvoices** | **int32** | Nombre de factures | 
**CutoffDays** | **int32** | Nombre de jours de la fenetre de recherche | 

## Methods

### NewLifecycleResponse

`func NewLifecycleResponse(totalInvoices int32, cutoffDays int32, ) *LifecycleResponse`

NewLifecycleResponse instantiates a new LifecycleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLifecycleResponseWithDefaults

`func NewLifecycleResponseWithDefaults() *LifecycleResponse`

NewLifecycleResponseWithDefaults instantiates a new LifecycleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoices

`func (o *LifecycleResponse) GetInvoices() []InvoiceLifecycle`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *LifecycleResponse) GetInvoicesOk() (*[]InvoiceLifecycle, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *LifecycleResponse) SetInvoices(v []InvoiceLifecycle)`

SetInvoices sets Invoices field to given value.

### HasInvoices

`func (o *LifecycleResponse) HasInvoices() bool`

HasInvoices returns a boolean if a field has been set.

### GetTotalInvoices

`func (o *LifecycleResponse) GetTotalInvoices() int32`

GetTotalInvoices returns the TotalInvoices field if non-nil, zero value otherwise.

### GetTotalInvoicesOk

`func (o *LifecycleResponse) GetTotalInvoicesOk() (*int32, bool)`

GetTotalInvoicesOk returns a tuple with the TotalInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInvoices

`func (o *LifecycleResponse) SetTotalInvoices(v int32)`

SetTotalInvoices sets TotalInvoices field to given value.


### GetCutoffDays

`func (o *LifecycleResponse) GetCutoffDays() int32`

GetCutoffDays returns the CutoffDays field if non-nil, zero value otherwise.

### GetCutoffDaysOk

`func (o *LifecycleResponse) GetCutoffDaysOk() (*int32, bool)`

GetCutoffDaysOk returns a tuple with the CutoffDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutoffDays

`func (o *LifecycleResponse) SetCutoffDays(v int32)`

SetCutoffDays sets CutoffDays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



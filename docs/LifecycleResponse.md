# LifecycleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoices** | Pointer to [**[]InvoiceLifecycle**](InvoiceLifecycle.md) | Cycles de vie par facture | [optional] 
**TotalInvoices** | **int32** | Nombre de factures | 
**CutoffDays** | **int32** | Nombre de jours de la fenetre de recherche | 
**TotalFlows** | Pointer to **int32** | Nombre de flux de statut (CDAR) lus sur la PA | [optional] [default to 0]
**Truncated** | Pointer to **bool** | True si la limite de flux lus par appel a ete atteinte : la reponse est incomplete, reduire la fenetre (dateFrom/dateTo ou days) | [optional] [default to false]
**DateFrom** | Pointer to **NullableString** |  | [optional] 
**DateTo** | Pointer to **NullableString** |  | [optional] 

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


### GetTotalFlows

`func (o *LifecycleResponse) GetTotalFlows() int32`

GetTotalFlows returns the TotalFlows field if non-nil, zero value otherwise.

### GetTotalFlowsOk

`func (o *LifecycleResponse) GetTotalFlowsOk() (*int32, bool)`

GetTotalFlowsOk returns a tuple with the TotalFlows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFlows

`func (o *LifecycleResponse) SetTotalFlows(v int32)`

SetTotalFlows sets TotalFlows field to given value.

### HasTotalFlows

`func (o *LifecycleResponse) HasTotalFlows() bool`

HasTotalFlows returns a boolean if a field has been set.

### GetTruncated

`func (o *LifecycleResponse) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *LifecycleResponse) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *LifecycleResponse) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *LifecycleResponse) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.

### GetDateFrom

`func (o *LifecycleResponse) GetDateFrom() string`

GetDateFrom returns the DateFrom field if non-nil, zero value otherwise.

### GetDateFromOk

`func (o *LifecycleResponse) GetDateFromOk() (*string, bool)`

GetDateFromOk returns a tuple with the DateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFrom

`func (o *LifecycleResponse) SetDateFrom(v string)`

SetDateFrom sets DateFrom field to given value.

### HasDateFrom

`func (o *LifecycleResponse) HasDateFrom() bool`

HasDateFrom returns a boolean if a field has been set.

### SetDateFromNil

`func (o *LifecycleResponse) SetDateFromNil(b bool)`

 SetDateFromNil sets the value for DateFrom to be an explicit nil

### UnsetDateFrom
`func (o *LifecycleResponse) UnsetDateFrom()`

UnsetDateFrom ensures that no value is present for DateFrom, not even an explicit nil
### GetDateTo

`func (o *LifecycleResponse) GetDateTo() string`

GetDateTo returns the DateTo field if non-nil, zero value otherwise.

### GetDateToOk

`func (o *LifecycleResponse) GetDateToOk() (*string, bool)`

GetDateToOk returns a tuple with the DateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTo

`func (o *LifecycleResponse) SetDateTo(v string)`

SetDateTo sets DateTo field to given value.

### HasDateTo

`func (o *LifecycleResponse) HasDateTo() bool`

HasDateTo returns a boolean if a field has been set.

### SetDateToNil

`func (o *LifecycleResponse) SetDateToNil(b bool)`

 SetDateToNil sets the value for DateTo to be an explicit nil

### UnsetDateTo
`func (o *LifecycleResponse) UnsetDateTo()`

UnsetDateTo ensures that no value is present for DateTo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



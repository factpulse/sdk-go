# InvoiceLifecycle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SellerId** | Pointer to **NullableString** |  | [optional] 
**InvoiceId** | **string** | Reference de la facture (IssuerAssignedID du CDAR) | 
**Events** | Pointer to [**[]LifecycleEvent**](LifecycleEvent.md) | Evenements de cycle de vie tries chronologiquement | [optional] 
**TotalEvents** | **int32** | Nombre total d&#39;evenements | 

## Methods

### NewInvoiceLifecycle

`func NewInvoiceLifecycle(invoiceId string, totalEvents int32, ) *InvoiceLifecycle`

NewInvoiceLifecycle instantiates a new InvoiceLifecycle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLifecycleWithDefaults

`func NewInvoiceLifecycleWithDefaults() *InvoiceLifecycle`

NewInvoiceLifecycleWithDefaults instantiates a new InvoiceLifecycle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSellerId

`func (o *InvoiceLifecycle) GetSellerId() string`

GetSellerId returns the SellerId field if non-nil, zero value otherwise.

### GetSellerIdOk

`func (o *InvoiceLifecycle) GetSellerIdOk() (*string, bool)`

GetSellerIdOk returns a tuple with the SellerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerId

`func (o *InvoiceLifecycle) SetSellerId(v string)`

SetSellerId sets SellerId field to given value.

### HasSellerId

`func (o *InvoiceLifecycle) HasSellerId() bool`

HasSellerId returns a boolean if a field has been set.

### SetSellerIdNil

`func (o *InvoiceLifecycle) SetSellerIdNil(b bool)`

 SetSellerIdNil sets the value for SellerId to be an explicit nil

### UnsetSellerId
`func (o *InvoiceLifecycle) UnsetSellerId()`

UnsetSellerId ensures that no value is present for SellerId, not even an explicit nil
### GetInvoiceId

`func (o *InvoiceLifecycle) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *InvoiceLifecycle) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *InvoiceLifecycle) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetEvents

`func (o *InvoiceLifecycle) GetEvents() []LifecycleEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *InvoiceLifecycle) GetEventsOk() (*[]LifecycleEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *InvoiceLifecycle) SetEvents(v []LifecycleEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *InvoiceLifecycle) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTotalEvents

`func (o *InvoiceLifecycle) GetTotalEvents() int32`

GetTotalEvents returns the TotalEvents field if non-nil, zero value otherwise.

### GetTotalEventsOk

`func (o *InvoiceLifecycle) GetTotalEventsOk() (*int32, bool)`

GetTotalEventsOk returns a tuple with the TotalEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEvents

`func (o *InvoiceLifecycle) SetTotalEvents(v int32)`

SetTotalEvents sets TotalEvents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



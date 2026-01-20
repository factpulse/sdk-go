# SimplifiedCDARResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowId** | **string** | Identifiant du flux AFNOR | 
**DocumentId** | **string** | Identifiant du message CDAR généré | 
**Status** | **string** | Code statut soumis (210 ou 212) | 
**InvoiceId** | **string** | Identifiant de la facture | 
**Message** | **string** | Message de confirmation | 

## Methods

### NewSimplifiedCDARResponse

`func NewSimplifiedCDARResponse(flowId string, documentId string, status string, invoiceId string, message string, ) *SimplifiedCDARResponse`

NewSimplifiedCDARResponse instantiates a new SimplifiedCDARResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedCDARResponseWithDefaults

`func NewSimplifiedCDARResponseWithDefaults() *SimplifiedCDARResponse`

NewSimplifiedCDARResponseWithDefaults instantiates a new SimplifiedCDARResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowId

`func (o *SimplifiedCDARResponse) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *SimplifiedCDARResponse) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *SimplifiedCDARResponse) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.


### GetDocumentId

`func (o *SimplifiedCDARResponse) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *SimplifiedCDARResponse) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *SimplifiedCDARResponse) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.


### GetStatus

`func (o *SimplifiedCDARResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SimplifiedCDARResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SimplifiedCDARResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetInvoiceId

`func (o *SimplifiedCDARResponse) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *SimplifiedCDARResponse) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *SimplifiedCDARResponse) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetMessage

`func (o *SimplifiedCDARResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SimplifiedCDARResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SimplifiedCDARResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



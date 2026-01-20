# EncaisseeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** | Identifiant de la facture (BT-1) | 
**InvoiceIssueDate** | **string** | Date d&#39;émission de la facture (YYYY-MM-DD) | 
**SenderSiren** | Pointer to **NullableString** |  | [optional] 
**FlowType** | Pointer to **string** | Type de flux: SupplierInvoiceLC (acheteur) ou CustomerInvoiceLC (vendeur) | [optional] [default to "SupplierInvoiceLC"]
**PdpFlowServiceUrl** | Pointer to **NullableString** |  | [optional] 
**PdpTokenUrl** | Pointer to **NullableString** |  | [optional] 
**PdpClientId** | Pointer to **NullableString** |  | [optional] 
**PdpClientSecret** | Pointer to **NullableString** |  | [optional] 
**Amount** | [**Amount**](Amount.md) |  | 
**Currency** | Pointer to **string** | Code devise ISO 4217 | [optional] [default to "EUR"]

## Methods

### NewEncaisseeRequest

`func NewEncaisseeRequest(invoiceId string, invoiceIssueDate string, amount Amount, ) *EncaisseeRequest`

NewEncaisseeRequest instantiates a new EncaisseeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncaisseeRequestWithDefaults

`func NewEncaisseeRequestWithDefaults() *EncaisseeRequest`

NewEncaisseeRequestWithDefaults instantiates a new EncaisseeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *EncaisseeRequest) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *EncaisseeRequest) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *EncaisseeRequest) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetInvoiceIssueDate

`func (o *EncaisseeRequest) GetInvoiceIssueDate() string`

GetInvoiceIssueDate returns the InvoiceIssueDate field if non-nil, zero value otherwise.

### GetInvoiceIssueDateOk

`func (o *EncaisseeRequest) GetInvoiceIssueDateOk() (*string, bool)`

GetInvoiceIssueDateOk returns a tuple with the InvoiceIssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceIssueDate

`func (o *EncaisseeRequest) SetInvoiceIssueDate(v string)`

SetInvoiceIssueDate sets InvoiceIssueDate field to given value.


### GetSenderSiren

`func (o *EncaisseeRequest) GetSenderSiren() string`

GetSenderSiren returns the SenderSiren field if non-nil, zero value otherwise.

### GetSenderSirenOk

`func (o *EncaisseeRequest) GetSenderSirenOk() (*string, bool)`

GetSenderSirenOk returns a tuple with the SenderSiren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderSiren

`func (o *EncaisseeRequest) SetSenderSiren(v string)`

SetSenderSiren sets SenderSiren field to given value.

### HasSenderSiren

`func (o *EncaisseeRequest) HasSenderSiren() bool`

HasSenderSiren returns a boolean if a field has been set.

### SetSenderSirenNil

`func (o *EncaisseeRequest) SetSenderSirenNil(b bool)`

 SetSenderSirenNil sets the value for SenderSiren to be an explicit nil

### UnsetSenderSiren
`func (o *EncaisseeRequest) UnsetSenderSiren()`

UnsetSenderSiren ensures that no value is present for SenderSiren, not even an explicit nil
### GetFlowType

`func (o *EncaisseeRequest) GetFlowType() string`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *EncaisseeRequest) GetFlowTypeOk() (*string, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *EncaisseeRequest) SetFlowType(v string)`

SetFlowType sets FlowType field to given value.

### HasFlowType

`func (o *EncaisseeRequest) HasFlowType() bool`

HasFlowType returns a boolean if a field has been set.

### GetPdpFlowServiceUrl

`func (o *EncaisseeRequest) GetPdpFlowServiceUrl() string`

GetPdpFlowServiceUrl returns the PdpFlowServiceUrl field if non-nil, zero value otherwise.

### GetPdpFlowServiceUrlOk

`func (o *EncaisseeRequest) GetPdpFlowServiceUrlOk() (*string, bool)`

GetPdpFlowServiceUrlOk returns a tuple with the PdpFlowServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpFlowServiceUrl

`func (o *EncaisseeRequest) SetPdpFlowServiceUrl(v string)`

SetPdpFlowServiceUrl sets PdpFlowServiceUrl field to given value.

### HasPdpFlowServiceUrl

`func (o *EncaisseeRequest) HasPdpFlowServiceUrl() bool`

HasPdpFlowServiceUrl returns a boolean if a field has been set.

### SetPdpFlowServiceUrlNil

`func (o *EncaisseeRequest) SetPdpFlowServiceUrlNil(b bool)`

 SetPdpFlowServiceUrlNil sets the value for PdpFlowServiceUrl to be an explicit nil

### UnsetPdpFlowServiceUrl
`func (o *EncaisseeRequest) UnsetPdpFlowServiceUrl()`

UnsetPdpFlowServiceUrl ensures that no value is present for PdpFlowServiceUrl, not even an explicit nil
### GetPdpTokenUrl

`func (o *EncaisseeRequest) GetPdpTokenUrl() string`

GetPdpTokenUrl returns the PdpTokenUrl field if non-nil, zero value otherwise.

### GetPdpTokenUrlOk

`func (o *EncaisseeRequest) GetPdpTokenUrlOk() (*string, bool)`

GetPdpTokenUrlOk returns a tuple with the PdpTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpTokenUrl

`func (o *EncaisseeRequest) SetPdpTokenUrl(v string)`

SetPdpTokenUrl sets PdpTokenUrl field to given value.

### HasPdpTokenUrl

`func (o *EncaisseeRequest) HasPdpTokenUrl() bool`

HasPdpTokenUrl returns a boolean if a field has been set.

### SetPdpTokenUrlNil

`func (o *EncaisseeRequest) SetPdpTokenUrlNil(b bool)`

 SetPdpTokenUrlNil sets the value for PdpTokenUrl to be an explicit nil

### UnsetPdpTokenUrl
`func (o *EncaisseeRequest) UnsetPdpTokenUrl()`

UnsetPdpTokenUrl ensures that no value is present for PdpTokenUrl, not even an explicit nil
### GetPdpClientId

`func (o *EncaisseeRequest) GetPdpClientId() string`

GetPdpClientId returns the PdpClientId field if non-nil, zero value otherwise.

### GetPdpClientIdOk

`func (o *EncaisseeRequest) GetPdpClientIdOk() (*string, bool)`

GetPdpClientIdOk returns a tuple with the PdpClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpClientId

`func (o *EncaisseeRequest) SetPdpClientId(v string)`

SetPdpClientId sets PdpClientId field to given value.

### HasPdpClientId

`func (o *EncaisseeRequest) HasPdpClientId() bool`

HasPdpClientId returns a boolean if a field has been set.

### SetPdpClientIdNil

`func (o *EncaisseeRequest) SetPdpClientIdNil(b bool)`

 SetPdpClientIdNil sets the value for PdpClientId to be an explicit nil

### UnsetPdpClientId
`func (o *EncaisseeRequest) UnsetPdpClientId()`

UnsetPdpClientId ensures that no value is present for PdpClientId, not even an explicit nil
### GetPdpClientSecret

`func (o *EncaisseeRequest) GetPdpClientSecret() string`

GetPdpClientSecret returns the PdpClientSecret field if non-nil, zero value otherwise.

### GetPdpClientSecretOk

`func (o *EncaisseeRequest) GetPdpClientSecretOk() (*string, bool)`

GetPdpClientSecretOk returns a tuple with the PdpClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpClientSecret

`func (o *EncaisseeRequest) SetPdpClientSecret(v string)`

SetPdpClientSecret sets PdpClientSecret field to given value.

### HasPdpClientSecret

`func (o *EncaisseeRequest) HasPdpClientSecret() bool`

HasPdpClientSecret returns a boolean if a field has been set.

### SetPdpClientSecretNil

`func (o *EncaisseeRequest) SetPdpClientSecretNil(b bool)`

 SetPdpClientSecretNil sets the value for PdpClientSecret to be an explicit nil

### UnsetPdpClientSecret
`func (o *EncaisseeRequest) UnsetPdpClientSecret()`

UnsetPdpClientSecret ensures that no value is present for PdpClientSecret, not even an explicit nil
### GetAmount

`func (o *EncaisseeRequest) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EncaisseeRequest) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EncaisseeRequest) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *EncaisseeRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *EncaisseeRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *EncaisseeRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *EncaisseeRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



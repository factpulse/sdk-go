# RefuseeRequest

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
**ReasonCode** | **string** | Code motif du refus (obligatoire). Valeurs autorisées: TX_TVA_ERR, MONTANTTOTAL_ERR, CALCUL_ERR, NON_CONFORME, DOUBLON, DEST_ERR, TRANSAC_INC, EMMET_INC, CONTRAT_TERM, DOUBLE_FACT, CMD_ERR, ADR_ERR, REF_CT_ABSENT | 
**ReasonText** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRefuseeRequest

`func NewRefuseeRequest(invoiceId string, invoiceIssueDate string, reasonCode string, ) *RefuseeRequest`

NewRefuseeRequest instantiates a new RefuseeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefuseeRequestWithDefaults

`func NewRefuseeRequestWithDefaults() *RefuseeRequest`

NewRefuseeRequestWithDefaults instantiates a new RefuseeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *RefuseeRequest) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *RefuseeRequest) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *RefuseeRequest) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetInvoiceIssueDate

`func (o *RefuseeRequest) GetInvoiceIssueDate() string`

GetInvoiceIssueDate returns the InvoiceIssueDate field if non-nil, zero value otherwise.

### GetInvoiceIssueDateOk

`func (o *RefuseeRequest) GetInvoiceIssueDateOk() (*string, bool)`

GetInvoiceIssueDateOk returns a tuple with the InvoiceIssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceIssueDate

`func (o *RefuseeRequest) SetInvoiceIssueDate(v string)`

SetInvoiceIssueDate sets InvoiceIssueDate field to given value.


### GetSenderSiren

`func (o *RefuseeRequest) GetSenderSiren() string`

GetSenderSiren returns the SenderSiren field if non-nil, zero value otherwise.

### GetSenderSirenOk

`func (o *RefuseeRequest) GetSenderSirenOk() (*string, bool)`

GetSenderSirenOk returns a tuple with the SenderSiren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderSiren

`func (o *RefuseeRequest) SetSenderSiren(v string)`

SetSenderSiren sets SenderSiren field to given value.

### HasSenderSiren

`func (o *RefuseeRequest) HasSenderSiren() bool`

HasSenderSiren returns a boolean if a field has been set.

### SetSenderSirenNil

`func (o *RefuseeRequest) SetSenderSirenNil(b bool)`

 SetSenderSirenNil sets the value for SenderSiren to be an explicit nil

### UnsetSenderSiren
`func (o *RefuseeRequest) UnsetSenderSiren()`

UnsetSenderSiren ensures that no value is present for SenderSiren, not even an explicit nil
### GetFlowType

`func (o *RefuseeRequest) GetFlowType() string`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *RefuseeRequest) GetFlowTypeOk() (*string, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *RefuseeRequest) SetFlowType(v string)`

SetFlowType sets FlowType field to given value.

### HasFlowType

`func (o *RefuseeRequest) HasFlowType() bool`

HasFlowType returns a boolean if a field has been set.

### GetPdpFlowServiceUrl

`func (o *RefuseeRequest) GetPdpFlowServiceUrl() string`

GetPdpFlowServiceUrl returns the PdpFlowServiceUrl field if non-nil, zero value otherwise.

### GetPdpFlowServiceUrlOk

`func (o *RefuseeRequest) GetPdpFlowServiceUrlOk() (*string, bool)`

GetPdpFlowServiceUrlOk returns a tuple with the PdpFlowServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpFlowServiceUrl

`func (o *RefuseeRequest) SetPdpFlowServiceUrl(v string)`

SetPdpFlowServiceUrl sets PdpFlowServiceUrl field to given value.

### HasPdpFlowServiceUrl

`func (o *RefuseeRequest) HasPdpFlowServiceUrl() bool`

HasPdpFlowServiceUrl returns a boolean if a field has been set.

### SetPdpFlowServiceUrlNil

`func (o *RefuseeRequest) SetPdpFlowServiceUrlNil(b bool)`

 SetPdpFlowServiceUrlNil sets the value for PdpFlowServiceUrl to be an explicit nil

### UnsetPdpFlowServiceUrl
`func (o *RefuseeRequest) UnsetPdpFlowServiceUrl()`

UnsetPdpFlowServiceUrl ensures that no value is present for PdpFlowServiceUrl, not even an explicit nil
### GetPdpTokenUrl

`func (o *RefuseeRequest) GetPdpTokenUrl() string`

GetPdpTokenUrl returns the PdpTokenUrl field if non-nil, zero value otherwise.

### GetPdpTokenUrlOk

`func (o *RefuseeRequest) GetPdpTokenUrlOk() (*string, bool)`

GetPdpTokenUrlOk returns a tuple with the PdpTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpTokenUrl

`func (o *RefuseeRequest) SetPdpTokenUrl(v string)`

SetPdpTokenUrl sets PdpTokenUrl field to given value.

### HasPdpTokenUrl

`func (o *RefuseeRequest) HasPdpTokenUrl() bool`

HasPdpTokenUrl returns a boolean if a field has been set.

### SetPdpTokenUrlNil

`func (o *RefuseeRequest) SetPdpTokenUrlNil(b bool)`

 SetPdpTokenUrlNil sets the value for PdpTokenUrl to be an explicit nil

### UnsetPdpTokenUrl
`func (o *RefuseeRequest) UnsetPdpTokenUrl()`

UnsetPdpTokenUrl ensures that no value is present for PdpTokenUrl, not even an explicit nil
### GetPdpClientId

`func (o *RefuseeRequest) GetPdpClientId() string`

GetPdpClientId returns the PdpClientId field if non-nil, zero value otherwise.

### GetPdpClientIdOk

`func (o *RefuseeRequest) GetPdpClientIdOk() (*string, bool)`

GetPdpClientIdOk returns a tuple with the PdpClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpClientId

`func (o *RefuseeRequest) SetPdpClientId(v string)`

SetPdpClientId sets PdpClientId field to given value.

### HasPdpClientId

`func (o *RefuseeRequest) HasPdpClientId() bool`

HasPdpClientId returns a boolean if a field has been set.

### SetPdpClientIdNil

`func (o *RefuseeRequest) SetPdpClientIdNil(b bool)`

 SetPdpClientIdNil sets the value for PdpClientId to be an explicit nil

### UnsetPdpClientId
`func (o *RefuseeRequest) UnsetPdpClientId()`

UnsetPdpClientId ensures that no value is present for PdpClientId, not even an explicit nil
### GetPdpClientSecret

`func (o *RefuseeRequest) GetPdpClientSecret() string`

GetPdpClientSecret returns the PdpClientSecret field if non-nil, zero value otherwise.

### GetPdpClientSecretOk

`func (o *RefuseeRequest) GetPdpClientSecretOk() (*string, bool)`

GetPdpClientSecretOk returns a tuple with the PdpClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpClientSecret

`func (o *RefuseeRequest) SetPdpClientSecret(v string)`

SetPdpClientSecret sets PdpClientSecret field to given value.

### HasPdpClientSecret

`func (o *RefuseeRequest) HasPdpClientSecret() bool`

HasPdpClientSecret returns a boolean if a field has been set.

### SetPdpClientSecretNil

`func (o *RefuseeRequest) SetPdpClientSecretNil(b bool)`

 SetPdpClientSecretNil sets the value for PdpClientSecret to be an explicit nil

### UnsetPdpClientSecret
`func (o *RefuseeRequest) UnsetPdpClientSecret()`

UnsetPdpClientSecret ensures that no value is present for PdpClientSecret, not even an explicit nil
### GetReasonCode

`func (o *RefuseeRequest) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *RefuseeRequest) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *RefuseeRequest) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.


### GetReasonText

`func (o *RefuseeRequest) GetReasonText() string`

GetReasonText returns the ReasonText field if non-nil, zero value otherwise.

### GetReasonTextOk

`func (o *RefuseeRequest) GetReasonTextOk() (*string, bool)`

GetReasonTextOk returns a tuple with the ReasonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonText

`func (o *RefuseeRequest) SetReasonText(v string)`

SetReasonText sets ReasonText field to given value.

### HasReasonText

`func (o *RefuseeRequest) HasReasonText() bool`

HasReasonText returns a boolean if a field has been set.

### SetReasonTextNil

`func (o *RefuseeRequest) SetReasonTextNil(b bool)`

 SetReasonTextNil sets the value for ReasonText to be an explicit nil

### UnsetReasonText
`func (o *RefuseeRequest) UnsetReasonText()`

UnsetReasonText ensures that no value is present for ReasonText, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



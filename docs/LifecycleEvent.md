# LifecycleEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowId** | **string** | ID du flux lifecycle | 
**StatusCode** | **string** | Code statut (200-601) | 
**StatusDescription** | Pointer to **NullableString** |  | [optional] 
**AckStatus** | Pointer to **NullableString** |  | [optional] 
**At** | Pointer to **NullableString** |  | [optional] 
**DocumentId** | Pointer to **NullableString** |  | [optional] 
**Amount** | Pointer to **NullableString** |  | [optional] 
**Currency** | Pointer to **NullableString** |  | [optional] 
**IssuerSiren** | Pointer to **NullableString** |  | [optional] 
**IssuerRole** | Pointer to **NullableString** |  | [optional] 
**ReasonCode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLifecycleEvent

`func NewLifecycleEvent(flowId string, statusCode string, ) *LifecycleEvent`

NewLifecycleEvent instantiates a new LifecycleEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLifecycleEventWithDefaults

`func NewLifecycleEventWithDefaults() *LifecycleEvent`

NewLifecycleEventWithDefaults instantiates a new LifecycleEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowId

`func (o *LifecycleEvent) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *LifecycleEvent) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *LifecycleEvent) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.


### GetStatusCode

`func (o *LifecycleEvent) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *LifecycleEvent) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *LifecycleEvent) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.


### GetStatusDescription

`func (o *LifecycleEvent) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *LifecycleEvent) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *LifecycleEvent) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *LifecycleEvent) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.

### SetStatusDescriptionNil

`func (o *LifecycleEvent) SetStatusDescriptionNil(b bool)`

 SetStatusDescriptionNil sets the value for StatusDescription to be an explicit nil

### UnsetStatusDescription
`func (o *LifecycleEvent) UnsetStatusDescription()`

UnsetStatusDescription ensures that no value is present for StatusDescription, not even an explicit nil
### GetAckStatus

`func (o *LifecycleEvent) GetAckStatus() string`

GetAckStatus returns the AckStatus field if non-nil, zero value otherwise.

### GetAckStatusOk

`func (o *LifecycleEvent) GetAckStatusOk() (*string, bool)`

GetAckStatusOk returns a tuple with the AckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckStatus

`func (o *LifecycleEvent) SetAckStatus(v string)`

SetAckStatus sets AckStatus field to given value.

### HasAckStatus

`func (o *LifecycleEvent) HasAckStatus() bool`

HasAckStatus returns a boolean if a field has been set.

### SetAckStatusNil

`func (o *LifecycleEvent) SetAckStatusNil(b bool)`

 SetAckStatusNil sets the value for AckStatus to be an explicit nil

### UnsetAckStatus
`func (o *LifecycleEvent) UnsetAckStatus()`

UnsetAckStatus ensures that no value is present for AckStatus, not even an explicit nil
### GetAt

`func (o *LifecycleEvent) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *LifecycleEvent) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *LifecycleEvent) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *LifecycleEvent) HasAt() bool`

HasAt returns a boolean if a field has been set.

### SetAtNil

`func (o *LifecycleEvent) SetAtNil(b bool)`

 SetAtNil sets the value for At to be an explicit nil

### UnsetAt
`func (o *LifecycleEvent) UnsetAt()`

UnsetAt ensures that no value is present for At, not even an explicit nil
### GetDocumentId

`func (o *LifecycleEvent) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *LifecycleEvent) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *LifecycleEvent) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *LifecycleEvent) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### SetDocumentIdNil

`func (o *LifecycleEvent) SetDocumentIdNil(b bool)`

 SetDocumentIdNil sets the value for DocumentId to be an explicit nil

### UnsetDocumentId
`func (o *LifecycleEvent) UnsetDocumentId()`

UnsetDocumentId ensures that no value is present for DocumentId, not even an explicit nil
### GetAmount

`func (o *LifecycleEvent) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *LifecycleEvent) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *LifecycleEvent) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *LifecycleEvent) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *LifecycleEvent) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *LifecycleEvent) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetCurrency

`func (o *LifecycleEvent) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *LifecycleEvent) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *LifecycleEvent) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *LifecycleEvent) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *LifecycleEvent) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *LifecycleEvent) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetIssuerSiren

`func (o *LifecycleEvent) GetIssuerSiren() string`

GetIssuerSiren returns the IssuerSiren field if non-nil, zero value otherwise.

### GetIssuerSirenOk

`func (o *LifecycleEvent) GetIssuerSirenOk() (*string, bool)`

GetIssuerSirenOk returns a tuple with the IssuerSiren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerSiren

`func (o *LifecycleEvent) SetIssuerSiren(v string)`

SetIssuerSiren sets IssuerSiren field to given value.

### HasIssuerSiren

`func (o *LifecycleEvent) HasIssuerSiren() bool`

HasIssuerSiren returns a boolean if a field has been set.

### SetIssuerSirenNil

`func (o *LifecycleEvent) SetIssuerSirenNil(b bool)`

 SetIssuerSirenNil sets the value for IssuerSiren to be an explicit nil

### UnsetIssuerSiren
`func (o *LifecycleEvent) UnsetIssuerSiren()`

UnsetIssuerSiren ensures that no value is present for IssuerSiren, not even an explicit nil
### GetIssuerRole

`func (o *LifecycleEvent) GetIssuerRole() string`

GetIssuerRole returns the IssuerRole field if non-nil, zero value otherwise.

### GetIssuerRoleOk

`func (o *LifecycleEvent) GetIssuerRoleOk() (*string, bool)`

GetIssuerRoleOk returns a tuple with the IssuerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerRole

`func (o *LifecycleEvent) SetIssuerRole(v string)`

SetIssuerRole sets IssuerRole field to given value.

### HasIssuerRole

`func (o *LifecycleEvent) HasIssuerRole() bool`

HasIssuerRole returns a boolean if a field has been set.

### SetIssuerRoleNil

`func (o *LifecycleEvent) SetIssuerRoleNil(b bool)`

 SetIssuerRoleNil sets the value for IssuerRole to be an explicit nil

### UnsetIssuerRole
`func (o *LifecycleEvent) UnsetIssuerRole()`

UnsetIssuerRole ensures that no value is present for IssuerRole, not even an explicit nil
### GetReasonCode

`func (o *LifecycleEvent) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *LifecycleEvent) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *LifecycleEvent) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *LifecycleEvent) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### SetReasonCodeNil

`func (o *LifecycleEvent) SetReasonCodeNil(b bool)`

 SetReasonCodeNil sets the value for ReasonCode to be an explicit nil

### UnsetReasonCode
`func (o *LifecycleEvent) UnsetReasonCode()`

UnsetReasonCode ensures that no value is present for ReasonCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



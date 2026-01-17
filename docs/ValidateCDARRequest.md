# ValidateCDARRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | Pointer to **NullableString** |  | [optional] 
**SenderSiren** | Pointer to **NullableString** |  | [optional] 
**SenderRole** | Pointer to **NullableString** |  | [optional] 
**InvoiceId** | Pointer to **NullableString** |  | [optional] 
**InvoiceIssueDate** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**ReasonCode** | Pointer to **NullableString** |  | [optional] 
**EncaisseAmount** | Pointer to [**NullableEncaisseamount1**](Encaisseamount1.md) |  | [optional] 

## Methods

### NewValidateCDARRequest

`func NewValidateCDARRequest() *ValidateCDARRequest`

NewValidateCDARRequest instantiates a new ValidateCDARRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateCDARRequestWithDefaults

`func NewValidateCDARRequestWithDefaults() *ValidateCDARRequest`

NewValidateCDARRequestWithDefaults instantiates a new ValidateCDARRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *ValidateCDARRequest) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *ValidateCDARRequest) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *ValidateCDARRequest) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *ValidateCDARRequest) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### SetDocumentIdNil

`func (o *ValidateCDARRequest) SetDocumentIdNil(b bool)`

 SetDocumentIdNil sets the value for DocumentId to be an explicit nil

### UnsetDocumentId
`func (o *ValidateCDARRequest) UnsetDocumentId()`

UnsetDocumentId ensures that no value is present for DocumentId, not even an explicit nil
### GetSenderSiren

`func (o *ValidateCDARRequest) GetSenderSiren() string`

GetSenderSiren returns the SenderSiren field if non-nil, zero value otherwise.

### GetSenderSirenOk

`func (o *ValidateCDARRequest) GetSenderSirenOk() (*string, bool)`

GetSenderSirenOk returns a tuple with the SenderSiren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderSiren

`func (o *ValidateCDARRequest) SetSenderSiren(v string)`

SetSenderSiren sets SenderSiren field to given value.

### HasSenderSiren

`func (o *ValidateCDARRequest) HasSenderSiren() bool`

HasSenderSiren returns a boolean if a field has been set.

### SetSenderSirenNil

`func (o *ValidateCDARRequest) SetSenderSirenNil(b bool)`

 SetSenderSirenNil sets the value for SenderSiren to be an explicit nil

### UnsetSenderSiren
`func (o *ValidateCDARRequest) UnsetSenderSiren()`

UnsetSenderSiren ensures that no value is present for SenderSiren, not even an explicit nil
### GetSenderRole

`func (o *ValidateCDARRequest) GetSenderRole() string`

GetSenderRole returns the SenderRole field if non-nil, zero value otherwise.

### GetSenderRoleOk

`func (o *ValidateCDARRequest) GetSenderRoleOk() (*string, bool)`

GetSenderRoleOk returns a tuple with the SenderRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderRole

`func (o *ValidateCDARRequest) SetSenderRole(v string)`

SetSenderRole sets SenderRole field to given value.

### HasSenderRole

`func (o *ValidateCDARRequest) HasSenderRole() bool`

HasSenderRole returns a boolean if a field has been set.

### SetSenderRoleNil

`func (o *ValidateCDARRequest) SetSenderRoleNil(b bool)`

 SetSenderRoleNil sets the value for SenderRole to be an explicit nil

### UnsetSenderRole
`func (o *ValidateCDARRequest) UnsetSenderRole()`

UnsetSenderRole ensures that no value is present for SenderRole, not even an explicit nil
### GetInvoiceId

`func (o *ValidateCDARRequest) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *ValidateCDARRequest) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *ValidateCDARRequest) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *ValidateCDARRequest) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### SetInvoiceIdNil

`func (o *ValidateCDARRequest) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *ValidateCDARRequest) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetInvoiceIssueDate

`func (o *ValidateCDARRequest) GetInvoiceIssueDate() string`

GetInvoiceIssueDate returns the InvoiceIssueDate field if non-nil, zero value otherwise.

### GetInvoiceIssueDateOk

`func (o *ValidateCDARRequest) GetInvoiceIssueDateOk() (*string, bool)`

GetInvoiceIssueDateOk returns a tuple with the InvoiceIssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceIssueDate

`func (o *ValidateCDARRequest) SetInvoiceIssueDate(v string)`

SetInvoiceIssueDate sets InvoiceIssueDate field to given value.

### HasInvoiceIssueDate

`func (o *ValidateCDARRequest) HasInvoiceIssueDate() bool`

HasInvoiceIssueDate returns a boolean if a field has been set.

### SetInvoiceIssueDateNil

`func (o *ValidateCDARRequest) SetInvoiceIssueDateNil(b bool)`

 SetInvoiceIssueDateNil sets the value for InvoiceIssueDate to be an explicit nil

### UnsetInvoiceIssueDate
`func (o *ValidateCDARRequest) UnsetInvoiceIssueDate()`

UnsetInvoiceIssueDate ensures that no value is present for InvoiceIssueDate, not even an explicit nil
### GetStatus

`func (o *ValidateCDARRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ValidateCDARRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ValidateCDARRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ValidateCDARRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ValidateCDARRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ValidateCDARRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetReasonCode

`func (o *ValidateCDARRequest) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *ValidateCDARRequest) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *ValidateCDARRequest) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *ValidateCDARRequest) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### SetReasonCodeNil

`func (o *ValidateCDARRequest) SetReasonCodeNil(b bool)`

 SetReasonCodeNil sets the value for ReasonCode to be an explicit nil

### UnsetReasonCode
`func (o *ValidateCDARRequest) UnsetReasonCode()`

UnsetReasonCode ensures that no value is present for ReasonCode, not even an explicit nil
### GetEncaisseAmount

`func (o *ValidateCDARRequest) GetEncaisseAmount() Encaisseamount1`

GetEncaisseAmount returns the EncaisseAmount field if non-nil, zero value otherwise.

### GetEncaisseAmountOk

`func (o *ValidateCDARRequest) GetEncaisseAmountOk() (*Encaisseamount1, bool)`

GetEncaisseAmountOk returns a tuple with the EncaisseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncaisseAmount

`func (o *ValidateCDARRequest) SetEncaisseAmount(v Encaisseamount1)`

SetEncaisseAmount sets EncaisseAmount field to given value.

### HasEncaisseAmount

`func (o *ValidateCDARRequest) HasEncaisseAmount() bool`

HasEncaisseAmount returns a boolean if a field has been set.

### SetEncaisseAmountNil

`func (o *ValidateCDARRequest) SetEncaisseAmountNil(b bool)`

 SetEncaisseAmountNil sets the value for EncaisseAmount to be an explicit nil

### UnsetEncaisseAmount
`func (o *ValidateCDARRequest) UnsetEncaisseAmount()`

UnsetEncaisseAmount ensures that no value is present for EncaisseAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



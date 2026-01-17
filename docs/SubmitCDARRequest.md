# SubmitCDARRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | **string** | Identifiant unique du message CDAR | 
**BusinessProcess** | Pointer to **string** | Code processus métier (REGULATED, B2C, B2BINT, etc.) | [optional] [default to "REGULATED"]
**TypeCode** | Pointer to **string** | Type de message (23&#x3D;Traitement, 305&#x3D;Transmission) | [optional] [default to "23"]
**SenderSiren** | **string** | SIREN de l&#39;émetteur (9 chiffres) | 
**SenderRole** | Pointer to **string** | Rôle de l&#39;émetteur (WK, SE, BY, etc.) | [optional] [default to "WK"]
**SenderName** | Pointer to **NullableString** |  | [optional] 
**SenderEmail** | Pointer to **NullableString** |  | [optional] 
**Recipients** | Pointer to [**[]RecipientInput**](RecipientInput.md) | Liste des destinataires | [optional] 
**InvoiceId** | **string** | Identifiant de la facture (BT-1) | 
**InvoiceIssueDate** | **string** | Date d&#39;émission de la facture (YYYY-MM-DD) | 
**InvoiceTypeCode** | Pointer to **string** | Type de document (380&#x3D;Facture, 381&#x3D;Avoir) | [optional] [default to "380"]
**InvoiceSellerSiren** | Pointer to **NullableString** |  | [optional] 
**InvoiceBuyerSiren** | Pointer to **NullableString** |  | [optional] 
**Status** | **string** | Code statut de la facture (200-601) | 
**ReasonCode** | Pointer to **NullableString** |  | [optional] 
**ReasonText** | Pointer to **NullableString** |  | [optional] 
**ActionCode** | Pointer to **NullableString** |  | [optional] 
**EncaisseAmount** | Pointer to [**NullableEncaisseamount**](Encaisseamount.md) |  | [optional] 
**FlowType** | Pointer to **string** | Type de flux AFNOR (CustomerInvoiceLC, SupplierInvoiceLC, etc.) | [optional] [default to "CustomerInvoiceLC"]

## Methods

### NewSubmitCDARRequest

`func NewSubmitCDARRequest(documentId string, senderSiren string, invoiceId string, invoiceIssueDate string, status string, ) *SubmitCDARRequest`

NewSubmitCDARRequest instantiates a new SubmitCDARRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitCDARRequestWithDefaults

`func NewSubmitCDARRequestWithDefaults() *SubmitCDARRequest`

NewSubmitCDARRequestWithDefaults instantiates a new SubmitCDARRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *SubmitCDARRequest) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *SubmitCDARRequest) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *SubmitCDARRequest) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.


### GetBusinessProcess

`func (o *SubmitCDARRequest) GetBusinessProcess() string`

GetBusinessProcess returns the BusinessProcess field if non-nil, zero value otherwise.

### GetBusinessProcessOk

`func (o *SubmitCDARRequest) GetBusinessProcessOk() (*string, bool)`

GetBusinessProcessOk returns a tuple with the BusinessProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProcess

`func (o *SubmitCDARRequest) SetBusinessProcess(v string)`

SetBusinessProcess sets BusinessProcess field to given value.

### HasBusinessProcess

`func (o *SubmitCDARRequest) HasBusinessProcess() bool`

HasBusinessProcess returns a boolean if a field has been set.

### GetTypeCode

`func (o *SubmitCDARRequest) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *SubmitCDARRequest) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *SubmitCDARRequest) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.

### HasTypeCode

`func (o *SubmitCDARRequest) HasTypeCode() bool`

HasTypeCode returns a boolean if a field has been set.

### GetSenderSiren

`func (o *SubmitCDARRequest) GetSenderSiren() string`

GetSenderSiren returns the SenderSiren field if non-nil, zero value otherwise.

### GetSenderSirenOk

`func (o *SubmitCDARRequest) GetSenderSirenOk() (*string, bool)`

GetSenderSirenOk returns a tuple with the SenderSiren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderSiren

`func (o *SubmitCDARRequest) SetSenderSiren(v string)`

SetSenderSiren sets SenderSiren field to given value.


### GetSenderRole

`func (o *SubmitCDARRequest) GetSenderRole() string`

GetSenderRole returns the SenderRole field if non-nil, zero value otherwise.

### GetSenderRoleOk

`func (o *SubmitCDARRequest) GetSenderRoleOk() (*string, bool)`

GetSenderRoleOk returns a tuple with the SenderRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderRole

`func (o *SubmitCDARRequest) SetSenderRole(v string)`

SetSenderRole sets SenderRole field to given value.

### HasSenderRole

`func (o *SubmitCDARRequest) HasSenderRole() bool`

HasSenderRole returns a boolean if a field has been set.

### GetSenderName

`func (o *SubmitCDARRequest) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *SubmitCDARRequest) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *SubmitCDARRequest) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.

### HasSenderName

`func (o *SubmitCDARRequest) HasSenderName() bool`

HasSenderName returns a boolean if a field has been set.

### SetSenderNameNil

`func (o *SubmitCDARRequest) SetSenderNameNil(b bool)`

 SetSenderNameNil sets the value for SenderName to be an explicit nil

### UnsetSenderName
`func (o *SubmitCDARRequest) UnsetSenderName()`

UnsetSenderName ensures that no value is present for SenderName, not even an explicit nil
### GetSenderEmail

`func (o *SubmitCDARRequest) GetSenderEmail() string`

GetSenderEmail returns the SenderEmail field if non-nil, zero value otherwise.

### GetSenderEmailOk

`func (o *SubmitCDARRequest) GetSenderEmailOk() (*string, bool)`

GetSenderEmailOk returns a tuple with the SenderEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderEmail

`func (o *SubmitCDARRequest) SetSenderEmail(v string)`

SetSenderEmail sets SenderEmail field to given value.

### HasSenderEmail

`func (o *SubmitCDARRequest) HasSenderEmail() bool`

HasSenderEmail returns a boolean if a field has been set.

### SetSenderEmailNil

`func (o *SubmitCDARRequest) SetSenderEmailNil(b bool)`

 SetSenderEmailNil sets the value for SenderEmail to be an explicit nil

### UnsetSenderEmail
`func (o *SubmitCDARRequest) UnsetSenderEmail()`

UnsetSenderEmail ensures that no value is present for SenderEmail, not even an explicit nil
### GetRecipients

`func (o *SubmitCDARRequest) GetRecipients() []RecipientInput`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *SubmitCDARRequest) GetRecipientsOk() (*[]RecipientInput, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *SubmitCDARRequest) SetRecipients(v []RecipientInput)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *SubmitCDARRequest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetInvoiceId

`func (o *SubmitCDARRequest) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *SubmitCDARRequest) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *SubmitCDARRequest) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetInvoiceIssueDate

`func (o *SubmitCDARRequest) GetInvoiceIssueDate() string`

GetInvoiceIssueDate returns the InvoiceIssueDate field if non-nil, zero value otherwise.

### GetInvoiceIssueDateOk

`func (o *SubmitCDARRequest) GetInvoiceIssueDateOk() (*string, bool)`

GetInvoiceIssueDateOk returns a tuple with the InvoiceIssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceIssueDate

`func (o *SubmitCDARRequest) SetInvoiceIssueDate(v string)`

SetInvoiceIssueDate sets InvoiceIssueDate field to given value.


### GetInvoiceTypeCode

`func (o *SubmitCDARRequest) GetInvoiceTypeCode() string`

GetInvoiceTypeCode returns the InvoiceTypeCode field if non-nil, zero value otherwise.

### GetInvoiceTypeCodeOk

`func (o *SubmitCDARRequest) GetInvoiceTypeCodeOk() (*string, bool)`

GetInvoiceTypeCodeOk returns a tuple with the InvoiceTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTypeCode

`func (o *SubmitCDARRequest) SetInvoiceTypeCode(v string)`

SetInvoiceTypeCode sets InvoiceTypeCode field to given value.

### HasInvoiceTypeCode

`func (o *SubmitCDARRequest) HasInvoiceTypeCode() bool`

HasInvoiceTypeCode returns a boolean if a field has been set.

### GetInvoiceSellerSiren

`func (o *SubmitCDARRequest) GetInvoiceSellerSiren() string`

GetInvoiceSellerSiren returns the InvoiceSellerSiren field if non-nil, zero value otherwise.

### GetInvoiceSellerSirenOk

`func (o *SubmitCDARRequest) GetInvoiceSellerSirenOk() (*string, bool)`

GetInvoiceSellerSirenOk returns a tuple with the InvoiceSellerSiren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceSellerSiren

`func (o *SubmitCDARRequest) SetInvoiceSellerSiren(v string)`

SetInvoiceSellerSiren sets InvoiceSellerSiren field to given value.

### HasInvoiceSellerSiren

`func (o *SubmitCDARRequest) HasInvoiceSellerSiren() bool`

HasInvoiceSellerSiren returns a boolean if a field has been set.

### SetInvoiceSellerSirenNil

`func (o *SubmitCDARRequest) SetInvoiceSellerSirenNil(b bool)`

 SetInvoiceSellerSirenNil sets the value for InvoiceSellerSiren to be an explicit nil

### UnsetInvoiceSellerSiren
`func (o *SubmitCDARRequest) UnsetInvoiceSellerSiren()`

UnsetInvoiceSellerSiren ensures that no value is present for InvoiceSellerSiren, not even an explicit nil
### GetInvoiceBuyerSiren

`func (o *SubmitCDARRequest) GetInvoiceBuyerSiren() string`

GetInvoiceBuyerSiren returns the InvoiceBuyerSiren field if non-nil, zero value otherwise.

### GetInvoiceBuyerSirenOk

`func (o *SubmitCDARRequest) GetInvoiceBuyerSirenOk() (*string, bool)`

GetInvoiceBuyerSirenOk returns a tuple with the InvoiceBuyerSiren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceBuyerSiren

`func (o *SubmitCDARRequest) SetInvoiceBuyerSiren(v string)`

SetInvoiceBuyerSiren sets InvoiceBuyerSiren field to given value.

### HasInvoiceBuyerSiren

`func (o *SubmitCDARRequest) HasInvoiceBuyerSiren() bool`

HasInvoiceBuyerSiren returns a boolean if a field has been set.

### SetInvoiceBuyerSirenNil

`func (o *SubmitCDARRequest) SetInvoiceBuyerSirenNil(b bool)`

 SetInvoiceBuyerSirenNil sets the value for InvoiceBuyerSiren to be an explicit nil

### UnsetInvoiceBuyerSiren
`func (o *SubmitCDARRequest) UnsetInvoiceBuyerSiren()`

UnsetInvoiceBuyerSiren ensures that no value is present for InvoiceBuyerSiren, not even an explicit nil
### GetStatus

`func (o *SubmitCDARRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubmitCDARRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubmitCDARRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReasonCode

`func (o *SubmitCDARRequest) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *SubmitCDARRequest) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *SubmitCDARRequest) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *SubmitCDARRequest) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### SetReasonCodeNil

`func (o *SubmitCDARRequest) SetReasonCodeNil(b bool)`

 SetReasonCodeNil sets the value for ReasonCode to be an explicit nil

### UnsetReasonCode
`func (o *SubmitCDARRequest) UnsetReasonCode()`

UnsetReasonCode ensures that no value is present for ReasonCode, not even an explicit nil
### GetReasonText

`func (o *SubmitCDARRequest) GetReasonText() string`

GetReasonText returns the ReasonText field if non-nil, zero value otherwise.

### GetReasonTextOk

`func (o *SubmitCDARRequest) GetReasonTextOk() (*string, bool)`

GetReasonTextOk returns a tuple with the ReasonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonText

`func (o *SubmitCDARRequest) SetReasonText(v string)`

SetReasonText sets ReasonText field to given value.

### HasReasonText

`func (o *SubmitCDARRequest) HasReasonText() bool`

HasReasonText returns a boolean if a field has been set.

### SetReasonTextNil

`func (o *SubmitCDARRequest) SetReasonTextNil(b bool)`

 SetReasonTextNil sets the value for ReasonText to be an explicit nil

### UnsetReasonText
`func (o *SubmitCDARRequest) UnsetReasonText()`

UnsetReasonText ensures that no value is present for ReasonText, not even an explicit nil
### GetActionCode

`func (o *SubmitCDARRequest) GetActionCode() string`

GetActionCode returns the ActionCode field if non-nil, zero value otherwise.

### GetActionCodeOk

`func (o *SubmitCDARRequest) GetActionCodeOk() (*string, bool)`

GetActionCodeOk returns a tuple with the ActionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionCode

`func (o *SubmitCDARRequest) SetActionCode(v string)`

SetActionCode sets ActionCode field to given value.

### HasActionCode

`func (o *SubmitCDARRequest) HasActionCode() bool`

HasActionCode returns a boolean if a field has been set.

### SetActionCodeNil

`func (o *SubmitCDARRequest) SetActionCodeNil(b bool)`

 SetActionCodeNil sets the value for ActionCode to be an explicit nil

### UnsetActionCode
`func (o *SubmitCDARRequest) UnsetActionCode()`

UnsetActionCode ensures that no value is present for ActionCode, not even an explicit nil
### GetEncaisseAmount

`func (o *SubmitCDARRequest) GetEncaisseAmount() Encaisseamount`

GetEncaisseAmount returns the EncaisseAmount field if non-nil, zero value otherwise.

### GetEncaisseAmountOk

`func (o *SubmitCDARRequest) GetEncaisseAmountOk() (*Encaisseamount, bool)`

GetEncaisseAmountOk returns a tuple with the EncaisseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncaisseAmount

`func (o *SubmitCDARRequest) SetEncaisseAmount(v Encaisseamount)`

SetEncaisseAmount sets EncaisseAmount field to given value.

### HasEncaisseAmount

`func (o *SubmitCDARRequest) HasEncaisseAmount() bool`

HasEncaisseAmount returns a boolean if a field has been set.

### SetEncaisseAmountNil

`func (o *SubmitCDARRequest) SetEncaisseAmountNil(b bool)`

 SetEncaisseAmountNil sets the value for EncaisseAmount to be an explicit nil

### UnsetEncaisseAmount
`func (o *SubmitCDARRequest) UnsetEncaisseAmount()`

UnsetEncaisseAmount ensures that no value is present for EncaisseAmount, not even an explicit nil
### GetFlowType

`func (o *SubmitCDARRequest) GetFlowType() string`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *SubmitCDARRequest) GetFlowTypeOk() (*string, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *SubmitCDARRequest) SetFlowType(v string)`

SetFlowType sets FlowType field to given value.

### HasFlowType

`func (o *SubmitCDARRequest) HasFlowType() bool`

HasFlowType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



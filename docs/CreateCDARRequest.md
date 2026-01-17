# CreateCDARRequest

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

## Methods

### NewCreateCDARRequest

`func NewCreateCDARRequest(documentId string, senderSiren string, invoiceId string, invoiceIssueDate string, status string, ) *CreateCDARRequest`

NewCreateCDARRequest instantiates a new CreateCDARRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCDARRequestWithDefaults

`func NewCreateCDARRequestWithDefaults() *CreateCDARRequest`

NewCreateCDARRequestWithDefaults instantiates a new CreateCDARRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *CreateCDARRequest) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *CreateCDARRequest) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *CreateCDARRequest) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.


### GetBusinessProcess

`func (o *CreateCDARRequest) GetBusinessProcess() string`

GetBusinessProcess returns the BusinessProcess field if non-nil, zero value otherwise.

### GetBusinessProcessOk

`func (o *CreateCDARRequest) GetBusinessProcessOk() (*string, bool)`

GetBusinessProcessOk returns a tuple with the BusinessProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProcess

`func (o *CreateCDARRequest) SetBusinessProcess(v string)`

SetBusinessProcess sets BusinessProcess field to given value.

### HasBusinessProcess

`func (o *CreateCDARRequest) HasBusinessProcess() bool`

HasBusinessProcess returns a boolean if a field has been set.

### GetTypeCode

`func (o *CreateCDARRequest) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *CreateCDARRequest) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *CreateCDARRequest) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.

### HasTypeCode

`func (o *CreateCDARRequest) HasTypeCode() bool`

HasTypeCode returns a boolean if a field has been set.

### GetSenderSiren

`func (o *CreateCDARRequest) GetSenderSiren() string`

GetSenderSiren returns the SenderSiren field if non-nil, zero value otherwise.

### GetSenderSirenOk

`func (o *CreateCDARRequest) GetSenderSirenOk() (*string, bool)`

GetSenderSirenOk returns a tuple with the SenderSiren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderSiren

`func (o *CreateCDARRequest) SetSenderSiren(v string)`

SetSenderSiren sets SenderSiren field to given value.


### GetSenderRole

`func (o *CreateCDARRequest) GetSenderRole() string`

GetSenderRole returns the SenderRole field if non-nil, zero value otherwise.

### GetSenderRoleOk

`func (o *CreateCDARRequest) GetSenderRoleOk() (*string, bool)`

GetSenderRoleOk returns a tuple with the SenderRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderRole

`func (o *CreateCDARRequest) SetSenderRole(v string)`

SetSenderRole sets SenderRole field to given value.

### HasSenderRole

`func (o *CreateCDARRequest) HasSenderRole() bool`

HasSenderRole returns a boolean if a field has been set.

### GetSenderName

`func (o *CreateCDARRequest) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *CreateCDARRequest) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *CreateCDARRequest) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.

### HasSenderName

`func (o *CreateCDARRequest) HasSenderName() bool`

HasSenderName returns a boolean if a field has been set.

### SetSenderNameNil

`func (o *CreateCDARRequest) SetSenderNameNil(b bool)`

 SetSenderNameNil sets the value for SenderName to be an explicit nil

### UnsetSenderName
`func (o *CreateCDARRequest) UnsetSenderName()`

UnsetSenderName ensures that no value is present for SenderName, not even an explicit nil
### GetSenderEmail

`func (o *CreateCDARRequest) GetSenderEmail() string`

GetSenderEmail returns the SenderEmail field if non-nil, zero value otherwise.

### GetSenderEmailOk

`func (o *CreateCDARRequest) GetSenderEmailOk() (*string, bool)`

GetSenderEmailOk returns a tuple with the SenderEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderEmail

`func (o *CreateCDARRequest) SetSenderEmail(v string)`

SetSenderEmail sets SenderEmail field to given value.

### HasSenderEmail

`func (o *CreateCDARRequest) HasSenderEmail() bool`

HasSenderEmail returns a boolean if a field has been set.

### SetSenderEmailNil

`func (o *CreateCDARRequest) SetSenderEmailNil(b bool)`

 SetSenderEmailNil sets the value for SenderEmail to be an explicit nil

### UnsetSenderEmail
`func (o *CreateCDARRequest) UnsetSenderEmail()`

UnsetSenderEmail ensures that no value is present for SenderEmail, not even an explicit nil
### GetRecipients

`func (o *CreateCDARRequest) GetRecipients() []RecipientInput`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *CreateCDARRequest) GetRecipientsOk() (*[]RecipientInput, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *CreateCDARRequest) SetRecipients(v []RecipientInput)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *CreateCDARRequest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetInvoiceId

`func (o *CreateCDARRequest) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CreateCDARRequest) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CreateCDARRequest) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetInvoiceIssueDate

`func (o *CreateCDARRequest) GetInvoiceIssueDate() string`

GetInvoiceIssueDate returns the InvoiceIssueDate field if non-nil, zero value otherwise.

### GetInvoiceIssueDateOk

`func (o *CreateCDARRequest) GetInvoiceIssueDateOk() (*string, bool)`

GetInvoiceIssueDateOk returns a tuple with the InvoiceIssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceIssueDate

`func (o *CreateCDARRequest) SetInvoiceIssueDate(v string)`

SetInvoiceIssueDate sets InvoiceIssueDate field to given value.


### GetInvoiceTypeCode

`func (o *CreateCDARRequest) GetInvoiceTypeCode() string`

GetInvoiceTypeCode returns the InvoiceTypeCode field if non-nil, zero value otherwise.

### GetInvoiceTypeCodeOk

`func (o *CreateCDARRequest) GetInvoiceTypeCodeOk() (*string, bool)`

GetInvoiceTypeCodeOk returns a tuple with the InvoiceTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTypeCode

`func (o *CreateCDARRequest) SetInvoiceTypeCode(v string)`

SetInvoiceTypeCode sets InvoiceTypeCode field to given value.

### HasInvoiceTypeCode

`func (o *CreateCDARRequest) HasInvoiceTypeCode() bool`

HasInvoiceTypeCode returns a boolean if a field has been set.

### GetInvoiceSellerSiren

`func (o *CreateCDARRequest) GetInvoiceSellerSiren() string`

GetInvoiceSellerSiren returns the InvoiceSellerSiren field if non-nil, zero value otherwise.

### GetInvoiceSellerSirenOk

`func (o *CreateCDARRequest) GetInvoiceSellerSirenOk() (*string, bool)`

GetInvoiceSellerSirenOk returns a tuple with the InvoiceSellerSiren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceSellerSiren

`func (o *CreateCDARRequest) SetInvoiceSellerSiren(v string)`

SetInvoiceSellerSiren sets InvoiceSellerSiren field to given value.

### HasInvoiceSellerSiren

`func (o *CreateCDARRequest) HasInvoiceSellerSiren() bool`

HasInvoiceSellerSiren returns a boolean if a field has been set.

### SetInvoiceSellerSirenNil

`func (o *CreateCDARRequest) SetInvoiceSellerSirenNil(b bool)`

 SetInvoiceSellerSirenNil sets the value for InvoiceSellerSiren to be an explicit nil

### UnsetInvoiceSellerSiren
`func (o *CreateCDARRequest) UnsetInvoiceSellerSiren()`

UnsetInvoiceSellerSiren ensures that no value is present for InvoiceSellerSiren, not even an explicit nil
### GetInvoiceBuyerSiren

`func (o *CreateCDARRequest) GetInvoiceBuyerSiren() string`

GetInvoiceBuyerSiren returns the InvoiceBuyerSiren field if non-nil, zero value otherwise.

### GetInvoiceBuyerSirenOk

`func (o *CreateCDARRequest) GetInvoiceBuyerSirenOk() (*string, bool)`

GetInvoiceBuyerSirenOk returns a tuple with the InvoiceBuyerSiren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceBuyerSiren

`func (o *CreateCDARRequest) SetInvoiceBuyerSiren(v string)`

SetInvoiceBuyerSiren sets InvoiceBuyerSiren field to given value.

### HasInvoiceBuyerSiren

`func (o *CreateCDARRequest) HasInvoiceBuyerSiren() bool`

HasInvoiceBuyerSiren returns a boolean if a field has been set.

### SetInvoiceBuyerSirenNil

`func (o *CreateCDARRequest) SetInvoiceBuyerSirenNil(b bool)`

 SetInvoiceBuyerSirenNil sets the value for InvoiceBuyerSiren to be an explicit nil

### UnsetInvoiceBuyerSiren
`func (o *CreateCDARRequest) UnsetInvoiceBuyerSiren()`

UnsetInvoiceBuyerSiren ensures that no value is present for InvoiceBuyerSiren, not even an explicit nil
### GetStatus

`func (o *CreateCDARRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateCDARRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateCDARRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReasonCode

`func (o *CreateCDARRequest) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *CreateCDARRequest) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *CreateCDARRequest) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *CreateCDARRequest) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### SetReasonCodeNil

`func (o *CreateCDARRequest) SetReasonCodeNil(b bool)`

 SetReasonCodeNil sets the value for ReasonCode to be an explicit nil

### UnsetReasonCode
`func (o *CreateCDARRequest) UnsetReasonCode()`

UnsetReasonCode ensures that no value is present for ReasonCode, not even an explicit nil
### GetReasonText

`func (o *CreateCDARRequest) GetReasonText() string`

GetReasonText returns the ReasonText field if non-nil, zero value otherwise.

### GetReasonTextOk

`func (o *CreateCDARRequest) GetReasonTextOk() (*string, bool)`

GetReasonTextOk returns a tuple with the ReasonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonText

`func (o *CreateCDARRequest) SetReasonText(v string)`

SetReasonText sets ReasonText field to given value.

### HasReasonText

`func (o *CreateCDARRequest) HasReasonText() bool`

HasReasonText returns a boolean if a field has been set.

### SetReasonTextNil

`func (o *CreateCDARRequest) SetReasonTextNil(b bool)`

 SetReasonTextNil sets the value for ReasonText to be an explicit nil

### UnsetReasonText
`func (o *CreateCDARRequest) UnsetReasonText()`

UnsetReasonText ensures that no value is present for ReasonText, not even an explicit nil
### GetActionCode

`func (o *CreateCDARRequest) GetActionCode() string`

GetActionCode returns the ActionCode field if non-nil, zero value otherwise.

### GetActionCodeOk

`func (o *CreateCDARRequest) GetActionCodeOk() (*string, bool)`

GetActionCodeOk returns a tuple with the ActionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionCode

`func (o *CreateCDARRequest) SetActionCode(v string)`

SetActionCode sets ActionCode field to given value.

### HasActionCode

`func (o *CreateCDARRequest) HasActionCode() bool`

HasActionCode returns a boolean if a field has been set.

### SetActionCodeNil

`func (o *CreateCDARRequest) SetActionCodeNil(b bool)`

 SetActionCodeNil sets the value for ActionCode to be an explicit nil

### UnsetActionCode
`func (o *CreateCDARRequest) UnsetActionCode()`

UnsetActionCode ensures that no value is present for ActionCode, not even an explicit nil
### GetEncaisseAmount

`func (o *CreateCDARRequest) GetEncaisseAmount() Encaisseamount`

GetEncaisseAmount returns the EncaisseAmount field if non-nil, zero value otherwise.

### GetEncaisseAmountOk

`func (o *CreateCDARRequest) GetEncaisseAmountOk() (*Encaisseamount, bool)`

GetEncaisseAmountOk returns a tuple with the EncaisseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncaisseAmount

`func (o *CreateCDARRequest) SetEncaisseAmount(v Encaisseamount)`

SetEncaisseAmount sets EncaisseAmount field to given value.

### HasEncaisseAmount

`func (o *CreateCDARRequest) HasEncaisseAmount() bool`

HasEncaisseAmount returns a boolean if a field has been set.

### SetEncaisseAmountNil

`func (o *CreateCDARRequest) SetEncaisseAmountNil(b bool)`

 SetEncaisseAmountNil sets the value for EncaisseAmount to be an explicit nil

### UnsetEncaisseAmount
`func (o *CreateCDARRequest) UnsetEncaisseAmount()`

UnsetEncaisseAmount ensures that no value is present for EncaisseAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SubmitCDARResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowId** | **string** | Identifiant du flux AFNOR | 
**Status** | **string** | Statut de la soumission | 
**Message** | Pointer to **NullableString** |  | [optional] 
**DocumentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSubmitCDARResponse

`func NewSubmitCDARResponse(flowId string, status string, ) *SubmitCDARResponse`

NewSubmitCDARResponse instantiates a new SubmitCDARResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitCDARResponseWithDefaults

`func NewSubmitCDARResponseWithDefaults() *SubmitCDARResponse`

NewSubmitCDARResponseWithDefaults instantiates a new SubmitCDARResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowId

`func (o *SubmitCDARResponse) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *SubmitCDARResponse) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *SubmitCDARResponse) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.


### GetStatus

`func (o *SubmitCDARResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubmitCDARResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubmitCDARResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *SubmitCDARResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SubmitCDARResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SubmitCDARResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SubmitCDARResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *SubmitCDARResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *SubmitCDARResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetDocumentId

`func (o *SubmitCDARResponse) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *SubmitCDARResponse) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *SubmitCDARResponse) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *SubmitCDARResponse) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### SetDocumentIdNil

`func (o *SubmitCDARResponse) SetDocumentIdNil(b bool)`

 SetDocumentIdNil sets the value for DocumentId to be an explicit nil

### UnsetDocumentId
`func (o *SubmitCDARResponse) UnsetDocumentId()`

UnsetDocumentId ensures that no value is present for DocumentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



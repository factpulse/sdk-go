# SubmitInvoiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnCode** | **int32** | Return code (0 &#x3D; success) | 
**Message** | **string** | Return message | 
**ChorusInvoiceId** | Pointer to **NullableInt32** |  | [optional] 
**DepositFlowNumber** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSubmitInvoiceResponse

`func NewSubmitInvoiceResponse(returnCode int32, message string, ) *SubmitInvoiceResponse`

NewSubmitInvoiceResponse instantiates a new SubmitInvoiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitInvoiceResponseWithDefaults

`func NewSubmitInvoiceResponseWithDefaults() *SubmitInvoiceResponse`

NewSubmitInvoiceResponseWithDefaults instantiates a new SubmitInvoiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnCode

`func (o *SubmitInvoiceResponse) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *SubmitInvoiceResponse) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *SubmitInvoiceResponse) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.


### GetMessage

`func (o *SubmitInvoiceResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SubmitInvoiceResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SubmitInvoiceResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetChorusInvoiceId

`func (o *SubmitInvoiceResponse) GetChorusInvoiceId() int32`

GetChorusInvoiceId returns the ChorusInvoiceId field if non-nil, zero value otherwise.

### GetChorusInvoiceIdOk

`func (o *SubmitInvoiceResponse) GetChorusInvoiceIdOk() (*int32, bool)`

GetChorusInvoiceIdOk returns a tuple with the ChorusInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChorusInvoiceId

`func (o *SubmitInvoiceResponse) SetChorusInvoiceId(v int32)`

SetChorusInvoiceId sets ChorusInvoiceId field to given value.

### HasChorusInvoiceId

`func (o *SubmitInvoiceResponse) HasChorusInvoiceId() bool`

HasChorusInvoiceId returns a boolean if a field has been set.

### SetChorusInvoiceIdNil

`func (o *SubmitInvoiceResponse) SetChorusInvoiceIdNil(b bool)`

 SetChorusInvoiceIdNil sets the value for ChorusInvoiceId to be an explicit nil

### UnsetChorusInvoiceId
`func (o *SubmitInvoiceResponse) UnsetChorusInvoiceId()`

UnsetChorusInvoiceId ensures that no value is present for ChorusInvoiceId, not even an explicit nil
### GetDepositFlowNumber

`func (o *SubmitInvoiceResponse) GetDepositFlowNumber() string`

GetDepositFlowNumber returns the DepositFlowNumber field if non-nil, zero value otherwise.

### GetDepositFlowNumberOk

`func (o *SubmitInvoiceResponse) GetDepositFlowNumberOk() (*string, bool)`

GetDepositFlowNumberOk returns a tuple with the DepositFlowNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositFlowNumber

`func (o *SubmitInvoiceResponse) SetDepositFlowNumber(v string)`

SetDepositFlowNumber sets DepositFlowNumber field to given value.

### HasDepositFlowNumber

`func (o *SubmitInvoiceResponse) HasDepositFlowNumber() bool`

HasDepositFlowNumber returns a boolean if a field has been set.

### SetDepositFlowNumberNil

`func (o *SubmitInvoiceResponse) SetDepositFlowNumberNil(b bool)`

 SetDepositFlowNumberNil sets the value for DepositFlowNumber to be an explicit nil

### UnsetDepositFlowNumber
`func (o *SubmitInvoiceResponse) UnsetDepositFlowNumber()`

UnsetDepositFlowNumber ensures that no value is present for DepositFlowNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



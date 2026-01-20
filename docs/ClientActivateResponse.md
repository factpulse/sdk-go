# ClientActivateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | Unique client identifier | 
**IsActive** | **bool** | New status | 
**Message** | **string** | Confirmation message | 

## Methods

### NewClientActivateResponse

`func NewClientActivateResponse(uid string, isActive bool, message string, ) *ClientActivateResponse`

NewClientActivateResponse instantiates a new ClientActivateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientActivateResponseWithDefaults

`func NewClientActivateResponseWithDefaults() *ClientActivateResponse`

NewClientActivateResponseWithDefaults instantiates a new ClientActivateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *ClientActivateResponse) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ClientActivateResponse) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ClientActivateResponse) SetUid(v string)`

SetUid sets Uid field to given value.


### GetIsActive

`func (o *ClientActivateResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ClientActivateResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ClientActivateResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetMessage

`func (o *ClientActivateResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClientActivateResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClientActivateResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SuperPDPDisconnectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Whether the disconnection completed successfully | 
**Deleted** | **bool** | Whether a SuperPDP connection existed and was removed | 

## Methods

### NewSuperPDPDisconnectResponse

`func NewSuperPDPDisconnectResponse(success bool, deleted bool, ) *SuperPDPDisconnectResponse`

NewSuperPDPDisconnectResponse instantiates a new SuperPDPDisconnectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuperPDPDisconnectResponseWithDefaults

`func NewSuperPDPDisconnectResponseWithDefaults() *SuperPDPDisconnectResponse`

NewSuperPDPDisconnectResponseWithDefaults instantiates a new SuperPDPDisconnectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *SuperPDPDisconnectResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SuperPDPDisconnectResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SuperPDPDisconnectResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetDeleted

`func (o *SuperPDPDisconnectResponse) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *SuperPDPDisconnectResponse) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *SuperPDPDisconnectResponse) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



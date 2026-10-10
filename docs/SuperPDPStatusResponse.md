# SuperPDPStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connected** | **bool** | Whether this client has connected its SuperPDP account | 
**ConnectedAt** | Pointer to **NullableTime** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**ModeSandbox** | Pointer to **NullableBool** |  | [optional] 
**LastError** | Pointer to **string** | Last token refresh error, if any | [optional] [default to ""]

## Methods

### NewSuperPDPStatusResponse

`func NewSuperPDPStatusResponse(connected bool, ) *SuperPDPStatusResponse`

NewSuperPDPStatusResponse instantiates a new SuperPDPStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuperPDPStatusResponseWithDefaults

`func NewSuperPDPStatusResponseWithDefaults() *SuperPDPStatusResponse`

NewSuperPDPStatusResponseWithDefaults instantiates a new SuperPDPStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnected

`func (o *SuperPDPStatusResponse) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *SuperPDPStatusResponse) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *SuperPDPStatusResponse) SetConnected(v bool)`

SetConnected sets Connected field to given value.


### GetConnectedAt

`func (o *SuperPDPStatusResponse) GetConnectedAt() time.Time`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *SuperPDPStatusResponse) GetConnectedAtOk() (*time.Time, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *SuperPDPStatusResponse) SetConnectedAt(v time.Time)`

SetConnectedAt sets ConnectedAt field to given value.

### HasConnectedAt

`func (o *SuperPDPStatusResponse) HasConnectedAt() bool`

HasConnectedAt returns a boolean if a field has been set.

### SetConnectedAtNil

`func (o *SuperPDPStatusResponse) SetConnectedAtNil(b bool)`

 SetConnectedAtNil sets the value for ConnectedAt to be an explicit nil

### UnsetConnectedAt
`func (o *SuperPDPStatusResponse) UnsetConnectedAt()`

UnsetConnectedAt ensures that no value is present for ConnectedAt, not even an explicit nil
### GetExpiresAt

`func (o *SuperPDPStatusResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SuperPDPStatusResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SuperPDPStatusResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *SuperPDPStatusResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *SuperPDPStatusResponse) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *SuperPDPStatusResponse) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetModeSandbox

`func (o *SuperPDPStatusResponse) GetModeSandbox() bool`

GetModeSandbox returns the ModeSandbox field if non-nil, zero value otherwise.

### GetModeSandboxOk

`func (o *SuperPDPStatusResponse) GetModeSandboxOk() (*bool, bool)`

GetModeSandboxOk returns a tuple with the ModeSandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeSandbox

`func (o *SuperPDPStatusResponse) SetModeSandbox(v bool)`

SetModeSandbox sets ModeSandbox field to given value.

### HasModeSandbox

`func (o *SuperPDPStatusResponse) HasModeSandbox() bool`

HasModeSandbox returns a boolean if a field has been set.

### SetModeSandboxNil

`func (o *SuperPDPStatusResponse) SetModeSandboxNil(b bool)`

 SetModeSandboxNil sets the value for ModeSandbox to be an explicit nil

### UnsetModeSandbox
`func (o *SuperPDPStatusResponse) UnsetModeSandbox()`

UnsetModeSandbox ensures that no value is present for ModeSandbox, not even an explicit nil
### GetLastError

`func (o *SuperPDPStatusResponse) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *SuperPDPStatusResponse) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *SuperPDPStatusResponse) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *SuperPDPStatusResponse) HasLastError() bool`

HasLastError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



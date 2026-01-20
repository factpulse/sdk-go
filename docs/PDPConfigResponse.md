# PDPConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsConfigured** | **bool** | Whether PDP config exists | 
**Id** | Pointer to **NullableInt32** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**ModeSandbox** | Pointer to **NullableBool** |  | [optional] 
**FlowServiceUrl** | Pointer to **NullableString** |  | [optional] 
**TokenUrl** | Pointer to **NullableString** |  | [optional] 
**OauthClientId** | Pointer to **NullableString** |  | [optional] 
**SecretStatus** | Pointer to [**NullableSecretStatus**](SecretStatus.md) |  | [optional] 
**LastTestAt** | Pointer to **NullableTime** |  | [optional] 
**LastTestSuccess** | Pointer to **NullableBool** |  | [optional] 
**LastTestError** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPDPConfigResponse

`func NewPDPConfigResponse(isConfigured bool, ) *PDPConfigResponse`

NewPDPConfigResponse instantiates a new PDPConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPDPConfigResponseWithDefaults

`func NewPDPConfigResponseWithDefaults() *PDPConfigResponse`

NewPDPConfigResponseWithDefaults instantiates a new PDPConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsConfigured

`func (o *PDPConfigResponse) GetIsConfigured() bool`

GetIsConfigured returns the IsConfigured field if non-nil, zero value otherwise.

### GetIsConfiguredOk

`func (o *PDPConfigResponse) GetIsConfiguredOk() (*bool, bool)`

GetIsConfiguredOk returns a tuple with the IsConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfigured

`func (o *PDPConfigResponse) SetIsConfigured(v bool)`

SetIsConfigured sets IsConfigured field to given value.


### GetId

`func (o *PDPConfigResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PDPConfigResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PDPConfigResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PDPConfigResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PDPConfigResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PDPConfigResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsActive

`func (o *PDPConfigResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PDPConfigResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PDPConfigResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PDPConfigResponse) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *PDPConfigResponse) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *PDPConfigResponse) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetModeSandbox

`func (o *PDPConfigResponse) GetModeSandbox() bool`

GetModeSandbox returns the ModeSandbox field if non-nil, zero value otherwise.

### GetModeSandboxOk

`func (o *PDPConfigResponse) GetModeSandboxOk() (*bool, bool)`

GetModeSandboxOk returns a tuple with the ModeSandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeSandbox

`func (o *PDPConfigResponse) SetModeSandbox(v bool)`

SetModeSandbox sets ModeSandbox field to given value.

### HasModeSandbox

`func (o *PDPConfigResponse) HasModeSandbox() bool`

HasModeSandbox returns a boolean if a field has been set.

### SetModeSandboxNil

`func (o *PDPConfigResponse) SetModeSandboxNil(b bool)`

 SetModeSandboxNil sets the value for ModeSandbox to be an explicit nil

### UnsetModeSandbox
`func (o *PDPConfigResponse) UnsetModeSandbox()`

UnsetModeSandbox ensures that no value is present for ModeSandbox, not even an explicit nil
### GetFlowServiceUrl

`func (o *PDPConfigResponse) GetFlowServiceUrl() string`

GetFlowServiceUrl returns the FlowServiceUrl field if non-nil, zero value otherwise.

### GetFlowServiceUrlOk

`func (o *PDPConfigResponse) GetFlowServiceUrlOk() (*string, bool)`

GetFlowServiceUrlOk returns a tuple with the FlowServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowServiceUrl

`func (o *PDPConfigResponse) SetFlowServiceUrl(v string)`

SetFlowServiceUrl sets FlowServiceUrl field to given value.

### HasFlowServiceUrl

`func (o *PDPConfigResponse) HasFlowServiceUrl() bool`

HasFlowServiceUrl returns a boolean if a field has been set.

### SetFlowServiceUrlNil

`func (o *PDPConfigResponse) SetFlowServiceUrlNil(b bool)`

 SetFlowServiceUrlNil sets the value for FlowServiceUrl to be an explicit nil

### UnsetFlowServiceUrl
`func (o *PDPConfigResponse) UnsetFlowServiceUrl()`

UnsetFlowServiceUrl ensures that no value is present for FlowServiceUrl, not even an explicit nil
### GetTokenUrl

`func (o *PDPConfigResponse) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *PDPConfigResponse) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *PDPConfigResponse) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *PDPConfigResponse) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.

### SetTokenUrlNil

`func (o *PDPConfigResponse) SetTokenUrlNil(b bool)`

 SetTokenUrlNil sets the value for TokenUrl to be an explicit nil

### UnsetTokenUrl
`func (o *PDPConfigResponse) UnsetTokenUrl()`

UnsetTokenUrl ensures that no value is present for TokenUrl, not even an explicit nil
### GetOauthClientId

`func (o *PDPConfigResponse) GetOauthClientId() string`

GetOauthClientId returns the OauthClientId field if non-nil, zero value otherwise.

### GetOauthClientIdOk

`func (o *PDPConfigResponse) GetOauthClientIdOk() (*string, bool)`

GetOauthClientIdOk returns a tuple with the OauthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientId

`func (o *PDPConfigResponse) SetOauthClientId(v string)`

SetOauthClientId sets OauthClientId field to given value.

### HasOauthClientId

`func (o *PDPConfigResponse) HasOauthClientId() bool`

HasOauthClientId returns a boolean if a field has been set.

### SetOauthClientIdNil

`func (o *PDPConfigResponse) SetOauthClientIdNil(b bool)`

 SetOauthClientIdNil sets the value for OauthClientId to be an explicit nil

### UnsetOauthClientId
`func (o *PDPConfigResponse) UnsetOauthClientId()`

UnsetOauthClientId ensures that no value is present for OauthClientId, not even an explicit nil
### GetSecretStatus

`func (o *PDPConfigResponse) GetSecretStatus() SecretStatus`

GetSecretStatus returns the SecretStatus field if non-nil, zero value otherwise.

### GetSecretStatusOk

`func (o *PDPConfigResponse) GetSecretStatusOk() (*SecretStatus, bool)`

GetSecretStatusOk returns a tuple with the SecretStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretStatus

`func (o *PDPConfigResponse) SetSecretStatus(v SecretStatus)`

SetSecretStatus sets SecretStatus field to given value.

### HasSecretStatus

`func (o *PDPConfigResponse) HasSecretStatus() bool`

HasSecretStatus returns a boolean if a field has been set.

### SetSecretStatusNil

`func (o *PDPConfigResponse) SetSecretStatusNil(b bool)`

 SetSecretStatusNil sets the value for SecretStatus to be an explicit nil

### UnsetSecretStatus
`func (o *PDPConfigResponse) UnsetSecretStatus()`

UnsetSecretStatus ensures that no value is present for SecretStatus, not even an explicit nil
### GetLastTestAt

`func (o *PDPConfigResponse) GetLastTestAt() time.Time`

GetLastTestAt returns the LastTestAt field if non-nil, zero value otherwise.

### GetLastTestAtOk

`func (o *PDPConfigResponse) GetLastTestAtOk() (*time.Time, bool)`

GetLastTestAtOk returns a tuple with the LastTestAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestAt

`func (o *PDPConfigResponse) SetLastTestAt(v time.Time)`

SetLastTestAt sets LastTestAt field to given value.

### HasLastTestAt

`func (o *PDPConfigResponse) HasLastTestAt() bool`

HasLastTestAt returns a boolean if a field has been set.

### SetLastTestAtNil

`func (o *PDPConfigResponse) SetLastTestAtNil(b bool)`

 SetLastTestAtNil sets the value for LastTestAt to be an explicit nil

### UnsetLastTestAt
`func (o *PDPConfigResponse) UnsetLastTestAt()`

UnsetLastTestAt ensures that no value is present for LastTestAt, not even an explicit nil
### GetLastTestSuccess

`func (o *PDPConfigResponse) GetLastTestSuccess() bool`

GetLastTestSuccess returns the LastTestSuccess field if non-nil, zero value otherwise.

### GetLastTestSuccessOk

`func (o *PDPConfigResponse) GetLastTestSuccessOk() (*bool, bool)`

GetLastTestSuccessOk returns a tuple with the LastTestSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestSuccess

`func (o *PDPConfigResponse) SetLastTestSuccess(v bool)`

SetLastTestSuccess sets LastTestSuccess field to given value.

### HasLastTestSuccess

`func (o *PDPConfigResponse) HasLastTestSuccess() bool`

HasLastTestSuccess returns a boolean if a field has been set.

### SetLastTestSuccessNil

`func (o *PDPConfigResponse) SetLastTestSuccessNil(b bool)`

 SetLastTestSuccessNil sets the value for LastTestSuccess to be an explicit nil

### UnsetLastTestSuccess
`func (o *PDPConfigResponse) UnsetLastTestSuccess()`

UnsetLastTestSuccess ensures that no value is present for LastTestSuccess, not even an explicit nil
### GetLastTestError

`func (o *PDPConfigResponse) GetLastTestError() string`

GetLastTestError returns the LastTestError field if non-nil, zero value otherwise.

### GetLastTestErrorOk

`func (o *PDPConfigResponse) GetLastTestErrorOk() (*string, bool)`

GetLastTestErrorOk returns a tuple with the LastTestError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestError

`func (o *PDPConfigResponse) SetLastTestError(v string)`

SetLastTestError sets LastTestError field to given value.

### HasLastTestError

`func (o *PDPConfigResponse) HasLastTestError() bool`

HasLastTestError returns a boolean if a field has been set.

### SetLastTestErrorNil

`func (o *PDPConfigResponse) SetLastTestErrorNil(b bool)`

 SetLastTestErrorNil sets the value for LastTestError to be an explicit nil

### UnsetLastTestError
`func (o *PDPConfigResponse) UnsetLastTestError()`

UnsetLastTestError ensures that no value is present for LastTestError, not even an explicit nil
### GetCreatedAt

`func (o *PDPConfigResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PDPConfigResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PDPConfigResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PDPConfigResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *PDPConfigResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *PDPConfigResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *PDPConfigResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PDPConfigResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PDPConfigResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PDPConfigResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *PDPConfigResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *PDPConfigResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetMessage

`func (o *PDPConfigResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PDPConfigResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PDPConfigResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PDPConfigResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *PDPConfigResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *PDPConfigResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



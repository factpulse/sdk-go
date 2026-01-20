# PDPConfigUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | Pointer to **bool** | Whether config is active | [optional] [default to true]
**ModeSandbox** | Pointer to **bool** | Sandbox mode | [optional] [default to false]
**FlowServiceUrl** | **string** | PDP Flow Service URL | 
**TokenUrl** | **string** | PDP OAuth token URL | 
**OauthClientId** | **string** | OAuth Client ID | 
**ClientSecret** | **string** | OAuth Client Secret (sent but never returned) | 

## Methods

### NewPDPConfigUpdateRequest

`func NewPDPConfigUpdateRequest(flowServiceUrl string, tokenUrl string, oauthClientId string, clientSecret string, ) *PDPConfigUpdateRequest`

NewPDPConfigUpdateRequest instantiates a new PDPConfigUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPDPConfigUpdateRequestWithDefaults

`func NewPDPConfigUpdateRequestWithDefaults() *PDPConfigUpdateRequest`

NewPDPConfigUpdateRequestWithDefaults instantiates a new PDPConfigUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActive

`func (o *PDPConfigUpdateRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PDPConfigUpdateRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PDPConfigUpdateRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PDPConfigUpdateRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetModeSandbox

`func (o *PDPConfigUpdateRequest) GetModeSandbox() bool`

GetModeSandbox returns the ModeSandbox field if non-nil, zero value otherwise.

### GetModeSandboxOk

`func (o *PDPConfigUpdateRequest) GetModeSandboxOk() (*bool, bool)`

GetModeSandboxOk returns a tuple with the ModeSandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeSandbox

`func (o *PDPConfigUpdateRequest) SetModeSandbox(v bool)`

SetModeSandbox sets ModeSandbox field to given value.

### HasModeSandbox

`func (o *PDPConfigUpdateRequest) HasModeSandbox() bool`

HasModeSandbox returns a boolean if a field has been set.

### GetFlowServiceUrl

`func (o *PDPConfigUpdateRequest) GetFlowServiceUrl() string`

GetFlowServiceUrl returns the FlowServiceUrl field if non-nil, zero value otherwise.

### GetFlowServiceUrlOk

`func (o *PDPConfigUpdateRequest) GetFlowServiceUrlOk() (*string, bool)`

GetFlowServiceUrlOk returns a tuple with the FlowServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowServiceUrl

`func (o *PDPConfigUpdateRequest) SetFlowServiceUrl(v string)`

SetFlowServiceUrl sets FlowServiceUrl field to given value.


### GetTokenUrl

`func (o *PDPConfigUpdateRequest) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *PDPConfigUpdateRequest) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *PDPConfigUpdateRequest) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.


### GetOauthClientId

`func (o *PDPConfigUpdateRequest) GetOauthClientId() string`

GetOauthClientId returns the OauthClientId field if non-nil, zero value otherwise.

### GetOauthClientIdOk

`func (o *PDPConfigUpdateRequest) GetOauthClientIdOk() (*string, bool)`

GetOauthClientIdOk returns a tuple with the OauthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientId

`func (o *PDPConfigUpdateRequest) SetOauthClientId(v string)`

SetOauthClientId sets OauthClientId field to given value.


### GetClientSecret

`func (o *PDPConfigUpdateRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PDPConfigUpdateRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PDPConfigUpdateRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



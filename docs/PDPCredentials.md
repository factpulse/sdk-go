# PDPCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowServiceUrl** | **string** | URL de base du Flow Service AFNOR | 
**DirectoryServiceUrl** | Pointer to **NullableString** |  | [optional] 
**TokenUrl** | **string** | URL du serveur OAuth2 | 
**ClientId** | **string** | Client ID OAuth2 | 
**ClientSecret** | **string** | Client Secret OAuth2 (sensible) | 

## Methods

### NewPDPCredentials

`func NewPDPCredentials(flowServiceUrl string, tokenUrl string, clientId string, clientSecret string, ) *PDPCredentials`

NewPDPCredentials instantiates a new PDPCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPDPCredentialsWithDefaults

`func NewPDPCredentialsWithDefaults() *PDPCredentials`

NewPDPCredentialsWithDefaults instantiates a new PDPCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowServiceUrl

`func (o *PDPCredentials) GetFlowServiceUrl() string`

GetFlowServiceUrl returns the FlowServiceUrl field if non-nil, zero value otherwise.

### GetFlowServiceUrlOk

`func (o *PDPCredentials) GetFlowServiceUrlOk() (*string, bool)`

GetFlowServiceUrlOk returns a tuple with the FlowServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowServiceUrl

`func (o *PDPCredentials) SetFlowServiceUrl(v string)`

SetFlowServiceUrl sets FlowServiceUrl field to given value.


### GetDirectoryServiceUrl

`func (o *PDPCredentials) GetDirectoryServiceUrl() string`

GetDirectoryServiceUrl returns the DirectoryServiceUrl field if non-nil, zero value otherwise.

### GetDirectoryServiceUrlOk

`func (o *PDPCredentials) GetDirectoryServiceUrlOk() (*string, bool)`

GetDirectoryServiceUrlOk returns a tuple with the DirectoryServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryServiceUrl

`func (o *PDPCredentials) SetDirectoryServiceUrl(v string)`

SetDirectoryServiceUrl sets DirectoryServiceUrl field to given value.

### HasDirectoryServiceUrl

`func (o *PDPCredentials) HasDirectoryServiceUrl() bool`

HasDirectoryServiceUrl returns a boolean if a field has been set.

### SetDirectoryServiceUrlNil

`func (o *PDPCredentials) SetDirectoryServiceUrlNil(b bool)`

 SetDirectoryServiceUrlNil sets the value for DirectoryServiceUrl to be an explicit nil

### UnsetDirectoryServiceUrl
`func (o *PDPCredentials) UnsetDirectoryServiceUrl()`

UnsetDirectoryServiceUrl ensures that no value is present for DirectoryServiceUrl, not even an explicit nil
### GetTokenUrl

`func (o *PDPCredentials) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *PDPCredentials) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *PDPCredentials) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.


### GetClientId

`func (o *PDPCredentials) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PDPCredentials) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PDPCredentials) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *PDPCredentials) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PDPCredentials) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PDPCredentials) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



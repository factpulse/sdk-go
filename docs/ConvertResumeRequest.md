# ConvertResumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overrides** | Pointer to **map[string]interface{}** | Sous-ensemble de FacturXInvoice a mettre a jour (merge profond) | [optional] 
**CallbackUrl** | Pointer to **NullableString** |  | [optional] 
**WebhookMode** | Pointer to **string** | Mode de livraison webhook: &#39;inline&#39; ou &#39;download_url&#39; | [optional] [default to "inline"]

## Methods

### NewConvertResumeRequest

`func NewConvertResumeRequest() *ConvertResumeRequest`

NewConvertResumeRequest instantiates a new ConvertResumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertResumeRequestWithDefaults

`func NewConvertResumeRequestWithDefaults() *ConvertResumeRequest`

NewConvertResumeRequestWithDefaults instantiates a new ConvertResumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverrides

`func (o *ConvertResumeRequest) GetOverrides() map[string]interface{}`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *ConvertResumeRequest) GetOverridesOk() (*map[string]interface{}, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *ConvertResumeRequest) SetOverrides(v map[string]interface{})`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *ConvertResumeRequest) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *ConvertResumeRequest) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *ConvertResumeRequest) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *ConvertResumeRequest) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *ConvertResumeRequest) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### SetCallbackUrlNil

`func (o *ConvertResumeRequest) SetCallbackUrlNil(b bool)`

 SetCallbackUrlNil sets the value for CallbackUrl to be an explicit nil

### UnsetCallbackUrl
`func (o *ConvertResumeRequest) UnsetCallbackUrl()`

UnsetCallbackUrl ensures that no value is present for CallbackUrl, not even an explicit nil
### GetWebhookMode

`func (o *ConvertResumeRequest) GetWebhookMode() string`

GetWebhookMode returns the WebhookMode field if non-nil, zero value otherwise.

### GetWebhookModeOk

`func (o *ConvertResumeRequest) GetWebhookModeOk() (*string, bool)`

GetWebhookModeOk returns a tuple with the WebhookMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookMode

`func (o *ConvertResumeRequest) SetWebhookMode(v string)`

SetWebhookMode sets WebhookMode field to given value.

### HasWebhookMode

`func (o *ConvertResumeRequest) HasWebhookMode() bool`

HasWebhookMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



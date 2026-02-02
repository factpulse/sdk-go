# WebhookSecretGenerateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Whether the secret was generated successfully | 
**WebhookSecret** | **string** | The generated webhook secret (save it, it will never be shown again) | 
**Message** | **string** | Result message | 
**CreatedAt** | **time.Time** | When the secret was created | 

## Methods

### NewWebhookSecretGenerateResponse

`func NewWebhookSecretGenerateResponse(success bool, webhookSecret string, message string, createdAt time.Time, ) *WebhookSecretGenerateResponse`

NewWebhookSecretGenerateResponse instantiates a new WebhookSecretGenerateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSecretGenerateResponseWithDefaults

`func NewWebhookSecretGenerateResponseWithDefaults() *WebhookSecretGenerateResponse`

NewWebhookSecretGenerateResponseWithDefaults instantiates a new WebhookSecretGenerateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *WebhookSecretGenerateResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *WebhookSecretGenerateResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *WebhookSecretGenerateResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetWebhookSecret

`func (o *WebhookSecretGenerateResponse) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *WebhookSecretGenerateResponse) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *WebhookSecretGenerateResponse) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.


### GetMessage

`func (o *WebhookSecretGenerateResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WebhookSecretGenerateResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WebhookSecretGenerateResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCreatedAt

`func (o *WebhookSecretGenerateResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookSecretGenerateResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookSecretGenerateResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# WebhookSecretStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasSecret** | **bool** | Whether a webhook secret is configured | 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewWebhookSecretStatusResponse

`func NewWebhookSecretStatusResponse(hasSecret bool, ) *WebhookSecretStatusResponse`

NewWebhookSecretStatusResponse instantiates a new WebhookSecretStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSecretStatusResponseWithDefaults

`func NewWebhookSecretStatusResponseWithDefaults() *WebhookSecretStatusResponse`

NewWebhookSecretStatusResponseWithDefaults instantiates a new WebhookSecretStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasSecret

`func (o *WebhookSecretStatusResponse) GetHasSecret() bool`

GetHasSecret returns the HasSecret field if non-nil, zero value otherwise.

### GetHasSecretOk

`func (o *WebhookSecretStatusResponse) GetHasSecretOk() (*bool, bool)`

GetHasSecretOk returns a tuple with the HasSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSecret

`func (o *WebhookSecretStatusResponse) SetHasSecret(v bool)`

SetHasSecret sets HasSecret field to given value.


### GetCreatedAt

`func (o *WebhookSecretStatusResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookSecretStatusResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookSecretStatusResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WebhookSecretStatusResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *WebhookSecretStatusResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *WebhookSecretStatusResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AFNORWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookId** | Pointer to **string** |  | [optional] 
**SigningKey** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CallbackUrl** | **string** |  | 
**FlowTypes** | Pointer to [**[]AFNORFlowType**](AFNORFlowType.md) |  | [optional] 
**FlowDirection** | Pointer to [**AFNORFlowDirection**](AFNORFlowDirection.md) |  | [optional] 
**AckStatus** | Pointer to [**AFNORFlowAckStatus**](AFNORFlowAckStatus.md) |  | [optional] 

## Methods

### NewAFNORWebhook

`func NewAFNORWebhook(callbackUrl string, ) *AFNORWebhook`

NewAFNORWebhook instantiates a new AFNORWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORWebhookWithDefaults

`func NewAFNORWebhookWithDefaults() *AFNORWebhook`

NewAFNORWebhookWithDefaults instantiates a new AFNORWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookId

`func (o *AFNORWebhook) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *AFNORWebhook) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *AFNORWebhook) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *AFNORWebhook) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetSigningKey

`func (o *AFNORWebhook) GetSigningKey() string`

GetSigningKey returns the SigningKey field if non-nil, zero value otherwise.

### GetSigningKeyOk

`func (o *AFNORWebhook) GetSigningKeyOk() (*string, bool)`

GetSigningKeyOk returns a tuple with the SigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKey

`func (o *AFNORWebhook) SetSigningKey(v string)`

SetSigningKey sets SigningKey field to given value.

### HasSigningKey

`func (o *AFNORWebhook) HasSigningKey() bool`

HasSigningKey returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AFNORWebhook) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AFNORWebhook) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AFNORWebhook) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AFNORWebhook) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *AFNORWebhook) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *AFNORWebhook) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *AFNORWebhook) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetFlowTypes

`func (o *AFNORWebhook) GetFlowTypes() []AFNORFlowType`

GetFlowTypes returns the FlowTypes field if non-nil, zero value otherwise.

### GetFlowTypesOk

`func (o *AFNORWebhook) GetFlowTypesOk() (*[]AFNORFlowType, bool)`

GetFlowTypesOk returns a tuple with the FlowTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowTypes

`func (o *AFNORWebhook) SetFlowTypes(v []AFNORFlowType)`

SetFlowTypes sets FlowTypes field to given value.

### HasFlowTypes

`func (o *AFNORWebhook) HasFlowTypes() bool`

HasFlowTypes returns a boolean if a field has been set.

### GetFlowDirection

`func (o *AFNORWebhook) GetFlowDirection() AFNORFlowDirection`

GetFlowDirection returns the FlowDirection field if non-nil, zero value otherwise.

### GetFlowDirectionOk

`func (o *AFNORWebhook) GetFlowDirectionOk() (*AFNORFlowDirection, bool)`

GetFlowDirectionOk returns a tuple with the FlowDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDirection

`func (o *AFNORWebhook) SetFlowDirection(v AFNORFlowDirection)`

SetFlowDirection sets FlowDirection field to given value.

### HasFlowDirection

`func (o *AFNORWebhook) HasFlowDirection() bool`

HasFlowDirection returns a boolean if a field has been set.

### GetAckStatus

`func (o *AFNORWebhook) GetAckStatus() AFNORFlowAckStatus`

GetAckStatus returns the AckStatus field if non-nil, zero value otherwise.

### GetAckStatusOk

`func (o *AFNORWebhook) GetAckStatusOk() (*AFNORFlowAckStatus, bool)`

GetAckStatusOk returns a tuple with the AckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckStatus

`func (o *AFNORWebhook) SetAckStatus(v AFNORFlowAckStatus)`

SetAckStatus sets AckStatus field to given value.

### HasAckStatus

`func (o *AFNORWebhook) HasAckStatus() bool`

HasAckStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



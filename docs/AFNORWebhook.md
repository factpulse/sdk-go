# AFNORWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookId** | Pointer to **string** |  | [optional] 
**Callback** | [**AFNORCallbackParameters**](AFNORCallbackParameters.md) |  | 
**Metadata** | [**AFNORWebhookMetadata**](AFNORWebhookMetadata.md) |  | 

## Methods

### NewAFNORWebhook

`func NewAFNORWebhook(callback AFNORCallbackParameters, metadata AFNORWebhookMetadata, ) *AFNORWebhook`

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

### GetCallback

`func (o *AFNORWebhook) GetCallback() AFNORCallbackParameters`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *AFNORWebhook) GetCallbackOk() (*AFNORCallbackParameters, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *AFNORWebhook) SetCallback(v AFNORCallbackParameters)`

SetCallback sets Callback field to given value.


### GetMetadata

`func (o *AFNORWebhook) GetMetadata() AFNORWebhookMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AFNORWebhook) GetMetadataOk() (*AFNORWebhookMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AFNORWebhook) SetMetadata(v AFNORWebhookMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



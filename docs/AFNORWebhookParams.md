# AFNORWebhookParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUrl** | **string** |  | 
**FlowTypes** | Pointer to [**[]AFNORFlowType**](AFNORFlowType.md) |  | [optional] 
**FlowDirection** | Pointer to [**AFNORFlowDirection**](AFNORFlowDirection.md) |  | [optional] 
**AckStatus** | Pointer to [**AFNORFlowAckStatus**](AFNORFlowAckStatus.md) |  | [optional] 

## Methods

### NewAFNORWebhookParams

`func NewAFNORWebhookParams(callbackUrl string, ) *AFNORWebhookParams`

NewAFNORWebhookParams instantiates a new AFNORWebhookParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORWebhookParamsWithDefaults

`func NewAFNORWebhookParamsWithDefaults() *AFNORWebhookParams`

NewAFNORWebhookParamsWithDefaults instantiates a new AFNORWebhookParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackUrl

`func (o *AFNORWebhookParams) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *AFNORWebhookParams) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *AFNORWebhookParams) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetFlowTypes

`func (o *AFNORWebhookParams) GetFlowTypes() []AFNORFlowType`

GetFlowTypes returns the FlowTypes field if non-nil, zero value otherwise.

### GetFlowTypesOk

`func (o *AFNORWebhookParams) GetFlowTypesOk() (*[]AFNORFlowType, bool)`

GetFlowTypesOk returns a tuple with the FlowTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowTypes

`func (o *AFNORWebhookParams) SetFlowTypes(v []AFNORFlowType)`

SetFlowTypes sets FlowTypes field to given value.

### HasFlowTypes

`func (o *AFNORWebhookParams) HasFlowTypes() bool`

HasFlowTypes returns a boolean if a field has been set.

### GetFlowDirection

`func (o *AFNORWebhookParams) GetFlowDirection() AFNORFlowDirection`

GetFlowDirection returns the FlowDirection field if non-nil, zero value otherwise.

### GetFlowDirectionOk

`func (o *AFNORWebhookParams) GetFlowDirectionOk() (*AFNORFlowDirection, bool)`

GetFlowDirectionOk returns a tuple with the FlowDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDirection

`func (o *AFNORWebhookParams) SetFlowDirection(v AFNORFlowDirection)`

SetFlowDirection sets FlowDirection field to given value.

### HasFlowDirection

`func (o *AFNORWebhookParams) HasFlowDirection() bool`

HasFlowDirection returns a boolean if a field has been set.

### GetAckStatus

`func (o *AFNORWebhookParams) GetAckStatus() AFNORFlowAckStatus`

GetAckStatus returns the AckStatus field if non-nil, zero value otherwise.

### GetAckStatusOk

`func (o *AFNORWebhookParams) GetAckStatusOk() (*AFNORFlowAckStatus, bool)`

GetAckStatusOk returns a tuple with the AckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckStatus

`func (o *AFNORWebhookParams) SetAckStatus(v AFNORFlowAckStatus)`

SetAckStatus sets AckStatus field to given value.

### HasAckStatus

`func (o *AFNORWebhookParams) HasAckStatus() bool`

HasAckStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



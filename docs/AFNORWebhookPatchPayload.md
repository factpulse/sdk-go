# AFNORWebhookPatchPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | Pointer to [**[]AFNORCallbackHeader**](AFNORCallbackHeader.md) |  | [optional] 
**Authentication** | Pointer to [**AFNORCallbackAuthentication**](AFNORCallbackAuthentication.md) |  | [optional] 
**Signature** | Pointer to [**AFNORCallbackSignature**](AFNORCallbackSignature.md) |  | [optional] 

## Methods

### NewAFNORWebhookPatchPayload

`func NewAFNORWebhookPatchPayload() *AFNORWebhookPatchPayload`

NewAFNORWebhookPatchPayload instantiates a new AFNORWebhookPatchPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORWebhookPatchPayloadWithDefaults

`func NewAFNORWebhookPatchPayloadWithDefaults() *AFNORWebhookPatchPayload`

NewAFNORWebhookPatchPayloadWithDefaults instantiates a new AFNORWebhookPatchPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *AFNORWebhookPatchPayload) GetHeaders() []AFNORCallbackHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *AFNORWebhookPatchPayload) GetHeadersOk() (*[]AFNORCallbackHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *AFNORWebhookPatchPayload) SetHeaders(v []AFNORCallbackHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *AFNORWebhookPatchPayload) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetAuthentication

`func (o *AFNORWebhookPatchPayload) GetAuthentication() AFNORCallbackAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *AFNORWebhookPatchPayload) GetAuthenticationOk() (*AFNORCallbackAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *AFNORWebhookPatchPayload) SetAuthentication(v AFNORCallbackAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *AFNORWebhookPatchPayload) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetSignature

`func (o *AFNORWebhookPatchPayload) GetSignature() AFNORCallbackSignature`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *AFNORWebhookPatchPayload) GetSignatureOk() (*AFNORCallbackSignature, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *AFNORWebhookPatchPayload) SetSignature(v AFNORCallbackSignature)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *AFNORWebhookPatchPayload) HasSignature() bool`

HasSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SecretStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Secret status: &#39;active&#39;, &#39;missing&#39;, etc. | 
**Message** | **string** | Descriptive status message | 
**EncryptionMode** | Pointer to **NullableString** |  | [optional] 
**RequiresClientKey** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewSecretStatus

`func NewSecretStatus(status string, message string, ) *SecretStatus`

NewSecretStatus instantiates a new SecretStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretStatusWithDefaults

`func NewSecretStatusWithDefaults() *SecretStatus`

NewSecretStatusWithDefaults instantiates a new SecretStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SecretStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecretStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecretStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *SecretStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SecretStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SecretStatus) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetEncryptionMode

`func (o *SecretStatus) GetEncryptionMode() string`

GetEncryptionMode returns the EncryptionMode field if non-nil, zero value otherwise.

### GetEncryptionModeOk

`func (o *SecretStatus) GetEncryptionModeOk() (*string, bool)`

GetEncryptionModeOk returns a tuple with the EncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionMode

`func (o *SecretStatus) SetEncryptionMode(v string)`

SetEncryptionMode sets EncryptionMode field to given value.

### HasEncryptionMode

`func (o *SecretStatus) HasEncryptionMode() bool`

HasEncryptionMode returns a boolean if a field has been set.

### SetEncryptionModeNil

`func (o *SecretStatus) SetEncryptionModeNil(b bool)`

 SetEncryptionModeNil sets the value for EncryptionMode to be an explicit nil

### UnsetEncryptionMode
`func (o *SecretStatus) UnsetEncryptionMode()`

UnsetEncryptionMode ensures that no value is present for EncryptionMode, not even an explicit nil
### GetRequiresClientKey

`func (o *SecretStatus) GetRequiresClientKey() bool`

GetRequiresClientKey returns the RequiresClientKey field if non-nil, zero value otherwise.

### GetRequiresClientKeyOk

`func (o *SecretStatus) GetRequiresClientKeyOk() (*bool, bool)`

GetRequiresClientKeyOk returns a tuple with the RequiresClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresClientKey

`func (o *SecretStatus) SetRequiresClientKey(v bool)`

SetRequiresClientKey sets RequiresClientKey field to given value.

### HasRequiresClientKey

`func (o *SecretStatus) HasRequiresClientKey() bool`

HasRequiresClientKey returns a boolean if a field has been set.

### SetRequiresClientKeyNil

`func (o *SecretStatus) SetRequiresClientKeyNil(b bool)`

 SetRequiresClientKeyNil sets the value for RequiresClientKey to be an explicit nil

### UnsetRequiresClientKey
`func (o *SecretStatus) UnsetRequiresClientKey()`

UnsetRequiresClientKey ensures that no value is present for RequiresClientKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



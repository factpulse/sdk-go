# SecretStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Secret status: &#39;active&#39;, &#39;missing&#39;, etc. | 
**Message** | **string** | Descriptive status message | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



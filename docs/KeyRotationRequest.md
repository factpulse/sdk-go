# KeyRotationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OldKey** | **string** | Current encryption key (base64-encoded AES-256) | 
**NewKey** | **string** | New encryption key (base64-encoded AES-256) | 

## Methods

### NewKeyRotationRequest

`func NewKeyRotationRequest(oldKey string, newKey string, ) *KeyRotationRequest`

NewKeyRotationRequest instantiates a new KeyRotationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyRotationRequestWithDefaults

`func NewKeyRotationRequestWithDefaults() *KeyRotationRequest`

NewKeyRotationRequestWithDefaults instantiates a new KeyRotationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOldKey

`func (o *KeyRotationRequest) GetOldKey() string`

GetOldKey returns the OldKey field if non-nil, zero value otherwise.

### GetOldKeyOk

`func (o *KeyRotationRequest) GetOldKeyOk() (*string, bool)`

GetOldKeyOk returns a tuple with the OldKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldKey

`func (o *KeyRotationRequest) SetOldKey(v string)`

SetOldKey sets OldKey field to given value.


### GetNewKey

`func (o *KeyRotationRequest) GetNewKey() string`

GetNewKey returns the NewKey field if non-nil, zero value otherwise.

### GetNewKeyOk

`func (o *KeyRotationRequest) GetNewKeyOk() (*string, bool)`

GetNewKeyOk returns a tuple with the NewKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewKey

`func (o *KeyRotationRequest) SetNewKey(v string)`

SetNewKey sets NewKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# KeyRotationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Whether rotation was successful | 
**Message** | **string** | Result message | 
**RotatedCount** | **int32** | Number of secrets that were rotated | 
**PartialErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewKeyRotationResponse

`func NewKeyRotationResponse(success bool, message string, rotatedCount int32, ) *KeyRotationResponse`

NewKeyRotationResponse instantiates a new KeyRotationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyRotationResponseWithDefaults

`func NewKeyRotationResponseWithDefaults() *KeyRotationResponse`

NewKeyRotationResponseWithDefaults instantiates a new KeyRotationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *KeyRotationResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *KeyRotationResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *KeyRotationResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetMessage

`func (o *KeyRotationResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KeyRotationResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KeyRotationResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRotatedCount

`func (o *KeyRotationResponse) GetRotatedCount() int32`

GetRotatedCount returns the RotatedCount field if non-nil, zero value otherwise.

### GetRotatedCountOk

`func (o *KeyRotationResponse) GetRotatedCountOk() (*int32, bool)`

GetRotatedCountOk returns a tuple with the RotatedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedCount

`func (o *KeyRotationResponse) SetRotatedCount(v int32)`

SetRotatedCount sets RotatedCount field to given value.


### GetPartialErrors

`func (o *KeyRotationResponse) GetPartialErrors() []string`

GetPartialErrors returns the PartialErrors field if non-nil, zero value otherwise.

### GetPartialErrorsOk

`func (o *KeyRotationResponse) GetPartialErrorsOk() (*[]string, bool)`

GetPartialErrorsOk returns a tuple with the PartialErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialErrors

`func (o *KeyRotationResponse) SetPartialErrors(v []string)`

SetPartialErrors sets PartialErrors field to given value.

### HasPartialErrors

`func (o *KeyRotationResponse) HasPartialErrors() bool`

HasPartialErrors returns a boolean if a field has been set.

### SetPartialErrorsNil

`func (o *KeyRotationResponse) SetPartialErrorsNil(b bool)`

 SetPartialErrorsNil sets the value for PartialErrors to be an explicit nil

### UnsetPartialErrors
`func (o *KeyRotationResponse) UnsetPartialErrors()`

UnsetPartialErrors ensures that no value is present for PartialErrors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



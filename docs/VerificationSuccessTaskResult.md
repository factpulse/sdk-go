# VerificationSuccessTaskResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "SUCCESS"]
**VerificationResult** | [**VerificationSuccessResponse**](VerificationSuccessResponse.md) |  | 

## Methods

### NewVerificationSuccessTaskResult

`func NewVerificationSuccessTaskResult(verificationResult VerificationSuccessResponse, ) *VerificationSuccessTaskResult`

NewVerificationSuccessTaskResult instantiates a new VerificationSuccessTaskResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationSuccessTaskResultWithDefaults

`func NewVerificationSuccessTaskResultWithDefaults() *VerificationSuccessTaskResult`

NewVerificationSuccessTaskResultWithDefaults instantiates a new VerificationSuccessTaskResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VerificationSuccessTaskResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VerificationSuccessTaskResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VerificationSuccessTaskResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VerificationSuccessTaskResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVerificationResult

`func (o *VerificationSuccessTaskResult) GetVerificationResult() VerificationSuccessResponse`

GetVerificationResult returns the VerificationResult field if non-nil, zero value otherwise.

### GetVerificationResultOk

`func (o *VerificationSuccessTaskResult) GetVerificationResultOk() (*VerificationSuccessResponse, bool)`

GetVerificationResultOk returns a tuple with the VerificationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationResult

`func (o *VerificationSuccessTaskResult) SetVerificationResult(v VerificationSuccessResponse)`

SetVerificationResult sets VerificationResult field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# VerificationTypedTaskStatusResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "ERROR"]
**VerificationResult** | [**VerificationSuccessResponse**](VerificationSuccessResponse.md) |  | 
**ErrorCode** | **string** |  | 
**ErrorMessage** | **string** |  | 
**Details** | Pointer to [**[]AFNORErrorDetail**](AFNORErrorDetail.md) |  | [optional] 
**Traceback** | Pointer to **string** |  | [optional] 

## Methods

### NewVerificationTypedTaskStatusResult

`func NewVerificationTypedTaskStatusResult(verificationResult VerificationSuccessResponse, errorCode string, errorMessage string, ) *VerificationTypedTaskStatusResult`

NewVerificationTypedTaskStatusResult instantiates a new VerificationTypedTaskStatusResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationTypedTaskStatusResultWithDefaults

`func NewVerificationTypedTaskStatusResultWithDefaults() *VerificationTypedTaskStatusResult`

NewVerificationTypedTaskStatusResultWithDefaults instantiates a new VerificationTypedTaskStatusResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VerificationTypedTaskStatusResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VerificationTypedTaskStatusResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VerificationTypedTaskStatusResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VerificationTypedTaskStatusResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVerificationResult

`func (o *VerificationTypedTaskStatusResult) GetVerificationResult() VerificationSuccessResponse`

GetVerificationResult returns the VerificationResult field if non-nil, zero value otherwise.

### GetVerificationResultOk

`func (o *VerificationTypedTaskStatusResult) GetVerificationResultOk() (*VerificationSuccessResponse, bool)`

GetVerificationResultOk returns a tuple with the VerificationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationResult

`func (o *VerificationTypedTaskStatusResult) SetVerificationResult(v VerificationSuccessResponse)`

SetVerificationResult sets VerificationResult field to given value.


### GetErrorCode

`func (o *VerificationTypedTaskStatusResult) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *VerificationTypedTaskStatusResult) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *VerificationTypedTaskStatusResult) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorMessage

`func (o *VerificationTypedTaskStatusResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *VerificationTypedTaskStatusResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *VerificationTypedTaskStatusResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### GetDetails

`func (o *VerificationTypedTaskStatusResult) GetDetails() []AFNORErrorDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *VerificationTypedTaskStatusResult) GetDetailsOk() (*[]AFNORErrorDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *VerificationTypedTaskStatusResult) SetDetails(v []AFNORErrorDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *VerificationTypedTaskStatusResult) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTraceback

`func (o *VerificationTypedTaskStatusResult) GetTraceback() string`

GetTraceback returns the Traceback field if non-nil, zero value otherwise.

### GetTracebackOk

`func (o *VerificationTypedTaskStatusResult) GetTracebackOk() (*string, bool)`

GetTracebackOk returns a tuple with the Traceback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceback

`func (o *VerificationTypedTaskStatusResult) SetTraceback(v string)`

SetTraceback sets Traceback field to given value.

### HasTraceback

`func (o *VerificationTypedTaskStatusResult) HasTraceback() bool`

HasTraceback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



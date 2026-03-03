# AFNORError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | **string** | Short numerical or alphanumerical code that identifies precisely a unique error. | 
**ErrorMessage** | Pointer to **string** | Contains information on the error. Not intended to be displayed to an end user. For security reasons, a tradeoff between clarity &amp; security shall be found. | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "about:blank"]
**Details** | Pointer to **string** |  | [optional] 
**Instance** | Pointer to **string** |  | [optional] 

## Methods

### NewAFNORError

`func NewAFNORError(errorCode string, ) *AFNORError`

NewAFNORError instantiates a new AFNORError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORErrorWithDefaults

`func NewAFNORErrorWithDefaults() *AFNORError`

NewAFNORErrorWithDefaults instantiates a new AFNORError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *AFNORError) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *AFNORError) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *AFNORError) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorMessage

`func (o *AFNORError) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AFNORError) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AFNORError) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AFNORError) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetType

`func (o *AFNORError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AFNORError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AFNORError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AFNORError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDetails

`func (o *AFNORError) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AFNORError) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AFNORError) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AFNORError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetInstance

`func (o *AFNORError) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *AFNORError) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *AFNORError) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *AFNORError) HasInstance() bool`

HasInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



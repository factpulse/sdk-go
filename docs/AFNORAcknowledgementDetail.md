# AFNORAcknowledgementDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | **string** |  | 
**Item** | **string** | Item on which the error refers | 
**ReasonCode** | [**AFNORReasonCode**](AFNORReasonCode.md) |  | 
**ReasonMessage** | **string** |  | 

## Methods

### NewAFNORAcknowledgementDetail

`func NewAFNORAcknowledgementDetail(level string, item string, reasonCode AFNORReasonCode, reasonMessage string, ) *AFNORAcknowledgementDetail`

NewAFNORAcknowledgementDetail instantiates a new AFNORAcknowledgementDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORAcknowledgementDetailWithDefaults

`func NewAFNORAcknowledgementDetailWithDefaults() *AFNORAcknowledgementDetail`

NewAFNORAcknowledgementDetailWithDefaults instantiates a new AFNORAcknowledgementDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *AFNORAcknowledgementDetail) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AFNORAcknowledgementDetail) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AFNORAcknowledgementDetail) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetItem

`func (o *AFNORAcknowledgementDetail) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *AFNORAcknowledgementDetail) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *AFNORAcknowledgementDetail) SetItem(v string)`

SetItem sets Item field to given value.


### GetReasonCode

`func (o *AFNORAcknowledgementDetail) GetReasonCode() AFNORReasonCode`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *AFNORAcknowledgementDetail) GetReasonCodeOk() (*AFNORReasonCode, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *AFNORAcknowledgementDetail) SetReasonCode(v AFNORReasonCode)`

SetReasonCode sets ReasonCode field to given value.


### GetReasonMessage

`func (o *AFNORAcknowledgementDetail) GetReasonMessage() string`

GetReasonMessage returns the ReasonMessage field if non-nil, zero value otherwise.

### GetReasonMessageOk

`func (o *AFNORAcknowledgementDetail) GetReasonMessageOk() (*string, bool)`

GetReasonMessageOk returns a tuple with the ReasonMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonMessage

`func (o *AFNORAcknowledgementDetail) SetReasonMessage(v string)`

SetReasonMessage sets ReasonMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



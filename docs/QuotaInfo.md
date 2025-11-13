# QuotaInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** |  | 
**Usage** | **int32** |  | 
**Remaining** | **int32** |  | 
**ResetDate** | **string** |  | 
**Plan** | **string** |  | 

## Methods

### NewQuotaInfo

`func NewQuotaInfo(limit int32, usage int32, remaining int32, resetDate string, plan string, ) *QuotaInfo`

NewQuotaInfo instantiates a new QuotaInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaInfoWithDefaults

`func NewQuotaInfoWithDefaults() *QuotaInfo`

NewQuotaInfoWithDefaults instantiates a new QuotaInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *QuotaInfo) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QuotaInfo) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QuotaInfo) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetUsage

`func (o *QuotaInfo) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *QuotaInfo) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *QuotaInfo) SetUsage(v int32)`

SetUsage sets Usage field to given value.


### GetRemaining

`func (o *QuotaInfo) GetRemaining() int32`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *QuotaInfo) GetRemainingOk() (*int32, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *QuotaInfo) SetRemaining(v int32)`

SetRemaining sets Remaining field to given value.


### GetResetDate

`func (o *QuotaInfo) GetResetDate() string`

GetResetDate returns the ResetDate field if non-nil, zero value otherwise.

### GetResetDateOk

`func (o *QuotaInfo) GetResetDateOk() (*string, bool)`

GetResetDateOk returns a tuple with the ResetDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetDate

`func (o *QuotaInfo) SetResetDate(v string)`

SetResetDate sets ResetDate field to given value.


### GetPlan

`func (o *QuotaInfo) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *QuotaInfo) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *QuotaInfo) SetPlan(v string)`

SetPlan sets Plan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



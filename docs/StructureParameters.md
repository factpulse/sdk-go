# StructureParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceCodeRequired** | Pointer to **bool** | Service code is mandatory | [optional] [default to false]
**EngagementNumberRequired** | Pointer to **bool** | Engagement number is mandatory | [optional] [default to false]
**EngagementOrServiceManagement** | Pointer to **bool** | EJ or service code management enabled | [optional] [default to false]

## Methods

### NewStructureParameters

`func NewStructureParameters() *StructureParameters`

NewStructureParameters instantiates a new StructureParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructureParametersWithDefaults

`func NewStructureParametersWithDefaults() *StructureParameters`

NewStructureParametersWithDefaults instantiates a new StructureParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceCodeRequired

`func (o *StructureParameters) GetServiceCodeRequired() bool`

GetServiceCodeRequired returns the ServiceCodeRequired field if non-nil, zero value otherwise.

### GetServiceCodeRequiredOk

`func (o *StructureParameters) GetServiceCodeRequiredOk() (*bool, bool)`

GetServiceCodeRequiredOk returns a tuple with the ServiceCodeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCodeRequired

`func (o *StructureParameters) SetServiceCodeRequired(v bool)`

SetServiceCodeRequired sets ServiceCodeRequired field to given value.

### HasServiceCodeRequired

`func (o *StructureParameters) HasServiceCodeRequired() bool`

HasServiceCodeRequired returns a boolean if a field has been set.

### GetEngagementNumberRequired

`func (o *StructureParameters) GetEngagementNumberRequired() bool`

GetEngagementNumberRequired returns the EngagementNumberRequired field if non-nil, zero value otherwise.

### GetEngagementNumberRequiredOk

`func (o *StructureParameters) GetEngagementNumberRequiredOk() (*bool, bool)`

GetEngagementNumberRequiredOk returns a tuple with the EngagementNumberRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementNumberRequired

`func (o *StructureParameters) SetEngagementNumberRequired(v bool)`

SetEngagementNumberRequired sets EngagementNumberRequired field to given value.

### HasEngagementNumberRequired

`func (o *StructureParameters) HasEngagementNumberRequired() bool`

HasEngagementNumberRequired returns a boolean if a field has been set.

### GetEngagementOrServiceManagement

`func (o *StructureParameters) GetEngagementOrServiceManagement() bool`

GetEngagementOrServiceManagement returns the EngagementOrServiceManagement field if non-nil, zero value otherwise.

### GetEngagementOrServiceManagementOk

`func (o *StructureParameters) GetEngagementOrServiceManagementOk() (*bool, bool)`

GetEngagementOrServiceManagementOk returns a tuple with the EngagementOrServiceManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementOrServiceManagement

`func (o *StructureParameters) SetEngagementOrServiceManagement(v bool)`

SetEngagementOrServiceManagement sets EngagementOrServiceManagement field to given value.

### HasEngagementOrServiceManagement

`func (o *StructureParameters) HasEngagementOrServiceManagement() bool`

HasEngagementOrServiceManagement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



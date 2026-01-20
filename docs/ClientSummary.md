# ClientSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | Unique client identifier | 
**Name** | **string** | Client name | 
**Siret** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsActive** | **bool** | Whether the client is active | 
**HasConfigPdp** | **bool** | Whether PDP config exists | 
**PdpIsActive** | Pointer to **NullableBool** |  | [optional] 
**PdpIsMock** | Pointer to **NullableBool** |  | [optional] 
**HasConfigChorus** | **bool** | Whether Chorus Pro config exists | 
**CreatedAt** | **time.Time** | Creation date | 
**UpdatedAt** | **time.Time** | Last update date | 

## Methods

### NewClientSummary

`func NewClientSummary(uid string, name string, isActive bool, hasConfigPdp bool, hasConfigChorus bool, createdAt time.Time, updatedAt time.Time, ) *ClientSummary`

NewClientSummary instantiates a new ClientSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientSummaryWithDefaults

`func NewClientSummaryWithDefaults() *ClientSummary`

NewClientSummaryWithDefaults instantiates a new ClientSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *ClientSummary) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ClientSummary) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ClientSummary) SetUid(v string)`

SetUid sets Uid field to given value.


### GetName

`func (o *ClientSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientSummary) SetName(v string)`

SetName sets Name field to given value.


### GetSiret

`func (o *ClientSummary) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *ClientSummary) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *ClientSummary) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *ClientSummary) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### SetSiretNil

`func (o *ClientSummary) SetSiretNil(b bool)`

 SetSiretNil sets the value for Siret to be an explicit nil

### UnsetSiret
`func (o *ClientSummary) UnsetSiret()`

UnsetSiret ensures that no value is present for Siret, not even an explicit nil
### GetDescription

`func (o *ClientSummary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClientSummary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClientSummary) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClientSummary) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ClientSummary) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ClientSummary) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsActive

`func (o *ClientSummary) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ClientSummary) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ClientSummary) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetHasConfigPdp

`func (o *ClientSummary) GetHasConfigPdp() bool`

GetHasConfigPdp returns the HasConfigPdp field if non-nil, zero value otherwise.

### GetHasConfigPdpOk

`func (o *ClientSummary) GetHasConfigPdpOk() (*bool, bool)`

GetHasConfigPdpOk returns a tuple with the HasConfigPdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConfigPdp

`func (o *ClientSummary) SetHasConfigPdp(v bool)`

SetHasConfigPdp sets HasConfigPdp field to given value.


### GetPdpIsActive

`func (o *ClientSummary) GetPdpIsActive() bool`

GetPdpIsActive returns the PdpIsActive field if non-nil, zero value otherwise.

### GetPdpIsActiveOk

`func (o *ClientSummary) GetPdpIsActiveOk() (*bool, bool)`

GetPdpIsActiveOk returns a tuple with the PdpIsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpIsActive

`func (o *ClientSummary) SetPdpIsActive(v bool)`

SetPdpIsActive sets PdpIsActive field to given value.

### HasPdpIsActive

`func (o *ClientSummary) HasPdpIsActive() bool`

HasPdpIsActive returns a boolean if a field has been set.

### SetPdpIsActiveNil

`func (o *ClientSummary) SetPdpIsActiveNil(b bool)`

 SetPdpIsActiveNil sets the value for PdpIsActive to be an explicit nil

### UnsetPdpIsActive
`func (o *ClientSummary) UnsetPdpIsActive()`

UnsetPdpIsActive ensures that no value is present for PdpIsActive, not even an explicit nil
### GetPdpIsMock

`func (o *ClientSummary) GetPdpIsMock() bool`

GetPdpIsMock returns the PdpIsMock field if non-nil, zero value otherwise.

### GetPdpIsMockOk

`func (o *ClientSummary) GetPdpIsMockOk() (*bool, bool)`

GetPdpIsMockOk returns a tuple with the PdpIsMock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpIsMock

`func (o *ClientSummary) SetPdpIsMock(v bool)`

SetPdpIsMock sets PdpIsMock field to given value.

### HasPdpIsMock

`func (o *ClientSummary) HasPdpIsMock() bool`

HasPdpIsMock returns a boolean if a field has been set.

### SetPdpIsMockNil

`func (o *ClientSummary) SetPdpIsMockNil(b bool)`

 SetPdpIsMockNil sets the value for PdpIsMock to be an explicit nil

### UnsetPdpIsMock
`func (o *ClientSummary) UnsetPdpIsMock()`

UnsetPdpIsMock ensures that no value is present for PdpIsMock, not even an explicit nil
### GetHasConfigChorus

`func (o *ClientSummary) GetHasConfigChorus() bool`

GetHasConfigChorus returns the HasConfigChorus field if non-nil, zero value otherwise.

### GetHasConfigChorusOk

`func (o *ClientSummary) GetHasConfigChorusOk() (*bool, bool)`

GetHasConfigChorusOk returns a tuple with the HasConfigChorus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConfigChorus

`func (o *ClientSummary) SetHasConfigChorus(v bool)`

SetHasConfigChorus sets HasConfigChorus field to given value.


### GetCreatedAt

`func (o *ClientSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClientSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClientSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ClientSummary) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ClientSummary) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ClientSummary) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



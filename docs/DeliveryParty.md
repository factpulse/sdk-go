# DeliveryParty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**GlobalId** | Pointer to [**NullableElectronicAddress**](ElectronicAddress.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**PostalAddress** | Pointer to [**NullablePostalAddress**](PostalAddress.md) |  | [optional] 

## Methods

### NewDeliveryParty

`func NewDeliveryParty() *DeliveryParty`

NewDeliveryParty instantiates a new DeliveryParty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryPartyWithDefaults

`func NewDeliveryPartyWithDefaults() *DeliveryParty`

NewDeliveryPartyWithDefaults instantiates a new DeliveryParty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeliveryParty) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeliveryParty) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeliveryParty) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeliveryParty) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DeliveryParty) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DeliveryParty) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetGlobalId

`func (o *DeliveryParty) GetGlobalId() ElectronicAddress`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *DeliveryParty) GetGlobalIdOk() (*ElectronicAddress, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *DeliveryParty) SetGlobalId(v ElectronicAddress)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *DeliveryParty) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### SetGlobalIdNil

`func (o *DeliveryParty) SetGlobalIdNil(b bool)`

 SetGlobalIdNil sets the value for GlobalId to be an explicit nil

### UnsetGlobalId
`func (o *DeliveryParty) UnsetGlobalId()`

UnsetGlobalId ensures that no value is present for GlobalId, not even an explicit nil
### GetName

`func (o *DeliveryParty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeliveryParty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeliveryParty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeliveryParty) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DeliveryParty) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeliveryParty) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPostalAddress

`func (o *DeliveryParty) GetPostalAddress() PostalAddress`

GetPostalAddress returns the PostalAddress field if non-nil, zero value otherwise.

### GetPostalAddressOk

`func (o *DeliveryParty) GetPostalAddressOk() (*PostalAddress, bool)`

GetPostalAddressOk returns a tuple with the PostalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalAddress

`func (o *DeliveryParty) SetPostalAddress(v PostalAddress)`

SetPostalAddress sets PostalAddress field to given value.

### HasPostalAddress

`func (o *DeliveryParty) HasPostalAddress() bool`

HasPostalAddress returns a boolean if a field has been set.

### SetPostalAddressNil

`func (o *DeliveryParty) SetPostalAddressNil(b bool)`

 SetPostalAddressNil sets the value for PostalAddress to be an explicit nil

### UnsetPostalAddress
`func (o *DeliveryParty) UnsetPostalAddress()`

UnsetPostalAddress ensures that no value is present for PostalAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



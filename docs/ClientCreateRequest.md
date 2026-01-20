# ClientCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Client name | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Siret** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewClientCreateRequest

`func NewClientCreateRequest(name string, ) *ClientCreateRequest`

NewClientCreateRequest instantiates a new ClientCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCreateRequestWithDefaults

`func NewClientCreateRequestWithDefaults() *ClientCreateRequest`

NewClientCreateRequestWithDefaults instantiates a new ClientCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClientCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ClientCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClientCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClientCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClientCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ClientCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ClientCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSiret

`func (o *ClientCreateRequest) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *ClientCreateRequest) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *ClientCreateRequest) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *ClientCreateRequest) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### SetSiretNil

`func (o *ClientCreateRequest) SetSiretNil(b bool)`

 SetSiretNil sets the value for Siret to be an explicit nil

### UnsetSiret
`func (o *ClientCreateRequest) UnsetSiret()`

UnsetSiret ensures that no value is present for Siret, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



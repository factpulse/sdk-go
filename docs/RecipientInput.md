# RecipientInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Siren** | Pointer to **NullableString** |  | [optional] 
**Siret** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Role** | Pointer to **string** | Code rôle (BY, SE, WK, etc.) | [optional] [default to "BY"]
**Email** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRecipientInput

`func NewRecipientInput() *RecipientInput`

NewRecipientInput instantiates a new RecipientInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipientInputWithDefaults

`func NewRecipientInputWithDefaults() *RecipientInput`

NewRecipientInputWithDefaults instantiates a new RecipientInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiren

`func (o *RecipientInput) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *RecipientInput) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *RecipientInput) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *RecipientInput) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### SetSirenNil

`func (o *RecipientInput) SetSirenNil(b bool)`

 SetSirenNil sets the value for Siren to be an explicit nil

### UnsetSiren
`func (o *RecipientInput) UnsetSiren()`

UnsetSiren ensures that no value is present for Siren, not even an explicit nil
### GetSiret

`func (o *RecipientInput) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *RecipientInput) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *RecipientInput) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *RecipientInput) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### SetSiretNil

`func (o *RecipientInput) SetSiretNil(b bool)`

 SetSiretNil sets the value for Siret to be an explicit nil

### UnsetSiret
`func (o *RecipientInput) UnsetSiret()`

UnsetSiret ensures that no value is present for Siret, not even an explicit nil
### GetName

`func (o *RecipientInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecipientInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecipientInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecipientInput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RecipientInput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RecipientInput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRole

`func (o *RecipientInput) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RecipientInput) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RecipientInput) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *RecipientInput) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetEmail

`func (o *RecipientInput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RecipientInput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RecipientInput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *RecipientInput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *RecipientInput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *RecipientInput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



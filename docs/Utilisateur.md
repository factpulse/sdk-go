# Utilisateur

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Username** | **string** |  | 
**Email** | **string** |  | 
**IsActive** | **bool** |  | 
**IsSuperuser** | Pointer to **bool** |  | [optional] [default to false]
**IsStaff** | Pointer to **bool** |  | [optional] [default to false]
**BypassQuota** | Pointer to **bool** |  | [optional] [default to false]
**TeamId** | Pointer to **NullableInt32** |  | [optional] 
**HasQuota** | Pointer to **bool** |  | [optional] [default to true]
**QuotaInfo** | Pointer to [**NullableQuotaInfo**](QuotaInfo.md) |  | [optional] 
**IsTrial** | Pointer to **bool** |  | [optional] [default to false]
**ClientUid** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUtilisateur

`func NewUtilisateur(id int32, username string, email string, isActive bool, ) *Utilisateur`

NewUtilisateur instantiates a new Utilisateur object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilisateurWithDefaults

`func NewUtilisateurWithDefaults() *Utilisateur`

NewUtilisateurWithDefaults instantiates a new Utilisateur object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Utilisateur) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Utilisateur) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Utilisateur) SetId(v int32)`

SetId sets Id field to given value.


### GetUsername

`func (o *Utilisateur) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Utilisateur) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Utilisateur) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *Utilisateur) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Utilisateur) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Utilisateur) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetIsActive

`func (o *Utilisateur) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Utilisateur) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Utilisateur) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetIsSuperuser

`func (o *Utilisateur) GetIsSuperuser() bool`

GetIsSuperuser returns the IsSuperuser field if non-nil, zero value otherwise.

### GetIsSuperuserOk

`func (o *Utilisateur) GetIsSuperuserOk() (*bool, bool)`

GetIsSuperuserOk returns a tuple with the IsSuperuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperuser

`func (o *Utilisateur) SetIsSuperuser(v bool)`

SetIsSuperuser sets IsSuperuser field to given value.

### HasIsSuperuser

`func (o *Utilisateur) HasIsSuperuser() bool`

HasIsSuperuser returns a boolean if a field has been set.

### GetIsStaff

`func (o *Utilisateur) GetIsStaff() bool`

GetIsStaff returns the IsStaff field if non-nil, zero value otherwise.

### GetIsStaffOk

`func (o *Utilisateur) GetIsStaffOk() (*bool, bool)`

GetIsStaffOk returns a tuple with the IsStaff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStaff

`func (o *Utilisateur) SetIsStaff(v bool)`

SetIsStaff sets IsStaff field to given value.

### HasIsStaff

`func (o *Utilisateur) HasIsStaff() bool`

HasIsStaff returns a boolean if a field has been set.

### GetBypassQuota

`func (o *Utilisateur) GetBypassQuota() bool`

GetBypassQuota returns the BypassQuota field if non-nil, zero value otherwise.

### GetBypassQuotaOk

`func (o *Utilisateur) GetBypassQuotaOk() (*bool, bool)`

GetBypassQuotaOk returns a tuple with the BypassQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassQuota

`func (o *Utilisateur) SetBypassQuota(v bool)`

SetBypassQuota sets BypassQuota field to given value.

### HasBypassQuota

`func (o *Utilisateur) HasBypassQuota() bool`

HasBypassQuota returns a boolean if a field has been set.

### GetTeamId

`func (o *Utilisateur) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *Utilisateur) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *Utilisateur) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *Utilisateur) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### SetTeamIdNil

`func (o *Utilisateur) SetTeamIdNil(b bool)`

 SetTeamIdNil sets the value for TeamId to be an explicit nil

### UnsetTeamId
`func (o *Utilisateur) UnsetTeamId()`

UnsetTeamId ensures that no value is present for TeamId, not even an explicit nil
### GetHasQuota

`func (o *Utilisateur) GetHasQuota() bool`

GetHasQuota returns the HasQuota field if non-nil, zero value otherwise.

### GetHasQuotaOk

`func (o *Utilisateur) GetHasQuotaOk() (*bool, bool)`

GetHasQuotaOk returns a tuple with the HasQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasQuota

`func (o *Utilisateur) SetHasQuota(v bool)`

SetHasQuota sets HasQuota field to given value.

### HasHasQuota

`func (o *Utilisateur) HasHasQuota() bool`

HasHasQuota returns a boolean if a field has been set.

### GetQuotaInfo

`func (o *Utilisateur) GetQuotaInfo() QuotaInfo`

GetQuotaInfo returns the QuotaInfo field if non-nil, zero value otherwise.

### GetQuotaInfoOk

`func (o *Utilisateur) GetQuotaInfoOk() (*QuotaInfo, bool)`

GetQuotaInfoOk returns a tuple with the QuotaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaInfo

`func (o *Utilisateur) SetQuotaInfo(v QuotaInfo)`

SetQuotaInfo sets QuotaInfo field to given value.

### HasQuotaInfo

`func (o *Utilisateur) HasQuotaInfo() bool`

HasQuotaInfo returns a boolean if a field has been set.

### SetQuotaInfoNil

`func (o *Utilisateur) SetQuotaInfoNil(b bool)`

 SetQuotaInfoNil sets the value for QuotaInfo to be an explicit nil

### UnsetQuotaInfo
`func (o *Utilisateur) UnsetQuotaInfo()`

UnsetQuotaInfo ensures that no value is present for QuotaInfo, not even an explicit nil
### GetIsTrial

`func (o *Utilisateur) GetIsTrial() bool`

GetIsTrial returns the IsTrial field if non-nil, zero value otherwise.

### GetIsTrialOk

`func (o *Utilisateur) GetIsTrialOk() (*bool, bool)`

GetIsTrialOk returns a tuple with the IsTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrial

`func (o *Utilisateur) SetIsTrial(v bool)`

SetIsTrial sets IsTrial field to given value.

### HasIsTrial

`func (o *Utilisateur) HasIsTrial() bool`

HasIsTrial returns a boolean if a field has been set.

### GetClientUid

`func (o *Utilisateur) GetClientUid() string`

GetClientUid returns the ClientUid field if non-nil, zero value otherwise.

### GetClientUidOk

`func (o *Utilisateur) GetClientUidOk() (*string, bool)`

GetClientUidOk returns a tuple with the ClientUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUid

`func (o *Utilisateur) SetClientUid(v string)`

SetClientUid sets ClientUid field to given value.

### HasClientUid

`func (o *Utilisateur) HasClientUid() bool`

HasClientUid returns a boolean if a field has been set.

### SetClientUidNil

`func (o *Utilisateur) SetClientUidNil(b bool)`

 SetClientUidNil sets the value for ClientUid to be an explicit nil

### UnsetClientUid
`func (o *Utilisateur) UnsetClientUid()`

UnsetClientUid ensures that no value is present for ClientUid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



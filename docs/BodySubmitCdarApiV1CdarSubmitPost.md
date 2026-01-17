# BodySubmitCdarApiV1CdarSubmitPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | [**SubmitCDARRequest**](SubmitCDARRequest.md) |  | 
**PdpCredentials** | Pointer to [**NullablePDPCredentials**](PDPCredentials.md) |  | [optional] 

## Methods

### NewBodySubmitCdarApiV1CdarSubmitPost

`func NewBodySubmitCdarApiV1CdarSubmitPost(request SubmitCDARRequest, ) *BodySubmitCdarApiV1CdarSubmitPost`

NewBodySubmitCdarApiV1CdarSubmitPost instantiates a new BodySubmitCdarApiV1CdarSubmitPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBodySubmitCdarApiV1CdarSubmitPostWithDefaults

`func NewBodySubmitCdarApiV1CdarSubmitPostWithDefaults() *BodySubmitCdarApiV1CdarSubmitPost`

NewBodySubmitCdarApiV1CdarSubmitPostWithDefaults instantiates a new BodySubmitCdarApiV1CdarSubmitPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *BodySubmitCdarApiV1CdarSubmitPost) GetRequest() SubmitCDARRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *BodySubmitCdarApiV1CdarSubmitPost) GetRequestOk() (*SubmitCDARRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *BodySubmitCdarApiV1CdarSubmitPost) SetRequest(v SubmitCDARRequest)`

SetRequest sets Request field to given value.


### GetPdpCredentials

`func (o *BodySubmitCdarApiV1CdarSubmitPost) GetPdpCredentials() PDPCredentials`

GetPdpCredentials returns the PdpCredentials field if non-nil, zero value otherwise.

### GetPdpCredentialsOk

`func (o *BodySubmitCdarApiV1CdarSubmitPost) GetPdpCredentialsOk() (*PDPCredentials, bool)`

GetPdpCredentialsOk returns a tuple with the PdpCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpCredentials

`func (o *BodySubmitCdarApiV1CdarSubmitPost) SetPdpCredentials(v PDPCredentials)`

SetPdpCredentials sets PdpCredentials field to given value.

### HasPdpCredentials

`func (o *BodySubmitCdarApiV1CdarSubmitPost) HasPdpCredentials() bool`

HasPdpCredentials returns a boolean if a field has been set.

### SetPdpCredentialsNil

`func (o *BodySubmitCdarApiV1CdarSubmitPost) SetPdpCredentialsNil(b bool)`

 SetPdpCredentialsNil sets the value for PdpCredentials to be an explicit nil

### UnsetPdpCredentials
`func (o *BodySubmitCdarApiV1CdarSubmitPost) UnsetPdpCredentials()`

UnsetPdpCredentials ensures that no value is present for PdpCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



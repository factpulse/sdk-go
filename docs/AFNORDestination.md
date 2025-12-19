# AFNORDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "afnor"]
**Credentials** | Pointer to [**NullableAFNORCredentials**](AFNORCredentials.md) |  | [optional] 
**FlowSyntax** | Pointer to **string** | Flow syntax to send | [optional] [default to "Factur-X"]
**TrackingId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAFNORDestination

`func NewAFNORDestination() *AFNORDestination`

NewAFNORDestination instantiates a new AFNORDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORDestinationWithDefaults

`func NewAFNORDestinationWithDefaults() *AFNORDestination`

NewAFNORDestinationWithDefaults instantiates a new AFNORDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AFNORDestination) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AFNORDestination) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AFNORDestination) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AFNORDestination) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCredentials

`func (o *AFNORDestination) GetCredentials() AFNORCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *AFNORDestination) GetCredentialsOk() (*AFNORCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *AFNORDestination) SetCredentials(v AFNORCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *AFNORDestination) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *AFNORDestination) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *AFNORDestination) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetFlowSyntax

`func (o *AFNORDestination) GetFlowSyntax() string`

GetFlowSyntax returns the FlowSyntax field if non-nil, zero value otherwise.

### GetFlowSyntaxOk

`func (o *AFNORDestination) GetFlowSyntaxOk() (*string, bool)`

GetFlowSyntaxOk returns a tuple with the FlowSyntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSyntax

`func (o *AFNORDestination) SetFlowSyntax(v string)`

SetFlowSyntax sets FlowSyntax field to given value.

### HasFlowSyntax

`func (o *AFNORDestination) HasFlowSyntax() bool`

HasFlowSyntax returns a boolean if a field has been set.

### GetTrackingId

`func (o *AFNORDestination) GetTrackingId() string`

GetTrackingId returns the TrackingId field if non-nil, zero value otherwise.

### GetTrackingIdOk

`func (o *AFNORDestination) GetTrackingIdOk() (*string, bool)`

GetTrackingIdOk returns a tuple with the TrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingId

`func (o *AFNORDestination) SetTrackingId(v string)`

SetTrackingId sets TrackingId field to given value.

### HasTrackingId

`func (o *AFNORDestination) HasTrackingId() bool`

HasTrackingId returns a boolean if a field has been set.

### SetTrackingIdNil

`func (o *AFNORDestination) SetTrackingIdNil(b bool)`

 SetTrackingIdNil sets the value for TrackingId to be an explicit nil

### UnsetTrackingId
`func (o *AFNORDestination) UnsetTrackingId()`

UnsetTrackingId ensures that no value is present for TrackingId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DestinationAFNOR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "afnor"]
**Credentials** | Pointer to [**NullableCredentialsAFNOR**](CredentialsAFNOR.md) |  | [optional] 
**FlowSyntax** | Pointer to **string** | Syntaxe du flux à envoyer | [optional] [default to "Factur-X"]
**TrackingId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDestinationAFNOR

`func NewDestinationAFNOR() *DestinationAFNOR`

NewDestinationAFNOR instantiates a new DestinationAFNOR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationAFNORWithDefaults

`func NewDestinationAFNORWithDefaults() *DestinationAFNOR`

NewDestinationAFNORWithDefaults instantiates a new DestinationAFNOR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DestinationAFNOR) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationAFNOR) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationAFNOR) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DestinationAFNOR) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCredentials

`func (o *DestinationAFNOR) GetCredentials() CredentialsAFNOR`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *DestinationAFNOR) GetCredentialsOk() (*CredentialsAFNOR, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *DestinationAFNOR) SetCredentials(v CredentialsAFNOR)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *DestinationAFNOR) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *DestinationAFNOR) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *DestinationAFNOR) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetFlowSyntax

`func (o *DestinationAFNOR) GetFlowSyntax() string`

GetFlowSyntax returns the FlowSyntax field if non-nil, zero value otherwise.

### GetFlowSyntaxOk

`func (o *DestinationAFNOR) GetFlowSyntaxOk() (*string, bool)`

GetFlowSyntaxOk returns a tuple with the FlowSyntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSyntax

`func (o *DestinationAFNOR) SetFlowSyntax(v string)`

SetFlowSyntax sets FlowSyntax field to given value.

### HasFlowSyntax

`func (o *DestinationAFNOR) HasFlowSyntax() bool`

HasFlowSyntax returns a boolean if a field has been set.

### GetTrackingId

`func (o *DestinationAFNOR) GetTrackingId() string`

GetTrackingId returns the TrackingId field if non-nil, zero value otherwise.

### GetTrackingIdOk

`func (o *DestinationAFNOR) GetTrackingIdOk() (*string, bool)`

GetTrackingIdOk returns a tuple with the TrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingId

`func (o *DestinationAFNOR) SetTrackingId(v string)`

SetTrackingId sets TrackingId field to given value.

### HasTrackingId

`func (o *DestinationAFNOR) HasTrackingId() bool`

HasTrackingId returns a boolean if a field has been set.

### SetTrackingIdNil

`func (o *DestinationAFNOR) SetTrackingIdNil(b bool)`

 SetTrackingIdNil sets the value for TrackingId to be an explicit nil

### UnsetTrackingId
`func (o *DestinationAFNOR) UnsetTrackingId()`

UnsetTrackingId ensures that no value is present for TrackingId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



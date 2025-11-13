# Destination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "chorus_pro"]
**Credentials** | Pointer to [**CredentialsAFNOR**](CredentialsAFNOR.md) |  | [optional] 
**FlowSyntax** | Pointer to **string** | Syntaxe du flux à envoyer | [optional] [default to "Factur-X"]
**TrackingId** | Pointer to **string** |  | [optional] 

## Methods

### NewDestination

`func NewDestination() *Destination`

NewDestination instantiates a new Destination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationWithDefaults

`func NewDestinationWithDefaults() *Destination`

NewDestinationWithDefaults instantiates a new Destination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Destination) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Destination) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Destination) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Destination) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCredentials

`func (o *Destination) GetCredentials() CredentialsAFNOR`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *Destination) GetCredentialsOk() (*CredentialsAFNOR, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *Destination) SetCredentials(v CredentialsAFNOR)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *Destination) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetFlowSyntax

`func (o *Destination) GetFlowSyntax() string`

GetFlowSyntax returns the FlowSyntax field if non-nil, zero value otherwise.

### GetFlowSyntaxOk

`func (o *Destination) GetFlowSyntaxOk() (*string, bool)`

GetFlowSyntaxOk returns a tuple with the FlowSyntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSyntax

`func (o *Destination) SetFlowSyntax(v string)`

SetFlowSyntax sets FlowSyntax field to given value.

### HasFlowSyntax

`func (o *Destination) HasFlowSyntax() bool`

HasFlowSyntax returns a boolean if a field has been set.

### GetTrackingId

`func (o *Destination) GetTrackingId() string`

GetTrackingId returns the TrackingId field if non-nil, zero value otherwise.

### GetTrackingIdOk

`func (o *Destination) GetTrackingIdOk() (*string, bool)`

GetTrackingIdOk returns a tuple with the TrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingId

`func (o *Destination) SetTrackingId(v string)`

SetTrackingId sets TrackingId field to given value.

### HasTrackingId

`func (o *Destination) HasTrackingId() bool`

HasTrackingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



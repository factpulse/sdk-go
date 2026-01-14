# Destination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "chorus_pro"]
**Credentials** | Pointer to [**AFNORCredentials**](AFNORCredentials.md) |  | [optional] 
**FlowSyntax** | Pointer to [**FlowSyntax**](FlowSyntax.md) | Flow syntax (AFNOR XP Z12-013) | [optional] [default to FACTUR_X]
**TrackingId** | Pointer to **string** |  | [optional] 
**ProcessingRule** | Pointer to [**ProcessingRule**](ProcessingRule.md) |  | [optional] 

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

`func (o *Destination) GetCredentials() AFNORCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *Destination) GetCredentialsOk() (*AFNORCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *Destination) SetCredentials(v AFNORCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *Destination) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetFlowSyntax

`func (o *Destination) GetFlowSyntax() FlowSyntax`

GetFlowSyntax returns the FlowSyntax field if non-nil, zero value otherwise.

### GetFlowSyntaxOk

`func (o *Destination) GetFlowSyntaxOk() (*FlowSyntax, bool)`

GetFlowSyntaxOk returns a tuple with the FlowSyntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSyntax

`func (o *Destination) SetFlowSyntax(v FlowSyntax)`

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

### GetProcessingRule

`func (o *Destination) GetProcessingRule() ProcessingRule`

GetProcessingRule returns the ProcessingRule field if non-nil, zero value otherwise.

### GetProcessingRuleOk

`func (o *Destination) GetProcessingRuleOk() (*ProcessingRule, bool)`

GetProcessingRuleOk returns a tuple with the ProcessingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingRule

`func (o *Destination) SetProcessingRule(v ProcessingRule)`

SetProcessingRule sets ProcessingRule field to given value.

### HasProcessingRule

`func (o *Destination) HasProcessingRule() bool`

HasProcessingRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



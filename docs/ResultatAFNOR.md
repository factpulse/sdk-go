# ResultatAFNOR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowId** | **string** | Identifiant du flux soumis | 
**TrackingId** | Pointer to **NullableString** |  | [optional] 
**Sha256** | **string** | Hash SHA-256 du fichier soumis | 
**FlowSyntax** | **string** | Syntaxe du flux | 
**FlowProfile** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewResultatAFNOR

`func NewResultatAFNOR(flowId string, sha256 string, flowSyntax string, ) *ResultatAFNOR`

NewResultatAFNOR instantiates a new ResultatAFNOR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultatAFNORWithDefaults

`func NewResultatAFNORWithDefaults() *ResultatAFNOR`

NewResultatAFNORWithDefaults instantiates a new ResultatAFNOR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowId

`func (o *ResultatAFNOR) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *ResultatAFNOR) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *ResultatAFNOR) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.


### GetTrackingId

`func (o *ResultatAFNOR) GetTrackingId() string`

GetTrackingId returns the TrackingId field if non-nil, zero value otherwise.

### GetTrackingIdOk

`func (o *ResultatAFNOR) GetTrackingIdOk() (*string, bool)`

GetTrackingIdOk returns a tuple with the TrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingId

`func (o *ResultatAFNOR) SetTrackingId(v string)`

SetTrackingId sets TrackingId field to given value.

### HasTrackingId

`func (o *ResultatAFNOR) HasTrackingId() bool`

HasTrackingId returns a boolean if a field has been set.

### SetTrackingIdNil

`func (o *ResultatAFNOR) SetTrackingIdNil(b bool)`

 SetTrackingIdNil sets the value for TrackingId to be an explicit nil

### UnsetTrackingId
`func (o *ResultatAFNOR) UnsetTrackingId()`

UnsetTrackingId ensures that no value is present for TrackingId, not even an explicit nil
### GetSha256

`func (o *ResultatAFNOR) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ResultatAFNOR) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ResultatAFNOR) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.


### GetFlowSyntax

`func (o *ResultatAFNOR) GetFlowSyntax() string`

GetFlowSyntax returns the FlowSyntax field if non-nil, zero value otherwise.

### GetFlowSyntaxOk

`func (o *ResultatAFNOR) GetFlowSyntaxOk() (*string, bool)`

GetFlowSyntaxOk returns a tuple with the FlowSyntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSyntax

`func (o *ResultatAFNOR) SetFlowSyntax(v string)`

SetFlowSyntax sets FlowSyntax field to given value.


### GetFlowProfile

`func (o *ResultatAFNOR) GetFlowProfile() string`

GetFlowProfile returns the FlowProfile field if non-nil, zero value otherwise.

### GetFlowProfileOk

`func (o *ResultatAFNOR) GetFlowProfileOk() (*string, bool)`

GetFlowProfileOk returns a tuple with the FlowProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowProfile

`func (o *ResultatAFNOR) SetFlowProfile(v string)`

SetFlowProfile sets FlowProfile field to given value.

### HasFlowProfile

`func (o *ResultatAFNOR) HasFlowProfile() bool`

HasFlowProfile returns a boolean if a field has been set.

### SetFlowProfileNil

`func (o *ResultatAFNOR) SetFlowProfileNil(b bool)`

 SetFlowProfileNil sets the value for FlowProfile to be an explicit nil

### UnsetFlowProfile
`func (o *ResultatAFNOR) UnsetFlowProfile()`

UnsetFlowProfile ensures that no value is present for FlowProfile, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



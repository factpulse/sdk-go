# AFNORFlowInfoExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha256** | Pointer to **string** | The sha256 is the fingerprint of the attached file: - if provided in the request: it should be checked once received - if not provided in the request: it may be computed and returned in the response  | [optional] 

## Methods

### NewAFNORFlowInfoExtension

`func NewAFNORFlowInfoExtension() *AFNORFlowInfoExtension`

NewAFNORFlowInfoExtension instantiates a new AFNORFlowInfoExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORFlowInfoExtensionWithDefaults

`func NewAFNORFlowInfoExtensionWithDefaults() *AFNORFlowInfoExtension`

NewAFNORFlowInfoExtensionWithDefaults instantiates a new AFNORFlowInfoExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha256

`func (o *AFNORFlowInfoExtension) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *AFNORFlowInfoExtension) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *AFNORFlowInfoExtension) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *AFNORFlowInfoExtension) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



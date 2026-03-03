# AFNORCallbackSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algo** | [**AFNORAlgorithm**](AFNORAlgorithm.md) |  | 
**Key** | **string** | Base 64 encoded string | 

## Methods

### NewAFNORCallbackSignature

`func NewAFNORCallbackSignature(algo AFNORAlgorithm, key string, ) *AFNORCallbackSignature`

NewAFNORCallbackSignature instantiates a new AFNORCallbackSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORCallbackSignatureWithDefaults

`func NewAFNORCallbackSignatureWithDefaults() *AFNORCallbackSignature`

NewAFNORCallbackSignatureWithDefaults instantiates a new AFNORCallbackSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgo

`func (o *AFNORCallbackSignature) GetAlgo() AFNORAlgorithm`

GetAlgo returns the Algo field if non-nil, zero value otherwise.

### GetAlgoOk

`func (o *AFNORCallbackSignature) GetAlgoOk() (*AFNORAlgorithm, bool)`

GetAlgoOk returns a tuple with the Algo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgo

`func (o *AFNORCallbackSignature) SetAlgo(v AFNORAlgorithm)`

SetAlgo sets Algo field to given value.


### GetKey

`func (o *AFNORCallbackSignature) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AFNORCallbackSignature) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AFNORCallbackSignature) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



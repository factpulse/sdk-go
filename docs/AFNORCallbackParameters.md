# AFNORCallbackParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Headers** | Pointer to [**[]AFNORCallbackHeader**](AFNORCallbackHeader.md) |  | [optional] 
**Authentication** | Pointer to [**AFNORCallbackAuthentication**](AFNORCallbackAuthentication.md) |  | [optional] 
**Signature** | Pointer to [**AFNORCallbackSignature**](AFNORCallbackSignature.md) |  | [optional] 

## Methods

### NewAFNORCallbackParameters

`func NewAFNORCallbackParameters(url string, ) *AFNORCallbackParameters`

NewAFNORCallbackParameters instantiates a new AFNORCallbackParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAFNORCallbackParametersWithDefaults

`func NewAFNORCallbackParametersWithDefaults() *AFNORCallbackParameters`

NewAFNORCallbackParametersWithDefaults instantiates a new AFNORCallbackParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *AFNORCallbackParameters) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AFNORCallbackParameters) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AFNORCallbackParameters) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaders

`func (o *AFNORCallbackParameters) GetHeaders() []AFNORCallbackHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *AFNORCallbackParameters) GetHeadersOk() (*[]AFNORCallbackHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *AFNORCallbackParameters) SetHeaders(v []AFNORCallbackHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *AFNORCallbackParameters) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetAuthentication

`func (o *AFNORCallbackParameters) GetAuthentication() AFNORCallbackAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *AFNORCallbackParameters) GetAuthenticationOk() (*AFNORCallbackAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *AFNORCallbackParameters) SetAuthentication(v AFNORCallbackAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *AFNORCallbackParameters) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetSignature

`func (o *AFNORCallbackParameters) GetSignature() AFNORCallbackSignature`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *AFNORCallbackParameters) GetSignatureOk() (*AFNORCallbackSignature, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *AFNORCallbackParameters) SetSignature(v AFNORCallbackSignature)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *AFNORCallbackParameters) HasSignature() bool`

HasSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



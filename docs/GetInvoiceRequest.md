# GetInvoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**NullableChorusProCredentials**](ChorusProCredentials.md) |  | [optional] 
**ChorusInvoiceId** | **int32** | Chorus Pro invoice ID | 

## Methods

### NewGetInvoiceRequest

`func NewGetInvoiceRequest(chorusInvoiceId int32, ) *GetInvoiceRequest`

NewGetInvoiceRequest instantiates a new GetInvoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvoiceRequestWithDefaults

`func NewGetInvoiceRequestWithDefaults() *GetInvoiceRequest`

NewGetInvoiceRequestWithDefaults instantiates a new GetInvoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *GetInvoiceRequest) GetCredentials() ChorusProCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GetInvoiceRequest) GetCredentialsOk() (*ChorusProCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GetInvoiceRequest) SetCredentials(v ChorusProCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *GetInvoiceRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *GetInvoiceRequest) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *GetInvoiceRequest) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetChorusInvoiceId

`func (o *GetInvoiceRequest) GetChorusInvoiceId() int32`

GetChorusInvoiceId returns the ChorusInvoiceId field if non-nil, zero value otherwise.

### GetChorusInvoiceIdOk

`func (o *GetInvoiceRequest) GetChorusInvoiceIdOk() (*int32, bool)`

GetChorusInvoiceIdOk returns a tuple with the ChorusInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChorusInvoiceId

`func (o *GetInvoiceRequest) SetChorusInvoiceId(v int32)`

SetChorusInvoiceId sets ChorusInvoiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PaymentCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardId** | Pointer to **NullableString** |  | [optional] 
**CardholderName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaymentCard

`func NewPaymentCard() *PaymentCard`

NewPaymentCard instantiates a new PaymentCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentCardWithDefaults

`func NewPaymentCardWithDefaults() *PaymentCard`

NewPaymentCardWithDefaults instantiates a new PaymentCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardId

`func (o *PaymentCard) GetCardId() string`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *PaymentCard) GetCardIdOk() (*string, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *PaymentCard) SetCardId(v string)`

SetCardId sets CardId field to given value.

### HasCardId

`func (o *PaymentCard) HasCardId() bool`

HasCardId returns a boolean if a field has been set.

### SetCardIdNil

`func (o *PaymentCard) SetCardIdNil(b bool)`

 SetCardIdNil sets the value for CardId to be an explicit nil

### UnsetCardId
`func (o *PaymentCard) UnsetCardId()`

UnsetCardId ensures that no value is present for CardId, not even an explicit nil
### GetCardholderName

`func (o *PaymentCard) GetCardholderName() string`

GetCardholderName returns the CardholderName field if non-nil, zero value otherwise.

### GetCardholderNameOk

`func (o *PaymentCard) GetCardholderNameOk() (*string, bool)`

GetCardholderNameOk returns a tuple with the CardholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderName

`func (o *PaymentCard) SetCardholderName(v string)`

SetCardholderName sets CardholderName field to given value.

### HasCardholderName

`func (o *PaymentCard) HasCardholderName() bool`

HasCardholderName returns a boolean if a field has been set.

### SetCardholderNameNil

`func (o *PaymentCard) SetCardholderNameNil(b bool)`

 SetCardholderNameNil sets the value for CardholderName to be an explicit nil

### UnsetCardholderName
`func (o *PaymentCard) UnsetCardholderName()`

UnsetCardholderName ensures that no value is present for CardholderName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



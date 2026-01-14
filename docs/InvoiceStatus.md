# InvoiceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Status code (SOUMISE, VALIDEE, REJETEE, SUSPENDUE, MANDATEE, MISE_EN_PAIEMENT, etc.) | 
**Label** | **string** | Status label | 
**Date** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInvoiceStatus

`func NewInvoiceStatus(code string, label string, ) *InvoiceStatus`

NewInvoiceStatus instantiates a new InvoiceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceStatusWithDefaults

`func NewInvoiceStatusWithDefaults() *InvoiceStatus`

NewInvoiceStatusWithDefaults instantiates a new InvoiceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *InvoiceStatus) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InvoiceStatus) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InvoiceStatus) SetCode(v string)`

SetCode sets Code field to given value.


### GetLabel

`func (o *InvoiceStatus) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InvoiceStatus) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InvoiceStatus) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDate

`func (o *InvoiceStatus) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *InvoiceStatus) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *InvoiceStatus) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *InvoiceStatus) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *InvoiceStatus) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *InvoiceStatus) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetInvoiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnCode** | **int32** |  | 
**Message** | **string** |  | 
**ChorusInvoiceId** | Pointer to **NullableInt32** |  | [optional] 
**InvoiceNumber** | Pointer to **NullableString** |  | [optional] 
**InvoiceDate** | Pointer to **NullableString** |  | [optional] 
**TotalGrossAmount** | Pointer to **NullableString** |  | [optional] 
**CurrentStatus** | Pointer to [**NullableInvoiceStatus**](InvoiceStatus.md) |  | [optional] 
**RecipientStructureId** | Pointer to **NullableInt32** |  | [optional] 
**RecipientStructureName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetInvoiceResponse

`func NewGetInvoiceResponse(returnCode int32, message string, ) *GetInvoiceResponse`

NewGetInvoiceResponse instantiates a new GetInvoiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvoiceResponseWithDefaults

`func NewGetInvoiceResponseWithDefaults() *GetInvoiceResponse`

NewGetInvoiceResponseWithDefaults instantiates a new GetInvoiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnCode

`func (o *GetInvoiceResponse) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *GetInvoiceResponse) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *GetInvoiceResponse) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.


### GetMessage

`func (o *GetInvoiceResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetInvoiceResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetInvoiceResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetChorusInvoiceId

`func (o *GetInvoiceResponse) GetChorusInvoiceId() int32`

GetChorusInvoiceId returns the ChorusInvoiceId field if non-nil, zero value otherwise.

### GetChorusInvoiceIdOk

`func (o *GetInvoiceResponse) GetChorusInvoiceIdOk() (*int32, bool)`

GetChorusInvoiceIdOk returns a tuple with the ChorusInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChorusInvoiceId

`func (o *GetInvoiceResponse) SetChorusInvoiceId(v int32)`

SetChorusInvoiceId sets ChorusInvoiceId field to given value.

### HasChorusInvoiceId

`func (o *GetInvoiceResponse) HasChorusInvoiceId() bool`

HasChorusInvoiceId returns a boolean if a field has been set.

### SetChorusInvoiceIdNil

`func (o *GetInvoiceResponse) SetChorusInvoiceIdNil(b bool)`

 SetChorusInvoiceIdNil sets the value for ChorusInvoiceId to be an explicit nil

### UnsetChorusInvoiceId
`func (o *GetInvoiceResponse) UnsetChorusInvoiceId()`

UnsetChorusInvoiceId ensures that no value is present for ChorusInvoiceId, not even an explicit nil
### GetInvoiceNumber

`func (o *GetInvoiceResponse) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *GetInvoiceResponse) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *GetInvoiceResponse) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *GetInvoiceResponse) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### SetInvoiceNumberNil

`func (o *GetInvoiceResponse) SetInvoiceNumberNil(b bool)`

 SetInvoiceNumberNil sets the value for InvoiceNumber to be an explicit nil

### UnsetInvoiceNumber
`func (o *GetInvoiceResponse) UnsetInvoiceNumber()`

UnsetInvoiceNumber ensures that no value is present for InvoiceNumber, not even an explicit nil
### GetInvoiceDate

`func (o *GetInvoiceResponse) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *GetInvoiceResponse) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *GetInvoiceResponse) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *GetInvoiceResponse) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### SetInvoiceDateNil

`func (o *GetInvoiceResponse) SetInvoiceDateNil(b bool)`

 SetInvoiceDateNil sets the value for InvoiceDate to be an explicit nil

### UnsetInvoiceDate
`func (o *GetInvoiceResponse) UnsetInvoiceDate()`

UnsetInvoiceDate ensures that no value is present for InvoiceDate, not even an explicit nil
### GetTotalGrossAmount

`func (o *GetInvoiceResponse) GetTotalGrossAmount() string`

GetTotalGrossAmount returns the TotalGrossAmount field if non-nil, zero value otherwise.

### GetTotalGrossAmountOk

`func (o *GetInvoiceResponse) GetTotalGrossAmountOk() (*string, bool)`

GetTotalGrossAmountOk returns a tuple with the TotalGrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossAmount

`func (o *GetInvoiceResponse) SetTotalGrossAmount(v string)`

SetTotalGrossAmount sets TotalGrossAmount field to given value.

### HasTotalGrossAmount

`func (o *GetInvoiceResponse) HasTotalGrossAmount() bool`

HasTotalGrossAmount returns a boolean if a field has been set.

### SetTotalGrossAmountNil

`func (o *GetInvoiceResponse) SetTotalGrossAmountNil(b bool)`

 SetTotalGrossAmountNil sets the value for TotalGrossAmount to be an explicit nil

### UnsetTotalGrossAmount
`func (o *GetInvoiceResponse) UnsetTotalGrossAmount()`

UnsetTotalGrossAmount ensures that no value is present for TotalGrossAmount, not even an explicit nil
### GetCurrentStatus

`func (o *GetInvoiceResponse) GetCurrentStatus() InvoiceStatus`

GetCurrentStatus returns the CurrentStatus field if non-nil, zero value otherwise.

### GetCurrentStatusOk

`func (o *GetInvoiceResponse) GetCurrentStatusOk() (*InvoiceStatus, bool)`

GetCurrentStatusOk returns a tuple with the CurrentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStatus

`func (o *GetInvoiceResponse) SetCurrentStatus(v InvoiceStatus)`

SetCurrentStatus sets CurrentStatus field to given value.

### HasCurrentStatus

`func (o *GetInvoiceResponse) HasCurrentStatus() bool`

HasCurrentStatus returns a boolean if a field has been set.

### SetCurrentStatusNil

`func (o *GetInvoiceResponse) SetCurrentStatusNil(b bool)`

 SetCurrentStatusNil sets the value for CurrentStatus to be an explicit nil

### UnsetCurrentStatus
`func (o *GetInvoiceResponse) UnsetCurrentStatus()`

UnsetCurrentStatus ensures that no value is present for CurrentStatus, not even an explicit nil
### GetRecipientStructureId

`func (o *GetInvoiceResponse) GetRecipientStructureId() int32`

GetRecipientStructureId returns the RecipientStructureId field if non-nil, zero value otherwise.

### GetRecipientStructureIdOk

`func (o *GetInvoiceResponse) GetRecipientStructureIdOk() (*int32, bool)`

GetRecipientStructureIdOk returns a tuple with the RecipientStructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientStructureId

`func (o *GetInvoiceResponse) SetRecipientStructureId(v int32)`

SetRecipientStructureId sets RecipientStructureId field to given value.

### HasRecipientStructureId

`func (o *GetInvoiceResponse) HasRecipientStructureId() bool`

HasRecipientStructureId returns a boolean if a field has been set.

### SetRecipientStructureIdNil

`func (o *GetInvoiceResponse) SetRecipientStructureIdNil(b bool)`

 SetRecipientStructureIdNil sets the value for RecipientStructureId to be an explicit nil

### UnsetRecipientStructureId
`func (o *GetInvoiceResponse) UnsetRecipientStructureId()`

UnsetRecipientStructureId ensures that no value is present for RecipientStructureId, not even an explicit nil
### GetRecipientStructureName

`func (o *GetInvoiceResponse) GetRecipientStructureName() string`

GetRecipientStructureName returns the RecipientStructureName field if non-nil, zero value otherwise.

### GetRecipientStructureNameOk

`func (o *GetInvoiceResponse) GetRecipientStructureNameOk() (*string, bool)`

GetRecipientStructureNameOk returns a tuple with the RecipientStructureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientStructureName

`func (o *GetInvoiceResponse) SetRecipientStructureName(v string)`

SetRecipientStructureName sets RecipientStructureName field to given value.

### HasRecipientStructureName

`func (o *GetInvoiceResponse) HasRecipientStructureName() bool`

HasRecipientStructureName returns a boolean if a field has been set.

### SetRecipientStructureNameNil

`func (o *GetInvoiceResponse) SetRecipientStructureNameNil(b bool)`

 SetRecipientStructureNameNil sets the value for RecipientStructureName to be an explicit nil

### UnsetRecipientStructureName
`func (o *GetInvoiceResponse) UnsetRecipientStructureName()`

UnsetRecipientStructureName ensures that no value is present for RecipientStructureName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



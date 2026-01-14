# InvoiceNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectCode** | Pointer to **NullableString** |  | [optional] 
**Content** | **string** | Note content | 

## Methods

### NewInvoiceNote

`func NewInvoiceNote(content string, ) *InvoiceNote`

NewInvoiceNote instantiates a new InvoiceNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceNoteWithDefaults

`func NewInvoiceNoteWithDefaults() *InvoiceNote`

NewInvoiceNoteWithDefaults instantiates a new InvoiceNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectCode

`func (o *InvoiceNote) GetSubjectCode() string`

GetSubjectCode returns the SubjectCode field if non-nil, zero value otherwise.

### GetSubjectCodeOk

`func (o *InvoiceNote) GetSubjectCodeOk() (*string, bool)`

GetSubjectCodeOk returns a tuple with the SubjectCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCode

`func (o *InvoiceNote) SetSubjectCode(v string)`

SetSubjectCode sets SubjectCode field to given value.

### HasSubjectCode

`func (o *InvoiceNote) HasSubjectCode() bool`

HasSubjectCode returns a boolean if a field has been set.

### SetSubjectCodeNil

`func (o *InvoiceNote) SetSubjectCodeNil(b bool)`

 SetSubjectCodeNil sets the value for SubjectCode to be an explicit nil

### UnsetSubjectCode
`func (o *InvoiceNote) UnsetSubjectCode()`

UnsetSubjectCode ensures that no value is present for SubjectCode, not even an explicit nil
### GetContent

`func (o *InvoiceNote) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *InvoiceNote) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *InvoiceNote) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



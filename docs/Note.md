# Note

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectCode** | Pointer to **NullableString** |  | [optional] 
**Content** | **string** |  | 

## Methods

### NewNote

`func NewNote(content string, ) *Note`

NewNote instantiates a new Note object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteWithDefaults

`func NewNoteWithDefaults() *Note`

NewNoteWithDefaults instantiates a new Note object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectCode

`func (o *Note) GetSubjectCode() string`

GetSubjectCode returns the SubjectCode field if non-nil, zero value otherwise.

### GetSubjectCodeOk

`func (o *Note) GetSubjectCodeOk() (*string, bool)`

GetSubjectCodeOk returns a tuple with the SubjectCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCode

`func (o *Note) SetSubjectCode(v string)`

SetSubjectCode sets SubjectCode field to given value.

### HasSubjectCode

`func (o *Note) HasSubjectCode() bool`

HasSubjectCode returns a boolean if a field has been set.

### SetSubjectCodeNil

`func (o *Note) SetSubjectCodeNil(b bool)`

 SetSubjectCodeNil sets the value for SubjectCode to be an explicit nil

### UnsetSubjectCode
`func (o *Note) UnsetSubjectCode()`

UnsetSubjectCode ensures that no value is present for SubjectCode, not even an explicit nil
### GetContent

`func (o *Note) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Note) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Note) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



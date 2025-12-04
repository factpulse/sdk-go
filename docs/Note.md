# Note

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodeSujet** | Pointer to **NullableString** |  | [optional] 
**Contenu** | **string** |  | 

## Methods

### NewNote

`func NewNote(contenu string, ) *Note`

NewNote instantiates a new Note object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteWithDefaults

`func NewNoteWithDefaults() *Note`

NewNoteWithDefaults instantiates a new Note object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodeSujet

`func (o *Note) GetCodeSujet() string`

GetCodeSujet returns the CodeSujet field if non-nil, zero value otherwise.

### GetCodeSujetOk

`func (o *Note) GetCodeSujetOk() (*string, bool)`

GetCodeSujetOk returns a tuple with the CodeSujet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeSujet

`func (o *Note) SetCodeSujet(v string)`

SetCodeSujet sets CodeSujet field to given value.

### HasCodeSujet

`func (o *Note) HasCodeSujet() bool`

HasCodeSujet returns a boolean if a field has been set.

### SetCodeSujetNil

`func (o *Note) SetCodeSujetNil(b bool)`

 SetCodeSujetNil sets the value for CodeSujet to be an explicit nil

### UnsetCodeSujet
`func (o *Note) UnsetCodeSujet()`

UnsetCodeSujet ensures that no value is present for CodeSujet, not even an explicit nil
### GetContenu

`func (o *Note) GetContenu() string`

GetContenu returns the Contenu field if non-nil, zero value otherwise.

### GetContenuOk

`func (o *Note) GetContenuOk() (*string, bool)`

GetContenuOk returns a tuple with the Contenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContenu

`func (o *Note) SetContenu(v string)`

SetContenu sets Contenu field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



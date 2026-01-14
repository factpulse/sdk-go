# SearchStructureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnCode** | **int32** | Return code (0 &#x3D; success) | 
**Message** | **string** | Return message | 
**Structures** | Pointer to [**[]StructureInfo**](StructureInfo.md) |  | [optional] 
**Total** | Pointer to **int32** | Total number of results | [optional] [default to 0]

## Methods

### NewSearchStructureResponse

`func NewSearchStructureResponse(returnCode int32, message string, ) *SearchStructureResponse`

NewSearchStructureResponse instantiates a new SearchStructureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchStructureResponseWithDefaults

`func NewSearchStructureResponseWithDefaults() *SearchStructureResponse`

NewSearchStructureResponseWithDefaults instantiates a new SearchStructureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnCode

`func (o *SearchStructureResponse) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *SearchStructureResponse) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *SearchStructureResponse) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.


### GetMessage

`func (o *SearchStructureResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SearchStructureResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SearchStructureResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStructures

`func (o *SearchStructureResponse) GetStructures() []StructureInfo`

GetStructures returns the Structures field if non-nil, zero value otherwise.

### GetStructuresOk

`func (o *SearchStructureResponse) GetStructuresOk() (*[]StructureInfo, bool)`

GetStructuresOk returns a tuple with the Structures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructures

`func (o *SearchStructureResponse) SetStructures(v []StructureInfo)`

SetStructures sets Structures field to given value.

### HasStructures

`func (o *SearchStructureResponse) HasStructures() bool`

HasStructures returns a boolean if a field has been set.

### GetTotal

`func (o *SearchStructureResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SearchStructureResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SearchStructureResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SearchStructureResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



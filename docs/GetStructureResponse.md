# GetStructureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnCode** | **int32** |  | 
**Message** | **string** |  | 
**StructureId** | Pointer to **NullableInt32** |  | [optional] 
**StructureIdentifier** | Pointer to **NullableString** |  | [optional] 
**StructureLabel** | Pointer to **NullableString** |  | [optional] 
**CompanyName** | Pointer to **NullableString** |  | [optional] 
**VatNumber** | Pointer to **NullableString** |  | [optional] 
**StructureEmail** | Pointer to **NullableString** |  | [optional] 
**Parameters** | Pointer to [**NullableStructureParameters**](StructureParameters.md) |  | [optional] 

## Methods

### NewGetStructureResponse

`func NewGetStructureResponse(returnCode int32, message string, ) *GetStructureResponse`

NewGetStructureResponse instantiates a new GetStructureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStructureResponseWithDefaults

`func NewGetStructureResponseWithDefaults() *GetStructureResponse`

NewGetStructureResponseWithDefaults instantiates a new GetStructureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnCode

`func (o *GetStructureResponse) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *GetStructureResponse) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *GetStructureResponse) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.


### GetMessage

`func (o *GetStructureResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetStructureResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetStructureResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStructureId

`func (o *GetStructureResponse) GetStructureId() int32`

GetStructureId returns the StructureId field if non-nil, zero value otherwise.

### GetStructureIdOk

`func (o *GetStructureResponse) GetStructureIdOk() (*int32, bool)`

GetStructureIdOk returns a tuple with the StructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureId

`func (o *GetStructureResponse) SetStructureId(v int32)`

SetStructureId sets StructureId field to given value.

### HasStructureId

`func (o *GetStructureResponse) HasStructureId() bool`

HasStructureId returns a boolean if a field has been set.

### SetStructureIdNil

`func (o *GetStructureResponse) SetStructureIdNil(b bool)`

 SetStructureIdNil sets the value for StructureId to be an explicit nil

### UnsetStructureId
`func (o *GetStructureResponse) UnsetStructureId()`

UnsetStructureId ensures that no value is present for StructureId, not even an explicit nil
### GetStructureIdentifier

`func (o *GetStructureResponse) GetStructureIdentifier() string`

GetStructureIdentifier returns the StructureIdentifier field if non-nil, zero value otherwise.

### GetStructureIdentifierOk

`func (o *GetStructureResponse) GetStructureIdentifierOk() (*string, bool)`

GetStructureIdentifierOk returns a tuple with the StructureIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureIdentifier

`func (o *GetStructureResponse) SetStructureIdentifier(v string)`

SetStructureIdentifier sets StructureIdentifier field to given value.

### HasStructureIdentifier

`func (o *GetStructureResponse) HasStructureIdentifier() bool`

HasStructureIdentifier returns a boolean if a field has been set.

### SetStructureIdentifierNil

`func (o *GetStructureResponse) SetStructureIdentifierNil(b bool)`

 SetStructureIdentifierNil sets the value for StructureIdentifier to be an explicit nil

### UnsetStructureIdentifier
`func (o *GetStructureResponse) UnsetStructureIdentifier()`

UnsetStructureIdentifier ensures that no value is present for StructureIdentifier, not even an explicit nil
### GetStructureLabel

`func (o *GetStructureResponse) GetStructureLabel() string`

GetStructureLabel returns the StructureLabel field if non-nil, zero value otherwise.

### GetStructureLabelOk

`func (o *GetStructureResponse) GetStructureLabelOk() (*string, bool)`

GetStructureLabelOk returns a tuple with the StructureLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureLabel

`func (o *GetStructureResponse) SetStructureLabel(v string)`

SetStructureLabel sets StructureLabel field to given value.

### HasStructureLabel

`func (o *GetStructureResponse) HasStructureLabel() bool`

HasStructureLabel returns a boolean if a field has been set.

### SetStructureLabelNil

`func (o *GetStructureResponse) SetStructureLabelNil(b bool)`

 SetStructureLabelNil sets the value for StructureLabel to be an explicit nil

### UnsetStructureLabel
`func (o *GetStructureResponse) UnsetStructureLabel()`

UnsetStructureLabel ensures that no value is present for StructureLabel, not even an explicit nil
### GetCompanyName

`func (o *GetStructureResponse) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *GetStructureResponse) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *GetStructureResponse) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *GetStructureResponse) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *GetStructureResponse) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *GetStructureResponse) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetVatNumber

`func (o *GetStructureResponse) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *GetStructureResponse) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *GetStructureResponse) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *GetStructureResponse) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### SetVatNumberNil

`func (o *GetStructureResponse) SetVatNumberNil(b bool)`

 SetVatNumberNil sets the value for VatNumber to be an explicit nil

### UnsetVatNumber
`func (o *GetStructureResponse) UnsetVatNumber()`

UnsetVatNumber ensures that no value is present for VatNumber, not even an explicit nil
### GetStructureEmail

`func (o *GetStructureResponse) GetStructureEmail() string`

GetStructureEmail returns the StructureEmail field if non-nil, zero value otherwise.

### GetStructureEmailOk

`func (o *GetStructureResponse) GetStructureEmailOk() (*string, bool)`

GetStructureEmailOk returns a tuple with the StructureEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureEmail

`func (o *GetStructureResponse) SetStructureEmail(v string)`

SetStructureEmail sets StructureEmail field to given value.

### HasStructureEmail

`func (o *GetStructureResponse) HasStructureEmail() bool`

HasStructureEmail returns a boolean if a field has been set.

### SetStructureEmailNil

`func (o *GetStructureResponse) SetStructureEmailNil(b bool)`

 SetStructureEmailNil sets the value for StructureEmail to be an explicit nil

### UnsetStructureEmail
`func (o *GetStructureResponse) UnsetStructureEmail()`

UnsetStructureEmail ensures that no value is present for StructureEmail, not even an explicit nil
### GetParameters

`func (o *GetStructureResponse) GetParameters() StructureParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *GetStructureResponse) GetParametersOk() (*StructureParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *GetStructureResponse) SetParameters(v StructureParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *GetStructureResponse) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *GetStructureResponse) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *GetStructureResponse) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



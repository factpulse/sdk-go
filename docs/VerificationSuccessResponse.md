# VerificationSuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCompliant** | **bool** | True if no critical discrepancy | 
**ComplianceScore** | **float32** | Compliance score (0-100%) | 
**VerifiedFieldsCount** | Pointer to **int32** | Number of verified fields | [optional] [default to 0]
**CompliantFieldsCount** | Pointer to **int32** | Number of compliant fields | [optional] [default to 0]
**IsFacturx** | Pointer to **bool** | True if PDF contains Factur-X XML | [optional] [default to false]
**FacturxProfile** | Pointer to **NullableString** |  | [optional] 
**Fields** | Pointer to [**[]VerifiedFieldSchema**](VerifiedFieldSchema.md) | List of verified fields with values, statuses and PDF coordinates | [optional] 
**MandatoryNotes** | Pointer to [**[]MandatoryNoteSchema**](MandatoryNoteSchema.md) | Mandatory notes (PMT, PMD, AAB) with PDF location | [optional] 
**PageDimensions** | Pointer to [**[]PageDimensionsSchema**](PageDimensionsSchema.md) | Dimensions of each PDF page (width, height) | [optional] 
**Warnings** | Pointer to **[]string** | Non-blocking warnings | [optional] 

## Methods

### NewVerificationSuccessResponse

`func NewVerificationSuccessResponse(isCompliant bool, complianceScore float32, ) *VerificationSuccessResponse`

NewVerificationSuccessResponse instantiates a new VerificationSuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationSuccessResponseWithDefaults

`func NewVerificationSuccessResponseWithDefaults() *VerificationSuccessResponse`

NewVerificationSuccessResponseWithDefaults instantiates a new VerificationSuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCompliant

`func (o *VerificationSuccessResponse) GetIsCompliant() bool`

GetIsCompliant returns the IsCompliant field if non-nil, zero value otherwise.

### GetIsCompliantOk

`func (o *VerificationSuccessResponse) GetIsCompliantOk() (*bool, bool)`

GetIsCompliantOk returns a tuple with the IsCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompliant

`func (o *VerificationSuccessResponse) SetIsCompliant(v bool)`

SetIsCompliant sets IsCompliant field to given value.


### GetComplianceScore

`func (o *VerificationSuccessResponse) GetComplianceScore() float32`

GetComplianceScore returns the ComplianceScore field if non-nil, zero value otherwise.

### GetComplianceScoreOk

`func (o *VerificationSuccessResponse) GetComplianceScoreOk() (*float32, bool)`

GetComplianceScoreOk returns a tuple with the ComplianceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceScore

`func (o *VerificationSuccessResponse) SetComplianceScore(v float32)`

SetComplianceScore sets ComplianceScore field to given value.


### GetVerifiedFieldsCount

`func (o *VerificationSuccessResponse) GetVerifiedFieldsCount() int32`

GetVerifiedFieldsCount returns the VerifiedFieldsCount field if non-nil, zero value otherwise.

### GetVerifiedFieldsCountOk

`func (o *VerificationSuccessResponse) GetVerifiedFieldsCountOk() (*int32, bool)`

GetVerifiedFieldsCountOk returns a tuple with the VerifiedFieldsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedFieldsCount

`func (o *VerificationSuccessResponse) SetVerifiedFieldsCount(v int32)`

SetVerifiedFieldsCount sets VerifiedFieldsCount field to given value.

### HasVerifiedFieldsCount

`func (o *VerificationSuccessResponse) HasVerifiedFieldsCount() bool`

HasVerifiedFieldsCount returns a boolean if a field has been set.

### GetCompliantFieldsCount

`func (o *VerificationSuccessResponse) GetCompliantFieldsCount() int32`

GetCompliantFieldsCount returns the CompliantFieldsCount field if non-nil, zero value otherwise.

### GetCompliantFieldsCountOk

`func (o *VerificationSuccessResponse) GetCompliantFieldsCountOk() (*int32, bool)`

GetCompliantFieldsCountOk returns a tuple with the CompliantFieldsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliantFieldsCount

`func (o *VerificationSuccessResponse) SetCompliantFieldsCount(v int32)`

SetCompliantFieldsCount sets CompliantFieldsCount field to given value.

### HasCompliantFieldsCount

`func (o *VerificationSuccessResponse) HasCompliantFieldsCount() bool`

HasCompliantFieldsCount returns a boolean if a field has been set.

### GetIsFacturx

`func (o *VerificationSuccessResponse) GetIsFacturx() bool`

GetIsFacturx returns the IsFacturx field if non-nil, zero value otherwise.

### GetIsFacturxOk

`func (o *VerificationSuccessResponse) GetIsFacturxOk() (*bool, bool)`

GetIsFacturxOk returns a tuple with the IsFacturx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFacturx

`func (o *VerificationSuccessResponse) SetIsFacturx(v bool)`

SetIsFacturx sets IsFacturx field to given value.

### HasIsFacturx

`func (o *VerificationSuccessResponse) HasIsFacturx() bool`

HasIsFacturx returns a boolean if a field has been set.

### GetFacturxProfile

`func (o *VerificationSuccessResponse) GetFacturxProfile() string`

GetFacturxProfile returns the FacturxProfile field if non-nil, zero value otherwise.

### GetFacturxProfileOk

`func (o *VerificationSuccessResponse) GetFacturxProfileOk() (*string, bool)`

GetFacturxProfileOk returns a tuple with the FacturxProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacturxProfile

`func (o *VerificationSuccessResponse) SetFacturxProfile(v string)`

SetFacturxProfile sets FacturxProfile field to given value.

### HasFacturxProfile

`func (o *VerificationSuccessResponse) HasFacturxProfile() bool`

HasFacturxProfile returns a boolean if a field has been set.

### SetFacturxProfileNil

`func (o *VerificationSuccessResponse) SetFacturxProfileNil(b bool)`

 SetFacturxProfileNil sets the value for FacturxProfile to be an explicit nil

### UnsetFacturxProfile
`func (o *VerificationSuccessResponse) UnsetFacturxProfile()`

UnsetFacturxProfile ensures that no value is present for FacturxProfile, not even an explicit nil
### GetFields

`func (o *VerificationSuccessResponse) GetFields() []VerifiedFieldSchema`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *VerificationSuccessResponse) GetFieldsOk() (*[]VerifiedFieldSchema, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *VerificationSuccessResponse) SetFields(v []VerifiedFieldSchema)`

SetFields sets Fields field to given value.

### HasFields

`func (o *VerificationSuccessResponse) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetMandatoryNotes

`func (o *VerificationSuccessResponse) GetMandatoryNotes() []MandatoryNoteSchema`

GetMandatoryNotes returns the MandatoryNotes field if non-nil, zero value otherwise.

### GetMandatoryNotesOk

`func (o *VerificationSuccessResponse) GetMandatoryNotesOk() (*[]MandatoryNoteSchema, bool)`

GetMandatoryNotesOk returns a tuple with the MandatoryNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryNotes

`func (o *VerificationSuccessResponse) SetMandatoryNotes(v []MandatoryNoteSchema)`

SetMandatoryNotes sets MandatoryNotes field to given value.

### HasMandatoryNotes

`func (o *VerificationSuccessResponse) HasMandatoryNotes() bool`

HasMandatoryNotes returns a boolean if a field has been set.

### GetPageDimensions

`func (o *VerificationSuccessResponse) GetPageDimensions() []PageDimensionsSchema`

GetPageDimensions returns the PageDimensions field if non-nil, zero value otherwise.

### GetPageDimensionsOk

`func (o *VerificationSuccessResponse) GetPageDimensionsOk() (*[]PageDimensionsSchema, bool)`

GetPageDimensionsOk returns a tuple with the PageDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageDimensions

`func (o *VerificationSuccessResponse) SetPageDimensions(v []PageDimensionsSchema)`

SetPageDimensions sets PageDimensions field to given value.

### HasPageDimensions

`func (o *VerificationSuccessResponse) HasPageDimensions() bool`

HasPageDimensions returns a boolean if a field has been set.

### GetWarnings

`func (o *VerificationSuccessResponse) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *VerificationSuccessResponse) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *VerificationSuccessResponse) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *VerificationSuccessResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



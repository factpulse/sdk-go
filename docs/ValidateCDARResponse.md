# ValidateCDARResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | **bool** | Résultat de validation | 
**Errors** | Pointer to [**[]ValidationErrorResponse**](ValidationErrorResponse.md) | Liste des erreurs | [optional] 
**Warnings** | Pointer to [**[]ValidationErrorResponse**](ValidationErrorResponse.md) | Liste des avertissements | [optional] 

## Methods

### NewValidateCDARResponse

`func NewValidateCDARResponse(valid bool, ) *ValidateCDARResponse`

NewValidateCDARResponse instantiates a new ValidateCDARResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateCDARResponseWithDefaults

`func NewValidateCDARResponseWithDefaults() *ValidateCDARResponse`

NewValidateCDARResponseWithDefaults instantiates a new ValidateCDARResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *ValidateCDARResponse) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *ValidateCDARResponse) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *ValidateCDARResponse) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetErrors

`func (o *ValidateCDARResponse) GetErrors() []ValidationErrorResponse`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ValidateCDARResponse) GetErrorsOk() (*[]ValidationErrorResponse, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ValidateCDARResponse) SetErrors(v []ValidationErrorResponse)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ValidateCDARResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *ValidateCDARResponse) GetWarnings() []ValidationErrorResponse`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ValidateCDARResponse) GetWarningsOk() (*[]ValidationErrorResponse, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ValidateCDARResponse) SetWarnings(v []ValidationErrorResponse)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ValidateCDARResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



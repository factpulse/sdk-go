# StatutFacture

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Code statut (SOUMISE, VALIDEE, REJETEE, SUSPENDUE, MANDATEE, MISE_EN_PAIEMENT, etc.) | 
**Libelle** | **string** | Libellé du statut | 
**Date** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewStatutFacture

`func NewStatutFacture(code string, libelle string, ) *StatutFacture`

NewStatutFacture instantiates a new StatutFacture object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatutFactureWithDefaults

`func NewStatutFactureWithDefaults() *StatutFacture`

NewStatutFactureWithDefaults instantiates a new StatutFacture object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *StatutFacture) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *StatutFacture) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *StatutFacture) SetCode(v string)`

SetCode sets Code field to given value.


### GetLibelle

`func (o *StatutFacture) GetLibelle() string`

GetLibelle returns the Libelle field if non-nil, zero value otherwise.

### GetLibelleOk

`func (o *StatutFacture) GetLibelleOk() (*string, bool)`

GetLibelleOk returns a tuple with the Libelle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibelle

`func (o *StatutFacture) SetLibelle(v string)`

SetLibelle sets Libelle field to given value.


### GetDate

`func (o *StatutFacture) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *StatutFacture) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *StatutFacture) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *StatutFacture) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *StatutFacture) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *StatutFacture) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



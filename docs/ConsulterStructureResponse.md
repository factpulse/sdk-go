# ConsulterStructureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodeRetour** | **int32** |  | 
**Libelle** | **string** |  | 
**IdStructureCpp** | Pointer to **NullableInt32** |  | [optional] 
**IdentifiantStructure** | Pointer to **NullableString** |  | [optional] 
**LibelleStructure** | Pointer to **NullableString** |  | [optional] 
**RaisonSocialeStructure** | Pointer to **NullableString** |  | [optional] 
**NumeroTva** | Pointer to **NullableString** |  | [optional] 
**EmailStructure** | Pointer to **NullableString** |  | [optional] 
**Parametres** | Pointer to [**NullableParametresStructure**](ParametresStructure.md) |  | [optional] 

## Methods

### NewConsulterStructureResponse

`func NewConsulterStructureResponse(codeRetour int32, libelle string, ) *ConsulterStructureResponse`

NewConsulterStructureResponse instantiates a new ConsulterStructureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsulterStructureResponseWithDefaults

`func NewConsulterStructureResponseWithDefaults() *ConsulterStructureResponse`

NewConsulterStructureResponseWithDefaults instantiates a new ConsulterStructureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodeRetour

`func (o *ConsulterStructureResponse) GetCodeRetour() int32`

GetCodeRetour returns the CodeRetour field if non-nil, zero value otherwise.

### GetCodeRetourOk

`func (o *ConsulterStructureResponse) GetCodeRetourOk() (*int32, bool)`

GetCodeRetourOk returns a tuple with the CodeRetour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeRetour

`func (o *ConsulterStructureResponse) SetCodeRetour(v int32)`

SetCodeRetour sets CodeRetour field to given value.


### GetLibelle

`func (o *ConsulterStructureResponse) GetLibelle() string`

GetLibelle returns the Libelle field if non-nil, zero value otherwise.

### GetLibelleOk

`func (o *ConsulterStructureResponse) GetLibelleOk() (*string, bool)`

GetLibelleOk returns a tuple with the Libelle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibelle

`func (o *ConsulterStructureResponse) SetLibelle(v string)`

SetLibelle sets Libelle field to given value.


### GetIdStructureCpp

`func (o *ConsulterStructureResponse) GetIdStructureCpp() int32`

GetIdStructureCpp returns the IdStructureCpp field if non-nil, zero value otherwise.

### GetIdStructureCppOk

`func (o *ConsulterStructureResponse) GetIdStructureCppOk() (*int32, bool)`

GetIdStructureCppOk returns a tuple with the IdStructureCpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdStructureCpp

`func (o *ConsulterStructureResponse) SetIdStructureCpp(v int32)`

SetIdStructureCpp sets IdStructureCpp field to given value.

### HasIdStructureCpp

`func (o *ConsulterStructureResponse) HasIdStructureCpp() bool`

HasIdStructureCpp returns a boolean if a field has been set.

### SetIdStructureCppNil

`func (o *ConsulterStructureResponse) SetIdStructureCppNil(b bool)`

 SetIdStructureCppNil sets the value for IdStructureCpp to be an explicit nil

### UnsetIdStructureCpp
`func (o *ConsulterStructureResponse) UnsetIdStructureCpp()`

UnsetIdStructureCpp ensures that no value is present for IdStructureCpp, not even an explicit nil
### GetIdentifiantStructure

`func (o *ConsulterStructureResponse) GetIdentifiantStructure() string`

GetIdentifiantStructure returns the IdentifiantStructure field if non-nil, zero value otherwise.

### GetIdentifiantStructureOk

`func (o *ConsulterStructureResponse) GetIdentifiantStructureOk() (*string, bool)`

GetIdentifiantStructureOk returns a tuple with the IdentifiantStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiantStructure

`func (o *ConsulterStructureResponse) SetIdentifiantStructure(v string)`

SetIdentifiantStructure sets IdentifiantStructure field to given value.

### HasIdentifiantStructure

`func (o *ConsulterStructureResponse) HasIdentifiantStructure() bool`

HasIdentifiantStructure returns a boolean if a field has been set.

### SetIdentifiantStructureNil

`func (o *ConsulterStructureResponse) SetIdentifiantStructureNil(b bool)`

 SetIdentifiantStructureNil sets the value for IdentifiantStructure to be an explicit nil

### UnsetIdentifiantStructure
`func (o *ConsulterStructureResponse) UnsetIdentifiantStructure()`

UnsetIdentifiantStructure ensures that no value is present for IdentifiantStructure, not even an explicit nil
### GetLibelleStructure

`func (o *ConsulterStructureResponse) GetLibelleStructure() string`

GetLibelleStructure returns the LibelleStructure field if non-nil, zero value otherwise.

### GetLibelleStructureOk

`func (o *ConsulterStructureResponse) GetLibelleStructureOk() (*string, bool)`

GetLibelleStructureOk returns a tuple with the LibelleStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibelleStructure

`func (o *ConsulterStructureResponse) SetLibelleStructure(v string)`

SetLibelleStructure sets LibelleStructure field to given value.

### HasLibelleStructure

`func (o *ConsulterStructureResponse) HasLibelleStructure() bool`

HasLibelleStructure returns a boolean if a field has been set.

### SetLibelleStructureNil

`func (o *ConsulterStructureResponse) SetLibelleStructureNil(b bool)`

 SetLibelleStructureNil sets the value for LibelleStructure to be an explicit nil

### UnsetLibelleStructure
`func (o *ConsulterStructureResponse) UnsetLibelleStructure()`

UnsetLibelleStructure ensures that no value is present for LibelleStructure, not even an explicit nil
### GetRaisonSocialeStructure

`func (o *ConsulterStructureResponse) GetRaisonSocialeStructure() string`

GetRaisonSocialeStructure returns the RaisonSocialeStructure field if non-nil, zero value otherwise.

### GetRaisonSocialeStructureOk

`func (o *ConsulterStructureResponse) GetRaisonSocialeStructureOk() (*string, bool)`

GetRaisonSocialeStructureOk returns a tuple with the RaisonSocialeStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaisonSocialeStructure

`func (o *ConsulterStructureResponse) SetRaisonSocialeStructure(v string)`

SetRaisonSocialeStructure sets RaisonSocialeStructure field to given value.

### HasRaisonSocialeStructure

`func (o *ConsulterStructureResponse) HasRaisonSocialeStructure() bool`

HasRaisonSocialeStructure returns a boolean if a field has been set.

### SetRaisonSocialeStructureNil

`func (o *ConsulterStructureResponse) SetRaisonSocialeStructureNil(b bool)`

 SetRaisonSocialeStructureNil sets the value for RaisonSocialeStructure to be an explicit nil

### UnsetRaisonSocialeStructure
`func (o *ConsulterStructureResponse) UnsetRaisonSocialeStructure()`

UnsetRaisonSocialeStructure ensures that no value is present for RaisonSocialeStructure, not even an explicit nil
### GetNumeroTva

`func (o *ConsulterStructureResponse) GetNumeroTva() string`

GetNumeroTva returns the NumeroTva field if non-nil, zero value otherwise.

### GetNumeroTvaOk

`func (o *ConsulterStructureResponse) GetNumeroTvaOk() (*string, bool)`

GetNumeroTvaOk returns a tuple with the NumeroTva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroTva

`func (o *ConsulterStructureResponse) SetNumeroTva(v string)`

SetNumeroTva sets NumeroTva field to given value.

### HasNumeroTva

`func (o *ConsulterStructureResponse) HasNumeroTva() bool`

HasNumeroTva returns a boolean if a field has been set.

### SetNumeroTvaNil

`func (o *ConsulterStructureResponse) SetNumeroTvaNil(b bool)`

 SetNumeroTvaNil sets the value for NumeroTva to be an explicit nil

### UnsetNumeroTva
`func (o *ConsulterStructureResponse) UnsetNumeroTva()`

UnsetNumeroTva ensures that no value is present for NumeroTva, not even an explicit nil
### GetEmailStructure

`func (o *ConsulterStructureResponse) GetEmailStructure() string`

GetEmailStructure returns the EmailStructure field if non-nil, zero value otherwise.

### GetEmailStructureOk

`func (o *ConsulterStructureResponse) GetEmailStructureOk() (*string, bool)`

GetEmailStructureOk returns a tuple with the EmailStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailStructure

`func (o *ConsulterStructureResponse) SetEmailStructure(v string)`

SetEmailStructure sets EmailStructure field to given value.

### HasEmailStructure

`func (o *ConsulterStructureResponse) HasEmailStructure() bool`

HasEmailStructure returns a boolean if a field has been set.

### SetEmailStructureNil

`func (o *ConsulterStructureResponse) SetEmailStructureNil(b bool)`

 SetEmailStructureNil sets the value for EmailStructure to be an explicit nil

### UnsetEmailStructure
`func (o *ConsulterStructureResponse) UnsetEmailStructure()`

UnsetEmailStructure ensures that no value is present for EmailStructure, not even an explicit nil
### GetParametres

`func (o *ConsulterStructureResponse) GetParametres() ParametresStructure`

GetParametres returns the Parametres field if non-nil, zero value otherwise.

### GetParametresOk

`func (o *ConsulterStructureResponse) GetParametresOk() (*ParametresStructure, bool)`

GetParametresOk returns a tuple with the Parametres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParametres

`func (o *ConsulterStructureResponse) SetParametres(v ParametresStructure)`

SetParametres sets Parametres field to given value.

### HasParametres

`func (o *ConsulterStructureResponse) HasParametres() bool`

HasParametres returns a boolean if a field has been set.

### SetParametresNil

`func (o *ConsulterStructureResponse) SetParametresNil(b bool)`

 SetParametresNil sets the value for Parametres to be an explicit nil

### UnsetParametres
`func (o *ConsulterStructureResponse) UnsetParametres()`

UnsetParametres ensures that no value is present for Parametres, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



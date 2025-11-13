# ServiceStructure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdService** | **int32** | ID du service | 
**CodeService** | **string** | Code du service | 
**LibelleService** | **string** | Libellé du service | 
**EstActif** | **bool** | Service actif | 

## Methods

### NewServiceStructure

`func NewServiceStructure(idService int32, codeService string, libelleService string, estActif bool, ) *ServiceStructure`

NewServiceStructure instantiates a new ServiceStructure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceStructureWithDefaults

`func NewServiceStructureWithDefaults() *ServiceStructure`

NewServiceStructureWithDefaults instantiates a new ServiceStructure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdService

`func (o *ServiceStructure) GetIdService() int32`

GetIdService returns the IdService field if non-nil, zero value otherwise.

### GetIdServiceOk

`func (o *ServiceStructure) GetIdServiceOk() (*int32, bool)`

GetIdServiceOk returns a tuple with the IdService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdService

`func (o *ServiceStructure) SetIdService(v int32)`

SetIdService sets IdService field to given value.


### GetCodeService

`func (o *ServiceStructure) GetCodeService() string`

GetCodeService returns the CodeService field if non-nil, zero value otherwise.

### GetCodeServiceOk

`func (o *ServiceStructure) GetCodeServiceOk() (*string, bool)`

GetCodeServiceOk returns a tuple with the CodeService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeService

`func (o *ServiceStructure) SetCodeService(v string)`

SetCodeService sets CodeService field to given value.


### GetLibelleService

`func (o *ServiceStructure) GetLibelleService() string`

GetLibelleService returns the LibelleService field if non-nil, zero value otherwise.

### GetLibelleServiceOk

`func (o *ServiceStructure) GetLibelleServiceOk() (*string, bool)`

GetLibelleServiceOk returns a tuple with the LibelleService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibelleService

`func (o *ServiceStructure) SetLibelleService(v string)`

SetLibelleService sets LibelleService field to given value.


### GetEstActif

`func (o *ServiceStructure) GetEstActif() bool`

GetEstActif returns the EstActif field if non-nil, zero value otherwise.

### GetEstActifOk

`func (o *ServiceStructure) GetEstActifOk() (*bool, bool)`

GetEstActifOk returns a tuple with the EstActif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstActif

`func (o *ServiceStructure) SetEstActif(v bool)`

SetEstActif sets EstActif field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



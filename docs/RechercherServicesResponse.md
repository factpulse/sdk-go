# RechercherServicesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodeRetour** | **int32** |  | 
**Libelle** | **string** |  | 
**ListeServices** | Pointer to [**[]ServiceStructure**](ServiceStructure.md) |  | [optional] 
**Total** | Pointer to **int32** | Nombre de services | [optional] [default to 0]

## Methods

### NewRechercherServicesResponse

`func NewRechercherServicesResponse(codeRetour int32, libelle string, ) *RechercherServicesResponse`

NewRechercherServicesResponse instantiates a new RechercherServicesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRechercherServicesResponseWithDefaults

`func NewRechercherServicesResponseWithDefaults() *RechercherServicesResponse`

NewRechercherServicesResponseWithDefaults instantiates a new RechercherServicesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodeRetour

`func (o *RechercherServicesResponse) GetCodeRetour() int32`

GetCodeRetour returns the CodeRetour field if non-nil, zero value otherwise.

### GetCodeRetourOk

`func (o *RechercherServicesResponse) GetCodeRetourOk() (*int32, bool)`

GetCodeRetourOk returns a tuple with the CodeRetour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeRetour

`func (o *RechercherServicesResponse) SetCodeRetour(v int32)`

SetCodeRetour sets CodeRetour field to given value.


### GetLibelle

`func (o *RechercherServicesResponse) GetLibelle() string`

GetLibelle returns the Libelle field if non-nil, zero value otherwise.

### GetLibelleOk

`func (o *RechercherServicesResponse) GetLibelleOk() (*string, bool)`

GetLibelleOk returns a tuple with the Libelle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibelle

`func (o *RechercherServicesResponse) SetLibelle(v string)`

SetLibelle sets Libelle field to given value.


### GetListeServices

`func (o *RechercherServicesResponse) GetListeServices() []ServiceStructure`

GetListeServices returns the ListeServices field if non-nil, zero value otherwise.

### GetListeServicesOk

`func (o *RechercherServicesResponse) GetListeServicesOk() (*[]ServiceStructure, bool)`

GetListeServicesOk returns a tuple with the ListeServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeServices

`func (o *RechercherServicesResponse) SetListeServices(v []ServiceStructure)`

SetListeServices sets ListeServices field to given value.

### HasListeServices

`func (o *RechercherServicesResponse) HasListeServices() bool`

HasListeServices returns a boolean if a field has been set.

### GetTotal

`func (o *RechercherServicesResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RechercherServicesResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RechercherServicesResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *RechercherServicesResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



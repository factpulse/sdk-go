# ConsulterFactureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodeRetour** | **int32** |  | 
**Libelle** | **string** |  | 
**IdentifiantFactureCpp** | Pointer to **NullableInt32** |  | [optional] 
**NumeroFacture** | Pointer to **NullableString** |  | [optional] 
**DateFacture** | Pointer to **NullableString** |  | [optional] 
**MontantTtcTotal** | Pointer to **NullableString** |  | [optional] 
**StatutCourant** | Pointer to [**NullableStatutFacture**](StatutFacture.md) |  | [optional] 
**IdStructureCppDestinataire** | Pointer to **NullableInt32** |  | [optional] 
**DesignationStructureDestinataire** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewConsulterFactureResponse

`func NewConsulterFactureResponse(codeRetour int32, libelle string, ) *ConsulterFactureResponse`

NewConsulterFactureResponse instantiates a new ConsulterFactureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsulterFactureResponseWithDefaults

`func NewConsulterFactureResponseWithDefaults() *ConsulterFactureResponse`

NewConsulterFactureResponseWithDefaults instantiates a new ConsulterFactureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodeRetour

`func (o *ConsulterFactureResponse) GetCodeRetour() int32`

GetCodeRetour returns the CodeRetour field if non-nil, zero value otherwise.

### GetCodeRetourOk

`func (o *ConsulterFactureResponse) GetCodeRetourOk() (*int32, bool)`

GetCodeRetourOk returns a tuple with the CodeRetour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeRetour

`func (o *ConsulterFactureResponse) SetCodeRetour(v int32)`

SetCodeRetour sets CodeRetour field to given value.


### GetLibelle

`func (o *ConsulterFactureResponse) GetLibelle() string`

GetLibelle returns the Libelle field if non-nil, zero value otherwise.

### GetLibelleOk

`func (o *ConsulterFactureResponse) GetLibelleOk() (*string, bool)`

GetLibelleOk returns a tuple with the Libelle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibelle

`func (o *ConsulterFactureResponse) SetLibelle(v string)`

SetLibelle sets Libelle field to given value.


### GetIdentifiantFactureCpp

`func (o *ConsulterFactureResponse) GetIdentifiantFactureCpp() int32`

GetIdentifiantFactureCpp returns the IdentifiantFactureCpp field if non-nil, zero value otherwise.

### GetIdentifiantFactureCppOk

`func (o *ConsulterFactureResponse) GetIdentifiantFactureCppOk() (*int32, bool)`

GetIdentifiantFactureCppOk returns a tuple with the IdentifiantFactureCpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiantFactureCpp

`func (o *ConsulterFactureResponse) SetIdentifiantFactureCpp(v int32)`

SetIdentifiantFactureCpp sets IdentifiantFactureCpp field to given value.

### HasIdentifiantFactureCpp

`func (o *ConsulterFactureResponse) HasIdentifiantFactureCpp() bool`

HasIdentifiantFactureCpp returns a boolean if a field has been set.

### SetIdentifiantFactureCppNil

`func (o *ConsulterFactureResponse) SetIdentifiantFactureCppNil(b bool)`

 SetIdentifiantFactureCppNil sets the value for IdentifiantFactureCpp to be an explicit nil

### UnsetIdentifiantFactureCpp
`func (o *ConsulterFactureResponse) UnsetIdentifiantFactureCpp()`

UnsetIdentifiantFactureCpp ensures that no value is present for IdentifiantFactureCpp, not even an explicit nil
### GetNumeroFacture

`func (o *ConsulterFactureResponse) GetNumeroFacture() string`

GetNumeroFacture returns the NumeroFacture field if non-nil, zero value otherwise.

### GetNumeroFactureOk

`func (o *ConsulterFactureResponse) GetNumeroFactureOk() (*string, bool)`

GetNumeroFactureOk returns a tuple with the NumeroFacture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroFacture

`func (o *ConsulterFactureResponse) SetNumeroFacture(v string)`

SetNumeroFacture sets NumeroFacture field to given value.

### HasNumeroFacture

`func (o *ConsulterFactureResponse) HasNumeroFacture() bool`

HasNumeroFacture returns a boolean if a field has been set.

### SetNumeroFactureNil

`func (o *ConsulterFactureResponse) SetNumeroFactureNil(b bool)`

 SetNumeroFactureNil sets the value for NumeroFacture to be an explicit nil

### UnsetNumeroFacture
`func (o *ConsulterFactureResponse) UnsetNumeroFacture()`

UnsetNumeroFacture ensures that no value is present for NumeroFacture, not even an explicit nil
### GetDateFacture

`func (o *ConsulterFactureResponse) GetDateFacture() string`

GetDateFacture returns the DateFacture field if non-nil, zero value otherwise.

### GetDateFactureOk

`func (o *ConsulterFactureResponse) GetDateFactureOk() (*string, bool)`

GetDateFactureOk returns a tuple with the DateFacture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFacture

`func (o *ConsulterFactureResponse) SetDateFacture(v string)`

SetDateFacture sets DateFacture field to given value.

### HasDateFacture

`func (o *ConsulterFactureResponse) HasDateFacture() bool`

HasDateFacture returns a boolean if a field has been set.

### SetDateFactureNil

`func (o *ConsulterFactureResponse) SetDateFactureNil(b bool)`

 SetDateFactureNil sets the value for DateFacture to be an explicit nil

### UnsetDateFacture
`func (o *ConsulterFactureResponse) UnsetDateFacture()`

UnsetDateFacture ensures that no value is present for DateFacture, not even an explicit nil
### GetMontantTtcTotal

`func (o *ConsulterFactureResponse) GetMontantTtcTotal() string`

GetMontantTtcTotal returns the MontantTtcTotal field if non-nil, zero value otherwise.

### GetMontantTtcTotalOk

`func (o *ConsulterFactureResponse) GetMontantTtcTotalOk() (*string, bool)`

GetMontantTtcTotalOk returns a tuple with the MontantTtcTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantTtcTotal

`func (o *ConsulterFactureResponse) SetMontantTtcTotal(v string)`

SetMontantTtcTotal sets MontantTtcTotal field to given value.

### HasMontantTtcTotal

`func (o *ConsulterFactureResponse) HasMontantTtcTotal() bool`

HasMontantTtcTotal returns a boolean if a field has been set.

### SetMontantTtcTotalNil

`func (o *ConsulterFactureResponse) SetMontantTtcTotalNil(b bool)`

 SetMontantTtcTotalNil sets the value for MontantTtcTotal to be an explicit nil

### UnsetMontantTtcTotal
`func (o *ConsulterFactureResponse) UnsetMontantTtcTotal()`

UnsetMontantTtcTotal ensures that no value is present for MontantTtcTotal, not even an explicit nil
### GetStatutCourant

`func (o *ConsulterFactureResponse) GetStatutCourant() StatutFacture`

GetStatutCourant returns the StatutCourant field if non-nil, zero value otherwise.

### GetStatutCourantOk

`func (o *ConsulterFactureResponse) GetStatutCourantOk() (*StatutFacture, bool)`

GetStatutCourantOk returns a tuple with the StatutCourant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatutCourant

`func (o *ConsulterFactureResponse) SetStatutCourant(v StatutFacture)`

SetStatutCourant sets StatutCourant field to given value.

### HasStatutCourant

`func (o *ConsulterFactureResponse) HasStatutCourant() bool`

HasStatutCourant returns a boolean if a field has been set.

### SetStatutCourantNil

`func (o *ConsulterFactureResponse) SetStatutCourantNil(b bool)`

 SetStatutCourantNil sets the value for StatutCourant to be an explicit nil

### UnsetStatutCourant
`func (o *ConsulterFactureResponse) UnsetStatutCourant()`

UnsetStatutCourant ensures that no value is present for StatutCourant, not even an explicit nil
### GetIdStructureCppDestinataire

`func (o *ConsulterFactureResponse) GetIdStructureCppDestinataire() int32`

GetIdStructureCppDestinataire returns the IdStructureCppDestinataire field if non-nil, zero value otherwise.

### GetIdStructureCppDestinataireOk

`func (o *ConsulterFactureResponse) GetIdStructureCppDestinataireOk() (*int32, bool)`

GetIdStructureCppDestinataireOk returns a tuple with the IdStructureCppDestinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdStructureCppDestinataire

`func (o *ConsulterFactureResponse) SetIdStructureCppDestinataire(v int32)`

SetIdStructureCppDestinataire sets IdStructureCppDestinataire field to given value.

### HasIdStructureCppDestinataire

`func (o *ConsulterFactureResponse) HasIdStructureCppDestinataire() bool`

HasIdStructureCppDestinataire returns a boolean if a field has been set.

### SetIdStructureCppDestinataireNil

`func (o *ConsulterFactureResponse) SetIdStructureCppDestinataireNil(b bool)`

 SetIdStructureCppDestinataireNil sets the value for IdStructureCppDestinataire to be an explicit nil

### UnsetIdStructureCppDestinataire
`func (o *ConsulterFactureResponse) UnsetIdStructureCppDestinataire()`

UnsetIdStructureCppDestinataire ensures that no value is present for IdStructureCppDestinataire, not even an explicit nil
### GetDesignationStructureDestinataire

`func (o *ConsulterFactureResponse) GetDesignationStructureDestinataire() string`

GetDesignationStructureDestinataire returns the DesignationStructureDestinataire field if non-nil, zero value otherwise.

### GetDesignationStructureDestinataireOk

`func (o *ConsulterFactureResponse) GetDesignationStructureDestinataireOk() (*string, bool)`

GetDesignationStructureDestinataireOk returns a tuple with the DesignationStructureDestinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignationStructureDestinataire

`func (o *ConsulterFactureResponse) SetDesignationStructureDestinataire(v string)`

SetDesignationStructureDestinataire sets DesignationStructureDestinataire field to given value.

### HasDesignationStructureDestinataire

`func (o *ConsulterFactureResponse) HasDesignationStructureDestinataire() bool`

HasDesignationStructureDestinataire returns a boolean if a field has been set.

### SetDesignationStructureDestinataireNil

`func (o *ConsulterFactureResponse) SetDesignationStructureDestinataireNil(b bool)`

 SetDesignationStructureDestinataireNil sets the value for DesignationStructureDestinataire to be an explicit nil

### UnsetDesignationStructureDestinataire
`func (o *ConsulterFactureResponse) UnsetDesignationStructureDestinataire()`

UnsetDesignationStructureDestinataire ensures that no value is present for DesignationStructureDestinataire, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



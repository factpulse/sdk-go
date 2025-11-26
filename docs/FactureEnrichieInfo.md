# FactureEnrichieInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumeroFacture** | **string** |  | 
**IdEmetteur** | Pointer to **NullableInt32** |  | [optional] 
**IdDestinataire** | Pointer to **NullableInt32** |  | [optional] 
**NomEmetteur** | **string** |  | 
**NomDestinataire** | **string** |  | 
**MontantHtTotal** | **string** |  | 
**MontantTva** | **string** |  | 
**MontantTtcTotal** | **string** |  | 

## Methods

### NewFactureEnrichieInfo

`func NewFactureEnrichieInfo(numeroFacture string, nomEmetteur string, nomDestinataire string, montantHtTotal string, montantTva string, montantTtcTotal string, ) *FactureEnrichieInfo`

NewFactureEnrichieInfo instantiates a new FactureEnrichieInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFactureEnrichieInfoWithDefaults

`func NewFactureEnrichieInfoWithDefaults() *FactureEnrichieInfo`

NewFactureEnrichieInfoWithDefaults instantiates a new FactureEnrichieInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumeroFacture

`func (o *FactureEnrichieInfo) GetNumeroFacture() string`

GetNumeroFacture returns the NumeroFacture field if non-nil, zero value otherwise.

### GetNumeroFactureOk

`func (o *FactureEnrichieInfo) GetNumeroFactureOk() (*string, bool)`

GetNumeroFactureOk returns a tuple with the NumeroFacture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroFacture

`func (o *FactureEnrichieInfo) SetNumeroFacture(v string)`

SetNumeroFacture sets NumeroFacture field to given value.


### GetIdEmetteur

`func (o *FactureEnrichieInfo) GetIdEmetteur() int32`

GetIdEmetteur returns the IdEmetteur field if non-nil, zero value otherwise.

### GetIdEmetteurOk

`func (o *FactureEnrichieInfo) GetIdEmetteurOk() (*int32, bool)`

GetIdEmetteurOk returns a tuple with the IdEmetteur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdEmetteur

`func (o *FactureEnrichieInfo) SetIdEmetteur(v int32)`

SetIdEmetteur sets IdEmetteur field to given value.

### HasIdEmetteur

`func (o *FactureEnrichieInfo) HasIdEmetteur() bool`

HasIdEmetteur returns a boolean if a field has been set.

### SetIdEmetteurNil

`func (o *FactureEnrichieInfo) SetIdEmetteurNil(b bool)`

 SetIdEmetteurNil sets the value for IdEmetteur to be an explicit nil

### UnsetIdEmetteur
`func (o *FactureEnrichieInfo) UnsetIdEmetteur()`

UnsetIdEmetteur ensures that no value is present for IdEmetteur, not even an explicit nil
### GetIdDestinataire

`func (o *FactureEnrichieInfo) GetIdDestinataire() int32`

GetIdDestinataire returns the IdDestinataire field if non-nil, zero value otherwise.

### GetIdDestinataireOk

`func (o *FactureEnrichieInfo) GetIdDestinataireOk() (*int32, bool)`

GetIdDestinataireOk returns a tuple with the IdDestinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdDestinataire

`func (o *FactureEnrichieInfo) SetIdDestinataire(v int32)`

SetIdDestinataire sets IdDestinataire field to given value.

### HasIdDestinataire

`func (o *FactureEnrichieInfo) HasIdDestinataire() bool`

HasIdDestinataire returns a boolean if a field has been set.

### SetIdDestinataireNil

`func (o *FactureEnrichieInfo) SetIdDestinataireNil(b bool)`

 SetIdDestinataireNil sets the value for IdDestinataire to be an explicit nil

### UnsetIdDestinataire
`func (o *FactureEnrichieInfo) UnsetIdDestinataire()`

UnsetIdDestinataire ensures that no value is present for IdDestinataire, not even an explicit nil
### GetNomEmetteur

`func (o *FactureEnrichieInfo) GetNomEmetteur() string`

GetNomEmetteur returns the NomEmetteur field if non-nil, zero value otherwise.

### GetNomEmetteurOk

`func (o *FactureEnrichieInfo) GetNomEmetteurOk() (*string, bool)`

GetNomEmetteurOk returns a tuple with the NomEmetteur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomEmetteur

`func (o *FactureEnrichieInfo) SetNomEmetteur(v string)`

SetNomEmetteur sets NomEmetteur field to given value.


### GetNomDestinataire

`func (o *FactureEnrichieInfo) GetNomDestinataire() string`

GetNomDestinataire returns the NomDestinataire field if non-nil, zero value otherwise.

### GetNomDestinataireOk

`func (o *FactureEnrichieInfo) GetNomDestinataireOk() (*string, bool)`

GetNomDestinataireOk returns a tuple with the NomDestinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomDestinataire

`func (o *FactureEnrichieInfo) SetNomDestinataire(v string)`

SetNomDestinataire sets NomDestinataire field to given value.


### GetMontantHtTotal

`func (o *FactureEnrichieInfo) GetMontantHtTotal() string`

GetMontantHtTotal returns the MontantHtTotal field if non-nil, zero value otherwise.

### GetMontantHtTotalOk

`func (o *FactureEnrichieInfo) GetMontantHtTotalOk() (*string, bool)`

GetMontantHtTotalOk returns a tuple with the MontantHtTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantHtTotal

`func (o *FactureEnrichieInfo) SetMontantHtTotal(v string)`

SetMontantHtTotal sets MontantHtTotal field to given value.


### GetMontantTva

`func (o *FactureEnrichieInfo) GetMontantTva() string`

GetMontantTva returns the MontantTva field if non-nil, zero value otherwise.

### GetMontantTvaOk

`func (o *FactureEnrichieInfo) GetMontantTvaOk() (*string, bool)`

GetMontantTvaOk returns a tuple with the MontantTva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantTva

`func (o *FactureEnrichieInfo) SetMontantTva(v string)`

SetMontantTva sets MontantTva field to given value.


### GetMontantTtcTotal

`func (o *FactureEnrichieInfo) GetMontantTtcTotal() string`

GetMontantTtcTotal returns the MontantTtcTotal field if non-nil, zero value otherwise.

### GetMontantTtcTotalOk

`func (o *FactureEnrichieInfo) GetMontantTtcTotalOk() (*string, bool)`

GetMontantTtcTotalOk returns a tuple with the MontantTtcTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantTtcTotal

`func (o *FactureEnrichieInfo) SetMontantTtcTotal(v string)`

SetMontantTtcTotal sets MontantTtcTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



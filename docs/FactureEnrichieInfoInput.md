# FactureEnrichieInfoInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumeroFacture** | **string** |  | 
**IdEmetteur** | Pointer to **NullableInt32** |  | [optional] 
**IdDestinataire** | Pointer to **NullableInt32** |  | [optional] 
**NomEmetteur** | **string** |  | 
**NomDestinataire** | **string** |  | 
**MontantHtTotal** | [**MontantHtTotal**](MontantHtTotal.md) |  | 
**MontantTva** | [**MontantTva**](MontantTva.md) |  | 
**MontantTtcTotal** | [**MontantTtcTotal**](MontantTtcTotal.md) |  | 

## Methods

### NewFactureEnrichieInfoInput

`func NewFactureEnrichieInfoInput(numeroFacture string, nomEmetteur string, nomDestinataire string, montantHtTotal MontantHtTotal, montantTva MontantTva, montantTtcTotal MontantTtcTotal, ) *FactureEnrichieInfoInput`

NewFactureEnrichieInfoInput instantiates a new FactureEnrichieInfoInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFactureEnrichieInfoInputWithDefaults

`func NewFactureEnrichieInfoInputWithDefaults() *FactureEnrichieInfoInput`

NewFactureEnrichieInfoInputWithDefaults instantiates a new FactureEnrichieInfoInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumeroFacture

`func (o *FactureEnrichieInfoInput) GetNumeroFacture() string`

GetNumeroFacture returns the NumeroFacture field if non-nil, zero value otherwise.

### GetNumeroFactureOk

`func (o *FactureEnrichieInfoInput) GetNumeroFactureOk() (*string, bool)`

GetNumeroFactureOk returns a tuple with the NumeroFacture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroFacture

`func (o *FactureEnrichieInfoInput) SetNumeroFacture(v string)`

SetNumeroFacture sets NumeroFacture field to given value.


### GetIdEmetteur

`func (o *FactureEnrichieInfoInput) GetIdEmetteur() int32`

GetIdEmetteur returns the IdEmetteur field if non-nil, zero value otherwise.

### GetIdEmetteurOk

`func (o *FactureEnrichieInfoInput) GetIdEmetteurOk() (*int32, bool)`

GetIdEmetteurOk returns a tuple with the IdEmetteur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdEmetteur

`func (o *FactureEnrichieInfoInput) SetIdEmetteur(v int32)`

SetIdEmetteur sets IdEmetteur field to given value.

### HasIdEmetteur

`func (o *FactureEnrichieInfoInput) HasIdEmetteur() bool`

HasIdEmetteur returns a boolean if a field has been set.

### SetIdEmetteurNil

`func (o *FactureEnrichieInfoInput) SetIdEmetteurNil(b bool)`

 SetIdEmetteurNil sets the value for IdEmetteur to be an explicit nil

### UnsetIdEmetteur
`func (o *FactureEnrichieInfoInput) UnsetIdEmetteur()`

UnsetIdEmetteur ensures that no value is present for IdEmetteur, not even an explicit nil
### GetIdDestinataire

`func (o *FactureEnrichieInfoInput) GetIdDestinataire() int32`

GetIdDestinataire returns the IdDestinataire field if non-nil, zero value otherwise.

### GetIdDestinataireOk

`func (o *FactureEnrichieInfoInput) GetIdDestinataireOk() (*int32, bool)`

GetIdDestinataireOk returns a tuple with the IdDestinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdDestinataire

`func (o *FactureEnrichieInfoInput) SetIdDestinataire(v int32)`

SetIdDestinataire sets IdDestinataire field to given value.

### HasIdDestinataire

`func (o *FactureEnrichieInfoInput) HasIdDestinataire() bool`

HasIdDestinataire returns a boolean if a field has been set.

### SetIdDestinataireNil

`func (o *FactureEnrichieInfoInput) SetIdDestinataireNil(b bool)`

 SetIdDestinataireNil sets the value for IdDestinataire to be an explicit nil

### UnsetIdDestinataire
`func (o *FactureEnrichieInfoInput) UnsetIdDestinataire()`

UnsetIdDestinataire ensures that no value is present for IdDestinataire, not even an explicit nil
### GetNomEmetteur

`func (o *FactureEnrichieInfoInput) GetNomEmetteur() string`

GetNomEmetteur returns the NomEmetteur field if non-nil, zero value otherwise.

### GetNomEmetteurOk

`func (o *FactureEnrichieInfoInput) GetNomEmetteurOk() (*string, bool)`

GetNomEmetteurOk returns a tuple with the NomEmetteur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomEmetteur

`func (o *FactureEnrichieInfoInput) SetNomEmetteur(v string)`

SetNomEmetteur sets NomEmetteur field to given value.


### GetNomDestinataire

`func (o *FactureEnrichieInfoInput) GetNomDestinataire() string`

GetNomDestinataire returns the NomDestinataire field if non-nil, zero value otherwise.

### GetNomDestinataireOk

`func (o *FactureEnrichieInfoInput) GetNomDestinataireOk() (*string, bool)`

GetNomDestinataireOk returns a tuple with the NomDestinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomDestinataire

`func (o *FactureEnrichieInfoInput) SetNomDestinataire(v string)`

SetNomDestinataire sets NomDestinataire field to given value.


### GetMontantHtTotal

`func (o *FactureEnrichieInfoInput) GetMontantHtTotal() MontantHtTotal`

GetMontantHtTotal returns the MontantHtTotal field if non-nil, zero value otherwise.

### GetMontantHtTotalOk

`func (o *FactureEnrichieInfoInput) GetMontantHtTotalOk() (*MontantHtTotal, bool)`

GetMontantHtTotalOk returns a tuple with the MontantHtTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantHtTotal

`func (o *FactureEnrichieInfoInput) SetMontantHtTotal(v MontantHtTotal)`

SetMontantHtTotal sets MontantHtTotal field to given value.


### GetMontantTva

`func (o *FactureEnrichieInfoInput) GetMontantTva() MontantTva`

GetMontantTva returns the MontantTva field if non-nil, zero value otherwise.

### GetMontantTvaOk

`func (o *FactureEnrichieInfoInput) GetMontantTvaOk() (*MontantTva, bool)`

GetMontantTvaOk returns a tuple with the MontantTva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantTva

`func (o *FactureEnrichieInfoInput) SetMontantTva(v MontantTva)`

SetMontantTva sets MontantTva field to given value.


### GetMontantTtcTotal

`func (o *FactureEnrichieInfoInput) GetMontantTtcTotal() MontantTtcTotal`

GetMontantTtcTotal returns the MontantTtcTotal field if non-nil, zero value otherwise.

### GetMontantTtcTotalOk

`func (o *FactureEnrichieInfoInput) GetMontantTtcTotalOk() (*MontantTtcTotal, bool)`

GetMontantTtcTotalOk returns a tuple with the MontantTtcTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantTtcTotal

`func (o *FactureEnrichieInfoInput) SetMontantTtcTotal(v MontantTtcTotal)`

SetMontantTtcTotal sets MontantTtcTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



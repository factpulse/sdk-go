# MontantTotal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MontantHtTotal** | **float64** | Montant total HT. | 
**MontantTva** | **float64** | Montant total de la TVA. | 
**MontantTtcTotal** | **float64** | Montant total TTC. | 
**MontantAPayer** | **float64** | Montant à payer. | 
**Acompte** | Pointer to **NullableFloat64** | Acompte versé. | [optional] 
**MontantRemiseGlobaleTtc** | Pointer to **float64** | Montant de la remise globale TTC. | [optional] 
**MotifRemiseGlobaleTtc** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMontantTotal

`func NewMontantTotal(montantHtTotal float64, montantTva float64, montantTtcTotal float64, montantAPayer float64, ) *MontantTotal`

NewMontantTotal instantiates a new MontantTotal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMontantTotalWithDefaults

`func NewMontantTotalWithDefaults() *MontantTotal`

NewMontantTotalWithDefaults instantiates a new MontantTotal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMontantHtTotal

`func (o *MontantTotal) GetMontantHtTotal() float64`

GetMontantHtTotal returns the MontantHtTotal field if non-nil, zero value otherwise.

### GetMontantHtTotalOk

`func (o *MontantTotal) GetMontantHtTotalOk() (*float64, bool)`

GetMontantHtTotalOk returns a tuple with the MontantHtTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantHtTotal

`func (o *MontantTotal) SetMontantHtTotal(v float64)`

SetMontantHtTotal sets MontantHtTotal field to given value.


### GetMontantTva

`func (o *MontantTotal) GetMontantTva() float64`

GetMontantTva returns the MontantTva field if non-nil, zero value otherwise.

### GetMontantTvaOk

`func (o *MontantTotal) GetMontantTvaOk() (*float64, bool)`

GetMontantTvaOk returns a tuple with the MontantTva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantTva

`func (o *MontantTotal) SetMontantTva(v float64)`

SetMontantTva sets MontantTva field to given value.


### GetMontantTtcTotal

`func (o *MontantTotal) GetMontantTtcTotal() float64`

GetMontantTtcTotal returns the MontantTtcTotal field if non-nil, zero value otherwise.

### GetMontantTtcTotalOk

`func (o *MontantTotal) GetMontantTtcTotalOk() (*float64, bool)`

GetMontantTtcTotalOk returns a tuple with the MontantTtcTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantTtcTotal

`func (o *MontantTotal) SetMontantTtcTotal(v float64)`

SetMontantTtcTotal sets MontantTtcTotal field to given value.


### GetMontantAPayer

`func (o *MontantTotal) GetMontantAPayer() float64`

GetMontantAPayer returns the MontantAPayer field if non-nil, zero value otherwise.

### GetMontantAPayerOk

`func (o *MontantTotal) GetMontantAPayerOk() (*float64, bool)`

GetMontantAPayerOk returns a tuple with the MontantAPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantAPayer

`func (o *MontantTotal) SetMontantAPayer(v float64)`

SetMontantAPayer sets MontantAPayer field to given value.


### GetAcompte

`func (o *MontantTotal) GetAcompte() float64`

GetAcompte returns the Acompte field if non-nil, zero value otherwise.

### GetAcompteOk

`func (o *MontantTotal) GetAcompteOk() (*float64, bool)`

GetAcompteOk returns a tuple with the Acompte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcompte

`func (o *MontantTotal) SetAcompte(v float64)`

SetAcompte sets Acompte field to given value.

### HasAcompte

`func (o *MontantTotal) HasAcompte() bool`

HasAcompte returns a boolean if a field has been set.

### SetAcompteNil

`func (o *MontantTotal) SetAcompteNil(b bool)`

 SetAcompteNil sets the value for Acompte to be an explicit nil

### UnsetAcompte
`func (o *MontantTotal) UnsetAcompte()`

UnsetAcompte ensures that no value is present for Acompte, not even an explicit nil
### GetMontantRemiseGlobaleTtc

`func (o *MontantTotal) GetMontantRemiseGlobaleTtc() float64`

GetMontantRemiseGlobaleTtc returns the MontantRemiseGlobaleTtc field if non-nil, zero value otherwise.

### GetMontantRemiseGlobaleTtcOk

`func (o *MontantTotal) GetMontantRemiseGlobaleTtcOk() (*float64, bool)`

GetMontantRemiseGlobaleTtcOk returns a tuple with the MontantRemiseGlobaleTtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantRemiseGlobaleTtc

`func (o *MontantTotal) SetMontantRemiseGlobaleTtc(v float64)`

SetMontantRemiseGlobaleTtc sets MontantRemiseGlobaleTtc field to given value.

### HasMontantRemiseGlobaleTtc

`func (o *MontantTotal) HasMontantRemiseGlobaleTtc() bool`

HasMontantRemiseGlobaleTtc returns a boolean if a field has been set.

### GetMotifRemiseGlobaleTtc

`func (o *MontantTotal) GetMotifRemiseGlobaleTtc() string`

GetMotifRemiseGlobaleTtc returns the MotifRemiseGlobaleTtc field if non-nil, zero value otherwise.

### GetMotifRemiseGlobaleTtcOk

`func (o *MontantTotal) GetMotifRemiseGlobaleTtcOk() (*string, bool)`

GetMotifRemiseGlobaleTtcOk returns a tuple with the MotifRemiseGlobaleTtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotifRemiseGlobaleTtc

`func (o *MontantTotal) SetMotifRemiseGlobaleTtc(v string)`

SetMotifRemiseGlobaleTtc sets MotifRemiseGlobaleTtc field to given value.

### HasMotifRemiseGlobaleTtc

`func (o *MontantTotal) HasMotifRemiseGlobaleTtc() bool`

HasMotifRemiseGlobaleTtc returns a boolean if a field has been set.

### SetMotifRemiseGlobaleTtcNil

`func (o *MontantTotal) SetMotifRemiseGlobaleTtcNil(b bool)`

 SetMotifRemiseGlobaleTtcNil sets the value for MotifRemiseGlobaleTtc to be an explicit nil

### UnsetMotifRemiseGlobaleTtc
`func (o *MontantTotal) UnsetMotifRemiseGlobaleTtc()`

UnsetMotifRemiseGlobaleTtc ensures that no value is present for MotifRemiseGlobaleTtc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



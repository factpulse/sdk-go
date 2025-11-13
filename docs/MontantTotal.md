# MontantTotal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MontantHtTotal** | [**Montanthttotal**](Montanthttotal.md) |  | 
**MontantTva** | [**Montanttva1**](Montanttva1.md) |  | 
**MontantTtcTotal** | [**Montantttctotal**](Montantttctotal.md) |  | 
**MontantAPayer** | [**Montantapayer**](Montantapayer.md) |  | 
**Acompte** | Pointer to [**NullableMontantTotalAcompte**](MontantTotalAcompte.md) |  | [optional] 
**MontantRemiseGlobaleTtc** | Pointer to [**NullableMontantTotalMontantRemiseGlobaleTtc**](MontantTotalMontantRemiseGlobaleTtc.md) |  | [optional] 
**MotifRemiseGlobaleTtc** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMontantTotal

`func NewMontantTotal(montantHtTotal Montanthttotal, montantTva Montanttva1, montantTtcTotal Montantttctotal, montantAPayer Montantapayer, ) *MontantTotal`

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

`func (o *MontantTotal) GetMontantHtTotal() Montanthttotal`

GetMontantHtTotal returns the MontantHtTotal field if non-nil, zero value otherwise.

### GetMontantHtTotalOk

`func (o *MontantTotal) GetMontantHtTotalOk() (*Montanthttotal, bool)`

GetMontantHtTotalOk returns a tuple with the MontantHtTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantHtTotal

`func (o *MontantTotal) SetMontantHtTotal(v Montanthttotal)`

SetMontantHtTotal sets MontantHtTotal field to given value.


### GetMontantTva

`func (o *MontantTotal) GetMontantTva() Montanttva1`

GetMontantTva returns the MontantTva field if non-nil, zero value otherwise.

### GetMontantTvaOk

`func (o *MontantTotal) GetMontantTvaOk() (*Montanttva1, bool)`

GetMontantTvaOk returns a tuple with the MontantTva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantTva

`func (o *MontantTotal) SetMontantTva(v Montanttva1)`

SetMontantTva sets MontantTva field to given value.


### GetMontantTtcTotal

`func (o *MontantTotal) GetMontantTtcTotal() Montantttctotal`

GetMontantTtcTotal returns the MontantTtcTotal field if non-nil, zero value otherwise.

### GetMontantTtcTotalOk

`func (o *MontantTotal) GetMontantTtcTotalOk() (*Montantttctotal, bool)`

GetMontantTtcTotalOk returns a tuple with the MontantTtcTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantTtcTotal

`func (o *MontantTotal) SetMontantTtcTotal(v Montantttctotal)`

SetMontantTtcTotal sets MontantTtcTotal field to given value.


### GetMontantAPayer

`func (o *MontantTotal) GetMontantAPayer() Montantapayer`

GetMontantAPayer returns the MontantAPayer field if non-nil, zero value otherwise.

### GetMontantAPayerOk

`func (o *MontantTotal) GetMontantAPayerOk() (*Montantapayer, bool)`

GetMontantAPayerOk returns a tuple with the MontantAPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantAPayer

`func (o *MontantTotal) SetMontantAPayer(v Montantapayer)`

SetMontantAPayer sets MontantAPayer field to given value.


### GetAcompte

`func (o *MontantTotal) GetAcompte() MontantTotalAcompte`

GetAcompte returns the Acompte field if non-nil, zero value otherwise.

### GetAcompteOk

`func (o *MontantTotal) GetAcompteOk() (*MontantTotalAcompte, bool)`

GetAcompteOk returns a tuple with the Acompte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcompte

`func (o *MontantTotal) SetAcompte(v MontantTotalAcompte)`

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

`func (o *MontantTotal) GetMontantRemiseGlobaleTtc() MontantTotalMontantRemiseGlobaleTtc`

GetMontantRemiseGlobaleTtc returns the MontantRemiseGlobaleTtc field if non-nil, zero value otherwise.

### GetMontantRemiseGlobaleTtcOk

`func (o *MontantTotal) GetMontantRemiseGlobaleTtcOk() (*MontantTotalMontantRemiseGlobaleTtc, bool)`

GetMontantRemiseGlobaleTtcOk returns a tuple with the MontantRemiseGlobaleTtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantRemiseGlobaleTtc

`func (o *MontantTotal) SetMontantRemiseGlobaleTtc(v MontantTotalMontantRemiseGlobaleTtc)`

SetMontantRemiseGlobaleTtc sets MontantRemiseGlobaleTtc field to given value.

### HasMontantRemiseGlobaleTtc

`func (o *MontantTotal) HasMontantRemiseGlobaleTtc() bool`

HasMontantRemiseGlobaleTtc returns a boolean if a field has been set.

### SetMontantRemiseGlobaleTtcNil

`func (o *MontantTotal) SetMontantRemiseGlobaleTtcNil(b bool)`

 SetMontantRemiseGlobaleTtcNil sets the value for MontantRemiseGlobaleTtc to be an explicit nil

### UnsetMontantRemiseGlobaleTtc
`func (o *MontantTotal) UnsetMontantRemiseGlobaleTtc()`

UnsetMontantRemiseGlobaleTtc ensures that no value is present for MontantRemiseGlobaleTtc, not even an explicit nil
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



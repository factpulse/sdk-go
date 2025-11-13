# InformationSignatureAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NomChamp** | **string** | Nom du champ de signature dans le PDF | 
**Signataire** | Pointer to **NullableString** |  | [optional] 
**DateSignature** | Pointer to **NullableString** |  | [optional] 
**Raison** | Pointer to **NullableString** |  | [optional] 
**Localisation** | Pointer to **NullableString** |  | [optional] 
**EstValide** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewInformationSignatureAPI

`func NewInformationSignatureAPI(nomChamp string, ) *InformationSignatureAPI`

NewInformationSignatureAPI instantiates a new InformationSignatureAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInformationSignatureAPIWithDefaults

`func NewInformationSignatureAPIWithDefaults() *InformationSignatureAPI`

NewInformationSignatureAPIWithDefaults instantiates a new InformationSignatureAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNomChamp

`func (o *InformationSignatureAPI) GetNomChamp() string`

GetNomChamp returns the NomChamp field if non-nil, zero value otherwise.

### GetNomChampOk

`func (o *InformationSignatureAPI) GetNomChampOk() (*string, bool)`

GetNomChampOk returns a tuple with the NomChamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomChamp

`func (o *InformationSignatureAPI) SetNomChamp(v string)`

SetNomChamp sets NomChamp field to given value.


### GetSignataire

`func (o *InformationSignatureAPI) GetSignataire() string`

GetSignataire returns the Signataire field if non-nil, zero value otherwise.

### GetSignataireOk

`func (o *InformationSignatureAPI) GetSignataireOk() (*string, bool)`

GetSignataireOk returns a tuple with the Signataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignataire

`func (o *InformationSignatureAPI) SetSignataire(v string)`

SetSignataire sets Signataire field to given value.

### HasSignataire

`func (o *InformationSignatureAPI) HasSignataire() bool`

HasSignataire returns a boolean if a field has been set.

### SetSignataireNil

`func (o *InformationSignatureAPI) SetSignataireNil(b bool)`

 SetSignataireNil sets the value for Signataire to be an explicit nil

### UnsetSignataire
`func (o *InformationSignatureAPI) UnsetSignataire()`

UnsetSignataire ensures that no value is present for Signataire, not even an explicit nil
### GetDateSignature

`func (o *InformationSignatureAPI) GetDateSignature() string`

GetDateSignature returns the DateSignature field if non-nil, zero value otherwise.

### GetDateSignatureOk

`func (o *InformationSignatureAPI) GetDateSignatureOk() (*string, bool)`

GetDateSignatureOk returns a tuple with the DateSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateSignature

`func (o *InformationSignatureAPI) SetDateSignature(v string)`

SetDateSignature sets DateSignature field to given value.

### HasDateSignature

`func (o *InformationSignatureAPI) HasDateSignature() bool`

HasDateSignature returns a boolean if a field has been set.

### SetDateSignatureNil

`func (o *InformationSignatureAPI) SetDateSignatureNil(b bool)`

 SetDateSignatureNil sets the value for DateSignature to be an explicit nil

### UnsetDateSignature
`func (o *InformationSignatureAPI) UnsetDateSignature()`

UnsetDateSignature ensures that no value is present for DateSignature, not even an explicit nil
### GetRaison

`func (o *InformationSignatureAPI) GetRaison() string`

GetRaison returns the Raison field if non-nil, zero value otherwise.

### GetRaisonOk

`func (o *InformationSignatureAPI) GetRaisonOk() (*string, bool)`

GetRaisonOk returns a tuple with the Raison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaison

`func (o *InformationSignatureAPI) SetRaison(v string)`

SetRaison sets Raison field to given value.

### HasRaison

`func (o *InformationSignatureAPI) HasRaison() bool`

HasRaison returns a boolean if a field has been set.

### SetRaisonNil

`func (o *InformationSignatureAPI) SetRaisonNil(b bool)`

 SetRaisonNil sets the value for Raison to be an explicit nil

### UnsetRaison
`func (o *InformationSignatureAPI) UnsetRaison()`

UnsetRaison ensures that no value is present for Raison, not even an explicit nil
### GetLocalisation

`func (o *InformationSignatureAPI) GetLocalisation() string`

GetLocalisation returns the Localisation field if non-nil, zero value otherwise.

### GetLocalisationOk

`func (o *InformationSignatureAPI) GetLocalisationOk() (*string, bool)`

GetLocalisationOk returns a tuple with the Localisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalisation

`func (o *InformationSignatureAPI) SetLocalisation(v string)`

SetLocalisation sets Localisation field to given value.

### HasLocalisation

`func (o *InformationSignatureAPI) HasLocalisation() bool`

HasLocalisation returns a boolean if a field has been set.

### SetLocalisationNil

`func (o *InformationSignatureAPI) SetLocalisationNil(b bool)`

 SetLocalisationNil sets the value for Localisation to be an explicit nil

### UnsetLocalisation
`func (o *InformationSignatureAPI) UnsetLocalisation()`

UnsetLocalisation ensures that no value is present for Localisation, not even an explicit nil
### GetEstValide

`func (o *InformationSignatureAPI) GetEstValide() bool`

GetEstValide returns the EstValide field if non-nil, zero value otherwise.

### GetEstValideOk

`func (o *InformationSignatureAPI) GetEstValideOk() (*bool, bool)`

GetEstValideOk returns a tuple with the EstValide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstValide

`func (o *InformationSignatureAPI) SetEstValide(v bool)`

SetEstValide sets EstValide field to given value.

### HasEstValide

`func (o *InformationSignatureAPI) HasEstValide() bool`

HasEstValide returns a boolean if a field has been set.

### SetEstValideNil

`func (o *InformationSignatureAPI) SetEstValideNil(b bool)`

 SetEstValideNil sets the value for EstValide to be an explicit nil

### UnsetEstValide
`func (o *InformationSignatureAPI) UnsetEstValide()`

UnsetEstValide ensures that no value is present for EstValide, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



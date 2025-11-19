# OptionsProcessing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfilFacturx** | Pointer to [**ProfilAPI**](ProfilAPI.md) | Profil Factur-X à utiliser | [optional] [default to EN16931]
**AutoEnrichir** | Pointer to **bool** | Auto-enrichir les données (APIs Entreprises, Chorus Pro, etc.) | [optional] [default to true]
**Valider** | Pointer to **bool** | Valider le XML Factur-X avec Schematron | [optional] [default to true]
**VerifierParametresDestination** | Pointer to **bool** | Vérifier les paramètres requis par la destination (ex: code_service pour Chorus) | [optional] [default to true]

## Methods

### NewOptionsProcessing

`func NewOptionsProcessing() *OptionsProcessing`

NewOptionsProcessing instantiates a new OptionsProcessing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionsProcessingWithDefaults

`func NewOptionsProcessingWithDefaults() *OptionsProcessing`

NewOptionsProcessingWithDefaults instantiates a new OptionsProcessing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfilFacturx

`func (o *OptionsProcessing) GetProfilFacturx() ProfilAPI`

GetProfilFacturx returns the ProfilFacturx field if non-nil, zero value otherwise.

### GetProfilFacturxOk

`func (o *OptionsProcessing) GetProfilFacturxOk() (*ProfilAPI, bool)`

GetProfilFacturxOk returns a tuple with the ProfilFacturx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilFacturx

`func (o *OptionsProcessing) SetProfilFacturx(v ProfilAPI)`

SetProfilFacturx sets ProfilFacturx field to given value.

### HasProfilFacturx

`func (o *OptionsProcessing) HasProfilFacturx() bool`

HasProfilFacturx returns a boolean if a field has been set.

### GetAutoEnrichir

`func (o *OptionsProcessing) GetAutoEnrichir() bool`

GetAutoEnrichir returns the AutoEnrichir field if non-nil, zero value otherwise.

### GetAutoEnrichirOk

`func (o *OptionsProcessing) GetAutoEnrichirOk() (*bool, bool)`

GetAutoEnrichirOk returns a tuple with the AutoEnrichir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoEnrichir

`func (o *OptionsProcessing) SetAutoEnrichir(v bool)`

SetAutoEnrichir sets AutoEnrichir field to given value.

### HasAutoEnrichir

`func (o *OptionsProcessing) HasAutoEnrichir() bool`

HasAutoEnrichir returns a boolean if a field has been set.

### GetValider

`func (o *OptionsProcessing) GetValider() bool`

GetValider returns the Valider field if non-nil, zero value otherwise.

### GetValiderOk

`func (o *OptionsProcessing) GetValiderOk() (*bool, bool)`

GetValiderOk returns a tuple with the Valider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValider

`func (o *OptionsProcessing) SetValider(v bool)`

SetValider sets Valider field to given value.

### HasValider

`func (o *OptionsProcessing) HasValider() bool`

HasValider returns a boolean if a field has been set.

### GetVerifierParametresDestination

`func (o *OptionsProcessing) GetVerifierParametresDestination() bool`

GetVerifierParametresDestination returns the VerifierParametresDestination field if non-nil, zero value otherwise.

### GetVerifierParametresDestinationOk

`func (o *OptionsProcessing) GetVerifierParametresDestinationOk() (*bool, bool)`

GetVerifierParametresDestinationOk returns a tuple with the VerifierParametresDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifierParametresDestination

`func (o *OptionsProcessing) SetVerifierParametresDestination(v bool)`

SetVerifierParametresDestination sets VerifierParametresDestination field to given value.

### HasVerifierParametresDestination

`func (o *OptionsProcessing) HasVerifierParametresDestination() bool`

HasVerifierParametresDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



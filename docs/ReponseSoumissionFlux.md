# ReponseSoumissionFlux

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowId** | **string** | Identifiant unique du flux généré par la PDP | 
**TrackingId** | Pointer to **NullableString** |  | [optional] 
**Nom** | **string** | Nom du flux | 
**SyntaxeFlux** | **string** | Syntaxe du flux (CII, UBL, etc.) | 
**ProfilFlux** | Pointer to **NullableString** |  | [optional] 
**Sha256** | **string** | Hash SHA256 du fichier soumis | 
**Message** | **string** | Message de confirmation | 

## Methods

### NewReponseSoumissionFlux

`func NewReponseSoumissionFlux(flowId string, nom string, syntaxeFlux string, sha256 string, message string, ) *ReponseSoumissionFlux`

NewReponseSoumissionFlux instantiates a new ReponseSoumissionFlux object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReponseSoumissionFluxWithDefaults

`func NewReponseSoumissionFluxWithDefaults() *ReponseSoumissionFlux`

NewReponseSoumissionFluxWithDefaults instantiates a new ReponseSoumissionFlux object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowId

`func (o *ReponseSoumissionFlux) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *ReponseSoumissionFlux) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *ReponseSoumissionFlux) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.


### GetTrackingId

`func (o *ReponseSoumissionFlux) GetTrackingId() string`

GetTrackingId returns the TrackingId field if non-nil, zero value otherwise.

### GetTrackingIdOk

`func (o *ReponseSoumissionFlux) GetTrackingIdOk() (*string, bool)`

GetTrackingIdOk returns a tuple with the TrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingId

`func (o *ReponseSoumissionFlux) SetTrackingId(v string)`

SetTrackingId sets TrackingId field to given value.

### HasTrackingId

`func (o *ReponseSoumissionFlux) HasTrackingId() bool`

HasTrackingId returns a boolean if a field has been set.

### SetTrackingIdNil

`func (o *ReponseSoumissionFlux) SetTrackingIdNil(b bool)`

 SetTrackingIdNil sets the value for TrackingId to be an explicit nil

### UnsetTrackingId
`func (o *ReponseSoumissionFlux) UnsetTrackingId()`

UnsetTrackingId ensures that no value is present for TrackingId, not even an explicit nil
### GetNom

`func (o *ReponseSoumissionFlux) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *ReponseSoumissionFlux) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *ReponseSoumissionFlux) SetNom(v string)`

SetNom sets Nom field to given value.


### GetSyntaxeFlux

`func (o *ReponseSoumissionFlux) GetSyntaxeFlux() string`

GetSyntaxeFlux returns the SyntaxeFlux field if non-nil, zero value otherwise.

### GetSyntaxeFluxOk

`func (o *ReponseSoumissionFlux) GetSyntaxeFluxOk() (*string, bool)`

GetSyntaxeFluxOk returns a tuple with the SyntaxeFlux field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntaxeFlux

`func (o *ReponseSoumissionFlux) SetSyntaxeFlux(v string)`

SetSyntaxeFlux sets SyntaxeFlux field to given value.


### GetProfilFlux

`func (o *ReponseSoumissionFlux) GetProfilFlux() string`

GetProfilFlux returns the ProfilFlux field if non-nil, zero value otherwise.

### GetProfilFluxOk

`func (o *ReponseSoumissionFlux) GetProfilFluxOk() (*string, bool)`

GetProfilFluxOk returns a tuple with the ProfilFlux field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilFlux

`func (o *ReponseSoumissionFlux) SetProfilFlux(v string)`

SetProfilFlux sets ProfilFlux field to given value.

### HasProfilFlux

`func (o *ReponseSoumissionFlux) HasProfilFlux() bool`

HasProfilFlux returns a boolean if a field has been set.

### SetProfilFluxNil

`func (o *ReponseSoumissionFlux) SetProfilFluxNil(b bool)`

 SetProfilFluxNil sets the value for ProfilFlux to be an explicit nil

### UnsetProfilFlux
`func (o *ReponseSoumissionFlux) UnsetProfilFlux()`

UnsetProfilFlux ensures that no value is present for ProfilFlux, not even an explicit nil
### GetSha256

`func (o *ReponseSoumissionFlux) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ReponseSoumissionFlux) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ReponseSoumissionFlux) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.


### GetMessage

`func (o *ReponseSoumissionFlux) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ReponseSoumissionFlux) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ReponseSoumissionFlux) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



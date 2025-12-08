# RequeteSoumissionFlux

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NomFlux** | **string** | Nom du flux (ex: &#39;Facture 2025-001&#39;) | 
**SyntaxeFlux** | Pointer to [**SyntaxeFlux**](SyntaxeFlux.md) | Syntaxe du flux (CII pour Factur-X) | [optional] [default to CII]
**ProfilFlux** | Pointer to [**NullableProfilFlux**](ProfilFlux.md) |  | [optional] 
**TrackingId** | Pointer to **NullableString** |  | [optional] 
**RequestId** | Pointer to **NullableString** |  | [optional] 
**PdpCredentials** | Pointer to [**NullablePDPCredentials**](PDPCredentials.md) |  | [optional] 

## Methods

### NewRequeteSoumissionFlux

`func NewRequeteSoumissionFlux(nomFlux string, ) *RequeteSoumissionFlux`

NewRequeteSoumissionFlux instantiates a new RequeteSoumissionFlux object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequeteSoumissionFluxWithDefaults

`func NewRequeteSoumissionFluxWithDefaults() *RequeteSoumissionFlux`

NewRequeteSoumissionFluxWithDefaults instantiates a new RequeteSoumissionFlux object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNomFlux

`func (o *RequeteSoumissionFlux) GetNomFlux() string`

GetNomFlux returns the NomFlux field if non-nil, zero value otherwise.

### GetNomFluxOk

`func (o *RequeteSoumissionFlux) GetNomFluxOk() (*string, bool)`

GetNomFluxOk returns a tuple with the NomFlux field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomFlux

`func (o *RequeteSoumissionFlux) SetNomFlux(v string)`

SetNomFlux sets NomFlux field to given value.


### GetSyntaxeFlux

`func (o *RequeteSoumissionFlux) GetSyntaxeFlux() SyntaxeFlux`

GetSyntaxeFlux returns the SyntaxeFlux field if non-nil, zero value otherwise.

### GetSyntaxeFluxOk

`func (o *RequeteSoumissionFlux) GetSyntaxeFluxOk() (*SyntaxeFlux, bool)`

GetSyntaxeFluxOk returns a tuple with the SyntaxeFlux field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntaxeFlux

`func (o *RequeteSoumissionFlux) SetSyntaxeFlux(v SyntaxeFlux)`

SetSyntaxeFlux sets SyntaxeFlux field to given value.

### HasSyntaxeFlux

`func (o *RequeteSoumissionFlux) HasSyntaxeFlux() bool`

HasSyntaxeFlux returns a boolean if a field has been set.

### GetProfilFlux

`func (o *RequeteSoumissionFlux) GetProfilFlux() ProfilFlux`

GetProfilFlux returns the ProfilFlux field if non-nil, zero value otherwise.

### GetProfilFluxOk

`func (o *RequeteSoumissionFlux) GetProfilFluxOk() (*ProfilFlux, bool)`

GetProfilFluxOk returns a tuple with the ProfilFlux field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilFlux

`func (o *RequeteSoumissionFlux) SetProfilFlux(v ProfilFlux)`

SetProfilFlux sets ProfilFlux field to given value.

### HasProfilFlux

`func (o *RequeteSoumissionFlux) HasProfilFlux() bool`

HasProfilFlux returns a boolean if a field has been set.

### SetProfilFluxNil

`func (o *RequeteSoumissionFlux) SetProfilFluxNil(b bool)`

 SetProfilFluxNil sets the value for ProfilFlux to be an explicit nil

### UnsetProfilFlux
`func (o *RequeteSoumissionFlux) UnsetProfilFlux()`

UnsetProfilFlux ensures that no value is present for ProfilFlux, not even an explicit nil
### GetTrackingId

`func (o *RequeteSoumissionFlux) GetTrackingId() string`

GetTrackingId returns the TrackingId field if non-nil, zero value otherwise.

### GetTrackingIdOk

`func (o *RequeteSoumissionFlux) GetTrackingIdOk() (*string, bool)`

GetTrackingIdOk returns a tuple with the TrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingId

`func (o *RequeteSoumissionFlux) SetTrackingId(v string)`

SetTrackingId sets TrackingId field to given value.

### HasTrackingId

`func (o *RequeteSoumissionFlux) HasTrackingId() bool`

HasTrackingId returns a boolean if a field has been set.

### SetTrackingIdNil

`func (o *RequeteSoumissionFlux) SetTrackingIdNil(b bool)`

 SetTrackingIdNil sets the value for TrackingId to be an explicit nil

### UnsetTrackingId
`func (o *RequeteSoumissionFlux) UnsetTrackingId()`

UnsetTrackingId ensures that no value is present for TrackingId, not even an explicit nil
### GetRequestId

`func (o *RequeteSoumissionFlux) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *RequeteSoumissionFlux) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *RequeteSoumissionFlux) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *RequeteSoumissionFlux) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestIdNil

`func (o *RequeteSoumissionFlux) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *RequeteSoumissionFlux) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetPdpCredentials

`func (o *RequeteSoumissionFlux) GetPdpCredentials() PDPCredentials`

GetPdpCredentials returns the PdpCredentials field if non-nil, zero value otherwise.

### GetPdpCredentialsOk

`func (o *RequeteSoumissionFlux) GetPdpCredentialsOk() (*PDPCredentials, bool)`

GetPdpCredentialsOk returns a tuple with the PdpCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpCredentials

`func (o *RequeteSoumissionFlux) SetPdpCredentials(v PDPCredentials)`

SetPdpCredentials sets PdpCredentials field to given value.

### HasPdpCredentials

`func (o *RequeteSoumissionFlux) HasPdpCredentials() bool`

HasPdpCredentials returns a boolean if a field has been set.

### SetPdpCredentialsNil

`func (o *RequeteSoumissionFlux) SetPdpCredentialsNil(b bool)`

 SetPdpCredentialsNil sets the value for PdpCredentials to be an explicit nil

### UnsetPdpCredentials
`func (o *RequeteSoumissionFlux) UnsetPdpCredentials()`

UnsetPdpCredentials ensures that no value is present for PdpCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



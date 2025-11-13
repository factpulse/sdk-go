# SoumettreFactureCompleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DonneesFacture** | [**DonneesFactureSimplifiees**](DonneesFactureSimplifiees.md) | Données de la facture au format simplifié (voir exemples) | 
**PdfSource** | **string** | PDF source encodé en base64 (sera transformé en Factur-X) | 
**Destination** | [**Destination**](Destination.md) |  | 
**Signature** | Pointer to [**NullableParametresSignature**](ParametresSignature.md) |  | [optional] 
**Options** | Pointer to [**OptionsProcessing**](OptionsProcessing.md) | Options de traitement | [optional] 

## Methods

### NewSoumettreFactureCompleteRequest

`func NewSoumettreFactureCompleteRequest(donneesFacture DonneesFactureSimplifiees, pdfSource string, destination Destination, ) *SoumettreFactureCompleteRequest`

NewSoumettreFactureCompleteRequest instantiates a new SoumettreFactureCompleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoumettreFactureCompleteRequestWithDefaults

`func NewSoumettreFactureCompleteRequestWithDefaults() *SoumettreFactureCompleteRequest`

NewSoumettreFactureCompleteRequestWithDefaults instantiates a new SoumettreFactureCompleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDonneesFacture

`func (o *SoumettreFactureCompleteRequest) GetDonneesFacture() DonneesFactureSimplifiees`

GetDonneesFacture returns the DonneesFacture field if non-nil, zero value otherwise.

### GetDonneesFactureOk

`func (o *SoumettreFactureCompleteRequest) GetDonneesFactureOk() (*DonneesFactureSimplifiees, bool)`

GetDonneesFactureOk returns a tuple with the DonneesFacture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDonneesFacture

`func (o *SoumettreFactureCompleteRequest) SetDonneesFacture(v DonneesFactureSimplifiees)`

SetDonneesFacture sets DonneesFacture field to given value.


### GetPdfSource

`func (o *SoumettreFactureCompleteRequest) GetPdfSource() string`

GetPdfSource returns the PdfSource field if non-nil, zero value otherwise.

### GetPdfSourceOk

`func (o *SoumettreFactureCompleteRequest) GetPdfSourceOk() (*string, bool)`

GetPdfSourceOk returns a tuple with the PdfSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfSource

`func (o *SoumettreFactureCompleteRequest) SetPdfSource(v string)`

SetPdfSource sets PdfSource field to given value.


### GetDestination

`func (o *SoumettreFactureCompleteRequest) GetDestination() Destination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *SoumettreFactureCompleteRequest) GetDestinationOk() (*Destination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *SoumettreFactureCompleteRequest) SetDestination(v Destination)`

SetDestination sets Destination field to given value.


### GetSignature

`func (o *SoumettreFactureCompleteRequest) GetSignature() ParametresSignature`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *SoumettreFactureCompleteRequest) GetSignatureOk() (*ParametresSignature, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *SoumettreFactureCompleteRequest) SetSignature(v ParametresSignature)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *SoumettreFactureCompleteRequest) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### SetSignatureNil

`func (o *SoumettreFactureCompleteRequest) SetSignatureNil(b bool)`

 SetSignatureNil sets the value for Signature to be an explicit nil

### UnsetSignature
`func (o *SoumettreFactureCompleteRequest) UnsetSignature()`

UnsetSignature ensures that no value is present for Signature, not even an explicit nil
### GetOptions

`func (o *SoumettreFactureCompleteRequest) GetOptions() OptionsProcessing`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SoumettreFactureCompleteRequest) GetOptionsOk() (*OptionsProcessing, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SoumettreFactureCompleteRequest) SetOptions(v OptionsProcessing)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SoumettreFactureCompleteRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



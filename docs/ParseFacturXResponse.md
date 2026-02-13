# ParseFacturXResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Parsing status | 
**Invoice** | **map[string]interface{}** | Parsed invoice data. For CII/Factur-X: FacturXInvoice format (round-trip with /generate-invoice). For UBL: IncomingInvoice format (summary extraction). | 
**DetectedFormat** | Pointer to **NullableString** |  | [optional] 
**DetectedProfile** | Pointer to **NullableString** |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewParseFacturXResponse

`func NewParseFacturXResponse(status string, invoice map[string]interface{}, ) *ParseFacturXResponse`

NewParseFacturXResponse instantiates a new ParseFacturXResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParseFacturXResponseWithDefaults

`func NewParseFacturXResponseWithDefaults() *ParseFacturXResponse`

NewParseFacturXResponseWithDefaults instantiates a new ParseFacturXResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ParseFacturXResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ParseFacturXResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ParseFacturXResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetInvoice

`func (o *ParseFacturXResponse) GetInvoice() map[string]interface{}`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *ParseFacturXResponse) GetInvoiceOk() (*map[string]interface{}, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *ParseFacturXResponse) SetInvoice(v map[string]interface{})`

SetInvoice sets Invoice field to given value.


### GetDetectedFormat

`func (o *ParseFacturXResponse) GetDetectedFormat() string`

GetDetectedFormat returns the DetectedFormat field if non-nil, zero value otherwise.

### GetDetectedFormatOk

`func (o *ParseFacturXResponse) GetDetectedFormatOk() (*string, bool)`

GetDetectedFormatOk returns a tuple with the DetectedFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedFormat

`func (o *ParseFacturXResponse) SetDetectedFormat(v string)`

SetDetectedFormat sets DetectedFormat field to given value.

### HasDetectedFormat

`func (o *ParseFacturXResponse) HasDetectedFormat() bool`

HasDetectedFormat returns a boolean if a field has been set.

### SetDetectedFormatNil

`func (o *ParseFacturXResponse) SetDetectedFormatNil(b bool)`

 SetDetectedFormatNil sets the value for DetectedFormat to be an explicit nil

### UnsetDetectedFormat
`func (o *ParseFacturXResponse) UnsetDetectedFormat()`

UnsetDetectedFormat ensures that no value is present for DetectedFormat, not even an explicit nil
### GetDetectedProfile

`func (o *ParseFacturXResponse) GetDetectedProfile() string`

GetDetectedProfile returns the DetectedProfile field if non-nil, zero value otherwise.

### GetDetectedProfileOk

`func (o *ParseFacturXResponse) GetDetectedProfileOk() (*string, bool)`

GetDetectedProfileOk returns a tuple with the DetectedProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedProfile

`func (o *ParseFacturXResponse) SetDetectedProfile(v string)`

SetDetectedProfile sets DetectedProfile field to given value.

### HasDetectedProfile

`func (o *ParseFacturXResponse) HasDetectedProfile() bool`

HasDetectedProfile returns a boolean if a field has been set.

### SetDetectedProfileNil

`func (o *ParseFacturXResponse) SetDetectedProfileNil(b bool)`

 SetDetectedProfileNil sets the value for DetectedProfile to be an explicit nil

### UnsetDetectedProfile
`func (o *ParseFacturXResponse) UnsetDetectedProfile()`

UnsetDetectedProfile ensures that no value is present for DetectedProfile, not even an explicit nil
### GetError

`func (o *ParseFacturXResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ParseFacturXResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ParseFacturXResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ParseFacturXResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ParseFacturXResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ParseFacturXResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



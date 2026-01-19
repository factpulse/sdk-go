# FactureElectroniqueRestApiSchemasCdarValidationErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Champ concerné | 
**Message** | **string** | Message d&#39;erreur | 
**Rule** | Pointer to **NullableString** |  | [optional] 
**Severity** | Pointer to **string** | Sévérité (error/warning) | [optional] [default to "error"]

## Methods

### NewFactureElectroniqueRestApiSchemasCdarValidationErrorResponse

`func NewFactureElectroniqueRestApiSchemasCdarValidationErrorResponse(field string, message string, ) *FactureElectroniqueRestApiSchemasCdarValidationErrorResponse`

NewFactureElectroniqueRestApiSchemasCdarValidationErrorResponse instantiates a new FactureElectroniqueRestApiSchemasCdarValidationErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFactureElectroniqueRestApiSchemasCdarValidationErrorResponseWithDefaults

`func NewFactureElectroniqueRestApiSchemasCdarValidationErrorResponseWithDefaults() *FactureElectroniqueRestApiSchemasCdarValidationErrorResponse`

NewFactureElectroniqueRestApiSchemasCdarValidationErrorResponseWithDefaults instantiates a new FactureElectroniqueRestApiSchemasCdarValidationErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *FactureElectroniqueRestApiSchemasCdarValidationErrorResponse) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FactureElectroniqueRestApiSchemasCdarValidationErrorResponse) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FactureElectroniqueRestApiSchemasCdarValidationErrorResponse) SetField(v string)`

SetField sets Field field to given value.


### GetMessage

`func (o *FactureElectroniqueRestApiSchemasCdarValidationErrorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FactureElectroniqueRestApiSchemasCdarValidationErrorResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FactureElectroniqueRestApiSchemasCdarValidationErrorResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRule

`func (o *FactureElectroniqueRestApiSchemasCdarValidationErrorResponse) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *FactureElectroniqueRestApiSchemasCdarValidationErrorResponse) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *FactureElectroniqueRestApiSchemasCdarValidationErrorResponse) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *FactureElectroniqueRestApiSchemasCdarValidationErrorResponse) HasRule() bool`

HasRule returns a boolean if a field has been set.

### SetRuleNil

`func (o *FactureElectroniqueRestApiSchemasCdarValidationErrorResponse) SetRuleNil(b bool)`

 SetRuleNil sets the value for Rule to be an explicit nil

### UnsetRule
`func (o *FactureElectroniqueRestApiSchemasCdarValidationErrorResponse) UnsetRule()`

UnsetRule ensures that no value is present for Rule, not even an explicit nil
### GetSeverity

`func (o *FactureElectroniqueRestApiSchemasCdarValidationErrorResponse) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *FactureElectroniqueRestApiSchemasCdarValidationErrorResponse) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *FactureElectroniqueRestApiSchemasCdarValidationErrorResponse) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *FactureElectroniqueRestApiSchemasCdarValidationErrorResponse) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



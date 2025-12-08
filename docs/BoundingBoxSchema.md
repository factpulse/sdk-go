# BoundingBoxSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X0** | **float32** | Coordonnée X gauche | 
**Y0** | **float32** | Coordonnée Y bas | 
**X1** | **float32** | Coordonnée X droite | 
**Y1** | **float32** | Coordonnée Y haut | 
**Page** | Pointer to **int32** | Numéro de page (0-indexed) | [optional] [default to 0]
**Width** | **float32** | Largeur de la zone | 
**Height** | **float32** | Hauteur de la zone | 

## Methods

### NewBoundingBoxSchema

`func NewBoundingBoxSchema(x0 float32, y0 float32, x1 float32, y1 float32, width float32, height float32, ) *BoundingBoxSchema`

NewBoundingBoxSchema instantiates a new BoundingBoxSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoundingBoxSchemaWithDefaults

`func NewBoundingBoxSchemaWithDefaults() *BoundingBoxSchema`

NewBoundingBoxSchemaWithDefaults instantiates a new BoundingBoxSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX0

`func (o *BoundingBoxSchema) GetX0() float32`

GetX0 returns the X0 field if non-nil, zero value otherwise.

### GetX0Ok

`func (o *BoundingBoxSchema) GetX0Ok() (*float32, bool)`

GetX0Ok returns a tuple with the X0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX0

`func (o *BoundingBoxSchema) SetX0(v float32)`

SetX0 sets X0 field to given value.


### GetY0

`func (o *BoundingBoxSchema) GetY0() float32`

GetY0 returns the Y0 field if non-nil, zero value otherwise.

### GetY0Ok

`func (o *BoundingBoxSchema) GetY0Ok() (*float32, bool)`

GetY0Ok returns a tuple with the Y0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY0

`func (o *BoundingBoxSchema) SetY0(v float32)`

SetY0 sets Y0 field to given value.


### GetX1

`func (o *BoundingBoxSchema) GetX1() float32`

GetX1 returns the X1 field if non-nil, zero value otherwise.

### GetX1Ok

`func (o *BoundingBoxSchema) GetX1Ok() (*float32, bool)`

GetX1Ok returns a tuple with the X1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX1

`func (o *BoundingBoxSchema) SetX1(v float32)`

SetX1 sets X1 field to given value.


### GetY1

`func (o *BoundingBoxSchema) GetY1() float32`

GetY1 returns the Y1 field if non-nil, zero value otherwise.

### GetY1Ok

`func (o *BoundingBoxSchema) GetY1Ok() (*float32, bool)`

GetY1Ok returns a tuple with the Y1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY1

`func (o *BoundingBoxSchema) SetY1(v float32)`

SetY1 sets Y1 field to given value.


### GetPage

`func (o *BoundingBoxSchema) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *BoundingBoxSchema) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *BoundingBoxSchema) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *BoundingBoxSchema) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetWidth

`func (o *BoundingBoxSchema) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *BoundingBoxSchema) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *BoundingBoxSchema) SetWidth(v float32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *BoundingBoxSchema) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *BoundingBoxSchema) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *BoundingBoxSchema) SetHeight(v float32)`

SetHeight sets Height field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



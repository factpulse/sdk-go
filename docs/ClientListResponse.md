# ClientListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]ClientSummary**](ClientSummary.md) | List of clients | 
**Total** | **int32** | Total number of clients | 
**Page** | **int32** | Current page | 
**PageSize** | **int32** | Page size | 

## Methods

### NewClientListResponse

`func NewClientListResponse(results []ClientSummary, total int32, page int32, pageSize int32, ) *ClientListResponse`

NewClientListResponse instantiates a new ClientListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientListResponseWithDefaults

`func NewClientListResponseWithDefaults() *ClientListResponse`

NewClientListResponseWithDefaults instantiates a new ClientListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ClientListResponse) GetResults() []ClientSummary`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ClientListResponse) GetResultsOk() (*[]ClientSummary, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ClientListResponse) SetResults(v []ClientSummary)`

SetResults sets Results field to given value.


### GetTotal

`func (o *ClientListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ClientListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ClientListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPage

`func (o *ClientListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ClientListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ClientListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *ClientListResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ClientListResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ClientListResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



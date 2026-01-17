# \CDARCycleDeVieAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateCdarApiV1CdarGeneratePost**](CDARCycleDeVieAPI.md#GenerateCdarApiV1CdarGeneratePost) | **Post** /api/v1/cdar/generate | Générer un message CDAR
[**GetActionCodesApiV1CdarActionCodesGet**](CDARCycleDeVieAPI.md#GetActionCodesApiV1CdarActionCodesGet) | **Get** /api/v1/cdar/action-codes | Liste des codes action CDAR
[**GetReasonCodesApiV1CdarReasonCodesGet**](CDARCycleDeVieAPI.md#GetReasonCodesApiV1CdarReasonCodesGet) | **Get** /api/v1/cdar/reason-codes | Liste des codes motif CDAR
[**GetStatusCodesApiV1CdarStatusCodesGet**](CDARCycleDeVieAPI.md#GetStatusCodesApiV1CdarStatusCodesGet) | **Get** /api/v1/cdar/status-codes | Liste des codes statut CDAR
[**SubmitCdarApiV1CdarSubmitPost**](CDARCycleDeVieAPI.md#SubmitCdarApiV1CdarSubmitPost) | **Post** /api/v1/cdar/submit | Générer et soumettre un message CDAR
[**SubmitCdarXmlApiV1CdarSubmitXmlPost**](CDARCycleDeVieAPI.md#SubmitCdarXmlApiV1CdarSubmitXmlPost) | **Post** /api/v1/cdar/submit-xml | Soumettre un XML CDAR pré-généré
[**ValidateCdarApiV1CdarValidatePost**](CDARCycleDeVieAPI.md#ValidateCdarApiV1CdarValidatePost) | **Post** /api/v1/cdar/validate | Valider des données CDAR



## GenerateCdarApiV1CdarGeneratePost

> GenerateCDARResponse GenerateCdarApiV1CdarGeneratePost(ctx).CreateCDARRequest(createCDARRequest).Execute()

Générer un message CDAR



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/factpulse/sdk-go/v3"
)

func main() {
	createCDARRequest := *openapiclient.NewCreateCDARRequest("DocumentId_example", "SenderSiren_example", "InvoiceId_example", time.Now(), "Status_example") // CreateCDARRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CDARCycleDeVieAPI.GenerateCdarApiV1CdarGeneratePost(context.Background()).CreateCDARRequest(createCDARRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CDARCycleDeVieAPI.GenerateCdarApiV1CdarGeneratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateCdarApiV1CdarGeneratePost`: GenerateCDARResponse
	fmt.Fprintf(os.Stdout, "Response from `CDARCycleDeVieAPI.GenerateCdarApiV1CdarGeneratePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateCdarApiV1CdarGeneratePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCDARRequest** | [**CreateCDARRequest**](CreateCDARRequest.md) |  | 

### Return type

[**GenerateCDARResponse**](GenerateCDARResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActionCodesApiV1CdarActionCodesGet

> ActionCodesResponse GetActionCodesApiV1CdarActionCodesGet(ctx).Execute()

Liste des codes action CDAR



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/factpulse/sdk-go/v3"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CDARCycleDeVieAPI.GetActionCodesApiV1CdarActionCodesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CDARCycleDeVieAPI.GetActionCodesApiV1CdarActionCodesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActionCodesApiV1CdarActionCodesGet`: ActionCodesResponse
	fmt.Fprintf(os.Stdout, "Response from `CDARCycleDeVieAPI.GetActionCodesApiV1CdarActionCodesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetActionCodesApiV1CdarActionCodesGetRequest struct via the builder pattern


### Return type

[**ActionCodesResponse**](ActionCodesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReasonCodesApiV1CdarReasonCodesGet

> ReasonCodesResponse GetReasonCodesApiV1CdarReasonCodesGet(ctx).Execute()

Liste des codes motif CDAR



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/factpulse/sdk-go/v3"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CDARCycleDeVieAPI.GetReasonCodesApiV1CdarReasonCodesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CDARCycleDeVieAPI.GetReasonCodesApiV1CdarReasonCodesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReasonCodesApiV1CdarReasonCodesGet`: ReasonCodesResponse
	fmt.Fprintf(os.Stdout, "Response from `CDARCycleDeVieAPI.GetReasonCodesApiV1CdarReasonCodesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetReasonCodesApiV1CdarReasonCodesGetRequest struct via the builder pattern


### Return type

[**ReasonCodesResponse**](ReasonCodesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatusCodesApiV1CdarStatusCodesGet

> StatusCodesResponse GetStatusCodesApiV1CdarStatusCodesGet(ctx).Execute()

Liste des codes statut CDAR



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/factpulse/sdk-go/v3"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CDARCycleDeVieAPI.GetStatusCodesApiV1CdarStatusCodesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CDARCycleDeVieAPI.GetStatusCodesApiV1CdarStatusCodesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatusCodesApiV1CdarStatusCodesGet`: StatusCodesResponse
	fmt.Fprintf(os.Stdout, "Response from `CDARCycleDeVieAPI.GetStatusCodesApiV1CdarStatusCodesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusCodesApiV1CdarStatusCodesGetRequest struct via the builder pattern


### Return type

[**StatusCodesResponse**](StatusCodesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitCdarApiV1CdarSubmitPost

> SubmitCDARResponse SubmitCdarApiV1CdarSubmitPost(ctx).UserId(userId).BodySubmitCdarApiV1CdarSubmitPost(bodySubmitCdarApiV1CdarSubmitPost).JwtToken(jwtToken).ClientUid(clientUid).Execute()

Générer et soumettre un message CDAR



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/factpulse/sdk-go/v3"
)

func main() {
	userId := int32(56) // int32 | 
	bodySubmitCdarApiV1CdarSubmitPost := *openapiclient.NewBodySubmitCdarApiV1CdarSubmitPost(*openapiclient.NewSubmitCDARRequest("DocumentId_example", "SenderSiren_example", "InvoiceId_example", time.Now(), "Status_example")) // BodySubmitCdarApiV1CdarSubmitPost | 
	jwtToken := "jwtToken_example" // string |  (optional)
	clientUid := "clientUid_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CDARCycleDeVieAPI.SubmitCdarApiV1CdarSubmitPost(context.Background()).UserId(userId).BodySubmitCdarApiV1CdarSubmitPost(bodySubmitCdarApiV1CdarSubmitPost).JwtToken(jwtToken).ClientUid(clientUid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CDARCycleDeVieAPI.SubmitCdarApiV1CdarSubmitPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitCdarApiV1CdarSubmitPost`: SubmitCDARResponse
	fmt.Fprintf(os.Stdout, "Response from `CDARCycleDeVieAPI.SubmitCdarApiV1CdarSubmitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitCdarApiV1CdarSubmitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32** |  | 
 **bodySubmitCdarApiV1CdarSubmitPost** | [**BodySubmitCdarApiV1CdarSubmitPost**](BodySubmitCdarApiV1CdarSubmitPost.md) |  | 
 **jwtToken** | **string** |  | 
 **clientUid** | **string** |  | 

### Return type

[**SubmitCDARResponse**](SubmitCDARResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitCdarXmlApiV1CdarSubmitXmlPost

> SubmitCDARResponse SubmitCdarXmlApiV1CdarSubmitXmlPost(ctx).UserId(userId).BodySubmitCdarXmlApiV1CdarSubmitXmlPost(bodySubmitCdarXmlApiV1CdarSubmitXmlPost).JwtToken(jwtToken).ClientUid(clientUid).Execute()

Soumettre un XML CDAR pré-généré



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/factpulse/sdk-go/v3"
)

func main() {
	userId := int32(56) // int32 | 
	bodySubmitCdarXmlApiV1CdarSubmitXmlPost := *openapiclient.NewBodySubmitCdarXmlApiV1CdarSubmitXmlPost(*openapiclient.NewSubmitCDARXMLRequest("Xml_example")) // BodySubmitCdarXmlApiV1CdarSubmitXmlPost | 
	jwtToken := "jwtToken_example" // string |  (optional)
	clientUid := "clientUid_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CDARCycleDeVieAPI.SubmitCdarXmlApiV1CdarSubmitXmlPost(context.Background()).UserId(userId).BodySubmitCdarXmlApiV1CdarSubmitXmlPost(bodySubmitCdarXmlApiV1CdarSubmitXmlPost).JwtToken(jwtToken).ClientUid(clientUid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CDARCycleDeVieAPI.SubmitCdarXmlApiV1CdarSubmitXmlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitCdarXmlApiV1CdarSubmitXmlPost`: SubmitCDARResponse
	fmt.Fprintf(os.Stdout, "Response from `CDARCycleDeVieAPI.SubmitCdarXmlApiV1CdarSubmitXmlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitCdarXmlApiV1CdarSubmitXmlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32** |  | 
 **bodySubmitCdarXmlApiV1CdarSubmitXmlPost** | [**BodySubmitCdarXmlApiV1CdarSubmitXmlPost**](BodySubmitCdarXmlApiV1CdarSubmitXmlPost.md) |  | 
 **jwtToken** | **string** |  | 
 **clientUid** | **string** |  | 

### Return type

[**SubmitCDARResponse**](SubmitCDARResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateCdarApiV1CdarValidatePost

> ValidateCDARResponse ValidateCdarApiV1CdarValidatePost(ctx).ValidateCDARRequest(validateCDARRequest).Execute()

Valider des données CDAR



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/factpulse/sdk-go/v3"
)

func main() {
	validateCDARRequest := *openapiclient.NewValidateCDARRequest() // ValidateCDARRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CDARCycleDeVieAPI.ValidateCdarApiV1CdarValidatePost(context.Background()).ValidateCDARRequest(validateCDARRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CDARCycleDeVieAPI.ValidateCdarApiV1CdarValidatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateCdarApiV1CdarValidatePost`: ValidateCDARResponse
	fmt.Fprintf(os.Stdout, "Response from `CDARCycleDeVieAPI.ValidateCdarApiV1CdarValidatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCdarApiV1CdarValidatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateCDARRequest** | [**ValidateCDARRequest**](ValidateCDARRequest.md) |  | 

### Return type

[**ValidateCDARResponse**](ValidateCDARResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


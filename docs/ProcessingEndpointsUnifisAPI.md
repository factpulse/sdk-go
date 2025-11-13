# \ProcessingEndpointsUnifisAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePost**](ProcessingEndpointsUnifisAPI.md#SoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePost) | **Post** /api/v1/traitement/factures/soumettre-complete | Soumettre une facture complète (génération + signature + soumission)
[**SoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPost**](ProcessingEndpointsUnifisAPI.md#SoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPost) | **Post** /api/v1/traitement/factures/soumettre-complete-async | Soumettre une facture complète (asynchrone avec Celery)



## SoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePost

> SoumettreFactureCompleteResponse SoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePost(ctx).SoumettreFactureCompleteRequest(soumettreFactureCompleteRequest).Execute()

Soumettre une facture complète (génération + signature + soumission)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	soumettreFactureCompleteRequest := *openapiclient.NewSoumettreFactureCompleteRequest(*openapiclient.NewDonneesFactureSimplifiees("Numero_example", map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}), "PdfSource_example", openapiclient.Destination{DestinationAFNOR: openapiclient.NewDestinationAFNOR()}) // SoumettreFactureCompleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessingEndpointsUnifisAPI.SoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePost(context.Background()).SoumettreFactureCompleteRequest(soumettreFactureCompleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessingEndpointsUnifisAPI.SoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePost`: SoumettreFactureCompleteResponse
	fmt.Fprintf(os.Stdout, "Response from `ProcessingEndpointsUnifisAPI.SoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSoumettreFactureCompleteApiV1TraitementFacturesSoumettreCompletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **soumettreFactureCompleteRequest** | [**SoumettreFactureCompleteRequest**](SoumettreFactureCompleteRequest.md) |  | 

### Return type

[**SoumettreFactureCompleteResponse**](SoumettreFactureCompleteResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPost

> ReponseTache SoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPost(ctx).SoumettreFactureCompleteRequest(soumettreFactureCompleteRequest).Execute()

Soumettre une facture complète (asynchrone avec Celery)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	soumettreFactureCompleteRequest := *openapiclient.NewSoumettreFactureCompleteRequest(*openapiclient.NewDonneesFactureSimplifiees("Numero_example", map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}), "PdfSource_example", openapiclient.Destination{DestinationAFNOR: openapiclient.NewDestinationAFNOR()}) // SoumettreFactureCompleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessingEndpointsUnifisAPI.SoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPost(context.Background()).SoumettreFactureCompleteRequest(soumettreFactureCompleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessingEndpointsUnifisAPI.SoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPost`: ReponseTache
	fmt.Fprintf(os.Stdout, "Response from `ProcessingEndpointsUnifisAPI.SoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSoumettreFactureCompleteAsyncApiV1TraitementFacturesSoumettreCompleteAsyncPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **soumettreFactureCompleteRequest** | [**SoumettreFactureCompleteRequest**](SoumettreFactureCompleteRequest.md) |  | 

### Return type

[**ReponseTache**](ReponseTache.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


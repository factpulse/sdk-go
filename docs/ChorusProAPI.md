# \ChorusProAPI

All URIs are relative to *https://factpulse.fr*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AjouterFichierApiV1ChorusProTransversesAjouterFichierPost**](ChorusProAPI.md#AjouterFichierApiV1ChorusProTransversesAjouterFichierPost) | **Post** /api/v1/chorus-pro/transverses/ajouter-fichier | Add an attachment
[**CompleterFactureApiV1ChorusProFacturesCompleterPost**](ChorusProAPI.md#CompleterFactureApiV1ChorusProFacturesCompleterPost) | **Post** /api/v1/chorus-pro/factures/completer | Complete a suspended invoice (Supplier)
[**ConsulterFactureApiV1ChorusProFacturesConsulterPost**](ChorusProAPI.md#ConsulterFactureApiV1ChorusProFacturesConsulterPost) | **Post** /api/v1/chorus-pro/factures/consulter | Consult invoice status
[**ConsulterStructureApiV1ChorusProStructuresConsulterPost**](ChorusProAPI.md#ConsulterStructureApiV1ChorusProStructuresConsulterPost) | **Post** /api/v1/chorus-pro/structures/consulter | Consult structure details
[**ListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet**](ChorusProAPI.md#ListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet) | **Get** /api/v1/chorus-pro/structures/{id_structure_cpp}/services | List structure services
[**ObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPost**](ChorusProAPI.md#ObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPost) | **Post** /api/v1/chorus-pro/structures/obtenir-id-depuis-siret | Utility: Get Chorus Pro ID from SIRET
[**RechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost**](ChorusProAPI.md#RechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost) | **Post** /api/v1/chorus-pro/factures/rechercher-destinataire | Search received invoices (Recipient)
[**RechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost**](ChorusProAPI.md#RechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost) | **Post** /api/v1/chorus-pro/factures/rechercher-fournisseur | Search issued invoices (Supplier)
[**RechercherStructuresApiV1ChorusProStructuresRechercherPost**](ChorusProAPI.md#RechercherStructuresApiV1ChorusProStructuresRechercherPost) | **Post** /api/v1/chorus-pro/structures/rechercher | Search Chorus Pro structures
[**RecyclerFactureApiV1ChorusProFacturesRecyclerPost**](ChorusProAPI.md#RecyclerFactureApiV1ChorusProFacturesRecyclerPost) | **Post** /api/v1/chorus-pro/factures/recycler | Recycle an invoice (Supplier)
[**SoumettreFactureApiV1ChorusProFacturesSoumettrePost**](ChorusProAPI.md#SoumettreFactureApiV1ChorusProFacturesSoumettrePost) | **Post** /api/v1/chorus-pro/factures/soumettre | Submit an invoice to Chorus Pro
[**TelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost**](ChorusProAPI.md#TelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost) | **Post** /api/v1/chorus-pro/factures/telecharger-groupe | Download a group of invoices
[**TraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost**](ChorusProAPI.md#TraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost) | **Post** /api/v1/chorus-pro/factures/traiter-facture-recue | Process a received invoice (Recipient)
[**ValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost**](ChorusProAPI.md#ValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost) | **Post** /api/v1/chorus-pro/factures/valideur/consulter | Consult an invoice (Validator)
[**ValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost**](ChorusProAPI.md#ValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost) | **Post** /api/v1/chorus-pro/factures/valideur/rechercher | Search invoices to validate (Validator)
[**ValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost**](ChorusProAPI.md#ValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost) | **Post** /api/v1/chorus-pro/factures/valideur/traiter | Validate or reject an invoice (Validator)



## AjouterFichierApiV1ChorusProTransversesAjouterFichierPost

> interface{} AjouterFichierApiV1ChorusProTransversesAjouterFichierPost(ctx).RequestBody(requestBody).Execute()

Add an attachment



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
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.AjouterFichierApiV1ChorusProTransversesAjouterFichierPost(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChorusProAPI.AjouterFichierApiV1ChorusProTransversesAjouterFichierPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AjouterFichierApiV1ChorusProTransversesAjouterFichierPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChorusProAPI.AjouterFichierApiV1ChorusProTransversesAjouterFichierPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAjouterFichierApiV1ChorusProTransversesAjouterFichierPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleterFactureApiV1ChorusProFacturesCompleterPost

> interface{} CompleterFactureApiV1ChorusProFacturesCompleterPost(ctx).RequestBody(requestBody).Execute()

Complete a suspended invoice (Supplier)



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
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.CompleterFactureApiV1ChorusProFacturesCompleterPost(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChorusProAPI.CompleterFactureApiV1ChorusProFacturesCompleterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleterFactureApiV1ChorusProFacturesCompleterPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChorusProAPI.CompleterFactureApiV1ChorusProFacturesCompleterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleterFactureApiV1ChorusProFacturesCompleterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsulterFactureApiV1ChorusProFacturesConsulterPost

> GetInvoiceResponse ConsulterFactureApiV1ChorusProFacturesConsulterPost(ctx).GetInvoiceRequest(getInvoiceRequest).Execute()

Consult invoice status



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
	getInvoiceRequest := *openapiclient.NewGetInvoiceRequest(int32(123)) // GetInvoiceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.ConsulterFactureApiV1ChorusProFacturesConsulterPost(context.Background()).GetInvoiceRequest(getInvoiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChorusProAPI.ConsulterFactureApiV1ChorusProFacturesConsulterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsulterFactureApiV1ChorusProFacturesConsulterPost`: GetInvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `ChorusProAPI.ConsulterFactureApiV1ChorusProFacturesConsulterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsulterFactureApiV1ChorusProFacturesConsulterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getInvoiceRequest** | [**GetInvoiceRequest**](GetInvoiceRequest.md) |  | 

### Return type

[**GetInvoiceResponse**](GetInvoiceResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsulterStructureApiV1ChorusProStructuresConsulterPost

> GetStructureResponse ConsulterStructureApiV1ChorusProStructuresConsulterPost(ctx).GetStructureRequest(getStructureRequest).Execute()

Consult structure details



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
	getStructureRequest := *openapiclient.NewGetStructureRequest(int32(123)) // GetStructureRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.ConsulterStructureApiV1ChorusProStructuresConsulterPost(context.Background()).GetStructureRequest(getStructureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChorusProAPI.ConsulterStructureApiV1ChorusProStructuresConsulterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsulterStructureApiV1ChorusProStructuresConsulterPost`: GetStructureResponse
	fmt.Fprintf(os.Stdout, "Response from `ChorusProAPI.ConsulterStructureApiV1ChorusProStructuresConsulterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsulterStructureApiV1ChorusProStructuresConsulterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getStructureRequest** | [**GetStructureRequest**](GetStructureRequest.md) |  | 

### Return type

[**GetStructureResponse**](GetStructureResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet

> SearchServicesResponse ListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet(ctx, idStructureCpp).Execute()

List structure services



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
	idStructureCpp := int32(56) // int32 | Chorus Pro structure ID (idStructureCPP)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.ListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet(context.Background(), idStructureCpp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChorusProAPI.ListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet`: SearchServicesResponse
	fmt.Fprintf(os.Stdout, "Response from `ChorusProAPI.ListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idStructureCpp** | **int32** | Chorus Pro structure ID (idStructureCPP) | 

### Other Parameters

Other parameters are passed through a pointer to a apiListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SearchServicesResponse**](SearchServicesResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPost

> GetChorusProIdResponse ObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPost(ctx).GetChorusProIdRequest(getChorusProIdRequest).Execute()

Utility: Get Chorus Pro ID from SIRET



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
	getChorusProIdRequest := *openapiclient.NewGetChorusProIdRequest("Siret_example") // GetChorusProIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.ObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPost(context.Background()).GetChorusProIdRequest(getChorusProIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChorusProAPI.ObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPost`: GetChorusProIdResponse
	fmt.Fprintf(os.Stdout, "Response from `ChorusProAPI.ObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getChorusProIdRequest** | [**GetChorusProIdRequest**](GetChorusProIdRequest.md) |  | 

### Return type

[**GetChorusProIdResponse**](GetChorusProIdResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost

> interface{} RechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost(ctx).RequestBody(requestBody).Execute()

Search received invoices (Recipient)



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
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.RechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChorusProAPI.RechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChorusProAPI.RechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost

> interface{} RechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost(ctx).RequestBody(requestBody).Execute()

Search issued invoices (Supplier)



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
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.RechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChorusProAPI.RechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChorusProAPI.RechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RechercherStructuresApiV1ChorusProStructuresRechercherPost

> SearchStructureResponse RechercherStructuresApiV1ChorusProStructuresRechercherPost(ctx).SearchStructureRequest(searchStructureRequest).Execute()

Search Chorus Pro structures



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
	searchStructureRequest := *openapiclient.NewSearchStructureRequest() // SearchStructureRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.RechercherStructuresApiV1ChorusProStructuresRechercherPost(context.Background()).SearchStructureRequest(searchStructureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChorusProAPI.RechercherStructuresApiV1ChorusProStructuresRechercherPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RechercherStructuresApiV1ChorusProStructuresRechercherPost`: SearchStructureResponse
	fmt.Fprintf(os.Stdout, "Response from `ChorusProAPI.RechercherStructuresApiV1ChorusProStructuresRechercherPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRechercherStructuresApiV1ChorusProStructuresRechercherPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchStructureRequest** | [**SearchStructureRequest**](SearchStructureRequest.md) |  | 

### Return type

[**SearchStructureResponse**](SearchStructureResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecyclerFactureApiV1ChorusProFacturesRecyclerPost

> interface{} RecyclerFactureApiV1ChorusProFacturesRecyclerPost(ctx).RequestBody(requestBody).Execute()

Recycle an invoice (Supplier)



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
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.RecyclerFactureApiV1ChorusProFacturesRecyclerPost(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChorusProAPI.RecyclerFactureApiV1ChorusProFacturesRecyclerPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecyclerFactureApiV1ChorusProFacturesRecyclerPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChorusProAPI.RecyclerFactureApiV1ChorusProFacturesRecyclerPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecyclerFactureApiV1ChorusProFacturesRecyclerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoumettreFactureApiV1ChorusProFacturesSoumettrePost

> SubmitInvoiceResponse SoumettreFactureApiV1ChorusProFacturesSoumettrePost(ctx).SubmitInvoiceRequest(submitInvoiceRequest).Execute()

Submit an invoice to Chorus Pro



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
	submitInvoiceRequest := *openapiclient.NewSubmitInvoiceRequest("InvoiceNumber_example", "InvoiceDate_example", int32(123), *openapiclient.NewSubmitNetAmount(), *openapiclient.NewSubmitVatAmount(), *openapiclient.NewSubmitGrossAmount()) // SubmitInvoiceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.SoumettreFactureApiV1ChorusProFacturesSoumettrePost(context.Background()).SubmitInvoiceRequest(submitInvoiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChorusProAPI.SoumettreFactureApiV1ChorusProFacturesSoumettrePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SoumettreFactureApiV1ChorusProFacturesSoumettrePost`: SubmitInvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `ChorusProAPI.SoumettreFactureApiV1ChorusProFacturesSoumettrePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSoumettreFactureApiV1ChorusProFacturesSoumettrePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitInvoiceRequest** | [**SubmitInvoiceRequest**](SubmitInvoiceRequest.md) |  | 

### Return type

[**SubmitInvoiceResponse**](SubmitInvoiceResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost

> interface{} TelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost(ctx).RequestBody(requestBody).Execute()

Download a group of invoices



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
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.TelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChorusProAPI.TelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChorusProAPI.TelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost

> interface{} TraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost(ctx).RequestBody(requestBody).Execute()

Process a received invoice (Recipient)



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
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.TraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChorusProAPI.TraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChorusProAPI.TraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost

> interface{} ValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost(ctx).RequestBody(requestBody).Execute()

Consult an invoice (Validator)



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
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.ValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChorusProAPI.ValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChorusProAPI.ValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost

> interface{} ValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost(ctx).RequestBody(requestBody).Execute()

Search invoices to validate (Validator)



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
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.ValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChorusProAPI.ValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChorusProAPI.ValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost

> interface{} ValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost(ctx).RequestBody(requestBody).Execute()

Validate or reject an invoice (Validator)



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
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.ValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChorusProAPI.ValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChorusProAPI.ValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


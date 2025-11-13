# \ChorusProAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AjouterFichierApiV1ChorusProTransversesAjouterFichierPost**](ChorusProAPI.md#AjouterFichierApiV1ChorusProTransversesAjouterFichierPost) | **Post** /api/v1/chorus-pro/transverses/ajouter-fichier | Ajouter une pièce jointe
[**CompleterFactureApiV1ChorusProFacturesCompleterPost**](ChorusProAPI.md#CompleterFactureApiV1ChorusProFacturesCompleterPost) | **Post** /api/v1/chorus-pro/factures/completer | Compléter une facture suspendue (Fournisseur)
[**ConsulterFactureApiV1ChorusProFacturesConsulterPost**](ChorusProAPI.md#ConsulterFactureApiV1ChorusProFacturesConsulterPost) | **Post** /api/v1/chorus-pro/factures/consulter | Consulter le statut d&#39;une facture
[**ConsulterStructureApiV1ChorusProStructuresConsulterPost**](ChorusProAPI.md#ConsulterStructureApiV1ChorusProStructuresConsulterPost) | **Post** /api/v1/chorus-pro/structures/consulter | Consulter les détails d&#39;une structure
[**ListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet**](ChorusProAPI.md#ListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet) | **Get** /api/v1/chorus-pro/structures/{id_structure_cpp}/services | Lister les services d&#39;une structure
[**ObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPost**](ChorusProAPI.md#ObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPost) | **Post** /api/v1/chorus-pro/structures/obtenir-id-depuis-siret | Utilitaire : Obtenir l&#39;ID Chorus Pro depuis un SIRET
[**RechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost**](ChorusProAPI.md#RechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost) | **Post** /api/v1/chorus-pro/factures/rechercher-destinataire | Rechercher factures reçues (Destinataire)
[**RechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost**](ChorusProAPI.md#RechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost) | **Post** /api/v1/chorus-pro/factures/rechercher-fournisseur | Rechercher factures émises (Fournisseur)
[**RechercherStructuresApiV1ChorusProStructuresRechercherPost**](ChorusProAPI.md#RechercherStructuresApiV1ChorusProStructuresRechercherPost) | **Post** /api/v1/chorus-pro/structures/rechercher | Rechercher des structures Chorus Pro
[**RecyclerFactureApiV1ChorusProFacturesRecyclerPost**](ChorusProAPI.md#RecyclerFactureApiV1ChorusProFacturesRecyclerPost) | **Post** /api/v1/chorus-pro/factures/recycler | Recycler une facture (Fournisseur)
[**SoumettreFactureApiV1ChorusProFacturesSoumettrePost**](ChorusProAPI.md#SoumettreFactureApiV1ChorusProFacturesSoumettrePost) | **Post** /api/v1/chorus-pro/factures/soumettre | Soumettre une facture à Chorus Pro
[**TelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost**](ChorusProAPI.md#TelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost) | **Post** /api/v1/chorus-pro/factures/telecharger-groupe | Télécharger un groupe de factures
[**TraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost**](ChorusProAPI.md#TraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost) | **Post** /api/v1/chorus-pro/factures/traiter-facture-recue | Traiter une facture reçue (Destinataire)
[**ValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost**](ChorusProAPI.md#ValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost) | **Post** /api/v1/chorus-pro/factures/valideur/consulter | Consulter une facture (Valideur)
[**ValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost**](ChorusProAPI.md#ValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost) | **Post** /api/v1/chorus-pro/factures/valideur/rechercher | Rechercher factures à valider (Valideur)
[**ValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost**](ChorusProAPI.md#ValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost) | **Post** /api/v1/chorus-pro/factures/valideur/traiter | Valider ou refuser une facture (Valideur)



## AjouterFichierApiV1ChorusProTransversesAjouterFichierPost

> interface{} AjouterFichierApiV1ChorusProTransversesAjouterFichierPost(ctx).BodyAjouterFichierApiV1ChorusProTransversesAjouterFichierPost(bodyAjouterFichierApiV1ChorusProTransversesAjouterFichierPost).Execute()

Ajouter une pièce jointe



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
	bodyAjouterFichierApiV1ChorusProTransversesAjouterFichierPost := *openapiclient.NewBodyAjouterFichierApiV1ChorusProTransversesAjouterFichierPost(map[string]interface{}{"key": interface{}(123)}, *openapiclient.NewUtilisateur(int32(123), "Username_example", "Email_example", false)) // BodyAjouterFichierApiV1ChorusProTransversesAjouterFichierPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.AjouterFichierApiV1ChorusProTransversesAjouterFichierPost(context.Background()).BodyAjouterFichierApiV1ChorusProTransversesAjouterFichierPost(bodyAjouterFichierApiV1ChorusProTransversesAjouterFichierPost).Execute()
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
 **bodyAjouterFichierApiV1ChorusProTransversesAjouterFichierPost** | [**BodyAjouterFichierApiV1ChorusProTransversesAjouterFichierPost**](BodyAjouterFichierApiV1ChorusProTransversesAjouterFichierPost.md) |  | 

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

> interface{} CompleterFactureApiV1ChorusProFacturesCompleterPost(ctx).BodyCompleterFactureApiV1ChorusProFacturesCompleterPost(bodyCompleterFactureApiV1ChorusProFacturesCompleterPost).Execute()

Compléter une facture suspendue (Fournisseur)



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
	bodyCompleterFactureApiV1ChorusProFacturesCompleterPost := *openapiclient.NewBodyCompleterFactureApiV1ChorusProFacturesCompleterPost(map[string]interface{}{"key": interface{}(123)}, *openapiclient.NewUtilisateur(int32(123), "Username_example", "Email_example", false)) // BodyCompleterFactureApiV1ChorusProFacturesCompleterPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.CompleterFactureApiV1ChorusProFacturesCompleterPost(context.Background()).BodyCompleterFactureApiV1ChorusProFacturesCompleterPost(bodyCompleterFactureApiV1ChorusProFacturesCompleterPost).Execute()
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
 **bodyCompleterFactureApiV1ChorusProFacturesCompleterPost** | [**BodyCompleterFactureApiV1ChorusProFacturesCompleterPost**](BodyCompleterFactureApiV1ChorusProFacturesCompleterPost.md) |  | 

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

> ConsulterFactureResponse ConsulterFactureApiV1ChorusProFacturesConsulterPost(ctx).ConsulterFactureRequest(consulterFactureRequest).Execute()

Consulter le statut d'une facture



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
	consulterFactureRequest := *openapiclient.NewConsulterFactureRequest(int32(123)) // ConsulterFactureRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.ConsulterFactureApiV1ChorusProFacturesConsulterPost(context.Background()).ConsulterFactureRequest(consulterFactureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChorusProAPI.ConsulterFactureApiV1ChorusProFacturesConsulterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsulterFactureApiV1ChorusProFacturesConsulterPost`: ConsulterFactureResponse
	fmt.Fprintf(os.Stdout, "Response from `ChorusProAPI.ConsulterFactureApiV1ChorusProFacturesConsulterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsulterFactureApiV1ChorusProFacturesConsulterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consulterFactureRequest** | [**ConsulterFactureRequest**](ConsulterFactureRequest.md) |  | 

### Return type

[**ConsulterFactureResponse**](ConsulterFactureResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsulterStructureApiV1ChorusProStructuresConsulterPost

> ConsulterStructureResponse ConsulterStructureApiV1ChorusProStructuresConsulterPost(ctx).ConsulterStructureRequest(consulterStructureRequest).Execute()

Consulter les détails d'une structure



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
	consulterStructureRequest := *openapiclient.NewConsulterStructureRequest(int32(123)) // ConsulterStructureRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.ConsulterStructureApiV1ChorusProStructuresConsulterPost(context.Background()).ConsulterStructureRequest(consulterStructureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChorusProAPI.ConsulterStructureApiV1ChorusProStructuresConsulterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsulterStructureApiV1ChorusProStructuresConsulterPost`: ConsulterStructureResponse
	fmt.Fprintf(os.Stdout, "Response from `ChorusProAPI.ConsulterStructureApiV1ChorusProStructuresConsulterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsulterStructureApiV1ChorusProStructuresConsulterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consulterStructureRequest** | [**ConsulterStructureRequest**](ConsulterStructureRequest.md) |  | 

### Return type

[**ConsulterStructureResponse**](ConsulterStructureResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet

> RechercherServicesResponse ListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet(ctx, idStructureCpp).BodyListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet(bodyListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet).Execute()

Lister les services d'une structure



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
	idStructureCpp := int32(56) // int32 | 
	bodyListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet := *openapiclient.NewBodyListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet(*openapiclient.NewUtilisateur(int32(123), "Username_example", "Email_example", false)) // BodyListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.ListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet(context.Background(), idStructureCpp).BodyListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet(bodyListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChorusProAPI.ListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet`: RechercherServicesResponse
	fmt.Fprintf(os.Stdout, "Response from `ChorusProAPI.ListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idStructureCpp** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet** | [**BodyListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet**](BodyListerServicesStructureApiV1ChorusProStructuresIdStructureCppServicesGet.md) |  | 

### Return type

[**RechercherServicesResponse**](RechercherServicesResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPost

> ObtenirIdChorusProResponse ObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPost(ctx).ObtenirIdChorusProRequest(obtenirIdChorusProRequest).Execute()

Utilitaire : Obtenir l'ID Chorus Pro depuis un SIRET



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
	obtenirIdChorusProRequest := *openapiclient.NewObtenirIdChorusProRequest("Siret_example") // ObtenirIdChorusProRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.ObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPost(context.Background()).ObtenirIdChorusProRequest(obtenirIdChorusProRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChorusProAPI.ObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPost`: ObtenirIdChorusProResponse
	fmt.Fprintf(os.Stdout, "Response from `ChorusProAPI.ObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiObtenirIdChorusProDepuisSiretApiV1ChorusProStructuresObtenirIdDepuisSiretPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **obtenirIdChorusProRequest** | [**ObtenirIdChorusProRequest**](ObtenirIdChorusProRequest.md) |  | 

### Return type

[**ObtenirIdChorusProResponse**](ObtenirIdChorusProResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost

> interface{} RechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost(ctx).BodyRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost(bodyRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost).Execute()

Rechercher factures reçues (Destinataire)



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
	bodyRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost := *openapiclient.NewBodyRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost(map[string]interface{}{"key": interface{}(123)}, *openapiclient.NewUtilisateur(int32(123), "Username_example", "Email_example", false)) // BodyRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.RechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost(context.Background()).BodyRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost(bodyRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost).Execute()
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
 **bodyRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost** | [**BodyRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost**](BodyRechercherFacturesDestinataireApiV1ChorusProFacturesRechercherDestinatairePost.md) |  | 

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

> interface{} RechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost(ctx).BodyRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost(bodyRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost).Execute()

Rechercher factures émises (Fournisseur)



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
	bodyRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost := *openapiclient.NewBodyRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost(map[string]interface{}{"key": interface{}(123)}, *openapiclient.NewUtilisateur(int32(123), "Username_example", "Email_example", false)) // BodyRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.RechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost(context.Background()).BodyRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost(bodyRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost).Execute()
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
 **bodyRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost** | [**BodyRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost**](BodyRechercherFacturesFournisseurApiV1ChorusProFacturesRechercherFournisseurPost.md) |  | 

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

> RechercherStructureResponse RechercherStructuresApiV1ChorusProStructuresRechercherPost(ctx).RechercherStructureRequest(rechercherStructureRequest).Execute()

Rechercher des structures Chorus Pro



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
	rechercherStructureRequest := *openapiclient.NewRechercherStructureRequest() // RechercherStructureRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.RechercherStructuresApiV1ChorusProStructuresRechercherPost(context.Background()).RechercherStructureRequest(rechercherStructureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChorusProAPI.RechercherStructuresApiV1ChorusProStructuresRechercherPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RechercherStructuresApiV1ChorusProStructuresRechercherPost`: RechercherStructureResponse
	fmt.Fprintf(os.Stdout, "Response from `ChorusProAPI.RechercherStructuresApiV1ChorusProStructuresRechercherPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRechercherStructuresApiV1ChorusProStructuresRechercherPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rechercherStructureRequest** | [**RechercherStructureRequest**](RechercherStructureRequest.md) |  | 

### Return type

[**RechercherStructureResponse**](RechercherStructureResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecyclerFactureApiV1ChorusProFacturesRecyclerPost

> interface{} RecyclerFactureApiV1ChorusProFacturesRecyclerPost(ctx).BodyRecyclerFactureApiV1ChorusProFacturesRecyclerPost(bodyRecyclerFactureApiV1ChorusProFacturesRecyclerPost).Execute()

Recycler une facture (Fournisseur)



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
	bodyRecyclerFactureApiV1ChorusProFacturesRecyclerPost := *openapiclient.NewBodyRecyclerFactureApiV1ChorusProFacturesRecyclerPost(map[string]interface{}{"key": interface{}(123)}, *openapiclient.NewUtilisateur(int32(123), "Username_example", "Email_example", false)) // BodyRecyclerFactureApiV1ChorusProFacturesRecyclerPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.RecyclerFactureApiV1ChorusProFacturesRecyclerPost(context.Background()).BodyRecyclerFactureApiV1ChorusProFacturesRecyclerPost(bodyRecyclerFactureApiV1ChorusProFacturesRecyclerPost).Execute()
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
 **bodyRecyclerFactureApiV1ChorusProFacturesRecyclerPost** | [**BodyRecyclerFactureApiV1ChorusProFacturesRecyclerPost**](BodyRecyclerFactureApiV1ChorusProFacturesRecyclerPost.md) |  | 

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

> SoumettreFactureResponse SoumettreFactureApiV1ChorusProFacturesSoumettrePost(ctx).SoumettreFactureRequest(soumettreFactureRequest).Execute()

Soumettre une facture à Chorus Pro



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
	soumettreFactureRequest := *openapiclient.NewSoumettreFactureRequest("NumeroFacture_example", "DateFacture_example", int32(123), "MontantHtTotal_example", "MontantTva_example", "MontantTtcTotal_example") // SoumettreFactureRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.SoumettreFactureApiV1ChorusProFacturesSoumettrePost(context.Background()).SoumettreFactureRequest(soumettreFactureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChorusProAPI.SoumettreFactureApiV1ChorusProFacturesSoumettrePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SoumettreFactureApiV1ChorusProFacturesSoumettrePost`: SoumettreFactureResponse
	fmt.Fprintf(os.Stdout, "Response from `ChorusProAPI.SoumettreFactureApiV1ChorusProFacturesSoumettrePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSoumettreFactureApiV1ChorusProFacturesSoumettrePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **soumettreFactureRequest** | [**SoumettreFactureRequest**](SoumettreFactureRequest.md) |  | 

### Return type

[**SoumettreFactureResponse**](SoumettreFactureResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost

> interface{} TelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost(ctx).BodyTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost(bodyTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost).Execute()

Télécharger un groupe de factures



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
	bodyTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost := *openapiclient.NewBodyTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost(map[string]interface{}{"key": interface{}(123)}, *openapiclient.NewUtilisateur(int32(123), "Username_example", "Email_example", false)) // BodyTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.TelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost(context.Background()).BodyTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost(bodyTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost).Execute()
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
 **bodyTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost** | [**BodyTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost**](BodyTelechargerGroupeFacturesApiV1ChorusProFacturesTelechargerGroupePost.md) |  | 

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

> interface{} TraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost(ctx).BodyTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost(bodyTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost).Execute()

Traiter une facture reçue (Destinataire)



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
	bodyTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost := *openapiclient.NewBodyTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost(map[string]interface{}{"key": interface{}(123)}, *openapiclient.NewUtilisateur(int32(123), "Username_example", "Email_example", false)) // BodyTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.TraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost(context.Background()).BodyTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost(bodyTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost).Execute()
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
 **bodyTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost** | [**BodyTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost**](BodyTraiterFactureRecueApiV1ChorusProFacturesTraiterFactureRecuePost.md) |  | 

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

> interface{} ValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost(ctx).BodyValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost(bodyValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost).Execute()

Consulter une facture (Valideur)



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
	bodyValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost := *openapiclient.NewBodyValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost(map[string]interface{}{"key": interface{}(123)}, *openapiclient.NewUtilisateur(int32(123), "Username_example", "Email_example", false)) // BodyValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.ValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost(context.Background()).BodyValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost(bodyValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost).Execute()
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
 **bodyValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost** | [**BodyValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost**](BodyValideurConsulterFactureApiV1ChorusProFacturesValideurConsulterPost.md) |  | 

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

> interface{} ValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost(ctx).BodyValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost(bodyValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost).Execute()

Rechercher factures à valider (Valideur)



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
	bodyValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost := *openapiclient.NewBodyValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost(map[string]interface{}{"key": interface{}(123)}, *openapiclient.NewUtilisateur(int32(123), "Username_example", "Email_example", false)) // BodyValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.ValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost(context.Background()).BodyValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost(bodyValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost).Execute()
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
 **bodyValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost** | [**BodyValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost**](BodyValideurRechercherFacturesApiV1ChorusProFacturesValideurRechercherPost.md) |  | 

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

> interface{} ValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost(ctx).BodyValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost(bodyValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost).Execute()

Valider ou refuser une facture (Valideur)



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
	bodyValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost := *openapiclient.NewBodyValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost(map[string]interface{}{"key": interface{}(123)}, *openapiclient.NewUtilisateur(int32(123), "Username_example", "Email_example", false)) // BodyValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChorusProAPI.ValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost(context.Background()).BodyValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost(bodyValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost).Execute()
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
 **bodyValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost** | [**BodyValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost**](BodyValideurTraiterFactureApiV1ChorusProFacturesValideurTraiterPost.md) |  | 

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


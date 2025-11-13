# \UtilisateurAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ObtenirInfosUtilisateurApiV1MoiGet**](UtilisateurAPI.md#ObtenirInfosUtilisateurApiV1MoiGet) | **Get** /api/v1/moi | Obtenir les informations de l&#39;utilisateur connecté



## ObtenirInfosUtilisateurApiV1MoiGet

> interface{} ObtenirInfosUtilisateurApiV1MoiGet(ctx).Execute()

Obtenir les informations de l'utilisateur connecté



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilisateurAPI.ObtenirInfosUtilisateurApiV1MoiGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilisateurAPI.ObtenirInfosUtilisateurApiV1MoiGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObtenirInfosUtilisateurApiV1MoiGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UtilisateurAPI.ObtenirInfosUtilisateurApiV1MoiGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiObtenirInfosUtilisateurApiV1MoiGetRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


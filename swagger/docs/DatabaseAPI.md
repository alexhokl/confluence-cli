# \DatabaseAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDatabase**](DatabaseAPI.md#CreateDatabase) | **Post** /databases | Create database
[**DeleteDatabase**](DatabaseAPI.md#DeleteDatabase) | **Delete** /databases/{id} | Delete database
[**GetDatabaseById**](DatabaseAPI.md#GetDatabaseById) | **Get** /databases/{id} | Get database by id



## CreateDatabase

> CreateDatabase200Response CreateDatabase(ctx).CreateDatabaseRequest(createDatabaseRequest).Private(private).Execute()

Create database



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/confluence-cli/swagger"
)

func main() {
	createDatabaseRequest := *openapiclient.NewCreateDatabaseRequest("SpaceId_example") // CreateDatabaseRequest | 
	private := true // bool | The database will be private. Only the user who creates this database will have permission to view and edit one. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabaseAPI.CreateDatabase(context.Background()).CreateDatabaseRequest(createDatabaseRequest).Private(private).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabaseAPI.CreateDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDatabase`: CreateDatabase200Response
	fmt.Fprintf(os.Stdout, "Response from `DatabaseAPI.CreateDatabase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDatabaseRequest** | [**CreateDatabaseRequest**](CreateDatabaseRequest.md) |  | 
 **private** | **bool** | The database will be private. Only the user who creates this database will have permission to view and edit one. | [default to false]

### Return type

[**CreateDatabase200Response**](CreateDatabase200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDatabase

> DeleteDatabase(ctx, id).Execute()

Delete database



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/confluence-cli/swagger"
)

func main() {
	id := int64(789) // int64 | The ID of the database to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DatabaseAPI.DeleteDatabase(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabaseAPI.DeleteDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the database to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseById

> CreateDatabase200Response GetDatabaseById(ctx, id).IncludeCollaborators(includeCollaborators).IncludeDirectChildren(includeDirectChildren).IncludeOperations(includeOperations).IncludeProperties(includeProperties).Execute()

Get database by id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/confluence-cli/swagger"
)

func main() {
	id := int64(789) // int64 | The ID of the database to be returned
	includeCollaborators := true // bool | Includes collaborators on the database. (optional) (default to false)
	includeDirectChildren := true // bool | Includes direct children of the database, as defined in the `ChildrenResponse` object. (optional) (default to false)
	includeOperations := true // bool | Includes operations associated with this database in the response, as defined in the `Operation` object. The number of results will be limited to 50 and sorted in the default sort order. A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeProperties := true // bool | Includes content properties associated with this database in the response. The number of results will be limited to 50 and sorted in the default sort order. A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabaseAPI.GetDatabaseById(context.Background(), id).IncludeCollaborators(includeCollaborators).IncludeDirectChildren(includeDirectChildren).IncludeOperations(includeOperations).IncludeProperties(includeProperties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabaseAPI.GetDatabaseById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabaseById`: CreateDatabase200Response
	fmt.Fprintf(os.Stdout, "Response from `DatabaseAPI.GetDatabaseById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the database to be returned | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCollaborators** | **bool** | Includes collaborators on the database. | [default to false]
 **includeDirectChildren** | **bool** | Includes direct children of the database, as defined in the &#x60;ChildrenResponse&#x60; object. | [default to false]
 **includeOperations** | **bool** | Includes operations associated with this database in the response, as defined in the &#x60;Operation&#x60; object. The number of results will be limited to 50 and sorted in the default sort order. A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeProperties** | **bool** | Includes content properties associated with this database in the response. The number of results will be limited to 50 and sorted in the default sort order. A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]

### Return type

[**CreateDatabase200Response**](CreateDatabase200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


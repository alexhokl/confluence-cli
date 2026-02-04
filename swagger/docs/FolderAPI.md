# \FolderAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFolder**](FolderAPI.md#CreateFolder) | **Post** /folders | Create folder
[**DeleteFolder**](FolderAPI.md#DeleteFolder) | **Delete** /folders/{id} | Delete folder
[**GetFolderById**](FolderAPI.md#GetFolderById) | **Get** /folders/{id} | Get folder by id



## CreateFolder

> CreateFolder200Response CreateFolder(ctx).CreateFolderRequest(createFolderRequest).Execute()

Create folder



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
	createFolderRequest := *openapiclient.NewCreateFolderRequest("SpaceId_example") // CreateFolderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FolderAPI.CreateFolder(context.Background()).CreateFolderRequest(createFolderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FolderAPI.CreateFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFolder`: CreateFolder200Response
	fmt.Fprintf(os.Stdout, "Response from `FolderAPI.CreateFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFolderRequest** | [**CreateFolderRequest**](CreateFolderRequest.md) |  | 

### Return type

[**CreateFolder200Response**](CreateFolder200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFolder

> DeleteFolder(ctx, id).Execute()

Delete folder



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
	id := int64(789) // int64 | The ID of the folder to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FolderAPI.DeleteFolder(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FolderAPI.DeleteFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the folder to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFolderRequest struct via the builder pattern


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


## GetFolderById

> CreateFolder200Response GetFolderById(ctx, id).IncludeCollaborators(includeCollaborators).IncludeDirectChildren(includeDirectChildren).IncludeOperations(includeOperations).IncludeProperties(includeProperties).Execute()

Get folder by id



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
	id := int64(789) // int64 | The ID of the folder to be returned.
	includeCollaborators := true // bool | Includes collaborators on the folder. (optional) (default to false)
	includeDirectChildren := true // bool | Includes direct children of the folder, as defined in the `ChildrenResponse` object. (optional) (default to false)
	includeOperations := true // bool | Includes operations associated with this folder in the response, as defined in the `Operation` object. The number of results will be limited to 50 and sorted in the default sort order. A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeProperties := true // bool | Includes content properties associated with this folder in the response. The number of results will be limited to 50 and sorted in the default sort order. A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FolderAPI.GetFolderById(context.Background(), id).IncludeCollaborators(includeCollaborators).IncludeDirectChildren(includeDirectChildren).IncludeOperations(includeOperations).IncludeProperties(includeProperties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FolderAPI.GetFolderById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolderById`: CreateFolder200Response
	fmt.Fprintf(os.Stdout, "Response from `FolderAPI.GetFolderById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the folder to be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCollaborators** | **bool** | Includes collaborators on the folder. | [default to false]
 **includeDirectChildren** | **bool** | Includes direct children of the folder, as defined in the &#x60;ChildrenResponse&#x60; object. | [default to false]
 **includeOperations** | **bool** | Includes operations associated with this folder in the response, as defined in the &#x60;Operation&#x60; object. The number of results will be limited to 50 and sorted in the default sort order. A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeProperties** | **bool** | Includes content properties associated with this folder in the response. The number of results will be limited to 50 and sorted in the default sort order. A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]

### Return type

[**CreateFolder200Response**](CreateFolder200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


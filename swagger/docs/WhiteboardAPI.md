# \WhiteboardAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWhiteboard**](WhiteboardAPI.md#CreateWhiteboard) | **Post** /whiteboards | Create whiteboard
[**DeleteWhiteboard**](WhiteboardAPI.md#DeleteWhiteboard) | **Delete** /whiteboards/{id} | Delete whiteboard
[**GetWhiteboardById**](WhiteboardAPI.md#GetWhiteboardById) | **Get** /whiteboards/{id} | Get whiteboard by id



## CreateWhiteboard

> CreateWhiteboard200Response CreateWhiteboard(ctx).CreateWhiteboardRequest(createWhiteboardRequest).Private(private).Execute()

Create whiteboard



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
	createWhiteboardRequest := *openapiclient.NewCreateWhiteboardRequest("SpaceId_example") // CreateWhiteboardRequest | 
	private := true // bool | The whiteboard will be private. Only the user who creates this whiteboard will have permission to view and edit one. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhiteboardAPI.CreateWhiteboard(context.Background()).CreateWhiteboardRequest(createWhiteboardRequest).Private(private).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhiteboardAPI.CreateWhiteboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWhiteboard`: CreateWhiteboard200Response
	fmt.Fprintf(os.Stdout, "Response from `WhiteboardAPI.CreateWhiteboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWhiteboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWhiteboardRequest** | [**CreateWhiteboardRequest**](CreateWhiteboardRequest.md) |  | 
 **private** | **bool** | The whiteboard will be private. Only the user who creates this whiteboard will have permission to view and edit one. | [default to false]

### Return type

[**CreateWhiteboard200Response**](CreateWhiteboard200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWhiteboard

> DeleteWhiteboard(ctx, id).Execute()

Delete whiteboard



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
	id := int64(789) // int64 | The ID of the whiteboard to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WhiteboardAPI.DeleteWhiteboard(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhiteboardAPI.DeleteWhiteboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the whiteboard to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWhiteboardRequest struct via the builder pattern


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


## GetWhiteboardById

> CreateWhiteboard200Response GetWhiteboardById(ctx, id).IncludeCollaborators(includeCollaborators).IncludeDirectChildren(includeDirectChildren).IncludeOperations(includeOperations).IncludeProperties(includeProperties).Execute()

Get whiteboard by id



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
	id := int64(789) // int64 | The ID of the whiteboard to be returned
	includeCollaborators := true // bool | Includes collaborators on the whiteboard. (optional) (default to false)
	includeDirectChildren := true // bool | Includes direct children of the whiteboard, as defined in the `ChildrenResponse` object. (optional) (default to false)
	includeOperations := true // bool | Includes operations associated with this whiteboard in the response, as defined in the `Operation` object. The number of results will be limited to 50 and sorted in the default sort order. A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeProperties := true // bool | Includes content properties associated with this whiteboard in the response. The number of results will be limited to 50 and sorted in the default sort order. A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhiteboardAPI.GetWhiteboardById(context.Background(), id).IncludeCollaborators(includeCollaborators).IncludeDirectChildren(includeDirectChildren).IncludeOperations(includeOperations).IncludeProperties(includeProperties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhiteboardAPI.GetWhiteboardById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhiteboardById`: CreateWhiteboard200Response
	fmt.Fprintf(os.Stdout, "Response from `WhiteboardAPI.GetWhiteboardById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the whiteboard to be returned | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhiteboardByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCollaborators** | **bool** | Includes collaborators on the whiteboard. | [default to false]
 **includeDirectChildren** | **bool** | Includes direct children of the whiteboard, as defined in the &#x60;ChildrenResponse&#x60; object. | [default to false]
 **includeOperations** | **bool** | Includes operations associated with this whiteboard in the response, as defined in the &#x60;Operation&#x60; object. The number of results will be limited to 50 and sorted in the default sort order. A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeProperties** | **bool** | Includes content properties associated with this whiteboard in the response. The number of results will be limited to 50 and sorted in the default sort order. A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]

### Return type

[**CreateWhiteboard200Response**](CreateWhiteboard200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


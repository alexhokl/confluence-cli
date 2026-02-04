# \SmartLinkAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSmartLink**](SmartLinkAPI.md#CreateSmartLink) | **Post** /embeds | Create Smart Link in the content tree
[**DeleteSmartLink**](SmartLinkAPI.md#DeleteSmartLink) | **Delete** /embeds/{id} | Delete Smart Link in the content tree
[**GetSmartLinkById**](SmartLinkAPI.md#GetSmartLinkById) | **Get** /embeds/{id} | Get Smart Link in the content tree by id



## CreateSmartLink

> CreateSmartLink200Response CreateSmartLink(ctx).CreateSmartLinkRequest(createSmartLinkRequest).Execute()

Create Smart Link in the content tree



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
	createSmartLinkRequest := *openapiclient.NewCreateSmartLinkRequest("SpaceId_example") // CreateSmartLinkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartLinkAPI.CreateSmartLink(context.Background()).CreateSmartLinkRequest(createSmartLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartLinkAPI.CreateSmartLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSmartLink`: CreateSmartLink200Response
	fmt.Fprintf(os.Stdout, "Response from `SmartLinkAPI.CreateSmartLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSmartLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSmartLinkRequest** | [**CreateSmartLinkRequest**](CreateSmartLinkRequest.md) |  | 

### Return type

[**CreateSmartLink200Response**](CreateSmartLink200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSmartLink

> DeleteSmartLink(ctx, id).Execute()

Delete Smart Link in the content tree



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
	id := int64(789) // int64 | The ID of the Smart Link in the content tree to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SmartLinkAPI.DeleteSmartLink(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartLinkAPI.DeleteSmartLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the Smart Link in the content tree to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSmartLinkRequest struct via the builder pattern


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


## GetSmartLinkById

> CreateSmartLink200Response GetSmartLinkById(ctx, id).IncludeCollaborators(includeCollaborators).IncludeDirectChildren(includeDirectChildren).IncludeOperations(includeOperations).IncludeProperties(includeProperties).Execute()

Get Smart Link in the content tree by id



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
	id := int64(789) // int64 | The ID of the Smart Link in the content tree to be returned.
	includeCollaborators := true // bool | Includes collaborators on the Smart Link. (optional) (default to false)
	includeDirectChildren := true // bool | Includes direct children of the Smart Link, as defined in the `ChildrenResponse` object. (optional) (default to false)
	includeOperations := true // bool | Includes operations associated with this Smart Link in the response, as defined in the `Operation` object. The number of results will be limited to 50 and sorted in the default sort order. A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeProperties := true // bool | Includes content properties associated with this Smart Link in the response. The number of results will be limited to 50 and sorted in the default sort order. A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartLinkAPI.GetSmartLinkById(context.Background(), id).IncludeCollaborators(includeCollaborators).IncludeDirectChildren(includeDirectChildren).IncludeOperations(includeOperations).IncludeProperties(includeProperties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartLinkAPI.GetSmartLinkById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSmartLinkById`: CreateSmartLink200Response
	fmt.Fprintf(os.Stdout, "Response from `SmartLinkAPI.GetSmartLinkById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the Smart Link in the content tree to be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmartLinkByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCollaborators** | **bool** | Includes collaborators on the Smart Link. | [default to false]
 **includeDirectChildren** | **bool** | Includes direct children of the Smart Link, as defined in the &#x60;ChildrenResponse&#x60; object. | [default to false]
 **includeOperations** | **bool** | Includes operations associated with this Smart Link in the response, as defined in the &#x60;Operation&#x60; object. The number of results will be limited to 50 and sorted in the default sort order. A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeProperties** | **bool** | Includes content properties associated with this Smart Link in the response. The number of results will be limited to 50 and sorted in the default sort order. A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]

### Return type

[**CreateSmartLink200Response**](CreateSmartLink200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


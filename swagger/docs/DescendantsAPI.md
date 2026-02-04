# \DescendantsAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDatabaseDescendants**](DescendantsAPI.md#GetDatabaseDescendants) | **Get** /databases/{id}/descendants | Get descendants of a database
[**GetFolderDescendants**](DescendantsAPI.md#GetFolderDescendants) | **Get** /folders/{id}/descendants | Get descendants of folder
[**GetPageDescendants**](DescendantsAPI.md#GetPageDescendants) | **Get** /pages/{id}/descendants | Get descendants of page
[**GetSmartLinkDescendants**](DescendantsAPI.md#GetSmartLinkDescendants) | **Get** /embeds/{id}/descendants | Get descendants of a smart link
[**GetWhiteboardDescendants**](DescendantsAPI.md#GetWhiteboardDescendants) | **Get** /whiteboards/{id}/descendants | Get descendants of a whiteboard



## GetDatabaseDescendants

> MultiEntityResultDescendantsResponse GetDatabaseDescendants(ctx, id).Limit(limit).Depth(depth).Cursor(cursor).Execute()

Get descendants of a database



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
	id := int64(789) // int64 | The ID of the database.
	limit := int32(56) // int32 | Maximum number of items per result to return. If more results exist, call the endpoint with the cursor to fetch the next set of results. (optional) (default to 25)
	depth := int32(56) // int32 | Maximum depth of descendants to return. If more results are required, use the endpoint corresponding to the content type of the deepest descendant to fetch more descendants. (optional) (default to 2)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DescendantsAPI.GetDatabaseDescendants(context.Background(), id).Limit(limit).Depth(depth).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DescendantsAPI.GetDatabaseDescendants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabaseDescendants`: MultiEntityResultDescendantsResponse
	fmt.Fprintf(os.Stdout, "Response from `DescendantsAPI.GetDatabaseDescendants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseDescendantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of items per result to return. If more results exist, call the endpoint with the cursor to fetch the next set of results. | [default to 25]
 **depth** | **int32** | Maximum depth of descendants to return. If more results are required, use the endpoint corresponding to the content type of the deepest descendant to fetch more descendants. | [default to 2]
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 

### Return type

[**MultiEntityResultDescendantsResponse**](MultiEntityResultDescendantsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolderDescendants

> MultiEntityResultDescendantsResponse GetFolderDescendants(ctx, id).Limit(limit).Depth(depth).Cursor(cursor).Execute()

Get descendants of folder



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
	id := int64(789) // int64 | The ID of the folder.
	limit := int32(56) // int32 | Maximum number of items per result to return. If more results exist, call the endpoint with the cursor to fetch the next set of results. (optional) (default to 25)
	depth := int32(56) // int32 | Maximum depth of descendants to return. If more results are required, use the endpoint corresponding to the content type of the deepest descendant to fetch more descendants. (optional) (default to 2)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DescendantsAPI.GetFolderDescendants(context.Background(), id).Limit(limit).Depth(depth).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DescendantsAPI.GetFolderDescendants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolderDescendants`: MultiEntityResultDescendantsResponse
	fmt.Fprintf(os.Stdout, "Response from `DescendantsAPI.GetFolderDescendants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the folder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderDescendantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of items per result to return. If more results exist, call the endpoint with the cursor to fetch the next set of results. | [default to 25]
 **depth** | **int32** | Maximum depth of descendants to return. If more results are required, use the endpoint corresponding to the content type of the deepest descendant to fetch more descendants. | [default to 2]
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 

### Return type

[**MultiEntityResultDescendantsResponse**](MultiEntityResultDescendantsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageDescendants

> MultiEntityResultDescendantsResponse GetPageDescendants(ctx, id).Limit(limit).Depth(depth).Cursor(cursor).Execute()

Get descendants of page



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
	id := int64(789) // int64 | The ID of the page.
	limit := int32(56) // int32 | Maximum number of items per result to return. If more results exist, call the endpoint with the cursor to fetch the next set of results. (optional) (default to 25)
	depth := int32(56) // int32 | Maximum depth of descendants to return. If more results are required, use the endpoint corresponding to the content type of the deepest descendant to fetch more descendants. (optional) (default to 2)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DescendantsAPI.GetPageDescendants(context.Background(), id).Limit(limit).Depth(depth).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DescendantsAPI.GetPageDescendants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPageDescendants`: MultiEntityResultDescendantsResponse
	fmt.Fprintf(os.Stdout, "Response from `DescendantsAPI.GetPageDescendants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the page. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageDescendantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of items per result to return. If more results exist, call the endpoint with the cursor to fetch the next set of results. | [default to 25]
 **depth** | **int32** | Maximum depth of descendants to return. If more results are required, use the endpoint corresponding to the content type of the deepest descendant to fetch more descendants. | [default to 2]
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 

### Return type

[**MultiEntityResultDescendantsResponse**](MultiEntityResultDescendantsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmartLinkDescendants

> MultiEntityResultDescendantsResponse GetSmartLinkDescendants(ctx, id).Limit(limit).Depth(depth).Cursor(cursor).Execute()

Get descendants of a smart link



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
	id := int64(789) // int64 | The ID of the smart link.
	limit := int32(56) // int32 | Maximum number of items per result to return. If more results exist, call the endpoint with the cursor to fetch the next set of results. (optional) (default to 25)
	depth := int32(56) // int32 | Maximum depth of descendants to return. If more results are required, use the endpoint corresponding to the content type of the deepest descendant to fetch more descendants. (optional) (default to 2)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DescendantsAPI.GetSmartLinkDescendants(context.Background(), id).Limit(limit).Depth(depth).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DescendantsAPI.GetSmartLinkDescendants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSmartLinkDescendants`: MultiEntityResultDescendantsResponse
	fmt.Fprintf(os.Stdout, "Response from `DescendantsAPI.GetSmartLinkDescendants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the smart link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmartLinkDescendantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of items per result to return. If more results exist, call the endpoint with the cursor to fetch the next set of results. | [default to 25]
 **depth** | **int32** | Maximum depth of descendants to return. If more results are required, use the endpoint corresponding to the content type of the deepest descendant to fetch more descendants. | [default to 2]
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 

### Return type

[**MultiEntityResultDescendantsResponse**](MultiEntityResultDescendantsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhiteboardDescendants

> MultiEntityResultDescendantsResponse GetWhiteboardDescendants(ctx, id).Limit(limit).Depth(depth).Cursor(cursor).Execute()

Get descendants of a whiteboard



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
	id := int64(789) // int64 | The ID of the whiteboard.
	limit := int32(56) // int32 | Maximum number of items per result to return. If more results exist, call the endpoint with the cursor to fetch the next set of results. (optional) (default to 25)
	depth := int32(56) // int32 | Maximum depth of descendants to return. If more results are required, use the endpoint corresponding to the content type of the deepest descendant to fetch more descendants. (optional) (default to 2)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DescendantsAPI.GetWhiteboardDescendants(context.Background(), id).Limit(limit).Depth(depth).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DescendantsAPI.GetWhiteboardDescendants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhiteboardDescendants`: MultiEntityResultDescendantsResponse
	fmt.Fprintf(os.Stdout, "Response from `DescendantsAPI.GetWhiteboardDescendants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the whiteboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhiteboardDescendantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of items per result to return. If more results exist, call the endpoint with the cursor to fetch the next set of results. | [default to 25]
 **depth** | **int32** | Maximum depth of descendants to return. If more results are required, use the endpoint corresponding to the content type of the deepest descendant to fetch more descendants. | [default to 2]
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 

### Return type

[**MultiEntityResultDescendantsResponse**](MultiEntityResultDescendantsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


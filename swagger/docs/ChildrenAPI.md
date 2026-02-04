# \ChildrenAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChildCustomContent**](ChildrenAPI.md#GetChildCustomContent) | **Get** /custom-content/{id}/children | Get child custom content
[**GetChildPages**](ChildrenAPI.md#GetChildPages) | **Get** /pages/{id}/children | Get child pages
[**GetDatabaseDirectChildren**](ChildrenAPI.md#GetDatabaseDirectChildren) | **Get** /databases/{id}/direct-children | Get direct children of a database
[**GetFolderDirectChildren**](ChildrenAPI.md#GetFolderDirectChildren) | **Get** /folders/{id}/direct-children | Get direct children of a folder
[**GetPageDirectChildren**](ChildrenAPI.md#GetPageDirectChildren) | **Get** /pages/{id}/direct-children | Get direct children of a page
[**GetSmartLinkDirectChildren**](ChildrenAPI.md#GetSmartLinkDirectChildren) | **Get** /embeds/{id}/direct-children | Get direct children of a Smart Link
[**GetWhiteboardDirectChildren**](ChildrenAPI.md#GetWhiteboardDirectChildren) | **Get** /whiteboards/{id}/direct-children | Get direct children of a whiteboard



## GetChildCustomContent

> MultiEntityResultChildCustomContent GetChildCustomContent(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).Execute()

Get child custom content



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
	id := int64(789) // int64 | The ID of the parent custom content. If you don't know the custom content ID, use Get custom-content and filter the results.
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of pages per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
	sort := []openapiclient.ChildCustomContentSortOrder{openapiclient.ChildCustomContentSortOrder("created-date")} // []ChildCustomContentSortOrder | Used to sort the result by a particular field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChildrenAPI.GetChildCustomContent(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChildrenAPI.GetChildCustomContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChildCustomContent`: MultiEntityResultChildCustomContent
	fmt.Fprintf(os.Stdout, "Response from `ChildrenAPI.GetChildCustomContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the parent custom content. If you don&#39;t know the custom content ID, use Get custom-content and filter the results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChildCustomContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of pages per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **sort** | [**[]ChildCustomContentSortOrder**](ChildCustomContentSortOrder.md) | Used to sort the result by a particular field. | 

### Return type

[**MultiEntityResultChildCustomContent**](MultiEntityResultChildCustomContent.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChildPages

> MultiEntityResultChildPage GetChildPages(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).Execute()

Get child pages



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
	id := int64(789) // int64 | The ID of the parent page. If you don't know the page ID, use Get pages and filter the results.
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of pages per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
	sort := []openapiclient.ChildPageSortOrder{openapiclient.ChildPageSortOrder("created-date")} // []ChildPageSortOrder | Used to sort the result by a particular field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChildrenAPI.GetChildPages(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChildrenAPI.GetChildPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChildPages`: MultiEntityResultChildPage
	fmt.Fprintf(os.Stdout, "Response from `ChildrenAPI.GetChildPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the parent page. If you don&#39;t know the page ID, use Get pages and filter the results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChildPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of pages per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **sort** | [**[]ChildPageSortOrder**](ChildPageSortOrder.md) | Used to sort the result by a particular field. | 

### Return type

[**MultiEntityResultChildPage**](MultiEntityResultChildPage.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseDirectChildren

> MultiEntityResultChildrenResponse GetDatabaseDirectChildren(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).Execute()

Get direct children of a database



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
	id := int64(789) // int64 | The ID of the parent database.
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of items per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
	sort := []openapiclient.ContentSortOrder{openapiclient.ContentSortOrder("created-date")} // []ContentSortOrder | Used to sort the result by a particular field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChildrenAPI.GetDatabaseDirectChildren(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChildrenAPI.GetDatabaseDirectChildren``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabaseDirectChildren`: MultiEntityResultChildrenResponse
	fmt.Fprintf(os.Stdout, "Response from `ChildrenAPI.GetDatabaseDirectChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the parent database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseDirectChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of items per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **sort** | [**[]ContentSortOrder**](ContentSortOrder.md) | Used to sort the result by a particular field. | 

### Return type

[**MultiEntityResultChildrenResponse**](MultiEntityResultChildrenResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolderDirectChildren

> MultiEntityResultChildrenResponse GetFolderDirectChildren(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).Execute()

Get direct children of a folder



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
	id := int64(789) // int64 | The ID of the parent folder.
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of items per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
	sort := []openapiclient.ContentSortOrder{openapiclient.ContentSortOrder("created-date")} // []ContentSortOrder | Used to sort the result by a particular field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChildrenAPI.GetFolderDirectChildren(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChildrenAPI.GetFolderDirectChildren``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolderDirectChildren`: MultiEntityResultChildrenResponse
	fmt.Fprintf(os.Stdout, "Response from `ChildrenAPI.GetFolderDirectChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the parent folder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderDirectChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of items per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **sort** | [**[]ContentSortOrder**](ContentSortOrder.md) | Used to sort the result by a particular field. | 

### Return type

[**MultiEntityResultChildrenResponse**](MultiEntityResultChildrenResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageDirectChildren

> MultiEntityResultChildrenResponse GetPageDirectChildren(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).Execute()

Get direct children of a page



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
	id := int64(789) // int64 | The ID of the parent page.
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of items per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
	sort := []openapiclient.ContentSortOrder{openapiclient.ContentSortOrder("created-date")} // []ContentSortOrder | Used to sort the result by a particular field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChildrenAPI.GetPageDirectChildren(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChildrenAPI.GetPageDirectChildren``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPageDirectChildren`: MultiEntityResultChildrenResponse
	fmt.Fprintf(os.Stdout, "Response from `ChildrenAPI.GetPageDirectChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the parent page. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageDirectChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of items per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **sort** | [**[]ContentSortOrder**](ContentSortOrder.md) | Used to sort the result by a particular field. | 

### Return type

[**MultiEntityResultChildrenResponse**](MultiEntityResultChildrenResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmartLinkDirectChildren

> MultiEntityResultChildrenResponse GetSmartLinkDirectChildren(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).Execute()

Get direct children of a Smart Link



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
	id := int64(789) // int64 | The ID of the parent smart link.
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of items per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
	sort := []openapiclient.ContentSortOrder{openapiclient.ContentSortOrder("created-date")} // []ContentSortOrder | Used to sort the result by a particular field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChildrenAPI.GetSmartLinkDirectChildren(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChildrenAPI.GetSmartLinkDirectChildren``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSmartLinkDirectChildren`: MultiEntityResultChildrenResponse
	fmt.Fprintf(os.Stdout, "Response from `ChildrenAPI.GetSmartLinkDirectChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the parent smart link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmartLinkDirectChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of items per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **sort** | [**[]ContentSortOrder**](ContentSortOrder.md) | Used to sort the result by a particular field. | 

### Return type

[**MultiEntityResultChildrenResponse**](MultiEntityResultChildrenResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhiteboardDirectChildren

> MultiEntityResultChildrenResponse GetWhiteboardDirectChildren(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).Execute()

Get direct children of a whiteboard



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
	id := int64(789) // int64 | The ID of the parent whiteboard.
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of items per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
	sort := []openapiclient.ContentSortOrder{openapiclient.ContentSortOrder("created-date")} // []ContentSortOrder | Used to sort the result by a particular field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChildrenAPI.GetWhiteboardDirectChildren(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChildrenAPI.GetWhiteboardDirectChildren``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhiteboardDirectChildren`: MultiEntityResultChildrenResponse
	fmt.Fprintf(os.Stdout, "Response from `ChildrenAPI.GetWhiteboardDirectChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the parent whiteboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhiteboardDirectChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of items per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **sort** | [**[]ContentSortOrder**](ContentSortOrder.md) | Used to sort the result by a particular field. | 

### Return type

[**MultiEntityResultChildrenResponse**](MultiEntityResultChildrenResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


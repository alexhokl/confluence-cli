# \LabelAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAttachmentLabels**](LabelAPI.md#GetAttachmentLabels) | **Get** /attachments/{id}/labels | Get labels for attachment
[**GetBlogPostLabels**](LabelAPI.md#GetBlogPostLabels) | **Get** /blogposts/{id}/labels | Get labels for blog post
[**GetCustomContentLabels**](LabelAPI.md#GetCustomContentLabels) | **Get** /custom-content/{id}/labels | Get labels for custom content
[**GetLabels**](LabelAPI.md#GetLabels) | **Get** /labels | Get labels
[**GetPageLabels**](LabelAPI.md#GetPageLabels) | **Get** /pages/{id}/labels | Get labels for page
[**GetSpaceContentLabels**](LabelAPI.md#GetSpaceContentLabels) | **Get** /spaces/{id}/content/labels | Get labels for space content
[**GetSpaceLabels**](LabelAPI.md#GetSpaceLabels) | **Get** /spaces/{id}/labels | Get labels for space



## GetAttachmentLabels

> MultiEntityResultLabel GetAttachmentLabels(ctx, id).Prefix(prefix).Sort(sort).Cursor(cursor).Limit(limit).Execute()

Get labels for attachment



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
	id := int64(789) // int64 | The ID of the attachment for which labels should be returned.
	prefix := "prefix_example" // string | Filter the results to labels based on their prefix. (optional)
	sort := []openapiclient.LabelSortOrder{openapiclient.LabelSortOrder("created-date")} // []LabelSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of labels per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelAPI.GetAttachmentLabels(context.Background(), id).Prefix(prefix).Sort(sort).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelAPI.GetAttachmentLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachmentLabels`: MultiEntityResultLabel
	fmt.Fprintf(os.Stdout, "Response from `LabelAPI.GetAttachmentLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the attachment for which labels should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefix** | **string** | Filter the results to labels based on their prefix. | 
 **sort** | [**[]LabelSortOrder**](LabelSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of labels per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultLabel**](MultiEntityResultLabel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlogPostLabels

> MultiEntityResultLabel GetBlogPostLabels(ctx, id).Prefix(prefix).Sort(sort).Cursor(cursor).Limit(limit).Execute()

Get labels for blog post



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
	id := int64(789) // int64 | The ID of the blog post for which labels should be returned.
	prefix := "prefix_example" // string | Filter the results to labels based on their prefix. (optional)
	sort := []openapiclient.LabelSortOrder{openapiclient.LabelSortOrder("created-date")} // []LabelSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of labels per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelAPI.GetBlogPostLabels(context.Background(), id).Prefix(prefix).Sort(sort).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelAPI.GetBlogPostLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlogPostLabels`: MultiEntityResultLabel
	fmt.Fprintf(os.Stdout, "Response from `LabelAPI.GetBlogPostLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the blog post for which labels should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogPostLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefix** | **string** | Filter the results to labels based on their prefix. | 
 **sort** | [**[]LabelSortOrder**](LabelSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of labels per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultLabel**](MultiEntityResultLabel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomContentLabels

> MultiEntityResultLabel GetCustomContentLabels(ctx, id).Prefix(prefix).Sort(sort).Cursor(cursor).Limit(limit).Execute()

Get labels for custom content



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
	id := int64(789) // int64 | The ID of the custom content for which labels should be returned.
	prefix := "prefix_example" // string | Filter the results to labels based on their prefix. (optional)
	sort := []openapiclient.LabelSortOrder{openapiclient.LabelSortOrder("created-date")} // []LabelSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of labels per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelAPI.GetCustomContentLabels(context.Background(), id).Prefix(prefix).Sort(sort).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelAPI.GetCustomContentLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomContentLabels`: MultiEntityResultLabel
	fmt.Fprintf(os.Stdout, "Response from `LabelAPI.GetCustomContentLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the custom content for which labels should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomContentLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefix** | **string** | Filter the results to labels based on their prefix. | 
 **sort** | [**[]LabelSortOrder**](LabelSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of labels per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultLabel**](MultiEntityResultLabel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLabels

> MultiEntityResultLabel GetLabels(ctx).LabelId(labelId).Prefix(prefix).Cursor(cursor).Sort(sort).Limit(limit).Execute()

Get labels



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
	labelId := []int64{int64(123)} // []int64 | Filters on label ID. Multiple IDs can be specified as a comma-separated list. (optional)
	prefix := []string{"Inner_example"} // []string | Filters on label prefix. Multiple IDs can be specified as a comma-separated list. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	sort := []openapiclient.LabelSortOrder{openapiclient.LabelSortOrder("created-date")} // []LabelSortOrder | Used to sort the result by a particular field. (optional)
	limit := int32(56) // int32 | Maximum number of labels per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelAPI.GetLabels(context.Background()).LabelId(labelId).Prefix(prefix).Cursor(cursor).Sort(sort).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelAPI.GetLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLabels`: MultiEntityResultLabel
	fmt.Fprintf(os.Stdout, "Response from `LabelAPI.GetLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **labelId** | **[]int64** | Filters on label ID. Multiple IDs can be specified as a comma-separated list. | 
 **prefix** | **[]string** | Filters on label prefix. Multiple IDs can be specified as a comma-separated list. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **sort** | [**[]LabelSortOrder**](LabelSortOrder.md) | Used to sort the result by a particular field. | 
 **limit** | **int32** | Maximum number of labels per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultLabel**](MultiEntityResultLabel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageLabels

> MultiEntityResultLabel GetPageLabels(ctx, id).Prefix(prefix).Sort(sort).Cursor(cursor).Limit(limit).Execute()

Get labels for page



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
	id := int64(789) // int64 | The ID of the page for which labels should be returned.
	prefix := "prefix_example" // string | Filter the results to labels based on their prefix. (optional)
	sort := []openapiclient.LabelSortOrder{openapiclient.LabelSortOrder("created-date")} // []LabelSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of labels per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelAPI.GetPageLabels(context.Background(), id).Prefix(prefix).Sort(sort).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelAPI.GetPageLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPageLabels`: MultiEntityResultLabel
	fmt.Fprintf(os.Stdout, "Response from `LabelAPI.GetPageLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the page for which labels should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefix** | **string** | Filter the results to labels based on their prefix. | 
 **sort** | [**[]LabelSortOrder**](LabelSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of labels per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultLabel**](MultiEntityResultLabel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpaceContentLabels

> MultiEntityResultLabel GetSpaceContentLabels(ctx, id).Prefix(prefix).Sort(sort).Cursor(cursor).Limit(limit).Execute()

Get labels for space content



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
	id := int64(789) // int64 | The ID of the space for which labels should be returned.
	prefix := "prefix_example" // string | Filter the results to labels based on their prefix. (optional) (default to "my, team")
	sort := []openapiclient.LabelSortOrder{openapiclient.LabelSortOrder("created-date")} // []LabelSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of labels per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelAPI.GetSpaceContentLabels(context.Background(), id).Prefix(prefix).Sort(sort).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelAPI.GetSpaceContentLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpaceContentLabels`: MultiEntityResultLabel
	fmt.Fprintf(os.Stdout, "Response from `LabelAPI.GetSpaceContentLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the space for which labels should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpaceContentLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefix** | **string** | Filter the results to labels based on their prefix. | [default to &quot;my, team&quot;]
 **sort** | [**[]LabelSortOrder**](LabelSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of labels per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultLabel**](MultiEntityResultLabel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpaceLabels

> MultiEntityResultLabel GetSpaceLabels(ctx, id).Prefix(prefix).Sort(sort).Cursor(cursor).Limit(limit).Execute()

Get labels for space



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
	id := int64(789) // int64 | The ID of the space for which labels should be returned.
	prefix := "prefix_example" // string | Filter the results to labels based on their prefix. (optional) (default to "my, team")
	sort := []openapiclient.LabelSortOrder{openapiclient.LabelSortOrder("created-date")} // []LabelSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of labels per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelAPI.GetSpaceLabels(context.Background(), id).Prefix(prefix).Sort(sort).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelAPI.GetSpaceLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpaceLabels`: MultiEntityResultLabel
	fmt.Fprintf(os.Stdout, "Response from `LabelAPI.GetSpaceLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the space for which labels should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpaceLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefix** | **string** | Filter the results to labels based on their prefix. | [default to &quot;my, team&quot;]
 **sort** | [**[]LabelSortOrder**](LabelSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of labels per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultLabel**](MultiEntityResultLabel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


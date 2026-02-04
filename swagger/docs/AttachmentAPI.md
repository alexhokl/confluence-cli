# \AttachmentAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAttachment**](AttachmentAPI.md#DeleteAttachment) | **Delete** /attachments/{id} | Delete attachment
[**GetAttachmentById**](AttachmentAPI.md#GetAttachmentById) | **Get** /attachments/{id} | Get attachment by id
[**GetAttachments**](AttachmentAPI.md#GetAttachments) | **Get** /attachments | Get attachments
[**GetBlogpostAttachments**](AttachmentAPI.md#GetBlogpostAttachments) | **Get** /blogposts/{id}/attachments | Get attachments for blog post
[**GetCustomContentAttachments**](AttachmentAPI.md#GetCustomContentAttachments) | **Get** /custom-content/{id}/attachments | Get attachments for custom content
[**GetLabelAttachments**](AttachmentAPI.md#GetLabelAttachments) | **Get** /labels/{id}/attachments | Get attachments for label
[**GetPageAttachments**](AttachmentAPI.md#GetPageAttachments) | **Get** /pages/{id}/attachments | Get attachments for page



## DeleteAttachment

> DeleteAttachment(ctx, id).Purge(purge).Execute()

Delete attachment



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
	id := int64(789) // int64 | The ID of the attachment to be deleted.
	purge := true // bool | If attempting to purge the attachment. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AttachmentAPI.DeleteAttachment(context.Background(), id).Purge(purge).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentAPI.DeleteAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the attachment to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purge** | **bool** | If attempting to purge the attachment. | [default to false]

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


## GetAttachmentById

> GetAttachmentById200Response GetAttachmentById(ctx, id).Version(version).IncludeLabels(includeLabels).IncludeProperties(includeProperties).IncludeOperations(includeOperations).IncludeVersions(includeVersions).IncludeVersion(includeVersion).IncludeCollaborators(includeCollaborators).Execute()

Get attachment by id



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
	id := "id_example" // string | The ID of the attachment to be returned. If you don't know the attachment's ID, use Get attachments for page/blogpost/custom content.
	version := int32(56) // int32 | Allows you to retrieve a previously published version. Specify the previous version's number to retrieve its details. (optional)
	includeLabels := true // bool | Includes labels associated with this attachment in the response. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeProperties := true // bool | Includes content properties associated with this attachment in the response. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeOperations := true // bool | Includes operations associated with this attachment in the response, as defined in the `Operation` object. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeVersions := true // bool | Includes versions associated with this attachment in the response. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeVersion := true // bool | Includes the current version associated with this attachment in the response. By default this is included and can be omitted by setting the value to `false`. (optional) (default to true)
	includeCollaborators := true // bool | Includes collaborators on the attachment. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentAPI.GetAttachmentById(context.Background(), id).Version(version).IncludeLabels(includeLabels).IncludeProperties(includeProperties).IncludeOperations(includeOperations).IncludeVersions(includeVersions).IncludeVersion(includeVersion).IncludeCollaborators(includeCollaborators).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentAPI.GetAttachmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachmentById`: GetAttachmentById200Response
	fmt.Fprintf(os.Stdout, "Response from `AttachmentAPI.GetAttachmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the attachment to be returned. If you don&#39;t know the attachment&#39;s ID, use Get attachments for page/blogpost/custom content. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** | Allows you to retrieve a previously published version. Specify the previous version&#39;s number to retrieve its details. | 
 **includeLabels** | **bool** | Includes labels associated with this attachment in the response. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeProperties** | **bool** | Includes content properties associated with this attachment in the response. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeOperations** | **bool** | Includes operations associated with this attachment in the response, as defined in the &#x60;Operation&#x60; object. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeVersions** | **bool** | Includes versions associated with this attachment in the response. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeVersion** | **bool** | Includes the current version associated with this attachment in the response. By default this is included and can be omitted by setting the value to &#x60;false&#x60;. | [default to true]
 **includeCollaborators** | **bool** | Includes collaborators on the attachment. | [default to false]

### Return type

[**GetAttachmentById200Response**](GetAttachmentById200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachments

> MultiEntityResultAttachment GetAttachments(ctx).Sort(sort).Cursor(cursor).Status(status).MediaType(mediaType).Filename(filename).Limit(limit).Execute()

Get attachments



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
	sort := openapiclient.AttachmentSortOrder("created-date") // AttachmentSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	status := []string{"Status_example"} // []string | Filter the results to attachments based on their status. By default, `current` and `archived` are used. (optional)
	mediaType := "mediaType_example" // string | Filters on the mediaType of attachments. Only one may be specified. (optional)
	filename := "filename_example" // string | Filters on the file-name of attachments. Only one may be specified. (optional)
	limit := int32(56) // int32 | Maximum number of attachments per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentAPI.GetAttachments(context.Background()).Sort(sort).Cursor(cursor).Status(status).MediaType(mediaType).Filename(filename).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentAPI.GetAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachments`: MultiEntityResultAttachment
	fmt.Fprintf(os.Stdout, "Response from `AttachmentAPI.GetAttachments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | [**AttachmentSortOrder**](AttachmentSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **status** | **[]string** | Filter the results to attachments based on their status. By default, &#x60;current&#x60; and &#x60;archived&#x60; are used. | 
 **mediaType** | **string** | Filters on the mediaType of attachments. Only one may be specified. | 
 **filename** | **string** | Filters on the file-name of attachments. Only one may be specified. | 
 **limit** | **int32** | Maximum number of attachments per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 50]

### Return type

[**MultiEntityResultAttachment**](MultiEntityResultAttachment.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlogpostAttachments

> MultiEntityResultAttachment GetBlogpostAttachments(ctx, id).Sort(sort).Cursor(cursor).Status(status).MediaType(mediaType).Filename(filename).Limit(limit).Execute()

Get attachments for blog post



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
	id := int64(789) // int64 | The ID of the blog post for which attachments should be returned.
	sort := openapiclient.AttachmentSortOrder("created-date") // AttachmentSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	status := []string{"Status_example"} // []string | Filter the results to attachments based on their status. By default, `current` and `archived` are used. (optional)
	mediaType := "mediaType_example" // string | Filters on the mediaType of attachments. Only one may be specified. (optional)
	filename := "filename_example" // string | Filters on the file-name of attachments. Only one may be specified. (optional)
	limit := int32(56) // int32 | Maximum number of attachments per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentAPI.GetBlogpostAttachments(context.Background(), id).Sort(sort).Cursor(cursor).Status(status).MediaType(mediaType).Filename(filename).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentAPI.GetBlogpostAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlogpostAttachments`: MultiEntityResultAttachment
	fmt.Fprintf(os.Stdout, "Response from `AttachmentAPI.GetBlogpostAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the blog post for which attachments should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogpostAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | [**AttachmentSortOrder**](AttachmentSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **status** | **[]string** | Filter the results to attachments based on their status. By default, &#x60;current&#x60; and &#x60;archived&#x60; are used. | 
 **mediaType** | **string** | Filters on the mediaType of attachments. Only one may be specified. | 
 **filename** | **string** | Filters on the file-name of attachments. Only one may be specified. | 
 **limit** | **int32** | Maximum number of attachments per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 50]

### Return type

[**MultiEntityResultAttachment**](MultiEntityResultAttachment.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomContentAttachments

> MultiEntityResultAttachment GetCustomContentAttachments(ctx, id).Sort(sort).Cursor(cursor).Status(status).MediaType(mediaType).Filename(filename).Limit(limit).Execute()

Get attachments for custom content



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
	id := int64(789) // int64 | The ID of the custom content for which attachments should be returned.
	sort := openapiclient.AttachmentSortOrder("created-date") // AttachmentSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	status := []string{"Status_example"} // []string | Filter the results to attachments based on their status. By default, `current` and `archived` are used. (optional)
	mediaType := "mediaType_example" // string | Filters on the mediaType of attachments. Only one may be specified. (optional)
	filename := "filename_example" // string | Filters on the file-name of attachments. Only one may be specified. (optional)
	limit := int32(56) // int32 | Maximum number of attachments per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentAPI.GetCustomContentAttachments(context.Background(), id).Sort(sort).Cursor(cursor).Status(status).MediaType(mediaType).Filename(filename).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentAPI.GetCustomContentAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomContentAttachments`: MultiEntityResultAttachment
	fmt.Fprintf(os.Stdout, "Response from `AttachmentAPI.GetCustomContentAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the custom content for which attachments should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomContentAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | [**AttachmentSortOrder**](AttachmentSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **status** | **[]string** | Filter the results to attachments based on their status. By default, &#x60;current&#x60; and &#x60;archived&#x60; are used. | 
 **mediaType** | **string** | Filters on the mediaType of attachments. Only one may be specified. | 
 **filename** | **string** | Filters on the file-name of attachments. Only one may be specified. | 
 **limit** | **int32** | Maximum number of attachments per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 50]

### Return type

[**MultiEntityResultAttachment**](MultiEntityResultAttachment.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLabelAttachments

> MultiEntityResultAttachment GetLabelAttachments(ctx, id).Sort(sort).Cursor(cursor).Limit(limit).Execute()

Get attachments for label



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
	id := int64(789) // int64 | The ID of the label for which attachments should be returned.
	sort := openapiclient.AttachmentSortOrder("created-date") // AttachmentSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of pages per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentAPI.GetLabelAttachments(context.Background(), id).Sort(sort).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentAPI.GetLabelAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLabelAttachments`: MultiEntityResultAttachment
	fmt.Fprintf(os.Stdout, "Response from `AttachmentAPI.GetLabelAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the label for which attachments should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLabelAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | [**AttachmentSortOrder**](AttachmentSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of pages per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultAttachment**](MultiEntityResultAttachment.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageAttachments

> MultiEntityResultAttachment GetPageAttachments(ctx, id).Sort(sort).Cursor(cursor).Status(status).MediaType(mediaType).Filename(filename).Limit(limit).Execute()

Get attachments for page



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
	id := int64(789) // int64 | The ID of the page for which attachments should be returned.
	sort := openapiclient.AttachmentSortOrder("created-date") // AttachmentSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	status := []string{"Status_example"} // []string | Filter the results to attachments based on their status. By default, `current` and `archived` are used. (optional)
	mediaType := "mediaType_example" // string | Filters on the mediaType of attachments. Only one may be specified. (optional)
	filename := "filename_example" // string | Filters on the file-name of attachments. Only one may be specified. (optional)
	limit := int32(56) // int32 | Maximum number of attachments per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentAPI.GetPageAttachments(context.Background(), id).Sort(sort).Cursor(cursor).Status(status).MediaType(mediaType).Filename(filename).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentAPI.GetPageAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPageAttachments`: MultiEntityResultAttachment
	fmt.Fprintf(os.Stdout, "Response from `AttachmentAPI.GetPageAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the page for which attachments should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | [**AttachmentSortOrder**](AttachmentSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **status** | **[]string** | Filter the results to attachments based on their status. By default, &#x60;current&#x60; and &#x60;archived&#x60; are used. | 
 **mediaType** | **string** | Filters on the mediaType of attachments. Only one may be specified. | 
 **filename** | **string** | Filters on the file-name of attachments. Only one may be specified. | 
 **limit** | **int32** | Maximum number of attachments per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 50]

### Return type

[**MultiEntityResultAttachment**](MultiEntityResultAttachment.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


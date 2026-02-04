# \VersionAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAttachmentVersionDetails**](VersionAPI.md#GetAttachmentVersionDetails) | **Get** /attachments/{attachment-id}/versions/{version-number} | Get version details for attachment version
[**GetAttachmentVersions**](VersionAPI.md#GetAttachmentVersions) | **Get** /attachments/{id}/versions | Get attachment versions
[**GetBlogPostVersionDetails**](VersionAPI.md#GetBlogPostVersionDetails) | **Get** /blogposts/{blogpost-id}/versions/{version-number} | Get version details for blog post version
[**GetBlogPostVersions**](VersionAPI.md#GetBlogPostVersions) | **Get** /blogposts/{id}/versions | Get blog post versions
[**GetCustomContentVersionDetails**](VersionAPI.md#GetCustomContentVersionDetails) | **Get** /custom-content/{custom-content-id}/versions/{version-number} | Get version details for custom content version
[**GetCustomContentVersions**](VersionAPI.md#GetCustomContentVersions) | **Get** /custom-content/{custom-content-id}/versions | Get custom content versions
[**GetFooterCommentVersionDetails**](VersionAPI.md#GetFooterCommentVersionDetails) | **Get** /footer-comments/{id}/versions/{version-number} | Get version details for footer comment version
[**GetFooterCommentVersions**](VersionAPI.md#GetFooterCommentVersions) | **Get** /footer-comments/{id}/versions | Get footer comment versions
[**GetInlineCommentVersionDetails**](VersionAPI.md#GetInlineCommentVersionDetails) | **Get** /inline-comments/{id}/versions/{version-number} | Get version details for inline comment version
[**GetInlineCommentVersions**](VersionAPI.md#GetInlineCommentVersions) | **Get** /inline-comments/{id}/versions | Get inline comment versions
[**GetPageVersionDetails**](VersionAPI.md#GetPageVersionDetails) | **Get** /pages/{page-id}/versions/{version-number} | Get version details for page version
[**GetPageVersions**](VersionAPI.md#GetPageVersions) | **Get** /pages/{id}/versions | Get page versions



## GetAttachmentVersionDetails

> DetailedVersion GetAttachmentVersionDetails(ctx, attachmentId, versionNumber).Execute()

Get version details for attachment version



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
	attachmentId := "attachmentId_example" // string | The ID of the attachment for which version details should be returned.
	versionNumber := int64(789) // int64 | The version number of the attachment to be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VersionAPI.GetAttachmentVersionDetails(context.Background(), attachmentId, versionNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionAPI.GetAttachmentVersionDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachmentVersionDetails`: DetailedVersion
	fmt.Fprintf(os.Stdout, "Response from `VersionAPI.GetAttachmentVersionDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** | The ID of the attachment for which version details should be returned. | 
**versionNumber** | **int64** | The version number of the attachment to be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentVersionDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DetailedVersion**](DetailedVersion.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachmentVersions

> MultiEntityResultVersion GetAttachmentVersions(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).Execute()

Get attachment versions



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
	id := "id_example" // string | The ID of the attachment to be queried for its versions. If you don't know the attachment ID, use Get attachments and filter the results.
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of versions per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
	sort := openapiclient.VersionSortOrder("modified-date") // VersionSortOrder | Used to sort the result by a particular field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VersionAPI.GetAttachmentVersions(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionAPI.GetAttachmentVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachmentVersions`: MultiEntityResultVersion
	fmt.Fprintf(os.Stdout, "Response from `VersionAPI.GetAttachmentVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the attachment to be queried for its versions. If you don&#39;t know the attachment ID, use Get attachments and filter the results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of versions per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **sort** | [**VersionSortOrder**](VersionSortOrder.md) | Used to sort the result by a particular field. | 

### Return type

[**MultiEntityResultVersion**](MultiEntityResultVersion.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlogPostVersionDetails

> DetailedVersion GetBlogPostVersionDetails(ctx, blogpostId, versionNumber).Execute()

Get version details for blog post version



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
	blogpostId := int64(789) // int64 | The ID of the blog post for which version details should be returned.
	versionNumber := int64(789) // int64 | The version number of the blog post to be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VersionAPI.GetBlogPostVersionDetails(context.Background(), blogpostId, versionNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionAPI.GetBlogPostVersionDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlogPostVersionDetails`: DetailedVersion
	fmt.Fprintf(os.Stdout, "Response from `VersionAPI.GetBlogPostVersionDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogpostId** | **int64** | The ID of the blog post for which version details should be returned. | 
**versionNumber** | **int64** | The version number of the blog post to be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogPostVersionDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DetailedVersion**](DetailedVersion.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlogPostVersions

> MultiEntityResultVersion1 GetBlogPostVersions(ctx, id).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).Sort(sort).Execute()

Get blog post versions



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
	id := int64(789) // int64 | The ID of the blog post to be queried for its versions. If you don't know the blog post ID, use Get blog posts and filter the results.
	bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of versions per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
	sort := openapiclient.VersionSortOrder("modified-date") // VersionSortOrder | Used to sort the result by a particular field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VersionAPI.GetBlogPostVersions(context.Background(), id).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionAPI.GetBlogPostVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlogPostVersions`: MultiEntityResultVersion1
	fmt.Fprintf(os.Stdout, "Response from `VersionAPI.GetBlogPostVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the blog post to be queried for its versions. If you don&#39;t know the blog post ID, use Get blog posts and filter the results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogPostVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of versions per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **sort** | [**VersionSortOrder**](VersionSortOrder.md) | Used to sort the result by a particular field. | 

### Return type

[**MultiEntityResultVersion1**](MultiEntityResultVersion1.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomContentVersionDetails

> DetailedVersion GetCustomContentVersionDetails(ctx, customContentId, versionNumber).Execute()

Get version details for custom content version



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
	customContentId := int64(789) // int64 | The ID of the custom content for which version details should be returned.
	versionNumber := int64(789) // int64 | The version number of the custom content to be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VersionAPI.GetCustomContentVersionDetails(context.Background(), customContentId, versionNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionAPI.GetCustomContentVersionDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomContentVersionDetails`: DetailedVersion
	fmt.Fprintf(os.Stdout, "Response from `VersionAPI.GetCustomContentVersionDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customContentId** | **int64** | The ID of the custom content for which version details should be returned. | 
**versionNumber** | **int64** | The version number of the custom content to be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomContentVersionDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DetailedVersion**](DetailedVersion.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomContentVersions

> MultiEntityResultVersion3 GetCustomContentVersions(ctx, customContentId).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).Sort(sort).Execute()

Get custom content versions



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
	customContentId := int64(789) // int64 | The ID of the custom content to be queried for its versions. If you don't know the custom content ID, use Get custom-content by type and filter the results.
	bodyFormat := openapiclient.CustomContentBodyRepresentation("raw") // CustomContentBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field.  Note: If the custom content body type is `storage`, the `storage` and `atlas_doc_format` body formats are able to be returned. If the custom content body type is `raw`, only the `raw` body format is able to be returned. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of versions per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
	sort := openapiclient.VersionSortOrder("modified-date") // VersionSortOrder | Used to sort the result by a particular field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VersionAPI.GetCustomContentVersions(context.Background(), customContentId).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionAPI.GetCustomContentVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomContentVersions`: MultiEntityResultVersion3
	fmt.Fprintf(os.Stdout, "Response from `VersionAPI.GetCustomContentVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customContentId** | **int64** | The ID of the custom content to be queried for its versions. If you don&#39;t know the custom content ID, use Get custom-content by type and filter the results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomContentVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**CustomContentBodyRepresentation**](CustomContentBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field.  Note: If the custom content body type is &#x60;storage&#x60;, the &#x60;storage&#x60; and &#x60;atlas_doc_format&#x60; body formats are able to be returned. If the custom content body type is &#x60;raw&#x60;, only the &#x60;raw&#x60; body format is able to be returned. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of versions per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **sort** | [**VersionSortOrder**](VersionSortOrder.md) | Used to sort the result by a particular field. | 

### Return type

[**MultiEntityResultVersion3**](MultiEntityResultVersion3.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFooterCommentVersionDetails

> DetailedVersion GetFooterCommentVersionDetails(ctx, id, versionNumber).Execute()

Get version details for footer comment version



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
	id := int64(789) // int64 | The ID of the footer comment for which version details should be returned.
	versionNumber := int64(789) // int64 | The version number of the footer comment to be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VersionAPI.GetFooterCommentVersionDetails(context.Background(), id, versionNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionAPI.GetFooterCommentVersionDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFooterCommentVersionDetails`: DetailedVersion
	fmt.Fprintf(os.Stdout, "Response from `VersionAPI.GetFooterCommentVersionDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the footer comment for which version details should be returned. | 
**versionNumber** | **int64** | The version number of the footer comment to be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFooterCommentVersionDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DetailedVersion**](DetailedVersion.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFooterCommentVersions

> MultiEntityResultVersion4 GetFooterCommentVersions(ctx, id).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).Sort(sort).Execute()

Get footer comment versions



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
	id := int64(789) // int64 | The ID of the footer comment for which versions should be returned
	bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of versions per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
	sort := openapiclient.VersionSortOrder("modified-date") // VersionSortOrder | Used to sort the result by a particular field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VersionAPI.GetFooterCommentVersions(context.Background(), id).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionAPI.GetFooterCommentVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFooterCommentVersions`: MultiEntityResultVersion4
	fmt.Fprintf(os.Stdout, "Response from `VersionAPI.GetFooterCommentVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the footer comment for which versions should be returned | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFooterCommentVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of versions per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **sort** | [**VersionSortOrder**](VersionSortOrder.md) | Used to sort the result by a particular field. | 

### Return type

[**MultiEntityResultVersion4**](MultiEntityResultVersion4.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInlineCommentVersionDetails

> DetailedVersion GetInlineCommentVersionDetails(ctx, id, versionNumber).Execute()

Get version details for inline comment version



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
	id := int64(789) // int64 | The ID of the inline comment for which version details should be returned.
	versionNumber := int64(789) // int64 | The version number of the inline comment to be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VersionAPI.GetInlineCommentVersionDetails(context.Background(), id, versionNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionAPI.GetInlineCommentVersionDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInlineCommentVersionDetails`: DetailedVersion
	fmt.Fprintf(os.Stdout, "Response from `VersionAPI.GetInlineCommentVersionDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the inline comment for which version details should be returned. | 
**versionNumber** | **int64** | The version number of the inline comment to be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInlineCommentVersionDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DetailedVersion**](DetailedVersion.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInlineCommentVersions

> MultiEntityResultVersion4 GetInlineCommentVersions(ctx, id).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).Sort(sort).Execute()

Get inline comment versions



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
	id := int64(789) // int64 | The ID of the inline comment for which versions should be returned
	bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of versions per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
	sort := openapiclient.VersionSortOrder("modified-date") // VersionSortOrder | Used to sort the result by a particular field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VersionAPI.GetInlineCommentVersions(context.Background(), id).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionAPI.GetInlineCommentVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInlineCommentVersions`: MultiEntityResultVersion4
	fmt.Fprintf(os.Stdout, "Response from `VersionAPI.GetInlineCommentVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the inline comment for which versions should be returned | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInlineCommentVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of versions per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **sort** | [**VersionSortOrder**](VersionSortOrder.md) | Used to sort the result by a particular field. | 

### Return type

[**MultiEntityResultVersion4**](MultiEntityResultVersion4.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageVersionDetails

> DetailedVersion GetPageVersionDetails(ctx, pageId, versionNumber).Execute()

Get version details for page version



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
	pageId := int64(789) // int64 | The ID of the page for which version details should be returned.
	versionNumber := int64(789) // int64 | The version number of the page to be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VersionAPI.GetPageVersionDetails(context.Background(), pageId, versionNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionAPI.GetPageVersionDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPageVersionDetails`: DetailedVersion
	fmt.Fprintf(os.Stdout, "Response from `VersionAPI.GetPageVersionDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **int64** | The ID of the page for which version details should be returned. | 
**versionNumber** | **int64** | The version number of the page to be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageVersionDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DetailedVersion**](DetailedVersion.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageVersions

> MultiEntityResultVersion2 GetPageVersions(ctx, id).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).Sort(sort).Execute()

Get page versions



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
	id := int64(789) // int64 | The ID of the page to be queried for its versions. If you don't know the page ID, use Get pages and filter the results.
	bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of versions per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
	sort := openapiclient.VersionSortOrder("modified-date") // VersionSortOrder | Used to sort the result by a particular field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VersionAPI.GetPageVersions(context.Background(), id).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionAPI.GetPageVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPageVersions`: MultiEntityResultVersion2
	fmt.Fprintf(os.Stdout, "Response from `VersionAPI.GetPageVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the page to be queried for its versions. If you don&#39;t know the page ID, use Get pages and filter the results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of versions per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **sort** | [**VersionSortOrder**](VersionSortOrder.md) | Used to sort the result by a particular field. | 

### Return type

[**MultiEntityResultVersion2**](MultiEntityResultVersion2.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


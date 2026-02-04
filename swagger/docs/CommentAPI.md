# \CommentAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFooterComment**](CommentAPI.md#CreateFooterComment) | **Post** /footer-comments | Create footer comment
[**CreateInlineComment**](CommentAPI.md#CreateInlineComment) | **Post** /inline-comments | Create inline comment
[**DeleteFooterComment**](CommentAPI.md#DeleteFooterComment) | **Delete** /footer-comments/{comment-id} | Delete footer comment
[**DeleteInlineComment**](CommentAPI.md#DeleteInlineComment) | **Delete** /inline-comments/{comment-id} | Delete inline comment
[**GetAttachmentComments**](CommentAPI.md#GetAttachmentComments) | **Get** /attachments/{id}/footer-comments | Get attachment comments
[**GetBlogPostFooterComments**](CommentAPI.md#GetBlogPostFooterComments) | **Get** /blogposts/{id}/footer-comments | Get footer comments for blog post
[**GetBlogPostInlineComments**](CommentAPI.md#GetBlogPostInlineComments) | **Get** /blogposts/{id}/inline-comments | Get inline comments for blog post
[**GetCustomContentComments**](CommentAPI.md#GetCustomContentComments) | **Get** /custom-content/{id}/footer-comments | Get custom content comments
[**GetFooterCommentById**](CommentAPI.md#GetFooterCommentById) | **Get** /footer-comments/{comment-id} | Get footer comment by id
[**GetFooterCommentChildren**](CommentAPI.md#GetFooterCommentChildren) | **Get** /footer-comments/{id}/children | Get children footer comments
[**GetFooterComments**](CommentAPI.md#GetFooterComments) | **Get** /footer-comments | Get footer comments
[**GetInlineCommentById**](CommentAPI.md#GetInlineCommentById) | **Get** /inline-comments/{comment-id} | Get inline comment by id
[**GetInlineCommentChildren**](CommentAPI.md#GetInlineCommentChildren) | **Get** /inline-comments/{id}/children | Get children inline comments
[**GetInlineComments**](CommentAPI.md#GetInlineComments) | **Get** /inline-comments | Get inline comments
[**GetPageFooterComments**](CommentAPI.md#GetPageFooterComments) | **Get** /pages/{id}/footer-comments | Get footer comments for page
[**GetPageInlineComments**](CommentAPI.md#GetPageInlineComments) | **Get** /pages/{id}/inline-comments | Get inline comments for page
[**UpdateFooterComment**](CommentAPI.md#UpdateFooterComment) | **Put** /footer-comments/{comment-id} | Update footer comment
[**UpdateInlineComment**](CommentAPI.md#UpdateInlineComment) | **Put** /inline-comments/{comment-id} | Update inline comment



## CreateFooterComment

> CreateFooterComment201Response CreateFooterComment(ctx).CreateFooterCommentModel(createFooterCommentModel).Execute()

Create footer comment



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
	createFooterCommentModel := *openapiclient.NewCreateFooterCommentModel() // CreateFooterCommentModel | The footer comment to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentAPI.CreateFooterComment(context.Background()).CreateFooterCommentModel(createFooterCommentModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAPI.CreateFooterComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFooterComment`: CreateFooterComment201Response
	fmt.Fprintf(os.Stdout, "Response from `CommentAPI.CreateFooterComment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFooterCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFooterCommentModel** | [**CreateFooterCommentModel**](CreateFooterCommentModel.md) | The footer comment to be created | 

### Return type

[**CreateFooterComment201Response**](CreateFooterComment201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInlineComment

> CreateInlineComment201Response CreateInlineComment(ctx).CreateInlineCommentModel(createInlineCommentModel).Execute()

Create inline comment



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
	createInlineCommentModel := *openapiclient.NewCreateInlineCommentModel() // CreateInlineCommentModel | The inline comment to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentAPI.CreateInlineComment(context.Background()).CreateInlineCommentModel(createInlineCommentModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAPI.CreateInlineComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInlineComment`: CreateInlineComment201Response
	fmt.Fprintf(os.Stdout, "Response from `CommentAPI.CreateInlineComment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInlineCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInlineCommentModel** | [**CreateInlineCommentModel**](CreateInlineCommentModel.md) | The inline comment to be created | 

### Return type

[**CreateInlineComment201Response**](CreateInlineComment201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFooterComment

> DeleteFooterComment(ctx, commentId).Execute()

Delete footer comment



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
	commentId := int64(789) // int64 | The ID of the comment to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommentAPI.DeleteFooterComment(context.Background(), commentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAPI.DeleteFooterComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int64** | The ID of the comment to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFooterCommentRequest struct via the builder pattern


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


## DeleteInlineComment

> DeleteInlineComment(ctx, commentId).Execute()

Delete inline comment



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
	commentId := int64(789) // int64 | The ID of the comment to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommentAPI.DeleteInlineComment(context.Background(), commentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAPI.DeleteInlineComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int64** | The ID of the comment to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInlineCommentRequest struct via the builder pattern


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


## GetAttachmentComments

> MultiEntityResultAttachmentCommentModel GetAttachmentComments(ctx, id).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).Sort(sort).Version(version).Execute()

Get attachment comments



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
	id := "id_example" // string | The ID of the attachment for which comments should be returned.
	bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format type to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of comments per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
	sort := openapiclient.CommentSortOrder("created-date") // CommentSortOrder | Used to sort the result by a particular field. (optional)
	version := int64(789) // int64 | Version number of the attachment to retrieve comments for. If no version provided, retrieves comments for the latest version. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentAPI.GetAttachmentComments(context.Background(), id).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).Sort(sort).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAPI.GetAttachmentComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachmentComments`: MultiEntityResultAttachmentCommentModel
	fmt.Fprintf(os.Stdout, "Response from `CommentAPI.GetAttachmentComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the attachment for which comments should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format type to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of comments per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **sort** | [**CommentSortOrder**](CommentSortOrder.md) | Used to sort the result by a particular field. | 
 **version** | **int64** | Version number of the attachment to retrieve comments for. If no version provided, retrieves comments for the latest version. | 

### Return type

[**MultiEntityResultAttachmentCommentModel**](MultiEntityResultAttachmentCommentModel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlogPostFooterComments

> MultiEntityResultBlogPostCommentModel GetBlogPostFooterComments(ctx, id).BodyFormat(bodyFormat).Status(status).Sort(sort).Cursor(cursor).Limit(limit).Execute()

Get footer comments for blog post



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
	id := int64(789) // int64 | The ID of the blog post for which footer comments should be returned.
	bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format type to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	status := []string{"Status_example"} // []string | Filter the footer comment being retrieved by its status. (optional)
	sort := openapiclient.CommentSortOrder("created-date") // CommentSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of footer comments per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentAPI.GetBlogPostFooterComments(context.Background(), id).BodyFormat(bodyFormat).Status(status).Sort(sort).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAPI.GetBlogPostFooterComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlogPostFooterComments`: MultiEntityResultBlogPostCommentModel
	fmt.Fprintf(os.Stdout, "Response from `CommentAPI.GetBlogPostFooterComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the blog post for which footer comments should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogPostFooterCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format type to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **status** | **[]string** | Filter the footer comment being retrieved by its status. | 
 **sort** | [**CommentSortOrder**](CommentSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of footer comments per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultBlogPostCommentModel**](MultiEntityResultBlogPostCommentModel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlogPostInlineComments

> MultiEntityResultBlogPostInlineCommentModel GetBlogPostInlineComments(ctx, id).BodyFormat(bodyFormat).Status(status).ResolutionStatus(resolutionStatus).Sort(sort).Cursor(cursor).Limit(limit).Execute()

Get inline comments for blog post



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
	id := int64(789) // int64 | The ID of the blog post for which inline comments should be returned.
	bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format type to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	status := []string{"Status_example"} // []string | Filter the inline comment being retrieved by its status. (optional)
	resolutionStatus := []string{"ResolutionStatus_example"} // []string | Filter the inline comment being retrieved by its resolution status. (optional)
	sort := openapiclient.CommentSortOrder("created-date") // CommentSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of inline comments per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentAPI.GetBlogPostInlineComments(context.Background(), id).BodyFormat(bodyFormat).Status(status).ResolutionStatus(resolutionStatus).Sort(sort).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAPI.GetBlogPostInlineComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlogPostInlineComments`: MultiEntityResultBlogPostInlineCommentModel
	fmt.Fprintf(os.Stdout, "Response from `CommentAPI.GetBlogPostInlineComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the blog post for which inline comments should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogPostInlineCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format type to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **status** | **[]string** | Filter the inline comment being retrieved by its status. | 
 **resolutionStatus** | **[]string** | Filter the inline comment being retrieved by its resolution status. | 
 **sort** | [**CommentSortOrder**](CommentSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of inline comments per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultBlogPostInlineCommentModel**](MultiEntityResultBlogPostInlineCommentModel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomContentComments

> MultiEntityResultCustomContentCommentModel GetCustomContentComments(ctx, id).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).Sort(sort).Execute()

Get custom content comments



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
	id := int64(789) // int64 | The ID of the custom content for which comments should be returned.
	bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format type to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of comments per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
	sort := openapiclient.CommentSortOrder("created-date") // CommentSortOrder | Used to sort the result by a particular field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentAPI.GetCustomContentComments(context.Background(), id).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAPI.GetCustomContentComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomContentComments`: MultiEntityResultCustomContentCommentModel
	fmt.Fprintf(os.Stdout, "Response from `CommentAPI.GetCustomContentComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the custom content for which comments should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomContentCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format type to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of comments per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **sort** | [**CommentSortOrder**](CommentSortOrder.md) | Used to sort the result by a particular field. | 

### Return type

[**MultiEntityResultCustomContentCommentModel**](MultiEntityResultCustomContentCommentModel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFooterCommentById

> CreateFooterComment201Response GetFooterCommentById(ctx, commentId).BodyFormat(bodyFormat).Version(version).IncludeProperties(includeProperties).IncludeOperations(includeOperations).IncludeLikes(includeLikes).IncludeVersions(includeVersions).IncludeVersion(includeVersion).Execute()

Get footer comment by id



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
	commentId := int64(789) // int64 | The ID of the comment to be retrieved.
	bodyFormat := openapiclient.PrimaryBodyRepresentationSingle("storage") // PrimaryBodyRepresentationSingle | The content format type to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	version := int32(56) // int32 | Allows you to retrieve a previously published version. Specify the previous version's number to retrieve its details. (optional)
	includeProperties := true // bool | Includes content properties associated with this footer comment in the response. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeOperations := true // bool | Includes operations associated with this footer comment in the response, as defined in the `Operation` object. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeLikes := true // bool | Includes likes associated with this footer comment in the response. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeVersions := true // bool | Includes versions associated with this footer comment in the response. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeVersion := true // bool | Includes the current version associated with this footer comment in the response. By default this is included and can be omitted by setting the value to `false`. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentAPI.GetFooterCommentById(context.Background(), commentId).BodyFormat(bodyFormat).Version(version).IncludeProperties(includeProperties).IncludeOperations(includeOperations).IncludeLikes(includeLikes).IncludeVersions(includeVersions).IncludeVersion(includeVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAPI.GetFooterCommentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFooterCommentById`: CreateFooterComment201Response
	fmt.Fprintf(os.Stdout, "Response from `CommentAPI.GetFooterCommentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int64** | The ID of the comment to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFooterCommentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**PrimaryBodyRepresentationSingle**](PrimaryBodyRepresentationSingle.md) | The content format type to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **version** | **int32** | Allows you to retrieve a previously published version. Specify the previous version&#39;s number to retrieve its details. | 
 **includeProperties** | **bool** | Includes content properties associated with this footer comment in the response. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeOperations** | **bool** | Includes operations associated with this footer comment in the response, as defined in the &#x60;Operation&#x60; object. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeLikes** | **bool** | Includes likes associated with this footer comment in the response. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeVersions** | **bool** | Includes versions associated with this footer comment in the response. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeVersion** | **bool** | Includes the current version associated with this footer comment in the response. By default this is included and can be omitted by setting the value to &#x60;false&#x60;. | [default to true]

### Return type

[**CreateFooterComment201Response**](CreateFooterComment201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFooterCommentChildren

> MultiEntityResultChildrenCommentModel GetFooterCommentChildren(ctx, id).BodyFormat(bodyFormat).Sort(sort).Cursor(cursor).Limit(limit).Execute()

Get children footer comments



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
	id := int64(789) // int64 | The ID of the parent comment for which footer comment children should be returned.
	bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format type to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	sort := openapiclient.CommentSortOrder("created-date") // CommentSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of footer comments per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentAPI.GetFooterCommentChildren(context.Background(), id).BodyFormat(bodyFormat).Sort(sort).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAPI.GetFooterCommentChildren``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFooterCommentChildren`: MultiEntityResultChildrenCommentModel
	fmt.Fprintf(os.Stdout, "Response from `CommentAPI.GetFooterCommentChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the parent comment for which footer comment children should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFooterCommentChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format type to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **sort** | [**CommentSortOrder**](CommentSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of footer comments per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultChildrenCommentModel**](MultiEntityResultChildrenCommentModel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFooterComments

> MultiEntityResultFooterCommentModel GetFooterComments(ctx).BodyFormat(bodyFormat).Sort(sort).Cursor(cursor).Limit(limit).Execute()

Get footer comments



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
	bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format type to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	sort := openapiclient.CommentSortOrder("created-date") // CommentSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of footer comments per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentAPI.GetFooterComments(context.Background()).BodyFormat(bodyFormat).Sort(sort).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAPI.GetFooterComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFooterComments`: MultiEntityResultFooterCommentModel
	fmt.Fprintf(os.Stdout, "Response from `CommentAPI.GetFooterComments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFooterCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format type to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **sort** | [**CommentSortOrder**](CommentSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of footer comments per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultFooterCommentModel**](MultiEntityResultFooterCommentModel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInlineCommentById

> CreateInlineComment201Response GetInlineCommentById(ctx, commentId).BodyFormat(bodyFormat).Version(version).IncludeProperties(includeProperties).IncludeOperations(includeOperations).IncludeLikes(includeLikes).IncludeVersions(includeVersions).IncludeVersion(includeVersion).Execute()

Get inline comment by id



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
	commentId := int64(789) // int64 | The ID of the comment to be retrieved.
	bodyFormat := openapiclient.PrimaryBodyRepresentationSingle("storage") // PrimaryBodyRepresentationSingle | The content format type to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	version := int32(56) // int32 | Allows you to retrieve a previously published version. Specify the previous version's number to retrieve its details. (optional)
	includeProperties := true // bool | Includes content properties associated with this inline comment in the response. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeOperations := true // bool | Includes operations associated with this inline comment in the response, as defined in the `Operation` object. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeLikes := true // bool | Includes likes associated with this inline comment in the response. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeVersions := true // bool | Includes versions associated with this inline comment in the response. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeVersion := true // bool | Includes the current version associated with this inline comment in the response. By default this is included and can be omitted by setting the value to `false`. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentAPI.GetInlineCommentById(context.Background(), commentId).BodyFormat(bodyFormat).Version(version).IncludeProperties(includeProperties).IncludeOperations(includeOperations).IncludeLikes(includeLikes).IncludeVersions(includeVersions).IncludeVersion(includeVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAPI.GetInlineCommentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInlineCommentById`: CreateInlineComment201Response
	fmt.Fprintf(os.Stdout, "Response from `CommentAPI.GetInlineCommentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int64** | The ID of the comment to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInlineCommentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**PrimaryBodyRepresentationSingle**](PrimaryBodyRepresentationSingle.md) | The content format type to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **version** | **int32** | Allows you to retrieve a previously published version. Specify the previous version&#39;s number to retrieve its details. | 
 **includeProperties** | **bool** | Includes content properties associated with this inline comment in the response. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeOperations** | **bool** | Includes operations associated with this inline comment in the response, as defined in the &#x60;Operation&#x60; object. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeLikes** | **bool** | Includes likes associated with this inline comment in the response. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeVersions** | **bool** | Includes versions associated with this inline comment in the response. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeVersion** | **bool** | Includes the current version associated with this inline comment in the response. By default this is included and can be omitted by setting the value to &#x60;false&#x60;. | [default to true]

### Return type

[**CreateInlineComment201Response**](CreateInlineComment201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInlineCommentChildren

> MultiEntityResultInlineCommentChildrenModel GetInlineCommentChildren(ctx, id).BodyFormat(bodyFormat).Sort(sort).Cursor(cursor).Limit(limit).Execute()

Get children inline comments



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
	id := int64(789) // int64 | The ID of the parent comment for which inline comment children should be returned.
	bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format type to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	sort := openapiclient.CommentSortOrder("created-date") // CommentSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of footer comments per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentAPI.GetInlineCommentChildren(context.Background(), id).BodyFormat(bodyFormat).Sort(sort).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAPI.GetInlineCommentChildren``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInlineCommentChildren`: MultiEntityResultInlineCommentChildrenModel
	fmt.Fprintf(os.Stdout, "Response from `CommentAPI.GetInlineCommentChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the parent comment for which inline comment children should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInlineCommentChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format type to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **sort** | [**CommentSortOrder**](CommentSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of footer comments per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultInlineCommentChildrenModel**](MultiEntityResultInlineCommentChildrenModel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInlineComments

> MultiEntityResultInlineCommentModel GetInlineComments(ctx).BodyFormat(bodyFormat).Sort(sort).Cursor(cursor).Limit(limit).Execute()

Get inline comments



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
	bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format type to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	sort := openapiclient.CommentSortOrder("created-date") // CommentSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of footer comments per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentAPI.GetInlineComments(context.Background()).BodyFormat(bodyFormat).Sort(sort).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAPI.GetInlineComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInlineComments`: MultiEntityResultInlineCommentModel
	fmt.Fprintf(os.Stdout, "Response from `CommentAPI.GetInlineComments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInlineCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format type to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **sort** | [**CommentSortOrder**](CommentSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of footer comments per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultInlineCommentModel**](MultiEntityResultInlineCommentModel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageFooterComments

> MultiEntityResultPageCommentModel GetPageFooterComments(ctx, id).BodyFormat(bodyFormat).Status(status).Sort(sort).Cursor(cursor).Limit(limit).Execute()

Get footer comments for page



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
	id := int64(789) // int64 | The ID of the page for which footer comments should be returned.
	bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format type to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	status := []string{"Status_example"} // []string | Filter the footer comment being retrieved by its status. (optional)
	sort := openapiclient.CommentSortOrder("created-date") // CommentSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of footer comments per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentAPI.GetPageFooterComments(context.Background(), id).BodyFormat(bodyFormat).Status(status).Sort(sort).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAPI.GetPageFooterComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPageFooterComments`: MultiEntityResultPageCommentModel
	fmt.Fprintf(os.Stdout, "Response from `CommentAPI.GetPageFooterComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the page for which footer comments should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageFooterCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format type to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **status** | **[]string** | Filter the footer comment being retrieved by its status. | 
 **sort** | [**CommentSortOrder**](CommentSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of footer comments per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultPageCommentModel**](MultiEntityResultPageCommentModel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageInlineComments

> MultiEntityResultPageInlineCommentModel GetPageInlineComments(ctx, id).BodyFormat(bodyFormat).Status(status).ResolutionStatus(resolutionStatus).Sort(sort).Cursor(cursor).Limit(limit).Execute()

Get inline comments for page



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
	id := int64(789) // int64 | The ID of the page for which inline comments should be returned.
	bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format type to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	status := []string{"Status_example"} // []string | Filter the inline comment being retrieved by its status. (optional)
	resolutionStatus := []string{"ResolutionStatus_example"} // []string | Filter the inline comment being retrieved by its resolution status. (optional)
	sort := openapiclient.CommentSortOrder("created-date") // CommentSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of inline comments per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentAPI.GetPageInlineComments(context.Background(), id).BodyFormat(bodyFormat).Status(status).ResolutionStatus(resolutionStatus).Sort(sort).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAPI.GetPageInlineComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPageInlineComments`: MultiEntityResultPageInlineCommentModel
	fmt.Fprintf(os.Stdout, "Response from `CommentAPI.GetPageInlineComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the page for which inline comments should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageInlineCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format type to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **status** | **[]string** | Filter the inline comment being retrieved by its status. | 
 **resolutionStatus** | **[]string** | Filter the inline comment being retrieved by its resolution status. | 
 **sort** | [**CommentSortOrder**](CommentSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of inline comments per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultPageInlineCommentModel**](MultiEntityResultPageInlineCommentModel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFooterComment

> FooterCommentModel UpdateFooterComment(ctx, commentId).UpdateFooterCommentRequest(updateFooterCommentRequest).Execute()

Update footer comment



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
	commentId := int64(789) // int64 | The ID of the comment to be retrieved.
	updateFooterCommentRequest := *openapiclient.NewUpdateFooterCommentRequest() // UpdateFooterCommentRequest | The footer comment to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentAPI.UpdateFooterComment(context.Background(), commentId).UpdateFooterCommentRequest(updateFooterCommentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAPI.UpdateFooterComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFooterComment`: FooterCommentModel
	fmt.Fprintf(os.Stdout, "Response from `CommentAPI.UpdateFooterComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int64** | The ID of the comment to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFooterCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFooterCommentRequest** | [**UpdateFooterCommentRequest**](UpdateFooterCommentRequest.md) | The footer comment to be created | 

### Return type

[**FooterCommentModel**](FooterCommentModel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInlineComment

> CreateInlineComment201Response UpdateInlineComment(ctx, commentId).UpdateInlineCommentModel(updateInlineCommentModel).Execute()

Update inline comment



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
	commentId := int64(789) // int64 | The ID of the comment to be retrieved.
	updateInlineCommentModel := *openapiclient.NewUpdateInlineCommentModel() // UpdateInlineCommentModel | The inline comment to be updated

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentAPI.UpdateInlineComment(context.Background(), commentId).UpdateInlineCommentModel(updateInlineCommentModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAPI.UpdateInlineComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInlineComment`: CreateInlineComment201Response
	fmt.Fprintf(os.Stdout, "Response from `CommentAPI.UpdateInlineComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int64** | The ID of the comment to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInlineCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInlineCommentModel** | [**UpdateInlineCommentModel**](UpdateInlineCommentModel.md) | The inline comment to be updated | 

### Return type

[**CreateInlineComment201Response**](CreateInlineComment201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


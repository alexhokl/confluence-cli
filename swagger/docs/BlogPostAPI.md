# \BlogPostAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBlogPost**](BlogPostAPI.md#CreateBlogPost) | **Post** /blogposts | Create blog post
[**DeleteBlogPost**](BlogPostAPI.md#DeleteBlogPost) | **Delete** /blogposts/{id} | Delete blog post
[**GetBlogPostById**](BlogPostAPI.md#GetBlogPostById) | **Get** /blogposts/{id} | Get blog post by id
[**GetBlogPosts**](BlogPostAPI.md#GetBlogPosts) | **Get** /blogposts | Get blog posts
[**GetBlogPostsInSpace**](BlogPostAPI.md#GetBlogPostsInSpace) | **Get** /spaces/{id}/blogposts | Get blog posts in space
[**GetLabelBlogPosts**](BlogPostAPI.md#GetLabelBlogPosts) | **Get** /labels/{id}/blogposts | Get blog posts for label
[**UpdateBlogPost**](BlogPostAPI.md#UpdateBlogPost) | **Put** /blogposts/{id} | Update blog post



## CreateBlogPost

> CreateBlogPost200Response CreateBlogPost(ctx).CreateBlogPostRequest(createBlogPostRequest).Private(private).Execute()

Create blog post



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
	createBlogPostRequest := *openapiclient.NewCreateBlogPostRequest("SpaceId_example") // CreateBlogPostRequest | 
	private := true // bool | The blog post will be private. Only the user who creates this blog post will have permission to view and edit one. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostAPI.CreateBlogPost(context.Background()).CreateBlogPostRequest(createBlogPostRequest).Private(private).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostAPI.CreateBlogPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBlogPost`: CreateBlogPost200Response
	fmt.Fprintf(os.Stdout, "Response from `BlogPostAPI.CreateBlogPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlogPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBlogPostRequest** | [**CreateBlogPostRequest**](CreateBlogPostRequest.md) |  | 
 **private** | **bool** | The blog post will be private. Only the user who creates this blog post will have permission to view and edit one. | [default to false]

### Return type

[**CreateBlogPost200Response**](CreateBlogPost200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBlogPost

> DeleteBlogPost(ctx, id).Purge(purge).Draft(draft).Execute()

Delete blog post



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
	id := int64(789) // int64 | The ID of the blog post to be deleted.
	purge := true // bool | If attempting to purge the blog post. (optional) (default to false)
	draft := true // bool | If attempting to delete a blog post that is a draft. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogPostAPI.DeleteBlogPost(context.Background(), id).Purge(purge).Draft(draft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostAPI.DeleteBlogPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the blog post to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlogPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purge** | **bool** | If attempting to purge the blog post. | [default to false]
 **draft** | **bool** | If attempting to delete a blog post that is a draft. | [default to false]

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


## GetBlogPostById

> CreateBlogPost200Response GetBlogPostById(ctx, id).BodyFormat(bodyFormat).GetDraft(getDraft).Status(status).Version(version).IncludeLabels(includeLabels).IncludeProperties(includeProperties).IncludeOperations(includeOperations).IncludeLikes(includeLikes).IncludeVersions(includeVersions).IncludeVersion(includeVersion).IncludeFavoritedByCurrentUserStatus(includeFavoritedByCurrentUserStatus).IncludeWebresources(includeWebresources).IncludeCollaborators(includeCollaborators).Execute()

Get blog post by id



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
	id := int64(789) // int64 | The ID of the blog post to be returned. If you don't know the blog post ID, use Get blog posts and filter the results.
	bodyFormat := openapiclient.PrimaryBodyRepresentationSingle("storage") // PrimaryBodyRepresentationSingle | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	getDraft := true // bool | Retrieve the draft version of this blog post. (optional) (default to false)
	status := []string{"Status_example"} // []string | Filter the blog post being retrieved by its status. (optional)
	version := int32(56) // int32 | Allows you to retrieve a previously published version. Specify the previous version's number to retrieve its details. (optional)
	includeLabels := true // bool | Includes labels associated with this blog post in the response. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeProperties := true // bool | Includes content properties associated with this blog post in the response. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeOperations := true // bool | Includes operations associated with this blog post in the response, as defined in the `Operation` object. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeLikes := true // bool | Includes likes associated with this blog post in the response. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeVersions := true // bool | Includes versions associated with this blog post in the response. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeVersion := true // bool | Includes the current version associated with this blog post in the response. By default this is included and can be omitted by setting the value to `false`. (optional) (default to true)
	includeFavoritedByCurrentUserStatus := true // bool | Includes whether this blog post has been favorited by the current user. (optional) (default to false)
	includeWebresources := true // bool | Includes web resources that can be used to render blog post content on a client. (optional) (default to false)
	includeCollaborators := true // bool | Includes collaborators on the blog post. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostAPI.GetBlogPostById(context.Background(), id).BodyFormat(bodyFormat).GetDraft(getDraft).Status(status).Version(version).IncludeLabels(includeLabels).IncludeProperties(includeProperties).IncludeOperations(includeOperations).IncludeLikes(includeLikes).IncludeVersions(includeVersions).IncludeVersion(includeVersion).IncludeFavoritedByCurrentUserStatus(includeFavoritedByCurrentUserStatus).IncludeWebresources(includeWebresources).IncludeCollaborators(includeCollaborators).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostAPI.GetBlogPostById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlogPostById`: CreateBlogPost200Response
	fmt.Fprintf(os.Stdout, "Response from `BlogPostAPI.GetBlogPostById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the blog post to be returned. If you don&#39;t know the blog post ID, use Get blog posts and filter the results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogPostByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**PrimaryBodyRepresentationSingle**](PrimaryBodyRepresentationSingle.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **getDraft** | **bool** | Retrieve the draft version of this blog post. | [default to false]
 **status** | **[]string** | Filter the blog post being retrieved by its status. | 
 **version** | **int32** | Allows you to retrieve a previously published version. Specify the previous version&#39;s number to retrieve its details. | 
 **includeLabels** | **bool** | Includes labels associated with this blog post in the response. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeProperties** | **bool** | Includes content properties associated with this blog post in the response. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeOperations** | **bool** | Includes operations associated with this blog post in the response, as defined in the &#x60;Operation&#x60; object. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeLikes** | **bool** | Includes likes associated with this blog post in the response. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeVersions** | **bool** | Includes versions associated with this blog post in the response. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeVersion** | **bool** | Includes the current version associated with this blog post in the response. By default this is included and can be omitted by setting the value to &#x60;false&#x60;. | [default to true]
 **includeFavoritedByCurrentUserStatus** | **bool** | Includes whether this blog post has been favorited by the current user. | [default to false]
 **includeWebresources** | **bool** | Includes web resources that can be used to render blog post content on a client. | [default to false]
 **includeCollaborators** | **bool** | Includes collaborators on the blog post. | [default to false]

### Return type

[**CreateBlogPost200Response**](CreateBlogPost200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlogPosts

> MultiEntityResultBlogPost GetBlogPosts(ctx).Id(id).SpaceId(spaceId).Sort(sort).Status(status).Title(title).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).Execute()

Get blog posts



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
	id := []int64{int64(123)} // []int64 | Filter the results based on blog post ids. Multiple blog post ids can be specified as a comma-separated list. (optional)
	spaceId := []int64{int64(123)} // []int64 | Filter the results based on space ids. Multiple space ids can be specified as a comma-separated list. (optional)
	sort := openapiclient.BlogPostSortOrder("id") // BlogPostSortOrder | Used to sort the result by a particular field. (optional)
	status := []string{"Status_example"} // []string | Filter the results to blog posts based on their status. By default, `current` is used. (optional)
	title := "title_example" // string | Filter the results to blog posts based on their title. (optional)
	bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of blog posts per result to return. If more results exist, use the `Link` response header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostAPI.GetBlogPosts(context.Background()).Id(id).SpaceId(spaceId).Sort(sort).Status(status).Title(title).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostAPI.GetBlogPosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlogPosts`: MultiEntityResultBlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostAPI.GetBlogPosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogPostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]int64** | Filter the results based on blog post ids. Multiple blog post ids can be specified as a comma-separated list. | 
 **spaceId** | **[]int64** | Filter the results based on space ids. Multiple space ids can be specified as a comma-separated list. | 
 **sort** | [**BlogPostSortOrder**](BlogPostSortOrder.md) | Used to sort the result by a particular field. | 
 **status** | **[]string** | Filter the results to blog posts based on their status. By default, &#x60;current&#x60; is used. | 
 **title** | **string** | Filter the results to blog posts based on their title. | 
 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of blog posts per result to return. If more results exist, use the &#x60;Link&#x60; response header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultBlogPost**](MultiEntityResultBlogPost.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlogPostsInSpace

> MultiEntityResultBlogPost GetBlogPostsInSpace(ctx, id).Sort(sort).Status(status).Title(title).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).Execute()

Get blog posts in space



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
	id := int64(789) // int64 | The ID of the space for which blog posts should be returned.
	sort := openapiclient.BlogPostSortOrder("id") // BlogPostSortOrder | Used to sort the result by a particular field. (optional)
	status := []string{"Status_example"} // []string | Filter the results to blog posts based on their status. By default, `current` is used. (optional)
	title := "title_example" // string | Filter the results to blog posts based on their title. (optional)
	bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of blog posts per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostAPI.GetBlogPostsInSpace(context.Background(), id).Sort(sort).Status(status).Title(title).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostAPI.GetBlogPostsInSpace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlogPostsInSpace`: MultiEntityResultBlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostAPI.GetBlogPostsInSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the space for which blog posts should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogPostsInSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | [**BlogPostSortOrder**](BlogPostSortOrder.md) | Used to sort the result by a particular field. | 
 **status** | **[]string** | Filter the results to blog posts based on their status. By default, &#x60;current&#x60; is used. | 
 **title** | **string** | Filter the results to blog posts based on their title. | 
 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of blog posts per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultBlogPost**](MultiEntityResultBlogPost.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLabelBlogPosts

> MultiEntityResultBlogPost GetLabelBlogPosts(ctx, id).SpaceId(spaceId).BodyFormat(bodyFormat).Sort(sort).Cursor(cursor).Limit(limit).Execute()

Get blog posts for label



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
	id := int64(789) // int64 | The ID of the label for which blog posts should be returned.
	spaceId := []int64{int64(123)} // []int64 | Filter the results based on space ids. Multiple space ids can be specified as a comma-separated list. (optional)
	bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	sort := openapiclient.BlogPostSortOrder("id") // BlogPostSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of blog posts per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostAPI.GetLabelBlogPosts(context.Background(), id).SpaceId(spaceId).BodyFormat(bodyFormat).Sort(sort).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostAPI.GetLabelBlogPosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLabelBlogPosts`: MultiEntityResultBlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostAPI.GetLabelBlogPosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the label for which blog posts should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLabelBlogPostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spaceId** | **[]int64** | Filter the results based on space ids. Multiple space ids can be specified as a comma-separated list. | 
 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **sort** | [**BlogPostSortOrder**](BlogPostSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of blog posts per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultBlogPost**](MultiEntityResultBlogPost.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBlogPost

> CreateBlogPost200Response UpdateBlogPost(ctx, id).UpdateBlogPostRequest(updateBlogPostRequest).Execute()

Update blog post



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
	id := int64(789) // int64 | The ID of the blog post to be updated. If you don't know the blog post ID, use Get Blog Posts and filter the results.
	updateBlogPostRequest := *openapiclient.NewUpdateBlogPostRequest("Id_example", "Status_example", "Title_example", openapiclient.createBlogPost_request_body{BlogPostBodyWrite: openapiclient.NewBlogPostBodyWrite()}, *openapiclient.NewUpdateBlogPostRequestVersion()) // UpdateBlogPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostAPI.UpdateBlogPost(context.Background(), id).UpdateBlogPostRequest(updateBlogPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostAPI.UpdateBlogPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBlogPost`: CreateBlogPost200Response
	fmt.Fprintf(os.Stdout, "Response from `BlogPostAPI.UpdateBlogPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the blog post to be updated. If you don&#39;t know the blog post ID, use Get Blog Posts and filter the results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBlogPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBlogPostRequest** | [**UpdateBlogPostRequest**](UpdateBlogPostRequest.md) |  | 

### Return type

[**CreateBlogPost200Response**](CreateBlogPost200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


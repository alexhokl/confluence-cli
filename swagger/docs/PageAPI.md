# \PageAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePage**](PageAPI.md#CreatePage) | **Post** /pages | Create page
[**DeletePage**](PageAPI.md#DeletePage) | **Delete** /pages/{id} | Delete page
[**GetLabelPages**](PageAPI.md#GetLabelPages) | **Get** /labels/{id}/pages | Get pages for label
[**GetPageById**](PageAPI.md#GetPageById) | **Get** /pages/{id} | Get page by id
[**GetPages**](PageAPI.md#GetPages) | **Get** /pages | Get pages
[**GetPagesInSpace**](PageAPI.md#GetPagesInSpace) | **Get** /spaces/{id}/pages | Get pages in space
[**UpdatePage**](PageAPI.md#UpdatePage) | **Put** /pages/{id} | Update page
[**UpdatePageTitle**](PageAPI.md#UpdatePageTitle) | **Put** /pages/{id}/title | Update page title



## CreatePage

> CreatePage200Response CreatePage(ctx).CreatePageRequest(createPageRequest).Embedded(embedded).Private(private).RootLevel(rootLevel).Execute()

Create page



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
	createPageRequest := *openapiclient.NewCreatePageRequest("SpaceId_example") // CreatePageRequest | 
	embedded := true // bool | Tag the content as embedded and content will be created in NCS. (optional) (default to false)
	private := true // bool | The page will be private. Only the user who creates this page will have permission to view and edit one. (optional) (default to false)
	rootLevel := true // bool | The page will be created at the root level of the space (outside the space homepage tree). If true, then a  value may not be supplied for the `parentId` body parameter. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PageAPI.CreatePage(context.Background()).CreatePageRequest(createPageRequest).Embedded(embedded).Private(private).RootLevel(rootLevel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageAPI.CreatePage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePage`: CreatePage200Response
	fmt.Fprintf(os.Stdout, "Response from `PageAPI.CreatePage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPageRequest** | [**CreatePageRequest**](CreatePageRequest.md) |  | 
 **embedded** | **bool** | Tag the content as embedded and content will be created in NCS. | [default to false]
 **private** | **bool** | The page will be private. Only the user who creates this page will have permission to view and edit one. | [default to false]
 **rootLevel** | **bool** | The page will be created at the root level of the space (outside the space homepage tree). If true, then a  value may not be supplied for the &#x60;parentId&#x60; body parameter. | [default to false]

### Return type

[**CreatePage200Response**](CreatePage200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePage

> DeletePage(ctx, id).Purge(purge).Draft(draft).Execute()

Delete page



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
	id := int64(789) // int64 | The ID of the page to be deleted.
	purge := true // bool | If attempting to purge the page. (optional) (default to false)
	draft := true // bool | If attempting to delete a page that is a draft. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PageAPI.DeletePage(context.Background(), id).Purge(purge).Draft(draft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageAPI.DeletePage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the page to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purge** | **bool** | If attempting to purge the page. | [default to false]
 **draft** | **bool** | If attempting to delete a page that is a draft. | [default to false]

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


## GetLabelPages

> MultiEntityResultPage GetLabelPages(ctx, id).SpaceId(spaceId).BodyFormat(bodyFormat).Sort(sort).Cursor(cursor).Limit(limit).Execute()

Get pages for label



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
	id := int64(789) // int64 | The ID of the label for which pages should be returned.
	spaceId := []int64{int64(123)} // []int64 | Filter the results based on space ids. Multiple space ids can be specified as a comma-separated list. (optional)
	bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	sort := openapiclient.PageSortOrder("id") // PageSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of pages per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PageAPI.GetLabelPages(context.Background(), id).SpaceId(spaceId).BodyFormat(bodyFormat).Sort(sort).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageAPI.GetLabelPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLabelPages`: MultiEntityResultPage
	fmt.Fprintf(os.Stdout, "Response from `PageAPI.GetLabelPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the label for which pages should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLabelPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spaceId** | **[]int64** | Filter the results based on space ids. Multiple space ids can be specified as a comma-separated list. | 
 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **sort** | [**PageSortOrder**](PageSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of pages per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultPage**](MultiEntityResultPage.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageById

> CreatePage200Response GetPageById(ctx, id).BodyFormat(bodyFormat).GetDraft(getDraft).Status(status).Version(version).IncludeLabels(includeLabels).IncludeProperties(includeProperties).IncludeOperations(includeOperations).IncludeLikes(includeLikes).IncludeVersions(includeVersions).IncludeVersion(includeVersion).IncludeFavoritedByCurrentUserStatus(includeFavoritedByCurrentUserStatus).IncludeWebresources(includeWebresources).IncludeCollaborators(includeCollaborators).IncludeDirectChildren(includeDirectChildren).Execute()

Get page by id



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
	id := int64(789) // int64 | The ID of the page to be returned. If you don't know the page ID, use Get pages and filter the results.
	bodyFormat := openapiclient.PrimaryBodyRepresentationSingle("storage") // PrimaryBodyRepresentationSingle | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	getDraft := true // bool | Retrieve the draft version of this page. (optional) (default to false)
	status := []string{"Status_example"} // []string | Filter the page being retrieved by its status. (optional)
	version := int32(56) // int32 | Allows you to retrieve a previously published version. Specify the previous version's number to retrieve its details. (optional)
	includeLabels := true // bool | Includes labels associated with this page in the response. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeProperties := true // bool | Includes content properties associated with this page in the response. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeOperations := true // bool | Includes operations associated with this page in the response, as defined in the `Operation` object. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeLikes := true // bool | Includes likes associated with this page in the response. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeVersions := true // bool | Includes versions associated with this page in the response. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeVersion := true // bool | Includes the current version associated with this page in the response. By default this is included and can be omitted by setting the value to `false`. (optional) (default to true)
	includeFavoritedByCurrentUserStatus := true // bool | Includes whether this page has been favorited by the current user. (optional) (default to false)
	includeWebresources := true // bool | Includes web resources that can be used to render page content on a client. (optional) (default to false)
	includeCollaborators := true // bool | Includes collaborators on the page. (optional) (default to false)
	includeDirectChildren := true // bool | Includes direct children of the page, as defined in the `ChildrenResponse` object. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PageAPI.GetPageById(context.Background(), id).BodyFormat(bodyFormat).GetDraft(getDraft).Status(status).Version(version).IncludeLabels(includeLabels).IncludeProperties(includeProperties).IncludeOperations(includeOperations).IncludeLikes(includeLikes).IncludeVersions(includeVersions).IncludeVersion(includeVersion).IncludeFavoritedByCurrentUserStatus(includeFavoritedByCurrentUserStatus).IncludeWebresources(includeWebresources).IncludeCollaborators(includeCollaborators).IncludeDirectChildren(includeDirectChildren).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageAPI.GetPageById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPageById`: CreatePage200Response
	fmt.Fprintf(os.Stdout, "Response from `PageAPI.GetPageById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the page to be returned. If you don&#39;t know the page ID, use Get pages and filter the results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**PrimaryBodyRepresentationSingle**](PrimaryBodyRepresentationSingle.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **getDraft** | **bool** | Retrieve the draft version of this page. | [default to false]
 **status** | **[]string** | Filter the page being retrieved by its status. | 
 **version** | **int32** | Allows you to retrieve a previously published version. Specify the previous version&#39;s number to retrieve its details. | 
 **includeLabels** | **bool** | Includes labels associated with this page in the response. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeProperties** | **bool** | Includes content properties associated with this page in the response. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeOperations** | **bool** | Includes operations associated with this page in the response, as defined in the &#x60;Operation&#x60; object. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeLikes** | **bool** | Includes likes associated with this page in the response. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeVersions** | **bool** | Includes versions associated with this page in the response. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeVersion** | **bool** | Includes the current version associated with this page in the response. By default this is included and can be omitted by setting the value to &#x60;false&#x60;. | [default to true]
 **includeFavoritedByCurrentUserStatus** | **bool** | Includes whether this page has been favorited by the current user. | [default to false]
 **includeWebresources** | **bool** | Includes web resources that can be used to render page content on a client. | [default to false]
 **includeCollaborators** | **bool** | Includes collaborators on the page. | [default to false]
 **includeDirectChildren** | **bool** | Includes direct children of the page, as defined in the &#x60;ChildrenResponse&#x60; object. | [default to false]

### Return type

[**CreatePage200Response**](CreatePage200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPages

> MultiEntityResultPage GetPages(ctx).Id(id).SpaceId(spaceId).Sort(sort).Status(status).Title(title).BodyFormat(bodyFormat).Subtype(subtype).Cursor(cursor).Limit(limit).Execute()

Get pages



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
	id := []int64{int64(123)} // []int64 | Filter the results based on page ids. Multiple page ids can be specified as a comma-separated list. (optional)
	spaceId := []int64{int64(123)} // []int64 | Filter the results based on space ids. Multiple space ids can be specified as a comma-separated list. (optional)
	sort := openapiclient.PageSortOrder("id") // PageSortOrder | Used to sort the result by a particular field. (optional)
	status := []string{"Status_example"} // []string | Filter the results to pages based on their status. By default, `current` and `archived` are used. (optional)
	title := "title_example" // string | Filter the results to pages based on their title. (optional)
	bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	subtype := "subtype_example" // string | Filter the results to pages based on their subtype. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of pages per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PageAPI.GetPages(context.Background()).Id(id).SpaceId(spaceId).Sort(sort).Status(status).Title(title).BodyFormat(bodyFormat).Subtype(subtype).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageAPI.GetPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPages`: MultiEntityResultPage
	fmt.Fprintf(os.Stdout, "Response from `PageAPI.GetPages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]int64** | Filter the results based on page ids. Multiple page ids can be specified as a comma-separated list. | 
 **spaceId** | **[]int64** | Filter the results based on space ids. Multiple space ids can be specified as a comma-separated list. | 
 **sort** | [**PageSortOrder**](PageSortOrder.md) | Used to sort the result by a particular field. | 
 **status** | **[]string** | Filter the results to pages based on their status. By default, &#x60;current&#x60; and &#x60;archived&#x60; are used. | 
 **title** | **string** | Filter the results to pages based on their title. | 
 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **subtype** | **string** | Filter the results to pages based on their subtype. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of pages per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultPage**](MultiEntityResultPage.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesInSpace

> MultiEntityResultPage GetPagesInSpace(ctx, id).Depth(depth).Sort(sort).Status(status).Title(title).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).Execute()

Get pages in space



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
	id := int64(789) // int64 | The ID of the space for which pages should be returned.
	depth := "depth_example" // string | Filter the results to pages at the root level of the space or to all pages in the space. (optional) (default to "all")
	sort := openapiclient.PageSortOrder("id") // PageSortOrder | Used to sort the result by a particular field. (optional)
	status := []string{"Status_example"} // []string | Filter the results to pages based on their status. By default, `current` and `archived` are used. (optional)
	title := "title_example" // string | Filter the results to pages based on their title. (optional)
	bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of pages per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PageAPI.GetPagesInSpace(context.Background(), id).Depth(depth).Sort(sort).Status(status).Title(title).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageAPI.GetPagesInSpace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagesInSpace`: MultiEntityResultPage
	fmt.Fprintf(os.Stdout, "Response from `PageAPI.GetPagesInSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the space for which pages should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesInSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **depth** | **string** | Filter the results to pages at the root level of the space or to all pages in the space. | [default to &quot;all&quot;]
 **sort** | [**PageSortOrder**](PageSortOrder.md) | Used to sort the result by a particular field. | 
 **status** | **[]string** | Filter the results to pages based on their status. By default, &#x60;current&#x60; and &#x60;archived&#x60; are used. | 
 **title** | **string** | Filter the results to pages based on their title. | 
 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of pages per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultPage**](MultiEntityResultPage.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePage

> CreatePage200Response UpdatePage(ctx, id).UpdatePageRequest(updatePageRequest).Execute()

Update page



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
	id := int64(789) // int64 | The ID of the page to be updated. If you don't know the page ID, use Get Pages and filter the results.
	updatePageRequest := *openapiclient.NewUpdatePageRequest("Id_example", "Status_example", "Title_example", openapiclient.createPage_request_body{PageBodyWrite: openapiclient.NewPageBodyWrite()}, *openapiclient.NewUpdatePageRequestVersion()) // UpdatePageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PageAPI.UpdatePage(context.Background(), id).UpdatePageRequest(updatePageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageAPI.UpdatePage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePage`: CreatePage200Response
	fmt.Fprintf(os.Stdout, "Response from `PageAPI.UpdatePage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the page to be updated. If you don&#39;t know the page ID, use Get Pages and filter the results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePageRequest** | [**UpdatePageRequest**](UpdatePageRequest.md) |  | 

### Return type

[**CreatePage200Response**](CreatePage200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePageTitle

> CreatePage200Response UpdatePageTitle(ctx, id).UpdatePageTitleRequest(updatePageTitleRequest).Execute()

Update page title



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
	id := int64(789) // int64 | The ID of the page to be updated. If you don't know the page ID, use Get Pages and filter the results
	updatePageTitleRequest := *openapiclient.NewUpdatePageTitleRequest("Status_example", "Title_example") // UpdatePageTitleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PageAPI.UpdatePageTitle(context.Background(), id).UpdatePageTitleRequest(updatePageTitleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageAPI.UpdatePageTitle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePageTitle`: CreatePage200Response
	fmt.Fprintf(os.Stdout, "Response from `PageAPI.UpdatePageTitle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the page to be updated. If you don&#39;t know the page ID, use Get Pages and filter the results | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePageTitleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePageTitleRequest** | [**UpdatePageTitleRequest**](UpdatePageTitleRequest.md) |  | 

### Return type

[**CreatePage200Response**](CreatePage200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


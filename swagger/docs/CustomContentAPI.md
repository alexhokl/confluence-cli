# \CustomContentAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomContent**](CustomContentAPI.md#CreateCustomContent) | **Post** /custom-content | Create custom content
[**DeleteCustomContent**](CustomContentAPI.md#DeleteCustomContent) | **Delete** /custom-content/{id} | Delete custom content
[**GetCustomContentById**](CustomContentAPI.md#GetCustomContentById) | **Get** /custom-content/{id} | Get custom content by id
[**GetCustomContentByType**](CustomContentAPI.md#GetCustomContentByType) | **Get** /custom-content | Get custom content by type
[**GetCustomContentByTypeInBlogPost**](CustomContentAPI.md#GetCustomContentByTypeInBlogPost) | **Get** /blogposts/{id}/custom-content | Get custom content by type in blog post
[**GetCustomContentByTypeInPage**](CustomContentAPI.md#GetCustomContentByTypeInPage) | **Get** /pages/{id}/custom-content | Get custom content by type in page
[**GetCustomContentByTypeInSpace**](CustomContentAPI.md#GetCustomContentByTypeInSpace) | **Get** /spaces/{id}/custom-content | Get custom content by type in space
[**UpdateCustomContent**](CustomContentAPI.md#UpdateCustomContent) | **Put** /custom-content/{id} | Update custom content



## CreateCustomContent

> CreateCustomContent201Response CreateCustomContent(ctx).CreateCustomContentRequest(createCustomContentRequest).Execute()

Create custom content



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
	createCustomContentRequest := *openapiclient.NewCreateCustomContentRequest("Type_example", "Title_example", openapiclient.createCustomContent_request_body{CustomContentBodyWrite: openapiclient.NewCustomContentBodyWrite()}) // CreateCustomContentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomContentAPI.CreateCustomContent(context.Background()).CreateCustomContentRequest(createCustomContentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomContentAPI.CreateCustomContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomContent`: CreateCustomContent201Response
	fmt.Fprintf(os.Stdout, "Response from `CustomContentAPI.CreateCustomContent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCustomContentRequest** | [**CreateCustomContentRequest**](CreateCustomContentRequest.md) |  | 

### Return type

[**CreateCustomContent201Response**](CreateCustomContent201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomContent

> DeleteCustomContent(ctx, id).Purge(purge).Execute()

Delete custom content



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
	id := int64(789) // int64 | The ID of the custom content to be deleted.
	purge := true // bool | If attempting to purge the custom content. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomContentAPI.DeleteCustomContent(context.Background(), id).Purge(purge).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomContentAPI.DeleteCustomContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the custom content to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purge** | **bool** | If attempting to purge the custom content. | [default to false]

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


## GetCustomContentById

> CreateCustomContent201Response GetCustomContentById(ctx, id).BodyFormat(bodyFormat).Version(version).IncludeLabels(includeLabels).IncludeProperties(includeProperties).IncludeOperations(includeOperations).IncludeVersions(includeVersions).IncludeVersion(includeVersion).IncludeCollaborators(includeCollaborators).Execute()

Get custom content by id



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
	id := int64(789) // int64 | The ID of the custom content to be returned. If you don't know the custom content ID, use Get Custom Content by Type and filter the results.
	bodyFormat := openapiclient.CustomContentBodyRepresentationSingle("raw") // CustomContentBodyRepresentationSingle | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field.  Note: If the custom content body type is `storage`, the `storage` and `atlas_doc_format` body formats are able to be returned. If the custom content body type is `raw`, only the `raw` body format is able to be returned. (optional)
	version := int32(56) // int32 | Allows you to retrieve a previously published version. Specify the previous version's number to retrieve its details. (optional)
	includeLabels := true // bool | Includes labels associated with this custom content in the response. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeProperties := true // bool | Includes content properties associated with this custom content in the response. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeOperations := true // bool | Includes operations associated with this custom content in the response, as defined in the `Operation` object. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeVersions := true // bool | Includes versions associated with this custom content in the response. The number of results will be limited to 50 and sorted in the default sort order.  A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeVersion := true // bool | Includes the current version associated with this custom content in the response. By default this is included and can be omitted by setting the value to `false`. (optional) (default to true)
	includeCollaborators := true // bool | Includes collaborators on the custom content. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomContentAPI.GetCustomContentById(context.Background(), id).BodyFormat(bodyFormat).Version(version).IncludeLabels(includeLabels).IncludeProperties(includeProperties).IncludeOperations(includeOperations).IncludeVersions(includeVersions).IncludeVersion(includeVersion).IncludeCollaborators(includeCollaborators).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomContentAPI.GetCustomContentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomContentById`: CreateCustomContent201Response
	fmt.Fprintf(os.Stdout, "Response from `CustomContentAPI.GetCustomContentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the custom content to be returned. If you don&#39;t know the custom content ID, use Get Custom Content by Type and filter the results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomContentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**CustomContentBodyRepresentationSingle**](CustomContentBodyRepresentationSingle.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field.  Note: If the custom content body type is &#x60;storage&#x60;, the &#x60;storage&#x60; and &#x60;atlas_doc_format&#x60; body formats are able to be returned. If the custom content body type is &#x60;raw&#x60;, only the &#x60;raw&#x60; body format is able to be returned. | 
 **version** | **int32** | Allows you to retrieve a previously published version. Specify the previous version&#39;s number to retrieve its details. | 
 **includeLabels** | **bool** | Includes labels associated with this custom content in the response. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeProperties** | **bool** | Includes content properties associated with this custom content in the response. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeOperations** | **bool** | Includes operations associated with this custom content in the response, as defined in the &#x60;Operation&#x60; object. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeVersions** | **bool** | Includes versions associated with this custom content in the response. The number of results will be limited to 50 and sorted in the default sort order.  A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeVersion** | **bool** | Includes the current version associated with this custom content in the response. By default this is included and can be omitted by setting the value to &#x60;false&#x60;. | [default to true]
 **includeCollaborators** | **bool** | Includes collaborators on the custom content. | [default to false]

### Return type

[**CreateCustomContent201Response**](CreateCustomContent201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomContentByType

> MultiEntityResultCustomContent GetCustomContentByType(ctx).Type_(type_).Id(id).SpaceId(spaceId).Sort(sort).Cursor(cursor).Limit(limit).BodyFormat(bodyFormat).Execute()

Get custom content by type



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
	type_ := "type__example" // string | The type of custom content being requested. See: https://developer.atlassian.com/cloud/confluence/custom-content/ for additional details on custom content.
	id := []int64{int64(123)} // []int64 | Filter the results based on custom content ids. Multiple custom content ids can be specified as a comma-separated list. (optional)
	spaceId := []int64{int64(123)} // []int64 | Filter the results based on space ids. Multiple space ids can be specified as a comma-separated list. (optional)
	sort := openapiclient.CustomContentSortOrder("id") // CustomContentSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of pages per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
	bodyFormat := openapiclient.CustomContentBodyRepresentation("raw") // CustomContentBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field.  Note: If the custom content body type is `storage`, the `storage` and `atlas_doc_format` body formats are able to be returned. If the custom content body type is `raw`, only the `raw` body format is able to be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomContentAPI.GetCustomContentByType(context.Background()).Type_(type_).Id(id).SpaceId(spaceId).Sort(sort).Cursor(cursor).Limit(limit).BodyFormat(bodyFormat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomContentAPI.GetCustomContentByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomContentByType`: MultiEntityResultCustomContent
	fmt.Fprintf(os.Stdout, "Response from `CustomContentAPI.GetCustomContentByType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomContentByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | The type of custom content being requested. See: https://developer.atlassian.com/cloud/confluence/custom-content/ for additional details on custom content. | 
 **id** | **[]int64** | Filter the results based on custom content ids. Multiple custom content ids can be specified as a comma-separated list. | 
 **spaceId** | **[]int64** | Filter the results based on space ids. Multiple space ids can be specified as a comma-separated list. | 
 **sort** | [**CustomContentSortOrder**](CustomContentSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of pages per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **bodyFormat** | [**CustomContentBodyRepresentation**](CustomContentBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field.  Note: If the custom content body type is &#x60;storage&#x60;, the &#x60;storage&#x60; and &#x60;atlas_doc_format&#x60; body formats are able to be returned. If the custom content body type is &#x60;raw&#x60;, only the &#x60;raw&#x60; body format is able to be returned. | 

### Return type

[**MultiEntityResultCustomContent**](MultiEntityResultCustomContent.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomContentByTypeInBlogPost

> MultiEntityResultCustomContent GetCustomContentByTypeInBlogPost(ctx, id).Type_(type_).Sort(sort).Cursor(cursor).Limit(limit).BodyFormat(bodyFormat).Execute()

Get custom content by type in blog post



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
	id := int64(789) // int64 | The ID of the blog post for which custom content should be returned.
	type_ := "type__example" // string | The type of custom content being requested. See: https://developer.atlassian.com/cloud/confluence/custom-content/ for additional details on custom content.
	sort := openapiclient.CustomContentSortOrder("id") // CustomContentSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of pages per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
	bodyFormat := openapiclient.CustomContentBodyRepresentation("raw") // CustomContentBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field.  Note: If the custom content body type is `storage`, the `storage` and `atlas_doc_format` body formats are able to be returned. If the custom content body type is `raw`, only the `raw` body format is able to be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomContentAPI.GetCustomContentByTypeInBlogPost(context.Background(), id).Type_(type_).Sort(sort).Cursor(cursor).Limit(limit).BodyFormat(bodyFormat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomContentAPI.GetCustomContentByTypeInBlogPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomContentByTypeInBlogPost`: MultiEntityResultCustomContent
	fmt.Fprintf(os.Stdout, "Response from `CustomContentAPI.GetCustomContentByTypeInBlogPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the blog post for which custom content should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomContentByTypeInBlogPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | The type of custom content being requested. See: https://developer.atlassian.com/cloud/confluence/custom-content/ for additional details on custom content. | 
 **sort** | [**CustomContentSortOrder**](CustomContentSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of pages per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **bodyFormat** | [**CustomContentBodyRepresentation**](CustomContentBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field.  Note: If the custom content body type is &#x60;storage&#x60;, the &#x60;storage&#x60; and &#x60;atlas_doc_format&#x60; body formats are able to be returned. If the custom content body type is &#x60;raw&#x60;, only the &#x60;raw&#x60; body format is able to be returned. | 

### Return type

[**MultiEntityResultCustomContent**](MultiEntityResultCustomContent.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomContentByTypeInPage

> MultiEntityResultCustomContent GetCustomContentByTypeInPage(ctx, id).Type_(type_).Sort(sort).Cursor(cursor).Limit(limit).BodyFormat(bodyFormat).Execute()

Get custom content by type in page



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
	id := int64(789) // int64 | The ID of the page for which custom content should be returned.
	type_ := "type__example" // string | The type of custom content being requested. See: https://developer.atlassian.com/cloud/confluence/custom-content/ for additional details on custom content.
	sort := openapiclient.CustomContentSortOrder("id") // CustomContentSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of pages per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
	bodyFormat := openapiclient.CustomContentBodyRepresentation("raw") // CustomContentBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field.  Note: If the custom content body type is `storage`, the `storage` and `atlas_doc_format` body formats are able to be returned. If the custom content body type is `raw`, only the `raw` body format is able to be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomContentAPI.GetCustomContentByTypeInPage(context.Background(), id).Type_(type_).Sort(sort).Cursor(cursor).Limit(limit).BodyFormat(bodyFormat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomContentAPI.GetCustomContentByTypeInPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomContentByTypeInPage`: MultiEntityResultCustomContent
	fmt.Fprintf(os.Stdout, "Response from `CustomContentAPI.GetCustomContentByTypeInPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the page for which custom content should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomContentByTypeInPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | The type of custom content being requested. See: https://developer.atlassian.com/cloud/confluence/custom-content/ for additional details on custom content. | 
 **sort** | [**CustomContentSortOrder**](CustomContentSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of pages per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **bodyFormat** | [**CustomContentBodyRepresentation**](CustomContentBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field.  Note: If the custom content body type is &#x60;storage&#x60;, the &#x60;storage&#x60; and &#x60;atlas_doc_format&#x60; body formats are able to be returned. If the custom content body type is &#x60;raw&#x60;, only the &#x60;raw&#x60; body format is able to be returned. | 

### Return type

[**MultiEntityResultCustomContent**](MultiEntityResultCustomContent.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomContentByTypeInSpace

> MultiEntityResultCustomContent GetCustomContentByTypeInSpace(ctx, id).Type_(type_).Cursor(cursor).Limit(limit).BodyFormat(bodyFormat).Execute()

Get custom content by type in space



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
	id := int64(789) // int64 | The ID of the space for which custom content should be returned.
	type_ := "type__example" // string | The type of custom content being requested. See: https://developer.atlassian.com/cloud/confluence/custom-content/ for additional details on custom content.
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of pages per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
	bodyFormat := openapiclient.CustomContentBodyRepresentation("raw") // CustomContentBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field.  Note: If the custom content body type is `storage`, the `storage` and `atlas_doc_format` body formats are able to be returned. If the custom content body type is `raw`, only the `raw` body format is able to be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomContentAPI.GetCustomContentByTypeInSpace(context.Background(), id).Type_(type_).Cursor(cursor).Limit(limit).BodyFormat(bodyFormat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomContentAPI.GetCustomContentByTypeInSpace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomContentByTypeInSpace`: MultiEntityResultCustomContent
	fmt.Fprintf(os.Stdout, "Response from `CustomContentAPI.GetCustomContentByTypeInSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the space for which custom content should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomContentByTypeInSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | The type of custom content being requested. See: https://developer.atlassian.com/cloud/confluence/custom-content/ for additional details on custom content. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of pages per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **bodyFormat** | [**CustomContentBodyRepresentation**](CustomContentBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field.  Note: If the custom content body type is &#x60;storage&#x60;, the &#x60;storage&#x60; and &#x60;atlas_doc_format&#x60; body formats are able to be returned. If the custom content body type is &#x60;raw&#x60;, only the &#x60;raw&#x60; body format is able to be returned. | 

### Return type

[**MultiEntityResultCustomContent**](MultiEntityResultCustomContent.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomContent

> CreateCustomContent201Response UpdateCustomContent(ctx, id).UpdateCustomContentRequest(updateCustomContentRequest).Execute()

Update custom content



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
	id := int64(789) // int64 | The ID of the custom content to be updated. If you don't know the custom content ID, use Get Custom Content by Type and filter the results.
	updateCustomContentRequest := *openapiclient.NewUpdateCustomContentRequest("Id_example", "Type_example", "Status_example", "Title_example", openapiclient.createCustomContent_request_body{CustomContentBodyWrite: openapiclient.NewCustomContentBodyWrite()}, *openapiclient.NewUpdateCustomContentRequestVersion()) // UpdateCustomContentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomContentAPI.UpdateCustomContent(context.Background(), id).UpdateCustomContentRequest(updateCustomContentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomContentAPI.UpdateCustomContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomContent`: CreateCustomContent201Response
	fmt.Fprintf(os.Stdout, "Response from `CustomContentAPI.UpdateCustomContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the custom content to be updated. If you don&#39;t know the custom content ID, use Get Custom Content by Type and filter the results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCustomContentRequest** | [**UpdateCustomContentRequest**](UpdateCustomContentRequest.md) |  | 

### Return type

[**CreateCustomContent201Response**](CreateCustomContent201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


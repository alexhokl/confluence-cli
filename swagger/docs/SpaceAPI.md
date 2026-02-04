# \SpaceAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSpace**](SpaceAPI.md#CreateSpace) | **Post** /spaces | Create space
[**GetSpaceById**](SpaceAPI.md#GetSpaceById) | **Get** /spaces/{id} | Get space by id
[**GetSpaces**](SpaceAPI.md#GetSpaces) | **Get** /spaces | Get spaces



## CreateSpace

> CreateSpace201Response CreateSpace(ctx).CreateSpaceRequest(createSpaceRequest).Execute()

Create space



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
	createSpaceRequest := *openapiclient.NewCreateSpaceRequest("Name_example") // CreateSpaceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceAPI.CreateSpace(context.Background()).CreateSpaceRequest(createSpaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceAPI.CreateSpace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSpace`: CreateSpace201Response
	fmt.Fprintf(os.Stdout, "Response from `SpaceAPI.CreateSpace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSpaceRequest** | [**CreateSpaceRequest**](CreateSpaceRequest.md) |  | 

### Return type

[**CreateSpace201Response**](CreateSpace201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpaceById

> GetSpaceById200Response GetSpaceById(ctx, id).DescriptionFormat(descriptionFormat).IncludeIcon(includeIcon).IncludeOperations(includeOperations).IncludeProperties(includeProperties).IncludePermissions(includePermissions).IncludeRoleAssignments(includeRoleAssignments).IncludeLabels(includeLabels).Execute()

Get space by id



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
	id := int64(789) // int64 | The ID of the space to be returned.
	descriptionFormat := openapiclient.SpaceDescriptionBodyRepresentation("plain") // SpaceDescriptionBodyRepresentation | The content format type to be returned in the `description` field of the response. If available, the representation will be available under a response field of the same name under the `description` field. (optional)
	includeIcon := true // bool | If the icon for the space should be fetched or not. (optional) (default to false)
	includeOperations := true // bool | Includes operations associated with this space in the response, as defined in the `Operation` object. The number of results will be limited to 50 and sorted in the default sort order. A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeProperties := true // bool | Includes space properties associated with this space in the response. The number of results will be limited to 50 and sorted in the default sort order. A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includePermissions := true // bool | Includes space permissions associated with this space in the response. The number of results will be limited to 50 and sorted in the default sort order. A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeRoleAssignments := true // bool | Includes role assignments associated with this space in the response. This parameter is only accepted for EAP sites. The number of results will be limited to 50 and sorted in the default sort order. A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)
	includeLabels := true // bool | Includes labels associated with this space in the response. The number of results will be limited to 50 and sorted in the default sort order. A `meta` and `_links` property will be present to indicate if more results are available and a link to retrieve the rest of the results. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceAPI.GetSpaceById(context.Background(), id).DescriptionFormat(descriptionFormat).IncludeIcon(includeIcon).IncludeOperations(includeOperations).IncludeProperties(includeProperties).IncludePermissions(includePermissions).IncludeRoleAssignments(includeRoleAssignments).IncludeLabels(includeLabels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceAPI.GetSpaceById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpaceById`: GetSpaceById200Response
	fmt.Fprintf(os.Stdout, "Response from `SpaceAPI.GetSpaceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the space to be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpaceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **descriptionFormat** | [**SpaceDescriptionBodyRepresentation**](SpaceDescriptionBodyRepresentation.md) | The content format type to be returned in the &#x60;description&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;description&#x60; field. | 
 **includeIcon** | **bool** | If the icon for the space should be fetched or not. | [default to false]
 **includeOperations** | **bool** | Includes operations associated with this space in the response, as defined in the &#x60;Operation&#x60; object. The number of results will be limited to 50 and sorted in the default sort order. A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeProperties** | **bool** | Includes space properties associated with this space in the response. The number of results will be limited to 50 and sorted in the default sort order. A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includePermissions** | **bool** | Includes space permissions associated with this space in the response. The number of results will be limited to 50 and sorted in the default sort order. A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeRoleAssignments** | **bool** | Includes role assignments associated with this space in the response. This parameter is only accepted for EAP sites. The number of results will be limited to 50 and sorted in the default sort order. A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]
 **includeLabels** | **bool** | Includes labels associated with this space in the response. The number of results will be limited to 50 and sorted in the default sort order. A &#x60;meta&#x60; and &#x60;_links&#x60; property will be present to indicate if more results are available and a link to retrieve the rest of the results. | [default to false]

### Return type

[**GetSpaceById200Response**](GetSpaceById200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpaces

> MultiEntityResultSpace GetSpaces(ctx).Ids(ids).Keys(keys).Type_(type_).Status(status).Labels(labels).FavoritedBy(favoritedBy).NotFavoritedBy(notFavoritedBy).Sort(sort).DescriptionFormat(descriptionFormat).IncludeIcon(includeIcon).Cursor(cursor).Limit(limit).Execute()

Get spaces



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
	ids := []int64{int64(123)} // []int64 | Filter the results to spaces based on their IDs. Multiple IDs can be specified as a comma-separated list. (optional)
	keys := []string{"Inner_example"} // []string | Filter the results to spaces based on their keys. Multiple keys can be specified as a comma-separated list. (optional)
	type_ := "type__example" // string | Filter the results to spaces based on their type. (optional)
	status := "status_example" // string | Filter the results to spaces based on their status. (optional)
	labels := []string{"Inner_example"} // []string | Filter the results to spaces based on their labels. Multiple labels can be specified as a comma-separated list. (optional)
	favoritedBy := "favoritedBy_example" // string | Filter the results to spaces favorited by the user with the specified account ID. (optional)
	notFavoritedBy := "notFavoritedBy_example" // string | Filter the results to spaces NOT favorited by the user with the specified account ID. (optional)
	sort := openapiclient.SpaceSortOrder("id") // SpaceSortOrder | Used to sort the result by a particular field. (optional)
	descriptionFormat := openapiclient.SpaceDescriptionBodyRepresentation("plain") // SpaceDescriptionBodyRepresentation | The content format type to be returned in the `description` field of the response. If available, the representation will be available under a response field of the same name under the `description` field. (optional)
	includeIcon := true // bool | If the icon for the space should be fetched or not. (optional) (default to false)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of spaces per result to return. If more results exist, use the `Link` response header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceAPI.GetSpaces(context.Background()).Ids(ids).Keys(keys).Type_(type_).Status(status).Labels(labels).FavoritedBy(favoritedBy).NotFavoritedBy(notFavoritedBy).Sort(sort).DescriptionFormat(descriptionFormat).IncludeIcon(includeIcon).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceAPI.GetSpaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpaces`: MultiEntityResultSpace
	fmt.Fprintf(os.Stdout, "Response from `SpaceAPI.GetSpaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSpacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int64** | Filter the results to spaces based on their IDs. Multiple IDs can be specified as a comma-separated list. | 
 **keys** | **[]string** | Filter the results to spaces based on their keys. Multiple keys can be specified as a comma-separated list. | 
 **type_** | **string** | Filter the results to spaces based on their type. | 
 **status** | **string** | Filter the results to spaces based on their status. | 
 **labels** | **[]string** | Filter the results to spaces based on their labels. Multiple labels can be specified as a comma-separated list. | 
 **favoritedBy** | **string** | Filter the results to spaces favorited by the user with the specified account ID. | 
 **notFavoritedBy** | **string** | Filter the results to spaces NOT favorited by the user with the specified account ID. | 
 **sort** | [**SpaceSortOrder**](SpaceSortOrder.md) | Used to sort the result by a particular field. | 
 **descriptionFormat** | [**SpaceDescriptionBodyRepresentation**](SpaceDescriptionBodyRepresentation.md) | The content format type to be returned in the &#x60;description&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;description&#x60; field. | 
 **includeIcon** | **bool** | If the icon for the space should be fetched or not. | [default to false]
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of spaces per result to return. If more results exist, use the &#x60;Link&#x60; response header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultSpace**](MultiEntityResultSpace.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


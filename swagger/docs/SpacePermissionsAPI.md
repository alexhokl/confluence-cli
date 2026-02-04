# \SpacePermissionsAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAvailableSpacePermissions**](SpacePermissionsAPI.md#GetAvailableSpacePermissions) | **Get** /space-permissions | Get available space permissions
[**GetSpacePermissionsAssignments**](SpacePermissionsAPI.md#GetSpacePermissionsAssignments) | **Get** /spaces/{id}/permissions | Get space permissions assignments



## GetAvailableSpacePermissions

> MultiEntityResultSpacePermission GetAvailableSpacePermissions(ctx).Cursor(cursor).Limit(limit).Execute()

Get available space permissions



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
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of space permissions to return. If more results exist, use the `Link` response header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacePermissionsAPI.GetAvailableSpacePermissions(context.Background()).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacePermissionsAPI.GetAvailableSpacePermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailableSpacePermissions`: MultiEntityResultSpacePermission
	fmt.Fprintf(os.Stdout, "Response from `SpacePermissionsAPI.GetAvailableSpacePermissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableSpacePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of space permissions to return. If more results exist, use the &#x60;Link&#x60; response header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultSpacePermission**](MultiEntityResultSpacePermission.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpacePermissionsAssignments

> MultiEntityResultSpacePermissionAssignment GetSpacePermissionsAssignments(ctx, id).Cursor(cursor).Limit(limit).Execute()

Get space permissions assignments



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
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of assignments to return. If more results exist, use the `Link` response header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacePermissionsAPI.GetSpacePermissionsAssignments(context.Background(), id).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacePermissionsAPI.GetSpacePermissionsAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpacePermissionsAssignments`: MultiEntityResultSpacePermissionAssignment
	fmt.Fprintf(os.Stdout, "Response from `SpacePermissionsAPI.GetSpacePermissionsAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the space to be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpacePermissionsAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of assignments to return. If more results exist, use the &#x60;Link&#x60; response header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultSpacePermissionAssignment**](MultiEntityResultSpacePermissionAssignment.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


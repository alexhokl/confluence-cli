# \SpaceRolesAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSpaceRole**](SpaceRolesAPI.md#CreateSpaceRole) | **Post** /space-roles | Create a space role
[**DeleteSpaceRole**](SpaceRolesAPI.md#DeleteSpaceRole) | **Delete** /space-roles/{id} | Delete a space role
[**GetAvailableSpaceRoles**](SpaceRolesAPI.md#GetAvailableSpaceRoles) | **Get** /space-roles | Get available space roles
[**GetSpaceRoleAssignments**](SpaceRolesAPI.md#GetSpaceRoleAssignments) | **Get** /spaces/{id}/role-assignments | Get space role assignments
[**GetSpaceRoleMode**](SpaceRolesAPI.md#GetSpaceRoleMode) | **Get** /space-role-mode | Get space role mode
[**GetSpaceRolesById**](SpaceRolesAPI.md#GetSpaceRolesById) | **Get** /space-roles/{id} | Get space role by ID
[**SetSpaceRoleAssignments**](SpaceRolesAPI.md#SetSpaceRoleAssignments) | **Post** /spaces/{id}/role-assignments | Set space role assignments
[**UpdateSpaceRole**](SpaceRolesAPI.md#UpdateSpaceRole) | **Put** /space-roles/{id} | Update a space role



## CreateSpaceRole

> SpaceRole CreateSpaceRole(ctx).CreateSpaceRoleRequest(createSpaceRoleRequest).Execute()

Create a space role



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
	createSpaceRoleRequest := *openapiclient.NewCreateSpaceRoleRequest("Name_example", "Description_example", []string{"SpacePermissions_example"}) // CreateSpaceRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceRolesAPI.CreateSpaceRole(context.Background()).CreateSpaceRoleRequest(createSpaceRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceRolesAPI.CreateSpaceRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSpaceRole`: SpaceRole
	fmt.Fprintf(os.Stdout, "Response from `SpaceRolesAPI.CreateSpaceRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpaceRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSpaceRoleRequest** | [**CreateSpaceRoleRequest**](CreateSpaceRoleRequest.md) |  | 

### Return type

[**SpaceRole**](SpaceRole.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSpaceRole

> DeleteSpaceRoleResponse DeleteSpaceRole(ctx, id).Execute()

Delete a space role



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
	id := "id_example" // string | Id of the space role

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceRolesAPI.DeleteSpaceRole(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceRolesAPI.DeleteSpaceRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSpaceRole`: DeleteSpaceRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `SpaceRolesAPI.DeleteSpaceRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the space role | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpaceRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteSpaceRoleResponse**](DeleteSpaceRoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailableSpaceRoles

> MultiEntityResultSpaceRole GetAvailableSpaceRoles(ctx).SpaceId(spaceId).RoleType(roleType).PrincipalId(principalId).PrincipalType(principalType).Cursor(cursor).Limit(limit).Execute()

Get available space roles



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
	spaceId := "spaceId_example" // string | The space ID for which to filter available space roles; if empty, return all available space roles for the tenant. (optional)
	roleType := "roleType_example" // string | The space role type to filter results by. (optional)
	principalId := "principalId_example" // string | The principal ID to filter results by. If specified, a principal-type must also be specified. Paired with a `principal-type` of `ACCESS_CLASS`, valid values include [`anonymous-users`, `jsm-project-admins`, `authenticated-users`, `all-licensed-users`, `all-product-admins`] (optional)
	principalType := openapiclient.PrincipalType("USER") // PrincipalType | The principal type to filter results by. If specified, a principal-id must also be specified. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of space roles to return. If more results exist, use the `Link` response header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceRolesAPI.GetAvailableSpaceRoles(context.Background()).SpaceId(spaceId).RoleType(roleType).PrincipalId(principalId).PrincipalType(principalType).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceRolesAPI.GetAvailableSpaceRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailableSpaceRoles`: MultiEntityResultSpaceRole
	fmt.Fprintf(os.Stdout, "Response from `SpaceRolesAPI.GetAvailableSpaceRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableSpaceRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spaceId** | **string** | The space ID for which to filter available space roles; if empty, return all available space roles for the tenant. | 
 **roleType** | **string** | The space role type to filter results by. | 
 **principalId** | **string** | The principal ID to filter results by. If specified, a principal-type must also be specified. Paired with a &#x60;principal-type&#x60; of &#x60;ACCESS_CLASS&#x60;, valid values include [&#x60;anonymous-users&#x60;, &#x60;jsm-project-admins&#x60;, &#x60;authenticated-users&#x60;, &#x60;all-licensed-users&#x60;, &#x60;all-product-admins&#x60;] | 
 **principalType** | [**PrincipalType**](PrincipalType.md) | The principal type to filter results by. If specified, a principal-id must also be specified. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of space roles to return. If more results exist, use the &#x60;Link&#x60; response header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultSpaceRole**](MultiEntityResultSpaceRole.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpaceRoleAssignments

> MultiEntityResultSpaceRoleAssignment GetSpaceRoleAssignments(ctx, id).RoleId(roleId).RoleType(roleType).PrincipalId(principalId).PrincipalType(principalType).Cursor(cursor).Limit(limit).Execute()

Get space role assignments



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
	id := int32(56) // int32 | The ID of the space for which to retrieve assignments.
	roleId := "roleId_example" // string | Filters the returned role assignments to the provided role ID. (optional)
	roleType := "roleType_example" // string | Filters the returned role assignments to the provided role type. (optional)
	principalId := "principalId_example" // string | Filters the returned role assignments to the provided principal id. If specified, a principal-type must also be specified. Paired with a `principal-type` of `ACCESS_CLASS`, valid values include [`anonymous-users`, `jsm-project-admins`, `authenticated-users`, `all-licensed-users`, `all-product-admins`] (optional)
	principalType := openapiclient.PrincipalType("USER") // PrincipalType | Filters the returned role assignments to the provided principal type. If specified, a principal-id must also be specified. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of space roles to return. If more results exist, use the `Link` response header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceRolesAPI.GetSpaceRoleAssignments(context.Background(), id).RoleId(roleId).RoleType(roleType).PrincipalId(principalId).PrincipalType(principalType).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceRolesAPI.GetSpaceRoleAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpaceRoleAssignments`: MultiEntityResultSpaceRoleAssignment
	fmt.Fprintf(os.Stdout, "Response from `SpaceRolesAPI.GetSpaceRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the space for which to retrieve assignments. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpaceRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleId** | **string** | Filters the returned role assignments to the provided role ID. | 
 **roleType** | **string** | Filters the returned role assignments to the provided role type. | 
 **principalId** | **string** | Filters the returned role assignments to the provided principal id. If specified, a principal-type must also be specified. Paired with a &#x60;principal-type&#x60; of &#x60;ACCESS_CLASS&#x60;, valid values include [&#x60;anonymous-users&#x60;, &#x60;jsm-project-admins&#x60;, &#x60;authenticated-users&#x60;, &#x60;all-licensed-users&#x60;, &#x60;all-product-admins&#x60;] | 
 **principalType** | [**PrincipalType**](PrincipalType.md) | Filters the returned role assignments to the provided principal type. If specified, a principal-id must also be specified. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of space roles to return. If more results exist, use the &#x60;Link&#x60; response header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultSpaceRoleAssignment**](MultiEntityResultSpaceRoleAssignment.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpaceRoleMode

> GetSpaceRoleMode200Response GetSpaceRoleMode(ctx).Execute()

Get space role mode



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceRolesAPI.GetSpaceRoleMode(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceRolesAPI.GetSpaceRoleMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpaceRoleMode`: GetSpaceRoleMode200Response
	fmt.Fprintf(os.Stdout, "Response from `SpaceRolesAPI.GetSpaceRoleMode`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpaceRoleModeRequest struct via the builder pattern


### Return type

[**GetSpaceRoleMode200Response**](GetSpaceRoleMode200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpaceRolesById

> GetSpaceRolesById200Response GetSpaceRolesById(ctx, id).Execute()

Get space role by ID



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
	id := int32(56) // int32 | The ID of the space role to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceRolesAPI.GetSpaceRolesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceRolesAPI.GetSpaceRolesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpaceRolesById`: GetSpaceRolesById200Response
	fmt.Fprintf(os.Stdout, "Response from `SpaceRolesAPI.GetSpaceRolesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the space role to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpaceRolesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSpaceRolesById200Response**](GetSpaceRolesById200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSpaceRoleAssignments

> MultiEntityResultSpaceRoleAssignment SetSpaceRoleAssignments(ctx, id).SetSpaceRoleAssignmentsRequestInner(setSpaceRoleAssignmentsRequestInner).Execute()

Set space role assignments



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
	id := int32(56) // int32 | The ID of the space for which to retrieve assignments.
	setSpaceRoleAssignmentsRequestInner := []openapiclient.SetSpaceRoleAssignmentsRequestInner{*openapiclient.NewSetSpaceRoleAssignmentsRequestInner(*openapiclient.NewPrincipal())} // []SetSpaceRoleAssignmentsRequestInner | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceRolesAPI.SetSpaceRoleAssignments(context.Background(), id).SetSpaceRoleAssignmentsRequestInner(setSpaceRoleAssignmentsRequestInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceRolesAPI.SetSpaceRoleAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetSpaceRoleAssignments`: MultiEntityResultSpaceRoleAssignment
	fmt.Fprintf(os.Stdout, "Response from `SpaceRolesAPI.SetSpaceRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the space for which to retrieve assignments. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSpaceRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setSpaceRoleAssignmentsRequestInner** | [**[]SetSpaceRoleAssignmentsRequestInner**](SetSpaceRoleAssignmentsRequestInner.md) |  | 

### Return type

[**MultiEntityResultSpaceRoleAssignment**](MultiEntityResultSpaceRoleAssignment.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpaceRole

> UpdateSpaceRoleResponse UpdateSpaceRole(ctx, id).UpdateSpaceRoleRequest(updateSpaceRoleRequest).Execute()

Update a space role



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
	id := "id_example" // string | Id of the space role
	updateSpaceRoleRequest := *openapiclient.NewUpdateSpaceRoleRequest("Name_example", "Description_example", []string{"SpacePermissions_example"}) // UpdateSpaceRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceRolesAPI.UpdateSpaceRole(context.Background(), id).UpdateSpaceRoleRequest(updateSpaceRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceRolesAPI.UpdateSpaceRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSpaceRole`: UpdateSpaceRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `SpaceRolesAPI.UpdateSpaceRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the space role | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpaceRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSpaceRoleRequest** | [**UpdateSpaceRoleRequest**](UpdateSpaceRoleRequest.md) |  | 

### Return type

[**UpdateSpaceRoleResponse**](UpdateSpaceRoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \SpacePropertiesAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSpaceProperty**](SpacePropertiesAPI.md#CreateSpaceProperty) | **Post** /spaces/{space-id}/properties | Create space property in space
[**DeleteSpacePropertyById**](SpacePropertiesAPI.md#DeleteSpacePropertyById) | **Delete** /spaces/{space-id}/properties/{property-id} | Delete space property by id
[**GetSpaceProperties**](SpacePropertiesAPI.md#GetSpaceProperties) | **Get** /spaces/{space-id}/properties | Get space properties in space
[**GetSpacePropertyById**](SpacePropertiesAPI.md#GetSpacePropertyById) | **Get** /spaces/{space-id}/properties/{property-id} | Get space property by id
[**UpdateSpacePropertyById**](SpacePropertiesAPI.md#UpdateSpacePropertyById) | **Put** /spaces/{space-id}/properties/{property-id} | Update space property by id



## CreateSpaceProperty

> SpaceProperty CreateSpaceProperty(ctx, spaceId).SpacePropertyCreateRequest(spacePropertyCreateRequest).Execute()

Create space property in space



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
	spaceId := int64(789) // int64 | The ID of the space for which space properties should be returned.
	spacePropertyCreateRequest := *openapiclient.NewSpacePropertyCreateRequest() // SpacePropertyCreateRequest | The space property to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacePropertiesAPI.CreateSpaceProperty(context.Background(), spaceId).SpacePropertyCreateRequest(spacePropertyCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacePropertiesAPI.CreateSpaceProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSpaceProperty`: SpaceProperty
	fmt.Fprintf(os.Stdout, "Response from `SpacePropertiesAPI.CreateSpaceProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **int64** | The ID of the space for which space properties should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpacePropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spacePropertyCreateRequest** | [**SpacePropertyCreateRequest**](SpacePropertyCreateRequest.md) | The space property to be created | 

### Return type

[**SpaceProperty**](SpaceProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSpacePropertyById

> DeleteSpacePropertyById(ctx, spaceId, propertyId).Execute()

Delete space property by id



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
	spaceId := int64(789) // int64 | The ID of the space the property belongs to.
	propertyId := int64(789) // int64 | The ID of the property to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SpacePropertiesAPI.DeleteSpacePropertyById(context.Background(), spaceId, propertyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacePropertiesAPI.DeleteSpacePropertyById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **int64** | The ID of the space the property belongs to. | 
**propertyId** | **int64** | The ID of the property to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpacePropertyByIdRequest struct via the builder pattern


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


## GetSpaceProperties

> MultiEntityResultSpaceProperty GetSpaceProperties(ctx, spaceId).Key(key).Cursor(cursor).Limit(limit).Execute()

Get space properties in space



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
	spaceId := int64(789) // int64 | The ID of the space for which space properties should be returned.
	key := "key_example" // string | The key of the space property to retrieve. This should be used when a user knows the key of their property, but needs to retrieve the id for use in other methods. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of pages per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacePropertiesAPI.GetSpaceProperties(context.Background(), spaceId).Key(key).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacePropertiesAPI.GetSpaceProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpaceProperties`: MultiEntityResultSpaceProperty
	fmt.Fprintf(os.Stdout, "Response from `SpacePropertiesAPI.GetSpaceProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **int64** | The ID of the space for which space properties should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpacePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** | The key of the space property to retrieve. This should be used when a user knows the key of their property, but needs to retrieve the id for use in other methods. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of pages per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultSpaceProperty**](MultiEntityResultSpaceProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpacePropertyById

> SpaceProperty GetSpacePropertyById(ctx, spaceId, propertyId).Execute()

Get space property by id



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
	spaceId := int64(789) // int64 | The ID of the space the property belongs to.
	propertyId := int64(789) // int64 | The ID of the property to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacePropertiesAPI.GetSpacePropertyById(context.Background(), spaceId, propertyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacePropertiesAPI.GetSpacePropertyById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpacePropertyById`: SpaceProperty
	fmt.Fprintf(os.Stdout, "Response from `SpacePropertiesAPI.GetSpacePropertyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **int64** | The ID of the space the property belongs to. | 
**propertyId** | **int64** | The ID of the property to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpacePropertyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SpaceProperty**](SpaceProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpacePropertyById

> SpaceProperty UpdateSpacePropertyById(ctx, spaceId, propertyId).SpacePropertyUpdateRequest(spacePropertyUpdateRequest).Execute()

Update space property by id



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
	spaceId := int64(789) // int64 | The ID of the space the property belongs to.
	propertyId := int64(789) // int64 | The ID of the property to be updated.
	spacePropertyUpdateRequest := *openapiclient.NewSpacePropertyUpdateRequest() // SpacePropertyUpdateRequest | The space property to be updated.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacePropertiesAPI.UpdateSpacePropertyById(context.Background(), spaceId, propertyId).SpacePropertyUpdateRequest(spacePropertyUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacePropertiesAPI.UpdateSpacePropertyById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSpacePropertyById`: SpaceProperty
	fmt.Fprintf(os.Stdout, "Response from `SpacePropertiesAPI.UpdateSpacePropertyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **int64** | The ID of the space the property belongs to. | 
**propertyId** | **int64** | The ID of the property to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpacePropertyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **spacePropertyUpdateRequest** | [**SpacePropertyUpdateRequest**](SpacePropertyUpdateRequest.md) | The space property to be updated. | 

### Return type

[**SpaceProperty**](SpaceProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


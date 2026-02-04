# \AppPropertiesAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteForgeAppProperty**](AppPropertiesAPI.md#DeleteForgeAppProperty) | **Delete** /app/properties/{propertyKey} | Deletes a Forge app property.
[**GetForgeAppProperties**](AppPropertiesAPI.md#GetForgeAppProperties) | **Get** /app/properties | Get Forge app properties.
[**GetForgeAppProperty**](AppPropertiesAPI.md#GetForgeAppProperty) | **Get** /app/properties/{propertyKey} | Get a Forge app property by key.
[**PutForgeAppProperty**](AppPropertiesAPI.md#PutForgeAppProperty) | **Put** /app/properties/{propertyKey} | Create or update a Forge app property.



## DeleteForgeAppProperty

> DeleteForgeAppProperty(ctx, propertyKey).Execute()

Deletes a Forge app property.



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
	propertyKey := "propertyKey_example" // string | The key of the property

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppPropertiesAPI.DeleteForgeAppProperty(context.Background(), propertyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPropertiesAPI.DeleteForgeAppProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propertyKey** | **string** | The key of the property | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteForgeAppPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetForgeAppProperties

> MultiEntityResultAppProperty GetForgeAppProperties(ctx).Cursor(cursor).Limit(limit).Execute()

Get Forge app properties.



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
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor represents the last returned property key. It will be included in the response body as the next link. Use this key to request the next set of results. (optional)
	limit := int32(56) // int32 | Maximum number of app properties per result to return. If more results exist, use the last returned property key from the Link field in the response body as a cursor to retrieve the next set of results. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppPropertiesAPI.GetForgeAppProperties(context.Background()).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPropertiesAPI.GetForgeAppProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetForgeAppProperties`: MultiEntityResultAppProperty
	fmt.Fprintf(os.Stdout, "Response from `AppPropertiesAPI.GetForgeAppProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetForgeAppPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Used for pagination, this opaque cursor represents the last returned property key. It will be included in the response body as the next link. Use this key to request the next set of results. | 
 **limit** | **int32** | Maximum number of app properties per result to return. If more results exist, use the last returned property key from the Link field in the response body as a cursor to retrieve the next set of results. | [default to 50]

### Return type

[**MultiEntityResultAppProperty**](MultiEntityResultAppProperty.md)

### Authorization

[oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetForgeAppProperty

> GetForgeAppProperty200Response GetForgeAppProperty(ctx, propertyKey).Execute()

Get a Forge app property by key.



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
	propertyKey := "propertyKey_example" // string | The key of the property

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppPropertiesAPI.GetForgeAppProperty(context.Background(), propertyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPropertiesAPI.GetForgeAppProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetForgeAppProperty`: GetForgeAppProperty200Response
	fmt.Fprintf(os.Stdout, "Response from `AppPropertiesAPI.GetForgeAppProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propertyKey** | **string** | The key of the property | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetForgeAppPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetForgeAppProperty200Response**](GetForgeAppProperty200Response.md)

### Authorization

[oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutForgeAppProperty

> PutForgeAppProperty(ctx, propertyKey).Body(body).Execute()

Create or update a Forge app property.



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
	propertyKey := "propertyKey_example" // string | The key of the property
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppPropertiesAPI.PutForgeAppProperty(context.Background(), propertyKey).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPropertiesAPI.PutForgeAppProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propertyKey** | **string** | The key of the property | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutForgeAppPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


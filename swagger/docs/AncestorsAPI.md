# \AncestorsAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDatabaseAncestors**](AncestorsAPI.md#GetDatabaseAncestors) | **Get** /databases/{id}/ancestors | Get all ancestors of database
[**GetFolderAncestors**](AncestorsAPI.md#GetFolderAncestors) | **Get** /folders/{id}/ancestors | Get all ancestors of folder
[**GetPageAncestors**](AncestorsAPI.md#GetPageAncestors) | **Get** /pages/{id}/ancestors | Get all ancestors of page
[**GetSmartLinkAncestors**](AncestorsAPI.md#GetSmartLinkAncestors) | **Get** /embeds/{id}/ancestors | Get all ancestors of Smart Link in content tree
[**GetWhiteboardAncestors**](AncestorsAPI.md#GetWhiteboardAncestors) | **Get** /whiteboards/{id}/ancestors | Get all ancestors of whiteboard



## GetDatabaseAncestors

> MultiEntityResultAncestor GetDatabaseAncestors(ctx, id).Limit(limit).Execute()

Get all ancestors of database



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
	id := int64(789) // int64 | The ID of the database.
	limit := int32(56) // int32 | Maximum number of items per result to return. If more results exist, call the endpoint with the highest ancestor's ID to fetch the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AncestorsAPI.GetDatabaseAncestors(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AncestorsAPI.GetDatabaseAncestors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabaseAncestors`: MultiEntityResultAncestor
	fmt.Fprintf(os.Stdout, "Response from `AncestorsAPI.GetDatabaseAncestors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseAncestorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of items per result to return. If more results exist, call the endpoint with the highest ancestor&#39;s ID to fetch the next set of results. | [default to 25]

### Return type

[**MultiEntityResultAncestor**](MultiEntityResultAncestor.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolderAncestors

> MultiEntityResultAncestor GetFolderAncestors(ctx, id).Limit(limit).Execute()

Get all ancestors of folder



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
	id := int64(789) // int64 | The ID of the folder.
	limit := int32(56) // int32 | Maximum number of items per result to return. If more results exist, call the endpoint with the highest ancestor's ID to fetch the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AncestorsAPI.GetFolderAncestors(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AncestorsAPI.GetFolderAncestors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolderAncestors`: MultiEntityResultAncestor
	fmt.Fprintf(os.Stdout, "Response from `AncestorsAPI.GetFolderAncestors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the folder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderAncestorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of items per result to return. If more results exist, call the endpoint with the highest ancestor&#39;s ID to fetch the next set of results. | [default to 25]

### Return type

[**MultiEntityResultAncestor**](MultiEntityResultAncestor.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageAncestors

> MultiEntityResultAncestor1 GetPageAncestors(ctx, id).Limit(limit).Execute()

Get all ancestors of page



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
	id := int64(789) // int64 | The ID of the page.
	limit := int32(56) // int32 | Maximum number of pages per result to return. If more results exist, call this endpoint with the highest ancestor's ID to fetch the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AncestorsAPI.GetPageAncestors(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AncestorsAPI.GetPageAncestors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPageAncestors`: MultiEntityResultAncestor1
	fmt.Fprintf(os.Stdout, "Response from `AncestorsAPI.GetPageAncestors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the page. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageAncestorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of pages per result to return. If more results exist, call this endpoint with the highest ancestor&#39;s ID to fetch the next set of results. | [default to 25]

### Return type

[**MultiEntityResultAncestor1**](MultiEntityResultAncestor1.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmartLinkAncestors

> MultiEntityResultAncestor GetSmartLinkAncestors(ctx, id).Limit(limit).Execute()

Get all ancestors of Smart Link in content tree



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
	id := int64(789) // int64 | The ID of the Smart Link in the content tree.
	limit := int32(56) // int32 | Maximum number of items per result to return. If more results exist, call the endpoint with the highest ancestor's ID to fetch the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AncestorsAPI.GetSmartLinkAncestors(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AncestorsAPI.GetSmartLinkAncestors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSmartLinkAncestors`: MultiEntityResultAncestor
	fmt.Fprintf(os.Stdout, "Response from `AncestorsAPI.GetSmartLinkAncestors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the Smart Link in the content tree. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmartLinkAncestorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of items per result to return. If more results exist, call the endpoint with the highest ancestor&#39;s ID to fetch the next set of results. | [default to 25]

### Return type

[**MultiEntityResultAncestor**](MultiEntityResultAncestor.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhiteboardAncestors

> MultiEntityResultAncestor GetWhiteboardAncestors(ctx, id).Limit(limit).Execute()

Get all ancestors of whiteboard



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
	id := int64(789) // int64 | The ID of the whiteboard.
	limit := int32(56) // int32 | Maximum number of items per result to return. If more results exist, call the endpoint with the highest ancestor's ID to fetch the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AncestorsAPI.GetWhiteboardAncestors(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AncestorsAPI.GetWhiteboardAncestors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhiteboardAncestors`: MultiEntityResultAncestor
	fmt.Fprintf(os.Stdout, "Response from `AncestorsAPI.GetWhiteboardAncestors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the whiteboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhiteboardAncestorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of items per result to return. If more results exist, call the endpoint with the highest ancestor&#39;s ID to fetch the next set of results. | [default to 25]

### Return type

[**MultiEntityResultAncestor**](MultiEntityResultAncestor.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


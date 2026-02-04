# \OperationAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAttachmentOperations**](OperationAPI.md#GetAttachmentOperations) | **Get** /attachments/{id}/operations | Get permitted operations for attachment
[**GetBlogPostOperations**](OperationAPI.md#GetBlogPostOperations) | **Get** /blogposts/{id}/operations | Get permitted operations for blog post
[**GetCustomContentOperations**](OperationAPI.md#GetCustomContentOperations) | **Get** /custom-content/{id}/operations | Get permitted operations for custom content
[**GetDatabaseOperations**](OperationAPI.md#GetDatabaseOperations) | **Get** /databases/{id}/operations | Get permitted operations for a database
[**GetFolderOperations**](OperationAPI.md#GetFolderOperations) | **Get** /folders/{id}/operations | Get permitted operations for a folder
[**GetFooterCommentOperations**](OperationAPI.md#GetFooterCommentOperations) | **Get** /footer-comments/{id}/operations | Get permitted operations for footer comment
[**GetInlineCommentOperations**](OperationAPI.md#GetInlineCommentOperations) | **Get** /inline-comments/{id}/operations | Get permitted operations for inline comment
[**GetPageOperations**](OperationAPI.md#GetPageOperations) | **Get** /pages/{id}/operations | Get permitted operations for page
[**GetSmartLinkOperations**](OperationAPI.md#GetSmartLinkOperations) | **Get** /embeds/{id}/operations | Get permitted operations for a Smart Link in the content tree
[**GetSpaceOperations**](OperationAPI.md#GetSpaceOperations) | **Get** /spaces/{id}/operations | Get permitted operations for space
[**GetWhiteboardOperations**](OperationAPI.md#GetWhiteboardOperations) | **Get** /whiteboards/{id}/operations | Get permitted operations for a whiteboard



## GetAttachmentOperations

> PermittedOperationsResponse GetAttachmentOperations(ctx, id).Execute()

Get permitted operations for attachment



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
	id := "id_example" // string | The ID of the attachment for which operations should be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationAPI.GetAttachmentOperations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationAPI.GetAttachmentOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachmentOperations`: PermittedOperationsResponse
	fmt.Fprintf(os.Stdout, "Response from `OperationAPI.GetAttachmentOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the attachment for which operations should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PermittedOperationsResponse**](PermittedOperationsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlogPostOperations

> PermittedOperationsResponse GetBlogPostOperations(ctx, id).Execute()

Get permitted operations for blog post



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
	id := int64(789) // int64 | The ID of the blog post for which operations should be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationAPI.GetBlogPostOperations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationAPI.GetBlogPostOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlogPostOperations`: PermittedOperationsResponse
	fmt.Fprintf(os.Stdout, "Response from `OperationAPI.GetBlogPostOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the blog post for which operations should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogPostOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PermittedOperationsResponse**](PermittedOperationsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomContentOperations

> PermittedOperationsResponse GetCustomContentOperations(ctx, id).Execute()

Get permitted operations for custom content



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
	id := int64(789) // int64 | The ID of the custom content for which operations should be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationAPI.GetCustomContentOperations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationAPI.GetCustomContentOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomContentOperations`: PermittedOperationsResponse
	fmt.Fprintf(os.Stdout, "Response from `OperationAPI.GetCustomContentOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the custom content for which operations should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomContentOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PermittedOperationsResponse**](PermittedOperationsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseOperations

> PermittedOperationsResponse GetDatabaseOperations(ctx, id).Execute()

Get permitted operations for a database



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
	id := int64(789) // int64 | The ID of the database for which operations should be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationAPI.GetDatabaseOperations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationAPI.GetDatabaseOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabaseOperations`: PermittedOperationsResponse
	fmt.Fprintf(os.Stdout, "Response from `OperationAPI.GetDatabaseOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the database for which operations should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PermittedOperationsResponse**](PermittedOperationsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolderOperations

> PermittedOperationsResponse GetFolderOperations(ctx, id).Execute()

Get permitted operations for a folder



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
	id := int64(789) // int64 | The ID of the folder for which operations should be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationAPI.GetFolderOperations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationAPI.GetFolderOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolderOperations`: PermittedOperationsResponse
	fmt.Fprintf(os.Stdout, "Response from `OperationAPI.GetFolderOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the folder for which operations should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PermittedOperationsResponse**](PermittedOperationsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFooterCommentOperations

> PermittedOperationsResponse GetFooterCommentOperations(ctx, id).Execute()

Get permitted operations for footer comment



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
	id := int64(789) // int64 | The ID of the footer comment for which operations should be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationAPI.GetFooterCommentOperations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationAPI.GetFooterCommentOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFooterCommentOperations`: PermittedOperationsResponse
	fmt.Fprintf(os.Stdout, "Response from `OperationAPI.GetFooterCommentOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the footer comment for which operations should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFooterCommentOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PermittedOperationsResponse**](PermittedOperationsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInlineCommentOperations

> PermittedOperationsResponse GetInlineCommentOperations(ctx, id).Execute()

Get permitted operations for inline comment



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
	id := int64(789) // int64 | The ID of the inline comment for which operations should be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationAPI.GetInlineCommentOperations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationAPI.GetInlineCommentOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInlineCommentOperations`: PermittedOperationsResponse
	fmt.Fprintf(os.Stdout, "Response from `OperationAPI.GetInlineCommentOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the inline comment for which operations should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInlineCommentOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PermittedOperationsResponse**](PermittedOperationsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageOperations

> PermittedOperationsResponse GetPageOperations(ctx, id).Execute()

Get permitted operations for page



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
	id := int64(789) // int64 | The ID of the page for which operations should be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationAPI.GetPageOperations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationAPI.GetPageOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPageOperations`: PermittedOperationsResponse
	fmt.Fprintf(os.Stdout, "Response from `OperationAPI.GetPageOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the page for which operations should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PermittedOperationsResponse**](PermittedOperationsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmartLinkOperations

> PermittedOperationsResponse GetSmartLinkOperations(ctx, id).Execute()

Get permitted operations for a Smart Link in the content tree



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
	id := int64(789) // int64 | The ID of the Smart Link in the content tree for which operations should be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationAPI.GetSmartLinkOperations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationAPI.GetSmartLinkOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSmartLinkOperations`: PermittedOperationsResponse
	fmt.Fprintf(os.Stdout, "Response from `OperationAPI.GetSmartLinkOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the Smart Link in the content tree for which operations should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmartLinkOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PermittedOperationsResponse**](PermittedOperationsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpaceOperations

> PermittedOperationsResponse GetSpaceOperations(ctx, id).Execute()

Get permitted operations for space



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
	id := int64(789) // int64 | The ID of the space for which operations should be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationAPI.GetSpaceOperations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationAPI.GetSpaceOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpaceOperations`: PermittedOperationsResponse
	fmt.Fprintf(os.Stdout, "Response from `OperationAPI.GetSpaceOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the space for which operations should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpaceOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PermittedOperationsResponse**](PermittedOperationsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhiteboardOperations

> PermittedOperationsResponse GetWhiteboardOperations(ctx, id).Execute()

Get permitted operations for a whiteboard



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
	id := int64(789) // int64 | The ID of the whiteboard for which operations should be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationAPI.GetWhiteboardOperations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationAPI.GetWhiteboardOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhiteboardOperations`: PermittedOperationsResponse
	fmt.Fprintf(os.Stdout, "Response from `OperationAPI.GetWhiteboardOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the whiteboard for which operations should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhiteboardOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PermittedOperationsResponse**](PermittedOperationsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


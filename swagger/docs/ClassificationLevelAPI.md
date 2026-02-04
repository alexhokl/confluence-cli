# \ClassificationLevelAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSpaceDefaultClassificationLevel**](ClassificationLevelAPI.md#DeleteSpaceDefaultClassificationLevel) | **Delete** /spaces/{id}/classification-level/default | Delete space default classification level
[**GetBlogPostClassificationLevel**](ClassificationLevelAPI.md#GetBlogPostClassificationLevel) | **Get** /blogposts/{id}/classification-level | Get blog post classification level
[**GetClassificationLevels**](ClassificationLevelAPI.md#GetClassificationLevels) | **Get** /classification-levels | Get list of classification levels
[**GetDatabaseClassificationLevel**](ClassificationLevelAPI.md#GetDatabaseClassificationLevel) | **Get** /databases/{id}/classification-level | Get database classification level
[**GetPageClassificationLevel**](ClassificationLevelAPI.md#GetPageClassificationLevel) | **Get** /pages/{id}/classification-level | Get page classification level
[**GetSpaceDefaultClassificationLevel**](ClassificationLevelAPI.md#GetSpaceDefaultClassificationLevel) | **Get** /spaces/{id}/classification-level/default | Get space default classification level
[**GetWhiteboardClassificationLevel**](ClassificationLevelAPI.md#GetWhiteboardClassificationLevel) | **Get** /whiteboards/{id}/classification-level | Get whiteboard classification level
[**PostBlogPostClassificationLevel**](ClassificationLevelAPI.md#PostBlogPostClassificationLevel) | **Post** /blogposts/{id}/classification-level/reset | Reset blog post classification level
[**PostDatabaseClassificationLevel**](ClassificationLevelAPI.md#PostDatabaseClassificationLevel) | **Post** /databases/{id}/classification-level/reset | Reset database classification level
[**PostPageClassificationLevel**](ClassificationLevelAPI.md#PostPageClassificationLevel) | **Post** /pages/{id}/classification-level/reset | Reset page classification level
[**PostWhiteboardClassificationLevel**](ClassificationLevelAPI.md#PostWhiteboardClassificationLevel) | **Post** /whiteboards/{id}/classification-level/reset | Reset whiteboard classification level
[**PutBlogPostClassificationLevel**](ClassificationLevelAPI.md#PutBlogPostClassificationLevel) | **Put** /blogposts/{id}/classification-level | Update blog post classification level
[**PutDatabaseClassificationLevel**](ClassificationLevelAPI.md#PutDatabaseClassificationLevel) | **Put** /databases/{id}/classification-level | Update database classification level
[**PutPageClassificationLevel**](ClassificationLevelAPI.md#PutPageClassificationLevel) | **Put** /pages/{id}/classification-level | Update page classification level
[**PutSpaceDefaultClassificationLevel**](ClassificationLevelAPI.md#PutSpaceDefaultClassificationLevel) | **Put** /spaces/{id}/classification-level/default | Update space default classification level
[**PutWhiteboardClassificationLevel**](ClassificationLevelAPI.md#PutWhiteboardClassificationLevel) | **Put** /whiteboards/{id}/classification-level | Update whiteboard classification level



## DeleteSpaceDefaultClassificationLevel

> DeleteSpaceDefaultClassificationLevel(ctx, id).Execute()

Delete space default classification level



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
	id := int64(789) // int64 | The ID of the space for which default classification level should be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClassificationLevelAPI.DeleteSpaceDefaultClassificationLevel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClassificationLevelAPI.DeleteSpaceDefaultClassificationLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the space for which default classification level should be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpaceDefaultClassificationLevelRequest struct via the builder pattern


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


## GetBlogPostClassificationLevel

> ClassificationLevel GetBlogPostClassificationLevel(ctx, id).Status(status).Execute()

Get blog post classification level



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
	id := int64(789) // int64 | The ID of the blog post for which classification level should be returned.
	status := "status_example" // string | Status of blog post from which classification level will fetched. (optional) (default to "current")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClassificationLevelAPI.GetBlogPostClassificationLevel(context.Background(), id).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClassificationLevelAPI.GetBlogPostClassificationLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlogPostClassificationLevel`: ClassificationLevel
	fmt.Fprintf(os.Stdout, "Response from `ClassificationLevelAPI.GetBlogPostClassificationLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the blog post for which classification level should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogPostClassificationLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Status of blog post from which classification level will fetched. | [default to &quot;current&quot;]

### Return type

[**ClassificationLevel**](ClassificationLevel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClassificationLevels

> []ClassificationLevel GetClassificationLevels(ctx).Execute()

Get list of classification levels



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
	resp, r, err := apiClient.ClassificationLevelAPI.GetClassificationLevels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClassificationLevelAPI.GetClassificationLevels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClassificationLevels`: []ClassificationLevel
	fmt.Fprintf(os.Stdout, "Response from `ClassificationLevelAPI.GetClassificationLevels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClassificationLevelsRequest struct via the builder pattern


### Return type

[**[]ClassificationLevel**](ClassificationLevel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseClassificationLevel

> ClassificationLevel GetDatabaseClassificationLevel(ctx, id).Execute()

Get database classification level



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
	id := int64(789) // int64 | The ID of the database for which classification level should be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClassificationLevelAPI.GetDatabaseClassificationLevel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClassificationLevelAPI.GetDatabaseClassificationLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabaseClassificationLevel`: ClassificationLevel
	fmt.Fprintf(os.Stdout, "Response from `ClassificationLevelAPI.GetDatabaseClassificationLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the database for which classification level should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseClassificationLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClassificationLevel**](ClassificationLevel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageClassificationLevel

> ClassificationLevel GetPageClassificationLevel(ctx, id).Status(status).Execute()

Get page classification level



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
	id := int64(789) // int64 | The ID of the page for which classification level should be returned.
	status := "status_example" // string | Status of page from which classification level will fetched. (optional) (default to "current")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClassificationLevelAPI.GetPageClassificationLevel(context.Background(), id).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClassificationLevelAPI.GetPageClassificationLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPageClassificationLevel`: ClassificationLevel
	fmt.Fprintf(os.Stdout, "Response from `ClassificationLevelAPI.GetPageClassificationLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the page for which classification level should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageClassificationLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Status of page from which classification level will fetched. | [default to &quot;current&quot;]

### Return type

[**ClassificationLevel**](ClassificationLevel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpaceDefaultClassificationLevel

> ClassificationLevel GetSpaceDefaultClassificationLevel(ctx, id).Execute()

Get space default classification level



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
	id := int64(789) // int64 | The ID of the space for which default classification level should be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClassificationLevelAPI.GetSpaceDefaultClassificationLevel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClassificationLevelAPI.GetSpaceDefaultClassificationLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpaceDefaultClassificationLevel`: ClassificationLevel
	fmt.Fprintf(os.Stdout, "Response from `ClassificationLevelAPI.GetSpaceDefaultClassificationLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the space for which default classification level should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpaceDefaultClassificationLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClassificationLevel**](ClassificationLevel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhiteboardClassificationLevel

> ClassificationLevel GetWhiteboardClassificationLevel(ctx, id).Execute()

Get whiteboard classification level



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
	id := int64(789) // int64 | The ID of the whiteboard for which classification level should be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClassificationLevelAPI.GetWhiteboardClassificationLevel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClassificationLevelAPI.GetWhiteboardClassificationLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhiteboardClassificationLevel`: ClassificationLevel
	fmt.Fprintf(os.Stdout, "Response from `ClassificationLevelAPI.GetWhiteboardClassificationLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the whiteboard for which classification level should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhiteboardClassificationLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClassificationLevel**](ClassificationLevel.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBlogPostClassificationLevel

> PostBlogPostClassificationLevel(ctx, id).PostPageClassificationLevelRequest(postPageClassificationLevelRequest).Execute()

Reset blog post classification level



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
	id := int64(789) // int64 | The ID of the blog post for which classification level should be updated.
	postPageClassificationLevelRequest := *openapiclient.NewPostPageClassificationLevelRequest("Status_example") // PostPageClassificationLevelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClassificationLevelAPI.PostBlogPostClassificationLevel(context.Background(), id).PostPageClassificationLevelRequest(postPageClassificationLevelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClassificationLevelAPI.PostBlogPostClassificationLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the blog post for which classification level should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBlogPostClassificationLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postPageClassificationLevelRequest** | [**PostPageClassificationLevelRequest**](PostPageClassificationLevelRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDatabaseClassificationLevel

> PostDatabaseClassificationLevel(ctx, id).PostWhiteboardClassificationLevelRequest(postWhiteboardClassificationLevelRequest).Execute()

Reset database classification level



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
	id := int64(789) // int64 | The ID of the database for which classification level should be updated.
	postWhiteboardClassificationLevelRequest := *openapiclient.NewPostWhiteboardClassificationLevelRequest("Status_example") // PostWhiteboardClassificationLevelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClassificationLevelAPI.PostDatabaseClassificationLevel(context.Background(), id).PostWhiteboardClassificationLevelRequest(postWhiteboardClassificationLevelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClassificationLevelAPI.PostDatabaseClassificationLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the database for which classification level should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDatabaseClassificationLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postWhiteboardClassificationLevelRequest** | [**PostWhiteboardClassificationLevelRequest**](PostWhiteboardClassificationLevelRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPageClassificationLevel

> PostPageClassificationLevel(ctx, id).PostPageClassificationLevelRequest(postPageClassificationLevelRequest).Execute()

Reset page classification level



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
	id := int64(789) // int64 | The ID of the page for which classification level should be updated.
	postPageClassificationLevelRequest := *openapiclient.NewPostPageClassificationLevelRequest("Status_example") // PostPageClassificationLevelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClassificationLevelAPI.PostPageClassificationLevel(context.Background(), id).PostPageClassificationLevelRequest(postPageClassificationLevelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClassificationLevelAPI.PostPageClassificationLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the page for which classification level should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPageClassificationLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postPageClassificationLevelRequest** | [**PostPageClassificationLevelRequest**](PostPageClassificationLevelRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWhiteboardClassificationLevel

> PostWhiteboardClassificationLevel(ctx, id).PostWhiteboardClassificationLevelRequest(postWhiteboardClassificationLevelRequest).Execute()

Reset whiteboard classification level



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
	id := int64(789) // int64 | The ID of the whiteboard for which classification level should be updated.
	postWhiteboardClassificationLevelRequest := *openapiclient.NewPostWhiteboardClassificationLevelRequest("Status_example") // PostWhiteboardClassificationLevelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClassificationLevelAPI.PostWhiteboardClassificationLevel(context.Background(), id).PostWhiteboardClassificationLevelRequest(postWhiteboardClassificationLevelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClassificationLevelAPI.PostWhiteboardClassificationLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the whiteboard for which classification level should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWhiteboardClassificationLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postWhiteboardClassificationLevelRequest** | [**PostWhiteboardClassificationLevelRequest**](PostWhiteboardClassificationLevelRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBlogPostClassificationLevel

> PutBlogPostClassificationLevel(ctx, id).PutPageClassificationLevelRequest(putPageClassificationLevelRequest).Execute()

Update blog post classification level



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
	id := int64(789) // int64 | The ID of the blog post for which classification level should be updated.
	putPageClassificationLevelRequest := *openapiclient.NewPutPageClassificationLevelRequest("Id_example", "Status_example") // PutPageClassificationLevelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClassificationLevelAPI.PutBlogPostClassificationLevel(context.Background(), id).PutPageClassificationLevelRequest(putPageClassificationLevelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClassificationLevelAPI.PutBlogPostClassificationLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the blog post for which classification level should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutBlogPostClassificationLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putPageClassificationLevelRequest** | [**PutPageClassificationLevelRequest**](PutPageClassificationLevelRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDatabaseClassificationLevel

> PutDatabaseClassificationLevel(ctx, id).PutWhiteboardClassificationLevelRequest(putWhiteboardClassificationLevelRequest).Execute()

Update database classification level



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
	id := int64(789) // int64 | The ID of the database for which classification level should be updated.
	putWhiteboardClassificationLevelRequest := *openapiclient.NewPutWhiteboardClassificationLevelRequest("Id_example", "Status_example") // PutWhiteboardClassificationLevelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClassificationLevelAPI.PutDatabaseClassificationLevel(context.Background(), id).PutWhiteboardClassificationLevelRequest(putWhiteboardClassificationLevelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClassificationLevelAPI.PutDatabaseClassificationLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the database for which classification level should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDatabaseClassificationLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putWhiteboardClassificationLevelRequest** | [**PutWhiteboardClassificationLevelRequest**](PutWhiteboardClassificationLevelRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPageClassificationLevel

> PutPageClassificationLevel(ctx, id).PutPageClassificationLevelRequest(putPageClassificationLevelRequest).Execute()

Update page classification level



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
	id := int64(789) // int64 | The ID of the page for which classification level should be updated.
	putPageClassificationLevelRequest := *openapiclient.NewPutPageClassificationLevelRequest("Id_example", "Status_example") // PutPageClassificationLevelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClassificationLevelAPI.PutPageClassificationLevel(context.Background(), id).PutPageClassificationLevelRequest(putPageClassificationLevelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClassificationLevelAPI.PutPageClassificationLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the page for which classification level should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPageClassificationLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putPageClassificationLevelRequest** | [**PutPageClassificationLevelRequest**](PutPageClassificationLevelRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSpaceDefaultClassificationLevel

> PutSpaceDefaultClassificationLevel(ctx, id).PutSpaceDefaultClassificationLevelRequest(putSpaceDefaultClassificationLevelRequest).Execute()

Update space default classification level



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
	id := int64(789) // int64 | The ID of the space for which default classification level should be updated.
	putSpaceDefaultClassificationLevelRequest := *openapiclient.NewPutSpaceDefaultClassificationLevelRequest("Id_example") // PutSpaceDefaultClassificationLevelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClassificationLevelAPI.PutSpaceDefaultClassificationLevel(context.Background(), id).PutSpaceDefaultClassificationLevelRequest(putSpaceDefaultClassificationLevelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClassificationLevelAPI.PutSpaceDefaultClassificationLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the space for which default classification level should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSpaceDefaultClassificationLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putSpaceDefaultClassificationLevelRequest** | [**PutSpaceDefaultClassificationLevelRequest**](PutSpaceDefaultClassificationLevelRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutWhiteboardClassificationLevel

> PutWhiteboardClassificationLevel(ctx, id).PutWhiteboardClassificationLevelRequest(putWhiteboardClassificationLevelRequest).Execute()

Update whiteboard classification level



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
	id := int64(789) // int64 | The ID of the whiteboard for which classification level should be updated.
	putWhiteboardClassificationLevelRequest := *openapiclient.NewPutWhiteboardClassificationLevelRequest("Id_example", "Status_example") // PutWhiteboardClassificationLevelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClassificationLevelAPI.PutWhiteboardClassificationLevel(context.Background(), id).PutWhiteboardClassificationLevelRequest(putWhiteboardClassificationLevelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClassificationLevelAPI.PutWhiteboardClassificationLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the whiteboard for which classification level should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutWhiteboardClassificationLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putWhiteboardClassificationLevelRequest** | [**PutWhiteboardClassificationLevelRequest**](PutWhiteboardClassificationLevelRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


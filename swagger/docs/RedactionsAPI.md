# \RedactionsAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostRedactBlog**](RedactionsAPI.md#PostRedactBlog) | **Post** /blogposts/{id}/redact | Redact Content in a Confluence Blog Post
[**PostRedactPage**](RedactionsAPI.md#PostRedactPage) | **Post** /pages/{id}/redact | Redact Content in a Confluence Page



## PostRedactBlog

> RedactionResponse PostRedactBlog(ctx, id).PostRedactPageRequest(postRedactPageRequest).Execute()

Redact Content in a Confluence Blog Post



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/alexhokl/confluence-cli/swagger"
)

func main() {
	id := int64(789) // int64 | The ID of the blog post to redact content from.
	postRedactPageRequest := *openapiclient.NewPostRedactPageRequest(time.Now()) // PostRedactPageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RedactionsAPI.PostRedactBlog(context.Background(), id).PostRedactPageRequest(postRedactPageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedactionsAPI.PostRedactBlog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRedactBlog`: RedactionResponse
	fmt.Fprintf(os.Stdout, "Response from `RedactionsAPI.PostRedactBlog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the blog post to redact content from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRedactBlogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postRedactPageRequest** | [**PostRedactPageRequest**](PostRedactPageRequest.md) |  | 

### Return type

[**RedactionResponse**](RedactionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedactPage

> RedactionResponse PostRedactPage(ctx, id).PostRedactPageRequest(postRedactPageRequest).Execute()

Redact Content in a Confluence Page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/alexhokl/confluence-cli/swagger"
)

func main() {
	id := int64(789) // int64 | The ID of the page to redact content from.
	postRedactPageRequest := *openapiclient.NewPostRedactPageRequest(time.Now()) // PostRedactPageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RedactionsAPI.PostRedactPage(context.Background(), id).PostRedactPageRequest(postRedactPageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedactionsAPI.PostRedactPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRedactPage`: RedactionResponse
	fmt.Fprintf(os.Stdout, "Response from `RedactionsAPI.PostRedactPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the page to redact content from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRedactPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postRedactPageRequest** | [**PostRedactPageRequest**](PostRedactPageRequest.md) |  | 

### Return type

[**RedactionResponse**](RedactionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


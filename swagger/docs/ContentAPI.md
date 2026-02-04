# \ContentAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConvertContentIdsToContentTypes**](ContentAPI.md#ConvertContentIdsToContentTypes) | **Post** /content/convert-ids-to-types | Convert content ids to content types



## ConvertContentIdsToContentTypes

> ContentIdToContentTypeResponse ConvertContentIdsToContentTypes(ctx).ConvertContentIdsToContentTypesRequest(convertContentIdsToContentTypesRequest).Execute()

Convert content ids to content types



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
	convertContentIdsToContentTypesRequest := *openapiclient.NewConvertContentIdsToContentTypesRequest([]openapiclient.ConvertContentIdsToContentTypesRequestContentIdsInner{*openapiclient.NewConvertContentIdsToContentTypesRequestContentIdsInner()}) // ConvertContentIdsToContentTypesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.ConvertContentIdsToContentTypes(context.Background()).ConvertContentIdsToContentTypesRequest(convertContentIdsToContentTypesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.ConvertContentIdsToContentTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConvertContentIdsToContentTypes`: ContentIdToContentTypeResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.ConvertContentIdsToContentTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConvertContentIdsToContentTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **convertContentIdsToContentTypesRequest** | [**ConvertContentIdsToContentTypesRequest**](ConvertContentIdsToContentTypesRequest.md) |  | 

### Return type

[**ContentIdToContentTypeResponse**](ContentIdToContentTypeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


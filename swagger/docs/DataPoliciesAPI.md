# \DataPoliciesAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDataPolicyMetadata**](DataPoliciesAPI.md#GetDataPolicyMetadata) | **Get** /data-policies/metadata | Get data policy metadata for the workspace
[**GetDataPolicySpaces**](DataPoliciesAPI.md#GetDataPolicySpaces) | **Get** /data-policies/spaces | Get spaces with data policies



## GetDataPolicyMetadata

> DataPolicyMetadata GetDataPolicyMetadata(ctx).Execute()

Get data policy metadata for the workspace



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
	resp, r, err := apiClient.DataPoliciesAPI.GetDataPolicyMetadata(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataPoliciesAPI.GetDataPolicyMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataPolicyMetadata`: DataPolicyMetadata
	fmt.Fprintf(os.Stdout, "Response from `DataPoliciesAPI.GetDataPolicyMetadata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataPolicyMetadataRequest struct via the builder pattern


### Return type

[**DataPolicyMetadata**](DataPolicyMetadata.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataPolicySpaces

> MultiEntityResultDataPolicySpace GetDataPolicySpaces(ctx).Ids(ids).Keys(keys).Sort(sort).Cursor(cursor).Limit(limit).Execute()

Get spaces with data policies



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
	sort := openapiclient.SpaceSortOrder("id") // SpaceSortOrder | Used to sort the result by a particular field. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of spaces per result to return. If more results exist, use the `Link` response header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataPoliciesAPI.GetDataPolicySpaces(context.Background()).Ids(ids).Keys(keys).Sort(sort).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataPoliciesAPI.GetDataPolicySpaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataPolicySpaces`: MultiEntityResultDataPolicySpace
	fmt.Fprintf(os.Stdout, "Response from `DataPoliciesAPI.GetDataPolicySpaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataPolicySpacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int64** | Filter the results to spaces based on their IDs. Multiple IDs can be specified as a comma-separated list. | 
 **keys** | **[]string** | Filter the results to spaces based on their keys. Multiple keys can be specified as a comma-separated list. | 
 **sort** | [**SpaceSortOrder**](SpaceSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of spaces per result to return. If more results exist, use the &#x60;Link&#x60; response header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultDataPolicySpace**](MultiEntityResultDataPolicySpace.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \AdminKeyAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableAdminKey**](AdminKeyAPI.md#DisableAdminKey) | **Delete** /admin-key | Disable Admin Key
[**EnableAdminKey**](AdminKeyAPI.md#EnableAdminKey) | **Post** /admin-key | Enable Admin Key
[**GetAdminKey**](AdminKeyAPI.md#GetAdminKey) | **Get** /admin-key | Get Admin Key



## DisableAdminKey

> DisableAdminKey(ctx).Execute()

Disable Admin Key



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
	r, err := apiClient.AdminKeyAPI.DisableAdminKey(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminKeyAPI.DisableAdminKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDisableAdminKeyRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableAdminKey

> AdminKeyResponse EnableAdminKey(ctx).EnableAdminKeyRequest(enableAdminKeyRequest).Execute()

Enable Admin Key



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
	enableAdminKeyRequest := *openapiclient.NewEnableAdminKeyRequest() // EnableAdminKeyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminKeyAPI.EnableAdminKey(context.Background()).EnableAdminKeyRequest(enableAdminKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminKeyAPI.EnableAdminKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableAdminKey`: AdminKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminKeyAPI.EnableAdminKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableAdminKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableAdminKeyRequest** | [**EnableAdminKeyRequest**](EnableAdminKeyRequest.md) |  | 

### Return type

[**AdminKeyResponse**](AdminKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminKey

> AdminKeyResponse GetAdminKey(ctx).Execute()

Get Admin Key



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
	resp, r, err := apiClient.AdminKeyAPI.GetAdminKey(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminKeyAPI.GetAdminKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminKey`: AdminKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminKeyAPI.GetAdminKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminKeyRequest struct via the builder pattern


### Return type

[**AdminKeyResponse**](AdminKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


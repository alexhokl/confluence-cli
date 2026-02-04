# \UserAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckAccessByEmail**](UserAPI.md#CheckAccessByEmail) | **Post** /user/access/check-access-by-email | Check site access for a list of emails
[**CreateBulkUserLookup**](UserAPI.md#CreateBulkUserLookup) | **Post** /users-bulk | Create bulk user lookup using ids
[**InviteByEmail**](UserAPI.md#InviteByEmail) | **Post** /user/access/invite-by-email | Invite a list of emails to the site



## CheckAccessByEmail

> CheckAccessByEmail200Response CheckAccessByEmail(ctx).CheckAccessByEmailRequest(checkAccessByEmailRequest).Execute()

Check site access for a list of emails



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
	checkAccessByEmailRequest := *openapiclient.NewCheckAccessByEmailRequest([]string{"Emails_example"}) // CheckAccessByEmailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.CheckAccessByEmail(context.Background()).CheckAccessByEmailRequest(checkAccessByEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.CheckAccessByEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckAccessByEmail`: CheckAccessByEmail200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.CheckAccessByEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckAccessByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkAccessByEmailRequest** | [**CheckAccessByEmailRequest**](CheckAccessByEmailRequest.md) |  | 

### Return type

[**CheckAccessByEmail200Response**](CheckAccessByEmail200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBulkUserLookup

> MultiEntityResultUser CreateBulkUserLookup(ctx).CreateBulkUserLookupRequest(createBulkUserLookupRequest).Execute()

Create bulk user lookup using ids



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
	createBulkUserLookupRequest := *openapiclient.NewCreateBulkUserLookupRequest([]string{"AccountIds_example"}) // CreateBulkUserLookupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.CreateBulkUserLookup(context.Background()).CreateBulkUserLookupRequest(createBulkUserLookupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.CreateBulkUserLookup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBulkUserLookup`: MultiEntityResultUser
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.CreateBulkUserLookup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBulkUserLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBulkUserLookupRequest** | [**CreateBulkUserLookupRequest**](CreateBulkUserLookupRequest.md) |  | 

### Return type

[**MultiEntityResultUser**](MultiEntityResultUser.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteByEmail

> InviteByEmail(ctx).CheckAccessByEmailRequest(checkAccessByEmailRequest).Execute()

Invite a list of emails to the site



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
	checkAccessByEmailRequest := *openapiclient.NewCheckAccessByEmailRequest([]string{"Emails_example"}) // CheckAccessByEmailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.InviteByEmail(context.Background()).CheckAccessByEmailRequest(checkAccessByEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.InviteByEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInviteByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkAccessByEmailRequest** | [**CheckAccessByEmailRequest**](CheckAccessByEmailRequest.md) |  | 

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


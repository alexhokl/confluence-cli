# \TaskAPI

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTaskById**](TaskAPI.md#GetTaskById) | **Get** /tasks/{id} | Get task by id
[**GetTasks**](TaskAPI.md#GetTasks) | **Get** /tasks | Get tasks
[**UpdateTask**](TaskAPI.md#UpdateTask) | **Put** /tasks/{id} | Update task



## GetTaskById

> Task GetTaskById(ctx, id).BodyFormat(bodyFormat).Execute()

Get task by id



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
	id := int64(789) // int64 | The ID of the task to be returned. If you don't know the task ID, use Get tasks and filter the results.
	bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.GetTaskById(context.Background(), id).BodyFormat(bodyFormat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.GetTaskById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskById`: Task
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.GetTaskById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the task to be returned. If you don&#39;t know the task ID, use Get tasks and filter the results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 

### Return type

[**Task**](Task.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasks

> MultiEntityResultTask GetTasks(ctx).BodyFormat(bodyFormat).IncludeBlankTasks(includeBlankTasks).Status(status).TaskId(taskId).SpaceId(spaceId).PageId(pageId).BlogpostId(blogpostId).CreatedBy(createdBy).AssignedTo(assignedTo).CompletedBy(completedBy).CreatedAtFrom(createdAtFrom).CreatedAtTo(createdAtTo).DueAtFrom(dueAtFrom).DueAtTo(dueAtTo).CompletedAtFrom(completedAtFrom).CompletedAtTo(completedAtTo).Cursor(cursor).Limit(limit).Execute()

Get tasks



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
	bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
	includeBlankTasks := true // bool | Specifies whether to include blank tasks in the response. Defaults to `true`. (optional)
	status := "status_example" // string | Filters on the status of the task. (optional)
	taskId := []int64{int64(123)} // []int64 | Filters on task ID. Multiple IDs can be specified. (optional)
	spaceId := []int64{int64(123)} // []int64 | Filters on the space ID of the task. Multiple IDs can be specified. (optional)
	pageId := []int64{int64(123)} // []int64 | Filters on the page ID of the task. Multiple IDs can be specified. Note - page and blog post filters can be used in conjunction. (optional)
	blogpostId := []int64{int64(123)} // []int64 | Filters on the blog post ID of the task. Multiple IDs can be specified. Note - page and blog post filters can be used in conjunction. (optional)
	createdBy := []string{"Inner_example"} // []string | Filters on the Account ID of the user who created this task. Multiple IDs can be specified. (optional)
	assignedTo := []string{"Inner_example"} // []string | Filters on the Account ID of the user to whom this task is assigned. Multiple IDs can be specified. (optional)
	completedBy := []string{"Inner_example"} // []string | Filters on the Account ID of the user who completed this task. Multiple IDs can be specified. (optional)
	createdAtFrom := int64(789) // int64 | Filters on start of date-time range of task based on creation date (inclusive). Input is epoch time in milliseconds. (optional)
	createdAtTo := int64(789) // int64 | Filters on end of date-time range of task based on creation date (inclusive). Input is epoch time in milliseconds. (optional)
	dueAtFrom := int64(789) // int64 | Filters on start of date-time range of task based on due date (inclusive). Input is epoch time in milliseconds. (optional)
	dueAtTo := int64(789) // int64 | Filters on end of date-time range of task based on due date (inclusive). Input is epoch time in milliseconds. (optional)
	completedAtFrom := int64(789) // int64 | Filters on start of date-time range of task based on completion date (inclusive). Input is epoch time in milliseconds. (optional)
	completedAtTo := int64(789) // int64 | Filters on end of date-time range of task based on completion date (inclusive). Input is epoch time in milliseconds. (optional)
	cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
	limit := int32(56) // int32 | Maximum number of tasks per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.GetTasks(context.Background()).BodyFormat(bodyFormat).IncludeBlankTasks(includeBlankTasks).Status(status).TaskId(taskId).SpaceId(spaceId).PageId(pageId).BlogpostId(blogpostId).CreatedBy(createdBy).AssignedTo(assignedTo).CompletedBy(completedBy).CreatedAtFrom(createdAtFrom).CreatedAtTo(createdAtTo).DueAtFrom(dueAtFrom).DueAtTo(dueAtTo).CompletedAtFrom(completedAtFrom).CompletedAtTo(completedAtTo).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.GetTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTasks`: MultiEntityResultTask
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.GetTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **includeBlankTasks** | **bool** | Specifies whether to include blank tasks in the response. Defaults to &#x60;true&#x60;. | 
 **status** | **string** | Filters on the status of the task. | 
 **taskId** | **[]int64** | Filters on task ID. Multiple IDs can be specified. | 
 **spaceId** | **[]int64** | Filters on the space ID of the task. Multiple IDs can be specified. | 
 **pageId** | **[]int64** | Filters on the page ID of the task. Multiple IDs can be specified. Note - page and blog post filters can be used in conjunction. | 
 **blogpostId** | **[]int64** | Filters on the blog post ID of the task. Multiple IDs can be specified. Note - page and blog post filters can be used in conjunction. | 
 **createdBy** | **[]string** | Filters on the Account ID of the user who created this task. Multiple IDs can be specified. | 
 **assignedTo** | **[]string** | Filters on the Account ID of the user to whom this task is assigned. Multiple IDs can be specified. | 
 **completedBy** | **[]string** | Filters on the Account ID of the user who completed this task. Multiple IDs can be specified. | 
 **createdAtFrom** | **int64** | Filters on start of date-time range of task based on creation date (inclusive). Input is epoch time in milliseconds. | 
 **createdAtTo** | **int64** | Filters on end of date-time range of task based on creation date (inclusive). Input is epoch time in milliseconds. | 
 **dueAtFrom** | **int64** | Filters on start of date-time range of task based on due date (inclusive). Input is epoch time in milliseconds. | 
 **dueAtTo** | **int64** | Filters on end of date-time range of task based on due date (inclusive). Input is epoch time in milliseconds. | 
 **completedAtFrom** | **int64** | Filters on start of date-time range of task based on completion date (inclusive). Input is epoch time in milliseconds. | 
 **completedAtTo** | **int64** | Filters on end of date-time range of task based on completion date (inclusive). Input is epoch time in milliseconds. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of tasks per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]

### Return type

[**MultiEntityResultTask**](MultiEntityResultTask.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTask

> Task UpdateTask(ctx, id).UpdateTaskRequest(updateTaskRequest).BodyFormat(bodyFormat).Execute()

Update task



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
	id := int64(789) // int64 | The ID of the task to be updated. If you don't know the task ID, use Get tasks and filter the results.
	updateTaskRequest := *openapiclient.NewUpdateTaskRequest("Status_example") // UpdateTaskRequest | 
	bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.UpdateTask(context.Background(), id).UpdateTaskRequest(updateTaskRequest).BodyFormat(bodyFormat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.UpdateTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTask`: Task
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.UpdateTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the task to be updated. If you don&#39;t know the task ID, use Get tasks and filter the results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTaskRequest** | [**UpdateTaskRequest**](UpdateTaskRequest.md) |  | 
 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 

### Return type

[**Task**](Task.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


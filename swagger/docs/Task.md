# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the task. | [optional] 
**LocalId** | Pointer to **string** | Local ID of the task. This ID is local to the corresponding page or blog post. | [optional] 
**SpaceId** | Pointer to **string** | ID of the space the task is in. | [optional] 
**PageId** | Pointer to **string** | ID of the page the task is in. | [optional] 
**BlogPostId** | Pointer to **string** | ID of the blog post the task is in. | [optional] 
**Status** | Pointer to **string** | Status of the task. | [optional] 
**Body** | Pointer to [**TaskBodySingle**](TaskBodySingle.md) |  | [optional] 
**CreatedBy** | Pointer to **string** | Account ID of the user who created this task. | [optional] 
**AssignedTo** | Pointer to **string** | Account ID of the user to whom this task is assigned. | [optional] 
**CompletedBy** | Pointer to **string** | Account ID of the user who completed this task. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time when the task was created. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Date and time when the task was updated. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 
**DueAt** | Pointer to **time.Time** | Date and time when the task is due. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 
**CompletedAt** | Pointer to **time.Time** | Date and time when the task was completed. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 

## Methods

### NewTask

`func NewTask() *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Task) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Task) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Task) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Task) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocalId

`func (o *Task) GetLocalId() string`

GetLocalId returns the LocalId field if non-nil, zero value otherwise.

### GetLocalIdOk

`func (o *Task) GetLocalIdOk() (*string, bool)`

GetLocalIdOk returns a tuple with the LocalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalId

`func (o *Task) SetLocalId(v string)`

SetLocalId sets LocalId field to given value.

### HasLocalId

`func (o *Task) HasLocalId() bool`

HasLocalId returns a boolean if a field has been set.

### GetSpaceId

`func (o *Task) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *Task) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *Task) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *Task) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetPageId

`func (o *Task) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *Task) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *Task) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *Task) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetBlogPostId

`func (o *Task) GetBlogPostId() string`

GetBlogPostId returns the BlogPostId field if non-nil, zero value otherwise.

### GetBlogPostIdOk

`func (o *Task) GetBlogPostIdOk() (*string, bool)`

GetBlogPostIdOk returns a tuple with the BlogPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostId

`func (o *Task) SetBlogPostId(v string)`

SetBlogPostId sets BlogPostId field to given value.

### HasBlogPostId

`func (o *Task) HasBlogPostId() bool`

HasBlogPostId returns a boolean if a field has been set.

### GetStatus

`func (o *Task) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Task) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Task) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Task) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBody

`func (o *Task) GetBody() TaskBodySingle`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Task) GetBodyOk() (*TaskBodySingle, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Task) SetBody(v TaskBodySingle)`

SetBody sets Body field to given value.

### HasBody

`func (o *Task) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Task) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Task) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Task) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Task) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetAssignedTo

`func (o *Task) GetAssignedTo() string`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *Task) GetAssignedToOk() (*string, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *Task) SetAssignedTo(v string)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *Task) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetCompletedBy

`func (o *Task) GetCompletedBy() string`

GetCompletedBy returns the CompletedBy field if non-nil, zero value otherwise.

### GetCompletedByOk

`func (o *Task) GetCompletedByOk() (*string, bool)`

GetCompletedByOk returns a tuple with the CompletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedBy

`func (o *Task) SetCompletedBy(v string)`

SetCompletedBy sets CompletedBy field to given value.

### HasCompletedBy

`func (o *Task) HasCompletedBy() bool`

HasCompletedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Task) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Task) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Task) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Task) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Task) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Task) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Task) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Task) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDueAt

`func (o *Task) GetDueAt() time.Time`

GetDueAt returns the DueAt field if non-nil, zero value otherwise.

### GetDueAtOk

`func (o *Task) GetDueAtOk() (*time.Time, bool)`

GetDueAtOk returns a tuple with the DueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAt

`func (o *Task) SetDueAt(v time.Time)`

SetDueAt sets DueAt field to given value.

### HasDueAt

`func (o *Task) HasDueAt() bool`

HasDueAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *Task) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *Task) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *Task) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *Task) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateCustomContentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of custom content. | 
**Status** | Pointer to **string** | The status of the custom content. Defaults to &#x60;current&#x60; when status not provided. | [optional] 
**SpaceId** | Pointer to **string** | ID of the containing space. | [optional] 
**PageId** | Pointer to **string** | ID of the containing page. | [optional] 
**BlogPostId** | Pointer to **string** | ID of the containing Blog Post. | [optional] 
**CustomContentId** | Pointer to **string** | ID of the containing custom content. | [optional] 
**Title** | **string** | Title of the custom content. | 
**Body** | [**CreateCustomContentRequestBody**](CreateCustomContentRequestBody.md) |  | 

## Methods

### NewCreateCustomContentRequest

`func NewCreateCustomContentRequest(type_ string, title string, body CreateCustomContentRequestBody, ) *CreateCustomContentRequest`

NewCreateCustomContentRequest instantiates a new CreateCustomContentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomContentRequestWithDefaults

`func NewCreateCustomContentRequestWithDefaults() *CreateCustomContentRequest`

NewCreateCustomContentRequestWithDefaults instantiates a new CreateCustomContentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateCustomContentRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateCustomContentRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateCustomContentRequest) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *CreateCustomContentRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateCustomContentRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateCustomContentRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateCustomContentRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpaceId

`func (o *CreateCustomContentRequest) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *CreateCustomContentRequest) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *CreateCustomContentRequest) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *CreateCustomContentRequest) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetPageId

`func (o *CreateCustomContentRequest) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *CreateCustomContentRequest) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *CreateCustomContentRequest) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *CreateCustomContentRequest) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetBlogPostId

`func (o *CreateCustomContentRequest) GetBlogPostId() string`

GetBlogPostId returns the BlogPostId field if non-nil, zero value otherwise.

### GetBlogPostIdOk

`func (o *CreateCustomContentRequest) GetBlogPostIdOk() (*string, bool)`

GetBlogPostIdOk returns a tuple with the BlogPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostId

`func (o *CreateCustomContentRequest) SetBlogPostId(v string)`

SetBlogPostId sets BlogPostId field to given value.

### HasBlogPostId

`func (o *CreateCustomContentRequest) HasBlogPostId() bool`

HasBlogPostId returns a boolean if a field has been set.

### GetCustomContentId

`func (o *CreateCustomContentRequest) GetCustomContentId() string`

GetCustomContentId returns the CustomContentId field if non-nil, zero value otherwise.

### GetCustomContentIdOk

`func (o *CreateCustomContentRequest) GetCustomContentIdOk() (*string, bool)`

GetCustomContentIdOk returns a tuple with the CustomContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContentId

`func (o *CreateCustomContentRequest) SetCustomContentId(v string)`

SetCustomContentId sets CustomContentId field to given value.

### HasCustomContentId

`func (o *CreateCustomContentRequest) HasCustomContentId() bool`

HasCustomContentId returns a boolean if a field has been set.

### GetTitle

`func (o *CreateCustomContentRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateCustomContentRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateCustomContentRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *CreateCustomContentRequest) GetBody() CreateCustomContentRequestBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreateCustomContentRequest) GetBodyOk() (*CreateCustomContentRequestBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreateCustomContentRequest) SetBody(v CreateCustomContentRequestBody)`

SetBody sets Body field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



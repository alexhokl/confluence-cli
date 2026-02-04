# UpdateCustomContentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of custom content. | 
**Type** | **string** | Type of custom content. | 
**Status** | **string** | The status of the custom content. | 
**SpaceId** | Pointer to **string** | ID of the containing space (must be the same as the spaceId of the space the custom content was created in). | [optional] 
**PageId** | Pointer to **string** | ID of the containing page. | [optional] 
**BlogPostId** | Pointer to **string** | ID of the containing Blog Post. | [optional] 
**CustomContentId** | Pointer to **string** | ID of the containing custom content. | [optional] 
**Title** | **string** | Title of the custom content. | 
**Body** | [**CreateCustomContentRequestBody**](CreateCustomContentRequestBody.md) |  | 
**Version** | [**UpdateCustomContentRequestVersion**](UpdateCustomContentRequestVersion.md) |  | 

## Methods

### NewUpdateCustomContentRequest

`func NewUpdateCustomContentRequest(id string, type_ string, status string, title string, body CreateCustomContentRequestBody, version UpdateCustomContentRequestVersion, ) *UpdateCustomContentRequest`

NewUpdateCustomContentRequest instantiates a new UpdateCustomContentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomContentRequestWithDefaults

`func NewUpdateCustomContentRequestWithDefaults() *UpdateCustomContentRequest`

NewUpdateCustomContentRequestWithDefaults instantiates a new UpdateCustomContentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateCustomContentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateCustomContentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateCustomContentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *UpdateCustomContentRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateCustomContentRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateCustomContentRequest) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *UpdateCustomContentRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateCustomContentRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateCustomContentRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSpaceId

`func (o *UpdateCustomContentRequest) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *UpdateCustomContentRequest) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *UpdateCustomContentRequest) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *UpdateCustomContentRequest) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetPageId

`func (o *UpdateCustomContentRequest) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *UpdateCustomContentRequest) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *UpdateCustomContentRequest) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *UpdateCustomContentRequest) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetBlogPostId

`func (o *UpdateCustomContentRequest) GetBlogPostId() string`

GetBlogPostId returns the BlogPostId field if non-nil, zero value otherwise.

### GetBlogPostIdOk

`func (o *UpdateCustomContentRequest) GetBlogPostIdOk() (*string, bool)`

GetBlogPostIdOk returns a tuple with the BlogPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostId

`func (o *UpdateCustomContentRequest) SetBlogPostId(v string)`

SetBlogPostId sets BlogPostId field to given value.

### HasBlogPostId

`func (o *UpdateCustomContentRequest) HasBlogPostId() bool`

HasBlogPostId returns a boolean if a field has been set.

### GetCustomContentId

`func (o *UpdateCustomContentRequest) GetCustomContentId() string`

GetCustomContentId returns the CustomContentId field if non-nil, zero value otherwise.

### GetCustomContentIdOk

`func (o *UpdateCustomContentRequest) GetCustomContentIdOk() (*string, bool)`

GetCustomContentIdOk returns a tuple with the CustomContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContentId

`func (o *UpdateCustomContentRequest) SetCustomContentId(v string)`

SetCustomContentId sets CustomContentId field to given value.

### HasCustomContentId

`func (o *UpdateCustomContentRequest) HasCustomContentId() bool`

HasCustomContentId returns a boolean if a field has been set.

### GetTitle

`func (o *UpdateCustomContentRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateCustomContentRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateCustomContentRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *UpdateCustomContentRequest) GetBody() CreateCustomContentRequestBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *UpdateCustomContentRequest) GetBodyOk() (*CreateCustomContentRequestBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *UpdateCustomContentRequest) SetBody(v CreateCustomContentRequestBody)`

SetBody sets Body field to given value.


### GetVersion

`func (o *UpdateCustomContentRequest) GetVersion() UpdateCustomContentRequestVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateCustomContentRequest) GetVersionOk() (*UpdateCustomContentRequestVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateCustomContentRequest) SetVersion(v UpdateCustomContentRequestVersion)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CustomContentBulk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the custom content. | [optional] 
**Type** | Pointer to **string** | The type of custom content. | [optional] 
**Status** | Pointer to [**ContentStatus**](ContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the custom content. | [optional] 
**SpaceId** | Pointer to **string** | ID of the space the custom content is in.  Note: This is always returned, regardless of if the custom content has a container that is a space. | [optional] 
**PageId** | Pointer to **string** | ID of the containing page.  Note: This is only returned if the custom content has a container that is a page. | [optional] 
**BlogPostId** | Pointer to **string** | ID of the containing blog post.  Note: This is only returned if the custom content has a container that is a blog post. | [optional] 
**CustomContentId** | Pointer to **string** | ID of the containing custom content.  Note: This is only returned if the custom content has a container that is custom content. | [optional] 
**AuthorId** | Pointer to **string** | The account ID of the user who created this custom content originally. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time when the custom content was created. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Body** | Pointer to [**CustomContentBodyBulk**](CustomContentBodyBulk.md) |  | [optional] 
**Links** | Pointer to [**CustomContentLinks**](CustomContentLinks.md) |  | [optional] 

## Methods

### NewCustomContentBulk

`func NewCustomContentBulk() *CustomContentBulk`

NewCustomContentBulk instantiates a new CustomContentBulk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomContentBulkWithDefaults

`func NewCustomContentBulkWithDefaults() *CustomContentBulk`

NewCustomContentBulkWithDefaults instantiates a new CustomContentBulk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomContentBulk) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomContentBulk) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomContentBulk) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomContentBulk) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CustomContentBulk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomContentBulk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomContentBulk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomContentBulk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *CustomContentBulk) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomContentBulk) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomContentBulk) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomContentBulk) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *CustomContentBulk) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CustomContentBulk) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CustomContentBulk) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CustomContentBulk) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSpaceId

`func (o *CustomContentBulk) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *CustomContentBulk) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *CustomContentBulk) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *CustomContentBulk) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetPageId

`func (o *CustomContentBulk) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *CustomContentBulk) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *CustomContentBulk) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *CustomContentBulk) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetBlogPostId

`func (o *CustomContentBulk) GetBlogPostId() string`

GetBlogPostId returns the BlogPostId field if non-nil, zero value otherwise.

### GetBlogPostIdOk

`func (o *CustomContentBulk) GetBlogPostIdOk() (*string, bool)`

GetBlogPostIdOk returns a tuple with the BlogPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostId

`func (o *CustomContentBulk) SetBlogPostId(v string)`

SetBlogPostId sets BlogPostId field to given value.

### HasBlogPostId

`func (o *CustomContentBulk) HasBlogPostId() bool`

HasBlogPostId returns a boolean if a field has been set.

### GetCustomContentId

`func (o *CustomContentBulk) GetCustomContentId() string`

GetCustomContentId returns the CustomContentId field if non-nil, zero value otherwise.

### GetCustomContentIdOk

`func (o *CustomContentBulk) GetCustomContentIdOk() (*string, bool)`

GetCustomContentIdOk returns a tuple with the CustomContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContentId

`func (o *CustomContentBulk) SetCustomContentId(v string)`

SetCustomContentId sets CustomContentId field to given value.

### HasCustomContentId

`func (o *CustomContentBulk) HasCustomContentId() bool`

HasCustomContentId returns a boolean if a field has been set.

### GetAuthorId

`func (o *CustomContentBulk) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *CustomContentBulk) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *CustomContentBulk) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *CustomContentBulk) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CustomContentBulk) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomContentBulk) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomContentBulk) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CustomContentBulk) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *CustomContentBulk) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CustomContentBulk) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CustomContentBulk) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CustomContentBulk) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBody

`func (o *CustomContentBulk) GetBody() CustomContentBodyBulk`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CustomContentBulk) GetBodyOk() (*CustomContentBodyBulk, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CustomContentBulk) SetBody(v CustomContentBodyBulk)`

SetBody sets Body field to given value.

### HasBody

`func (o *CustomContentBulk) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetLinks

`func (o *CustomContentBulk) GetLinks() CustomContentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CustomContentBulk) GetLinksOk() (*CustomContentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CustomContentBulk) SetLinks(v CustomContentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CustomContentBulk) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



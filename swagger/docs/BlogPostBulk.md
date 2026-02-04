# BlogPostBulk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the blog post. | [optional] 
**Status** | Pointer to [**BlogPostContentStatus**](BlogPostContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the blog post. | [optional] 
**SpaceId** | Pointer to **string** | ID of the space the blog post is in. | [optional] 
**AuthorId** | Pointer to **string** | The account ID of the user who created this blog post originally. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time when the blog post was created. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Body** | Pointer to [**BodyBulk**](BodyBulk.md) |  | [optional] 
**Links** | Pointer to [**AbstractPageLinks**](AbstractPageLinks.md) |  | [optional] 

## Methods

### NewBlogPostBulk

`func NewBlogPostBulk() *BlogPostBulk`

NewBlogPostBulk instantiates a new BlogPostBulk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlogPostBulkWithDefaults

`func NewBlogPostBulkWithDefaults() *BlogPostBulk`

NewBlogPostBulkWithDefaults instantiates a new BlogPostBulk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlogPostBulk) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlogPostBulk) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlogPostBulk) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BlogPostBulk) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *BlogPostBulk) GetStatus() BlogPostContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BlogPostBulk) GetStatusOk() (*BlogPostContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BlogPostBulk) SetStatus(v BlogPostContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BlogPostBulk) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *BlogPostBulk) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BlogPostBulk) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BlogPostBulk) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BlogPostBulk) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSpaceId

`func (o *BlogPostBulk) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *BlogPostBulk) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *BlogPostBulk) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *BlogPostBulk) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetAuthorId

`func (o *BlogPostBulk) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *BlogPostBulk) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *BlogPostBulk) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *BlogPostBulk) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BlogPostBulk) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BlogPostBulk) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BlogPostBulk) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BlogPostBulk) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *BlogPostBulk) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BlogPostBulk) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BlogPostBulk) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BlogPostBulk) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBody

`func (o *BlogPostBulk) GetBody() BodyBulk`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *BlogPostBulk) GetBodyOk() (*BodyBulk, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *BlogPostBulk) SetBody(v BodyBulk)`

SetBody sets Body field to given value.

### HasBody

`func (o *BlogPostBulk) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetLinks

`func (o *BlogPostBulk) GetLinks() AbstractPageLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BlogPostBulk) GetLinksOk() (*AbstractPageLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BlogPostBulk) SetLinks(v AbstractPageLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BlogPostBulk) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



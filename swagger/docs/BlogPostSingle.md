# BlogPostSingle

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
**Body** | Pointer to [**BodySingle**](BodySingle.md) |  | [optional] 
**Labels** | Pointer to [**AttachmentSingleLabels**](AttachmentSingleLabels.md) |  | [optional] 
**Properties** | Pointer to [**AttachmentSingleProperties**](AttachmentSingleProperties.md) |  | [optional] 
**Operations** | Pointer to [**AttachmentSingleOperations**](AttachmentSingleOperations.md) |  | [optional] 
**Likes** | Pointer to [**BlogPostSingleLikes**](BlogPostSingleLikes.md) |  | [optional] 
**Versions** | Pointer to [**AttachmentSingleVersions**](AttachmentSingleVersions.md) |  | [optional] 
**IsFavoritedByCurrentUser** | Pointer to **bool** | Whether the blog post has been favorited by the current user. | [optional] 
**Links** | Pointer to [**AbstractPageLinks**](AbstractPageLinks.md) |  | [optional] 

## Methods

### NewBlogPostSingle

`func NewBlogPostSingle() *BlogPostSingle`

NewBlogPostSingle instantiates a new BlogPostSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlogPostSingleWithDefaults

`func NewBlogPostSingleWithDefaults() *BlogPostSingle`

NewBlogPostSingleWithDefaults instantiates a new BlogPostSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlogPostSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlogPostSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlogPostSingle) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BlogPostSingle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *BlogPostSingle) GetStatus() BlogPostContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BlogPostSingle) GetStatusOk() (*BlogPostContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BlogPostSingle) SetStatus(v BlogPostContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BlogPostSingle) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *BlogPostSingle) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BlogPostSingle) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BlogPostSingle) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BlogPostSingle) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSpaceId

`func (o *BlogPostSingle) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *BlogPostSingle) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *BlogPostSingle) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *BlogPostSingle) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetAuthorId

`func (o *BlogPostSingle) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *BlogPostSingle) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *BlogPostSingle) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *BlogPostSingle) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BlogPostSingle) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BlogPostSingle) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BlogPostSingle) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BlogPostSingle) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *BlogPostSingle) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BlogPostSingle) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BlogPostSingle) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BlogPostSingle) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBody

`func (o *BlogPostSingle) GetBody() BodySingle`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *BlogPostSingle) GetBodyOk() (*BodySingle, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *BlogPostSingle) SetBody(v BodySingle)`

SetBody sets Body field to given value.

### HasBody

`func (o *BlogPostSingle) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetLabels

`func (o *BlogPostSingle) GetLabels() AttachmentSingleLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *BlogPostSingle) GetLabelsOk() (*AttachmentSingleLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *BlogPostSingle) SetLabels(v AttachmentSingleLabels)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *BlogPostSingle) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetProperties

`func (o *BlogPostSingle) GetProperties() AttachmentSingleProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BlogPostSingle) GetPropertiesOk() (*AttachmentSingleProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BlogPostSingle) SetProperties(v AttachmentSingleProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BlogPostSingle) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetOperations

`func (o *BlogPostSingle) GetOperations() AttachmentSingleOperations`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *BlogPostSingle) GetOperationsOk() (*AttachmentSingleOperations, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *BlogPostSingle) SetOperations(v AttachmentSingleOperations)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *BlogPostSingle) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetLikes

`func (o *BlogPostSingle) GetLikes() BlogPostSingleLikes`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *BlogPostSingle) GetLikesOk() (*BlogPostSingleLikes, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *BlogPostSingle) SetLikes(v BlogPostSingleLikes)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *BlogPostSingle) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetVersions

`func (o *BlogPostSingle) GetVersions() AttachmentSingleVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *BlogPostSingle) GetVersionsOk() (*AttachmentSingleVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *BlogPostSingle) SetVersions(v AttachmentSingleVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *BlogPostSingle) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetIsFavoritedByCurrentUser

`func (o *BlogPostSingle) GetIsFavoritedByCurrentUser() bool`

GetIsFavoritedByCurrentUser returns the IsFavoritedByCurrentUser field if non-nil, zero value otherwise.

### GetIsFavoritedByCurrentUserOk

`func (o *BlogPostSingle) GetIsFavoritedByCurrentUserOk() (*bool, bool)`

GetIsFavoritedByCurrentUserOk returns a tuple with the IsFavoritedByCurrentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavoritedByCurrentUser

`func (o *BlogPostSingle) SetIsFavoritedByCurrentUser(v bool)`

SetIsFavoritedByCurrentUser sets IsFavoritedByCurrentUser field to given value.

### HasIsFavoritedByCurrentUser

`func (o *BlogPostSingle) HasIsFavoritedByCurrentUser() bool`

HasIsFavoritedByCurrentUser returns a boolean if a field has been set.

### GetLinks

`func (o *BlogPostSingle) GetLinks() AbstractPageLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BlogPostSingle) GetLinksOk() (*AbstractPageLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BlogPostSingle) SetLinks(v AbstractPageLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BlogPostSingle) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InlineCommentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the comment. | [optional] 
**Status** | Pointer to [**ContentStatus**](ContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the comment. | [optional] 
**BlogPostId** | Pointer to **string** | ID of the blog post containing the comment if the comment is on a blog post. | [optional] 
**PageId** | Pointer to **string** | ID of the page containing the comment if the comment is on a page. | [optional] 
**ParentCommentId** | Pointer to **string** | ID of the parent comment if the comment is a reply. | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Body** | Pointer to [**BodySingle**](BodySingle.md) |  | [optional] 
**ResolutionLastModifierId** | Pointer to **string** | Atlassian Account ID of last person who modified the resolve state of the comment. Null until comment is resolved or reopened. | [optional] 
**ResolutionLastModifiedAt** | Pointer to **time.Time** | Timestamp of the last modification to the comment&#39;s resolution status. Null until comment is resolved or reopened. | [optional] 
**ResolutionStatus** | Pointer to [**InlineCommentResolutionStatus**](InlineCommentResolutionStatus.md) |  | [optional] 
**Properties** | Pointer to [**InlineCommentModelProperties**](InlineCommentModelProperties.md) |  | [optional] 
**Operations** | Pointer to [**AttachmentSingleOperations**](AttachmentSingleOperations.md) |  | [optional] 
**Likes** | Pointer to [**BlogPostSingleLikes**](BlogPostSingleLikes.md) |  | [optional] 
**Versions** | Pointer to [**AttachmentSingleVersions**](AttachmentSingleVersions.md) |  | [optional] 
**Links** | Pointer to [**CommentLinks**](CommentLinks.md) |  | [optional] 

## Methods

### NewInlineCommentModel

`func NewInlineCommentModel() *InlineCommentModel`

NewInlineCommentModel instantiates a new InlineCommentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineCommentModelWithDefaults

`func NewInlineCommentModelWithDefaults() *InlineCommentModel`

NewInlineCommentModelWithDefaults instantiates a new InlineCommentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineCommentModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineCommentModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineCommentModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineCommentModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *InlineCommentModel) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineCommentModel) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineCommentModel) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineCommentModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *InlineCommentModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InlineCommentModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InlineCommentModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InlineCommentModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBlogPostId

`func (o *InlineCommentModel) GetBlogPostId() string`

GetBlogPostId returns the BlogPostId field if non-nil, zero value otherwise.

### GetBlogPostIdOk

`func (o *InlineCommentModel) GetBlogPostIdOk() (*string, bool)`

GetBlogPostIdOk returns a tuple with the BlogPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostId

`func (o *InlineCommentModel) SetBlogPostId(v string)`

SetBlogPostId sets BlogPostId field to given value.

### HasBlogPostId

`func (o *InlineCommentModel) HasBlogPostId() bool`

HasBlogPostId returns a boolean if a field has been set.

### GetPageId

`func (o *InlineCommentModel) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *InlineCommentModel) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *InlineCommentModel) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *InlineCommentModel) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetParentCommentId

`func (o *InlineCommentModel) GetParentCommentId() string`

GetParentCommentId returns the ParentCommentId field if non-nil, zero value otherwise.

### GetParentCommentIdOk

`func (o *InlineCommentModel) GetParentCommentIdOk() (*string, bool)`

GetParentCommentIdOk returns a tuple with the ParentCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommentId

`func (o *InlineCommentModel) SetParentCommentId(v string)`

SetParentCommentId sets ParentCommentId field to given value.

### HasParentCommentId

`func (o *InlineCommentModel) HasParentCommentId() bool`

HasParentCommentId returns a boolean if a field has been set.

### GetVersion

`func (o *InlineCommentModel) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InlineCommentModel) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InlineCommentModel) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InlineCommentModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBody

`func (o *InlineCommentModel) GetBody() BodySingle`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *InlineCommentModel) GetBodyOk() (*BodySingle, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *InlineCommentModel) SetBody(v BodySingle)`

SetBody sets Body field to given value.

### HasBody

`func (o *InlineCommentModel) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetResolutionLastModifierId

`func (o *InlineCommentModel) GetResolutionLastModifierId() string`

GetResolutionLastModifierId returns the ResolutionLastModifierId field if non-nil, zero value otherwise.

### GetResolutionLastModifierIdOk

`func (o *InlineCommentModel) GetResolutionLastModifierIdOk() (*string, bool)`

GetResolutionLastModifierIdOk returns a tuple with the ResolutionLastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionLastModifierId

`func (o *InlineCommentModel) SetResolutionLastModifierId(v string)`

SetResolutionLastModifierId sets ResolutionLastModifierId field to given value.

### HasResolutionLastModifierId

`func (o *InlineCommentModel) HasResolutionLastModifierId() bool`

HasResolutionLastModifierId returns a boolean if a field has been set.

### GetResolutionLastModifiedAt

`func (o *InlineCommentModel) GetResolutionLastModifiedAt() time.Time`

GetResolutionLastModifiedAt returns the ResolutionLastModifiedAt field if non-nil, zero value otherwise.

### GetResolutionLastModifiedAtOk

`func (o *InlineCommentModel) GetResolutionLastModifiedAtOk() (*time.Time, bool)`

GetResolutionLastModifiedAtOk returns a tuple with the ResolutionLastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionLastModifiedAt

`func (o *InlineCommentModel) SetResolutionLastModifiedAt(v time.Time)`

SetResolutionLastModifiedAt sets ResolutionLastModifiedAt field to given value.

### HasResolutionLastModifiedAt

`func (o *InlineCommentModel) HasResolutionLastModifiedAt() bool`

HasResolutionLastModifiedAt returns a boolean if a field has been set.

### GetResolutionStatus

`func (o *InlineCommentModel) GetResolutionStatus() InlineCommentResolutionStatus`

GetResolutionStatus returns the ResolutionStatus field if non-nil, zero value otherwise.

### GetResolutionStatusOk

`func (o *InlineCommentModel) GetResolutionStatusOk() (*InlineCommentResolutionStatus, bool)`

GetResolutionStatusOk returns a tuple with the ResolutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionStatus

`func (o *InlineCommentModel) SetResolutionStatus(v InlineCommentResolutionStatus)`

SetResolutionStatus sets ResolutionStatus field to given value.

### HasResolutionStatus

`func (o *InlineCommentModel) HasResolutionStatus() bool`

HasResolutionStatus returns a boolean if a field has been set.

### GetProperties

`func (o *InlineCommentModel) GetProperties() InlineCommentModelProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *InlineCommentModel) GetPropertiesOk() (*InlineCommentModelProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *InlineCommentModel) SetProperties(v InlineCommentModelProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *InlineCommentModel) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetOperations

`func (o *InlineCommentModel) GetOperations() AttachmentSingleOperations`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *InlineCommentModel) GetOperationsOk() (*AttachmentSingleOperations, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *InlineCommentModel) SetOperations(v AttachmentSingleOperations)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *InlineCommentModel) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetLikes

`func (o *InlineCommentModel) GetLikes() BlogPostSingleLikes`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *InlineCommentModel) GetLikesOk() (*BlogPostSingleLikes, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *InlineCommentModel) SetLikes(v BlogPostSingleLikes)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *InlineCommentModel) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetVersions

`func (o *InlineCommentModel) GetVersions() AttachmentSingleVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *InlineCommentModel) GetVersionsOk() (*AttachmentSingleVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *InlineCommentModel) SetVersions(v AttachmentSingleVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *InlineCommentModel) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetLinks

`func (o *InlineCommentModel) GetLinks() CommentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InlineCommentModel) GetLinksOk() (*CommentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InlineCommentModel) SetLinks(v CommentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *InlineCommentModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



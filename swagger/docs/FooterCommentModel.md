# FooterCommentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the comment. | [optional] 
**Status** | Pointer to [**ContentStatus**](ContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the comment. | [optional] 
**BlogPostId** | Pointer to **string** | ID of the blog post containing the comment if the comment is on a blog post. | [optional] 
**PageId** | Pointer to **string** | ID of the page containing the comment if the comment is on a page. | [optional] 
**AttachmentId** | Pointer to **string** | ID of the attachment containing the comment if the comment is on an attachment. | [optional] 
**CustomContentId** | Pointer to **string** | ID of the custom content containing the comment if the comment is on a custom content. | [optional] 
**ParentCommentId** | Pointer to **string** | ID of the parent comment if the comment is a reply. | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Properties** | Pointer to [**AttachmentSingleProperties**](AttachmentSingleProperties.md) |  | [optional] 
**Operations** | Pointer to [**AttachmentSingleOperations**](AttachmentSingleOperations.md) |  | [optional] 
**Likes** | Pointer to [**BlogPostSingleLikes**](BlogPostSingleLikes.md) |  | [optional] 
**Versions** | Pointer to [**AttachmentSingleVersions**](AttachmentSingleVersions.md) |  | [optional] 
**Body** | Pointer to [**BodySingle**](BodySingle.md) |  | [optional] 
**Links** | Pointer to [**CommentLinks**](CommentLinks.md) |  | [optional] 

## Methods

### NewFooterCommentModel

`func NewFooterCommentModel() *FooterCommentModel`

NewFooterCommentModel instantiates a new FooterCommentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFooterCommentModelWithDefaults

`func NewFooterCommentModelWithDefaults() *FooterCommentModel`

NewFooterCommentModelWithDefaults instantiates a new FooterCommentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FooterCommentModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FooterCommentModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FooterCommentModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FooterCommentModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *FooterCommentModel) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FooterCommentModel) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FooterCommentModel) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FooterCommentModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *FooterCommentModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FooterCommentModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FooterCommentModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FooterCommentModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBlogPostId

`func (o *FooterCommentModel) GetBlogPostId() string`

GetBlogPostId returns the BlogPostId field if non-nil, zero value otherwise.

### GetBlogPostIdOk

`func (o *FooterCommentModel) GetBlogPostIdOk() (*string, bool)`

GetBlogPostIdOk returns a tuple with the BlogPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostId

`func (o *FooterCommentModel) SetBlogPostId(v string)`

SetBlogPostId sets BlogPostId field to given value.

### HasBlogPostId

`func (o *FooterCommentModel) HasBlogPostId() bool`

HasBlogPostId returns a boolean if a field has been set.

### GetPageId

`func (o *FooterCommentModel) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *FooterCommentModel) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *FooterCommentModel) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *FooterCommentModel) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetAttachmentId

`func (o *FooterCommentModel) GetAttachmentId() string`

GetAttachmentId returns the AttachmentId field if non-nil, zero value otherwise.

### GetAttachmentIdOk

`func (o *FooterCommentModel) GetAttachmentIdOk() (*string, bool)`

GetAttachmentIdOk returns a tuple with the AttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentId

`func (o *FooterCommentModel) SetAttachmentId(v string)`

SetAttachmentId sets AttachmentId field to given value.

### HasAttachmentId

`func (o *FooterCommentModel) HasAttachmentId() bool`

HasAttachmentId returns a boolean if a field has been set.

### GetCustomContentId

`func (o *FooterCommentModel) GetCustomContentId() string`

GetCustomContentId returns the CustomContentId field if non-nil, zero value otherwise.

### GetCustomContentIdOk

`func (o *FooterCommentModel) GetCustomContentIdOk() (*string, bool)`

GetCustomContentIdOk returns a tuple with the CustomContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContentId

`func (o *FooterCommentModel) SetCustomContentId(v string)`

SetCustomContentId sets CustomContentId field to given value.

### HasCustomContentId

`func (o *FooterCommentModel) HasCustomContentId() bool`

HasCustomContentId returns a boolean if a field has been set.

### GetParentCommentId

`func (o *FooterCommentModel) GetParentCommentId() string`

GetParentCommentId returns the ParentCommentId field if non-nil, zero value otherwise.

### GetParentCommentIdOk

`func (o *FooterCommentModel) GetParentCommentIdOk() (*string, bool)`

GetParentCommentIdOk returns a tuple with the ParentCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommentId

`func (o *FooterCommentModel) SetParentCommentId(v string)`

SetParentCommentId sets ParentCommentId field to given value.

### HasParentCommentId

`func (o *FooterCommentModel) HasParentCommentId() bool`

HasParentCommentId returns a boolean if a field has been set.

### GetVersion

`func (o *FooterCommentModel) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FooterCommentModel) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FooterCommentModel) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FooterCommentModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetProperties

`func (o *FooterCommentModel) GetProperties() AttachmentSingleProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *FooterCommentModel) GetPropertiesOk() (*AttachmentSingleProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *FooterCommentModel) SetProperties(v AttachmentSingleProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *FooterCommentModel) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetOperations

`func (o *FooterCommentModel) GetOperations() AttachmentSingleOperations`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *FooterCommentModel) GetOperationsOk() (*AttachmentSingleOperations, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *FooterCommentModel) SetOperations(v AttachmentSingleOperations)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *FooterCommentModel) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetLikes

`func (o *FooterCommentModel) GetLikes() BlogPostSingleLikes`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *FooterCommentModel) GetLikesOk() (*BlogPostSingleLikes, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *FooterCommentModel) SetLikes(v BlogPostSingleLikes)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *FooterCommentModel) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetVersions

`func (o *FooterCommentModel) GetVersions() AttachmentSingleVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *FooterCommentModel) GetVersionsOk() (*AttachmentSingleVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *FooterCommentModel) SetVersions(v AttachmentSingleVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *FooterCommentModel) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetBody

`func (o *FooterCommentModel) GetBody() BodySingle`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *FooterCommentModel) GetBodyOk() (*BodySingle, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *FooterCommentModel) SetBody(v BodySingle)`

SetBody sets Body field to given value.

### HasBody

`func (o *FooterCommentModel) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetLinks

`func (o *FooterCommentModel) GetLinks() CommentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FooterCommentModel) GetLinksOk() (*CommentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FooterCommentModel) SetLinks(v CommentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FooterCommentModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



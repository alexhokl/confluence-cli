# CreateInlineComment201Response

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
**Links** | Pointer to [**GetAttachmentById200ResponseAllOfLinks**](GetAttachmentById200ResponseAllOfLinks.md) |  | [optional] 

## Methods

### NewCreateInlineComment201Response

`func NewCreateInlineComment201Response() *CreateInlineComment201Response`

NewCreateInlineComment201Response instantiates a new CreateInlineComment201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInlineComment201ResponseWithDefaults

`func NewCreateInlineComment201ResponseWithDefaults() *CreateInlineComment201Response`

NewCreateInlineComment201ResponseWithDefaults instantiates a new CreateInlineComment201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateInlineComment201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateInlineComment201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateInlineComment201Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateInlineComment201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *CreateInlineComment201Response) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateInlineComment201Response) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateInlineComment201Response) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateInlineComment201Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *CreateInlineComment201Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateInlineComment201Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateInlineComment201Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateInlineComment201Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBlogPostId

`func (o *CreateInlineComment201Response) GetBlogPostId() string`

GetBlogPostId returns the BlogPostId field if non-nil, zero value otherwise.

### GetBlogPostIdOk

`func (o *CreateInlineComment201Response) GetBlogPostIdOk() (*string, bool)`

GetBlogPostIdOk returns a tuple with the BlogPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostId

`func (o *CreateInlineComment201Response) SetBlogPostId(v string)`

SetBlogPostId sets BlogPostId field to given value.

### HasBlogPostId

`func (o *CreateInlineComment201Response) HasBlogPostId() bool`

HasBlogPostId returns a boolean if a field has been set.

### GetPageId

`func (o *CreateInlineComment201Response) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *CreateInlineComment201Response) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *CreateInlineComment201Response) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *CreateInlineComment201Response) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetParentCommentId

`func (o *CreateInlineComment201Response) GetParentCommentId() string`

GetParentCommentId returns the ParentCommentId field if non-nil, zero value otherwise.

### GetParentCommentIdOk

`func (o *CreateInlineComment201Response) GetParentCommentIdOk() (*string, bool)`

GetParentCommentIdOk returns a tuple with the ParentCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommentId

`func (o *CreateInlineComment201Response) SetParentCommentId(v string)`

SetParentCommentId sets ParentCommentId field to given value.

### HasParentCommentId

`func (o *CreateInlineComment201Response) HasParentCommentId() bool`

HasParentCommentId returns a boolean if a field has been set.

### GetVersion

`func (o *CreateInlineComment201Response) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateInlineComment201Response) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateInlineComment201Response) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreateInlineComment201Response) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBody

`func (o *CreateInlineComment201Response) GetBody() BodySingle`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreateInlineComment201Response) GetBodyOk() (*BodySingle, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreateInlineComment201Response) SetBody(v BodySingle)`

SetBody sets Body field to given value.

### HasBody

`func (o *CreateInlineComment201Response) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetResolutionLastModifierId

`func (o *CreateInlineComment201Response) GetResolutionLastModifierId() string`

GetResolutionLastModifierId returns the ResolutionLastModifierId field if non-nil, zero value otherwise.

### GetResolutionLastModifierIdOk

`func (o *CreateInlineComment201Response) GetResolutionLastModifierIdOk() (*string, bool)`

GetResolutionLastModifierIdOk returns a tuple with the ResolutionLastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionLastModifierId

`func (o *CreateInlineComment201Response) SetResolutionLastModifierId(v string)`

SetResolutionLastModifierId sets ResolutionLastModifierId field to given value.

### HasResolutionLastModifierId

`func (o *CreateInlineComment201Response) HasResolutionLastModifierId() bool`

HasResolutionLastModifierId returns a boolean if a field has been set.

### GetResolutionLastModifiedAt

`func (o *CreateInlineComment201Response) GetResolutionLastModifiedAt() time.Time`

GetResolutionLastModifiedAt returns the ResolutionLastModifiedAt field if non-nil, zero value otherwise.

### GetResolutionLastModifiedAtOk

`func (o *CreateInlineComment201Response) GetResolutionLastModifiedAtOk() (*time.Time, bool)`

GetResolutionLastModifiedAtOk returns a tuple with the ResolutionLastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionLastModifiedAt

`func (o *CreateInlineComment201Response) SetResolutionLastModifiedAt(v time.Time)`

SetResolutionLastModifiedAt sets ResolutionLastModifiedAt field to given value.

### HasResolutionLastModifiedAt

`func (o *CreateInlineComment201Response) HasResolutionLastModifiedAt() bool`

HasResolutionLastModifiedAt returns a boolean if a field has been set.

### GetResolutionStatus

`func (o *CreateInlineComment201Response) GetResolutionStatus() InlineCommentResolutionStatus`

GetResolutionStatus returns the ResolutionStatus field if non-nil, zero value otherwise.

### GetResolutionStatusOk

`func (o *CreateInlineComment201Response) GetResolutionStatusOk() (*InlineCommentResolutionStatus, bool)`

GetResolutionStatusOk returns a tuple with the ResolutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionStatus

`func (o *CreateInlineComment201Response) SetResolutionStatus(v InlineCommentResolutionStatus)`

SetResolutionStatus sets ResolutionStatus field to given value.

### HasResolutionStatus

`func (o *CreateInlineComment201Response) HasResolutionStatus() bool`

HasResolutionStatus returns a boolean if a field has been set.

### GetProperties

`func (o *CreateInlineComment201Response) GetProperties() InlineCommentModelProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CreateInlineComment201Response) GetPropertiesOk() (*InlineCommentModelProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CreateInlineComment201Response) SetProperties(v InlineCommentModelProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CreateInlineComment201Response) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetOperations

`func (o *CreateInlineComment201Response) GetOperations() AttachmentSingleOperations`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *CreateInlineComment201Response) GetOperationsOk() (*AttachmentSingleOperations, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *CreateInlineComment201Response) SetOperations(v AttachmentSingleOperations)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *CreateInlineComment201Response) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetLikes

`func (o *CreateInlineComment201Response) GetLikes() BlogPostSingleLikes`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *CreateInlineComment201Response) GetLikesOk() (*BlogPostSingleLikes, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *CreateInlineComment201Response) SetLikes(v BlogPostSingleLikes)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *CreateInlineComment201Response) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetVersions

`func (o *CreateInlineComment201Response) GetVersions() AttachmentSingleVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *CreateInlineComment201Response) GetVersionsOk() (*AttachmentSingleVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *CreateInlineComment201Response) SetVersions(v AttachmentSingleVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *CreateInlineComment201Response) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetLinks

`func (o *CreateInlineComment201Response) GetLinks() GetAttachmentById200ResponseAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateInlineComment201Response) GetLinksOk() (*GetAttachmentById200ResponseAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateInlineComment201Response) SetLinks(v GetAttachmentById200ResponseAllOfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateInlineComment201Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



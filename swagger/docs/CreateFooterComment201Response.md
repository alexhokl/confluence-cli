# CreateFooterComment201Response

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
**Links** | Pointer to [**GetAttachmentById200ResponseAllOfLinks**](GetAttachmentById200ResponseAllOfLinks.md) |  | [optional] 

## Methods

### NewCreateFooterComment201Response

`func NewCreateFooterComment201Response() *CreateFooterComment201Response`

NewCreateFooterComment201Response instantiates a new CreateFooterComment201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFooterComment201ResponseWithDefaults

`func NewCreateFooterComment201ResponseWithDefaults() *CreateFooterComment201Response`

NewCreateFooterComment201ResponseWithDefaults instantiates a new CreateFooterComment201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateFooterComment201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateFooterComment201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateFooterComment201Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateFooterComment201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *CreateFooterComment201Response) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateFooterComment201Response) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateFooterComment201Response) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateFooterComment201Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *CreateFooterComment201Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateFooterComment201Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateFooterComment201Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateFooterComment201Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBlogPostId

`func (o *CreateFooterComment201Response) GetBlogPostId() string`

GetBlogPostId returns the BlogPostId field if non-nil, zero value otherwise.

### GetBlogPostIdOk

`func (o *CreateFooterComment201Response) GetBlogPostIdOk() (*string, bool)`

GetBlogPostIdOk returns a tuple with the BlogPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostId

`func (o *CreateFooterComment201Response) SetBlogPostId(v string)`

SetBlogPostId sets BlogPostId field to given value.

### HasBlogPostId

`func (o *CreateFooterComment201Response) HasBlogPostId() bool`

HasBlogPostId returns a boolean if a field has been set.

### GetPageId

`func (o *CreateFooterComment201Response) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *CreateFooterComment201Response) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *CreateFooterComment201Response) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *CreateFooterComment201Response) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetAttachmentId

`func (o *CreateFooterComment201Response) GetAttachmentId() string`

GetAttachmentId returns the AttachmentId field if non-nil, zero value otherwise.

### GetAttachmentIdOk

`func (o *CreateFooterComment201Response) GetAttachmentIdOk() (*string, bool)`

GetAttachmentIdOk returns a tuple with the AttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentId

`func (o *CreateFooterComment201Response) SetAttachmentId(v string)`

SetAttachmentId sets AttachmentId field to given value.

### HasAttachmentId

`func (o *CreateFooterComment201Response) HasAttachmentId() bool`

HasAttachmentId returns a boolean if a field has been set.

### GetCustomContentId

`func (o *CreateFooterComment201Response) GetCustomContentId() string`

GetCustomContentId returns the CustomContentId field if non-nil, zero value otherwise.

### GetCustomContentIdOk

`func (o *CreateFooterComment201Response) GetCustomContentIdOk() (*string, bool)`

GetCustomContentIdOk returns a tuple with the CustomContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContentId

`func (o *CreateFooterComment201Response) SetCustomContentId(v string)`

SetCustomContentId sets CustomContentId field to given value.

### HasCustomContentId

`func (o *CreateFooterComment201Response) HasCustomContentId() bool`

HasCustomContentId returns a boolean if a field has been set.

### GetParentCommentId

`func (o *CreateFooterComment201Response) GetParentCommentId() string`

GetParentCommentId returns the ParentCommentId field if non-nil, zero value otherwise.

### GetParentCommentIdOk

`func (o *CreateFooterComment201Response) GetParentCommentIdOk() (*string, bool)`

GetParentCommentIdOk returns a tuple with the ParentCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommentId

`func (o *CreateFooterComment201Response) SetParentCommentId(v string)`

SetParentCommentId sets ParentCommentId field to given value.

### HasParentCommentId

`func (o *CreateFooterComment201Response) HasParentCommentId() bool`

HasParentCommentId returns a boolean if a field has been set.

### GetVersion

`func (o *CreateFooterComment201Response) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateFooterComment201Response) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateFooterComment201Response) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreateFooterComment201Response) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetProperties

`func (o *CreateFooterComment201Response) GetProperties() AttachmentSingleProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CreateFooterComment201Response) GetPropertiesOk() (*AttachmentSingleProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CreateFooterComment201Response) SetProperties(v AttachmentSingleProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CreateFooterComment201Response) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetOperations

`func (o *CreateFooterComment201Response) GetOperations() AttachmentSingleOperations`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *CreateFooterComment201Response) GetOperationsOk() (*AttachmentSingleOperations, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *CreateFooterComment201Response) SetOperations(v AttachmentSingleOperations)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *CreateFooterComment201Response) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetLikes

`func (o *CreateFooterComment201Response) GetLikes() BlogPostSingleLikes`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *CreateFooterComment201Response) GetLikesOk() (*BlogPostSingleLikes, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *CreateFooterComment201Response) SetLikes(v BlogPostSingleLikes)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *CreateFooterComment201Response) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetVersions

`func (o *CreateFooterComment201Response) GetVersions() AttachmentSingleVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *CreateFooterComment201Response) GetVersionsOk() (*AttachmentSingleVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *CreateFooterComment201Response) SetVersions(v AttachmentSingleVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *CreateFooterComment201Response) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetBody

`func (o *CreateFooterComment201Response) GetBody() BodySingle`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreateFooterComment201Response) GetBodyOk() (*BodySingle, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreateFooterComment201Response) SetBody(v BodySingle)`

SetBody sets Body field to given value.

### HasBody

`func (o *CreateFooterComment201Response) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetLinks

`func (o *CreateFooterComment201Response) GetLinks() GetAttachmentById200ResponseAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateFooterComment201Response) GetLinksOk() (*GetAttachmentById200ResponseAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateFooterComment201Response) SetLinks(v GetAttachmentById200ResponseAllOfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateFooterComment201Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



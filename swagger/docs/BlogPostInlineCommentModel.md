# BlogPostInlineCommentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the comment. | [optional] 
**Status** | Pointer to [**ContentStatus**](ContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the comment. | [optional] 
**BlogPostId** | Pointer to **string** | ID of the blog post the comment is in. | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Body** | Pointer to [**BodyBulk**](BodyBulk.md) |  | [optional] 
**ResolutionStatus** | Pointer to [**InlineCommentResolutionStatus**](InlineCommentResolutionStatus.md) |  | [optional] 
**Properties** | Pointer to [**InlineCommentProperties**](InlineCommentProperties.md) |  | [optional] 
**Links** | Pointer to [**CommentLinks**](CommentLinks.md) |  | [optional] 

## Methods

### NewBlogPostInlineCommentModel

`func NewBlogPostInlineCommentModel() *BlogPostInlineCommentModel`

NewBlogPostInlineCommentModel instantiates a new BlogPostInlineCommentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlogPostInlineCommentModelWithDefaults

`func NewBlogPostInlineCommentModelWithDefaults() *BlogPostInlineCommentModel`

NewBlogPostInlineCommentModelWithDefaults instantiates a new BlogPostInlineCommentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlogPostInlineCommentModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlogPostInlineCommentModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlogPostInlineCommentModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BlogPostInlineCommentModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *BlogPostInlineCommentModel) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BlogPostInlineCommentModel) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BlogPostInlineCommentModel) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BlogPostInlineCommentModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *BlogPostInlineCommentModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BlogPostInlineCommentModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BlogPostInlineCommentModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BlogPostInlineCommentModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBlogPostId

`func (o *BlogPostInlineCommentModel) GetBlogPostId() string`

GetBlogPostId returns the BlogPostId field if non-nil, zero value otherwise.

### GetBlogPostIdOk

`func (o *BlogPostInlineCommentModel) GetBlogPostIdOk() (*string, bool)`

GetBlogPostIdOk returns a tuple with the BlogPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostId

`func (o *BlogPostInlineCommentModel) SetBlogPostId(v string)`

SetBlogPostId sets BlogPostId field to given value.

### HasBlogPostId

`func (o *BlogPostInlineCommentModel) HasBlogPostId() bool`

HasBlogPostId returns a boolean if a field has been set.

### GetVersion

`func (o *BlogPostInlineCommentModel) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BlogPostInlineCommentModel) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BlogPostInlineCommentModel) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BlogPostInlineCommentModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBody

`func (o *BlogPostInlineCommentModel) GetBody() BodyBulk`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *BlogPostInlineCommentModel) GetBodyOk() (*BodyBulk, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *BlogPostInlineCommentModel) SetBody(v BodyBulk)`

SetBody sets Body field to given value.

### HasBody

`func (o *BlogPostInlineCommentModel) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetResolutionStatus

`func (o *BlogPostInlineCommentModel) GetResolutionStatus() InlineCommentResolutionStatus`

GetResolutionStatus returns the ResolutionStatus field if non-nil, zero value otherwise.

### GetResolutionStatusOk

`func (o *BlogPostInlineCommentModel) GetResolutionStatusOk() (*InlineCommentResolutionStatus, bool)`

GetResolutionStatusOk returns a tuple with the ResolutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionStatus

`func (o *BlogPostInlineCommentModel) SetResolutionStatus(v InlineCommentResolutionStatus)`

SetResolutionStatus sets ResolutionStatus field to given value.

### HasResolutionStatus

`func (o *BlogPostInlineCommentModel) HasResolutionStatus() bool`

HasResolutionStatus returns a boolean if a field has been set.

### GetProperties

`func (o *BlogPostInlineCommentModel) GetProperties() InlineCommentProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BlogPostInlineCommentModel) GetPropertiesOk() (*InlineCommentProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BlogPostInlineCommentModel) SetProperties(v InlineCommentProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BlogPostInlineCommentModel) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetLinks

`func (o *BlogPostInlineCommentModel) GetLinks() CommentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BlogPostInlineCommentModel) GetLinksOk() (*CommentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BlogPostInlineCommentModel) SetLinks(v CommentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BlogPostInlineCommentModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateFooterCommentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlogPostId** | Pointer to **string** | ID of the containing blog post, if intending to create a top level footer comment. Do not provide if creating a reply. | [optional] 
**PageId** | Pointer to **string** | ID of the containing page, if intending to create a top level footer comment. Do not provide if creating a reply. | [optional] 
**ParentCommentId** | Pointer to **string** | ID of the parent comment, if intending to create a reply. Do not provide if creating a top level comment. | [optional] 
**AttachmentId** | Pointer to **string** | ID of the attachment, if intending to create a comment against an attachment. | [optional] 
**CustomContentId** | Pointer to **string** | ID of the custom content, if intending to create a comment against a custom content. | [optional] 
**Body** | Pointer to [**CreateFooterCommentModelBody**](CreateFooterCommentModelBody.md) |  | [optional] 

## Methods

### NewCreateFooterCommentModel

`func NewCreateFooterCommentModel() *CreateFooterCommentModel`

NewCreateFooterCommentModel instantiates a new CreateFooterCommentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFooterCommentModelWithDefaults

`func NewCreateFooterCommentModelWithDefaults() *CreateFooterCommentModel`

NewCreateFooterCommentModelWithDefaults instantiates a new CreateFooterCommentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlogPostId

`func (o *CreateFooterCommentModel) GetBlogPostId() string`

GetBlogPostId returns the BlogPostId field if non-nil, zero value otherwise.

### GetBlogPostIdOk

`func (o *CreateFooterCommentModel) GetBlogPostIdOk() (*string, bool)`

GetBlogPostIdOk returns a tuple with the BlogPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostId

`func (o *CreateFooterCommentModel) SetBlogPostId(v string)`

SetBlogPostId sets BlogPostId field to given value.

### HasBlogPostId

`func (o *CreateFooterCommentModel) HasBlogPostId() bool`

HasBlogPostId returns a boolean if a field has been set.

### GetPageId

`func (o *CreateFooterCommentModel) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *CreateFooterCommentModel) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *CreateFooterCommentModel) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *CreateFooterCommentModel) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetParentCommentId

`func (o *CreateFooterCommentModel) GetParentCommentId() string`

GetParentCommentId returns the ParentCommentId field if non-nil, zero value otherwise.

### GetParentCommentIdOk

`func (o *CreateFooterCommentModel) GetParentCommentIdOk() (*string, bool)`

GetParentCommentIdOk returns a tuple with the ParentCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommentId

`func (o *CreateFooterCommentModel) SetParentCommentId(v string)`

SetParentCommentId sets ParentCommentId field to given value.

### HasParentCommentId

`func (o *CreateFooterCommentModel) HasParentCommentId() bool`

HasParentCommentId returns a boolean if a field has been set.

### GetAttachmentId

`func (o *CreateFooterCommentModel) GetAttachmentId() string`

GetAttachmentId returns the AttachmentId field if non-nil, zero value otherwise.

### GetAttachmentIdOk

`func (o *CreateFooterCommentModel) GetAttachmentIdOk() (*string, bool)`

GetAttachmentIdOk returns a tuple with the AttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentId

`func (o *CreateFooterCommentModel) SetAttachmentId(v string)`

SetAttachmentId sets AttachmentId field to given value.

### HasAttachmentId

`func (o *CreateFooterCommentModel) HasAttachmentId() bool`

HasAttachmentId returns a boolean if a field has been set.

### GetCustomContentId

`func (o *CreateFooterCommentModel) GetCustomContentId() string`

GetCustomContentId returns the CustomContentId field if non-nil, zero value otherwise.

### GetCustomContentIdOk

`func (o *CreateFooterCommentModel) GetCustomContentIdOk() (*string, bool)`

GetCustomContentIdOk returns a tuple with the CustomContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContentId

`func (o *CreateFooterCommentModel) SetCustomContentId(v string)`

SetCustomContentId sets CustomContentId field to given value.

### HasCustomContentId

`func (o *CreateFooterCommentModel) HasCustomContentId() bool`

HasCustomContentId returns a boolean if a field has been set.

### GetBody

`func (o *CreateFooterCommentModel) GetBody() CreateFooterCommentModelBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreateFooterCommentModel) GetBodyOk() (*CreateFooterCommentModelBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreateFooterCommentModel) SetBody(v CreateFooterCommentModelBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *CreateFooterCommentModel) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateInlineCommentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlogPostId** | Pointer to **string** | ID of the containing blog post, if intending to create a top level footer comment. Do not provide if creating a reply. | [optional] 
**PageId** | Pointer to **string** | ID of the containing page, if intending to create a top level footer comment. Do not provide if creating a reply. | [optional] 
**ParentCommentId** | Pointer to **string** | ID of the parent comment, if intending to create a reply. Do not provide if creating a top level comment. | [optional] 
**Body** | Pointer to [**CreateFooterCommentModelBody**](CreateFooterCommentModelBody.md) |  | [optional] 
**InlineCommentProperties** | Pointer to [**CreateInlineCommentModelInlineCommentProperties**](CreateInlineCommentModelInlineCommentProperties.md) |  | [optional] 

## Methods

### NewCreateInlineCommentModel

`func NewCreateInlineCommentModel() *CreateInlineCommentModel`

NewCreateInlineCommentModel instantiates a new CreateInlineCommentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInlineCommentModelWithDefaults

`func NewCreateInlineCommentModelWithDefaults() *CreateInlineCommentModel`

NewCreateInlineCommentModelWithDefaults instantiates a new CreateInlineCommentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlogPostId

`func (o *CreateInlineCommentModel) GetBlogPostId() string`

GetBlogPostId returns the BlogPostId field if non-nil, zero value otherwise.

### GetBlogPostIdOk

`func (o *CreateInlineCommentModel) GetBlogPostIdOk() (*string, bool)`

GetBlogPostIdOk returns a tuple with the BlogPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostId

`func (o *CreateInlineCommentModel) SetBlogPostId(v string)`

SetBlogPostId sets BlogPostId field to given value.

### HasBlogPostId

`func (o *CreateInlineCommentModel) HasBlogPostId() bool`

HasBlogPostId returns a boolean if a field has been set.

### GetPageId

`func (o *CreateInlineCommentModel) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *CreateInlineCommentModel) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *CreateInlineCommentModel) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *CreateInlineCommentModel) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetParentCommentId

`func (o *CreateInlineCommentModel) GetParentCommentId() string`

GetParentCommentId returns the ParentCommentId field if non-nil, zero value otherwise.

### GetParentCommentIdOk

`func (o *CreateInlineCommentModel) GetParentCommentIdOk() (*string, bool)`

GetParentCommentIdOk returns a tuple with the ParentCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommentId

`func (o *CreateInlineCommentModel) SetParentCommentId(v string)`

SetParentCommentId sets ParentCommentId field to given value.

### HasParentCommentId

`func (o *CreateInlineCommentModel) HasParentCommentId() bool`

HasParentCommentId returns a boolean if a field has been set.

### GetBody

`func (o *CreateInlineCommentModel) GetBody() CreateFooterCommentModelBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreateInlineCommentModel) GetBodyOk() (*CreateFooterCommentModelBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreateInlineCommentModel) SetBody(v CreateFooterCommentModelBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *CreateInlineCommentModel) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetInlineCommentProperties

`func (o *CreateInlineCommentModel) GetInlineCommentProperties() CreateInlineCommentModelInlineCommentProperties`

GetInlineCommentProperties returns the InlineCommentProperties field if non-nil, zero value otherwise.

### GetInlineCommentPropertiesOk

`func (o *CreateInlineCommentModel) GetInlineCommentPropertiesOk() (*CreateInlineCommentModelInlineCommentProperties, bool)`

GetInlineCommentPropertiesOk returns a tuple with the InlineCommentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineCommentProperties

`func (o *CreateInlineCommentModel) SetInlineCommentProperties(v CreateInlineCommentModelInlineCommentProperties)`

SetInlineCommentProperties sets InlineCommentProperties field to given value.

### HasInlineCommentProperties

`func (o *CreateInlineCommentModel) HasInlineCommentProperties() bool`

HasInlineCommentProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



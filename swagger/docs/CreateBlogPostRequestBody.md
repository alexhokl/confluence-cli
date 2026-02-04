# CreateBlogPostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Representation** | Pointer to **string** | Type of content representation used for the value field. | [optional] 
**Value** | Pointer to **string** | Body of the blog post, in the format found in the representation field. | [optional] 
**Storage** | Pointer to [**BlogPostBodyWrite**](BlogPostBodyWrite.md) |  | [optional] 
**AtlasDocFormat** | Pointer to [**BlogPostBodyWrite**](BlogPostBodyWrite.md) |  | [optional] 
**Wiki** | Pointer to [**BlogPostBodyWrite**](BlogPostBodyWrite.md) |  | [optional] 

## Methods

### NewCreateBlogPostRequestBody

`func NewCreateBlogPostRequestBody() *CreateBlogPostRequestBody`

NewCreateBlogPostRequestBody instantiates a new CreateBlogPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBlogPostRequestBodyWithDefaults

`func NewCreateBlogPostRequestBodyWithDefaults() *CreateBlogPostRequestBody`

NewCreateBlogPostRequestBodyWithDefaults instantiates a new CreateBlogPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepresentation

`func (o *CreateBlogPostRequestBody) GetRepresentation() string`

GetRepresentation returns the Representation field if non-nil, zero value otherwise.

### GetRepresentationOk

`func (o *CreateBlogPostRequestBody) GetRepresentationOk() (*string, bool)`

GetRepresentationOk returns a tuple with the Representation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepresentation

`func (o *CreateBlogPostRequestBody) SetRepresentation(v string)`

SetRepresentation sets Representation field to given value.

### HasRepresentation

`func (o *CreateBlogPostRequestBody) HasRepresentation() bool`

HasRepresentation returns a boolean if a field has been set.

### GetValue

`func (o *CreateBlogPostRequestBody) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateBlogPostRequestBody) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateBlogPostRequestBody) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CreateBlogPostRequestBody) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetStorage

`func (o *CreateBlogPostRequestBody) GetStorage() BlogPostBodyWrite`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CreateBlogPostRequestBody) GetStorageOk() (*BlogPostBodyWrite, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CreateBlogPostRequestBody) SetStorage(v BlogPostBodyWrite)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *CreateBlogPostRequestBody) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetAtlasDocFormat

`func (o *CreateBlogPostRequestBody) GetAtlasDocFormat() BlogPostBodyWrite`

GetAtlasDocFormat returns the AtlasDocFormat field if non-nil, zero value otherwise.

### GetAtlasDocFormatOk

`func (o *CreateBlogPostRequestBody) GetAtlasDocFormatOk() (*BlogPostBodyWrite, bool)`

GetAtlasDocFormatOk returns a tuple with the AtlasDocFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasDocFormat

`func (o *CreateBlogPostRequestBody) SetAtlasDocFormat(v BlogPostBodyWrite)`

SetAtlasDocFormat sets AtlasDocFormat field to given value.

### HasAtlasDocFormat

`func (o *CreateBlogPostRequestBody) HasAtlasDocFormat() bool`

HasAtlasDocFormat returns a boolean if a field has been set.

### GetWiki

`func (o *CreateBlogPostRequestBody) GetWiki() BlogPostBodyWrite`

GetWiki returns the Wiki field if non-nil, zero value otherwise.

### GetWikiOk

`func (o *CreateBlogPostRequestBody) GetWikiOk() (*BlogPostBodyWrite, bool)`

GetWikiOk returns a tuple with the Wiki field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiki

`func (o *CreateBlogPostRequestBody) SetWiki(v BlogPostBodyWrite)`

SetWiki sets Wiki field to given value.

### HasWiki

`func (o *CreateBlogPostRequestBody) HasWiki() bool`

HasWiki returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



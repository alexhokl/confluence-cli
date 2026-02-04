# CreateFooterCommentModelBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Representation** | Pointer to **string** | Type of content representation used for the value field. | [optional] 
**Value** | Pointer to **string** | Body of the comment, in the format found in the representation field. | [optional] 
**Storage** | Pointer to [**CommentBodyWrite**](CommentBodyWrite.md) |  | [optional] 
**AtlasDocFormat** | Pointer to [**CommentBodyWrite**](CommentBodyWrite.md) |  | [optional] 
**Wiki** | Pointer to [**CommentBodyWrite**](CommentBodyWrite.md) |  | [optional] 

## Methods

### NewCreateFooterCommentModelBody

`func NewCreateFooterCommentModelBody() *CreateFooterCommentModelBody`

NewCreateFooterCommentModelBody instantiates a new CreateFooterCommentModelBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFooterCommentModelBodyWithDefaults

`func NewCreateFooterCommentModelBodyWithDefaults() *CreateFooterCommentModelBody`

NewCreateFooterCommentModelBodyWithDefaults instantiates a new CreateFooterCommentModelBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepresentation

`func (o *CreateFooterCommentModelBody) GetRepresentation() string`

GetRepresentation returns the Representation field if non-nil, zero value otherwise.

### GetRepresentationOk

`func (o *CreateFooterCommentModelBody) GetRepresentationOk() (*string, bool)`

GetRepresentationOk returns a tuple with the Representation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepresentation

`func (o *CreateFooterCommentModelBody) SetRepresentation(v string)`

SetRepresentation sets Representation field to given value.

### HasRepresentation

`func (o *CreateFooterCommentModelBody) HasRepresentation() bool`

HasRepresentation returns a boolean if a field has been set.

### GetValue

`func (o *CreateFooterCommentModelBody) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateFooterCommentModelBody) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateFooterCommentModelBody) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CreateFooterCommentModelBody) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetStorage

`func (o *CreateFooterCommentModelBody) GetStorage() CommentBodyWrite`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CreateFooterCommentModelBody) GetStorageOk() (*CommentBodyWrite, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CreateFooterCommentModelBody) SetStorage(v CommentBodyWrite)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *CreateFooterCommentModelBody) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetAtlasDocFormat

`func (o *CreateFooterCommentModelBody) GetAtlasDocFormat() CommentBodyWrite`

GetAtlasDocFormat returns the AtlasDocFormat field if non-nil, zero value otherwise.

### GetAtlasDocFormatOk

`func (o *CreateFooterCommentModelBody) GetAtlasDocFormatOk() (*CommentBodyWrite, bool)`

GetAtlasDocFormatOk returns a tuple with the AtlasDocFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasDocFormat

`func (o *CreateFooterCommentModelBody) SetAtlasDocFormat(v CommentBodyWrite)`

SetAtlasDocFormat sets AtlasDocFormat field to given value.

### HasAtlasDocFormat

`func (o *CreateFooterCommentModelBody) HasAtlasDocFormat() bool`

HasAtlasDocFormat returns a boolean if a field has been set.

### GetWiki

`func (o *CreateFooterCommentModelBody) GetWiki() CommentBodyWrite`

GetWiki returns the Wiki field if non-nil, zero value otherwise.

### GetWikiOk

`func (o *CreateFooterCommentModelBody) GetWikiOk() (*CommentBodyWrite, bool)`

GetWikiOk returns a tuple with the Wiki field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiki

`func (o *CreateFooterCommentModelBody) SetWiki(v CommentBodyWrite)`

SetWiki sets Wiki field to given value.

### HasWiki

`func (o *CreateFooterCommentModelBody) HasWiki() bool`

HasWiki returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CommentBodyWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Representation** | Pointer to **string** | Type of content representation used for the value field. | [optional] 
**Value** | Pointer to **string** | Body of the comment, in the format found in the representation field. | [optional] 

## Methods

### NewCommentBodyWrite

`func NewCommentBodyWrite() *CommentBodyWrite`

NewCommentBodyWrite instantiates a new CommentBodyWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentBodyWriteWithDefaults

`func NewCommentBodyWriteWithDefaults() *CommentBodyWrite`

NewCommentBodyWriteWithDefaults instantiates a new CommentBodyWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepresentation

`func (o *CommentBodyWrite) GetRepresentation() string`

GetRepresentation returns the Representation field if non-nil, zero value otherwise.

### GetRepresentationOk

`func (o *CommentBodyWrite) GetRepresentationOk() (*string, bool)`

GetRepresentationOk returns a tuple with the Representation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepresentation

`func (o *CommentBodyWrite) SetRepresentation(v string)`

SetRepresentation sets Representation field to given value.

### HasRepresentation

`func (o *CommentBodyWrite) HasRepresentation() bool`

HasRepresentation returns a boolean if a field has been set.

### GetValue

`func (o *CommentBodyWrite) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CommentBodyWrite) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CommentBodyWrite) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CommentBodyWrite) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateInlineCommentModelInlineCommentProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TextSelection** | Pointer to **string** | The text to highlight | [optional] 
**TextSelectionMatchCount** | Pointer to **int32** | The number of matches for the selected text on the page (should be strictly greater than textSelectionMatchIndex) | [optional] 
**TextSelectionMatchIndex** | Pointer to **int32** | The match index to highlight. This is zero-based. E.g. if you have 3 occurrences of \&quot;hello world\&quot; on a page  and you want to highlight the second occurrence, you should pass 1 for textSelectionMatchIndex and 3 for textSelectionMatchCount. | [optional] 

## Methods

### NewCreateInlineCommentModelInlineCommentProperties

`func NewCreateInlineCommentModelInlineCommentProperties() *CreateInlineCommentModelInlineCommentProperties`

NewCreateInlineCommentModelInlineCommentProperties instantiates a new CreateInlineCommentModelInlineCommentProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInlineCommentModelInlineCommentPropertiesWithDefaults

`func NewCreateInlineCommentModelInlineCommentPropertiesWithDefaults() *CreateInlineCommentModelInlineCommentProperties`

NewCreateInlineCommentModelInlineCommentPropertiesWithDefaults instantiates a new CreateInlineCommentModelInlineCommentProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTextSelection

`func (o *CreateInlineCommentModelInlineCommentProperties) GetTextSelection() string`

GetTextSelection returns the TextSelection field if non-nil, zero value otherwise.

### GetTextSelectionOk

`func (o *CreateInlineCommentModelInlineCommentProperties) GetTextSelectionOk() (*string, bool)`

GetTextSelectionOk returns a tuple with the TextSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextSelection

`func (o *CreateInlineCommentModelInlineCommentProperties) SetTextSelection(v string)`

SetTextSelection sets TextSelection field to given value.

### HasTextSelection

`func (o *CreateInlineCommentModelInlineCommentProperties) HasTextSelection() bool`

HasTextSelection returns a boolean if a field has been set.

### GetTextSelectionMatchCount

`func (o *CreateInlineCommentModelInlineCommentProperties) GetTextSelectionMatchCount() int32`

GetTextSelectionMatchCount returns the TextSelectionMatchCount field if non-nil, zero value otherwise.

### GetTextSelectionMatchCountOk

`func (o *CreateInlineCommentModelInlineCommentProperties) GetTextSelectionMatchCountOk() (*int32, bool)`

GetTextSelectionMatchCountOk returns a tuple with the TextSelectionMatchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextSelectionMatchCount

`func (o *CreateInlineCommentModelInlineCommentProperties) SetTextSelectionMatchCount(v int32)`

SetTextSelectionMatchCount sets TextSelectionMatchCount field to given value.

### HasTextSelectionMatchCount

`func (o *CreateInlineCommentModelInlineCommentProperties) HasTextSelectionMatchCount() bool`

HasTextSelectionMatchCount returns a boolean if a field has been set.

### GetTextSelectionMatchIndex

`func (o *CreateInlineCommentModelInlineCommentProperties) GetTextSelectionMatchIndex() int32`

GetTextSelectionMatchIndex returns the TextSelectionMatchIndex field if non-nil, zero value otherwise.

### GetTextSelectionMatchIndexOk

`func (o *CreateInlineCommentModelInlineCommentProperties) GetTextSelectionMatchIndexOk() (*int32, bool)`

GetTextSelectionMatchIndexOk returns a tuple with the TextSelectionMatchIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextSelectionMatchIndex

`func (o *CreateInlineCommentModelInlineCommentProperties) SetTextSelectionMatchIndex(v int32)`

SetTextSelectionMatchIndex sets TextSelectionMatchIndex field to given value.

### HasTextSelectionMatchIndex

`func (o *CreateInlineCommentModelInlineCommentProperties) HasTextSelectionMatchIndex() bool`

HasTextSelectionMatchIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



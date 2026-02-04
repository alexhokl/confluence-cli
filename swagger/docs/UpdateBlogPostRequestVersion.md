# UpdateBlogPostRequestVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | The new version number of the updated blog post.  Set this to the current version number plus one, unless you are updating the status to &#39;draft&#39; which requires a version number of 1.  If you don&#39;t know the current version number, use Get blog post by id. | [optional] 
**Message** | Pointer to **string** | An optional message to be stored with the version. | [optional] 

## Methods

### NewUpdateBlogPostRequestVersion

`func NewUpdateBlogPostRequestVersion() *UpdateBlogPostRequestVersion`

NewUpdateBlogPostRequestVersion instantiates a new UpdateBlogPostRequestVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBlogPostRequestVersionWithDefaults

`func NewUpdateBlogPostRequestVersionWithDefaults() *UpdateBlogPostRequestVersion`

NewUpdateBlogPostRequestVersionWithDefaults instantiates a new UpdateBlogPostRequestVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *UpdateBlogPostRequestVersion) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *UpdateBlogPostRequestVersion) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *UpdateBlogPostRequestVersion) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *UpdateBlogPostRequestVersion) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetMessage

`func (o *UpdateBlogPostRequestVersion) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateBlogPostRequestVersion) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateBlogPostRequestVersion) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpdateBlogPostRequestVersion) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



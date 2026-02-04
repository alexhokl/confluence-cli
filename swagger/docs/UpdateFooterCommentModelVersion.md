# UpdateFooterCommentModelVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | Number of new version. Should be 1 higher than current version of the comment. | [optional] 
**Message** | Pointer to **string** | Optional message store for the new version. | [optional] 

## Methods

### NewUpdateFooterCommentModelVersion

`func NewUpdateFooterCommentModelVersion() *UpdateFooterCommentModelVersion`

NewUpdateFooterCommentModelVersion instantiates a new UpdateFooterCommentModelVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFooterCommentModelVersionWithDefaults

`func NewUpdateFooterCommentModelVersionWithDefaults() *UpdateFooterCommentModelVersion`

NewUpdateFooterCommentModelVersionWithDefaults instantiates a new UpdateFooterCommentModelVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *UpdateFooterCommentModelVersion) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *UpdateFooterCommentModelVersion) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *UpdateFooterCommentModelVersion) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *UpdateFooterCommentModelVersion) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetMessage

`func (o *UpdateFooterCommentModelVersion) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateFooterCommentModelVersion) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateFooterCommentModelVersion) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpdateFooterCommentModelVersion) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



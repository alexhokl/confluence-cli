# CommentVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | Date and time when the version was created. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 
**Message** | Pointer to **string** | Message associated with the current version. | [optional] 
**Number** | Pointer to **int32** | The version number. | [optional] 
**MinorEdit** | Pointer to **bool** | Describes if this version is a minor version. Email notifications and activity stream updates are not created for minor versions. | [optional] 
**AuthorId** | Pointer to **string** | The account ID of the user who created this version. | [optional] 
**Comment** | Pointer to [**VersionedEntity**](VersionedEntity.md) |  | [optional] 

## Methods

### NewCommentVersion

`func NewCommentVersion() *CommentVersion`

NewCommentVersion instantiates a new CommentVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentVersionWithDefaults

`func NewCommentVersionWithDefaults() *CommentVersion`

NewCommentVersionWithDefaults instantiates a new CommentVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CommentVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommentVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommentVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CommentVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMessage

`func (o *CommentVersion) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CommentVersion) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CommentVersion) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CommentVersion) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNumber

`func (o *CommentVersion) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CommentVersion) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CommentVersion) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CommentVersion) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetMinorEdit

`func (o *CommentVersion) GetMinorEdit() bool`

GetMinorEdit returns the MinorEdit field if non-nil, zero value otherwise.

### GetMinorEditOk

`func (o *CommentVersion) GetMinorEditOk() (*bool, bool)`

GetMinorEditOk returns a tuple with the MinorEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorEdit

`func (o *CommentVersion) SetMinorEdit(v bool)`

SetMinorEdit sets MinorEdit field to given value.

### HasMinorEdit

`func (o *CommentVersion) HasMinorEdit() bool`

HasMinorEdit returns a boolean if a field has been set.

### GetAuthorId

`func (o *CommentVersion) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *CommentVersion) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *CommentVersion) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *CommentVersion) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetComment

`func (o *CommentVersion) GetComment() VersionedEntity`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CommentVersion) GetCommentOk() (*VersionedEntity, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CommentVersion) SetComment(v VersionedEntity)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CommentVersion) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



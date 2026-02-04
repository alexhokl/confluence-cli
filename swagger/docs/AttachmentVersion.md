# AttachmentVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | Date and time when the version was created. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 
**Message** | Pointer to **string** | Message associated with the current version. | [optional] 
**Number** | Pointer to **int32** | The version number. | [optional] 
**MinorEdit** | Pointer to **bool** | Describes if this version is a minor version. Email notifications and activity stream updates are not created for minor versions. | [optional] 
**AuthorId** | Pointer to **string** | The account ID of the user who created this version. | [optional] 
**Attachment** | Pointer to [**VersionedEntity**](VersionedEntity.md) |  | [optional] 

## Methods

### NewAttachmentVersion

`func NewAttachmentVersion() *AttachmentVersion`

NewAttachmentVersion instantiates a new AttachmentVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentVersionWithDefaults

`func NewAttachmentVersionWithDefaults() *AttachmentVersion`

NewAttachmentVersionWithDefaults instantiates a new AttachmentVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AttachmentVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AttachmentVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AttachmentVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AttachmentVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMessage

`func (o *AttachmentVersion) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AttachmentVersion) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AttachmentVersion) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AttachmentVersion) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNumber

`func (o *AttachmentVersion) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *AttachmentVersion) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *AttachmentVersion) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *AttachmentVersion) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetMinorEdit

`func (o *AttachmentVersion) GetMinorEdit() bool`

GetMinorEdit returns the MinorEdit field if non-nil, zero value otherwise.

### GetMinorEditOk

`func (o *AttachmentVersion) GetMinorEditOk() (*bool, bool)`

GetMinorEditOk returns a tuple with the MinorEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorEdit

`func (o *AttachmentVersion) SetMinorEdit(v bool)`

SetMinorEdit sets MinorEdit field to given value.

### HasMinorEdit

`func (o *AttachmentVersion) HasMinorEdit() bool`

HasMinorEdit returns a boolean if a field has been set.

### GetAuthorId

`func (o *AttachmentVersion) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *AttachmentVersion) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *AttachmentVersion) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *AttachmentVersion) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetAttachment

`func (o *AttachmentVersion) GetAttachment() VersionedEntity`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *AttachmentVersion) GetAttachmentOk() (*VersionedEntity, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *AttachmentVersion) SetAttachment(v VersionedEntity)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *AttachmentVersion) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



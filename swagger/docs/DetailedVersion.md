# DetailedVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | The current version number. | [optional] 
**AuthorId** | Pointer to **string** | The account ID of the user who created this version. | [optional] 
**Message** | Pointer to **string** | Message associated with the current version. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time when the version was created. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 
**MinorEdit** | Pointer to **bool** | Describes if this version is a minor version. Email notifications and activity stream updates are not created for minor versions. | [optional] 
**ContentTypeModified** | Pointer to **bool** | Describes if the content type is modified in this version (e.g. page to blog) | [optional] 
**Collaborators** | Pointer to **[]string** | The account IDs of users that collaborated on this version. | [optional] 
**PrevVersion** | Pointer to **int32** | The version number of the version prior to this current content update. | [optional] 
**NextVersion** | Pointer to **int32** | The version number of the version after this current content update. | [optional] 

## Methods

### NewDetailedVersion

`func NewDetailedVersion() *DetailedVersion`

NewDetailedVersion instantiates a new DetailedVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedVersionWithDefaults

`func NewDetailedVersionWithDefaults() *DetailedVersion`

NewDetailedVersionWithDefaults instantiates a new DetailedVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *DetailedVersion) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *DetailedVersion) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *DetailedVersion) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *DetailedVersion) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetAuthorId

`func (o *DetailedVersion) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *DetailedVersion) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *DetailedVersion) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *DetailedVersion) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetMessage

`func (o *DetailedVersion) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DetailedVersion) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DetailedVersion) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DetailedVersion) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DetailedVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DetailedVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DetailedVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DetailedVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMinorEdit

`func (o *DetailedVersion) GetMinorEdit() bool`

GetMinorEdit returns the MinorEdit field if non-nil, zero value otherwise.

### GetMinorEditOk

`func (o *DetailedVersion) GetMinorEditOk() (*bool, bool)`

GetMinorEditOk returns a tuple with the MinorEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorEdit

`func (o *DetailedVersion) SetMinorEdit(v bool)`

SetMinorEdit sets MinorEdit field to given value.

### HasMinorEdit

`func (o *DetailedVersion) HasMinorEdit() bool`

HasMinorEdit returns a boolean if a field has been set.

### GetContentTypeModified

`func (o *DetailedVersion) GetContentTypeModified() bool`

GetContentTypeModified returns the ContentTypeModified field if non-nil, zero value otherwise.

### GetContentTypeModifiedOk

`func (o *DetailedVersion) GetContentTypeModifiedOk() (*bool, bool)`

GetContentTypeModifiedOk returns a tuple with the ContentTypeModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypeModified

`func (o *DetailedVersion) SetContentTypeModified(v bool)`

SetContentTypeModified sets ContentTypeModified field to given value.

### HasContentTypeModified

`func (o *DetailedVersion) HasContentTypeModified() bool`

HasContentTypeModified returns a boolean if a field has been set.

### GetCollaborators

`func (o *DetailedVersion) GetCollaborators() []string`

GetCollaborators returns the Collaborators field if non-nil, zero value otherwise.

### GetCollaboratorsOk

`func (o *DetailedVersion) GetCollaboratorsOk() (*[]string, bool)`

GetCollaboratorsOk returns a tuple with the Collaborators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaborators

`func (o *DetailedVersion) SetCollaborators(v []string)`

SetCollaborators sets Collaborators field to given value.

### HasCollaborators

`func (o *DetailedVersion) HasCollaborators() bool`

HasCollaborators returns a boolean if a field has been set.

### GetPrevVersion

`func (o *DetailedVersion) GetPrevVersion() int32`

GetPrevVersion returns the PrevVersion field if non-nil, zero value otherwise.

### GetPrevVersionOk

`func (o *DetailedVersion) GetPrevVersionOk() (*int32, bool)`

GetPrevVersionOk returns a tuple with the PrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevVersion

`func (o *DetailedVersion) SetPrevVersion(v int32)`

SetPrevVersion sets PrevVersion field to given value.

### HasPrevVersion

`func (o *DetailedVersion) HasPrevVersion() bool`

HasPrevVersion returns a boolean if a field has been set.

### GetNextVersion

`func (o *DetailedVersion) GetNextVersion() int32`

GetNextVersion returns the NextVersion field if non-nil, zero value otherwise.

### GetNextVersionOk

`func (o *DetailedVersion) GetNextVersionOk() (*int32, bool)`

GetNextVersionOk returns a tuple with the NextVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextVersion

`func (o *DetailedVersion) SetNextVersion(v int32)`

SetNextVersion sets NextVersion field to given value.

### HasNextVersion

`func (o *DetailedVersion) HasNextVersion() bool`

HasNextVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PageVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | Date and time when the version was created. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 
**Message** | Pointer to **string** | Message associated with the current version. | [optional] 
**Number** | Pointer to **int32** | The version number. | [optional] 
**MinorEdit** | Pointer to **bool** | Describes if this version is a minor version. Email notifications and activity stream updates are not created for minor versions. | [optional] 
**AuthorId** | Pointer to **string** | The account ID of the user who created this version. | [optional] 
**Page** | Pointer to [**VersionedEntity**](VersionedEntity.md) |  | [optional] 

## Methods

### NewPageVersion

`func NewPageVersion() *PageVersion`

NewPageVersion instantiates a new PageVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageVersionWithDefaults

`func NewPageVersionWithDefaults() *PageVersion`

NewPageVersionWithDefaults instantiates a new PageVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PageVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PageVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PageVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PageVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMessage

`func (o *PageVersion) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PageVersion) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PageVersion) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PageVersion) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNumber

`func (o *PageVersion) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PageVersion) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PageVersion) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PageVersion) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetMinorEdit

`func (o *PageVersion) GetMinorEdit() bool`

GetMinorEdit returns the MinorEdit field if non-nil, zero value otherwise.

### GetMinorEditOk

`func (o *PageVersion) GetMinorEditOk() (*bool, bool)`

GetMinorEditOk returns a tuple with the MinorEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorEdit

`func (o *PageVersion) SetMinorEdit(v bool)`

SetMinorEdit sets MinorEdit field to given value.

### HasMinorEdit

`func (o *PageVersion) HasMinorEdit() bool`

HasMinorEdit returns a boolean if a field has been set.

### GetAuthorId

`func (o *PageVersion) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *PageVersion) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *PageVersion) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *PageVersion) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetPage

`func (o *PageVersion) GetPage() VersionedEntity`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PageVersion) GetPageOk() (*VersionedEntity, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PageVersion) SetPage(v VersionedEntity)`

SetPage sets Page field to given value.

### HasPage

`func (o *PageVersion) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



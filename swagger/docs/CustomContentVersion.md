# CustomContentVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | Date and time when the version was created. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 
**Message** | Pointer to **string** | Message associated with the current version. | [optional] 
**Number** | Pointer to **int32** | The version number. | [optional] 
**MinorEdit** | Pointer to **bool** | Describes if this version is a minor version. Email notifications and activity stream updates are not created for minor versions. | [optional] 
**AuthorId** | Pointer to **string** | The account ID of the user who created this version. | [optional] 
**Custom** | Pointer to [**VersionedEntity**](VersionedEntity.md) |  | [optional] 

## Methods

### NewCustomContentVersion

`func NewCustomContentVersion() *CustomContentVersion`

NewCustomContentVersion instantiates a new CustomContentVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomContentVersionWithDefaults

`func NewCustomContentVersionWithDefaults() *CustomContentVersion`

NewCustomContentVersionWithDefaults instantiates a new CustomContentVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CustomContentVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomContentVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomContentVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CustomContentVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMessage

`func (o *CustomContentVersion) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CustomContentVersion) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CustomContentVersion) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CustomContentVersion) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNumber

`func (o *CustomContentVersion) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CustomContentVersion) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CustomContentVersion) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CustomContentVersion) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetMinorEdit

`func (o *CustomContentVersion) GetMinorEdit() bool`

GetMinorEdit returns the MinorEdit field if non-nil, zero value otherwise.

### GetMinorEditOk

`func (o *CustomContentVersion) GetMinorEditOk() (*bool, bool)`

GetMinorEditOk returns a tuple with the MinorEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorEdit

`func (o *CustomContentVersion) SetMinorEdit(v bool)`

SetMinorEdit sets MinorEdit field to given value.

### HasMinorEdit

`func (o *CustomContentVersion) HasMinorEdit() bool`

HasMinorEdit returns a boolean if a field has been set.

### GetAuthorId

`func (o *CustomContentVersion) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *CustomContentVersion) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *CustomContentVersion) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *CustomContentVersion) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetCustom

`func (o *CustomContentVersion) GetCustom() VersionedEntity`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *CustomContentVersion) GetCustomOk() (*VersionedEntity, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *CustomContentVersion) SetCustom(v VersionedEntity)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *CustomContentVersion) HasCustom() bool`

HasCustom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



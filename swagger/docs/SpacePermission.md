# SpacePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The identifier for the space permission. | [optional] 
**DisplayName** | Pointer to **string** | The display name for the space permission. | [optional] 
**Description** | Pointer to **string** | Describes the space permission’s usage. | [optional] 
**RequiredPermissionIds** | Pointer to **[]string** | The permissions required for this permission to be enabled. | [optional] 

## Methods

### NewSpacePermission

`func NewSpacePermission() *SpacePermission`

NewSpacePermission instantiates a new SpacePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpacePermissionWithDefaults

`func NewSpacePermissionWithDefaults() *SpacePermission`

NewSpacePermissionWithDefaults instantiates a new SpacePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SpacePermission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpacePermission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpacePermission) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SpacePermission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *SpacePermission) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SpacePermission) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SpacePermission) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SpacePermission) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *SpacePermission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SpacePermission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SpacePermission) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SpacePermission) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequiredPermissionIds

`func (o *SpacePermission) GetRequiredPermissionIds() []string`

GetRequiredPermissionIds returns the RequiredPermissionIds field if non-nil, zero value otherwise.

### GetRequiredPermissionIdsOk

`func (o *SpacePermission) GetRequiredPermissionIdsOk() (*[]string, bool)`

GetRequiredPermissionIdsOk returns a tuple with the RequiredPermissionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredPermissionIds

`func (o *SpacePermission) SetRequiredPermissionIds(v []string)`

SetRequiredPermissionIds sets RequiredPermissionIds field to given value.

### HasRequiredPermissionIds

`func (o *SpacePermission) HasRequiredPermissionIds() bool`

HasRequiredPermissionIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



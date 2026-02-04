# SpaceRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The identifier for the space role. | [optional] 
**Type** | Pointer to [**RoleType**](RoleType.md) |  | [optional] 
**Name** | Pointer to **string** | The name for the space role. | [optional] 
**Description** | Pointer to **string** | The description for the space role’s usage. | [optional] 
**SpacePermissions** | Pointer to **[]string** | The space permissions the space role is comprised of. | [optional] 

## Methods

### NewSpaceRole

`func NewSpaceRole() *SpaceRole`

NewSpaceRole instantiates a new SpaceRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceRoleWithDefaults

`func NewSpaceRoleWithDefaults() *SpaceRole`

NewSpaceRoleWithDefaults instantiates a new SpaceRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SpaceRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpaceRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpaceRole) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SpaceRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SpaceRole) GetType() RoleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpaceRole) GetTypeOk() (*RoleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpaceRole) SetType(v RoleType)`

SetType sets Type field to given value.

### HasType

`func (o *SpaceRole) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *SpaceRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpaceRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpaceRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SpaceRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SpaceRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SpaceRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SpaceRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SpaceRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSpacePermissions

`func (o *SpaceRole) GetSpacePermissions() []string`

GetSpacePermissions returns the SpacePermissions field if non-nil, zero value otherwise.

### GetSpacePermissionsOk

`func (o *SpaceRole) GetSpacePermissionsOk() (*[]string, bool)`

GetSpacePermissionsOk returns a tuple with the SpacePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpacePermissions

`func (o *SpaceRole) SetSpacePermissions(v []string)`

SetSpacePermissions sets SpacePermissions field to given value.

### HasSpacePermissions

`func (o *SpaceRole) HasSpacePermissions() bool`

HasSpacePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



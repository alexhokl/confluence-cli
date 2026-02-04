# UpdateSpaceRoleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the space role | [optional] 
**Type** | Pointer to [**RoleType**](RoleType.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the space role | [optional] 
**Description** | Pointer to **string** | Description for the space role | [optional] 
**TaskId** | Pointer to **string** | Id of the task to update the space permissions associated with the space role | [optional] 

## Methods

### NewUpdateSpaceRoleResponse

`func NewUpdateSpaceRoleResponse() *UpdateSpaceRoleResponse`

NewUpdateSpaceRoleResponse instantiates a new UpdateSpaceRoleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSpaceRoleResponseWithDefaults

`func NewUpdateSpaceRoleResponseWithDefaults() *UpdateSpaceRoleResponse`

NewUpdateSpaceRoleResponseWithDefaults instantiates a new UpdateSpaceRoleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateSpaceRoleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateSpaceRoleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateSpaceRoleResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateSpaceRoleResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UpdateSpaceRoleResponse) GetType() RoleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateSpaceRoleResponse) GetTypeOk() (*RoleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateSpaceRoleResponse) SetType(v RoleType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateSpaceRoleResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *UpdateSpaceRoleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSpaceRoleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSpaceRoleResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSpaceRoleResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateSpaceRoleResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSpaceRoleResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSpaceRoleResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSpaceRoleResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTaskId

`func (o *UpdateSpaceRoleResponse) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *UpdateSpaceRoleResponse) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *UpdateSpaceRoleResponse) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *UpdateSpaceRoleResponse) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



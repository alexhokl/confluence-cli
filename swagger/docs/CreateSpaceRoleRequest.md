# CreateSpaceRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the space role | 
**Description** | **string** | Description for the space role | 
**SpacePermissions** | **[]string** | The ids of the space permissions associated with the space role. Sample value \&quot;read/space\&quot;; retrieve ids from responses returned by [GET /space-permissions](https://developer.atlassian.com/cloud/confluence/rest/v2/api-group-space-permissions/#api-space-permissions-get) endpoint | 

## Methods

### NewCreateSpaceRoleRequest

`func NewCreateSpaceRoleRequest(name string, description string, spacePermissions []string, ) *CreateSpaceRoleRequest`

NewCreateSpaceRoleRequest instantiates a new CreateSpaceRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSpaceRoleRequestWithDefaults

`func NewCreateSpaceRoleRequestWithDefaults() *CreateSpaceRoleRequest`

NewCreateSpaceRoleRequestWithDefaults instantiates a new CreateSpaceRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSpaceRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSpaceRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSpaceRoleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateSpaceRoleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSpaceRoleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSpaceRoleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSpacePermissions

`func (o *CreateSpaceRoleRequest) GetSpacePermissions() []string`

GetSpacePermissions returns the SpacePermissions field if non-nil, zero value otherwise.

### GetSpacePermissionsOk

`func (o *CreateSpaceRoleRequest) GetSpacePermissionsOk() (*[]string, bool)`

GetSpacePermissionsOk returns a tuple with the SpacePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpacePermissions

`func (o *CreateSpaceRoleRequest) SetSpacePermissions(v []string)`

SetSpacePermissions sets SpacePermissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UpdateSpaceRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the space role | 
**Description** | **string** | Description for the space role | 
**SpacePermissions** | **[]string** | The ids of the space permissions associated with the space role. Sample value \&quot;read/space\&quot;; retrieve ids from responses returned by [GET /space-permissions](https://developer.atlassian.com/cloud/confluence/rest/v2/api-group-space-permissions/#api-space-permissions-get) endpoint | 
**AnonymousReassignmentRoleId** | Pointer to **string** | If space anonymous access is assigned to the role being modified, the Id of a role to migrate those assignments to can be specified. Anonymous access role assignments left unchanged if unspecified. | [optional] 
**GuestReassignmentRoleId** | Pointer to **string** | If guests are assigned to the role being modified, the Id of a role to migrate those assignments to can be specified. Guest role assignments left unchanged if unspecified. | [optional] 

## Methods

### NewUpdateSpaceRoleRequest

`func NewUpdateSpaceRoleRequest(name string, description string, spacePermissions []string, ) *UpdateSpaceRoleRequest`

NewUpdateSpaceRoleRequest instantiates a new UpdateSpaceRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSpaceRoleRequestWithDefaults

`func NewUpdateSpaceRoleRequestWithDefaults() *UpdateSpaceRoleRequest`

NewUpdateSpaceRoleRequestWithDefaults instantiates a new UpdateSpaceRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateSpaceRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSpaceRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSpaceRoleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateSpaceRoleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSpaceRoleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSpaceRoleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSpacePermissions

`func (o *UpdateSpaceRoleRequest) GetSpacePermissions() []string`

GetSpacePermissions returns the SpacePermissions field if non-nil, zero value otherwise.

### GetSpacePermissionsOk

`func (o *UpdateSpaceRoleRequest) GetSpacePermissionsOk() (*[]string, bool)`

GetSpacePermissionsOk returns a tuple with the SpacePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpacePermissions

`func (o *UpdateSpaceRoleRequest) SetSpacePermissions(v []string)`

SetSpacePermissions sets SpacePermissions field to given value.


### GetAnonymousReassignmentRoleId

`func (o *UpdateSpaceRoleRequest) GetAnonymousReassignmentRoleId() string`

GetAnonymousReassignmentRoleId returns the AnonymousReassignmentRoleId field if non-nil, zero value otherwise.

### GetAnonymousReassignmentRoleIdOk

`func (o *UpdateSpaceRoleRequest) GetAnonymousReassignmentRoleIdOk() (*string, bool)`

GetAnonymousReassignmentRoleIdOk returns a tuple with the AnonymousReassignmentRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousReassignmentRoleId

`func (o *UpdateSpaceRoleRequest) SetAnonymousReassignmentRoleId(v string)`

SetAnonymousReassignmentRoleId sets AnonymousReassignmentRoleId field to given value.

### HasAnonymousReassignmentRoleId

`func (o *UpdateSpaceRoleRequest) HasAnonymousReassignmentRoleId() bool`

HasAnonymousReassignmentRoleId returns a boolean if a field has been set.

### GetGuestReassignmentRoleId

`func (o *UpdateSpaceRoleRequest) GetGuestReassignmentRoleId() string`

GetGuestReassignmentRoleId returns the GuestReassignmentRoleId field if non-nil, zero value otherwise.

### GetGuestReassignmentRoleIdOk

`func (o *UpdateSpaceRoleRequest) GetGuestReassignmentRoleIdOk() (*string, bool)`

GetGuestReassignmentRoleIdOk returns a tuple with the GuestReassignmentRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestReassignmentRoleId

`func (o *UpdateSpaceRoleRequest) SetGuestReassignmentRoleId(v string)`

SetGuestReassignmentRoleId sets GuestReassignmentRoleId field to given value.

### HasGuestReassignmentRoleId

`func (o *UpdateSpaceRoleRequest) HasGuestReassignmentRoleId() bool`

HasGuestReassignmentRoleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



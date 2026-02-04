# SpaceRoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Principal** | Pointer to [**Principal**](Principal.md) |  | [optional] 
**RoleId** | Pointer to **string** | The role to which the principal is assigned. | [optional] 

## Methods

### NewSpaceRoleAssignment

`func NewSpaceRoleAssignment() *SpaceRoleAssignment`

NewSpaceRoleAssignment instantiates a new SpaceRoleAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceRoleAssignmentWithDefaults

`func NewSpaceRoleAssignmentWithDefaults() *SpaceRoleAssignment`

NewSpaceRoleAssignmentWithDefaults instantiates a new SpaceRoleAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipal

`func (o *SpaceRoleAssignment) GetPrincipal() Principal`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *SpaceRoleAssignment) GetPrincipalOk() (*Principal, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *SpaceRoleAssignment) SetPrincipal(v Principal)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *SpaceRoleAssignment) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetRoleId

`func (o *SpaceRoleAssignment) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *SpaceRoleAssignment) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *SpaceRoleAssignment) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *SpaceRoleAssignment) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



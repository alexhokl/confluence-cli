# SetSpaceRoleAssignmentsRequestInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Principal** | [**Principal**](Principal.md) |  | 
**RoleId** | Pointer to **string** | The role to which the principal is assigned. | [optional] 

## Methods

### NewSetSpaceRoleAssignmentsRequestInner

`func NewSetSpaceRoleAssignmentsRequestInner(principal Principal, ) *SetSpaceRoleAssignmentsRequestInner`

NewSetSpaceRoleAssignmentsRequestInner instantiates a new SetSpaceRoleAssignmentsRequestInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetSpaceRoleAssignmentsRequestInnerWithDefaults

`func NewSetSpaceRoleAssignmentsRequestInnerWithDefaults() *SetSpaceRoleAssignmentsRequestInner`

NewSetSpaceRoleAssignmentsRequestInnerWithDefaults instantiates a new SetSpaceRoleAssignmentsRequestInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipal

`func (o *SetSpaceRoleAssignmentsRequestInner) GetPrincipal() Principal`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *SetSpaceRoleAssignmentsRequestInner) GetPrincipalOk() (*Principal, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *SetSpaceRoleAssignmentsRequestInner) SetPrincipal(v Principal)`

SetPrincipal sets Principal field to given value.


### GetRoleId

`func (o *SetSpaceRoleAssignmentsRequestInner) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *SetSpaceRoleAssignmentsRequestInner) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *SetSpaceRoleAssignmentsRequestInner) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *SetSpaceRoleAssignmentsRequestInner) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SpacePermissionAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the space permission. | [optional] 
**Principal** | Pointer to [**SpacePermissionAssignmentPrincipal**](SpacePermissionAssignmentPrincipal.md) |  | [optional] 
**Operation** | Pointer to [**SpacePermissionAssignmentOperation**](SpacePermissionAssignmentOperation.md) |  | [optional] 

## Methods

### NewSpacePermissionAssignment

`func NewSpacePermissionAssignment() *SpacePermissionAssignment`

NewSpacePermissionAssignment instantiates a new SpacePermissionAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpacePermissionAssignmentWithDefaults

`func NewSpacePermissionAssignmentWithDefaults() *SpacePermissionAssignment`

NewSpacePermissionAssignmentWithDefaults instantiates a new SpacePermissionAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SpacePermissionAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpacePermissionAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpacePermissionAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SpacePermissionAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrincipal

`func (o *SpacePermissionAssignment) GetPrincipal() SpacePermissionAssignmentPrincipal`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *SpacePermissionAssignment) GetPrincipalOk() (*SpacePermissionAssignmentPrincipal, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *SpacePermissionAssignment) SetPrincipal(v SpacePermissionAssignmentPrincipal)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *SpacePermissionAssignment) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetOperation

`func (o *SpacePermissionAssignment) GetOperation() SpacePermissionAssignmentOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *SpacePermissionAssignment) GetOperationOk() (*SpacePermissionAssignmentOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *SpacePermissionAssignment) SetOperation(v SpacePermissionAssignmentOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *SpacePermissionAssignment) HasOperation() bool`

HasOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



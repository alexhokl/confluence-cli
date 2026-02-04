# SpaceSinglePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]SpacePermissionAssignment**](SpacePermissionAssignment.md) |  | [optional] 
**Meta** | Pointer to [**OptionalFieldMeta**](OptionalFieldMeta.md) |  | [optional] 
**Links** | Pointer to [**OptionalFieldLinks**](OptionalFieldLinks.md) |  | [optional] 

## Methods

### NewSpaceSinglePermissions

`func NewSpaceSinglePermissions() *SpaceSinglePermissions`

NewSpaceSinglePermissions instantiates a new SpaceSinglePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceSinglePermissionsWithDefaults

`func NewSpaceSinglePermissionsWithDefaults() *SpaceSinglePermissions`

NewSpaceSinglePermissionsWithDefaults instantiates a new SpaceSinglePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *SpaceSinglePermissions) GetResults() []SpacePermissionAssignment`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SpaceSinglePermissions) GetResultsOk() (*[]SpacePermissionAssignment, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SpaceSinglePermissions) SetResults(v []SpacePermissionAssignment)`

SetResults sets Results field to given value.

### HasResults

`func (o *SpaceSinglePermissions) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetMeta

`func (o *SpaceSinglePermissions) GetMeta() OptionalFieldMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SpaceSinglePermissions) GetMetaOk() (*OptionalFieldMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SpaceSinglePermissions) SetMeta(v OptionalFieldMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SpaceSinglePermissions) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *SpaceSinglePermissions) GetLinks() OptionalFieldLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SpaceSinglePermissions) GetLinksOk() (*OptionalFieldLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SpaceSinglePermissions) SetLinks(v OptionalFieldLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SpaceSinglePermissions) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MultiEntityResultSpacePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]SpacePermission**](SpacePermission.md) |  | [optional] 
**Links** | Pointer to [**MultiEntityLinks**](MultiEntityLinks.md) |  | [optional] 

## Methods

### NewMultiEntityResultSpacePermission

`func NewMultiEntityResultSpacePermission() *MultiEntityResultSpacePermission`

NewMultiEntityResultSpacePermission instantiates a new MultiEntityResultSpacePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiEntityResultSpacePermissionWithDefaults

`func NewMultiEntityResultSpacePermissionWithDefaults() *MultiEntityResultSpacePermission`

NewMultiEntityResultSpacePermissionWithDefaults instantiates a new MultiEntityResultSpacePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *MultiEntityResultSpacePermission) GetResults() []SpacePermission`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *MultiEntityResultSpacePermission) GetResultsOk() (*[]SpacePermission, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *MultiEntityResultSpacePermission) SetResults(v []SpacePermission)`

SetResults sets Results field to given value.

### HasResults

`func (o *MultiEntityResultSpacePermission) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetLinks

`func (o *MultiEntityResultSpacePermission) GetLinks() MultiEntityLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MultiEntityResultSpacePermission) GetLinksOk() (*MultiEntityLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MultiEntityResultSpacePermission) SetLinks(v MultiEntityLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MultiEntityResultSpacePermission) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



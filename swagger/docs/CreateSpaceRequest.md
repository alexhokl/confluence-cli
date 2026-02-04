# CreateSpaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the space to be created. | 
**Key** | Pointer to **string** | The key for the new space. See [Space Keys](https://support.atlassian.com/confluence-cloud/docs/create-a-space/). If the key property is not provided, the alias property is required to be used instead. | [optional] 
**Alias** | Pointer to **string** | This field will be used as the new identifier for the space in confluence page URLs. If the alias property is not provided, the key property is required to be used instead. Maximum 255 alphanumeric characters in length. | [optional] 
**Description** | Pointer to [**CreateSpaceRequestDescription**](CreateSpaceRequestDescription.md) |  | [optional] 
**RoleAssignments** | Pointer to [**[]CreateSpaceRequestRoleAssignmentsInner**](CreateSpaceRequestRoleAssignmentsInner.md) |  | [optional] 
**CopySpaceAccessConfiguration** | Pointer to **int32** | The id of the space to copy the space access configuration from. | [optional] 
**CreatePrivateSpace** | Pointer to **bool** | Whether to create the space as private. | [optional] 
**TemplateKey** | Pointer to **string** | The key of the template to use. | [optional] 

## Methods

### NewCreateSpaceRequest

`func NewCreateSpaceRequest(name string, ) *CreateSpaceRequest`

NewCreateSpaceRequest instantiates a new CreateSpaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSpaceRequestWithDefaults

`func NewCreateSpaceRequestWithDefaults() *CreateSpaceRequest`

NewCreateSpaceRequestWithDefaults instantiates a new CreateSpaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSpaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSpaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSpaceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *CreateSpaceRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateSpaceRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateSpaceRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateSpaceRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetAlias

`func (o *CreateSpaceRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *CreateSpaceRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *CreateSpaceRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *CreateSpaceRequest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetDescription

`func (o *CreateSpaceRequest) GetDescription() CreateSpaceRequestDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSpaceRequest) GetDescriptionOk() (*CreateSpaceRequestDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSpaceRequest) SetDescription(v CreateSpaceRequestDescription)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSpaceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRoleAssignments

`func (o *CreateSpaceRequest) GetRoleAssignments() []CreateSpaceRequestRoleAssignmentsInner`

GetRoleAssignments returns the RoleAssignments field if non-nil, zero value otherwise.

### GetRoleAssignmentsOk

`func (o *CreateSpaceRequest) GetRoleAssignmentsOk() (*[]CreateSpaceRequestRoleAssignmentsInner, bool)`

GetRoleAssignmentsOk returns a tuple with the RoleAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAssignments

`func (o *CreateSpaceRequest) SetRoleAssignments(v []CreateSpaceRequestRoleAssignmentsInner)`

SetRoleAssignments sets RoleAssignments field to given value.

### HasRoleAssignments

`func (o *CreateSpaceRequest) HasRoleAssignments() bool`

HasRoleAssignments returns a boolean if a field has been set.

### GetCopySpaceAccessConfiguration

`func (o *CreateSpaceRequest) GetCopySpaceAccessConfiguration() int32`

GetCopySpaceAccessConfiguration returns the CopySpaceAccessConfiguration field if non-nil, zero value otherwise.

### GetCopySpaceAccessConfigurationOk

`func (o *CreateSpaceRequest) GetCopySpaceAccessConfigurationOk() (*int32, bool)`

GetCopySpaceAccessConfigurationOk returns a tuple with the CopySpaceAccessConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopySpaceAccessConfiguration

`func (o *CreateSpaceRequest) SetCopySpaceAccessConfiguration(v int32)`

SetCopySpaceAccessConfiguration sets CopySpaceAccessConfiguration field to given value.

### HasCopySpaceAccessConfiguration

`func (o *CreateSpaceRequest) HasCopySpaceAccessConfiguration() bool`

HasCopySpaceAccessConfiguration returns a boolean if a field has been set.

### GetCreatePrivateSpace

`func (o *CreateSpaceRequest) GetCreatePrivateSpace() bool`

GetCreatePrivateSpace returns the CreatePrivateSpace field if non-nil, zero value otherwise.

### GetCreatePrivateSpaceOk

`func (o *CreateSpaceRequest) GetCreatePrivateSpaceOk() (*bool, bool)`

GetCreatePrivateSpaceOk returns a tuple with the CreatePrivateSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatePrivateSpace

`func (o *CreateSpaceRequest) SetCreatePrivateSpace(v bool)`

SetCreatePrivateSpace sets CreatePrivateSpace field to given value.

### HasCreatePrivateSpace

`func (o *CreateSpaceRequest) HasCreatePrivateSpace() bool`

HasCreatePrivateSpace returns a boolean if a field has been set.

### GetTemplateKey

`func (o *CreateSpaceRequest) GetTemplateKey() string`

GetTemplateKey returns the TemplateKey field if non-nil, zero value otherwise.

### GetTemplateKeyOk

`func (o *CreateSpaceRequest) GetTemplateKeyOk() (*string, bool)`

GetTemplateKeyOk returns a tuple with the TemplateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateKey

`func (o *CreateSpaceRequest) SetTemplateKey(v string)`

SetTemplateKey sets TemplateKey field to given value.

### HasTemplateKey

`func (o *CreateSpaceRequest) HasTemplateKey() bool`

HasTemplateKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



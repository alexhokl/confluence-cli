# GetSpaceRolesById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The identifier for the space role. | [optional] 
**Type** | Pointer to [**RoleType**](RoleType.md) |  | [optional] 
**Name** | Pointer to **string** | The name for the space role. | [optional] 
**Description** | Pointer to **string** | The description for the space role’s usage. | [optional] 
**SpacePermissions** | Pointer to **[]string** | The space permissions the space role is comprised of. | [optional] 
**Links** | Pointer to [**GetAttachmentById200ResponseAllOfLinks**](GetAttachmentById200ResponseAllOfLinks.md) |  | [optional] 

## Methods

### NewGetSpaceRolesById200Response

`func NewGetSpaceRolesById200Response() *GetSpaceRolesById200Response`

NewGetSpaceRolesById200Response instantiates a new GetSpaceRolesById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSpaceRolesById200ResponseWithDefaults

`func NewGetSpaceRolesById200ResponseWithDefaults() *GetSpaceRolesById200Response`

NewGetSpaceRolesById200ResponseWithDefaults instantiates a new GetSpaceRolesById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSpaceRolesById200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSpaceRolesById200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSpaceRolesById200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetSpaceRolesById200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetSpaceRolesById200Response) GetType() RoleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetSpaceRolesById200Response) GetTypeOk() (*RoleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetSpaceRolesById200Response) SetType(v RoleType)`

SetType sets Type field to given value.

### HasType

`func (o *GetSpaceRolesById200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetSpaceRolesById200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSpaceRolesById200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSpaceRolesById200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSpaceRolesById200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetSpaceRolesById200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetSpaceRolesById200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetSpaceRolesById200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetSpaceRolesById200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSpacePermissions

`func (o *GetSpaceRolesById200Response) GetSpacePermissions() []string`

GetSpacePermissions returns the SpacePermissions field if non-nil, zero value otherwise.

### GetSpacePermissionsOk

`func (o *GetSpaceRolesById200Response) GetSpacePermissionsOk() (*[]string, bool)`

GetSpacePermissionsOk returns a tuple with the SpacePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpacePermissions

`func (o *GetSpaceRolesById200Response) SetSpacePermissions(v []string)`

SetSpacePermissions sets SpacePermissions field to given value.

### HasSpacePermissions

`func (o *GetSpaceRolesById200Response) HasSpacePermissions() bool`

HasSpacePermissions returns a boolean if a field has been set.

### GetLinks

`func (o *GetSpaceRolesById200Response) GetLinks() GetAttachmentById200ResponseAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetSpaceRolesById200Response) GetLinksOk() (*GetAttachmentById200ResponseAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetSpaceRolesById200Response) SetLinks(v GetAttachmentById200ResponseAllOfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetSpaceRolesById200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



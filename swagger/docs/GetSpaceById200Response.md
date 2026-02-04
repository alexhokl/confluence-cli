# GetSpaceById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the space. | [optional] 
**Key** | Pointer to **string** | Key of the space. | [optional] 
**Name** | Pointer to **string** | Name of the space. | [optional] 
**Type** | Pointer to [**SpaceType**](SpaceType.md) |  | [optional] 
**Status** | Pointer to [**SpaceStatus**](SpaceStatus.md) |  | [optional] 
**AuthorId** | Pointer to **string** | The account ID of the user who created this space originally. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time when the space was created. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 
**HomepageId** | Pointer to **string** | ID of the space&#39;s homepage. | [optional] 
**Description** | Pointer to [**SpaceDescription**](SpaceDescription.md) |  | [optional] 
**Icon** | Pointer to [**SpaceIcon**](SpaceIcon.md) |  | [optional] 
**Labels** | Pointer to [**AttachmentSingleLabels**](AttachmentSingleLabels.md) |  | [optional] 
**Properties** | Pointer to [**SpaceSingleProperties**](SpaceSingleProperties.md) |  | [optional] 
**Operations** | Pointer to [**AttachmentSingleOperations**](AttachmentSingleOperations.md) |  | [optional] 
**Permissions** | Pointer to [**SpaceSinglePermissions**](SpaceSinglePermissions.md) |  | [optional] 
**Links** | Pointer to [**GetAttachmentById200ResponseAllOfLinks**](GetAttachmentById200ResponseAllOfLinks.md) |  | [optional] 

## Methods

### NewGetSpaceById200Response

`func NewGetSpaceById200Response() *GetSpaceById200Response`

NewGetSpaceById200Response instantiates a new GetSpaceById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSpaceById200ResponseWithDefaults

`func NewGetSpaceById200ResponseWithDefaults() *GetSpaceById200Response`

NewGetSpaceById200ResponseWithDefaults instantiates a new GetSpaceById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSpaceById200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSpaceById200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSpaceById200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetSpaceById200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *GetSpaceById200Response) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetSpaceById200Response) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetSpaceById200Response) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GetSpaceById200Response) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *GetSpaceById200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSpaceById200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSpaceById200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSpaceById200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *GetSpaceById200Response) GetType() SpaceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetSpaceById200Response) GetTypeOk() (*SpaceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetSpaceById200Response) SetType(v SpaceType)`

SetType sets Type field to given value.

### HasType

`func (o *GetSpaceById200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *GetSpaceById200Response) GetStatus() SpaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSpaceById200Response) GetStatusOk() (*SpaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSpaceById200Response) SetStatus(v SpaceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetSpaceById200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAuthorId

`func (o *GetSpaceById200Response) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *GetSpaceById200Response) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *GetSpaceById200Response) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *GetSpaceById200Response) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetSpaceById200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetSpaceById200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetSpaceById200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetSpaceById200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHomepageId

`func (o *GetSpaceById200Response) GetHomepageId() string`

GetHomepageId returns the HomepageId field if non-nil, zero value otherwise.

### GetHomepageIdOk

`func (o *GetSpaceById200Response) GetHomepageIdOk() (*string, bool)`

GetHomepageIdOk returns a tuple with the HomepageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepageId

`func (o *GetSpaceById200Response) SetHomepageId(v string)`

SetHomepageId sets HomepageId field to given value.

### HasHomepageId

`func (o *GetSpaceById200Response) HasHomepageId() bool`

HasHomepageId returns a boolean if a field has been set.

### GetDescription

`func (o *GetSpaceById200Response) GetDescription() SpaceDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetSpaceById200Response) GetDescriptionOk() (*SpaceDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetSpaceById200Response) SetDescription(v SpaceDescription)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetSpaceById200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *GetSpaceById200Response) GetIcon() SpaceIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *GetSpaceById200Response) GetIconOk() (*SpaceIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *GetSpaceById200Response) SetIcon(v SpaceIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *GetSpaceById200Response) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetLabels

`func (o *GetSpaceById200Response) GetLabels() AttachmentSingleLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetSpaceById200Response) GetLabelsOk() (*AttachmentSingleLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetSpaceById200Response) SetLabels(v AttachmentSingleLabels)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetSpaceById200Response) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetProperties

`func (o *GetSpaceById200Response) GetProperties() SpaceSingleProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *GetSpaceById200Response) GetPropertiesOk() (*SpaceSingleProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *GetSpaceById200Response) SetProperties(v SpaceSingleProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *GetSpaceById200Response) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetOperations

`func (o *GetSpaceById200Response) GetOperations() AttachmentSingleOperations`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *GetSpaceById200Response) GetOperationsOk() (*AttachmentSingleOperations, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *GetSpaceById200Response) SetOperations(v AttachmentSingleOperations)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *GetSpaceById200Response) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetPermissions

`func (o *GetSpaceById200Response) GetPermissions() SpaceSinglePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GetSpaceById200Response) GetPermissionsOk() (*SpaceSinglePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GetSpaceById200Response) SetPermissions(v SpaceSinglePermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GetSpaceById200Response) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetLinks

`func (o *GetSpaceById200Response) GetLinks() GetAttachmentById200ResponseAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetSpaceById200Response) GetLinksOk() (*GetAttachmentById200ResponseAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetSpaceById200Response) SetLinks(v GetAttachmentById200ResponseAllOfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetSpaceById200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



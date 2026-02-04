# SpaceSingle

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
**Links** | Pointer to [**SpaceLinks**](SpaceLinks.md) |  | [optional] 

## Methods

### NewSpaceSingle

`func NewSpaceSingle() *SpaceSingle`

NewSpaceSingle instantiates a new SpaceSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceSingleWithDefaults

`func NewSpaceSingleWithDefaults() *SpaceSingle`

NewSpaceSingleWithDefaults instantiates a new SpaceSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SpaceSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpaceSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpaceSingle) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SpaceSingle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *SpaceSingle) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SpaceSingle) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SpaceSingle) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SpaceSingle) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *SpaceSingle) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpaceSingle) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpaceSingle) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SpaceSingle) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SpaceSingle) GetType() SpaceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpaceSingle) GetTypeOk() (*SpaceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpaceSingle) SetType(v SpaceType)`

SetType sets Type field to given value.

### HasType

`func (o *SpaceSingle) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *SpaceSingle) GetStatus() SpaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SpaceSingle) GetStatusOk() (*SpaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SpaceSingle) SetStatus(v SpaceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SpaceSingle) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAuthorId

`func (o *SpaceSingle) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *SpaceSingle) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *SpaceSingle) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *SpaceSingle) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SpaceSingle) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SpaceSingle) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SpaceSingle) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SpaceSingle) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHomepageId

`func (o *SpaceSingle) GetHomepageId() string`

GetHomepageId returns the HomepageId field if non-nil, zero value otherwise.

### GetHomepageIdOk

`func (o *SpaceSingle) GetHomepageIdOk() (*string, bool)`

GetHomepageIdOk returns a tuple with the HomepageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepageId

`func (o *SpaceSingle) SetHomepageId(v string)`

SetHomepageId sets HomepageId field to given value.

### HasHomepageId

`func (o *SpaceSingle) HasHomepageId() bool`

HasHomepageId returns a boolean if a field has been set.

### GetDescription

`func (o *SpaceSingle) GetDescription() SpaceDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SpaceSingle) GetDescriptionOk() (*SpaceDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SpaceSingle) SetDescription(v SpaceDescription)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SpaceSingle) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *SpaceSingle) GetIcon() SpaceIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *SpaceSingle) GetIconOk() (*SpaceIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *SpaceSingle) SetIcon(v SpaceIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *SpaceSingle) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetLabels

`func (o *SpaceSingle) GetLabels() AttachmentSingleLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SpaceSingle) GetLabelsOk() (*AttachmentSingleLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SpaceSingle) SetLabels(v AttachmentSingleLabels)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SpaceSingle) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetProperties

`func (o *SpaceSingle) GetProperties() SpaceSingleProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SpaceSingle) GetPropertiesOk() (*SpaceSingleProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SpaceSingle) SetProperties(v SpaceSingleProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SpaceSingle) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetOperations

`func (o *SpaceSingle) GetOperations() AttachmentSingleOperations`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *SpaceSingle) GetOperationsOk() (*AttachmentSingleOperations, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *SpaceSingle) SetOperations(v AttachmentSingleOperations)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *SpaceSingle) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetPermissions

`func (o *SpaceSingle) GetPermissions() SpaceSinglePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SpaceSingle) GetPermissionsOk() (*SpaceSinglePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SpaceSingle) SetPermissions(v SpaceSinglePermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *SpaceSingle) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetLinks

`func (o *SpaceSingle) GetLinks() SpaceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SpaceSingle) GetLinksOk() (*SpaceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SpaceSingle) SetLinks(v SpaceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SpaceSingle) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateSpace201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the space. | [optional] 
**Key** | Pointer to **string** | Key of the space. | [optional] 
**Name** | Pointer to **string** | Name of the space. | [optional] 
**Type** | Pointer to [**SpaceType**](SpaceType.md) |  | [optional] 
**Status** | Pointer to [**SpaceStatus**](SpaceStatus.md) |  | [optional] 
**AuthorId** | Pointer to **string** | The account ID of the user who created this space originally. | [optional] 
**CurrentActiveAlias** | Pointer to **string** | Currently active alias for a Confluence space. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time when the space was created. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 
**HomepageId** | Pointer to **string** | ID of the space&#39;s homepage. | [optional] 
**Description** | Pointer to [**SpaceDescription**](SpaceDescription.md) |  | [optional] 
**Icon** | Pointer to [**SpaceIcon**](SpaceIcon.md) |  | [optional] 
**Links** | Pointer to [**GetAttachmentById200ResponseAllOfLinks**](GetAttachmentById200ResponseAllOfLinks.md) |  | [optional] 

## Methods

### NewCreateSpace201Response

`func NewCreateSpace201Response() *CreateSpace201Response`

NewCreateSpace201Response instantiates a new CreateSpace201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSpace201ResponseWithDefaults

`func NewCreateSpace201ResponseWithDefaults() *CreateSpace201Response`

NewCreateSpace201ResponseWithDefaults instantiates a new CreateSpace201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateSpace201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateSpace201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateSpace201Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateSpace201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *CreateSpace201Response) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateSpace201Response) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateSpace201Response) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateSpace201Response) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *CreateSpace201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSpace201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSpace201Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateSpace201Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *CreateSpace201Response) GetType() SpaceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateSpace201Response) GetTypeOk() (*SpaceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateSpace201Response) SetType(v SpaceType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateSpace201Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *CreateSpace201Response) GetStatus() SpaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateSpace201Response) GetStatusOk() (*SpaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateSpace201Response) SetStatus(v SpaceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateSpace201Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAuthorId

`func (o *CreateSpace201Response) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *CreateSpace201Response) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *CreateSpace201Response) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *CreateSpace201Response) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetCurrentActiveAlias

`func (o *CreateSpace201Response) GetCurrentActiveAlias() string`

GetCurrentActiveAlias returns the CurrentActiveAlias field if non-nil, zero value otherwise.

### GetCurrentActiveAliasOk

`func (o *CreateSpace201Response) GetCurrentActiveAliasOk() (*string, bool)`

GetCurrentActiveAliasOk returns a tuple with the CurrentActiveAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentActiveAlias

`func (o *CreateSpace201Response) SetCurrentActiveAlias(v string)`

SetCurrentActiveAlias sets CurrentActiveAlias field to given value.

### HasCurrentActiveAlias

`func (o *CreateSpace201Response) HasCurrentActiveAlias() bool`

HasCurrentActiveAlias returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateSpace201Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateSpace201Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateSpace201Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateSpace201Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHomepageId

`func (o *CreateSpace201Response) GetHomepageId() string`

GetHomepageId returns the HomepageId field if non-nil, zero value otherwise.

### GetHomepageIdOk

`func (o *CreateSpace201Response) GetHomepageIdOk() (*string, bool)`

GetHomepageIdOk returns a tuple with the HomepageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepageId

`func (o *CreateSpace201Response) SetHomepageId(v string)`

SetHomepageId sets HomepageId field to given value.

### HasHomepageId

`func (o *CreateSpace201Response) HasHomepageId() bool`

HasHomepageId returns a boolean if a field has been set.

### GetDescription

`func (o *CreateSpace201Response) GetDescription() SpaceDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSpace201Response) GetDescriptionOk() (*SpaceDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSpace201Response) SetDescription(v SpaceDescription)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSpace201Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *CreateSpace201Response) GetIcon() SpaceIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CreateSpace201Response) GetIconOk() (*SpaceIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CreateSpace201Response) SetIcon(v SpaceIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CreateSpace201Response) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetLinks

`func (o *CreateSpace201Response) GetLinks() GetAttachmentById200ResponseAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateSpace201Response) GetLinksOk() (*GetAttachmentById200ResponseAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateSpace201Response) SetLinks(v GetAttachmentById200ResponseAllOfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateSpace201Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



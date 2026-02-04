# SpaceBulk

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
**Links** | Pointer to [**SpaceLinks**](SpaceLinks.md) |  | [optional] 

## Methods

### NewSpaceBulk

`func NewSpaceBulk() *SpaceBulk`

NewSpaceBulk instantiates a new SpaceBulk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceBulkWithDefaults

`func NewSpaceBulkWithDefaults() *SpaceBulk`

NewSpaceBulkWithDefaults instantiates a new SpaceBulk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SpaceBulk) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpaceBulk) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpaceBulk) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SpaceBulk) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *SpaceBulk) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SpaceBulk) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SpaceBulk) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SpaceBulk) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *SpaceBulk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpaceBulk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpaceBulk) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SpaceBulk) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SpaceBulk) GetType() SpaceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpaceBulk) GetTypeOk() (*SpaceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpaceBulk) SetType(v SpaceType)`

SetType sets Type field to given value.

### HasType

`func (o *SpaceBulk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *SpaceBulk) GetStatus() SpaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SpaceBulk) GetStatusOk() (*SpaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SpaceBulk) SetStatus(v SpaceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SpaceBulk) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAuthorId

`func (o *SpaceBulk) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *SpaceBulk) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *SpaceBulk) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *SpaceBulk) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetCurrentActiveAlias

`func (o *SpaceBulk) GetCurrentActiveAlias() string`

GetCurrentActiveAlias returns the CurrentActiveAlias field if non-nil, zero value otherwise.

### GetCurrentActiveAliasOk

`func (o *SpaceBulk) GetCurrentActiveAliasOk() (*string, bool)`

GetCurrentActiveAliasOk returns a tuple with the CurrentActiveAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentActiveAlias

`func (o *SpaceBulk) SetCurrentActiveAlias(v string)`

SetCurrentActiveAlias sets CurrentActiveAlias field to given value.

### HasCurrentActiveAlias

`func (o *SpaceBulk) HasCurrentActiveAlias() bool`

HasCurrentActiveAlias returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SpaceBulk) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SpaceBulk) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SpaceBulk) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SpaceBulk) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHomepageId

`func (o *SpaceBulk) GetHomepageId() string`

GetHomepageId returns the HomepageId field if non-nil, zero value otherwise.

### GetHomepageIdOk

`func (o *SpaceBulk) GetHomepageIdOk() (*string, bool)`

GetHomepageIdOk returns a tuple with the HomepageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepageId

`func (o *SpaceBulk) SetHomepageId(v string)`

SetHomepageId sets HomepageId field to given value.

### HasHomepageId

`func (o *SpaceBulk) HasHomepageId() bool`

HasHomepageId returns a boolean if a field has been set.

### GetDescription

`func (o *SpaceBulk) GetDescription() SpaceDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SpaceBulk) GetDescriptionOk() (*SpaceDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SpaceBulk) SetDescription(v SpaceDescription)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SpaceBulk) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *SpaceBulk) GetIcon() SpaceIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *SpaceBulk) GetIconOk() (*SpaceIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *SpaceBulk) SetIcon(v SpaceIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *SpaceBulk) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetLinks

`func (o *SpaceBulk) GetLinks() SpaceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SpaceBulk) GetLinksOk() (*SpaceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SpaceBulk) SetLinks(v SpaceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SpaceBulk) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



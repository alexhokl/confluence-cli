# FolderSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the folder. | [optional] 
**Type** | Pointer to **string** | The content type of the object. | [optional] 
**Status** | Pointer to [**ContentStatus**](ContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the folder. | [optional] 
**ParentId** | Pointer to **string** | ID of the parent content, or null if there is no parent content. | [optional] 
**ParentType** | Pointer to [**ParentContentType**](ParentContentType.md) |  | [optional] 
**Position** | Pointer to **NullableInt32** | Position of the folder within the given parent page tree. | [optional] 
**AuthorId** | Pointer to **string** | The account ID of the user who created this folder. | [optional] 
**OwnerId** | Pointer to **string** | The account ID of the user who owns this folder. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time when the folder was created. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 
**SpaceId** | Pointer to **string** | ID of the space the folder is in. | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Links** | Pointer to [**FolderLinks**](FolderLinks.md) |  | [optional] 

## Methods

### NewFolderSingle

`func NewFolderSingle() *FolderSingle`

NewFolderSingle instantiates a new FolderSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderSingleWithDefaults

`func NewFolderSingleWithDefaults() *FolderSingle`

NewFolderSingleWithDefaults instantiates a new FolderSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FolderSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FolderSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FolderSingle) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FolderSingle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *FolderSingle) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FolderSingle) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FolderSingle) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FolderSingle) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *FolderSingle) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FolderSingle) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FolderSingle) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FolderSingle) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *FolderSingle) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FolderSingle) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FolderSingle) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FolderSingle) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetParentId

`func (o *FolderSingle) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *FolderSingle) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *FolderSingle) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *FolderSingle) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetParentType

`func (o *FolderSingle) GetParentType() ParentContentType`

GetParentType returns the ParentType field if non-nil, zero value otherwise.

### GetParentTypeOk

`func (o *FolderSingle) GetParentTypeOk() (*ParentContentType, bool)`

GetParentTypeOk returns a tuple with the ParentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentType

`func (o *FolderSingle) SetParentType(v ParentContentType)`

SetParentType sets ParentType field to given value.

### HasParentType

`func (o *FolderSingle) HasParentType() bool`

HasParentType returns a boolean if a field has been set.

### GetPosition

`func (o *FolderSingle) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FolderSingle) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FolderSingle) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *FolderSingle) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *FolderSingle) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *FolderSingle) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetAuthorId

`func (o *FolderSingle) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *FolderSingle) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *FolderSingle) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *FolderSingle) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetOwnerId

`func (o *FolderSingle) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *FolderSingle) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *FolderSingle) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *FolderSingle) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FolderSingle) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FolderSingle) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FolderSingle) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FolderSingle) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSpaceId

`func (o *FolderSingle) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *FolderSingle) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *FolderSingle) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *FolderSingle) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetVersion

`func (o *FolderSingle) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FolderSingle) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FolderSingle) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FolderSingle) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLinks

`func (o *FolderSingle) GetLinks() FolderLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FolderSingle) GetLinksOk() (*FolderLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FolderSingle) SetLinks(v FolderLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FolderSingle) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



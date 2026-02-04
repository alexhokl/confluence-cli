# DatabaseSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the database. | [optional] 
**Type** | Pointer to **string** | The content type of the object. | [optional] 
**Status** | Pointer to [**ContentStatus**](ContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the database. | [optional] 
**ParentId** | Pointer to **string** | ID of the parent content, or null if there is no parent content. | [optional] 
**ParentType** | Pointer to [**ParentContentType**](ParentContentType.md) |  | [optional] 
**Position** | Pointer to **NullableInt32** | Position of the database within the given parent page tree. | [optional] 
**AuthorId** | Pointer to **string** | The account ID of the user who created this database originally. | [optional] 
**OwnerId** | Pointer to **string** | The account ID of the user who owns this database. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time when the database was created. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 
**SpaceId** | Pointer to **string** | ID of the space the database is in. | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Links** | Pointer to [**DatabaseLinks**](DatabaseLinks.md) |  | [optional] 

## Methods

### NewDatabaseSingle

`func NewDatabaseSingle() *DatabaseSingle`

NewDatabaseSingle instantiates a new DatabaseSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseSingleWithDefaults

`func NewDatabaseSingleWithDefaults() *DatabaseSingle`

NewDatabaseSingleWithDefaults instantiates a new DatabaseSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatabaseSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseSingle) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DatabaseSingle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *DatabaseSingle) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatabaseSingle) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatabaseSingle) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DatabaseSingle) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *DatabaseSingle) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatabaseSingle) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatabaseSingle) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DatabaseSingle) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *DatabaseSingle) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DatabaseSingle) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DatabaseSingle) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DatabaseSingle) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetParentId

`func (o *DatabaseSingle) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *DatabaseSingle) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *DatabaseSingle) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *DatabaseSingle) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetParentType

`func (o *DatabaseSingle) GetParentType() ParentContentType`

GetParentType returns the ParentType field if non-nil, zero value otherwise.

### GetParentTypeOk

`func (o *DatabaseSingle) GetParentTypeOk() (*ParentContentType, bool)`

GetParentTypeOk returns a tuple with the ParentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentType

`func (o *DatabaseSingle) SetParentType(v ParentContentType)`

SetParentType sets ParentType field to given value.

### HasParentType

`func (o *DatabaseSingle) HasParentType() bool`

HasParentType returns a boolean if a field has been set.

### GetPosition

`func (o *DatabaseSingle) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *DatabaseSingle) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *DatabaseSingle) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *DatabaseSingle) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *DatabaseSingle) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *DatabaseSingle) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetAuthorId

`func (o *DatabaseSingle) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *DatabaseSingle) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *DatabaseSingle) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *DatabaseSingle) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetOwnerId

`func (o *DatabaseSingle) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *DatabaseSingle) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *DatabaseSingle) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *DatabaseSingle) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DatabaseSingle) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DatabaseSingle) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DatabaseSingle) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DatabaseSingle) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSpaceId

`func (o *DatabaseSingle) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *DatabaseSingle) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *DatabaseSingle) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *DatabaseSingle) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetVersion

`func (o *DatabaseSingle) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DatabaseSingle) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DatabaseSingle) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DatabaseSingle) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLinks

`func (o *DatabaseSingle) GetLinks() DatabaseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DatabaseSingle) GetLinksOk() (*DatabaseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DatabaseSingle) SetLinks(v DatabaseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DatabaseSingle) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



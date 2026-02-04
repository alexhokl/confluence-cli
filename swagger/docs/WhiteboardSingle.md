# WhiteboardSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the whiteboard. | [optional] 
**Type** | Pointer to **string** | The content type of the object. | [optional] 
**Status** | Pointer to [**ContentStatus**](ContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the whiteboard. | [optional] 
**ParentId** | Pointer to **string** | ID of the parent content, or null if there is no parent content. | [optional] 
**ParentType** | Pointer to [**ParentContentType**](ParentContentType.md) |  | [optional] 
**Position** | Pointer to **NullableInt32** | Position of the whiteboard within the given parent page tree. | [optional] 
**AuthorId** | Pointer to **string** | The account ID of the user who created this whiteboard originally. | [optional] 
**OwnerId** | Pointer to **string** | The account ID of the user who owns this whiteboard. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time when the whiteboard was created. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 
**SpaceId** | Pointer to **string** | ID of the space the whiteboard is in. | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Links** | Pointer to [**WhiteboardLinks**](WhiteboardLinks.md) |  | [optional] 

## Methods

### NewWhiteboardSingle

`func NewWhiteboardSingle() *WhiteboardSingle`

NewWhiteboardSingle instantiates a new WhiteboardSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhiteboardSingleWithDefaults

`func NewWhiteboardSingleWithDefaults() *WhiteboardSingle`

NewWhiteboardSingleWithDefaults instantiates a new WhiteboardSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WhiteboardSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhiteboardSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhiteboardSingle) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WhiteboardSingle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *WhiteboardSingle) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhiteboardSingle) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhiteboardSingle) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WhiteboardSingle) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *WhiteboardSingle) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WhiteboardSingle) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WhiteboardSingle) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WhiteboardSingle) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *WhiteboardSingle) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WhiteboardSingle) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WhiteboardSingle) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WhiteboardSingle) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetParentId

`func (o *WhiteboardSingle) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *WhiteboardSingle) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *WhiteboardSingle) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *WhiteboardSingle) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetParentType

`func (o *WhiteboardSingle) GetParentType() ParentContentType`

GetParentType returns the ParentType field if non-nil, zero value otherwise.

### GetParentTypeOk

`func (o *WhiteboardSingle) GetParentTypeOk() (*ParentContentType, bool)`

GetParentTypeOk returns a tuple with the ParentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentType

`func (o *WhiteboardSingle) SetParentType(v ParentContentType)`

SetParentType sets ParentType field to given value.

### HasParentType

`func (o *WhiteboardSingle) HasParentType() bool`

HasParentType returns a boolean if a field has been set.

### GetPosition

`func (o *WhiteboardSingle) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *WhiteboardSingle) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *WhiteboardSingle) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *WhiteboardSingle) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *WhiteboardSingle) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *WhiteboardSingle) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetAuthorId

`func (o *WhiteboardSingle) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *WhiteboardSingle) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *WhiteboardSingle) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *WhiteboardSingle) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetOwnerId

`func (o *WhiteboardSingle) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *WhiteboardSingle) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *WhiteboardSingle) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *WhiteboardSingle) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WhiteboardSingle) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WhiteboardSingle) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WhiteboardSingle) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WhiteboardSingle) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSpaceId

`func (o *WhiteboardSingle) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *WhiteboardSingle) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *WhiteboardSingle) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *WhiteboardSingle) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetVersion

`func (o *WhiteboardSingle) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WhiteboardSingle) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WhiteboardSingle) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WhiteboardSingle) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLinks

`func (o *WhiteboardSingle) GetLinks() WhiteboardLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WhiteboardSingle) GetLinksOk() (*WhiteboardLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WhiteboardSingle) SetLinks(v WhiteboardLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WhiteboardSingle) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



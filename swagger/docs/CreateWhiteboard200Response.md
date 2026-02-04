# CreateWhiteboard200Response

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
**Links** | Pointer to [**GetAttachmentById200ResponseAllOfLinks**](GetAttachmentById200ResponseAllOfLinks.md) |  | [optional] 

## Methods

### NewCreateWhiteboard200Response

`func NewCreateWhiteboard200Response() *CreateWhiteboard200Response`

NewCreateWhiteboard200Response instantiates a new CreateWhiteboard200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWhiteboard200ResponseWithDefaults

`func NewCreateWhiteboard200ResponseWithDefaults() *CreateWhiteboard200Response`

NewCreateWhiteboard200ResponseWithDefaults instantiates a new CreateWhiteboard200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateWhiteboard200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateWhiteboard200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateWhiteboard200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateWhiteboard200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CreateWhiteboard200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateWhiteboard200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateWhiteboard200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateWhiteboard200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *CreateWhiteboard200Response) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateWhiteboard200Response) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateWhiteboard200Response) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateWhiteboard200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *CreateWhiteboard200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateWhiteboard200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateWhiteboard200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateWhiteboard200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetParentId

`func (o *CreateWhiteboard200Response) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateWhiteboard200Response) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateWhiteboard200Response) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateWhiteboard200Response) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetParentType

`func (o *CreateWhiteboard200Response) GetParentType() ParentContentType`

GetParentType returns the ParentType field if non-nil, zero value otherwise.

### GetParentTypeOk

`func (o *CreateWhiteboard200Response) GetParentTypeOk() (*ParentContentType, bool)`

GetParentTypeOk returns a tuple with the ParentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentType

`func (o *CreateWhiteboard200Response) SetParentType(v ParentContentType)`

SetParentType sets ParentType field to given value.

### HasParentType

`func (o *CreateWhiteboard200Response) HasParentType() bool`

HasParentType returns a boolean if a field has been set.

### GetPosition

`func (o *CreateWhiteboard200Response) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CreateWhiteboard200Response) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CreateWhiteboard200Response) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *CreateWhiteboard200Response) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *CreateWhiteboard200Response) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *CreateWhiteboard200Response) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetAuthorId

`func (o *CreateWhiteboard200Response) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *CreateWhiteboard200Response) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *CreateWhiteboard200Response) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *CreateWhiteboard200Response) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetOwnerId

`func (o *CreateWhiteboard200Response) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *CreateWhiteboard200Response) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *CreateWhiteboard200Response) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *CreateWhiteboard200Response) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateWhiteboard200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateWhiteboard200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateWhiteboard200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateWhiteboard200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSpaceId

`func (o *CreateWhiteboard200Response) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *CreateWhiteboard200Response) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *CreateWhiteboard200Response) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *CreateWhiteboard200Response) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetVersion

`func (o *CreateWhiteboard200Response) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateWhiteboard200Response) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateWhiteboard200Response) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreateWhiteboard200Response) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLinks

`func (o *CreateWhiteboard200Response) GetLinks() GetAttachmentById200ResponseAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateWhiteboard200Response) GetLinksOk() (*GetAttachmentById200ResponseAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateWhiteboard200Response) SetLinks(v GetAttachmentById200ResponseAllOfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateWhiteboard200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



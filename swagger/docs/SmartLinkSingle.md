# SmartLinkSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the Smart Link in the content tree. | [optional] 
**Type** | Pointer to **string** | The content type of the object. | [optional] 
**Status** | Pointer to [**ContentStatus**](ContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the Smart Link in the content tree. | [optional] 
**ParentId** | Pointer to **string** | ID of the parent content, or null if there is no parent content. | [optional] 
**ParentType** | Pointer to [**ParentContentType**](ParentContentType.md) |  | [optional] 
**Position** | Pointer to **NullableInt32** | Position of the Smart Link within the given parent page tree. | [optional] 
**AuthorId** | Pointer to **string** | The account ID of the user who created this Smart Link in the content tree originally. | [optional] 
**OwnerId** | Pointer to **string** | The account ID of the user who owns this Smart Link in the content tree. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time when the Smart Link in the content tree was created. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 
**EmbedUrl** | Pointer to **string** | The embedded URL of the Smart Link. If the Smart Link does not have an embedded URL, this property will not be included in the response. | [optional] 
**SpaceId** | Pointer to **string** | ID of the space the Smart Link is in. | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Links** | Pointer to [**SmartLinkLinks**](SmartLinkLinks.md) |  | [optional] 

## Methods

### NewSmartLinkSingle

`func NewSmartLinkSingle() *SmartLinkSingle`

NewSmartLinkSingle instantiates a new SmartLinkSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartLinkSingleWithDefaults

`func NewSmartLinkSingleWithDefaults() *SmartLinkSingle`

NewSmartLinkSingleWithDefaults instantiates a new SmartLinkSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SmartLinkSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SmartLinkSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SmartLinkSingle) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SmartLinkSingle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SmartLinkSingle) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SmartLinkSingle) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SmartLinkSingle) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SmartLinkSingle) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *SmartLinkSingle) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SmartLinkSingle) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SmartLinkSingle) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SmartLinkSingle) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *SmartLinkSingle) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SmartLinkSingle) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SmartLinkSingle) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SmartLinkSingle) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetParentId

`func (o *SmartLinkSingle) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *SmartLinkSingle) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *SmartLinkSingle) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *SmartLinkSingle) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetParentType

`func (o *SmartLinkSingle) GetParentType() ParentContentType`

GetParentType returns the ParentType field if non-nil, zero value otherwise.

### GetParentTypeOk

`func (o *SmartLinkSingle) GetParentTypeOk() (*ParentContentType, bool)`

GetParentTypeOk returns a tuple with the ParentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentType

`func (o *SmartLinkSingle) SetParentType(v ParentContentType)`

SetParentType sets ParentType field to given value.

### HasParentType

`func (o *SmartLinkSingle) HasParentType() bool`

HasParentType returns a boolean if a field has been set.

### GetPosition

`func (o *SmartLinkSingle) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *SmartLinkSingle) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *SmartLinkSingle) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *SmartLinkSingle) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *SmartLinkSingle) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *SmartLinkSingle) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetAuthorId

`func (o *SmartLinkSingle) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *SmartLinkSingle) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *SmartLinkSingle) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *SmartLinkSingle) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetOwnerId

`func (o *SmartLinkSingle) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *SmartLinkSingle) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *SmartLinkSingle) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *SmartLinkSingle) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SmartLinkSingle) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SmartLinkSingle) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SmartLinkSingle) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SmartLinkSingle) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmbedUrl

`func (o *SmartLinkSingle) GetEmbedUrl() string`

GetEmbedUrl returns the EmbedUrl field if non-nil, zero value otherwise.

### GetEmbedUrlOk

`func (o *SmartLinkSingle) GetEmbedUrlOk() (*string, bool)`

GetEmbedUrlOk returns a tuple with the EmbedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedUrl

`func (o *SmartLinkSingle) SetEmbedUrl(v string)`

SetEmbedUrl sets EmbedUrl field to given value.

### HasEmbedUrl

`func (o *SmartLinkSingle) HasEmbedUrl() bool`

HasEmbedUrl returns a boolean if a field has been set.

### GetSpaceId

`func (o *SmartLinkSingle) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *SmartLinkSingle) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *SmartLinkSingle) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *SmartLinkSingle) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetVersion

`func (o *SmartLinkSingle) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SmartLinkSingle) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SmartLinkSingle) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SmartLinkSingle) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLinks

`func (o *SmartLinkSingle) GetLinks() SmartLinkLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SmartLinkSingle) GetLinksOk() (*SmartLinkLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SmartLinkSingle) SetLinks(v SmartLinkLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SmartLinkSingle) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



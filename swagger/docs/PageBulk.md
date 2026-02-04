# PageBulk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the page. | [optional] 
**Status** | Pointer to [**ContentStatus**](ContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the page. | [optional] 
**SpaceId** | Pointer to **string** | ID of the space the page is in. | [optional] 
**ParentId** | Pointer to **string** | ID of the parent page, or null if there is no parent page. | [optional] 
**ParentType** | Pointer to [**ParentContentType**](ParentContentType.md) |  | [optional] 
**Position** | Pointer to **NullableInt32** | Position of child page within the given parent page tree. | [optional] 
**AuthorId** | Pointer to **string** | The account ID of the user who created this page originally. | [optional] 
**OwnerId** | Pointer to **NullableString** | The account ID of the user who owns this page. | [optional] 
**LastOwnerId** | Pointer to **NullableString** | The account ID of the user who owned this page previously, or null if there is no previous owner. | [optional] 
**Subtype** | Pointer to **NullableString** | The subtype of the page. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time when the page was created. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Body** | Pointer to [**BodyBulk**](BodyBulk.md) |  | [optional] 
**Links** | Pointer to [**AbstractPageLinks**](AbstractPageLinks.md) |  | [optional] 

## Methods

### NewPageBulk

`func NewPageBulk() *PageBulk`

NewPageBulk instantiates a new PageBulk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageBulkWithDefaults

`func NewPageBulkWithDefaults() *PageBulk`

NewPageBulkWithDefaults instantiates a new PageBulk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PageBulk) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PageBulk) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PageBulk) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PageBulk) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *PageBulk) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PageBulk) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PageBulk) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PageBulk) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *PageBulk) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PageBulk) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PageBulk) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PageBulk) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSpaceId

`func (o *PageBulk) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *PageBulk) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *PageBulk) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *PageBulk) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetParentId

`func (o *PageBulk) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *PageBulk) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *PageBulk) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *PageBulk) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetParentType

`func (o *PageBulk) GetParentType() ParentContentType`

GetParentType returns the ParentType field if non-nil, zero value otherwise.

### GetParentTypeOk

`func (o *PageBulk) GetParentTypeOk() (*ParentContentType, bool)`

GetParentTypeOk returns a tuple with the ParentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentType

`func (o *PageBulk) SetParentType(v ParentContentType)`

SetParentType sets ParentType field to given value.

### HasParentType

`func (o *PageBulk) HasParentType() bool`

HasParentType returns a boolean if a field has been set.

### GetPosition

`func (o *PageBulk) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PageBulk) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PageBulk) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PageBulk) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *PageBulk) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *PageBulk) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetAuthorId

`func (o *PageBulk) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *PageBulk) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *PageBulk) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *PageBulk) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetOwnerId

`func (o *PageBulk) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *PageBulk) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *PageBulk) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *PageBulk) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *PageBulk) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *PageBulk) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetLastOwnerId

`func (o *PageBulk) GetLastOwnerId() string`

GetLastOwnerId returns the LastOwnerId field if non-nil, zero value otherwise.

### GetLastOwnerIdOk

`func (o *PageBulk) GetLastOwnerIdOk() (*string, bool)`

GetLastOwnerIdOk returns a tuple with the LastOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOwnerId

`func (o *PageBulk) SetLastOwnerId(v string)`

SetLastOwnerId sets LastOwnerId field to given value.

### HasLastOwnerId

`func (o *PageBulk) HasLastOwnerId() bool`

HasLastOwnerId returns a boolean if a field has been set.

### SetLastOwnerIdNil

`func (o *PageBulk) SetLastOwnerIdNil(b bool)`

 SetLastOwnerIdNil sets the value for LastOwnerId to be an explicit nil

### UnsetLastOwnerId
`func (o *PageBulk) UnsetLastOwnerId()`

UnsetLastOwnerId ensures that no value is present for LastOwnerId, not even an explicit nil
### GetSubtype

`func (o *PageBulk) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *PageBulk) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *PageBulk) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *PageBulk) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### SetSubtypeNil

`func (o *PageBulk) SetSubtypeNil(b bool)`

 SetSubtypeNil sets the value for Subtype to be an explicit nil

### UnsetSubtype
`func (o *PageBulk) UnsetSubtype()`

UnsetSubtype ensures that no value is present for Subtype, not even an explicit nil
### GetCreatedAt

`func (o *PageBulk) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PageBulk) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PageBulk) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PageBulk) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *PageBulk) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PageBulk) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PageBulk) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PageBulk) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBody

`func (o *PageBulk) GetBody() BodyBulk`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PageBulk) GetBodyOk() (*BodyBulk, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PageBulk) SetBody(v BodyBulk)`

SetBody sets Body field to given value.

### HasBody

`func (o *PageBulk) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetLinks

`func (o *PageBulk) GetLinks() AbstractPageLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PageBulk) GetLinksOk() (*AbstractPageLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PageBulk) SetLinks(v AbstractPageLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PageBulk) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



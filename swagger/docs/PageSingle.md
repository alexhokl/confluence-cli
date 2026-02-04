# PageSingle

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
**CreatedAt** | Pointer to **time.Time** | Date and time when the page was created. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Body** | Pointer to [**BodySingle**](BodySingle.md) |  | [optional] 
**Labels** | Pointer to [**AttachmentSingleLabels**](AttachmentSingleLabels.md) |  | [optional] 
**Properties** | Pointer to [**AttachmentSingleProperties**](AttachmentSingleProperties.md) |  | [optional] 
**Operations** | Pointer to [**AttachmentSingleOperations**](AttachmentSingleOperations.md) |  | [optional] 
**Likes** | Pointer to [**BlogPostSingleLikes**](BlogPostSingleLikes.md) |  | [optional] 
**Versions** | Pointer to [**AttachmentSingleVersions**](AttachmentSingleVersions.md) |  | [optional] 
**IsFavoritedByCurrentUser** | Pointer to **bool** | Whether the page has been favorited by the current user. | [optional] 
**Links** | Pointer to [**AbstractPageLinks**](AbstractPageLinks.md) |  | [optional] 

## Methods

### NewPageSingle

`func NewPageSingle() *PageSingle`

NewPageSingle instantiates a new PageSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageSingleWithDefaults

`func NewPageSingleWithDefaults() *PageSingle`

NewPageSingleWithDefaults instantiates a new PageSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PageSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PageSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PageSingle) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PageSingle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *PageSingle) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PageSingle) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PageSingle) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PageSingle) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *PageSingle) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PageSingle) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PageSingle) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PageSingle) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSpaceId

`func (o *PageSingle) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *PageSingle) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *PageSingle) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *PageSingle) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetParentId

`func (o *PageSingle) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *PageSingle) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *PageSingle) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *PageSingle) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetParentType

`func (o *PageSingle) GetParentType() ParentContentType`

GetParentType returns the ParentType field if non-nil, zero value otherwise.

### GetParentTypeOk

`func (o *PageSingle) GetParentTypeOk() (*ParentContentType, bool)`

GetParentTypeOk returns a tuple with the ParentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentType

`func (o *PageSingle) SetParentType(v ParentContentType)`

SetParentType sets ParentType field to given value.

### HasParentType

`func (o *PageSingle) HasParentType() bool`

HasParentType returns a boolean if a field has been set.

### GetPosition

`func (o *PageSingle) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PageSingle) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PageSingle) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PageSingle) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *PageSingle) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *PageSingle) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetAuthorId

`func (o *PageSingle) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *PageSingle) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *PageSingle) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *PageSingle) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetOwnerId

`func (o *PageSingle) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *PageSingle) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *PageSingle) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *PageSingle) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *PageSingle) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *PageSingle) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetLastOwnerId

`func (o *PageSingle) GetLastOwnerId() string`

GetLastOwnerId returns the LastOwnerId field if non-nil, zero value otherwise.

### GetLastOwnerIdOk

`func (o *PageSingle) GetLastOwnerIdOk() (*string, bool)`

GetLastOwnerIdOk returns a tuple with the LastOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOwnerId

`func (o *PageSingle) SetLastOwnerId(v string)`

SetLastOwnerId sets LastOwnerId field to given value.

### HasLastOwnerId

`func (o *PageSingle) HasLastOwnerId() bool`

HasLastOwnerId returns a boolean if a field has been set.

### SetLastOwnerIdNil

`func (o *PageSingle) SetLastOwnerIdNil(b bool)`

 SetLastOwnerIdNil sets the value for LastOwnerId to be an explicit nil

### UnsetLastOwnerId
`func (o *PageSingle) UnsetLastOwnerId()`

UnsetLastOwnerId ensures that no value is present for LastOwnerId, not even an explicit nil
### GetCreatedAt

`func (o *PageSingle) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PageSingle) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PageSingle) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PageSingle) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *PageSingle) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PageSingle) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PageSingle) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PageSingle) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBody

`func (o *PageSingle) GetBody() BodySingle`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PageSingle) GetBodyOk() (*BodySingle, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PageSingle) SetBody(v BodySingle)`

SetBody sets Body field to given value.

### HasBody

`func (o *PageSingle) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetLabels

`func (o *PageSingle) GetLabels() AttachmentSingleLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PageSingle) GetLabelsOk() (*AttachmentSingleLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PageSingle) SetLabels(v AttachmentSingleLabels)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PageSingle) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetProperties

`func (o *PageSingle) GetProperties() AttachmentSingleProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PageSingle) GetPropertiesOk() (*AttachmentSingleProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PageSingle) SetProperties(v AttachmentSingleProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PageSingle) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetOperations

`func (o *PageSingle) GetOperations() AttachmentSingleOperations`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *PageSingle) GetOperationsOk() (*AttachmentSingleOperations, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *PageSingle) SetOperations(v AttachmentSingleOperations)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *PageSingle) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetLikes

`func (o *PageSingle) GetLikes() BlogPostSingleLikes`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *PageSingle) GetLikesOk() (*BlogPostSingleLikes, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *PageSingle) SetLikes(v BlogPostSingleLikes)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *PageSingle) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetVersions

`func (o *PageSingle) GetVersions() AttachmentSingleVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *PageSingle) GetVersionsOk() (*AttachmentSingleVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *PageSingle) SetVersions(v AttachmentSingleVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *PageSingle) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetIsFavoritedByCurrentUser

`func (o *PageSingle) GetIsFavoritedByCurrentUser() bool`

GetIsFavoritedByCurrentUser returns the IsFavoritedByCurrentUser field if non-nil, zero value otherwise.

### GetIsFavoritedByCurrentUserOk

`func (o *PageSingle) GetIsFavoritedByCurrentUserOk() (*bool, bool)`

GetIsFavoritedByCurrentUserOk returns a tuple with the IsFavoritedByCurrentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavoritedByCurrentUser

`func (o *PageSingle) SetIsFavoritedByCurrentUser(v bool)`

SetIsFavoritedByCurrentUser sets IsFavoritedByCurrentUser field to given value.

### HasIsFavoritedByCurrentUser

`func (o *PageSingle) HasIsFavoritedByCurrentUser() bool`

HasIsFavoritedByCurrentUser returns a boolean if a field has been set.

### GetLinks

`func (o *PageSingle) GetLinks() AbstractPageLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PageSingle) GetLinksOk() (*AbstractPageLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PageSingle) SetLinks(v AbstractPageLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PageSingle) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



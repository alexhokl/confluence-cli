# CreatePage200Response

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
**Links** | Pointer to [**GetAttachmentById200ResponseAllOfLinks**](GetAttachmentById200ResponseAllOfLinks.md) |  | [optional] 

## Methods

### NewCreatePage200Response

`func NewCreatePage200Response() *CreatePage200Response`

NewCreatePage200Response instantiates a new CreatePage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePage200ResponseWithDefaults

`func NewCreatePage200ResponseWithDefaults() *CreatePage200Response`

NewCreatePage200ResponseWithDefaults instantiates a new CreatePage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreatePage200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreatePage200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreatePage200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreatePage200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *CreatePage200Response) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreatePage200Response) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreatePage200Response) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreatePage200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *CreatePage200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreatePage200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreatePage200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreatePage200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSpaceId

`func (o *CreatePage200Response) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *CreatePage200Response) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *CreatePage200Response) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *CreatePage200Response) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetParentId

`func (o *CreatePage200Response) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreatePage200Response) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreatePage200Response) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreatePage200Response) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetParentType

`func (o *CreatePage200Response) GetParentType() ParentContentType`

GetParentType returns the ParentType field if non-nil, zero value otherwise.

### GetParentTypeOk

`func (o *CreatePage200Response) GetParentTypeOk() (*ParentContentType, bool)`

GetParentTypeOk returns a tuple with the ParentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentType

`func (o *CreatePage200Response) SetParentType(v ParentContentType)`

SetParentType sets ParentType field to given value.

### HasParentType

`func (o *CreatePage200Response) HasParentType() bool`

HasParentType returns a boolean if a field has been set.

### GetPosition

`func (o *CreatePage200Response) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CreatePage200Response) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CreatePage200Response) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *CreatePage200Response) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *CreatePage200Response) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *CreatePage200Response) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetAuthorId

`func (o *CreatePage200Response) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *CreatePage200Response) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *CreatePage200Response) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *CreatePage200Response) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetOwnerId

`func (o *CreatePage200Response) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *CreatePage200Response) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *CreatePage200Response) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *CreatePage200Response) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *CreatePage200Response) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *CreatePage200Response) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetLastOwnerId

`func (o *CreatePage200Response) GetLastOwnerId() string`

GetLastOwnerId returns the LastOwnerId field if non-nil, zero value otherwise.

### GetLastOwnerIdOk

`func (o *CreatePage200Response) GetLastOwnerIdOk() (*string, bool)`

GetLastOwnerIdOk returns a tuple with the LastOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOwnerId

`func (o *CreatePage200Response) SetLastOwnerId(v string)`

SetLastOwnerId sets LastOwnerId field to given value.

### HasLastOwnerId

`func (o *CreatePage200Response) HasLastOwnerId() bool`

HasLastOwnerId returns a boolean if a field has been set.

### SetLastOwnerIdNil

`func (o *CreatePage200Response) SetLastOwnerIdNil(b bool)`

 SetLastOwnerIdNil sets the value for LastOwnerId to be an explicit nil

### UnsetLastOwnerId
`func (o *CreatePage200Response) UnsetLastOwnerId()`

UnsetLastOwnerId ensures that no value is present for LastOwnerId, not even an explicit nil
### GetCreatedAt

`func (o *CreatePage200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreatePage200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreatePage200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreatePage200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *CreatePage200Response) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreatePage200Response) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreatePage200Response) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreatePage200Response) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBody

`func (o *CreatePage200Response) GetBody() BodySingle`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreatePage200Response) GetBodyOk() (*BodySingle, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreatePage200Response) SetBody(v BodySingle)`

SetBody sets Body field to given value.

### HasBody

`func (o *CreatePage200Response) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetLabels

`func (o *CreatePage200Response) GetLabels() AttachmentSingleLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreatePage200Response) GetLabelsOk() (*AttachmentSingleLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreatePage200Response) SetLabels(v AttachmentSingleLabels)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreatePage200Response) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetProperties

`func (o *CreatePage200Response) GetProperties() AttachmentSingleProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CreatePage200Response) GetPropertiesOk() (*AttachmentSingleProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CreatePage200Response) SetProperties(v AttachmentSingleProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CreatePage200Response) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetOperations

`func (o *CreatePage200Response) GetOperations() AttachmentSingleOperations`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *CreatePage200Response) GetOperationsOk() (*AttachmentSingleOperations, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *CreatePage200Response) SetOperations(v AttachmentSingleOperations)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *CreatePage200Response) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetLikes

`func (o *CreatePage200Response) GetLikes() BlogPostSingleLikes`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *CreatePage200Response) GetLikesOk() (*BlogPostSingleLikes, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *CreatePage200Response) SetLikes(v BlogPostSingleLikes)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *CreatePage200Response) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetVersions

`func (o *CreatePage200Response) GetVersions() AttachmentSingleVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *CreatePage200Response) GetVersionsOk() (*AttachmentSingleVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *CreatePage200Response) SetVersions(v AttachmentSingleVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *CreatePage200Response) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetIsFavoritedByCurrentUser

`func (o *CreatePage200Response) GetIsFavoritedByCurrentUser() bool`

GetIsFavoritedByCurrentUser returns the IsFavoritedByCurrentUser field if non-nil, zero value otherwise.

### GetIsFavoritedByCurrentUserOk

`func (o *CreatePage200Response) GetIsFavoritedByCurrentUserOk() (*bool, bool)`

GetIsFavoritedByCurrentUserOk returns a tuple with the IsFavoritedByCurrentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavoritedByCurrentUser

`func (o *CreatePage200Response) SetIsFavoritedByCurrentUser(v bool)`

SetIsFavoritedByCurrentUser sets IsFavoritedByCurrentUser field to given value.

### HasIsFavoritedByCurrentUser

`func (o *CreatePage200Response) HasIsFavoritedByCurrentUser() bool`

HasIsFavoritedByCurrentUser returns a boolean if a field has been set.

### GetLinks

`func (o *CreatePage200Response) GetLinks() GetAttachmentById200ResponseAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreatePage200Response) GetLinksOk() (*GetAttachmentById200ResponseAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreatePage200Response) SetLinks(v GetAttachmentById200ResponseAllOfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreatePage200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



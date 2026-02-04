# CreateCustomContent201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the custom content. | [optional] 
**Type** | Pointer to **string** | The type of custom content. | [optional] 
**Status** | Pointer to [**ContentStatus**](ContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the custom content. | [optional] 
**SpaceId** | Pointer to **string** | ID of the space the custom content is in.  Note: This is always returned, regardless of if the custom content has a container that is a space. | [optional] 
**PageId** | Pointer to **string** | ID of the containing page.  Note: This is only returned if the custom content has a container that is a page. | [optional] 
**BlogPostId** | Pointer to **string** | ID of the containing blog post.  Note: This is only returned if the custom content has a container that is a blog post. | [optional] 
**CustomContentId** | Pointer to **string** | ID of the containing custom content.  Note: This is only returned if the custom content has a container that is custom content. | [optional] 
**AuthorId** | Pointer to **string** | The account ID of the user who created this custom content originally. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time when the custom content was created. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Labels** | Pointer to [**AttachmentSingleLabels**](AttachmentSingleLabels.md) |  | [optional] 
**Properties** | Pointer to [**AttachmentSingleProperties**](AttachmentSingleProperties.md) |  | [optional] 
**Operations** | Pointer to [**AttachmentSingleOperations**](AttachmentSingleOperations.md) |  | [optional] 
**Versions** | Pointer to [**AttachmentSingleVersions**](AttachmentSingleVersions.md) |  | [optional] 
**Body** | Pointer to [**CustomContentBodySingle**](CustomContentBodySingle.md) |  | [optional] 
**Links** | Pointer to [**GetAttachmentById200ResponseAllOfLinks**](GetAttachmentById200ResponseAllOfLinks.md) |  | [optional] 

## Methods

### NewCreateCustomContent201Response

`func NewCreateCustomContent201Response() *CreateCustomContent201Response`

NewCreateCustomContent201Response instantiates a new CreateCustomContent201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomContent201ResponseWithDefaults

`func NewCreateCustomContent201ResponseWithDefaults() *CreateCustomContent201Response`

NewCreateCustomContent201ResponseWithDefaults instantiates a new CreateCustomContent201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateCustomContent201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateCustomContent201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateCustomContent201Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateCustomContent201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CreateCustomContent201Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateCustomContent201Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateCustomContent201Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateCustomContent201Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *CreateCustomContent201Response) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateCustomContent201Response) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateCustomContent201Response) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateCustomContent201Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *CreateCustomContent201Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateCustomContent201Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateCustomContent201Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateCustomContent201Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSpaceId

`func (o *CreateCustomContent201Response) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *CreateCustomContent201Response) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *CreateCustomContent201Response) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *CreateCustomContent201Response) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetPageId

`func (o *CreateCustomContent201Response) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *CreateCustomContent201Response) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *CreateCustomContent201Response) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *CreateCustomContent201Response) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetBlogPostId

`func (o *CreateCustomContent201Response) GetBlogPostId() string`

GetBlogPostId returns the BlogPostId field if non-nil, zero value otherwise.

### GetBlogPostIdOk

`func (o *CreateCustomContent201Response) GetBlogPostIdOk() (*string, bool)`

GetBlogPostIdOk returns a tuple with the BlogPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostId

`func (o *CreateCustomContent201Response) SetBlogPostId(v string)`

SetBlogPostId sets BlogPostId field to given value.

### HasBlogPostId

`func (o *CreateCustomContent201Response) HasBlogPostId() bool`

HasBlogPostId returns a boolean if a field has been set.

### GetCustomContentId

`func (o *CreateCustomContent201Response) GetCustomContentId() string`

GetCustomContentId returns the CustomContentId field if non-nil, zero value otherwise.

### GetCustomContentIdOk

`func (o *CreateCustomContent201Response) GetCustomContentIdOk() (*string, bool)`

GetCustomContentIdOk returns a tuple with the CustomContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContentId

`func (o *CreateCustomContent201Response) SetCustomContentId(v string)`

SetCustomContentId sets CustomContentId field to given value.

### HasCustomContentId

`func (o *CreateCustomContent201Response) HasCustomContentId() bool`

HasCustomContentId returns a boolean if a field has been set.

### GetAuthorId

`func (o *CreateCustomContent201Response) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *CreateCustomContent201Response) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *CreateCustomContent201Response) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *CreateCustomContent201Response) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateCustomContent201Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateCustomContent201Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateCustomContent201Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateCustomContent201Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *CreateCustomContent201Response) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateCustomContent201Response) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateCustomContent201Response) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreateCustomContent201Response) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLabels

`func (o *CreateCustomContent201Response) GetLabels() AttachmentSingleLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreateCustomContent201Response) GetLabelsOk() (*AttachmentSingleLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreateCustomContent201Response) SetLabels(v AttachmentSingleLabels)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreateCustomContent201Response) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetProperties

`func (o *CreateCustomContent201Response) GetProperties() AttachmentSingleProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CreateCustomContent201Response) GetPropertiesOk() (*AttachmentSingleProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CreateCustomContent201Response) SetProperties(v AttachmentSingleProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CreateCustomContent201Response) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetOperations

`func (o *CreateCustomContent201Response) GetOperations() AttachmentSingleOperations`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *CreateCustomContent201Response) GetOperationsOk() (*AttachmentSingleOperations, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *CreateCustomContent201Response) SetOperations(v AttachmentSingleOperations)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *CreateCustomContent201Response) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetVersions

`func (o *CreateCustomContent201Response) GetVersions() AttachmentSingleVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *CreateCustomContent201Response) GetVersionsOk() (*AttachmentSingleVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *CreateCustomContent201Response) SetVersions(v AttachmentSingleVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *CreateCustomContent201Response) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetBody

`func (o *CreateCustomContent201Response) GetBody() CustomContentBodySingle`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreateCustomContent201Response) GetBodyOk() (*CustomContentBodySingle, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreateCustomContent201Response) SetBody(v CustomContentBodySingle)`

SetBody sets Body field to given value.

### HasBody

`func (o *CreateCustomContent201Response) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetLinks

`func (o *CreateCustomContent201Response) GetLinks() GetAttachmentById200ResponseAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateCustomContent201Response) GetLinksOk() (*GetAttachmentById200ResponseAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateCustomContent201Response) SetLinks(v GetAttachmentById200ResponseAllOfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateCustomContent201Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



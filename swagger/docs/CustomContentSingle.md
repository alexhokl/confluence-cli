# CustomContentSingle

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
**Links** | Pointer to [**CustomContentLinks**](CustomContentLinks.md) |  | [optional] 

## Methods

### NewCustomContentSingle

`func NewCustomContentSingle() *CustomContentSingle`

NewCustomContentSingle instantiates a new CustomContentSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomContentSingleWithDefaults

`func NewCustomContentSingleWithDefaults() *CustomContentSingle`

NewCustomContentSingleWithDefaults instantiates a new CustomContentSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomContentSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomContentSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomContentSingle) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomContentSingle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CustomContentSingle) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomContentSingle) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomContentSingle) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomContentSingle) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *CustomContentSingle) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomContentSingle) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomContentSingle) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomContentSingle) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *CustomContentSingle) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CustomContentSingle) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CustomContentSingle) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CustomContentSingle) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSpaceId

`func (o *CustomContentSingle) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *CustomContentSingle) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *CustomContentSingle) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *CustomContentSingle) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetPageId

`func (o *CustomContentSingle) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *CustomContentSingle) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *CustomContentSingle) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *CustomContentSingle) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetBlogPostId

`func (o *CustomContentSingle) GetBlogPostId() string`

GetBlogPostId returns the BlogPostId field if non-nil, zero value otherwise.

### GetBlogPostIdOk

`func (o *CustomContentSingle) GetBlogPostIdOk() (*string, bool)`

GetBlogPostIdOk returns a tuple with the BlogPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostId

`func (o *CustomContentSingle) SetBlogPostId(v string)`

SetBlogPostId sets BlogPostId field to given value.

### HasBlogPostId

`func (o *CustomContentSingle) HasBlogPostId() bool`

HasBlogPostId returns a boolean if a field has been set.

### GetCustomContentId

`func (o *CustomContentSingle) GetCustomContentId() string`

GetCustomContentId returns the CustomContentId field if non-nil, zero value otherwise.

### GetCustomContentIdOk

`func (o *CustomContentSingle) GetCustomContentIdOk() (*string, bool)`

GetCustomContentIdOk returns a tuple with the CustomContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContentId

`func (o *CustomContentSingle) SetCustomContentId(v string)`

SetCustomContentId sets CustomContentId field to given value.

### HasCustomContentId

`func (o *CustomContentSingle) HasCustomContentId() bool`

HasCustomContentId returns a boolean if a field has been set.

### GetAuthorId

`func (o *CustomContentSingle) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *CustomContentSingle) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *CustomContentSingle) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *CustomContentSingle) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CustomContentSingle) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomContentSingle) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomContentSingle) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CustomContentSingle) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *CustomContentSingle) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CustomContentSingle) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CustomContentSingle) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CustomContentSingle) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLabels

`func (o *CustomContentSingle) GetLabels() AttachmentSingleLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CustomContentSingle) GetLabelsOk() (*AttachmentSingleLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CustomContentSingle) SetLabels(v AttachmentSingleLabels)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CustomContentSingle) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetProperties

`func (o *CustomContentSingle) GetProperties() AttachmentSingleProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CustomContentSingle) GetPropertiesOk() (*AttachmentSingleProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CustomContentSingle) SetProperties(v AttachmentSingleProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CustomContentSingle) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetOperations

`func (o *CustomContentSingle) GetOperations() AttachmentSingleOperations`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *CustomContentSingle) GetOperationsOk() (*AttachmentSingleOperations, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *CustomContentSingle) SetOperations(v AttachmentSingleOperations)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *CustomContentSingle) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetVersions

`func (o *CustomContentSingle) GetVersions() AttachmentSingleVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *CustomContentSingle) GetVersionsOk() (*AttachmentSingleVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *CustomContentSingle) SetVersions(v AttachmentSingleVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *CustomContentSingle) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetBody

`func (o *CustomContentSingle) GetBody() CustomContentBodySingle`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CustomContentSingle) GetBodyOk() (*CustomContentBodySingle, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CustomContentSingle) SetBody(v CustomContentBodySingle)`

SetBody sets Body field to given value.

### HasBody

`func (o *CustomContentSingle) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetLinks

`func (o *CustomContentSingle) GetLinks() CustomContentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CustomContentSingle) GetLinksOk() (*CustomContentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CustomContentSingle) SetLinks(v CustomContentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CustomContentSingle) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



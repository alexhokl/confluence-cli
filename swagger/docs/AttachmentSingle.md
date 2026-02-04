# AttachmentSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the attachment. | [optional] 
**Status** | Pointer to [**ContentStatus**](ContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the comment. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time when the attachment was created. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 
**PageId** | Pointer to **string** | ID of the containing page.  Note: This is only returned if the attachment has a container that is a page. | [optional] 
**BlogPostId** | Pointer to **string** | ID of the containing blog post.  Note: This is only returned if the attachment has a container that is a blog post. | [optional] 
**CustomContentId** | Pointer to **string** | ID of the containing custom content.  Note: This is only returned if the attachment has a container that is custom content. | [optional] 
**MediaType** | Pointer to **string** | Media Type for the attachment. | [optional] 
**MediaTypeDescription** | Pointer to **string** | Media Type description for the attachment. | [optional] 
**Comment** | Pointer to **string** | Comment for the attachment. | [optional] 
**FileId** | Pointer to **string** | File ID of the attachment. This is the ID referenced in &#x60;atlas_doc_format&#x60; bodies and is distinct from the attachment ID. | [optional] 
**FileSize** | Pointer to **int64** | File size of the attachment. | [optional] 
**WebuiLink** | Pointer to **string** | WebUI link of the attachment. | [optional] 
**DownloadLink** | Pointer to **string** | Download link of the attachment. | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Labels** | Pointer to [**AttachmentSingleLabels**](AttachmentSingleLabels.md) |  | [optional] 
**Properties** | Pointer to [**AttachmentSingleProperties**](AttachmentSingleProperties.md) |  | [optional] 
**Operations** | Pointer to [**AttachmentSingleOperations**](AttachmentSingleOperations.md) |  | [optional] 
**Versions** | Pointer to [**AttachmentSingleVersions**](AttachmentSingleVersions.md) |  | [optional] 
**Links** | Pointer to [**AttachmentLinks**](AttachmentLinks.md) |  | [optional] 

## Methods

### NewAttachmentSingle

`func NewAttachmentSingle() *AttachmentSingle`

NewAttachmentSingle instantiates a new AttachmentSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentSingleWithDefaults

`func NewAttachmentSingleWithDefaults() *AttachmentSingle`

NewAttachmentSingleWithDefaults instantiates a new AttachmentSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AttachmentSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttachmentSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttachmentSingle) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AttachmentSingle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *AttachmentSingle) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AttachmentSingle) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AttachmentSingle) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AttachmentSingle) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *AttachmentSingle) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AttachmentSingle) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AttachmentSingle) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AttachmentSingle) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AttachmentSingle) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AttachmentSingle) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AttachmentSingle) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AttachmentSingle) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPageId

`func (o *AttachmentSingle) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *AttachmentSingle) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *AttachmentSingle) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *AttachmentSingle) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetBlogPostId

`func (o *AttachmentSingle) GetBlogPostId() string`

GetBlogPostId returns the BlogPostId field if non-nil, zero value otherwise.

### GetBlogPostIdOk

`func (o *AttachmentSingle) GetBlogPostIdOk() (*string, bool)`

GetBlogPostIdOk returns a tuple with the BlogPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostId

`func (o *AttachmentSingle) SetBlogPostId(v string)`

SetBlogPostId sets BlogPostId field to given value.

### HasBlogPostId

`func (o *AttachmentSingle) HasBlogPostId() bool`

HasBlogPostId returns a boolean if a field has been set.

### GetCustomContentId

`func (o *AttachmentSingle) GetCustomContentId() string`

GetCustomContentId returns the CustomContentId field if non-nil, zero value otherwise.

### GetCustomContentIdOk

`func (o *AttachmentSingle) GetCustomContentIdOk() (*string, bool)`

GetCustomContentIdOk returns a tuple with the CustomContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContentId

`func (o *AttachmentSingle) SetCustomContentId(v string)`

SetCustomContentId sets CustomContentId field to given value.

### HasCustomContentId

`func (o *AttachmentSingle) HasCustomContentId() bool`

HasCustomContentId returns a boolean if a field has been set.

### GetMediaType

`func (o *AttachmentSingle) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *AttachmentSingle) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *AttachmentSingle) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *AttachmentSingle) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetMediaTypeDescription

`func (o *AttachmentSingle) GetMediaTypeDescription() string`

GetMediaTypeDescription returns the MediaTypeDescription field if non-nil, zero value otherwise.

### GetMediaTypeDescriptionOk

`func (o *AttachmentSingle) GetMediaTypeDescriptionOk() (*string, bool)`

GetMediaTypeDescriptionOk returns a tuple with the MediaTypeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaTypeDescription

`func (o *AttachmentSingle) SetMediaTypeDescription(v string)`

SetMediaTypeDescription sets MediaTypeDescription field to given value.

### HasMediaTypeDescription

`func (o *AttachmentSingle) HasMediaTypeDescription() bool`

HasMediaTypeDescription returns a boolean if a field has been set.

### GetComment

`func (o *AttachmentSingle) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AttachmentSingle) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AttachmentSingle) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AttachmentSingle) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetFileId

`func (o *AttachmentSingle) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *AttachmentSingle) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *AttachmentSingle) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *AttachmentSingle) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetFileSize

`func (o *AttachmentSingle) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *AttachmentSingle) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *AttachmentSingle) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *AttachmentSingle) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetWebuiLink

`func (o *AttachmentSingle) GetWebuiLink() string`

GetWebuiLink returns the WebuiLink field if non-nil, zero value otherwise.

### GetWebuiLinkOk

`func (o *AttachmentSingle) GetWebuiLinkOk() (*string, bool)`

GetWebuiLinkOk returns a tuple with the WebuiLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebuiLink

`func (o *AttachmentSingle) SetWebuiLink(v string)`

SetWebuiLink sets WebuiLink field to given value.

### HasWebuiLink

`func (o *AttachmentSingle) HasWebuiLink() bool`

HasWebuiLink returns a boolean if a field has been set.

### GetDownloadLink

`func (o *AttachmentSingle) GetDownloadLink() string`

GetDownloadLink returns the DownloadLink field if non-nil, zero value otherwise.

### GetDownloadLinkOk

`func (o *AttachmentSingle) GetDownloadLinkOk() (*string, bool)`

GetDownloadLinkOk returns a tuple with the DownloadLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadLink

`func (o *AttachmentSingle) SetDownloadLink(v string)`

SetDownloadLink sets DownloadLink field to given value.

### HasDownloadLink

`func (o *AttachmentSingle) HasDownloadLink() bool`

HasDownloadLink returns a boolean if a field has been set.

### GetVersion

`func (o *AttachmentSingle) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AttachmentSingle) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AttachmentSingle) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AttachmentSingle) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLabels

`func (o *AttachmentSingle) GetLabels() AttachmentSingleLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AttachmentSingle) GetLabelsOk() (*AttachmentSingleLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AttachmentSingle) SetLabels(v AttachmentSingleLabels)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AttachmentSingle) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetProperties

`func (o *AttachmentSingle) GetProperties() AttachmentSingleProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AttachmentSingle) GetPropertiesOk() (*AttachmentSingleProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AttachmentSingle) SetProperties(v AttachmentSingleProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *AttachmentSingle) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetOperations

`func (o *AttachmentSingle) GetOperations() AttachmentSingleOperations`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *AttachmentSingle) GetOperationsOk() (*AttachmentSingleOperations, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *AttachmentSingle) SetOperations(v AttachmentSingleOperations)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *AttachmentSingle) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetVersions

`func (o *AttachmentSingle) GetVersions() AttachmentSingleVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AttachmentSingle) GetVersionsOk() (*AttachmentSingleVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AttachmentSingle) SetVersions(v AttachmentSingleVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *AttachmentSingle) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetLinks

`func (o *AttachmentSingle) GetLinks() AttachmentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AttachmentSingle) GetLinksOk() (*AttachmentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AttachmentSingle) SetLinks(v AttachmentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AttachmentSingle) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



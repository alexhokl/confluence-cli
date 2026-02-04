# AttachmentBulk

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
**Links** | Pointer to [**AttachmentLinks**](AttachmentLinks.md) |  | [optional] 

## Methods

### NewAttachmentBulk

`func NewAttachmentBulk() *AttachmentBulk`

NewAttachmentBulk instantiates a new AttachmentBulk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentBulkWithDefaults

`func NewAttachmentBulkWithDefaults() *AttachmentBulk`

NewAttachmentBulkWithDefaults instantiates a new AttachmentBulk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AttachmentBulk) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttachmentBulk) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttachmentBulk) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AttachmentBulk) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *AttachmentBulk) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AttachmentBulk) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AttachmentBulk) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AttachmentBulk) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *AttachmentBulk) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AttachmentBulk) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AttachmentBulk) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AttachmentBulk) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AttachmentBulk) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AttachmentBulk) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AttachmentBulk) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AttachmentBulk) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPageId

`func (o *AttachmentBulk) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *AttachmentBulk) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *AttachmentBulk) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *AttachmentBulk) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetBlogPostId

`func (o *AttachmentBulk) GetBlogPostId() string`

GetBlogPostId returns the BlogPostId field if non-nil, zero value otherwise.

### GetBlogPostIdOk

`func (o *AttachmentBulk) GetBlogPostIdOk() (*string, bool)`

GetBlogPostIdOk returns a tuple with the BlogPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostId

`func (o *AttachmentBulk) SetBlogPostId(v string)`

SetBlogPostId sets BlogPostId field to given value.

### HasBlogPostId

`func (o *AttachmentBulk) HasBlogPostId() bool`

HasBlogPostId returns a boolean if a field has been set.

### GetCustomContentId

`func (o *AttachmentBulk) GetCustomContentId() string`

GetCustomContentId returns the CustomContentId field if non-nil, zero value otherwise.

### GetCustomContentIdOk

`func (o *AttachmentBulk) GetCustomContentIdOk() (*string, bool)`

GetCustomContentIdOk returns a tuple with the CustomContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContentId

`func (o *AttachmentBulk) SetCustomContentId(v string)`

SetCustomContentId sets CustomContentId field to given value.

### HasCustomContentId

`func (o *AttachmentBulk) HasCustomContentId() bool`

HasCustomContentId returns a boolean if a field has been set.

### GetMediaType

`func (o *AttachmentBulk) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *AttachmentBulk) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *AttachmentBulk) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *AttachmentBulk) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetMediaTypeDescription

`func (o *AttachmentBulk) GetMediaTypeDescription() string`

GetMediaTypeDescription returns the MediaTypeDescription field if non-nil, zero value otherwise.

### GetMediaTypeDescriptionOk

`func (o *AttachmentBulk) GetMediaTypeDescriptionOk() (*string, bool)`

GetMediaTypeDescriptionOk returns a tuple with the MediaTypeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaTypeDescription

`func (o *AttachmentBulk) SetMediaTypeDescription(v string)`

SetMediaTypeDescription sets MediaTypeDescription field to given value.

### HasMediaTypeDescription

`func (o *AttachmentBulk) HasMediaTypeDescription() bool`

HasMediaTypeDescription returns a boolean if a field has been set.

### GetComment

`func (o *AttachmentBulk) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AttachmentBulk) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AttachmentBulk) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AttachmentBulk) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetFileId

`func (o *AttachmentBulk) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *AttachmentBulk) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *AttachmentBulk) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *AttachmentBulk) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetFileSize

`func (o *AttachmentBulk) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *AttachmentBulk) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *AttachmentBulk) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *AttachmentBulk) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetWebuiLink

`func (o *AttachmentBulk) GetWebuiLink() string`

GetWebuiLink returns the WebuiLink field if non-nil, zero value otherwise.

### GetWebuiLinkOk

`func (o *AttachmentBulk) GetWebuiLinkOk() (*string, bool)`

GetWebuiLinkOk returns a tuple with the WebuiLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebuiLink

`func (o *AttachmentBulk) SetWebuiLink(v string)`

SetWebuiLink sets WebuiLink field to given value.

### HasWebuiLink

`func (o *AttachmentBulk) HasWebuiLink() bool`

HasWebuiLink returns a boolean if a field has been set.

### GetDownloadLink

`func (o *AttachmentBulk) GetDownloadLink() string`

GetDownloadLink returns the DownloadLink field if non-nil, zero value otherwise.

### GetDownloadLinkOk

`func (o *AttachmentBulk) GetDownloadLinkOk() (*string, bool)`

GetDownloadLinkOk returns a tuple with the DownloadLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadLink

`func (o *AttachmentBulk) SetDownloadLink(v string)`

SetDownloadLink sets DownloadLink field to given value.

### HasDownloadLink

`func (o *AttachmentBulk) HasDownloadLink() bool`

HasDownloadLink returns a boolean if a field has been set.

### GetVersion

`func (o *AttachmentBulk) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AttachmentBulk) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AttachmentBulk) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AttachmentBulk) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLinks

`func (o *AttachmentBulk) GetLinks() AttachmentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AttachmentBulk) GetLinksOk() (*AttachmentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AttachmentBulk) SetLinks(v AttachmentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AttachmentBulk) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



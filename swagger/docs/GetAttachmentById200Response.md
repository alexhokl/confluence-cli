# GetAttachmentById200Response

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
**Links** | Pointer to [**GetAttachmentById200ResponseAllOfLinks**](GetAttachmentById200ResponseAllOfLinks.md) |  | [optional] 

## Methods

### NewGetAttachmentById200Response

`func NewGetAttachmentById200Response() *GetAttachmentById200Response`

NewGetAttachmentById200Response instantiates a new GetAttachmentById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAttachmentById200ResponseWithDefaults

`func NewGetAttachmentById200ResponseWithDefaults() *GetAttachmentById200Response`

NewGetAttachmentById200ResponseWithDefaults instantiates a new GetAttachmentById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAttachmentById200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAttachmentById200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAttachmentById200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetAttachmentById200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *GetAttachmentById200Response) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAttachmentById200Response) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAttachmentById200Response) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetAttachmentById200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *GetAttachmentById200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetAttachmentById200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetAttachmentById200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetAttachmentById200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetAttachmentById200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetAttachmentById200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetAttachmentById200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetAttachmentById200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPageId

`func (o *GetAttachmentById200Response) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *GetAttachmentById200Response) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *GetAttachmentById200Response) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *GetAttachmentById200Response) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetBlogPostId

`func (o *GetAttachmentById200Response) GetBlogPostId() string`

GetBlogPostId returns the BlogPostId field if non-nil, zero value otherwise.

### GetBlogPostIdOk

`func (o *GetAttachmentById200Response) GetBlogPostIdOk() (*string, bool)`

GetBlogPostIdOk returns a tuple with the BlogPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostId

`func (o *GetAttachmentById200Response) SetBlogPostId(v string)`

SetBlogPostId sets BlogPostId field to given value.

### HasBlogPostId

`func (o *GetAttachmentById200Response) HasBlogPostId() bool`

HasBlogPostId returns a boolean if a field has been set.

### GetCustomContentId

`func (o *GetAttachmentById200Response) GetCustomContentId() string`

GetCustomContentId returns the CustomContentId field if non-nil, zero value otherwise.

### GetCustomContentIdOk

`func (o *GetAttachmentById200Response) GetCustomContentIdOk() (*string, bool)`

GetCustomContentIdOk returns a tuple with the CustomContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContentId

`func (o *GetAttachmentById200Response) SetCustomContentId(v string)`

SetCustomContentId sets CustomContentId field to given value.

### HasCustomContentId

`func (o *GetAttachmentById200Response) HasCustomContentId() bool`

HasCustomContentId returns a boolean if a field has been set.

### GetMediaType

`func (o *GetAttachmentById200Response) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *GetAttachmentById200Response) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *GetAttachmentById200Response) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *GetAttachmentById200Response) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetMediaTypeDescription

`func (o *GetAttachmentById200Response) GetMediaTypeDescription() string`

GetMediaTypeDescription returns the MediaTypeDescription field if non-nil, zero value otherwise.

### GetMediaTypeDescriptionOk

`func (o *GetAttachmentById200Response) GetMediaTypeDescriptionOk() (*string, bool)`

GetMediaTypeDescriptionOk returns a tuple with the MediaTypeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaTypeDescription

`func (o *GetAttachmentById200Response) SetMediaTypeDescription(v string)`

SetMediaTypeDescription sets MediaTypeDescription field to given value.

### HasMediaTypeDescription

`func (o *GetAttachmentById200Response) HasMediaTypeDescription() bool`

HasMediaTypeDescription returns a boolean if a field has been set.

### GetComment

`func (o *GetAttachmentById200Response) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetAttachmentById200Response) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetAttachmentById200Response) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetAttachmentById200Response) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetFileId

`func (o *GetAttachmentById200Response) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *GetAttachmentById200Response) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *GetAttachmentById200Response) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *GetAttachmentById200Response) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetFileSize

`func (o *GetAttachmentById200Response) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *GetAttachmentById200Response) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *GetAttachmentById200Response) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *GetAttachmentById200Response) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetWebuiLink

`func (o *GetAttachmentById200Response) GetWebuiLink() string`

GetWebuiLink returns the WebuiLink field if non-nil, zero value otherwise.

### GetWebuiLinkOk

`func (o *GetAttachmentById200Response) GetWebuiLinkOk() (*string, bool)`

GetWebuiLinkOk returns a tuple with the WebuiLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebuiLink

`func (o *GetAttachmentById200Response) SetWebuiLink(v string)`

SetWebuiLink sets WebuiLink field to given value.

### HasWebuiLink

`func (o *GetAttachmentById200Response) HasWebuiLink() bool`

HasWebuiLink returns a boolean if a field has been set.

### GetDownloadLink

`func (o *GetAttachmentById200Response) GetDownloadLink() string`

GetDownloadLink returns the DownloadLink field if non-nil, zero value otherwise.

### GetDownloadLinkOk

`func (o *GetAttachmentById200Response) GetDownloadLinkOk() (*string, bool)`

GetDownloadLinkOk returns a tuple with the DownloadLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadLink

`func (o *GetAttachmentById200Response) SetDownloadLink(v string)`

SetDownloadLink sets DownloadLink field to given value.

### HasDownloadLink

`func (o *GetAttachmentById200Response) HasDownloadLink() bool`

HasDownloadLink returns a boolean if a field has been set.

### GetVersion

`func (o *GetAttachmentById200Response) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetAttachmentById200Response) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetAttachmentById200Response) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetAttachmentById200Response) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLabels

`func (o *GetAttachmentById200Response) GetLabels() AttachmentSingleLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetAttachmentById200Response) GetLabelsOk() (*AttachmentSingleLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetAttachmentById200Response) SetLabels(v AttachmentSingleLabels)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetAttachmentById200Response) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetProperties

`func (o *GetAttachmentById200Response) GetProperties() AttachmentSingleProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *GetAttachmentById200Response) GetPropertiesOk() (*AttachmentSingleProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *GetAttachmentById200Response) SetProperties(v AttachmentSingleProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *GetAttachmentById200Response) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetOperations

`func (o *GetAttachmentById200Response) GetOperations() AttachmentSingleOperations`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *GetAttachmentById200Response) GetOperationsOk() (*AttachmentSingleOperations, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *GetAttachmentById200Response) SetOperations(v AttachmentSingleOperations)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *GetAttachmentById200Response) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetVersions

`func (o *GetAttachmentById200Response) GetVersions() AttachmentSingleVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *GetAttachmentById200Response) GetVersionsOk() (*AttachmentSingleVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *GetAttachmentById200Response) SetVersions(v AttachmentSingleVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *GetAttachmentById200Response) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetLinks

`func (o *GetAttachmentById200Response) GetLinks() GetAttachmentById200ResponseAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetAttachmentById200Response) GetLinksOk() (*GetAttachmentById200ResponseAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetAttachmentById200Response) SetLinks(v GetAttachmentById200ResponseAllOfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetAttachmentById200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



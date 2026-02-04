# AttachmentCommentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the comment. | [optional] 
**Status** | Pointer to [**ContentStatus**](ContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the comment. | [optional] 
**AttachmentId** | Pointer to **string** | ID of the attachment containing the comment. | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Body** | Pointer to [**BodySingle**](BodySingle.md) |  | [optional] 
**Links** | Pointer to [**CommentLinks**](CommentLinks.md) |  | [optional] 

## Methods

### NewAttachmentCommentModel

`func NewAttachmentCommentModel() *AttachmentCommentModel`

NewAttachmentCommentModel instantiates a new AttachmentCommentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentCommentModelWithDefaults

`func NewAttachmentCommentModelWithDefaults() *AttachmentCommentModel`

NewAttachmentCommentModelWithDefaults instantiates a new AttachmentCommentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AttachmentCommentModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttachmentCommentModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttachmentCommentModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AttachmentCommentModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *AttachmentCommentModel) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AttachmentCommentModel) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AttachmentCommentModel) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AttachmentCommentModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *AttachmentCommentModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AttachmentCommentModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AttachmentCommentModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AttachmentCommentModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAttachmentId

`func (o *AttachmentCommentModel) GetAttachmentId() string`

GetAttachmentId returns the AttachmentId field if non-nil, zero value otherwise.

### GetAttachmentIdOk

`func (o *AttachmentCommentModel) GetAttachmentIdOk() (*string, bool)`

GetAttachmentIdOk returns a tuple with the AttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentId

`func (o *AttachmentCommentModel) SetAttachmentId(v string)`

SetAttachmentId sets AttachmentId field to given value.

### HasAttachmentId

`func (o *AttachmentCommentModel) HasAttachmentId() bool`

HasAttachmentId returns a boolean if a field has been set.

### GetVersion

`func (o *AttachmentCommentModel) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AttachmentCommentModel) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AttachmentCommentModel) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AttachmentCommentModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBody

`func (o *AttachmentCommentModel) GetBody() BodySingle`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *AttachmentCommentModel) GetBodyOk() (*BodySingle, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *AttachmentCommentModel) SetBody(v BodySingle)`

SetBody sets Body field to given value.

### HasBody

`func (o *AttachmentCommentModel) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetLinks

`func (o *AttachmentCommentModel) GetLinks() CommentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AttachmentCommentModel) GetLinksOk() (*CommentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AttachmentCommentModel) SetLinks(v CommentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AttachmentCommentModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



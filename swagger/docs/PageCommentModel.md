# PageCommentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the comment. | [optional] 
**Status** | Pointer to [**ContentStatus**](ContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the comment. | [optional] 
**PageId** | Pointer to **string** | ID of the page the comment is in. | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Body** | Pointer to [**BodyBulk**](BodyBulk.md) |  | [optional] 
**Links** | Pointer to [**CommentLinks**](CommentLinks.md) |  | [optional] 

## Methods

### NewPageCommentModel

`func NewPageCommentModel() *PageCommentModel`

NewPageCommentModel instantiates a new PageCommentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageCommentModelWithDefaults

`func NewPageCommentModelWithDefaults() *PageCommentModel`

NewPageCommentModelWithDefaults instantiates a new PageCommentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PageCommentModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PageCommentModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PageCommentModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PageCommentModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *PageCommentModel) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PageCommentModel) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PageCommentModel) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PageCommentModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *PageCommentModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PageCommentModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PageCommentModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PageCommentModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPageId

`func (o *PageCommentModel) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *PageCommentModel) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *PageCommentModel) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *PageCommentModel) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetVersion

`func (o *PageCommentModel) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PageCommentModel) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PageCommentModel) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PageCommentModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBody

`func (o *PageCommentModel) GetBody() BodyBulk`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PageCommentModel) GetBodyOk() (*BodyBulk, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PageCommentModel) SetBody(v BodyBulk)`

SetBody sets Body field to given value.

### HasBody

`func (o *PageCommentModel) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetLinks

`func (o *PageCommentModel) GetLinks() CommentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PageCommentModel) GetLinksOk() (*CommentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PageCommentModel) SetLinks(v CommentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PageCommentModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PageInlineCommentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the comment. | [optional] 
**Status** | Pointer to [**ContentStatus**](ContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the comment. | [optional] 
**PageId** | Pointer to **string** | ID of the page the comment is in. | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Body** | Pointer to [**BodyBulk**](BodyBulk.md) |  | [optional] 
**ResolutionStatus** | Pointer to [**InlineCommentResolutionStatus**](InlineCommentResolutionStatus.md) |  | [optional] 
**Properties** | Pointer to [**InlineCommentProperties**](InlineCommentProperties.md) |  | [optional] 
**Links** | Pointer to [**CommentLinks**](CommentLinks.md) |  | [optional] 

## Methods

### NewPageInlineCommentModel

`func NewPageInlineCommentModel() *PageInlineCommentModel`

NewPageInlineCommentModel instantiates a new PageInlineCommentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageInlineCommentModelWithDefaults

`func NewPageInlineCommentModelWithDefaults() *PageInlineCommentModel`

NewPageInlineCommentModelWithDefaults instantiates a new PageInlineCommentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PageInlineCommentModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PageInlineCommentModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PageInlineCommentModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PageInlineCommentModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *PageInlineCommentModel) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PageInlineCommentModel) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PageInlineCommentModel) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PageInlineCommentModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *PageInlineCommentModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PageInlineCommentModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PageInlineCommentModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PageInlineCommentModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPageId

`func (o *PageInlineCommentModel) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *PageInlineCommentModel) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *PageInlineCommentModel) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *PageInlineCommentModel) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetVersion

`func (o *PageInlineCommentModel) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PageInlineCommentModel) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PageInlineCommentModel) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PageInlineCommentModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBody

`func (o *PageInlineCommentModel) GetBody() BodyBulk`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PageInlineCommentModel) GetBodyOk() (*BodyBulk, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PageInlineCommentModel) SetBody(v BodyBulk)`

SetBody sets Body field to given value.

### HasBody

`func (o *PageInlineCommentModel) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetResolutionStatus

`func (o *PageInlineCommentModel) GetResolutionStatus() InlineCommentResolutionStatus`

GetResolutionStatus returns the ResolutionStatus field if non-nil, zero value otherwise.

### GetResolutionStatusOk

`func (o *PageInlineCommentModel) GetResolutionStatusOk() (*InlineCommentResolutionStatus, bool)`

GetResolutionStatusOk returns a tuple with the ResolutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionStatus

`func (o *PageInlineCommentModel) SetResolutionStatus(v InlineCommentResolutionStatus)`

SetResolutionStatus sets ResolutionStatus field to given value.

### HasResolutionStatus

`func (o *PageInlineCommentModel) HasResolutionStatus() bool`

HasResolutionStatus returns a boolean if a field has been set.

### GetProperties

`func (o *PageInlineCommentModel) GetProperties() InlineCommentProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PageInlineCommentModel) GetPropertiesOk() (*InlineCommentProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PageInlineCommentModel) SetProperties(v InlineCommentProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PageInlineCommentModel) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetLinks

`func (o *PageInlineCommentModel) GetLinks() CommentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PageInlineCommentModel) GetLinksOk() (*CommentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PageInlineCommentModel) SetLinks(v CommentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PageInlineCommentModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



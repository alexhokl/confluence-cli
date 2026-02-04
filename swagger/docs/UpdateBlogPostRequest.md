# UpdateBlogPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the blog post. | 
**Status** | **string** | The updated status of the blog post.  Note, if you change the status of a blog post from &#39;current&#39; to &#39;draft&#39; and it has an existing draft, the existing draft will be deleted in favor of the updated draft. Additionally, this endpoint can be used to restore a &#39;trashed&#39; or &#39;deleted&#39; blog post to &#39;current&#39; status. For restoration, blog post contents will not be updated and only the blog post status will be changed. | 
**Title** | **string** | Title of the blog post. | 
**SpaceId** | Pointer to **string** | ID of the containing space.  This currently **does not support moving the blog post to a different space**. | [optional] 
**Body** | [**CreateBlogPostRequestBody**](CreateBlogPostRequestBody.md) |  | 
**Version** | [**UpdateBlogPostRequestVersion**](UpdateBlogPostRequestVersion.md) |  | 
**CreatedAt** | Pointer to **string** | Created date of the blog post in the format of \&quot;yyyy-MM-ddTHH:mm:ss.SSSZ\&quot;. | [optional] 

## Methods

### NewUpdateBlogPostRequest

`func NewUpdateBlogPostRequest(id string, status string, title string, body CreateBlogPostRequestBody, version UpdateBlogPostRequestVersion, ) *UpdateBlogPostRequest`

NewUpdateBlogPostRequest instantiates a new UpdateBlogPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBlogPostRequestWithDefaults

`func NewUpdateBlogPostRequestWithDefaults() *UpdateBlogPostRequest`

NewUpdateBlogPostRequestWithDefaults instantiates a new UpdateBlogPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateBlogPostRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateBlogPostRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateBlogPostRequest) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *UpdateBlogPostRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateBlogPostRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateBlogPostRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *UpdateBlogPostRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateBlogPostRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateBlogPostRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSpaceId

`func (o *UpdateBlogPostRequest) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *UpdateBlogPostRequest) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *UpdateBlogPostRequest) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *UpdateBlogPostRequest) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetBody

`func (o *UpdateBlogPostRequest) GetBody() CreateBlogPostRequestBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *UpdateBlogPostRequest) GetBodyOk() (*CreateBlogPostRequestBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *UpdateBlogPostRequest) SetBody(v CreateBlogPostRequestBody)`

SetBody sets Body field to given value.


### GetVersion

`func (o *UpdateBlogPostRequest) GetVersion() UpdateBlogPostRequestVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateBlogPostRequest) GetVersionOk() (*UpdateBlogPostRequestVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateBlogPostRequest) SetVersion(v UpdateBlogPostRequestVersion)`

SetVersion sets Version field to given value.


### GetCreatedAt

`func (o *UpdateBlogPostRequest) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdateBlogPostRequest) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdateBlogPostRequest) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UpdateBlogPostRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



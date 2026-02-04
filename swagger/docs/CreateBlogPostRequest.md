# CreateBlogPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpaceId** | **string** | ID of the space | 
**Status** | Pointer to **string** | The status of the blog post, specifies if the blog post will be created as a new blog post or a draft | [optional] 
**Title** | Pointer to **string** | Title of the blog post, required if creating non-draft. | [optional] 
**Body** | Pointer to [**CreateBlogPostRequestBody**](CreateBlogPostRequestBody.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | Created date of the blog post in the format of \&quot;yyyy-MM-ddTHH:mm:ss.SSSZ\&quot;. | [optional] 

## Methods

### NewCreateBlogPostRequest

`func NewCreateBlogPostRequest(spaceId string, ) *CreateBlogPostRequest`

NewCreateBlogPostRequest instantiates a new CreateBlogPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBlogPostRequestWithDefaults

`func NewCreateBlogPostRequestWithDefaults() *CreateBlogPostRequest`

NewCreateBlogPostRequestWithDefaults instantiates a new CreateBlogPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpaceId

`func (o *CreateBlogPostRequest) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *CreateBlogPostRequest) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *CreateBlogPostRequest) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetStatus

`func (o *CreateBlogPostRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateBlogPostRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateBlogPostRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateBlogPostRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *CreateBlogPostRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateBlogPostRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateBlogPostRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateBlogPostRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBody

`func (o *CreateBlogPostRequest) GetBody() CreateBlogPostRequestBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreateBlogPostRequest) GetBodyOk() (*CreateBlogPostRequestBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreateBlogPostRequest) SetBody(v CreateBlogPostRequestBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *CreateBlogPostRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateBlogPostRequest) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateBlogPostRequest) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateBlogPostRequest) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateBlogPostRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateSmartLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpaceId** | **string** | ID of the space. | 
**Title** | Pointer to **string** | Title of the Smart Link in the content tree. | [optional] 
**ParentId** | Pointer to **string** | The parent content ID of the Smart Link in the content tree. | [optional] 
**EmbedUrl** | Pointer to **string** | The URL that the Smart Link in the content tree should be populated with. | [optional] 

## Methods

### NewCreateSmartLinkRequest

`func NewCreateSmartLinkRequest(spaceId string, ) *CreateSmartLinkRequest`

NewCreateSmartLinkRequest instantiates a new CreateSmartLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSmartLinkRequestWithDefaults

`func NewCreateSmartLinkRequestWithDefaults() *CreateSmartLinkRequest`

NewCreateSmartLinkRequestWithDefaults instantiates a new CreateSmartLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpaceId

`func (o *CreateSmartLinkRequest) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *CreateSmartLinkRequest) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *CreateSmartLinkRequest) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetTitle

`func (o *CreateSmartLinkRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateSmartLinkRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateSmartLinkRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateSmartLinkRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetParentId

`func (o *CreateSmartLinkRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateSmartLinkRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateSmartLinkRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateSmartLinkRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetEmbedUrl

`func (o *CreateSmartLinkRequest) GetEmbedUrl() string`

GetEmbedUrl returns the EmbedUrl field if non-nil, zero value otherwise.

### GetEmbedUrlOk

`func (o *CreateSmartLinkRequest) GetEmbedUrlOk() (*string, bool)`

GetEmbedUrlOk returns a tuple with the EmbedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedUrl

`func (o *CreateSmartLinkRequest) SetEmbedUrl(v string)`

SetEmbedUrl sets EmbedUrl field to given value.

### HasEmbedUrl

`func (o *CreateSmartLinkRequest) HasEmbedUrl() bool`

HasEmbedUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



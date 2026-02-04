# UpdatePageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the page. | 
**Status** | **string** | The updated status of the page.  Note, if you change the status of a page from &#39;current&#39; to &#39;draft&#39; and it has an existing draft, the existing draft will be deleted in favor of the updated draft. Additionally, this endpoint can be used to restore a &#39;trashed&#39; or &#39;deleted&#39; page to &#39;current&#39; status. For restoration, page contents will not be updated and only the page status will be changed. | 
**Title** | **string** | Title of the page. | 
**SpaceId** | Pointer to **interface{}** | ID of the containing space.  This currently **does not support moving the page to a different space**. | [optional] 
**ParentId** | Pointer to **interface{}** | ID of the parent content.  This allows the page to be moved under a different parent within the same space. | [optional] 
**OwnerId** | Pointer to **interface{}** | Account ID of the page owner.  This allows page ownership to be transferred to another user. | [optional] 
**Body** | [**CreatePageRequestBody**](CreatePageRequestBody.md) |  | 
**Version** | [**UpdatePageRequestVersion**](UpdatePageRequestVersion.md) |  | 

## Methods

### NewUpdatePageRequest

`func NewUpdatePageRequest(id string, status string, title string, body CreatePageRequestBody, version UpdatePageRequestVersion, ) *UpdatePageRequest`

NewUpdatePageRequest instantiates a new UpdatePageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePageRequestWithDefaults

`func NewUpdatePageRequestWithDefaults() *UpdatePageRequest`

NewUpdatePageRequestWithDefaults instantiates a new UpdatePageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdatePageRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdatePageRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdatePageRequest) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *UpdatePageRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdatePageRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdatePageRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *UpdatePageRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdatePageRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdatePageRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSpaceId

`func (o *UpdatePageRequest) GetSpaceId() interface{}`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *UpdatePageRequest) GetSpaceIdOk() (*interface{}, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *UpdatePageRequest) SetSpaceId(v interface{})`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *UpdatePageRequest) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### SetSpaceIdNil

`func (o *UpdatePageRequest) SetSpaceIdNil(b bool)`

 SetSpaceIdNil sets the value for SpaceId to be an explicit nil

### UnsetSpaceId
`func (o *UpdatePageRequest) UnsetSpaceId()`

UnsetSpaceId ensures that no value is present for SpaceId, not even an explicit nil
### GetParentId

`func (o *UpdatePageRequest) GetParentId() interface{}`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *UpdatePageRequest) GetParentIdOk() (*interface{}, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *UpdatePageRequest) SetParentId(v interface{})`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *UpdatePageRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *UpdatePageRequest) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *UpdatePageRequest) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetOwnerId

`func (o *UpdatePageRequest) GetOwnerId() interface{}`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *UpdatePageRequest) GetOwnerIdOk() (*interface{}, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *UpdatePageRequest) SetOwnerId(v interface{})`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *UpdatePageRequest) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *UpdatePageRequest) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *UpdatePageRequest) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetBody

`func (o *UpdatePageRequest) GetBody() CreatePageRequestBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *UpdatePageRequest) GetBodyOk() (*CreatePageRequestBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *UpdatePageRequest) SetBody(v CreatePageRequestBody)`

SetBody sets Body field to given value.


### GetVersion

`func (o *UpdatePageRequest) GetVersion() UpdatePageRequestVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdatePageRequest) GetVersionOk() (*UpdatePageRequestVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdatePageRequest) SetVersion(v UpdatePageRequestVersion)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



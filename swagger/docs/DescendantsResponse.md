# DescendantsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the descendant. | [optional] 
**Status** | Pointer to [**OnlyArchivedAndCurrentContentStatus**](OnlyArchivedAndCurrentContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the descendant. | [optional] 
**Type** | Pointer to **string** | Hierarchical content type (database/embed/folder/page/whiteboard). | [optional] 
**ParentId** | Pointer to **string** | ID of the parent content. | [optional] 
**Depth** | Pointer to **int32** | Depth of the descendant in the content tree relative to the content specified in the request. | [optional] 
**ChildPosition** | Pointer to **NullableInt32** | Numerical value indicating position of the content relative to its siblings (with the same parentId) within the content tree. If the content is sorted by childPosition, it will reflect the default content ordering within the content tree. | [optional] 

## Methods

### NewDescendantsResponse

`func NewDescendantsResponse() *DescendantsResponse`

NewDescendantsResponse instantiates a new DescendantsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescendantsResponseWithDefaults

`func NewDescendantsResponseWithDefaults() *DescendantsResponse`

NewDescendantsResponseWithDefaults instantiates a new DescendantsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DescendantsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescendantsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescendantsResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DescendantsResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *DescendantsResponse) GetStatus() OnlyArchivedAndCurrentContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescendantsResponse) GetStatusOk() (*OnlyArchivedAndCurrentContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescendantsResponse) SetStatus(v OnlyArchivedAndCurrentContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DescendantsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *DescendantsResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DescendantsResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DescendantsResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DescendantsResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *DescendantsResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DescendantsResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DescendantsResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DescendantsResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetParentId

`func (o *DescendantsResponse) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *DescendantsResponse) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *DescendantsResponse) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *DescendantsResponse) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetDepth

`func (o *DescendantsResponse) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *DescendantsResponse) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *DescendantsResponse) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *DescendantsResponse) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetChildPosition

`func (o *DescendantsResponse) GetChildPosition() int32`

GetChildPosition returns the ChildPosition field if non-nil, zero value otherwise.

### GetChildPositionOk

`func (o *DescendantsResponse) GetChildPositionOk() (*int32, bool)`

GetChildPositionOk returns a tuple with the ChildPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildPosition

`func (o *DescendantsResponse) SetChildPosition(v int32)`

SetChildPosition sets ChildPosition field to given value.

### HasChildPosition

`func (o *DescendantsResponse) HasChildPosition() bool`

HasChildPosition returns a boolean if a field has been set.

### SetChildPositionNil

`func (o *DescendantsResponse) SetChildPositionNil(b bool)`

 SetChildPositionNil sets the value for ChildPosition to be an explicit nil

### UnsetChildPosition
`func (o *DescendantsResponse) UnsetChildPosition()`

UnsetChildPosition ensures that no value is present for ChildPosition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



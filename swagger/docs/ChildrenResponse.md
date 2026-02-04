# ChildrenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the child content. | [optional] 
**Status** | Pointer to [**OnlyArchivedAndCurrentContentStatus**](OnlyArchivedAndCurrentContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the child content. | [optional] 
**Type** | Pointer to **string** | Hierarchical content type (database/embed/folder/page/whiteboard). | [optional] 
**SpaceId** | Pointer to **string** | ID of the space the content is in. | [optional] 
**ChildPosition** | Pointer to **NullableInt32** | Numerical value indicating position of the content relative to its siblings (with the same parentId) within the content tree. If the content is sorted by childPosition, it will reflect the default content ordering within the content tree. | [optional] 

## Methods

### NewChildrenResponse

`func NewChildrenResponse() *ChildrenResponse`

NewChildrenResponse instantiates a new ChildrenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChildrenResponseWithDefaults

`func NewChildrenResponseWithDefaults() *ChildrenResponse`

NewChildrenResponseWithDefaults instantiates a new ChildrenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChildrenResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChildrenResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChildrenResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChildrenResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ChildrenResponse) GetStatus() OnlyArchivedAndCurrentContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChildrenResponse) GetStatusOk() (*OnlyArchivedAndCurrentContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChildrenResponse) SetStatus(v OnlyArchivedAndCurrentContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChildrenResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *ChildrenResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChildrenResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChildrenResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ChildrenResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *ChildrenResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChildrenResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChildrenResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ChildrenResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSpaceId

`func (o *ChildrenResponse) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *ChildrenResponse) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *ChildrenResponse) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *ChildrenResponse) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetChildPosition

`func (o *ChildrenResponse) GetChildPosition() int32`

GetChildPosition returns the ChildPosition field if non-nil, zero value otherwise.

### GetChildPositionOk

`func (o *ChildrenResponse) GetChildPositionOk() (*int32, bool)`

GetChildPositionOk returns a tuple with the ChildPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildPosition

`func (o *ChildrenResponse) SetChildPosition(v int32)`

SetChildPosition sets ChildPosition field to given value.

### HasChildPosition

`func (o *ChildrenResponse) HasChildPosition() bool`

HasChildPosition returns a boolean if a field has been set.

### SetChildPositionNil

`func (o *ChildrenResponse) SetChildPositionNil(b bool)`

 SetChildPositionNil sets the value for ChildPosition to be an explicit nil

### UnsetChildPosition
`func (o *ChildrenResponse) UnsetChildPosition()`

UnsetChildPosition ensures that no value is present for ChildPosition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



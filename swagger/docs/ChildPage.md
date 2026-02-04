# ChildPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the page. | [optional] 
**Status** | Pointer to [**OnlyArchivedAndCurrentContentStatus**](OnlyArchivedAndCurrentContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the page. | [optional] 
**SpaceId** | Pointer to **string** | ID of the space the page is in. | [optional] 
**ChildPosition** | Pointer to **NullableInt32** | Position of child page within the given parent page tree. | [optional] 

## Methods

### NewChildPage

`func NewChildPage() *ChildPage`

NewChildPage instantiates a new ChildPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChildPageWithDefaults

`func NewChildPageWithDefaults() *ChildPage`

NewChildPageWithDefaults instantiates a new ChildPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChildPage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChildPage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChildPage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChildPage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ChildPage) GetStatus() OnlyArchivedAndCurrentContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChildPage) GetStatusOk() (*OnlyArchivedAndCurrentContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChildPage) SetStatus(v OnlyArchivedAndCurrentContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChildPage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *ChildPage) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChildPage) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChildPage) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ChildPage) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSpaceId

`func (o *ChildPage) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *ChildPage) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *ChildPage) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *ChildPage) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetChildPosition

`func (o *ChildPage) GetChildPosition() int32`

GetChildPosition returns the ChildPosition field if non-nil, zero value otherwise.

### GetChildPositionOk

`func (o *ChildPage) GetChildPositionOk() (*int32, bool)`

GetChildPositionOk returns a tuple with the ChildPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildPosition

`func (o *ChildPage) SetChildPosition(v int32)`

SetChildPosition sets ChildPosition field to given value.

### HasChildPosition

`func (o *ChildPage) HasChildPosition() bool`

HasChildPosition returns a boolean if a field has been set.

### SetChildPositionNil

`func (o *ChildPage) SetChildPositionNil(b bool)`

 SetChildPositionNil sets the value for ChildPosition to be an explicit nil

### UnsetChildPosition
`func (o *ChildPage) UnsetChildPosition()`

UnsetChildPosition ensures that no value is present for ChildPosition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



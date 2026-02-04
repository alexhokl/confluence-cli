# CreateFolderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpaceId** | **string** | ID of the space. | 
**Title** | Pointer to **string** | Title of the folder. | [optional] 
**ParentId** | Pointer to **string** | The parent content ID of the folder. | [optional] 

## Methods

### NewCreateFolderRequest

`func NewCreateFolderRequest(spaceId string, ) *CreateFolderRequest`

NewCreateFolderRequest instantiates a new CreateFolderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFolderRequestWithDefaults

`func NewCreateFolderRequestWithDefaults() *CreateFolderRequest`

NewCreateFolderRequestWithDefaults instantiates a new CreateFolderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpaceId

`func (o *CreateFolderRequest) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *CreateFolderRequest) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *CreateFolderRequest) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetTitle

`func (o *CreateFolderRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateFolderRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateFolderRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateFolderRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetParentId

`func (o *CreateFolderRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateFolderRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateFolderRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateFolderRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



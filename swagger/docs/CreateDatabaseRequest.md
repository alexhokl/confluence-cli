# CreateDatabaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpaceId** | **string** | ID of the space. | 
**Title** | Pointer to **string** | Title of the database. | [optional] 
**ParentId** | Pointer to **string** | The parent content ID of the database. | [optional] 

## Methods

### NewCreateDatabaseRequest

`func NewCreateDatabaseRequest(spaceId string, ) *CreateDatabaseRequest`

NewCreateDatabaseRequest instantiates a new CreateDatabaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDatabaseRequestWithDefaults

`func NewCreateDatabaseRequestWithDefaults() *CreateDatabaseRequest`

NewCreateDatabaseRequestWithDefaults instantiates a new CreateDatabaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpaceId

`func (o *CreateDatabaseRequest) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *CreateDatabaseRequest) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *CreateDatabaseRequest) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetTitle

`func (o *CreateDatabaseRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateDatabaseRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateDatabaseRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateDatabaseRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetParentId

`func (o *CreateDatabaseRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateDatabaseRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateDatabaseRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateDatabaseRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



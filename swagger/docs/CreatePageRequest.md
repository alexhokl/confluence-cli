# CreatePageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpaceId** | **string** | ID of the space. | 
**Status** | Pointer to **string** | The status of the page, published or draft. | [optional] 
**Title** | Pointer to **string** | Title of the page, required if page status is not draft. | [optional] 
**ParentId** | Pointer to **string** | The parent content ID of the page. If the &#x60;root-level&#x60; query parameter is set to false and a value is  not supplied for this parameter, then the space homepage&#39;s ID will be used. If the &#x60;root-level&#x60; query  parameter is set to true, then a value may not be supplied for this parameter. | [optional] 
**Body** | Pointer to [**CreatePageRequestBody**](CreatePageRequestBody.md) |  | [optional] 
**Subtype** | Pointer to **string** | The subtype of the page. Provide the subtype live to create a live doc or no subtype to create a page. | [optional] 

## Methods

### NewCreatePageRequest

`func NewCreatePageRequest(spaceId string, ) *CreatePageRequest`

NewCreatePageRequest instantiates a new CreatePageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePageRequestWithDefaults

`func NewCreatePageRequestWithDefaults() *CreatePageRequest`

NewCreatePageRequestWithDefaults instantiates a new CreatePageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpaceId

`func (o *CreatePageRequest) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *CreatePageRequest) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *CreatePageRequest) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetStatus

`func (o *CreatePageRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreatePageRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreatePageRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreatePageRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *CreatePageRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreatePageRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreatePageRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreatePageRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetParentId

`func (o *CreatePageRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreatePageRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreatePageRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreatePageRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetBody

`func (o *CreatePageRequest) GetBody() CreatePageRequestBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreatePageRequest) GetBodyOk() (*CreatePageRequestBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreatePageRequest) SetBody(v CreatePageRequestBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *CreatePageRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetSubtype

`func (o *CreatePageRequest) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *CreatePageRequest) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *CreatePageRequest) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *CreatePageRequest) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



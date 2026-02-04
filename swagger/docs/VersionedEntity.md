# VersionedEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title of the entity. | [optional] 
**Id** | Pointer to **string** | ID of the entity. | [optional] 
**Body** | Pointer to [**BodyBulk**](BodyBulk.md) |  | [optional] 

## Methods

### NewVersionedEntity

`func NewVersionedEntity() *VersionedEntity`

NewVersionedEntity instantiates a new VersionedEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionedEntityWithDefaults

`func NewVersionedEntityWithDefaults() *VersionedEntity`

NewVersionedEntityWithDefaults instantiates a new VersionedEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *VersionedEntity) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VersionedEntity) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VersionedEntity) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *VersionedEntity) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetId

`func (o *VersionedEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VersionedEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VersionedEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VersionedEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBody

`func (o *VersionedEntity) GetBody() BodyBulk`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *VersionedEntity) GetBodyOk() (*BodyBulk, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *VersionedEntity) SetBody(v BodyBulk)`

SetBody sets Body field to given value.

### HasBody

`func (o *VersionedEntity) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



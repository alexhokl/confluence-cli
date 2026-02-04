# OptionalFieldMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** | Indicates if there are more available results that can be fetched. | [optional] 
**Cursor** | Pointer to **string** | A token that can be used in the query parameter of the endpoint returned in the &#x60;_links&#x60; property to retrieve the next set of results. | [optional] 

## Methods

### NewOptionalFieldMeta

`func NewOptionalFieldMeta() *OptionalFieldMeta`

NewOptionalFieldMeta instantiates a new OptionalFieldMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionalFieldMetaWithDefaults

`func NewOptionalFieldMetaWithDefaults() *OptionalFieldMeta`

NewOptionalFieldMetaWithDefaults instantiates a new OptionalFieldMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *OptionalFieldMeta) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *OptionalFieldMeta) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *OptionalFieldMeta) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *OptionalFieldMeta) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetCursor

`func (o *OptionalFieldMeta) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *OptionalFieldMeta) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *OptionalFieldMeta) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *OptionalFieldMeta) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



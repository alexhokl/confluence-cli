# SpacePropertyVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | RFC3339 compliant date time at which the property&#39;s current version was created. | [optional] 
**CreatedBy** | Pointer to **string** | Atlassian account ID of the user that created the space property&#39;s current version. | [optional] 
**Message** | Pointer to **string** | Message associated with the current version. | [optional] 
**Number** | Pointer to **int32** | The space property&#39;s current version number. | [optional] 

## Methods

### NewSpacePropertyVersion

`func NewSpacePropertyVersion() *SpacePropertyVersion`

NewSpacePropertyVersion instantiates a new SpacePropertyVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpacePropertyVersionWithDefaults

`func NewSpacePropertyVersionWithDefaults() *SpacePropertyVersion`

NewSpacePropertyVersionWithDefaults instantiates a new SpacePropertyVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SpacePropertyVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SpacePropertyVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SpacePropertyVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SpacePropertyVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *SpacePropertyVersion) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SpacePropertyVersion) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SpacePropertyVersion) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SpacePropertyVersion) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetMessage

`func (o *SpacePropertyVersion) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SpacePropertyVersion) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SpacePropertyVersion) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SpacePropertyVersion) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNumber

`func (o *SpacePropertyVersion) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *SpacePropertyVersion) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *SpacePropertyVersion) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *SpacePropertyVersion) HasNumber() bool`

HasNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



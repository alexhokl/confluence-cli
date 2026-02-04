# UpdateCustomContentRequestVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | The version number, must be incremented by one. | [optional] 
**Message** | Pointer to **string** | An optional message to be stored with the version. | [optional] 

## Methods

### NewUpdateCustomContentRequestVersion

`func NewUpdateCustomContentRequestVersion() *UpdateCustomContentRequestVersion`

NewUpdateCustomContentRequestVersion instantiates a new UpdateCustomContentRequestVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomContentRequestVersionWithDefaults

`func NewUpdateCustomContentRequestVersionWithDefaults() *UpdateCustomContentRequestVersion`

NewUpdateCustomContentRequestVersionWithDefaults instantiates a new UpdateCustomContentRequestVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *UpdateCustomContentRequestVersion) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *UpdateCustomContentRequestVersion) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *UpdateCustomContentRequestVersion) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *UpdateCustomContentRequestVersion) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetMessage

`func (o *UpdateCustomContentRequestVersion) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateCustomContentRequestVersion) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateCustomContentRequestVersion) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpdateCustomContentRequestVersion) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



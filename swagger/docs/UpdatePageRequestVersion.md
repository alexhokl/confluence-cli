# UpdatePageRequestVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | The new version of the updated page.  Set this to the current version number plus one, unless you are updating the status to &#39;draft&#39; which requires a version number of 1.  If you don&#39;t know the current version number, use Get page by id. | [optional] 
**Message** | Pointer to **string** | An optional message to be stored with the version. | [optional] 

## Methods

### NewUpdatePageRequestVersion

`func NewUpdatePageRequestVersion() *UpdatePageRequestVersion`

NewUpdatePageRequestVersion instantiates a new UpdatePageRequestVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePageRequestVersionWithDefaults

`func NewUpdatePageRequestVersionWithDefaults() *UpdatePageRequestVersion`

NewUpdatePageRequestVersionWithDefaults instantiates a new UpdatePageRequestVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *UpdatePageRequestVersion) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *UpdatePageRequestVersion) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *UpdatePageRequestVersion) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *UpdatePageRequestVersion) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetMessage

`func (o *UpdatePageRequestVersion) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdatePageRequestVersion) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdatePageRequestVersion) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpdatePageRequestVersion) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ContentPropertyUpdateRequestVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | Version number of the new version. Should be 1 more than the current version number. | [optional] 
**Message** | Pointer to **string** | Message to be associated with the new version. | [optional] 

## Methods

### NewContentPropertyUpdateRequestVersion

`func NewContentPropertyUpdateRequestVersion() *ContentPropertyUpdateRequestVersion`

NewContentPropertyUpdateRequestVersion instantiates a new ContentPropertyUpdateRequestVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentPropertyUpdateRequestVersionWithDefaults

`func NewContentPropertyUpdateRequestVersionWithDefaults() *ContentPropertyUpdateRequestVersion`

NewContentPropertyUpdateRequestVersionWithDefaults instantiates a new ContentPropertyUpdateRequestVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *ContentPropertyUpdateRequestVersion) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ContentPropertyUpdateRequestVersion) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ContentPropertyUpdateRequestVersion) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *ContentPropertyUpdateRequestVersion) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetMessage

`func (o *ContentPropertyUpdateRequestVersion) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ContentPropertyUpdateRequestVersion) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ContentPropertyUpdateRequestVersion) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ContentPropertyUpdateRequestVersion) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



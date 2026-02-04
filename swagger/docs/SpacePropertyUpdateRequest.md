# SpacePropertyUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key of the space property | [optional] 
**Value** | Pointer to **interface{}** | Value of the space property. | [optional] 
**Version** | Pointer to [**ContentPropertyUpdateRequestVersion**](ContentPropertyUpdateRequestVersion.md) |  | [optional] 

## Methods

### NewSpacePropertyUpdateRequest

`func NewSpacePropertyUpdateRequest() *SpacePropertyUpdateRequest`

NewSpacePropertyUpdateRequest instantiates a new SpacePropertyUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpacePropertyUpdateRequestWithDefaults

`func NewSpacePropertyUpdateRequestWithDefaults() *SpacePropertyUpdateRequest`

NewSpacePropertyUpdateRequestWithDefaults instantiates a new SpacePropertyUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SpacePropertyUpdateRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SpacePropertyUpdateRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SpacePropertyUpdateRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SpacePropertyUpdateRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *SpacePropertyUpdateRequest) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SpacePropertyUpdateRequest) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SpacePropertyUpdateRequest) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *SpacePropertyUpdateRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *SpacePropertyUpdateRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *SpacePropertyUpdateRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetVersion

`func (o *SpacePropertyUpdateRequest) GetVersion() ContentPropertyUpdateRequestVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SpacePropertyUpdateRequest) GetVersionOk() (*ContentPropertyUpdateRequestVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SpacePropertyUpdateRequest) SetVersion(v ContentPropertyUpdateRequestVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SpacePropertyUpdateRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



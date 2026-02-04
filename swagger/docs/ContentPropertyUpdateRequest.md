# ContentPropertyUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key of the content property | [optional] 
**Value** | Pointer to **interface{}** | Value of the content property. | [optional] 
**Version** | Pointer to [**ContentPropertyUpdateRequestVersion**](ContentPropertyUpdateRequestVersion.md) |  | [optional] 

## Methods

### NewContentPropertyUpdateRequest

`func NewContentPropertyUpdateRequest() *ContentPropertyUpdateRequest`

NewContentPropertyUpdateRequest instantiates a new ContentPropertyUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentPropertyUpdateRequestWithDefaults

`func NewContentPropertyUpdateRequestWithDefaults() *ContentPropertyUpdateRequest`

NewContentPropertyUpdateRequestWithDefaults instantiates a new ContentPropertyUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ContentPropertyUpdateRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ContentPropertyUpdateRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ContentPropertyUpdateRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ContentPropertyUpdateRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *ContentPropertyUpdateRequest) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContentPropertyUpdateRequest) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContentPropertyUpdateRequest) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ContentPropertyUpdateRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ContentPropertyUpdateRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ContentPropertyUpdateRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetVersion

`func (o *ContentPropertyUpdateRequest) GetVersion() ContentPropertyUpdateRequestVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ContentPropertyUpdateRequest) GetVersionOk() (*ContentPropertyUpdateRequestVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ContentPropertyUpdateRequest) SetVersion(v ContentPropertyUpdateRequestVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ContentPropertyUpdateRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



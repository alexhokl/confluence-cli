# SpacePropertyCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key of the space property | [optional] 
**Value** | Pointer to **interface{}** | Value of the space property. | [optional] 

## Methods

### NewSpacePropertyCreateRequest

`func NewSpacePropertyCreateRequest() *SpacePropertyCreateRequest`

NewSpacePropertyCreateRequest instantiates a new SpacePropertyCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpacePropertyCreateRequestWithDefaults

`func NewSpacePropertyCreateRequestWithDefaults() *SpacePropertyCreateRequest`

NewSpacePropertyCreateRequestWithDefaults instantiates a new SpacePropertyCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SpacePropertyCreateRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SpacePropertyCreateRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SpacePropertyCreateRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SpacePropertyCreateRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *SpacePropertyCreateRequest) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SpacePropertyCreateRequest) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SpacePropertyCreateRequest) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *SpacePropertyCreateRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *SpacePropertyCreateRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *SpacePropertyCreateRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



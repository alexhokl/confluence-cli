# ContentPropertyCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key of the content property | [optional] 
**Value** | Pointer to **interface{}** | Value of the content property. | [optional] 

## Methods

### NewContentPropertyCreateRequest

`func NewContentPropertyCreateRequest() *ContentPropertyCreateRequest`

NewContentPropertyCreateRequest instantiates a new ContentPropertyCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentPropertyCreateRequestWithDefaults

`func NewContentPropertyCreateRequestWithDefaults() *ContentPropertyCreateRequest`

NewContentPropertyCreateRequestWithDefaults instantiates a new ContentPropertyCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ContentPropertyCreateRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ContentPropertyCreateRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ContentPropertyCreateRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ContentPropertyCreateRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *ContentPropertyCreateRequest) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContentPropertyCreateRequest) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContentPropertyCreateRequest) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ContentPropertyCreateRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ContentPropertyCreateRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ContentPropertyCreateRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



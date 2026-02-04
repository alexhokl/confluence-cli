# ContentProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the property | [optional] 
**Key** | Pointer to **string** | Key of the property | [optional] 
**Value** | Pointer to **interface{}** | Value of the property. Must be a valid JSON value. | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 

## Methods

### NewContentProperty

`func NewContentProperty() *ContentProperty`

NewContentProperty instantiates a new ContentProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentPropertyWithDefaults

`func NewContentPropertyWithDefaults() *ContentProperty`

NewContentPropertyWithDefaults instantiates a new ContentProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContentProperty) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContentProperty) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContentProperty) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContentProperty) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *ContentProperty) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ContentProperty) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ContentProperty) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ContentProperty) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *ContentProperty) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContentProperty) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContentProperty) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ContentProperty) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ContentProperty) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ContentProperty) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetVersion

`func (o *ContentProperty) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ContentProperty) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ContentProperty) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ContentProperty) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



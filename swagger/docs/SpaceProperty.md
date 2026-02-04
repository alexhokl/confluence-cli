# SpaceProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the space property. | [optional] 
**Key** | Pointer to **string** | Key of the space property. | [optional] 
**Value** | Pointer to **interface{}** | Value of the space property. | [optional] 
**CreatedAt** | Pointer to **time.Time** | RFC3339 compliant date time at which the property was created. | [optional] 
**CreatedBy** | Pointer to **string** | Atlassian account ID of the user that created the space property. | [optional] 
**Version** | Pointer to [**SpacePropertyVersion**](SpacePropertyVersion.md) |  | [optional] 

## Methods

### NewSpaceProperty

`func NewSpaceProperty() *SpaceProperty`

NewSpaceProperty instantiates a new SpaceProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpacePropertyWithDefaults

`func NewSpacePropertyWithDefaults() *SpaceProperty`

NewSpacePropertyWithDefaults instantiates a new SpaceProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SpaceProperty) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpaceProperty) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpaceProperty) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SpaceProperty) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *SpaceProperty) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SpaceProperty) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SpaceProperty) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SpaceProperty) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *SpaceProperty) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SpaceProperty) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SpaceProperty) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *SpaceProperty) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *SpaceProperty) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *SpaceProperty) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetCreatedAt

`func (o *SpaceProperty) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SpaceProperty) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SpaceProperty) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SpaceProperty) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *SpaceProperty) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SpaceProperty) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SpaceProperty) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SpaceProperty) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetVersion

`func (o *SpaceProperty) GetVersion() SpacePropertyVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SpaceProperty) GetVersionOk() (*SpacePropertyVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SpaceProperty) SetVersion(v SpacePropertyVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SpaceProperty) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



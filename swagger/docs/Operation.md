# Operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | Pointer to **string** | The type of operation. | [optional] 
**TargetType** | Pointer to **string** | The type of entity the operation type targets. | [optional] 

## Methods

### NewOperation

`func NewOperation() *Operation`

NewOperation instantiates a new Operation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationWithDefaults

`func NewOperationWithDefaults() *Operation`

NewOperationWithDefaults instantiates a new Operation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *Operation) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *Operation) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *Operation) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *Operation) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetTargetType

`func (o *Operation) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *Operation) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *Operation) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *Operation) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Principal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalType** | Pointer to [**PrincipalType**](PrincipalType.md) |  | [optional] 
**PrincipalId** | Pointer to **string** | The principal ID. | [optional] 

## Methods

### NewPrincipal

`func NewPrincipal() *Principal`

NewPrincipal instantiates a new Principal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalWithDefaults

`func NewPrincipalWithDefaults() *Principal`

NewPrincipalWithDefaults instantiates a new Principal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalType

`func (o *Principal) GetPrincipalType() PrincipalType`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *Principal) GetPrincipalTypeOk() (*PrincipalType, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *Principal) SetPrincipalType(v PrincipalType)`

SetPrincipalType sets PrincipalType field to given value.

### HasPrincipalType

`func (o *Principal) HasPrincipalType() bool`

HasPrincipalType returns a boolean if a field has been set.

### GetPrincipalId

`func (o *Principal) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *Principal) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *Principal) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.

### HasPrincipalId

`func (o *Principal) HasPrincipalId() bool`

HasPrincipalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



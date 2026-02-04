# EnableAdminKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationInMinutes** | Pointer to **int32** | The requested duration of admin key access in minutes, up to a maximum of 60 minutes, after which the issued admin key will automatically expire. | [optional] 

## Methods

### NewEnableAdminKeyRequest

`func NewEnableAdminKeyRequest() *EnableAdminKeyRequest`

NewEnableAdminKeyRequest instantiates a new EnableAdminKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableAdminKeyRequestWithDefaults

`func NewEnableAdminKeyRequestWithDefaults() *EnableAdminKeyRequest`

NewEnableAdminKeyRequestWithDefaults instantiates a new EnableAdminKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationInMinutes

`func (o *EnableAdminKeyRequest) GetDurationInMinutes() int32`

GetDurationInMinutes returns the DurationInMinutes field if non-nil, zero value otherwise.

### GetDurationInMinutesOk

`func (o *EnableAdminKeyRequest) GetDurationInMinutesOk() (*int32, bool)`

GetDurationInMinutesOk returns a tuple with the DurationInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMinutes

`func (o *EnableAdminKeyRequest) SetDurationInMinutes(v int32)`

SetDurationInMinutes sets DurationInMinutes field to given value.

### HasDurationInMinutes

`func (o *EnableAdminKeyRequest) HasDurationInMinutes() bool`

HasDurationInMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AdminKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | User identifier. | [optional] 
**ExpirationTime** | Pointer to **time.Time** | Timestamp in UTC that represents when the admin key will expire. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 

## Methods

### NewAdminKeyResponse

`func NewAdminKeyResponse() *AdminKeyResponse`

NewAdminKeyResponse instantiates a new AdminKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminKeyResponseWithDefaults

`func NewAdminKeyResponseWithDefaults() *AdminKeyResponse`

NewAdminKeyResponseWithDefaults instantiates a new AdminKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AdminKeyResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AdminKeyResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AdminKeyResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AdminKeyResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetExpirationTime

`func (o *AdminKeyResponse) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *AdminKeyResponse) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *AdminKeyResponse) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *AdminKeyResponse) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



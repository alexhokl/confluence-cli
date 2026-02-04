# CreateCustomContentRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Representation** | Pointer to **string** | Type of content representation used for the value field. | [optional] 
**Value** | Pointer to **string** | Body of the custom content, in the format found in the representation field. | [optional] 
**Storage** | Pointer to [**CustomContentBodyWrite**](CustomContentBodyWrite.md) |  | [optional] 
**AtlasDocFormat** | Pointer to [**CustomContentBodyWrite**](CustomContentBodyWrite.md) |  | [optional] 
**Raw** | Pointer to [**CustomContentBodyWrite**](CustomContentBodyWrite.md) |  | [optional] 

## Methods

### NewCreateCustomContentRequestBody

`func NewCreateCustomContentRequestBody() *CreateCustomContentRequestBody`

NewCreateCustomContentRequestBody instantiates a new CreateCustomContentRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomContentRequestBodyWithDefaults

`func NewCreateCustomContentRequestBodyWithDefaults() *CreateCustomContentRequestBody`

NewCreateCustomContentRequestBodyWithDefaults instantiates a new CreateCustomContentRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepresentation

`func (o *CreateCustomContentRequestBody) GetRepresentation() string`

GetRepresentation returns the Representation field if non-nil, zero value otherwise.

### GetRepresentationOk

`func (o *CreateCustomContentRequestBody) GetRepresentationOk() (*string, bool)`

GetRepresentationOk returns a tuple with the Representation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepresentation

`func (o *CreateCustomContentRequestBody) SetRepresentation(v string)`

SetRepresentation sets Representation field to given value.

### HasRepresentation

`func (o *CreateCustomContentRequestBody) HasRepresentation() bool`

HasRepresentation returns a boolean if a field has been set.

### GetValue

`func (o *CreateCustomContentRequestBody) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateCustomContentRequestBody) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateCustomContentRequestBody) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CreateCustomContentRequestBody) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetStorage

`func (o *CreateCustomContentRequestBody) GetStorage() CustomContentBodyWrite`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CreateCustomContentRequestBody) GetStorageOk() (*CustomContentBodyWrite, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CreateCustomContentRequestBody) SetStorage(v CustomContentBodyWrite)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *CreateCustomContentRequestBody) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetAtlasDocFormat

`func (o *CreateCustomContentRequestBody) GetAtlasDocFormat() CustomContentBodyWrite`

GetAtlasDocFormat returns the AtlasDocFormat field if non-nil, zero value otherwise.

### GetAtlasDocFormatOk

`func (o *CreateCustomContentRequestBody) GetAtlasDocFormatOk() (*CustomContentBodyWrite, bool)`

GetAtlasDocFormatOk returns a tuple with the AtlasDocFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasDocFormat

`func (o *CreateCustomContentRequestBody) SetAtlasDocFormat(v CustomContentBodyWrite)`

SetAtlasDocFormat sets AtlasDocFormat field to given value.

### HasAtlasDocFormat

`func (o *CreateCustomContentRequestBody) HasAtlasDocFormat() bool`

HasAtlasDocFormat returns a boolean if a field has been set.

### GetRaw

`func (o *CreateCustomContentRequestBody) GetRaw() CustomContentBodyWrite`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *CreateCustomContentRequestBody) GetRawOk() (*CustomContentBodyWrite, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *CreateCustomContentRequestBody) SetRaw(v CustomContentBodyWrite)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *CreateCustomContentRequestBody) HasRaw() bool`

HasRaw returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreatePageRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Representation** | Pointer to **string** | Type of content representation used for the value field. | [optional] 
**Value** | Pointer to **string** | Body of the page, in the format found in the representation field. | [optional] 
**Storage** | Pointer to [**PageBodyWrite**](PageBodyWrite.md) |  | [optional] 
**AtlasDocFormat** | Pointer to [**PageBodyWrite**](PageBodyWrite.md) |  | [optional] 
**Wiki** | Pointer to [**PageBodyWrite**](PageBodyWrite.md) |  | [optional] 

## Methods

### NewCreatePageRequestBody

`func NewCreatePageRequestBody() *CreatePageRequestBody`

NewCreatePageRequestBody instantiates a new CreatePageRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePageRequestBodyWithDefaults

`func NewCreatePageRequestBodyWithDefaults() *CreatePageRequestBody`

NewCreatePageRequestBodyWithDefaults instantiates a new CreatePageRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepresentation

`func (o *CreatePageRequestBody) GetRepresentation() string`

GetRepresentation returns the Representation field if non-nil, zero value otherwise.

### GetRepresentationOk

`func (o *CreatePageRequestBody) GetRepresentationOk() (*string, bool)`

GetRepresentationOk returns a tuple with the Representation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepresentation

`func (o *CreatePageRequestBody) SetRepresentation(v string)`

SetRepresentation sets Representation field to given value.

### HasRepresentation

`func (o *CreatePageRequestBody) HasRepresentation() bool`

HasRepresentation returns a boolean if a field has been set.

### GetValue

`func (o *CreatePageRequestBody) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreatePageRequestBody) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreatePageRequestBody) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CreatePageRequestBody) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetStorage

`func (o *CreatePageRequestBody) GetStorage() PageBodyWrite`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CreatePageRequestBody) GetStorageOk() (*PageBodyWrite, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CreatePageRequestBody) SetStorage(v PageBodyWrite)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *CreatePageRequestBody) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetAtlasDocFormat

`func (o *CreatePageRequestBody) GetAtlasDocFormat() PageBodyWrite`

GetAtlasDocFormat returns the AtlasDocFormat field if non-nil, zero value otherwise.

### GetAtlasDocFormatOk

`func (o *CreatePageRequestBody) GetAtlasDocFormatOk() (*PageBodyWrite, bool)`

GetAtlasDocFormatOk returns a tuple with the AtlasDocFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasDocFormat

`func (o *CreatePageRequestBody) SetAtlasDocFormat(v PageBodyWrite)`

SetAtlasDocFormat sets AtlasDocFormat field to given value.

### HasAtlasDocFormat

`func (o *CreatePageRequestBody) HasAtlasDocFormat() bool`

HasAtlasDocFormat returns a boolean if a field has been set.

### GetWiki

`func (o *CreatePageRequestBody) GetWiki() PageBodyWrite`

GetWiki returns the Wiki field if non-nil, zero value otherwise.

### GetWikiOk

`func (o *CreatePageRequestBody) GetWikiOk() (*PageBodyWrite, bool)`

GetWikiOk returns a tuple with the Wiki field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiki

`func (o *CreatePageRequestBody) SetWiki(v PageBodyWrite)`

SetWiki sets Wiki field to given value.

### HasWiki

`func (o *CreatePageRequestBody) HasWiki() bool`

HasWiki returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



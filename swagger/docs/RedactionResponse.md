# RedactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to [**RedactionSectionResponse**](RedactionSectionResponse.md) |  | [optional] 
**Title** | Pointer to [**RedactionSectionResponse**](RedactionSectionResponse.md) |  | [optional] 

## Methods

### NewRedactionResponse

`func NewRedactionResponse() *RedactionResponse`

NewRedactionResponse instantiates a new RedactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedactionResponseWithDefaults

`func NewRedactionResponseWithDefaults() *RedactionResponse`

NewRedactionResponseWithDefaults instantiates a new RedactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *RedactionResponse) GetBody() RedactionSectionResponse`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *RedactionResponse) GetBodyOk() (*RedactionSectionResponse, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *RedactionResponse) SetBody(v RedactionSectionResponse)`

SetBody sets Body field to given value.

### HasBody

`func (o *RedactionResponse) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetTitle

`func (o *RedactionResponse) GetTitle() RedactionSectionResponse`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RedactionResponse) GetTitleOk() (*RedactionSectionResponse, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RedactionResponse) SetTitle(v RedactionSectionResponse)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RedactionResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



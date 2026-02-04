# PostRedactPageRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Redactions** | Pointer to [**[]RedactionPointer**](RedactionPointer.md) |  | [optional] 

## Methods

### NewPostRedactPageRequestBody

`func NewPostRedactPageRequestBody() *PostRedactPageRequestBody`

NewPostRedactPageRequestBody instantiates a new PostRedactPageRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRedactPageRequestBodyWithDefaults

`func NewPostRedactPageRequestBodyWithDefaults() *PostRedactPageRequestBody`

NewPostRedactPageRequestBodyWithDefaults instantiates a new PostRedactPageRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedactions

`func (o *PostRedactPageRequestBody) GetRedactions() []RedactionPointer`

GetRedactions returns the Redactions field if non-nil, zero value otherwise.

### GetRedactionsOk

`func (o *PostRedactPageRequestBody) GetRedactionsOk() (*[]RedactionPointer, bool)`

GetRedactionsOk returns a tuple with the Redactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactions

`func (o *PostRedactPageRequestBody) SetRedactions(v []RedactionPointer)`

SetRedactions sets Redactions field to given value.

### HasRedactions

`func (o *PostRedactPageRequestBody) HasRedactions() bool`

HasRedactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



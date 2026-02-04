# RedactionSectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Redactions** | Pointer to [**[]RedactionPointerResponse**](RedactionPointerResponse.md) | List of redactions that were applied to this section | [optional] 

## Methods

### NewRedactionSectionResponse

`func NewRedactionSectionResponse() *RedactionSectionResponse`

NewRedactionSectionResponse instantiates a new RedactionSectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedactionSectionResponseWithDefaults

`func NewRedactionSectionResponseWithDefaults() *RedactionSectionResponse`

NewRedactionSectionResponseWithDefaults instantiates a new RedactionSectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedactions

`func (o *RedactionSectionResponse) GetRedactions() []RedactionPointerResponse`

GetRedactions returns the Redactions field if non-nil, zero value otherwise.

### GetRedactionsOk

`func (o *RedactionSectionResponse) GetRedactionsOk() (*[]RedactionPointerResponse, bool)`

GetRedactionsOk returns a tuple with the Redactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactions

`func (o *RedactionSectionResponse) SetRedactions(v []RedactionPointerResponse)`

SetRedactions sets Redactions field to given value.

### HasRedactions

`func (o *RedactionSectionResponse) HasRedactions() bool`

HasRedactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



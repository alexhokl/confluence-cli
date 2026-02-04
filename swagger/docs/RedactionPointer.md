# RedactionPointer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pointer** | **string** | JSON pointer indicating the exact location within the content structure  where redaction should be applied. Points to the text node containing the content to redact.  | 
**From** | Pointer to **int32** | Starting character index (zero-based) within the target text where redaction begins.  | [optional] 
**To** | Pointer to **int32** | Ending character index (zero-based) within the target text where redaction ends (exclusive). Must be greater than or equal to &#39;from&#39; value.  | [optional] 
**Reason** | Pointer to **NullableString** | Optional human-readable reason for the redaction. Used for audit trails and compliance documentation.  | [optional] 

## Methods

### NewRedactionPointer

`func NewRedactionPointer(pointer string, ) *RedactionPointer`

NewRedactionPointer instantiates a new RedactionPointer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedactionPointerWithDefaults

`func NewRedactionPointerWithDefaults() *RedactionPointer`

NewRedactionPointerWithDefaults instantiates a new RedactionPointer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPointer

`func (o *RedactionPointer) GetPointer() string`

GetPointer returns the Pointer field if non-nil, zero value otherwise.

### GetPointerOk

`func (o *RedactionPointer) GetPointerOk() (*string, bool)`

GetPointerOk returns a tuple with the Pointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointer

`func (o *RedactionPointer) SetPointer(v string)`

SetPointer sets Pointer field to given value.


### GetFrom

`func (o *RedactionPointer) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *RedactionPointer) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *RedactionPointer) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *RedactionPointer) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *RedactionPointer) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *RedactionPointer) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *RedactionPointer) SetTo(v int32)`

SetTo sets To field to given value.

### HasTo

`func (o *RedactionPointer) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetReason

`func (o *RedactionPointer) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RedactionPointer) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RedactionPointer) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RedactionPointer) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *RedactionPointer) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *RedactionPointer) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



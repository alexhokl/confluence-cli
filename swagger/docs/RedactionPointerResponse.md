# RedactionPointerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pointer** | Pointer to **string** | JSON pointer indicating where the redaction was applied | [optional] 
**From** | Pointer to **int32** | Starting character index where redaction was applied | [optional] 
**To** | Pointer to **int32** | Ending character index where redaction was applied | [optional] 
**Reason** | Pointer to **string** | Reason for the redaction | [optional] 
**RedactionId** | Pointer to **string** | Unique identifier for this redaction. Can be used to restore the redacted content later.  | [optional] 

## Methods

### NewRedactionPointerResponse

`func NewRedactionPointerResponse() *RedactionPointerResponse`

NewRedactionPointerResponse instantiates a new RedactionPointerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedactionPointerResponseWithDefaults

`func NewRedactionPointerResponseWithDefaults() *RedactionPointerResponse`

NewRedactionPointerResponseWithDefaults instantiates a new RedactionPointerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPointer

`func (o *RedactionPointerResponse) GetPointer() string`

GetPointer returns the Pointer field if non-nil, zero value otherwise.

### GetPointerOk

`func (o *RedactionPointerResponse) GetPointerOk() (*string, bool)`

GetPointerOk returns a tuple with the Pointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointer

`func (o *RedactionPointerResponse) SetPointer(v string)`

SetPointer sets Pointer field to given value.

### HasPointer

`func (o *RedactionPointerResponse) HasPointer() bool`

HasPointer returns a boolean if a field has been set.

### GetFrom

`func (o *RedactionPointerResponse) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *RedactionPointerResponse) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *RedactionPointerResponse) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *RedactionPointerResponse) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *RedactionPointerResponse) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *RedactionPointerResponse) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *RedactionPointerResponse) SetTo(v int32)`

SetTo sets To field to given value.

### HasTo

`func (o *RedactionPointerResponse) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetReason

`func (o *RedactionPointerResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RedactionPointerResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RedactionPointerResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RedactionPointerResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRedactionId

`func (o *RedactionPointerResponse) GetRedactionId() string`

GetRedactionId returns the RedactionId field if non-nil, zero value otherwise.

### GetRedactionIdOk

`func (o *RedactionPointerResponse) GetRedactionIdOk() (*string, bool)`

GetRedactionIdOk returns a tuple with the RedactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactionId

`func (o *RedactionPointerResponse) SetRedactionId(v string)`

SetRedactionId sets RedactionId field to given value.

### HasRedactionId

`func (o *RedactionPointerResponse) HasRedactionId() bool`

HasRedactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



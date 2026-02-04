# MultiEntityResultChildrenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]ChildrenResponse**](ChildrenResponse.md) |  | [optional] 
**Links** | Pointer to [**MultiEntityLinks**](MultiEntityLinks.md) |  | [optional] 

## Methods

### NewMultiEntityResultChildrenResponse

`func NewMultiEntityResultChildrenResponse() *MultiEntityResultChildrenResponse`

NewMultiEntityResultChildrenResponse instantiates a new MultiEntityResultChildrenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiEntityResultChildrenResponseWithDefaults

`func NewMultiEntityResultChildrenResponseWithDefaults() *MultiEntityResultChildrenResponse`

NewMultiEntityResultChildrenResponseWithDefaults instantiates a new MultiEntityResultChildrenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *MultiEntityResultChildrenResponse) GetResults() []ChildrenResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *MultiEntityResultChildrenResponse) GetResultsOk() (*[]ChildrenResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *MultiEntityResultChildrenResponse) SetResults(v []ChildrenResponse)`

SetResults sets Results field to given value.

### HasResults

`func (o *MultiEntityResultChildrenResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetLinks

`func (o *MultiEntityResultChildrenResponse) GetLinks() MultiEntityLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MultiEntityResultChildrenResponse) GetLinksOk() (*MultiEntityLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MultiEntityResultChildrenResponse) SetLinks(v MultiEntityLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MultiEntityResultChildrenResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MultiEntityResultVersion2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]PageVersion**](PageVersion.md) |  | [optional] 
**Links** | Pointer to [**MultiEntityLinks**](MultiEntityLinks.md) |  | [optional] 

## Methods

### NewMultiEntityResultVersion2

`func NewMultiEntityResultVersion2() *MultiEntityResultVersion2`

NewMultiEntityResultVersion2 instantiates a new MultiEntityResultVersion2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiEntityResultVersion2WithDefaults

`func NewMultiEntityResultVersion2WithDefaults() *MultiEntityResultVersion2`

NewMultiEntityResultVersion2WithDefaults instantiates a new MultiEntityResultVersion2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *MultiEntityResultVersion2) GetResults() []PageVersion`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *MultiEntityResultVersion2) GetResultsOk() (*[]PageVersion, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *MultiEntityResultVersion2) SetResults(v []PageVersion)`

SetResults sets Results field to given value.

### HasResults

`func (o *MultiEntityResultVersion2) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetLinks

`func (o *MultiEntityResultVersion2) GetLinks() MultiEntityLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MultiEntityResultVersion2) GetLinksOk() (*MultiEntityLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MultiEntityResultVersion2) SetLinks(v MultiEntityLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MultiEntityResultVersion2) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MultiEntityResultSpaceProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]SpaceProperty**](SpaceProperty.md) |  | [optional] 
**Links** | Pointer to [**MultiEntityLinks**](MultiEntityLinks.md) |  | [optional] 

## Methods

### NewMultiEntityResultSpaceProperty

`func NewMultiEntityResultSpaceProperty() *MultiEntityResultSpaceProperty`

NewMultiEntityResultSpaceProperty instantiates a new MultiEntityResultSpaceProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiEntityResultSpacePropertyWithDefaults

`func NewMultiEntityResultSpacePropertyWithDefaults() *MultiEntityResultSpaceProperty`

NewMultiEntityResultSpacePropertyWithDefaults instantiates a new MultiEntityResultSpaceProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *MultiEntityResultSpaceProperty) GetResults() []SpaceProperty`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *MultiEntityResultSpaceProperty) GetResultsOk() (*[]SpaceProperty, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *MultiEntityResultSpaceProperty) SetResults(v []SpaceProperty)`

SetResults sets Results field to given value.

### HasResults

`func (o *MultiEntityResultSpaceProperty) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetLinks

`func (o *MultiEntityResultSpaceProperty) GetLinks() MultiEntityLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MultiEntityResultSpaceProperty) GetLinksOk() (*MultiEntityLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MultiEntityResultSpaceProperty) SetLinks(v MultiEntityLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MultiEntityResultSpaceProperty) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



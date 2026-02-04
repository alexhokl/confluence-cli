# SpaceSingleProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]SpaceProperty**](SpaceProperty.md) |  | [optional] 
**Meta** | Pointer to [**OptionalFieldMeta**](OptionalFieldMeta.md) |  | [optional] 
**Links** | Pointer to [**OptionalFieldLinks**](OptionalFieldLinks.md) |  | [optional] 

## Methods

### NewSpaceSingleProperties

`func NewSpaceSingleProperties() *SpaceSingleProperties`

NewSpaceSingleProperties instantiates a new SpaceSingleProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceSinglePropertiesWithDefaults

`func NewSpaceSinglePropertiesWithDefaults() *SpaceSingleProperties`

NewSpaceSinglePropertiesWithDefaults instantiates a new SpaceSingleProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *SpaceSingleProperties) GetResults() []SpaceProperty`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SpaceSingleProperties) GetResultsOk() (*[]SpaceProperty, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SpaceSingleProperties) SetResults(v []SpaceProperty)`

SetResults sets Results field to given value.

### HasResults

`func (o *SpaceSingleProperties) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetMeta

`func (o *SpaceSingleProperties) GetMeta() OptionalFieldMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SpaceSingleProperties) GetMetaOk() (*OptionalFieldMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SpaceSingleProperties) SetMeta(v OptionalFieldMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SpaceSingleProperties) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *SpaceSingleProperties) GetLinks() OptionalFieldLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SpaceSingleProperties) GetLinksOk() (*OptionalFieldLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SpaceSingleProperties) SetLinks(v OptionalFieldLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SpaceSingleProperties) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



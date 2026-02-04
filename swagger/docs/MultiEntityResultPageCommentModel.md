# MultiEntityResultPageCommentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]PageCommentModel**](PageCommentModel.md) |  | [optional] 
**Links** | Pointer to [**MultiEntityLinks**](MultiEntityLinks.md) |  | [optional] 

## Methods

### NewMultiEntityResultPageCommentModel

`func NewMultiEntityResultPageCommentModel() *MultiEntityResultPageCommentModel`

NewMultiEntityResultPageCommentModel instantiates a new MultiEntityResultPageCommentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiEntityResultPageCommentModelWithDefaults

`func NewMultiEntityResultPageCommentModelWithDefaults() *MultiEntityResultPageCommentModel`

NewMultiEntityResultPageCommentModelWithDefaults instantiates a new MultiEntityResultPageCommentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *MultiEntityResultPageCommentModel) GetResults() []PageCommentModel`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *MultiEntityResultPageCommentModel) GetResultsOk() (*[]PageCommentModel, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *MultiEntityResultPageCommentModel) SetResults(v []PageCommentModel)`

SetResults sets Results field to given value.

### HasResults

`func (o *MultiEntityResultPageCommentModel) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetLinks

`func (o *MultiEntityResultPageCommentModel) GetLinks() MultiEntityLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MultiEntityResultPageCommentModel) GetLinksOk() (*MultiEntityLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MultiEntityResultPageCommentModel) SetLinks(v MultiEntityLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MultiEntityResultPageCommentModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



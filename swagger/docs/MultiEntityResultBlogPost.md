# MultiEntityResultBlogPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]BlogPostBulk**](BlogPostBulk.md) |  | [optional] 
**Links** | Pointer to [**MultiEntityLinks**](MultiEntityLinks.md) |  | [optional] 

## Methods

### NewMultiEntityResultBlogPost

`func NewMultiEntityResultBlogPost() *MultiEntityResultBlogPost`

NewMultiEntityResultBlogPost instantiates a new MultiEntityResultBlogPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiEntityResultBlogPostWithDefaults

`func NewMultiEntityResultBlogPostWithDefaults() *MultiEntityResultBlogPost`

NewMultiEntityResultBlogPostWithDefaults instantiates a new MultiEntityResultBlogPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *MultiEntityResultBlogPost) GetResults() []BlogPostBulk`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *MultiEntityResultBlogPost) GetResultsOk() (*[]BlogPostBulk, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *MultiEntityResultBlogPost) SetResults(v []BlogPostBulk)`

SetResults sets Results field to given value.

### HasResults

`func (o *MultiEntityResultBlogPost) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetLinks

`func (o *MultiEntityResultBlogPost) GetLinks() MultiEntityLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MultiEntityResultBlogPost) GetLinksOk() (*MultiEntityLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MultiEntityResultBlogPost) SetLinks(v MultiEntityLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MultiEntityResultBlogPost) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



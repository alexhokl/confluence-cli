# InlineCommentModelProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]ContentProperty**](ContentProperty.md) |  | [optional] 
**Meta** | Pointer to [**OptionalFieldMeta**](OptionalFieldMeta.md) |  | [optional] 
**Links** | Pointer to [**OptionalFieldLinks**](OptionalFieldLinks.md) |  | [optional] 
**InlineMarkerRef** | Pointer to **string** | Property value used to reference the highlighted element in DOM. | [optional] 
**InlineOriginalSelection** | Pointer to **string** | Text that is highlighted. | [optional] 

## Methods

### NewInlineCommentModelProperties

`func NewInlineCommentModelProperties() *InlineCommentModelProperties`

NewInlineCommentModelProperties instantiates a new InlineCommentModelProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineCommentModelPropertiesWithDefaults

`func NewInlineCommentModelPropertiesWithDefaults() *InlineCommentModelProperties`

NewInlineCommentModelPropertiesWithDefaults instantiates a new InlineCommentModelProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *InlineCommentModelProperties) GetResults() []ContentProperty`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *InlineCommentModelProperties) GetResultsOk() (*[]ContentProperty, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *InlineCommentModelProperties) SetResults(v []ContentProperty)`

SetResults sets Results field to given value.

### HasResults

`func (o *InlineCommentModelProperties) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetMeta

`func (o *InlineCommentModelProperties) GetMeta() OptionalFieldMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineCommentModelProperties) GetMetaOk() (*OptionalFieldMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineCommentModelProperties) SetMeta(v OptionalFieldMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineCommentModelProperties) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *InlineCommentModelProperties) GetLinks() OptionalFieldLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InlineCommentModelProperties) GetLinksOk() (*OptionalFieldLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InlineCommentModelProperties) SetLinks(v OptionalFieldLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *InlineCommentModelProperties) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetInlineMarkerRef

`func (o *InlineCommentModelProperties) GetInlineMarkerRef() string`

GetInlineMarkerRef returns the InlineMarkerRef field if non-nil, zero value otherwise.

### GetInlineMarkerRefOk

`func (o *InlineCommentModelProperties) GetInlineMarkerRefOk() (*string, bool)`

GetInlineMarkerRefOk returns a tuple with the InlineMarkerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineMarkerRef

`func (o *InlineCommentModelProperties) SetInlineMarkerRef(v string)`

SetInlineMarkerRef sets InlineMarkerRef field to given value.

### HasInlineMarkerRef

`func (o *InlineCommentModelProperties) HasInlineMarkerRef() bool`

HasInlineMarkerRef returns a boolean if a field has been set.

### GetInlineOriginalSelection

`func (o *InlineCommentModelProperties) GetInlineOriginalSelection() string`

GetInlineOriginalSelection returns the InlineOriginalSelection field if non-nil, zero value otherwise.

### GetInlineOriginalSelectionOk

`func (o *InlineCommentModelProperties) GetInlineOriginalSelectionOk() (*string, bool)`

GetInlineOriginalSelectionOk returns a tuple with the InlineOriginalSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineOriginalSelection

`func (o *InlineCommentModelProperties) SetInlineOriginalSelection(v string)`

SetInlineOriginalSelection sets InlineOriginalSelection field to given value.

### HasInlineOriginalSelection

`func (o *InlineCommentModelProperties) HasInlineOriginalSelection() bool`

HasInlineOriginalSelection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



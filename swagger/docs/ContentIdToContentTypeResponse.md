# ContentIdToContentTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**map[string]ContentIdToContentTypeResponseResultsValue**](ContentIdToContentTypeResponseResultsValue.md) | JSON object containing all requested content ids as keys and their associated content types as the values. Duplicate content ids in the request will be returned under a single key in the response. For built-in content types, the enumerations are as specified. Custom content ids will be mapped to their associated type. | [optional] 

## Methods

### NewContentIdToContentTypeResponse

`func NewContentIdToContentTypeResponse() *ContentIdToContentTypeResponse`

NewContentIdToContentTypeResponse instantiates a new ContentIdToContentTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentIdToContentTypeResponseWithDefaults

`func NewContentIdToContentTypeResponseWithDefaults() *ContentIdToContentTypeResponse`

NewContentIdToContentTypeResponseWithDefaults instantiates a new ContentIdToContentTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ContentIdToContentTypeResponse) GetResults() map[string]ContentIdToContentTypeResponseResultsValue`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ContentIdToContentTypeResponse) GetResultsOk() (*map[string]ContentIdToContentTypeResponseResultsValue, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ContentIdToContentTypeResponse) SetResults(v map[string]ContentIdToContentTypeResponseResultsValue)`

SetResults sets Results field to given value.

### HasResults

`func (o *ContentIdToContentTypeResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



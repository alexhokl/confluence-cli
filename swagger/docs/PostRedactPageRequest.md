# PostRedactPageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Timestamp when the content was last updated. | 
**CleanHistory** | Pointer to **NullableBool** | Whether to clean up previous versions containing the redaction. When true, historical versions of the content that contain the redacted text will be squashed. | [optional] 
**VersionNumber** | Pointer to **NullableInt32** | Optional version number of the content to redact. When specified, the redaction will target  a specific historical version of the content rather than the current version.  - If omitted or null, the redaction applies to the current (latest) version of the content. - When provided, must be a valid version number that exists for the content.  **Note**: Version numbers start at 1 and increment with each content update.  | [optional] 
**Body** | Pointer to [**PostRedactPageRequestBody**](PostRedactPageRequestBody.md) |  | [optional] 
**Title** | Pointer to [**PostRedactPageRequestBody**](PostRedactPageRequestBody.md) |  | [optional] 

## Methods

### NewPostRedactPageRequest

`func NewPostRedactPageRequest(createdAt time.Time, ) *PostRedactPageRequest`

NewPostRedactPageRequest instantiates a new PostRedactPageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRedactPageRequestWithDefaults

`func NewPostRedactPageRequestWithDefaults() *PostRedactPageRequest`

NewPostRedactPageRequestWithDefaults instantiates a new PostRedactPageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PostRedactPageRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PostRedactPageRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PostRedactPageRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCleanHistory

`func (o *PostRedactPageRequest) GetCleanHistory() bool`

GetCleanHistory returns the CleanHistory field if non-nil, zero value otherwise.

### GetCleanHistoryOk

`func (o *PostRedactPageRequest) GetCleanHistoryOk() (*bool, bool)`

GetCleanHistoryOk returns a tuple with the CleanHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanHistory

`func (o *PostRedactPageRequest) SetCleanHistory(v bool)`

SetCleanHistory sets CleanHistory field to given value.

### HasCleanHistory

`func (o *PostRedactPageRequest) HasCleanHistory() bool`

HasCleanHistory returns a boolean if a field has been set.

### SetCleanHistoryNil

`func (o *PostRedactPageRequest) SetCleanHistoryNil(b bool)`

 SetCleanHistoryNil sets the value for CleanHistory to be an explicit nil

### UnsetCleanHistory
`func (o *PostRedactPageRequest) UnsetCleanHistory()`

UnsetCleanHistory ensures that no value is present for CleanHistory, not even an explicit nil
### GetVersionNumber

`func (o *PostRedactPageRequest) GetVersionNumber() int32`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *PostRedactPageRequest) GetVersionNumberOk() (*int32, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *PostRedactPageRequest) SetVersionNumber(v int32)`

SetVersionNumber sets VersionNumber field to given value.

### HasVersionNumber

`func (o *PostRedactPageRequest) HasVersionNumber() bool`

HasVersionNumber returns a boolean if a field has been set.

### SetVersionNumberNil

`func (o *PostRedactPageRequest) SetVersionNumberNil(b bool)`

 SetVersionNumberNil sets the value for VersionNumber to be an explicit nil

### UnsetVersionNumber
`func (o *PostRedactPageRequest) UnsetVersionNumber()`

UnsetVersionNumber ensures that no value is present for VersionNumber, not even an explicit nil
### GetBody

`func (o *PostRedactPageRequest) GetBody() PostRedactPageRequestBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PostRedactPageRequest) GetBodyOk() (*PostRedactPageRequestBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PostRedactPageRequest) SetBody(v PostRedactPageRequestBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *PostRedactPageRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetTitle

`func (o *PostRedactPageRequest) GetTitle() PostRedactPageRequestBody`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PostRedactPageRequest) GetTitleOk() (*PostRedactPageRequestBody, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PostRedactPageRequest) SetTitle(v PostRedactPageRequestBody)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PostRedactPageRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



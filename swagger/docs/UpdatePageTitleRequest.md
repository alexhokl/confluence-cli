# UpdatePageTitleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The status of the page, current or draft. | 
**Title** | **string** | The updated title for the page | 

## Methods

### NewUpdatePageTitleRequest

`func NewUpdatePageTitleRequest(status string, title string, ) *UpdatePageTitleRequest`

NewUpdatePageTitleRequest instantiates a new UpdatePageTitleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePageTitleRequestWithDefaults

`func NewUpdatePageTitleRequestWithDefaults() *UpdatePageTitleRequest`

NewUpdatePageTitleRequestWithDefaults instantiates a new UpdatePageTitleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdatePageTitleRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdatePageTitleRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdatePageTitleRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *UpdatePageTitleRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdatePageTitleRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdatePageTitleRequest) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateWhiteboardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpaceId** | **string** | ID of the space. | 
**Title** | Pointer to **string** | Title of the whiteboard. | [optional] 
**ParentId** | Pointer to **string** | The parent content ID of the whiteboard. | [optional] 
**TemplateKey** | Pointer to **string** | Providing a template key will add that template to the new whiteboard. | [optional] 
**Locale** | Pointer to **string** | If templateKey is provided, locale will decide which language the template will be created with. If locale is omitted, the user&#39;s locale will be used. | [optional] 

## Methods

### NewCreateWhiteboardRequest

`func NewCreateWhiteboardRequest(spaceId string, ) *CreateWhiteboardRequest`

NewCreateWhiteboardRequest instantiates a new CreateWhiteboardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWhiteboardRequestWithDefaults

`func NewCreateWhiteboardRequestWithDefaults() *CreateWhiteboardRequest`

NewCreateWhiteboardRequestWithDefaults instantiates a new CreateWhiteboardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpaceId

`func (o *CreateWhiteboardRequest) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *CreateWhiteboardRequest) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *CreateWhiteboardRequest) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetTitle

`func (o *CreateWhiteboardRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateWhiteboardRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateWhiteboardRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateWhiteboardRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetParentId

`func (o *CreateWhiteboardRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateWhiteboardRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateWhiteboardRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateWhiteboardRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetTemplateKey

`func (o *CreateWhiteboardRequest) GetTemplateKey() string`

GetTemplateKey returns the TemplateKey field if non-nil, zero value otherwise.

### GetTemplateKeyOk

`func (o *CreateWhiteboardRequest) GetTemplateKeyOk() (*string, bool)`

GetTemplateKeyOk returns a tuple with the TemplateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateKey

`func (o *CreateWhiteboardRequest) SetTemplateKey(v string)`

SetTemplateKey sets TemplateKey field to given value.

### HasTemplateKey

`func (o *CreateWhiteboardRequest) HasTemplateKey() bool`

HasTemplateKey returns a boolean if a field has been set.

### GetLocale

`func (o *CreateWhiteboardRequest) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *CreateWhiteboardRequest) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *CreateWhiteboardRequest) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *CreateWhiteboardRequest) HasLocale() bool`

HasLocale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



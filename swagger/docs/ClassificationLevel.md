# ClassificationLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the classification level. | [optional] 
**Status** | Pointer to [**ClassificationLevelStatus**](ClassificationLevelStatus.md) |  | [optional] 
**Order** | Pointer to **float32** | The order of the classification level object. | [optional] 
**Name** | Pointer to **string** | The name of the classification level object. | [optional] 
**Description** | Pointer to **string** | The description of the classification level object. | [optional] 
**Guideline** | Pointer to **string** | The guideline of the classification level object. | [optional] 
**Color** | Pointer to [**ClassificationLevelColor**](ClassificationLevelColor.md) |  | [optional] 

## Methods

### NewClassificationLevel

`func NewClassificationLevel() *ClassificationLevel`

NewClassificationLevel instantiates a new ClassificationLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassificationLevelWithDefaults

`func NewClassificationLevelWithDefaults() *ClassificationLevel`

NewClassificationLevelWithDefaults instantiates a new ClassificationLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClassificationLevel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClassificationLevel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClassificationLevel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClassificationLevel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ClassificationLevel) GetStatus() ClassificationLevelStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClassificationLevel) GetStatusOk() (*ClassificationLevelStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClassificationLevel) SetStatus(v ClassificationLevelStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClassificationLevel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOrder

`func (o *ClassificationLevel) GetOrder() float32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ClassificationLevel) GetOrderOk() (*float32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ClassificationLevel) SetOrder(v float32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ClassificationLevel) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetName

`func (o *ClassificationLevel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClassificationLevel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClassificationLevel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClassificationLevel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ClassificationLevel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClassificationLevel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClassificationLevel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClassificationLevel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGuideline

`func (o *ClassificationLevel) GetGuideline() string`

GetGuideline returns the Guideline field if non-nil, zero value otherwise.

### GetGuidelineOk

`func (o *ClassificationLevel) GetGuidelineOk() (*string, bool)`

GetGuidelineOk returns a tuple with the Guideline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuideline

`func (o *ClassificationLevel) SetGuideline(v string)`

SetGuideline sets Guideline field to given value.

### HasGuideline

`func (o *ClassificationLevel) HasGuideline() bool`

HasGuideline returns a boolean if a field has been set.

### GetColor

`func (o *ClassificationLevel) GetColor() ClassificationLevelColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ClassificationLevel) GetColorOk() (*ClassificationLevelColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ClassificationLevel) SetColor(v ClassificationLevelColor)`

SetColor sets Color field to given value.

### HasColor

`func (o *ClassificationLevel) HasColor() bool`

HasColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



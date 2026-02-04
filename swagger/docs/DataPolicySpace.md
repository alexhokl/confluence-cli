# DataPolicySpace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the space. | [optional] 
**Key** | Pointer to **string** | Key of the space. | [optional] 
**Name** | Pointer to **string** | Name of the space. | [optional] 
**Description** | Pointer to [**SpaceDescription**](SpaceDescription.md) |  | [optional] 
**DataPolicy** | Pointer to [**DataPolicySpaceDataPolicy**](DataPolicySpaceDataPolicy.md) |  | [optional] 
**Icon** | Pointer to [**SpaceIcon**](SpaceIcon.md) |  | [optional] 
**Links** | Pointer to [**SpaceLinks**](SpaceLinks.md) |  | [optional] 

## Methods

### NewDataPolicySpace

`func NewDataPolicySpace() *DataPolicySpace`

NewDataPolicySpace instantiates a new DataPolicySpace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataPolicySpaceWithDefaults

`func NewDataPolicySpaceWithDefaults() *DataPolicySpace`

NewDataPolicySpaceWithDefaults instantiates a new DataPolicySpace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataPolicySpace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataPolicySpace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataPolicySpace) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataPolicySpace) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *DataPolicySpace) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DataPolicySpace) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DataPolicySpace) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *DataPolicySpace) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *DataPolicySpace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataPolicySpace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataPolicySpace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataPolicySpace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *DataPolicySpace) GetDescription() SpaceDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataPolicySpace) GetDescriptionOk() (*SpaceDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataPolicySpace) SetDescription(v SpaceDescription)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DataPolicySpace) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDataPolicy

`func (o *DataPolicySpace) GetDataPolicy() DataPolicySpaceDataPolicy`

GetDataPolicy returns the DataPolicy field if non-nil, zero value otherwise.

### GetDataPolicyOk

`func (o *DataPolicySpace) GetDataPolicyOk() (*DataPolicySpaceDataPolicy, bool)`

GetDataPolicyOk returns a tuple with the DataPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPolicy

`func (o *DataPolicySpace) SetDataPolicy(v DataPolicySpaceDataPolicy)`

SetDataPolicy sets DataPolicy field to given value.

### HasDataPolicy

`func (o *DataPolicySpace) HasDataPolicy() bool`

HasDataPolicy returns a boolean if a field has been set.

### GetIcon

`func (o *DataPolicySpace) GetIcon() SpaceIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *DataPolicySpace) GetIconOk() (*SpaceIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *DataPolicySpace) SetIcon(v SpaceIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *DataPolicySpace) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetLinks

`func (o *DataPolicySpace) GetLinks() SpaceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DataPolicySpace) GetLinksOk() (*SpaceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DataPolicySpace) SetLinks(v SpaceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DataPolicySpace) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



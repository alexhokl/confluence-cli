# UpdateInlineCommentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to [**UpdateFooterCommentModelVersion**](UpdateFooterCommentModelVersion.md) |  | [optional] 
**Body** | Pointer to [**CreateFooterCommentModelBody**](CreateFooterCommentModelBody.md) |  | [optional] 
**Resolved** | Pointer to **bool** | Resolved state of the comment. Set to true to resolve the comment, set to false to reopen it. If matching the existing state (i.e. true -&gt; resolved or false -&gt; open/reopened) , no change will occur. A dangling comment cannot be updated. | [optional] 

## Methods

### NewUpdateInlineCommentModel

`func NewUpdateInlineCommentModel() *UpdateInlineCommentModel`

NewUpdateInlineCommentModel instantiates a new UpdateInlineCommentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInlineCommentModelWithDefaults

`func NewUpdateInlineCommentModelWithDefaults() *UpdateInlineCommentModel`

NewUpdateInlineCommentModelWithDefaults instantiates a new UpdateInlineCommentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *UpdateInlineCommentModel) GetVersion() UpdateFooterCommentModelVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateInlineCommentModel) GetVersionOk() (*UpdateFooterCommentModelVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateInlineCommentModel) SetVersion(v UpdateFooterCommentModelVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdateInlineCommentModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBody

`func (o *UpdateInlineCommentModel) GetBody() CreateFooterCommentModelBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *UpdateInlineCommentModel) GetBodyOk() (*CreateFooterCommentModelBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *UpdateInlineCommentModel) SetBody(v CreateFooterCommentModelBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *UpdateInlineCommentModel) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetResolved

`func (o *UpdateInlineCommentModel) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *UpdateInlineCommentModel) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *UpdateInlineCommentModel) SetResolved(v bool)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *UpdateInlineCommentModel) HasResolved() bool`

HasResolved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



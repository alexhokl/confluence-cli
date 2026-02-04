# UpdateFooterCommentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to [**UpdateFooterCommentModelVersion**](UpdateFooterCommentModelVersion.md) |  | [optional] 
**Body** | Pointer to [**CreateFooterCommentModelBody**](CreateFooterCommentModelBody.md) |  | [optional] 
**Links** | Pointer to [**GetAttachmentById200ResponseAllOfLinks**](GetAttachmentById200ResponseAllOfLinks.md) |  | [optional] 

## Methods

### NewUpdateFooterCommentRequest

`func NewUpdateFooterCommentRequest() *UpdateFooterCommentRequest`

NewUpdateFooterCommentRequest instantiates a new UpdateFooterCommentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFooterCommentRequestWithDefaults

`func NewUpdateFooterCommentRequestWithDefaults() *UpdateFooterCommentRequest`

NewUpdateFooterCommentRequestWithDefaults instantiates a new UpdateFooterCommentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *UpdateFooterCommentRequest) GetVersion() UpdateFooterCommentModelVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateFooterCommentRequest) GetVersionOk() (*UpdateFooterCommentModelVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateFooterCommentRequest) SetVersion(v UpdateFooterCommentModelVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdateFooterCommentRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBody

`func (o *UpdateFooterCommentRequest) GetBody() CreateFooterCommentModelBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *UpdateFooterCommentRequest) GetBodyOk() (*CreateFooterCommentModelBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *UpdateFooterCommentRequest) SetBody(v CreateFooterCommentModelBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *UpdateFooterCommentRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetLinks

`func (o *UpdateFooterCommentRequest) GetLinks() GetAttachmentById200ResponseAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UpdateFooterCommentRequest) GetLinksOk() (*GetAttachmentById200ResponseAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UpdateFooterCommentRequest) SetLinks(v GetAttachmentById200ResponseAllOfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UpdateFooterCommentRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



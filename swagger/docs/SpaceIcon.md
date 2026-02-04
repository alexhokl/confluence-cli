# SpaceIcon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | The path (relative to base URL) at which the space&#39;s icon can be retrieved. The format should be like &#x60;/wiki/download/...&#x60; or &#x60;/wiki/aa-avatar/...&#x60; | [optional] 
**ApiDownloadLink** | Pointer to **string** | The path (relative to base URL) that can be used to retrieve a link to download the space icon. 3LO apps should use this link instead of the value provided in the &#x60;path&#x60; property to retrieve the icon.  Currently this field is only returned for &#x60;global&#x60; spaces and not &#x60;personal&#x60; spaces.  | [optional] 

## Methods

### NewSpaceIcon

`func NewSpaceIcon() *SpaceIcon`

NewSpaceIcon instantiates a new SpaceIcon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceIconWithDefaults

`func NewSpaceIconWithDefaults() *SpaceIcon`

NewSpaceIconWithDefaults instantiates a new SpaceIcon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *SpaceIcon) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SpaceIcon) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SpaceIcon) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SpaceIcon) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetApiDownloadLink

`func (o *SpaceIcon) GetApiDownloadLink() string`

GetApiDownloadLink returns the ApiDownloadLink field if non-nil, zero value otherwise.

### GetApiDownloadLinkOk

`func (o *SpaceIcon) GetApiDownloadLinkOk() (*string, bool)`

GetApiDownloadLinkOk returns a tuple with the ApiDownloadLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiDownloadLink

`func (o *SpaceIcon) SetApiDownloadLink(v string)`

SetApiDownloadLink sets ApiDownloadLink field to given value.

### HasApiDownloadLink

`func (o *SpaceIcon) HasApiDownloadLink() bool`

HasApiDownloadLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



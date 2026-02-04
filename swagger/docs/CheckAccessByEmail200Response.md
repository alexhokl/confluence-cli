# CheckAccessByEmail200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailsWithoutAccess** | Pointer to **[]string** | List of emails that do not have access to site. | [optional] 
**InvalidEmails** | Pointer to **[]string** | List of invalid emails provided in the request. | [optional] 

## Methods

### NewCheckAccessByEmail200Response

`func NewCheckAccessByEmail200Response() *CheckAccessByEmail200Response`

NewCheckAccessByEmail200Response instantiates a new CheckAccessByEmail200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckAccessByEmail200ResponseWithDefaults

`func NewCheckAccessByEmail200ResponseWithDefaults() *CheckAccessByEmail200Response`

NewCheckAccessByEmail200ResponseWithDefaults instantiates a new CheckAccessByEmail200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailsWithoutAccess

`func (o *CheckAccessByEmail200Response) GetEmailsWithoutAccess() []string`

GetEmailsWithoutAccess returns the EmailsWithoutAccess field if non-nil, zero value otherwise.

### GetEmailsWithoutAccessOk

`func (o *CheckAccessByEmail200Response) GetEmailsWithoutAccessOk() (*[]string, bool)`

GetEmailsWithoutAccessOk returns a tuple with the EmailsWithoutAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailsWithoutAccess

`func (o *CheckAccessByEmail200Response) SetEmailsWithoutAccess(v []string)`

SetEmailsWithoutAccess sets EmailsWithoutAccess field to given value.

### HasEmailsWithoutAccess

`func (o *CheckAccessByEmail200Response) HasEmailsWithoutAccess() bool`

HasEmailsWithoutAccess returns a boolean if a field has been set.

### GetInvalidEmails

`func (o *CheckAccessByEmail200Response) GetInvalidEmails() []string`

GetInvalidEmails returns the InvalidEmails field if non-nil, zero value otherwise.

### GetInvalidEmailsOk

`func (o *CheckAccessByEmail200Response) GetInvalidEmailsOk() (*[]string, bool)`

GetInvalidEmailsOk returns a tuple with the InvalidEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidEmails

`func (o *CheckAccessByEmail200Response) SetInvalidEmails(v []string)`

SetInvalidEmails sets InvalidEmails field to given value.

### HasInvalidEmails

`func (o *CheckAccessByEmail200Response) HasInvalidEmails() bool`

HasInvalidEmails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



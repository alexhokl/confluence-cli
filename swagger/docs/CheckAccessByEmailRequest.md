# CheckAccessByEmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | **[]string** | List of emails to check access to site. | 

## Methods

### NewCheckAccessByEmailRequest

`func NewCheckAccessByEmailRequest(emails []string, ) *CheckAccessByEmailRequest`

NewCheckAccessByEmailRequest instantiates a new CheckAccessByEmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckAccessByEmailRequestWithDefaults

`func NewCheckAccessByEmailRequestWithDefaults() *CheckAccessByEmailRequest`

NewCheckAccessByEmailRequestWithDefaults instantiates a new CheckAccessByEmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *CheckAccessByEmailRequest) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *CheckAccessByEmailRequest) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *CheckAccessByEmailRequest) SetEmails(v []string)`

SetEmails sets Emails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



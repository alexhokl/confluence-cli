# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Display name of the user. | [optional] 
**TimeZone** | Pointer to **string** | Time zone of the user. Depending on the user&#39;s privacy setting, this may return null. | [optional] 
**PersonalSpaceId** | Pointer to **string** | Space ID of the user&#39;s personal space. Returns null, if no personal space for the user. | [optional] 
**IsExternalCollaborator** | Pointer to **bool** | Whether the user is an external collaborator. | [optional] 
**AccountStatus** | Pointer to [**AccountStatus**](AccountStatus.md) |  | [optional] 
**AccountId** | Pointer to **string** | Account ID of the user. | [optional] 
**Email** | Pointer to **string** | The email address of the user. Depending on the user&#39;s privacy setting, this may return an empty string. | [optional] 
**AccountType** | Pointer to [**AccountType**](AccountType.md) |  | [optional] 
**PublicName** | Pointer to **string** | Public name of the user. | [optional] 
**ProfilePicture** | Pointer to [**NullableIcon**](Icon.md) |  | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *User) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *User) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *User) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *User) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetTimeZone

`func (o *User) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *User) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *User) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *User) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetPersonalSpaceId

`func (o *User) GetPersonalSpaceId() string`

GetPersonalSpaceId returns the PersonalSpaceId field if non-nil, zero value otherwise.

### GetPersonalSpaceIdOk

`func (o *User) GetPersonalSpaceIdOk() (*string, bool)`

GetPersonalSpaceIdOk returns a tuple with the PersonalSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalSpaceId

`func (o *User) SetPersonalSpaceId(v string)`

SetPersonalSpaceId sets PersonalSpaceId field to given value.

### HasPersonalSpaceId

`func (o *User) HasPersonalSpaceId() bool`

HasPersonalSpaceId returns a boolean if a field has been set.

### GetIsExternalCollaborator

`func (o *User) GetIsExternalCollaborator() bool`

GetIsExternalCollaborator returns the IsExternalCollaborator field if non-nil, zero value otherwise.

### GetIsExternalCollaboratorOk

`func (o *User) GetIsExternalCollaboratorOk() (*bool, bool)`

GetIsExternalCollaboratorOk returns a tuple with the IsExternalCollaborator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternalCollaborator

`func (o *User) SetIsExternalCollaborator(v bool)`

SetIsExternalCollaborator sets IsExternalCollaborator field to given value.

### HasIsExternalCollaborator

`func (o *User) HasIsExternalCollaborator() bool`

HasIsExternalCollaborator returns a boolean if a field has been set.

### GetAccountStatus

`func (o *User) GetAccountStatus() AccountStatus`

GetAccountStatus returns the AccountStatus field if non-nil, zero value otherwise.

### GetAccountStatusOk

`func (o *User) GetAccountStatusOk() (*AccountStatus, bool)`

GetAccountStatusOk returns a tuple with the AccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatus

`func (o *User) SetAccountStatus(v AccountStatus)`

SetAccountStatus sets AccountStatus field to given value.

### HasAccountStatus

`func (o *User) HasAccountStatus() bool`

HasAccountStatus returns a boolean if a field has been set.

### GetAccountId

`func (o *User) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *User) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *User) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *User) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAccountType

`func (o *User) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *User) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *User) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *User) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetPublicName

`func (o *User) GetPublicName() string`

GetPublicName returns the PublicName field if non-nil, zero value otherwise.

### GetPublicNameOk

`func (o *User) GetPublicNameOk() (*string, bool)`

GetPublicNameOk returns a tuple with the PublicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicName

`func (o *User) SetPublicName(v string)`

SetPublicName sets PublicName field to given value.

### HasPublicName

`func (o *User) HasPublicName() bool`

HasPublicName returns a boolean if a field has been set.

### GetProfilePicture

`func (o *User) GetProfilePicture() Icon`

GetProfilePicture returns the ProfilePicture field if non-nil, zero value otherwise.

### GetProfilePictureOk

`func (o *User) GetProfilePictureOk() (*Icon, bool)`

GetProfilePictureOk returns a tuple with the ProfilePicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicture

`func (o *User) SetProfilePicture(v Icon)`

SetProfilePicture sets ProfilePicture field to given value.

### HasProfilePicture

`func (o *User) HasProfilePicture() bool`

HasProfilePicture returns a boolean if a field has been set.

### SetProfilePictureNil

`func (o *User) SetProfilePictureNil(b bool)`

 SetProfilePictureNil sets the value for ProfilePicture to be an explicit nil

### UnsetProfilePicture
`func (o *User) UnsetProfilePicture()`

UnsetProfilePicture ensures that no value is present for ProfilePicture, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



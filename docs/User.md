# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Email** | **string** | The new users email address. | 
**Blocked** | Pointer to **bool** | Boolean to indicate if the user is blocked. | [optional] 
**BlockedCode** | Pointer to [**NullableUserBlockedCodeProperty**](UserBlockedCodeProperty.md) |  | [optional] 
**Role** | Pointer to [**NullableUserRoleProperty**](UserRoleProperty.md) |  | [optional] 

## Methods

### NewUser

`func NewUser(email string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *User) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *User) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *User) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *User) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *User) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *User) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

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


### GetBlocked

`func (o *User) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *User) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *User) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *User) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetBlockedCode

`func (o *User) GetBlockedCode() UserBlockedCodeProperty`

GetBlockedCode returns the BlockedCode field if non-nil, zero value otherwise.

### GetBlockedCodeOk

`func (o *User) GetBlockedCodeOk() (*UserBlockedCodeProperty, bool)`

GetBlockedCodeOk returns a tuple with the BlockedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedCode

`func (o *User) SetBlockedCode(v UserBlockedCodeProperty)`

SetBlockedCode sets BlockedCode field to given value.

### HasBlockedCode

`func (o *User) HasBlockedCode() bool`

HasBlockedCode returns a boolean if a field has been set.

### SetBlockedCodeNil

`func (o *User) SetBlockedCodeNil(b bool)`

 SetBlockedCodeNil sets the value for BlockedCode to be an explicit nil

### UnsetBlockedCode
`func (o *User) UnsetBlockedCode()`

UnsetBlockedCode ensures that no value is present for BlockedCode, not even an explicit nil
### GetRole

`func (o *User) GetRole() UserRoleProperty`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *User) GetRoleOk() (*UserRoleProperty, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *User) SetRole(v UserRoleProperty)`

SetRole sets Role field to given value.

### HasRole

`func (o *User) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *User) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *User) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



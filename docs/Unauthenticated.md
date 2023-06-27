# Unauthenticated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Exception** | Pointer to **string** |  | [optional] 

## Methods

### NewUnauthenticated

`func NewUnauthenticated() *Unauthenticated`

NewUnauthenticated instantiates a new Unauthenticated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnauthenticatedWithDefaults

`func NewUnauthenticatedWithDefaults() *Unauthenticated`

NewUnauthenticatedWithDefaults instantiates a new Unauthenticated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *Unauthenticated) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Unauthenticated) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Unauthenticated) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Unauthenticated) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetException

`func (o *Unauthenticated) GetException() string`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *Unauthenticated) GetExceptionOk() (*string, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *Unauthenticated) SetException(v string)`

SetException sets Exception field to given value.

### HasException

`func (o *Unauthenticated) HasException() bool`

HasException returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



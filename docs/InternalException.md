# InternalException

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Exception** | Pointer to **string** |  | [optional] 

## Methods

### NewInternalException

`func NewInternalException() *InternalException`

NewInternalException instantiates a new InternalException object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalExceptionWithDefaults

`func NewInternalExceptionWithDefaults() *InternalException`

NewInternalExceptionWithDefaults instantiates a new InternalException object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *InternalException) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InternalException) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InternalException) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InternalException) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetException

`func (o *InternalException) GetException() string`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *InternalException) GetExceptionOk() (*string, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *InternalException) SetException(v string)`

SetException sets Exception field to given value.

### HasException

`func (o *InternalException) HasException() bool`

HasException returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



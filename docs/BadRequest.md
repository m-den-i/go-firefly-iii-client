# BadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Exception** | Pointer to **string** |  | [optional] 

## Methods

### NewBadRequest

`func NewBadRequest() *BadRequest`

NewBadRequest instantiates a new BadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadRequestWithDefaults

`func NewBadRequestWithDefaults() *BadRequest`

NewBadRequestWithDefaults instantiates a new BadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *BadRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BadRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BadRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BadRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetException

`func (o *BadRequest) GetException() string`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *BadRequest) GetExceptionOk() (*string, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *BadRequest) SetException(v string)`

SetException sets Exception field to given value.

### HasException

`func (o *BadRequest) HasException() bool`

HasException returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



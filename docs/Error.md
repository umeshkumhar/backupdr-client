# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrCode** | **int32** | Error code | 
**ErrMessage** | Pointer to **string** | Optional error message | [optional] 

## Methods

### NewError

`func NewError(errCode int32, ) *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrCode

`func (o *Error) GetErrCode() int32`

GetErrCode returns the ErrCode field if non-nil, zero value otherwise.

### GetErrCodeOk

`func (o *Error) GetErrCodeOk() (*int32, bool)`

GetErrCodeOk returns a tuple with the ErrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrCode

`func (o *Error) SetErrCode(v int32)`

SetErrCode sets ErrCode field to given value.


### GetErrMessage

`func (o *Error) GetErrMessage() string`

GetErrMessage returns the ErrMessage field if non-nil, zero value otherwise.

### GetErrMessageOk

`func (o *Error) GetErrMessageOk() (*string, bool)`

GetErrMessageOk returns a tuple with the ErrMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrMessage

`func (o *Error) SetErrMessage(v string)`

SetErrMessage sets ErrMessage field to given value.

### HasErrMessage

`func (o *Error) HasErrMessage() bool`

HasErrMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



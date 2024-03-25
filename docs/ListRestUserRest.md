# ListRestUserRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]UserRest**](UserRest.md) |  | [optional] 

## Methods

### NewListRestUserRest

`func NewListRestUserRest() *ListRestUserRest`

NewListRestUserRest instantiates a new ListRestUserRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRestUserRestWithDefaults

`func NewListRestUserRestWithDefaults() *ListRestUserRest`

NewListRestUserRestWithDefaults instantiates a new ListRestUserRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ListRestUserRest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListRestUserRest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListRestUserRest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListRestUserRest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetItems

`func (o *ListRestUserRest) GetItems() []UserRest`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListRestUserRest) GetItemsOk() (*[]UserRest, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListRestUserRest) SetItems(v []UserRest)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListRestUserRest) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



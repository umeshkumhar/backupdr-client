# ListDiskPoolRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]DiskPoolRest**](DiskPoolRest.md) |  | [optional] 

## Methods

### NewListDiskPoolRest

`func NewListDiskPoolRest() *ListDiskPoolRest`

NewListDiskPoolRest instantiates a new ListDiskPoolRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDiskPoolRestWithDefaults

`func NewListDiskPoolRestWithDefaults() *ListDiskPoolRest`

NewListDiskPoolRestWithDefaults instantiates a new ListDiskPoolRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ListDiskPoolRest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListDiskPoolRest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListDiskPoolRest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListDiskPoolRest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetItems

`func (o *ListDiskPoolRest) GetItems() []DiskPoolRest`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListDiskPoolRest) GetItemsOk() (*[]DiskPoolRest, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListDiskPoolRest) SetItems(v []DiskPoolRest)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListDiskPoolRest) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



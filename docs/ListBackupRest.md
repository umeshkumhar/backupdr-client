# ListBackupRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]BackupRest**](BackupRest.md) |  | [optional] 

## Methods

### NewListBackupRest

`func NewListBackupRest() *ListBackupRest`

NewListBackupRest instantiates a new ListBackupRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBackupRestWithDefaults

`func NewListBackupRestWithDefaults() *ListBackupRest`

NewListBackupRestWithDefaults instantiates a new ListBackupRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ListBackupRest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListBackupRest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListBackupRest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListBackupRest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetItems

`func (o *ListBackupRest) GetItems() []BackupRest`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListBackupRest) GetItemsOk() (*[]BackupRest, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListBackupRest) SetItems(v []BackupRest)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListBackupRest) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



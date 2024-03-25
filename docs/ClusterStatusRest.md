# ClusterStatusRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessible** | Pointer to **bool** |  | [optional] 
**Vdiskcount** | Pointer to **int32** |  | [optional] 
**Vdisklimit** | Pointer to **int32** |  | [optional] 
**Vdiskcolor** | Pointer to **string** |  | [optional] 
**Localsnapshotcolor** | Pointer to **string** |  | [optional] 
**Localdedupcolor** | Pointer to **string** |  | [optional] 
**Remoteprotectioncolor** | Pointer to **string** |  | [optional] 
**Copydata** | Pointer to **int64** |  | [optional] 

## Methods

### NewClusterStatusRest

`func NewClusterStatusRest() *ClusterStatusRest`

NewClusterStatusRest instantiates a new ClusterStatusRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStatusRestWithDefaults

`func NewClusterStatusRestWithDefaults() *ClusterStatusRest`

NewClusterStatusRestWithDefaults instantiates a new ClusterStatusRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessible

`func (o *ClusterStatusRest) GetAccessible() bool`

GetAccessible returns the Accessible field if non-nil, zero value otherwise.

### GetAccessibleOk

`func (o *ClusterStatusRest) GetAccessibleOk() (*bool, bool)`

GetAccessibleOk returns a tuple with the Accessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessible

`func (o *ClusterStatusRest) SetAccessible(v bool)`

SetAccessible sets Accessible field to given value.

### HasAccessible

`func (o *ClusterStatusRest) HasAccessible() bool`

HasAccessible returns a boolean if a field has been set.

### GetVdiskcount

`func (o *ClusterStatusRest) GetVdiskcount() int32`

GetVdiskcount returns the Vdiskcount field if non-nil, zero value otherwise.

### GetVdiskcountOk

`func (o *ClusterStatusRest) GetVdiskcountOk() (*int32, bool)`

GetVdiskcountOk returns a tuple with the Vdiskcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiskcount

`func (o *ClusterStatusRest) SetVdiskcount(v int32)`

SetVdiskcount sets Vdiskcount field to given value.

### HasVdiskcount

`func (o *ClusterStatusRest) HasVdiskcount() bool`

HasVdiskcount returns a boolean if a field has been set.

### GetVdisklimit

`func (o *ClusterStatusRest) GetVdisklimit() int32`

GetVdisklimit returns the Vdisklimit field if non-nil, zero value otherwise.

### GetVdisklimitOk

`func (o *ClusterStatusRest) GetVdisklimitOk() (*int32, bool)`

GetVdisklimitOk returns a tuple with the Vdisklimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdisklimit

`func (o *ClusterStatusRest) SetVdisklimit(v int32)`

SetVdisklimit sets Vdisklimit field to given value.

### HasVdisklimit

`func (o *ClusterStatusRest) HasVdisklimit() bool`

HasVdisklimit returns a boolean if a field has been set.

### GetVdiskcolor

`func (o *ClusterStatusRest) GetVdiskcolor() string`

GetVdiskcolor returns the Vdiskcolor field if non-nil, zero value otherwise.

### GetVdiskcolorOk

`func (o *ClusterStatusRest) GetVdiskcolorOk() (*string, bool)`

GetVdiskcolorOk returns a tuple with the Vdiskcolor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiskcolor

`func (o *ClusterStatusRest) SetVdiskcolor(v string)`

SetVdiskcolor sets Vdiskcolor field to given value.

### HasVdiskcolor

`func (o *ClusterStatusRest) HasVdiskcolor() bool`

HasVdiskcolor returns a boolean if a field has been set.

### GetLocalsnapshotcolor

`func (o *ClusterStatusRest) GetLocalsnapshotcolor() string`

GetLocalsnapshotcolor returns the Localsnapshotcolor field if non-nil, zero value otherwise.

### GetLocalsnapshotcolorOk

`func (o *ClusterStatusRest) GetLocalsnapshotcolorOk() (*string, bool)`

GetLocalsnapshotcolorOk returns a tuple with the Localsnapshotcolor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalsnapshotcolor

`func (o *ClusterStatusRest) SetLocalsnapshotcolor(v string)`

SetLocalsnapshotcolor sets Localsnapshotcolor field to given value.

### HasLocalsnapshotcolor

`func (o *ClusterStatusRest) HasLocalsnapshotcolor() bool`

HasLocalsnapshotcolor returns a boolean if a field has been set.

### GetLocaldedupcolor

`func (o *ClusterStatusRest) GetLocaldedupcolor() string`

GetLocaldedupcolor returns the Localdedupcolor field if non-nil, zero value otherwise.

### GetLocaldedupcolorOk

`func (o *ClusterStatusRest) GetLocaldedupcolorOk() (*string, bool)`

GetLocaldedupcolorOk returns a tuple with the Localdedupcolor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocaldedupcolor

`func (o *ClusterStatusRest) SetLocaldedupcolor(v string)`

SetLocaldedupcolor sets Localdedupcolor field to given value.

### HasLocaldedupcolor

`func (o *ClusterStatusRest) HasLocaldedupcolor() bool`

HasLocaldedupcolor returns a boolean if a field has been set.

### GetRemoteprotectioncolor

`func (o *ClusterStatusRest) GetRemoteprotectioncolor() string`

GetRemoteprotectioncolor returns the Remoteprotectioncolor field if non-nil, zero value otherwise.

### GetRemoteprotectioncolorOk

`func (o *ClusterStatusRest) GetRemoteprotectioncolorOk() (*string, bool)`

GetRemoteprotectioncolorOk returns a tuple with the Remoteprotectioncolor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteprotectioncolor

`func (o *ClusterStatusRest) SetRemoteprotectioncolor(v string)`

SetRemoteprotectioncolor sets Remoteprotectioncolor field to given value.

### HasRemoteprotectioncolor

`func (o *ClusterStatusRest) HasRemoteprotectioncolor() bool`

HasRemoteprotectioncolor returns a boolean if a field has been set.

### GetCopydata

`func (o *ClusterStatusRest) GetCopydata() int64`

GetCopydata returns the Copydata field if non-nil, zero value otherwise.

### GetCopydataOk

`func (o *ClusterStatusRest) GetCopydataOk() (*int64, bool)`

GetCopydataOk returns a tuple with the Copydata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopydata

`func (o *ClusterStatusRest) SetCopydata(v int64)`

SetCopydata sets Copydata field to given value.

### HasCopydata

`func (o *ClusterStatusRest) HasCopydata() bool`

HasCopydata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



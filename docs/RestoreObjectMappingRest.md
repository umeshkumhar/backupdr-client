# RestoreObjectMappingRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vmdkprovisionformat** | Pointer to **string** |  | [optional] 
**Diskpool** | Pointer to [**DiskPoolRest**](DiskPoolRest.md) |  | [optional] 
**Appliancemountpoint** | Pointer to **string** |  | [optional] 
**Datastore** | Pointer to **string** |  | [optional] 
**Restoreobject** | Pointer to **string** |  | [optional] 
**Mountdrive** | Pointer to **string** |  | [optional] 
**Mountpoint** | Pointer to **string** |  | [optional] 
**Targetvg** | Pointer to **string** |  | [optional] 

## Methods

### NewRestoreObjectMappingRest

`func NewRestoreObjectMappingRest() *RestoreObjectMappingRest`

NewRestoreObjectMappingRest instantiates a new RestoreObjectMappingRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreObjectMappingRestWithDefaults

`func NewRestoreObjectMappingRestWithDefaults() *RestoreObjectMappingRest`

NewRestoreObjectMappingRestWithDefaults instantiates a new RestoreObjectMappingRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmdkprovisionformat

`func (o *RestoreObjectMappingRest) GetVmdkprovisionformat() string`

GetVmdkprovisionformat returns the Vmdkprovisionformat field if non-nil, zero value otherwise.

### GetVmdkprovisionformatOk

`func (o *RestoreObjectMappingRest) GetVmdkprovisionformatOk() (*string, bool)`

GetVmdkprovisionformatOk returns a tuple with the Vmdkprovisionformat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmdkprovisionformat

`func (o *RestoreObjectMappingRest) SetVmdkprovisionformat(v string)`

SetVmdkprovisionformat sets Vmdkprovisionformat field to given value.

### HasVmdkprovisionformat

`func (o *RestoreObjectMappingRest) HasVmdkprovisionformat() bool`

HasVmdkprovisionformat returns a boolean if a field has been set.

### GetDiskpool

`func (o *RestoreObjectMappingRest) GetDiskpool() DiskPoolRest`

GetDiskpool returns the Diskpool field if non-nil, zero value otherwise.

### GetDiskpoolOk

`func (o *RestoreObjectMappingRest) GetDiskpoolOk() (*DiskPoolRest, bool)`

GetDiskpoolOk returns a tuple with the Diskpool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskpool

`func (o *RestoreObjectMappingRest) SetDiskpool(v DiskPoolRest)`

SetDiskpool sets Diskpool field to given value.

### HasDiskpool

`func (o *RestoreObjectMappingRest) HasDiskpool() bool`

HasDiskpool returns a boolean if a field has been set.

### GetAppliancemountpoint

`func (o *RestoreObjectMappingRest) GetAppliancemountpoint() string`

GetAppliancemountpoint returns the Appliancemountpoint field if non-nil, zero value otherwise.

### GetAppliancemountpointOk

`func (o *RestoreObjectMappingRest) GetAppliancemountpointOk() (*string, bool)`

GetAppliancemountpointOk returns a tuple with the Appliancemountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliancemountpoint

`func (o *RestoreObjectMappingRest) SetAppliancemountpoint(v string)`

SetAppliancemountpoint sets Appliancemountpoint field to given value.

### HasAppliancemountpoint

`func (o *RestoreObjectMappingRest) HasAppliancemountpoint() bool`

HasAppliancemountpoint returns a boolean if a field has been set.

### GetDatastore

`func (o *RestoreObjectMappingRest) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *RestoreObjectMappingRest) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *RestoreObjectMappingRest) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *RestoreObjectMappingRest) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetRestoreobject

`func (o *RestoreObjectMappingRest) GetRestoreobject() string`

GetRestoreobject returns the Restoreobject field if non-nil, zero value otherwise.

### GetRestoreobjectOk

`func (o *RestoreObjectMappingRest) GetRestoreobjectOk() (*string, bool)`

GetRestoreobjectOk returns a tuple with the Restoreobject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreobject

`func (o *RestoreObjectMappingRest) SetRestoreobject(v string)`

SetRestoreobject sets Restoreobject field to given value.

### HasRestoreobject

`func (o *RestoreObjectMappingRest) HasRestoreobject() bool`

HasRestoreobject returns a boolean if a field has been set.

### GetMountdrive

`func (o *RestoreObjectMappingRest) GetMountdrive() string`

GetMountdrive returns the Mountdrive field if non-nil, zero value otherwise.

### GetMountdriveOk

`func (o *RestoreObjectMappingRest) GetMountdriveOk() (*string, bool)`

GetMountdriveOk returns a tuple with the Mountdrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountdrive

`func (o *RestoreObjectMappingRest) SetMountdrive(v string)`

SetMountdrive sets Mountdrive field to given value.

### HasMountdrive

`func (o *RestoreObjectMappingRest) HasMountdrive() bool`

HasMountdrive returns a boolean if a field has been set.

### GetMountpoint

`func (o *RestoreObjectMappingRest) GetMountpoint() string`

GetMountpoint returns the Mountpoint field if non-nil, zero value otherwise.

### GetMountpointOk

`func (o *RestoreObjectMappingRest) GetMountpointOk() (*string, bool)`

GetMountpointOk returns a tuple with the Mountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountpoint

`func (o *RestoreObjectMappingRest) SetMountpoint(v string)`

SetMountpoint sets Mountpoint field to given value.

### HasMountpoint

`func (o *RestoreObjectMappingRest) HasMountpoint() bool`

HasMountpoint returns a boolean if a field has been set.

### GetTargetvg

`func (o *RestoreObjectMappingRest) GetTargetvg() string`

GetTargetvg returns the Targetvg field if non-nil, zero value otherwise.

### GetTargetvgOk

`func (o *RestoreObjectMappingRest) GetTargetvgOk() (*string, bool)`

GetTargetvgOk returns a tuple with the Targetvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetvg

`func (o *RestoreObjectMappingRest) SetTargetvg(v string)`

SetTargetvg sets Targetvg field to given value.

### HasTargetvg

`func (o *RestoreObjectMappingRest) HasTargetvg() bool`

HasTargetvg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



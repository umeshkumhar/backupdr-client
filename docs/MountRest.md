# MountRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**Host** | Pointer to [**HostRest**](HostRest.md) |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** | Mutually exclusive with host, when mounting as a brand new host(VM) | [optional] 
**Container** | Pointer to **bool** |  | [optional] 
**Diskpool** | Pointer to [**DiskPoolRest**](DiskPoolRest.md) |  | [optional] 
**Cloudvmoptions** | Pointer to [**CloudVmMountRest**](CloudVmMountRest.md) |  | [optional] 
**Rdmmode** | Pointer to **string** |  | [optional] 
**Rehydrationmode** | Pointer to **string** |  | [optional] 
**Appaware** | Pointer to **bool** |  | [optional] 
**Systemstateoptions** | Pointer to [**[]SystemStateOptionRest**](SystemStateOptionRest.md) |  | [optional] 
**Migratevm** | Pointer to **bool** |  | [optional] 
**Allowedips** | Pointer to **[]string** |  | [optional] 
**Datastore** | Pointer to **string** | Mutually exclusive with diskpool | [optional] 
**Recoverytime** | Pointer to **int64** | Can be used when log backup is available. Microseconds since Epoch in UTC | [optional] 
**Poweronvm** | Pointer to **bool** |  | [optional] 
**Scripts** | Pointer to [**[]ScriptRest**](ScriptRest.md) |  | [optional] 
**Restoreobjectmappings** | Pointer to [**[]RestoreObjectMappingRest**](RestoreObjectMappingRest.md) | A customized mapping from the volumes to their mount points | [optional] 
**Restoreoptions** | Pointer to [**[]RestoreOptionRest**](RestoreOptionRest.md) |  | [optional] 
**Provisioningoptions** | Pointer to [**[]ProvisioningOptionRest**](ProvisioningOptionRest.md) | List of provisioning options when appaware is true | [optional] 
**Selectedobjects** | Pointer to [**[]SelectedObjectRest**](SelectedObjectRest.md) | A subset of restorable objects to be mounted | [optional] 
**Nfsoptions** | Pointer to [**NfsOptionsRest**](NfsOptionsRest.md) |  | [optional] 
**Instantmount** | Pointer to **bool** |  | [optional] 
**Mgmtserver** | Pointer to [**HostRest**](HostRest.md) |  | [optional] 
**Hypervisor** | Pointer to [**HostRest**](HostRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewMountRest

`func NewMountRest() *MountRest`

NewMountRest instantiates a new MountRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMountRestWithDefaults

`func NewMountRestWithDefaults() *MountRest`

NewMountRestWithDefaults instantiates a new MountRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *MountRest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MountRest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MountRest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MountRest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetHost

`func (o *MountRest) GetHost() HostRest`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MountRest) GetHostOk() (*HostRest, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MountRest) SetHost(v HostRest)`

SetHost sets Host field to given value.

### HasHost

`func (o *MountRest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetLabel

`func (o *MountRest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MountRest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MountRest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *MountRest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetHostname

`func (o *MountRest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *MountRest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *MountRest) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *MountRest) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetContainer

`func (o *MountRest) GetContainer() bool`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *MountRest) GetContainerOk() (*bool, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *MountRest) SetContainer(v bool)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *MountRest) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetDiskpool

`func (o *MountRest) GetDiskpool() DiskPoolRest`

GetDiskpool returns the Diskpool field if non-nil, zero value otherwise.

### GetDiskpoolOk

`func (o *MountRest) GetDiskpoolOk() (*DiskPoolRest, bool)`

GetDiskpoolOk returns a tuple with the Diskpool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskpool

`func (o *MountRest) SetDiskpool(v DiskPoolRest)`

SetDiskpool sets Diskpool field to given value.

### HasDiskpool

`func (o *MountRest) HasDiskpool() bool`

HasDiskpool returns a boolean if a field has been set.

### GetCloudvmoptions

`func (o *MountRest) GetCloudvmoptions() CloudVmMountRest`

GetCloudvmoptions returns the Cloudvmoptions field if non-nil, zero value otherwise.

### GetCloudvmoptionsOk

`func (o *MountRest) GetCloudvmoptionsOk() (*CloudVmMountRest, bool)`

GetCloudvmoptionsOk returns a tuple with the Cloudvmoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudvmoptions

`func (o *MountRest) SetCloudvmoptions(v CloudVmMountRest)`

SetCloudvmoptions sets Cloudvmoptions field to given value.

### HasCloudvmoptions

`func (o *MountRest) HasCloudvmoptions() bool`

HasCloudvmoptions returns a boolean if a field has been set.

### GetRdmmode

`func (o *MountRest) GetRdmmode() string`

GetRdmmode returns the Rdmmode field if non-nil, zero value otherwise.

### GetRdmmodeOk

`func (o *MountRest) GetRdmmodeOk() (*string, bool)`

GetRdmmodeOk returns a tuple with the Rdmmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdmmode

`func (o *MountRest) SetRdmmode(v string)`

SetRdmmode sets Rdmmode field to given value.

### HasRdmmode

`func (o *MountRest) HasRdmmode() bool`

HasRdmmode returns a boolean if a field has been set.

### GetRehydrationmode

`func (o *MountRest) GetRehydrationmode() string`

GetRehydrationmode returns the Rehydrationmode field if non-nil, zero value otherwise.

### GetRehydrationmodeOk

`func (o *MountRest) GetRehydrationmodeOk() (*string, bool)`

GetRehydrationmodeOk returns a tuple with the Rehydrationmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRehydrationmode

`func (o *MountRest) SetRehydrationmode(v string)`

SetRehydrationmode sets Rehydrationmode field to given value.

### HasRehydrationmode

`func (o *MountRest) HasRehydrationmode() bool`

HasRehydrationmode returns a boolean if a field has been set.

### GetAppaware

`func (o *MountRest) GetAppaware() bool`

GetAppaware returns the Appaware field if non-nil, zero value otherwise.

### GetAppawareOk

`func (o *MountRest) GetAppawareOk() (*bool, bool)`

GetAppawareOk returns a tuple with the Appaware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppaware

`func (o *MountRest) SetAppaware(v bool)`

SetAppaware sets Appaware field to given value.

### HasAppaware

`func (o *MountRest) HasAppaware() bool`

HasAppaware returns a boolean if a field has been set.

### GetSystemstateoptions

`func (o *MountRest) GetSystemstateoptions() []SystemStateOptionRest`

GetSystemstateoptions returns the Systemstateoptions field if non-nil, zero value otherwise.

### GetSystemstateoptionsOk

`func (o *MountRest) GetSystemstateoptionsOk() (*[]SystemStateOptionRest, bool)`

GetSystemstateoptionsOk returns a tuple with the Systemstateoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemstateoptions

`func (o *MountRest) SetSystemstateoptions(v []SystemStateOptionRest)`

SetSystemstateoptions sets Systemstateoptions field to given value.

### HasSystemstateoptions

`func (o *MountRest) HasSystemstateoptions() bool`

HasSystemstateoptions returns a boolean if a field has been set.

### GetMigratevm

`func (o *MountRest) GetMigratevm() bool`

GetMigratevm returns the Migratevm field if non-nil, zero value otherwise.

### GetMigratevmOk

`func (o *MountRest) GetMigratevmOk() (*bool, bool)`

GetMigratevmOk returns a tuple with the Migratevm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratevm

`func (o *MountRest) SetMigratevm(v bool)`

SetMigratevm sets Migratevm field to given value.

### HasMigratevm

`func (o *MountRest) HasMigratevm() bool`

HasMigratevm returns a boolean if a field has been set.

### GetAllowedips

`func (o *MountRest) GetAllowedips() []string`

GetAllowedips returns the Allowedips field if non-nil, zero value otherwise.

### GetAllowedipsOk

`func (o *MountRest) GetAllowedipsOk() (*[]string, bool)`

GetAllowedipsOk returns a tuple with the Allowedips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedips

`func (o *MountRest) SetAllowedips(v []string)`

SetAllowedips sets Allowedips field to given value.

### HasAllowedips

`func (o *MountRest) HasAllowedips() bool`

HasAllowedips returns a boolean if a field has been set.

### GetDatastore

`func (o *MountRest) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *MountRest) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *MountRest) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *MountRest) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetRecoverytime

`func (o *MountRest) GetRecoverytime() int64`

GetRecoverytime returns the Recoverytime field if non-nil, zero value otherwise.

### GetRecoverytimeOk

`func (o *MountRest) GetRecoverytimeOk() (*int64, bool)`

GetRecoverytimeOk returns a tuple with the Recoverytime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverytime

`func (o *MountRest) SetRecoverytime(v int64)`

SetRecoverytime sets Recoverytime field to given value.

### HasRecoverytime

`func (o *MountRest) HasRecoverytime() bool`

HasRecoverytime returns a boolean if a field has been set.

### GetPoweronvm

`func (o *MountRest) GetPoweronvm() bool`

GetPoweronvm returns the Poweronvm field if non-nil, zero value otherwise.

### GetPoweronvmOk

`func (o *MountRest) GetPoweronvmOk() (*bool, bool)`

GetPoweronvmOk returns a tuple with the Poweronvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoweronvm

`func (o *MountRest) SetPoweronvm(v bool)`

SetPoweronvm sets Poweronvm field to given value.

### HasPoweronvm

`func (o *MountRest) HasPoweronvm() bool`

HasPoweronvm returns a boolean if a field has been set.

### GetScripts

`func (o *MountRest) GetScripts() []ScriptRest`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *MountRest) GetScriptsOk() (*[]ScriptRest, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *MountRest) SetScripts(v []ScriptRest)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *MountRest) HasScripts() bool`

HasScripts returns a boolean if a field has been set.

### GetRestoreobjectmappings

`func (o *MountRest) GetRestoreobjectmappings() []RestoreObjectMappingRest`

GetRestoreobjectmappings returns the Restoreobjectmappings field if non-nil, zero value otherwise.

### GetRestoreobjectmappingsOk

`func (o *MountRest) GetRestoreobjectmappingsOk() (*[]RestoreObjectMappingRest, bool)`

GetRestoreobjectmappingsOk returns a tuple with the Restoreobjectmappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreobjectmappings

`func (o *MountRest) SetRestoreobjectmappings(v []RestoreObjectMappingRest)`

SetRestoreobjectmappings sets Restoreobjectmappings field to given value.

### HasRestoreobjectmappings

`func (o *MountRest) HasRestoreobjectmappings() bool`

HasRestoreobjectmappings returns a boolean if a field has been set.

### GetRestoreoptions

`func (o *MountRest) GetRestoreoptions() []RestoreOptionRest`

GetRestoreoptions returns the Restoreoptions field if non-nil, zero value otherwise.

### GetRestoreoptionsOk

`func (o *MountRest) GetRestoreoptionsOk() (*[]RestoreOptionRest, bool)`

GetRestoreoptionsOk returns a tuple with the Restoreoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreoptions

`func (o *MountRest) SetRestoreoptions(v []RestoreOptionRest)`

SetRestoreoptions sets Restoreoptions field to given value.

### HasRestoreoptions

`func (o *MountRest) HasRestoreoptions() bool`

HasRestoreoptions returns a boolean if a field has been set.

### GetProvisioningoptions

`func (o *MountRest) GetProvisioningoptions() []ProvisioningOptionRest`

GetProvisioningoptions returns the Provisioningoptions field if non-nil, zero value otherwise.

### GetProvisioningoptionsOk

`func (o *MountRest) GetProvisioningoptionsOk() (*[]ProvisioningOptionRest, bool)`

GetProvisioningoptionsOk returns a tuple with the Provisioningoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningoptions

`func (o *MountRest) SetProvisioningoptions(v []ProvisioningOptionRest)`

SetProvisioningoptions sets Provisioningoptions field to given value.

### HasProvisioningoptions

`func (o *MountRest) HasProvisioningoptions() bool`

HasProvisioningoptions returns a boolean if a field has been set.

### GetSelectedobjects

`func (o *MountRest) GetSelectedobjects() []SelectedObjectRest`

GetSelectedobjects returns the Selectedobjects field if non-nil, zero value otherwise.

### GetSelectedobjectsOk

`func (o *MountRest) GetSelectedobjectsOk() (*[]SelectedObjectRest, bool)`

GetSelectedobjectsOk returns a tuple with the Selectedobjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedobjects

`func (o *MountRest) SetSelectedobjects(v []SelectedObjectRest)`

SetSelectedobjects sets Selectedobjects field to given value.

### HasSelectedobjects

`func (o *MountRest) HasSelectedobjects() bool`

HasSelectedobjects returns a boolean if a field has been set.

### GetNfsoptions

`func (o *MountRest) GetNfsoptions() NfsOptionsRest`

GetNfsoptions returns the Nfsoptions field if non-nil, zero value otherwise.

### GetNfsoptionsOk

`func (o *MountRest) GetNfsoptionsOk() (*NfsOptionsRest, bool)`

GetNfsoptionsOk returns a tuple with the Nfsoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsoptions

`func (o *MountRest) SetNfsoptions(v NfsOptionsRest)`

SetNfsoptions sets Nfsoptions field to given value.

### HasNfsoptions

`func (o *MountRest) HasNfsoptions() bool`

HasNfsoptions returns a boolean if a field has been set.

### GetInstantmount

`func (o *MountRest) GetInstantmount() bool`

GetInstantmount returns the Instantmount field if non-nil, zero value otherwise.

### GetInstantmountOk

`func (o *MountRest) GetInstantmountOk() (*bool, bool)`

GetInstantmountOk returns a tuple with the Instantmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantmount

`func (o *MountRest) SetInstantmount(v bool)`

SetInstantmount sets Instantmount field to given value.

### HasInstantmount

`func (o *MountRest) HasInstantmount() bool`

HasInstantmount returns a boolean if a field has been set.

### GetMgmtserver

`func (o *MountRest) GetMgmtserver() HostRest`

GetMgmtserver returns the Mgmtserver field if non-nil, zero value otherwise.

### GetMgmtserverOk

`func (o *MountRest) GetMgmtserverOk() (*HostRest, bool)`

GetMgmtserverOk returns a tuple with the Mgmtserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtserver

`func (o *MountRest) SetMgmtserver(v HostRest)`

SetMgmtserver sets Mgmtserver field to given value.

### HasMgmtserver

`func (o *MountRest) HasMgmtserver() bool`

HasMgmtserver returns a boolean if a field has been set.

### GetHypervisor

`func (o *MountRest) GetHypervisor() HostRest`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *MountRest) GetHypervisorOk() (*HostRest, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisor

`func (o *MountRest) SetHypervisor(v HostRest)`

SetHypervisor sets Hypervisor field to given value.

### HasHypervisor

`func (o *MountRest) HasHypervisor() bool`

HasHypervisor returns a boolean if a field has been set.

### GetId

`func (o *MountRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MountRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MountRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MountRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *MountRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *MountRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *MountRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *MountRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *MountRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *MountRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *MountRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *MountRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *MountRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *MountRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *MountRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *MountRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



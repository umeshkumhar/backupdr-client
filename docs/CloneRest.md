# CloneRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**Host** | Pointer to [**HostRest**](HostRest.md) |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Diskpool** | Pointer to [**DiskPoolRest**](DiskPoolRest.md) |  | [optional] 
**Vcenter** | Pointer to [**HostRest**](HostRest.md) |  | [optional] 
**Datastore** | Pointer to **string** |  | [optional] 
**Recoverytime** | Pointer to **int64** |  | [optional] 
**Poweronvm** | Pointer to **bool** |  | [optional] 
**Scripts** | Pointer to [**[]ScriptRest**](ScriptRest.md) |  | [optional] 
**Restoreobjectmappings** | Pointer to [**[]RestoreObjectMappingRest**](RestoreObjectMappingRest.md) |  | [optional] 
**Restoreoptions** | Pointer to [**[]RestoreOptionRest**](RestoreOptionRest.md) |  | [optional] 
**Provisioningoptions** | Pointer to [**[]ProvisioningOptionRest**](ProvisioningOptionRest.md) |  | [optional] 
**Selectedobjects** | Pointer to [**[]SelectedObjectRest**](SelectedObjectRest.md) |  | [optional] 
**Nfsoptions** | Pointer to [**NfsOptionsRest**](NfsOptionsRest.md) |  | [optional] 
**Instantmount** | Pointer to **bool** |  | [optional] 
**Restorelocation** | Pointer to [**RestoreLocationRest**](RestoreLocationRest.md) |  | [optional] 
**Mgmtserver** | Pointer to [**HostRest**](HostRest.md) |  | [optional] 
**Hypervisor** | Pointer to [**HostRest**](HostRest.md) |  | [optional] 
**Vmname** | Pointer to **string** |  | [optional] 
**Esxhost** | Pointer to [**HostRest**](HostRest.md) |  | [optional] 
**Container** | Pointer to **bool** |  | [optional] 
**Allowedips** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewCloneRest

`func NewCloneRest() *CloneRest`

NewCloneRest instantiates a new CloneRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneRestWithDefaults

`func NewCloneRestWithDefaults() *CloneRest`

NewCloneRestWithDefaults instantiates a new CloneRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *CloneRest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CloneRest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CloneRest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CloneRest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetHost

`func (o *CloneRest) GetHost() HostRest`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CloneRest) GetHostOk() (*HostRest, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CloneRest) SetHost(v HostRest)`

SetHost sets Host field to given value.

### HasHost

`func (o *CloneRest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetLabel

`func (o *CloneRest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CloneRest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CloneRest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CloneRest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetHostname

`func (o *CloneRest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CloneRest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CloneRest) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *CloneRest) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetDiskpool

`func (o *CloneRest) GetDiskpool() DiskPoolRest`

GetDiskpool returns the Diskpool field if non-nil, zero value otherwise.

### GetDiskpoolOk

`func (o *CloneRest) GetDiskpoolOk() (*DiskPoolRest, bool)`

GetDiskpoolOk returns a tuple with the Diskpool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskpool

`func (o *CloneRest) SetDiskpool(v DiskPoolRest)`

SetDiskpool sets Diskpool field to given value.

### HasDiskpool

`func (o *CloneRest) HasDiskpool() bool`

HasDiskpool returns a boolean if a field has been set.

### GetVcenter

`func (o *CloneRest) GetVcenter() HostRest`

GetVcenter returns the Vcenter field if non-nil, zero value otherwise.

### GetVcenterOk

`func (o *CloneRest) GetVcenterOk() (*HostRest, bool)`

GetVcenterOk returns a tuple with the Vcenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenter

`func (o *CloneRest) SetVcenter(v HostRest)`

SetVcenter sets Vcenter field to given value.

### HasVcenter

`func (o *CloneRest) HasVcenter() bool`

HasVcenter returns a boolean if a field has been set.

### GetDatastore

`func (o *CloneRest) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *CloneRest) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *CloneRest) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *CloneRest) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetRecoverytime

`func (o *CloneRest) GetRecoverytime() int64`

GetRecoverytime returns the Recoverytime field if non-nil, zero value otherwise.

### GetRecoverytimeOk

`func (o *CloneRest) GetRecoverytimeOk() (*int64, bool)`

GetRecoverytimeOk returns a tuple with the Recoverytime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverytime

`func (o *CloneRest) SetRecoverytime(v int64)`

SetRecoverytime sets Recoverytime field to given value.

### HasRecoverytime

`func (o *CloneRest) HasRecoverytime() bool`

HasRecoverytime returns a boolean if a field has been set.

### GetPoweronvm

`func (o *CloneRest) GetPoweronvm() bool`

GetPoweronvm returns the Poweronvm field if non-nil, zero value otherwise.

### GetPoweronvmOk

`func (o *CloneRest) GetPoweronvmOk() (*bool, bool)`

GetPoweronvmOk returns a tuple with the Poweronvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoweronvm

`func (o *CloneRest) SetPoweronvm(v bool)`

SetPoweronvm sets Poweronvm field to given value.

### HasPoweronvm

`func (o *CloneRest) HasPoweronvm() bool`

HasPoweronvm returns a boolean if a field has been set.

### GetScripts

`func (o *CloneRest) GetScripts() []ScriptRest`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *CloneRest) GetScriptsOk() (*[]ScriptRest, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *CloneRest) SetScripts(v []ScriptRest)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *CloneRest) HasScripts() bool`

HasScripts returns a boolean if a field has been set.

### GetRestoreobjectmappings

`func (o *CloneRest) GetRestoreobjectmappings() []RestoreObjectMappingRest`

GetRestoreobjectmappings returns the Restoreobjectmappings field if non-nil, zero value otherwise.

### GetRestoreobjectmappingsOk

`func (o *CloneRest) GetRestoreobjectmappingsOk() (*[]RestoreObjectMappingRest, bool)`

GetRestoreobjectmappingsOk returns a tuple with the Restoreobjectmappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreobjectmappings

`func (o *CloneRest) SetRestoreobjectmappings(v []RestoreObjectMappingRest)`

SetRestoreobjectmappings sets Restoreobjectmappings field to given value.

### HasRestoreobjectmappings

`func (o *CloneRest) HasRestoreobjectmappings() bool`

HasRestoreobjectmappings returns a boolean if a field has been set.

### GetRestoreoptions

`func (o *CloneRest) GetRestoreoptions() []RestoreOptionRest`

GetRestoreoptions returns the Restoreoptions field if non-nil, zero value otherwise.

### GetRestoreoptionsOk

`func (o *CloneRest) GetRestoreoptionsOk() (*[]RestoreOptionRest, bool)`

GetRestoreoptionsOk returns a tuple with the Restoreoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreoptions

`func (o *CloneRest) SetRestoreoptions(v []RestoreOptionRest)`

SetRestoreoptions sets Restoreoptions field to given value.

### HasRestoreoptions

`func (o *CloneRest) HasRestoreoptions() bool`

HasRestoreoptions returns a boolean if a field has been set.

### GetProvisioningoptions

`func (o *CloneRest) GetProvisioningoptions() []ProvisioningOptionRest`

GetProvisioningoptions returns the Provisioningoptions field if non-nil, zero value otherwise.

### GetProvisioningoptionsOk

`func (o *CloneRest) GetProvisioningoptionsOk() (*[]ProvisioningOptionRest, bool)`

GetProvisioningoptionsOk returns a tuple with the Provisioningoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningoptions

`func (o *CloneRest) SetProvisioningoptions(v []ProvisioningOptionRest)`

SetProvisioningoptions sets Provisioningoptions field to given value.

### HasProvisioningoptions

`func (o *CloneRest) HasProvisioningoptions() bool`

HasProvisioningoptions returns a boolean if a field has been set.

### GetSelectedobjects

`func (o *CloneRest) GetSelectedobjects() []SelectedObjectRest`

GetSelectedobjects returns the Selectedobjects field if non-nil, zero value otherwise.

### GetSelectedobjectsOk

`func (o *CloneRest) GetSelectedobjectsOk() (*[]SelectedObjectRest, bool)`

GetSelectedobjectsOk returns a tuple with the Selectedobjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedobjects

`func (o *CloneRest) SetSelectedobjects(v []SelectedObjectRest)`

SetSelectedobjects sets Selectedobjects field to given value.

### HasSelectedobjects

`func (o *CloneRest) HasSelectedobjects() bool`

HasSelectedobjects returns a boolean if a field has been set.

### GetNfsoptions

`func (o *CloneRest) GetNfsoptions() NfsOptionsRest`

GetNfsoptions returns the Nfsoptions field if non-nil, zero value otherwise.

### GetNfsoptionsOk

`func (o *CloneRest) GetNfsoptionsOk() (*NfsOptionsRest, bool)`

GetNfsoptionsOk returns a tuple with the Nfsoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsoptions

`func (o *CloneRest) SetNfsoptions(v NfsOptionsRest)`

SetNfsoptions sets Nfsoptions field to given value.

### HasNfsoptions

`func (o *CloneRest) HasNfsoptions() bool`

HasNfsoptions returns a boolean if a field has been set.

### GetInstantmount

`func (o *CloneRest) GetInstantmount() bool`

GetInstantmount returns the Instantmount field if non-nil, zero value otherwise.

### GetInstantmountOk

`func (o *CloneRest) GetInstantmountOk() (*bool, bool)`

GetInstantmountOk returns a tuple with the Instantmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantmount

`func (o *CloneRest) SetInstantmount(v bool)`

SetInstantmount sets Instantmount field to given value.

### HasInstantmount

`func (o *CloneRest) HasInstantmount() bool`

HasInstantmount returns a boolean if a field has been set.

### GetRestorelocation

`func (o *CloneRest) GetRestorelocation() RestoreLocationRest`

GetRestorelocation returns the Restorelocation field if non-nil, zero value otherwise.

### GetRestorelocationOk

`func (o *CloneRest) GetRestorelocationOk() (*RestoreLocationRest, bool)`

GetRestorelocationOk returns a tuple with the Restorelocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestorelocation

`func (o *CloneRest) SetRestorelocation(v RestoreLocationRest)`

SetRestorelocation sets Restorelocation field to given value.

### HasRestorelocation

`func (o *CloneRest) HasRestorelocation() bool`

HasRestorelocation returns a boolean if a field has been set.

### GetMgmtserver

`func (o *CloneRest) GetMgmtserver() HostRest`

GetMgmtserver returns the Mgmtserver field if non-nil, zero value otherwise.

### GetMgmtserverOk

`func (o *CloneRest) GetMgmtserverOk() (*HostRest, bool)`

GetMgmtserverOk returns a tuple with the Mgmtserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtserver

`func (o *CloneRest) SetMgmtserver(v HostRest)`

SetMgmtserver sets Mgmtserver field to given value.

### HasMgmtserver

`func (o *CloneRest) HasMgmtserver() bool`

HasMgmtserver returns a boolean if a field has been set.

### GetHypervisor

`func (o *CloneRest) GetHypervisor() HostRest`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *CloneRest) GetHypervisorOk() (*HostRest, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisor

`func (o *CloneRest) SetHypervisor(v HostRest)`

SetHypervisor sets Hypervisor field to given value.

### HasHypervisor

`func (o *CloneRest) HasHypervisor() bool`

HasHypervisor returns a boolean if a field has been set.

### GetVmname

`func (o *CloneRest) GetVmname() string`

GetVmname returns the Vmname field if non-nil, zero value otherwise.

### GetVmnameOk

`func (o *CloneRest) GetVmnameOk() (*string, bool)`

GetVmnameOk returns a tuple with the Vmname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmname

`func (o *CloneRest) SetVmname(v string)`

SetVmname sets Vmname field to given value.

### HasVmname

`func (o *CloneRest) HasVmname() bool`

HasVmname returns a boolean if a field has been set.

### GetEsxhost

`func (o *CloneRest) GetEsxhost() HostRest`

GetEsxhost returns the Esxhost field if non-nil, zero value otherwise.

### GetEsxhostOk

`func (o *CloneRest) GetEsxhostOk() (*HostRest, bool)`

GetEsxhostOk returns a tuple with the Esxhost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsxhost

`func (o *CloneRest) SetEsxhost(v HostRest)`

SetEsxhost sets Esxhost field to given value.

### HasEsxhost

`func (o *CloneRest) HasEsxhost() bool`

HasEsxhost returns a boolean if a field has been set.

### GetContainer

`func (o *CloneRest) GetContainer() bool`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *CloneRest) GetContainerOk() (*bool, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *CloneRest) SetContainer(v bool)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *CloneRest) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetAllowedips

`func (o *CloneRest) GetAllowedips() []string`

GetAllowedips returns the Allowedips field if non-nil, zero value otherwise.

### GetAllowedipsOk

`func (o *CloneRest) GetAllowedipsOk() (*[]string, bool)`

GetAllowedipsOk returns a tuple with the Allowedips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedips

`func (o *CloneRest) SetAllowedips(v []string)`

SetAllowedips sets Allowedips field to given value.

### HasAllowedips

`func (o *CloneRest) HasAllowedips() bool`

HasAllowedips returns a boolean if a field has been set.

### GetId

`func (o *CloneRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloneRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloneRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloneRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *CloneRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CloneRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CloneRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CloneRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *CloneRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *CloneRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *CloneRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *CloneRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *CloneRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *CloneRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *CloneRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *CloneRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# BackupRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mountmigrateflag** | Pointer to **bool** |  | [optional] 
**Eligiblestoragetypes** | Pointer to **string** |  | [optional] 
**Targetstoragetype** | Pointer to **string** |  | [optional] 
**Migrationstate** | Pointer to **string** |  | [optional] 
**Datafilemovestarted** | Pointer to **bool** |  | [optional] 
**Migratemapping** | Pointer to **string** | source to target storage mapping | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Flags** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Host** | Pointer to [**HostRest**](HostRest.md) |  | [optional] 
**Expiration** | Pointer to **int64** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Application** | Pointer to [**ApplicationRest**](ApplicationRest.md) |  | [optional] 
**Jobclass** | Pointer to **string** |  | [optional] 
**Srcid** | Pointer to **string** |  | [optional] 
**Clusterid** | Pointer to **string** |  | [optional] 
**Appname** | Pointer to **string** |  | [optional] 
**Modifydate** | Pointer to **int64** |  | [optional] 
**Isasm** | Pointer to **bool** |  | [optional] 
**Apptype** | Pointer to **string** |  | [optional] 
**Snapshottype** | Pointer to **string** |  | [optional] 
**Consumedsize** | Pointer to **int64** |  | [optional] 
**Catalogstate** | Pointer to **string** |  | [optional] 
**Backupname** | Pointer to **string** |  | [optional] 
**Sourceuds** | Pointer to **string** |  | [optional] 
**Originatinguds** | Pointer to **string** |  | [optional] 
**Targetuds** | Pointer to **string** |  | [optional] 
**Appclass** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterRest**](ClusterRest.md) |  | [optional] 
**Diskpool** | Pointer to [**DiskPoolRest**](DiskPoolRest.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Transport** | Pointer to **string** |  | [optional] 
**Multiregion** | Pointer to **string** |  | [optional] 
**Sensitivity** | Pointer to **int32** |  | [optional] 
**Originalbackupid** | Pointer to **string** |  | [optional] 
**Backupdate** | Pointer to **int64** |  | [optional] 
**Immutabilitydate** | Pointer to **int64** |  | [optional] 
**Consistencydate** | Pointer to **int64** |  | [optional] 
**Mappedhost** | Pointer to **string** |  | [optional] 
**Virtualsize** | Pointer to **string** |  | [optional] 
**Slpname** | Pointer to **string** |  | [optional] 
**Sltname** | Pointer to **string** |  | [optional] 
**Policyname** | Pointer to **string** |  | [optional] 
**Characteristic** | Pointer to **string** |  | [optional] 
**Sourceimage** | Pointer to [**BackupRest**](BackupRest.md) |  | [optional] 
**Snapshotlocation** | Pointer to **string** |  | [optional] 
**Endpit** | Pointer to **int64** |  | [optional] 
**Beginpit** | Pointer to **int64** |  | [optional] 
**Uniquehostname** | Pointer to **string** |  | [optional] 
**Modifiedbytes** | Pointer to **string** |  | [optional] 
**Hasdependency** | Pointer to **bool** |  | [optional] 
**Vaultowner** | Pointer to **bool** |  | [optional] 
**Mountedhost** | Pointer to [**[]HostRest**](HostRest.md) |  | [optional] 
**Jobclasscode** | Pointer to **int32** |  | [optional] 
**Cloudcredential** | Pointer to [**CloudCredentialRest**](CloudCredentialRest.md) |  | [optional] 
**Releaselogs** | Pointer to **bool** |  | [optional] 
**Retainlogs** | Pointer to **bool** |  | [optional] 
**Sourcediskpool** | Pointer to [**DiskPoolRest**](DiskPoolRest.md) |  | [optional] 
**Clonemigrateeligible** | Pointer to **bool** |  | [optional] 
**MigrateFrequency** | Pointer to **int32** |  | [optional] 
**MigrateCopythreadcount** | Pointer to **int32** |  | [optional] 
**MigrateConfigured** | Pointer to **bool** |  | [optional] 
**Yaml** | Pointer to **string** |  | [optional] 
**Yamlavailable** | Pointer to **bool** |  | [optional] 
**Allowedips** | Pointer to **[]string** |  | [optional] 
**Volgroupname** | Pointer to **string** |  | [optional] 
**Racnodelist** | Pointer to **[]string** |  | [optional] 
**Restorableobjects** | Pointer to [**[]RestorableRest**](RestorableRest.md) |  | [optional] 
**Powerfactor** | Pointer to **int32** |  | [optional] 
**Provisioningoptions** | Pointer to [**[]KeyValueRest**](KeyValueRest.md) |  | [optional] 
**Vmname** | Pointer to **string** |  | [optional] 
**Copies** | Pointer to [**[]BackupRest**](BackupRest.md) |  | [optional] 
**Hasmountedmap** | Pointer to **bool** |  | [optional] 
**Childapp** | Pointer to [**ApplicationRest**](ApplicationRest.md) |  | [optional] 
**Restorelock** | Pointer to **string** |  | [optional] 
**Mountcapacity** | Pointer to **int64** |  | [optional] 
**Nvolumes** | Pointer to **string** |  | [optional] 
**Originalbackup** | Pointer to [**BackupRest**](BackupRest.md) |  | [optional] 
**Dnsname** | Pointer to **string** |  | [optional] 
**Vmpath** | Pointer to **string** |  | [optional] 
**Esxhost** | Pointer to **string** |  | [optional] 
**Backuplock** | Pointer to **string** |  | [optional] 
**Mountedmountpoint** | Pointer to **string** |  | [optional] 
**Consistencymode** | Pointer to **string** |  | [optional] 
**Mountedvdisk** | Pointer to **string** |  | [optional] 
**Sourcemountpoint** | Pointer to **string** |  | [optional] 
**FlagsText** | Pointer to **[]string** |  | [optional] 
**Expirytries** | Pointer to **string** |  | [optional] 
**Logsequences** | Pointer to [**[]LogSequenceRest**](LogSequenceRest.md) |  | [optional] 
**Hosttimezone** | Pointer to **string** |  | [optional] 
**Hostisdst** | Pointer to **bool** |  | [optional] 
**Asmswitchcapable** | Pointer to **bool** |  | [optional] 
**Incarnation** | Pointer to **string** |  | [optional] 
**Hananodelist** | Pointer to **[]string** |  | [optional] 
**ApplicationAllocatedSize** | Pointer to **int64** |  | [optional] 
**Exportvolume** | Pointer to [**[]ExportVolumeInfoRest**](ExportVolumeInfoRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewBackupRest

`func NewBackupRest() *BackupRest`

NewBackupRest instantiates a new BackupRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestWithDefaults

`func NewBackupRestWithDefaults() *BackupRest`

NewBackupRestWithDefaults instantiates a new BackupRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountmigrateflag

`func (o *BackupRest) GetMountmigrateflag() bool`

GetMountmigrateflag returns the Mountmigrateflag field if non-nil, zero value otherwise.

### GetMountmigrateflagOk

`func (o *BackupRest) GetMountmigrateflagOk() (*bool, bool)`

GetMountmigrateflagOk returns a tuple with the Mountmigrateflag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountmigrateflag

`func (o *BackupRest) SetMountmigrateflag(v bool)`

SetMountmigrateflag sets Mountmigrateflag field to given value.

### HasMountmigrateflag

`func (o *BackupRest) HasMountmigrateflag() bool`

HasMountmigrateflag returns a boolean if a field has been set.

### GetEligiblestoragetypes

`func (o *BackupRest) GetEligiblestoragetypes() string`

GetEligiblestoragetypes returns the Eligiblestoragetypes field if non-nil, zero value otherwise.

### GetEligiblestoragetypesOk

`func (o *BackupRest) GetEligiblestoragetypesOk() (*string, bool)`

GetEligiblestoragetypesOk returns a tuple with the Eligiblestoragetypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligiblestoragetypes

`func (o *BackupRest) SetEligiblestoragetypes(v string)`

SetEligiblestoragetypes sets Eligiblestoragetypes field to given value.

### HasEligiblestoragetypes

`func (o *BackupRest) HasEligiblestoragetypes() bool`

HasEligiblestoragetypes returns a boolean if a field has been set.

### GetTargetstoragetype

`func (o *BackupRest) GetTargetstoragetype() string`

GetTargetstoragetype returns the Targetstoragetype field if non-nil, zero value otherwise.

### GetTargetstoragetypeOk

`func (o *BackupRest) GetTargetstoragetypeOk() (*string, bool)`

GetTargetstoragetypeOk returns a tuple with the Targetstoragetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetstoragetype

`func (o *BackupRest) SetTargetstoragetype(v string)`

SetTargetstoragetype sets Targetstoragetype field to given value.

### HasTargetstoragetype

`func (o *BackupRest) HasTargetstoragetype() bool`

HasTargetstoragetype returns a boolean if a field has been set.

### GetMigrationstate

`func (o *BackupRest) GetMigrationstate() string`

GetMigrationstate returns the Migrationstate field if non-nil, zero value otherwise.

### GetMigrationstateOk

`func (o *BackupRest) GetMigrationstateOk() (*string, bool)`

GetMigrationstateOk returns a tuple with the Migrationstate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationstate

`func (o *BackupRest) SetMigrationstate(v string)`

SetMigrationstate sets Migrationstate field to given value.

### HasMigrationstate

`func (o *BackupRest) HasMigrationstate() bool`

HasMigrationstate returns a boolean if a field has been set.

### GetDatafilemovestarted

`func (o *BackupRest) GetDatafilemovestarted() bool`

GetDatafilemovestarted returns the Datafilemovestarted field if non-nil, zero value otherwise.

### GetDatafilemovestartedOk

`func (o *BackupRest) GetDatafilemovestartedOk() (*bool, bool)`

GetDatafilemovestartedOk returns a tuple with the Datafilemovestarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatafilemovestarted

`func (o *BackupRest) SetDatafilemovestarted(v bool)`

SetDatafilemovestarted sets Datafilemovestarted field to given value.

### HasDatafilemovestarted

`func (o *BackupRest) HasDatafilemovestarted() bool`

HasDatafilemovestarted returns a boolean if a field has been set.

### GetMigratemapping

`func (o *BackupRest) GetMigratemapping() string`

GetMigratemapping returns the Migratemapping field if non-nil, zero value otherwise.

### GetMigratemappingOk

`func (o *BackupRest) GetMigratemappingOk() (*string, bool)`

GetMigratemappingOk returns a tuple with the Migratemapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratemapping

`func (o *BackupRest) SetMigratemapping(v string)`

SetMigratemapping sets Migratemapping field to given value.

### HasMigratemapping

`func (o *BackupRest) HasMigratemapping() bool`

HasMigratemapping returns a boolean if a field has been set.

### GetVersion

`func (o *BackupRest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BackupRest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BackupRest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BackupRest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetFlags

`func (o *BackupRest) GetFlags() string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *BackupRest) GetFlagsOk() (*string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *BackupRest) SetFlags(v string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *BackupRest) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetZone

`func (o *BackupRest) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *BackupRest) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *BackupRest) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *BackupRest) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetName

`func (o *BackupRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BackupRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHost

`func (o *BackupRest) GetHost() HostRest`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *BackupRest) GetHostOk() (*HostRest, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *BackupRest) SetHost(v HostRest)`

SetHost sets Host field to given value.

### HasHost

`func (o *BackupRest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetExpiration

`func (o *BackupRest) GetExpiration() int64`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *BackupRest) GetExpirationOk() (*int64, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *BackupRest) SetExpiration(v int64)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *BackupRest) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetLabel

`func (o *BackupRest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BackupRest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BackupRest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BackupRest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetApplication

`func (o *BackupRest) GetApplication() ApplicationRest`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *BackupRest) GetApplicationOk() (*ApplicationRest, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *BackupRest) SetApplication(v ApplicationRest)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *BackupRest) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetJobclass

`func (o *BackupRest) GetJobclass() string`

GetJobclass returns the Jobclass field if non-nil, zero value otherwise.

### GetJobclassOk

`func (o *BackupRest) GetJobclassOk() (*string, bool)`

GetJobclassOk returns a tuple with the Jobclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobclass

`func (o *BackupRest) SetJobclass(v string)`

SetJobclass sets Jobclass field to given value.

### HasJobclass

`func (o *BackupRest) HasJobclass() bool`

HasJobclass returns a boolean if a field has been set.

### GetSrcid

`func (o *BackupRest) GetSrcid() string`

GetSrcid returns the Srcid field if non-nil, zero value otherwise.

### GetSrcidOk

`func (o *BackupRest) GetSrcidOk() (*string, bool)`

GetSrcidOk returns a tuple with the Srcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcid

`func (o *BackupRest) SetSrcid(v string)`

SetSrcid sets Srcid field to given value.

### HasSrcid

`func (o *BackupRest) HasSrcid() bool`

HasSrcid returns a boolean if a field has been set.

### GetClusterid

`func (o *BackupRest) GetClusterid() string`

GetClusterid returns the Clusterid field if non-nil, zero value otherwise.

### GetClusteridOk

`func (o *BackupRest) GetClusteridOk() (*string, bool)`

GetClusteridOk returns a tuple with the Clusterid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterid

`func (o *BackupRest) SetClusterid(v string)`

SetClusterid sets Clusterid field to given value.

### HasClusterid

`func (o *BackupRest) HasClusterid() bool`

HasClusterid returns a boolean if a field has been set.

### GetAppname

`func (o *BackupRest) GetAppname() string`

GetAppname returns the Appname field if non-nil, zero value otherwise.

### GetAppnameOk

`func (o *BackupRest) GetAppnameOk() (*string, bool)`

GetAppnameOk returns a tuple with the Appname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppname

`func (o *BackupRest) SetAppname(v string)`

SetAppname sets Appname field to given value.

### HasAppname

`func (o *BackupRest) HasAppname() bool`

HasAppname returns a boolean if a field has been set.

### GetModifydate

`func (o *BackupRest) GetModifydate() int64`

GetModifydate returns the Modifydate field if non-nil, zero value otherwise.

### GetModifydateOk

`func (o *BackupRest) GetModifydateOk() (*int64, bool)`

GetModifydateOk returns a tuple with the Modifydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifydate

`func (o *BackupRest) SetModifydate(v int64)`

SetModifydate sets Modifydate field to given value.

### HasModifydate

`func (o *BackupRest) HasModifydate() bool`

HasModifydate returns a boolean if a field has been set.

### GetIsasm

`func (o *BackupRest) GetIsasm() bool`

GetIsasm returns the Isasm field if non-nil, zero value otherwise.

### GetIsasmOk

`func (o *BackupRest) GetIsasmOk() (*bool, bool)`

GetIsasmOk returns a tuple with the Isasm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsasm

`func (o *BackupRest) SetIsasm(v bool)`

SetIsasm sets Isasm field to given value.

### HasIsasm

`func (o *BackupRest) HasIsasm() bool`

HasIsasm returns a boolean if a field has been set.

### GetApptype

`func (o *BackupRest) GetApptype() string`

GetApptype returns the Apptype field if non-nil, zero value otherwise.

### GetApptypeOk

`func (o *BackupRest) GetApptypeOk() (*string, bool)`

GetApptypeOk returns a tuple with the Apptype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApptype

`func (o *BackupRest) SetApptype(v string)`

SetApptype sets Apptype field to given value.

### HasApptype

`func (o *BackupRest) HasApptype() bool`

HasApptype returns a boolean if a field has been set.

### GetSnapshottype

`func (o *BackupRest) GetSnapshottype() string`

GetSnapshottype returns the Snapshottype field if non-nil, zero value otherwise.

### GetSnapshottypeOk

`func (o *BackupRest) GetSnapshottypeOk() (*string, bool)`

GetSnapshottypeOk returns a tuple with the Snapshottype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshottype

`func (o *BackupRest) SetSnapshottype(v string)`

SetSnapshottype sets Snapshottype field to given value.

### HasSnapshottype

`func (o *BackupRest) HasSnapshottype() bool`

HasSnapshottype returns a boolean if a field has been set.

### GetConsumedsize

`func (o *BackupRest) GetConsumedsize() int64`

GetConsumedsize returns the Consumedsize field if non-nil, zero value otherwise.

### GetConsumedsizeOk

`func (o *BackupRest) GetConsumedsizeOk() (*int64, bool)`

GetConsumedsizeOk returns a tuple with the Consumedsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedsize

`func (o *BackupRest) SetConsumedsize(v int64)`

SetConsumedsize sets Consumedsize field to given value.

### HasConsumedsize

`func (o *BackupRest) HasConsumedsize() bool`

HasConsumedsize returns a boolean if a field has been set.

### GetCatalogstate

`func (o *BackupRest) GetCatalogstate() string`

GetCatalogstate returns the Catalogstate field if non-nil, zero value otherwise.

### GetCatalogstateOk

`func (o *BackupRest) GetCatalogstateOk() (*string, bool)`

GetCatalogstateOk returns a tuple with the Catalogstate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogstate

`func (o *BackupRest) SetCatalogstate(v string)`

SetCatalogstate sets Catalogstate field to given value.

### HasCatalogstate

`func (o *BackupRest) HasCatalogstate() bool`

HasCatalogstate returns a boolean if a field has been set.

### GetBackupname

`func (o *BackupRest) GetBackupname() string`

GetBackupname returns the Backupname field if non-nil, zero value otherwise.

### GetBackupnameOk

`func (o *BackupRest) GetBackupnameOk() (*string, bool)`

GetBackupnameOk returns a tuple with the Backupname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupname

`func (o *BackupRest) SetBackupname(v string)`

SetBackupname sets Backupname field to given value.

### HasBackupname

`func (o *BackupRest) HasBackupname() bool`

HasBackupname returns a boolean if a field has been set.

### GetSourceuds

`func (o *BackupRest) GetSourceuds() string`

GetSourceuds returns the Sourceuds field if non-nil, zero value otherwise.

### GetSourceudsOk

`func (o *BackupRest) GetSourceudsOk() (*string, bool)`

GetSourceudsOk returns a tuple with the Sourceuds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceuds

`func (o *BackupRest) SetSourceuds(v string)`

SetSourceuds sets Sourceuds field to given value.

### HasSourceuds

`func (o *BackupRest) HasSourceuds() bool`

HasSourceuds returns a boolean if a field has been set.

### GetOriginatinguds

`func (o *BackupRest) GetOriginatinguds() string`

GetOriginatinguds returns the Originatinguds field if non-nil, zero value otherwise.

### GetOriginatingudsOk

`func (o *BackupRest) GetOriginatingudsOk() (*string, bool)`

GetOriginatingudsOk returns a tuple with the Originatinguds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatinguds

`func (o *BackupRest) SetOriginatinguds(v string)`

SetOriginatinguds sets Originatinguds field to given value.

### HasOriginatinguds

`func (o *BackupRest) HasOriginatinguds() bool`

HasOriginatinguds returns a boolean if a field has been set.

### GetTargetuds

`func (o *BackupRest) GetTargetuds() string`

GetTargetuds returns the Targetuds field if non-nil, zero value otherwise.

### GetTargetudsOk

`func (o *BackupRest) GetTargetudsOk() (*string, bool)`

GetTargetudsOk returns a tuple with the Targetuds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetuds

`func (o *BackupRest) SetTargetuds(v string)`

SetTargetuds sets Targetuds field to given value.

### HasTargetuds

`func (o *BackupRest) HasTargetuds() bool`

HasTargetuds returns a boolean if a field has been set.

### GetAppclass

`func (o *BackupRest) GetAppclass() string`

GetAppclass returns the Appclass field if non-nil, zero value otherwise.

### GetAppclassOk

`func (o *BackupRest) GetAppclassOk() (*string, bool)`

GetAppclassOk returns a tuple with the Appclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppclass

`func (o *BackupRest) SetAppclass(v string)`

SetAppclass sets Appclass field to given value.

### HasAppclass

`func (o *BackupRest) HasAppclass() bool`

HasAppclass returns a boolean if a field has been set.

### GetCluster

`func (o *BackupRest) GetCluster() ClusterRest`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *BackupRest) GetClusterOk() (*ClusterRest, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *BackupRest) SetCluster(v ClusterRest)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *BackupRest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDiskpool

`func (o *BackupRest) GetDiskpool() DiskPoolRest`

GetDiskpool returns the Diskpool field if non-nil, zero value otherwise.

### GetDiskpoolOk

`func (o *BackupRest) GetDiskpoolOk() (*DiskPoolRest, bool)`

GetDiskpoolOk returns a tuple with the Diskpool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskpool

`func (o *BackupRest) SetDiskpool(v DiskPoolRest)`

SetDiskpool sets Diskpool field to given value.

### HasDiskpool

`func (o *BackupRest) HasDiskpool() bool`

HasDiskpool returns a boolean if a field has been set.

### GetStatus

`func (o *BackupRest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BackupRest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BackupRest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BackupRest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransport

`func (o *BackupRest) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *BackupRest) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *BackupRest) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *BackupRest) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetMultiregion

`func (o *BackupRest) GetMultiregion() string`

GetMultiregion returns the Multiregion field if non-nil, zero value otherwise.

### GetMultiregionOk

`func (o *BackupRest) GetMultiregionOk() (*string, bool)`

GetMultiregionOk returns a tuple with the Multiregion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiregion

`func (o *BackupRest) SetMultiregion(v string)`

SetMultiregion sets Multiregion field to given value.

### HasMultiregion

`func (o *BackupRest) HasMultiregion() bool`

HasMultiregion returns a boolean if a field has been set.

### GetSensitivity

`func (o *BackupRest) GetSensitivity() int32`

GetSensitivity returns the Sensitivity field if non-nil, zero value otherwise.

### GetSensitivityOk

`func (o *BackupRest) GetSensitivityOk() (*int32, bool)`

GetSensitivityOk returns a tuple with the Sensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitivity

`func (o *BackupRest) SetSensitivity(v int32)`

SetSensitivity sets Sensitivity field to given value.

### HasSensitivity

`func (o *BackupRest) HasSensitivity() bool`

HasSensitivity returns a boolean if a field has been set.

### GetOriginalbackupid

`func (o *BackupRest) GetOriginalbackupid() string`

GetOriginalbackupid returns the Originalbackupid field if non-nil, zero value otherwise.

### GetOriginalbackupidOk

`func (o *BackupRest) GetOriginalbackupidOk() (*string, bool)`

GetOriginalbackupidOk returns a tuple with the Originalbackupid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalbackupid

`func (o *BackupRest) SetOriginalbackupid(v string)`

SetOriginalbackupid sets Originalbackupid field to given value.

### HasOriginalbackupid

`func (o *BackupRest) HasOriginalbackupid() bool`

HasOriginalbackupid returns a boolean if a field has been set.

### GetBackupdate

`func (o *BackupRest) GetBackupdate() int64`

GetBackupdate returns the Backupdate field if non-nil, zero value otherwise.

### GetBackupdateOk

`func (o *BackupRest) GetBackupdateOk() (*int64, bool)`

GetBackupdateOk returns a tuple with the Backupdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupdate

`func (o *BackupRest) SetBackupdate(v int64)`

SetBackupdate sets Backupdate field to given value.

### HasBackupdate

`func (o *BackupRest) HasBackupdate() bool`

HasBackupdate returns a boolean if a field has been set.

### GetImmutabilitydate

`func (o *BackupRest) GetImmutabilitydate() int64`

GetImmutabilitydate returns the Immutabilitydate field if non-nil, zero value otherwise.

### GetImmutabilitydateOk

`func (o *BackupRest) GetImmutabilitydateOk() (*int64, bool)`

GetImmutabilitydateOk returns a tuple with the Immutabilitydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutabilitydate

`func (o *BackupRest) SetImmutabilitydate(v int64)`

SetImmutabilitydate sets Immutabilitydate field to given value.

### HasImmutabilitydate

`func (o *BackupRest) HasImmutabilitydate() bool`

HasImmutabilitydate returns a boolean if a field has been set.

### GetConsistencydate

`func (o *BackupRest) GetConsistencydate() int64`

GetConsistencydate returns the Consistencydate field if non-nil, zero value otherwise.

### GetConsistencydateOk

`func (o *BackupRest) GetConsistencydateOk() (*int64, bool)`

GetConsistencydateOk returns a tuple with the Consistencydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistencydate

`func (o *BackupRest) SetConsistencydate(v int64)`

SetConsistencydate sets Consistencydate field to given value.

### HasConsistencydate

`func (o *BackupRest) HasConsistencydate() bool`

HasConsistencydate returns a boolean if a field has been set.

### GetMappedhost

`func (o *BackupRest) GetMappedhost() string`

GetMappedhost returns the Mappedhost field if non-nil, zero value otherwise.

### GetMappedhostOk

`func (o *BackupRest) GetMappedhostOk() (*string, bool)`

GetMappedhostOk returns a tuple with the Mappedhost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedhost

`func (o *BackupRest) SetMappedhost(v string)`

SetMappedhost sets Mappedhost field to given value.

### HasMappedhost

`func (o *BackupRest) HasMappedhost() bool`

HasMappedhost returns a boolean if a field has been set.

### GetVirtualsize

`func (o *BackupRest) GetVirtualsize() string`

GetVirtualsize returns the Virtualsize field if non-nil, zero value otherwise.

### GetVirtualsizeOk

`func (o *BackupRest) GetVirtualsizeOk() (*string, bool)`

GetVirtualsizeOk returns a tuple with the Virtualsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualsize

`func (o *BackupRest) SetVirtualsize(v string)`

SetVirtualsize sets Virtualsize field to given value.

### HasVirtualsize

`func (o *BackupRest) HasVirtualsize() bool`

HasVirtualsize returns a boolean if a field has been set.

### GetSlpname

`func (o *BackupRest) GetSlpname() string`

GetSlpname returns the Slpname field if non-nil, zero value otherwise.

### GetSlpnameOk

`func (o *BackupRest) GetSlpnameOk() (*string, bool)`

GetSlpnameOk returns a tuple with the Slpname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlpname

`func (o *BackupRest) SetSlpname(v string)`

SetSlpname sets Slpname field to given value.

### HasSlpname

`func (o *BackupRest) HasSlpname() bool`

HasSlpname returns a boolean if a field has been set.

### GetSltname

`func (o *BackupRest) GetSltname() string`

GetSltname returns the Sltname field if non-nil, zero value otherwise.

### GetSltnameOk

`func (o *BackupRest) GetSltnameOk() (*string, bool)`

GetSltnameOk returns a tuple with the Sltname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSltname

`func (o *BackupRest) SetSltname(v string)`

SetSltname sets Sltname field to given value.

### HasSltname

`func (o *BackupRest) HasSltname() bool`

HasSltname returns a boolean if a field has been set.

### GetPolicyname

`func (o *BackupRest) GetPolicyname() string`

GetPolicyname returns the Policyname field if non-nil, zero value otherwise.

### GetPolicynameOk

`func (o *BackupRest) GetPolicynameOk() (*string, bool)`

GetPolicynameOk returns a tuple with the Policyname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyname

`func (o *BackupRest) SetPolicyname(v string)`

SetPolicyname sets Policyname field to given value.

### HasPolicyname

`func (o *BackupRest) HasPolicyname() bool`

HasPolicyname returns a boolean if a field has been set.

### GetCharacteristic

`func (o *BackupRest) GetCharacteristic() string`

GetCharacteristic returns the Characteristic field if non-nil, zero value otherwise.

### GetCharacteristicOk

`func (o *BackupRest) GetCharacteristicOk() (*string, bool)`

GetCharacteristicOk returns a tuple with the Characteristic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacteristic

`func (o *BackupRest) SetCharacteristic(v string)`

SetCharacteristic sets Characteristic field to given value.

### HasCharacteristic

`func (o *BackupRest) HasCharacteristic() bool`

HasCharacteristic returns a boolean if a field has been set.

### GetSourceimage

`func (o *BackupRest) GetSourceimage() BackupRest`

GetSourceimage returns the Sourceimage field if non-nil, zero value otherwise.

### GetSourceimageOk

`func (o *BackupRest) GetSourceimageOk() (*BackupRest, bool)`

GetSourceimageOk returns a tuple with the Sourceimage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceimage

`func (o *BackupRest) SetSourceimage(v BackupRest)`

SetSourceimage sets Sourceimage field to given value.

### HasSourceimage

`func (o *BackupRest) HasSourceimage() bool`

HasSourceimage returns a boolean if a field has been set.

### GetSnapshotlocation

`func (o *BackupRest) GetSnapshotlocation() string`

GetSnapshotlocation returns the Snapshotlocation field if non-nil, zero value otherwise.

### GetSnapshotlocationOk

`func (o *BackupRest) GetSnapshotlocationOk() (*string, bool)`

GetSnapshotlocationOk returns a tuple with the Snapshotlocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotlocation

`func (o *BackupRest) SetSnapshotlocation(v string)`

SetSnapshotlocation sets Snapshotlocation field to given value.

### HasSnapshotlocation

`func (o *BackupRest) HasSnapshotlocation() bool`

HasSnapshotlocation returns a boolean if a field has been set.

### GetEndpit

`func (o *BackupRest) GetEndpit() int64`

GetEndpit returns the Endpit field if non-nil, zero value otherwise.

### GetEndpitOk

`func (o *BackupRest) GetEndpitOk() (*int64, bool)`

GetEndpitOk returns a tuple with the Endpit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpit

`func (o *BackupRest) SetEndpit(v int64)`

SetEndpit sets Endpit field to given value.

### HasEndpit

`func (o *BackupRest) HasEndpit() bool`

HasEndpit returns a boolean if a field has been set.

### GetBeginpit

`func (o *BackupRest) GetBeginpit() int64`

GetBeginpit returns the Beginpit field if non-nil, zero value otherwise.

### GetBeginpitOk

`func (o *BackupRest) GetBeginpitOk() (*int64, bool)`

GetBeginpitOk returns a tuple with the Beginpit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginpit

`func (o *BackupRest) SetBeginpit(v int64)`

SetBeginpit sets Beginpit field to given value.

### HasBeginpit

`func (o *BackupRest) HasBeginpit() bool`

HasBeginpit returns a boolean if a field has been set.

### GetUniquehostname

`func (o *BackupRest) GetUniquehostname() string`

GetUniquehostname returns the Uniquehostname field if non-nil, zero value otherwise.

### GetUniquehostnameOk

`func (o *BackupRest) GetUniquehostnameOk() (*string, bool)`

GetUniquehostnameOk returns a tuple with the Uniquehostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniquehostname

`func (o *BackupRest) SetUniquehostname(v string)`

SetUniquehostname sets Uniquehostname field to given value.

### HasUniquehostname

`func (o *BackupRest) HasUniquehostname() bool`

HasUniquehostname returns a boolean if a field has been set.

### GetModifiedbytes

`func (o *BackupRest) GetModifiedbytes() string`

GetModifiedbytes returns the Modifiedbytes field if non-nil, zero value otherwise.

### GetModifiedbytesOk

`func (o *BackupRest) GetModifiedbytesOk() (*string, bool)`

GetModifiedbytesOk returns a tuple with the Modifiedbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedbytes

`func (o *BackupRest) SetModifiedbytes(v string)`

SetModifiedbytes sets Modifiedbytes field to given value.

### HasModifiedbytes

`func (o *BackupRest) HasModifiedbytes() bool`

HasModifiedbytes returns a boolean if a field has been set.

### GetHasdependency

`func (o *BackupRest) GetHasdependency() bool`

GetHasdependency returns the Hasdependency field if non-nil, zero value otherwise.

### GetHasdependencyOk

`func (o *BackupRest) GetHasdependencyOk() (*bool, bool)`

GetHasdependencyOk returns a tuple with the Hasdependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasdependency

`func (o *BackupRest) SetHasdependency(v bool)`

SetHasdependency sets Hasdependency field to given value.

### HasHasdependency

`func (o *BackupRest) HasHasdependency() bool`

HasHasdependency returns a boolean if a field has been set.

### GetVaultowner

`func (o *BackupRest) GetVaultowner() bool`

GetVaultowner returns the Vaultowner field if non-nil, zero value otherwise.

### GetVaultownerOk

`func (o *BackupRest) GetVaultownerOk() (*bool, bool)`

GetVaultownerOk returns a tuple with the Vaultowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultowner

`func (o *BackupRest) SetVaultowner(v bool)`

SetVaultowner sets Vaultowner field to given value.

### HasVaultowner

`func (o *BackupRest) HasVaultowner() bool`

HasVaultowner returns a boolean if a field has been set.

### GetMountedhost

`func (o *BackupRest) GetMountedhost() []HostRest`

GetMountedhost returns the Mountedhost field if non-nil, zero value otherwise.

### GetMountedhostOk

`func (o *BackupRest) GetMountedhostOk() (*[]HostRest, bool)`

GetMountedhostOk returns a tuple with the Mountedhost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountedhost

`func (o *BackupRest) SetMountedhost(v []HostRest)`

SetMountedhost sets Mountedhost field to given value.

### HasMountedhost

`func (o *BackupRest) HasMountedhost() bool`

HasMountedhost returns a boolean if a field has been set.

### GetJobclasscode

`func (o *BackupRest) GetJobclasscode() int32`

GetJobclasscode returns the Jobclasscode field if non-nil, zero value otherwise.

### GetJobclasscodeOk

`func (o *BackupRest) GetJobclasscodeOk() (*int32, bool)`

GetJobclasscodeOk returns a tuple with the Jobclasscode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobclasscode

`func (o *BackupRest) SetJobclasscode(v int32)`

SetJobclasscode sets Jobclasscode field to given value.

### HasJobclasscode

`func (o *BackupRest) HasJobclasscode() bool`

HasJobclasscode returns a boolean if a field has been set.

### GetCloudcredential

`func (o *BackupRest) GetCloudcredential() CloudCredentialRest`

GetCloudcredential returns the Cloudcredential field if non-nil, zero value otherwise.

### GetCloudcredentialOk

`func (o *BackupRest) GetCloudcredentialOk() (*CloudCredentialRest, bool)`

GetCloudcredentialOk returns a tuple with the Cloudcredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudcredential

`func (o *BackupRest) SetCloudcredential(v CloudCredentialRest)`

SetCloudcredential sets Cloudcredential field to given value.

### HasCloudcredential

`func (o *BackupRest) HasCloudcredential() bool`

HasCloudcredential returns a boolean if a field has been set.

### GetReleaselogs

`func (o *BackupRest) GetReleaselogs() bool`

GetReleaselogs returns the Releaselogs field if non-nil, zero value otherwise.

### GetReleaselogsOk

`func (o *BackupRest) GetReleaselogsOk() (*bool, bool)`

GetReleaselogsOk returns a tuple with the Releaselogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaselogs

`func (o *BackupRest) SetReleaselogs(v bool)`

SetReleaselogs sets Releaselogs field to given value.

### HasReleaselogs

`func (o *BackupRest) HasReleaselogs() bool`

HasReleaselogs returns a boolean if a field has been set.

### GetRetainlogs

`func (o *BackupRest) GetRetainlogs() bool`

GetRetainlogs returns the Retainlogs field if non-nil, zero value otherwise.

### GetRetainlogsOk

`func (o *BackupRest) GetRetainlogsOk() (*bool, bool)`

GetRetainlogsOk returns a tuple with the Retainlogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainlogs

`func (o *BackupRest) SetRetainlogs(v bool)`

SetRetainlogs sets Retainlogs field to given value.

### HasRetainlogs

`func (o *BackupRest) HasRetainlogs() bool`

HasRetainlogs returns a boolean if a field has been set.

### GetSourcediskpool

`func (o *BackupRest) GetSourcediskpool() DiskPoolRest`

GetSourcediskpool returns the Sourcediskpool field if non-nil, zero value otherwise.

### GetSourcediskpoolOk

`func (o *BackupRest) GetSourcediskpoolOk() (*DiskPoolRest, bool)`

GetSourcediskpoolOk returns a tuple with the Sourcediskpool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcediskpool

`func (o *BackupRest) SetSourcediskpool(v DiskPoolRest)`

SetSourcediskpool sets Sourcediskpool field to given value.

### HasSourcediskpool

`func (o *BackupRest) HasSourcediskpool() bool`

HasSourcediskpool returns a boolean if a field has been set.

### GetClonemigrateeligible

`func (o *BackupRest) GetClonemigrateeligible() bool`

GetClonemigrateeligible returns the Clonemigrateeligible field if non-nil, zero value otherwise.

### GetClonemigrateeligibleOk

`func (o *BackupRest) GetClonemigrateeligibleOk() (*bool, bool)`

GetClonemigrateeligibleOk returns a tuple with the Clonemigrateeligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonemigrateeligible

`func (o *BackupRest) SetClonemigrateeligible(v bool)`

SetClonemigrateeligible sets Clonemigrateeligible field to given value.

### HasClonemigrateeligible

`func (o *BackupRest) HasClonemigrateeligible() bool`

HasClonemigrateeligible returns a boolean if a field has been set.

### GetMigrateFrequency

`func (o *BackupRest) GetMigrateFrequency() int32`

GetMigrateFrequency returns the MigrateFrequency field if non-nil, zero value otherwise.

### GetMigrateFrequencyOk

`func (o *BackupRest) GetMigrateFrequencyOk() (*int32, bool)`

GetMigrateFrequencyOk returns a tuple with the MigrateFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateFrequency

`func (o *BackupRest) SetMigrateFrequency(v int32)`

SetMigrateFrequency sets MigrateFrequency field to given value.

### HasMigrateFrequency

`func (o *BackupRest) HasMigrateFrequency() bool`

HasMigrateFrequency returns a boolean if a field has been set.

### GetMigrateCopythreadcount

`func (o *BackupRest) GetMigrateCopythreadcount() int32`

GetMigrateCopythreadcount returns the MigrateCopythreadcount field if non-nil, zero value otherwise.

### GetMigrateCopythreadcountOk

`func (o *BackupRest) GetMigrateCopythreadcountOk() (*int32, bool)`

GetMigrateCopythreadcountOk returns a tuple with the MigrateCopythreadcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateCopythreadcount

`func (o *BackupRest) SetMigrateCopythreadcount(v int32)`

SetMigrateCopythreadcount sets MigrateCopythreadcount field to given value.

### HasMigrateCopythreadcount

`func (o *BackupRest) HasMigrateCopythreadcount() bool`

HasMigrateCopythreadcount returns a boolean if a field has been set.

### GetMigrateConfigured

`func (o *BackupRest) GetMigrateConfigured() bool`

GetMigrateConfigured returns the MigrateConfigured field if non-nil, zero value otherwise.

### GetMigrateConfiguredOk

`func (o *BackupRest) GetMigrateConfiguredOk() (*bool, bool)`

GetMigrateConfiguredOk returns a tuple with the MigrateConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateConfigured

`func (o *BackupRest) SetMigrateConfigured(v bool)`

SetMigrateConfigured sets MigrateConfigured field to given value.

### HasMigrateConfigured

`func (o *BackupRest) HasMigrateConfigured() bool`

HasMigrateConfigured returns a boolean if a field has been set.

### GetYaml

`func (o *BackupRest) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *BackupRest) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *BackupRest) SetYaml(v string)`

SetYaml sets Yaml field to given value.

### HasYaml

`func (o *BackupRest) HasYaml() bool`

HasYaml returns a boolean if a field has been set.

### GetYamlavailable

`func (o *BackupRest) GetYamlavailable() bool`

GetYamlavailable returns the Yamlavailable field if non-nil, zero value otherwise.

### GetYamlavailableOk

`func (o *BackupRest) GetYamlavailableOk() (*bool, bool)`

GetYamlavailableOk returns a tuple with the Yamlavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYamlavailable

`func (o *BackupRest) SetYamlavailable(v bool)`

SetYamlavailable sets Yamlavailable field to given value.

### HasYamlavailable

`func (o *BackupRest) HasYamlavailable() bool`

HasYamlavailable returns a boolean if a field has been set.

### GetAllowedips

`func (o *BackupRest) GetAllowedips() []string`

GetAllowedips returns the Allowedips field if non-nil, zero value otherwise.

### GetAllowedipsOk

`func (o *BackupRest) GetAllowedipsOk() (*[]string, bool)`

GetAllowedipsOk returns a tuple with the Allowedips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedips

`func (o *BackupRest) SetAllowedips(v []string)`

SetAllowedips sets Allowedips field to given value.

### HasAllowedips

`func (o *BackupRest) HasAllowedips() bool`

HasAllowedips returns a boolean if a field has been set.

### GetVolgroupname

`func (o *BackupRest) GetVolgroupname() string`

GetVolgroupname returns the Volgroupname field if non-nil, zero value otherwise.

### GetVolgroupnameOk

`func (o *BackupRest) GetVolgroupnameOk() (*string, bool)`

GetVolgroupnameOk returns a tuple with the Volgroupname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolgroupname

`func (o *BackupRest) SetVolgroupname(v string)`

SetVolgroupname sets Volgroupname field to given value.

### HasVolgroupname

`func (o *BackupRest) HasVolgroupname() bool`

HasVolgroupname returns a boolean if a field has been set.

### GetRacnodelist

`func (o *BackupRest) GetRacnodelist() []string`

GetRacnodelist returns the Racnodelist field if non-nil, zero value otherwise.

### GetRacnodelistOk

`func (o *BackupRest) GetRacnodelistOk() (*[]string, bool)`

GetRacnodelistOk returns a tuple with the Racnodelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRacnodelist

`func (o *BackupRest) SetRacnodelist(v []string)`

SetRacnodelist sets Racnodelist field to given value.

### HasRacnodelist

`func (o *BackupRest) HasRacnodelist() bool`

HasRacnodelist returns a boolean if a field has been set.

### GetRestorableobjects

`func (o *BackupRest) GetRestorableobjects() []RestorableRest`

GetRestorableobjects returns the Restorableobjects field if non-nil, zero value otherwise.

### GetRestorableobjectsOk

`func (o *BackupRest) GetRestorableobjectsOk() (*[]RestorableRest, bool)`

GetRestorableobjectsOk returns a tuple with the Restorableobjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestorableobjects

`func (o *BackupRest) SetRestorableobjects(v []RestorableRest)`

SetRestorableobjects sets Restorableobjects field to given value.

### HasRestorableobjects

`func (o *BackupRest) HasRestorableobjects() bool`

HasRestorableobjects returns a boolean if a field has been set.

### GetPowerfactor

`func (o *BackupRest) GetPowerfactor() int32`

GetPowerfactor returns the Powerfactor field if non-nil, zero value otherwise.

### GetPowerfactorOk

`func (o *BackupRest) GetPowerfactorOk() (*int32, bool)`

GetPowerfactorOk returns a tuple with the Powerfactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerfactor

`func (o *BackupRest) SetPowerfactor(v int32)`

SetPowerfactor sets Powerfactor field to given value.

### HasPowerfactor

`func (o *BackupRest) HasPowerfactor() bool`

HasPowerfactor returns a boolean if a field has been set.

### GetProvisioningoptions

`func (o *BackupRest) GetProvisioningoptions() []KeyValueRest`

GetProvisioningoptions returns the Provisioningoptions field if non-nil, zero value otherwise.

### GetProvisioningoptionsOk

`func (o *BackupRest) GetProvisioningoptionsOk() (*[]KeyValueRest, bool)`

GetProvisioningoptionsOk returns a tuple with the Provisioningoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningoptions

`func (o *BackupRest) SetProvisioningoptions(v []KeyValueRest)`

SetProvisioningoptions sets Provisioningoptions field to given value.

### HasProvisioningoptions

`func (o *BackupRest) HasProvisioningoptions() bool`

HasProvisioningoptions returns a boolean if a field has been set.

### GetVmname

`func (o *BackupRest) GetVmname() string`

GetVmname returns the Vmname field if non-nil, zero value otherwise.

### GetVmnameOk

`func (o *BackupRest) GetVmnameOk() (*string, bool)`

GetVmnameOk returns a tuple with the Vmname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmname

`func (o *BackupRest) SetVmname(v string)`

SetVmname sets Vmname field to given value.

### HasVmname

`func (o *BackupRest) HasVmname() bool`

HasVmname returns a boolean if a field has been set.

### GetCopies

`func (o *BackupRest) GetCopies() []BackupRest`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *BackupRest) GetCopiesOk() (*[]BackupRest, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *BackupRest) SetCopies(v []BackupRest)`

SetCopies sets Copies field to given value.

### HasCopies

`func (o *BackupRest) HasCopies() bool`

HasCopies returns a boolean if a field has been set.

### GetHasmountedmap

`func (o *BackupRest) GetHasmountedmap() bool`

GetHasmountedmap returns the Hasmountedmap field if non-nil, zero value otherwise.

### GetHasmountedmapOk

`func (o *BackupRest) GetHasmountedmapOk() (*bool, bool)`

GetHasmountedmapOk returns a tuple with the Hasmountedmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasmountedmap

`func (o *BackupRest) SetHasmountedmap(v bool)`

SetHasmountedmap sets Hasmountedmap field to given value.

### HasHasmountedmap

`func (o *BackupRest) HasHasmountedmap() bool`

HasHasmountedmap returns a boolean if a field has been set.

### GetChildapp

`func (o *BackupRest) GetChildapp() ApplicationRest`

GetChildapp returns the Childapp field if non-nil, zero value otherwise.

### GetChildappOk

`func (o *BackupRest) GetChildappOk() (*ApplicationRest, bool)`

GetChildappOk returns a tuple with the Childapp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildapp

`func (o *BackupRest) SetChildapp(v ApplicationRest)`

SetChildapp sets Childapp field to given value.

### HasChildapp

`func (o *BackupRest) HasChildapp() bool`

HasChildapp returns a boolean if a field has been set.

### GetRestorelock

`func (o *BackupRest) GetRestorelock() string`

GetRestorelock returns the Restorelock field if non-nil, zero value otherwise.

### GetRestorelockOk

`func (o *BackupRest) GetRestorelockOk() (*string, bool)`

GetRestorelockOk returns a tuple with the Restorelock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestorelock

`func (o *BackupRest) SetRestorelock(v string)`

SetRestorelock sets Restorelock field to given value.

### HasRestorelock

`func (o *BackupRest) HasRestorelock() bool`

HasRestorelock returns a boolean if a field has been set.

### GetMountcapacity

`func (o *BackupRest) GetMountcapacity() int64`

GetMountcapacity returns the Mountcapacity field if non-nil, zero value otherwise.

### GetMountcapacityOk

`func (o *BackupRest) GetMountcapacityOk() (*int64, bool)`

GetMountcapacityOk returns a tuple with the Mountcapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountcapacity

`func (o *BackupRest) SetMountcapacity(v int64)`

SetMountcapacity sets Mountcapacity field to given value.

### HasMountcapacity

`func (o *BackupRest) HasMountcapacity() bool`

HasMountcapacity returns a boolean if a field has been set.

### GetNvolumes

`func (o *BackupRest) GetNvolumes() string`

GetNvolumes returns the Nvolumes field if non-nil, zero value otherwise.

### GetNvolumesOk

`func (o *BackupRest) GetNvolumesOk() (*string, bool)`

GetNvolumesOk returns a tuple with the Nvolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvolumes

`func (o *BackupRest) SetNvolumes(v string)`

SetNvolumes sets Nvolumes field to given value.

### HasNvolumes

`func (o *BackupRest) HasNvolumes() bool`

HasNvolumes returns a boolean if a field has been set.

### GetOriginalbackup

`func (o *BackupRest) GetOriginalbackup() BackupRest`

GetOriginalbackup returns the Originalbackup field if non-nil, zero value otherwise.

### GetOriginalbackupOk

`func (o *BackupRest) GetOriginalbackupOk() (*BackupRest, bool)`

GetOriginalbackupOk returns a tuple with the Originalbackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalbackup

`func (o *BackupRest) SetOriginalbackup(v BackupRest)`

SetOriginalbackup sets Originalbackup field to given value.

### HasOriginalbackup

`func (o *BackupRest) HasOriginalbackup() bool`

HasOriginalbackup returns a boolean if a field has been set.

### GetDnsname

`func (o *BackupRest) GetDnsname() string`

GetDnsname returns the Dnsname field if non-nil, zero value otherwise.

### GetDnsnameOk

`func (o *BackupRest) GetDnsnameOk() (*string, bool)`

GetDnsnameOk returns a tuple with the Dnsname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsname

`func (o *BackupRest) SetDnsname(v string)`

SetDnsname sets Dnsname field to given value.

### HasDnsname

`func (o *BackupRest) HasDnsname() bool`

HasDnsname returns a boolean if a field has been set.

### GetVmpath

`func (o *BackupRest) GetVmpath() string`

GetVmpath returns the Vmpath field if non-nil, zero value otherwise.

### GetVmpathOk

`func (o *BackupRest) GetVmpathOk() (*string, bool)`

GetVmpathOk returns a tuple with the Vmpath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmpath

`func (o *BackupRest) SetVmpath(v string)`

SetVmpath sets Vmpath field to given value.

### HasVmpath

`func (o *BackupRest) HasVmpath() bool`

HasVmpath returns a boolean if a field has been set.

### GetEsxhost

`func (o *BackupRest) GetEsxhost() string`

GetEsxhost returns the Esxhost field if non-nil, zero value otherwise.

### GetEsxhostOk

`func (o *BackupRest) GetEsxhostOk() (*string, bool)`

GetEsxhostOk returns a tuple with the Esxhost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsxhost

`func (o *BackupRest) SetEsxhost(v string)`

SetEsxhost sets Esxhost field to given value.

### HasEsxhost

`func (o *BackupRest) HasEsxhost() bool`

HasEsxhost returns a boolean if a field has been set.

### GetBackuplock

`func (o *BackupRest) GetBackuplock() string`

GetBackuplock returns the Backuplock field if non-nil, zero value otherwise.

### GetBackuplockOk

`func (o *BackupRest) GetBackuplockOk() (*string, bool)`

GetBackuplockOk returns a tuple with the Backuplock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackuplock

`func (o *BackupRest) SetBackuplock(v string)`

SetBackuplock sets Backuplock field to given value.

### HasBackuplock

`func (o *BackupRest) HasBackuplock() bool`

HasBackuplock returns a boolean if a field has been set.

### GetMountedmountpoint

`func (o *BackupRest) GetMountedmountpoint() string`

GetMountedmountpoint returns the Mountedmountpoint field if non-nil, zero value otherwise.

### GetMountedmountpointOk

`func (o *BackupRest) GetMountedmountpointOk() (*string, bool)`

GetMountedmountpointOk returns a tuple with the Mountedmountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountedmountpoint

`func (o *BackupRest) SetMountedmountpoint(v string)`

SetMountedmountpoint sets Mountedmountpoint field to given value.

### HasMountedmountpoint

`func (o *BackupRest) HasMountedmountpoint() bool`

HasMountedmountpoint returns a boolean if a field has been set.

### GetConsistencymode

`func (o *BackupRest) GetConsistencymode() string`

GetConsistencymode returns the Consistencymode field if non-nil, zero value otherwise.

### GetConsistencymodeOk

`func (o *BackupRest) GetConsistencymodeOk() (*string, bool)`

GetConsistencymodeOk returns a tuple with the Consistencymode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistencymode

`func (o *BackupRest) SetConsistencymode(v string)`

SetConsistencymode sets Consistencymode field to given value.

### HasConsistencymode

`func (o *BackupRest) HasConsistencymode() bool`

HasConsistencymode returns a boolean if a field has been set.

### GetMountedvdisk

`func (o *BackupRest) GetMountedvdisk() string`

GetMountedvdisk returns the Mountedvdisk field if non-nil, zero value otherwise.

### GetMountedvdiskOk

`func (o *BackupRest) GetMountedvdiskOk() (*string, bool)`

GetMountedvdiskOk returns a tuple with the Mountedvdisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountedvdisk

`func (o *BackupRest) SetMountedvdisk(v string)`

SetMountedvdisk sets Mountedvdisk field to given value.

### HasMountedvdisk

`func (o *BackupRest) HasMountedvdisk() bool`

HasMountedvdisk returns a boolean if a field has been set.

### GetSourcemountpoint

`func (o *BackupRest) GetSourcemountpoint() string`

GetSourcemountpoint returns the Sourcemountpoint field if non-nil, zero value otherwise.

### GetSourcemountpointOk

`func (o *BackupRest) GetSourcemountpointOk() (*string, bool)`

GetSourcemountpointOk returns a tuple with the Sourcemountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcemountpoint

`func (o *BackupRest) SetSourcemountpoint(v string)`

SetSourcemountpoint sets Sourcemountpoint field to given value.

### HasSourcemountpoint

`func (o *BackupRest) HasSourcemountpoint() bool`

HasSourcemountpoint returns a boolean if a field has been set.

### GetFlagsText

`func (o *BackupRest) GetFlagsText() []string`

GetFlagsText returns the FlagsText field if non-nil, zero value otherwise.

### GetFlagsTextOk

`func (o *BackupRest) GetFlagsTextOk() (*[]string, bool)`

GetFlagsTextOk returns a tuple with the FlagsText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagsText

`func (o *BackupRest) SetFlagsText(v []string)`

SetFlagsText sets FlagsText field to given value.

### HasFlagsText

`func (o *BackupRest) HasFlagsText() bool`

HasFlagsText returns a boolean if a field has been set.

### GetExpirytries

`func (o *BackupRest) GetExpirytries() string`

GetExpirytries returns the Expirytries field if non-nil, zero value otherwise.

### GetExpirytriesOk

`func (o *BackupRest) GetExpirytriesOk() (*string, bool)`

GetExpirytriesOk returns a tuple with the Expirytries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirytries

`func (o *BackupRest) SetExpirytries(v string)`

SetExpirytries sets Expirytries field to given value.

### HasExpirytries

`func (o *BackupRest) HasExpirytries() bool`

HasExpirytries returns a boolean if a field has been set.

### GetLogsequences

`func (o *BackupRest) GetLogsequences() []LogSequenceRest`

GetLogsequences returns the Logsequences field if non-nil, zero value otherwise.

### GetLogsequencesOk

`func (o *BackupRest) GetLogsequencesOk() (*[]LogSequenceRest, bool)`

GetLogsequencesOk returns a tuple with the Logsequences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsequences

`func (o *BackupRest) SetLogsequences(v []LogSequenceRest)`

SetLogsequences sets Logsequences field to given value.

### HasLogsequences

`func (o *BackupRest) HasLogsequences() bool`

HasLogsequences returns a boolean if a field has been set.

### GetHosttimezone

`func (o *BackupRest) GetHosttimezone() string`

GetHosttimezone returns the Hosttimezone field if non-nil, zero value otherwise.

### GetHosttimezoneOk

`func (o *BackupRest) GetHosttimezoneOk() (*string, bool)`

GetHosttimezoneOk returns a tuple with the Hosttimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosttimezone

`func (o *BackupRest) SetHosttimezone(v string)`

SetHosttimezone sets Hosttimezone field to given value.

### HasHosttimezone

`func (o *BackupRest) HasHosttimezone() bool`

HasHosttimezone returns a boolean if a field has been set.

### GetHostisdst

`func (o *BackupRest) GetHostisdst() bool`

GetHostisdst returns the Hostisdst field if non-nil, zero value otherwise.

### GetHostisdstOk

`func (o *BackupRest) GetHostisdstOk() (*bool, bool)`

GetHostisdstOk returns a tuple with the Hostisdst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostisdst

`func (o *BackupRest) SetHostisdst(v bool)`

SetHostisdst sets Hostisdst field to given value.

### HasHostisdst

`func (o *BackupRest) HasHostisdst() bool`

HasHostisdst returns a boolean if a field has been set.

### GetAsmswitchcapable

`func (o *BackupRest) GetAsmswitchcapable() bool`

GetAsmswitchcapable returns the Asmswitchcapable field if non-nil, zero value otherwise.

### GetAsmswitchcapableOk

`func (o *BackupRest) GetAsmswitchcapableOk() (*bool, bool)`

GetAsmswitchcapableOk returns a tuple with the Asmswitchcapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmswitchcapable

`func (o *BackupRest) SetAsmswitchcapable(v bool)`

SetAsmswitchcapable sets Asmswitchcapable field to given value.

### HasAsmswitchcapable

`func (o *BackupRest) HasAsmswitchcapable() bool`

HasAsmswitchcapable returns a boolean if a field has been set.

### GetIncarnation

`func (o *BackupRest) GetIncarnation() string`

GetIncarnation returns the Incarnation field if non-nil, zero value otherwise.

### GetIncarnationOk

`func (o *BackupRest) GetIncarnationOk() (*string, bool)`

GetIncarnationOk returns a tuple with the Incarnation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncarnation

`func (o *BackupRest) SetIncarnation(v string)`

SetIncarnation sets Incarnation field to given value.

### HasIncarnation

`func (o *BackupRest) HasIncarnation() bool`

HasIncarnation returns a boolean if a field has been set.

### GetHananodelist

`func (o *BackupRest) GetHananodelist() []string`

GetHananodelist returns the Hananodelist field if non-nil, zero value otherwise.

### GetHananodelistOk

`func (o *BackupRest) GetHananodelistOk() (*[]string, bool)`

GetHananodelistOk returns a tuple with the Hananodelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHananodelist

`func (o *BackupRest) SetHananodelist(v []string)`

SetHananodelist sets Hananodelist field to given value.

### HasHananodelist

`func (o *BackupRest) HasHananodelist() bool`

HasHananodelist returns a boolean if a field has been set.

### GetApplicationAllocatedSize

`func (o *BackupRest) GetApplicationAllocatedSize() int64`

GetApplicationAllocatedSize returns the ApplicationAllocatedSize field if non-nil, zero value otherwise.

### GetApplicationAllocatedSizeOk

`func (o *BackupRest) GetApplicationAllocatedSizeOk() (*int64, bool)`

GetApplicationAllocatedSizeOk returns a tuple with the ApplicationAllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationAllocatedSize

`func (o *BackupRest) SetApplicationAllocatedSize(v int64)`

SetApplicationAllocatedSize sets ApplicationAllocatedSize field to given value.

### HasApplicationAllocatedSize

`func (o *BackupRest) HasApplicationAllocatedSize() bool`

HasApplicationAllocatedSize returns a boolean if a field has been set.

### GetExportvolume

`func (o *BackupRest) GetExportvolume() []ExportVolumeInfoRest`

GetExportvolume returns the Exportvolume field if non-nil, zero value otherwise.

### GetExportvolumeOk

`func (o *BackupRest) GetExportvolumeOk() (*[]ExportVolumeInfoRest, bool)`

GetExportvolumeOk returns a tuple with the Exportvolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportvolume

`func (o *BackupRest) SetExportvolume(v []ExportVolumeInfoRest)`

SetExportvolume sets Exportvolume field to given value.

### HasExportvolume

`func (o *BackupRest) HasExportvolume() bool`

HasExportvolume returns a boolean if a field has been set.

### GetId

`func (o *BackupRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BackupRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *BackupRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BackupRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BackupRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BackupRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *BackupRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *BackupRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *BackupRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *BackupRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *BackupRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *BackupRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *BackupRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *BackupRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



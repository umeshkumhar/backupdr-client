# ApplicationRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Immutable** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Sources** | Pointer to [**[]ApplicationRest**](ApplicationRest.md) | List of entries representing the consolidated application on particular appliances. | [optional] 
**Name** | Pointer to **string** | Application&#39;s name. It may not be the same as appname. | [optional] 
**Host** | Pointer to [**HostRest**](HostRest.md) |  | [optional] 
**Srcid** | Pointer to **string** | Application&#39;s id on the appliance. | [optional] 
**Uniquename** | Pointer to **string** |  | [optional] 
**Appname** | Pointer to **string** | Application name. | [optional] 
**Isvm** | Pointer to **bool** | If the application is a VM | [optional] 
**Managed** | Pointer to **bool** | If the application is protected. | [optional] 
**Scheduleoff** | Pointer to **bool** | If the backup schedule is turned off | [optional] 
**Apptype** | Pointer to **string** | Application type. | [optional] 
**Originalappid** | Pointer to **string** | If this application is a shadow, the value of this attribute will be the id of the original application on original appliance. | [optional] 
**Pathname** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Backup** | Pointer to [**[]BackupRest**](BackupRest.md) | List of backups, if available. | [optional] 
**Isorphan** | Pointer to **bool** | Indicates if the application is orphan, which still has backups but the original application was unproteced and removed on purpose. | [optional] 
**Appclass** | Pointer to **string** | Application Class. It&#39;s similar to application type. It is useful in app-aware mount | [optional] 
**Sla** | Pointer to [**SlaRest**](SlaRest.md) |  | [optional] 
**Cluster** | Pointer to [**ClusterRest**](ClusterRest.md) |  | [optional] 
**Friendlypath** | Pointer to **string** |  | [optional] 
**Sourcecluster** | Pointer to **string** | Application&#39;s original appliance&#39;s clusterid. | [optional] 
**Friendlytype** | Pointer to **string** |  | [optional] 
**Volumes** | Pointer to **[]string** |  | [optional] 
**Protectable** | Pointer to **string** | If the application can be protected. | [optional] 
**Failoverstate** | Pointer to **string** |  | [optional] 
**Auxinfo** | Pointer to **string** |  | [optional] 
**Appversion** | Pointer to **string** | Application version. | [optional] 
**Networkname** | Pointer to **string** | Network name. | [optional] 
**Networkip** | Pointer to **string** | Network IP address. | [optional] 
**Ignore** | Pointer to **bool** | If the application is marked as to be ignored. | [optional] 
**Isclustered** | Pointer to **bool** | If the application is clustered. | [optional] 
**Frommount** | Pointer to **bool** | Indicate if this application is a child-app from app-aware mount. | [optional] 
**Sensitivity** | Pointer to **int32** | Sensitivity level. | [optional] 
**Mountedhosts** | Pointer to [**[]HostRest**](HostRest.md) |  | [optional] 
**AvailableSlp** | Pointer to [**[]SlpRest**](SlpRest.md) | Available SLP (profiles) that this application can potentially use for protection. | [optional] 
**Orglist** | Pointer to [**[]OrganizationRest**](OrganizationRest.md) | List of organizations that the application belongs to. | [optional] 
**Isrestoring** | Pointer to **bool** | If the application is being restore replaced. | [optional] 
**Consistencygroup** | Pointer to [**ApplicationRest**](ApplicationRest.md) |  | [optional] 
**Logicalgroup** | Pointer to [**LogicalGroupRest**](LogicalGroupRest.md) |  | [optional] 
**AppstateText** | Pointer to **[]string** | Application&#39;s states. | [optional] 
**Diskpools** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewApplicationRest

`func NewApplicationRest() *ApplicationRest`

NewApplicationRest instantiates a new ApplicationRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRestWithDefaults

`func NewApplicationRestWithDefaults() *ApplicationRest`

NewApplicationRestWithDefaults instantiates a new ApplicationRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImmutable

`func (o *ApplicationRest) GetImmutable() bool`

GetImmutable returns the Immutable field if non-nil, zero value otherwise.

### GetImmutableOk

`func (o *ApplicationRest) GetImmutableOk() (*bool, bool)`

GetImmutableOk returns a tuple with the Immutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutable

`func (o *ApplicationRest) SetImmutable(v bool)`

SetImmutable sets Immutable field to given value.

### HasImmutable

`func (o *ApplicationRest) HasImmutable() bool`

HasImmutable returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationRest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationRest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationRest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationRest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSources

`func (o *ApplicationRest) GetSources() []ApplicationRest`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *ApplicationRest) GetSourcesOk() (*[]ApplicationRest, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *ApplicationRest) SetSources(v []ApplicationRest)`

SetSources sets Sources field to given value.

### HasSources

`func (o *ApplicationRest) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetName

`func (o *ApplicationRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHost

`func (o *ApplicationRest) GetHost() HostRest`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ApplicationRest) GetHostOk() (*HostRest, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ApplicationRest) SetHost(v HostRest)`

SetHost sets Host field to given value.

### HasHost

`func (o *ApplicationRest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetSrcid

`func (o *ApplicationRest) GetSrcid() string`

GetSrcid returns the Srcid field if non-nil, zero value otherwise.

### GetSrcidOk

`func (o *ApplicationRest) GetSrcidOk() (*string, bool)`

GetSrcidOk returns a tuple with the Srcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcid

`func (o *ApplicationRest) SetSrcid(v string)`

SetSrcid sets Srcid field to given value.

### HasSrcid

`func (o *ApplicationRest) HasSrcid() bool`

HasSrcid returns a boolean if a field has been set.

### GetUniquename

`func (o *ApplicationRest) GetUniquename() string`

GetUniquename returns the Uniquename field if non-nil, zero value otherwise.

### GetUniquenameOk

`func (o *ApplicationRest) GetUniquenameOk() (*string, bool)`

GetUniquenameOk returns a tuple with the Uniquename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniquename

`func (o *ApplicationRest) SetUniquename(v string)`

SetUniquename sets Uniquename field to given value.

### HasUniquename

`func (o *ApplicationRest) HasUniquename() bool`

HasUniquename returns a boolean if a field has been set.

### GetAppname

`func (o *ApplicationRest) GetAppname() string`

GetAppname returns the Appname field if non-nil, zero value otherwise.

### GetAppnameOk

`func (o *ApplicationRest) GetAppnameOk() (*string, bool)`

GetAppnameOk returns a tuple with the Appname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppname

`func (o *ApplicationRest) SetAppname(v string)`

SetAppname sets Appname field to given value.

### HasAppname

`func (o *ApplicationRest) HasAppname() bool`

HasAppname returns a boolean if a field has been set.

### GetIsvm

`func (o *ApplicationRest) GetIsvm() bool`

GetIsvm returns the Isvm field if non-nil, zero value otherwise.

### GetIsvmOk

`func (o *ApplicationRest) GetIsvmOk() (*bool, bool)`

GetIsvmOk returns a tuple with the Isvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsvm

`func (o *ApplicationRest) SetIsvm(v bool)`

SetIsvm sets Isvm field to given value.

### HasIsvm

`func (o *ApplicationRest) HasIsvm() bool`

HasIsvm returns a boolean if a field has been set.

### GetManaged

`func (o *ApplicationRest) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ApplicationRest) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ApplicationRest) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ApplicationRest) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetScheduleoff

`func (o *ApplicationRest) GetScheduleoff() bool`

GetScheduleoff returns the Scheduleoff field if non-nil, zero value otherwise.

### GetScheduleoffOk

`func (o *ApplicationRest) GetScheduleoffOk() (*bool, bool)`

GetScheduleoffOk returns a tuple with the Scheduleoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleoff

`func (o *ApplicationRest) SetScheduleoff(v bool)`

SetScheduleoff sets Scheduleoff field to given value.

### HasScheduleoff

`func (o *ApplicationRest) HasScheduleoff() bool`

HasScheduleoff returns a boolean if a field has been set.

### GetApptype

`func (o *ApplicationRest) GetApptype() string`

GetApptype returns the Apptype field if non-nil, zero value otherwise.

### GetApptypeOk

`func (o *ApplicationRest) GetApptypeOk() (*string, bool)`

GetApptypeOk returns a tuple with the Apptype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApptype

`func (o *ApplicationRest) SetApptype(v string)`

SetApptype sets Apptype field to given value.

### HasApptype

`func (o *ApplicationRest) HasApptype() bool`

HasApptype returns a boolean if a field has been set.

### GetOriginalappid

`func (o *ApplicationRest) GetOriginalappid() string`

GetOriginalappid returns the Originalappid field if non-nil, zero value otherwise.

### GetOriginalappidOk

`func (o *ApplicationRest) GetOriginalappidOk() (*string, bool)`

GetOriginalappidOk returns a tuple with the Originalappid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalappid

`func (o *ApplicationRest) SetOriginalappid(v string)`

SetOriginalappid sets Originalappid field to given value.

### HasOriginalappid

`func (o *ApplicationRest) HasOriginalappid() bool`

HasOriginalappid returns a boolean if a field has been set.

### GetPathname

`func (o *ApplicationRest) GetPathname() string`

GetPathname returns the Pathname field if non-nil, zero value otherwise.

### GetPathnameOk

`func (o *ApplicationRest) GetPathnameOk() (*string, bool)`

GetPathnameOk returns a tuple with the Pathname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathname

`func (o *ApplicationRest) SetPathname(v string)`

SetPathname sets Pathname field to given value.

### HasPathname

`func (o *ApplicationRest) HasPathname() bool`

HasPathname returns a boolean if a field has been set.

### GetUsername

`func (o *ApplicationRest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApplicationRest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApplicationRest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApplicationRest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetBackup

`func (o *ApplicationRest) GetBackup() []BackupRest`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *ApplicationRest) GetBackupOk() (*[]BackupRest, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *ApplicationRest) SetBackup(v []BackupRest)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *ApplicationRest) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetIsorphan

`func (o *ApplicationRest) GetIsorphan() bool`

GetIsorphan returns the Isorphan field if non-nil, zero value otherwise.

### GetIsorphanOk

`func (o *ApplicationRest) GetIsorphanOk() (*bool, bool)`

GetIsorphanOk returns a tuple with the Isorphan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsorphan

`func (o *ApplicationRest) SetIsorphan(v bool)`

SetIsorphan sets Isorphan field to given value.

### HasIsorphan

`func (o *ApplicationRest) HasIsorphan() bool`

HasIsorphan returns a boolean if a field has been set.

### GetAppclass

`func (o *ApplicationRest) GetAppclass() string`

GetAppclass returns the Appclass field if non-nil, zero value otherwise.

### GetAppclassOk

`func (o *ApplicationRest) GetAppclassOk() (*string, bool)`

GetAppclassOk returns a tuple with the Appclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppclass

`func (o *ApplicationRest) SetAppclass(v string)`

SetAppclass sets Appclass field to given value.

### HasAppclass

`func (o *ApplicationRest) HasAppclass() bool`

HasAppclass returns a boolean if a field has been set.

### GetSla

`func (o *ApplicationRest) GetSla() SlaRest`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *ApplicationRest) GetSlaOk() (*SlaRest, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *ApplicationRest) SetSla(v SlaRest)`

SetSla sets Sla field to given value.

### HasSla

`func (o *ApplicationRest) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetCluster

`func (o *ApplicationRest) GetCluster() ClusterRest`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ApplicationRest) GetClusterOk() (*ClusterRest, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ApplicationRest) SetCluster(v ClusterRest)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ApplicationRest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetFriendlypath

`func (o *ApplicationRest) GetFriendlypath() string`

GetFriendlypath returns the Friendlypath field if non-nil, zero value otherwise.

### GetFriendlypathOk

`func (o *ApplicationRest) GetFriendlypathOk() (*string, bool)`

GetFriendlypathOk returns a tuple with the Friendlypath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlypath

`func (o *ApplicationRest) SetFriendlypath(v string)`

SetFriendlypath sets Friendlypath field to given value.

### HasFriendlypath

`func (o *ApplicationRest) HasFriendlypath() bool`

HasFriendlypath returns a boolean if a field has been set.

### GetSourcecluster

`func (o *ApplicationRest) GetSourcecluster() string`

GetSourcecluster returns the Sourcecluster field if non-nil, zero value otherwise.

### GetSourceclusterOk

`func (o *ApplicationRest) GetSourceclusterOk() (*string, bool)`

GetSourceclusterOk returns a tuple with the Sourcecluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcecluster

`func (o *ApplicationRest) SetSourcecluster(v string)`

SetSourcecluster sets Sourcecluster field to given value.

### HasSourcecluster

`func (o *ApplicationRest) HasSourcecluster() bool`

HasSourcecluster returns a boolean if a field has been set.

### GetFriendlytype

`func (o *ApplicationRest) GetFriendlytype() string`

GetFriendlytype returns the Friendlytype field if non-nil, zero value otherwise.

### GetFriendlytypeOk

`func (o *ApplicationRest) GetFriendlytypeOk() (*string, bool)`

GetFriendlytypeOk returns a tuple with the Friendlytype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlytype

`func (o *ApplicationRest) SetFriendlytype(v string)`

SetFriendlytype sets Friendlytype field to given value.

### HasFriendlytype

`func (o *ApplicationRest) HasFriendlytype() bool`

HasFriendlytype returns a boolean if a field has been set.

### GetVolumes

`func (o *ApplicationRest) GetVolumes() []string`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ApplicationRest) GetVolumesOk() (*[]string, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ApplicationRest) SetVolumes(v []string)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ApplicationRest) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetProtectable

`func (o *ApplicationRest) GetProtectable() string`

GetProtectable returns the Protectable field if non-nil, zero value otherwise.

### GetProtectableOk

`func (o *ApplicationRest) GetProtectableOk() (*string, bool)`

GetProtectableOk returns a tuple with the Protectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectable

`func (o *ApplicationRest) SetProtectable(v string)`

SetProtectable sets Protectable field to given value.

### HasProtectable

`func (o *ApplicationRest) HasProtectable() bool`

HasProtectable returns a boolean if a field has been set.

### GetFailoverstate

`func (o *ApplicationRest) GetFailoverstate() string`

GetFailoverstate returns the Failoverstate field if non-nil, zero value otherwise.

### GetFailoverstateOk

`func (o *ApplicationRest) GetFailoverstateOk() (*string, bool)`

GetFailoverstateOk returns a tuple with the Failoverstate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverstate

`func (o *ApplicationRest) SetFailoverstate(v string)`

SetFailoverstate sets Failoverstate field to given value.

### HasFailoverstate

`func (o *ApplicationRest) HasFailoverstate() bool`

HasFailoverstate returns a boolean if a field has been set.

### GetAuxinfo

`func (o *ApplicationRest) GetAuxinfo() string`

GetAuxinfo returns the Auxinfo field if non-nil, zero value otherwise.

### GetAuxinfoOk

`func (o *ApplicationRest) GetAuxinfoOk() (*string, bool)`

GetAuxinfoOk returns a tuple with the Auxinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxinfo

`func (o *ApplicationRest) SetAuxinfo(v string)`

SetAuxinfo sets Auxinfo field to given value.

### HasAuxinfo

`func (o *ApplicationRest) HasAuxinfo() bool`

HasAuxinfo returns a boolean if a field has been set.

### GetAppversion

`func (o *ApplicationRest) GetAppversion() string`

GetAppversion returns the Appversion field if non-nil, zero value otherwise.

### GetAppversionOk

`func (o *ApplicationRest) GetAppversionOk() (*string, bool)`

GetAppversionOk returns a tuple with the Appversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppversion

`func (o *ApplicationRest) SetAppversion(v string)`

SetAppversion sets Appversion field to given value.

### HasAppversion

`func (o *ApplicationRest) HasAppversion() bool`

HasAppversion returns a boolean if a field has been set.

### GetNetworkname

`func (o *ApplicationRest) GetNetworkname() string`

GetNetworkname returns the Networkname field if non-nil, zero value otherwise.

### GetNetworknameOk

`func (o *ApplicationRest) GetNetworknameOk() (*string, bool)`

GetNetworknameOk returns a tuple with the Networkname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkname

`func (o *ApplicationRest) SetNetworkname(v string)`

SetNetworkname sets Networkname field to given value.

### HasNetworkname

`func (o *ApplicationRest) HasNetworkname() bool`

HasNetworkname returns a boolean if a field has been set.

### GetNetworkip

`func (o *ApplicationRest) GetNetworkip() string`

GetNetworkip returns the Networkip field if non-nil, zero value otherwise.

### GetNetworkipOk

`func (o *ApplicationRest) GetNetworkipOk() (*string, bool)`

GetNetworkipOk returns a tuple with the Networkip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkip

`func (o *ApplicationRest) SetNetworkip(v string)`

SetNetworkip sets Networkip field to given value.

### HasNetworkip

`func (o *ApplicationRest) HasNetworkip() bool`

HasNetworkip returns a boolean if a field has been set.

### GetIgnore

`func (o *ApplicationRest) GetIgnore() bool`

GetIgnore returns the Ignore field if non-nil, zero value otherwise.

### GetIgnoreOk

`func (o *ApplicationRest) GetIgnoreOk() (*bool, bool)`

GetIgnoreOk returns a tuple with the Ignore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnore

`func (o *ApplicationRest) SetIgnore(v bool)`

SetIgnore sets Ignore field to given value.

### HasIgnore

`func (o *ApplicationRest) HasIgnore() bool`

HasIgnore returns a boolean if a field has been set.

### GetIsclustered

`func (o *ApplicationRest) GetIsclustered() bool`

GetIsclustered returns the Isclustered field if non-nil, zero value otherwise.

### GetIsclusteredOk

`func (o *ApplicationRest) GetIsclusteredOk() (*bool, bool)`

GetIsclusteredOk returns a tuple with the Isclustered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsclustered

`func (o *ApplicationRest) SetIsclustered(v bool)`

SetIsclustered sets Isclustered field to given value.

### HasIsclustered

`func (o *ApplicationRest) HasIsclustered() bool`

HasIsclustered returns a boolean if a field has been set.

### GetFrommount

`func (o *ApplicationRest) GetFrommount() bool`

GetFrommount returns the Frommount field if non-nil, zero value otherwise.

### GetFrommountOk

`func (o *ApplicationRest) GetFrommountOk() (*bool, bool)`

GetFrommountOk returns a tuple with the Frommount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrommount

`func (o *ApplicationRest) SetFrommount(v bool)`

SetFrommount sets Frommount field to given value.

### HasFrommount

`func (o *ApplicationRest) HasFrommount() bool`

HasFrommount returns a boolean if a field has been set.

### GetSensitivity

`func (o *ApplicationRest) GetSensitivity() int32`

GetSensitivity returns the Sensitivity field if non-nil, zero value otherwise.

### GetSensitivityOk

`func (o *ApplicationRest) GetSensitivityOk() (*int32, bool)`

GetSensitivityOk returns a tuple with the Sensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitivity

`func (o *ApplicationRest) SetSensitivity(v int32)`

SetSensitivity sets Sensitivity field to given value.

### HasSensitivity

`func (o *ApplicationRest) HasSensitivity() bool`

HasSensitivity returns a boolean if a field has been set.

### GetMountedhosts

`func (o *ApplicationRest) GetMountedhosts() []HostRest`

GetMountedhosts returns the Mountedhosts field if non-nil, zero value otherwise.

### GetMountedhostsOk

`func (o *ApplicationRest) GetMountedhostsOk() (*[]HostRest, bool)`

GetMountedhostsOk returns a tuple with the Mountedhosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountedhosts

`func (o *ApplicationRest) SetMountedhosts(v []HostRest)`

SetMountedhosts sets Mountedhosts field to given value.

### HasMountedhosts

`func (o *ApplicationRest) HasMountedhosts() bool`

HasMountedhosts returns a boolean if a field has been set.

### GetAvailableSlp

`func (o *ApplicationRest) GetAvailableSlp() []SlpRest`

GetAvailableSlp returns the AvailableSlp field if non-nil, zero value otherwise.

### GetAvailableSlpOk

`func (o *ApplicationRest) GetAvailableSlpOk() (*[]SlpRest, bool)`

GetAvailableSlpOk returns a tuple with the AvailableSlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableSlp

`func (o *ApplicationRest) SetAvailableSlp(v []SlpRest)`

SetAvailableSlp sets AvailableSlp field to given value.

### HasAvailableSlp

`func (o *ApplicationRest) HasAvailableSlp() bool`

HasAvailableSlp returns a boolean if a field has been set.

### GetOrglist

`func (o *ApplicationRest) GetOrglist() []OrganizationRest`

GetOrglist returns the Orglist field if non-nil, zero value otherwise.

### GetOrglistOk

`func (o *ApplicationRest) GetOrglistOk() (*[]OrganizationRest, bool)`

GetOrglistOk returns a tuple with the Orglist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrglist

`func (o *ApplicationRest) SetOrglist(v []OrganizationRest)`

SetOrglist sets Orglist field to given value.

### HasOrglist

`func (o *ApplicationRest) HasOrglist() bool`

HasOrglist returns a boolean if a field has been set.

### GetIsrestoring

`func (o *ApplicationRest) GetIsrestoring() bool`

GetIsrestoring returns the Isrestoring field if non-nil, zero value otherwise.

### GetIsrestoringOk

`func (o *ApplicationRest) GetIsrestoringOk() (*bool, bool)`

GetIsrestoringOk returns a tuple with the Isrestoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsrestoring

`func (o *ApplicationRest) SetIsrestoring(v bool)`

SetIsrestoring sets Isrestoring field to given value.

### HasIsrestoring

`func (o *ApplicationRest) HasIsrestoring() bool`

HasIsrestoring returns a boolean if a field has been set.

### GetConsistencygroup

`func (o *ApplicationRest) GetConsistencygroup() ApplicationRest`

GetConsistencygroup returns the Consistencygroup field if non-nil, zero value otherwise.

### GetConsistencygroupOk

`func (o *ApplicationRest) GetConsistencygroupOk() (*ApplicationRest, bool)`

GetConsistencygroupOk returns a tuple with the Consistencygroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistencygroup

`func (o *ApplicationRest) SetConsistencygroup(v ApplicationRest)`

SetConsistencygroup sets Consistencygroup field to given value.

### HasConsistencygroup

`func (o *ApplicationRest) HasConsistencygroup() bool`

HasConsistencygroup returns a boolean if a field has been set.

### GetLogicalgroup

`func (o *ApplicationRest) GetLogicalgroup() LogicalGroupRest`

GetLogicalgroup returns the Logicalgroup field if non-nil, zero value otherwise.

### GetLogicalgroupOk

`func (o *ApplicationRest) GetLogicalgroupOk() (*LogicalGroupRest, bool)`

GetLogicalgroupOk returns a tuple with the Logicalgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalgroup

`func (o *ApplicationRest) SetLogicalgroup(v LogicalGroupRest)`

SetLogicalgroup sets Logicalgroup field to given value.

### HasLogicalgroup

`func (o *ApplicationRest) HasLogicalgroup() bool`

HasLogicalgroup returns a boolean if a field has been set.

### GetAppstateText

`func (o *ApplicationRest) GetAppstateText() []string`

GetAppstateText returns the AppstateText field if non-nil, zero value otherwise.

### GetAppstateTextOk

`func (o *ApplicationRest) GetAppstateTextOk() (*[]string, bool)`

GetAppstateTextOk returns a tuple with the AppstateText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppstateText

`func (o *ApplicationRest) SetAppstateText(v []string)`

SetAppstateText sets AppstateText field to given value.

### HasAppstateText

`func (o *ApplicationRest) HasAppstateText() bool`

HasAppstateText returns a boolean if a field has been set.

### GetDiskpools

`func (o *ApplicationRest) GetDiskpools() []string`

GetDiskpools returns the Diskpools field if non-nil, zero value otherwise.

### GetDiskpoolsOk

`func (o *ApplicationRest) GetDiskpoolsOk() (*[]string, bool)`

GetDiskpoolsOk returns a tuple with the Diskpools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskpools

`func (o *ApplicationRest) SetDiskpools(v []string)`

SetDiskpools sets Diskpools field to given value.

### HasDiskpools

`func (o *ApplicationRest) HasDiskpools() bool`

HasDiskpools returns a boolean if a field has been set.

### GetId

`func (o *ApplicationRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ApplicationRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ApplicationRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ApplicationRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ApplicationRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ApplicationRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ApplicationRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ApplicationRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ApplicationRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ApplicationRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ApplicationRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ApplicationRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ApplicationRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



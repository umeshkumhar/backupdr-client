# JobRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pid** | Pointer to **int32** |  | [optional] 
**Task** | Pointer to **map[string]interface{}** |  | [optional] 
**Flags** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**Expiration** | Pointer to **int64** |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Errorcode** | Pointer to **string** |  | [optional] 
**Jobclass** | Pointer to **string** |  | [optional] 
**Appid** | Pointer to **string** |  | [optional] 
**Hostid** | Pointer to **int64** |  | [optional] 
**Srcid** | Pointer to **string** |  | [optional] 
**Clusterid** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Appname** | Pointer to **string** |  | [optional] 
**Apptype** | Pointer to **string** |  | [optional] 
**Backup** | Pointer to [**BackupRest**](BackupRest.md) |  | [optional] 
**Sourceuds** | Pointer to **int64** |  | [optional] 
**Originatinguds** | Pointer to **int64** |  | [optional] 
**Targetuds** | Pointer to **int64** |  | [optional] 
**Parentid** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Jobname** | Pointer to **string** |  | [optional] 
**Transport** | Pointer to **string** |  | [optional] 
**Startdate** | Pointer to **int64** |  | [optional] 
**Enddate** | Pointer to **int64** |  | [optional] 
**Immutabilitydate** | Pointer to **int64** |  | [optional] 
**Consistencydate** | Pointer to **int64** |  | [optional] 
**Virtualsize** | Pointer to **int64** |  | [optional] 
**Sltname** | Pointer to **string** |  | [optional] 
**Policyname** | Pointer to **string** |  | [optional] 
**Isexpired** | Pointer to **bool** |  | [optional] 
**Targethost** | Pointer to **string** |  | [optional] 
**Sourceid** | Pointer to **string** |  | [optional] 
**Jobcount** | Pointer to **int32** |  | [optional] 
**Changerequest** | Pointer to **int32** |  | [optional] 
**Progress** | Pointer to **int64** |  | [optional] 
**Relativesize** | Pointer to **int64** |  | [optional] 
**Retrycount** | Pointer to **int32** |  | [optional] 
**Queuedate** | Pointer to **int64** |  | [optional] 
**Expirationdate** | Pointer to **int64** |  | [optional] 
**Isscheduled** | Pointer to **bool** |  | [optional] 
**Jobtag** | Pointer to **string** |  | [optional] 
**Event** | Pointer to [**EventRest**](EventRest.md) |  | [optional] 
**Id2** | Pointer to **string** |  | [optional] 
**Changerequesttext** | Pointer to **string** |  | [optional] 
**Logsmarttype** | Pointer to **string** |  | [optional] 
**Extrainfo** | Pointer to **string** |  | [optional] 
**Diskpools** | Pointer to [**[]DiskPoolRest**](DiskPoolRest.md) |  | [optional] 
**Targetdiskpool** | Pointer to [**DiskPoolRest**](DiskPoolRest.md) |  | [optional] 
**Jobclasscode** | Pointer to **int32** |  | [optional] 
**Sourcediskpool** | Pointer to [**DiskPoolRest**](DiskPoolRest.md) |  | [optional] 
**Yaml** | Pointer to **string** |  | [optional] 
**Appliance** | Pointer to [**ClusterRest**](ClusterRest.md) |  | [optional] 
**ConsistencyMode** | Pointer to **string** |  | [optional] 
**FlagsText** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewJobRest

`func NewJobRest() *JobRest`

NewJobRest instantiates a new JobRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobRestWithDefaults

`func NewJobRestWithDefaults() *JobRest`

NewJobRestWithDefaults instantiates a new JobRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPid

`func (o *JobRest) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *JobRest) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *JobRest) SetPid(v int32)`

SetPid sets Pid field to given value.

### HasPid

`func (o *JobRest) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetTask

`func (o *JobRest) GetTask() map[string]interface{}`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *JobRest) GetTaskOk() (*map[string]interface{}, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *JobRest) SetTask(v map[string]interface{})`

SetTask sets Task field to given value.

### HasTask

`func (o *JobRest) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetFlags

`func (o *JobRest) GetFlags() int64`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *JobRest) GetFlagsOk() (*int64, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *JobRest) SetFlags(v int64)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *JobRest) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetDescription

`func (o *JobRest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JobRest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JobRest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JobRest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMessage

`func (o *JobRest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *JobRest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *JobRest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *JobRest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPriority

`func (o *JobRest) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *JobRest) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *JobRest) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *JobRest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetExpiration

`func (o *JobRest) GetExpiration() int64`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *JobRest) GetExpirationOk() (*int64, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *JobRest) SetExpiration(v int64)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *JobRest) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetDuration

`func (o *JobRest) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *JobRest) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *JobRest) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *JobRest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetLabel

`func (o *JobRest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *JobRest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *JobRest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *JobRest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetErrorcode

`func (o *JobRest) GetErrorcode() string`

GetErrorcode returns the Errorcode field if non-nil, zero value otherwise.

### GetErrorcodeOk

`func (o *JobRest) GetErrorcodeOk() (*string, bool)`

GetErrorcodeOk returns a tuple with the Errorcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorcode

`func (o *JobRest) SetErrorcode(v string)`

SetErrorcode sets Errorcode field to given value.

### HasErrorcode

`func (o *JobRest) HasErrorcode() bool`

HasErrorcode returns a boolean if a field has been set.

### GetJobclass

`func (o *JobRest) GetJobclass() string`

GetJobclass returns the Jobclass field if non-nil, zero value otherwise.

### GetJobclassOk

`func (o *JobRest) GetJobclassOk() (*string, bool)`

GetJobclassOk returns a tuple with the Jobclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobclass

`func (o *JobRest) SetJobclass(v string)`

SetJobclass sets Jobclass field to given value.

### HasJobclass

`func (o *JobRest) HasJobclass() bool`

HasJobclass returns a boolean if a field has been set.

### GetAppid

`func (o *JobRest) GetAppid() string`

GetAppid returns the Appid field if non-nil, zero value otherwise.

### GetAppidOk

`func (o *JobRest) GetAppidOk() (*string, bool)`

GetAppidOk returns a tuple with the Appid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppid

`func (o *JobRest) SetAppid(v string)`

SetAppid sets Appid field to given value.

### HasAppid

`func (o *JobRest) HasAppid() bool`

HasAppid returns a boolean if a field has been set.

### GetHostid

`func (o *JobRest) GetHostid() int64`

GetHostid returns the Hostid field if non-nil, zero value otherwise.

### GetHostidOk

`func (o *JobRest) GetHostidOk() (*int64, bool)`

GetHostidOk returns a tuple with the Hostid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostid

`func (o *JobRest) SetHostid(v int64)`

SetHostid sets Hostid field to given value.

### HasHostid

`func (o *JobRest) HasHostid() bool`

HasHostid returns a boolean if a field has been set.

### GetSrcid

`func (o *JobRest) GetSrcid() string`

GetSrcid returns the Srcid field if non-nil, zero value otherwise.

### GetSrcidOk

`func (o *JobRest) GetSrcidOk() (*string, bool)`

GetSrcidOk returns a tuple with the Srcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcid

`func (o *JobRest) SetSrcid(v string)`

SetSrcid sets Srcid field to given value.

### HasSrcid

`func (o *JobRest) HasSrcid() bool`

HasSrcid returns a boolean if a field has been set.

### GetClusterid

`func (o *JobRest) GetClusterid() string`

GetClusterid returns the Clusterid field if non-nil, zero value otherwise.

### GetClusteridOk

`func (o *JobRest) GetClusteridOk() (*string, bool)`

GetClusteridOk returns a tuple with the Clusterid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterid

`func (o *JobRest) SetClusterid(v string)`

SetClusterid sets Clusterid field to given value.

### HasClusterid

`func (o *JobRest) HasClusterid() bool`

HasClusterid returns a boolean if a field has been set.

### GetHostname

`func (o *JobRest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *JobRest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *JobRest) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *JobRest) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetAppname

`func (o *JobRest) GetAppname() string`

GetAppname returns the Appname field if non-nil, zero value otherwise.

### GetAppnameOk

`func (o *JobRest) GetAppnameOk() (*string, bool)`

GetAppnameOk returns a tuple with the Appname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppname

`func (o *JobRest) SetAppname(v string)`

SetAppname sets Appname field to given value.

### HasAppname

`func (o *JobRest) HasAppname() bool`

HasAppname returns a boolean if a field has been set.

### GetApptype

`func (o *JobRest) GetApptype() string`

GetApptype returns the Apptype field if non-nil, zero value otherwise.

### GetApptypeOk

`func (o *JobRest) GetApptypeOk() (*string, bool)`

GetApptypeOk returns a tuple with the Apptype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApptype

`func (o *JobRest) SetApptype(v string)`

SetApptype sets Apptype field to given value.

### HasApptype

`func (o *JobRest) HasApptype() bool`

HasApptype returns a boolean if a field has been set.

### GetBackup

`func (o *JobRest) GetBackup() BackupRest`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *JobRest) GetBackupOk() (*BackupRest, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *JobRest) SetBackup(v BackupRest)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *JobRest) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetSourceuds

`func (o *JobRest) GetSourceuds() int64`

GetSourceuds returns the Sourceuds field if non-nil, zero value otherwise.

### GetSourceudsOk

`func (o *JobRest) GetSourceudsOk() (*int64, bool)`

GetSourceudsOk returns a tuple with the Sourceuds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceuds

`func (o *JobRest) SetSourceuds(v int64)`

SetSourceuds sets Sourceuds field to given value.

### HasSourceuds

`func (o *JobRest) HasSourceuds() bool`

HasSourceuds returns a boolean if a field has been set.

### GetOriginatinguds

`func (o *JobRest) GetOriginatinguds() int64`

GetOriginatinguds returns the Originatinguds field if non-nil, zero value otherwise.

### GetOriginatingudsOk

`func (o *JobRest) GetOriginatingudsOk() (*int64, bool)`

GetOriginatingudsOk returns a tuple with the Originatinguds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatinguds

`func (o *JobRest) SetOriginatinguds(v int64)`

SetOriginatinguds sets Originatinguds field to given value.

### HasOriginatinguds

`func (o *JobRest) HasOriginatinguds() bool`

HasOriginatinguds returns a boolean if a field has been set.

### GetTargetuds

`func (o *JobRest) GetTargetuds() int64`

GetTargetuds returns the Targetuds field if non-nil, zero value otherwise.

### GetTargetudsOk

`func (o *JobRest) GetTargetudsOk() (*int64, bool)`

GetTargetudsOk returns a tuple with the Targetuds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetuds

`func (o *JobRest) SetTargetuds(v int64)`

SetTargetuds sets Targetuds field to given value.

### HasTargetuds

`func (o *JobRest) HasTargetuds() bool`

HasTargetuds returns a boolean if a field has been set.

### GetParentid

`func (o *JobRest) GetParentid() string`

GetParentid returns the Parentid field if non-nil, zero value otherwise.

### GetParentidOk

`func (o *JobRest) GetParentidOk() (*string, bool)`

GetParentidOk returns a tuple with the Parentid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentid

`func (o *JobRest) SetParentid(v string)`

SetParentid sets Parentid field to given value.

### HasParentid

`func (o *JobRest) HasParentid() bool`

HasParentid returns a boolean if a field has been set.

### GetStatus

`func (o *JobRest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobRest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobRest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JobRest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetJobname

`func (o *JobRest) GetJobname() string`

GetJobname returns the Jobname field if non-nil, zero value otherwise.

### GetJobnameOk

`func (o *JobRest) GetJobnameOk() (*string, bool)`

GetJobnameOk returns a tuple with the Jobname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobname

`func (o *JobRest) SetJobname(v string)`

SetJobname sets Jobname field to given value.

### HasJobname

`func (o *JobRest) HasJobname() bool`

HasJobname returns a boolean if a field has been set.

### GetTransport

`func (o *JobRest) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *JobRest) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *JobRest) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *JobRest) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetStartdate

`func (o *JobRest) GetStartdate() int64`

GetStartdate returns the Startdate field if non-nil, zero value otherwise.

### GetStartdateOk

`func (o *JobRest) GetStartdateOk() (*int64, bool)`

GetStartdateOk returns a tuple with the Startdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartdate

`func (o *JobRest) SetStartdate(v int64)`

SetStartdate sets Startdate field to given value.

### HasStartdate

`func (o *JobRest) HasStartdate() bool`

HasStartdate returns a boolean if a field has been set.

### GetEnddate

`func (o *JobRest) GetEnddate() int64`

GetEnddate returns the Enddate field if non-nil, zero value otherwise.

### GetEnddateOk

`func (o *JobRest) GetEnddateOk() (*int64, bool)`

GetEnddateOk returns a tuple with the Enddate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnddate

`func (o *JobRest) SetEnddate(v int64)`

SetEnddate sets Enddate field to given value.

### HasEnddate

`func (o *JobRest) HasEnddate() bool`

HasEnddate returns a boolean if a field has been set.

### GetImmutabilitydate

`func (o *JobRest) GetImmutabilitydate() int64`

GetImmutabilitydate returns the Immutabilitydate field if non-nil, zero value otherwise.

### GetImmutabilitydateOk

`func (o *JobRest) GetImmutabilitydateOk() (*int64, bool)`

GetImmutabilitydateOk returns a tuple with the Immutabilitydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutabilitydate

`func (o *JobRest) SetImmutabilitydate(v int64)`

SetImmutabilitydate sets Immutabilitydate field to given value.

### HasImmutabilitydate

`func (o *JobRest) HasImmutabilitydate() bool`

HasImmutabilitydate returns a boolean if a field has been set.

### GetConsistencydate

`func (o *JobRest) GetConsistencydate() int64`

GetConsistencydate returns the Consistencydate field if non-nil, zero value otherwise.

### GetConsistencydateOk

`func (o *JobRest) GetConsistencydateOk() (*int64, bool)`

GetConsistencydateOk returns a tuple with the Consistencydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistencydate

`func (o *JobRest) SetConsistencydate(v int64)`

SetConsistencydate sets Consistencydate field to given value.

### HasConsistencydate

`func (o *JobRest) HasConsistencydate() bool`

HasConsistencydate returns a boolean if a field has been set.

### GetVirtualsize

`func (o *JobRest) GetVirtualsize() int64`

GetVirtualsize returns the Virtualsize field if non-nil, zero value otherwise.

### GetVirtualsizeOk

`func (o *JobRest) GetVirtualsizeOk() (*int64, bool)`

GetVirtualsizeOk returns a tuple with the Virtualsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualsize

`func (o *JobRest) SetVirtualsize(v int64)`

SetVirtualsize sets Virtualsize field to given value.

### HasVirtualsize

`func (o *JobRest) HasVirtualsize() bool`

HasVirtualsize returns a boolean if a field has been set.

### GetSltname

`func (o *JobRest) GetSltname() string`

GetSltname returns the Sltname field if non-nil, zero value otherwise.

### GetSltnameOk

`func (o *JobRest) GetSltnameOk() (*string, bool)`

GetSltnameOk returns a tuple with the Sltname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSltname

`func (o *JobRest) SetSltname(v string)`

SetSltname sets Sltname field to given value.

### HasSltname

`func (o *JobRest) HasSltname() bool`

HasSltname returns a boolean if a field has been set.

### GetPolicyname

`func (o *JobRest) GetPolicyname() string`

GetPolicyname returns the Policyname field if non-nil, zero value otherwise.

### GetPolicynameOk

`func (o *JobRest) GetPolicynameOk() (*string, bool)`

GetPolicynameOk returns a tuple with the Policyname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyname

`func (o *JobRest) SetPolicyname(v string)`

SetPolicyname sets Policyname field to given value.

### HasPolicyname

`func (o *JobRest) HasPolicyname() bool`

HasPolicyname returns a boolean if a field has been set.

### GetIsexpired

`func (o *JobRest) GetIsexpired() bool`

GetIsexpired returns the Isexpired field if non-nil, zero value otherwise.

### GetIsexpiredOk

`func (o *JobRest) GetIsexpiredOk() (*bool, bool)`

GetIsexpiredOk returns a tuple with the Isexpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsexpired

`func (o *JobRest) SetIsexpired(v bool)`

SetIsexpired sets Isexpired field to given value.

### HasIsexpired

`func (o *JobRest) HasIsexpired() bool`

HasIsexpired returns a boolean if a field has been set.

### GetTargethost

`func (o *JobRest) GetTargethost() string`

GetTargethost returns the Targethost field if non-nil, zero value otherwise.

### GetTargethostOk

`func (o *JobRest) GetTargethostOk() (*string, bool)`

GetTargethostOk returns a tuple with the Targethost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargethost

`func (o *JobRest) SetTargethost(v string)`

SetTargethost sets Targethost field to given value.

### HasTargethost

`func (o *JobRest) HasTargethost() bool`

HasTargethost returns a boolean if a field has been set.

### GetSourceid

`func (o *JobRest) GetSourceid() string`

GetSourceid returns the Sourceid field if non-nil, zero value otherwise.

### GetSourceidOk

`func (o *JobRest) GetSourceidOk() (*string, bool)`

GetSourceidOk returns a tuple with the Sourceid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceid

`func (o *JobRest) SetSourceid(v string)`

SetSourceid sets Sourceid field to given value.

### HasSourceid

`func (o *JobRest) HasSourceid() bool`

HasSourceid returns a boolean if a field has been set.

### GetJobcount

`func (o *JobRest) GetJobcount() int32`

GetJobcount returns the Jobcount field if non-nil, zero value otherwise.

### GetJobcountOk

`func (o *JobRest) GetJobcountOk() (*int32, bool)`

GetJobcountOk returns a tuple with the Jobcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobcount

`func (o *JobRest) SetJobcount(v int32)`

SetJobcount sets Jobcount field to given value.

### HasJobcount

`func (o *JobRest) HasJobcount() bool`

HasJobcount returns a boolean if a field has been set.

### GetChangerequest

`func (o *JobRest) GetChangerequest() int32`

GetChangerequest returns the Changerequest field if non-nil, zero value otherwise.

### GetChangerequestOk

`func (o *JobRest) GetChangerequestOk() (*int32, bool)`

GetChangerequestOk returns a tuple with the Changerequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangerequest

`func (o *JobRest) SetChangerequest(v int32)`

SetChangerequest sets Changerequest field to given value.

### HasChangerequest

`func (o *JobRest) HasChangerequest() bool`

HasChangerequest returns a boolean if a field has been set.

### GetProgress

`func (o *JobRest) GetProgress() int64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *JobRest) GetProgressOk() (*int64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *JobRest) SetProgress(v int64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *JobRest) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetRelativesize

`func (o *JobRest) GetRelativesize() int64`

GetRelativesize returns the Relativesize field if non-nil, zero value otherwise.

### GetRelativesizeOk

`func (o *JobRest) GetRelativesizeOk() (*int64, bool)`

GetRelativesizeOk returns a tuple with the Relativesize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativesize

`func (o *JobRest) SetRelativesize(v int64)`

SetRelativesize sets Relativesize field to given value.

### HasRelativesize

`func (o *JobRest) HasRelativesize() bool`

HasRelativesize returns a boolean if a field has been set.

### GetRetrycount

`func (o *JobRest) GetRetrycount() int32`

GetRetrycount returns the Retrycount field if non-nil, zero value otherwise.

### GetRetrycountOk

`func (o *JobRest) GetRetrycountOk() (*int32, bool)`

GetRetrycountOk returns a tuple with the Retrycount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrycount

`func (o *JobRest) SetRetrycount(v int32)`

SetRetrycount sets Retrycount field to given value.

### HasRetrycount

`func (o *JobRest) HasRetrycount() bool`

HasRetrycount returns a boolean if a field has been set.

### GetQueuedate

`func (o *JobRest) GetQueuedate() int64`

GetQueuedate returns the Queuedate field if non-nil, zero value otherwise.

### GetQueuedateOk

`func (o *JobRest) GetQueuedateOk() (*int64, bool)`

GetQueuedateOk returns a tuple with the Queuedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedate

`func (o *JobRest) SetQueuedate(v int64)`

SetQueuedate sets Queuedate field to given value.

### HasQueuedate

`func (o *JobRest) HasQueuedate() bool`

HasQueuedate returns a boolean if a field has been set.

### GetExpirationdate

`func (o *JobRest) GetExpirationdate() int64`

GetExpirationdate returns the Expirationdate field if non-nil, zero value otherwise.

### GetExpirationdateOk

`func (o *JobRest) GetExpirationdateOk() (*int64, bool)`

GetExpirationdateOk returns a tuple with the Expirationdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationdate

`func (o *JobRest) SetExpirationdate(v int64)`

SetExpirationdate sets Expirationdate field to given value.

### HasExpirationdate

`func (o *JobRest) HasExpirationdate() bool`

HasExpirationdate returns a boolean if a field has been set.

### GetIsscheduled

`func (o *JobRest) GetIsscheduled() bool`

GetIsscheduled returns the Isscheduled field if non-nil, zero value otherwise.

### GetIsscheduledOk

`func (o *JobRest) GetIsscheduledOk() (*bool, bool)`

GetIsscheduledOk returns a tuple with the Isscheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsscheduled

`func (o *JobRest) SetIsscheduled(v bool)`

SetIsscheduled sets Isscheduled field to given value.

### HasIsscheduled

`func (o *JobRest) HasIsscheduled() bool`

HasIsscheduled returns a boolean if a field has been set.

### GetJobtag

`func (o *JobRest) GetJobtag() string`

GetJobtag returns the Jobtag field if non-nil, zero value otherwise.

### GetJobtagOk

`func (o *JobRest) GetJobtagOk() (*string, bool)`

GetJobtagOk returns a tuple with the Jobtag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobtag

`func (o *JobRest) SetJobtag(v string)`

SetJobtag sets Jobtag field to given value.

### HasJobtag

`func (o *JobRest) HasJobtag() bool`

HasJobtag returns a boolean if a field has been set.

### GetEvent

`func (o *JobRest) GetEvent() EventRest`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *JobRest) GetEventOk() (*EventRest, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *JobRest) SetEvent(v EventRest)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *JobRest) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetId2

`func (o *JobRest) GetId2() string`

GetId2 returns the Id2 field if non-nil, zero value otherwise.

### GetId2Ok

`func (o *JobRest) GetId2Ok() (*string, bool)`

GetId2Ok returns a tuple with the Id2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId2

`func (o *JobRest) SetId2(v string)`

SetId2 sets Id2 field to given value.

### HasId2

`func (o *JobRest) HasId2() bool`

HasId2 returns a boolean if a field has been set.

### GetChangerequesttext

`func (o *JobRest) GetChangerequesttext() string`

GetChangerequesttext returns the Changerequesttext field if non-nil, zero value otherwise.

### GetChangerequesttextOk

`func (o *JobRest) GetChangerequesttextOk() (*string, bool)`

GetChangerequesttextOk returns a tuple with the Changerequesttext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangerequesttext

`func (o *JobRest) SetChangerequesttext(v string)`

SetChangerequesttext sets Changerequesttext field to given value.

### HasChangerequesttext

`func (o *JobRest) HasChangerequesttext() bool`

HasChangerequesttext returns a boolean if a field has been set.

### GetLogsmarttype

`func (o *JobRest) GetLogsmarttype() string`

GetLogsmarttype returns the Logsmarttype field if non-nil, zero value otherwise.

### GetLogsmarttypeOk

`func (o *JobRest) GetLogsmarttypeOk() (*string, bool)`

GetLogsmarttypeOk returns a tuple with the Logsmarttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsmarttype

`func (o *JobRest) SetLogsmarttype(v string)`

SetLogsmarttype sets Logsmarttype field to given value.

### HasLogsmarttype

`func (o *JobRest) HasLogsmarttype() bool`

HasLogsmarttype returns a boolean if a field has been set.

### GetExtrainfo

`func (o *JobRest) GetExtrainfo() string`

GetExtrainfo returns the Extrainfo field if non-nil, zero value otherwise.

### GetExtrainfoOk

`func (o *JobRest) GetExtrainfoOk() (*string, bool)`

GetExtrainfoOk returns a tuple with the Extrainfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtrainfo

`func (o *JobRest) SetExtrainfo(v string)`

SetExtrainfo sets Extrainfo field to given value.

### HasExtrainfo

`func (o *JobRest) HasExtrainfo() bool`

HasExtrainfo returns a boolean if a field has been set.

### GetDiskpools

`func (o *JobRest) GetDiskpools() []DiskPoolRest`

GetDiskpools returns the Diskpools field if non-nil, zero value otherwise.

### GetDiskpoolsOk

`func (o *JobRest) GetDiskpoolsOk() (*[]DiskPoolRest, bool)`

GetDiskpoolsOk returns a tuple with the Diskpools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskpools

`func (o *JobRest) SetDiskpools(v []DiskPoolRest)`

SetDiskpools sets Diskpools field to given value.

### HasDiskpools

`func (o *JobRest) HasDiskpools() bool`

HasDiskpools returns a boolean if a field has been set.

### GetTargetdiskpool

`func (o *JobRest) GetTargetdiskpool() DiskPoolRest`

GetTargetdiskpool returns the Targetdiskpool field if non-nil, zero value otherwise.

### GetTargetdiskpoolOk

`func (o *JobRest) GetTargetdiskpoolOk() (*DiskPoolRest, bool)`

GetTargetdiskpoolOk returns a tuple with the Targetdiskpool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetdiskpool

`func (o *JobRest) SetTargetdiskpool(v DiskPoolRest)`

SetTargetdiskpool sets Targetdiskpool field to given value.

### HasTargetdiskpool

`func (o *JobRest) HasTargetdiskpool() bool`

HasTargetdiskpool returns a boolean if a field has been set.

### GetJobclasscode

`func (o *JobRest) GetJobclasscode() int32`

GetJobclasscode returns the Jobclasscode field if non-nil, zero value otherwise.

### GetJobclasscodeOk

`func (o *JobRest) GetJobclasscodeOk() (*int32, bool)`

GetJobclasscodeOk returns a tuple with the Jobclasscode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobclasscode

`func (o *JobRest) SetJobclasscode(v int32)`

SetJobclasscode sets Jobclasscode field to given value.

### HasJobclasscode

`func (o *JobRest) HasJobclasscode() bool`

HasJobclasscode returns a boolean if a field has been set.

### GetSourcediskpool

`func (o *JobRest) GetSourcediskpool() DiskPoolRest`

GetSourcediskpool returns the Sourcediskpool field if non-nil, zero value otherwise.

### GetSourcediskpoolOk

`func (o *JobRest) GetSourcediskpoolOk() (*DiskPoolRest, bool)`

GetSourcediskpoolOk returns a tuple with the Sourcediskpool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcediskpool

`func (o *JobRest) SetSourcediskpool(v DiskPoolRest)`

SetSourcediskpool sets Sourcediskpool field to given value.

### HasSourcediskpool

`func (o *JobRest) HasSourcediskpool() bool`

HasSourcediskpool returns a boolean if a field has been set.

### GetYaml

`func (o *JobRest) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *JobRest) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *JobRest) SetYaml(v string)`

SetYaml sets Yaml field to given value.

### HasYaml

`func (o *JobRest) HasYaml() bool`

HasYaml returns a boolean if a field has been set.

### GetAppliance

`func (o *JobRest) GetAppliance() ClusterRest`

GetAppliance returns the Appliance field if non-nil, zero value otherwise.

### GetApplianceOk

`func (o *JobRest) GetApplianceOk() (*ClusterRest, bool)`

GetApplianceOk returns a tuple with the Appliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliance

`func (o *JobRest) SetAppliance(v ClusterRest)`

SetAppliance sets Appliance field to given value.

### HasAppliance

`func (o *JobRest) HasAppliance() bool`

HasAppliance returns a boolean if a field has been set.

### GetConsistencyMode

`func (o *JobRest) GetConsistencyMode() string`

GetConsistencyMode returns the ConsistencyMode field if non-nil, zero value otherwise.

### GetConsistencyModeOk

`func (o *JobRest) GetConsistencyModeOk() (*string, bool)`

GetConsistencyModeOk returns a tuple with the ConsistencyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistencyMode

`func (o *JobRest) SetConsistencyMode(v string)`

SetConsistencyMode sets ConsistencyMode field to given value.

### HasConsistencyMode

`func (o *JobRest) HasConsistencyMode() bool`

HasConsistencyMode returns a boolean if a field has been set.

### GetFlagsText

`func (o *JobRest) GetFlagsText() []string`

GetFlagsText returns the FlagsText field if non-nil, zero value otherwise.

### GetFlagsTextOk

`func (o *JobRest) GetFlagsTextOk() (*[]string, bool)`

GetFlagsTextOk returns a tuple with the FlagsText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagsText

`func (o *JobRest) SetFlagsText(v []string)`

SetFlagsText sets FlagsText field to given value.

### HasFlagsText

`func (o *JobRest) HasFlagsText() bool`

HasFlagsText returns a boolean if a field has been set.

### GetId

`func (o *JobRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JobRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *JobRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *JobRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *JobRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *JobRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *JobRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *JobRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *JobRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *JobRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *JobRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *JobRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *JobRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *JobRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



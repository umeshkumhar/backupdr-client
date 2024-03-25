# LogicalGroupRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Srcid** | Pointer to **string** |  | [optional] 
**Modifydate** | Pointer to **int64** |  | [optional] 
**Managed** | Pointer to **bool** |  | [optional] 
**Scheduleoff** | Pointer to **bool** |  | [optional] 
**Sla** | Pointer to [**SlaRest**](SlaRest.md) |  | [optional] 
**Cluster** | Pointer to [**ClusterRest**](ClusterRest.md) |  | [optional] 
**Membercount** | Pointer to **int32** |  | [optional] 
**Orglist** | Pointer to [**[]OrganizationRest**](OrganizationRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewLogicalGroupRest

`func NewLogicalGroupRest() *LogicalGroupRest`

NewLogicalGroupRest instantiates a new LogicalGroupRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicalGroupRestWithDefaults

`func NewLogicalGroupRestWithDefaults() *LogicalGroupRest`

NewLogicalGroupRestWithDefaults instantiates a new LogicalGroupRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *LogicalGroupRest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LogicalGroupRest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LogicalGroupRest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LogicalGroupRest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *LogicalGroupRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogicalGroupRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogicalGroupRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogicalGroupRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSrcid

`func (o *LogicalGroupRest) GetSrcid() string`

GetSrcid returns the Srcid field if non-nil, zero value otherwise.

### GetSrcidOk

`func (o *LogicalGroupRest) GetSrcidOk() (*string, bool)`

GetSrcidOk returns a tuple with the Srcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcid

`func (o *LogicalGroupRest) SetSrcid(v string)`

SetSrcid sets Srcid field to given value.

### HasSrcid

`func (o *LogicalGroupRest) HasSrcid() bool`

HasSrcid returns a boolean if a field has been set.

### GetModifydate

`func (o *LogicalGroupRest) GetModifydate() int64`

GetModifydate returns the Modifydate field if non-nil, zero value otherwise.

### GetModifydateOk

`func (o *LogicalGroupRest) GetModifydateOk() (*int64, bool)`

GetModifydateOk returns a tuple with the Modifydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifydate

`func (o *LogicalGroupRest) SetModifydate(v int64)`

SetModifydate sets Modifydate field to given value.

### HasModifydate

`func (o *LogicalGroupRest) HasModifydate() bool`

HasModifydate returns a boolean if a field has been set.

### GetManaged

`func (o *LogicalGroupRest) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *LogicalGroupRest) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *LogicalGroupRest) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *LogicalGroupRest) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetScheduleoff

`func (o *LogicalGroupRest) GetScheduleoff() bool`

GetScheduleoff returns the Scheduleoff field if non-nil, zero value otherwise.

### GetScheduleoffOk

`func (o *LogicalGroupRest) GetScheduleoffOk() (*bool, bool)`

GetScheduleoffOk returns a tuple with the Scheduleoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleoff

`func (o *LogicalGroupRest) SetScheduleoff(v bool)`

SetScheduleoff sets Scheduleoff field to given value.

### HasScheduleoff

`func (o *LogicalGroupRest) HasScheduleoff() bool`

HasScheduleoff returns a boolean if a field has been set.

### GetSla

`func (o *LogicalGroupRest) GetSla() SlaRest`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *LogicalGroupRest) GetSlaOk() (*SlaRest, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *LogicalGroupRest) SetSla(v SlaRest)`

SetSla sets Sla field to given value.

### HasSla

`func (o *LogicalGroupRest) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetCluster

`func (o *LogicalGroupRest) GetCluster() ClusterRest`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *LogicalGroupRest) GetClusterOk() (*ClusterRest, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *LogicalGroupRest) SetCluster(v ClusterRest)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *LogicalGroupRest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetMembercount

`func (o *LogicalGroupRest) GetMembercount() int32`

GetMembercount returns the Membercount field if non-nil, zero value otherwise.

### GetMembercountOk

`func (o *LogicalGroupRest) GetMembercountOk() (*int32, bool)`

GetMembercountOk returns a tuple with the Membercount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembercount

`func (o *LogicalGroupRest) SetMembercount(v int32)`

SetMembercount sets Membercount field to given value.

### HasMembercount

`func (o *LogicalGroupRest) HasMembercount() bool`

HasMembercount returns a boolean if a field has been set.

### GetOrglist

`func (o *LogicalGroupRest) GetOrglist() []OrganizationRest`

GetOrglist returns the Orglist field if non-nil, zero value otherwise.

### GetOrglistOk

`func (o *LogicalGroupRest) GetOrglistOk() (*[]OrganizationRest, bool)`

GetOrglistOk returns a tuple with the Orglist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrglist

`func (o *LogicalGroupRest) SetOrglist(v []OrganizationRest)`

SetOrglist sets Orglist field to given value.

### HasOrglist

`func (o *LogicalGroupRest) HasOrglist() bool`

HasOrglist returns a boolean if a field has been set.

### GetId

`func (o *LogicalGroupRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogicalGroupRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogicalGroupRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LogicalGroupRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *LogicalGroupRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LogicalGroupRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LogicalGroupRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *LogicalGroupRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *LogicalGroupRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *LogicalGroupRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *LogicalGroupRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *LogicalGroupRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *LogicalGroupRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *LogicalGroupRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *LogicalGroupRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *LogicalGroupRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



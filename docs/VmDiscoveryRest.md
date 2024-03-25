# VmDiscoveryRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to **string** |  | [optional] 
**Addvms** | Pointer to **bool** |  | [optional] 
**VcenterHostId** | Pointer to **string** |  | [optional] 
**Esxcluster** | Pointer to **string** |  | [optional] 
**Addall** | Pointer to **bool** |  | [optional] 
**Vms** | Pointer to **[]string** |  | [optional] 
**Dc** | Pointer to **string** |  | [optional] 
**Discoverclusters** | Pointer to **bool** |  | [optional] 
**Discovervms** | Pointer to **bool** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewVmDiscoveryRest

`func NewVmDiscoveryRest() *VmDiscoveryRest`

NewVmDiscoveryRest instantiates a new VmDiscoveryRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmDiscoveryRestWithDefaults

`func NewVmDiscoveryRestWithDefaults() *VmDiscoveryRest`

NewVmDiscoveryRestWithDefaults instantiates a new VmDiscoveryRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *VmDiscoveryRest) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VmDiscoveryRest) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VmDiscoveryRest) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VmDiscoveryRest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetAddvms

`func (o *VmDiscoveryRest) GetAddvms() bool`

GetAddvms returns the Addvms field if non-nil, zero value otherwise.

### GetAddvmsOk

`func (o *VmDiscoveryRest) GetAddvmsOk() (*bool, bool)`

GetAddvmsOk returns a tuple with the Addvms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddvms

`func (o *VmDiscoveryRest) SetAddvms(v bool)`

SetAddvms sets Addvms field to given value.

### HasAddvms

`func (o *VmDiscoveryRest) HasAddvms() bool`

HasAddvms returns a boolean if a field has been set.

### GetVcenterHostId

`func (o *VmDiscoveryRest) GetVcenterHostId() string`

GetVcenterHostId returns the VcenterHostId field if non-nil, zero value otherwise.

### GetVcenterHostIdOk

`func (o *VmDiscoveryRest) GetVcenterHostIdOk() (*string, bool)`

GetVcenterHostIdOk returns a tuple with the VcenterHostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterHostId

`func (o *VmDiscoveryRest) SetVcenterHostId(v string)`

SetVcenterHostId sets VcenterHostId field to given value.

### HasVcenterHostId

`func (o *VmDiscoveryRest) HasVcenterHostId() bool`

HasVcenterHostId returns a boolean if a field has been set.

### GetEsxcluster

`func (o *VmDiscoveryRest) GetEsxcluster() string`

GetEsxcluster returns the Esxcluster field if non-nil, zero value otherwise.

### GetEsxclusterOk

`func (o *VmDiscoveryRest) GetEsxclusterOk() (*string, bool)`

GetEsxclusterOk returns a tuple with the Esxcluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsxcluster

`func (o *VmDiscoveryRest) SetEsxcluster(v string)`

SetEsxcluster sets Esxcluster field to given value.

### HasEsxcluster

`func (o *VmDiscoveryRest) HasEsxcluster() bool`

HasEsxcluster returns a boolean if a field has been set.

### GetAddall

`func (o *VmDiscoveryRest) GetAddall() bool`

GetAddall returns the Addall field if non-nil, zero value otherwise.

### GetAddallOk

`func (o *VmDiscoveryRest) GetAddallOk() (*bool, bool)`

GetAddallOk returns a tuple with the Addall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddall

`func (o *VmDiscoveryRest) SetAddall(v bool)`

SetAddall sets Addall field to given value.

### HasAddall

`func (o *VmDiscoveryRest) HasAddall() bool`

HasAddall returns a boolean if a field has been set.

### GetVms

`func (o *VmDiscoveryRest) GetVms() []string`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *VmDiscoveryRest) GetVmsOk() (*[]string, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVms

`func (o *VmDiscoveryRest) SetVms(v []string)`

SetVms sets Vms field to given value.

### HasVms

`func (o *VmDiscoveryRest) HasVms() bool`

HasVms returns a boolean if a field has been set.

### GetDc

`func (o *VmDiscoveryRest) GetDc() string`

GetDc returns the Dc field if non-nil, zero value otherwise.

### GetDcOk

`func (o *VmDiscoveryRest) GetDcOk() (*string, bool)`

GetDcOk returns a tuple with the Dc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDc

`func (o *VmDiscoveryRest) SetDc(v string)`

SetDc sets Dc field to given value.

### HasDc

`func (o *VmDiscoveryRest) HasDc() bool`

HasDc returns a boolean if a field has been set.

### GetDiscoverclusters

`func (o *VmDiscoveryRest) GetDiscoverclusters() bool`

GetDiscoverclusters returns the Discoverclusters field if non-nil, zero value otherwise.

### GetDiscoverclustersOk

`func (o *VmDiscoveryRest) GetDiscoverclustersOk() (*bool, bool)`

GetDiscoverclustersOk returns a tuple with the Discoverclusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverclusters

`func (o *VmDiscoveryRest) SetDiscoverclusters(v bool)`

SetDiscoverclusters sets Discoverclusters field to given value.

### HasDiscoverclusters

`func (o *VmDiscoveryRest) HasDiscoverclusters() bool`

HasDiscoverclusters returns a boolean if a field has been set.

### GetDiscovervms

`func (o *VmDiscoveryRest) GetDiscovervms() bool`

GetDiscovervms returns the Discovervms field if non-nil, zero value otherwise.

### GetDiscovervmsOk

`func (o *VmDiscoveryRest) GetDiscovervmsOk() (*bool, bool)`

GetDiscovervmsOk returns a tuple with the Discovervms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovervms

`func (o *VmDiscoveryRest) SetDiscovervms(v bool)`

SetDiscovervms sets Discovervms field to given value.

### HasDiscovervms

`func (o *VmDiscoveryRest) HasDiscovervms() bool`

HasDiscovervms returns a boolean if a field has been set.

### GetOrg

`func (o *VmDiscoveryRest) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *VmDiscoveryRest) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *VmDiscoveryRest) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *VmDiscoveryRest) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetId

`func (o *VmDiscoveryRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VmDiscoveryRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VmDiscoveryRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VmDiscoveryRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *VmDiscoveryRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *VmDiscoveryRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *VmDiscoveryRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *VmDiscoveryRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *VmDiscoveryRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *VmDiscoveryRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *VmDiscoveryRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *VmDiscoveryRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *VmDiscoveryRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *VmDiscoveryRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *VmDiscoveryRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *VmDiscoveryRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



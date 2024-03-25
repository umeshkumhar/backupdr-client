# DiskPoolRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Usedefaultsa** | Pointer to **bool** |  | [optional] 
**Immutable** | Pointer to **bool** |  | [optional] 
**Metadataonly** | Pointer to **bool** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Array** | Pointer to [**ArrayRest**](ArrayRest.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to [**[]KeyValueRest**](KeyValueRest.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Srcid** | Pointer to **string** |  | [optional] 
**Modifydate** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**ClusterRest**](ClusterRest.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Mdiskgrp** | Pointer to **string** |  | [optional] 
**Warnpct** | Pointer to **int32** |  | [optional] 
**Safepct** | Pointer to **int32** |  | [optional] 
**Udsuid** | Pointer to **int32** |  | [optional] 
**Orglist** | Pointer to [**[]OrganizationRest**](OrganizationRest.md) |  | [optional] 
**Storage** | Pointer to **[]string** |  | [optional] 
**FreeMb** | Pointer to **int64** |  | [optional] 
**UsageMb** | Pointer to **int64** |  | [optional] 
**CapacityMb** | Pointer to **int64** |  | [optional] 
**Pct** | Pointer to **float64** |  | [optional] 
**Pooltypedisplayname** | Pointer to **string** |  | [optional] 
**Vaultprops** | Pointer to [**VaultPropsRest**](VaultPropsRest.md) |  | [optional] 
**Ext** | Pointer to **int32** |  | [optional] 
**Pooltype** | Pointer to **string** |  | [optional] 
**Cloudcredential** | Pointer to [**CloudCredentialRest**](CloudCredentialRest.md) |  | [optional] 
**Nocache** | Pointer to **bool** |  | [optional] 
**Grainsize** | Pointer to **int32** |  | [optional] 
**Copies** | Pointer to [**[]DiskPoolCopiesRest**](DiskPoolCopiesRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewDiskPoolRest

`func NewDiskPoolRest() *DiskPoolRest`

NewDiskPoolRest instantiates a new DiskPoolRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskPoolRestWithDefaults

`func NewDiskPoolRestWithDefaults() *DiskPoolRest`

NewDiskPoolRestWithDefaults instantiates a new DiskPoolRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsedefaultsa

`func (o *DiskPoolRest) GetUsedefaultsa() bool`

GetUsedefaultsa returns the Usedefaultsa field if non-nil, zero value otherwise.

### GetUsedefaultsaOk

`func (o *DiskPoolRest) GetUsedefaultsaOk() (*bool, bool)`

GetUsedefaultsaOk returns a tuple with the Usedefaultsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedefaultsa

`func (o *DiskPoolRest) SetUsedefaultsa(v bool)`

SetUsedefaultsa sets Usedefaultsa field to given value.

### HasUsedefaultsa

`func (o *DiskPoolRest) HasUsedefaultsa() bool`

HasUsedefaultsa returns a boolean if a field has been set.

### GetImmutable

`func (o *DiskPoolRest) GetImmutable() bool`

GetImmutable returns the Immutable field if non-nil, zero value otherwise.

### GetImmutableOk

`func (o *DiskPoolRest) GetImmutableOk() (*bool, bool)`

GetImmutableOk returns a tuple with the Immutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutable

`func (o *DiskPoolRest) SetImmutable(v bool)`

SetImmutable sets Immutable field to given value.

### HasImmutable

`func (o *DiskPoolRest) HasImmutable() bool`

HasImmutable returns a boolean if a field has been set.

### GetMetadataonly

`func (o *DiskPoolRest) GetMetadataonly() bool`

GetMetadataonly returns the Metadataonly field if non-nil, zero value otherwise.

### GetMetadataonlyOk

`func (o *DiskPoolRest) GetMetadataonlyOk() (*bool, bool)`

GetMetadataonlyOk returns a tuple with the Metadataonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataonly

`func (o *DiskPoolRest) SetMetadataonly(v bool)`

SetMetadataonly sets Metadataonly field to given value.

### HasMetadataonly

`func (o *DiskPoolRest) HasMetadataonly() bool`

HasMetadataonly returns a boolean if a field has been set.

### GetLocation

`func (o *DiskPoolRest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DiskPoolRest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DiskPoolRest) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DiskPoolRest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetArray

`func (o *DiskPoolRest) GetArray() ArrayRest`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *DiskPoolRest) GetArrayOk() (*ArrayRest, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *DiskPoolRest) SetArray(v ArrayRest)`

SetArray sets Array field to given value.

### HasArray

`func (o *DiskPoolRest) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetName

`func (o *DiskPoolRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiskPoolRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiskPoolRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DiskPoolRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *DiskPoolRest) GetProperties() []KeyValueRest`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DiskPoolRest) GetPropertiesOk() (*[]KeyValueRest, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DiskPoolRest) SetProperties(v []KeyValueRest)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DiskPoolRest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetState

`func (o *DiskPoolRest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DiskPoolRest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DiskPoolRest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DiskPoolRest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *DiskPoolRest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiskPoolRest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiskPoolRest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DiskPoolRest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSrcid

`func (o *DiskPoolRest) GetSrcid() string`

GetSrcid returns the Srcid field if non-nil, zero value otherwise.

### GetSrcidOk

`func (o *DiskPoolRest) GetSrcidOk() (*string, bool)`

GetSrcidOk returns a tuple with the Srcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcid

`func (o *DiskPoolRest) SetSrcid(v string)`

SetSrcid sets Srcid field to given value.

### HasSrcid

`func (o *DiskPoolRest) HasSrcid() bool`

HasSrcid returns a boolean if a field has been set.

### GetModifydate

`func (o *DiskPoolRest) GetModifydate() int64`

GetModifydate returns the Modifydate field if non-nil, zero value otherwise.

### GetModifydateOk

`func (o *DiskPoolRest) GetModifydateOk() (*int64, bool)`

GetModifydateOk returns a tuple with the Modifydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifydate

`func (o *DiskPoolRest) SetModifydate(v int64)`

SetModifydate sets Modifydate field to given value.

### HasModifydate

`func (o *DiskPoolRest) HasModifydate() bool`

HasModifydate returns a boolean if a field has been set.

### GetCluster

`func (o *DiskPoolRest) GetCluster() ClusterRest`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DiskPoolRest) GetClusterOk() (*ClusterRest, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DiskPoolRest) SetCluster(v ClusterRest)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DiskPoolRest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetStatus

`func (o *DiskPoolRest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DiskPoolRest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DiskPoolRest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DiskPoolRest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMdiskgrp

`func (o *DiskPoolRest) GetMdiskgrp() string`

GetMdiskgrp returns the Mdiskgrp field if non-nil, zero value otherwise.

### GetMdiskgrpOk

`func (o *DiskPoolRest) GetMdiskgrpOk() (*string, bool)`

GetMdiskgrpOk returns a tuple with the Mdiskgrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdiskgrp

`func (o *DiskPoolRest) SetMdiskgrp(v string)`

SetMdiskgrp sets Mdiskgrp field to given value.

### HasMdiskgrp

`func (o *DiskPoolRest) HasMdiskgrp() bool`

HasMdiskgrp returns a boolean if a field has been set.

### GetWarnpct

`func (o *DiskPoolRest) GetWarnpct() int32`

GetWarnpct returns the Warnpct field if non-nil, zero value otherwise.

### GetWarnpctOk

`func (o *DiskPoolRest) GetWarnpctOk() (*int32, bool)`

GetWarnpctOk returns a tuple with the Warnpct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnpct

`func (o *DiskPoolRest) SetWarnpct(v int32)`

SetWarnpct sets Warnpct field to given value.

### HasWarnpct

`func (o *DiskPoolRest) HasWarnpct() bool`

HasWarnpct returns a boolean if a field has been set.

### GetSafepct

`func (o *DiskPoolRest) GetSafepct() int32`

GetSafepct returns the Safepct field if non-nil, zero value otherwise.

### GetSafepctOk

`func (o *DiskPoolRest) GetSafepctOk() (*int32, bool)`

GetSafepctOk returns a tuple with the Safepct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafepct

`func (o *DiskPoolRest) SetSafepct(v int32)`

SetSafepct sets Safepct field to given value.

### HasSafepct

`func (o *DiskPoolRest) HasSafepct() bool`

HasSafepct returns a boolean if a field has been set.

### GetUdsuid

`func (o *DiskPoolRest) GetUdsuid() int32`

GetUdsuid returns the Udsuid field if non-nil, zero value otherwise.

### GetUdsuidOk

`func (o *DiskPoolRest) GetUdsuidOk() (*int32, bool)`

GetUdsuidOk returns a tuple with the Udsuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdsuid

`func (o *DiskPoolRest) SetUdsuid(v int32)`

SetUdsuid sets Udsuid field to given value.

### HasUdsuid

`func (o *DiskPoolRest) HasUdsuid() bool`

HasUdsuid returns a boolean if a field has been set.

### GetOrglist

`func (o *DiskPoolRest) GetOrglist() []OrganizationRest`

GetOrglist returns the Orglist field if non-nil, zero value otherwise.

### GetOrglistOk

`func (o *DiskPoolRest) GetOrglistOk() (*[]OrganizationRest, bool)`

GetOrglistOk returns a tuple with the Orglist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrglist

`func (o *DiskPoolRest) SetOrglist(v []OrganizationRest)`

SetOrglist sets Orglist field to given value.

### HasOrglist

`func (o *DiskPoolRest) HasOrglist() bool`

HasOrglist returns a boolean if a field has been set.

### GetStorage

`func (o *DiskPoolRest) GetStorage() []string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *DiskPoolRest) GetStorageOk() (*[]string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *DiskPoolRest) SetStorage(v []string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *DiskPoolRest) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetFreeMb

`func (o *DiskPoolRest) GetFreeMb() int64`

GetFreeMb returns the FreeMb field if non-nil, zero value otherwise.

### GetFreeMbOk

`func (o *DiskPoolRest) GetFreeMbOk() (*int64, bool)`

GetFreeMbOk returns a tuple with the FreeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeMb

`func (o *DiskPoolRest) SetFreeMb(v int64)`

SetFreeMb sets FreeMb field to given value.

### HasFreeMb

`func (o *DiskPoolRest) HasFreeMb() bool`

HasFreeMb returns a boolean if a field has been set.

### GetUsageMb

`func (o *DiskPoolRest) GetUsageMb() int64`

GetUsageMb returns the UsageMb field if non-nil, zero value otherwise.

### GetUsageMbOk

`func (o *DiskPoolRest) GetUsageMbOk() (*int64, bool)`

GetUsageMbOk returns a tuple with the UsageMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageMb

`func (o *DiskPoolRest) SetUsageMb(v int64)`

SetUsageMb sets UsageMb field to given value.

### HasUsageMb

`func (o *DiskPoolRest) HasUsageMb() bool`

HasUsageMb returns a boolean if a field has been set.

### GetCapacityMb

`func (o *DiskPoolRest) GetCapacityMb() int64`

GetCapacityMb returns the CapacityMb field if non-nil, zero value otherwise.

### GetCapacityMbOk

`func (o *DiskPoolRest) GetCapacityMbOk() (*int64, bool)`

GetCapacityMbOk returns a tuple with the CapacityMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityMb

`func (o *DiskPoolRest) SetCapacityMb(v int64)`

SetCapacityMb sets CapacityMb field to given value.

### HasCapacityMb

`func (o *DiskPoolRest) HasCapacityMb() bool`

HasCapacityMb returns a boolean if a field has been set.

### GetPct

`func (o *DiskPoolRest) GetPct() float64`

GetPct returns the Pct field if non-nil, zero value otherwise.

### GetPctOk

`func (o *DiskPoolRest) GetPctOk() (*float64, bool)`

GetPctOk returns a tuple with the Pct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPct

`func (o *DiskPoolRest) SetPct(v float64)`

SetPct sets Pct field to given value.

### HasPct

`func (o *DiskPoolRest) HasPct() bool`

HasPct returns a boolean if a field has been set.

### GetPooltypedisplayname

`func (o *DiskPoolRest) GetPooltypedisplayname() string`

GetPooltypedisplayname returns the Pooltypedisplayname field if non-nil, zero value otherwise.

### GetPooltypedisplaynameOk

`func (o *DiskPoolRest) GetPooltypedisplaynameOk() (*string, bool)`

GetPooltypedisplaynameOk returns a tuple with the Pooltypedisplayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPooltypedisplayname

`func (o *DiskPoolRest) SetPooltypedisplayname(v string)`

SetPooltypedisplayname sets Pooltypedisplayname field to given value.

### HasPooltypedisplayname

`func (o *DiskPoolRest) HasPooltypedisplayname() bool`

HasPooltypedisplayname returns a boolean if a field has been set.

### GetVaultprops

`func (o *DiskPoolRest) GetVaultprops() VaultPropsRest`

GetVaultprops returns the Vaultprops field if non-nil, zero value otherwise.

### GetVaultpropsOk

`func (o *DiskPoolRest) GetVaultpropsOk() (*VaultPropsRest, bool)`

GetVaultpropsOk returns a tuple with the Vaultprops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultprops

`func (o *DiskPoolRest) SetVaultprops(v VaultPropsRest)`

SetVaultprops sets Vaultprops field to given value.

### HasVaultprops

`func (o *DiskPoolRest) HasVaultprops() bool`

HasVaultprops returns a boolean if a field has been set.

### GetExt

`func (o *DiskPoolRest) GetExt() int32`

GetExt returns the Ext field if non-nil, zero value otherwise.

### GetExtOk

`func (o *DiskPoolRest) GetExtOk() (*int32, bool)`

GetExtOk returns a tuple with the Ext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExt

`func (o *DiskPoolRest) SetExt(v int32)`

SetExt sets Ext field to given value.

### HasExt

`func (o *DiskPoolRest) HasExt() bool`

HasExt returns a boolean if a field has been set.

### GetPooltype

`func (o *DiskPoolRest) GetPooltype() string`

GetPooltype returns the Pooltype field if non-nil, zero value otherwise.

### GetPooltypeOk

`func (o *DiskPoolRest) GetPooltypeOk() (*string, bool)`

GetPooltypeOk returns a tuple with the Pooltype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPooltype

`func (o *DiskPoolRest) SetPooltype(v string)`

SetPooltype sets Pooltype field to given value.

### HasPooltype

`func (o *DiskPoolRest) HasPooltype() bool`

HasPooltype returns a boolean if a field has been set.

### GetCloudcredential

`func (o *DiskPoolRest) GetCloudcredential() CloudCredentialRest`

GetCloudcredential returns the Cloudcredential field if non-nil, zero value otherwise.

### GetCloudcredentialOk

`func (o *DiskPoolRest) GetCloudcredentialOk() (*CloudCredentialRest, bool)`

GetCloudcredentialOk returns a tuple with the Cloudcredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudcredential

`func (o *DiskPoolRest) SetCloudcredential(v CloudCredentialRest)`

SetCloudcredential sets Cloudcredential field to given value.

### HasCloudcredential

`func (o *DiskPoolRest) HasCloudcredential() bool`

HasCloudcredential returns a boolean if a field has been set.

### GetNocache

`func (o *DiskPoolRest) GetNocache() bool`

GetNocache returns the Nocache field if non-nil, zero value otherwise.

### GetNocacheOk

`func (o *DiskPoolRest) GetNocacheOk() (*bool, bool)`

GetNocacheOk returns a tuple with the Nocache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNocache

`func (o *DiskPoolRest) SetNocache(v bool)`

SetNocache sets Nocache field to given value.

### HasNocache

`func (o *DiskPoolRest) HasNocache() bool`

HasNocache returns a boolean if a field has been set.

### GetGrainsize

`func (o *DiskPoolRest) GetGrainsize() int32`

GetGrainsize returns the Grainsize field if non-nil, zero value otherwise.

### GetGrainsizeOk

`func (o *DiskPoolRest) GetGrainsizeOk() (*int32, bool)`

GetGrainsizeOk returns a tuple with the Grainsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrainsize

`func (o *DiskPoolRest) SetGrainsize(v int32)`

SetGrainsize sets Grainsize field to given value.

### HasGrainsize

`func (o *DiskPoolRest) HasGrainsize() bool`

HasGrainsize returns a boolean if a field has been set.

### GetCopies

`func (o *DiskPoolRest) GetCopies() []DiskPoolCopiesRest`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *DiskPoolRest) GetCopiesOk() (*[]DiskPoolCopiesRest, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *DiskPoolRest) SetCopies(v []DiskPoolCopiesRest)`

SetCopies sets Copies field to given value.

### HasCopies

`func (o *DiskPoolRest) HasCopies() bool`

HasCopies returns a boolean if a field has been set.

### GetId

`func (o *DiskPoolRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiskPoolRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiskPoolRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DiskPoolRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *DiskPoolRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DiskPoolRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DiskPoolRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *DiskPoolRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *DiskPoolRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *DiskPoolRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *DiskPoolRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *DiskPoolRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *DiskPoolRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *DiskPoolRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *DiskPoolRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *DiskPoolRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



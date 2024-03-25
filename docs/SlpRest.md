# SlpRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Srcid** | Pointer to **string** |  | [optional] 
**Clusterid** | Pointer to **string** |  | [optional] 
**Modifydate** | Pointer to **int64** |  | [optional] 
**Cid** | Pointer to **string** |  | [optional] 
**Performancepool** | Pointer to **string** |  | [optional] 
**Primarystorage** | Pointer to **string** |  | [optional] 
**Remotenode** | Pointer to **string** |  | [optional] 
**Dedupasyncnode** | Pointer to **string** |  | [optional] 
**Vaultpool** | Pointer to [**DiskPoolRest**](DiskPoolRest.md) |  | [optional] 
**Vaultpool2** | Pointer to [**DiskPoolRest**](DiskPoolRest.md) |  | [optional] 
**Vaultpool3** | Pointer to [**DiskPoolRest**](DiskPoolRest.md) |  | [optional] 
**Vaultpool4** | Pointer to [**DiskPoolRest**](DiskPoolRest.md) |  | [optional] 
**Createdate** | Pointer to **int64** |  | [optional] 
**Localnode** | Pointer to **string** |  | [optional] 
**Orglist** | Pointer to [**[]OrganizationRest**](OrganizationRest.md) |  | [optional] 
**CloudCredential** | Pointer to [**CloudCredentialRest**](CloudCredentialRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewSlpRest

`func NewSlpRest() *SlpRest`

NewSlpRest instantiates a new SlpRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlpRestWithDefaults

`func NewSlpRestWithDefaults() *SlpRest`

NewSlpRestWithDefaults instantiates a new SlpRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SlpRest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SlpRest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SlpRest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SlpRest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *SlpRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlpRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlpRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SlpRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSrcid

`func (o *SlpRest) GetSrcid() string`

GetSrcid returns the Srcid field if non-nil, zero value otherwise.

### GetSrcidOk

`func (o *SlpRest) GetSrcidOk() (*string, bool)`

GetSrcidOk returns a tuple with the Srcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcid

`func (o *SlpRest) SetSrcid(v string)`

SetSrcid sets Srcid field to given value.

### HasSrcid

`func (o *SlpRest) HasSrcid() bool`

HasSrcid returns a boolean if a field has been set.

### GetClusterid

`func (o *SlpRest) GetClusterid() string`

GetClusterid returns the Clusterid field if non-nil, zero value otherwise.

### GetClusteridOk

`func (o *SlpRest) GetClusteridOk() (*string, bool)`

GetClusteridOk returns a tuple with the Clusterid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterid

`func (o *SlpRest) SetClusterid(v string)`

SetClusterid sets Clusterid field to given value.

### HasClusterid

`func (o *SlpRest) HasClusterid() bool`

HasClusterid returns a boolean if a field has been set.

### GetModifydate

`func (o *SlpRest) GetModifydate() int64`

GetModifydate returns the Modifydate field if non-nil, zero value otherwise.

### GetModifydateOk

`func (o *SlpRest) GetModifydateOk() (*int64, bool)`

GetModifydateOk returns a tuple with the Modifydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifydate

`func (o *SlpRest) SetModifydate(v int64)`

SetModifydate sets Modifydate field to given value.

### HasModifydate

`func (o *SlpRest) HasModifydate() bool`

HasModifydate returns a boolean if a field has been set.

### GetCid

`func (o *SlpRest) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *SlpRest) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *SlpRest) SetCid(v string)`

SetCid sets Cid field to given value.

### HasCid

`func (o *SlpRest) HasCid() bool`

HasCid returns a boolean if a field has been set.

### GetPerformancepool

`func (o *SlpRest) GetPerformancepool() string`

GetPerformancepool returns the Performancepool field if non-nil, zero value otherwise.

### GetPerformancepoolOk

`func (o *SlpRest) GetPerformancepoolOk() (*string, bool)`

GetPerformancepoolOk returns a tuple with the Performancepool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformancepool

`func (o *SlpRest) SetPerformancepool(v string)`

SetPerformancepool sets Performancepool field to given value.

### HasPerformancepool

`func (o *SlpRest) HasPerformancepool() bool`

HasPerformancepool returns a boolean if a field has been set.

### GetPrimarystorage

`func (o *SlpRest) GetPrimarystorage() string`

GetPrimarystorage returns the Primarystorage field if non-nil, zero value otherwise.

### GetPrimarystorageOk

`func (o *SlpRest) GetPrimarystorageOk() (*string, bool)`

GetPrimarystorageOk returns a tuple with the Primarystorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarystorage

`func (o *SlpRest) SetPrimarystorage(v string)`

SetPrimarystorage sets Primarystorage field to given value.

### HasPrimarystorage

`func (o *SlpRest) HasPrimarystorage() bool`

HasPrimarystorage returns a boolean if a field has been set.

### GetRemotenode

`func (o *SlpRest) GetRemotenode() string`

GetRemotenode returns the Remotenode field if non-nil, zero value otherwise.

### GetRemotenodeOk

`func (o *SlpRest) GetRemotenodeOk() (*string, bool)`

GetRemotenodeOk returns a tuple with the Remotenode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotenode

`func (o *SlpRest) SetRemotenode(v string)`

SetRemotenode sets Remotenode field to given value.

### HasRemotenode

`func (o *SlpRest) HasRemotenode() bool`

HasRemotenode returns a boolean if a field has been set.

### GetDedupasyncnode

`func (o *SlpRest) GetDedupasyncnode() string`

GetDedupasyncnode returns the Dedupasyncnode field if non-nil, zero value otherwise.

### GetDedupasyncnodeOk

`func (o *SlpRest) GetDedupasyncnodeOk() (*string, bool)`

GetDedupasyncnodeOk returns a tuple with the Dedupasyncnode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupasyncnode

`func (o *SlpRest) SetDedupasyncnode(v string)`

SetDedupasyncnode sets Dedupasyncnode field to given value.

### HasDedupasyncnode

`func (o *SlpRest) HasDedupasyncnode() bool`

HasDedupasyncnode returns a boolean if a field has been set.

### GetVaultpool

`func (o *SlpRest) GetVaultpool() DiskPoolRest`

GetVaultpool returns the Vaultpool field if non-nil, zero value otherwise.

### GetVaultpoolOk

`func (o *SlpRest) GetVaultpoolOk() (*DiskPoolRest, bool)`

GetVaultpoolOk returns a tuple with the Vaultpool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultpool

`func (o *SlpRest) SetVaultpool(v DiskPoolRest)`

SetVaultpool sets Vaultpool field to given value.

### HasVaultpool

`func (o *SlpRest) HasVaultpool() bool`

HasVaultpool returns a boolean if a field has been set.

### GetVaultpool2

`func (o *SlpRest) GetVaultpool2() DiskPoolRest`

GetVaultpool2 returns the Vaultpool2 field if non-nil, zero value otherwise.

### GetVaultpool2Ok

`func (o *SlpRest) GetVaultpool2Ok() (*DiskPoolRest, bool)`

GetVaultpool2Ok returns a tuple with the Vaultpool2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultpool2

`func (o *SlpRest) SetVaultpool2(v DiskPoolRest)`

SetVaultpool2 sets Vaultpool2 field to given value.

### HasVaultpool2

`func (o *SlpRest) HasVaultpool2() bool`

HasVaultpool2 returns a boolean if a field has been set.

### GetVaultpool3

`func (o *SlpRest) GetVaultpool3() DiskPoolRest`

GetVaultpool3 returns the Vaultpool3 field if non-nil, zero value otherwise.

### GetVaultpool3Ok

`func (o *SlpRest) GetVaultpool3Ok() (*DiskPoolRest, bool)`

GetVaultpool3Ok returns a tuple with the Vaultpool3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultpool3

`func (o *SlpRest) SetVaultpool3(v DiskPoolRest)`

SetVaultpool3 sets Vaultpool3 field to given value.

### HasVaultpool3

`func (o *SlpRest) HasVaultpool3() bool`

HasVaultpool3 returns a boolean if a field has been set.

### GetVaultpool4

`func (o *SlpRest) GetVaultpool4() DiskPoolRest`

GetVaultpool4 returns the Vaultpool4 field if non-nil, zero value otherwise.

### GetVaultpool4Ok

`func (o *SlpRest) GetVaultpool4Ok() (*DiskPoolRest, bool)`

GetVaultpool4Ok returns a tuple with the Vaultpool4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultpool4

`func (o *SlpRest) SetVaultpool4(v DiskPoolRest)`

SetVaultpool4 sets Vaultpool4 field to given value.

### HasVaultpool4

`func (o *SlpRest) HasVaultpool4() bool`

HasVaultpool4 returns a boolean if a field has been set.

### GetCreatedate

`func (o *SlpRest) GetCreatedate() int64`

GetCreatedate returns the Createdate field if non-nil, zero value otherwise.

### GetCreatedateOk

`func (o *SlpRest) GetCreatedateOk() (*int64, bool)`

GetCreatedateOk returns a tuple with the Createdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedate

`func (o *SlpRest) SetCreatedate(v int64)`

SetCreatedate sets Createdate field to given value.

### HasCreatedate

`func (o *SlpRest) HasCreatedate() bool`

HasCreatedate returns a boolean if a field has been set.

### GetLocalnode

`func (o *SlpRest) GetLocalnode() string`

GetLocalnode returns the Localnode field if non-nil, zero value otherwise.

### GetLocalnodeOk

`func (o *SlpRest) GetLocalnodeOk() (*string, bool)`

GetLocalnodeOk returns a tuple with the Localnode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalnode

`func (o *SlpRest) SetLocalnode(v string)`

SetLocalnode sets Localnode field to given value.

### HasLocalnode

`func (o *SlpRest) HasLocalnode() bool`

HasLocalnode returns a boolean if a field has been set.

### GetOrglist

`func (o *SlpRest) GetOrglist() []OrganizationRest`

GetOrglist returns the Orglist field if non-nil, zero value otherwise.

### GetOrglistOk

`func (o *SlpRest) GetOrglistOk() (*[]OrganizationRest, bool)`

GetOrglistOk returns a tuple with the Orglist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrglist

`func (o *SlpRest) SetOrglist(v []OrganizationRest)`

SetOrglist sets Orglist field to given value.

### HasOrglist

`func (o *SlpRest) HasOrglist() bool`

HasOrglist returns a boolean if a field has been set.

### GetCloudCredential

`func (o *SlpRest) GetCloudCredential() CloudCredentialRest`

GetCloudCredential returns the CloudCredential field if non-nil, zero value otherwise.

### GetCloudCredentialOk

`func (o *SlpRest) GetCloudCredentialOk() (*CloudCredentialRest, bool)`

GetCloudCredentialOk returns a tuple with the CloudCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredential

`func (o *SlpRest) SetCloudCredential(v CloudCredentialRest)`

SetCloudCredential sets CloudCredential field to given value.

### HasCloudCredential

`func (o *SlpRest) HasCloudCredential() bool`

HasCloudCredential returns a boolean if a field has been set.

### GetId

`func (o *SlpRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SlpRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SlpRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SlpRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *SlpRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SlpRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SlpRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SlpRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *SlpRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *SlpRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *SlpRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *SlpRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *SlpRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *SlpRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *SlpRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *SlpRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



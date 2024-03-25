# ArrayRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sources** | Pointer to [**[]ArrayRest**](ArrayRest.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to [**[]KeyValueRest**](KeyValueRest.md) |  | [optional] 
**Srcid** | Pointer to **int64** |  | [optional] 
**Clusterid** | Pointer to **int64** |  | [optional] 
**Modifydate** | Pointer to **int64** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Ipaddress** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Overallstatus** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Arraytype** | Pointer to **string** |  | [optional] 
**Orglist** | Pointer to [**[]OrganizationRest**](OrganizationRest.md) |  | [optional] 
**Storage** | Pointer to [**[]MdiskGroupRest**](MdiskGroupRest.md) |  | [optional] 
**Arraytypelabel** | Pointer to **string** |  | [optional] 
**Hostcount** | Pointer to **int32** |  | [optional] 
**Diskpools** | Pointer to [**[]DiskPoolRest**](DiskPoolRest.md) |  | [optional] 
**Reset** | Pointer to **bool** |  | [optional] 
**Appliance** | Pointer to [**ClusterRest**](ClusterRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewArrayRest

`func NewArrayRest() *ArrayRest`

NewArrayRest instantiates a new ArrayRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArrayRestWithDefaults

`func NewArrayRestWithDefaults() *ArrayRest`

NewArrayRestWithDefaults instantiates a new ArrayRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSources

`func (o *ArrayRest) GetSources() []ArrayRest`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *ArrayRest) GetSourcesOk() (*[]ArrayRest, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *ArrayRest) SetSources(v []ArrayRest)`

SetSources sets Sources field to given value.

### HasSources

`func (o *ArrayRest) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetName

`func (o *ArrayRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArrayRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArrayRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArrayRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *ArrayRest) GetProperties() []KeyValueRest`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ArrayRest) GetPropertiesOk() (*[]KeyValueRest, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ArrayRest) SetProperties(v []KeyValueRest)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ArrayRest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSrcid

`func (o *ArrayRest) GetSrcid() int64`

GetSrcid returns the Srcid field if non-nil, zero value otherwise.

### GetSrcidOk

`func (o *ArrayRest) GetSrcidOk() (*int64, bool)`

GetSrcidOk returns a tuple with the Srcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcid

`func (o *ArrayRest) SetSrcid(v int64)`

SetSrcid sets Srcid field to given value.

### HasSrcid

`func (o *ArrayRest) HasSrcid() bool`

HasSrcid returns a boolean if a field has been set.

### GetClusterid

`func (o *ArrayRest) GetClusterid() int64`

GetClusterid returns the Clusterid field if non-nil, zero value otherwise.

### GetClusteridOk

`func (o *ArrayRest) GetClusteridOk() (*int64, bool)`

GetClusteridOk returns a tuple with the Clusterid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterid

`func (o *ArrayRest) SetClusterid(v int64)`

SetClusterid sets Clusterid field to given value.

### HasClusterid

`func (o *ArrayRest) HasClusterid() bool`

HasClusterid returns a boolean if a field has been set.

### GetModifydate

`func (o *ArrayRest) GetModifydate() int64`

GetModifydate returns the Modifydate field if non-nil, zero value otherwise.

### GetModifydateOk

`func (o *ArrayRest) GetModifydateOk() (*int64, bool)`

GetModifydateOk returns a tuple with the Modifydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifydate

`func (o *ArrayRest) SetModifydate(v int64)`

SetModifydate sets Modifydate field to given value.

### HasModifydate

`func (o *ArrayRest) HasModifydate() bool`

HasModifydate returns a boolean if a field has been set.

### GetUsername

`func (o *ArrayRest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ArrayRest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ArrayRest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ArrayRest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetIpaddress

`func (o *ArrayRest) GetIpaddress() string`

GetIpaddress returns the Ipaddress field if non-nil, zero value otherwise.

### GetIpaddressOk

`func (o *ArrayRest) GetIpaddressOk() (*string, bool)`

GetIpaddressOk returns a tuple with the Ipaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddress

`func (o *ArrayRest) SetIpaddress(v string)`

SetIpaddress sets Ipaddress field to given value.

### HasIpaddress

`func (o *ArrayRest) HasIpaddress() bool`

HasIpaddress returns a boolean if a field has been set.

### GetStatus

`func (o *ArrayRest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArrayRest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArrayRest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ArrayRest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOverallstatus

`func (o *ArrayRest) GetOverallstatus() string`

GetOverallstatus returns the Overallstatus field if non-nil, zero value otherwise.

### GetOverallstatusOk

`func (o *ArrayRest) GetOverallstatusOk() (*string, bool)`

GetOverallstatusOk returns a tuple with the Overallstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallstatus

`func (o *ArrayRest) SetOverallstatus(v string)`

SetOverallstatus sets Overallstatus field to given value.

### HasOverallstatus

`func (o *ArrayRest) HasOverallstatus() bool`

HasOverallstatus returns a boolean if a field has been set.

### GetModel

`func (o *ArrayRest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ArrayRest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ArrayRest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ArrayRest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetArraytype

`func (o *ArrayRest) GetArraytype() string`

GetArraytype returns the Arraytype field if non-nil, zero value otherwise.

### GetArraytypeOk

`func (o *ArrayRest) GetArraytypeOk() (*string, bool)`

GetArraytypeOk returns a tuple with the Arraytype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArraytype

`func (o *ArrayRest) SetArraytype(v string)`

SetArraytype sets Arraytype field to given value.

### HasArraytype

`func (o *ArrayRest) HasArraytype() bool`

HasArraytype returns a boolean if a field has been set.

### GetOrglist

`func (o *ArrayRest) GetOrglist() []OrganizationRest`

GetOrglist returns the Orglist field if non-nil, zero value otherwise.

### GetOrglistOk

`func (o *ArrayRest) GetOrglistOk() (*[]OrganizationRest, bool)`

GetOrglistOk returns a tuple with the Orglist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrglist

`func (o *ArrayRest) SetOrglist(v []OrganizationRest)`

SetOrglist sets Orglist field to given value.

### HasOrglist

`func (o *ArrayRest) HasOrglist() bool`

HasOrglist returns a boolean if a field has been set.

### GetStorage

`func (o *ArrayRest) GetStorage() []MdiskGroupRest`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ArrayRest) GetStorageOk() (*[]MdiskGroupRest, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ArrayRest) SetStorage(v []MdiskGroupRest)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ArrayRest) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetArraytypelabel

`func (o *ArrayRest) GetArraytypelabel() string`

GetArraytypelabel returns the Arraytypelabel field if non-nil, zero value otherwise.

### GetArraytypelabelOk

`func (o *ArrayRest) GetArraytypelabelOk() (*string, bool)`

GetArraytypelabelOk returns a tuple with the Arraytypelabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArraytypelabel

`func (o *ArrayRest) SetArraytypelabel(v string)`

SetArraytypelabel sets Arraytypelabel field to given value.

### HasArraytypelabel

`func (o *ArrayRest) HasArraytypelabel() bool`

HasArraytypelabel returns a boolean if a field has been set.

### GetHostcount

`func (o *ArrayRest) GetHostcount() int32`

GetHostcount returns the Hostcount field if non-nil, zero value otherwise.

### GetHostcountOk

`func (o *ArrayRest) GetHostcountOk() (*int32, bool)`

GetHostcountOk returns a tuple with the Hostcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostcount

`func (o *ArrayRest) SetHostcount(v int32)`

SetHostcount sets Hostcount field to given value.

### HasHostcount

`func (o *ArrayRest) HasHostcount() bool`

HasHostcount returns a boolean if a field has been set.

### GetDiskpools

`func (o *ArrayRest) GetDiskpools() []DiskPoolRest`

GetDiskpools returns the Diskpools field if non-nil, zero value otherwise.

### GetDiskpoolsOk

`func (o *ArrayRest) GetDiskpoolsOk() (*[]DiskPoolRest, bool)`

GetDiskpoolsOk returns a tuple with the Diskpools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskpools

`func (o *ArrayRest) SetDiskpools(v []DiskPoolRest)`

SetDiskpools sets Diskpools field to given value.

### HasDiskpools

`func (o *ArrayRest) HasDiskpools() bool`

HasDiskpools returns a boolean if a field has been set.

### GetReset

`func (o *ArrayRest) GetReset() bool`

GetReset returns the Reset field if non-nil, zero value otherwise.

### GetResetOk

`func (o *ArrayRest) GetResetOk() (*bool, bool)`

GetResetOk returns a tuple with the Reset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReset

`func (o *ArrayRest) SetReset(v bool)`

SetReset sets Reset field to given value.

### HasReset

`func (o *ArrayRest) HasReset() bool`

HasReset returns a boolean if a field has been set.

### GetAppliance

`func (o *ArrayRest) GetAppliance() ClusterRest`

GetAppliance returns the Appliance field if non-nil, zero value otherwise.

### GetApplianceOk

`func (o *ArrayRest) GetApplianceOk() (*ClusterRest, bool)`

GetApplianceOk returns a tuple with the Appliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliance

`func (o *ArrayRest) SetAppliance(v ClusterRest)`

SetAppliance sets Appliance field to given value.

### HasAppliance

`func (o *ArrayRest) HasAppliance() bool`

HasAppliance returns a boolean if a field has been set.

### GetId

`func (o *ArrayRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArrayRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArrayRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArrayRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ArrayRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ArrayRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ArrayRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ArrayRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ArrayRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ArrayRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ArrayRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ArrayRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ArrayRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ArrayRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ArrayRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ArrayRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



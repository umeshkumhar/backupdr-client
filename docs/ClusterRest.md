# ClusterRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedSecret** | Pointer to **string** |  | [optional] 
**Serviceaccount** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Projectid** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **int32** |  | [optional] 
**Datacenter** | Pointer to **string** |  | [optional] 
**Masterid** | Pointer to **string** |  | [optional] 
**Clusterid** | Pointer to **string** |  | [optional] 
**Importstatus** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Ipaddress** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**Clusterstatus** | Pointer to [**ClusterStatusRest**](ClusterStatusRest.md) |  | [optional] 
**Lastsync** | Pointer to **int64** |  | [optional] 
**Publicip** | Pointer to **string** |  | [optional] 
**Secureconnect** | Pointer to **bool** |  | [optional] 
**PkiBootstrapped** | Pointer to **bool** |  | [optional] 
**Clusterlist** | Pointer to [**[]ClusterRest**](ClusterRest.md) |  | [optional] 
**CallhomeInfo** | Pointer to [**CallhomeInfoRest**](CallhomeInfoRest.md) |  | [optional] 
**Rmipaddress** | Pointer to **[]string** |  | [optional] 
**Supportstatus** | Pointer to **string** | Shows the support status of cluster | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewClusterRest

`func NewClusterRest() *ClusterRest`

NewClusterRest instantiates a new ClusterRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRestWithDefaults

`func NewClusterRestWithDefaults() *ClusterRest`

NewClusterRestWithDefaults instantiates a new ClusterRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedSecret

`func (o *ClusterRest) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *ClusterRest) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *ClusterRest) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *ClusterRest) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetServiceaccount

`func (o *ClusterRest) GetServiceaccount() string`

GetServiceaccount returns the Serviceaccount field if non-nil, zero value otherwise.

### GetServiceaccountOk

`func (o *ClusterRest) GetServiceaccountOk() (*string, bool)`

GetServiceaccountOk returns a tuple with the Serviceaccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceaccount

`func (o *ClusterRest) SetServiceaccount(v string)`

SetServiceaccount sets Serviceaccount field to given value.

### HasServiceaccount

`func (o *ClusterRest) HasServiceaccount() bool`

HasServiceaccount returns a boolean if a field has been set.

### GetZone

`func (o *ClusterRest) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ClusterRest) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ClusterRest) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ClusterRest) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetRegion

`func (o *ClusterRest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ClusterRest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ClusterRest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ClusterRest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetProjectid

`func (o *ClusterRest) GetProjectid() string`

GetProjectid returns the Projectid field if non-nil, zero value otherwise.

### GetProjectidOk

`func (o *ClusterRest) GetProjectidOk() (*string, bool)`

GetProjectidOk returns a tuple with the Projectid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectid

`func (o *ClusterRest) SetProjectid(v string)`

SetProjectid sets Projectid field to given value.

### HasProjectid

`func (o *ClusterRest) HasProjectid() bool`

HasProjectid returns a boolean if a field has been set.

### GetPassword

`func (o *ClusterRest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ClusterRest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ClusterRest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ClusterRest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetVersion

`func (o *ClusterRest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClusterRest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClusterRest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ClusterRest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDescription

`func (o *ClusterRest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterRest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterRest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterRest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ClusterRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ClusterRest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterRest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterRest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterRest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProtocol

`func (o *ClusterRest) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ClusterRest) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ClusterRest) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ClusterRest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetDatacenter

`func (o *ClusterRest) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *ClusterRest) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *ClusterRest) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *ClusterRest) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetMasterid

`func (o *ClusterRest) GetMasterid() string`

GetMasterid returns the Masterid field if non-nil, zero value otherwise.

### GetMasteridOk

`func (o *ClusterRest) GetMasteridOk() (*string, bool)`

GetMasteridOk returns a tuple with the Masterid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterid

`func (o *ClusterRest) SetMasterid(v string)`

SetMasterid sets Masterid field to given value.

### HasMasterid

`func (o *ClusterRest) HasMasterid() bool`

HasMasterid returns a boolean if a field has been set.

### GetClusterid

`func (o *ClusterRest) GetClusterid() string`

GetClusterid returns the Clusterid field if non-nil, zero value otherwise.

### GetClusteridOk

`func (o *ClusterRest) GetClusteridOk() (*string, bool)`

GetClusteridOk returns a tuple with the Clusterid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterid

`func (o *ClusterRest) SetClusterid(v string)`

SetClusterid sets Clusterid field to given value.

### HasClusterid

`func (o *ClusterRest) HasClusterid() bool`

HasClusterid returns a boolean if a field has been set.

### GetImportstatus

`func (o *ClusterRest) GetImportstatus() string`

GetImportstatus returns the Importstatus field if non-nil, zero value otherwise.

### GetImportstatusOk

`func (o *ClusterRest) GetImportstatusOk() (*string, bool)`

GetImportstatusOk returns a tuple with the Importstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportstatus

`func (o *ClusterRest) SetImportstatus(v string)`

SetImportstatus sets Importstatus field to given value.

### HasImportstatus

`func (o *ClusterRest) HasImportstatus() bool`

HasImportstatus returns a boolean if a field has been set.

### GetUsername

`func (o *ClusterRest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ClusterRest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ClusterRest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ClusterRest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetIpaddress

`func (o *ClusterRest) GetIpaddress() string`

GetIpaddress returns the Ipaddress field if non-nil, zero value otherwise.

### GetIpaddressOk

`func (o *ClusterRest) GetIpaddressOk() (*string, bool)`

GetIpaddressOk returns a tuple with the Ipaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddress

`func (o *ClusterRest) SetIpaddress(v string)`

SetIpaddress sets Ipaddress field to given value.

### HasIpaddress

`func (o *ClusterRest) HasIpaddress() bool`

HasIpaddress returns a boolean if a field has been set.

### GetTimezone

`func (o *ClusterRest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ClusterRest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ClusterRest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ClusterRest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetClusterstatus

`func (o *ClusterRest) GetClusterstatus() ClusterStatusRest`

GetClusterstatus returns the Clusterstatus field if non-nil, zero value otherwise.

### GetClusterstatusOk

`func (o *ClusterRest) GetClusterstatusOk() (*ClusterStatusRest, bool)`

GetClusterstatusOk returns a tuple with the Clusterstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterstatus

`func (o *ClusterRest) SetClusterstatus(v ClusterStatusRest)`

SetClusterstatus sets Clusterstatus field to given value.

### HasClusterstatus

`func (o *ClusterRest) HasClusterstatus() bool`

HasClusterstatus returns a boolean if a field has been set.

### GetLastsync

`func (o *ClusterRest) GetLastsync() int64`

GetLastsync returns the Lastsync field if non-nil, zero value otherwise.

### GetLastsyncOk

`func (o *ClusterRest) GetLastsyncOk() (*int64, bool)`

GetLastsyncOk returns a tuple with the Lastsync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastsync

`func (o *ClusterRest) SetLastsync(v int64)`

SetLastsync sets Lastsync field to given value.

### HasLastsync

`func (o *ClusterRest) HasLastsync() bool`

HasLastsync returns a boolean if a field has been set.

### GetPublicip

`func (o *ClusterRest) GetPublicip() string`

GetPublicip returns the Publicip field if non-nil, zero value otherwise.

### GetPublicipOk

`func (o *ClusterRest) GetPublicipOk() (*string, bool)`

GetPublicipOk returns a tuple with the Publicip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicip

`func (o *ClusterRest) SetPublicip(v string)`

SetPublicip sets Publicip field to given value.

### HasPublicip

`func (o *ClusterRest) HasPublicip() bool`

HasPublicip returns a boolean if a field has been set.

### GetSecureconnect

`func (o *ClusterRest) GetSecureconnect() bool`

GetSecureconnect returns the Secureconnect field if non-nil, zero value otherwise.

### GetSecureconnectOk

`func (o *ClusterRest) GetSecureconnectOk() (*bool, bool)`

GetSecureconnectOk returns a tuple with the Secureconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureconnect

`func (o *ClusterRest) SetSecureconnect(v bool)`

SetSecureconnect sets Secureconnect field to given value.

### HasSecureconnect

`func (o *ClusterRest) HasSecureconnect() bool`

HasSecureconnect returns a boolean if a field has been set.

### GetPkiBootstrapped

`func (o *ClusterRest) GetPkiBootstrapped() bool`

GetPkiBootstrapped returns the PkiBootstrapped field if non-nil, zero value otherwise.

### GetPkiBootstrappedOk

`func (o *ClusterRest) GetPkiBootstrappedOk() (*bool, bool)`

GetPkiBootstrappedOk returns a tuple with the PkiBootstrapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiBootstrapped

`func (o *ClusterRest) SetPkiBootstrapped(v bool)`

SetPkiBootstrapped sets PkiBootstrapped field to given value.

### HasPkiBootstrapped

`func (o *ClusterRest) HasPkiBootstrapped() bool`

HasPkiBootstrapped returns a boolean if a field has been set.

### GetClusterlist

`func (o *ClusterRest) GetClusterlist() []ClusterRest`

GetClusterlist returns the Clusterlist field if non-nil, zero value otherwise.

### GetClusterlistOk

`func (o *ClusterRest) GetClusterlistOk() (*[]ClusterRest, bool)`

GetClusterlistOk returns a tuple with the Clusterlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterlist

`func (o *ClusterRest) SetClusterlist(v []ClusterRest)`

SetClusterlist sets Clusterlist field to given value.

### HasClusterlist

`func (o *ClusterRest) HasClusterlist() bool`

HasClusterlist returns a boolean if a field has been set.

### GetCallhomeInfo

`func (o *ClusterRest) GetCallhomeInfo() CallhomeInfoRest`

GetCallhomeInfo returns the CallhomeInfo field if non-nil, zero value otherwise.

### GetCallhomeInfoOk

`func (o *ClusterRest) GetCallhomeInfoOk() (*CallhomeInfoRest, bool)`

GetCallhomeInfoOk returns a tuple with the CallhomeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallhomeInfo

`func (o *ClusterRest) SetCallhomeInfo(v CallhomeInfoRest)`

SetCallhomeInfo sets CallhomeInfo field to given value.

### HasCallhomeInfo

`func (o *ClusterRest) HasCallhomeInfo() bool`

HasCallhomeInfo returns a boolean if a field has been set.

### GetRmipaddress

`func (o *ClusterRest) GetRmipaddress() []string`

GetRmipaddress returns the Rmipaddress field if non-nil, zero value otherwise.

### GetRmipaddressOk

`func (o *ClusterRest) GetRmipaddressOk() (*[]string, bool)`

GetRmipaddressOk returns a tuple with the Rmipaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmipaddress

`func (o *ClusterRest) SetRmipaddress(v []string)`

SetRmipaddress sets Rmipaddress field to given value.

### HasRmipaddress

`func (o *ClusterRest) HasRmipaddress() bool`

HasRmipaddress returns a boolean if a field has been set.

### GetSupportstatus

`func (o *ClusterRest) GetSupportstatus() string`

GetSupportstatus returns the Supportstatus field if non-nil, zero value otherwise.

### GetSupportstatusOk

`func (o *ClusterRest) GetSupportstatusOk() (*string, bool)`

GetSupportstatusOk returns a tuple with the Supportstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportstatus

`func (o *ClusterRest) SetSupportstatus(v string)`

SetSupportstatus sets Supportstatus field to given value.

### HasSupportstatus

`func (o *ClusterRest) HasSupportstatus() bool`

HasSupportstatus returns a boolean if a field has been set.

### GetId

`func (o *ClusterRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ClusterRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ClusterRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ClusterRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ClusterRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ClusterRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ClusterRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ClusterRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ClusterRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ClusterRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ClusterRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ClusterRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ClusterRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



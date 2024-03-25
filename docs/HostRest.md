# HostRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertRevoked** | Pointer to **bool** | Indicates if a host&#39;s CA issued certificate has been revoked or not. | [optional] 
**PkiState** | Pointer to **string** | Indicates if the host is enabled to use CA signed certificate. Valid states are TRUSTED (has valid CA issued certificate for communication), UNTRUSTED (doesn&#39;t have valid CA issued certificate for communication) and  NOT_APPLICABLE (doesn&#39;t needs CA issued certificate for communication). | [optional] 
**PkiErrors** | Pointer to **[]string** | List of error messages while trying to bootstrap PKI on the host using managing backup appliance | [optional] 
**Autoupgrade** | Pointer to **string** |  | [optional] 
**Immutable** | Pointer to **bool** |  | [optional] 
**Mask** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**[]SourceRest**](SourceRest.md) |  | [optional] 
**Zone** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Sources** | Pointer to [**[]HostRest**](HostRest.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Srcid** | Pointer to **string** |  | [optional] 
**Clusterid** | Pointer to **string** |  | [optional] 
**Hosttype** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Uniquename** | Pointer to **string** |  | [optional] 
**Isvm** | Pointer to **bool** |  | [optional] 
**Vmtype** | Pointer to **string** |  | [optional] 
**Vcenterhostid** | Pointer to **string** |  | [optional] 
**Modifydate** | Pointer to **int64** |  | [optional] 
**Ipaddress** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Transport** | Pointer to **string** |  | [optional] 
**Friendlypath** | Pointer to **string** |  | [optional] 
**Svcname** | Pointer to **string** |  | [optional] 
**Dataip** | Pointer to **string** |  | [optional] 
**Alternateip** | Pointer to **[]string** |  | [optional] 
**Ostype** | Pointer to **string** |  | [optional] 
**Osversion** | Pointer to **string** |  | [optional] 
**Osrelease** | Pointer to **string** |  | [optional] 
**Sourcecluster** | Pointer to **string** |  | [optional] 
**Originalhostid** | Pointer to **string** |  | [optional] 
**Maxjobs** | Pointer to **int32** |  | [optional] 
**Hasagent** | Pointer to **bool** |  | [optional] 
**Isvcenterhost** | Pointer to **bool** |  | [optional] 
**Isclusterhost** | Pointer to **bool** |  | [optional] 
**Isesxhost** | Pointer to **bool** |  | [optional] 
**Isproxyhost** | Pointer to **bool** |  | [optional] 
**Dbauthentication** | Pointer to **bool** |  | [optional] 
**Diskpref** | Pointer to **string** |  | [optional] 
**Nfsoption** | Pointer to [**NfsOptionsRest**](NfsOptionsRest.md) |  | [optional] 
**Multiregion** | Pointer to **string** |  | [optional] 
**Machinetype** | Pointer to **string** |  | [optional] 
**Systemdetail** | Pointer to **string** |  | [optional] 
**Connect2actip** | Pointer to **string** |  | [optional] 
**IsAutoDiscoveryEnabled** | Pointer to **bool** |  | [optional] 
**IsShadowHost** | Pointer to **bool** |  | [optional] 
**Orglist** | Pointer to [**[]OrganizationRest**](OrganizationRest.md) |  | [optional] 
**Esxlist** | Pointer to [**[]HostRest**](HostRest.md) |  | [optional] 
**NodeList** | Pointer to [**[]HostRest**](HostRest.md) |  | [optional] 
**Datastorelist** | Pointer to [**[]DataStoreRest**](DataStoreRest.md) |  | [optional] 
**OstypeSpecial** | Pointer to **string** |  | [optional] 
**PortCount** | Pointer to **string** |  | [optional] 
**IogrpCount** | Pointer to **string** |  | [optional] 
**IscsiName** | Pointer to **[]string** |  | [optional] 
**WWPN** | Pointer to **[]string** |  | [optional] 
**Iogrp** | Pointer to **string** |  | [optional] 
**NodeLoggedInCount** | Pointer to **string** |  | [optional] 
**Connectorversion** | Pointer to **string** |  | [optional] 
**IsClusterNode** | Pointer to **bool** |  | [optional] 
**ReconciliationNeeded** | Pointer to **bool** |  | [optional] 
**ReconciliationFields** | Pointer to **[]string** |  | [optional] 
**Agents** | Pointer to [**[]AgentRest**](AgentRest.md) |  | [optional] 
**Arrays** | Pointer to [**[]ArrayRest**](ArrayRest.md) |  | [optional] 
**VcenterHost** | Pointer to [**HostRest**](HostRest.md) |  | [optional] 
**Nasdconfig** | Pointer to **map[string]interface{}** |  | [optional] 
**Cloudcredential** | Pointer to [**CloudCredentialRest**](CloudCredentialRest.md) |  | [optional] 
**Chapusername** | Pointer to **string** |  | [optional] 
**Chappassword** | Pointer to **string** |  | [optional] 
**Clearchap** | Pointer to **bool** |  | [optional] 
**Udsagent** | Pointer to [**AgentRest**](AgentRest.md) |  | [optional] 
**Applicationagent** | Pointer to [**AgentRest**](AgentRest.md) |  | [optional] 
**Hypervisoragent** | Pointer to [**AgentRest**](AgentRest.md) |  | [optional] 
**Hmcconnector** | Pointer to [**AgentRest**](AgentRest.md) |  | [optional] 
**Appliance** | Pointer to [**ClusterRest**](ClusterRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewHostRest

`func NewHostRest() *HostRest`

NewHostRest instantiates a new HostRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostRestWithDefaults

`func NewHostRestWithDefaults() *HostRest`

NewHostRestWithDefaults instantiates a new HostRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertRevoked

`func (o *HostRest) GetCertRevoked() bool`

GetCertRevoked returns the CertRevoked field if non-nil, zero value otherwise.

### GetCertRevokedOk

`func (o *HostRest) GetCertRevokedOk() (*bool, bool)`

GetCertRevokedOk returns a tuple with the CertRevoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertRevoked

`func (o *HostRest) SetCertRevoked(v bool)`

SetCertRevoked sets CertRevoked field to given value.

### HasCertRevoked

`func (o *HostRest) HasCertRevoked() bool`

HasCertRevoked returns a boolean if a field has been set.

### GetPkiState

`func (o *HostRest) GetPkiState() string`

GetPkiState returns the PkiState field if non-nil, zero value otherwise.

### GetPkiStateOk

`func (o *HostRest) GetPkiStateOk() (*string, bool)`

GetPkiStateOk returns a tuple with the PkiState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiState

`func (o *HostRest) SetPkiState(v string)`

SetPkiState sets PkiState field to given value.

### HasPkiState

`func (o *HostRest) HasPkiState() bool`

HasPkiState returns a boolean if a field has been set.

### GetPkiErrors

`func (o *HostRest) GetPkiErrors() []string`

GetPkiErrors returns the PkiErrors field if non-nil, zero value otherwise.

### GetPkiErrorsOk

`func (o *HostRest) GetPkiErrorsOk() (*[]string, bool)`

GetPkiErrorsOk returns a tuple with the PkiErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiErrors

`func (o *HostRest) SetPkiErrors(v []string)`

SetPkiErrors sets PkiErrors field to given value.

### HasPkiErrors

`func (o *HostRest) HasPkiErrors() bool`

HasPkiErrors returns a boolean if a field has been set.

### GetAutoupgrade

`func (o *HostRest) GetAutoupgrade() string`

GetAutoupgrade returns the Autoupgrade field if non-nil, zero value otherwise.

### GetAutoupgradeOk

`func (o *HostRest) GetAutoupgradeOk() (*string, bool)`

GetAutoupgradeOk returns a tuple with the Autoupgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoupgrade

`func (o *HostRest) SetAutoupgrade(v string)`

SetAutoupgrade sets Autoupgrade field to given value.

### HasAutoupgrade

`func (o *HostRest) HasAutoupgrade() bool`

HasAutoupgrade returns a boolean if a field has been set.

### GetImmutable

`func (o *HostRest) GetImmutable() bool`

GetImmutable returns the Immutable field if non-nil, zero value otherwise.

### GetImmutableOk

`func (o *HostRest) GetImmutableOk() (*bool, bool)`

GetImmutableOk returns a tuple with the Immutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutable

`func (o *HostRest) SetImmutable(v bool)`

SetImmutable sets Immutable field to given value.

### HasImmutable

`func (o *HostRest) HasImmutable() bool`

HasImmutable returns a boolean if a field has been set.

### GetMask

`func (o *HostRest) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *HostRest) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *HostRest) SetMask(v string)`

SetMask sets Mask field to given value.

### HasMask

`func (o *HostRest) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetSource

`func (o *HostRest) GetSource() []SourceRest`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *HostRest) GetSourceOk() (*[]SourceRest, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *HostRest) SetSource(v []SourceRest)`

SetSource sets Source field to given value.

### HasSource

`func (o *HostRest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetZone

`func (o *HostRest) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *HostRest) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *HostRest) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *HostRest) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetDescription

`func (o *HostRest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HostRest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HostRest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HostRest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSources

`func (o *HostRest) GetSources() []HostRest`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *HostRest) GetSourcesOk() (*[]HostRest, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *HostRest) SetSources(v []HostRest)`

SetSources sets Sources field to given value.

### HasSources

`func (o *HostRest) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetName

`func (o *HostRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HostRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HostRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HostRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *HostRest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HostRest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HostRest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HostRest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSrcid

`func (o *HostRest) GetSrcid() string`

GetSrcid returns the Srcid field if non-nil, zero value otherwise.

### GetSrcidOk

`func (o *HostRest) GetSrcidOk() (*string, bool)`

GetSrcidOk returns a tuple with the Srcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcid

`func (o *HostRest) SetSrcid(v string)`

SetSrcid sets Srcid field to given value.

### HasSrcid

`func (o *HostRest) HasSrcid() bool`

HasSrcid returns a boolean if a field has been set.

### GetClusterid

`func (o *HostRest) GetClusterid() string`

GetClusterid returns the Clusterid field if non-nil, zero value otherwise.

### GetClusteridOk

`func (o *HostRest) GetClusteridOk() (*string, bool)`

GetClusteridOk returns a tuple with the Clusterid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterid

`func (o *HostRest) SetClusterid(v string)`

SetClusterid sets Clusterid field to given value.

### HasClusterid

`func (o *HostRest) HasClusterid() bool`

HasClusterid returns a boolean if a field has been set.

### GetHosttype

`func (o *HostRest) GetHosttype() string`

GetHosttype returns the Hosttype field if non-nil, zero value otherwise.

### GetHosttypeOk

`func (o *HostRest) GetHosttypeOk() (*string, bool)`

GetHosttypeOk returns a tuple with the Hosttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosttype

`func (o *HostRest) SetHosttype(v string)`

SetHosttype sets Hosttype field to given value.

### HasHosttype

`func (o *HostRest) HasHosttype() bool`

HasHosttype returns a boolean if a field has been set.

### GetHostname

`func (o *HostRest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *HostRest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *HostRest) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *HostRest) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetUniquename

`func (o *HostRest) GetUniquename() string`

GetUniquename returns the Uniquename field if non-nil, zero value otherwise.

### GetUniquenameOk

`func (o *HostRest) GetUniquenameOk() (*string, bool)`

GetUniquenameOk returns a tuple with the Uniquename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniquename

`func (o *HostRest) SetUniquename(v string)`

SetUniquename sets Uniquename field to given value.

### HasUniquename

`func (o *HostRest) HasUniquename() bool`

HasUniquename returns a boolean if a field has been set.

### GetIsvm

`func (o *HostRest) GetIsvm() bool`

GetIsvm returns the Isvm field if non-nil, zero value otherwise.

### GetIsvmOk

`func (o *HostRest) GetIsvmOk() (*bool, bool)`

GetIsvmOk returns a tuple with the Isvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsvm

`func (o *HostRest) SetIsvm(v bool)`

SetIsvm sets Isvm field to given value.

### HasIsvm

`func (o *HostRest) HasIsvm() bool`

HasIsvm returns a boolean if a field has been set.

### GetVmtype

`func (o *HostRest) GetVmtype() string`

GetVmtype returns the Vmtype field if non-nil, zero value otherwise.

### GetVmtypeOk

`func (o *HostRest) GetVmtypeOk() (*string, bool)`

GetVmtypeOk returns a tuple with the Vmtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmtype

`func (o *HostRest) SetVmtype(v string)`

SetVmtype sets Vmtype field to given value.

### HasVmtype

`func (o *HostRest) HasVmtype() bool`

HasVmtype returns a boolean if a field has been set.

### GetVcenterhostid

`func (o *HostRest) GetVcenterhostid() string`

GetVcenterhostid returns the Vcenterhostid field if non-nil, zero value otherwise.

### GetVcenterhostidOk

`func (o *HostRest) GetVcenterhostidOk() (*string, bool)`

GetVcenterhostidOk returns a tuple with the Vcenterhostid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterhostid

`func (o *HostRest) SetVcenterhostid(v string)`

SetVcenterhostid sets Vcenterhostid field to given value.

### HasVcenterhostid

`func (o *HostRest) HasVcenterhostid() bool`

HasVcenterhostid returns a boolean if a field has been set.

### GetModifydate

`func (o *HostRest) GetModifydate() int64`

GetModifydate returns the Modifydate field if non-nil, zero value otherwise.

### GetModifydateOk

`func (o *HostRest) GetModifydateOk() (*int64, bool)`

GetModifydateOk returns a tuple with the Modifydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifydate

`func (o *HostRest) SetModifydate(v int64)`

SetModifydate sets Modifydate field to given value.

### HasModifydate

`func (o *HostRest) HasModifydate() bool`

HasModifydate returns a boolean if a field has been set.

### GetIpaddress

`func (o *HostRest) GetIpaddress() string`

GetIpaddress returns the Ipaddress field if non-nil, zero value otherwise.

### GetIpaddressOk

`func (o *HostRest) GetIpaddressOk() (*string, bool)`

GetIpaddressOk returns a tuple with the Ipaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddress

`func (o *HostRest) SetIpaddress(v string)`

SetIpaddress sets Ipaddress field to given value.

### HasIpaddress

`func (o *HostRest) HasIpaddress() bool`

HasIpaddress returns a boolean if a field has been set.

### GetStatus

`func (o *HostRest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HostRest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HostRest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HostRest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransport

`func (o *HostRest) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *HostRest) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *HostRest) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *HostRest) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetFriendlypath

`func (o *HostRest) GetFriendlypath() string`

GetFriendlypath returns the Friendlypath field if non-nil, zero value otherwise.

### GetFriendlypathOk

`func (o *HostRest) GetFriendlypathOk() (*string, bool)`

GetFriendlypathOk returns a tuple with the Friendlypath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlypath

`func (o *HostRest) SetFriendlypath(v string)`

SetFriendlypath sets Friendlypath field to given value.

### HasFriendlypath

`func (o *HostRest) HasFriendlypath() bool`

HasFriendlypath returns a boolean if a field has been set.

### GetSvcname

`func (o *HostRest) GetSvcname() string`

GetSvcname returns the Svcname field if non-nil, zero value otherwise.

### GetSvcnameOk

`func (o *HostRest) GetSvcnameOk() (*string, bool)`

GetSvcnameOk returns a tuple with the Svcname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcname

`func (o *HostRest) SetSvcname(v string)`

SetSvcname sets Svcname field to given value.

### HasSvcname

`func (o *HostRest) HasSvcname() bool`

HasSvcname returns a boolean if a field has been set.

### GetDataip

`func (o *HostRest) GetDataip() string`

GetDataip returns the Dataip field if non-nil, zero value otherwise.

### GetDataipOk

`func (o *HostRest) GetDataipOk() (*string, bool)`

GetDataipOk returns a tuple with the Dataip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataip

`func (o *HostRest) SetDataip(v string)`

SetDataip sets Dataip field to given value.

### HasDataip

`func (o *HostRest) HasDataip() bool`

HasDataip returns a boolean if a field has been set.

### GetAlternateip

`func (o *HostRest) GetAlternateip() []string`

GetAlternateip returns the Alternateip field if non-nil, zero value otherwise.

### GetAlternateipOk

`func (o *HostRest) GetAlternateipOk() (*[]string, bool)`

GetAlternateipOk returns a tuple with the Alternateip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateip

`func (o *HostRest) SetAlternateip(v []string)`

SetAlternateip sets Alternateip field to given value.

### HasAlternateip

`func (o *HostRest) HasAlternateip() bool`

HasAlternateip returns a boolean if a field has been set.

### GetOstype

`func (o *HostRest) GetOstype() string`

GetOstype returns the Ostype field if non-nil, zero value otherwise.

### GetOstypeOk

`func (o *HostRest) GetOstypeOk() (*string, bool)`

GetOstypeOk returns a tuple with the Ostype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOstype

`func (o *HostRest) SetOstype(v string)`

SetOstype sets Ostype field to given value.

### HasOstype

`func (o *HostRest) HasOstype() bool`

HasOstype returns a boolean if a field has been set.

### GetOsversion

`func (o *HostRest) GetOsversion() string`

GetOsversion returns the Osversion field if non-nil, zero value otherwise.

### GetOsversionOk

`func (o *HostRest) GetOsversionOk() (*string, bool)`

GetOsversionOk returns a tuple with the Osversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsversion

`func (o *HostRest) SetOsversion(v string)`

SetOsversion sets Osversion field to given value.

### HasOsversion

`func (o *HostRest) HasOsversion() bool`

HasOsversion returns a boolean if a field has been set.

### GetOsrelease

`func (o *HostRest) GetOsrelease() string`

GetOsrelease returns the Osrelease field if non-nil, zero value otherwise.

### GetOsreleaseOk

`func (o *HostRest) GetOsreleaseOk() (*string, bool)`

GetOsreleaseOk returns a tuple with the Osrelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsrelease

`func (o *HostRest) SetOsrelease(v string)`

SetOsrelease sets Osrelease field to given value.

### HasOsrelease

`func (o *HostRest) HasOsrelease() bool`

HasOsrelease returns a boolean if a field has been set.

### GetSourcecluster

`func (o *HostRest) GetSourcecluster() string`

GetSourcecluster returns the Sourcecluster field if non-nil, zero value otherwise.

### GetSourceclusterOk

`func (o *HostRest) GetSourceclusterOk() (*string, bool)`

GetSourceclusterOk returns a tuple with the Sourcecluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcecluster

`func (o *HostRest) SetSourcecluster(v string)`

SetSourcecluster sets Sourcecluster field to given value.

### HasSourcecluster

`func (o *HostRest) HasSourcecluster() bool`

HasSourcecluster returns a boolean if a field has been set.

### GetOriginalhostid

`func (o *HostRest) GetOriginalhostid() string`

GetOriginalhostid returns the Originalhostid field if non-nil, zero value otherwise.

### GetOriginalhostidOk

`func (o *HostRest) GetOriginalhostidOk() (*string, bool)`

GetOriginalhostidOk returns a tuple with the Originalhostid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalhostid

`func (o *HostRest) SetOriginalhostid(v string)`

SetOriginalhostid sets Originalhostid field to given value.

### HasOriginalhostid

`func (o *HostRest) HasOriginalhostid() bool`

HasOriginalhostid returns a boolean if a field has been set.

### GetMaxjobs

`func (o *HostRest) GetMaxjobs() int32`

GetMaxjobs returns the Maxjobs field if non-nil, zero value otherwise.

### GetMaxjobsOk

`func (o *HostRest) GetMaxjobsOk() (*int32, bool)`

GetMaxjobsOk returns a tuple with the Maxjobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxjobs

`func (o *HostRest) SetMaxjobs(v int32)`

SetMaxjobs sets Maxjobs field to given value.

### HasMaxjobs

`func (o *HostRest) HasMaxjobs() bool`

HasMaxjobs returns a boolean if a field has been set.

### GetHasagent

`func (o *HostRest) GetHasagent() bool`

GetHasagent returns the Hasagent field if non-nil, zero value otherwise.

### GetHasagentOk

`func (o *HostRest) GetHasagentOk() (*bool, bool)`

GetHasagentOk returns a tuple with the Hasagent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasagent

`func (o *HostRest) SetHasagent(v bool)`

SetHasagent sets Hasagent field to given value.

### HasHasagent

`func (o *HostRest) HasHasagent() bool`

HasHasagent returns a boolean if a field has been set.

### GetIsvcenterhost

`func (o *HostRest) GetIsvcenterhost() bool`

GetIsvcenterhost returns the Isvcenterhost field if non-nil, zero value otherwise.

### GetIsvcenterhostOk

`func (o *HostRest) GetIsvcenterhostOk() (*bool, bool)`

GetIsvcenterhostOk returns a tuple with the Isvcenterhost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsvcenterhost

`func (o *HostRest) SetIsvcenterhost(v bool)`

SetIsvcenterhost sets Isvcenterhost field to given value.

### HasIsvcenterhost

`func (o *HostRest) HasIsvcenterhost() bool`

HasIsvcenterhost returns a boolean if a field has been set.

### GetIsclusterhost

`func (o *HostRest) GetIsclusterhost() bool`

GetIsclusterhost returns the Isclusterhost field if non-nil, zero value otherwise.

### GetIsclusterhostOk

`func (o *HostRest) GetIsclusterhostOk() (*bool, bool)`

GetIsclusterhostOk returns a tuple with the Isclusterhost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsclusterhost

`func (o *HostRest) SetIsclusterhost(v bool)`

SetIsclusterhost sets Isclusterhost field to given value.

### HasIsclusterhost

`func (o *HostRest) HasIsclusterhost() bool`

HasIsclusterhost returns a boolean if a field has been set.

### GetIsesxhost

`func (o *HostRest) GetIsesxhost() bool`

GetIsesxhost returns the Isesxhost field if non-nil, zero value otherwise.

### GetIsesxhostOk

`func (o *HostRest) GetIsesxhostOk() (*bool, bool)`

GetIsesxhostOk returns a tuple with the Isesxhost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsesxhost

`func (o *HostRest) SetIsesxhost(v bool)`

SetIsesxhost sets Isesxhost field to given value.

### HasIsesxhost

`func (o *HostRest) HasIsesxhost() bool`

HasIsesxhost returns a boolean if a field has been set.

### GetIsproxyhost

`func (o *HostRest) GetIsproxyhost() bool`

GetIsproxyhost returns the Isproxyhost field if non-nil, zero value otherwise.

### GetIsproxyhostOk

`func (o *HostRest) GetIsproxyhostOk() (*bool, bool)`

GetIsproxyhostOk returns a tuple with the Isproxyhost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsproxyhost

`func (o *HostRest) SetIsproxyhost(v bool)`

SetIsproxyhost sets Isproxyhost field to given value.

### HasIsproxyhost

`func (o *HostRest) HasIsproxyhost() bool`

HasIsproxyhost returns a boolean if a field has been set.

### GetDbauthentication

`func (o *HostRest) GetDbauthentication() bool`

GetDbauthentication returns the Dbauthentication field if non-nil, zero value otherwise.

### GetDbauthenticationOk

`func (o *HostRest) GetDbauthenticationOk() (*bool, bool)`

GetDbauthenticationOk returns a tuple with the Dbauthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbauthentication

`func (o *HostRest) SetDbauthentication(v bool)`

SetDbauthentication sets Dbauthentication field to given value.

### HasDbauthentication

`func (o *HostRest) HasDbauthentication() bool`

HasDbauthentication returns a boolean if a field has been set.

### GetDiskpref

`func (o *HostRest) GetDiskpref() string`

GetDiskpref returns the Diskpref field if non-nil, zero value otherwise.

### GetDiskprefOk

`func (o *HostRest) GetDiskprefOk() (*string, bool)`

GetDiskprefOk returns a tuple with the Diskpref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskpref

`func (o *HostRest) SetDiskpref(v string)`

SetDiskpref sets Diskpref field to given value.

### HasDiskpref

`func (o *HostRest) HasDiskpref() bool`

HasDiskpref returns a boolean if a field has been set.

### GetNfsoption

`func (o *HostRest) GetNfsoption() NfsOptionsRest`

GetNfsoption returns the Nfsoption field if non-nil, zero value otherwise.

### GetNfsoptionOk

`func (o *HostRest) GetNfsoptionOk() (*NfsOptionsRest, bool)`

GetNfsoptionOk returns a tuple with the Nfsoption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsoption

`func (o *HostRest) SetNfsoption(v NfsOptionsRest)`

SetNfsoption sets Nfsoption field to given value.

### HasNfsoption

`func (o *HostRest) HasNfsoption() bool`

HasNfsoption returns a boolean if a field has been set.

### GetMultiregion

`func (o *HostRest) GetMultiregion() string`

GetMultiregion returns the Multiregion field if non-nil, zero value otherwise.

### GetMultiregionOk

`func (o *HostRest) GetMultiregionOk() (*string, bool)`

GetMultiregionOk returns a tuple with the Multiregion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiregion

`func (o *HostRest) SetMultiregion(v string)`

SetMultiregion sets Multiregion field to given value.

### HasMultiregion

`func (o *HostRest) HasMultiregion() bool`

HasMultiregion returns a boolean if a field has been set.

### GetMachinetype

`func (o *HostRest) GetMachinetype() string`

GetMachinetype returns the Machinetype field if non-nil, zero value otherwise.

### GetMachinetypeOk

`func (o *HostRest) GetMachinetypeOk() (*string, bool)`

GetMachinetypeOk returns a tuple with the Machinetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachinetype

`func (o *HostRest) SetMachinetype(v string)`

SetMachinetype sets Machinetype field to given value.

### HasMachinetype

`func (o *HostRest) HasMachinetype() bool`

HasMachinetype returns a boolean if a field has been set.

### GetSystemdetail

`func (o *HostRest) GetSystemdetail() string`

GetSystemdetail returns the Systemdetail field if non-nil, zero value otherwise.

### GetSystemdetailOk

`func (o *HostRest) GetSystemdetailOk() (*string, bool)`

GetSystemdetailOk returns a tuple with the Systemdetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemdetail

`func (o *HostRest) SetSystemdetail(v string)`

SetSystemdetail sets Systemdetail field to given value.

### HasSystemdetail

`func (o *HostRest) HasSystemdetail() bool`

HasSystemdetail returns a boolean if a field has been set.

### GetConnect2actip

`func (o *HostRest) GetConnect2actip() string`

GetConnect2actip returns the Connect2actip field if non-nil, zero value otherwise.

### GetConnect2actipOk

`func (o *HostRest) GetConnect2actipOk() (*string, bool)`

GetConnect2actipOk returns a tuple with the Connect2actip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnect2actip

`func (o *HostRest) SetConnect2actip(v string)`

SetConnect2actip sets Connect2actip field to given value.

### HasConnect2actip

`func (o *HostRest) HasConnect2actip() bool`

HasConnect2actip returns a boolean if a field has been set.

### GetIsAutoDiscoveryEnabled

`func (o *HostRest) GetIsAutoDiscoveryEnabled() bool`

GetIsAutoDiscoveryEnabled returns the IsAutoDiscoveryEnabled field if non-nil, zero value otherwise.

### GetIsAutoDiscoveryEnabledOk

`func (o *HostRest) GetIsAutoDiscoveryEnabledOk() (*bool, bool)`

GetIsAutoDiscoveryEnabledOk returns a tuple with the IsAutoDiscoveryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoDiscoveryEnabled

`func (o *HostRest) SetIsAutoDiscoveryEnabled(v bool)`

SetIsAutoDiscoveryEnabled sets IsAutoDiscoveryEnabled field to given value.

### HasIsAutoDiscoveryEnabled

`func (o *HostRest) HasIsAutoDiscoveryEnabled() bool`

HasIsAutoDiscoveryEnabled returns a boolean if a field has been set.

### GetIsShadowHost

`func (o *HostRest) GetIsShadowHost() bool`

GetIsShadowHost returns the IsShadowHost field if non-nil, zero value otherwise.

### GetIsShadowHostOk

`func (o *HostRest) GetIsShadowHostOk() (*bool, bool)`

GetIsShadowHostOk returns a tuple with the IsShadowHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShadowHost

`func (o *HostRest) SetIsShadowHost(v bool)`

SetIsShadowHost sets IsShadowHost field to given value.

### HasIsShadowHost

`func (o *HostRest) HasIsShadowHost() bool`

HasIsShadowHost returns a boolean if a field has been set.

### GetOrglist

`func (o *HostRest) GetOrglist() []OrganizationRest`

GetOrglist returns the Orglist field if non-nil, zero value otherwise.

### GetOrglistOk

`func (o *HostRest) GetOrglistOk() (*[]OrganizationRest, bool)`

GetOrglistOk returns a tuple with the Orglist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrglist

`func (o *HostRest) SetOrglist(v []OrganizationRest)`

SetOrglist sets Orglist field to given value.

### HasOrglist

`func (o *HostRest) HasOrglist() bool`

HasOrglist returns a boolean if a field has been set.

### GetEsxlist

`func (o *HostRest) GetEsxlist() []HostRest`

GetEsxlist returns the Esxlist field if non-nil, zero value otherwise.

### GetEsxlistOk

`func (o *HostRest) GetEsxlistOk() (*[]HostRest, bool)`

GetEsxlistOk returns a tuple with the Esxlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsxlist

`func (o *HostRest) SetEsxlist(v []HostRest)`

SetEsxlist sets Esxlist field to given value.

### HasEsxlist

`func (o *HostRest) HasEsxlist() bool`

HasEsxlist returns a boolean if a field has been set.

### GetNodeList

`func (o *HostRest) GetNodeList() []HostRest`

GetNodeList returns the NodeList field if non-nil, zero value otherwise.

### GetNodeListOk

`func (o *HostRest) GetNodeListOk() (*[]HostRest, bool)`

GetNodeListOk returns a tuple with the NodeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeList

`func (o *HostRest) SetNodeList(v []HostRest)`

SetNodeList sets NodeList field to given value.

### HasNodeList

`func (o *HostRest) HasNodeList() bool`

HasNodeList returns a boolean if a field has been set.

### GetDatastorelist

`func (o *HostRest) GetDatastorelist() []DataStoreRest`

GetDatastorelist returns the Datastorelist field if non-nil, zero value otherwise.

### GetDatastorelistOk

`func (o *HostRest) GetDatastorelistOk() (*[]DataStoreRest, bool)`

GetDatastorelistOk returns a tuple with the Datastorelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastorelist

`func (o *HostRest) SetDatastorelist(v []DataStoreRest)`

SetDatastorelist sets Datastorelist field to given value.

### HasDatastorelist

`func (o *HostRest) HasDatastorelist() bool`

HasDatastorelist returns a boolean if a field has been set.

### GetOstypeSpecial

`func (o *HostRest) GetOstypeSpecial() string`

GetOstypeSpecial returns the OstypeSpecial field if non-nil, zero value otherwise.

### GetOstypeSpecialOk

`func (o *HostRest) GetOstypeSpecialOk() (*string, bool)`

GetOstypeSpecialOk returns a tuple with the OstypeSpecial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOstypeSpecial

`func (o *HostRest) SetOstypeSpecial(v string)`

SetOstypeSpecial sets OstypeSpecial field to given value.

### HasOstypeSpecial

`func (o *HostRest) HasOstypeSpecial() bool`

HasOstypeSpecial returns a boolean if a field has been set.

### GetPortCount

`func (o *HostRest) GetPortCount() string`

GetPortCount returns the PortCount field if non-nil, zero value otherwise.

### GetPortCountOk

`func (o *HostRest) GetPortCountOk() (*string, bool)`

GetPortCountOk returns a tuple with the PortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortCount

`func (o *HostRest) SetPortCount(v string)`

SetPortCount sets PortCount field to given value.

### HasPortCount

`func (o *HostRest) HasPortCount() bool`

HasPortCount returns a boolean if a field has been set.

### GetIogrpCount

`func (o *HostRest) GetIogrpCount() string`

GetIogrpCount returns the IogrpCount field if non-nil, zero value otherwise.

### GetIogrpCountOk

`func (o *HostRest) GetIogrpCountOk() (*string, bool)`

GetIogrpCountOk returns a tuple with the IogrpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIogrpCount

`func (o *HostRest) SetIogrpCount(v string)`

SetIogrpCount sets IogrpCount field to given value.

### HasIogrpCount

`func (o *HostRest) HasIogrpCount() bool`

HasIogrpCount returns a boolean if a field has been set.

### GetIscsiName

`func (o *HostRest) GetIscsiName() []string`

GetIscsiName returns the IscsiName field if non-nil, zero value otherwise.

### GetIscsiNameOk

`func (o *HostRest) GetIscsiNameOk() (*[]string, bool)`

GetIscsiNameOk returns a tuple with the IscsiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiName

`func (o *HostRest) SetIscsiName(v []string)`

SetIscsiName sets IscsiName field to given value.

### HasIscsiName

`func (o *HostRest) HasIscsiName() bool`

HasIscsiName returns a boolean if a field has been set.

### GetWWPN

`func (o *HostRest) GetWWPN() []string`

GetWWPN returns the WWPN field if non-nil, zero value otherwise.

### GetWWPNOk

`func (o *HostRest) GetWWPNOk() (*[]string, bool)`

GetWWPNOk returns a tuple with the WWPN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWWPN

`func (o *HostRest) SetWWPN(v []string)`

SetWWPN sets WWPN field to given value.

### HasWWPN

`func (o *HostRest) HasWWPN() bool`

HasWWPN returns a boolean if a field has been set.

### GetIogrp

`func (o *HostRest) GetIogrp() string`

GetIogrp returns the Iogrp field if non-nil, zero value otherwise.

### GetIogrpOk

`func (o *HostRest) GetIogrpOk() (*string, bool)`

GetIogrpOk returns a tuple with the Iogrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIogrp

`func (o *HostRest) SetIogrp(v string)`

SetIogrp sets Iogrp field to given value.

### HasIogrp

`func (o *HostRest) HasIogrp() bool`

HasIogrp returns a boolean if a field has been set.

### GetNodeLoggedInCount

`func (o *HostRest) GetNodeLoggedInCount() string`

GetNodeLoggedInCount returns the NodeLoggedInCount field if non-nil, zero value otherwise.

### GetNodeLoggedInCountOk

`func (o *HostRest) GetNodeLoggedInCountOk() (*string, bool)`

GetNodeLoggedInCountOk returns a tuple with the NodeLoggedInCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeLoggedInCount

`func (o *HostRest) SetNodeLoggedInCount(v string)`

SetNodeLoggedInCount sets NodeLoggedInCount field to given value.

### HasNodeLoggedInCount

`func (o *HostRest) HasNodeLoggedInCount() bool`

HasNodeLoggedInCount returns a boolean if a field has been set.

### GetConnectorversion

`func (o *HostRest) GetConnectorversion() string`

GetConnectorversion returns the Connectorversion field if non-nil, zero value otherwise.

### GetConnectorversionOk

`func (o *HostRest) GetConnectorversionOk() (*string, bool)`

GetConnectorversionOk returns a tuple with the Connectorversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorversion

`func (o *HostRest) SetConnectorversion(v string)`

SetConnectorversion sets Connectorversion field to given value.

### HasConnectorversion

`func (o *HostRest) HasConnectorversion() bool`

HasConnectorversion returns a boolean if a field has been set.

### GetIsClusterNode

`func (o *HostRest) GetIsClusterNode() bool`

GetIsClusterNode returns the IsClusterNode field if non-nil, zero value otherwise.

### GetIsClusterNodeOk

`func (o *HostRest) GetIsClusterNodeOk() (*bool, bool)`

GetIsClusterNodeOk returns a tuple with the IsClusterNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClusterNode

`func (o *HostRest) SetIsClusterNode(v bool)`

SetIsClusterNode sets IsClusterNode field to given value.

### HasIsClusterNode

`func (o *HostRest) HasIsClusterNode() bool`

HasIsClusterNode returns a boolean if a field has been set.

### GetReconciliationNeeded

`func (o *HostRest) GetReconciliationNeeded() bool`

GetReconciliationNeeded returns the ReconciliationNeeded field if non-nil, zero value otherwise.

### GetReconciliationNeededOk

`func (o *HostRest) GetReconciliationNeededOk() (*bool, bool)`

GetReconciliationNeededOk returns a tuple with the ReconciliationNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciliationNeeded

`func (o *HostRest) SetReconciliationNeeded(v bool)`

SetReconciliationNeeded sets ReconciliationNeeded field to given value.

### HasReconciliationNeeded

`func (o *HostRest) HasReconciliationNeeded() bool`

HasReconciliationNeeded returns a boolean if a field has been set.

### GetReconciliationFields

`func (o *HostRest) GetReconciliationFields() []string`

GetReconciliationFields returns the ReconciliationFields field if non-nil, zero value otherwise.

### GetReconciliationFieldsOk

`func (o *HostRest) GetReconciliationFieldsOk() (*[]string, bool)`

GetReconciliationFieldsOk returns a tuple with the ReconciliationFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciliationFields

`func (o *HostRest) SetReconciliationFields(v []string)`

SetReconciliationFields sets ReconciliationFields field to given value.

### HasReconciliationFields

`func (o *HostRest) HasReconciliationFields() bool`

HasReconciliationFields returns a boolean if a field has been set.

### GetAgents

`func (o *HostRest) GetAgents() []AgentRest`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *HostRest) GetAgentsOk() (*[]AgentRest, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *HostRest) SetAgents(v []AgentRest)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *HostRest) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### GetArrays

`func (o *HostRest) GetArrays() []ArrayRest`

GetArrays returns the Arrays field if non-nil, zero value otherwise.

### GetArraysOk

`func (o *HostRest) GetArraysOk() (*[]ArrayRest, bool)`

GetArraysOk returns a tuple with the Arrays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrays

`func (o *HostRest) SetArrays(v []ArrayRest)`

SetArrays sets Arrays field to given value.

### HasArrays

`func (o *HostRest) HasArrays() bool`

HasArrays returns a boolean if a field has been set.

### GetVcenterHost

`func (o *HostRest) GetVcenterHost() HostRest`

GetVcenterHost returns the VcenterHost field if non-nil, zero value otherwise.

### GetVcenterHostOk

`func (o *HostRest) GetVcenterHostOk() (*HostRest, bool)`

GetVcenterHostOk returns a tuple with the VcenterHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterHost

`func (o *HostRest) SetVcenterHost(v HostRest)`

SetVcenterHost sets VcenterHost field to given value.

### HasVcenterHost

`func (o *HostRest) HasVcenterHost() bool`

HasVcenterHost returns a boolean if a field has been set.

### GetNasdconfig

`func (o *HostRest) GetNasdconfig() map[string]interface{}`

GetNasdconfig returns the Nasdconfig field if non-nil, zero value otherwise.

### GetNasdconfigOk

`func (o *HostRest) GetNasdconfigOk() (*map[string]interface{}, bool)`

GetNasdconfigOk returns a tuple with the Nasdconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasdconfig

`func (o *HostRest) SetNasdconfig(v map[string]interface{})`

SetNasdconfig sets Nasdconfig field to given value.

### HasNasdconfig

`func (o *HostRest) HasNasdconfig() bool`

HasNasdconfig returns a boolean if a field has been set.

### GetCloudcredential

`func (o *HostRest) GetCloudcredential() CloudCredentialRest`

GetCloudcredential returns the Cloudcredential field if non-nil, zero value otherwise.

### GetCloudcredentialOk

`func (o *HostRest) GetCloudcredentialOk() (*CloudCredentialRest, bool)`

GetCloudcredentialOk returns a tuple with the Cloudcredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudcredential

`func (o *HostRest) SetCloudcredential(v CloudCredentialRest)`

SetCloudcredential sets Cloudcredential field to given value.

### HasCloudcredential

`func (o *HostRest) HasCloudcredential() bool`

HasCloudcredential returns a boolean if a field has been set.

### GetChapusername

`func (o *HostRest) GetChapusername() string`

GetChapusername returns the Chapusername field if non-nil, zero value otherwise.

### GetChapusernameOk

`func (o *HostRest) GetChapusernameOk() (*string, bool)`

GetChapusernameOk returns a tuple with the Chapusername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapusername

`func (o *HostRest) SetChapusername(v string)`

SetChapusername sets Chapusername field to given value.

### HasChapusername

`func (o *HostRest) HasChapusername() bool`

HasChapusername returns a boolean if a field has been set.

### GetChappassword

`func (o *HostRest) GetChappassword() string`

GetChappassword returns the Chappassword field if non-nil, zero value otherwise.

### GetChappasswordOk

`func (o *HostRest) GetChappasswordOk() (*string, bool)`

GetChappasswordOk returns a tuple with the Chappassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChappassword

`func (o *HostRest) SetChappassword(v string)`

SetChappassword sets Chappassword field to given value.

### HasChappassword

`func (o *HostRest) HasChappassword() bool`

HasChappassword returns a boolean if a field has been set.

### GetClearchap

`func (o *HostRest) GetClearchap() bool`

GetClearchap returns the Clearchap field if non-nil, zero value otherwise.

### GetClearchapOk

`func (o *HostRest) GetClearchapOk() (*bool, bool)`

GetClearchapOk returns a tuple with the Clearchap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearchap

`func (o *HostRest) SetClearchap(v bool)`

SetClearchap sets Clearchap field to given value.

### HasClearchap

`func (o *HostRest) HasClearchap() bool`

HasClearchap returns a boolean if a field has been set.

### GetUdsagent

`func (o *HostRest) GetUdsagent() AgentRest`

GetUdsagent returns the Udsagent field if non-nil, zero value otherwise.

### GetUdsagentOk

`func (o *HostRest) GetUdsagentOk() (*AgentRest, bool)`

GetUdsagentOk returns a tuple with the Udsagent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdsagent

`func (o *HostRest) SetUdsagent(v AgentRest)`

SetUdsagent sets Udsagent field to given value.

### HasUdsagent

`func (o *HostRest) HasUdsagent() bool`

HasUdsagent returns a boolean if a field has been set.

### GetApplicationagent

`func (o *HostRest) GetApplicationagent() AgentRest`

GetApplicationagent returns the Applicationagent field if non-nil, zero value otherwise.

### GetApplicationagentOk

`func (o *HostRest) GetApplicationagentOk() (*AgentRest, bool)`

GetApplicationagentOk returns a tuple with the Applicationagent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationagent

`func (o *HostRest) SetApplicationagent(v AgentRest)`

SetApplicationagent sets Applicationagent field to given value.

### HasApplicationagent

`func (o *HostRest) HasApplicationagent() bool`

HasApplicationagent returns a boolean if a field has been set.

### GetHypervisoragent

`func (o *HostRest) GetHypervisoragent() AgentRest`

GetHypervisoragent returns the Hypervisoragent field if non-nil, zero value otherwise.

### GetHypervisoragentOk

`func (o *HostRest) GetHypervisoragentOk() (*AgentRest, bool)`

GetHypervisoragentOk returns a tuple with the Hypervisoragent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisoragent

`func (o *HostRest) SetHypervisoragent(v AgentRest)`

SetHypervisoragent sets Hypervisoragent field to given value.

### HasHypervisoragent

`func (o *HostRest) HasHypervisoragent() bool`

HasHypervisoragent returns a boolean if a field has been set.

### GetHmcconnector

`func (o *HostRest) GetHmcconnector() AgentRest`

GetHmcconnector returns the Hmcconnector field if non-nil, zero value otherwise.

### GetHmcconnectorOk

`func (o *HostRest) GetHmcconnectorOk() (*AgentRest, bool)`

GetHmcconnectorOk returns a tuple with the Hmcconnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmcconnector

`func (o *HostRest) SetHmcconnector(v AgentRest)`

SetHmcconnector sets Hmcconnector field to given value.

### HasHmcconnector

`func (o *HostRest) HasHmcconnector() bool`

HasHmcconnector returns a boolean if a field has been set.

### GetAppliance

`func (o *HostRest) GetAppliance() ClusterRest`

GetAppliance returns the Appliance field if non-nil, zero value otherwise.

### GetApplianceOk

`func (o *HostRest) GetApplianceOk() (*ClusterRest, bool)`

GetApplianceOk returns a tuple with the Appliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliance

`func (o *HostRest) SetAppliance(v ClusterRest)`

SetAppliance sets Appliance field to given value.

### HasAppliance

`func (o *HostRest) HasAppliance() bool`

HasAppliance returns a boolean if a field has been set.

### GetId

`func (o *HostRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HostRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *HostRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *HostRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *HostRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *HostRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *HostRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *HostRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *HostRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *HostRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *HostRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *HostRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *HostRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *HostRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



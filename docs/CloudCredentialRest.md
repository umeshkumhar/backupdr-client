# CloudCredentialRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Usedefaultsa** | Pointer to **bool** |  | [optional] 
**Immutable** | Pointer to **bool** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Sources** | Pointer to [**[]CloudCredentialRest**](CloudCredentialRest.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Publickey** | Pointer to **string** |  | [optional] 
**Vaultpool** | Pointer to [**DiskPoolRest**](DiskPoolRest.md) |  | [optional] 
**Projectid** | Pointer to **string** |  | [optional] 
**SrcId** | Pointer to **int64** |  | [optional] 
**Vaultudsuid** | Pointer to **int32** |  | [optional] 
**ClusterId** | Pointer to **int64** |  | [optional] 
**Clientid** | Pointer to **string** |  | [optional] 
**Privatekey** | Pointer to **string** |  | [optional] 
**Orglist** | Pointer to [**[]OrganizationRest**](OrganizationRest.md) |  | [optional] 
**Secretkey** | Pointer to **string** |  | [optional] 
**Serviceaccount** | Pointer to **string** |  | [optional] 
**Credential** | Pointer to **string** |  | [optional] 
**Cloudtype** | Pointer to **string** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**Subscriptionid** | Pointer to **string** |  | [optional] 
**Appliance** | Pointer to [**ClusterRest**](ClusterRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewCloudCredentialRest

`func NewCloudCredentialRest() *CloudCredentialRest`

NewCloudCredentialRest instantiates a new CloudCredentialRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialRestWithDefaults

`func NewCloudCredentialRestWithDefaults() *CloudCredentialRest`

NewCloudCredentialRestWithDefaults instantiates a new CloudCredentialRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsedefaultsa

`func (o *CloudCredentialRest) GetUsedefaultsa() bool`

GetUsedefaultsa returns the Usedefaultsa field if non-nil, zero value otherwise.

### GetUsedefaultsaOk

`func (o *CloudCredentialRest) GetUsedefaultsaOk() (*bool, bool)`

GetUsedefaultsaOk returns a tuple with the Usedefaultsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedefaultsa

`func (o *CloudCredentialRest) SetUsedefaultsa(v bool)`

SetUsedefaultsa sets Usedefaultsa field to given value.

### HasUsedefaultsa

`func (o *CloudCredentialRest) HasUsedefaultsa() bool`

HasUsedefaultsa returns a boolean if a field has been set.

### GetImmutable

`func (o *CloudCredentialRest) GetImmutable() bool`

GetImmutable returns the Immutable field if non-nil, zero value otherwise.

### GetImmutableOk

`func (o *CloudCredentialRest) GetImmutableOk() (*bool, bool)`

GetImmutableOk returns a tuple with the Immutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutable

`func (o *CloudCredentialRest) SetImmutable(v bool)`

SetImmutable sets Immutable field to given value.

### HasImmutable

`func (o *CloudCredentialRest) HasImmutable() bool`

HasImmutable returns a boolean if a field has been set.

### GetRegion

`func (o *CloudCredentialRest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudCredentialRest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudCredentialRest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudCredentialRest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetDomain

`func (o *CloudCredentialRest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CloudCredentialRest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CloudCredentialRest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CloudCredentialRest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetSources

`func (o *CloudCredentialRest) GetSources() []CloudCredentialRest`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *CloudCredentialRest) GetSourcesOk() (*[]CloudCredentialRest, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *CloudCredentialRest) SetSources(v []CloudCredentialRest)`

SetSources sets Sources field to given value.

### HasSources

`func (o *CloudCredentialRest) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetName

`func (o *CloudCredentialRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCredentialRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCredentialRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudCredentialRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublickey

`func (o *CloudCredentialRest) GetPublickey() string`

GetPublickey returns the Publickey field if non-nil, zero value otherwise.

### GetPublickeyOk

`func (o *CloudCredentialRest) GetPublickeyOk() (*string, bool)`

GetPublickeyOk returns a tuple with the Publickey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublickey

`func (o *CloudCredentialRest) SetPublickey(v string)`

SetPublickey sets Publickey field to given value.

### HasPublickey

`func (o *CloudCredentialRest) HasPublickey() bool`

HasPublickey returns a boolean if a field has been set.

### GetVaultpool

`func (o *CloudCredentialRest) GetVaultpool() DiskPoolRest`

GetVaultpool returns the Vaultpool field if non-nil, zero value otherwise.

### GetVaultpoolOk

`func (o *CloudCredentialRest) GetVaultpoolOk() (*DiskPoolRest, bool)`

GetVaultpoolOk returns a tuple with the Vaultpool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultpool

`func (o *CloudCredentialRest) SetVaultpool(v DiskPoolRest)`

SetVaultpool sets Vaultpool field to given value.

### HasVaultpool

`func (o *CloudCredentialRest) HasVaultpool() bool`

HasVaultpool returns a boolean if a field has been set.

### GetProjectid

`func (o *CloudCredentialRest) GetProjectid() string`

GetProjectid returns the Projectid field if non-nil, zero value otherwise.

### GetProjectidOk

`func (o *CloudCredentialRest) GetProjectidOk() (*string, bool)`

GetProjectidOk returns a tuple with the Projectid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectid

`func (o *CloudCredentialRest) SetProjectid(v string)`

SetProjectid sets Projectid field to given value.

### HasProjectid

`func (o *CloudCredentialRest) HasProjectid() bool`

HasProjectid returns a boolean if a field has been set.

### GetSrcId

`func (o *CloudCredentialRest) GetSrcId() int64`

GetSrcId returns the SrcId field if non-nil, zero value otherwise.

### GetSrcIdOk

`func (o *CloudCredentialRest) GetSrcIdOk() (*int64, bool)`

GetSrcIdOk returns a tuple with the SrcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcId

`func (o *CloudCredentialRest) SetSrcId(v int64)`

SetSrcId sets SrcId field to given value.

### HasSrcId

`func (o *CloudCredentialRest) HasSrcId() bool`

HasSrcId returns a boolean if a field has been set.

### GetVaultudsuid

`func (o *CloudCredentialRest) GetVaultudsuid() int32`

GetVaultudsuid returns the Vaultudsuid field if non-nil, zero value otherwise.

### GetVaultudsuidOk

`func (o *CloudCredentialRest) GetVaultudsuidOk() (*int32, bool)`

GetVaultudsuidOk returns a tuple with the Vaultudsuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultudsuid

`func (o *CloudCredentialRest) SetVaultudsuid(v int32)`

SetVaultudsuid sets Vaultudsuid field to given value.

### HasVaultudsuid

`func (o *CloudCredentialRest) HasVaultudsuid() bool`

HasVaultudsuid returns a boolean if a field has been set.

### GetClusterId

`func (o *CloudCredentialRest) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *CloudCredentialRest) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *CloudCredentialRest) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *CloudCredentialRest) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetClientid

`func (o *CloudCredentialRest) GetClientid() string`

GetClientid returns the Clientid field if non-nil, zero value otherwise.

### GetClientidOk

`func (o *CloudCredentialRest) GetClientidOk() (*string, bool)`

GetClientidOk returns a tuple with the Clientid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientid

`func (o *CloudCredentialRest) SetClientid(v string)`

SetClientid sets Clientid field to given value.

### HasClientid

`func (o *CloudCredentialRest) HasClientid() bool`

HasClientid returns a boolean if a field has been set.

### GetPrivatekey

`func (o *CloudCredentialRest) GetPrivatekey() string`

GetPrivatekey returns the Privatekey field if non-nil, zero value otherwise.

### GetPrivatekeyOk

`func (o *CloudCredentialRest) GetPrivatekeyOk() (*string, bool)`

GetPrivatekeyOk returns a tuple with the Privatekey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivatekey

`func (o *CloudCredentialRest) SetPrivatekey(v string)`

SetPrivatekey sets Privatekey field to given value.

### HasPrivatekey

`func (o *CloudCredentialRest) HasPrivatekey() bool`

HasPrivatekey returns a boolean if a field has been set.

### GetOrglist

`func (o *CloudCredentialRest) GetOrglist() []OrganizationRest`

GetOrglist returns the Orglist field if non-nil, zero value otherwise.

### GetOrglistOk

`func (o *CloudCredentialRest) GetOrglistOk() (*[]OrganizationRest, bool)`

GetOrglistOk returns a tuple with the Orglist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrglist

`func (o *CloudCredentialRest) SetOrglist(v []OrganizationRest)`

SetOrglist sets Orglist field to given value.

### HasOrglist

`func (o *CloudCredentialRest) HasOrglist() bool`

HasOrglist returns a boolean if a field has been set.

### GetSecretkey

`func (o *CloudCredentialRest) GetSecretkey() string`

GetSecretkey returns the Secretkey field if non-nil, zero value otherwise.

### GetSecretkeyOk

`func (o *CloudCredentialRest) GetSecretkeyOk() (*string, bool)`

GetSecretkeyOk returns a tuple with the Secretkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretkey

`func (o *CloudCredentialRest) SetSecretkey(v string)`

SetSecretkey sets Secretkey field to given value.

### HasSecretkey

`func (o *CloudCredentialRest) HasSecretkey() bool`

HasSecretkey returns a boolean if a field has been set.

### GetServiceaccount

`func (o *CloudCredentialRest) GetServiceaccount() string`

GetServiceaccount returns the Serviceaccount field if non-nil, zero value otherwise.

### GetServiceaccountOk

`func (o *CloudCredentialRest) GetServiceaccountOk() (*string, bool)`

GetServiceaccountOk returns a tuple with the Serviceaccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceaccount

`func (o *CloudCredentialRest) SetServiceaccount(v string)`

SetServiceaccount sets Serviceaccount field to given value.

### HasServiceaccount

`func (o *CloudCredentialRest) HasServiceaccount() bool`

HasServiceaccount returns a boolean if a field has been set.

### GetCredential

`func (o *CloudCredentialRest) GetCredential() string`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *CloudCredentialRest) GetCredentialOk() (*string, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *CloudCredentialRest) SetCredential(v string)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *CloudCredentialRest) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetCloudtype

`func (o *CloudCredentialRest) GetCloudtype() string`

GetCloudtype returns the Cloudtype field if non-nil, zero value otherwise.

### GetCloudtypeOk

`func (o *CloudCredentialRest) GetCloudtypeOk() (*string, bool)`

GetCloudtypeOk returns a tuple with the Cloudtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudtype

`func (o *CloudCredentialRest) SetCloudtype(v string)`

SetCloudtype sets Cloudtype field to given value.

### HasCloudtype

`func (o *CloudCredentialRest) HasCloudtype() bool`

HasCloudtype returns a boolean if a field has been set.

### GetEndpoint

`func (o *CloudCredentialRest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *CloudCredentialRest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *CloudCredentialRest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *CloudCredentialRest) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetSubscriptionid

`func (o *CloudCredentialRest) GetSubscriptionid() string`

GetSubscriptionid returns the Subscriptionid field if non-nil, zero value otherwise.

### GetSubscriptionidOk

`func (o *CloudCredentialRest) GetSubscriptionidOk() (*string, bool)`

GetSubscriptionidOk returns a tuple with the Subscriptionid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionid

`func (o *CloudCredentialRest) SetSubscriptionid(v string)`

SetSubscriptionid sets Subscriptionid field to given value.

### HasSubscriptionid

`func (o *CloudCredentialRest) HasSubscriptionid() bool`

HasSubscriptionid returns a boolean if a field has been set.

### GetAppliance

`func (o *CloudCredentialRest) GetAppliance() ClusterRest`

GetAppliance returns the Appliance field if non-nil, zero value otherwise.

### GetApplianceOk

`func (o *CloudCredentialRest) GetApplianceOk() (*ClusterRest, bool)`

GetApplianceOk returns a tuple with the Appliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliance

`func (o *CloudCredentialRest) SetAppliance(v ClusterRest)`

SetAppliance sets Appliance field to given value.

### HasAppliance

`func (o *CloudCredentialRest) HasAppliance() bool`

HasAppliance returns a boolean if a field has been set.

### GetId

`func (o *CloudCredentialRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCredentialRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCredentialRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudCredentialRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *CloudCredentialRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CloudCredentialRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CloudCredentialRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CloudCredentialRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *CloudCredentialRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *CloudCredentialRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *CloudCredentialRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *CloudCredentialRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *CloudCredentialRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *CloudCredentialRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *CloudCredentialRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *CloudCredentialRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



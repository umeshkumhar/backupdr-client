# AppDiscoveryRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apptypes** | Pointer to **[]string** |  | [optional] 
**Listonly** | Pointer to **bool** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Host** | Pointer to [**HostRest**](HostRest.md) |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Ipaddress** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to **string** |  | [optional] 
**Versiononly** | Pointer to **bool** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewAppDiscoveryRest

`func NewAppDiscoveryRest() *AppDiscoveryRest`

NewAppDiscoveryRest instantiates a new AppDiscoveryRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDiscoveryRestWithDefaults

`func NewAppDiscoveryRestWithDefaults() *AppDiscoveryRest`

NewAppDiscoveryRestWithDefaults instantiates a new AppDiscoveryRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApptypes

`func (o *AppDiscoveryRest) GetApptypes() []string`

GetApptypes returns the Apptypes field if non-nil, zero value otherwise.

### GetApptypesOk

`func (o *AppDiscoveryRest) GetApptypesOk() (*[]string, bool)`

GetApptypesOk returns a tuple with the Apptypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApptypes

`func (o *AppDiscoveryRest) SetApptypes(v []string)`

SetApptypes sets Apptypes field to given value.

### HasApptypes

`func (o *AppDiscoveryRest) HasApptypes() bool`

HasApptypes returns a boolean if a field has been set.

### GetListonly

`func (o *AppDiscoveryRest) GetListonly() bool`

GetListonly returns the Listonly field if non-nil, zero value otherwise.

### GetListonlyOk

`func (o *AppDiscoveryRest) GetListonlyOk() (*bool, bool)`

GetListonlyOk returns a tuple with the Listonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListonly

`func (o *AppDiscoveryRest) SetListonly(v bool)`

SetListonly sets Listonly field to given value.

### HasListonly

`func (o *AppDiscoveryRest) HasListonly() bool`

HasListonly returns a boolean if a field has been set.

### GetPassword

`func (o *AppDiscoveryRest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AppDiscoveryRest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AppDiscoveryRest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AppDiscoveryRest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetType

`func (o *AppDiscoveryRest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppDiscoveryRest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppDiscoveryRest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AppDiscoveryRest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPort

`func (o *AppDiscoveryRest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AppDiscoveryRest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AppDiscoveryRest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *AppDiscoveryRest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetHost

`func (o *AppDiscoveryRest) GetHost() HostRest`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AppDiscoveryRest) GetHostOk() (*HostRest, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AppDiscoveryRest) SetHost(v HostRest)`

SetHost sets Host field to given value.

### HasHost

`func (o *AppDiscoveryRest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetUsername

`func (o *AppDiscoveryRest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AppDiscoveryRest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AppDiscoveryRest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AppDiscoveryRest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetIpaddress

`func (o *AppDiscoveryRest) GetIpaddress() string`

GetIpaddress returns the Ipaddress field if non-nil, zero value otherwise.

### GetIpaddressOk

`func (o *AppDiscoveryRest) GetIpaddressOk() (*string, bool)`

GetIpaddressOk returns a tuple with the Ipaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddress

`func (o *AppDiscoveryRest) SetIpaddress(v string)`

SetIpaddress sets Ipaddress field to given value.

### HasIpaddress

`func (o *AppDiscoveryRest) HasIpaddress() bool`

HasIpaddress returns a boolean if a field has been set.

### GetCluster

`func (o *AppDiscoveryRest) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *AppDiscoveryRest) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *AppDiscoveryRest) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *AppDiscoveryRest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetVersiononly

`func (o *AppDiscoveryRest) GetVersiononly() bool`

GetVersiononly returns the Versiononly field if non-nil, zero value otherwise.

### GetVersiononlyOk

`func (o *AppDiscoveryRest) GetVersiononlyOk() (*bool, bool)`

GetVersiononlyOk returns a tuple with the Versiononly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersiononly

`func (o *AppDiscoveryRest) SetVersiononly(v bool)`

SetVersiononly sets Versiononly field to given value.

### HasVersiononly

`func (o *AppDiscoveryRest) HasVersiononly() bool`

HasVersiononly returns a boolean if a field has been set.

### GetOrg

`func (o *AppDiscoveryRest) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *AppDiscoveryRest) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *AppDiscoveryRest) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *AppDiscoveryRest) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetId

`func (o *AppDiscoveryRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppDiscoveryRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppDiscoveryRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppDiscoveryRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *AppDiscoveryRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AppDiscoveryRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AppDiscoveryRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AppDiscoveryRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *AppDiscoveryRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *AppDiscoveryRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *AppDiscoveryRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *AppDiscoveryRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *AppDiscoveryRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *AppDiscoveryRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *AppDiscoveryRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *AppDiscoveryRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ConsistencyGroupRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Host** | Pointer to [**HostRest**](HostRest.md) |  | [optional] 
**Groupname** | Pointer to **string** |  | [optional] 
**Application** | Pointer to [**ApplicationRest**](ApplicationRest.md) |  | [optional] 
**Cluster** | Pointer to [**ClusterRest**](ClusterRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewConsistencyGroupRest

`func NewConsistencyGroupRest() *ConsistencyGroupRest`

NewConsistencyGroupRest instantiates a new ConsistencyGroupRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsistencyGroupRestWithDefaults

`func NewConsistencyGroupRestWithDefaults() *ConsistencyGroupRest`

NewConsistencyGroupRestWithDefaults instantiates a new ConsistencyGroupRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ConsistencyGroupRest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConsistencyGroupRest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConsistencyGroupRest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConsistencyGroupRest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHost

`func (o *ConsistencyGroupRest) GetHost() HostRest`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ConsistencyGroupRest) GetHostOk() (*HostRest, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ConsistencyGroupRest) SetHost(v HostRest)`

SetHost sets Host field to given value.

### HasHost

`func (o *ConsistencyGroupRest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetGroupname

`func (o *ConsistencyGroupRest) GetGroupname() string`

GetGroupname returns the Groupname field if non-nil, zero value otherwise.

### GetGroupnameOk

`func (o *ConsistencyGroupRest) GetGroupnameOk() (*string, bool)`

GetGroupnameOk returns a tuple with the Groupname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupname

`func (o *ConsistencyGroupRest) SetGroupname(v string)`

SetGroupname sets Groupname field to given value.

### HasGroupname

`func (o *ConsistencyGroupRest) HasGroupname() bool`

HasGroupname returns a boolean if a field has been set.

### GetApplication

`func (o *ConsistencyGroupRest) GetApplication() ApplicationRest`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ConsistencyGroupRest) GetApplicationOk() (*ApplicationRest, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ConsistencyGroupRest) SetApplication(v ApplicationRest)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *ConsistencyGroupRest) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCluster

`func (o *ConsistencyGroupRest) GetCluster() ClusterRest`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ConsistencyGroupRest) GetClusterOk() (*ClusterRest, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ConsistencyGroupRest) SetCluster(v ClusterRest)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ConsistencyGroupRest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetId

`func (o *ConsistencyGroupRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsistencyGroupRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsistencyGroupRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConsistencyGroupRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ConsistencyGroupRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ConsistencyGroupRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ConsistencyGroupRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ConsistencyGroupRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ConsistencyGroupRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ConsistencyGroupRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ConsistencyGroupRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ConsistencyGroupRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ConsistencyGroupRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ConsistencyGroupRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ConsistencyGroupRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ConsistencyGroupRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



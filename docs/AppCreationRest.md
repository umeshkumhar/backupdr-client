# AppCreationRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterRest**](ClusterRest.md) |  | [optional] 
**Applications** | Pointer to [**[]ApplicationRest**](ApplicationRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewAppCreationRest

`func NewAppCreationRest() *AppCreationRest`

NewAppCreationRest instantiates a new AppCreationRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppCreationRestWithDefaults

`func NewAppCreationRestWithDefaults() *AppCreationRest`

NewAppCreationRestWithDefaults instantiates a new AppCreationRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *AppCreationRest) GetCluster() ClusterRest`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *AppCreationRest) GetClusterOk() (*ClusterRest, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *AppCreationRest) SetCluster(v ClusterRest)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *AppCreationRest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetApplications

`func (o *AppCreationRest) GetApplications() []ApplicationRest`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *AppCreationRest) GetApplicationsOk() (*[]ApplicationRest, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *AppCreationRest) SetApplications(v []ApplicationRest)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *AppCreationRest) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetId

`func (o *AppCreationRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppCreationRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppCreationRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppCreationRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *AppCreationRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AppCreationRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AppCreationRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AppCreationRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *AppCreationRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *AppCreationRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *AppCreationRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *AppCreationRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *AppCreationRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *AppCreationRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *AppCreationRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *AppCreationRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



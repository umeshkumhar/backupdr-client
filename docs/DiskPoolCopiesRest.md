# DiskPoolCopiesRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusterid** | Pointer to **int64** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Clustername** | Pointer to **string** |  | [optional] 
**Errormsg** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewDiskPoolCopiesRest

`func NewDiskPoolCopiesRest() *DiskPoolCopiesRest`

NewDiskPoolCopiesRest instantiates a new DiskPoolCopiesRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskPoolCopiesRestWithDefaults

`func NewDiskPoolCopiesRestWithDefaults() *DiskPoolCopiesRest`

NewDiskPoolCopiesRestWithDefaults instantiates a new DiskPoolCopiesRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterid

`func (o *DiskPoolCopiesRest) GetClusterid() int64`

GetClusterid returns the Clusterid field if non-nil, zero value otherwise.

### GetClusteridOk

`func (o *DiskPoolCopiesRest) GetClusteridOk() (*int64, bool)`

GetClusteridOk returns a tuple with the Clusterid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterid

`func (o *DiskPoolCopiesRest) SetClusterid(v int64)`

SetClusterid sets Clusterid field to given value.

### HasClusterid

`func (o *DiskPoolCopiesRest) HasClusterid() bool`

HasClusterid returns a boolean if a field has been set.

### GetSuccess

`func (o *DiskPoolCopiesRest) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *DiskPoolCopiesRest) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *DiskPoolCopiesRest) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *DiskPoolCopiesRest) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetClustername

`func (o *DiskPoolCopiesRest) GetClustername() string`

GetClustername returns the Clustername field if non-nil, zero value otherwise.

### GetClusternameOk

`func (o *DiskPoolCopiesRest) GetClusternameOk() (*string, bool)`

GetClusternameOk returns a tuple with the Clustername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClustername

`func (o *DiskPoolCopiesRest) SetClustername(v string)`

SetClustername sets Clustername field to given value.

### HasClustername

`func (o *DiskPoolCopiesRest) HasClustername() bool`

HasClustername returns a boolean if a field has been set.

### GetErrormsg

`func (o *DiskPoolCopiesRest) GetErrormsg() string`

GetErrormsg returns the Errormsg field if non-nil, zero value otherwise.

### GetErrormsgOk

`func (o *DiskPoolCopiesRest) GetErrormsgOk() (*string, bool)`

GetErrormsgOk returns a tuple with the Errormsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrormsg

`func (o *DiskPoolCopiesRest) SetErrormsg(v string)`

SetErrormsg sets Errormsg field to given value.

### HasErrormsg

`func (o *DiskPoolCopiesRest) HasErrormsg() bool`

HasErrormsg returns a boolean if a field has been set.

### GetId

`func (o *DiskPoolCopiesRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiskPoolCopiesRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiskPoolCopiesRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DiskPoolCopiesRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *DiskPoolCopiesRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DiskPoolCopiesRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DiskPoolCopiesRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *DiskPoolCopiesRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *DiskPoolCopiesRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *DiskPoolCopiesRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *DiskPoolCopiesRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *DiskPoolCopiesRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *DiskPoolCopiesRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *DiskPoolCopiesRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *DiskPoolCopiesRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *DiskPoolCopiesRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



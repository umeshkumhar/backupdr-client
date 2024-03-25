# SelectedObjectRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Diskpool** | Pointer to [**DiskPoolRest**](DiskPoolRest.md) |  | [optional] 
**Restorableobject** | Pointer to **string** |  | [optional] 
**Datastore** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewSelectedObjectRest

`func NewSelectedObjectRest() *SelectedObjectRest`

NewSelectedObjectRest instantiates a new SelectedObjectRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectedObjectRestWithDefaults

`func NewSelectedObjectRestWithDefaults() *SelectedObjectRest`

NewSelectedObjectRestWithDefaults instantiates a new SelectedObjectRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskpool

`func (o *SelectedObjectRest) GetDiskpool() DiskPoolRest`

GetDiskpool returns the Diskpool field if non-nil, zero value otherwise.

### GetDiskpoolOk

`func (o *SelectedObjectRest) GetDiskpoolOk() (*DiskPoolRest, bool)`

GetDiskpoolOk returns a tuple with the Diskpool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskpool

`func (o *SelectedObjectRest) SetDiskpool(v DiskPoolRest)`

SetDiskpool sets Diskpool field to given value.

### HasDiskpool

`func (o *SelectedObjectRest) HasDiskpool() bool`

HasDiskpool returns a boolean if a field has been set.

### GetRestorableobject

`func (o *SelectedObjectRest) GetRestorableobject() string`

GetRestorableobject returns the Restorableobject field if non-nil, zero value otherwise.

### GetRestorableobjectOk

`func (o *SelectedObjectRest) GetRestorableobjectOk() (*string, bool)`

GetRestorableobjectOk returns a tuple with the Restorableobject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestorableobject

`func (o *SelectedObjectRest) SetRestorableobject(v string)`

SetRestorableobject sets Restorableobject field to given value.

### HasRestorableobject

`func (o *SelectedObjectRest) HasRestorableobject() bool`

HasRestorableobject returns a boolean if a field has been set.

### GetDatastore

`func (o *SelectedObjectRest) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *SelectedObjectRest) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *SelectedObjectRest) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *SelectedObjectRest) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetId

`func (o *SelectedObjectRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SelectedObjectRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SelectedObjectRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SelectedObjectRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *SelectedObjectRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SelectedObjectRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SelectedObjectRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SelectedObjectRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *SelectedObjectRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *SelectedObjectRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *SelectedObjectRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *SelectedObjectRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *SelectedObjectRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *SelectedObjectRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *SelectedObjectRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *SelectedObjectRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



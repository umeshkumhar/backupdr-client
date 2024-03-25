# LiveCloneRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Diskpool** | Pointer to [**DiskPoolRest**](DiskPoolRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewLiveCloneRest

`func NewLiveCloneRest() *LiveCloneRest`

NewLiveCloneRest instantiates a new LiveCloneRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveCloneRestWithDefaults

`func NewLiveCloneRestWithDefaults() *LiveCloneRest`

NewLiveCloneRestWithDefaults instantiates a new LiveCloneRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *LiveCloneRest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LiveCloneRest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LiveCloneRest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *LiveCloneRest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDiskpool

`func (o *LiveCloneRest) GetDiskpool() DiskPoolRest`

GetDiskpool returns the Diskpool field if non-nil, zero value otherwise.

### GetDiskpoolOk

`func (o *LiveCloneRest) GetDiskpoolOk() (*DiskPoolRest, bool)`

GetDiskpoolOk returns a tuple with the Diskpool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskpool

`func (o *LiveCloneRest) SetDiskpool(v DiskPoolRest)`

SetDiskpool sets Diskpool field to given value.

### HasDiskpool

`func (o *LiveCloneRest) HasDiskpool() bool`

HasDiskpool returns a boolean if a field has been set.

### GetId

`func (o *LiveCloneRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LiveCloneRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LiveCloneRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LiveCloneRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *LiveCloneRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LiveCloneRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LiveCloneRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *LiveCloneRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *LiveCloneRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *LiveCloneRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *LiveCloneRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *LiveCloneRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *LiveCloneRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *LiveCloneRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *LiveCloneRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *LiveCloneRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# LogSequenceRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Beginlsn** | Pointer to **string** |  | [optional] 
**Endlsn** | Pointer to **string** |  | [optional] 
**Thread** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewLogSequenceRest

`func NewLogSequenceRest() *LogSequenceRest`

NewLogSequenceRest instantiates a new LogSequenceRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogSequenceRestWithDefaults

`func NewLogSequenceRestWithDefaults() *LogSequenceRest`

NewLogSequenceRestWithDefaults instantiates a new LogSequenceRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeginlsn

`func (o *LogSequenceRest) GetBeginlsn() string`

GetBeginlsn returns the Beginlsn field if non-nil, zero value otherwise.

### GetBeginlsnOk

`func (o *LogSequenceRest) GetBeginlsnOk() (*string, bool)`

GetBeginlsnOk returns a tuple with the Beginlsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginlsn

`func (o *LogSequenceRest) SetBeginlsn(v string)`

SetBeginlsn sets Beginlsn field to given value.

### HasBeginlsn

`func (o *LogSequenceRest) HasBeginlsn() bool`

HasBeginlsn returns a boolean if a field has been set.

### GetEndlsn

`func (o *LogSequenceRest) GetEndlsn() string`

GetEndlsn returns the Endlsn field if non-nil, zero value otherwise.

### GetEndlsnOk

`func (o *LogSequenceRest) GetEndlsnOk() (*string, bool)`

GetEndlsnOk returns a tuple with the Endlsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndlsn

`func (o *LogSequenceRest) SetEndlsn(v string)`

SetEndlsn sets Endlsn field to given value.

### HasEndlsn

`func (o *LogSequenceRest) HasEndlsn() bool`

HasEndlsn returns a boolean if a field has been set.

### GetThread

`func (o *LogSequenceRest) GetThread() string`

GetThread returns the Thread field if non-nil, zero value otherwise.

### GetThreadOk

`func (o *LogSequenceRest) GetThreadOk() (*string, bool)`

GetThreadOk returns a tuple with the Thread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThread

`func (o *LogSequenceRest) SetThread(v string)`

SetThread sets Thread field to given value.

### HasThread

`func (o *LogSequenceRest) HasThread() bool`

HasThread returns a boolean if a field has been set.

### GetId

`func (o *LogSequenceRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogSequenceRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogSequenceRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LogSequenceRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *LogSequenceRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LogSequenceRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LogSequenceRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *LogSequenceRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *LogSequenceRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *LogSequenceRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *LogSequenceRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *LogSequenceRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *LogSequenceRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *LogSequenceRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *LogSequenceRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *LogSequenceRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



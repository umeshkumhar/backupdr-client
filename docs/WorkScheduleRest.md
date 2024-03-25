# WorkScheduleRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **string** |  | [optional] 
**Frequency** | Pointer to **string** |  | [optional] 
**Interval** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewWorkScheduleRest

`func NewWorkScheduleRest() *WorkScheduleRest`

NewWorkScheduleRest instantiates a new WorkScheduleRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkScheduleRestWithDefaults

`func NewWorkScheduleRestWithDefaults() *WorkScheduleRest`

NewWorkScheduleRestWithDefaults instantiates a new WorkScheduleRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *WorkScheduleRest) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *WorkScheduleRest) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *WorkScheduleRest) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *WorkScheduleRest) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetFrequency

`func (o *WorkScheduleRest) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *WorkScheduleRest) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *WorkScheduleRest) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *WorkScheduleRest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetInterval

`func (o *WorkScheduleRest) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *WorkScheduleRest) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *WorkScheduleRest) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *WorkScheduleRest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetId

`func (o *WorkScheduleRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkScheduleRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkScheduleRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkScheduleRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *WorkScheduleRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *WorkScheduleRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *WorkScheduleRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *WorkScheduleRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *WorkScheduleRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *WorkScheduleRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *WorkScheduleRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *WorkScheduleRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *WorkScheduleRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *WorkScheduleRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *WorkScheduleRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *WorkScheduleRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



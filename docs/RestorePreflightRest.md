# RestorePreflightRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Testlist** | Pointer to [**[]TestRest**](TestRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewRestorePreflightRest

`func NewRestorePreflightRest() *RestorePreflightRest`

NewRestorePreflightRest instantiates a new RestorePreflightRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestorePreflightRestWithDefaults

`func NewRestorePreflightRestWithDefaults() *RestorePreflightRest`

NewRestorePreflightRestWithDefaults instantiates a new RestorePreflightRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RestorePreflightRest) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RestorePreflightRest) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RestorePreflightRest) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RestorePreflightRest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTestlist

`func (o *RestorePreflightRest) GetTestlist() []TestRest`

GetTestlist returns the Testlist field if non-nil, zero value otherwise.

### GetTestlistOk

`func (o *RestorePreflightRest) GetTestlistOk() (*[]TestRest, bool)`

GetTestlistOk returns a tuple with the Testlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestlist

`func (o *RestorePreflightRest) SetTestlist(v []TestRest)`

SetTestlist sets Testlist field to given value.

### HasTestlist

`func (o *RestorePreflightRest) HasTestlist() bool`

HasTestlist returns a boolean if a field has been set.

### GetId

`func (o *RestorePreflightRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestorePreflightRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestorePreflightRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RestorePreflightRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *RestorePreflightRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RestorePreflightRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RestorePreflightRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RestorePreflightRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *RestorePreflightRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *RestorePreflightRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *RestorePreflightRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *RestorePreflightRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *RestorePreflightRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *RestorePreflightRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *RestorePreflightRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *RestorePreflightRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ArrayTestResultRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**[]IscsiTestResultDetailRest**](IscsiTestResultDetailRest.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 
**Appliance** | Pointer to [**ClusterRest**](ClusterRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewArrayTestResultRest

`func NewArrayTestResultRest() *ArrayTestResultRest`

NewArrayTestResultRest instantiates a new ArrayTestResultRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArrayTestResultRestWithDefaults

`func NewArrayTestResultRestWithDefaults() *ArrayTestResultRest`

NewArrayTestResultRestWithDefaults instantiates a new ArrayTestResultRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ArrayTestResultRest) GetResult() []IscsiTestResultDetailRest`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ArrayTestResultRest) GetResultOk() (*[]IscsiTestResultDetailRest, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ArrayTestResultRest) SetResult(v []IscsiTestResultDetailRest)`

SetResult sets Result field to given value.

### HasResult

`func (o *ArrayTestResultRest) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetError

`func (o *ArrayTestResultRest) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ArrayTestResultRest) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ArrayTestResultRest) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *ArrayTestResultRest) HasError() bool`

HasError returns a boolean if a field has been set.

### GetAppliance

`func (o *ArrayTestResultRest) GetAppliance() ClusterRest`

GetAppliance returns the Appliance field if non-nil, zero value otherwise.

### GetApplianceOk

`func (o *ArrayTestResultRest) GetApplianceOk() (*ClusterRest, bool)`

GetApplianceOk returns a tuple with the Appliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliance

`func (o *ArrayTestResultRest) SetAppliance(v ClusterRest)`

SetAppliance sets Appliance field to given value.

### HasAppliance

`func (o *ArrayTestResultRest) HasAppliance() bool`

HasAppliance returns a boolean if a field has been set.

### GetId

`func (o *ArrayTestResultRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArrayTestResultRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArrayTestResultRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArrayTestResultRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ArrayTestResultRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ArrayTestResultRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ArrayTestResultRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ArrayTestResultRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ArrayTestResultRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ArrayTestResultRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ArrayTestResultRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ArrayTestResultRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ArrayTestResultRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ArrayTestResultRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ArrayTestResultRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ArrayTestResultRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# IscsiTestResultDetailRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Hint** | Pointer to **string** |  | [optional] 
**Test** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewIscsiTestResultDetailRest

`func NewIscsiTestResultDetailRest() *IscsiTestResultDetailRest`

NewIscsiTestResultDetailRest instantiates a new IscsiTestResultDetailRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIscsiTestResultDetailRestWithDefaults

`func NewIscsiTestResultDetailRestWithDefaults() *IscsiTestResultDetailRest`

NewIscsiTestResultDetailRestWithDefaults instantiates a new IscsiTestResultDetailRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *IscsiTestResultDetailRest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IscsiTestResultDetailRest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IscsiTestResultDetailRest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IscsiTestResultDetailRest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHint

`func (o *IscsiTestResultDetailRest) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *IscsiTestResultDetailRest) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *IscsiTestResultDetailRest) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *IscsiTestResultDetailRest) HasHint() bool`

HasHint returns a boolean if a field has been set.

### GetTest

`func (o *IscsiTestResultDetailRest) GetTest() string`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *IscsiTestResultDetailRest) GetTestOk() (*string, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *IscsiTestResultDetailRest) SetTest(v string)`

SetTest sets Test field to given value.

### HasTest

`func (o *IscsiTestResultDetailRest) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetId

`func (o *IscsiTestResultDetailRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IscsiTestResultDetailRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IscsiTestResultDetailRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IscsiTestResultDetailRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *IscsiTestResultDetailRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *IscsiTestResultDetailRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *IscsiTestResultDetailRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *IscsiTestResultDetailRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *IscsiTestResultDetailRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *IscsiTestResultDetailRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *IscsiTestResultDetailRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *IscsiTestResultDetailRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *IscsiTestResultDetailRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *IscsiTestResultDetailRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *IscsiTestResultDetailRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *IscsiTestResultDetailRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



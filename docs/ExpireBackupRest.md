# ExpireBackupRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backuptype** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewExpireBackupRest

`func NewExpireBackupRest() *ExpireBackupRest`

NewExpireBackupRest instantiates a new ExpireBackupRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpireBackupRestWithDefaults

`func NewExpireBackupRestWithDefaults() *ExpireBackupRest`

NewExpireBackupRestWithDefaults instantiates a new ExpireBackupRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackuptype

`func (o *ExpireBackupRest) GetBackuptype() string`

GetBackuptype returns the Backuptype field if non-nil, zero value otherwise.

### GetBackuptypeOk

`func (o *ExpireBackupRest) GetBackuptypeOk() (*string, bool)`

GetBackuptypeOk returns a tuple with the Backuptype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackuptype

`func (o *ExpireBackupRest) SetBackuptype(v string)`

SetBackuptype sets Backuptype field to given value.

### HasBackuptype

`func (o *ExpireBackupRest) HasBackuptype() bool`

HasBackuptype returns a boolean if a field has been set.

### GetId

`func (o *ExpireBackupRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpireBackupRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpireBackupRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExpireBackupRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ExpireBackupRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ExpireBackupRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ExpireBackupRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ExpireBackupRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ExpireBackupRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ExpireBackupRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ExpireBackupRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ExpireBackupRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ExpireBackupRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ExpireBackupRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ExpireBackupRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ExpireBackupRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



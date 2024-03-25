# NfsOptionsRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | Pointer to [**[]NameValueRest**](NameValueRest.md) |  | [optional] 
**Client** | Pointer to [**[]NameValueRest**](NameValueRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewNfsOptionsRest

`func NewNfsOptionsRest() *NfsOptionsRest`

NewNfsOptionsRest instantiates a new NfsOptionsRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfsOptionsRestWithDefaults

`func NewNfsOptionsRestWithDefaults() *NfsOptionsRest`

NewNfsOptionsRestWithDefaults instantiates a new NfsOptionsRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *NfsOptionsRest) GetServer() []NameValueRest`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *NfsOptionsRest) GetServerOk() (*[]NameValueRest, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *NfsOptionsRest) SetServer(v []NameValueRest)`

SetServer sets Server field to given value.

### HasServer

`func (o *NfsOptionsRest) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetClient

`func (o *NfsOptionsRest) GetClient() []NameValueRest`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *NfsOptionsRest) GetClientOk() (*[]NameValueRest, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *NfsOptionsRest) SetClient(v []NameValueRest)`

SetClient sets Client field to given value.

### HasClient

`func (o *NfsOptionsRest) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetId

`func (o *NfsOptionsRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NfsOptionsRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NfsOptionsRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NfsOptionsRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *NfsOptionsRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NfsOptionsRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NfsOptionsRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *NfsOptionsRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *NfsOptionsRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *NfsOptionsRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *NfsOptionsRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *NfsOptionsRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *NfsOptionsRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *NfsOptionsRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *NfsOptionsRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *NfsOptionsRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



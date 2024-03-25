# ListNameValueRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListNameValueRest** | Pointer to [**[]NameValueRest**](NameValueRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewListNameValueRest

`func NewListNameValueRest() *ListNameValueRest`

NewListNameValueRest instantiates a new ListNameValueRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNameValueRestWithDefaults

`func NewListNameValueRestWithDefaults() *ListNameValueRest`

NewListNameValueRestWithDefaults instantiates a new ListNameValueRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListNameValueRest

`func (o *ListNameValueRest) GetListNameValueRest() []NameValueRest`

GetListNameValueRest returns the ListNameValueRest field if non-nil, zero value otherwise.

### GetListNameValueRestOk

`func (o *ListNameValueRest) GetListNameValueRestOk() (*[]NameValueRest, bool)`

GetListNameValueRestOk returns a tuple with the ListNameValueRest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListNameValueRest

`func (o *ListNameValueRest) SetListNameValueRest(v []NameValueRest)`

SetListNameValueRest sets ListNameValueRest field to given value.

### HasListNameValueRest

`func (o *ListNameValueRest) HasListNameValueRest() bool`

HasListNameValueRest returns a boolean if a field has been set.

### GetId

`func (o *ListNameValueRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListNameValueRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListNameValueRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListNameValueRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ListNameValueRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ListNameValueRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ListNameValueRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ListNameValueRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ListNameValueRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ListNameValueRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ListNameValueRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ListNameValueRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ListNameValueRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ListNameValueRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ListNameValueRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ListNameValueRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



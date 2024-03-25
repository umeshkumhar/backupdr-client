# RestoreLocationRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mapping** | Pointer to [**[]LocationMappingItemRest**](LocationMappingItemRest.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewRestoreLocationRest

`func NewRestoreLocationRest() *RestoreLocationRest`

NewRestoreLocationRest instantiates a new RestoreLocationRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreLocationRestWithDefaults

`func NewRestoreLocationRestWithDefaults() *RestoreLocationRest`

NewRestoreLocationRestWithDefaults instantiates a new RestoreLocationRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMapping

`func (o *RestoreLocationRest) GetMapping() []LocationMappingItemRest`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *RestoreLocationRest) GetMappingOk() (*[]LocationMappingItemRest, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *RestoreLocationRest) SetMapping(v []LocationMappingItemRest)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *RestoreLocationRest) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### GetType

`func (o *RestoreLocationRest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestoreLocationRest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestoreLocationRest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RestoreLocationRest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *RestoreLocationRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestoreLocationRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestoreLocationRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RestoreLocationRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *RestoreLocationRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RestoreLocationRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RestoreLocationRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RestoreLocationRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *RestoreLocationRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *RestoreLocationRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *RestoreLocationRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *RestoreLocationRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *RestoreLocationRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *RestoreLocationRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *RestoreLocationRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *RestoreLocationRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



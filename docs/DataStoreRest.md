# DataStoreRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Capacity** | Pointer to **string** |  | [optional] 
**Freespace** | Pointer to **string** |  | [optional] 
**RdmSupported** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewDataStoreRest

`func NewDataStoreRest() *DataStoreRest`

NewDataStoreRest instantiates a new DataStoreRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStoreRestWithDefaults

`func NewDataStoreRestWithDefaults() *DataStoreRest`

NewDataStoreRestWithDefaults instantiates a new DataStoreRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DataStoreRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataStoreRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataStoreRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataStoreRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *DataStoreRest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataStoreRest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataStoreRest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DataStoreRest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCapacity

`func (o *DataStoreRest) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *DataStoreRest) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *DataStoreRest) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *DataStoreRest) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetFreespace

`func (o *DataStoreRest) GetFreespace() string`

GetFreespace returns the Freespace field if non-nil, zero value otherwise.

### GetFreespaceOk

`func (o *DataStoreRest) GetFreespaceOk() (*string, bool)`

GetFreespaceOk returns a tuple with the Freespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreespace

`func (o *DataStoreRest) SetFreespace(v string)`

SetFreespace sets Freespace field to given value.

### HasFreespace

`func (o *DataStoreRest) HasFreespace() bool`

HasFreespace returns a boolean if a field has been set.

### GetRdmSupported

`func (o *DataStoreRest) GetRdmSupported() bool`

GetRdmSupported returns the RdmSupported field if non-nil, zero value otherwise.

### GetRdmSupportedOk

`func (o *DataStoreRest) GetRdmSupportedOk() (*bool, bool)`

GetRdmSupportedOk returns a tuple with the RdmSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdmSupported

`func (o *DataStoreRest) SetRdmSupported(v bool)`

SetRdmSupported sets RdmSupported field to given value.

### HasRdmSupported

`func (o *DataStoreRest) HasRdmSupported() bool`

HasRdmSupported returns a boolean if a field has been set.

### GetId

`func (o *DataStoreRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataStoreRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataStoreRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataStoreRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *DataStoreRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DataStoreRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DataStoreRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *DataStoreRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *DataStoreRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *DataStoreRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *DataStoreRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *DataStoreRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *DataStoreRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *DataStoreRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *DataStoreRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *DataStoreRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



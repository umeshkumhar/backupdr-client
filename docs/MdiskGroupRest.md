# MdiskGroupRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Capacity** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewMdiskGroupRest

`func NewMdiskGroupRest() *MdiskGroupRest`

NewMdiskGroupRest instantiates a new MdiskGroupRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMdiskGroupRestWithDefaults

`func NewMdiskGroupRestWithDefaults() *MdiskGroupRest`

NewMdiskGroupRestWithDefaults instantiates a new MdiskGroupRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MdiskGroupRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MdiskGroupRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MdiskGroupRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MdiskGroupRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCapacity

`func (o *MdiskGroupRest) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *MdiskGroupRest) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *MdiskGroupRest) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *MdiskGroupRest) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetId

`func (o *MdiskGroupRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MdiskGroupRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MdiskGroupRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MdiskGroupRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *MdiskGroupRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *MdiskGroupRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *MdiskGroupRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *MdiskGroupRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *MdiskGroupRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *MdiskGroupRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *MdiskGroupRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *MdiskGroupRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *MdiskGroupRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *MdiskGroupRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *MdiskGroupRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *MdiskGroupRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# LocationMappingItemRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewLocationMappingItemRest

`func NewLocationMappingItemRest() *LocationMappingItemRest`

NewLocationMappingItemRest instantiates a new LocationMappingItemRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationMappingItemRestWithDefaults

`func NewLocationMappingItemRestWithDefaults() *LocationMappingItemRest`

NewLocationMappingItemRestWithDefaults instantiates a new LocationMappingItemRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *LocationMappingItemRest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LocationMappingItemRest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LocationMappingItemRest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *LocationMappingItemRest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetName

`func (o *LocationMappingItemRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocationMappingItemRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocationMappingItemRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LocationMappingItemRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTarget

`func (o *LocationMappingItemRest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *LocationMappingItemRest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *LocationMappingItemRest) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *LocationMappingItemRest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetId

`func (o *LocationMappingItemRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocationMappingItemRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocationMappingItemRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LocationMappingItemRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *LocationMappingItemRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LocationMappingItemRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LocationMappingItemRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *LocationMappingItemRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *LocationMappingItemRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *LocationMappingItemRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *LocationMappingItemRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *LocationMappingItemRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *LocationMappingItemRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *LocationMappingItemRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *LocationMappingItemRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *LocationMappingItemRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



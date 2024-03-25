# FeatureFlagsRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | feature value | [optional] 
**Featurename** | Pointer to **string** | Featurename | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewFeatureFlagsRest

`func NewFeatureFlagsRest() *FeatureFlagsRest`

NewFeatureFlagsRest instantiates a new FeatureFlagsRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagsRestWithDefaults

`func NewFeatureFlagsRestWithDefaults() *FeatureFlagsRest`

NewFeatureFlagsRestWithDefaults instantiates a new FeatureFlagsRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *FeatureFlagsRest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FeatureFlagsRest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FeatureFlagsRest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *FeatureFlagsRest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetFeaturename

`func (o *FeatureFlagsRest) GetFeaturename() string`

GetFeaturename returns the Featurename field if non-nil, zero value otherwise.

### GetFeaturenameOk

`func (o *FeatureFlagsRest) GetFeaturenameOk() (*string, bool)`

GetFeaturenameOk returns a tuple with the Featurename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturename

`func (o *FeatureFlagsRest) SetFeaturename(v string)`

SetFeaturename sets Featurename field to given value.

### HasFeaturename

`func (o *FeatureFlagsRest) HasFeaturename() bool`

HasFeaturename returns a boolean if a field has been set.

### GetId

`func (o *FeatureFlagsRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeatureFlagsRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeatureFlagsRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FeatureFlagsRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *FeatureFlagsRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FeatureFlagsRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FeatureFlagsRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FeatureFlagsRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *FeatureFlagsRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *FeatureFlagsRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *FeatureFlagsRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *FeatureFlagsRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *FeatureFlagsRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *FeatureFlagsRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *FeatureFlagsRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *FeatureFlagsRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



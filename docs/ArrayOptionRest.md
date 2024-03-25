# ArrayOptionRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Optiontype** | Pointer to **string** |  | [optional] 
**Updatable** | Pointer to **bool** |  | [optional] 
**Needmask** | Pointer to **bool** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Max** | Pointer to **int64** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Valuetype** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewArrayOptionRest

`func NewArrayOptionRest() *ArrayOptionRest`

NewArrayOptionRest instantiates a new ArrayOptionRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArrayOptionRestWithDefaults

`func NewArrayOptionRestWithDefaults() *ArrayOptionRest`

NewArrayOptionRestWithDefaults instantiates a new ArrayOptionRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptiontype

`func (o *ArrayOptionRest) GetOptiontype() string`

GetOptiontype returns the Optiontype field if non-nil, zero value otherwise.

### GetOptiontypeOk

`func (o *ArrayOptionRest) GetOptiontypeOk() (*string, bool)`

GetOptiontypeOk returns a tuple with the Optiontype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptiontype

`func (o *ArrayOptionRest) SetOptiontype(v string)`

SetOptiontype sets Optiontype field to given value.

### HasOptiontype

`func (o *ArrayOptionRest) HasOptiontype() bool`

HasOptiontype returns a boolean if a field has been set.

### GetUpdatable

`func (o *ArrayOptionRest) GetUpdatable() bool`

GetUpdatable returns the Updatable field if non-nil, zero value otherwise.

### GetUpdatableOk

`func (o *ArrayOptionRest) GetUpdatableOk() (*bool, bool)`

GetUpdatableOk returns a tuple with the Updatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatable

`func (o *ArrayOptionRest) SetUpdatable(v bool)`

SetUpdatable sets Updatable field to given value.

### HasUpdatable

`func (o *ArrayOptionRest) HasUpdatable() bool`

HasUpdatable returns a boolean if a field has been set.

### GetNeedmask

`func (o *ArrayOptionRest) GetNeedmask() bool`

GetNeedmask returns the Needmask field if non-nil, zero value otherwise.

### GetNeedmaskOk

`func (o *ArrayOptionRest) GetNeedmaskOk() (*bool, bool)`

GetNeedmaskOk returns a tuple with the Needmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedmask

`func (o *ArrayOptionRest) SetNeedmask(v bool)`

SetNeedmask sets Needmask field to given value.

### HasNeedmask

`func (o *ArrayOptionRest) HasNeedmask() bool`

HasNeedmask returns a boolean if a field has been set.

### GetRequired

`func (o *ArrayOptionRest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ArrayOptionRest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ArrayOptionRest) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ArrayOptionRest) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetName

`func (o *ArrayOptionRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArrayOptionRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArrayOptionRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArrayOptionRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMax

`func (o *ArrayOptionRest) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ArrayOptionRest) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ArrayOptionRest) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *ArrayOptionRest) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetLabel

`func (o *ArrayOptionRest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ArrayOptionRest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ArrayOptionRest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ArrayOptionRest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetValuetype

`func (o *ArrayOptionRest) GetValuetype() string`

GetValuetype returns the Valuetype field if non-nil, zero value otherwise.

### GetValuetypeOk

`func (o *ArrayOptionRest) GetValuetypeOk() (*string, bool)`

GetValuetypeOk returns a tuple with the Valuetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuetype

`func (o *ArrayOptionRest) SetValuetype(v string)`

SetValuetype sets Valuetype field to given value.

### HasValuetype

`func (o *ArrayOptionRest) HasValuetype() bool`

HasValuetype returns a boolean if a field has been set.

### GetId

`func (o *ArrayOptionRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArrayOptionRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArrayOptionRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArrayOptionRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ArrayOptionRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ArrayOptionRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ArrayOptionRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ArrayOptionRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ArrayOptionRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ArrayOptionRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ArrayOptionRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ArrayOptionRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ArrayOptionRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ArrayOptionRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ArrayOptionRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ArrayOptionRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



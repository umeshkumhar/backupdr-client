# ScriptFileRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Required** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewScriptFileRest

`func NewScriptFileRest() *ScriptFileRest`

NewScriptFileRest instantiates a new ScriptFileRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptFileRestWithDefaults

`func NewScriptFileRestWithDefaults() *ScriptFileRest`

NewScriptFileRestWithDefaults instantiates a new ScriptFileRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequired

`func (o *ScriptFileRest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ScriptFileRest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ScriptFileRest) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ScriptFileRest) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetName

`func (o *ScriptFileRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScriptFileRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScriptFileRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScriptFileRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *ScriptFileRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScriptFileRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScriptFileRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScriptFileRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ScriptFileRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ScriptFileRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ScriptFileRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ScriptFileRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ScriptFileRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ScriptFileRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ScriptFileRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ScriptFileRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ScriptFileRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ScriptFileRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ScriptFileRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ScriptFileRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



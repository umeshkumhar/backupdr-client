# UnmountRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | Pointer to **string** |  | [optional] 
**Delete** | Pointer to **bool** |  | [optional] 
**Force** | Pointer to **bool** |  | [optional] 
**Preservevm** | Pointer to **bool** |  | [optional] 
**Scripts** | Pointer to [**[]ScriptRest**](ScriptRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewUnmountRest

`func NewUnmountRest() *UnmountRest`

NewUnmountRest instantiates a new UnmountRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnmountRestWithDefaults

`func NewUnmountRestWithDefaults() *UnmountRest`

NewUnmountRestWithDefaults instantiates a new UnmountRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScript

`func (o *UnmountRest) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *UnmountRest) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *UnmountRest) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *UnmountRest) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetDelete

`func (o *UnmountRest) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *UnmountRest) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *UnmountRest) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *UnmountRest) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetForce

`func (o *UnmountRest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *UnmountRest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *UnmountRest) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *UnmountRest) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetPreservevm

`func (o *UnmountRest) GetPreservevm() bool`

GetPreservevm returns the Preservevm field if non-nil, zero value otherwise.

### GetPreservevmOk

`func (o *UnmountRest) GetPreservevmOk() (*bool, bool)`

GetPreservevmOk returns a tuple with the Preservevm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreservevm

`func (o *UnmountRest) SetPreservevm(v bool)`

SetPreservevm sets Preservevm field to given value.

### HasPreservevm

`func (o *UnmountRest) HasPreservevm() bool`

HasPreservevm returns a boolean if a field has been set.

### GetScripts

`func (o *UnmountRest) GetScripts() []ScriptRest`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *UnmountRest) GetScriptsOk() (*[]ScriptRest, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *UnmountRest) SetScripts(v []ScriptRest)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *UnmountRest) HasScripts() bool`

HasScripts returns a boolean if a field has been set.

### GetId

`func (o *UnmountRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnmountRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnmountRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UnmountRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *UnmountRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *UnmountRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *UnmountRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *UnmountRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *UnmountRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *UnmountRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *UnmountRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *UnmountRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *UnmountRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *UnmountRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *UnmountRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *UnmountRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



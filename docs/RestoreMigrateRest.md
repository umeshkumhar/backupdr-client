# RestoreMigrateRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Restoreoptions** | Pointer to [**PreflightRestoreOptionsRest**](PreflightRestoreOptionsRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewRestoreMigrateRest

`func NewRestoreMigrateRest() *RestoreMigrateRest`

NewRestoreMigrateRest instantiates a new RestoreMigrateRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreMigrateRestWithDefaults

`func NewRestoreMigrateRestWithDefaults() *RestoreMigrateRest`

NewRestoreMigrateRestWithDefaults instantiates a new RestoreMigrateRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RestoreMigrateRest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RestoreMigrateRest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RestoreMigrateRest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *RestoreMigrateRest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetLabel

`func (o *RestoreMigrateRest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RestoreMigrateRest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RestoreMigrateRest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *RestoreMigrateRest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRestoreoptions

`func (o *RestoreMigrateRest) GetRestoreoptions() PreflightRestoreOptionsRest`

GetRestoreoptions returns the Restoreoptions field if non-nil, zero value otherwise.

### GetRestoreoptionsOk

`func (o *RestoreMigrateRest) GetRestoreoptionsOk() (*PreflightRestoreOptionsRest, bool)`

GetRestoreoptionsOk returns a tuple with the Restoreoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreoptions

`func (o *RestoreMigrateRest) SetRestoreoptions(v PreflightRestoreOptionsRest)`

SetRestoreoptions sets Restoreoptions field to given value.

### HasRestoreoptions

`func (o *RestoreMigrateRest) HasRestoreoptions() bool`

HasRestoreoptions returns a boolean if a field has been set.

### GetId

`func (o *RestoreMigrateRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestoreMigrateRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestoreMigrateRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RestoreMigrateRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *RestoreMigrateRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RestoreMigrateRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RestoreMigrateRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RestoreMigrateRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *RestoreMigrateRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *RestoreMigrateRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *RestoreMigrateRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *RestoreMigrateRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *RestoreMigrateRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *RestoreMigrateRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *RestoreMigrateRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *RestoreMigrateRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PreflightRestoreOptionsRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ctrltarget** | Pointer to **string** |  | [optional] 
**Redotarget** | Pointer to **string** |  | [optional] 
**Targettype** | Pointer to **string** |  | [optional] 
**Standalone** | Pointer to **string** |  | [optional] 
**Migrationmapping** | Pointer to **string** |  | [optional] 
**Copythreadcount** | Pointer to **string** |  | [optional] 
**Asmracnodelist** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewPreflightRestoreOptionsRest

`func NewPreflightRestoreOptionsRest() *PreflightRestoreOptionsRest`

NewPreflightRestoreOptionsRest instantiates a new PreflightRestoreOptionsRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreflightRestoreOptionsRestWithDefaults

`func NewPreflightRestoreOptionsRestWithDefaults() *PreflightRestoreOptionsRest`

NewPreflightRestoreOptionsRestWithDefaults instantiates a new PreflightRestoreOptionsRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCtrltarget

`func (o *PreflightRestoreOptionsRest) GetCtrltarget() string`

GetCtrltarget returns the Ctrltarget field if non-nil, zero value otherwise.

### GetCtrltargetOk

`func (o *PreflightRestoreOptionsRest) GetCtrltargetOk() (*string, bool)`

GetCtrltargetOk returns a tuple with the Ctrltarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtrltarget

`func (o *PreflightRestoreOptionsRest) SetCtrltarget(v string)`

SetCtrltarget sets Ctrltarget field to given value.

### HasCtrltarget

`func (o *PreflightRestoreOptionsRest) HasCtrltarget() bool`

HasCtrltarget returns a boolean if a field has been set.

### GetRedotarget

`func (o *PreflightRestoreOptionsRest) GetRedotarget() string`

GetRedotarget returns the Redotarget field if non-nil, zero value otherwise.

### GetRedotargetOk

`func (o *PreflightRestoreOptionsRest) GetRedotargetOk() (*string, bool)`

GetRedotargetOk returns a tuple with the Redotarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedotarget

`func (o *PreflightRestoreOptionsRest) SetRedotarget(v string)`

SetRedotarget sets Redotarget field to given value.

### HasRedotarget

`func (o *PreflightRestoreOptionsRest) HasRedotarget() bool`

HasRedotarget returns a boolean if a field has been set.

### GetTargettype

`func (o *PreflightRestoreOptionsRest) GetTargettype() string`

GetTargettype returns the Targettype field if non-nil, zero value otherwise.

### GetTargettypeOk

`func (o *PreflightRestoreOptionsRest) GetTargettypeOk() (*string, bool)`

GetTargettypeOk returns a tuple with the Targettype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargettype

`func (o *PreflightRestoreOptionsRest) SetTargettype(v string)`

SetTargettype sets Targettype field to given value.

### HasTargettype

`func (o *PreflightRestoreOptionsRest) HasTargettype() bool`

HasTargettype returns a boolean if a field has been set.

### GetStandalone

`func (o *PreflightRestoreOptionsRest) GetStandalone() string`

GetStandalone returns the Standalone field if non-nil, zero value otherwise.

### GetStandaloneOk

`func (o *PreflightRestoreOptionsRest) GetStandaloneOk() (*string, bool)`

GetStandaloneOk returns a tuple with the Standalone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandalone

`func (o *PreflightRestoreOptionsRest) SetStandalone(v string)`

SetStandalone sets Standalone field to given value.

### HasStandalone

`func (o *PreflightRestoreOptionsRest) HasStandalone() bool`

HasStandalone returns a boolean if a field has been set.

### GetMigrationmapping

`func (o *PreflightRestoreOptionsRest) GetMigrationmapping() string`

GetMigrationmapping returns the Migrationmapping field if non-nil, zero value otherwise.

### GetMigrationmappingOk

`func (o *PreflightRestoreOptionsRest) GetMigrationmappingOk() (*string, bool)`

GetMigrationmappingOk returns a tuple with the Migrationmapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationmapping

`func (o *PreflightRestoreOptionsRest) SetMigrationmapping(v string)`

SetMigrationmapping sets Migrationmapping field to given value.

### HasMigrationmapping

`func (o *PreflightRestoreOptionsRest) HasMigrationmapping() bool`

HasMigrationmapping returns a boolean if a field has been set.

### GetCopythreadcount

`func (o *PreflightRestoreOptionsRest) GetCopythreadcount() string`

GetCopythreadcount returns the Copythreadcount field if non-nil, zero value otherwise.

### GetCopythreadcountOk

`func (o *PreflightRestoreOptionsRest) GetCopythreadcountOk() (*string, bool)`

GetCopythreadcountOk returns a tuple with the Copythreadcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopythreadcount

`func (o *PreflightRestoreOptionsRest) SetCopythreadcount(v string)`

SetCopythreadcount sets Copythreadcount field to given value.

### HasCopythreadcount

`func (o *PreflightRestoreOptionsRest) HasCopythreadcount() bool`

HasCopythreadcount returns a boolean if a field has been set.

### GetAsmracnodelist

`func (o *PreflightRestoreOptionsRest) GetAsmracnodelist() string`

GetAsmracnodelist returns the Asmracnodelist field if non-nil, zero value otherwise.

### GetAsmracnodelistOk

`func (o *PreflightRestoreOptionsRest) GetAsmracnodelistOk() (*string, bool)`

GetAsmracnodelistOk returns a tuple with the Asmracnodelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmracnodelist

`func (o *PreflightRestoreOptionsRest) SetAsmracnodelist(v string)`

SetAsmracnodelist sets Asmracnodelist field to given value.

### HasAsmracnodelist

`func (o *PreflightRestoreOptionsRest) HasAsmracnodelist() bool`

HasAsmracnodelist returns a boolean if a field has been set.

### GetId

`func (o *PreflightRestoreOptionsRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PreflightRestoreOptionsRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PreflightRestoreOptionsRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PreflightRestoreOptionsRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *PreflightRestoreOptionsRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PreflightRestoreOptionsRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PreflightRestoreOptionsRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PreflightRestoreOptionsRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *PreflightRestoreOptionsRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *PreflightRestoreOptionsRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *PreflightRestoreOptionsRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *PreflightRestoreOptionsRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *PreflightRestoreOptionsRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *PreflightRestoreOptionsRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *PreflightRestoreOptionsRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *PreflightRestoreOptionsRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



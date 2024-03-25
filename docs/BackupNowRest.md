# BackupNowRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to [**PolicyRest**](PolicyRest.md) |  | [optional] 
**Script** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Options** | Pointer to **string** |  | [optional] 
**Backuptype** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewBackupNowRest

`func NewBackupNowRest() *BackupNowRest`

NewBackupNowRest instantiates a new BackupNowRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupNowRestWithDefaults

`func NewBackupNowRestWithDefaults() *BackupNowRest`

NewBackupNowRestWithDefaults instantiates a new BackupNowRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *BackupNowRest) GetPolicy() PolicyRest`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *BackupNowRest) GetPolicyOk() (*PolicyRest, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *BackupNowRest) SetPolicy(v PolicyRest)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *BackupNowRest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetScript

`func (o *BackupNowRest) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *BackupNowRest) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *BackupNowRest) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *BackupNowRest) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetLabel

`func (o *BackupNowRest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BackupNowRest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BackupNowRest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BackupNowRest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetOptions

`func (o *BackupNowRest) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *BackupNowRest) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *BackupNowRest) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *BackupNowRest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetBackuptype

`func (o *BackupNowRest) GetBackuptype() string`

GetBackuptype returns the Backuptype field if non-nil, zero value otherwise.

### GetBackuptypeOk

`func (o *BackupNowRest) GetBackuptypeOk() (*string, bool)`

GetBackuptypeOk returns a tuple with the Backuptype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackuptype

`func (o *BackupNowRest) SetBackuptype(v string)`

SetBackuptype sets Backuptype field to given value.

### HasBackuptype

`func (o *BackupNowRest) HasBackuptype() bool`

HasBackuptype returns a boolean if a field has been set.

### GetId

`func (o *BackupNowRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupNowRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupNowRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BackupNowRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *BackupNowRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BackupNowRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BackupNowRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BackupNowRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *BackupNowRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *BackupNowRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *BackupNowRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *BackupNowRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *BackupNowRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *BackupNowRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *BackupNowRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *BackupNowRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



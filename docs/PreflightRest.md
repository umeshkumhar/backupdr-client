# PreflightRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Databasesid** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Orahome** | Pointer to **string** |  | [optional] 
**Hostid** | Pointer to **string** |  | [optional] 
**Restoreoptions** | Pointer to [**PreflightRestoreOptionsRest**](PreflightRestoreOptionsRest.md) |  | [optional] 
**Provisioningoptions** | Pointer to [**[]ProvisioningOptionRest**](ProvisioningOptionRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewPreflightRest

`func NewPreflightRest() *PreflightRest`

NewPreflightRest instantiates a new PreflightRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreflightRestWithDefaults

`func NewPreflightRestWithDefaults() *PreflightRest`

NewPreflightRestWithDefaults instantiates a new PreflightRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabasesid

`func (o *PreflightRest) GetDatabasesid() string`

GetDatabasesid returns the Databasesid field if non-nil, zero value otherwise.

### GetDatabasesidOk

`func (o *PreflightRest) GetDatabasesidOk() (*string, bool)`

GetDatabasesidOk returns a tuple with the Databasesid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabasesid

`func (o *PreflightRest) SetDatabasesid(v string)`

SetDatabasesid sets Databasesid field to given value.

### HasDatabasesid

`func (o *PreflightRest) HasDatabasesid() bool`

HasDatabasesid returns a boolean if a field has been set.

### GetUsername

`func (o *PreflightRest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PreflightRest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PreflightRest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PreflightRest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetOrahome

`func (o *PreflightRest) GetOrahome() string`

GetOrahome returns the Orahome field if non-nil, zero value otherwise.

### GetOrahomeOk

`func (o *PreflightRest) GetOrahomeOk() (*string, bool)`

GetOrahomeOk returns a tuple with the Orahome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrahome

`func (o *PreflightRest) SetOrahome(v string)`

SetOrahome sets Orahome field to given value.

### HasOrahome

`func (o *PreflightRest) HasOrahome() bool`

HasOrahome returns a boolean if a field has been set.

### GetHostid

`func (o *PreflightRest) GetHostid() string`

GetHostid returns the Hostid field if non-nil, zero value otherwise.

### GetHostidOk

`func (o *PreflightRest) GetHostidOk() (*string, bool)`

GetHostidOk returns a tuple with the Hostid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostid

`func (o *PreflightRest) SetHostid(v string)`

SetHostid sets Hostid field to given value.

### HasHostid

`func (o *PreflightRest) HasHostid() bool`

HasHostid returns a boolean if a field has been set.

### GetRestoreoptions

`func (o *PreflightRest) GetRestoreoptions() PreflightRestoreOptionsRest`

GetRestoreoptions returns the Restoreoptions field if non-nil, zero value otherwise.

### GetRestoreoptionsOk

`func (o *PreflightRest) GetRestoreoptionsOk() (*PreflightRestoreOptionsRest, bool)`

GetRestoreoptionsOk returns a tuple with the Restoreoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreoptions

`func (o *PreflightRest) SetRestoreoptions(v PreflightRestoreOptionsRest)`

SetRestoreoptions sets Restoreoptions field to given value.

### HasRestoreoptions

`func (o *PreflightRest) HasRestoreoptions() bool`

HasRestoreoptions returns a boolean if a field has been set.

### GetProvisioningoptions

`func (o *PreflightRest) GetProvisioningoptions() []ProvisioningOptionRest`

GetProvisioningoptions returns the Provisioningoptions field if non-nil, zero value otherwise.

### GetProvisioningoptionsOk

`func (o *PreflightRest) GetProvisioningoptionsOk() (*[]ProvisioningOptionRest, bool)`

GetProvisioningoptionsOk returns a tuple with the Provisioningoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningoptions

`func (o *PreflightRest) SetProvisioningoptions(v []ProvisioningOptionRest)`

SetProvisioningoptions sets Provisioningoptions field to given value.

### HasProvisioningoptions

`func (o *PreflightRest) HasProvisioningoptions() bool`

HasProvisioningoptions returns a boolean if a field has been set.

### GetId

`func (o *PreflightRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PreflightRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PreflightRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PreflightRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *PreflightRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PreflightRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PreflightRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PreflightRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *PreflightRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *PreflightRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *PreflightRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *PreflightRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *PreflightRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *PreflightRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *PreflightRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *PreflightRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



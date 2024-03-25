# ApplianceUpdateLogsRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to **[]string** |  | [optional] 
**Logfilename** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewApplianceUpdateLogsRest

`func NewApplianceUpdateLogsRest() *ApplianceUpdateLogsRest`

NewApplianceUpdateLogsRest instantiates a new ApplianceUpdateLogsRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceUpdateLogsRestWithDefaults

`func NewApplianceUpdateLogsRestWithDefaults() *ApplianceUpdateLogsRest`

NewApplianceUpdateLogsRestWithDefaults instantiates a new ApplianceUpdateLogsRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *ApplianceUpdateLogsRest) GetLogs() []string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *ApplianceUpdateLogsRest) GetLogsOk() (*[]string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *ApplianceUpdateLogsRest) SetLogs(v []string)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *ApplianceUpdateLogsRest) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetLogfilename

`func (o *ApplianceUpdateLogsRest) GetLogfilename() string`

GetLogfilename returns the Logfilename field if non-nil, zero value otherwise.

### GetLogfilenameOk

`func (o *ApplianceUpdateLogsRest) GetLogfilenameOk() (*string, bool)`

GetLogfilenameOk returns a tuple with the Logfilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogfilename

`func (o *ApplianceUpdateLogsRest) SetLogfilename(v string)`

SetLogfilename sets Logfilename field to given value.

### HasLogfilename

`func (o *ApplianceUpdateLogsRest) HasLogfilename() bool`

HasLogfilename returns a boolean if a field has been set.

### GetId

`func (o *ApplianceUpdateLogsRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplianceUpdateLogsRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplianceUpdateLogsRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplianceUpdateLogsRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ApplianceUpdateLogsRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ApplianceUpdateLogsRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ApplianceUpdateLogsRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ApplianceUpdateLogsRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ApplianceUpdateLogsRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ApplianceUpdateLogsRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ApplianceUpdateLogsRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ApplianceUpdateLogsRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ApplianceUpdateLogsRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ApplianceUpdateLogsRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ApplianceUpdateLogsRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ApplianceUpdateLogsRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



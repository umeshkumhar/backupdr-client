# ExportRuleInfoRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exporttype** | Pointer to **string** |  | [optional] 
**Exportedname** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**[]UserInfoRest**](UserInfoRest.md) |  | [optional] 
**Client** | Pointer to [**[]ClientInfoRest**](ClientInfoRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewExportRuleInfoRest

`func NewExportRuleInfoRest() *ExportRuleInfoRest`

NewExportRuleInfoRest instantiates a new ExportRuleInfoRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportRuleInfoRestWithDefaults

`func NewExportRuleInfoRestWithDefaults() *ExportRuleInfoRest`

NewExportRuleInfoRestWithDefaults instantiates a new ExportRuleInfoRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExporttype

`func (o *ExportRuleInfoRest) GetExporttype() string`

GetExporttype returns the Exporttype field if non-nil, zero value otherwise.

### GetExporttypeOk

`func (o *ExportRuleInfoRest) GetExporttypeOk() (*string, bool)`

GetExporttypeOk returns a tuple with the Exporttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExporttype

`func (o *ExportRuleInfoRest) SetExporttype(v string)`

SetExporttype sets Exporttype field to given value.

### HasExporttype

`func (o *ExportRuleInfoRest) HasExporttype() bool`

HasExporttype returns a boolean if a field has been set.

### GetExportedname

`func (o *ExportRuleInfoRest) GetExportedname() string`

GetExportedname returns the Exportedname field if non-nil, zero value otherwise.

### GetExportednameOk

`func (o *ExportRuleInfoRest) GetExportednameOk() (*string, bool)`

GetExportednameOk returns a tuple with the Exportedname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedname

`func (o *ExportRuleInfoRest) SetExportedname(v string)`

SetExportedname sets Exportedname field to given value.

### HasExportedname

`func (o *ExportRuleInfoRest) HasExportedname() bool`

HasExportedname returns a boolean if a field has been set.

### GetUser

`func (o *ExportRuleInfoRest) GetUser() []UserInfoRest`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ExportRuleInfoRest) GetUserOk() (*[]UserInfoRest, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ExportRuleInfoRest) SetUser(v []UserInfoRest)`

SetUser sets User field to given value.

### HasUser

`func (o *ExportRuleInfoRest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetClient

`func (o *ExportRuleInfoRest) GetClient() []ClientInfoRest`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *ExportRuleInfoRest) GetClientOk() (*[]ClientInfoRest, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *ExportRuleInfoRest) SetClient(v []ClientInfoRest)`

SetClient sets Client field to given value.

### HasClient

`func (o *ExportRuleInfoRest) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetId

`func (o *ExportRuleInfoRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExportRuleInfoRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExportRuleInfoRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExportRuleInfoRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ExportRuleInfoRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ExportRuleInfoRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ExportRuleInfoRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ExportRuleInfoRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ExportRuleInfoRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ExportRuleInfoRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ExportRuleInfoRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ExportRuleInfoRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ExportRuleInfoRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ExportRuleInfoRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ExportRuleInfoRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ExportRuleInfoRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



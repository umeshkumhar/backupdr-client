# ExportVolumeInfoRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exportrule** | Pointer to [**[]ExportRuleInfoRest**](ExportRuleInfoRest.md) |  | [optional] 
**Exportedpath** | Pointer to **string** |  | [optional] 
**Objecttype** | Pointer to **string** |  | [optional] 
**Objectid** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewExportVolumeInfoRest

`func NewExportVolumeInfoRest() *ExportVolumeInfoRest`

NewExportVolumeInfoRest instantiates a new ExportVolumeInfoRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportVolumeInfoRestWithDefaults

`func NewExportVolumeInfoRestWithDefaults() *ExportVolumeInfoRest`

NewExportVolumeInfoRestWithDefaults instantiates a new ExportVolumeInfoRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportrule

`func (o *ExportVolumeInfoRest) GetExportrule() []ExportRuleInfoRest`

GetExportrule returns the Exportrule field if non-nil, zero value otherwise.

### GetExportruleOk

`func (o *ExportVolumeInfoRest) GetExportruleOk() (*[]ExportRuleInfoRest, bool)`

GetExportruleOk returns a tuple with the Exportrule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportrule

`func (o *ExportVolumeInfoRest) SetExportrule(v []ExportRuleInfoRest)`

SetExportrule sets Exportrule field to given value.

### HasExportrule

`func (o *ExportVolumeInfoRest) HasExportrule() bool`

HasExportrule returns a boolean if a field has been set.

### GetExportedpath

`func (o *ExportVolumeInfoRest) GetExportedpath() string`

GetExportedpath returns the Exportedpath field if non-nil, zero value otherwise.

### GetExportedpathOk

`func (o *ExportVolumeInfoRest) GetExportedpathOk() (*string, bool)`

GetExportedpathOk returns a tuple with the Exportedpath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedpath

`func (o *ExportVolumeInfoRest) SetExportedpath(v string)`

SetExportedpath sets Exportedpath field to given value.

### HasExportedpath

`func (o *ExportVolumeInfoRest) HasExportedpath() bool`

HasExportedpath returns a boolean if a field has been set.

### GetObjecttype

`func (o *ExportVolumeInfoRest) GetObjecttype() string`

GetObjecttype returns the Objecttype field if non-nil, zero value otherwise.

### GetObjecttypeOk

`func (o *ExportVolumeInfoRest) GetObjecttypeOk() (*string, bool)`

GetObjecttypeOk returns a tuple with the Objecttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjecttype

`func (o *ExportVolumeInfoRest) SetObjecttype(v string)`

SetObjecttype sets Objecttype field to given value.

### HasObjecttype

`func (o *ExportVolumeInfoRest) HasObjecttype() bool`

HasObjecttype returns a boolean if a field has been set.

### GetObjectid

`func (o *ExportVolumeInfoRest) GetObjectid() string`

GetObjectid returns the Objectid field if non-nil, zero value otherwise.

### GetObjectidOk

`func (o *ExportVolumeInfoRest) GetObjectidOk() (*string, bool)`

GetObjectidOk returns a tuple with the Objectid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectid

`func (o *ExportVolumeInfoRest) SetObjectid(v string)`

SetObjectid sets Objectid field to given value.

### HasObjectid

`func (o *ExportVolumeInfoRest) HasObjectid() bool`

HasObjectid returns a boolean if a field has been set.

### GetId

`func (o *ExportVolumeInfoRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExportVolumeInfoRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExportVolumeInfoRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExportVolumeInfoRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ExportVolumeInfoRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ExportVolumeInfoRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ExportVolumeInfoRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ExportVolumeInfoRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ExportVolumeInfoRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ExportVolumeInfoRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ExportVolumeInfoRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ExportVolumeInfoRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ExportVolumeInfoRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ExportVolumeInfoRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ExportVolumeInfoRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ExportVolumeInfoRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



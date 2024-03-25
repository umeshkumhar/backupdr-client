# DbDumpUploadUrlRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbDumpUploadUrl** | Pointer to **string** |  | [optional] 
**DbName** | Pointer to **string** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Errormsg** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewDbDumpUploadUrlRest

`func NewDbDumpUploadUrlRest() *DbDumpUploadUrlRest`

NewDbDumpUploadUrlRest instantiates a new DbDumpUploadUrlRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbDumpUploadUrlRestWithDefaults

`func NewDbDumpUploadUrlRestWithDefaults() *DbDumpUploadUrlRest`

NewDbDumpUploadUrlRestWithDefaults instantiates a new DbDumpUploadUrlRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbDumpUploadUrl

`func (o *DbDumpUploadUrlRest) GetDbDumpUploadUrl() string`

GetDbDumpUploadUrl returns the DbDumpUploadUrl field if non-nil, zero value otherwise.

### GetDbDumpUploadUrlOk

`func (o *DbDumpUploadUrlRest) GetDbDumpUploadUrlOk() (*string, bool)`

GetDbDumpUploadUrlOk returns a tuple with the DbDumpUploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbDumpUploadUrl

`func (o *DbDumpUploadUrlRest) SetDbDumpUploadUrl(v string)`

SetDbDumpUploadUrl sets DbDumpUploadUrl field to given value.

### HasDbDumpUploadUrl

`func (o *DbDumpUploadUrlRest) HasDbDumpUploadUrl() bool`

HasDbDumpUploadUrl returns a boolean if a field has been set.

### GetDbName

`func (o *DbDumpUploadUrlRest) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *DbDumpUploadUrlRest) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *DbDumpUploadUrlRest) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *DbDumpUploadUrlRest) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetSuccess

`func (o *DbDumpUploadUrlRest) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *DbDumpUploadUrlRest) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *DbDumpUploadUrlRest) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *DbDumpUploadUrlRest) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetErrormsg

`func (o *DbDumpUploadUrlRest) GetErrormsg() string`

GetErrormsg returns the Errormsg field if non-nil, zero value otherwise.

### GetErrormsgOk

`func (o *DbDumpUploadUrlRest) GetErrormsgOk() (*string, bool)`

GetErrormsgOk returns a tuple with the Errormsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrormsg

`func (o *DbDumpUploadUrlRest) SetErrormsg(v string)`

SetErrormsg sets Errormsg field to given value.

### HasErrormsg

`func (o *DbDumpUploadUrlRest) HasErrormsg() bool`

HasErrormsg returns a boolean if a field has been set.

### GetId

`func (o *DbDumpUploadUrlRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbDumpUploadUrlRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbDumpUploadUrlRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DbDumpUploadUrlRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *DbDumpUploadUrlRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DbDumpUploadUrlRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DbDumpUploadUrlRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *DbDumpUploadUrlRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *DbDumpUploadUrlRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *DbDumpUploadUrlRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *DbDumpUploadUrlRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *DbDumpUploadUrlRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *DbDumpUploadUrlRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *DbDumpUploadUrlRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *DbDumpUploadUrlRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *DbDumpUploadUrlRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



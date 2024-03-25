# ClientInfoRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exportprop** | Pointer to [**[]NameValueRest**](NameValueRest.md) |  | [optional] 
**NasMountProp** | Pointer to [**[]NameValueRest**](NameValueRest.md) |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewClientInfoRest

`func NewClientInfoRest() *ClientInfoRest`

NewClientInfoRest instantiates a new ClientInfoRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientInfoRestWithDefaults

`func NewClientInfoRestWithDefaults() *ClientInfoRest`

NewClientInfoRestWithDefaults instantiates a new ClientInfoRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportprop

`func (o *ClientInfoRest) GetExportprop() []NameValueRest`

GetExportprop returns the Exportprop field if non-nil, zero value otherwise.

### GetExportpropOk

`func (o *ClientInfoRest) GetExportpropOk() (*[]NameValueRest, bool)`

GetExportpropOk returns a tuple with the Exportprop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportprop

`func (o *ClientInfoRest) SetExportprop(v []NameValueRest)`

SetExportprop sets Exportprop field to given value.

### HasExportprop

`func (o *ClientInfoRest) HasExportprop() bool`

HasExportprop returns a boolean if a field has been set.

### GetNasMountProp

`func (o *ClientInfoRest) GetNasMountProp() []NameValueRest`

GetNasMountProp returns the NasMountProp field if non-nil, zero value otherwise.

### GetNasMountPropOk

`func (o *ClientInfoRest) GetNasMountPropOk() (*[]NameValueRest, bool)`

GetNasMountPropOk returns a tuple with the NasMountProp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasMountProp

`func (o *ClientInfoRest) SetNasMountProp(v []NameValueRest)`

SetNasMountProp sets NasMountProp field to given value.

### HasNasMountProp

`func (o *ClientInfoRest) HasNasMountProp() bool`

HasNasMountProp returns a boolean if a field has been set.

### GetHost

`func (o *ClientInfoRest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ClientInfoRest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ClientInfoRest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ClientInfoRest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *ClientInfoRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientInfoRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientInfoRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClientInfoRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ClientInfoRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ClientInfoRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ClientInfoRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ClientInfoRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ClientInfoRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ClientInfoRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ClientInfoRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ClientInfoRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ClientInfoRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ClientInfoRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ClientInfoRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ClientInfoRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



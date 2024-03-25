# IAMPermissionsRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to **map[string]bool** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewIAMPermissionsRest

`func NewIAMPermissionsRest() *IAMPermissionsRest`

NewIAMPermissionsRest instantiates a new IAMPermissionsRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIAMPermissionsRestWithDefaults

`func NewIAMPermissionsRestWithDefaults() *IAMPermissionsRest`

NewIAMPermissionsRestWithDefaults instantiates a new IAMPermissionsRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *IAMPermissionsRest) GetPermissions() map[string]bool`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *IAMPermissionsRest) GetPermissionsOk() (*map[string]bool, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *IAMPermissionsRest) SetPermissions(v map[string]bool)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *IAMPermissionsRest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetId

`func (o *IAMPermissionsRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IAMPermissionsRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IAMPermissionsRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IAMPermissionsRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *IAMPermissionsRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *IAMPermissionsRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *IAMPermissionsRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *IAMPermissionsRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *IAMPermissionsRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *IAMPermissionsRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *IAMPermissionsRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *IAMPermissionsRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *IAMPermissionsRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *IAMPermissionsRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *IAMPermissionsRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *IAMPermissionsRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



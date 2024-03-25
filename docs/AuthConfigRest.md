# AuthConfigRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** | Authentication method currently configured | [optional] 
**AvailableMethod** | Pointer to **[]string** | Available authentication methods | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewAuthConfigRest

`func NewAuthConfigRest() *AuthConfigRest`

NewAuthConfigRest instantiates a new AuthConfigRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthConfigRestWithDefaults

`func NewAuthConfigRestWithDefaults() *AuthConfigRest`

NewAuthConfigRestWithDefaults instantiates a new AuthConfigRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *AuthConfigRest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AuthConfigRest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AuthConfigRest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *AuthConfigRest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetAvailableMethod

`func (o *AuthConfigRest) GetAvailableMethod() []string`

GetAvailableMethod returns the AvailableMethod field if non-nil, zero value otherwise.

### GetAvailableMethodOk

`func (o *AuthConfigRest) GetAvailableMethodOk() (*[]string, bool)`

GetAvailableMethodOk returns a tuple with the AvailableMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMethod

`func (o *AuthConfigRest) SetAvailableMethod(v []string)`

SetAvailableMethod sets AvailableMethod field to given value.

### HasAvailableMethod

`func (o *AuthConfigRest) HasAvailableMethod() bool`

HasAvailableMethod returns a boolean if a field has been set.

### GetId

`func (o *AuthConfigRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthConfigRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthConfigRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthConfigRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *AuthConfigRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AuthConfigRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AuthConfigRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AuthConfigRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *AuthConfigRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *AuthConfigRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *AuthConfigRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *AuthConfigRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *AuthConfigRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *AuthConfigRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *AuthConfigRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *AuthConfigRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



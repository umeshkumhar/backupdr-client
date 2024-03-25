# SessionRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** | region | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** | User&#39;s timezone, if set. | [optional] 
**Userpref** | Pointer to **string** | This object contains user preferences that client programs can utilize. The content is opaque to API service. | [optional] 
**User** | Pointer to [**UserRest**](UserRest.md) |  | [optional] 
**SessionId** | Pointer to **string** | Session id of the current login | [optional] 
**Rights** | Pointer to [**[]RightRest**](RightRest.md) | Effective access rights that this login session contains. | [optional] 
**Authconfig** | Pointer to [**AuthConfigRest**](AuthConfigRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewSessionRest

`func NewSessionRest() *SessionRest`

NewSessionRest instantiates a new SessionRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionRestWithDefaults

`func NewSessionRestWithDefaults() *SessionRest`

NewSessionRestWithDefaults instantiates a new SessionRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *SessionRest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SessionRest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SessionRest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *SessionRest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetUsername

`func (o *SessionRest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SessionRest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SessionRest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SessionRest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetTimezone

`func (o *SessionRest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *SessionRest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *SessionRest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *SessionRest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetUserpref

`func (o *SessionRest) GetUserpref() string`

GetUserpref returns the Userpref field if non-nil, zero value otherwise.

### GetUserprefOk

`func (o *SessionRest) GetUserprefOk() (*string, bool)`

GetUserprefOk returns a tuple with the Userpref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserpref

`func (o *SessionRest) SetUserpref(v string)`

SetUserpref sets Userpref field to given value.

### HasUserpref

`func (o *SessionRest) HasUserpref() bool`

HasUserpref returns a boolean if a field has been set.

### GetUser

`func (o *SessionRest) GetUser() UserRest`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SessionRest) GetUserOk() (*UserRest, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SessionRest) SetUser(v UserRest)`

SetUser sets User field to given value.

### HasUser

`func (o *SessionRest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetSessionId

`func (o *SessionRest) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *SessionRest) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *SessionRest) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *SessionRest) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetRights

`func (o *SessionRest) GetRights() []RightRest`

GetRights returns the Rights field if non-nil, zero value otherwise.

### GetRightsOk

`func (o *SessionRest) GetRightsOk() (*[]RightRest, bool)`

GetRightsOk returns a tuple with the Rights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRights

`func (o *SessionRest) SetRights(v []RightRest)`

SetRights sets Rights field to given value.

### HasRights

`func (o *SessionRest) HasRights() bool`

HasRights returns a boolean if a field has been set.

### GetAuthconfig

`func (o *SessionRest) GetAuthconfig() AuthConfigRest`

GetAuthconfig returns the Authconfig field if non-nil, zero value otherwise.

### GetAuthconfigOk

`func (o *SessionRest) GetAuthconfigOk() (*AuthConfigRest, bool)`

GetAuthconfigOk returns a tuple with the Authconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthconfig

`func (o *SessionRest) SetAuthconfig(v AuthConfigRest)`

SetAuthconfig sets Authconfig field to given value.

### HasAuthconfig

`func (o *SessionRest) HasAuthconfig() bool`

HasAuthconfig returns a boolean if a field has been set.

### GetId

`func (o *SessionRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SessionRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *SessionRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SessionRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SessionRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SessionRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *SessionRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *SessionRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *SessionRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *SessionRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *SessionRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *SessionRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *SessionRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *SessionRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UserInfoRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessmode** | Pointer to **[]string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Haspassword** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewUserInfoRest

`func NewUserInfoRest() *UserInfoRest`

NewUserInfoRest instantiates a new UserInfoRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInfoRestWithDefaults

`func NewUserInfoRestWithDefaults() *UserInfoRest`

NewUserInfoRestWithDefaults instantiates a new UserInfoRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessmode

`func (o *UserInfoRest) GetAccessmode() []string`

GetAccessmode returns the Accessmode field if non-nil, zero value otherwise.

### GetAccessmodeOk

`func (o *UserInfoRest) GetAccessmodeOk() (*[]string, bool)`

GetAccessmodeOk returns a tuple with the Accessmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessmode

`func (o *UserInfoRest) SetAccessmode(v []string)`

SetAccessmode sets Accessmode field to given value.

### HasAccessmode

`func (o *UserInfoRest) HasAccessmode() bool`

HasAccessmode returns a boolean if a field has been set.

### GetUsername

`func (o *UserInfoRest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserInfoRest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserInfoRest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserInfoRest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetGroup

`func (o *UserInfoRest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *UserInfoRest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *UserInfoRest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *UserInfoRest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHaspassword

`func (o *UserInfoRest) GetHaspassword() bool`

GetHaspassword returns the Haspassword field if non-nil, zero value otherwise.

### GetHaspasswordOk

`func (o *UserInfoRest) GetHaspasswordOk() (*bool, bool)`

GetHaspasswordOk returns a tuple with the Haspassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaspassword

`func (o *UserInfoRest) SetHaspassword(v bool)`

SetHaspassword sets Haspassword field to given value.

### HasHaspassword

`func (o *UserInfoRest) HasHaspassword() bool`

HasHaspassword returns a boolean if a field has been set.

### GetId

`func (o *UserInfoRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserInfoRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserInfoRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserInfoRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *UserInfoRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *UserInfoRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *UserInfoRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *UserInfoRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *UserInfoRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *UserInfoRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *UserInfoRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *UserInfoRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *UserInfoRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *UserInfoRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *UserInfoRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *UserInfoRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



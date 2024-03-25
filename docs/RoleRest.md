# RoleRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Modifydate** | Pointer to **int64** |  | [optional] 
**Createdate** | Pointer to **int64** |  | [optional] 
**RightList** | Pointer to [**[]RightRest**](RightRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewRoleRest

`func NewRoleRest() *RoleRest`

NewRoleRest instantiates a new RoleRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleRestWithDefaults

`func NewRoleRestWithDefaults() *RoleRest`

NewRoleRestWithDefaults instantiates a new RoleRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *RoleRest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RoleRest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RoleRest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RoleRest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetName

`func (o *RoleRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoleRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModifydate

`func (o *RoleRest) GetModifydate() int64`

GetModifydate returns the Modifydate field if non-nil, zero value otherwise.

### GetModifydateOk

`func (o *RoleRest) GetModifydateOk() (*int64, bool)`

GetModifydateOk returns a tuple with the Modifydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifydate

`func (o *RoleRest) SetModifydate(v int64)`

SetModifydate sets Modifydate field to given value.

### HasModifydate

`func (o *RoleRest) HasModifydate() bool`

HasModifydate returns a boolean if a field has been set.

### GetCreatedate

`func (o *RoleRest) GetCreatedate() int64`

GetCreatedate returns the Createdate field if non-nil, zero value otherwise.

### GetCreatedateOk

`func (o *RoleRest) GetCreatedateOk() (*int64, bool)`

GetCreatedateOk returns a tuple with the Createdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedate

`func (o *RoleRest) SetCreatedate(v int64)`

SetCreatedate sets Createdate field to given value.

### HasCreatedate

`func (o *RoleRest) HasCreatedate() bool`

HasCreatedate returns a boolean if a field has been set.

### GetRightList

`func (o *RoleRest) GetRightList() []RightRest`

GetRightList returns the RightList field if non-nil, zero value otherwise.

### GetRightListOk

`func (o *RoleRest) GetRightListOk() (*[]RightRest, bool)`

GetRightListOk returns a tuple with the RightList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightList

`func (o *RoleRest) SetRightList(v []RightRest)`

SetRightList sets RightList field to given value.

### HasRightList

`func (o *RoleRest) HasRightList() bool`

HasRightList returns a boolean if a field has been set.

### GetId

`func (o *RoleRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoleRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *RoleRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RoleRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RoleRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RoleRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *RoleRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *RoleRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *RoleRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *RoleRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *RoleRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *RoleRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *RoleRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *RoleRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



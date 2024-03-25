# MembershipChangeRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to **[]int64** |  | [optional] 
**Action** | Pointer to **string** | add | remove | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewMembershipChangeRest

`func NewMembershipChangeRest() *MembershipChangeRest`

NewMembershipChangeRest instantiates a new MembershipChangeRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembershipChangeRestWithDefaults

`func NewMembershipChangeRestWithDefaults() *MembershipChangeRest`

NewMembershipChangeRestWithDefaults instantiates a new MembershipChangeRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *MembershipChangeRest) GetMembers() []int64`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *MembershipChangeRest) GetMembersOk() (*[]int64, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *MembershipChangeRest) SetMembers(v []int64)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *MembershipChangeRest) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetAction

`func (o *MembershipChangeRest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *MembershipChangeRest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *MembershipChangeRest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *MembershipChangeRest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetId

`func (o *MembershipChangeRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MembershipChangeRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MembershipChangeRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MembershipChangeRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *MembershipChangeRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *MembershipChangeRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *MembershipChangeRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *MembershipChangeRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *MembershipChangeRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *MembershipChangeRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *MembershipChangeRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *MembershipChangeRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *MembershipChangeRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *MembershipChangeRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *MembershipChangeRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *MembershipChangeRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SlaRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Immutable** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Application** | Pointer to [**ApplicationRest**](ApplicationRest.md) |  | [optional] 
**Slt** | Pointer to [**SltRest**](SltRest.md) |  | [optional] 
**Options** | Pointer to [**[]AdvancedOptionRest**](AdvancedOptionRest.md) |  | [optional] 
**Modifydate** | Pointer to **int64** |  | [optional] 
**Scheduleoff** | Pointer to **string** |  | [optional] 
**Slp** | Pointer to [**SlpRest**](SlpRest.md) |  | [optional] 
**Logexpirationoff** | Pointer to **bool** |  | [optional] 
**Dedupasyncoff** | Pointer to **string** |  | [optional] 
**Expirationoff** | Pointer to **string** |  | [optional] 
**Group** | Pointer to [**LogicalGroupRest**](LogicalGroupRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewSlaRest

`func NewSlaRest() *SlaRest`

NewSlaRest instantiates a new SlaRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlaRestWithDefaults

`func NewSlaRestWithDefaults() *SlaRest`

NewSlaRestWithDefaults instantiates a new SlaRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImmutable

`func (o *SlaRest) GetImmutable() bool`

GetImmutable returns the Immutable field if non-nil, zero value otherwise.

### GetImmutableOk

`func (o *SlaRest) GetImmutableOk() (*bool, bool)`

GetImmutableOk returns a tuple with the Immutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutable

`func (o *SlaRest) SetImmutable(v bool)`

SetImmutable sets Immutable field to given value.

### HasImmutable

`func (o *SlaRest) HasImmutable() bool`

HasImmutable returns a boolean if a field has been set.

### GetDescription

`func (o *SlaRest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SlaRest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SlaRest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SlaRest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetApplication

`func (o *SlaRest) GetApplication() ApplicationRest`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *SlaRest) GetApplicationOk() (*ApplicationRest, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *SlaRest) SetApplication(v ApplicationRest)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *SlaRest) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetSlt

`func (o *SlaRest) GetSlt() SltRest`

GetSlt returns the Slt field if non-nil, zero value otherwise.

### GetSltOk

`func (o *SlaRest) GetSltOk() (*SltRest, bool)`

GetSltOk returns a tuple with the Slt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlt

`func (o *SlaRest) SetSlt(v SltRest)`

SetSlt sets Slt field to given value.

### HasSlt

`func (o *SlaRest) HasSlt() bool`

HasSlt returns a boolean if a field has been set.

### GetOptions

`func (o *SlaRest) GetOptions() []AdvancedOptionRest`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SlaRest) GetOptionsOk() (*[]AdvancedOptionRest, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SlaRest) SetOptions(v []AdvancedOptionRest)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SlaRest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetModifydate

`func (o *SlaRest) GetModifydate() int64`

GetModifydate returns the Modifydate field if non-nil, zero value otherwise.

### GetModifydateOk

`func (o *SlaRest) GetModifydateOk() (*int64, bool)`

GetModifydateOk returns a tuple with the Modifydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifydate

`func (o *SlaRest) SetModifydate(v int64)`

SetModifydate sets Modifydate field to given value.

### HasModifydate

`func (o *SlaRest) HasModifydate() bool`

HasModifydate returns a boolean if a field has been set.

### GetScheduleoff

`func (o *SlaRest) GetScheduleoff() string`

GetScheduleoff returns the Scheduleoff field if non-nil, zero value otherwise.

### GetScheduleoffOk

`func (o *SlaRest) GetScheduleoffOk() (*string, bool)`

GetScheduleoffOk returns a tuple with the Scheduleoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleoff

`func (o *SlaRest) SetScheduleoff(v string)`

SetScheduleoff sets Scheduleoff field to given value.

### HasScheduleoff

`func (o *SlaRest) HasScheduleoff() bool`

HasScheduleoff returns a boolean if a field has been set.

### GetSlp

`func (o *SlaRest) GetSlp() SlpRest`

GetSlp returns the Slp field if non-nil, zero value otherwise.

### GetSlpOk

`func (o *SlaRest) GetSlpOk() (*SlpRest, bool)`

GetSlpOk returns a tuple with the Slp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlp

`func (o *SlaRest) SetSlp(v SlpRest)`

SetSlp sets Slp field to given value.

### HasSlp

`func (o *SlaRest) HasSlp() bool`

HasSlp returns a boolean if a field has been set.

### GetLogexpirationoff

`func (o *SlaRest) GetLogexpirationoff() bool`

GetLogexpirationoff returns the Logexpirationoff field if non-nil, zero value otherwise.

### GetLogexpirationoffOk

`func (o *SlaRest) GetLogexpirationoffOk() (*bool, bool)`

GetLogexpirationoffOk returns a tuple with the Logexpirationoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogexpirationoff

`func (o *SlaRest) SetLogexpirationoff(v bool)`

SetLogexpirationoff sets Logexpirationoff field to given value.

### HasLogexpirationoff

`func (o *SlaRest) HasLogexpirationoff() bool`

HasLogexpirationoff returns a boolean if a field has been set.

### GetDedupasyncoff

`func (o *SlaRest) GetDedupasyncoff() string`

GetDedupasyncoff returns the Dedupasyncoff field if non-nil, zero value otherwise.

### GetDedupasyncoffOk

`func (o *SlaRest) GetDedupasyncoffOk() (*string, bool)`

GetDedupasyncoffOk returns a tuple with the Dedupasyncoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupasyncoff

`func (o *SlaRest) SetDedupasyncoff(v string)`

SetDedupasyncoff sets Dedupasyncoff field to given value.

### HasDedupasyncoff

`func (o *SlaRest) HasDedupasyncoff() bool`

HasDedupasyncoff returns a boolean if a field has been set.

### GetExpirationoff

`func (o *SlaRest) GetExpirationoff() string`

GetExpirationoff returns the Expirationoff field if non-nil, zero value otherwise.

### GetExpirationoffOk

`func (o *SlaRest) GetExpirationoffOk() (*string, bool)`

GetExpirationoffOk returns a tuple with the Expirationoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationoff

`func (o *SlaRest) SetExpirationoff(v string)`

SetExpirationoff sets Expirationoff field to given value.

### HasExpirationoff

`func (o *SlaRest) HasExpirationoff() bool`

HasExpirationoff returns a boolean if a field has been set.

### GetGroup

`func (o *SlaRest) GetGroup() LogicalGroupRest`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SlaRest) GetGroupOk() (*LogicalGroupRest, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SlaRest) SetGroup(v LogicalGroupRest)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *SlaRest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetId

`func (o *SlaRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SlaRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SlaRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SlaRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *SlaRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SlaRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SlaRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SlaRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *SlaRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *SlaRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *SlaRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *SlaRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *SlaRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *SlaRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *SlaRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *SlaRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



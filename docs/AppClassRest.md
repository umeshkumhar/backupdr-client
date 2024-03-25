# AppClassRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**FriendyName** | Pointer to **string** |  | [optional] 
**ScriptFile** | Pointer to [**ScriptFileRest**](ScriptFileRest.md) |  | [optional] 
**ProvisioningOptions** | Pointer to [**[]ProvisioningOptionMetaRest**](ProvisioningOptionMetaRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewAppClassRest

`func NewAppClassRest() *AppClassRest`

NewAppClassRest instantiates a new AppClassRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppClassRestWithDefaults

`func NewAppClassRestWithDefaults() *AppClassRest`

NewAppClassRestWithDefaults instantiates a new AppClassRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AppClassRest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppClassRest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppClassRest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppClassRest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *AppClassRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppClassRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppClassRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppClassRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFriendyName

`func (o *AppClassRest) GetFriendyName() string`

GetFriendyName returns the FriendyName field if non-nil, zero value otherwise.

### GetFriendyNameOk

`func (o *AppClassRest) GetFriendyNameOk() (*string, bool)`

GetFriendyNameOk returns a tuple with the FriendyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendyName

`func (o *AppClassRest) SetFriendyName(v string)`

SetFriendyName sets FriendyName field to given value.

### HasFriendyName

`func (o *AppClassRest) HasFriendyName() bool`

HasFriendyName returns a boolean if a field has been set.

### GetScriptFile

`func (o *AppClassRest) GetScriptFile() ScriptFileRest`

GetScriptFile returns the ScriptFile field if non-nil, zero value otherwise.

### GetScriptFileOk

`func (o *AppClassRest) GetScriptFileOk() (*ScriptFileRest, bool)`

GetScriptFileOk returns a tuple with the ScriptFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptFile

`func (o *AppClassRest) SetScriptFile(v ScriptFileRest)`

SetScriptFile sets ScriptFile field to given value.

### HasScriptFile

`func (o *AppClassRest) HasScriptFile() bool`

HasScriptFile returns a boolean if a field has been set.

### GetProvisioningOptions

`func (o *AppClassRest) GetProvisioningOptions() []ProvisioningOptionMetaRest`

GetProvisioningOptions returns the ProvisioningOptions field if non-nil, zero value otherwise.

### GetProvisioningOptionsOk

`func (o *AppClassRest) GetProvisioningOptionsOk() (*[]ProvisioningOptionMetaRest, bool)`

GetProvisioningOptionsOk returns a tuple with the ProvisioningOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningOptions

`func (o *AppClassRest) SetProvisioningOptions(v []ProvisioningOptionMetaRest)`

SetProvisioningOptions sets ProvisioningOptions field to given value.

### HasProvisioningOptions

`func (o *AppClassRest) HasProvisioningOptions() bool`

HasProvisioningOptions returns a boolean if a field has been set.

### GetId

`func (o *AppClassRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppClassRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppClassRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppClassRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *AppClassRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AppClassRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AppClassRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AppClassRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *AppClassRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *AppClassRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *AppClassRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *AppClassRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *AppClassRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *AppClassRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *AppClassRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *AppClassRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



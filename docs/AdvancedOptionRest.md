# AdvancedOptionRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to [**PolicyRest**](PolicyRest.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Slt** | Pointer to [**SltRest**](SltRest.md) |  | [optional] 
**Sla** | Pointer to [**SlaRest**](SlaRest.md) |  | [optional] 
**App** | Pointer to [**ApplicationRest**](ApplicationRest.md) |  | [optional] 
**Various** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewAdvancedOptionRest

`func NewAdvancedOptionRest() *AdvancedOptionRest`

NewAdvancedOptionRest instantiates a new AdvancedOptionRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedOptionRestWithDefaults

`func NewAdvancedOptionRestWithDefaults() *AdvancedOptionRest`

NewAdvancedOptionRestWithDefaults instantiates a new AdvancedOptionRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *AdvancedOptionRest) GetPolicy() PolicyRest`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *AdvancedOptionRest) GetPolicyOk() (*PolicyRest, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *AdvancedOptionRest) SetPolicy(v PolicyRest)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *AdvancedOptionRest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetName

`func (o *AdvancedOptionRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvancedOptionRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvancedOptionRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdvancedOptionRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *AdvancedOptionRest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AdvancedOptionRest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AdvancedOptionRest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AdvancedOptionRest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetSlt

`func (o *AdvancedOptionRest) GetSlt() SltRest`

GetSlt returns the Slt field if non-nil, zero value otherwise.

### GetSltOk

`func (o *AdvancedOptionRest) GetSltOk() (*SltRest, bool)`

GetSltOk returns a tuple with the Slt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlt

`func (o *AdvancedOptionRest) SetSlt(v SltRest)`

SetSlt sets Slt field to given value.

### HasSlt

`func (o *AdvancedOptionRest) HasSlt() bool`

HasSlt returns a boolean if a field has been set.

### GetSla

`func (o *AdvancedOptionRest) GetSla() SlaRest`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *AdvancedOptionRest) GetSlaOk() (*SlaRest, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *AdvancedOptionRest) SetSla(v SlaRest)`

SetSla sets Sla field to given value.

### HasSla

`func (o *AdvancedOptionRest) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetApp

`func (o *AdvancedOptionRest) GetApp() ApplicationRest`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *AdvancedOptionRest) GetAppOk() (*ApplicationRest, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *AdvancedOptionRest) SetApp(v ApplicationRest)`

SetApp sets App field to given value.

### HasApp

`func (o *AdvancedOptionRest) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetVarious

`func (o *AdvancedOptionRest) GetVarious() bool`

GetVarious returns the Various field if non-nil, zero value otherwise.

### GetVariousOk

`func (o *AdvancedOptionRest) GetVariousOk() (*bool, bool)`

GetVariousOk returns a tuple with the Various field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVarious

`func (o *AdvancedOptionRest) SetVarious(v bool)`

SetVarious sets Various field to given value.

### HasVarious

`func (o *AdvancedOptionRest) HasVarious() bool`

HasVarious returns a boolean if a field has been set.

### GetId

`func (o *AdvancedOptionRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvancedOptionRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvancedOptionRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvancedOptionRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *AdvancedOptionRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AdvancedOptionRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AdvancedOptionRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AdvancedOptionRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *AdvancedOptionRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *AdvancedOptionRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *AdvancedOptionRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *AdvancedOptionRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *AdvancedOptionRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *AdvancedOptionRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *AdvancedOptionRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *AdvancedOptionRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



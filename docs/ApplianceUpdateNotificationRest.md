# ApplianceUpdateNotificationRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewApplianceUpdateNotificationRest

`func NewApplianceUpdateNotificationRest() *ApplianceUpdateNotificationRest`

NewApplianceUpdateNotificationRest instantiates a new ApplianceUpdateNotificationRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceUpdateNotificationRestWithDefaults

`func NewApplianceUpdateNotificationRestWithDefaults() *ApplianceUpdateNotificationRest`

NewApplianceUpdateNotificationRestWithDefaults instantiates a new ApplianceUpdateNotificationRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplianceUpdateNotificationRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplianceUpdateNotificationRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplianceUpdateNotificationRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplianceUpdateNotificationRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ApplianceUpdateNotificationRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ApplianceUpdateNotificationRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ApplianceUpdateNotificationRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ApplianceUpdateNotificationRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ApplianceUpdateNotificationRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ApplianceUpdateNotificationRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ApplianceUpdateNotificationRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ApplianceUpdateNotificationRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ApplianceUpdateNotificationRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ApplianceUpdateNotificationRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ApplianceUpdateNotificationRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ApplianceUpdateNotificationRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



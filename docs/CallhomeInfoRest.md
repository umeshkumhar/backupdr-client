# CallhomeInfoRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disable** | Pointer to **bool** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewCallhomeInfoRest

`func NewCallhomeInfoRest() *CallhomeInfoRest`

NewCallhomeInfoRest instantiates a new CallhomeInfoRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallhomeInfoRestWithDefaults

`func NewCallhomeInfoRestWithDefaults() *CallhomeInfoRest`

NewCallhomeInfoRestWithDefaults instantiates a new CallhomeInfoRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisable

`func (o *CallhomeInfoRest) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *CallhomeInfoRest) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *CallhomeInfoRest) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *CallhomeInfoRest) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetMode

`func (o *CallhomeInfoRest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CallhomeInfoRest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CallhomeInfoRest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CallhomeInfoRest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetId

`func (o *CallhomeInfoRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CallhomeInfoRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CallhomeInfoRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CallhomeInfoRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *CallhomeInfoRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CallhomeInfoRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CallhomeInfoRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CallhomeInfoRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *CallhomeInfoRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *CallhomeInfoRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *CallhomeInfoRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *CallhomeInfoRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *CallhomeInfoRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *CallhomeInfoRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *CallhomeInfoRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *CallhomeInfoRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



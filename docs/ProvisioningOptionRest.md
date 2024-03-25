# ProvisioningOptionRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewProvisioningOptionRest

`func NewProvisioningOptionRest() *ProvisioningOptionRest`

NewProvisioningOptionRest instantiates a new ProvisioningOptionRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningOptionRestWithDefaults

`func NewProvisioningOptionRestWithDefaults() *ProvisioningOptionRest`

NewProvisioningOptionRestWithDefaults instantiates a new ProvisioningOptionRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProvisioningOptionRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisioningOptionRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisioningOptionRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProvisioningOptionRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *ProvisioningOptionRest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProvisioningOptionRest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProvisioningOptionRest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProvisioningOptionRest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetId

`func (o *ProvisioningOptionRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvisioningOptionRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvisioningOptionRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProvisioningOptionRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ProvisioningOptionRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ProvisioningOptionRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ProvisioningOptionRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ProvisioningOptionRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ProvisioningOptionRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ProvisioningOptionRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ProvisioningOptionRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ProvisioningOptionRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ProvisioningOptionRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ProvisioningOptionRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ProvisioningOptionRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ProvisioningOptionRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ComplianceSettingsRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to [**PolicyRest**](PolicyRest.md) |  | [optional] 
**WarnThresholdType** | Pointer to **string** |  | [optional] 
**WarnThresholdCustom** | Pointer to **int32** |  | [optional] 
**ErrorThresholdType** | Pointer to **string** |  | [optional] 
**ErrorThresholdCustom** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewComplianceSettingsRest

`func NewComplianceSettingsRest() *ComplianceSettingsRest`

NewComplianceSettingsRest instantiates a new ComplianceSettingsRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplianceSettingsRestWithDefaults

`func NewComplianceSettingsRestWithDefaults() *ComplianceSettingsRest`

NewComplianceSettingsRestWithDefaults instantiates a new ComplianceSettingsRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *ComplianceSettingsRest) GetPolicy() PolicyRest`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *ComplianceSettingsRest) GetPolicyOk() (*PolicyRest, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *ComplianceSettingsRest) SetPolicy(v PolicyRest)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *ComplianceSettingsRest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetWarnThresholdType

`func (o *ComplianceSettingsRest) GetWarnThresholdType() string`

GetWarnThresholdType returns the WarnThresholdType field if non-nil, zero value otherwise.

### GetWarnThresholdTypeOk

`func (o *ComplianceSettingsRest) GetWarnThresholdTypeOk() (*string, bool)`

GetWarnThresholdTypeOk returns a tuple with the WarnThresholdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnThresholdType

`func (o *ComplianceSettingsRest) SetWarnThresholdType(v string)`

SetWarnThresholdType sets WarnThresholdType field to given value.

### HasWarnThresholdType

`func (o *ComplianceSettingsRest) HasWarnThresholdType() bool`

HasWarnThresholdType returns a boolean if a field has been set.

### GetWarnThresholdCustom

`func (o *ComplianceSettingsRest) GetWarnThresholdCustom() int32`

GetWarnThresholdCustom returns the WarnThresholdCustom field if non-nil, zero value otherwise.

### GetWarnThresholdCustomOk

`func (o *ComplianceSettingsRest) GetWarnThresholdCustomOk() (*int32, bool)`

GetWarnThresholdCustomOk returns a tuple with the WarnThresholdCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnThresholdCustom

`func (o *ComplianceSettingsRest) SetWarnThresholdCustom(v int32)`

SetWarnThresholdCustom sets WarnThresholdCustom field to given value.

### HasWarnThresholdCustom

`func (o *ComplianceSettingsRest) HasWarnThresholdCustom() bool`

HasWarnThresholdCustom returns a boolean if a field has been set.

### GetErrorThresholdType

`func (o *ComplianceSettingsRest) GetErrorThresholdType() string`

GetErrorThresholdType returns the ErrorThresholdType field if non-nil, zero value otherwise.

### GetErrorThresholdTypeOk

`func (o *ComplianceSettingsRest) GetErrorThresholdTypeOk() (*string, bool)`

GetErrorThresholdTypeOk returns a tuple with the ErrorThresholdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorThresholdType

`func (o *ComplianceSettingsRest) SetErrorThresholdType(v string)`

SetErrorThresholdType sets ErrorThresholdType field to given value.

### HasErrorThresholdType

`func (o *ComplianceSettingsRest) HasErrorThresholdType() bool`

HasErrorThresholdType returns a boolean if a field has been set.

### GetErrorThresholdCustom

`func (o *ComplianceSettingsRest) GetErrorThresholdCustom() int32`

GetErrorThresholdCustom returns the ErrorThresholdCustom field if non-nil, zero value otherwise.

### GetErrorThresholdCustomOk

`func (o *ComplianceSettingsRest) GetErrorThresholdCustomOk() (*int32, bool)`

GetErrorThresholdCustomOk returns a tuple with the ErrorThresholdCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorThresholdCustom

`func (o *ComplianceSettingsRest) SetErrorThresholdCustom(v int32)`

SetErrorThresholdCustom sets ErrorThresholdCustom field to given value.

### HasErrorThresholdCustom

`func (o *ComplianceSettingsRest) HasErrorThresholdCustom() bool`

HasErrorThresholdCustom returns a boolean if a field has been set.

### GetId

`func (o *ComplianceSettingsRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComplianceSettingsRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComplianceSettingsRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComplianceSettingsRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ComplianceSettingsRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ComplianceSettingsRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ComplianceSettingsRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ComplianceSettingsRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ComplianceSettingsRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ComplianceSettingsRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ComplianceSettingsRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ComplianceSettingsRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ComplianceSettingsRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ComplianceSettingsRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ComplianceSettingsRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ComplianceSettingsRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



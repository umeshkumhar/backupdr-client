# MissingPermissionsRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceProject** | Pointer to [**ProjectPermissionsRest**](ProjectPermissionsRest.md) |  | [optional] 
**TargetProject** | Pointer to [**ProjectPermissionsRest**](ProjectPermissionsRest.md) |  | [optional] 
**Cloudcredential** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewMissingPermissionsRest

`func NewMissingPermissionsRest() *MissingPermissionsRest`

NewMissingPermissionsRest instantiates a new MissingPermissionsRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMissingPermissionsRestWithDefaults

`func NewMissingPermissionsRestWithDefaults() *MissingPermissionsRest`

NewMissingPermissionsRestWithDefaults instantiates a new MissingPermissionsRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceProject

`func (o *MissingPermissionsRest) GetSourceProject() ProjectPermissionsRest`

GetSourceProject returns the SourceProject field if non-nil, zero value otherwise.

### GetSourceProjectOk

`func (o *MissingPermissionsRest) GetSourceProjectOk() (*ProjectPermissionsRest, bool)`

GetSourceProjectOk returns a tuple with the SourceProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceProject

`func (o *MissingPermissionsRest) SetSourceProject(v ProjectPermissionsRest)`

SetSourceProject sets SourceProject field to given value.

### HasSourceProject

`func (o *MissingPermissionsRest) HasSourceProject() bool`

HasSourceProject returns a boolean if a field has been set.

### GetTargetProject

`func (o *MissingPermissionsRest) GetTargetProject() ProjectPermissionsRest`

GetTargetProject returns the TargetProject field if non-nil, zero value otherwise.

### GetTargetProjectOk

`func (o *MissingPermissionsRest) GetTargetProjectOk() (*ProjectPermissionsRest, bool)`

GetTargetProjectOk returns a tuple with the TargetProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetProject

`func (o *MissingPermissionsRest) SetTargetProject(v ProjectPermissionsRest)`

SetTargetProject sets TargetProject field to given value.

### HasTargetProject

`func (o *MissingPermissionsRest) HasTargetProject() bool`

HasTargetProject returns a boolean if a field has been set.

### GetCloudcredential

`func (o *MissingPermissionsRest) GetCloudcredential() string`

GetCloudcredential returns the Cloudcredential field if non-nil, zero value otherwise.

### GetCloudcredentialOk

`func (o *MissingPermissionsRest) GetCloudcredentialOk() (*string, bool)`

GetCloudcredentialOk returns a tuple with the Cloudcredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudcredential

`func (o *MissingPermissionsRest) SetCloudcredential(v string)`

SetCloudcredential sets Cloudcredential field to given value.

### HasCloudcredential

`func (o *MissingPermissionsRest) HasCloudcredential() bool`

HasCloudcredential returns a boolean if a field has been set.

### GetId

`func (o *MissingPermissionsRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MissingPermissionsRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MissingPermissionsRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MissingPermissionsRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *MissingPermissionsRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *MissingPermissionsRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *MissingPermissionsRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *MissingPermissionsRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *MissingPermissionsRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *MissingPermissionsRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *MissingPermissionsRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *MissingPermissionsRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *MissingPermissionsRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *MissingPermissionsRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *MissingPermissionsRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *MissingPermissionsRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ProjectPermissionsRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requiredpermissions** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewProjectPermissionsRest

`func NewProjectPermissionsRest() *ProjectPermissionsRest`

NewProjectPermissionsRest instantiates a new ProjectPermissionsRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectPermissionsRestWithDefaults

`func NewProjectPermissionsRestWithDefaults() *ProjectPermissionsRest`

NewProjectPermissionsRestWithDefaults instantiates a new ProjectPermissionsRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequiredpermissions

`func (o *ProjectPermissionsRest) GetRequiredpermissions() []string`

GetRequiredpermissions returns the Requiredpermissions field if non-nil, zero value otherwise.

### GetRequiredpermissionsOk

`func (o *ProjectPermissionsRest) GetRequiredpermissionsOk() (*[]string, bool)`

GetRequiredpermissionsOk returns a tuple with the Requiredpermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredpermissions

`func (o *ProjectPermissionsRest) SetRequiredpermissions(v []string)`

SetRequiredpermissions sets Requiredpermissions field to given value.

### HasRequiredpermissions

`func (o *ProjectPermissionsRest) HasRequiredpermissions() bool`

HasRequiredpermissions returns a boolean if a field has been set.

### GetName

`func (o *ProjectPermissionsRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectPermissionsRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectPermissionsRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectPermissionsRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *ProjectPermissionsRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectPermissionsRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectPermissionsRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectPermissionsRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ProjectPermissionsRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ProjectPermissionsRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ProjectPermissionsRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ProjectPermissionsRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ProjectPermissionsRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ProjectPermissionsRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ProjectPermissionsRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ProjectPermissionsRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ProjectPermissionsRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ProjectPermissionsRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ProjectPermissionsRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ProjectPermissionsRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



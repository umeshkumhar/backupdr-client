# CloudVmMountRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to [**[]FormFieldRest**](FormFieldRest.md) |  | [optional] 
**Cloudtype** | Pointer to **string** |  | [optional] 
**Transaction** | Pointer to **int32** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Getformtype** | Pointer to **string** |  | [optional] 
**Jobstatus** | Pointer to **string** |  | [optional] 
**MissingPermissions** | Pointer to [**MissingPermissionsRest**](MissingPermissionsRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewCloudVmMountRest

`func NewCloudVmMountRest() *CloudVmMountRest`

NewCloudVmMountRest instantiates a new CloudVmMountRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudVmMountRestWithDefaults

`func NewCloudVmMountRestWithDefaults() *CloudVmMountRest`

NewCloudVmMountRestWithDefaults instantiates a new CloudVmMountRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *CloudVmMountRest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CloudVmMountRest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CloudVmMountRest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CloudVmMountRest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetName

`func (o *CloudVmMountRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudVmMountRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudVmMountRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudVmMountRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFields

`func (o *CloudVmMountRest) GetFields() []FormFieldRest`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CloudVmMountRest) GetFieldsOk() (*[]FormFieldRest, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CloudVmMountRest) SetFields(v []FormFieldRest)`

SetFields sets Fields field to given value.

### HasFields

`func (o *CloudVmMountRest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetCloudtype

`func (o *CloudVmMountRest) GetCloudtype() string`

GetCloudtype returns the Cloudtype field if non-nil, zero value otherwise.

### GetCloudtypeOk

`func (o *CloudVmMountRest) GetCloudtypeOk() (*string, bool)`

GetCloudtypeOk returns a tuple with the Cloudtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudtype

`func (o *CloudVmMountRest) SetCloudtype(v string)`

SetCloudtype sets Cloudtype field to given value.

### HasCloudtype

`func (o *CloudVmMountRest) HasCloudtype() bool`

HasCloudtype returns a boolean if a field has been set.

### GetTransaction

`func (o *CloudVmMountRest) GetTransaction() int32`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *CloudVmMountRest) GetTransactionOk() (*int32, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *CloudVmMountRest) SetTransaction(v int32)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *CloudVmMountRest) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.

### GetImage

`func (o *CloudVmMountRest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CloudVmMountRest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CloudVmMountRest) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CloudVmMountRest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetGetformtype

`func (o *CloudVmMountRest) GetGetformtype() string`

GetGetformtype returns the Getformtype field if non-nil, zero value otherwise.

### GetGetformtypeOk

`func (o *CloudVmMountRest) GetGetformtypeOk() (*string, bool)`

GetGetformtypeOk returns a tuple with the Getformtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetformtype

`func (o *CloudVmMountRest) SetGetformtype(v string)`

SetGetformtype sets Getformtype field to given value.

### HasGetformtype

`func (o *CloudVmMountRest) HasGetformtype() bool`

HasGetformtype returns a boolean if a field has been set.

### GetJobstatus

`func (o *CloudVmMountRest) GetJobstatus() string`

GetJobstatus returns the Jobstatus field if non-nil, zero value otherwise.

### GetJobstatusOk

`func (o *CloudVmMountRest) GetJobstatusOk() (*string, bool)`

GetJobstatusOk returns a tuple with the Jobstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobstatus

`func (o *CloudVmMountRest) SetJobstatus(v string)`

SetJobstatus sets Jobstatus field to given value.

### HasJobstatus

`func (o *CloudVmMountRest) HasJobstatus() bool`

HasJobstatus returns a boolean if a field has been set.

### GetMissingPermissions

`func (o *CloudVmMountRest) GetMissingPermissions() MissingPermissionsRest`

GetMissingPermissions returns the MissingPermissions field if non-nil, zero value otherwise.

### GetMissingPermissionsOk

`func (o *CloudVmMountRest) GetMissingPermissionsOk() (*MissingPermissionsRest, bool)`

GetMissingPermissionsOk returns a tuple with the MissingPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingPermissions

`func (o *CloudVmMountRest) SetMissingPermissions(v MissingPermissionsRest)`

SetMissingPermissions sets MissingPermissions field to given value.

### HasMissingPermissions

`func (o *CloudVmMountRest) HasMissingPermissions() bool`

HasMissingPermissions returns a boolean if a field has been set.

### GetId

`func (o *CloudVmMountRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudVmMountRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudVmMountRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudVmMountRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *CloudVmMountRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CloudVmMountRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CloudVmMountRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CloudVmMountRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *CloudVmMountRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *CloudVmMountRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *CloudVmMountRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *CloudVmMountRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *CloudVmMountRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *CloudVmMountRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *CloudVmMountRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *CloudVmMountRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



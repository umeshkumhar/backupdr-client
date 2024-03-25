# VolumeSelectionRowRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volumeid** | Pointer to **string** |  | [optional] 
**Devicename** | Pointer to **string** |  | [optional] 
**Disktype** | Pointer to [**[]ChoiceValueRest**](ChoiceValueRest.md) |  | [optional] 
**Sourcemountpath** | Pointer to **string** |  | [optional] 
**Devicetype** | Pointer to **string** |  | [optional] 
**Deviceindex** | Pointer to **string** |  | [optional] 
**DisktypeDisabled** | Pointer to **bool** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Selected** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewVolumeSelectionRowRest

`func NewVolumeSelectionRowRest() *VolumeSelectionRowRest`

NewVolumeSelectionRowRest instantiates a new VolumeSelectionRowRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeSelectionRowRestWithDefaults

`func NewVolumeSelectionRowRestWithDefaults() *VolumeSelectionRowRest`

NewVolumeSelectionRowRestWithDefaults instantiates a new VolumeSelectionRowRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumeid

`func (o *VolumeSelectionRowRest) GetVolumeid() string`

GetVolumeid returns the Volumeid field if non-nil, zero value otherwise.

### GetVolumeidOk

`func (o *VolumeSelectionRowRest) GetVolumeidOk() (*string, bool)`

GetVolumeidOk returns a tuple with the Volumeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeid

`func (o *VolumeSelectionRowRest) SetVolumeid(v string)`

SetVolumeid sets Volumeid field to given value.

### HasVolumeid

`func (o *VolumeSelectionRowRest) HasVolumeid() bool`

HasVolumeid returns a boolean if a field has been set.

### GetDevicename

`func (o *VolumeSelectionRowRest) GetDevicename() string`

GetDevicename returns the Devicename field if non-nil, zero value otherwise.

### GetDevicenameOk

`func (o *VolumeSelectionRowRest) GetDevicenameOk() (*string, bool)`

GetDevicenameOk returns a tuple with the Devicename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicename

`func (o *VolumeSelectionRowRest) SetDevicename(v string)`

SetDevicename sets Devicename field to given value.

### HasDevicename

`func (o *VolumeSelectionRowRest) HasDevicename() bool`

HasDevicename returns a boolean if a field has been set.

### GetDisktype

`func (o *VolumeSelectionRowRest) GetDisktype() []ChoiceValueRest`

GetDisktype returns the Disktype field if non-nil, zero value otherwise.

### GetDisktypeOk

`func (o *VolumeSelectionRowRest) GetDisktypeOk() (*[]ChoiceValueRest, bool)`

GetDisktypeOk returns a tuple with the Disktype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisktype

`func (o *VolumeSelectionRowRest) SetDisktype(v []ChoiceValueRest)`

SetDisktype sets Disktype field to given value.

### HasDisktype

`func (o *VolumeSelectionRowRest) HasDisktype() bool`

HasDisktype returns a boolean if a field has been set.

### GetSourcemountpath

`func (o *VolumeSelectionRowRest) GetSourcemountpath() string`

GetSourcemountpath returns the Sourcemountpath field if non-nil, zero value otherwise.

### GetSourcemountpathOk

`func (o *VolumeSelectionRowRest) GetSourcemountpathOk() (*string, bool)`

GetSourcemountpathOk returns a tuple with the Sourcemountpath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcemountpath

`func (o *VolumeSelectionRowRest) SetSourcemountpath(v string)`

SetSourcemountpath sets Sourcemountpath field to given value.

### HasSourcemountpath

`func (o *VolumeSelectionRowRest) HasSourcemountpath() bool`

HasSourcemountpath returns a boolean if a field has been set.

### GetDevicetype

`func (o *VolumeSelectionRowRest) GetDevicetype() string`

GetDevicetype returns the Devicetype field if non-nil, zero value otherwise.

### GetDevicetypeOk

`func (o *VolumeSelectionRowRest) GetDevicetypeOk() (*string, bool)`

GetDevicetypeOk returns a tuple with the Devicetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicetype

`func (o *VolumeSelectionRowRest) SetDevicetype(v string)`

SetDevicetype sets Devicetype field to given value.

### HasDevicetype

`func (o *VolumeSelectionRowRest) HasDevicetype() bool`

HasDevicetype returns a boolean if a field has been set.

### GetDeviceindex

`func (o *VolumeSelectionRowRest) GetDeviceindex() string`

GetDeviceindex returns the Deviceindex field if non-nil, zero value otherwise.

### GetDeviceindexOk

`func (o *VolumeSelectionRowRest) GetDeviceindexOk() (*string, bool)`

GetDeviceindexOk returns a tuple with the Deviceindex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceindex

`func (o *VolumeSelectionRowRest) SetDeviceindex(v string)`

SetDeviceindex sets Deviceindex field to given value.

### HasDeviceindex

`func (o *VolumeSelectionRowRest) HasDeviceindex() bool`

HasDeviceindex returns a boolean if a field has been set.

### GetDisktypeDisabled

`func (o *VolumeSelectionRowRest) GetDisktypeDisabled() bool`

GetDisktypeDisabled returns the DisktypeDisabled field if non-nil, zero value otherwise.

### GetDisktypeDisabledOk

`func (o *VolumeSelectionRowRest) GetDisktypeDisabledOk() (*bool, bool)`

GetDisktypeDisabledOk returns a tuple with the DisktypeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisktypeDisabled

`func (o *VolumeSelectionRowRest) SetDisktypeDisabled(v bool)`

SetDisktypeDisabled sets DisktypeDisabled field to given value.

### HasDisktypeDisabled

`func (o *VolumeSelectionRowRest) HasDisktypeDisabled() bool`

HasDisktypeDisabled returns a boolean if a field has been set.

### GetSize

`func (o *VolumeSelectionRowRest) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VolumeSelectionRowRest) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VolumeSelectionRowRest) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *VolumeSelectionRowRest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetDisabled

`func (o *VolumeSelectionRowRest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *VolumeSelectionRowRest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *VolumeSelectionRowRest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *VolumeSelectionRowRest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetSelected

`func (o *VolumeSelectionRowRest) GetSelected() bool`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *VolumeSelectionRowRest) GetSelectedOk() (*bool, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *VolumeSelectionRowRest) SetSelected(v bool)`

SetSelected sets Selected field to given value.

### HasSelected

`func (o *VolumeSelectionRowRest) HasSelected() bool`

HasSelected returns a boolean if a field has been set.

### GetId

`func (o *VolumeSelectionRowRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeSelectionRowRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeSelectionRowRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VolumeSelectionRowRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *VolumeSelectionRowRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *VolumeSelectionRowRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *VolumeSelectionRowRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *VolumeSelectionRowRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *VolumeSelectionRowRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *VolumeSelectionRowRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *VolumeSelectionRowRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *VolumeSelectionRowRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *VolumeSelectionRowRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *VolumeSelectionRowRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *VolumeSelectionRowRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *VolumeSelectionRowRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



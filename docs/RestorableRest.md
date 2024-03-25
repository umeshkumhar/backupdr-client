# RestorableRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppstateText** | Pointer to **[]string** |  | [optional] 
**Fileinfo** | Pointer to [**[]FileinfoRest**](FileinfoRest.md) |  | [optional] 
**Backedupdb** | Pointer to **string** |  | [optional] 
**Skippeddb** | Pointer to **string** |  | [optional] 
**Faileddb** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Application** | Pointer to [**ApplicationRest**](ApplicationRest.md) |  | [optional] 
**Pathname** | Pointer to **string** |  | [optional] 
**Systemdb** | Pointer to **bool** |  | [optional] 
**AppState** | Pointer to **int64** |  | [optional] 
**Fullpath** | Pointer to **string** |  | [optional] 
**Volumeinfo** | Pointer to [**[]VolumeinfoRest**](VolumeinfoRest.md) |  | [optional] 

## Methods

### NewRestorableRest

`func NewRestorableRest() *RestorableRest`

NewRestorableRest instantiates a new RestorableRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestorableRestWithDefaults

`func NewRestorableRestWithDefaults() *RestorableRest`

NewRestorableRestWithDefaults instantiates a new RestorableRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppstateText

`func (o *RestorableRest) GetAppstateText() []string`

GetAppstateText returns the AppstateText field if non-nil, zero value otherwise.

### GetAppstateTextOk

`func (o *RestorableRest) GetAppstateTextOk() (*[]string, bool)`

GetAppstateTextOk returns a tuple with the AppstateText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppstateText

`func (o *RestorableRest) SetAppstateText(v []string)`

SetAppstateText sets AppstateText field to given value.

### HasAppstateText

`func (o *RestorableRest) HasAppstateText() bool`

HasAppstateText returns a boolean if a field has been set.

### GetFileinfo

`func (o *RestorableRest) GetFileinfo() []FileinfoRest`

GetFileinfo returns the Fileinfo field if non-nil, zero value otherwise.

### GetFileinfoOk

`func (o *RestorableRest) GetFileinfoOk() (*[]FileinfoRest, bool)`

GetFileinfoOk returns a tuple with the Fileinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileinfo

`func (o *RestorableRest) SetFileinfo(v []FileinfoRest)`

SetFileinfo sets Fileinfo field to given value.

### HasFileinfo

`func (o *RestorableRest) HasFileinfo() bool`

HasFileinfo returns a boolean if a field has been set.

### GetBackedupdb

`func (o *RestorableRest) GetBackedupdb() string`

GetBackedupdb returns the Backedupdb field if non-nil, zero value otherwise.

### GetBackedupdbOk

`func (o *RestorableRest) GetBackedupdbOk() (*string, bool)`

GetBackedupdbOk returns a tuple with the Backedupdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackedupdb

`func (o *RestorableRest) SetBackedupdb(v string)`

SetBackedupdb sets Backedupdb field to given value.

### HasBackedupdb

`func (o *RestorableRest) HasBackedupdb() bool`

HasBackedupdb returns a boolean if a field has been set.

### GetSkippeddb

`func (o *RestorableRest) GetSkippeddb() string`

GetSkippeddb returns the Skippeddb field if non-nil, zero value otherwise.

### GetSkippeddbOk

`func (o *RestorableRest) GetSkippeddbOk() (*string, bool)`

GetSkippeddbOk returns a tuple with the Skippeddb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippeddb

`func (o *RestorableRest) SetSkippeddb(v string)`

SetSkippeddb sets Skippeddb field to given value.

### HasSkippeddb

`func (o *RestorableRest) HasSkippeddb() bool`

HasSkippeddb returns a boolean if a field has been set.

### GetFaileddb

`func (o *RestorableRest) GetFaileddb() string`

GetFaileddb returns the Faileddb field if non-nil, zero value otherwise.

### GetFaileddbOk

`func (o *RestorableRest) GetFaileddbOk() (*string, bool)`

GetFaileddbOk returns a tuple with the Faileddb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaileddb

`func (o *RestorableRest) SetFaileddb(v string)`

SetFaileddb sets Faileddb field to given value.

### HasFaileddb

`func (o *RestorableRest) HasFaileddb() bool`

HasFaileddb returns a boolean if a field has been set.

### GetName

`func (o *RestorableRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestorableRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestorableRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RestorableRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *RestorableRest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestorableRest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestorableRest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RestorableRest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetApplication

`func (o *RestorableRest) GetApplication() ApplicationRest`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *RestorableRest) GetApplicationOk() (*ApplicationRest, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *RestorableRest) SetApplication(v ApplicationRest)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *RestorableRest) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetPathname

`func (o *RestorableRest) GetPathname() string`

GetPathname returns the Pathname field if non-nil, zero value otherwise.

### GetPathnameOk

`func (o *RestorableRest) GetPathnameOk() (*string, bool)`

GetPathnameOk returns a tuple with the Pathname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathname

`func (o *RestorableRest) SetPathname(v string)`

SetPathname sets Pathname field to given value.

### HasPathname

`func (o *RestorableRest) HasPathname() bool`

HasPathname returns a boolean if a field has been set.

### GetSystemdb

`func (o *RestorableRest) GetSystemdb() bool`

GetSystemdb returns the Systemdb field if non-nil, zero value otherwise.

### GetSystemdbOk

`func (o *RestorableRest) GetSystemdbOk() (*bool, bool)`

GetSystemdbOk returns a tuple with the Systemdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemdb

`func (o *RestorableRest) SetSystemdb(v bool)`

SetSystemdb sets Systemdb field to given value.

### HasSystemdb

`func (o *RestorableRest) HasSystemdb() bool`

HasSystemdb returns a boolean if a field has been set.

### GetAppState

`func (o *RestorableRest) GetAppState() int64`

GetAppState returns the AppState field if non-nil, zero value otherwise.

### GetAppStateOk

`func (o *RestorableRest) GetAppStateOk() (*int64, bool)`

GetAppStateOk returns a tuple with the AppState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppState

`func (o *RestorableRest) SetAppState(v int64)`

SetAppState sets AppState field to given value.

### HasAppState

`func (o *RestorableRest) HasAppState() bool`

HasAppState returns a boolean if a field has been set.

### GetFullpath

`func (o *RestorableRest) GetFullpath() string`

GetFullpath returns the Fullpath field if non-nil, zero value otherwise.

### GetFullpathOk

`func (o *RestorableRest) GetFullpathOk() (*string, bool)`

GetFullpathOk returns a tuple with the Fullpath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullpath

`func (o *RestorableRest) SetFullpath(v string)`

SetFullpath sets Fullpath field to given value.

### HasFullpath

`func (o *RestorableRest) HasFullpath() bool`

HasFullpath returns a boolean if a field has been set.

### GetVolumeinfo

`func (o *RestorableRest) GetVolumeinfo() []VolumeinfoRest`

GetVolumeinfo returns the Volumeinfo field if non-nil, zero value otherwise.

### GetVolumeinfoOk

`func (o *RestorableRest) GetVolumeinfoOk() (*[]VolumeinfoRest, bool)`

GetVolumeinfoOk returns a tuple with the Volumeinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeinfo

`func (o *RestorableRest) SetVolumeinfo(v []VolumeinfoRest)`

SetVolumeinfo sets Volumeinfo field to given value.

### HasVolumeinfo

`func (o *RestorableRest) HasVolumeinfo() bool`

HasVolumeinfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



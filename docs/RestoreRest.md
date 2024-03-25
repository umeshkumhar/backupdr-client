# RestoreRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** |  | [optional] 
**Host** | Pointer to [**HostRest**](HostRest.md) |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Cloudvmoptions** | Pointer to [**CloudVmMountRest**](CloudVmMountRest.md) |  | [optional] 
**Rdmmode** | Pointer to **string** |  | [optional] 
**Physicalrdm** | Pointer to **string** |  | [optional] 
**Targetnode** | Pointer to **string** |  | [optional] 
**Replacesource** | Pointer to **bool** |  | [optional] 
**Parts** | Pointer to **string** |  | [optional] 
**Datastore** | Pointer to **string** |  | [optional] 
**Recover** | Pointer to **bool** |  | [optional] 
**Recoverytime** | Pointer to **int64** |  | [optional] 
**Poweronvm** | Pointer to **bool** |  | [optional] 
**Scripts** | Pointer to [**[]ScriptRest**](ScriptRest.md) |  | [optional] 
**Restoreobjectmappings** | Pointer to [**[]RestoreObjectMappingRest**](RestoreObjectMappingRest.md) |  | [optional] 
**Restoreoptions** | Pointer to [**[]RestoreOptionRest**](RestoreOptionRest.md) |  | [optional] 
**Provisioningoptions** | Pointer to [**[]ProvisioningOptionRest**](ProvisioningOptionRest.md) |  | [optional] 
**Selectedobjects** | Pointer to [**[]SelectedObjectRest**](SelectedObjectRest.md) |  | [optional] 
**Nfsoptions** | Pointer to [**NfsOptionsRest**](NfsOptionsRest.md) |  | [optional] 
**Instantmount** | Pointer to **bool** |  | [optional] 
**Notdisableschedule** | Pointer to **bool** |  | [optional] 
**Container** | Pointer to **bool** |  | [optional] 
**Allowedips** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewRestoreRest

`func NewRestoreRest() *RestoreRest`

NewRestoreRest instantiates a new RestoreRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreRestWithDefaults

`func NewRestoreRestWithDefaults() *RestoreRest`

NewRestoreRestWithDefaults instantiates a new RestoreRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *RestoreRest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RestoreRest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RestoreRest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RestoreRest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetHost

`func (o *RestoreRest) GetHost() HostRest`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RestoreRest) GetHostOk() (*HostRest, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RestoreRest) SetHost(v HostRest)`

SetHost sets Host field to given value.

### HasHost

`func (o *RestoreRest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetLabel

`func (o *RestoreRest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RestoreRest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RestoreRest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *RestoreRest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetUsername

`func (o *RestoreRest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RestoreRest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RestoreRest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RestoreRest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetCloudvmoptions

`func (o *RestoreRest) GetCloudvmoptions() CloudVmMountRest`

GetCloudvmoptions returns the Cloudvmoptions field if non-nil, zero value otherwise.

### GetCloudvmoptionsOk

`func (o *RestoreRest) GetCloudvmoptionsOk() (*CloudVmMountRest, bool)`

GetCloudvmoptionsOk returns a tuple with the Cloudvmoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudvmoptions

`func (o *RestoreRest) SetCloudvmoptions(v CloudVmMountRest)`

SetCloudvmoptions sets Cloudvmoptions field to given value.

### HasCloudvmoptions

`func (o *RestoreRest) HasCloudvmoptions() bool`

HasCloudvmoptions returns a boolean if a field has been set.

### GetRdmmode

`func (o *RestoreRest) GetRdmmode() string`

GetRdmmode returns the Rdmmode field if non-nil, zero value otherwise.

### GetRdmmodeOk

`func (o *RestoreRest) GetRdmmodeOk() (*string, bool)`

GetRdmmodeOk returns a tuple with the Rdmmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdmmode

`func (o *RestoreRest) SetRdmmode(v string)`

SetRdmmode sets Rdmmode field to given value.

### HasRdmmode

`func (o *RestoreRest) HasRdmmode() bool`

HasRdmmode returns a boolean if a field has been set.

### GetPhysicalrdm

`func (o *RestoreRest) GetPhysicalrdm() string`

GetPhysicalrdm returns the Physicalrdm field if non-nil, zero value otherwise.

### GetPhysicalrdmOk

`func (o *RestoreRest) GetPhysicalrdmOk() (*string, bool)`

GetPhysicalrdmOk returns a tuple with the Physicalrdm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalrdm

`func (o *RestoreRest) SetPhysicalrdm(v string)`

SetPhysicalrdm sets Physicalrdm field to given value.

### HasPhysicalrdm

`func (o *RestoreRest) HasPhysicalrdm() bool`

HasPhysicalrdm returns a boolean if a field has been set.

### GetTargetnode

`func (o *RestoreRest) GetTargetnode() string`

GetTargetnode returns the Targetnode field if non-nil, zero value otherwise.

### GetTargetnodeOk

`func (o *RestoreRest) GetTargetnodeOk() (*string, bool)`

GetTargetnodeOk returns a tuple with the Targetnode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetnode

`func (o *RestoreRest) SetTargetnode(v string)`

SetTargetnode sets Targetnode field to given value.

### HasTargetnode

`func (o *RestoreRest) HasTargetnode() bool`

HasTargetnode returns a boolean if a field has been set.

### GetReplacesource

`func (o *RestoreRest) GetReplacesource() bool`

GetReplacesource returns the Replacesource field if non-nil, zero value otherwise.

### GetReplacesourceOk

`func (o *RestoreRest) GetReplacesourceOk() (*bool, bool)`

GetReplacesourceOk returns a tuple with the Replacesource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacesource

`func (o *RestoreRest) SetReplacesource(v bool)`

SetReplacesource sets Replacesource field to given value.

### HasReplacesource

`func (o *RestoreRest) HasReplacesource() bool`

HasReplacesource returns a boolean if a field has been set.

### GetParts

`func (o *RestoreRest) GetParts() string`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *RestoreRest) GetPartsOk() (*string, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *RestoreRest) SetParts(v string)`

SetParts sets Parts field to given value.

### HasParts

`func (o *RestoreRest) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetDatastore

`func (o *RestoreRest) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *RestoreRest) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *RestoreRest) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *RestoreRest) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetRecover

`func (o *RestoreRest) GetRecover() bool`

GetRecover returns the Recover field if non-nil, zero value otherwise.

### GetRecoverOk

`func (o *RestoreRest) GetRecoverOk() (*bool, bool)`

GetRecoverOk returns a tuple with the Recover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecover

`func (o *RestoreRest) SetRecover(v bool)`

SetRecover sets Recover field to given value.

### HasRecover

`func (o *RestoreRest) HasRecover() bool`

HasRecover returns a boolean if a field has been set.

### GetRecoverytime

`func (o *RestoreRest) GetRecoverytime() int64`

GetRecoverytime returns the Recoverytime field if non-nil, zero value otherwise.

### GetRecoverytimeOk

`func (o *RestoreRest) GetRecoverytimeOk() (*int64, bool)`

GetRecoverytimeOk returns a tuple with the Recoverytime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverytime

`func (o *RestoreRest) SetRecoverytime(v int64)`

SetRecoverytime sets Recoverytime field to given value.

### HasRecoverytime

`func (o *RestoreRest) HasRecoverytime() bool`

HasRecoverytime returns a boolean if a field has been set.

### GetPoweronvm

`func (o *RestoreRest) GetPoweronvm() bool`

GetPoweronvm returns the Poweronvm field if non-nil, zero value otherwise.

### GetPoweronvmOk

`func (o *RestoreRest) GetPoweronvmOk() (*bool, bool)`

GetPoweronvmOk returns a tuple with the Poweronvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoweronvm

`func (o *RestoreRest) SetPoweronvm(v bool)`

SetPoweronvm sets Poweronvm field to given value.

### HasPoweronvm

`func (o *RestoreRest) HasPoweronvm() bool`

HasPoweronvm returns a boolean if a field has been set.

### GetScripts

`func (o *RestoreRest) GetScripts() []ScriptRest`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *RestoreRest) GetScriptsOk() (*[]ScriptRest, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *RestoreRest) SetScripts(v []ScriptRest)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *RestoreRest) HasScripts() bool`

HasScripts returns a boolean if a field has been set.

### GetRestoreobjectmappings

`func (o *RestoreRest) GetRestoreobjectmappings() []RestoreObjectMappingRest`

GetRestoreobjectmappings returns the Restoreobjectmappings field if non-nil, zero value otherwise.

### GetRestoreobjectmappingsOk

`func (o *RestoreRest) GetRestoreobjectmappingsOk() (*[]RestoreObjectMappingRest, bool)`

GetRestoreobjectmappingsOk returns a tuple with the Restoreobjectmappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreobjectmappings

`func (o *RestoreRest) SetRestoreobjectmappings(v []RestoreObjectMappingRest)`

SetRestoreobjectmappings sets Restoreobjectmappings field to given value.

### HasRestoreobjectmappings

`func (o *RestoreRest) HasRestoreobjectmappings() bool`

HasRestoreobjectmappings returns a boolean if a field has been set.

### GetRestoreoptions

`func (o *RestoreRest) GetRestoreoptions() []RestoreOptionRest`

GetRestoreoptions returns the Restoreoptions field if non-nil, zero value otherwise.

### GetRestoreoptionsOk

`func (o *RestoreRest) GetRestoreoptionsOk() (*[]RestoreOptionRest, bool)`

GetRestoreoptionsOk returns a tuple with the Restoreoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreoptions

`func (o *RestoreRest) SetRestoreoptions(v []RestoreOptionRest)`

SetRestoreoptions sets Restoreoptions field to given value.

### HasRestoreoptions

`func (o *RestoreRest) HasRestoreoptions() bool`

HasRestoreoptions returns a boolean if a field has been set.

### GetProvisioningoptions

`func (o *RestoreRest) GetProvisioningoptions() []ProvisioningOptionRest`

GetProvisioningoptions returns the Provisioningoptions field if non-nil, zero value otherwise.

### GetProvisioningoptionsOk

`func (o *RestoreRest) GetProvisioningoptionsOk() (*[]ProvisioningOptionRest, bool)`

GetProvisioningoptionsOk returns a tuple with the Provisioningoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningoptions

`func (o *RestoreRest) SetProvisioningoptions(v []ProvisioningOptionRest)`

SetProvisioningoptions sets Provisioningoptions field to given value.

### HasProvisioningoptions

`func (o *RestoreRest) HasProvisioningoptions() bool`

HasProvisioningoptions returns a boolean if a field has been set.

### GetSelectedobjects

`func (o *RestoreRest) GetSelectedobjects() []SelectedObjectRest`

GetSelectedobjects returns the Selectedobjects field if non-nil, zero value otherwise.

### GetSelectedobjectsOk

`func (o *RestoreRest) GetSelectedobjectsOk() (*[]SelectedObjectRest, bool)`

GetSelectedobjectsOk returns a tuple with the Selectedobjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedobjects

`func (o *RestoreRest) SetSelectedobjects(v []SelectedObjectRest)`

SetSelectedobjects sets Selectedobjects field to given value.

### HasSelectedobjects

`func (o *RestoreRest) HasSelectedobjects() bool`

HasSelectedobjects returns a boolean if a field has been set.

### GetNfsoptions

`func (o *RestoreRest) GetNfsoptions() NfsOptionsRest`

GetNfsoptions returns the Nfsoptions field if non-nil, zero value otherwise.

### GetNfsoptionsOk

`func (o *RestoreRest) GetNfsoptionsOk() (*NfsOptionsRest, bool)`

GetNfsoptionsOk returns a tuple with the Nfsoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsoptions

`func (o *RestoreRest) SetNfsoptions(v NfsOptionsRest)`

SetNfsoptions sets Nfsoptions field to given value.

### HasNfsoptions

`func (o *RestoreRest) HasNfsoptions() bool`

HasNfsoptions returns a boolean if a field has been set.

### GetInstantmount

`func (o *RestoreRest) GetInstantmount() bool`

GetInstantmount returns the Instantmount field if non-nil, zero value otherwise.

### GetInstantmountOk

`func (o *RestoreRest) GetInstantmountOk() (*bool, bool)`

GetInstantmountOk returns a tuple with the Instantmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantmount

`func (o *RestoreRest) SetInstantmount(v bool)`

SetInstantmount sets Instantmount field to given value.

### HasInstantmount

`func (o *RestoreRest) HasInstantmount() bool`

HasInstantmount returns a boolean if a field has been set.

### GetNotdisableschedule

`func (o *RestoreRest) GetNotdisableschedule() bool`

GetNotdisableschedule returns the Notdisableschedule field if non-nil, zero value otherwise.

### GetNotdisablescheduleOk

`func (o *RestoreRest) GetNotdisablescheduleOk() (*bool, bool)`

GetNotdisablescheduleOk returns a tuple with the Notdisableschedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotdisableschedule

`func (o *RestoreRest) SetNotdisableschedule(v bool)`

SetNotdisableschedule sets Notdisableschedule field to given value.

### HasNotdisableschedule

`func (o *RestoreRest) HasNotdisableschedule() bool`

HasNotdisableschedule returns a boolean if a field has been set.

### GetContainer

`func (o *RestoreRest) GetContainer() bool`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *RestoreRest) GetContainerOk() (*bool, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *RestoreRest) SetContainer(v bool)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *RestoreRest) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetAllowedips

`func (o *RestoreRest) GetAllowedips() []string`

GetAllowedips returns the Allowedips field if non-nil, zero value otherwise.

### GetAllowedipsOk

`func (o *RestoreRest) GetAllowedipsOk() (*[]string, bool)`

GetAllowedipsOk returns a tuple with the Allowedips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedips

`func (o *RestoreRest) SetAllowedips(v []string)`

SetAllowedips sets Allowedips field to given value.

### HasAllowedips

`func (o *RestoreRest) HasAllowedips() bool`

HasAllowedips returns a boolean if a field has been set.

### GetId

`func (o *RestoreRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestoreRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestoreRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RestoreRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *RestoreRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RestoreRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RestoreRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RestoreRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *RestoreRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *RestoreRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *RestoreRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *RestoreRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *RestoreRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *RestoreRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *RestoreRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *RestoreRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


